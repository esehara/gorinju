package gorinju

import (
	"fmt"
	"net/http"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm alive.", r.URL.Path[1:])
}

func server() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9999", nil)
}

var is_runserver bool = false

func runserver() {
	if (! is_runserver ) {
		go server()
		is_runserver = true
	}
}

func TestHandGun(t *testing.T) {
	runserver()
	HandGun("http://localhost:9999", 200)
}

func TestMachineGun (t *testing.T) {
	runserver()
	MachineGun("http://localhost:9999", 200)
}
