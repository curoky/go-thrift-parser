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
#pragma once
#include <folly/GLog.h>
#include <folly/SocketAddress.h>
#include <folly/io/async/AsyncServerSocket.h>
#include <folly/io/async/EventBase.h>
#include <thrift/lib/cpp/TProcessor.h>
#include <wangle/acceptor/ConnectionManager.h>
#include <wangle/acceptor/ServerSocketConfig.h>

#ifndef NEW_FACEBOOK
namespace folly {
using NetworkSocket = int;
}
#endif

namespace wolic {

class Acceptor : public folly::AsyncServerSocket::AcceptCallback,
                 public wangle::ConnectionManager::Callback {
 public:
  explicit Acceptor(apache::thrift::TProcessor* processor,
                    const wangle::ServerSocketConfig& accConfig, folly::EventBase* evb)
      : processor_(processor), accConfig_(accConfig), base_(evb) {
    downstreamConnectionManager_ =
        wangle::ConnectionManager::makeUnique(evb, accConfig_.connectionIdleTimeout, this);
  }

  ~Acceptor() {}

  // Inherited from folly::AsyncServerSocket::AcceptCallback
  void connectionAccepted(folly::NetworkSocket fd,
                          const folly::SocketAddress& clientAddr) noexcept override;

  void acceptError(const std::exception& ex) noexcept override {
    FB_LOG_EVERY_MS(ERROR, 1000) << "error accepting on acceptor socket: " << ex.what();
  }

  // void acceptStarted() noexcept override {}

  void acceptStopped() noexcept override { VLOG(3) << "Acceptor " << this << " acceptStopped()"; }

  // Inherit from wangle::ConnectionManager::Callback
  void onEmpty(const wangle::ConnectionManager& cm) override {
    VLOG(3) << "Acceptor=" << this << " onEmpty()";
  }
  void onConnectionAdded(const wangle::ManagedConnection* conn) override {}
  void onConnectionRemoved(const wangle::ManagedConnection* conn) override {}

  void addConnection(wangle::ManagedConnection* conn) {
    downstreamConnectionManager_->addConnection(conn, true);
  }

 private:
  apache::thrift::TProcessor* processor_;
  wangle::ServerSocketConfig accConfig_;
  folly::EventBase* base_{nullptr};
  // folly::SocketOptionMap socketOptions_;

  wangle::ConnectionManager::UniquePtr downstreamConnectionManager_;
};

}  // namespace blink
