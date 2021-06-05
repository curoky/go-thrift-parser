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
#include <folly/futures/Future.h>
#include <folly/io/IOBufQueue.h>
#include <folly/io/async/AsyncSocket.h>
#include <thrift/lib/cpp/transport/TTransportException.h>

#include <memory>
#include <utility>
namespace blink {

class Connection;
class Channel;
class AsyncSocketHandler : public folly::AsyncReader::ReadCallback {
 public:
  AsyncSocketHandler(Connection* conn, Channel* channel, folly::AsyncSocket::UniquePtr socket)
      : conn_(conn), channel_(channel), socket_(std::move(socket)) {}

  ~AsyncSocketHandler() { close(false); }

  void getReadBuffer(void** bufReturn, size_t* lenReturn) override {
    const auto ret = bufQueue_.preallocate(readBufferSettings_.first, readBufferSettings_.second);
    *bufReturn = ret.first;
    *lenReturn = ret.second;
  }

  void readDataAvailable(size_t len) noexcept override;

  bool isBufferMovable() noexcept override { return true; }

  void readBufferAvailable(std::unique_ptr<folly::IOBuf> buf) noexcept override {
    VLOG(3) << buf->computeChainDataLength();
    bufQueue_.append(std::move(buf));
  }

  void readErr(const folly::AsyncSocketException& ex) noexcept override;

  void readEOF() noexcept override;

  folly::Future<folly::Unit> write(std::unique_ptr<folly::IOBuf> buf) {
    if (UNLIKELY(!buf)) {
      return folly::makeFuture();
    }

    if (!socket_->good()) {
      VLOG(5) << "transport is closed in write()";
      return folly::makeFuture<folly::Unit>(
          apache::thrift::transport::TTransportException("transport is closed in write()"));
    }

    auto cb = new WriteCallback();
    auto future = cb->promise_.getFuture();
    socket_->writeChain(cb, std::move(buf), writeFlags_);
    return future;
  }

  const std::pair<uint64_t, uint64_t>& getReadBufferSettings() { return readBufferSettings_; }

  void setReadBufferSettings(uint64_t x, uint64_t y) { readBufferSettings_ = {x, y}; }

  void start() { socket_->setReadCB(socket_->good() ? this : nullptr); }
  void stop() { close(false); }

 private:
  void close(bool closeWithReset) {
    if (socket_) {
      if (socket_->getReadCallback() == this) {
        socket_->setReadCB(nullptr);
      }
      if (closeWithReset && socket_->good()) {
        socket_->closeWithReset();
      } else {
        socket_->closeNow();
      }
    }
  }

  void refreshTimeout() {
    // TODO(curoky):
    // auto manager = getContext()->getPipeline()->getPipelineManager();
    // if (manager) {
    //     manager->refreshTimeout();
    // }
  }
  class WriteCallback : private folly::AsyncWriter::WriteCallback {
    void writeSuccess() noexcept override {
      promise_.setValue();
      delete this;
    }

    void writeErr(size_t /*bytesWritten*/,
                  const folly::AsyncSocketException& ex) noexcept override {
      apache::thrift::transport::TTransportException te(
          apache::thrift::transport::TTransportException::TTransportExceptionType(ex.getType()),
          ex.what(), ex.getErrno());
      promise_.setException(te);
      delete this;
    }

   private:
    friend class AsyncSocketHandler;
    folly::Promise<folly::Unit> promise_;
  };

  Connection* conn_;
  Channel* channel_;
  folly::WriteFlags writeFlags_{folly::WriteFlags::NONE};
  std::pair<uint64_t, uint64_t> readBufferSettings_{2048, 2048};
  folly::IOBufQueue bufQueue_{folly::IOBufQueue::cacheChainLength()};
  folly::AsyncSocket::UniquePtr socket_{nullptr};
};

}  // namespace blink
