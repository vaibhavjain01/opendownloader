package main

import (
	"log"
)

type Logger struct {
	ToFile	bool
	FilePath	string
}

func (logger Logger) LogToConsole(str string) bool {
	log.Println(str)
	return true
}