package zouyu

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"
)

// our parsing function "searchAndReplaceAll"
// allows for keywords as large as 3 chinese characters,
// but it could change to have to have more,
// without out too much effort.
var conversionTable = map[string]string{
	"走":  "go",
	"包裹": "package",
	"进口": "import",
	"程序": "func", //routine, taken from subroutine
	"主要": "main",
	"退还": "return",
	"如果": "if",
	"否则": "else",
	"出":  "A", // to export a function / struct-field
	"做":  "make",
}

func ConvertAll() {

	// get all files
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	// iterate through all files,
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
		cleanUpZhFile(name) //clean up file BEFORE converting it
		newfilename := "en_" + name[3:]
		writeConvertedFile(name, newfilename)
	}
}

// Adds build constraint to zh_ files
//
// It is best to clean up the file
// BEFORE converting it.
// Otherwise, there could be unnecessary build difficulties.
//
// Contribute: anything else necessary to make sure that the build process is as smooth as possible
func cleanUpZhFile(filename string) {

	// start by loading in the file
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	file := string(dat)

	containsBuildConstraint := strings.HasPrefix(file, "// +build ignore")
	if containsBuildConstraint == false {
		var re = regexp.MustCompile("包裹")
		file = re.ReplaceAllString(file, "// +build ignore\n\n包裹")
	}

	fmt.Println(file)
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

	// do all the translations for the file
	file = searchAndReplaceAll(file)

	// remove the build constraint from the en_ version of the file
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

// Traverses the source code (assumed to be UTF-8)
// and replaces every chinese keyword with its
// compiler-ready english keyword
//
// uses the package var conversionTable
//
// Resources:
// https://github.com/steele232/zouyu/wiki/GoLang-Source-Code-is-DEFINED-as-UTF-8
// https://blog.golang.org/strings
func searchAndReplaceAll(input string) string {

	// https://blog.golang.org/strings
	//
	// Besides the axiomatic detail that Go source code is UTF-8,
	// there's really only one way that Go treats UTF-8 specially,
	// and that is when using a for range loop on a string.
	// We've seen what happens with a regular for loop.
	// A for range loop, by contrast, decodes one UTF-8-encoded
	// rune ('rune' is synonymous with a 'UTF-8 code-point')
	// on each iteration.
	var isInQuotationBlock bool // starts at 0/false and flips on and off while parsing
	var sb strings.Builder
	s := []byte(input)
	for utf8.RuneCount(s) > 1 {
		runeValue, rSize := utf8.DecodeRune(s)
		quoteRune, _ := utf8.DecodeRuneInString(`"`)
		backQuoteRune, _ := utf8.DecodeRuneInString("`")

		if runeValue == quoteRune || // determine if we are in a quotation block
			runeValue == backQuoteRune {
			isInQuotationBlock = !isInQuotationBlock
			sb.WriteRune(runeValue)
			s = s[rSize:]
			continue
		}
		if isInQuotationBlock { // write without converting
			sb.WriteRune(runeValue)
			s = s[rSize:]
			continue
		}

		// get 1st rune (runeValue)
		// runeValue

		// get 2nd rune (r2)
		nextChar := s[rSize:]
		r2, r2Size := utf8.DecodeRune(nextChar)

		// get 3rd rune (r3)
		nextChar = nextChar[r2Size:]
		r3, r3Size := utf8.DecodeRune(nextChar)

		// now in source code ...
		char := string(runeValue)
		char2 := string(r2)
		char3 := string(r3)

		oneRune := char
		twoRunes := char + char2
		threeRunes := char + char2 + char3

		sizeOneRune := rSize
		sizeTwoRune := rSize + r2Size
		sizeThreeRune := rSize + r2Size + r3Size

		if val, ok := conversionTable[oneRune]; ok {
			// if this rune is in the conversionTable,
			// write the translated keyword
			sb.WriteString(val)
			// then move 1 rune down and start over.
			s = s[sizeOneRune:]
			continue
		} else if val, ok := conversionTable[twoRunes]; ok &&
			utf8.RuneCount(s) > 2 {
			// if these two runes are in the conversionTable,
			// write the translated keyword
			sb.WriteString(val)
			// then move 2 runes down and start over.
			s = s[sizeTwoRune:]
			continue
		} else if val, ok := conversionTable[threeRunes]; ok &&
			utf8.RuneCount(s) > 3 {
			// if these three runes are in the conversionTable,
			// write the translated keyword
			sb.WriteString(val)
			// then move 3 runes down and start over.
			s = s[sizeThreeRune:]
			continue
		} else {
			// otherwise, just write the character.
			sb.WriteRune(runeValue)
			// then move 1 rune down and start over.
			s = s[rSize:]
			continue
		}
	}
	sb.Write(s) // get that last character; the for (while) loop ended at len(s) == 1
	return sb.String()
}
