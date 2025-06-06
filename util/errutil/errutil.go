/*
 * Copyright 2024 The Go-Spring Authors.
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

package errutil

import (
	"fmt"
)

// LineBreak defines the separator used between errors with hierarchical relationships.
var LineBreak = " << "

// WrapError wraps an existing error, creating a new error with hierarchical relationships.
func WrapError(err error, format string, args ...any) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s"+LineBreak+"%w", msg, err)
}
