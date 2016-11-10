//examples and parts of the code taken from: https://github.com/fogleman/ln

package main

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"math"
	"os"

	"github.com/fogleman/gg"
	"github.com/fogleman/ln/ln"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	_ "github.com/therecipe/qt/svg"
	"github.com/therecipe/qt/widgets"
)

var (
	sliderX *widgets.QSlider
	sliderY *widgets.QSlider
	sliderZ *widgets.QSlider

	svgCB       *widgets.QCheckBox
	exampleCB   *widgets.QCheckBox
	colorDialog *widgets.QColorDialog
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Renderer")
	window.SetMinimumSize2(512, 512)

	var mainLayout = widgets.NewQVBoxLayout()
	window.Layout().DestroyQObject()
	window.SetLayout(mainLayout)

	var scene = widgets.NewQGraphicsScene(nil)

	//sliders
	sliderX = widgets.NewQSlider2(core.Qt__Horizontal, nil)
	sliderX.SetRange(0, 360)
	sliderX.SetValue(1)
	mainLayout.AddWidget(sliderX, 0, 0)

	sliderY = widgets.NewQSlider2(core.Qt__Horizontal, nil)
	sliderY.SetRange(0, 360)
	sliderY.SetValue(2)
	mainLayout.AddWidget(sliderY, 0, 0)

	sliderZ = widgets.NewQSlider2(core.Qt__Horizontal, nil)
	sliderZ.SetRange(0, 360)
	sliderZ.SetValue(3)
	mainLayout.AddWidget(sliderZ, 0, 0)

	sliderX.ConnectValueChanged(func(value int) { renderWith(scene, value, sliderY.Value(), sliderZ.Value()) })
	sliderY.ConnectValueChanged(func(value int) { renderWith(scene, sliderX.Value(), value, sliderZ.Value()) })
	sliderZ.ConnectValueChanged(func(value int) { renderWith(scene, sliderX.Value(), sliderY.Value(), value) })

	var hLayout = widgets.NewQHBoxLayout()

	//svg
	svgCB = widgets.NewQCheckBox2("SVG", nil)
	svgCB.ConnectClicked(func(checked bool) { render(scene) })
	hLayout.AddWidget(svgCB, 0, 0)

	//example
	exampleCB = widgets.NewQCheckBox2("Example 2", nil)
	exampleCB.ConnectClicked(func(checked bool) { render(scene) })
	hLayout.AddWidget(exampleCB, 0, 0)

	//colorPB
	colorDialog = widgets.NewQColorDialog2(gui.NewQColor6("purple"), nil)
	colorDialog.ConnectColorSelected(func(colorPB *gui.QColor) {
		render(scene)
	})

	var colorPB = widgets.NewQPushButton2("Change Color", nil)
	colorPB.ConnectClicked(func(checked bool) {
		colorDialog.Show()
	})
	hLayout.AddWidget(colorPB, 0, 0)

	//savePB
	var savePB = widgets.NewQPushButton2("Save", nil)
	savePB.ConnectClicked(func(checked bool) {
		ioutil.WriteFile(widgets.QFileDialog_GetSaveFileName(window, "Save File", core.QDir_HomePath(), "Images (*.png *.svg)", "", 0), render(scene), 0644)
	})
	hLayout.AddWidget(savePB, 0, 0)
	mainLayout.AddLayout(hLayout, 0)

	var view = widgets.NewQGraphicsView(nil)
	view.SetScene(scene)
	mainLayout.AddWidget(view, 0, 0)

	render(scene)
	window.Show()
	widgets.QApplication_Exec()
}

func render(scene *widgets.QGraphicsScene) []byte {
	return renderWith(scene, sliderX.Value(), sliderY.Value(), sliderZ.Value())
}

func renderWith(scene *widgets.QGraphicsScene, x, y, z int) []byte {
	scene.Clear()

	var data []byte
	if exampleCB.IsChecked() {
		data = example2(x, y, z)
	} else {
		data = example1(x, y, z)
	}

	var pix = gui.QPixmap_FromImage(gui.QImage_FromData2(core.NewQByteArray2(string(data), len(data)), func() string {
		if svgCB.IsChecked() {
			return "SVG"
		}
		return "PNG"
	}()), 0)

	scene.AddPixmap(pix)
	scene.SetSceneRect(core.NewQRectF5(pix.Rect()))

	return data
}

func pathsToBytes(paths ln.Paths, width, height float64) []byte {
	if svgCB.IsChecked() {
		return []byte(paths.ToSVG(width, height))
	}

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	// renderWith the paths in an image
	scale := 1.0
	w, h := int(width*scale), int(height*scale)
	dc := gg.NewContext(w, h)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(colorDialog.CurrentColor().RedF(), colorDialog.CurrentColor().GreenF(), colorDialog.CurrentColor().BlueF())
	dc.SetLineWidth(3)
	for _, path := range paths {
		for _, v := range path {
			dc.LineTo(v.X*scale, v.Y*scale)
		}
		dc.NewSubPath()
	}
	dc.Stroke()

	png.Encode(bb, dc.Image())
	return bb.Bytes()
}

func example1(frameX, frameY, frameZ int) []byte {
	// create a scene and add a single cube

	scene := ln.Scene{}
	scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))

	cx := math.Cos(ln.Radians(float64(frameX)))
	cy := math.Sin(ln.Radians(float64(frameY)))
	cz := math.Cos(ln.Radians(float64(frameZ)))

	// define camera parameters
	eye := ln.Vector{cx, cy, cz}.MulScalar(5) // camera position
	center := ln.Vector{0, 0, 0}              // camera looks at
	up := ln.Vector{0, 0, 1}                  // up direction

	// define rendering parameters
	width := 384.0  // rendered width
	height := 384.0 // rendered height
	fovy := 50.0    // vertical field of view, degrees
	znear := 0.1    // near z plane
	zfar := 10.0    // far z plane
	step := 0.01    // how finely to chop the paths for visibility testing

	// compute 2D paths that depict the 3D scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	return pathsToBytes(paths, width, height)
}

func example2(frameX, frameY, frameZ int) []byte {
	cx := math.Cos(ln.Radians(float64(frameX)))
	cy := math.Sin(ln.Radians(float64(frameY)))
	cz := math.Cos(ln.Radians(float64(frameZ)))
	scene := ln.Scene{}
	eye := ln.Vector{cx, cy, cz}.MulScalar(10)
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}

	nodes := []ln.Vector{
		{1.047, -0.000, -1.312},
		{-0.208, -0.000, -1.790},
		{2.176, 0.000, -2.246},
		{1.285, -0.001, 0.016},
		{-1.276, -0.000, -0.971},
		{-0.384, 0.000, -2.993},
		{-2.629, -0.000, -1.533},
		{-1.098, -0.000, 0.402},
		{0.193, 0.005, 0.911},
		{-1.934, -0.000, 1.444},
		{2.428, -0.000, 0.437},
		{0.068, -0.000, 2.286},
		{-1.251, -0.000, 2.560},
		{1.161, -0.000, 3.261},
		{1.800, 0.001, -3.269},
		{2.783, 0.890, -2.082},
		{2.783, -0.889, -2.083},
		{-2.570, -0.000, -2.622},
		{-3.162, -0.890, -1.198},
		{-3.162, 0.889, -1.198},
		{-1.679, 0.000, 3.552},
		{1.432, -1.028, 3.503},
		{2.024, 0.513, 2.839},
		{0.839, 0.513, 4.167},
		// {0.000000, 0.000000, 0.000000},
		// {0.000000, 0.000000, 1.089000},
		// {1.026719, 0.000000, -0.363000},
		// {-0.513360, -0.889165, -0.363000},
		// {-0.513360, 0.889165, -0.363000},
		//
		// {0, 0, 0},
		// {-1, 0, 0},
		// {1, 0, 0},
		// {0, 1, 0},
		// {0, -1, 0},
		// {0, 0, 1},
		// {0, 0, -1},
		//
		// {-1, 1, 1},
		// {-1, 1, -1},
		// {-1, -1, 1},
		// {-1, -1, -1},
		// {1, 1, 1},
		// {1, 1, -1},
		// {1, -1, 1},
		// {1, -1, -1},
	}

	edges := [][2]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 4},
		{1, 5},
		{2, 14},
		{2, 15},
		{2, 16},
		{3, 8},
		{3, 10},
		{4, 6},
		{4, 7},
		{6, 17},
		{6, 18},
		{6, 19},
		{7, 8},
		{7, 9},
		{8, 11},
		{9, 12},
		{11, 12},
		{11, 13},
		{12, 20},
		{13, 21},
		{13, 22},
		{13, 23},
	}

	for _, v := range nodes {
		scene.Add(ln.NewOutlineSphere(eye, up, v, 0.333))
	}

	// for _, v0 := range nodes {
	// 	for _, v1 := range nodes {
	// 		if v0 == v1 {
	// 			continue
	// 		}
	for _, edge := range edges {
		v0 := nodes[edge[0]]
		v1 := nodes[edge[1]]
		d := v1.Sub(v0)
		z := d.Length()
		u := d.Cross(up).Normalize()
		a := math.Acos(d.Normalize().Dot(up))
		m := ln.Translate(v0)
		if a != 0 {
			m = ln.Rotate(u, a).Translate(v0)
		}
		c := ln.NewOutlineCylinder(m.Inverse().MulPosition(eye), up, 0.1, 0, z)
		scene.Add(ln.NewTransformedShape(c, m))
	}
	// }

	width := 384.0
	height := 384.0
	paths := scene.Render(eye, center, up, width, height, 60, 0.1, 100, 0.01)

	return pathsToBytes(paths, width, height)
}
