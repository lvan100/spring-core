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

package log_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-spring/spring-core/log"
	"github.com/lvan100/go-assert"
)

const TraceID = "trace_id"
const SpanID = "span_id"

var TagDefault = log.GetTag("_def")
var TagRequestIn = log.GetTag("_com_request_in")
var TagRequestOut = log.GetTag("_com_request_out")

func TestLog(t *testing.T) {
	ctx := t.Context()

	log.TimeNow = func(ctx context.Context) time.Time {
		return time.Now()
	}

	log.StringFromContext = func(ctx context.Context) string {
		return ""
	}

	log.FieldsFromContext = func(ctx context.Context) []log.Field {
		traceID, _ := ctx.Value(TraceID).(string)
		spanID, _ := ctx.Value(SpanID).(string)
		return []log.Field{
			log.String(TraceID, traceID),
			log.String(SpanID, spanID),
		}
	}

	log.Debug(ctx, TagRequestOut, func() []log.Field {
		return []log.Field{
			log.Msgf("hello %s", "world"),
		}
	})

	log.Info(ctx, TagDefault, log.Msgf("hello %s", "world"))
	log.Info(ctx, TagRequestIn, log.Msgf("hello %s", "world"))

	err := log.RefreshFile("testdata/log.xml")
	assert.Nil(t, err)

	ctx = context.WithValue(ctx, TraceID, "0a882193682db71edd48044db54cae88")
	ctx = context.WithValue(ctx, SpanID, "50ef0724418c0a66")

	log.Trace(ctx, TagRequestOut, func() []log.Field {
		return []log.Field{
			log.Msgf("hello %s", "world"),
		}
	})

	log.Debug(ctx, TagRequestOut, func() []log.Field {
		return []log.Field{
			log.Msgf("hello %s", "world"),
		}
	})

	log.Info(ctx, TagRequestIn, log.Msgf("hello %s", "world"))
	log.Warn(ctx, TagRequestIn, log.Msgf("hello %s", "world"))
	log.Error(ctx, TagRequestIn, log.Msgf("hello %s", "world"))
	log.Panic(ctx, TagRequestIn, log.Msgf("hello %s", "world"))
	log.Fatal(ctx, TagRequestIn, log.Msgf("hello %s", "world"))

	log.Info(ctx, TagDefault, log.Msgf("hello %s", "world"))
	log.Warn(ctx, TagDefault, log.Msgf("hello %s", "world"))
	log.Error(ctx, TagDefault, log.Msgf("hello %s", "world"))
	log.Panic(ctx, TagDefault, log.Msgf("hello %s", "world"))
}
