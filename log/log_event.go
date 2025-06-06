/*
 * Copyright 2025 The Go-Spring Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package log

import (
	"sync"
	"time"
)

var eventPool = sync.Pool{
	New: func() any {
		return &Event{}
	},
}

// Event provides contextual information about a log message.
type Event struct {
	Level     Level     // The severity level of the log (e.g., INFO, ERROR, DEBUG)
	Time      time.Time // The timestamp when the event occurred
	File      string    // The source file where the log was triggered
	Line      int       // The line number in the source file
	Tag       string    // A tag used to categorize the log (e.g., subsystem name)
	Fields    []Field   // Custom fields provided specifically for this log event
	CtxString string    // The string representation of the context
	CtxFields []Field   // Additional fields derived from the context (e.g., request ID, user ID)
}

func (e *Event) Reset() {
	e.Level = NoneLevel
	e.Time = time.Time{}
	e.File = ""
	e.Line = 0
	e.Tag = ""
	e.Fields = nil
	e.CtxString = ""
	e.CtxFields = nil
}

func GetEvent() *Event {
	return eventPool.Get().(*Event)
}

func PutEvent(e *Event) {
	e.Reset()
	eventPool.Put(e)
}
