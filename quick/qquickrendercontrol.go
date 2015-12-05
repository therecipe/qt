package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QQuickRenderControl struct {
	core.QObject
}

type QQuickRenderControl_ITF interface {
	core.QObject_ITF
	QQuickRenderControl_PTR() *QQuickRenderControl
}

func PointerFromQQuickRenderControl(ptr QQuickRenderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickRenderControl_PTR().Pointer()
	}
	return nil
}

func NewQQuickRenderControlFromPointer(ptr unsafe.Pointer) *QQuickRenderControl {
	var n = new(QQuickRenderControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickRenderControl_") {
		n.SetObjectName("QQuickRenderControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickRenderControl) QQuickRenderControl_PTR() *QQuickRenderControl {
	return ptr
}

func NewQQuickRenderControl(parent core.QObject_ITF) *QQuickRenderControl {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::QQuickRenderControl")
		}
	}()

	return NewQQuickRenderControlFromPointer(C.QQuickRenderControl_NewQQuickRenderControl(core.PointerFromQObject(parent)))
}

func (ptr *QQuickRenderControl) Initialize(gl gui.QOpenGLContext_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::initialize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Initialize(ptr.Pointer(), gui.PointerFromQOpenGLContext(gl))
	}
}

func (ptr *QQuickRenderControl) Invalidate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::invalidate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Invalidate(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) PolishItems() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::polishItems")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PolishItems(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) PrepareThread(targetThread core.QThread_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::prepareThread")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PrepareThread(ptr.Pointer(), core.PointerFromQThread(targetThread))
	}
}

func (ptr *QQuickRenderControl) Render() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::render")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Render(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) ConnectRenderRequested(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::renderRequested")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectRenderRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "renderRequested", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectRenderRequested() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::renderRequested")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectRenderRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "renderRequested")
	}
}

//export callbackQQuickRenderControlRenderRequested
func callbackQQuickRenderControlRenderRequested(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::renderRequested")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "renderRequested").(func())()
}

func (ptr *QQuickRenderControl) RenderWindow(offset core.QPoint_ITF) *gui.QWindow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::renderWindow")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QQuickRenderControl_RenderWindow(ptr.Pointer(), core.PointerFromQPoint(offset)))
	}
	return nil
}

func QQuickRenderControl_RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::renderWindowFor")
		}
	}()

	return gui.NewQWindowFromPointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(PointerFromQQuickWindow(win), core.PointerFromQPoint(offset)))
}

func (ptr *QQuickRenderControl) ConnectSceneChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::sceneChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectSceneChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneChanged", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectSceneChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::sceneChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectSceneChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneChanged")
	}
}

//export callbackQQuickRenderControlSceneChanged
func callbackQQuickRenderControlSceneChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::sceneChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sceneChanged").(func())()
}

func (ptr *QQuickRenderControl) Sync() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::sync")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_Sync(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) DestroyQQuickRenderControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickRenderControl::~QQuickRenderControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DestroyQQuickRenderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
