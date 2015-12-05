package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QKeySequenceEdit struct {
	QWidget
}

type QKeySequenceEdit_ITF interface {
	QWidget_ITF
	QKeySequenceEdit_PTR() *QKeySequenceEdit
}

func PointerFromQKeySequenceEdit(ptr QKeySequenceEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeySequenceEdit_PTR().Pointer()
	}
	return nil
}

func NewQKeySequenceEditFromPointer(ptr unsafe.Pointer) *QKeySequenceEdit {
	var n = new(QKeySequenceEdit)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QKeySequenceEdit_") {
		n.SetObjectName("QKeySequenceEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QKeySequenceEdit) QKeySequenceEdit_PTR() *QKeySequenceEdit {
	return ptr
}

func (ptr *QKeySequenceEdit) SetKeySequence(keySequence gui.QKeySequence_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeySequenceEdit::setKeySequence")
		}
	}()

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_SetKeySequence(ptr.Pointer(), gui.PointerFromQKeySequence(keySequence))
	}
}

func NewQKeySequenceEdit(parent QWidget_ITF) *QKeySequenceEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeySequenceEdit::QKeySequenceEdit")
		}
	}()

	return NewQKeySequenceEditFromPointer(C.QKeySequenceEdit_NewQKeySequenceEdit(PointerFromQWidget(parent)))
}

func NewQKeySequenceEdit2(keySequence gui.QKeySequence_ITF, parent QWidget_ITF) *QKeySequenceEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeySequenceEdit::QKeySequenceEdit")
		}
	}()

	return NewQKeySequenceEditFromPointer(C.QKeySequenceEdit_NewQKeySequenceEdit2(gui.PointerFromQKeySequence(keySequence), PointerFromQWidget(parent)))
}

func (ptr *QKeySequenceEdit) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeySequenceEdit::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QKeySequenceEdit) ConnectEditingFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeySequenceEdit::editingFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ConnectEditingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectEditingFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeySequenceEdit::editingFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQKeySequenceEditEditingFinished
func callbackQKeySequenceEditEditingFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeySequenceEdit::editingFinished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "editingFinished").(func())()
}

func (ptr *QKeySequenceEdit) DestroyQKeySequenceEdit() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeySequenceEdit::~QKeySequenceEdit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DestroyQKeySequenceEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
