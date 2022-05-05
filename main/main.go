package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/auyer/steganography"
	"html/template"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

const (
	CC_BY         = "CC-BY"
	CC_BY_SA      = "CC-BY-SA"
	CC_BY_ND      = "CC-BY-ND"
	CC_BY_NC      = "CC-BY-NC"
	CC_BY_NC_SA   = "CC-BY-NC-SA"
	CC_BY_NC_ND   = "CC-BY-NC-ND"
	ccby_path     = "/CC-BY/CC-BY.txt"
	ccbync_path   = "/CC-BY-NC.txt"
	ccbyncnd_path = "/CC-BY-NC-ND.txt"
	ccbyncsa_path = "/CC-BY-NC-SA.txt"
	ccbynd_path   = "/CC-BY-ND.txt"
	ccbysa_path   = "/CC-BY-SA.txt"
	dir_path_out  = "/img"
)

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
	byte_msg := []byte(message)
	buf := new(bytes.Buffer)

	if uint32(len(byte_msg)) < steganography.MaxEncodeSize(img) {
		err := steganography.Encode(buf, img, byte_msg)
		if err != nil {
			log.Printf("Error %v", err)
		}
		outFile, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		buf.WriteTo(outFile)
		outFile.Close()
	} else {
		log.Printf("Message size to large for %v \n Skipping Input", input)
	}

}

func imgStegDecode(input string, output string) (msg string) {
	inFile, _ := os.Open(input)
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, _ := jpeg.Decode(reader)

	size := steganography.GetMessageSizeFromImage(img)

	dec_msg := steganography.Decode(size, img)
	fmt.Printf("Decoded message: %v", string(dec_msg))
	return string(dec_msg)
}

//Takes slice of file paths, encodes each of them with appropriate
//CC license and move the encoded images to the
func encode_Imgs_upload(root, pattern string) {

	//Make the output Directory

	if _, err := os.Stat((root + dir_path_out)); !os.IsNotExist(err) {
		log.Printf("Output Directory exists")
	} else {
		err := os.Mkdir((root + dir_path_out), 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	//Get all image filepaths for encoding
	//images, err := find_img(root, pattern)
	p := root + "/*/*.jpg"
	images, err := filepath.Glob(p)
	if err != nil {
		log.Fatal(err)
	}

	//Parse the filepaths, encode images with approrpiate CC, write to output dir
	for _, s := range images {
		log.Printf("Encoding Image %v \n", s)
		var copyright []byte
		if strings.Contains(s, CC_BY) {
			copyright, err = ioutil.ReadFile((root + ccby_path)) //Get CC text
			if err != nil {
				log.Fatal(err)
			}
		} else if strings.Contains(s, CC_BY_NC) {
			copyright, err = ioutil.ReadFile((root + ccbync_path)) //Get CC text
			if err != nil {
				log.Fatal(err)
			}
		} else if strings.Contains(s, CC_BY_NC_ND) {
			copyright, err = ioutil.ReadFile((root + ccbyncnd_path)) //Get CC text
			if err != nil {
				log.Fatal(err)
			}
		} else if strings.Contains(s, CC_BY_NC_SA) {
			copyright, err = ioutil.ReadFile((root + ccbyncsa_path)) //Get CC text
			if err != nil {
				log.Fatal(err)
			}
		} else if strings.Contains(s, CC_BY_ND) {
			copyright, err = ioutil.ReadFile((root + ccbynd_path)) //Get CC text
			if err != nil {
				log.Fatal(err)
			}
		} else if strings.Contains(s, CC_BY_SA) {
			copyright, err = ioutil.ReadFile((root + ccbysa_path)) //Get CC text
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("Could not find the copyright file specified")
		}

		msg := string(copyright) //Convert to string, just to be sure
		log.Printf("Copyright : %v \n", msg)

		file_name := filepath.Base(s)
		outfile := root + dir_path_out + "/cpoyrighted-" + file_name
		//Create the output file name
		if _, err := os.Stat(outfile); !os.IsNotExist(err) {
			log.Printf("Encoded File exists\n")
		} else {
			log.Printf("Output File: %v \n", outfile)
			imgStegEncode(s, msg, outfile) //encode CC to current image
		}

	}
}

func main() {
	//log.Printf("Please input user go root (i.e. C:/username/go/src/)") //For the future we want this to just be a filepath
	args := os.Args

	root := args[1] + "/Cyber-Security-Final-Project/"
	encode_Imgs_upload(root, "*.jpg")
	//log.Printf("%v", matches)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
