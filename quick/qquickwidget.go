package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
	"log"
	"unsafe"
)

type QQuickWidget struct {
	widgets.QWidget
}

type QQuickWidget_ITF interface {
	widgets.QWidget_ITF
	QQuickWidget_PTR() *QQuickWidget
}

func PointerFromQQuickWidget(ptr QQuickWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWidget_PTR().Pointer()
	}
	return nil
}

func NewQQuickWidgetFromPointer(ptr unsafe.Pointer) *QQuickWidget {
	var n = new(QQuickWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickWidget_") {
		n.SetObjectName("QQuickWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickWidget) QQuickWidget_PTR() *QQuickWidget {
	return ptr
}

//QQuickWidget::ResizeMode
type QQuickWidget__ResizeMode int64

const (
	QQuickWidget__SizeViewToRootObject = QQuickWidget__ResizeMode(0)
	QQuickWidget__SizeRootObjectToView = QQuickWidget__ResizeMode(1)
)

//QQuickWidget::Status
type QQuickWidget__Status int64

const (
	QQuickWidget__Null    = QQuickWidget__Status(0)
	QQuickWidget__Ready   = QQuickWidget__Status(1)
	QQuickWidget__Loading = QQuickWidget__Status(2)
	QQuickWidget__Error   = QQuickWidget__Status(3)
)

func (ptr *QQuickWidget) ResizeMode() QQuickWidget__ResizeMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::resizeMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QQuickWidget__ResizeMode(C.QQuickWidget_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWidget) SetResizeMode(v QQuickWidget__ResizeMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::setResizeMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetResizeMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QQuickWidget) Status() QQuickWidget__Status {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::status")
		}
	}()

	if ptr.Pointer() != nil {
		return QQuickWidget__Status(C.QQuickWidget_Status(ptr.Pointer()))
	}
	return 0
}

func NewQQuickWidget2(engine qml.QQmlEngine_ITF, parent widgets.QWidget_ITF) *QQuickWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::QQuickWidget")
		}
	}()

	return NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget2(qml.PointerFromQQmlEngine(engine), widgets.PointerFromQWidget(parent)))
}

func NewQQuickWidget(parent widgets.QWidget_ITF) *QQuickWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::QQuickWidget")
		}
	}()

	return NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget(widgets.PointerFromQWidget(parent)))
}

func NewQQuickWidget3(source core.QUrl_ITF, parent widgets.QWidget_ITF) *QQuickWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::QQuickWidget")
		}
	}()

	return NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget3(core.PointerFromQUrl(source), widgets.PointerFromQWidget(parent)))
}

func (ptr *QQuickWidget) Engine() *qml.QQmlEngine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::engine")
		}
	}()

	if ptr.Pointer() != nil {
		return qml.NewQQmlEngineFromPointer(C.QQuickWidget_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) QuickWindow() *QQuickWindow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::quickWindow")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQQuickWindowFromPointer(C.QQuickWidget_QuickWindow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) RootContext() *qml.QQmlContext {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::rootContext")
		}
	}()

	if ptr.Pointer() != nil {
		return qml.NewQQmlContextFromPointer(C.QQuickWidget_RootContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) RootObject() *QQuickItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::rootObject")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWidget_RootObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::sceneGraphError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectSceneGraphError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphError", f)
	}
}

func (ptr *QQuickWidget) DisconnectSceneGraphError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::sceneGraphError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphError")
	}
}

//export callbackQQuickWidgetSceneGraphError
func callbackQQuickWidgetSceneGraphError(ptrName *C.char, error C.int, message *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::sceneGraphError")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sceneGraphError").(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
}

func (ptr *QQuickWidget) SetClearColor(color gui.QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::setClearColor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetClearColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QQuickWidget) SetFormat(format gui.QSurfaceFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::setFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetFormat(ptr.Pointer(), gui.PointerFromQSurfaceFormat(format))
	}
}

func (ptr *QQuickWidget) SetSource(url core.QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::setSource")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQuickWidget) ConnectStatusChanged(f func(status QQuickWidget__Status)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::statusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QQuickWidget) DisconnectStatusChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::statusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQQuickWidgetStatusChanged
func callbackQQuickWidgetStatusChanged(ptrName *C.char, status C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::statusChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QQuickWidget__Status))(QQuickWidget__Status(status))
}

func (ptr *QQuickWidget) DestroyQQuickWidget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickWidget::~QQuickWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickWidget_DestroyQQuickWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
