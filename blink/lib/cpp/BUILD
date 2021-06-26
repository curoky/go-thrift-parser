# Copyright 2021 curoky(cccuroky@gmail.com).
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("@rules_cc//cc:defs.bzl", "cc_library")
load("//bazel:copts.bzl", "TEST_COPTS", "TEST_LINKOPTS")

cc_library(
    name = "blink",
    srcs = glob(["blink/**/*.cc"]),
    hdrs = glob(["blink/**/*.h"]),
    copts = TEST_COPTS,
    defines = ["NEW_FACEBOOK"],
    includes = ["."],
    linkopts = TEST_LINKOPTS,
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_facebook_fbthrift//:fbthrift",
        "@com_github_facebook_folly//:folly",
        "@com_github_gabime_spdlog//:spdlog",
        "@com_github_gflags_gflags//:gflags",
        "@com_github_google_glog//:glog",
        "@com_github_nlohmann_json//:json",
        "@com_google_protobuf//:protobuf",
    ],
)
