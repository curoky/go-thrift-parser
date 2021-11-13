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

// ref: https://github.com/meixiu/pongo2-packr/blob/master/loader.go
package template

import (
	"bytes"
	"io"
	"path/filepath"

	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
)

// Fileb0xLoader provides a Pongo2 loader for fileb0x template files.
type Pongo2PackrLoader struct {
	Box *packr.Box
}

func NewPongo2PackrLoader(box *packr.Box) *Pongo2PackrLoader {
	return &Pongo2PackrLoader{box}
}

// Abs returns the absolute path to a template file.
func (loader *Pongo2PackrLoader) Abs(base, name string) string {
	if filepath.IsAbs(name) {
		return name
	}
	p := filepath.Join(filepath.Dir(base), name)
	return p
}

// Get retrieves a reader for the specified path.
func (loader *Pongo2PackrLoader) Get(path string) (io.Reader, error) {
	log.Debugf("Pongo2PackrLoader:Get %s", path)

	b, err := loader.Box.Find(path)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
