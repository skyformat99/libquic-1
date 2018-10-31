# Copyright Go-IIoT (https://github.com/goiiot)
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

.PHONY: test ensure style_check fmt bench_utils profile

TEST_FLAGS = -v -count=1 -race -coverprofile=coverage.txt -covermode=atomic

PROFILE_OUT_DIR = profile
BENCH_FLAGS = -v -count=3 -benchmem -outputdir=${PROFILE_OUT_DIR} -cpuprofile=cpu.out -bench Benchmark*

test:
	go test ${TEST_FLAGS} ./...

ensure:
	go mod download

style_check:
	gofmt -s -e -d . 1>&2

fmt:
	gofmt -s -w .

bench_utils:
	go test ${BENCH_FLAGS} ./utils

profile:
	go tool pprof -http localhost:8080 libquic.test ${PROFILE_OUT_DIR}/cpu.out
