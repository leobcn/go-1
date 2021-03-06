/*
Copyright 2014 Tamás Gulácsi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package log15hlp

import (
	"io"

	log15 "gopkg.in/inconshreveable/log15.v2"

	"github.com/tgulacsi/go/term"
)

// UseWriter will use the given writer for log15.StderrHandler.
func UseWriter(w io.Writer) {
	if w == nil {
		return
	}
	logfmt := log15.LogfmtFormat()
	if term.IsTTY {
		logfmt = log15.TerminalFormat()
	}
	log15.StderrHandler = log15.StreamHandler(w, logfmt)
}
