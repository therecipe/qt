package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QSplashScreen struct {
	QWidget
}

type QSplashScreen_ITF interface {
	QWidget_ITF
	QSplashScreen_PTR() *QSplashScreen
}

func PointerFromQSplashScreen(ptr QSplashScreen_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplashScreen_PTR().Pointer()
	}
	return nil
}

func NewQSplashScreenFromPointer(ptr unsafe.Pointer) *QSplashScreen {
	var n = new(QSplashScreen)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSplashScreen_") {
		n.SetObjectName("QSplashScreen_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSplashScreen) QSplashScreen_PTR() *QSplashScreen {
	return ptr
}

func NewQSplashScreen2(parent QWidget_ITF, pixmap gui.QPixmap_ITF, f core.Qt__WindowType) *QSplashScreen {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::QSplashScreen")
		}
	}()

	return NewQSplashScreenFromPointer(C.QSplashScreen_NewQSplashScreen2(PointerFromQWidget(parent), gui.PointerFromQPixmap(pixmap), C.int(f)))
}

func NewQSplashScreen(pixmap gui.QPixmap_ITF, f core.Qt__WindowType) *QSplashScreen {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::QSplashScreen")
		}
	}()

	return NewQSplashScreenFromPointer(C.QSplashScreen_NewQSplashScreen(gui.PointerFromQPixmap(pixmap), C.int(f)))
}

func (ptr *QSplashScreen) ClearMessage() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::clearMessage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplashScreen_ClearMessage(ptr.Pointer())
	}
}

func (ptr *QSplashScreen) Finish(mainWin QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::finish")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplashScreen_Finish(ptr.Pointer(), PointerFromQWidget(mainWin))
	}
}

func (ptr *QSplashScreen) Message() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::message")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSplashScreen_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSplashScreen) ConnectMessageChanged(f func(message string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::messageChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplashScreen_ConnectMessageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageChanged", f)
	}
}

func (ptr *QSplashScreen) DisconnectMessageChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::messageChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplashScreen_DisconnectMessageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageChanged")
	}
}

//export callbackQSplashScreenMessageChanged
func callbackQSplashScreenMessageChanged(ptrName *C.char, message *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::messageChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "messageChanged").(func(string))(C.GoString(message))
}

func (ptr *QSplashScreen) Repaint() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::repaint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplashScreen_Repaint(ptr.Pointer())
	}
}

func (ptr *QSplashScreen) SetPixmap(pixmap gui.QPixmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::setPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplashScreen_SetPixmap(ptr.Pointer(), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QSplashScreen) ShowMessage(message string, alignment int, color gui.QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::showMessage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplashScreen_ShowMessage(ptr.Pointer(), C.CString(message), C.int(alignment), gui.PointerFromQColor(color))
	}
}

func (ptr *QSplashScreen) DestroyQSplashScreen() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplashScreen::~QSplashScreen")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplashScreen_DestroyQSplashScreen(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
