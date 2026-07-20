package main

import "log"

func exit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
