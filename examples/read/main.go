package main

import (
	"fmt"

	"github.com/mishamyrt/go-keychron"
	"github.com/mishamyrt/go-keychron/pkg/color"
	"github.com/mishamyrt/go-keychron/pkg/hid"
	"github.com/mishamyrt/go-keychron/pkg/keyboard"
	"github.com/mishamyrt/go-keychron/pkg/mode"
	"github.com/mishamyrt/go-keychron/pkg/preset"
)

func formatColor(c color.RGB) string {
	if c.IsRandom() {
		return "random"
	}
	return fmt.Sprintf("rgb(%v, %v, %v)", c.R, c.G, c.B)
}

func printPreset(p *preset.Preset) {
	fmt.Printf("  Mode: %v\n", p.Mode().Name)
	fmt.Printf("  Color: %v\n", formatColor(p.Color()))
	fmt.Printf("  Speed: %v\n", p.Speed())
	fmt.Printf("  Brightness: %v\n", p.Brightness())
	f := p.Mode().Features
	if f.SupportsAny(mode.VerticalDirection, mode.HorizontalDirection) {
		fmt.Printf("  Direction: %v\n", p.Direction().String())
	}
}

func main() {
	hid.Init()
	b, err := keyboard.Open(keychron.K3v2Optical)
	if err != nil {
		panic(err)
	}
	b.SetDebug(true)
	current, all, err := b.GetPresets()
	if err != nil {
		panic("cannot read presets: " + err.Error())
	}
	// Empty line before output
	fmt.Println("")

	for i, p := range all {
		fmt.Printf("Preset %v:\n", i+1)
		printPreset(&p)
	}
	fmt.Println("Current:")
	printPreset(&current)
}
