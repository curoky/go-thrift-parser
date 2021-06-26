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
#include <thrift/lib/cpp/transport/THeader.h>

#include <memory>
#include <tuple>
#include <utility>

namespace blink {

class FrameUtils {
 public:
  using THeader = apache::thrift::transport::THeader;

  static std::tuple<std::unique_ptr<folly::IOBuf>, size_t, std::unique_ptr<THeader>> removeFrame(
      folly::IOBufQueue* q) {
    if (!q || !q->front() || q->front()->empty()) {
      return make_tuple(std::unique_ptr<folly::IOBuf>(), 0, nullptr);
    }

    auto header = std::make_unique<THeader>(THeader::ALLOW_BIG_FRAMES);
    std::unique_ptr<folly::IOBuf> buf;
    size_t remaining = 0;
    try {
      THeader::StringToStringMap persistentReadHeaders;
      buf = header->removeHeader(q, remaining, persistentReadHeaders);
    } catch (const std::exception& e) {
      LOG(ERROR) << "Received invalid request from client: " << folly::exceptionStr(e);
      throw;
    }
    if (!buf) {
      return make_tuple(std::unique_ptr<folly::IOBuf>(), remaining, nullptr);
    }
    return make_tuple(std::move(buf), 0, std::move(header));
  }

  static std::unique_ptr<folly::IOBuf> addFrame(std::unique_ptr<folly::IOBuf> buf,
                                                THeader* header) {
    THeader::StringToStringMap persistentWriteHeaders;
    return header->addHeader(std::move(buf), persistentWriteHeaders, false);
  }
};

}  // namespace blink
