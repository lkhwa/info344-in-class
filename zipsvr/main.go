package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//Using pointers are efficient, can see changes on other side
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	//w.Write([]byte("Hello, World!"))
	fmt.Fprintf(w, "Hello %s!", name)
	/*Fprintf will take a template string and values and write the result to
	any I/O writer (the first parameter), such as ResponseWriter.*/
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC()
	stats := &runtime.MemStats{}
	//& creates a struct on the heap (memory), get a pointer to the block of memory in RAM.
	//stats var is a pointer to the MemStats{} struct, not an instance itself.
	runtime.ReadMemStats(stats)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	//fmt.Println("Hello World!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)

	//mux.HandleFunc("/hello/", helloHandler)
	//Putting a / at the end of the resource path allows the user to specify a unique identifier.
	//So if the requested resource path starts with the parameter, you can then look at what the last unique bit is.

	fmt.Printf("server is listening at http://localhost:4000\n")
	log.Fatal(http.ListenAndServe("localhost:4000", mux))
}
