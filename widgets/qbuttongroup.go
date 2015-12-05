package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QButtonGroup struct {
	core.QObject
}

type QButtonGroup_ITF interface {
	core.QObject_ITF
	QButtonGroup_PTR() *QButtonGroup
}

func PointerFromQButtonGroup(ptr QButtonGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QButtonGroup_PTR().Pointer()
	}
	return nil
}

func NewQButtonGroupFromPointer(ptr unsafe.Pointer) *QButtonGroup {
	var n = new(QButtonGroup)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QButtonGroup_") {
		n.SetObjectName("QButtonGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QButtonGroup) QButtonGroup_PTR() *QButtonGroup {
	return ptr
}

func NewQButtonGroup(parent core.QObject_ITF) *QButtonGroup {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::QButtonGroup")
		}
	}()

	return NewQButtonGroupFromPointer(C.QButtonGroup_NewQButtonGroup(core.PointerFromQObject(parent)))
}

func (ptr *QButtonGroup) AddButton(button QAbstractButton_ITF, id int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::addButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_AddButton(ptr.Pointer(), PointerFromQAbstractButton(button), C.int(id))
	}
}

func (ptr *QButtonGroup) Button(id int) *QAbstractButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::button")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QButtonGroup_Button(ptr.Pointer(), C.int(id)))
	}
	return nil
}

func (ptr *QButtonGroup) CheckedButton() *QAbstractButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::checkedButton")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QButtonGroup_CheckedButton(ptr.Pointer()))
	}
	return nil
}

func (ptr *QButtonGroup) CheckedId() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::checkedId")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QButtonGroup_CheckedId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QButtonGroup) Exclusive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::exclusive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QButtonGroup_Exclusive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QButtonGroup) Id(button QAbstractButton_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::id")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QButtonGroup_Id(ptr.Pointer(), PointerFromQAbstractButton(button)))
	}
	return 0
}

func (ptr *QButtonGroup) RemoveButton(button QAbstractButton_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::removeButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_RemoveButton(ptr.Pointer(), PointerFromQAbstractButton(button))
	}
}

func (ptr *QButtonGroup) SetExclusive(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::setExclusive")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_SetExclusive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QButtonGroup) SetId(button QAbstractButton_ITF, id int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::setId")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_SetId(ptr.Pointer(), PointerFromQAbstractButton(button), C.int(id))
	}
}

func (ptr *QButtonGroup) DestroyQButtonGroup() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::~QButtonGroup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_DestroyQButtonGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QButtonGroup) ConnectButtonClicked(f func(button *QAbstractButton)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonClicked", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonClicked")
	}
}

//export callbackQButtonGroupButtonClicked
func callbackQButtonGroupButtonClicked(ptrName *C.char, button unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "buttonClicked").(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
}

func (ptr *QButtonGroup) ConnectButtonPressed(f func(button *QAbstractButton)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonPressed", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonPressed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonPressed")
	}
}

//export callbackQButtonGroupButtonPressed
func callbackQButtonGroupButtonPressed(ptrName *C.char, button unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonPressed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "buttonPressed").(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
}

func (ptr *QButtonGroup) ConnectButtonReleased(f func(button *QAbstractButton)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonReleased")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonReleased(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonReleased", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonReleased() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonReleased")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonReleased")
	}
}

//export callbackQButtonGroupButtonReleased
func callbackQButtonGroupButtonReleased(ptrName *C.char, button unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonReleased")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "buttonReleased").(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
}

func (ptr *QButtonGroup) ConnectButtonToggled(f func(button *QAbstractButton, checked bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonToggled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonToggled", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonToggled() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonToggled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonToggled")
	}
}

//export callbackQButtonGroupButtonToggled
func callbackQButtonGroupButtonToggled(ptrName *C.char, button unsafe.Pointer, checked C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QButtonGroup::buttonToggled")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "buttonToggled").(func(*QAbstractButton, bool))(NewQAbstractButtonFromPointer(button), int(checked) != 0)
}
