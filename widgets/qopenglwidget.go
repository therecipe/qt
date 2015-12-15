package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QOpenGLWidget struct {
	QWidget
}

type QOpenGLWidget_ITF interface {
	QWidget_ITF
	QOpenGLWidget_PTR() *QOpenGLWidget
}

func PointerFromQOpenGLWidget(ptr QOpenGLWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLWidget_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLWidgetFromPointer(ptr unsafe.Pointer) *QOpenGLWidget {
	var n = new(QOpenGLWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QOpenGLWidget_") {
		n.SetObjectName("QOpenGLWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QOpenGLWidget) QOpenGLWidget_PTR() *QOpenGLWidget {
	return ptr
}

//QOpenGLWidget::UpdateBehavior
type QOpenGLWidget__UpdateBehavior int64

const (
	QOpenGLWidget__NoPartialUpdate = QOpenGLWidget__UpdateBehavior(0)
	QOpenGLWidget__PartialUpdate   = QOpenGLWidget__UpdateBehavior(1)
)
