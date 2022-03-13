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

load("@rules_cc//cc:defs.bzl", "cc_test")
load("//bazel:copts.bzl", "TEST_COPTS", "TEST_LINKOPTS")

cc_test(
    name = "type",
    testonly = False,
    srcs = glob(["*.cc"]),
    copts = TEST_COPTS + [
        "-DUBPA_USREFL_USE_NAMEOF",
    ],
    linkopts = TEST_LINKOPTS,
    deps = [
        "@com_github_aantron_better_enums//:better-enums",
        "@com_github_curoky_wolic//:catch2_main",
        "@com_github_neargye_nameof//:nameof",
        "@com_github_rttrorg_rttr//:rttr",
        "@com_github_ubpa_usrefl//:usrefl",
        "@com_github_veselink1_refl_cpp//:refl-cpp",
    ],
)
