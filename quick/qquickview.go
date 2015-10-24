package quick

//#include "qquickview.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"unsafe"
)

type QQuickView struct {
	QQuickWindow
}

type QQuickViewITF interface {
	QQuickWindowITF
	QQuickViewPTR() *QQuickView
}

func PointerFromQQuickView(ptr QQuickViewITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickViewPTR().Pointer()
	}
	return nil
}

func QQuickViewFromPointer(ptr unsafe.Pointer) *QQuickView {
	var n = new(QQuickView)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickView) QQuickViewPTR() *QQuickView {
	return ptr
}

//QQuickView::ResizeMode
type QQuickView__ResizeMode int

var (
	QQuickView__SizeViewToRootObject = QQuickView__ResizeMode(0)
	QQuickView__SizeRootObjectToView = QQuickView__ResizeMode(1)
)

//QQuickView::Status
type QQuickView__Status int

var (
	QQuickView__Null    = QQuickView__Status(0)
	QQuickView__Ready   = QQuickView__Status(1)
	QQuickView__Loading = QQuickView__Status(2)
	QQuickView__Error   = QQuickView__Status(3)
)

func (ptr *QQuickView) ResizeMode() QQuickView__ResizeMode {
	if ptr.Pointer() != nil {
		return QQuickView__ResizeMode(C.QQuickView_ResizeMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickView) SetResizeMode(v QQuickView__ResizeMode) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetResizeMode(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QQuickView) Status() QQuickView__Status {
	if ptr.Pointer() != nil {
		return QQuickView__Status(C.QQuickView_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQQuickView2(engine qml.QQmlEngineITF, parent gui.QWindowITF) *QQuickView {
	return QQuickViewFromPointer(unsafe.Pointer(C.QQuickView_NewQQuickView2(C.QtObjectPtr(qml.PointerFromQQmlEngine(engine)), C.QtObjectPtr(gui.PointerFromQWindow(parent)))))
}

func NewQQuickView(parent gui.QWindowITF) *QQuickView {
	return QQuickViewFromPointer(unsafe.Pointer(C.QQuickView_NewQQuickView(C.QtObjectPtr(gui.PointerFromQWindow(parent)))))
}

func NewQQuickView3(source string, parent gui.QWindowITF) *QQuickView {
	return QQuickViewFromPointer(unsafe.Pointer(C.QQuickView_NewQQuickView3(C.CString(source), C.QtObjectPtr(gui.PointerFromQWindow(parent)))))
}

func (ptr *QQuickView) Engine() *qml.QQmlEngine {
	if ptr.Pointer() != nil {
		return qml.QQmlEngineFromPointer(unsafe.Pointer(C.QQuickView_Engine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickView) RootContext() *qml.QQmlContext {
	if ptr.Pointer() != nil {
		return qml.QQmlContextFromPointer(unsafe.Pointer(C.QQuickView_RootContext(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickView) RootObject() *QQuickItem {
	if ptr.Pointer() != nil {
		return QQuickItemFromPointer(unsafe.Pointer(C.QQuickView_RootObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickView) SetSource(url string) {
	if ptr.Pointer() != nil {
		C.QQuickView_SetSource(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QQuickView) Source() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickView_Source(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQuickView) ConnectStatusChanged(f func(status QQuickView__Status)) {
	if ptr.Pointer() != nil {
		C.QQuickView_ConnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QQuickView) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QQuickView_DisconnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQQuickViewStatusChanged
func callbackQQuickViewStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QQuickView__Status))(QQuickView__Status(status))
}

func (ptr *QQuickView) DestroyQQuickView() {
	if ptr.Pointer() != nil {
		C.QQuickView_DestroyQQuickView(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
