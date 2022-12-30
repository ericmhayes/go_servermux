package main

import (
	"log"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (th timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func timeHandlerRefactored(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func timeHandlerPassVariables(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func main() {
	// Use the http.NewServeMux() function to create an empty servemux.
	mux := http.NewServeMux()

	// Initialise the timeHandler in exactly the same way we would any normal
	// struct.
	th := timeHandler{format: time.RFC1123}

	//Refactoring from function above (eg. timeHandler -> timeHandlerRefactored)
	th_refactored := http.HandlerFunc(timeHandlerRefactored)

	//Refactored from function above to utlize max.HandleFunc() method

	th_pass_variable := timeHandlerPassVariables(time.RFC1123)

	mux.HandleFunc("/anothertimerefactor", timeHandlerRefactored)

	//Refactored from function above to prevent hardcoding time format

	// Use the http.RedirectHandler() function to create a handler which 307
	// redirects all requests it receives to http://example.org.
	rh := http.RedirectHandler("http://example.org", 307)

	// Next we use the mux.Handle() function to register this with our new
	// servemux, so it acts as the handler for all incoming requests with the URL
	// path /foo.
	mux.Handle("/foo", rh)

	mux.Handle("/time", th)

	mux.Handle("/timerefactored", th_refactored)

	mux.Handle("/passvariable", th_pass_variable)

	log.Print("Listening...")

	// Then we create a new server and start listening for incoming requests
	// with the http.ListenAndServe() function, passing in our servemux for it to
	// match requests against as the second parameter.
	http.ListenAndServe(":3000", mux)
}
