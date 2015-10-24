package widgets

//#include "qsplashscreen.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSplashScreen struct {
	QWidget
}

type QSplashScreenITF interface {
	QWidgetITF
	QSplashScreenPTR() *QSplashScreen
}

func PointerFromQSplashScreen(ptr QSplashScreenITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplashScreenPTR().Pointer()
	}
	return nil
}

func QSplashScreenFromPointer(ptr unsafe.Pointer) *QSplashScreen {
	var n = new(QSplashScreen)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSplashScreen_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSplashScreen) QSplashScreenPTR() *QSplashScreen {
	return ptr
}

func NewQSplashScreen2(parent QWidgetITF, pixmap gui.QPixmapITF, f core.Qt__WindowType) *QSplashScreen {
	return QSplashScreenFromPointer(unsafe.Pointer(C.QSplashScreen_NewQSplashScreen2(C.QtObjectPtr(PointerFromQWidget(parent)), C.QtObjectPtr(gui.PointerFromQPixmap(pixmap)), C.int(f))))
}

func NewQSplashScreen(pixmap gui.QPixmapITF, f core.Qt__WindowType) *QSplashScreen {
	return QSplashScreenFromPointer(unsafe.Pointer(C.QSplashScreen_NewQSplashScreen(C.QtObjectPtr(gui.PointerFromQPixmap(pixmap)), C.int(f))))
}

func (ptr *QSplashScreen) ClearMessage() {
	if ptr.Pointer() != nil {
		C.QSplashScreen_ClearMessage(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSplashScreen) Finish(mainWin QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QSplashScreen_Finish(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(mainWin)))
	}
}

func (ptr *QSplashScreen) Message() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSplashScreen_Message(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSplashScreen) ConnectMessageChanged(f func(message string)) {
	if ptr.Pointer() != nil {
		C.QSplashScreen_ConnectMessageChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "messageChanged", f)
	}
}

func (ptr *QSplashScreen) DisconnectMessageChanged() {
	if ptr.Pointer() != nil {
		C.QSplashScreen_DisconnectMessageChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "messageChanged")
	}
}

//export callbackQSplashScreenMessageChanged
func callbackQSplashScreenMessageChanged(ptrName *C.char, message *C.char) {
	qt.GetSignal(C.GoString(ptrName), "messageChanged").(func(string))(C.GoString(message))
}

func (ptr *QSplashScreen) Repaint() {
	if ptr.Pointer() != nil {
		C.QSplashScreen_Repaint(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSplashScreen) SetPixmap(pixmap gui.QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QSplashScreen_SetPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QSplashScreen) ShowMessage(message string, alignment int, color gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QSplashScreen_ShowMessage(C.QtObjectPtr(ptr.Pointer()), C.CString(message), C.int(alignment), C.QtObjectPtr(gui.PointerFromQColor(color)))
	}
}

func (ptr *QSplashScreen) DestroyQSplashScreen() {
	if ptr.Pointer() != nil {
		C.QSplashScreen_DestroyQSplashScreen(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
