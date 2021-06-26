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
#include <folly/io/async/AsyncSocket.h>
#include <thrift/lib/cpp/TProcessor.h>

#include <memory>
#include <utility>

#include "blink/server/AsyncSocketHandler.h"
#include "blink/server/FrameUtils.h"

namespace blink {

class Connection;
class Channel {
 public:
  using THeader = apache::thrift::transport::THeader;

  Channel(apache::thrift::TProcessor* processor, folly::AsyncSocket::UniquePtr socket,
          Connection* conn)
      : processor_(processor), conn_(conn) {
    asyncSocketHandler_ = std::make_shared<AsyncSocketHandler>(conn, this, std::move(socket));
  }

  void read(folly::IOBufQueue& q);

  void process(std::unique_ptr<folly::IOBuf> unframed, std::unique_ptr<THeader> header);

  folly::Future<folly::Unit> write(std::unique_ptr<folly::IOBuf> buf, THeader* header);

  void start() { asyncSocketHandler_->start(); }
  void stop() {
    closing_ = true;
    asyncSocketHandler_->stop();
  }

 private:
  apache::thrift::TProcessor* processor_;
  Connection* conn_;
  bool closing_{false};
  std::shared_ptr<AsyncSocketHandler> asyncSocketHandler_{nullptr};
};

}  // namespace blink
