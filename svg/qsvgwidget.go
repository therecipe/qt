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

type QSvgWidgetITF interface {
	widgets.QWidgetITF
	QSvgWidgetPTR() *QSvgWidget
}

func PointerFromQSvgWidget(ptr QSvgWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgWidgetPTR().Pointer()
	}
	return nil
}

func QSvgWidgetFromPointer(ptr unsafe.Pointer) *QSvgWidget {
	var n = new(QSvgWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSvgWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSvgWidget) QSvgWidgetPTR() *QSvgWidget {
	return ptr
}

func NewQSvgWidget(parent widgets.QWidgetITF) *QSvgWidget {
	return QSvgWidgetFromPointer(unsafe.Pointer(C.QSvgWidget_NewQSvgWidget(C.QtObjectPtr(widgets.PointerFromQWidget(parent)))))
}

func NewQSvgWidget2(file string, parent widgets.QWidgetITF) *QSvgWidget {
	return QSvgWidgetFromPointer(unsafe.Pointer(C.QSvgWidget_NewQSvgWidget2(C.CString(file), C.QtObjectPtr(widgets.PointerFromQWidget(parent)))))
}

func (ptr *QSvgWidget) Load2(contents core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_Load2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(contents)))
	}
}

func (ptr *QSvgWidget) Load(file string) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_Load(C.QtObjectPtr(ptr.Pointer()), C.CString(file))
	}
}

func (ptr *QSvgWidget) Renderer() *QSvgRenderer {
	if ptr.Pointer() != nil {
		return QSvgRendererFromPointer(unsafe.Pointer(C.QSvgWidget_Renderer(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSvgWidget) DestroyQSvgWidget() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_DestroyQSvgWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
