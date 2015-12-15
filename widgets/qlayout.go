package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLayout struct {
	QLayoutItem
	core.QObject
}

type QLayout_ITF interface {
	QLayoutItem_ITF
	core.QObject_ITF
	QLayout_PTR() *QLayout
}

func (p *QLayout) Pointer() unsafe.Pointer {
	return p.QLayoutItem_PTR().Pointer()
}

func (p *QLayout) SetPointer(ptr unsafe.Pointer) {
	p.QLayoutItem_PTR().SetPointer(ptr)
	p.QObject_PTR().SetPointer(ptr)
}

func PointerFromQLayout(ptr QLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayout_PTR().Pointer()
	}
	return nil
}

func NewQLayoutFromPointer(ptr unsafe.Pointer) *QLayout {
	var n = new(QLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLayout_") {
		n.SetObjectName("QLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QLayout) QLayout_PTR() *QLayout {
	return ptr
}

//QLayout::SizeConstraint
type QLayout__SizeConstraint int64

const (
	QLayout__SetDefaultConstraint = QLayout__SizeConstraint(0)
	QLayout__SetNoConstraint      = QLayout__SizeConstraint(1)
	QLayout__SetMinimumSize       = QLayout__SizeConstraint(2)
	QLayout__SetFixedSize         = QLayout__SizeConstraint(3)
	QLayout__SetMaximumSize       = QLayout__SizeConstraint(4)
	QLayout__SetMinAndMaxSize     = QLayout__SizeConstraint(5)
)

func (ptr *QLayout) SetSizeConstraint(v QLayout__SizeConstraint) {
	defer qt.Recovering("QLayout::setSizeConstraint")

	if ptr.Pointer() != nil {
		C.QLayout_SetSizeConstraint(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLayout) SetSpacing(v int) {
	defer qt.Recovering("QLayout::setSpacing")

	if ptr.Pointer() != nil {
		C.QLayout_SetSpacing(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLayout) SizeConstraint() QLayout__SizeConstraint {
	defer qt.Recovering("QLayout::sizeConstraint")

	if ptr.Pointer() != nil {
		return QLayout__SizeConstraint(C.QLayout_SizeConstraint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) Spacing() int {
	defer qt.Recovering("QLayout::spacing")

	if ptr.Pointer() != nil {
		return int(C.QLayout_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) Activate() bool {
	defer qt.Recovering("QLayout::activate")

	if ptr.Pointer() != nil {
		return C.QLayout_Activate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayout) AddItem(item QLayoutItem_ITF) {
	defer qt.Recovering("QLayout::addItem")

	if ptr.Pointer() != nil {
		C.QLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QLayout) AddWidget(w QWidget_ITF) {
	defer qt.Recovering("QLayout::addWidget")

	if ptr.Pointer() != nil {
		C.QLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(w))
	}
}

func (ptr *QLayout) ConnectChildEvent(f func(e *core.QChildEvent)) {
	defer qt.Recovering("connect QLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLayout) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLayoutChildEvent
func callbackQLayoutChildEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLayout::childEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "childEvent")
	if signal != nil {
		defer signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QLayout) ControlTypes() QSizePolicy__ControlType {
	defer qt.Recovering("QLayout::controlTypes")

	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QLayout_ControlTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) Count() int {
	defer qt.Recovering("QLayout::count")

	if ptr.Pointer() != nil {
		return int(C.QLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) ExpandingDirections() core.Qt__Orientation {
	defer qt.Recovering("QLayout::expandingDirections")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayout_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) GetContentsMargins(left int, top int, right int, bottom int) {
	defer qt.Recovering("QLayout::getContentsMargins")

	if ptr.Pointer() != nil {
		C.QLayout_GetContentsMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLayout) IndexOf(widget QWidget_ITF) int {
	defer qt.Recovering("QLayout::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QLayout_IndexOf(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QLayout) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "invalidate", f)
	}
}

func (ptr *QLayout) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "invalidate")
	}
}

//export callbackQLayoutInvalidate
func callbackQLayoutInvalidate(ptrName *C.char) bool {
	defer qt.Recovering("callback QLayout::invalidate")

	var signal = qt.GetSignal(C.GoString(ptrName), "invalidate")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QLayout) IsEmpty() bool {
	defer qt.Recovering("QLayout::isEmpty")

	if ptr.Pointer() != nil {
		return C.QLayout_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayout) IsEnabled() bool {
	defer qt.Recovering("QLayout::isEnabled")

	if ptr.Pointer() != nil {
		return C.QLayout_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayout) ItemAt(index int) *QLayoutItem {
	defer qt.Recovering("QLayout::itemAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QLayout) Layout() *QLayout {
	defer qt.Recovering("QLayout::layout")

	if ptr.Pointer() != nil {
		return NewQLayoutFromPointer(C.QLayout_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) MenuBar() *QWidget {
	defer qt.Recovering("QLayout::menuBar")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLayout_MenuBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) ParentWidget() *QWidget {
	defer qt.Recovering("QLayout::parentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLayout_ParentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) RemoveItem(item QLayoutItem_ITF) {
	defer qt.Recovering("QLayout::removeItem")

	if ptr.Pointer() != nil {
		C.QLayout_RemoveItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QLayout) RemoveWidget(widget QWidget_ITF) {
	defer qt.Recovering("QLayout::removeWidget")

	if ptr.Pointer() != nil {
		C.QLayout_RemoveWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QLayout) ReplaceWidget(from QWidget_ITF, to QWidget_ITF, options core.Qt__FindChildOption) *QLayoutItem {
	defer qt.Recovering("QLayout::replaceWidget")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_ReplaceWidget(ptr.Pointer(), PointerFromQWidget(from), PointerFromQWidget(to), C.int(options)))
	}
	return nil
}

func (ptr *QLayout) SetAlignment2(l QLayout_ITF, alignment core.Qt__AlignmentFlag) bool {
	defer qt.Recovering("QLayout::setAlignment")

	if ptr.Pointer() != nil {
		return C.QLayout_SetAlignment2(ptr.Pointer(), PointerFromQLayout(l), C.int(alignment)) != 0
	}
	return false
}

func (ptr *QLayout) SetAlignment(w QWidget_ITF, alignment core.Qt__AlignmentFlag) bool {
	defer qt.Recovering("QLayout::setAlignment")

	if ptr.Pointer() != nil {
		return C.QLayout_SetAlignment(ptr.Pointer(), PointerFromQWidget(w), C.int(alignment)) != 0
	}
	return false
}

func (ptr *QLayout) SetContentsMargins2(margins core.QMargins_ITF) {
	defer qt.Recovering("QLayout::setContentsMargins")

	if ptr.Pointer() != nil {
		C.QLayout_SetContentsMargins2(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	defer qt.Recovering("QLayout::setContentsMargins")

	if ptr.Pointer() != nil {
		C.QLayout_SetContentsMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLayout) SetEnabled(enable bool) {
	defer qt.Recovering("QLayout::setEnabled")

	if ptr.Pointer() != nil {
		C.QLayout_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QLayout) SetMenuBar(widget QWidget_ITF) {
	defer qt.Recovering("QLayout::setMenuBar")

	if ptr.Pointer() != nil {
		C.QLayout_SetMenuBar(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QLayout) TakeAt(index int) *QLayoutItem {
	defer qt.Recovering("QLayout::takeAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QLayout) Update() {
	defer qt.Recovering("QLayout::update")

	if ptr.Pointer() != nil {
		C.QLayout_Update(ptr.Pointer())
	}
}
