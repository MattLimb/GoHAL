# GoHAL

`GoHAL` is an interpreter for the Esoteric Programming Language [`2001: An Esolang Odyssey`](https://esolangs.org/wiki/2001:_An_Esolang_Odyssey) and its Parent [`brainfuck`](https://esolangs.org/wiki/Brainfuck).

This is a toy project and was built for fun. If you run this and find bugs, please drop an issue, although be aware I might not be able to fix them in a timely mannar.

If you have any feedback or improvements, feel free to write up an issue for them or submit a PR. I'm learning here, so any thoughts are appreicated.


> [!CAUTION]
> This project provides NO WARRANTY for usage. USE AT YOUR OWN RISK.

> [!NOTE]
> This programming Language was written by [PythonshellDebugwindow](https://esolangs.org/wiki/User:PythonshellDebugwindow) on [EsoLangs.org](https://esolangs.org)


## Running the Project

This project is written in Google's GoLang. To build this project you'll need to install Go.

1. Download and Install GoLang for your operating system: [https://go.dev/doc/install](https://go.dev/doc/install)
2. From the current working directory: `go run main.go ./examples/lang_2001/hello_world.go`

## Usage

```sh
> go run main.go -h
Usage: gohal [flags] [filename]

Arguments:
  filename  The script file you want HAL to execute.

Flags:
  -d    Enable debug mode.
  -debug
        Enable debug mode.
  -l string
        specify which language to try to parse and run. (default "2001")
  -language string
        specify which language to try to parse and run. (default "2001")
  -v    display the current version and exit.
  -version
        display the current version and exit.
```

### Currently Available Languages

| Language String | Language | Examples Directory |
| :-------------: | :------: | :----------------: |
| `2001`    | 2001: An Esolang Odyssey | `examples/2001` |
| `brainfuck`    | brainfuck | `examples/brianfuck` |


## Running the Tests

1. Ensure GoLang is installed for your operating system: [https://go.dev/doc/install](https://go.dev/doc/install)
2. `go test ./internal ./languages/lang_2001 ./languages/lang_brainfuck`


# Credits

- [PythonshellDebugwindow](https://esolangs.org/wiki/User:PythonshellDebugwindow) for creating `2001: An Esolang Odyssey` and implementing the "Hello World!" example.
- [Urban Müller](https://esolangs.org/wiki/Urban_M%C3%BCller) for creating `brainfuck`.