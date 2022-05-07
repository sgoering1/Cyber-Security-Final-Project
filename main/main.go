package main

//Image Stenography to Embed Creative Copyright information
//Authors: Sam Goering, Colin Woods
//Uwyo Cybersecurity-4010 2022

import (
	"bufio"
	"bytes"
	"github.com/auyer/steganography"
	"html/template"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//Html template variable
var tpl = template.Must(template.ParseFiles("index.html"))

//Filepath constants
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
	dir_path_out  = "main/img/"
)

//Sends html to webserver
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

// Encoding Function takes in filepath to image, message to be encoded, and the desired output
// string, returns nothing.

// Note: The encode will not work with jpegs as the encoding of a jpeg is much more specific
// than a png images.
func imgStegEncode(input string, message string, output string) {

	//Create file readers and writer for the encoding
	inFile, _ := os.Open(input)
	reader := bufio.NewReader(inFile)
	img, _ := png.Decode(reader) //Needs to be png
	byte_msg := []byte(message)
	buf := new(bytes.Buffer)

	if uint32(len(byte_msg)) < steganography.MaxEncodeSize(img) {
		err := steganography.Encode(buf, img, byte_msg) //Actual encoding
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

//Decode function just take in a filepath and returns the
//String of the decoded images copyright
func imgStegDecode(input string) (msg string) {
	inFile, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, err := png.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	size := steganography.GetMessageSizeFromImage(img) //Needed for the decode to know

	dec_msg := steganography.Decode(size, img)
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
	p := root + "/CC*/*.png"
	images, err := filepath.Glob(p)
	if err != nil {
		log.Fatal(err)
	}

	//Parse the filepaths, encode images with approrpiate CC, write to output dir
	index := 0
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
			log.Printf("Could not find the copyright file specified")
		}

		msg := string(copyright) //Convert to string, just to be sure
		log.Printf("Copyright : %v \n", msg)

		//file_name := filepath.Base(s)
		i_s := strconv.Itoa(index)
		outfile := root + dir_path_out + i_s + ".png"
		//Create the output file name
		log.Printf("Index %v", index)
		if _, err := os.Stat(outfile); !os.IsNotExist(err) {
			log.Printf("Encoded File exists\n")
		} else {
			log.Printf("Output File: %v \n", outfile)
			imgStegEncode(s, msg, outfile) //encode CC to current image
		}
		index += 1
	}
}

func main() {
	args := os.Args

	//Simple CLI handling just 2 commands -d and -e
	if args[1] == "-d" {
		log.Printf(args[2])
		cc_msg := imgStegDecode(args[2])
		log.Printf("Cpoyright Message: %v", cc_msg)
		os.Exit(3)
	} else if args[1] == "-e" { //This has to be C:/user/go/src/
		root := args[2] + "/Cyber-Security-Final-Project/"
		encode_Imgs_upload(root, "*.png")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8887"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
