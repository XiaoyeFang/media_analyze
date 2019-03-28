// Copyright © 2018 joy  <lzy@spf13.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"flag"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"pure-media/cmd"
	"pure-media/config"
	"pure-media/gogrpc"
	"pure-media/protos"
)

func main() {
	flag.Parse()
	cmd.Execute()
	grpcStart()
	glog.Flush()
}
func grpcStart() {
	lis, err := net.Listen("tcp", config.MediaConfig.GrpcListen)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	srv := gogrpc.PureMedia{}
	protos.RegisterPureMediaServiceServer(server, &srv)
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
