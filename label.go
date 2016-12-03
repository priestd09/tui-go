package tui

import (
	"image"
	"strings"

	wordwrap "github.com/mitchellh/go-wordwrap"
)

var _ Widget = &Label{}

// Label is a widget to display read-only text.
type Label struct {
	text string

	wordWrap bool

	size image.Point

	sizePolicyX SizePolicy
	sizePolicyY SizePolicy
}

// NewLabel returns a new Label.
func NewLabel(text string) *Label {
	return &Label{
		text: text,
		size: image.Point{-1, -1},
	}
}

// Draw draws the label.
func (l *Label) Draw(p *Painter) {
	lines := strings.Split(l.text, "\n")

	if l.wordWrap {
		lines = strings.Split(wordwrap.WrapString(l.text, uint(l.Size().X)), "\n")
	}

	for i, line := range lines {
		p.DrawText(0, i, line)
	}
}

// Size returns the size of the label.
func (l *Label) Size() image.Point {
	return l.size
}

func (l *Label) MinSize() image.Point {
	return image.Point{}
}

// SizeHint returns the recommended size for the label.
func (l *Label) SizeHint() image.Point {
	if l.size.X < 0 && l.size.Y < 0 {
		var max int
		lines := strings.Split(l.text, "\n")
		for _, line := range lines {
			if w := stringWidth(line); w > max {
				max = w
			}
		}
		return image.Point{max, l.heightForWidth(max)}
	}
	return image.Point{l.size.X, l.heightForWidth(l.size.X)}
}

// SizePolicy returns the default layout behavior.
func (l *Label) SizePolicy() (SizePolicy, SizePolicy) {
	return l.sizePolicyX, l.sizePolicyY
}

func (l *Label) heightForWidth(w int) int {
	return len(strings.Split(wordwrap.WrapString(l.text, uint(w)), "\n"))
}

// Resize updates the size of the label.
func (l *Label) Resize(size image.Point) {
	hpol, vpol := l.SizePolicy()

	switch hpol {
	case Minimum:
		l.size.X = l.SizeHint().X
	case Expanding:
		l.size.X = size.X
	}

	switch vpol {
	case Minimum:
		l.size.Y = l.SizeHint().Y
	case Expanding:
		l.size.Y = size.Y
	}
}

func (l *Label) OnEvent(_ Event) {
}

func (l *Label) SetText(text string) {
	l.text = text
}

func (l *Label) SetWordWrap(enabled bool) {
	l.wordWrap = enabled
}

// SetSizePolicy sets the size policy for each axis.
func (l *Label) SetSizePolicy(horizontal, vertical SizePolicy) {
	l.sizePolicyX = horizontal
	l.sizePolicyY = vertical
}
