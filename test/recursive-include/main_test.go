/*
 * Copyright (c) 2020-2023 curoky(cccuroky@gmail.com).
 *
 * This file is part of go-thrift-parser.
 * See https://github.com/curoky/blink for further info.
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

package main

import (
	"testing"

	"github.com/curoky/go-thrift-parser/parser"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	_, err := parser.ParseThriftFile("b/b.thrift", []string{}, true, false)
	require.NoError(t, err)
	_, err = parser.ParseThriftFile("a/a.thrift", []string{}, true, false)
	require.NoError(t, err)
}
