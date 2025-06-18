# hackabin
a simple CLI app that lets you save code snippets for you to copy and paste wherever needed.

## commands
- `hackabin add [title:string] [code:string|file]` - add a snippet
- `hackabin remove [id]` - remove a snippet
- `hackabin list` - view all snippets
- `hackabin view [id]` - view and copy snippet to clipboard
- `hackabin help` - [stop it get some help](https://imgur.com/gallery/stop-get-some-help-DJIkoRf)

## installation
just download the binary for your OS from the [`build/`](https://github.com/ronykax/hackabin/tree/master/build) folder!

## building
follow the steps below if you prefer building the binary for your OS yourself:
1. download and install a non-ancient version [golang](https://go.dev/doc/install) on your computer
2. clone this repo and go to that directory
3. run `go build -o hackabin` and you will see a `hackabin` or `hackabin.exe` file appear in the current directory!