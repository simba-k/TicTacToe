package main

import (
	"fmt"
	"github.com/go-gl/gl"
)

//Check if there are any OpenGL errors enqueued
func CheckError(tag string) {
	if er := gl.GetError(); er != gl.NO_ERROR {
		fmt.Printf("%s: %d\n", tag, er)
	}
}
