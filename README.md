# ZouYu (走语/走語)

## Overview

#### Purpose 

Help young chinese people learn to program in their native mandarin

## Usage

- Install it (will be go gettable)
- ```zou``` takes all "zh_..." (e.g. "zh_main.go") files in the current directory and converts them to "en_..." files, which are then ready for compilation.
- You can combine ```zou``` with go commands on bash like so: ```zou && go run .``` or ```zou && go build .```. You can 
- [**Not Started**] (Create a Youtube/Youku video showing how to use it)

## Installation

- ```go install "github.com/steele232/zouyu/zou"```

## Roadmap

- [**Complete**] Make a way to translate and make compiler-ready source files from completely or mostly chinese source files
- [**Complete**] 
- [**In Progress**] Traditional characters
- [**In Progress**] Builtin functions, types [(Builtin package)](https://golang.org/pkg/builtin/)
- [**In Progress**] Full Test Coverage (tests using all the keywords)
- [**Not Started**] Chinese README
- [**Not Started**] A few nice CLI features (Verbosity, Multiple Directories)
- [**Not Started**] Wrapper library around stdlib functions (including useful comments) so that we can use stdlib functions but still write in mandarin.
- [**Not Started**] Educational Resources to use while teaching Programming to young native Mandarin speakers. (Getting Started Guide. Guide for Teachers. Other Resources)

## Disclaimers & Known Issues

- I am working on linux and mac computers so far, so I don't know exactly how this will work on windows.

## Contributing

Feel free to help this project succeed by

- Giving the Repo a Star on [Github](www.github.com/steele232/zouyu)!
- Subscribe to notifications by "Watching" the Repo on [Github](www.github.com/steele232/zouyu)!
- Consider helping with Translations if you speak Mandarin!
- **Try using the tool!**
- Give Feedback! (Find bugs, Report Issues, Consider possible Feature Requests etc)

### Translation Quality Confidence Levels

Here is a table of the Confidence levels of Translations between ZouYu keywords and original GoLang keywords:

| Character(s) | Go Keyword | Confidence Level (%) | Links |
| --- | --- | --- | --- |
| `"走"` | "go" | 100 | |
| `"包裹"` | "package" | 100 | |
| `"进口"` |  "import" | 100 | |
| `"進口"` |  "import" | 100 | |
| `"函数"` |  "func" | 99 | [1][1] |
| `"函數"` |  "func" | 99 | [1][1] |
| `"主要"` |  "main" | 100 | |
| `"返回"` |  "return" | 100 | [1][1], [2][2] |
| `"如果"` |  "if" | 100 | |
| `"否则"` |  "else" | 99 | |
| `"否則"` |  "else" | 99 | |
| `"出"` |  "A" (to export) | 50 | |
| `"初始化"` |  "make" | 70 | [1][1] |
| `"循环"` |  "for" | 99 | [3][3] |
| `"循環"` |  "for" | 99 | [3][3] |
| `"范围"` |  "range" | 90 | |
| `"範圍"` |  "range" | 90 | |
| `"打断"` |  "break" | 70 | |
| `"中断"` |  "break" | 90 | [5][5] |
| `"跳出"` |  "break" | 30 | [6][6] |
| `"继续"` |  "continue" | 100 |  |
| `"繼續"` |  "continue" | 100 |  |
| `"开关语句"` |  "switch" | 80 | [6][6] |
| `"開關語句"` |  "switch" | 80 | [6][6] |
| `"切换"` |  "switch" | 30 | [6][6] |
| `"选择"` |  "select" | 99 | [8][8] |
| `"選擇"` |  "select" | 99 | [8][8] |
| `"假如"` |  "case" | 99 | [9][9] |
| `"默认"` |  "default" | 99 | [10][10] |
| `"默認"` |  "default" | 99 | [10][10] |
| `"去"` |  "goto" | 95 | [11][11]|
| `"落下"` |  "fallthrough" | 90 | [12][12] |
| `"推迟"` |  "defer" | 95 | [13][13] |
| `"恐慌"` |  "panic" | 100 | [14][14] |
| `"整数"` |  "int" | 99 | |
| `"整數"` |  "int" | 99 | |
| `"浮点数"` |  "float" | 99 | [15][15] |
| `"浮點數"` |  "float" | 99 | [15][15] |
| `"字符串"` |  "string" | 99 | |
| `"字节"` |  "byte" | 100 | |
| `"字節"` |  "byte" | 100 | |
| `"无符号"` |  "uint" (shorten to "u"?) | 50 | |
| `"無符號"` |  "uint" (shorten to "u"?) | 50 | |
| `"常量"` |  "const" | 99 | [16][16] |
| `"变量"` |  "var" | 99 | [17][17], [1][1] |
| `"變量"` |  "var" | 99 | [17][17], [1][1] |
| `"映射"` |  "map" | 95 | [2][2] |
| `"信道"` |  "chan" | 95 | [18][18] |
| `"结构体"` |  "struct" | 100 | [15][15] |
| `"結構體"` |  "struct" | 100 | [15][15] |
| `"类型"` |  "type" | 100 | [19][19] |
| `"類型"` |  "type" | 100 | [19][19] |
| `"接口"` |  "interface" | 99 | [2][2], [1][1] |


[1]: https://blog.csdn.net/tzs919/article/details/53571632 
[2]: https://go-zh.org/doc/codewalk/sharemem/ "Go Docs - ShareMem"
[3]: https://fanyi.baidu.com/#en/zh/for%20loop "Baidu Translate @ Loop"
[4]: https://fanyi.baidu.com/#en/zh/range "Baid Translate @ Range"
[5]: https://blog.csdn.net/u014805066/article/details/50587309 "Dictionary of Sorts .. Professional Language.. ???"
[6]: https://baike.baidu.com/item/switch/18601752 "Switch article"
[7]: https://fanyi.baidu.com/#zh/en/%E5%89%8D%E8%BF%9B "Baidu Translate @ Continue"
[8]: https://go-zh.org/ref/spec.old#Select%E8%AF%AD%E5%8F%A5 "Chinese Go Documentation @ Switch"
[9]: https://translate.google.com/#view=home&op=translate&sl=en&tl=zh-CN&text=if "Google Translate @ Case"
[10]: https://fanyi.baidu.com/#en/zh/Default%20Settings "Baidu Translate @ Default"
[11]: https://golang.org/ref/spec#Goto_statements "Golang Documentation @ Goto Statements"
[12]: https://fanyi.baidu.com/#en/zh/fall "Baidu Translate @ Fall"
[13]: https://fanyi.baidu.com/#en/zh/defer "Baidu Translate @ Defer"
[14]: https://go-zh.org/ref/spec.old "Chinese Go Documentation @ Panic"
[15]: https://blog.csdn.net/u014805066/article/details/50587309 "Blog @ Floating Point Number"
[16]: https://fanyi.baidu.com/#en/zh/const "Baidu Translate @ Const"
[17]: https://fanyi.baidu.com/#en/zh/variable "Baidu Translate @ Variable"
[18]: https://go-zh.org/ref/spec.old#%E4%BF%A1%E9%81%93%E7%B1%BB%E5%9E%8B "Chinese Go Documentation @ Chan (Channels)"
[19]: https://fanyi.baidu.com/#en/zh/struct "Baidu Translate @ Struct"

## Thanks for looking at this Repo!

Please think about [Contributing](Contributing)
