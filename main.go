package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/auyer/steganography"
)

/*
TO-DO

- Detect when a file is being downloaded from the website
- Embed copy right info to file after the download request
- Determine how much we should change image
	- Which position of bit to we flip based on
	  the copy right of work.
- Get the website to work (localhost:3000)
- Decode??
*/

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

//Not the final function just an example of how to use the library
//Input = filepath
//Message = message to be encoded
func imgStegEncode(input string, message string, output string) {
	inFile, _ := os.Open(input)
	reader := bufio.NewReader(inFile)
	img, _ := jpeg.Decode(reader) //Needs logic for png images if we want

	buf := new(bytes.Buffer)
	err := steganography.Encode(buf, img, []byte(message))

	if err != nil {
		log.Printf("Error %v", err)
	}

	outFile, _ := os.Create(output)
	buf.WriteTo(outFile)
	outFile.Close()
}

func imgStegDecode(input string, message string, output string) {
	inFile, _ := os.Open(input)
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, _ := jpeg.Decode(reader)

	size := steganography.GetMessageSizeFromImage(img)

	msg := steganography.Decode(size, img)
	fmt.Printf(string(msg))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
