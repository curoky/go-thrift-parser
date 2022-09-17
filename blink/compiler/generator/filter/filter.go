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
package filter

import (
	"github.com/flosch/pongo2/v4"
	log "github.com/sirupsen/logrus"
)

func Init() {
	err := pongo2.RegisterFilter("base_name", BaseName)
	if err != nil {
		log.Error(err)
	}
	err = pongo2.RegisterFilter("cpp_type", CppType)
	if err != nil {
		log.Error(err)
	}
	err = pongo2.RegisterFilter("ann_cpp_type", AnnCppType)
	if err != nil {
		log.Error(err)
	}
	err = pongo2.RegisterFilter("cpp_value", CppValue)
	if err != nil {
		log.Error(err)
	}
	err = pongo2.RegisterFilter("expandCategory", expandCategory)
	if err != nil {
		log.Error(err)
	}
}
