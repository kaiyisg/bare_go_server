package api

import (
  "fmt"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

// var HttpRouter

// func main() {
//
//   HttpRouter := httprouter.new()
//
// }

type apiRoute struct {
	Method string
	Path   string
}
