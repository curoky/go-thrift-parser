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

load("@com_grail_bazel_compdb//:defs.bzl", "compilation_database")

compilation_database(
    name = "compdb",
    output_base = "/shm/cache/bazel",
    # bazel query '//blink/...' 2>&1 | xargs -n 1 | awk '{print "\""$1"\","}' | grep "\"//" | sort -n
    targets = [
        "//blink/lib/cpp:blink",
    ],
)
