package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/lkhwa/info344-in-class/zipsvr/handlers"
	"github.com/lkhwa/info344-in-class/zipsvr/models"
)

const zipsPath = "/zips/"

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
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	tlskey := os.Getenv("TLSKEY")
	tlscert := os.Getenv("TLSCERT")
	if len(tlskey) == 0 || len(tlscert) == 0 {
		log.Fatal("please set TLSKEY and TLSCERT")
	}
	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatal("error loading zips: %v", err)
	}
	log.Printf("loaded %d zips", len(zips))

	//TASK: return to retrieve all city=Seattle zip codes
	//HOW: use a map, getting constant-time access to a value, given a key
	cityIndex := models.ZipIndex{} //{} creates a static instance of the map
	for _, z := range zips {       // right side is something that's iterable, such as a slice; left side: index, ref to item
		cityLower := strings.ToLower(z.City) //convert to lowercase, using strings pkg and ToLower func in the pkg
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	//fmt.Println("Hello World!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)

	//mux.HandleFunc("/hello/", helloHandler)
	//Putting a / at the end of the resource path allows the user to specify a unique identifier.
	//So if the requested resource path starts with the parameter, you can then look at what the last unique bit is.

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: zipsPath,
	}
	mux.Handle(zipsPath, cityHandler)

	fmt.Printf("server is listening at http://%s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlscert, tlskey, mux))
}
