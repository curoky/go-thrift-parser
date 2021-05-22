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
#include "Channel.h"

#include <thrift/lib/cpp/protocol/TBinaryProtocol.h>
#include <thrift/lib/cpp/protocol/TProtocol.h>
#include <thrift/lib/cpp/transport/TBufferTransports.h>
#include <thrift/lib/cpp/transport/TTransport.h>

#include <memory>
#include <utility>

#include "Connection.h"

namespace blink {

void Channel::process(std::unique_ptr<folly::IOBuf> unframed, std::unique_ptr<THeader> header) {
  CHECK(unframed->isChained() == false);

  VLOG(3) << "process len: " << unframed->computeChainDataLength();
  static auto inProtoFac = std::make_shared<apache::thrift::protocol::TBinaryProtocolFactory>();
  static auto outProtoFac = std::make_shared<apache::thrift::protocol::TBinaryProtocolFactory>();

  auto in = std::make_shared<apache::thrift::transport::TMemoryBuffer>(unframed->writableData(),
                                                                       unframed->length());
  auto out =
      std::make_shared<apache::thrift::transport::TMemoryBuffer>(static_cast<uint32_t>(1024));
  auto inProto = inProtoFac->getProtocol(in);
  auto outProto = outProtoFac->getProtocol(out);
  processor_->process(inProto, outProto, nullptr);

  VLOG(3) << "out size: " << out->getBufferSize();

  auto buf = out->wrapBufferAsIOBuf();
  auto f = write(std::move(buf), header.get());
  std::move(f).thenTry([this](folly::Try<folly::Unit>&& t) {
    if (t.withException<apache::thrift::transport::TTransportException>(
            [&](const apache::thrift::transport::TTransportException& ex) {
              LOG(ERROR) << "write with: " << ex.what();
              this->conn_->stop();
            }) ||
        t.withException<std::exception>([&](const std::exception& ex) {
          LOG(ERROR) << "write with: " << ex.what();
          this->conn_->stop();
        })) {
      return;
    } else {
      VLOG(3) << "write success";
    }
  });
}
void Channel::read(folly::IOBufQueue& q) {
  size_t remaining = 0;
  while (!closing_) {
    std::unique_ptr<folly::IOBuf> unframed;
    std::unique_ptr<THeader> header;
    auto ex = folly::try_and_catch([&]() {
      // got a decrypted message
      std::tie(unframed, remaining, header) = FrameUtils::removeFrame(&q);
    });

    if (ex) {
      VLOG(5) << "Failed to read a message header";
      close(true);
      return;
    }

    if (header) {
      VLOG(3) << "getClientType: " << static_cast<int>(header->getClientType());
    }

    if (!unframed) {
      const auto& s = asyncSocketHandler_->getReadBufferSettings();
      asyncSocketHandler_->setReadBufferSettings(s.first, remaining ? remaining : s.second);
      return;
    } else {
      process(std::move(unframed), std::move(header));
    }
  }
}

folly::Future<folly::Unit> Channel::write(std::unique_ptr<folly::IOBuf> buf, THeader* header) {
  auto bufWithHeader = FrameUtils::addFrame(std::move(buf), header);
  return asyncSocketHandler_->write(std::move(bufWithHeader));
}
}  // namespace blink
