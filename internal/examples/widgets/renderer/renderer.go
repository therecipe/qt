//examples and parts of the code taken from: https://github.com/fogleman/ln and https://github.com/fogleman/pt
//3D Gopher Model by Takuya Ueda: http://u.hinoichi.net

package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"
	"github.com/fogleman/ln/ln"
	"github.com/fogleman/pt/pt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	_ "github.com/therecipe/qt/svg"
	"github.com/therecipe/qt/widgets"
)

var (
	sliderX *widgets.QSlider
	sliderY *widgets.QSlider
	sliderZ *widgets.QSlider

	svgCB     *widgets.QCheckBox
	exampleCB *widgets.QComboBox
	timedCB   *widgets.QCheckBox

	colorPB     *widgets.QPushButton
	colorDialog *widgets.QColorDialog
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Renderer")
	window.SetMinimumSize2(512, 512)

	var mainLayout = widgets.NewQVBoxLayout()

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(mainLayout)
	window.SetCentralWidget(centralWidget)

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

	//timed
	timedCB = widgets.NewQCheckBox2("Timed", nil)
	timedCB.SetDisabled(true)
	timedCB.ConnectClicked(func(checked bool) { render(scene) })
	hLayout.AddWidget(timedCB, 0, 0)

	//example
	exampleCB = widgets.NewQComboBox(nil)
	exampleCB.AddItems([]string{"Example 1 [LN]", "Example 2 [LN]", "Example 3 [PT]", "Example 4 [PT]", "Example 5 [PT]"})
	exampleCB.SetCurrentIndex(0)
	exampleCB.ConnectCurrentIndexChanged(func(index int) {
		switch index {
		case 0, 1:
			{
				svgCB.SetEnabled(true)
				timedCB.SetDisabled(true)
				colorPB.SetEnabled(true)
			}

		default:
			{
				svgCB.SetDisabled(true)
				svgCB.SetChecked(false)
				timedCB.SetEnabled(true)
				colorPB.SetDisabled(true)
			}
		}

		timedCB.SetChecked(false)
		render(scene)
	})
	hLayout.AddWidget(exampleCB, 0, 0)

	//colorPB
	colorDialog = widgets.NewQColorDialog2(gui.NewQColor6("purple"), nil)
	colorDialog.ConnectColorSelected(func(colorPB *gui.QColor) { render(scene) })

	colorPB = widgets.NewQPushButton2("Change Color", nil)
	colorPB.ConnectClicked(func(checked bool) { colorDialog.Show() })
	hLayout.AddWidget(colorPB, 0, 0)

	//savePB
	var savePB = widgets.NewQPushButton2("Save", nil)
	savePB.ConnectClicked(func(checked bool) {
		var path = widgets.QFileDialog_GetSaveFileName(window, "Save File", core.QDir_HomePath(), "Images (*.png *.svg)", "", 0)
		if path != "" {
			ioutil.WriteFile(path, render(scene), 0644)
		}
	})
	hLayout.AddWidget(savePB, 0, 0)
	mainLayout.AddLayout(hLayout, 0)

	var view = widgets.NewQGraphicsView(nil)
	view.SetRenderHints(gui.QPainter__Antialiasing | gui.QPainter__SmoothPixmapTransform)
	view.SetScene(scene)
	mainLayout.AddWidget(view, 0, 0)

	render(scene)
	window.Show()
	widgets.QApplication_Exec()
}

func render(scene *widgets.QGraphicsScene) []byte {
	return renderWith(scene, sliderX.Value(), sliderY.Value(), sliderZ.Value())
}

func renderWith(scene *widgets.QGraphicsScene, frameX, frameY, frameZ int) []byte {
	scene.Clear()

	var data []byte
	switch exampleCB.CurrentIndex() {
	case 0:
		{
			data = example1(frameX, frameY, frameZ)
		}

	case 1:
		{
			data = example2(frameX, frameY, frameZ)
		}

	case 2:
		{
			data = example3(frameX, frameY, frameZ, timedCB.IsChecked())
		}

	case 3:
		{
			data = example4(frameX, frameY, frameZ, timedCB.IsChecked())
		}

	case 4:
		{
			data = example5(frameX, frameY, frameZ, timedCB.IsChecked())
		}
	}

	var pix = gui.NewQPixmap()
	pix.LoadFromData(data, uint(len(data)), "", 0)

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

//example 1

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

//example 2

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

//example 3

var example3Scene *pt.Scene

func example3(frameX, frameY, frameZ int, timed bool) []byte {

	if example3Scene == nil {

		example3Scene = new(pt.Scene)

		// create materials
		gopher := pt.GlossyMaterial(pt.Black, 1.2, pt.Radians(30))
		wall := pt.GlossyMaterial(pt.HexColor(0xFCFAE1), 1.5, pt.Radians(10))
		light := pt.LightMaterial(pt.White, 80)

		// add walls and lights
		example3Scene.Add(pt.NewCube(pt.V(-10, -1, -10), pt.V(-2, 10, 10), wall))
		example3Scene.Add(pt.NewCube(pt.V(-10, -1, -10), pt.V(10, 0, 10), wall))
		example3Scene.Add(pt.NewSphere(pt.V(4, 10, 1), 1, light))

		// load and transform gopher mesh

		for _, s := range []string{"obj", "mtl"} {
			var file = core.NewQFile2(fmt.Sprintf("://qml/gopher.%v", s))
			if file.Open(core.QIODevice__ReadOnly) {
				ioutil.WriteFile(filepath.Join(os.TempDir(), fmt.Sprintf("gopher.%v", s)), []byte(file.ReadAll().ConstData()), 0644)
				file.Close()
			} else {
				panic(fmt.Sprintf("failed to open gopher.%v file", s))
			}
		}

		mesh, err := pt.LoadOBJ(filepath.Join(os.TempDir(), "gopher.obj"), gopher)
		if err != nil {
			panic(err)
		}
		mesh.Transform(pt.Rotate(pt.V(0, 1, 0), pt.Radians(-10)))
		mesh.SmoothNormals()
		mesh.FitInside(pt.Box{pt.V(-1, 0, -1), pt.V(1, 2, 1)}, pt.V(0.5, 0, 0.5))
		example3Scene.Add(mesh)
	}

	cx := math.Cos(ln.Radians(float64(frameX))) * -1
	cy := math.Sin(ln.Radians(float64(frameY))) * -1
	cz := math.Cos(ln.Radians(float64(frameZ))) * -1

	// position camera
	camera := pt.LookAt(pt.V(cx, cy, cz).MulScalar(6), pt.V(0, 0.9, 0), pt.V(0, 1, 0), 40)

	// render the scene
	sampler := pt.NewSampler(4, 4)

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if timed {
		renderer := pt.NewRenderer(example3Scene, &camera, sampler, 512, 512)
		png.Encode(bb, renderer.TimedRender(4*time.Second))
	} else {
		renderer := pt.NewRenderer(example3Scene, &camera, sampler, 128, 128)
		png.Encode(bb, renderer.Render())
	}

	return bb.Bytes()
}

//example 4

var example4Scene *pt.Scene

func example4(frameX, frameY, frameZ int, timed bool) []byte {

	if example4Scene == nil {
		example4Scene = new(pt.Scene)
		example4Scene.Add(pt.NewSphere(pt.V(1.5, 1.25, 0), 1.25, pt.SpecularMaterial(pt.HexColor(0x004358), 1.3)))
		example4Scene.Add(pt.NewSphere(pt.V(-1, 1, 2), 1, pt.SpecularMaterial(pt.HexColor(0xFFE11A), 1.3)))
		example4Scene.Add(pt.NewSphere(pt.V(-2.5, 0.75, 0), 0.75, pt.SpecularMaterial(pt.HexColor(0xFD7400), 1.3)))
		example4Scene.Add(pt.NewSphere(pt.V(-0.75, 0.5, -1), 0.5, pt.ClearMaterial(1.5, 0)))
		example4Scene.Add(pt.NewCube(pt.V(-10, -1, -10), pt.V(10, 0, 10), pt.GlossyMaterial(pt.White, 1.1, pt.Radians(10))))
		example4Scene.Add(pt.NewSphere(pt.V(-1.5, 4, 0), 0.5, pt.LightMaterial(pt.White, 30)))
		camera := pt.LookAt(pt.V(0, 2, -5), pt.V(0, 0.25, 3), pt.V(0, 1, 0), 45)
		camera.SetFocus(pt.V(-0.75, 1, -1), 0.1)
	}

	cx := math.Cos(ln.Radians(float64(frameX)))
	cy := math.Sin(ln.Radians(float64(frameY)))
	cz := math.Cos(ln.Radians(float64(frameZ)))

	// position camera
	camera := pt.LookAt(pt.V(cx, cy, cz).MulScalar(6), pt.V(0, 0.25, 3), pt.V(0, 1, 0), 45)

	// render the scene
	sampler := pt.NewSampler(4, 4)
	sampler.SpecularMode = pt.SpecularModeFirst

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if timed {
		renderer := pt.NewRenderer(example4Scene, &camera, sampler, 512, 512)
		png.Encode(bb, renderer.TimedRender(4*time.Second))
	} else {
		renderer := pt.NewRenderer(example4Scene, &camera, sampler, 128, 128)
		png.Encode(bb, renderer.Render())
	}

	return bb.Bytes()
}

//example 5

var example5Scene *pt.Scene

func example5(frameX, frameY, frameZ int, timed bool) []byte {

	if example5Scene == nil {
		example5Scene = new(pt.Scene)
		r := 0.4
		var material pt.Material

		material = pt.DiffuseMaterial(pt.HexColor(0x334D5C))
		example5Scene.Add(pt.NewSphere(pt.V(-2, r, 0), r, material))

		material = pt.SpecularMaterial(pt.HexColor(0x334D5C), 2)
		example5Scene.Add(pt.NewSphere(pt.V(-1, r, 0), r, material))

		material = pt.GlossyMaterial(pt.HexColor(0x334D5C), 2, pt.Radians(50))
		example5Scene.Add(pt.NewSphere(pt.V(0, r, 0), r, material))

		material = pt.TransparentMaterial(pt.HexColor(0x334D5C), 2, pt.Radians(20), 1)
		example5Scene.Add(pt.NewSphere(pt.V(1, r, 0), r, material))

		material = pt.ClearMaterial(2, 0)
		example5Scene.Add(pt.NewSphere(pt.V(2, r, 0), r, material))

		material = pt.MetallicMaterial(pt.HexColor(0xFFFFFF), 0, 1)
		example5Scene.Add(pt.NewSphere(pt.V(0, 1.5, -4), 1.5, material))

		example5Scene.Add(pt.NewCube(pt.V(-1000, -1, -1000), pt.V(1000, 0, 1000), pt.GlossyMaterial(pt.HexColor(0xFFFFFF), 1.4, pt.Radians(20))))
		example5Scene.Add(pt.NewSphere(pt.V(0, 5, 0), 1, pt.LightMaterial(pt.White, 25)))
	}

	cx := math.Cos(ln.Radians(float64(frameX)))
	cy := math.Sin(ln.Radians(float64(frameY)))
	cz := math.Cos(ln.Radians(float64(frameZ)))

	// position camera
	camera := pt.LookAt(pt.V(cx, cy, cz).MulScalar(8), pt.V(0, 1, 0), pt.V(0, 1, 0), 30)

	// render the scene
	sampler := pt.NewSampler(4, 4)
	sampler.SpecularMode = pt.SpecularModeFirst

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if timed {
		renderer := pt.NewRenderer(example5Scene, &camera, sampler, 512, 512)
		png.Encode(bb, renderer.TimedRender(4*time.Second))
	} else {
		renderer := pt.NewRenderer(example5Scene, &camera, sampler, 128, 128)
		png.Encode(bb, renderer.Render())
	}

	return bb.Bytes()
}
