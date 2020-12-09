package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	portBegin, portEnd := 8081, 8083
	stop := make(chan string, portEnd-portBegin)

	x := func(port int) {
		g.Go(func() error {
			return serveAt(ctx, port, stop)
		})
	}
	for i := portBegin; i <= portEnd; i++ {
		x(i)
	}

	if err := g.Wait(); err != nil {
		stop <- err.Error()
		fmt.Println("Error occured:", err)
	}
}

func serve(ctx context.Context, addr string, handler http.Handler, stop <-chan string) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		msg := <-stop
		fmt.Println("Shutting down due to:", msg)
		s.Shutdown(ctx)
	}()
	return s.ListenAndServe()
}

func serveAt(ctx context.Context, port int, stop <-chan string) error {
	helloMsg := fmt.Sprintf("Hello World @%d!", port)
	serveAddr := fmt.Sprintf("127.0.0.1:%d", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(helloMsg))
	})
	return serve(ctx, serveAddr, mux, stop)
}
