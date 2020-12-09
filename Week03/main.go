package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	port1 := 8081
	port2 := 8085

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	dummyError := make(chan error)

	g.Go(func() error {
		select {
		case <-sigs:
			dummyError <- fmt.Errorf("exit signal")
		}
		return nil
	})

	for i := port1; i <= port2; i++ {
		i := i
		g.Go(func() error {
			err := serveAt(ctx, i, dummyError)
			return err
		})
	}

	// g.Go(func() error {
	// 	err := serveAt(ctx, port2)
	// 	return err
	// })

	if err := g.Wait(); err != nil {
		fmt.Println("Error occured:", err)
	}
	fmt.Println("end main")
}

func serve(ctx context.Context, addr string, handler http.Handler, errOccured chan error) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("shutting down " + addr)
			s.Shutdown(ctx)
		case err := <-errOccured:
			fmt.Println("error occured at " + addr + ": " + err.Error())
			s.Shutdown(ctx)
		}
	}()
	fmt.Println("Start listening on " + addr)
	err := s.ListenAndServe()
	return err
}
func serveAt(ctx context.Context, port int, errOccured chan error) error {
	helloMsg := fmt.Sprintf("Hello World @%d!", port)
	serveAddr := fmt.Sprintf("127.0.0.1:%d", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(helloMsg))
	})
	return serve(ctx, serveAddr, mux, errOccured)
}
