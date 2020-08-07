package main

import (
	"net/http"

	routes "OrderApi/router"
)

func main() {

	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)

}
