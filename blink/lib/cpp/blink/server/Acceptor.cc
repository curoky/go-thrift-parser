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

#include "blink/server/Acceptor.h"

#include <folly/io/async/EventBaseManager.h>
#include <wangle/acceptor/TransportInfo.h>

#include <memory>
#include <utility>

#include "Connection.h"

namespace blink {

#ifndef NEW_FACEBOOK
void Acceptor::connectionAccepted(int fd, const folly::SocketAddress& clientAddr) noexcept {
#else
void Acceptor::connectionAccepted(
    folly::NetworkSocket fd, const folly::SocketAddress& clientAddr,
    folly::AsyncServerSocket::AcceptCallback::AcceptInfo info) noexcept {
#endif
  VLOG(3) << "connectionAccepted " << fd << ", " << clientAddr.describe();

  // QM: 这个已经是在IO线程里面了
  CHECK(base_ == folly::EventBaseManager::get()->getEventBase());

  // TODO(curoky): use socketOptions_
  // for (const auto& opt : socketOptions_){
  //   opt.first.apply()
  // }
#ifdef NEW_FACEBOOK
  auto socket = folly::AsyncSocket::newSocket(base_, fd);
#else
  auto socket = folly::AsyncSocket::UniquePtr(new folly::AsyncSocket(base_, fd),
                                              folly::AsyncSocket::Destructor());
#endif
  wangle::TransportInfo tinfo;
  tinfo.secure = false;
  tinfo.acceptTime = std::chrono::steady_clock::now();
  tinfo.tfoSucceded = socket->getTFOSucceded();
  tinfo.localAddr = std::make_shared<folly::SocketAddress>(accConfig_.bindAddress);
  socket->getLocalAddress(tinfo.localAddr.get());
  tinfo.remoteAddr = std::make_shared<folly::SocketAddress>(clientAddr);
  tinfo.initWithSocket(socket.get());

  auto conn = std::make_shared<Connection>(processor_, std::move(socket));
  conn->addConnection(conn);
  addConnection(conn.get());
  conn->start();
}

}  // namespace blink
