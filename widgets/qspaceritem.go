package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSpacerItem struct {
	QLayoutItem
}

type QSpacerItem_ITF interface {
	QLayoutItem_ITF
	QSpacerItem_PTR() *QSpacerItem
}

func PointerFromQSpacerItem(ptr QSpacerItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSpacerItem_PTR().Pointer()
	}
	return nil
}

func NewQSpacerItemFromPointer(ptr unsafe.Pointer) *QSpacerItem {
	var n = new(QSpacerItem)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSpacerItem_") {
		n.SetObjectNameAbs("QSpacerItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QSpacerItem) QSpacerItem_PTR() *QSpacerItem {
	return ptr
}

func NewQSpacerItem(w int, h int, hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) *QSpacerItem {
	defer qt.Recovering("QSpacerItem::QSpacerItem")

	return NewQSpacerItemFromPointer(C.QSpacerItem_NewQSpacerItem(C.int(w), C.int(h), C.int(hPolicy), C.int(vPolicy)))
}

func (ptr *QSpacerItem) ChangeSize(w int, h int, hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) {
	defer qt.Recovering("QSpacerItem::changeSize")

	if ptr.Pointer() != nil {
		C.QSpacerItem_ChangeSize(ptr.Pointer(), C.int(w), C.int(h), C.int(hPolicy), C.int(vPolicy))
	}
}

func (ptr *QSpacerItem) ExpandingDirections() core.Qt__Orientation {
	defer qt.Recovering("QSpacerItem::expandingDirections")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSpacerItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpacerItem) Geometry() *core.QRect {
	defer qt.Recovering("QSpacerItem::geometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QSpacerItem_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSpacerItem) IsEmpty() bool {
	defer qt.Recovering("QSpacerItem::isEmpty")

	if ptr.Pointer() != nil {
		return C.QSpacerItem_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSpacerItem) MaximumSize() *core.QSize {
	defer qt.Recovering("QSpacerItem::maximumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSpacerItem_MaximumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSpacerItem) MinimumSize() *core.QSize {
	defer qt.Recovering("QSpacerItem::minimumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSpacerItem_MinimumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSpacerItem) ConnectSetGeometry(f func(r *core.QRect)) {
	defer qt.Recovering("connect QSpacerItem::setGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setGeometry", f)
	}
}

func (ptr *QSpacerItem) DisconnectSetGeometry() {
	defer qt.Recovering("disconnect QSpacerItem::setGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setGeometry")
	}
}

//export callbackQSpacerItemSetGeometry
func callbackQSpacerItemSetGeometry(ptr unsafe.Pointer, ptrName *C.char, r unsafe.Pointer) {
	defer qt.Recovering("callback QSpacerItem::setGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "setGeometry"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(r))
	} else {
		NewQSpacerItemFromPointer(ptr).SetGeometryDefault(core.NewQRectFromPointer(r))
	}
}

func (ptr *QSpacerItem) SetGeometry(r core.QRect_ITF) {
	defer qt.Recovering("QSpacerItem::setGeometry")

	if ptr.Pointer() != nil {
		C.QSpacerItem_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QSpacerItem) SetGeometryDefault(r core.QRect_ITF) {
	defer qt.Recovering("QSpacerItem::setGeometry")

	if ptr.Pointer() != nil {
		C.QSpacerItem_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QSpacerItem) SizeHint() *core.QSize {
	defer qt.Recovering("QSpacerItem::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSpacerItem_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSpacerItem) SpacerItem() *QSpacerItem {
	defer qt.Recovering("QSpacerItem::spacerItem")

	if ptr.Pointer() != nil {
		return NewQSpacerItemFromPointer(C.QSpacerItem_SpacerItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSpacerItem) DestroyQSpacerItem() {
	defer qt.Recovering("QSpacerItem::~QSpacerItem")

	if ptr.Pointer() != nil {
		C.QSpacerItem_DestroyQSpacerItem(ptr.Pointer())
	}
}

func (ptr *QSpacerItem) ObjectNameAbs() string {
	defer qt.Recovering("QSpacerItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSpacerItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpacerItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSpacerItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSpacerItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSpacerItem) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QSpacerItem::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "invalidate", f)
	}
}

func (ptr *QSpacerItem) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QSpacerItem::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "invalidate")
	}
}

//export callbackQSpacerItemInvalidate
func callbackQSpacerItemInvalidate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSpacerItem::invalidate")

	if signal := qt.GetSignal(C.GoString(ptrName), "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQSpacerItemFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QSpacerItem) Invalidate() {
	defer qt.Recovering("QSpacerItem::invalidate")

	if ptr.Pointer() != nil {
		C.QSpacerItem_Invalidate(ptr.Pointer())
	}
}

func (ptr *QSpacerItem) InvalidateDefault() {
	defer qt.Recovering("QSpacerItem::invalidate")

	if ptr.Pointer() != nil {
		C.QSpacerItem_InvalidateDefault(ptr.Pointer())
	}
}
