package widget

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

// Box widget is a simple list where the child elements are arranged in a single column
// for vertical or a single row for horizontal arrangement
type Box struct {
	BaseWidget
	background color.Color

	Horizontal bool
	Children   []fyne.CanvasObject
}

// ApplyTheme updates this box to match the current theme
func (b *Box) ApplyTheme() {
	b.background = theme.BackgroundColor()
}

// Prepend inserts a new CanvasObject at the top/left of the box
func (b *Box) Prepend(object fyne.CanvasObject) {
	b.Children = append([]fyne.CanvasObject{object}, b.Children...)

	Refresh(b)
}

// Append adds a new CanvasObject to the end/right of the box
func (b *Box) Append(object fyne.CanvasObject) {
	b.Children = append(b.Children, object)

	Refresh(b)
}

// MinSize returns the size that this widget should not shrink below
func (b *Box) MinSize() fyne.Size {
	b.ExtendBaseWidget(b)
	return b.BaseWidget.MinSize()
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (b *Box) CreateRenderer() fyne.WidgetRenderer {
	b.ExtendBaseWidget(b)
	var lay fyne.Layout
	if b.Horizontal {
		lay = layout.NewHBoxLayout()
	} else {
		lay = layout.NewVBoxLayout()
	}

	return &boxRenderer{objects: b.Children, layout: lay, box: b}
}

func (b *Box) setBackgroundColor(bg color.Color) {
	b.background = bg
}

// NewHBox creates a new horizontally aligned box widget with the specified list of child objects
func NewHBox(children ...fyne.CanvasObject) *Box {
	box := &Box{BaseWidget: BaseWidget{}, Horizontal: true, Children: children}

	Renderer(box).Layout(box.MinSize())
	return box
}

// NewVBox creates a new vertically aligned box widget with the specified list of child objects
func NewVBox(children ...fyne.CanvasObject) *Box {
	box := &Box{BaseWidget: BaseWidget{}, Horizontal: false, Children: children}

	Renderer(box).Layout(box.MinSize())
	return box
}

type boxRenderer struct {
	layout fyne.Layout

	objects []fyne.CanvasObject
	box     *Box
}

func (b *boxRenderer) MinSize() fyne.Size {
	return b.layout.MinSize(b.objects)
}

func (b *boxRenderer) Layout(size fyne.Size) {
	b.layout.Layout(b.objects, size)
}

func (b *boxRenderer) ApplyTheme() {
}

func (b *boxRenderer) BackgroundColor() color.Color {
	if b.box.background == nil {
		return theme.BackgroundColor()
	}

	return b.box.background
}

func (b *boxRenderer) Objects() []fyne.CanvasObject {
	return b.objects
}

func (b *boxRenderer) Refresh() {
	b.objects = b.box.Children
	b.Layout(b.box.Size())

	canvas.Refresh(b.box)
}

func (b *boxRenderer) Destroy() {
}
