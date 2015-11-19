package svg

//#include "qsvgwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QSvgWidget struct {
	widgets.QWidget
}

type QSvgWidget_ITF interface {
	widgets.QWidget_ITF
	QSvgWidget_PTR() *QSvgWidget
}

func PointerFromQSvgWidget(ptr QSvgWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgWidget_PTR().Pointer()
	}
	return nil
}

func NewQSvgWidgetFromPointer(ptr unsafe.Pointer) *QSvgWidget {
	var n = new(QSvgWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSvgWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSvgWidget) QSvgWidget_PTR() *QSvgWidget {
	return ptr
}

func NewQSvgWidget(parent widgets.QWidget_ITF) *QSvgWidget {
	return NewQSvgWidgetFromPointer(C.QSvgWidget_NewQSvgWidget(widgets.PointerFromQWidget(parent)))
}

func NewQSvgWidget2(file string, parent widgets.QWidget_ITF) *QSvgWidget {
	return NewQSvgWidgetFromPointer(C.QSvgWidget_NewQSvgWidget2(C.CString(file), widgets.PointerFromQWidget(parent)))
}

func (ptr *QSvgWidget) Load2(contents core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_Load2(ptr.Pointer(), core.PointerFromQByteArray(contents))
	}
}

func (ptr *QSvgWidget) Load(file string) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_Load(ptr.Pointer(), C.CString(file))
	}
}

func (ptr *QSvgWidget) Renderer() *QSvgRenderer {
	if ptr.Pointer() != nil {
		return NewQSvgRendererFromPointer(C.QSvgWidget_Renderer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgWidget) DestroyQSvgWidget() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_DestroyQSvgWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
