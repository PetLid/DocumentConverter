# Document converter

Simple document converter that breaks down a file line by line into elements and converts into a different format (XML).
Element definitions must be supplied via a file. Comes with example input file and definition file.

## Usage
[Install go](https://go.dev/dl/), run `go run .` in the root. Use flag "--help" for options. 

## How to add additional output formats
If you want to convert to another format, you can a new printer implementation, and allow users to add a flag for output format,
then create a factory to supply the converter with a printer matching the flag.
