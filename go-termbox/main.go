package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Press ESC to exit.")
	drawLine(2, 1, fmt.Sprintf("date: %v", time.Now()))
	drawLine(0, 2, "hogehogehoge")

	termbox.Flush()
}

func pollEvent() {
	draw()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			default:
				draw()
			}
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	pollEvent()
}
