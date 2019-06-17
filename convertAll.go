package zouyu

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"
)

// our parsing function "SearchAndReplaceAll"
// allows for keywords as large as 3 chinese characters,
// but it could change to have to have more,
// without out too much effort.
//
// Possible Changes
// "u" / "uint" -> maybe should just be "u" so that the full 无符号整数 can be used for "uint", ... but IDK.
// (No) I won't put in "error". That would be weird because it's part of the std library, not a keyword..  https://go-zh.org/doc/codewalk/sharemem/
// "case" 字句 ??
// "switch" vs "select" .. A little fuzzy.. I differentiate that select choose channels of communication.
// "range" a little weird.. but IDK
var conversionTable = map[string]string{
	"走":  "go",          //zou3
	"包":  "package",     // bao3 (guo3?) ; https://github.com/zhanming/go-tour-cn
	"导入": "import",      // dao3 ru4 ; https://github.com/zhanming/go-tour-cn
	"導入": "import",      // jin4 kou3 traditional ; https://github.com/zhanming/go-tour-cn
	"函数": "func",        // han2 shu4 ; function (math) https://blog.csdn.net/tzs919/article/details/53571632, https://github.com/zhanming/go-tour-cn
	"函數": "func",        // han2 shu4 traditional ; https://github.com/zhanming/go-tour-cn
	"主要": "main",        // zhu3 yao4 ; ???
	"返回": "return",      // fan3 hui2 ; https://go-zh.org/doc/codewalk/sharemem/, https://blog.csdn.net/tzs919/article/details/53571632, https://github.com/zhanming/go-tour-cn
	"如果": "if",          // ru2 guo3
	"否则": "else",        // fou3 ze2
	"否則": "else",        // fou3 ze2 traditional
	"导出": "A",           // dao3 chu1 ; to export a function / struct-field ; https://github.com/zhanming/go-tour-cn
	"创建": "make",        // chuang4 jian4 ; create ; https://github.com/zhanming/go-tour-cn
	"創建": "make",        // chuang4 jian4 ; create ; traditional
	"循环": "for",         // xun2 huan2 ; loop ... https://fanyi.baidu.com/#en/zh/for%20loop
	"循環": "for",         // xun2 huan2 traditional
	"范围": "range",       // fan4 wei2 ; range ... https://fanyi.baidu.com/#en/zh/range
	"範圍": "range",       // fan4 wei2 traditional
	"终止": "break",       // zhong1 zhi3 ; https://github.com/zhanming/go-tour-cn
	"終止": "break",       // zhong1 zhi3 ; traditional
	"继续": "continue",    // ji4 xu4 ;
	"繼續": "continue",    // ji4 xu4 ; traditional
	"开关": "switch",      // kai1 guan1 ; https://baike.baidu.com/item/switch/18601752
	"開關": "switch",      // kai1 guan1 traditional
	"选择": "select",      // xuan3 ze2 ; https://go-zh.org/ref/spec.old#Select%E8%AF%AD%E5%8F%A5
	"選擇": "select",      // xuan3 ze2 traditional
	"假如": "case",        // jia3 ru2 ; https://translate.google.com/#view=home&op=translate&sl=en&tl=zh-CN&text=if
	"默认": "default",     // mo4 ren4 ; https://fanyi.baidu.com/#en/zh/Default%20Settings, https://github.com/zhanming/go-tour-cn
	"默認": "default",     // mo4 ren4  traditional
	"去到": "goto",        // qu4 ; https://golang.org/ref/spec#Goto_statements
	"落下": "fallthrough", // luo4 xia4 ; https://fanyi.baidu.com/#en/zh/fall
	"推迟": "defer",       // tui1 chi2 ; https://fanyi.baidu.com/#en/zh/defer
	"恐慌": "panic",       // kuang3 huang1 ; https://go-zh.org/ref/spec.old

	"整数":  "int",       // zheng3 shu4 ;
	"整數":  "int",       // zheng3 shu4 traditional
	"浮点数": "float",     // fu2 dian3 shu4 ; https://blog.csdn.net/u014805066/article/details/50587309
	"浮點數": "float",     // fu2 dian3 shu4 traditional
	"字符串": "string",    // zi4 fu2 chuan4
	"字节":  "byte",      // zi4 jie2
	"字節":  "byte",      // zi4 jie2 traditional
	"无符号": "uint",      // wu2 fu2 hao4 ; unsigned
	"無符號": "uint",      // wu2 fu2 hao4 traditional
	"常量":  "const",     // chang2 liang4 ; https://fanyi.baidu.com/#en/zh/const, https://github.com/zhanming/go-tour-cn
	"变量":  "var",       // bian4 liang4 ; https://fanyi.baidu.com/#en/zh/variable, https://blog.csdn.net/tzs919/article/details/53571632, https://github.com/zhanming/go-tour-cn
	"變量":  "var",       // bian4 liang4 traditional
	"映射":  "map",       // ying4 she4 ; https://go-zh.org/doc/codewalk/sharemem/
	"管道":  "chan",      // guan3 dao4 ; https://github.com/zhanming/go-tour-cn
	"结构体": "struct",    // jie2 gou4 ti3; struct https://blog.csdn.net/u014805066/article/details/50587309, https://github.com/zhanming/go-tour-cn
	"結構體": "struct",    // jie2 gou4 ti3 traditional
	"类型":  "type",      // lei4 xing2 ; type https://fanyi.baidu.com/#en/zh/struct, https://github.com/zhanming/go-tour-cn
	"類型":  "type",      // lei4 xing2 tradtional
	"接口":  "interface", // jie1 kou3 ; interface https://go-zh.org/doc/codewalk/sharemem/, https://blog.csdn.net/tzs919/article/details/53571632, https://github.com/zhanming/go-tour-cn

	"指针": "pointer", // zhi3 zhen1 ; pointer, ??? will need to be adapted for uintptr, etc. ; https://github.com/zhanming/go-tour-cn

	/*
		"切片": "slice",       // qie1 pian4 ; https://github.com/zhanming/go-tour-cn
		"数组": "array",       // shu4 zu3 ; array ;
		"數組": "array",       // shu4 zu3 ; array ; traditional
		"容量": "cap",         // rong2 liang4 ; (building) capacity ; https://github.com/zhanming/go-tour-cn
		"值":  "value",       // zhi2 ; value ; https://github.com/zhanming/go-tour-cn
		"闭包": "closure",     // bi4 bao1 ; closure (mathematics) ; https://github.com/zhanming/go-tour-cn
		"条件": "condition",   // tiao2 jian4 ;
		"错误": "error",       // cuo4 wu4 ; (builtin) https://github.com/zhanming/go-tour-cn, etc (type error interface { Error() string })
		"并发": "concurrency", // bing4 fa1 ; concurrency --
		"环境": "environment", // environment ;
		"缓冲": "buffer",      // buffer(ed) ; as in bufferend channel
	*/
}

func ConvertAllFilesInDir() {

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

	//convert the file.
	file = ConvertFile(file)

	fmt.Println(file)
	// then write converted text to the new filename
	bytes := []byte(file)
	err = ioutil.WriteFile(newfilename, bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conversion completed")
}

// ConvertFile takes an "zh_" file as a string
// and converts it to an "en_" file as a string,
// which means that it is ready to be compiled.
//
// Another useful detail is that if there is a
// build ignore ("// +build ignore") at the
// beginning of the "zh_" file, then it will be
// removed.
func ConvertFile(file string) string {

	// do all the translations for the file
	file = SearchAndReplaceAll(file)

	// remove the build constraint from the en_ version of the file
	file = removeBuildConstraints(file)

	return file
}

// removeBuildConstraints assumes that an english
// source file with a build ignore constraint is
// inputted as a string.
// In the context of this program, that would mean
// that the "zh_" file has been converted into a
// "zh_" file but didn't have the build constraints
// removed yet.
func removeBuildConstraints(file string) string {

	var re = regexp.MustCompile(`\/\/\ \+build\ ignore\n+package`)
	file = re.ReplaceAllString(file, `package`)
	return file
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
func SearchAndReplaceAll(input string) string {

	// https://blog.golang.org/strings
	//
	// Besides the axiomatic detail that Go source code is UTF-8,
	// there's really only one way that Go treats UTF-8 specially,
	// and that is when using a for range loop on a string.
	// We've seen what happens with a regular for loop.
	// A for-range-loop, by contrast, decodes one UTF-8-encoded
	// rune ('rune' is synonymous with a 'UTF-8 code-point')
	// on each iteration.
	var isInQuotationBlock bool // starts at 0/false and flips on and off while parsing
	var isInLineComment bool    // "//"
	var isInBlockComment bool   // "/* */"
	var sb strings.Builder
	s := []byte(input)
	for utf8.RuneCount(s) > 1 {
		runeValue, rSize := utf8.DecodeRune(s)
		quoteRune, _ := utf8.DecodeRuneInString(`"`)
		backQuoteRune, _ := utf8.DecodeRuneInString("`")
		quote1Rune, _ := utf8.DecodeRuneInString("“")
		quote2Rune, _ := utf8.DecodeRuneInString("”")
		quote3Rune, _ := utf8.DecodeRuneInString("‘")
		quote4Rune, _ := utf8.DecodeRuneInString("’")
		quote5Rune, _ := utf8.DecodeRuneInString("'")
		newLineRune, _ := utf8.DecodeRuneInString("\n")
		leftSlashRune, _ := utf8.DecodeRuneInString("/")
		asteriskRune, _ := utf8.DecodeRuneInString("*")

		// get 1st rune (runeValue)
		// runeValue

		// get 2nd rune (r2)
		nextChar := s[rSize:]
		r2, r2Size := utf8.DecodeRune(nextChar)

		// get 3rd rune (r3)
		nextChar = nextChar[r2Size:]
		r3, r3Size := utf8.DecodeRune(nextChar)

		/* determine if we are in a comment block */
		// // line comment
		if runeValue == leftSlashRune &&
			r2 == leftSlashRune &&
			isInBlockComment == false {

			isInLineComment = true
		}
		if runeValue == newLineRune {
			isInLineComment = false
		}

		// /**/ block comment
		if runeValue == leftSlashRune &&
			r2 == asteriskRune {

			isInBlockComment = true
		}
		if runeValue == asteriskRune &&
			r2 == leftSlashRune {

			isInBlockComment = false
		}

		// write and don't convert if in comments
		if isInLineComment || isInBlockComment {
			sb.WriteRune(runeValue)
			s = s[rSize:]
			continue
		}

		// determine if we are in a quotation block
		if runeValue == quoteRune ||
			runeValue == backQuoteRune ||
			runeValue == quote1Rune ||
			runeValue == quote2Rune ||
			runeValue == quote3Rune ||
			runeValue == quote4Rune ||
			runeValue == quote5Rune {

			isInQuotationBlock = !isInQuotationBlock
			sb.WriteRune(runeValue)
			s = s[rSize:]
			continue
		}
		// write and don't convert if in quotation block
		if isInQuotationBlock {
			sb.WriteRune(runeValue)
			s = s[rSize:]
			continue
		}

		// get 1st rune (runeValue)
		// runeValue

		// get 2nd rune (r2)
		nextChar = s[rSize:]
		r2, r2Size = utf8.DecodeRune(nextChar)

		// get 3rd rune (r3)
		nextChar = nextChar[r2Size:]
		r3, r3Size = utf8.DecodeRune(nextChar)

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
