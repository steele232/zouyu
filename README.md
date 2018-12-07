# ZouYu (走语)

## Overview

#### Purpose 

Help young chinese people learn to program in their native mandarin

## Roadmap

### Completed
- Zou -> Go
- // +build ignore (delete from english ones)
- zh_ file for chinese, en_ file for english (compilable)
- package
- import
- func
- main
- return
- if
- else
- export variables, functions, etc
- make

### Future Developments

#### All Keywords
- for 
- range
- break
- continue
- switch
- type
- interface

#### All Data Types
- string
- channel
- int (32,64)
- float (32,64)

#### Other
- traditional characters
- wrapper library around stdlib functions so that we can use stdlib functions but still write in mandarin.

## Usage

- Install it (will be go gettable)
- ```zou``` takes all "zh_..." (e.g. "zh_main.go") files in the current directory and converts them to "en_..." files, which are then ready for compilation.
- You can combine ```zou``` with go commands on bash like so: ```zou && go build .``` or ```zou && go run .```. You can 
- (Create a Youtube/Youku video showing how to use it)

## Installation

- TODO -> go gettable 
- ? Try ```go get "github.com/steele232/zouyu"```


## Contributing

I will be reaching out to my own friends about help with translations and usage. 

If you stumble upon this project, feel free to start using it, and I would love to hear how it's working for you. I have a lot of plans for this project and other projects to go with it, so I could use some help :D

## Disclaimers & Known Issues

- I am working on linux and mac computers so far, so I don't know exactly how this will work on windows.
- Because I'm doing a pretty simple search and replace. Any strings with any of the keywords in them will have those words replaced with their translated english keywords. Sorry! Someone could help me fix that. 