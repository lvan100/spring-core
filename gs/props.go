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

package gs

const (
	// AllowCircularReferencesProp controls whether the container
	// allows circular dependencies between beans.
	AllowCircularReferencesProp = "spring.allow-circular-references"

	// ForceAutowireIsNullableProp forces autowired dependencies
	// to be treated as nullable (i.e. allowed to be nil).
	ForceAutowireIsNullableProp = "spring.force-autowire-is-nullable"

	// ActiveProfilesProp defines the active application profiles
	// (e.g. "dev", "test", "prod").
	ActiveProfilesProp = "spring.profiles.active"

	// EnableJobsProp enables or disables scheduled job execution.
	EnableJobsProp = "spring.app.enable-jobs"

	// EnableServersProp enables or disables all server components.
	EnableServersProp = "spring.app.enable-servers"

	// EnableSimpleHttpServerProp enables or disables the built-in
	// lightweight HTTP server.
	EnableSimpleHttpServerProp = "spring.enable.simple-http-server"

	// EnableSimplePProfServerProp enables or disables the built-in
	// lightweight pprof server.
	EnableSimplePProfServerProp = "spring.enable.simple-pprof-server"
)
