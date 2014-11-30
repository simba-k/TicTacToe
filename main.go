package main

import (
	"fmt"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glh"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
)

func init() {
	if !glfw.Init() {
		panic("Can't init glfw!")
	}
}



//TODO
//refactor into window method later
func SetResizable(window *glfw.Window, rzable bool) {
}

func main() {
	fmt.Println(gl.TEXTURE0, gl.TEXTURE1, gl.TEXTURE2, gl.TEXTURE10, gl.TEXTURE20, gl.TEXTURE30)
	tri_data := []float32{-1., -1., 0.,
		1., -1., 0.,
		1., 1., 0.,
		1., 1., 0.,
		-1., 1., 0.,
		-1., -1., 0.}
	tex_data := []float32{0., 1.,
		1., 1.,
		1., 0.,
		1., 0.,
		0., 0.,
		0., 1.}
	window := CreateWindow(640, 480, "Game", false)
	SetResizable(window, false)
	window.SetFramebufferSizeCallback(OnResize)
	gl.Init()
	gl.GetError()
	vao := gl.GenVertexArray()
	vao.Bind()
	_ = CreateVBOxyz(tri_data)
	_ = CreateVBOst(tex_data)
	vao.Unbind()

	prog := LoadShaders("vert.glsl", "frag.glsl")
	texSampler := prog.GetUniformLocation("myTextureSampler")
	tex := CreateTexturePNG("arch.png")
	gl.Enable(gl.CULL_FACE)
	BindToTexture(tex, 0, texSampler)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		vao.Bind()
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
		vao.Unbind()
		window.SwapBuffers()
		glfw.PollEvents()
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			break
		}
	}
}

func OnResize(window *glfw.Window, width, height int) {
	gl.Viewport(0, 0, width, height)
}

func CreateWindow(w, h int, title string, rzable bool) *glfw.Window {
	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

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
	window.SetInputMode(glfw.StickyKeys, 1)
	return window
}

func BindToTexture(tex *glh.Texture, num int, texSampler gl.UniformLocation) {
	//TODO better error handleing
	if(num < 32) {
		return;
	}
	gl.ActiveTexture(gl.GLenum(gl.TEXTURE0 + num))
	tex.Bind(gl.TEXTURE_2D)
	texSampler.Uniform1i(num)
}

func CreateVBOxyz(data []float32) (vbo gl.Buffer){
	vbo = createVBO(data)
	attribLoc := gl.AttribLocation(0)
	attribLoc.EnableArray()
	attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)
	vbo.Unbind(gl.ARRAY_BUFFER)
	return vbo
}

func CreateVBOxy(data []float32) (vbo gl.Buffer){
	vbo = createVBO(data)
	attribLoc := gl.AttribLocation(0)
	attribLoc.EnableArray()
	attribLoc.AttribPointer(2, gl.FLOAT, false, 0, nil)
	vbo.Unbind(gl.ARRAY_BUFFER)
	return vbo
}

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

func LoadShaders(vShade, fShade string) gl.Program {
	vdat, _ := ioutil.ReadFile(vShade)
	var vShaderSrc string = string(vdat)
	fdat, _ := ioutil.ReadFile(fShade)
	var fShaderSrc string = string(fdat)
	vShader := gl.CreateShader(gl.VERTEX_SHADER)
	fShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	vShader.Source(vShaderSrc)
	vShader.Compile()
	fShader.Source(fShaderSrc)
	fShader.Compile()
	prog := gl.CreateProgram()
	prog.AttachShader(vShader)
	prog.AttachShader(fShader)
	prog.Link()
	prog.Use()
	return prog
}

func CheckError(tag string) {
	if er := gl.GetError(); er != gl.NO_ERROR {
		fmt.Printf("%s: %d\n", tag, er)
	}
}

func CreateTexturePNG(path string) *glh.Texture {
	img, _ := os.Open(path)
	im, _ := png.Decode(img)
	w := im.Bounds().Dx()
	h := im.Bounds().Dy()
	tex := &glh.Texture{gl.GenTexture(), w, h}
	tex.Bind(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.REPEAT)

	switch trueim := im.(type) {
	case *image.RGBA:
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
			trueim.Bounds().Dx(), trueim.Bounds().Dy(),
			0, gl.RGBA, gl.UNSIGNED_BYTE, trueim.Pix)

	default:
		copy := image.NewRGBA(trueim.Bounds())
		draw.Draw(copy, trueim.Bounds(), trueim, image.Pt(0, 0), draw.Src)
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
			copy.Bounds().Dx(), copy.Bounds().Dy(),
			0, gl.RGBA, gl.UNSIGNED_BYTE, copy.Pix)
	}

	gl.GenerateMipmap(gl.TEXTURE_2D)
	tex.Unbind(gl.TEXTURE_2D)
	return tex;
}
