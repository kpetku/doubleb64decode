package main

/* This program is free software. It comes without any warranty, to
 * the extent permitted by applicable law. You can redistribute it
 * and/or modify it under the terms of the Do What The Fuck You Want
 * To Public License, Version 2, as published by Sam Hocevar. See
 * http://www.wtfpl.net/ for more details. */

import (
	"log"
	"os"
	"time"
)

func main() {
	d := NewDoubleDecoder()
	log.Printf("initialized")
	err := d.Start()
	if err != nil {
		log.Fatalf("unable to open: %s", err.Error())
	}
	time.Sleep(time.Millisecond * 100)
	log.Printf("EOF")
	os.Exit(0)
}
