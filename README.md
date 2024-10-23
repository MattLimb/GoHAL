# GoHAL

`GoHAL` is a new interpreter for tape driven (Esoteric)[https://esolangs.org] programming languages.

Currently, `GoHAL` supports: 

- [`2001: An Esolang Odyssey`](https://esolangs.org/wiki/2001:_An_Esolang_Odyssey) - An Esoteric programming language based off of the book and film `2001: A Space Odyssey`. The programming language was created by [PythonshellDebugwindow](https://esolangs.org/wiki/User:PythonshellDebugwindow). `GoHAL` was originally written for this language.
- [`brainfuck`](https://esolangs.org/wiki/Brainfuck) - A well known and influencial esoteric programming language created by [Urban Müller](https://esolangs.org/wiki/Urban_M%C3%BCller)
- [`morsefuck`](https://esolangs.org/wiki/Morsefuck) - An Esoteric programming language based off of the Morse Code Encoding Scheme, created by [Thoga31](https://esolangs.org/wiki/User:Thoga31)

This is a toy project and was built for fun. If you run this and find bugs, please drop an issue, although be aware I might not be able to fix them in a timely mannar.

If you have any feedback or improvements or want to contribute, feel free to write up an issue for them or submit a PR. I'm learning here, so any thoughts are appreicated.


> [!CAUTION]
> This project provides NO WARRANTY for usage. USE AT YOUR OWN RISK.

## Running the Project

This project is written in Google's GoLang. To build this project you'll need to install Go.

1. Download and Install GoLang for your operating system: [https://go.dev/doc/install](https://go.dev/doc/install)
2. From the current working directory: `go run main.go run ./examples/lang_2001/hello_world.hal`

## Usage

```
$ go run main.go --help
Usage:
  gohal -v
  gohal [--debug,--language] run [inputFilename]
  gohal [--debug,--language,--outputLanguage] transpile [inputFilename] [outputFilename]

Arguments:
  filename  The script file you want HAL to execute.

Flags:
  -d    Global flag to enable debug mode.
  -debug
        Global flag to enable debug mode.
  -l string
        specify which language to parse the input with. (default "2001")
  -language string
        specify which language to parse the input with. (default "2001")
  -o string
        specify which language to try to parse and run. (default "brainfuck")
  -outputLanguage string
        specify which language to try to parse and run. (default "brainfuck")
  -v    display the current version and exit.
  -version
        display the current version and exit. 
```

### Available Esoteric Languages

This table is to be used with the flags: `-l/--language` and `-o/--outputLanguage`

| Language String | Language | Examples Directory |
| :-------------: | :------: | :----------------: |
| `2001`    | 2001: An Esolang Odyssey | `examples/2001` |
| `brainfuck`    | brainfuck | `examples/brianfuck` |
| `morsefuck` | morsefuck | `examples/morsefuck` |


## Running the Tests

1. Ensure GoLang is installed for your operating system: [https://go.dev/doc/install](https://go.dev/doc/install)
2. `go test ./internal ./languages/lang_2001 ./languages/lang_brainfuck`


# Credits

- [PythonshellDebugwindow](https://esolangs.org/wiki/User:PythonshellDebugwindow) for creating `2001: An Esolang Odyssey` and implementing the "Hello World!" example.
- [Urban Müller](https://esolangs.org/wiki/Urban_M%C3%BCller) for creating `brainfuck`.
- [Thoga31](https://esolangs.org/wiki/User:Thoga31) for creating Morsefuck.