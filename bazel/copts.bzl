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

"""blink specific copts.
"""

load("@com_curoky_rules_pkg//bazel:copts.bzl", "BASE_CPP_COPTS", _DEFAULT_LINKOPTS = "DEFAULT_LINKOPTS", _TEST_CPP_COPTS = "TEST_CPP_COPTS", _TEST_LINKOPTS = "TEST_LINKOPTS")

DEFAULT_COPTS = BASE_CPP_COPTS
TEST_COPTS = _TEST_CPP_COPTS

DEFAULT_LINKOPTS = _DEFAULT_LINKOPTS
TEST_LINKOPTS = _TEST_LINKOPTS
