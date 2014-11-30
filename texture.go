package main

import (
	"fmt"
	"github.com/go-gl/glh"
	"image"
	"image/draw"
	"image/png"
	"os"
	"github.com/go-gl/gl"
)

//Binds texture to GL_TEXTUREx where x is the parameter num
//sets the sampler uniform to the TEXTURE value give
func BindToTexture(tex *glh.Texture, num int, texSampler gl.UniformLocation) {
	if(num >= 32 || num <0) {
		fmt.Printf("ERROR at bind tex out of bounds")
		return;
	}
	gl.ActiveTexture(gl.GLenum(gl.TEXTURE0 + num))
	tex.Bind(gl.TEXTURE_2D)
}

//TODO
//Refactor into method of program
//Activates the texture in the program
func ActivateTexture(prog gl.Program, smpler2Dname string) {
	texSampler := prog.GetUniformLocation(smpler2Dname)
	texSampler.Uniform1i(num)
}

//Turns a PNG image into a GL Texture
//Takes the path to the PNG image as a parameter
func CreateTexturePNG(path string) *glh.Texture {
	//Open and decode image
	img, _ := os.Open(path)
	im, _ := png.Decode(img)

	//get width and height
	w := im.Bounds().Dx()
	h := im.Bounds().Dy()

	tex := &glh.Texture{gl.GenTexture(), w, h}

	tex.Bind(gl.TEXTURE_2D)
	//set paramters for OGL
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.REPEAT)

	//Check if image is in RGBA if it is load as a GL Texture
	//otherwise turn it into the RGBA type then load it as a
	//Texture
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

