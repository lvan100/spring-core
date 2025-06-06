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

package main

import (
	"context"
	"log"
	"net"

	"github.com/go-spring/spring-core/gs"
	"google.golang.org/grpc"
)

func init() {
	gs.Object(&SimpleGrpcServer{}).AsServer().Condition(
		gs.OnBean[GrpcServerConfiger](),
	)
}

type GrpcServerConfiger func(svr *grpc.Server)

type SimpleGrpcServer struct {
	Addr string               `value:"${grpc.server.addr:=0.0.0.0:9494}"`
	Cfgs []GrpcServerConfiger `autowire:""`
	svr  *grpc.Server
}

func (s *SimpleGrpcServer) ListenAndServe(sig gs.ReadySignal) error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s.svr = grpc.NewServer()
	for _, cfg := range s.Cfgs {
		cfg(s.svr)
	}
	<-sig.TriggerAndWait()
	return s.svr.Serve(listener)
}

func (s *SimpleGrpcServer) Shutdown(ctx context.Context) error {
	s.svr.GracefulStop()
	return nil
}
