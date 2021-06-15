# doubleb64decode

## About

**doubleb64decode** is a basic command line utility used to display greppable double base64 encoded strings from within a file.

The brute force results recursively returned from this program intentionally ommit quotes, double quotes, spaces, and newlines.

## Usage
    Usage of ./doubleb64decode:
    -asciiOnly
    	Only return ASII printable characters? (default true)
    -buffSize int
    	Size of buffer to use in bytes (default 134217728)
    -file string
    	Filename to open
    -minLength int
    	The minimum length of a potential base64 string to decode (default 6)

## Example
    ./doubleb64decode -filename=database.sql

## License
 [![License: WTFPL](https://img.shields.io/badge/License-WTFPL-brightgreen.svg)](http://www.wtfpl.net/about/) 

This program is free software. It comes without any warranty, to
the extent permitted by applicable law. You can redistribute it
and/or modify it under the terms of the Do What The Fuck You Want
To Public License, Version 2, as published by Sam Hocevar. See
http://www.wtfpl.net/ for more details.