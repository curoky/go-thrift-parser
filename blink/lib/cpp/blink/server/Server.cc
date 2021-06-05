/*
 * Copyright 2021 curoky(cccuroky@gmail.com).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
#include "Server.h"

#include <folly/executors/CPUThreadPoolExecutor.h>
#include <folly/executors/IOThreadPoolExecutor.h>
#include <folly/io/async/AsyncServerSocket.h>
#include <wangle/bootstrap/ServerSocketFactory.h>

#include <utility>

#include "Acceptor.h"
#include "AsyncSocketHandler.h"
#include "Channel.h"
#include "Connection.h"
namespace blink {

Server::Server(std::shared_ptr<apache::thrift::TProcessor> processor)
    : processor_(std::move(processor)) {
  serveEventBase_ = folly::EventBaseManager::get()->getEventBase();
  socketFactory_ = std::make_shared<wangle::AsyncServerSocketFactory>();
}

void Server::serve() {
  serveEventBase_->loopForever();
  join();
}

void Server::stop() { serveEventBase_->terminateLoopSoon(); }

void Server::join() {
  acceptorGroup_->join();
  ioGroup_->join();
}

void Server::bind(folly::SocketAddress& address) {
  // 1. 创建 accept / io 线程池
  acceptorGroup_ = std::make_shared<folly::IOThreadPoolExecutor>(
      1, std::make_shared<folly::NamedThreadFactory>("Acceptor Thread"));
  ioGroup_ = std::make_shared<folly::IOThreadPoolExecutor>(
      std::thread::hardware_concurrency(),
      std::make_shared<folly::NamedThreadFactory>("IO Thread"));

  // 2. 在 accept 线程里面创建 AsyncServerSocket, 每个 socket
  // 会绑定到当前 accept 线程的 evb 上
  sockets_->resize(acceptorGroup_->numThreads());
  for (size_t i = 0; i < acceptorGroup_->numThreads(); i++) {
    auto barrier = std::make_shared<folly::Baton<>>();
    acceptorGroup_->add([this, barrier, i, &address] {
      bool reusePort = this->acceptorGroup_->numThreads() > 1;
      (*this->sockets_)[i] = std::dynamic_pointer_cast<folly::AsyncServerSocket>(
          this->socketFactory_->newSocket(address, 0, reusePort, this->accConfig_));
      barrier->post();
    });
    barrier->wait();
  }

  acceptors_.resize(ioGroup_->numThreads());
  // 3. 在 io 线程里设置 AsyncServerSocket 的 accept callback
  for (auto& socket : *sockets_) {
    for (size_t i = 0; i < ioGroup_->numThreads(); i++) {
      auto barrier = std::make_shared<folly::Baton<>>();
      ioGroup_->add([this, barrier, socket, i] {
        folly::EventBase* evb = folly::EventBaseManager::get()->getEventBase();
        if (this->acceptors_[i] == nullptr) {
          this->acceptors_[i] = std::make_shared<Acceptor>(processor_.get(), this->accConfig_, evb);
        }
        socket->getEventBase()->runInEventBaseThreadAndWait(
            [&]() { socket->addAcceptCallback(this->acceptors_[i].get(), evb); });
        barrier->post();
      });
      barrier->wait();
    }
  }
}

}  // namespace blink
