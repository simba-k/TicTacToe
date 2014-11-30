package main

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/gl"
)

//Initialize glfw, automatically run on the import of the package
func init() {
	if !glfw.Init() {
		panic("Can't init glfw!")
	}
}

func OnResize(window *glfw.Window, width, height int) {
	//Gl Viewport is as large as screen
	gl.Viewport(0, 0, width, height)
}

func CreateWindow(w, h int, title string, rzable bool) *glfw.Window {
	//MultiSample
	glfw.WindowHint(glfw.Samples, 4)
	//Use GL3 Core, forward compatible hint for Mac
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	//Set Resizable
	if(rzable) {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	}else{
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}


	window, err := glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	//Catch Sticky Key input
	window.SetInputMode(glfw.StickyKeys, 1)
	//set onResize callback
	window.SetFramebufferSizeCallback(OnResize)

	//Init OpenGL in this context
	gl.Init()
	//if gl.Init throws an error that is irrelevant to the program working. Side effect of GLEW
	gl.GetError()
	return window
}
