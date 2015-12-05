package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QStatusBar struct {
	QWidget
}

type QStatusBar_ITF interface {
	QWidget_ITF
	QStatusBar_PTR() *QStatusBar
}

func PointerFromQStatusBar(ptr QStatusBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStatusBar_PTR().Pointer()
	}
	return nil
}

func NewQStatusBarFromPointer(ptr unsafe.Pointer) *QStatusBar {
	var n = new(QStatusBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStatusBar_") {
		n.SetObjectName("QStatusBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStatusBar) QStatusBar_PTR() *QStatusBar {
	return ptr
}

func (ptr *QStatusBar) IsSizeGripEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::isSizeGripEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStatusBar_IsSizeGripEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStatusBar) SetSizeGripEnabled(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::setSizeGripEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStatusBar_SetSizeGripEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQStatusBar(parent QWidget_ITF) *QStatusBar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::QStatusBar")
		}
	}()

	return NewQStatusBarFromPointer(C.QStatusBar_NewQStatusBar(PointerFromQWidget(parent)))
}

func (ptr *QStatusBar) AddPermanentWidget(widget QWidget_ITF, stretch int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::addPermanentWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStatusBar_AddPermanentWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch))
	}
}

func (ptr *QStatusBar) AddWidget(widget QWidget_ITF, stretch int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::addWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStatusBar_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch))
	}
}

func (ptr *QStatusBar) ClearMessage() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::clearMessage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStatusBar_ClearMessage(ptr.Pointer())
	}
}

func (ptr *QStatusBar) CurrentMessage() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::currentMessage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStatusBar_CurrentMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStatusBar) InsertPermanentWidget(index int, widget QWidget_ITF, stretch int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::insertPermanentWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertPermanentWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) InsertWidget(index int, widget QWidget_ITF, stretch int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::insertWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) ConnectMessageChanged(f func(message string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::messageChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStatusBar_ConnectMessageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageChanged", f)
	}
}

func (ptr *QStatusBar) DisconnectMessageChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::messageChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStatusBar_DisconnectMessageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageChanged")
	}
}

//export callbackQStatusBarMessageChanged
func callbackQStatusBarMessageChanged(ptrName *C.char, message *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::messageChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "messageChanged").(func(string))(C.GoString(message))
}

func (ptr *QStatusBar) RemoveWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::removeWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStatusBar_RemoveWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStatusBar) ShowMessage(message string, timeout int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::showMessage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStatusBar_ShowMessage(ptr.Pointer(), C.CString(message), C.int(timeout))
	}
}

func (ptr *QStatusBar) DestroyQStatusBar() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusBar::~QStatusBar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStatusBar_DestroyQStatusBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
