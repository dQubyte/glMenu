package main

import (
	"fmt"
	"github.com/Dumkin/glMenu"
	"github.com/Dumkin/glText"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"golang.org/x/image/math/fixed"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

var window *glfw.Window
var menuManager *glMenu.MenuManager

func main() {
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
	window.SetKeyCallback(keyCallback)
	window.SetMouseButtonCallback(mouseButtonCallback)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	fmt.Println("Opengl version", gl.GoStr(gl.GetString(gl.VERSION)))

	font, err := glText.LoadTruetype("fonts", "Roboto")
	defer font.Release()

	if err == nil {
		fmt.Println("Font loaded from disk...")
	} else {
		fd, err := os.Open("fonts/Roboto.ttf")
		if err != nil {
			panic(err)
		}
		defer fd.Close()

		scale := fixed.Int26_6(25)
		runesPerRow := fixed.Int26_6(128)

		runeRanges := make(glText.RuneRanges, 0)

		runeRange := glText.RuneRange{Low: 0x0001, High: 0x0570}
		runeRanges = append(runeRanges, runeRange)

		font, err = glText.NewTruetype(fd, scale, runeRanges, runesPerRow)
		if err != nil {
			panic(err)
		}

		font.Config.Name = "Roboto"

		err = font.Config.Save("fonts")
		if err != nil {
			panic(err)
		}
	}

	// load menus
	MenuInit(window, font)

	menuManager.Show("main")
	defer menuManager.Release()

	gl.ClearColor(0, 0, 0, 0.0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		xPos, yPos := window.GetCursorPos()
		menuManager.MouseHover(xPos, yPos)
		menuManager.Draw()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func keyCallback(
	w *glfw.Window,
	key glfw.Key,
	scancode int,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	if action != glfw.Release {
		if mods == glfw.ModShift {
			menuManager.KeyRelease(key, true)
		} else {
			menuManager.KeyRelease(key, false)
		}
	}
}

func mouseButtonCallback(
	w *glfw.Window,
	button glfw.MouseButton,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	xPos, yPos := w.GetCursorPos()
	if button == glfw.MouseButtonLeft && action == glfw.Press {
		menuManager.MouseClick(xPos, yPos, glMenu.MouseLeft)
	}
	if button == glfw.MouseButtonLeft && action == glfw.Release {
		menuManager.MouseRelease(xPos, yPos, glMenu.MouseLeft)
	}
}