package zouyu

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func ConvertAll() {

	// iterate through all files,
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		name := f.Name()
		if len(name) < 6 {
			continue
		}
		if name[:3] != "zh_" {
			continue
		}
		if name[len(name)-3:] != ".go" {
			continue
		}
		writeConvertedFile(name, name[3:])
	}
}

// replace all chinese keywords with the english keywords
// so that the normal go build tool can read it and run it.
func writeConvertedFile(filename, newfilename string) {

	fmt.Println(filename)
	defer fmt.Println("")

	// start by loading in the file
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dat))

	cleanUpFile(filename)

	// use regexp to search and replace "走" with zou.
	var re = regexp.MustCompile(`走`)
	s := re.ReplaceAllString(string(dat), `go`)

	re = regexp.MustCompile(`\/\/\ \+build\ ignore\n+package`)
	convertedText := re.ReplaceAllString(s, `package`)

	fmt.Println(convertedText)

	// then write converted text to the new filename
	bytes := []byte(convertedText)
	err = ioutil.WriteFile("en_"+newfilename, bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conversion completed")
}

// cleanUpFile
// TODO maybe implement some of the features here if they appeal to users.
func cleanUpFile(filename string) {
	//IDEAS:
	// make sure there is a "// +build ignore" at the beginning of all "zh_" file
	// and put it in, if there isn't one already...

	// anything else necessary to make sure that the build process is as smooth as possible
}
