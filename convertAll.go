package zouyu

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
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
	"做":   "make",
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
		cleanUpFile(name) //clean up file BEFORE converting it
		newfilename := "en_" + name[3:]
		writeConvertedFile(name, newfilename)
	}
}

// It is best to clean up the file
// BEFORE converting it.
// Otherwise, there could be unnecessary build difficulties.
//
// Contribute: anything else necessary to make sure that the build process is as smooth as possible
func cleanUpFile(filename string) {

	// start by loading in the file
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	file := string(dat)

	containsBuildConstraint := strings.HasPrefix(file, "// +build ignore")
	if containsBuildConstraint == false {
		file = searchAndReplace(file, "包裹", "// +build ignore\n\n包裹")
	}

	// then write converted text to the new filename
	bytes := []byte(file)
	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		log.Fatal(err)
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

	file := string(dat)

	fmt.Println(file)

	// makes sure that the zh_ file has a build constraint on it
	cleanUpFile(filename)

	// do all the translations for the file
	for k, v := range conversionTable {
		file = searchAndReplace(file, k, v)
	}

	// remove the build constraint by itself
	var re = regexp.MustCompile(`\/\/\ \+build\ ignore\n+package`)
	file = re.ReplaceAllString(file, `package`)
	fmt.Println(file)

	// then write converted text to the new filename
	bytes := []byte(file)
	err = ioutil.WriteFile(newfilename, bytes, 0644)
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
