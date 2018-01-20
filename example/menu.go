package main

import (
	"fmt"
	"github.com/Dumkin/glMenu"
	"github.com/Dumkin/glText"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"os"
)

func MenuInit(window *glfw.Window, font *glText.Font) {
	menuManager = glMenu.NewMenuManager(font, glfw.KeyM, "main")

	defaults := glMenu.MenuDefaults{
		TextColor:       mgl32.Vec3{1, 1, 1},
		TextClick:       mgl32.Vec3{250.0 / 255.0, 0, 154.0 / 255.0},
		TextHover:       mgl32.Vec3{0, 250.0 / 255.0, 154.0 / 255.0},
		BackgroundColor: mgl32.Vec4{0.5, 0.5, 0.5, 1.0},
		BorderColor:     mgl32.Vec4{1, 1, 1, 1.0},
		Border:          mgl32.Vec2{2, 2},
		Dimensions:      mgl32.Vec2{0, 0},
		Padding:         mgl32.Vec2{10, 10},
		HoverPadding:    mgl32.Vec2{10, 10},
	}

	// menu 1
	mainMenu, err := menuManager.NewMenu(window, "main", defaults, glMenu.ScreenLeft)
	if err != nil {
		fmt.Println("error loading the font")
		os.Exit(1)
	}
	textbox := mainMenu.NewTextBox("127.0.0.1", 250, 40, 1)
	textbox.Text.MaxRuneCount = 16
	mainMenu.NewLabel("Options", glMenu.LabelConfig{Action: glMenu.GOTO_MENU, Goto: "option"})
	mainMenu.NewLabel("Disabled", glMenu.LabelConfig{Action: glMenu.NOOP})
	mainMenu.NewLabel("Quit", glMenu.LabelConfig{Action: glMenu.EXIT_GAME})

	// menu 2
	defaults = glMenu.MenuDefaults{
		BackgroundColor: mgl32.Vec4{0, 1, 1, 1},
		Dimensions:      mgl32.Vec2{200, 200},
		Padding:         mgl32.Vec2{10, 10},
		HoverPadding:    mgl32.Vec2{10, 10},
	}
	optionMenu, err := menuManager.NewMenu(window, "option", defaults, glMenu.ScreenTopCenter)
	if err != nil {
		fmt.Println("error loading font")
		os.Exit(1)
	}
	optionMenu.NewLabel("Back", glMenu.LabelConfig{Action: glMenu.GOTO_MENU, Goto: "main"})

	// complete setup
	menuManager.Finalize(glMenu.AlignRight)
}