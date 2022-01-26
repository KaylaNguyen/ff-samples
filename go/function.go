package hello

import (
	"net/http"
)

// HelloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello, World!\n")
	panic("Intended panic")
}
