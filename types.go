package main

const singleQuote byte = '\''
const doubleQuote byte = '"'
const empty string = ""
const decodingStr string = "INFO: Decoding base64 string of: %s"

type doubleDecoder struct {
	cfg *config
}

type config struct {
	fileName  string
	minLength int
	asciiOnly bool
	buffSize  int
}
