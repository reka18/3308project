package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start(config Config) *HTMLServer {

	context.WithCancel(context.Background())

	router := mux.NewRouter()

	router.PathPrefix("/css/").Handler(
		http.StripPrefix("/css",
			http.FileServer(
				http.Dir("web/css/"))))
	router.PathPrefix("/js/").Handler(
		http.StripPrefix("/js",
			http.FileServer(
				http.Dir("web/js/"))))
	router.PathPrefix("/vendor/").Handler(
		http.StripPrefix("/vendor",
			http.FileServer(
				http.Dir("web/vendor/"))))

	router.PathPrefix("/images/").Handler(
		http.StripPrefix("/images",
			http.FileServer(
				http.Dir("web/images/"))))

	router.PathPrefix("/fonts/").Handler(
		http.StripPrefix("/fonts",
			http.FileServer(
				http.Dir("web/fonts/"))))

	router.HandleFunc("/", BaseUrlHandler)
	router.HandleFunc("/login", UserLoginHandler)
	router.HandleFunc("/create", CreateAccountHandler)
	router.HandleFunc("/logout", UserLogoutHandler)
	router.HandleFunc("/avatar", AvatarHandler)
	router.HandleFunc("/search", SearchHandler)
	router.HandleFunc("/user", UserHandler)
	router.HandleFunc("/{user}", UserLandingHandler)
	router.HandleFunc("/{user}/post", UserPostHandler)
	router.HandleFunc("/{user}/follow", FollowHandler)
	router.HandleFunc("/{user}/vote", VoteHandler)
	router.HandleFunc("/{user}/message", MessageHandler)

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
		log.Printf(Detail("SERVER : Service connection started : Host=%v"), config.Host)
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

	log.Println(Detail("SERVER : Service stopping."))

	if e := htmlServer.server.Shutdown(ctx); e != nil {

		if e := htmlServer.server.Close(); e != nil {
			log.Printf(Warn("SERVER : Service stopping : Error=%s"), e)
			return e
		}
	}

	htmlServer.wg.Wait()
	log.Println(Detail("SERVER : Stopped"))
	return nil
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

func pushAllResources(w http.ResponseWriter) {
	/*
		THIS FUNCTION APPLIES THE "PUSH" FUNC TO ALL NEEDED RESOURCE
	*/
	push(w, "css/main.css")
	push(w, "css/util.css")
	push(w, "js/main.js")
	push(w, "vendor/animate/animate.css")
	push(w, "vendor/animsition/css/animsition.css")
	push(w, "vendor/animsition/js/animsition.js")
	push(w, "vendor/bootstrap/css/bootstrap.css")
	push(w, "vendor/bootstrap/css/bootstrap-grid.css")
	push(w, "vendor/bootstrap/css/bootstrap-reboot.css")
	push(w, "vendor/bootstrap/js/bootstrap.js")
	push(w, "vendor/bootstrap/js/popper.js")
	push(w, "vendor/bootstrap/js/tooltip.js")
	push(w, "vendor/countdowntime/countdowntime.js")
	push(w, "vendor/css-hamburgers/hamburgers.css")
	push(w, "vendor/daterangepicker/daterangepicker.css")
	push(w, "vendor/daterangepicker/daterangepicker.js")
	push(w, "vendor/daterangepicker/moment.js")
	push(w, "vendor/jquery/jquery-3.2.1.min.js")
	push(w, "vendor/perfect-scrollbar/perfect-scrollbar.css")
	push(w, "vendor/perfect-scrollbar/perfect-scrollbar.min.js")
	push(w, "vendor/select2/select2.css")
	push(w, "vendor/select2/select2.js")
	push(w, "vendor/SlickGrid-2.4.14/lib/jquery-3.1.0.js")

	push(w, "vendor/SlickGrid-2.4.15/lib/jquery-3.1.0.js")
	push(w, "vendor/SlickGrid-2.4.15/controls/slick.pager.js")
	push(w, "vendor/SlickGrid-2.4.15/lib/jquery.event.drag-2.3.0.js")
	push(w, "vendor/SlickGrid-2.4.15/slick.core.js")
	push(w, "vendor/SlickGrid-2.4.15/slick.grid.js")
	push(w, "vendor/SlickGrid-2.4.15/lib/firebugx.js")

	push(w, "vendor/SlickGrid-2.4.15/lib/jquery.event.drop-2.3.0.js")
	push(w, "vendor/SlickGrid-2.4.15/plugins/slick.cellrangeselector.js")
	push(w, "vendor/SlickGrid-2.4.15/plugins/slick.cellselectionmodel.js")
	push(w, "vendor/SlickGrid-2.4.15/plugins/slick.rowselectionmodel.js")

	push(w, "vendor/SlickGrid-2.4.15/plugins/slick.rowmovemanager.js")
	push(w, "vendor/SlickGrid-2.4.15/slick.formatters.js")
	push(w, "vendor/SlickGrid-2.4.15/slick.editors.js")
	push(w, "vendor/SlickGrid-2.4.15/slick.dataview.js")

}
