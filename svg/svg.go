// +build !minimal

package svg

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "svg.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtSvg_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtSvg_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QGraphicsSvgItem struct {
	widgets.QGraphicsObject
}

type QGraphicsSvgItem_ITF interface {
	widgets.QGraphicsObject_ITF
	QGraphicsSvgItem_PTR() *QGraphicsSvgItem
}

func (ptr *QGraphicsSvgItem) QGraphicsSvgItem_PTR() *QGraphicsSvgItem {
	return ptr
}

func (ptr *QGraphicsSvgItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsSvgItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsObject_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsSvgItem(ptr QGraphicsSvgItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSvgItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSvgItemFromPointer(ptr unsafe.Pointer) (n *QGraphicsSvgItem) {
	n = new(QGraphicsSvgItem)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QGraphicsSvgItem__anonymous
//QGraphicsSvgItem::anonymous
type QGraphicsSvgItem__anonymous int64

const (
	QGraphicsSvgItem__Type QGraphicsSvgItem__anonymous = QGraphicsSvgItem__anonymous(13)
)

func NewQGraphicsSvgItem(parent widgets.QGraphicsItem_ITF) *QGraphicsSvgItem {
	tmpValue := NewQGraphicsSvgItemFromPointer(C.QGraphicsSvgItem_NewQGraphicsSvgItem(widgets.PointerFromQGraphicsItem(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQGraphicsSvgItem2(fileName string, parent widgets.QGraphicsItem_ITF) *QGraphicsSvgItem {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	tmpValue := NewQGraphicsSvgItemFromPointer(C.QGraphicsSvgItem_NewQGraphicsSvgItem2(C.struct_QtSvg_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, widgets.PointerFromQGraphicsItem(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGraphicsSvgItem_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGraphicsSvgItem_QGraphicsSvgItem_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QGraphicsSvgItem) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGraphicsSvgItem_QGraphicsSvgItem_Tr(sC, cC, C.int(int32(n))))
}

func QGraphicsSvgItem_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGraphicsSvgItem_QGraphicsSvgItem_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QGraphicsSvgItem) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGraphicsSvgItem_QGraphicsSvgItem_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQGraphicsSvgItem_Paint
func callbackQGraphicsSvgItem_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsSvgItem) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "paint", func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget) {
				signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(painter, option, widget)
				f(painter, option, widget)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paint", f)
		}
	}
}

func (ptr *QGraphicsSvgItem) DisconnectPaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paint")
	}
}

func (ptr *QGraphicsSvgItem) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsSvgItem) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsSvgItem) SetElementId(id string) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		C.QGraphicsSvgItem_SetElementId(ptr.Pointer(), C.struct_QtSvg_PackedString{data: idC, len: C.longlong(len(id))})
	}
}

func (ptr *QGraphicsSvgItem) SetMaximumCacheSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetMaximumCacheSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QGraphicsSvgItem) SetSharedRenderer(renderer QSvgRenderer_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetSharedRenderer(ptr.Pointer(), PointerFromQSvgRenderer(renderer))
	}
}

//export callbackQGraphicsSvgItem_BoundingRect
func callbackQGraphicsSvgItem_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQGraphicsSvgItemFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QGraphicsSvgItem) ConnectBoundingRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "boundingRect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "boundingRect", func() *core.QRectF {
				signal.(func() *core.QRectF)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "boundingRect", f)
		}
	}
}

func (ptr *QGraphicsSvgItem) DisconnectBoundingRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "boundingRect")
	}
}

func (ptr *QGraphicsSvgItem) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsSvgItem_BoundingRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) BoundingRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsSvgItem_BoundingRectDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) MaximumCacheSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QGraphicsSvgItem_MaximumCacheSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) ElementId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGraphicsSvgItem_ElementId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsSvgItem) Renderer() *QSvgRenderer {
	if ptr.Pointer() != nil {
		tmpValue := NewQSvgRendererFromPointer(C.QGraphicsSvgItem_Renderer(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSvgItem_MetaObject
func callbackQGraphicsSvgItem_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGraphicsSvgItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGraphicsSvgItem) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGraphicsSvgItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsSvgItem_Type
func callbackQGraphicsSvgItem_Type(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQGraphicsSvgItemFromPointer(ptr).TypeDefault()))
}

func (ptr *QGraphicsSvgItem) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGraphicsSvgItem_TypeDefault(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSvgItem) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGraphicsSvgItem___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGraphicsSvgItem) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGraphicsSvgItem___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGraphicsSvgItem) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsSvgItem___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsSvgItem) __findChildren_newList2() unsafe.Pointer {
	return C.QGraphicsSvgItem___findChildren_newList2(ptr.Pointer())
}

func (ptr *QGraphicsSvgItem) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsSvgItem___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsSvgItem) __findChildren_newList3() unsafe.Pointer {
	return C.QGraphicsSvgItem___findChildren_newList3(ptr.Pointer())
}

func (ptr *QGraphicsSvgItem) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsSvgItem___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsSvgItem) __findChildren_newList() unsafe.Pointer {
	return C.QGraphicsSvgItem___findChildren_newList(ptr.Pointer())
}

func (ptr *QGraphicsSvgItem) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsSvgItem___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsSvgItem) __children_newList() unsafe.Pointer {
	return C.QGraphicsSvgItem___children_newList(ptr.Pointer())
}

func (ptr *QGraphicsSvgItem) __setTransformations_transformations_atList(i int) *widgets.QGraphicsTransform {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQGraphicsTransformFromPointer(C.QGraphicsSvgItem___setTransformations_transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) __setTransformations_transformations_setList(i widgets.QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem___setTransformations_transformations_setList(ptr.Pointer(), widgets.PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QGraphicsSvgItem) __setTransformations_transformations_newList() unsafe.Pointer {
	return C.QGraphicsSvgItem___setTransformations_transformations_newList(ptr.Pointer())
}

func (ptr *QGraphicsSvgItem) __childItems_atList(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QGraphicsSvgItem___childItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QGraphicsSvgItem) __childItems_setList(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem___childItems_setList(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QGraphicsSvgItem) __childItems_newList() unsafe.Pointer {
	return C.QGraphicsSvgItem___childItems_newList(ptr.Pointer())
}

func (ptr *QGraphicsSvgItem) __collidingItems_atList(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QGraphicsSvgItem___collidingItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QGraphicsSvgItem) __collidingItems_setList(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem___collidingItems_setList(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QGraphicsSvgItem) __collidingItems_newList() unsafe.Pointer {
	return C.QGraphicsSvgItem___collidingItems_newList(ptr.Pointer())
}

func (ptr *QGraphicsSvgItem) __transformations_atList(i int) *widgets.QGraphicsTransform {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQGraphicsTransformFromPointer(C.QGraphicsSvgItem___transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) __transformations_setList(i widgets.QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem___transformations_setList(ptr.Pointer(), widgets.PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QGraphicsSvgItem) __transformations_newList() unsafe.Pointer {
	return C.QGraphicsSvgItem___transformations_newList(ptr.Pointer())
}

//export callbackQGraphicsSvgItem_Event
func callbackQGraphicsSvgItem_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev)))))
}

func (ptr *QGraphicsSvgItem) EventDefault(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsSvgItem_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_EnabledChanged
func callbackQGraphicsSvgItem_EnabledChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enabledChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsSvgItem_OpacityChanged
func callbackQGraphicsSvgItem_OpacityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsSvgItem_ParentChanged
func callbackQGraphicsSvgItem_ParentChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "parentChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsSvgItem_RotationChanged
func callbackQGraphicsSvgItem_RotationChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rotationChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsSvgItem_ScaleChanged
func callbackQGraphicsSvgItem_ScaleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scaleChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsSvgItem_UpdateMicroFocus
func callbackQGraphicsSvgItem_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QGraphicsSvgItem) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsSvgItem_VisibleChanged
func callbackQGraphicsSvgItem_VisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsSvgItem_XChanged
func callbackQGraphicsSvgItem_XChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "xChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsSvgItem_YChanged
func callbackQGraphicsSvgItem_YChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "yChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsSvgItem_ZChanged
func callbackQGraphicsSvgItem_ZChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "zChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsSvgItem_EventFilter
func callbackQGraphicsSvgItem_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsSvgItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsSvgItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_ChildEvent
func callbackQGraphicsSvgItem_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGraphicsSvgItem_ConnectNotify
func callbackQGraphicsSvgItem_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsSvgItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsSvgItem_CustomEvent
func callbackQGraphicsSvgItem_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsSvgItem_DeleteLater
func callbackQGraphicsSvgItem_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGraphicsSvgItem) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGraphicsSvgItem_Destroyed
func callbackQGraphicsSvgItem_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGraphicsSvgItem_DisconnectNotify
func callbackQGraphicsSvgItem_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsSvgItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsSvgItem_ObjectNameChanged
func callbackQGraphicsSvgItem_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSvg_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQGraphicsSvgItem_TimerEvent
func callbackQGraphicsSvgItem_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGraphicsSvgItem_ItemChange
func callbackQGraphicsSvgItem_ItemChange(ptr unsafe.Pointer, change C.longlong, value unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemChange"); signal != nil {
		return core.PointerFromQVariant(signal.(func(widgets.QGraphicsItem__GraphicsItemChange, *core.QVariant) *core.QVariant)(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
	}

	return core.PointerFromQVariant(NewQGraphicsSvgItemFromPointer(ptr).ItemChangeDefault(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
}

func (ptr *QGraphicsSvgItem) ItemChangeDefault(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QGraphicsSvgItem_ItemChangeDefault(ptr.Pointer(), C.longlong(change), core.PointerFromQVariant(value)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSvgItem_SceneEvent
func callbackQGraphicsSvgItem_SceneEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "sceneEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).SceneEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsSvgItem) SceneEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsSvgItem_SceneEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_SceneEventFilter
func callbackQGraphicsSvgItem_SceneEventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "sceneEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem, *core.QEvent) bool)(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).SceneEventFilterDefault(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsSvgItem) SceneEventFilterDefault(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsSvgItem_SceneEventFilterDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_Advance
func callbackQGraphicsSvgItem_Advance(ptr unsafe.Pointer, phase C.int) {
	if signal := qt.GetSignal(ptr, "advance"); signal != nil {
		signal.(func(int))(int(int32(phase)))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).AdvanceDefault(int(int32(phase)))
	}
}

func (ptr *QGraphicsSvgItem) AdvanceDefault(phase int) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_AdvanceDefault(ptr.Pointer(), C.int(int32(phase)))
	}
}

//export callbackQGraphicsSvgItem_ContextMenuEvent
func callbackQGraphicsSvgItem_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneContextMenuEvent))(widgets.NewQGraphicsSceneContextMenuEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).ContextMenuEventDefault(widgets.NewQGraphicsSceneContextMenuEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ContextMenuEventDefault(event widgets.QGraphicsSceneContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ContextMenuEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

//export callbackQGraphicsSvgItem_DragEnterEvent
func callbackQGraphicsSvgItem_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DragEnterEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) DragEnterEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DragEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsSvgItem_DragLeaveEvent
func callbackQGraphicsSvgItem_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DragLeaveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) DragLeaveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DragLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsSvgItem_DragMoveEvent
func callbackQGraphicsSvgItem_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DragMoveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) DragMoveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DragMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsSvgItem_DropEvent
func callbackQGraphicsSvgItem_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DropEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) DropEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DropEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsSvgItem_FocusInEvent
func callbackQGraphicsSvgItem_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQGraphicsSvgItem_FocusOutEvent
func callbackQGraphicsSvgItem_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQGraphicsSvgItem_HoverEnterEvent
func callbackQGraphicsSvgItem_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverEnterEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).HoverEnterEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) HoverEnterEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_HoverEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsSvgItem_HoverLeaveEvent
func callbackQGraphicsSvgItem_HoverLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverLeaveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).HoverLeaveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) HoverLeaveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_HoverLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsSvgItem_HoverMoveEvent
func callbackQGraphicsSvgItem_HoverMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).HoverMoveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) HoverMoveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_HoverMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsSvgItem_InputMethodEvent
func callbackQGraphicsSvgItem_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQGraphicsSvgItem_KeyPressEvent
func callbackQGraphicsSvgItem_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQGraphicsSvgItem_KeyReleaseEvent
func callbackQGraphicsSvgItem_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQGraphicsSvgItem_MouseDoubleClickEvent
func callbackQGraphicsSvgItem_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).MouseDoubleClickEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) MouseDoubleClickEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MouseDoubleClickEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsSvgItem_MouseMoveEvent
func callbackQGraphicsSvgItem_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).MouseMoveEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) MouseMoveEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MouseMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsSvgItem_MousePressEvent
func callbackQGraphicsSvgItem_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).MousePressEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) MousePressEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MousePressEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsSvgItem_MouseReleaseEvent
func callbackQGraphicsSvgItem_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).MouseReleaseEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) MouseReleaseEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MouseReleaseEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsSvgItem_WheelEvent
func callbackQGraphicsSvgItem_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneWheelEvent))(widgets.NewQGraphicsSceneWheelEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).WheelEventDefault(widgets.NewQGraphicsSceneWheelEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) WheelEventDefault(event widgets.QGraphicsSceneWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_WheelEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneWheelEvent(event))
	}
}

//export callbackQGraphicsSvgItem_OpaqueArea
func callbackQGraphicsSvgItem_OpaqueArea(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "opaqueArea"); signal != nil {
		return gui.PointerFromQPainterPath(signal.(func() *gui.QPainterPath)())
	}

	return gui.PointerFromQPainterPath(NewQGraphicsSvgItemFromPointer(ptr).OpaqueAreaDefault())
}

func (ptr *QGraphicsSvgItem) OpaqueAreaDefault() *gui.QPainterPath {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPainterPathFromPointer(C.QGraphicsSvgItem_OpaqueAreaDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSvgItem_Shape
func callbackQGraphicsSvgItem_Shape(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "shape"); signal != nil {
		return gui.PointerFromQPainterPath(signal.(func() *gui.QPainterPath)())
	}

	return gui.PointerFromQPainterPath(NewQGraphicsSvgItemFromPointer(ptr).ShapeDefault())
}

func (ptr *QGraphicsSvgItem) ShapeDefault() *gui.QPainterPath {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPainterPathFromPointer(C.QGraphicsSvgItem_ShapeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSvgItem_InputMethodQuery
func callbackQGraphicsSvgItem_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQGraphicsSvgItemFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QGraphicsSvgItem) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QGraphicsSvgItem_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSvgItem_CollidesWithItem
func callbackQGraphicsSvgItem_CollidesWithItem(ptr unsafe.Pointer, other unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "collidesWithItem"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem, core.Qt__ItemSelectionMode) bool)(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).CollidesWithItemDefault(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode)))))
}

func (ptr *QGraphicsSvgItem) CollidesWithItemDefault(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsSvgItem_CollidesWithItemDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(other), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_CollidesWithPath
func callbackQGraphicsSvgItem_CollidesWithPath(ptr unsafe.Pointer, path unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "collidesWithPath"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*gui.QPainterPath, core.Qt__ItemSelectionMode) bool)(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).CollidesWithPathDefault(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode)))))
}

func (ptr *QGraphicsSvgItem) CollidesWithPathDefault(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsSvgItem_CollidesWithPathDefault(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_Contains
func callbackQGraphicsSvgItem_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point)))))
}

func (ptr *QGraphicsSvgItem) ContainsDefault(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsSvgItem_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point))) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_IsObscuredBy
func callbackQGraphicsSvgItem_IsObscuredBy(ptr unsafe.Pointer, item unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isObscuredBy"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem) bool)(widgets.NewQGraphicsItemFromPointer(item)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).IsObscuredByDefault(widgets.NewQGraphicsItemFromPointer(item)))))
}

func (ptr *QGraphicsSvgItem) IsObscuredByDefault(item widgets.QGraphicsItem_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsSvgItem_IsObscuredByDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

type QSvgGenerator struct {
	gui.QPaintDevice
}

type QSvgGenerator_ITF interface {
	gui.QPaintDevice_ITF
	QSvgGenerator_PTR() *QSvgGenerator
}

func (ptr *QSvgGenerator) QSvgGenerator_PTR() *QSvgGenerator {
	return ptr
}

func (ptr *QSvgGenerator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QSvgGenerator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQSvgGenerator(ptr QSvgGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgGenerator_PTR().Pointer()
	}
	return nil
}

func NewQSvgGeneratorFromPointer(ptr unsafe.Pointer) (n *QSvgGenerator) {
	n = new(QSvgGenerator)
	n.SetPointer(ptr)
	return
}
func NewQSvgGenerator() *QSvgGenerator {
	return NewQSvgGeneratorFromPointer(C.QSvgGenerator_NewQSvgGenerator())
}

func (ptr *QSvgGenerator) SetDescription(description string) {
	if ptr.Pointer() != nil {
		var descriptionC *C.char
		if description != "" {
			descriptionC = C.CString(description)
			defer C.free(unsafe.Pointer(descriptionC))
		}
		C.QSvgGenerator_SetDescription(ptr.Pointer(), C.struct_QtSvg_PackedString{data: descriptionC, len: C.longlong(len(description))})
	}
}

func (ptr *QSvgGenerator) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QSvgGenerator_SetFileName(ptr.Pointer(), C.struct_QtSvg_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

func (ptr *QSvgGenerator) SetOutputDevice(outputDevice core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetOutputDevice(ptr.Pointer(), core.PointerFromQIODevice(outputDevice))
	}
}

func (ptr *QSvgGenerator) SetResolution(dpi int) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetResolution(ptr.Pointer(), C.int(int32(dpi)))
	}
}

func (ptr *QSvgGenerator) SetSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSvgGenerator) SetTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QSvgGenerator_SetTitle(ptr.Pointer(), C.struct_QtSvg_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *QSvgGenerator) SetViewBox(viewBox core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox(ptr.Pointer(), core.PointerFromQRect(viewBox))
	}
}

func (ptr *QSvgGenerator) SetViewBox2(viewBox core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox2(ptr.Pointer(), core.PointerFromQRectF(viewBox))
	}
}

//export callbackQSvgGenerator_DestroyQSvgGenerator
func callbackQSvgGenerator_DestroyQSvgGenerator(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSvgGenerator"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgGeneratorFromPointer(ptr).DestroyQSvgGeneratorDefault()
	}
}

func (ptr *QSvgGenerator) ConnectDestroyQSvgGenerator(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSvgGenerator"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSvgGenerator", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSvgGenerator", f)
		}
	}
}

func (ptr *QSvgGenerator) DisconnectDestroyQSvgGenerator() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSvgGenerator")
	}
}

func (ptr *QSvgGenerator) DestroyQSvgGenerator() {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_DestroyQSvgGenerator(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSvgGenerator) DestroyQSvgGeneratorDefault() {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_DestroyQSvgGeneratorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSvgGenerator) OutputDevice() *core.QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQIODeviceFromPointer(C.QSvgGenerator_OutputDevice(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSvgGenerator_PaintEngine
func callbackQSvgGenerator_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQSvgGeneratorFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QSvgGenerator) ConnectPaintEngine(f func() *gui.QPaintEngine) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paintEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "paintEngine", func() *gui.QPaintEngine {
				signal.(func() *gui.QPaintEngine)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paintEngine", f)
		}
	}
}

func (ptr *QSvgGenerator) DisconnectPaintEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paintEngine")
	}
}

func (ptr *QSvgGenerator) PaintEngine() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QSvgGenerator_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QSvgGenerator_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) ViewBox() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QSvgGenerator_ViewBox(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgGenerator) ViewBoxF() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSvgGenerator_ViewBoxF(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgGenerator) Size() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSvgGenerator_Size(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgGenerator) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSvgGenerator_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSvgGenerator_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSvgGenerator_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) Resolution() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSvgGenerator_Resolution(ptr.Pointer())))
	}
	return 0
}

type QSvgRenderer struct {
	core.QObject
}

type QSvgRenderer_ITF interface {
	core.QObject_ITF
	QSvgRenderer_PTR() *QSvgRenderer
}

func (ptr *QSvgRenderer) QSvgRenderer_PTR() *QSvgRenderer {
	return ptr
}

func (ptr *QSvgRenderer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSvgRenderer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSvgRenderer(ptr QSvgRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgRenderer_PTR().Pointer()
	}
	return nil
}

func NewQSvgRendererFromPointer(ptr unsafe.Pointer) (n *QSvgRenderer) {
	n = new(QSvgRenderer)
	n.SetPointer(ptr)
	return
}
func QSvgRenderer_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSvgRenderer_QSvgRenderer_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QSvgRenderer) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSvgRenderer_QSvgRenderer_Tr(sC, cC, C.int(int32(n))))
}

func QSvgRenderer_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSvgRenderer_QSvgRenderer_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QSvgRenderer) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSvgRenderer_QSvgRenderer_TrUtf8(sC, cC, C.int(int32(n))))
}

func NewQSvgRenderer(parent core.QObject_ITF) *QSvgRenderer {
	tmpValue := NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSvgRenderer4(contents core.QXmlStreamReader_ITF, parent core.QObject_ITF) *QSvgRenderer {
	tmpValue := NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer4(core.PointerFromQXmlStreamReader(contents), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSvgRenderer3(contents core.QByteArray_ITF, parent core.QObject_ITF) *QSvgRenderer {
	tmpValue := NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer3(core.PointerFromQByteArray(contents), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSvgRenderer2(filename string, parent core.QObject_ITF) *QSvgRenderer {
	var filenameC *C.char
	if filename != "" {
		filenameC = C.CString(filename)
		defer C.free(unsafe.Pointer(filenameC))
	}
	tmpValue := NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer2(C.struct_QtSvg_PackedString{data: filenameC, len: C.longlong(len(filename))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSvgRenderer_Load3
func callbackQSvgRenderer_Load3(ptr unsafe.Pointer, contents unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "load3"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QXmlStreamReader) bool)(core.NewQXmlStreamReaderFromPointer(contents)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgRendererFromPointer(ptr).Load3Default(core.NewQXmlStreamReaderFromPointer(contents)))))
}

func (ptr *QSvgRenderer) ConnectLoad3(f func(contents *core.QXmlStreamReader) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "load3"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "load3", func(contents *core.QXmlStreamReader) bool {
				signal.(func(*core.QXmlStreamReader) bool)(contents)
				return f(contents)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "load3", f)
		}
	}
}

func (ptr *QSvgRenderer) DisconnectLoad3() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "load3")
	}
}

func (ptr *QSvgRenderer) Load3(contents core.QXmlStreamReader_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgRenderer_Load3(ptr.Pointer(), core.PointerFromQXmlStreamReader(contents))) != 0
	}
	return false
}

func (ptr *QSvgRenderer) Load3Default(contents core.QXmlStreamReader_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgRenderer_Load3Default(ptr.Pointer(), core.PointerFromQXmlStreamReader(contents))) != 0
	}
	return false
}

//export callbackQSvgRenderer_Load2
func callbackQSvgRenderer_Load2(ptr unsafe.Pointer, contents unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "load2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray) bool)(core.NewQByteArrayFromPointer(contents)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgRendererFromPointer(ptr).Load2Default(core.NewQByteArrayFromPointer(contents)))))
}

func (ptr *QSvgRenderer) ConnectLoad2(f func(contents *core.QByteArray) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "load2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "load2", func(contents *core.QByteArray) bool {
				signal.(func(*core.QByteArray) bool)(contents)
				return f(contents)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "load2", f)
		}
	}
}

func (ptr *QSvgRenderer) DisconnectLoad2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "load2")
	}
}

func (ptr *QSvgRenderer) Load2(contents core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgRenderer_Load2(ptr.Pointer(), core.PointerFromQByteArray(contents))) != 0
	}
	return false
}

func (ptr *QSvgRenderer) Load2Default(contents core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgRenderer_Load2Default(ptr.Pointer(), core.PointerFromQByteArray(contents))) != 0
	}
	return false
}

//export callbackQSvgRenderer_Load
func callbackQSvgRenderer_Load(ptr unsafe.Pointer, filename C.struct_QtSvg_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "load"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(filename)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgRendererFromPointer(ptr).LoadDefault(cGoUnpackString(filename)))))
}

func (ptr *QSvgRenderer) ConnectLoad(f func(filename string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "load"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "load", func(filename string) bool {
				signal.(func(string) bool)(filename)
				return f(filename)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "load", f)
		}
	}
}

func (ptr *QSvgRenderer) DisconnectLoad() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "load")
	}
}

func (ptr *QSvgRenderer) Load(filename string) bool {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		return int8(C.QSvgRenderer_Load(ptr.Pointer(), C.struct_QtSvg_PackedString{data: filenameC, len: C.longlong(len(filename))})) != 0
	}
	return false
}

func (ptr *QSvgRenderer) LoadDefault(filename string) bool {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		return int8(C.QSvgRenderer_LoadDefault(ptr.Pointer(), C.struct_QtSvg_PackedString{data: filenameC, len: C.longlong(len(filename))})) != 0
	}
	return false
}

//export callbackQSvgRenderer_Render
func callbackQSvgRenderer_Render(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "render"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQSvgRendererFromPointer(ptr).RenderDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QSvgRenderer) ConnectRender(f func(painter *gui.QPainter)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "render"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "render", func(painter *gui.QPainter) {
				signal.(func(*gui.QPainter))(painter)
				f(painter)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "render", f)
		}
	}
}

func (ptr *QSvgRenderer) DisconnectRender() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "render")
	}
}

func (ptr *QSvgRenderer) Render(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSvgRenderer) RenderDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_RenderDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQSvgRenderer_Render2
func callbackQSvgRenderer_Render2(ptr unsafe.Pointer, painter unsafe.Pointer, bounds unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "render2"); signal != nil {
		signal.(func(*gui.QPainter, *core.QRectF))(gui.NewQPainterFromPointer(painter), core.NewQRectFFromPointer(bounds))
	} else {
		NewQSvgRendererFromPointer(ptr).Render2Default(gui.NewQPainterFromPointer(painter), core.NewQRectFFromPointer(bounds))
	}
}

func (ptr *QSvgRenderer) ConnectRender2(f func(painter *gui.QPainter, bounds *core.QRectF)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "render2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "render2", func(painter *gui.QPainter, bounds *core.QRectF) {
				signal.(func(*gui.QPainter, *core.QRectF))(painter, bounds)
				f(painter, bounds)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "render2", f)
		}
	}
}

func (ptr *QSvgRenderer) DisconnectRender2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "render2")
	}
}

func (ptr *QSvgRenderer) Render2(painter gui.QPainter_ITF, bounds core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render2(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRectF(bounds))
	}
}

func (ptr *QSvgRenderer) Render2Default(painter gui.QPainter_ITF, bounds core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render2Default(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRectF(bounds))
	}
}

//export callbackQSvgRenderer_Render3
func callbackQSvgRenderer_Render3(ptr unsafe.Pointer, painter unsafe.Pointer, elementId C.struct_QtSvg_PackedString, bounds unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "render3"); signal != nil {
		signal.(func(*gui.QPainter, string, *core.QRectF))(gui.NewQPainterFromPointer(painter), cGoUnpackString(elementId), core.NewQRectFFromPointer(bounds))
	} else {
		NewQSvgRendererFromPointer(ptr).Render3Default(gui.NewQPainterFromPointer(painter), cGoUnpackString(elementId), core.NewQRectFFromPointer(bounds))
	}
}

func (ptr *QSvgRenderer) ConnectRender3(f func(painter *gui.QPainter, elementId string, bounds *core.QRectF)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "render3"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "render3", func(painter *gui.QPainter, elementId string, bounds *core.QRectF) {
				signal.(func(*gui.QPainter, string, *core.QRectF))(painter, elementId, bounds)
				f(painter, elementId, bounds)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "render3", f)
		}
	}
}

func (ptr *QSvgRenderer) DisconnectRender3() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "render3")
	}
}

func (ptr *QSvgRenderer) Render3(painter gui.QPainter_ITF, elementId string, bounds core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		var elementIdC *C.char
		if elementId != "" {
			elementIdC = C.CString(elementId)
			defer C.free(unsafe.Pointer(elementIdC))
		}
		C.QSvgRenderer_Render3(ptr.Pointer(), gui.PointerFromQPainter(painter), C.struct_QtSvg_PackedString{data: elementIdC, len: C.longlong(len(elementId))}, core.PointerFromQRectF(bounds))
	}
}

func (ptr *QSvgRenderer) Render3Default(painter gui.QPainter_ITF, elementId string, bounds core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		var elementIdC *C.char
		if elementId != "" {
			elementIdC = C.CString(elementId)
			defer C.free(unsafe.Pointer(elementIdC))
		}
		C.QSvgRenderer_Render3Default(ptr.Pointer(), gui.PointerFromQPainter(painter), C.struct_QtSvg_PackedString{data: elementIdC, len: C.longlong(len(elementId))}, core.PointerFromQRectF(bounds))
	}
}

//export callbackQSvgRenderer_RepaintNeeded
func callbackQSvgRenderer_RepaintNeeded(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaintNeeded"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSvgRenderer) ConnectRepaintNeeded(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "repaintNeeded") {
			C.QSvgRenderer_ConnectRepaintNeeded(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "repaintNeeded"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "repaintNeeded", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "repaintNeeded", f)
		}
	}
}

func (ptr *QSvgRenderer) DisconnectRepaintNeeded() {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_DisconnectRepaintNeeded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "repaintNeeded")
	}
}

func (ptr *QSvgRenderer) RepaintNeeded() {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_RepaintNeeded(ptr.Pointer())
	}
}

func (ptr *QSvgRenderer) SetFramesPerSecond(num int) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetFramesPerSecond(ptr.Pointer(), C.int(int32(num)))
	}
}

func (ptr *QSvgRenderer) SetViewBox(viewbox core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetViewBox(ptr.Pointer(), core.PointerFromQRect(viewbox))
	}
}

func (ptr *QSvgRenderer) SetViewBox2(viewbox core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetViewBox2(ptr.Pointer(), core.PointerFromQRectF(viewbox))
	}
}

//export callbackQSvgRenderer_DestroyQSvgRenderer
func callbackQSvgRenderer_DestroyQSvgRenderer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSvgRenderer"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgRendererFromPointer(ptr).DestroyQSvgRendererDefault()
	}
}

func (ptr *QSvgRenderer) ConnectDestroyQSvgRenderer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSvgRenderer"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSvgRenderer", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSvgRenderer", f)
		}
	}
}

func (ptr *QSvgRenderer) DisconnectDestroyQSvgRenderer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSvgRenderer")
	}
}

func (ptr *QSvgRenderer) DestroyQSvgRenderer() {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_DestroyQSvgRenderer(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSvgRenderer) DestroyQSvgRendererDefault() {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_DestroyQSvgRendererDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSvgRenderer) MatrixForElement(id string) *gui.QMatrix {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		tmpValue := gui.NewQMatrixFromPointer(C.QSvgRenderer_MatrixForElement(ptr.Pointer(), C.struct_QtSvg_PackedString{data: idC, len: C.longlong(len(id))}))
		runtime.SetFinalizer(tmpValue, (*gui.QMatrix).DestroyQMatrix)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) ViewBox() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QSvgRenderer_ViewBox(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) BoundsOnElement(id string) *core.QRectF {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		tmpValue := core.NewQRectFFromPointer(C.QSvgRenderer_BoundsOnElement(ptr.Pointer(), C.struct_QtSvg_PackedString{data: idC, len: C.longlong(len(id))}))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) ViewBoxF() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QSvgRenderer_ViewBoxF(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) DefaultSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSvgRenderer_DefaultSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) Animated() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgRenderer_Animated(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSvgRenderer) ElementExists(id string) bool {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		return int8(C.QSvgRenderer_ElementExists(ptr.Pointer(), C.struct_QtSvg_PackedString{data: idC, len: C.longlong(len(id))})) != 0
	}
	return false
}

func (ptr *QSvgRenderer) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgRenderer_IsValid(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSvgRenderer_MetaObject
func callbackQSvgRenderer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSvgRendererFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSvgRenderer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSvgRenderer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgRenderer) FramesPerSecond() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSvgRenderer_FramesPerSecond(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSvgRenderer) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSvgRenderer___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSvgRenderer) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSvgRenderer___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSvgRenderer) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSvgRenderer___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSvgRenderer) __findChildren_newList2() unsafe.Pointer {
	return C.QSvgRenderer___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSvgRenderer) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSvgRenderer___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSvgRenderer) __findChildren_newList3() unsafe.Pointer {
	return C.QSvgRenderer___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSvgRenderer) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSvgRenderer___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSvgRenderer) __findChildren_newList() unsafe.Pointer {
	return C.QSvgRenderer___findChildren_newList(ptr.Pointer())
}

func (ptr *QSvgRenderer) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSvgRenderer___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSvgRenderer) __children_newList() unsafe.Pointer {
	return C.QSvgRenderer___children_newList(ptr.Pointer())
}

//export callbackQSvgRenderer_Event
func callbackQSvgRenderer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgRendererFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSvgRenderer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgRenderer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSvgRenderer_EventFilter
func callbackQSvgRenderer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgRendererFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSvgRenderer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgRenderer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSvgRenderer_ChildEvent
func callbackQSvgRenderer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSvgRendererFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSvgRenderer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSvgRenderer_ConnectNotify
func callbackQSvgRenderer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSvgRendererFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSvgRenderer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSvgRenderer_CustomEvent
func callbackQSvgRenderer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgRendererFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgRenderer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSvgRenderer_DeleteLater
func callbackQSvgRenderer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgRendererFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSvgRenderer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSvgRenderer_Destroyed
func callbackQSvgRenderer_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSvgRenderer_DisconnectNotify
func callbackQSvgRenderer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSvgRendererFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSvgRenderer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSvgRenderer_ObjectNameChanged
func callbackQSvgRenderer_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSvg_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSvgRenderer_TimerEvent
func callbackQSvgRenderer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSvgRendererFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSvgRenderer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSvgWidget struct {
	widgets.QWidget
}

type QSvgWidget_ITF interface {
	widgets.QWidget_ITF
	QSvgWidget_PTR() *QSvgWidget
}

func (ptr *QSvgWidget) QSvgWidget_PTR() *QSvgWidget {
	return ptr
}

func (ptr *QSvgWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QSvgWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQSvgWidget(ptr QSvgWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgWidget_PTR().Pointer()
	}
	return nil
}

func NewQSvgWidgetFromPointer(ptr unsafe.Pointer) (n *QSvgWidget) {
	n = new(QSvgWidget)
	n.SetPointer(ptr)
	return
}
func QSvgWidget_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSvgWidget_QSvgWidget_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QSvgWidget) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSvgWidget_QSvgWidget_Tr(sC, cC, C.int(int32(n))))
}

func QSvgWidget_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSvgWidget_QSvgWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QSvgWidget) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSvgWidget_QSvgWidget_TrUtf8(sC, cC, C.int(int32(n))))
}

func NewQSvgWidget(parent widgets.QWidget_ITF) *QSvgWidget {
	tmpValue := NewQSvgWidgetFromPointer(C.QSvgWidget_NewQSvgWidget(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSvgWidget2(file string, parent widgets.QWidget_ITF) *QSvgWidget {
	var fileC *C.char
	if file != "" {
		fileC = C.CString(file)
		defer C.free(unsafe.Pointer(fileC))
	}
	tmpValue := NewQSvgWidgetFromPointer(C.QSvgWidget_NewQSvgWidget2(C.struct_QtSvg_PackedString{data: fileC, len: C.longlong(len(file))}, widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSvgWidget_Load2
func callbackQSvgWidget_Load2(ptr unsafe.Pointer, contents unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "load2"); signal != nil {
		signal.(func(*core.QByteArray))(core.NewQByteArrayFromPointer(contents))
	} else {
		NewQSvgWidgetFromPointer(ptr).Load2Default(core.NewQByteArrayFromPointer(contents))
	}
}

func (ptr *QSvgWidget) ConnectLoad2(f func(contents *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "load2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "load2", func(contents *core.QByteArray) {
				signal.(func(*core.QByteArray))(contents)
				f(contents)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "load2", f)
		}
	}
}

func (ptr *QSvgWidget) DisconnectLoad2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "load2")
	}
}

func (ptr *QSvgWidget) Load2(contents core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_Load2(ptr.Pointer(), core.PointerFromQByteArray(contents))
	}
}

func (ptr *QSvgWidget) Load2Default(contents core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_Load2Default(ptr.Pointer(), core.PointerFromQByteArray(contents))
	}
}

//export callbackQSvgWidget_Load
func callbackQSvgWidget_Load(ptr unsafe.Pointer, file C.struct_QtSvg_PackedString) {
	if signal := qt.GetSignal(ptr, "load"); signal != nil {
		signal.(func(string))(cGoUnpackString(file))
	} else {
		NewQSvgWidgetFromPointer(ptr).LoadDefault(cGoUnpackString(file))
	}
}

func (ptr *QSvgWidget) ConnectLoad(f func(file string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "load"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "load", func(file string) {
				signal.(func(string))(file)
				f(file)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "load", f)
		}
	}
}

func (ptr *QSvgWidget) DisconnectLoad() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "load")
	}
}

func (ptr *QSvgWidget) Load(file string) {
	if ptr.Pointer() != nil {
		var fileC *C.char
		if file != "" {
			fileC = C.CString(file)
			defer C.free(unsafe.Pointer(fileC))
		}
		C.QSvgWidget_Load(ptr.Pointer(), C.struct_QtSvg_PackedString{data: fileC, len: C.longlong(len(file))})
	}
}

func (ptr *QSvgWidget) LoadDefault(file string) {
	if ptr.Pointer() != nil {
		var fileC *C.char
		if file != "" {
			fileC = C.CString(file)
			defer C.free(unsafe.Pointer(fileC))
		}
		C.QSvgWidget_LoadDefault(ptr.Pointer(), C.struct_QtSvg_PackedString{data: fileC, len: C.longlong(len(file))})
	}
}

//export callbackQSvgWidget_PaintEvent
func callbackQSvgWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQSvgWidget_DestroyQSvgWidget
func callbackQSvgWidget_DestroyQSvgWidget(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSvgWidget"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).DestroyQSvgWidgetDefault()
	}
}

func (ptr *QSvgWidget) ConnectDestroyQSvgWidget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSvgWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSvgWidget", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSvgWidget", f)
		}
	}
}

func (ptr *QSvgWidget) DisconnectDestroyQSvgWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSvgWidget")
	}
}

func (ptr *QSvgWidget) DestroyQSvgWidget() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_DestroyQSvgWidget(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSvgWidget) DestroyQSvgWidgetDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_DestroyQSvgWidgetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSvgWidget_SizeHint
func callbackQSvgWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQSvgWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QSvgWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSvgWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) Renderer() *QSvgRenderer {
	if ptr.Pointer() != nil {
		tmpValue := NewQSvgRendererFromPointer(C.QSvgWidget_Renderer(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSvgWidget_MetaObject
func callbackQSvgWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSvgWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSvgWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSvgWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgWidget) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QSvgWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QSvgWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QSvgWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QSvgWidget) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QSvgWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QSvgWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QSvgWidget___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QSvgWidget) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QSvgWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QSvgWidget) __actions_newList() unsafe.Pointer {
	return C.QSvgWidget___actions_newList(ptr.Pointer())
}

func (ptr *QSvgWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSvgWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSvgWidget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSvgWidget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSvgWidget) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSvgWidget___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSvgWidget) __findChildren_newList2() unsafe.Pointer {
	return C.QSvgWidget___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSvgWidget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSvgWidget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSvgWidget) __findChildren_newList3() unsafe.Pointer {
	return C.QSvgWidget___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSvgWidget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSvgWidget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSvgWidget) __findChildren_newList() unsafe.Pointer {
	return C.QSvgWidget___findChildren_newList(ptr.Pointer())
}

func (ptr *QSvgWidget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSvgWidget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSvgWidget) __children_newList() unsafe.Pointer {
	return C.QSvgWidget___children_newList(ptr.Pointer())
}

//export callbackQSvgWidget_Close
func callbackQSvgWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QSvgWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSvgWidget_Event
func callbackQSvgWidget_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QSvgWidget) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSvgWidget_FocusNextPrevChild
func callbackQSvgWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QSvgWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQSvgWidget_ActionEvent
func callbackQSvgWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQSvgWidget_ChangeEvent
func callbackQSvgWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSvgWidget_CloseEvent
func callbackQSvgWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQSvgWidget_ContextMenuEvent
func callbackQSvgWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQSvgWidget_CustomContextMenuRequested
func callbackQSvgWidget_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQSvgWidget_DragEnterEvent
func callbackQSvgWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQSvgWidget_DragLeaveEvent
func callbackQSvgWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQSvgWidget_DragMoveEvent
func callbackQSvgWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQSvgWidget_DropEvent
func callbackQSvgWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQSvgWidget_EnterEvent
func callbackQSvgWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSvgWidget_FocusInEvent
func callbackQSvgWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQSvgWidget_FocusOutEvent
func callbackQSvgWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQSvgWidget_Hide
func callbackQSvgWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QSvgWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_HideEvent
func callbackQSvgWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQSvgWidget_InputMethodEvent
func callbackQSvgWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQSvgWidget_KeyPressEvent
func callbackQSvgWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQSvgWidget_KeyReleaseEvent
func callbackQSvgWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQSvgWidget_LeaveEvent
func callbackQSvgWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSvgWidget_Lower
func callbackQSvgWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QSvgWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_MouseDoubleClickEvent
func callbackQSvgWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQSvgWidget_MouseMoveEvent
func callbackQSvgWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQSvgWidget_MousePressEvent
func callbackQSvgWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQSvgWidget_MouseReleaseEvent
func callbackQSvgWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQSvgWidget_MoveEvent
func callbackQSvgWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQSvgWidget_Raise
func callbackQSvgWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QSvgWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_Repaint
func callbackQSvgWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QSvgWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_ResizeEvent
func callbackQSvgWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQSvgWidget_SetDisabled
func callbackQSvgWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQSvgWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QSvgWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQSvgWidget_SetEnabled
func callbackQSvgWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQSvgWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QSvgWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQSvgWidget_SetFocus2
func callbackQSvgWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QSvgWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQSvgWidget_SetHidden
func callbackQSvgWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQSvgWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QSvgWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQSvgWidget_SetStyleSheet
func callbackQSvgWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtSvg_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQSvgWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QSvgWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QSvgWidget_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtSvg_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQSvgWidget_SetVisible
func callbackQSvgWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQSvgWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QSvgWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQSvgWidget_SetWindowModified
func callbackQSvgWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQSvgWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QSvgWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQSvgWidget_SetWindowTitle
func callbackQSvgWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtSvg_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQSvgWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QSvgWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QSvgWidget_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtSvg_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQSvgWidget_Show
func callbackQSvgWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QSvgWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_ShowEvent
func callbackQSvgWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQSvgWidget_ShowFullScreen
func callbackQSvgWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QSvgWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_ShowMaximized
func callbackQSvgWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QSvgWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_ShowMinimized
func callbackQSvgWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QSvgWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_ShowNormal
func callbackQSvgWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QSvgWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_TabletEvent
func callbackQSvgWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQSvgWidget_Update
func callbackQSvgWidget_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QSvgWidget) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_UpdateMicroFocus
func callbackQSvgWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QSvgWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_WheelEvent
func callbackQSvgWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQSvgWidget_WindowIconChanged
func callbackQSvgWidget_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQSvgWidget_WindowTitleChanged
func callbackQSvgWidget_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtSvg_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQSvgWidget_PaintEngine
func callbackQSvgWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQSvgWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QSvgWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QSvgWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSvgWidget_MinimumSizeHint
func callbackQSvgWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQSvgWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QSvgWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSvgWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSvgWidget_InputMethodQuery
func callbackQSvgWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQSvgWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QSvgWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSvgWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSvgWidget_HasHeightForWidth
func callbackQSvgWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QSvgWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgWidget_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSvgWidget_HeightForWidth
func callbackQSvgWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQSvgWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QSvgWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSvgWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQSvgWidget_Metric
func callbackQSvgWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQSvgWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QSvgWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSvgWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQSvgWidget_InitPainter
func callbackQSvgWidget_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQSvgWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QSvgWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQSvgWidget_EventFilter
func callbackQSvgWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSvgWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSvgWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSvgWidget_ChildEvent
func callbackQSvgWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSvgWidget_ConnectNotify
func callbackQSvgWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSvgWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSvgWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSvgWidget_CustomEvent
func callbackQSvgWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSvgWidget_DeleteLater
func callbackQSvgWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSvgWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSvgWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSvgWidget_Destroyed
func callbackQSvgWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSvgWidget_DisconnectNotify
func callbackQSvgWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSvgWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSvgWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSvgWidget_ObjectNameChanged
func callbackQSvgWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSvg_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSvgWidget_TimerEvent
func callbackQSvgWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}
