package zouyu

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

var conversionTable = map[string]string{
	"走":   "go",
	"包裹":  "package",
	"进口":  "import",
	"子程序": "func", //subroutine
	"主要":  "main",
	"退还":  "return",
	"如果":  "if",
	"否则":  "else",
	"出":   "A", // to export a function / struct-field
}

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

	// cleanUpFile(filename)

	file := string(dat)

	// do all the translations
	for k, v := range conversionTable {
		file = searchAndReplace(file, k, v)
	}

	// remove the build constraint by itself
	var re = regexp.MustCompile(`\/\/\ \+build\ ignore\n+package`)
	file = re.ReplaceAllString(file, `package`)
	fmt.Println(file)

	// then write converted text to the new filename
	bytes := []byte(file)
	err = ioutil.WriteFile("en_"+newfilename, bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conversion completed")
}

func searchAndReplace(input, findRegEx, replace string) string {
	var re = regexp.MustCompile(findRegEx)
	s := re.ReplaceAllString(input, replace)
	return s
}

// cleanUpFile
// TODO maybe implement some of the features here if they appeal to users.
func cleanUpFile(filename string) {
	//IDEAS:
	// make sure there is a "// +build ignore" at the beginning of all "zh_" file
	// and put it in, if there isn't one already...

	// anything else necessary to make sure that the build process is as smooth as possible
}
