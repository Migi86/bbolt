package main

import (
	"log"
	"flag"
	"net/http"
	"fmt"
	"bbolt/db"
	"bbolt/web"
)

var (
	dbLocation = flag.String("db-location", "/Users/michelegiglioni/desktop/bbolt/my.db", "The path to the bolt db database")
	httpAddr   = flag.String("http-addr", "127.0.0.1:8080", "HTTP host and port")
)

func parseFlags() {
	flag.Parse()

	if *dbLocation == "" {
		log.Fatalf("Must provide db-location")
	}
}

func main() {
	parseFlags()

	db, close, err := db.NewDatabase(*dbLocation)
	if err != nil {
		log.Fatalf("NewDatabase(%q): %v", *dbLocation, err)
	}
	defer close()
	fmt.Println("operation succeeded")
	srv := web.NewServer(db)

	http.HandleFunc("/get", srv.GetHandler)
	http.HandleFunc("/set", srv.SetHandler)

	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}