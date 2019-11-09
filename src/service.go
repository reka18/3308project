package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

func Start(config Config) *HTMLServer {
	context.WithCancel(context.Background())

	router := mux.NewRouter()
	router.PathPrefix("/css/").Handler(
		http.StripPrefix("/css",
			http.FileServer(
				http.Dir("web/css/"))))
	router.HandleFunc("/login", UserLoginHandler)
	router.HandleFunc("/create", CreateAccountHandler)
	router.HandleFunc("/", UserLandingHandler)

	htmlServer := HTMLServer{
		server: &http.Server{
			Addr:           config.Host,
			Handler:        router,
			ReadTimeout:    config.ReadTimeout,
			WriteTimeout:   config.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}

	htmlServer.wg.Add(1)

	go func() {
		fmt.Printf("\nHTMLServer : Service started : Host=%v\n", config.Host)
		_ = htmlServer.server.ListenAndServe()
		htmlServer.wg.Done()
	}()

	return &htmlServer
}

// Stop turns off the HTML Server
func (htmlServer *HTMLServer) Stop() error {

	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fmt.Printf("\nHTMLServer : Service stopping\n")

	if err := htmlServer.server.Shutdown(ctx); err != nil {

		if err := htmlServer.server.Close(); err != nil {
			fmt.Printf("\nHTMLServer : Service stopping : Error=%v\n", err)
			return err
		}
	}

	htmlServer.wg.Wait()
	fmt.Printf("\nHTMLServer : Stopped\n")
	return nil
}

// Render a template, or server error.
func render(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		fmt.Printf("\nRender Error: %v\n", err)
		return
	}
	w.Write(buf.Bytes())
}

// Push the given resource to the client.
func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}