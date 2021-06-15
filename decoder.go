package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func NewDoubleDecoder() *doubleDecoder {
	return new(doubleDecoder)
}

func (d *doubleDecoder) Start() error {
	fileName := flag.String("file", "", "Filename to open")
	minLength := flag.Int("minLength", 6, "The minimum length of a potential base64 string to decode")
	asciiOnly := flag.Bool("asciiOnly", true, "Only return ASII printable characters?")
	buffSize := flag.Int("buffSize", 128*1024*1024, "Size of buffer to use in bytes")
	flag.Parse()
	if *fileName == "" {
		log.Printf("Please specify a file to open with -file")
		flag.Usage()
		os.Exit(1)
	}
	d.cfg = &config{
		fileName:  *fileName,
		minLength: *minLength,
		asciiOnly: *asciiOnly,
		buffSize:  *buffSize,
	}
	f, err := os.Open(d.cfg.fileName)
	if err != nil {
		return err
	}
	d.Scan(bufio.NewReader(f))
	return nil
}

func (d *doubleDecoder) Scan(r io.Reader) error {
	var outerBuff, innerBuff []byte
	outerScanner := bufio.NewScanner(r)
	outerScanner.Buffer(outerBuff, d.cfg.buffSize)
	outerScanner.Split(bufio.ScanWords)
	for outerScanner.Scan() {
		innerScanner := bufio.NewScanner(bytes.NewReader(outerScanner.Bytes()))
		innerScanner.Buffer(innerBuff, d.cfg.buffSize)
		innerScanner.Split(splitFunc)
		for innerScanner.Scan() {
			if innerScanner.Text() != empty && len(innerScanner.Text()) >= d.cfg.minLength {
				d.decodeString(innerScanner.Text())
			}
		}
	}
	return nil
}

func (d *doubleDecoder) decodeString(s string) error {
	inner, err := base64.StdEncoding.DecodeString(s)
	if err == nil {
		if len(inner) > d.cfg.minLength {
			if isASCII(string(inner)) && d.cfg.asciiOnly {
				log.Printf(decodingStr, s)
				fmt.Println(string(inner))
				d.Scan(bytes.NewReader(inner))
			}
		}
	}
	return err
}
