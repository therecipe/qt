//source: http://doc.qt.io/qt-5/qtgui-openglwindow-example.html

package main

/*
#include <stdint.h>
float colors[] = { 1.0f, 0.0f, 0.0f, 0.0f, 1.0f, 0.0f, 0.0f, 0.0f, 1.0f };
float vertices[] = { 0.0f, 0.707f, -0.5f, -0.5f, 0.5f, -0.5f };
*/
import "C"
import (
	"os"
	"unsafe"

	"github.com/therecipe/qt/gui"
)

var (
	colors   = unsafe.Pointer(&C.colors)
	vertices = unsafe.Pointer(&C.vertices)
)

const (
	vertexShaderSource = "attribute highp vec4 posAttr;\n" +
		"attribute lowp vec4 colAttr;\n" +
		"varying lowp vec4 col;\n" +
		"uniform highp mat4 matrix;\n" +
		"void main() {\n" +
		"   col = colAttr;\n" +
		"   gl_Position = matrix * posAttr;\n" +
		"}\n"

	fragmentShaderSource = "varying lowp vec4 col;\n" +
		"void main() {\n" +
		"   gl_FragColor = col;\n" +
		"}\n"
)

type TriangleWindow struct {
	OpenGLWindow

	_ func() `constructor:"init"`

	m_posAttr       uint
	m_colAttr       uint
	m_matrixUniform int

	m_program *gui.QOpenGLShaderProgram
	m_frame   float32
}

func (w *TriangleWindow) init() {
	w.OpenGLWindow.initializeLazy = w.initialize
	w.OpenGLWindow.renderLazy = w.render

	w.ConnectEvent(w.OpenGLWindow.event)
	w.ConnectExposeEvent(w.OpenGLWindow.exposeEvent)
}

func (w *TriangleWindow) initialize() {
	w.m_program = gui.NewQOpenGLShaderProgram(w)
	w.m_program.AddShaderFromSourceCode(gui.QOpenGLShader__Vertex, vertexShaderSource)
	w.m_program.AddShaderFromSourceCode(gui.QOpenGLShader__Fragment, fragmentShaderSource)
	w.m_program.Link()
	w.m_posAttr = uint(w.m_program.AttributeLocation("posAttr"))
	w.m_colAttr = uint(w.m_program.AttributeLocation("colAttr"))
	w.m_matrixUniform = w.m_program.UniformLocation("matrix")
}

func (w *TriangleWindow) render(painter *gui.QPainter) {

	retinaScale := int(w.DevicePixelRatio())
	w.GlViewport(0, 0, w.Width()*retinaScale, w.Height()*retinaScale)

	w.GlClear(GL_COLOR_BUFFER_BIT)

	w.m_program.Bind()

	matrix := gui.NewQMatrix4x4()
	matrix.Perspective(60, float32(4)/float32(3), 0.1, 100)
	matrix.Translate3(0, 0, -2)
	matrix.Rotate2(100*w.m_frame/float32(w.Screen().RefreshRate()), 0, 1, 0)

	w.m_program.SetUniformValue23(w.m_matrixUniform, matrix)

	w.GlVertexAttribPointer(w.m_posAttr, 2, GL_FLOAT, GL_FALSE, 0, vertices)
	w.GlVertexAttribPointer(w.m_colAttr, 3, GL_FLOAT, GL_FALSE, 0, colors)

	w.GlEnableVertexAttribArray(0)
	w.GlEnableVertexAttribArray(1)

	w.GlDrawArrays(GL_TRIANGLES, 0, 3)

	w.GlDisableVertexAttribArray(1)
	w.GlDisableVertexAttribArray(0)

	w.m_program.Release()

	w.m_frame++
}

func main() {
	gui.NewQGuiApplication(len(os.Args), os.Args)

	format := gui.NewQSurfaceFormat()
	format.SetSamples(16)

	window := NewTriangleWindow(nil)
	window.SetFormat(format)
	window.Resize2(640, 480)
	window.Show()

	window.setAnimating(true)

	gui.QGuiApplication_Exec()
}
