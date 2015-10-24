package widgets

//#include "qlayout.h"
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

type QLayoutITF interface {
	QLayoutItemITF
	core.QObjectITF
	QLayoutPTR() *QLayout
}

func (p *QLayout) Pointer() unsafe.Pointer {
	return p.QLayoutItemPTR().Pointer()
}

func (p *QLayout) SetPointer(ptr unsafe.Pointer) {
	p.QLayoutItemPTR().SetPointer(ptr)
	p.QObjectPTR().SetPointer(ptr)
}

func PointerFromQLayout(ptr QLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayoutPTR().Pointer()
	}
	return nil
}

func QLayoutFromPointer(ptr unsafe.Pointer) *QLayout {
	var n = new(QLayout)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLayout) QLayoutPTR() *QLayout {
	return ptr
}

//QLayout::SizeConstraint
type QLayout__SizeConstraint int

var (
	QLayout__SetDefaultConstraint = QLayout__SizeConstraint(0)
	QLayout__SetNoConstraint      = QLayout__SizeConstraint(1)
	QLayout__SetMinimumSize       = QLayout__SizeConstraint(2)
	QLayout__SetFixedSize         = QLayout__SizeConstraint(3)
	QLayout__SetMaximumSize       = QLayout__SizeConstraint(4)
	QLayout__SetMinAndMaxSize     = QLayout__SizeConstraint(5)
)

func (ptr *QLayout) SetSizeConstraint(v QLayout__SizeConstraint) {
	if ptr.Pointer() != nil {
		C.QLayout_SetSizeConstraint(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLayout) SetSpacing(v int) {
	if ptr.Pointer() != nil {
		C.QLayout_SetSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLayout) SizeConstraint() QLayout__SizeConstraint {
	if ptr.Pointer() != nil {
		return QLayout__SizeConstraint(C.QLayout_SizeConstraint(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLayout) Spacing() int {
	if ptr.Pointer() != nil {
		return int(C.QLayout_Spacing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLayout) Activate() bool {
	if ptr.Pointer() != nil {
		return C.QLayout_Activate(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLayout) AddItem(item QLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QLayout_AddItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayoutItem(item)))
	}
}

func (ptr *QLayout) AddWidget(w QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QLayout_AddWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(w)))
	}
}

func (ptr *QLayout) ControlTypes() QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QLayout_ControlTypes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QLayout_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLayout) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayout_ExpandingDirections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLayout) GetContentsMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QLayout_GetContentsMargins(C.QtObjectPtr(ptr.Pointer()), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLayout) IndexOf(widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QLayout_IndexOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QLayout) Invalidate() {
	if ptr.Pointer() != nil {
		C.QLayout_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLayout) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QLayout_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLayout) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QLayout_IsEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QLayout_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QLayout) Layout() *QLayout {
	if ptr.Pointer() != nil {
		return QLayoutFromPointer(unsafe.Pointer(C.QLayout_Layout(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLayout) MenuBar() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QLayout_MenuBar(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLayout) ParentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QLayout_ParentWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLayout) RemoveItem(item QLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QLayout_RemoveItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayoutItem(item)))
	}
}

func (ptr *QLayout) RemoveWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QLayout_RemoveWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QLayout) ReplaceWidget(from QWidgetITF, to QWidgetITF, options core.Qt__FindChildOption) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QLayout_ReplaceWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(from)), C.QtObjectPtr(PointerFromQWidget(to)), C.int(options))))
	}
	return nil
}

func (ptr *QLayout) SetAlignment2(l QLayoutITF, alignment core.Qt__AlignmentFlag) bool {
	if ptr.Pointer() != nil {
		return C.QLayout_SetAlignment2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayout(l)), C.int(alignment)) != 0
	}
	return false
}

func (ptr *QLayout) SetAlignment(w QWidgetITF, alignment core.Qt__AlignmentFlag) bool {
	if ptr.Pointer() != nil {
		return C.QLayout_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(w)), C.int(alignment)) != 0
	}
	return false
}

func (ptr *QLayout) SetContentsMargins2(margins core.QMarginsITF) {
	if ptr.Pointer() != nil {
		C.QLayout_SetContentsMargins2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMargins(margins)))
	}
}

func (ptr *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QLayout_SetContentsMargins(C.QtObjectPtr(ptr.Pointer()), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLayout) SetEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QLayout_SetEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QLayout) SetGeometry(r core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QLayout_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(r)))
	}
}

func (ptr *QLayout) SetMenuBar(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QLayout_SetMenuBar(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QLayout_TakeAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QLayout) Update() {
	if ptr.Pointer() != nil {
		C.QLayout_Update(C.QtObjectPtr(ptr.Pointer()))
	}
}
