package main

import (
	"log"
	"net/http"

	"github.com/farwawaqar/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()              //creating r variable which will intialize the router here
	routes.RegisterBookStoreRoutes(r) // we are passing the router r to RegisterBookStoreRoutes inside pkg/routes
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	//ListenandServe is the function the helps use to create the server at our give port
	//it is inside http package

}
