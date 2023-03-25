package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"Hello" : "world!"}`))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

// package main

// import (
//   "fmt"
//   "log"
//   "net/http"
//   "os"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
//   log.Print("Hello world received a request.")
//   target := os.Getenv("TARGET")
//   if target == "" {
//     target = "World"
//   }
//   fmt.Fprintf(w, "Hello %s!\n", target)
// }

// func main() {
//   log.Print("Hello world sample started.")

//   http.HandleFunc("/", handler)

//   port := os.Getenv("PORT")
//   if port == "" {
//     port = "8080"
//   }

//   log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
// }
