package main

import (
	"fmt"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

//vertex position data
var tri_data []float32 = []float32{
	-1., -1., 0.,
	1., -1., 0.,
	1., 1., 0.,
	1., 1., 0.,
	-1., 1., 0.,
	-1., -1., 0.}

//texture data
var tex_data []float32 = []float32{
	0., 1.,
	1., 1.,
	1., 0.,
	1., 0.,
	0., 0.,
	0., 1.}

func main() {
	//create the game window
	window := CreateWindow(640, 480, "Game", false)

	//create VAO with the XYZ and ST data loaded in
	vao := gl.GenVertexArray()
	vao.Bind()
	_ = CreateVBOxyz(tri_data)
	_ = CreateVBOst(tex_data)
	vao.Unbind()

	//create program
	prog := LoadShaders("vert.glsl", "frag.glsl")

	//set up and bind texture
	tex := CreateTexturePNG("arch.png")
	BindToTexture(tex, 0, texSampler)
	ActiveTexture(prog, "myTextureSampler")

	//cull the backside of triangles
	gl.Enable(gl.CULL_FACE)

	//Draw loop
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		//draw data in VAO
		vao.Bind()
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
		vao.Unbind()

		//swap buffer and handle key events
		window.SwapBuffers()
		glfw.PollEvents()
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			break
		}
	}
}
