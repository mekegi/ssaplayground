// Copyright 2020 The golang.design Initiative authors.
// All rights reserved. Use of this source code is governed
// by a GPLv3 license that can be found in the LICENSE file.

package boot

import (
	"log"
	"net"
	"net/http"

	"golang.design/x/ssaplayground/src/config"
	"golang.design/x/ssaplayground/src/route"
)

func init() {
	log.SetPrefix("redir: ")
	log.SetFlags(log.Lmsgprefix | log.LstdFlags | log.Lshortfile)
	config.Init()
}

func Run() {

	// terminated := make(chan bool, 1)

	// go func() {
	// 	quit := make(chan os.Signal, 1)
	// 	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	// 	sig := <-quit

	// 	log.Printf("service is stopped with signal: %v", sig)

	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	if err := server.Shutdown(ctx); err != nil {
	// 		log.Printf("close ssaplayground with error: %v", err)
	// 	}

	// 	cancel()
	// 	terminated <- true
	// }()

	listener, err := net.Listen("tcp", config.Get().Addr)
	if err != nil {
		panic(err)
	}
	hnd := route.Register()
	addr := listener.Addr().(*net.TCPAddr)
	log.Printf("welcome to ssaplayground service... http://%s:%d/gossa",
		addr.IP, addr.Port)

	panic(http.Serve(listener, hnd))
}
