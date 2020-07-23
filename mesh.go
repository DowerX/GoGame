package main

import "github.com/go-gl/gl/v2.1/gl"

type mesh struct {
	Id  int
	VAO uint32
}

func (m *mesh) CreateVAO(vertices []float32, indicies []uint32) {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	var ebo uint32
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicies)*4, gl.Ptr(indicies), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(0)
	m.VAO = vao
}

func (m *mesh) Load(path string) {
	m.CreateVAO([]float32{0.0, 0.5, 0.0, 0.5, -0.5, 0.0, -0.5, -0.5, 0.0}, []uint32{0, 1, 2})
}
