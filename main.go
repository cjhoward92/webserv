package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func printHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Path)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

type myHandler struct{}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler(w, r)
}

func main() {
	//router := NewRouter()
	//router.AddRoute(Route{Name: "Home", Path: "/", Handler: &privateHandler{}})
	//Bind(router)

	h := myHandler{}

	server := new(http.Server)
	server.Addr = ":8080"
	http.Handle("/", printHandler(h))

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		fmt.Println(err)
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 10)
	}()

	go func() {
		wg.Wait()
		fmt.Println("Done waiting")
		ln.Close()
	}()

	sErr := server.Serve(ln)
	if sErr != nil {
		fmt.Println(sErr)
	}
}
