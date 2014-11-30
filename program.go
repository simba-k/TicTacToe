package main

import(
	"io/ioutil"
	"github.com/go-gl/gl"
)

//TODO
//Better error handling

//Create program from shader file paths
//Vertex Shader path first then Fragment Shader
func LoadShaders(vShade, fShade string) gl.Program {
	//create and compile shaders
	vShader := CreateShader(vShade, gl.VERTEX_SHADER)
	fShader := CreateShader(fShade, gl.FRAGMENT_SHADER)

	//attach them to a new program
	prog := gl.CreateProgram()
	prog.AttachShader(vShader)
	prog.AttachShader(fShader)
	prog.Link()

	//enable program
	prog.Use()
	return prog
}

//for internal use, gets the source from the file specified in the path
//then compiles
func createShader(path string, shtype gl.GLenum) gl.Shader {
	rawdata, _ := ioutil.ReadFile(path)
	src := string(rawdata)
	shader := gl.CreateShader(shtype)
	shader.Source(src)
	shader.Compile()
	return shader
}
