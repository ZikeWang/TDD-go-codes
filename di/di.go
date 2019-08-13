package main

import (
	"fmt"
	"io"
	"net/http"
)

//Greet uses io.writer to write string to buffer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//GreetHandler write string via http
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Web")
}

func main() {
	//Greet(os.Stdout, "Elise")
	http.ListenAndServe(":5000", http.HandlerFunc(GreetHandler))
}
