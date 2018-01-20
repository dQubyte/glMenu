package glMenu

import (
	"github.com/go-gl/mathgl/mgl32"
)

type MenuImage struct {
	MenuTexture      *MenuTexture
	MenuTextureIndex int

	// final position on screen
	finalPosition mgl32.Vec2

	// text color
	color mgl32.Vec3

	// general opengl values
	vao           uint32
	vbo           uint32
	ebo           uint32
	vboData       []float32
	vboIndexCount int
	eboData       []int32
	eboIndexCount int

	// X1, X2: the lower left and upper right points of a box that bounds the text with a center point (0,0)
	// lower left
	X1 Point
	// upper right
	X2 Point

	// Screen position away from center
	Position mgl32.Vec2
}