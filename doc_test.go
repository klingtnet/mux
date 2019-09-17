package mux

import (
	"testing"
)

func TestPerRouteDocumentation(t *testing.T) {
	// TODO: setup a more complicated example using subrouters
	router := NewRouter()
	router.HandleFunc("/", nil).Methods("GET").Doc("API overview")
	router.HandleFunc("/version", nil).Methods("GET").Doc("service version")
	router.HandleFunc("/status", nil).Methods("GET").Doc("service status")
	router.HandleFunc("/foo/{bar}", nil).Methods("POST, GET").Doc("update foo")

	err := router.Walk(func(route *Route, _ *Router, _ []*Route) error {
		doc := route.GetDoc()
		if doc == "" {
			t.Fatal("expected doc to be non-empty")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)		
	}
}