## Subtitles CLI

CLI tool to manipulate subtitle files.

## Installation

A pre-built Windows 64-bit executable can be found under releases.

You can also build an executable directly by installing Go 1.17.2+, cloning the
repo and running `go build`. Additional documentation can be found
[here](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies).

## Feature

- [x] Add/subtract offset time to all subtitle timestamps
- [x] Support `.srt` files
- [ ] Export/import from JSON format
- [ ] Support `.vtt` files

## Usage

```sh
# Add offset of 2 minute 3 seconds to all timestamps
subtitles offset sample.srt --duration 2m3s --output new.srt

# Subtract offset of 400 milliseconds to all timestamps
subtitles offset sample.srt --duration -400ms --output new.srt
```
