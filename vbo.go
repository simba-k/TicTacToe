package main

import (
	"github.com/go-gl/gl"
)

//Prereq: VAO is binded
//Create VBO from a float32 slice
//that is in form x1,y1,z1,x2,y2,z2
//and sets the attribute pointer
func CreateVBOxyz(data []float32) (vbo gl.Buffer){
	vbo = createVBO(data)
	attribLoc := gl.AttribLocation(0)
	attribLoc.EnableArray()
	attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)
	vbo.Unbind(gl.ARRAY_BUFFER)
	return vbo
}

//Prereq: VAO is binded
//Create VBO from a float32 slice
//that is in form x1,y1,x2,y2
//and sets the attribute pointer
func CreateVBOxy(data []float32) (vbo gl.Buffer){
	vbo = createVBO(data)
	attribLoc := gl.AttribLocation(0)
	attribLoc.EnableArray()
	attribLoc.AttribPointer(2, gl.FLOAT, false, 0, nil)
	vbo.Unbind(gl.ARRAY_BUFFER)
	return vbo
}

//Prereq: VAO is binded
//Create VBO from a float32 slice
//that is in form s1,t1,s2,t2
//and sets the attribute pointer
func CreateVBOst(data []float32)  (vbo gl.Buffer) {
	vbo = createVBO(data)
	attribLoc := gl.AttribLocation(1)
	attribLoc.EnableArray()
	attribLoc.AttribPointer(2, gl.FLOAT, false, 0, nil)
	vbo.Unbind(gl.ARRAY_BUFFER)
	return vbo
}

func createVBO(data []float32) (vbo gl.Buffer) {
	vbo = gl.GenBuffer()
	vbo.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, data, gl.DYNAMIC_DRAW)
	return vbo
}
