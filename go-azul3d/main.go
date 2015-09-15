package main

import (
	// 	"fmt"
	// 	"image"
	// 	"image/png"
	// 	"log"
	// 	"os"

	"azul3d.org/gfx.v1"
	"azul3d.org/gfx/window.v2"
	// 	"azul3d.org/keyboard.v1"
	math "azul3d.org/lmath.v1"
)

var glslVert = []byte(`
#version 120

attribute vec3 Vertex;
attribute vec4 Color;

uniform mat4 MVP;

verying vec4 frontColor;

void main()
{
	frontColor = Color;
	gl_Position = MVP * vec4(Vertex, 1.0);
}
`)

var glslFrag = []byte(`
#version 120

verying vec4 frontColor;

void main()
{
	gl_FragColor = frontColor;
}
`)

func gfxLoop(w window.Window, r gfx.Renderer) {
	camera := gfx.NewCamera()
	camFOV := 75.0
	camNear := 0.0001
	camFar := 1000.0
	camera.SetPersp(r.Bounds(), camFOV, camNear, camFar)

	camera.SetPos(math.Vec3{0, -2, 0})

	shader := gfx.NewShader("SimpleShader")
	shader.GLSLVert = glslVert
	shader.GLSLFrag = glslFrag

	triMesh := gfx.NewMesh()
	triMesh.Vertices = []gfx.Vec3{
		{0, 0, 1},
		{-.5, 0, 0},
		{.5, 0, 0},
		{-.5, 0, 0},
		{-1, 0, -1},
		{0, 0, -1},
		{.5, 0, 0},
		{0, 0, -1},
		{1, 0, -1},
	}

	triMesh.Colors = []gfx.Color{
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 1, 1},
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 1, 1},
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 1, 1},
	}

	triangle := gfx.NewObject()
	triangle.Shader = shader
	triangle.OcclusionTest = true
	triangle.State.FaceCulling = gfx.NoFaceCulling
	triangle.Meshes = []*gfx.Mesh{triMesh}

	right := gfx.NewTransform()
	right.SetRot(math.Vec3{0, 0, -45})

	left := gfx.NewTransform()
	left.SetRot(math.Vec3{0, 0, 45})
	left.SetParent(right)

	for {
		bounds := r.Bounds()
		r.Draw(bounds.Inset(50), triangle, camera)
		r.Render()
	}
}

func main() {
	window.Run(gfxLoop, nil)
}
