package glMenu

import (
	"github.com/Dumkin/glText"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
	"testing"
)

var window *glfw.Window

func openGLContext() {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		panic("glfw error")
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 5)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err = glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}
}

func TestTextBoxBackspace(t *testing.T) {
	openGLContext()

	tb := TextBox{}

	f := &glText.Font{}
	f.Config = &glText.FontConfig{}

	text := &glText.Text{}
	text.Font = f
	text.SetString("testing")
	tb.Text = text

	text = &glText.Text{}
	text.Font = f
	text.SetString("|")
	tb.Cursor = text

	tb.CursorIndex = 1
	tb.Backspace()
	if text.String != "esting" && tb.CursorIndex != 0 {
		t.Error(tb.Text.String, tb.CursorIndex)
	}
	tb.CursorIndex = 6
	tb.Backspace()
	if text.String != "estin" && tb.CursorIndex != 5 {
		t.Error(tb.Text.String, tb.CursorIndex)
	}
	tb.Backspace()
	if text.String != "esti" && tb.CursorIndex != 4 {
		t.Error(tb.Text.String, tb.CursorIndex)
	}
}
