package assert

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

func logCall(t testingT, funcName string) {
	var buffer bytes.Buffer
	insidePkg := true
	lastLine := 0
	lastFile := ""
	for i := 0; ; i += 1 {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if !insidePkg {
			buffer.WriteString(fmt.Sprintf("%s() call at \n%s:%d", funcName, lastFile, lastLine))
			break
		}
		if strings.Index(file, "github.com/emicklei/assert") == -1 {
			insidePkg = false
		}
		lastLine = line
		lastFile = file

	}
	Log(t, buffer.String())
}
