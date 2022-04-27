package findncode

import (
	"bufio"
	"bytes"
	"fmt"
	"image/jpeg"
	"log"
	"os"

	"github.com/auyer/steganography"
)

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

func imgStegDecode(input string) {
	inFile, _ := os.Open(input)
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, _ := jpeg.Decode(reader)

	size := steganography.GetMessageSizeFromImage(img)

	msg := steganography.Decode(size, img)
	fmt.Printf(string(msg))
}

//Finds all images in filepath
//Exports the paths to slice
func find_img(filepath string) {

}

//Takes slice of file paths, encodes each of them with appropriate
//CC license and move the encoded images to the
func PrepImgs() {

}
