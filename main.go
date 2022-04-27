package main

import (
	"html/template"
	"net/http"
	"os"
)

/*
TO-DO

- Go function to seek files to encode and encode them before site is launched
	- Has specific path passed to it (maybe CLI???)
- Only Using Creative commons
- Get the website to work (localhost:3000)
- Decode
*/

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

//Not the final function just an example of how to use the library
//Input = filepath
//Message = message to be encoded

func main() {

	filepath := "0853-Delta-5-Edit.jpg"
	message := "This is the message to be encoded"
	outfile := "encoded.jpg"

	imgStegEncode(filepath, message, outfile)
	//imgStegDecode("encoded.jpg")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
