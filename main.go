package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("sample-plugin starts.")
	http.ListenAndServe(":5000", nil)
}
