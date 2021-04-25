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
#include <wangle/acceptor/ManagedConnection.h>

#include <memory>
#include <string>
#include <utility>

#include "Channel.h"

namespace blink {

class Connection : public wangle::ManagedConnection {
 public:
  Connection(apache::thrift::TProcessor* processor, folly::AsyncSocket::UniquePtr socket) {
    channel_ = std::make_shared<Channel>(processor, std::move(socket), this);
  }

  ~Connection() { VLOG(3) << "~Connection"; }

  void timeoutExpired() noexcept override { stop(); }

  void start() { channel_->start(); }

  void stop();

  // wangle::ManagedConnection
  void describe(std::ostream&) const override {}
  bool isBusy() const override { return true; }  // TODO(curoky):
  void notifyPendingShutdown() override {}
  void closeWhenIdle() override { stop(); }
#ifdef NEW_FACEBOOK
  void dropConnection(const std::string& errorMsg = "") override { stop(); }
#else
  void dropConnection() override { stop(); }
#endif
  void dumpConnectionState(uint8_t /* loglevel */) override {}
  void addConnection(std::shared_ptr<Connection> conn) { this_ = conn; }

 private:
  std::shared_ptr<Connection> this_;
  std::shared_ptr<Channel> channel_;
};

}  // namespace blink
