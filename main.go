package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// generate a `Certificate` struct
	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")

	// create a custom server with `TLSConfig`
	s := &http.Server{
		Addr:    "",
		Handler: nil, // use `http.DefaultServeMux`
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// handle Deauthorize Callback URL
	http.HandleFunc("/deauth.go", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Deauthorize Callback URL [deauth.go]")
		fmt.Fprint(res, "Deauthorize Callback URL [deauth.go]")
	})

	// handle Valid OAuth Redirect URIs
	http.HandleFunc("/oauth.go", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Valid OAuth Redirect URIs [oauth.go]")
		fmt.Fprint(res, "Valid OAuth Redirect URIs [oauth.go]")
	})

	// handle Data Deletion Request URL
	http.HandleFunc("/datadelete.go", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Data Deletion Request URL[datadelete.go]")
		fmt.Fprint(res, "Data Deletion Request URL[datadelete.go]")
	})

	// handle Redirect URI to Check
	http.HandleFunc("/redirecturi.go", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Redirect URI to Check![redirecturi.go]")
		fmt.Fprint(res, "Redirect URI to Check![redirecturi.go]")
	})

	// run server on port "9000"
	log.Fatal(s.ListenAndServeTLS("", ""))

}
