package animation

import (
	"Ascii-art-PRO/style"
	"fmt"
	"time"
)

func Loading(text string) {

	frame := []string{
		".       ",
		"..      ",
		"...     ",
		"....    ",
		".....   ",
		"......  ",
		"....... ",
		"........",
		"....... ",
		"......  ",
		".....   ",
		"....    ",
		"...     ",
		"..      ",
		".       ",
		"        ",
	}

	for i := 0; i <= 31; i++ {
		fmt.Printf("\r%s%s", style.Yellow(style.Bold(text)), frame[i%len(frame)])
		time.Sleep(40 * time.Millisecond)
	}
	fmt.Println()
}

func Loading1(text string) {

	frame := []string{
		".       ",
		"..      ",
		"...     ",
		"....    ",
		".....   ",
		"......  ",
		"....... ",
		"........",
		"....... ",
		"......  ",
		".....   ",
		"....    ",
		"...     ",
		"..      ",
		".       ",
		"        ",
	}

	for i := 0; i <= 31; i++ {

		fmt.Printf(style.Red(style.Bold("\r%s%s")), text, frame[i%len(frame)])
		time.Sleep(40 * time.Millisecond)
	}
	fmt.Println()
}

func Loading2(text string) {

	frame := []string{
		".       ",
		"..      ",
		"...     ",
		"....    ",
		".....   ",
		"......  ",
		"....... ",
		"........",
		"....... ",
		"......  ",
		".....   ",
		"....    ",
		"...     ",
		"..      ",
		".       ",
		"        ",
	}

	for i := 0; i <= 31; i++ {

		fmt.Printf(style.Green(style.Bold("\r%s%s")), text, frame[i%len(frame)])
		time.Sleep(40 * time.Millisecond)
	}
}
