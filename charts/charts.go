// +build !minimal

package charts

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "charts.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtCharts_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtCharts_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QAbstractAxis struct {
	core.QObject
}

type QAbstractAxis_ITF interface {
	core.QObject_ITF
	QAbstractAxis_PTR() *QAbstractAxis
}

func (ptr *QAbstractAxis) QAbstractAxis_PTR() *QAbstractAxis {
	return ptr
}

func (ptr *QAbstractAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractAxis(ptr QAbstractAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractAxis_PTR().Pointer()
	}
	return nil
}

func NewQAbstractAxisFromPointer(ptr unsafe.Pointer) (n *QAbstractAxis) {
	n = new(QAbstractAxis)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAbstractAxis__AxisType
//QAbstractAxis::AxisType
type QAbstractAxis__AxisType int64

const (
	QAbstractAxis__AxisTypeNoAxis      QAbstractAxis__AxisType = QAbstractAxis__AxisType(0x0)
	QAbstractAxis__AxisTypeValue       QAbstractAxis__AxisType = QAbstractAxis__AxisType(0x1)
	QAbstractAxis__AxisTypeBarCategory QAbstractAxis__AxisType = QAbstractAxis__AxisType(0x2)
	QAbstractAxis__AxisTypeCategory    QAbstractAxis__AxisType = QAbstractAxis__AxisType(0x4)
	QAbstractAxis__AxisTypeDateTime    QAbstractAxis__AxisType = QAbstractAxis__AxisType(0x8)
	QAbstractAxis__AxisTypeLogValue    QAbstractAxis__AxisType = QAbstractAxis__AxisType(0x10)
)

func (ptr *QAbstractAxis) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QAbstractAxis_Alignment(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractAxis_ColorChanged
func callbackQAbstractAxis_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "colorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QAbstractAxis) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "colorChanged") {
			C.QAbstractAxis_ConnectColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "colorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "colorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "colorChanged")
	}
}

func (ptr *QAbstractAxis) ColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) GridLineColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QAbstractAxis_GridLineColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_GridLineColorChanged
func callbackQAbstractAxis_GridLineColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "gridLineColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstractAxis) ConnectGridLineColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gridLineColorChanged") {
			C.QAbstractAxis_ConnectGridLineColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "gridLineColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gridLineColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "gridLineColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gridLineColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectGridLineColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectGridLineColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gridLineColorChanged")
	}
}

func (ptr *QAbstractAxis) GridLineColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_GridLineColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) GridLinePen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QAbstractAxis_GridLinePen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_GridLinePenChanged
func callbackQAbstractAxis_GridLinePenChanged(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "gridLinePenChanged"); signal != nil {
		(*(*func(*gui.QPen))(signal))(gui.NewQPenFromPointer(pen))
	}

}

func (ptr *QAbstractAxis) ConnectGridLinePenChanged(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gridLinePenChanged") {
			C.QAbstractAxis_ConnectGridLinePenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "gridLinePenChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gridLinePenChanged"); signal != nil {
			f := func(pen *gui.QPen) {
				(*(*func(*gui.QPen))(signal))(pen)
				f(pen)
			}
			qt.ConnectSignal(ptr.Pointer(), "gridLinePenChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gridLinePenChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectGridLinePenChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectGridLinePenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gridLinePenChanged")
	}
}

func (ptr *QAbstractAxis) GridLinePenChanged(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_GridLinePenChanged(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

//export callbackQAbstractAxis_GridVisibleChanged
func callbackQAbstractAxis_GridVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "gridVisibleChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectGridVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gridVisibleChanged") {
			C.QAbstractAxis_ConnectGridVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "gridVisibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gridVisibleChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "gridVisibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gridVisibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectGridVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectGridVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gridVisibleChanged")
	}
}

func (ptr *QAbstractAxis) GridVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_GridVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) Hide() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_Hide(ptr.Pointer())
	}
}

func (ptr *QAbstractAxis) IsGridLineVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_IsGridLineVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractAxis) IsLineVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_IsLineVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractAxis) IsMinorGridLineVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_IsMinorGridLineVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractAxis) IsReverse() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_IsReverse(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractAxis) IsTitleVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_IsTitleVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractAxis) IsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_IsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractAxis) LabelsAngle() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractAxis_LabelsAngle(ptr.Pointer())))
	}
	return 0
}

//export callbackQAbstractAxis_LabelsAngleChanged
func callbackQAbstractAxis_LabelsAngleChanged(ptr unsafe.Pointer, angle C.int) {
	if signal := qt.GetSignal(ptr, "labelsAngleChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(angle)))
	}

}

func (ptr *QAbstractAxis) ConnectLabelsAngleChanged(f func(angle int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsAngleChanged") {
			C.QAbstractAxis_ConnectLabelsAngleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsAngleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsAngleChanged"); signal != nil {
			f := func(angle int) {
				(*(*func(int))(signal))(angle)
				f(angle)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsAngleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsAngleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsAngleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsAngleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsAngleChanged")
	}
}

func (ptr *QAbstractAxis) LabelsAngleChanged(angle int) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsAngleChanged(ptr.Pointer(), C.int(int32(angle)))
	}
}

func (ptr *QAbstractAxis) LabelsBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QAbstractAxis_LabelsBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_LabelsBrushChanged
func callbackQAbstractAxis_LabelsBrushChanged(ptr unsafe.Pointer, brush unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelsBrushChanged"); signal != nil {
		(*(*func(*gui.QBrush))(signal))(gui.NewQBrushFromPointer(brush))
	}

}

func (ptr *QAbstractAxis) ConnectLabelsBrushChanged(f func(brush *gui.QBrush)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsBrushChanged") {
			C.QAbstractAxis_ConnectLabelsBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsBrushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsBrushChanged"); signal != nil {
			f := func(brush *gui.QBrush) {
				(*(*func(*gui.QBrush))(signal))(brush)
				f(brush)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsBrushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsBrushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsBrushChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsBrushChanged")
	}
}

func (ptr *QAbstractAxis) LabelsBrushChanged(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsBrushChanged(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QAbstractAxis) LabelsColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QAbstractAxis_LabelsColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_LabelsColorChanged
func callbackQAbstractAxis_LabelsColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelsColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QAbstractAxis) ConnectLabelsColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsColorChanged") {
			C.QAbstractAxis_ConnectLabelsColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsColorChanged")
	}
}

func (ptr *QAbstractAxis) LabelsColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) LabelsEditable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_LabelsEditable(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAbstractAxis_LabelsEditableChanged
func callbackQAbstractAxis_LabelsEditableChanged(ptr unsafe.Pointer, editable C.char) {
	if signal := qt.GetSignal(ptr, "labelsEditableChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(editable) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectLabelsEditableChanged(f func(editable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsEditableChanged") {
			C.QAbstractAxis_ConnectLabelsEditableChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsEditableChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsEditableChanged"); signal != nil {
			f := func(editable bool) {
				(*(*func(bool))(signal))(editable)
				f(editable)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsEditableChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsEditableChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsEditableChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsEditableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsEditableChanged")
	}
}

func (ptr *QAbstractAxis) LabelsEditableChanged(editable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsEditableChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(editable))))
	}
}

func (ptr *QAbstractAxis) LabelsFont() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QAbstractAxis_LabelsFont(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_LabelsFontChanged
func callbackQAbstractAxis_LabelsFontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelsFontChanged"); signal != nil {
		(*(*func(*gui.QFont))(signal))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *QAbstractAxis) ConnectLabelsFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsFontChanged") {
			C.QAbstractAxis_ConnectLabelsFontChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsFontChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsFontChanged"); signal != nil {
			f := func(font *gui.QFont) {
				(*(*func(*gui.QFont))(signal))(font)
				f(font)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsFontChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsFontChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsFontChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsFontChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsFontChanged")
	}
}

func (ptr *QAbstractAxis) LabelsFontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsFontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QAbstractAxis) LabelsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_LabelsVisible(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAbstractAxis_LabelsVisibleChanged
func callbackQAbstractAxis_LabelsVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "labelsVisibleChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectLabelsVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsVisibleChanged") {
			C.QAbstractAxis_ConnectLabelsVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsVisibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsVisibleChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsVisibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsVisibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsVisibleChanged")
	}
}

func (ptr *QAbstractAxis) LabelsVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) LinePen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QAbstractAxis_LinePen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_LinePenChanged
func callbackQAbstractAxis_LinePenChanged(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "linePenChanged"); signal != nil {
		(*(*func(*gui.QPen))(signal))(gui.NewQPenFromPointer(pen))
	}

}

func (ptr *QAbstractAxis) ConnectLinePenChanged(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linePenChanged") {
			C.QAbstractAxis_ConnectLinePenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "linePenChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linePenChanged"); signal != nil {
			f := func(pen *gui.QPen) {
				(*(*func(*gui.QPen))(signal))(pen)
				f(pen)
			}
			qt.ConnectSignal(ptr.Pointer(), "linePenChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linePenChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectLinePenChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLinePenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "linePenChanged")
	}
}

func (ptr *QAbstractAxis) LinePenChanged(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LinePenChanged(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QAbstractAxis) LinePenColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QAbstractAxis_LinePenColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_LineVisibleChanged
func callbackQAbstractAxis_LineVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "lineVisibleChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectLineVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lineVisibleChanged") {
			C.QAbstractAxis_ConnectLineVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lineVisibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lineVisibleChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "lineVisibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lineVisibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectLineVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLineVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lineVisibleChanged")
	}
}

func (ptr *QAbstractAxis) LineVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LineVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) MinorGridLineColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QAbstractAxis_MinorGridLineColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_MinorGridLineColorChanged
func callbackQAbstractAxis_MinorGridLineColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "minorGridLineColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstractAxis) ConnectMinorGridLineColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minorGridLineColorChanged") {
			C.QAbstractAxis_ConnectMinorGridLineColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minorGridLineColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minorGridLineColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "minorGridLineColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minorGridLineColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectMinorGridLineColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectMinorGridLineColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minorGridLineColorChanged")
	}
}

func (ptr *QAbstractAxis) MinorGridLineColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_MinorGridLineColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) MinorGridLinePen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QAbstractAxis_MinorGridLinePen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_MinorGridLinePenChanged
func callbackQAbstractAxis_MinorGridLinePenChanged(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "minorGridLinePenChanged"); signal != nil {
		(*(*func(*gui.QPen))(signal))(gui.NewQPenFromPointer(pen))
	}

}

func (ptr *QAbstractAxis) ConnectMinorGridLinePenChanged(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minorGridLinePenChanged") {
			C.QAbstractAxis_ConnectMinorGridLinePenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minorGridLinePenChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minorGridLinePenChanged"); signal != nil {
			f := func(pen *gui.QPen) {
				(*(*func(*gui.QPen))(signal))(pen)
				f(pen)
			}
			qt.ConnectSignal(ptr.Pointer(), "minorGridLinePenChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minorGridLinePenChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectMinorGridLinePenChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectMinorGridLinePenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minorGridLinePenChanged")
	}
}

func (ptr *QAbstractAxis) MinorGridLinePenChanged(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_MinorGridLinePenChanged(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

//export callbackQAbstractAxis_MinorGridVisibleChanged
func callbackQAbstractAxis_MinorGridVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "minorGridVisibleChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectMinorGridVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minorGridVisibleChanged") {
			C.QAbstractAxis_ConnectMinorGridVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minorGridVisibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minorGridVisibleChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "minorGridVisibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minorGridVisibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectMinorGridVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectMinorGridVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minorGridVisibleChanged")
	}
}

func (ptr *QAbstractAxis) MinorGridVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_MinorGridVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QAbstractAxis_Orientation(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractAxis_ReverseChanged
func callbackQAbstractAxis_ReverseChanged(ptr unsafe.Pointer, reverse C.char) {
	if signal := qt.GetSignal(ptr, "reverseChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(reverse) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectReverseChanged(f func(reverse bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "reverseChanged") {
			C.QAbstractAxis_ConnectReverseChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "reverseChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "reverseChanged"); signal != nil {
			f := func(reverse bool) {
				(*(*func(bool))(signal))(reverse)
				f(reverse)
			}
			qt.ConnectSignal(ptr.Pointer(), "reverseChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reverseChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectReverseChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectReverseChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "reverseChanged")
	}
}

func (ptr *QAbstractAxis) ReverseChanged(reverse bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ReverseChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(reverse))))
	}
}

func (ptr *QAbstractAxis) SetGridLineColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetGridLineColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) SetGridLinePen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetGridLinePen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QAbstractAxis) SetGridLineVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetGridLineVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) SetLabelsAngle(angle int) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetLabelsAngle(ptr.Pointer(), C.int(int32(angle)))
	}
}

func (ptr *QAbstractAxis) SetLabelsBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetLabelsBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QAbstractAxis) SetLabelsColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetLabelsColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) SetLabelsEditable(editable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetLabelsEditable(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(editable))))
	}
}

func (ptr *QAbstractAxis) SetLabelsFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetLabelsFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QAbstractAxis) SetLabelsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetLabelsVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) SetLinePen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetLinePen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QAbstractAxis) SetLinePenColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetLinePenColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) SetLineVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetLineVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) SetMax(max core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetMax(ptr.Pointer(), core.PointerFromQVariant(max))
	}
}

func (ptr *QAbstractAxis) SetMin(min core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetMin(ptr.Pointer(), core.PointerFromQVariant(min))
	}
}

func (ptr *QAbstractAxis) SetMinorGridLineColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetMinorGridLineColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) SetMinorGridLinePen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetMinorGridLinePen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QAbstractAxis) SetMinorGridLineVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetMinorGridLineVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) SetRange(min core.QVariant_ITF, max core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetRange(ptr.Pointer(), core.PointerFromQVariant(min), core.PointerFromQVariant(max))
	}
}

func (ptr *QAbstractAxis) SetReverse(reverse bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetReverse(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(reverse))))
	}
}

func (ptr *QAbstractAxis) SetShadesBorderColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetShadesBorderColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) SetShadesBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetShadesBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QAbstractAxis) SetShadesColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetShadesColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) SetShadesPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetShadesPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QAbstractAxis) SetShadesVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetShadesVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) SetTitleBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetTitleBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QAbstractAxis) SetTitleFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetTitleFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QAbstractAxis) SetTitleText(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QAbstractAxis_SetTitleText(ptr.Pointer(), C.struct_QtCharts_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *QAbstractAxis) SetTitleVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetTitleVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) ShadesBorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QAbstractAxis_ShadesBorderColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_ShadesBorderColorChanged
func callbackQAbstractAxis_ShadesBorderColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "shadesBorderColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QAbstractAxis) ConnectShadesBorderColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "shadesBorderColorChanged") {
			C.QAbstractAxis_ConnectShadesBorderColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "shadesBorderColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "shadesBorderColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "shadesBorderColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shadesBorderColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectShadesBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectShadesBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "shadesBorderColorChanged")
	}
}

func (ptr *QAbstractAxis) ShadesBorderColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ShadesBorderColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) ShadesBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QAbstractAxis_ShadesBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_ShadesBrushChanged
func callbackQAbstractAxis_ShadesBrushChanged(ptr unsafe.Pointer, brush unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "shadesBrushChanged"); signal != nil {
		(*(*func(*gui.QBrush))(signal))(gui.NewQBrushFromPointer(brush))
	}

}

func (ptr *QAbstractAxis) ConnectShadesBrushChanged(f func(brush *gui.QBrush)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "shadesBrushChanged") {
			C.QAbstractAxis_ConnectShadesBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "shadesBrushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "shadesBrushChanged"); signal != nil {
			f := func(brush *gui.QBrush) {
				(*(*func(*gui.QBrush))(signal))(brush)
				f(brush)
			}
			qt.ConnectSignal(ptr.Pointer(), "shadesBrushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shadesBrushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectShadesBrushChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectShadesBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "shadesBrushChanged")
	}
}

func (ptr *QAbstractAxis) ShadesBrushChanged(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ShadesBrushChanged(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QAbstractAxis) ShadesColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QAbstractAxis_ShadesColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_ShadesColorChanged
func callbackQAbstractAxis_ShadesColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "shadesColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QAbstractAxis) ConnectShadesColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "shadesColorChanged") {
			C.QAbstractAxis_ConnectShadesColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "shadesColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "shadesColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "shadesColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shadesColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectShadesColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectShadesColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "shadesColorChanged")
	}
}

func (ptr *QAbstractAxis) ShadesColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ShadesColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstractAxis) ShadesPen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QAbstractAxis_ShadesPen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_ShadesPenChanged
func callbackQAbstractAxis_ShadesPenChanged(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "shadesPenChanged"); signal != nil {
		(*(*func(*gui.QPen))(signal))(gui.NewQPenFromPointer(pen))
	}

}

func (ptr *QAbstractAxis) ConnectShadesPenChanged(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "shadesPenChanged") {
			C.QAbstractAxis_ConnectShadesPenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "shadesPenChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "shadesPenChanged"); signal != nil {
			f := func(pen *gui.QPen) {
				(*(*func(*gui.QPen))(signal))(pen)
				f(pen)
			}
			qt.ConnectSignal(ptr.Pointer(), "shadesPenChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shadesPenChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectShadesPenChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectShadesPenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "shadesPenChanged")
	}
}

func (ptr *QAbstractAxis) ShadesPenChanged(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ShadesPenChanged(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QAbstractAxis) ShadesVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_ShadesVisible(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAbstractAxis_ShadesVisibleChanged
func callbackQAbstractAxis_ShadesVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "shadesVisibleChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectShadesVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "shadesVisibleChanged") {
			C.QAbstractAxis_ConnectShadesVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "shadesVisibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "shadesVisibleChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "shadesVisibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shadesVisibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectShadesVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectShadesVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "shadesVisibleChanged")
	}
}

func (ptr *QAbstractAxis) ShadesVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ShadesVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) Show() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_Show(ptr.Pointer())
	}
}

func (ptr *QAbstractAxis) TitleBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QAbstractAxis_TitleBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_TitleBrushChanged
func callbackQAbstractAxis_TitleBrushChanged(ptr unsafe.Pointer, brush unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "titleBrushChanged"); signal != nil {
		(*(*func(*gui.QBrush))(signal))(gui.NewQBrushFromPointer(brush))
	}

}

func (ptr *QAbstractAxis) ConnectTitleBrushChanged(f func(brush *gui.QBrush)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleBrushChanged") {
			C.QAbstractAxis_ConnectTitleBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "titleBrushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleBrushChanged"); signal != nil {
			f := func(brush *gui.QBrush) {
				(*(*func(*gui.QBrush))(signal))(brush)
				f(brush)
			}
			qt.ConnectSignal(ptr.Pointer(), "titleBrushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleBrushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectTitleBrushChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectTitleBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleBrushChanged")
	}
}

func (ptr *QAbstractAxis) TitleBrushChanged(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_TitleBrushChanged(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QAbstractAxis) TitleFont() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QAbstractAxis_TitleFont(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_TitleFontChanged
func callbackQAbstractAxis_TitleFontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "titleFontChanged"); signal != nil {
		(*(*func(*gui.QFont))(signal))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *QAbstractAxis) ConnectTitleFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleFontChanged") {
			C.QAbstractAxis_ConnectTitleFontChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "titleFontChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleFontChanged"); signal != nil {
			f := func(font *gui.QFont) {
				(*(*func(*gui.QFont))(signal))(font)
				f(font)
			}
			qt.ConnectSignal(ptr.Pointer(), "titleFontChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleFontChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectTitleFontChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectTitleFontChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleFontChanged")
	}
}

func (ptr *QAbstractAxis) TitleFontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_TitleFontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QAbstractAxis) TitleText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractAxis_TitleText(ptr.Pointer()))
	}
	return ""
}

//export callbackQAbstractAxis_TitleTextChanged
func callbackQAbstractAxis_TitleTextChanged(ptr unsafe.Pointer, text C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "titleTextChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	}

}

func (ptr *QAbstractAxis) ConnectTitleTextChanged(f func(text string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleTextChanged") {
			C.QAbstractAxis_ConnectTitleTextChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "titleTextChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleTextChanged"); signal != nil {
			f := func(text string) {
				(*(*func(string))(signal))(text)
				f(text)
			}
			qt.ConnectSignal(ptr.Pointer(), "titleTextChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleTextChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectTitleTextChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectTitleTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleTextChanged")
	}
}

func (ptr *QAbstractAxis) TitleTextChanged(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QAbstractAxis_TitleTextChanged(ptr.Pointer(), C.struct_QtCharts_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQAbstractAxis_TitleVisibleChanged
func callbackQAbstractAxis_TitleVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "titleVisibleChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectTitleVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleVisibleChanged") {
			C.QAbstractAxis_ConnectTitleVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "titleVisibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleVisibleChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "titleVisibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleVisibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectTitleVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectTitleVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleVisibleChanged")
	}
}

func (ptr *QAbstractAxis) TitleVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_TitleVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstractAxis_Type
func callbackQAbstractAxis_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractAxis__AxisType)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QAbstractAxis) ConnectType(f func() QAbstractAxis__AxisType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractAxis__AxisType {
				(*(*func() QAbstractAxis__AxisType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QAbstractAxis) Type() QAbstractAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstractAxis__AxisType(C.QAbstractAxis_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractAxis_VisibleChanged
func callbackQAbstractAxis_VisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "visibleChanged") {
			C.QAbstractAxis_ConnectVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "visibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "visibleChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "visibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "visibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "visibleChanged")
	}
}

func (ptr *QAbstractAxis) VisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_VisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstractAxis_DestroyQAbstractAxis
func callbackQAbstractAxis_DestroyQAbstractAxis(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractAxis"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractAxisFromPointer(ptr).DestroyQAbstractAxisDefault()
	}
}

func (ptr *QAbstractAxis) ConnectDestroyQAbstractAxis(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractAxis"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractAxis", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractAxis", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractAxis) DisconnectDestroyQAbstractAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractAxis")
	}
}

func (ptr *QAbstractAxis) DestroyQAbstractAxis() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractAxis_DestroyQAbstractAxis(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractAxis) DestroyQAbstractAxisDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractAxis_DestroyQAbstractAxisDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractAxis) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractAxis___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractAxis) __children_newList() unsafe.Pointer {
	return C.QAbstractAxis___children_newList(ptr.Pointer())
}

func (ptr *QAbstractAxis) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAbstractAxis___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QAbstractAxis) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QAbstractAxis___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QAbstractAxis) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractAxis___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractAxis) __findChildren_newList() unsafe.Pointer {
	return C.QAbstractAxis___findChildren_newList(ptr.Pointer())
}

func (ptr *QAbstractAxis) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractAxis___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractAxis) __findChildren_newList3() unsafe.Pointer {
	return C.QAbstractAxis___findChildren_newList3(ptr.Pointer())
}

//export callbackQAbstractAxis_ChildEvent
func callbackQAbstractAxis_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractAxisFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractAxis) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractAxis_ConnectNotify
func callbackQAbstractAxis_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractAxisFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractAxis) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractAxis_CustomEvent
func callbackQAbstractAxis_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractAxisFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractAxis) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractAxis_DeleteLater
func callbackQAbstractAxis_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractAxisFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractAxis) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractAxis_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQAbstractAxis_Destroyed
func callbackQAbstractAxis_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQAbstractAxis_DisconnectNotify
func callbackQAbstractAxis_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractAxisFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractAxis) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractAxis_Event
func callbackQAbstractAxis_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractAxisFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractAxis) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQAbstractAxis_EventFilter
func callbackQAbstractAxis_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractAxisFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractAxis) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractAxis_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAbstractAxis_MetaObject
func callbackQAbstractAxis_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQAbstractAxisFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractAxis) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractAxis_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQAbstractAxis_ObjectNameChanged
func callbackQAbstractAxis_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQAbstractAxis_TimerEvent
func callbackQAbstractAxis_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractAxisFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractAxis) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QAbstractBarSeries struct {
	QAbstractSeries
}

type QAbstractBarSeries_ITF interface {
	QAbstractSeries_ITF
	QAbstractBarSeries_PTR() *QAbstractBarSeries
}

func (ptr *QAbstractBarSeries) QAbstractBarSeries_PTR() *QAbstractBarSeries {
	return ptr
}

func (ptr *QAbstractBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSeries_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractBarSeries(ptr QAbstractBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQAbstractBarSeriesFromPointer(ptr unsafe.Pointer) (n *QAbstractBarSeries) {
	n = new(QAbstractBarSeries)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAbstractBarSeries__LabelsPosition
//QAbstractBarSeries::LabelsPosition
type QAbstractBarSeries__LabelsPosition int64

const (
	QAbstractBarSeries__LabelsCenter     QAbstractBarSeries__LabelsPosition = QAbstractBarSeries__LabelsPosition(0)
	QAbstractBarSeries__LabelsInsideEnd  QAbstractBarSeries__LabelsPosition = QAbstractBarSeries__LabelsPosition(1)
	QAbstractBarSeries__LabelsInsideBase QAbstractBarSeries__LabelsPosition = QAbstractBarSeries__LabelsPosition(2)
	QAbstractBarSeries__LabelsOutsideEnd QAbstractBarSeries__LabelsPosition = QAbstractBarSeries__LabelsPosition(3)
)

func (ptr *QAbstractBarSeries) Append(set QBarSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractBarSeries_Append(ptr.Pointer(), PointerFromQBarSet(set))) != 0
	}
	return false
}

func (ptr *QAbstractBarSeries) Append2(sets []*QBarSet) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractBarSeries_Append2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQAbstractBarSeriesFromPointer(NewQAbstractBarSeriesFromPointer(nil).__append_sets_newList2())
			for _, v := range sets {
				tmpList.__append_sets_setList2(v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

func (ptr *QAbstractBarSeries) BarSets() []*QBarSet {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QBarSet {
			out := make([]*QBarSet, int(l.len))
			tmpList := NewQAbstractBarSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__barSets_atList(i)
			}
			return out
		}(C.QAbstractBarSeries_BarSets(ptr.Pointer()))
	}
	return make([]*QBarSet, 0)
}

func (ptr *QAbstractBarSeries) BarWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstractBarSeries_BarWidth(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractBarSeries_BarsetsAdded
func callbackQAbstractBarSeries_BarsetsAdded(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(ptr, "barsetsAdded"); signal != nil {
		(*(*func([]*QBarSet))(signal))(func(l C.struct_QtCharts_PackedList) []*QBarSet {
			out := make([]*QBarSet, int(l.len))
			tmpList := NewQAbstractBarSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__barsetsAdded_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QAbstractBarSeries) ConnectBarsetsAdded(f func(sets []*QBarSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "barsetsAdded") {
			C.QAbstractBarSeries_ConnectBarsetsAdded(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "barsetsAdded")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "barsetsAdded"); signal != nil {
			f := func(sets []*QBarSet) {
				(*(*func([]*QBarSet))(signal))(sets)
				f(sets)
			}
			qt.ConnectSignal(ptr.Pointer(), "barsetsAdded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "barsetsAdded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectBarsetsAdded() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectBarsetsAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "barsetsAdded")
	}
}

func (ptr *QAbstractBarSeries) BarsetsAdded(sets []*QBarSet) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_BarsetsAdded(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQAbstractBarSeriesFromPointer(NewQAbstractBarSeriesFromPointer(nil).__barsetsAdded_sets_newList())
			for _, v := range sets {
				tmpList.__barsetsAdded_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQAbstractBarSeries_BarsetsRemoved
func callbackQAbstractBarSeries_BarsetsRemoved(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(ptr, "barsetsRemoved"); signal != nil {
		(*(*func([]*QBarSet))(signal))(func(l C.struct_QtCharts_PackedList) []*QBarSet {
			out := make([]*QBarSet, int(l.len))
			tmpList := NewQAbstractBarSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__barsetsRemoved_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QAbstractBarSeries) ConnectBarsetsRemoved(f func(sets []*QBarSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "barsetsRemoved") {
			C.QAbstractBarSeries_ConnectBarsetsRemoved(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "barsetsRemoved")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "barsetsRemoved"); signal != nil {
			f := func(sets []*QBarSet) {
				(*(*func([]*QBarSet))(signal))(sets)
				f(sets)
			}
			qt.ConnectSignal(ptr.Pointer(), "barsetsRemoved", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "barsetsRemoved", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectBarsetsRemoved() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectBarsetsRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "barsetsRemoved")
	}
}

func (ptr *QAbstractBarSeries) BarsetsRemoved(sets []*QBarSet) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_BarsetsRemoved(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQAbstractBarSeriesFromPointer(NewQAbstractBarSeriesFromPointer(nil).__barsetsRemoved_sets_newList())
			for _, v := range sets {
				tmpList.__barsetsRemoved_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QAbstractBarSeries) Clear() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_Clear(ptr.Pointer())
	}
}

//export callbackQAbstractBarSeries_Clicked
func callbackQAbstractBarSeries_Clicked(ptr unsafe.Pointer, index C.int, barset unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(int, *QBarSet))(signal))(int(int32(index)), NewQBarSetFromPointer(barset))
	}

}

func (ptr *QAbstractBarSeries) ConnectClicked(f func(index int, barset *QBarSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QAbstractBarSeries_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(index int, barset *QBarSet) {
				(*(*func(int, *QBarSet))(signal))(index, barset)
				f(index, barset)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QAbstractBarSeries) Clicked(index int, barset QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_Clicked(ptr.Pointer(), C.int(int32(index)), PointerFromQBarSet(barset))
	}
}

func (ptr *QAbstractBarSeries) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractBarSeries_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQAbstractBarSeries_CountChanged
func callbackQAbstractBarSeries_CountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "countChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractBarSeries) ConnectCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "countChanged") {
			C.QAbstractBarSeries_ConnectCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "countChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "countChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "countChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "countChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectCountChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "countChanged")
	}
}

func (ptr *QAbstractBarSeries) CountChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_CountChanged(ptr.Pointer())
	}
}

//export callbackQAbstractBarSeries_DoubleClicked
func callbackQAbstractBarSeries_DoubleClicked(ptr unsafe.Pointer, index C.int, barset unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		(*(*func(int, *QBarSet))(signal))(int(int32(index)), NewQBarSetFromPointer(barset))
	}

}

func (ptr *QAbstractBarSeries) ConnectDoubleClicked(f func(index int, barset *QBarSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "doubleClicked") {
			C.QAbstractBarSeries_ConnectDoubleClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "doubleClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "doubleClicked"); signal != nil {
			f := func(index int, barset *QBarSet) {
				(*(*func(int, *QBarSet))(signal))(index, barset)
				f(index, barset)
			}
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "doubleClicked")
	}
}

func (ptr *QAbstractBarSeries) DoubleClicked(index int, barset QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DoubleClicked(ptr.Pointer(), C.int(int32(index)), PointerFromQBarSet(barset))
	}
}

//export callbackQAbstractBarSeries_Hovered
func callbackQAbstractBarSeries_Hovered(ptr unsafe.Pointer, status C.char, index C.int, barset unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(bool, int, *QBarSet))(signal))(int8(status) != 0, int(int32(index)), NewQBarSetFromPointer(barset))
	}

}

func (ptr *QAbstractBarSeries) ConnectHovered(f func(status bool, index int, barset *QBarSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QAbstractBarSeries_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(status bool, index int, barset *QBarSet) {
				(*(*func(bool, int, *QBarSet))(signal))(status, index, barset)
				f(status, index, barset)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QAbstractBarSeries) Hovered(status bool, index int, barset QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))), C.int(int32(index)), PointerFromQBarSet(barset))
	}
}

func (ptr *QAbstractBarSeries) Insert(index int, set QBarSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractBarSeries_Insert(ptr.Pointer(), C.int(int32(index)), PointerFromQBarSet(set))) != 0
	}
	return false
}

func (ptr *QAbstractBarSeries) IsLabelsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractBarSeries_IsLabelsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractBarSeries) LabelsAngle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstractBarSeries_LabelsAngle(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractBarSeries_LabelsAngleChanged
func callbackQAbstractBarSeries_LabelsAngleChanged(ptr unsafe.Pointer, angle C.double) {
	if signal := qt.GetSignal(ptr, "labelsAngleChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(angle))
	}

}

func (ptr *QAbstractBarSeries) ConnectLabelsAngleChanged(f func(angle float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsAngleChanged") {
			C.QAbstractBarSeries_ConnectLabelsAngleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsAngleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsAngleChanged"); signal != nil {
			f := func(angle float64) {
				(*(*func(float64))(signal))(angle)
				f(angle)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsAngleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsAngleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectLabelsAngleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectLabelsAngleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsAngleChanged")
	}
}

func (ptr *QAbstractBarSeries) LabelsAngleChanged(angle float64) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_LabelsAngleChanged(ptr.Pointer(), C.double(angle))
	}
}

func (ptr *QAbstractBarSeries) LabelsFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractBarSeries_LabelsFormat(ptr.Pointer()))
	}
	return ""
}

//export callbackQAbstractBarSeries_LabelsFormatChanged
func callbackQAbstractBarSeries_LabelsFormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "labelsFormatChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(format))
	}

}

func (ptr *QAbstractBarSeries) ConnectLabelsFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsFormatChanged") {
			C.QAbstractBarSeries_ConnectLabelsFormatChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsFormatChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsFormatChanged"); signal != nil {
			f := func(format string) {
				(*(*func(string))(signal))(format)
				f(format)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsFormatChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsFormatChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectLabelsFormatChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectLabelsFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsFormatChanged")
	}
}

func (ptr *QAbstractBarSeries) LabelsFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QAbstractBarSeries_LabelsFormatChanged(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

func (ptr *QAbstractBarSeries) LabelsPosition() QAbstractBarSeries__LabelsPosition {
	if ptr.Pointer() != nil {
		return QAbstractBarSeries__LabelsPosition(C.QAbstractBarSeries_LabelsPosition(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractBarSeries_LabelsPositionChanged
func callbackQAbstractBarSeries_LabelsPositionChanged(ptr unsafe.Pointer, position C.longlong) {
	if signal := qt.GetSignal(ptr, "labelsPositionChanged"); signal != nil {
		(*(*func(QAbstractBarSeries__LabelsPosition))(signal))(QAbstractBarSeries__LabelsPosition(position))
	}

}

func (ptr *QAbstractBarSeries) ConnectLabelsPositionChanged(f func(position QAbstractBarSeries__LabelsPosition)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsPositionChanged") {
			C.QAbstractBarSeries_ConnectLabelsPositionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsPositionChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsPositionChanged"); signal != nil {
			f := func(position QAbstractBarSeries__LabelsPosition) {
				(*(*func(QAbstractBarSeries__LabelsPosition))(signal))(position)
				f(position)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsPositionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsPositionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectLabelsPositionChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectLabelsPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsPositionChanged")
	}
}

func (ptr *QAbstractBarSeries) LabelsPositionChanged(position QAbstractBarSeries__LabelsPosition) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_LabelsPositionChanged(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QAbstractBarSeries) LabelsPrecision() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractBarSeries_LabelsPrecision(ptr.Pointer())))
	}
	return 0
}

//export callbackQAbstractBarSeries_LabelsPrecisionChanged
func callbackQAbstractBarSeries_LabelsPrecisionChanged(ptr unsafe.Pointer, precision C.int) {
	if signal := qt.GetSignal(ptr, "labelsPrecisionChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(precision)))
	}

}

func (ptr *QAbstractBarSeries) ConnectLabelsPrecisionChanged(f func(precision int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsPrecisionChanged") {
			C.QAbstractBarSeries_ConnectLabelsPrecisionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsPrecisionChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsPrecisionChanged"); signal != nil {
			f := func(precision int) {
				(*(*func(int))(signal))(precision)
				f(precision)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsPrecisionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsPrecisionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectLabelsPrecisionChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectLabelsPrecisionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsPrecisionChanged")
	}
}

func (ptr *QAbstractBarSeries) LabelsPrecisionChanged(precision int) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_LabelsPrecisionChanged(ptr.Pointer(), C.int(int32(precision)))
	}
}

//export callbackQAbstractBarSeries_LabelsVisibleChanged
func callbackQAbstractBarSeries_LabelsVisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelsVisibleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractBarSeries) ConnectLabelsVisibleChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsVisibleChanged") {
			C.QAbstractBarSeries_ConnectLabelsVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsVisibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsVisibleChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsVisibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsVisibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectLabelsVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectLabelsVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsVisibleChanged")
	}
}

func (ptr *QAbstractBarSeries) LabelsVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_LabelsVisibleChanged(ptr.Pointer())
	}
}

//export callbackQAbstractBarSeries_Pressed
func callbackQAbstractBarSeries_Pressed(ptr unsafe.Pointer, index C.int, barset unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func(int, *QBarSet))(signal))(int(int32(index)), NewQBarSetFromPointer(barset))
	}

}

func (ptr *QAbstractBarSeries) ConnectPressed(f func(index int, barset *QBarSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pressed") {
			C.QAbstractBarSeries_ConnectPressed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pressed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pressed"); signal != nil {
			f := func(index int, barset *QBarSet) {
				(*(*func(int, *QBarSet))(signal))(index, barset)
				f(index, barset)
			}
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pressed")
	}
}

func (ptr *QAbstractBarSeries) Pressed(index int, barset QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_Pressed(ptr.Pointer(), C.int(int32(index)), PointerFromQBarSet(barset))
	}
}

//export callbackQAbstractBarSeries_Released
func callbackQAbstractBarSeries_Released(ptr unsafe.Pointer, index C.int, barset unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "released"); signal != nil {
		(*(*func(int, *QBarSet))(signal))(int(int32(index)), NewQBarSetFromPointer(barset))
	}

}

func (ptr *QAbstractBarSeries) ConnectReleased(f func(index int, barset *QBarSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "released") {
			C.QAbstractBarSeries_ConnectReleased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "released")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "released"); signal != nil {
			f := func(index int, barset *QBarSet) {
				(*(*func(int, *QBarSet))(signal))(index, barset)
				f(index, barset)
			}
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "released")
	}
}

func (ptr *QAbstractBarSeries) Released(index int, barset QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_Released(ptr.Pointer(), C.int(int32(index)), PointerFromQBarSet(barset))
	}
}

func (ptr *QAbstractBarSeries) Remove(set QBarSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractBarSeries_Remove(ptr.Pointer(), PointerFromQBarSet(set))) != 0
	}
	return false
}

func (ptr *QAbstractBarSeries) SetBarWidth(width float64) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_SetBarWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QAbstractBarSeries) SetLabelsAngle(angle float64) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_SetLabelsAngle(ptr.Pointer(), C.double(angle))
	}
}

func (ptr *QAbstractBarSeries) SetLabelsFormat(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QAbstractBarSeries_SetLabelsFormat(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

func (ptr *QAbstractBarSeries) SetLabelsPosition(position QAbstractBarSeries__LabelsPosition) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_SetLabelsPosition(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QAbstractBarSeries) SetLabelsPrecision(precision int) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_SetLabelsPrecision(ptr.Pointer(), C.int(int32(precision)))
	}
}

func (ptr *QAbstractBarSeries) SetLabelsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_SetLabelsVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractBarSeries) Take(set QBarSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractBarSeries_Take(ptr.Pointer(), PointerFromQBarSet(set))) != 0
	}
	return false
}

//export callbackQAbstractBarSeries_DestroyQAbstractBarSeries
func callbackQAbstractBarSeries_DestroyQAbstractBarSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractBarSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractBarSeriesFromPointer(ptr).DestroyQAbstractBarSeriesDefault()
	}
}

func (ptr *QAbstractBarSeries) ConnectDestroyQAbstractBarSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractBarSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractBarSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractBarSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractBarSeries) DisconnectDestroyQAbstractBarSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractBarSeries")
	}
}

func (ptr *QAbstractBarSeries) DestroyQAbstractBarSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractBarSeries_DestroyQAbstractBarSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractBarSeries) DestroyQAbstractBarSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractBarSeries_DestroyQAbstractBarSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractBarSeries) __append_sets_atList2(i int) *QBarSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQBarSetFromPointer(C.QAbstractBarSeries___append_sets_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractBarSeries) __append_sets_setList2(i QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries___append_sets_setList2(ptr.Pointer(), PointerFromQBarSet(i))
	}
}

func (ptr *QAbstractBarSeries) __append_sets_newList2() unsafe.Pointer {
	return C.QAbstractBarSeries___append_sets_newList2(ptr.Pointer())
}

func (ptr *QAbstractBarSeries) __barSets_atList(i int) *QBarSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQBarSetFromPointer(C.QAbstractBarSeries___barSets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractBarSeries) __barSets_setList(i QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries___barSets_setList(ptr.Pointer(), PointerFromQBarSet(i))
	}
}

func (ptr *QAbstractBarSeries) __barSets_newList() unsafe.Pointer {
	return C.QAbstractBarSeries___barSets_newList(ptr.Pointer())
}

func (ptr *QAbstractBarSeries) __barsetsAdded_sets_atList(i int) *QBarSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQBarSetFromPointer(C.QAbstractBarSeries___barsetsAdded_sets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractBarSeries) __barsetsAdded_sets_setList(i QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries___barsetsAdded_sets_setList(ptr.Pointer(), PointerFromQBarSet(i))
	}
}

func (ptr *QAbstractBarSeries) __barsetsAdded_sets_newList() unsafe.Pointer {
	return C.QAbstractBarSeries___barsetsAdded_sets_newList(ptr.Pointer())
}

func (ptr *QAbstractBarSeries) __barsetsRemoved_sets_atList(i int) *QBarSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQBarSetFromPointer(C.QAbstractBarSeries___barsetsRemoved_sets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractBarSeries) __barsetsRemoved_sets_setList(i QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries___barsetsRemoved_sets_setList(ptr.Pointer(), PointerFromQBarSet(i))
	}
}

func (ptr *QAbstractBarSeries) __barsetsRemoved_sets_newList() unsafe.Pointer {
	return C.QAbstractBarSeries___barsetsRemoved_sets_newList(ptr.Pointer())
}

//export callbackQAbstractBarSeries_Type
func callbackQAbstractBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQAbstractBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QAbstractBarSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QAbstractBarSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractBarSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QAbstractBarSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

type QAbstractSeries struct {
	core.QObject
}

type QAbstractSeries_ITF interface {
	core.QObject_ITF
	QAbstractSeries_PTR() *QAbstractSeries
}

func (ptr *QAbstractSeries) QAbstractSeries_PTR() *QAbstractSeries {
	return ptr
}

func (ptr *QAbstractSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractSeries(ptr QAbstractSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSeries_PTR().Pointer()
	}
	return nil
}

func NewQAbstractSeriesFromPointer(ptr unsafe.Pointer) (n *QAbstractSeries) {
	n = new(QAbstractSeries)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAbstractSeries__SeriesType
//QAbstractSeries::SeriesType
type QAbstractSeries__SeriesType int64

const (
	QAbstractSeries__SeriesTypeLine                 QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(0)
	QAbstractSeries__SeriesTypeArea                 QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(1)
	QAbstractSeries__SeriesTypeBar                  QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(2)
	QAbstractSeries__SeriesTypeStackedBar           QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(3)
	QAbstractSeries__SeriesTypePercentBar           QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(4)
	QAbstractSeries__SeriesTypePie                  QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(5)
	QAbstractSeries__SeriesTypeScatter              QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(6)
	QAbstractSeries__SeriesTypeSpline               QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(7)
	QAbstractSeries__SeriesTypeHorizontalBar        QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(8)
	QAbstractSeries__SeriesTypeHorizontalStackedBar QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(9)
	QAbstractSeries__SeriesTypeHorizontalPercentBar QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(10)
	QAbstractSeries__SeriesTypeBoxPlot              QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(11)
	QAbstractSeries__SeriesTypeCandlestick          QAbstractSeries__SeriesType = QAbstractSeries__SeriesType(12)
)

func (ptr *QAbstractSeries) AttachAxis(axis QAbstractAxis_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSeries_AttachAxis(ptr.Pointer(), PointerFromQAbstractAxis(axis))) != 0
	}
	return false
}

func (ptr *QAbstractSeries) AttachedAxes() []*QAbstractAxis {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QAbstractAxis {
			out := make([]*QAbstractAxis, int(l.len))
			tmpList := NewQAbstractSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__attachedAxes_atList(i)
			}
			return out
		}(C.QAbstractSeries_AttachedAxes(ptr.Pointer()))
	}
	return make([]*QAbstractAxis, 0)
}

func (ptr *QAbstractSeries) Chart() *QChart {
	if ptr.Pointer() != nil {
		tmpValue := NewQChartFromPointer(C.QAbstractSeries_Chart(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSeries) DetachAxis(axis QAbstractAxis_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSeries_DetachAxis(ptr.Pointer(), PointerFromQAbstractAxis(axis))) != 0
	}
	return false
}

func (ptr *QAbstractSeries) Hide() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_Hide(ptr.Pointer())
	}
}

func (ptr *QAbstractSeries) IsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSeries_IsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSeries) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractSeries_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQAbstractSeries_NameChanged
func callbackQAbstractSeries_NameChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "nameChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractSeries) ConnectNameChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nameChanged") {
			C.QAbstractSeries_ConnectNameChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "nameChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "nameChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "nameChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nameChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSeries) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "nameChanged")
	}
}

func (ptr *QAbstractSeries) NameChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_NameChanged(ptr.Pointer())
	}
}

func (ptr *QAbstractSeries) Opacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstractSeries_Opacity(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSeries_OpacityChanged
func callbackQAbstractSeries_OpacityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractSeries) ConnectOpacityChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "opacityChanged") {
			C.QAbstractSeries_ConnectOpacityChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "opacityChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "opacityChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "opacityChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "opacityChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSeries) DisconnectOpacityChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_DisconnectOpacityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "opacityChanged")
	}
}

func (ptr *QAbstractSeries) OpacityChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_OpacityChanged(ptr.Pointer())
	}
}

func (ptr *QAbstractSeries) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QAbstractSeries_SetName(ptr.Pointer(), C.struct_QtCharts_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QAbstractSeries) SetOpacity(opacity float64) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_SetOpacity(ptr.Pointer(), C.double(opacity))
	}
}

func (ptr *QAbstractSeries) SetUseOpenGL(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_SetUseOpenGL(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QAbstractSeries) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractSeries) Show() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_Show(ptr.Pointer())
	}
}

//export callbackQAbstractSeries_Type
func callbackQAbstractSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QAbstractSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QAbstractSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QAbstractSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSeries) UseOpenGL() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSeries_UseOpenGL(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAbstractSeries_UseOpenGLChanged
func callbackQAbstractSeries_UseOpenGLChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "useOpenGLChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractSeries) ConnectUseOpenGLChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "useOpenGLChanged") {
			C.QAbstractSeries_ConnectUseOpenGLChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "useOpenGLChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "useOpenGLChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "useOpenGLChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "useOpenGLChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSeries) DisconnectUseOpenGLChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_DisconnectUseOpenGLChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "useOpenGLChanged")
	}
}

func (ptr *QAbstractSeries) UseOpenGLChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_UseOpenGLChanged(ptr.Pointer())
	}
}

//export callbackQAbstractSeries_VisibleChanged
func callbackQAbstractSeries_VisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractSeries) ConnectVisibleChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "visibleChanged") {
			C.QAbstractSeries_ConnectVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "visibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "visibleChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "visibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "visibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSeries) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "visibleChanged")
	}
}

func (ptr *QAbstractSeries) VisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_VisibleChanged(ptr.Pointer())
	}
}

//export callbackQAbstractSeries_DestroyQAbstractSeries
func callbackQAbstractSeries_DestroyQAbstractSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractSeriesFromPointer(ptr).DestroyQAbstractSeriesDefault()
	}
}

func (ptr *QAbstractSeries) ConnectDestroyQAbstractSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSeries) DisconnectDestroyQAbstractSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractSeries")
	}
}

func (ptr *QAbstractSeries) DestroyQAbstractSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractSeries_DestroyQAbstractSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSeries) DestroyQAbstractSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractSeries_DestroyQAbstractSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSeries) __attachedAxes_atList(i int) *QAbstractAxis {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractAxisFromPointer(C.QAbstractSeries___attachedAxes_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSeries) __attachedAxes_setList(i QAbstractAxis_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries___attachedAxes_setList(ptr.Pointer(), PointerFromQAbstractAxis(i))
	}
}

func (ptr *QAbstractSeries) __attachedAxes_newList() unsafe.Pointer {
	return C.QAbstractSeries___attachedAxes_newList(ptr.Pointer())
}

func (ptr *QAbstractSeries) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractSeries___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSeries) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractSeries) __children_newList() unsafe.Pointer {
	return C.QAbstractSeries___children_newList(ptr.Pointer())
}

func (ptr *QAbstractSeries) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAbstractSeries___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSeries) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QAbstractSeries) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QAbstractSeries___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QAbstractSeries) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractSeries___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSeries) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractSeries) __findChildren_newList() unsafe.Pointer {
	return C.QAbstractSeries___findChildren_newList(ptr.Pointer())
}

func (ptr *QAbstractSeries) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractSeries___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSeries) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractSeries) __findChildren_newList3() unsafe.Pointer {
	return C.QAbstractSeries___findChildren_newList3(ptr.Pointer())
}

//export callbackQAbstractSeries_ChildEvent
func callbackQAbstractSeries_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractSeriesFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractSeries) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractSeries_ConnectNotify
func callbackQAbstractSeries_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractSeriesFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractSeries) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractSeries_CustomEvent
func callbackQAbstractSeries_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSeriesFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSeries) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractSeries_DeleteLater
func callbackQAbstractSeries_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractSeriesFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractSeries) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractSeries_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQAbstractSeries_Destroyed
func callbackQAbstractSeries_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQAbstractSeries_DisconnectNotify
func callbackQAbstractSeries_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractSeriesFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractSeries) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractSeries_Event
func callbackQAbstractSeries_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSeriesFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractSeries) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSeries_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQAbstractSeries_EventFilter
func callbackQAbstractSeries_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSeriesFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractSeries) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractSeries_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAbstractSeries_MetaObject
func callbackQAbstractSeries_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQAbstractSeriesFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractSeries) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractSeries_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQAbstractSeries_ObjectNameChanged
func callbackQAbstractSeries_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQAbstractSeries_TimerEvent
func callbackQAbstractSeries_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractSeriesFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractSeries) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QAreaLegendMarker struct {
	QLegendMarker
}

type QAreaLegendMarker_ITF interface {
	QLegendMarker_ITF
	QAreaLegendMarker_PTR() *QAreaLegendMarker
}

func (ptr *QAreaLegendMarker) QAreaLegendMarker_PTR() *QAreaLegendMarker {
	return ptr
}

func (ptr *QAreaLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLegendMarker_PTR().Pointer()
	}
	return nil
}

func (ptr *QAreaLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLegendMarker_PTR().SetPointer(p)
	}
}

func PointerFromQAreaLegendMarker(ptr QAreaLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAreaLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQAreaLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QAreaLegendMarker) {
	n = new(QAreaLegendMarker)
	n.SetPointer(ptr)
	return
}

//export callbackQAreaLegendMarker_Series
func callbackQAreaLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "series"); signal != nil {
		return PointerFromQAreaSeries((*(*func() *QAreaSeries)(signal))())
	}

	return PointerFromQAreaSeries(NewQAreaLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QAreaLegendMarker) ConnectSeries(f func() *QAreaSeries) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "series"); signal != nil {
			f := func() *QAreaSeries {
				(*(*func() *QAreaSeries)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "series")
	}
}

func (ptr *QAreaLegendMarker) Series() *QAreaSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQAreaSeriesFromPointer(C.QAreaLegendMarker_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAreaLegendMarker) SeriesDefault() *QAreaSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQAreaSeriesFromPointer(C.QAreaLegendMarker_SeriesDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAreaLegendMarker_Type
func callbackQAreaLegendMarker_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QLegendMarker__LegendMarkerType)(signal))())
	}

	return C.longlong(NewQAreaLegendMarkerFromPointer(ptr).TypeDefault())
}

func (ptr *QAreaLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QLegendMarker__LegendMarkerType {
				(*(*func() QLegendMarker__LegendMarkerType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaLegendMarker) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QAreaLegendMarker) Type() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QAreaLegendMarker_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAreaLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QAreaLegendMarker_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAreaLegendMarker_DestroyQAreaLegendMarker
func callbackQAreaLegendMarker_DestroyQAreaLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAreaLegendMarker"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAreaLegendMarkerFromPointer(ptr).DestroyQAreaLegendMarkerDefault()
	}
}

func (ptr *QAreaLegendMarker) ConnectDestroyQAreaLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAreaLegendMarker"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAreaLegendMarker", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAreaLegendMarker", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaLegendMarker) DisconnectDestroyQAreaLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAreaLegendMarker")
	}
}

func (ptr *QAreaLegendMarker) DestroyQAreaLegendMarker() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAreaLegendMarker_DestroyQAreaLegendMarker(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAreaLegendMarker) DestroyQAreaLegendMarkerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAreaLegendMarker_DestroyQAreaLegendMarkerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QAreaSeries struct {
	QAbstractSeries
}

type QAreaSeries_ITF interface {
	QAbstractSeries_ITF
	QAreaSeries_PTR() *QAreaSeries
}

func (ptr *QAreaSeries) QAreaSeries_PTR() *QAreaSeries {
	return ptr
}

func (ptr *QAreaSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QAreaSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSeries_PTR().SetPointer(p)
	}
}

func PointerFromQAreaSeries(ptr QAreaSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAreaSeries_PTR().Pointer()
	}
	return nil
}

func NewQAreaSeriesFromPointer(ptr unsafe.Pointer) (n *QAreaSeries) {
	n = new(QAreaSeries)
	n.SetPointer(ptr)
	return
}
func NewQAreaSeries(parent core.QObject_ITF) *QAreaSeries {
	tmpValue := NewQAreaSeriesFromPointer(C.QAreaSeries_NewQAreaSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQAreaSeries2(upperSeries QLineSeries_ITF, lowerSeries QLineSeries_ITF) *QAreaSeries {
	tmpValue := NewQAreaSeriesFromPointer(C.QAreaSeries_NewQAreaSeries2(PointerFromQLineSeries(upperSeries), PointerFromQLineSeries(lowerSeries)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAreaSeries) BorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QAreaSeries_BorderColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQAreaSeries_BorderColorChanged
func callbackQAreaSeries_BorderColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "borderColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QAreaSeries) ConnectBorderColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "borderColorChanged") {
			C.QAreaSeries_ConnectBorderColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "borderColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "borderColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "borderColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "borderColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "borderColorChanged")
	}
}

func (ptr *QAreaSeries) BorderColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_BorderColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAreaSeries) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QAreaSeries_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQAreaSeries_Clicked
func callbackQAreaSeries_Clicked(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(*core.QPointF))(signal))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QAreaSeries) ConnectClicked(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QAreaSeries_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(point *core.QPointF) {
				(*(*func(*core.QPointF))(signal))(point)
				f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QAreaSeries) Clicked(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_Clicked(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QAreaSeries) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QAreaSeries_Color(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQAreaSeries_ColorChanged
func callbackQAreaSeries_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "colorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QAreaSeries) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "colorChanged") {
			C.QAreaSeries_ConnectColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "colorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "colorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "colorChanged")
	}
}

func (ptr *QAreaSeries) ColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAreaSeries_DoubleClicked
func callbackQAreaSeries_DoubleClicked(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		(*(*func(*core.QPointF))(signal))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QAreaSeries) ConnectDoubleClicked(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "doubleClicked") {
			C.QAreaSeries_ConnectDoubleClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "doubleClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "doubleClicked"); signal != nil {
			f := func(point *core.QPointF) {
				(*(*func(*core.QPointF))(signal))(point)
				f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "doubleClicked")
	}
}

func (ptr *QAreaSeries) DoubleClicked(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DoubleClicked(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQAreaSeries_Hovered
func callbackQAreaSeries_Hovered(ptr unsafe.Pointer, point unsafe.Pointer, state C.char) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(*core.QPointF, bool))(signal))(core.NewQPointFFromPointer(point), int8(state) != 0)
	}

}

func (ptr *QAreaSeries) ConnectHovered(f func(point *core.QPointF, state bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QAreaSeries_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(point *core.QPointF, state bool) {
				(*(*func(*core.QPointF, bool))(signal))(point, state)
				f(point, state)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QAreaSeries) Hovered(point core.QPointF_ITF, state bool) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_Hovered(ptr.Pointer(), core.PointerFromQPointF(point), C.char(int8(qt.GoBoolToInt(state))))
	}
}

func (ptr *QAreaSeries) LowerSeries() *QLineSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQLineSeriesFromPointer(C.QAreaSeries_LowerSeries(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAreaSeries) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QAreaSeries_Pen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QAreaSeries) PointLabelsClipping() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAreaSeries_PointLabelsClipping(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAreaSeries_PointLabelsClippingChanged
func callbackQAreaSeries_PointLabelsClippingChanged(ptr unsafe.Pointer, clipping C.char) {
	if signal := qt.GetSignal(ptr, "pointLabelsClippingChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(clipping) != 0)
	}

}

func (ptr *QAreaSeries) ConnectPointLabelsClippingChanged(f func(clipping bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointLabelsClippingChanged") {
			C.QAreaSeries_ConnectPointLabelsClippingChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointLabelsClippingChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointLabelsClippingChanged"); signal != nil {
			f := func(clipping bool) {
				(*(*func(bool))(signal))(clipping)
				f(clipping)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsClippingChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsClippingChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectPointLabelsClippingChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPointLabelsClippingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointLabelsClippingChanged")
	}
}

func (ptr *QAreaSeries) PointLabelsClippingChanged(clipping bool) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_PointLabelsClippingChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(clipping))))
	}
}

func (ptr *QAreaSeries) PointLabelsColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QAreaSeries_PointLabelsColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQAreaSeries_PointLabelsColorChanged
func callbackQAreaSeries_PointLabelsColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pointLabelsColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAreaSeries) ConnectPointLabelsColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointLabelsColorChanged") {
			C.QAreaSeries_ConnectPointLabelsColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointLabelsColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointLabelsColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectPointLabelsColorChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPointLabelsColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointLabelsColorChanged")
	}
}

func (ptr *QAreaSeries) PointLabelsColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_PointLabelsColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAreaSeries) PointLabelsFont() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QAreaSeries_PointLabelsFont(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

//export callbackQAreaSeries_PointLabelsFontChanged
func callbackQAreaSeries_PointLabelsFontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pointLabelsFontChanged"); signal != nil {
		(*(*func(*gui.QFont))(signal))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *QAreaSeries) ConnectPointLabelsFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointLabelsFontChanged") {
			C.QAreaSeries_ConnectPointLabelsFontChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointLabelsFontChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointLabelsFontChanged"); signal != nil {
			f := func(font *gui.QFont) {
				(*(*func(*gui.QFont))(signal))(font)
				f(font)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsFontChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsFontChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectPointLabelsFontChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPointLabelsFontChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointLabelsFontChanged")
	}
}

func (ptr *QAreaSeries) PointLabelsFontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_PointLabelsFontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QAreaSeries) PointLabelsFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAreaSeries_PointLabelsFormat(ptr.Pointer()))
	}
	return ""
}

//export callbackQAreaSeries_PointLabelsFormatChanged
func callbackQAreaSeries_PointLabelsFormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "pointLabelsFormatChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(format))
	}

}

func (ptr *QAreaSeries) ConnectPointLabelsFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointLabelsFormatChanged") {
			C.QAreaSeries_ConnectPointLabelsFormatChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointLabelsFormatChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointLabelsFormatChanged"); signal != nil {
			f := func(format string) {
				(*(*func(string))(signal))(format)
				f(format)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsFormatChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsFormatChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectPointLabelsFormatChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPointLabelsFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointLabelsFormatChanged")
	}
}

func (ptr *QAreaSeries) PointLabelsFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QAreaSeries_PointLabelsFormatChanged(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

//export callbackQAreaSeries_PointLabelsVisibilityChanged
func callbackQAreaSeries_PointLabelsVisibilityChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "pointLabelsVisibilityChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QAreaSeries) ConnectPointLabelsVisibilityChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointLabelsVisibilityChanged") {
			C.QAreaSeries_ConnectPointLabelsVisibilityChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointLabelsVisibilityChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointLabelsVisibilityChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsVisibilityChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsVisibilityChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectPointLabelsVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPointLabelsVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointLabelsVisibilityChanged")
	}
}

func (ptr *QAreaSeries) PointLabelsVisibilityChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_PointLabelsVisibilityChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAreaSeries) PointLabelsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAreaSeries_PointLabelsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAreaSeries) PointsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAreaSeries_PointsVisible(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAreaSeries_Pressed
func callbackQAreaSeries_Pressed(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func(*core.QPointF))(signal))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QAreaSeries) ConnectPressed(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pressed") {
			C.QAreaSeries_ConnectPressed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pressed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pressed"); signal != nil {
			f := func(point *core.QPointF) {
				(*(*func(*core.QPointF))(signal))(point)
				f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pressed")
	}
}

func (ptr *QAreaSeries) Pressed(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_Pressed(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQAreaSeries_Released
func callbackQAreaSeries_Released(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "released"); signal != nil {
		(*(*func(*core.QPointF))(signal))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QAreaSeries) ConnectReleased(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "released") {
			C.QAreaSeries_ConnectReleased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "released")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "released"); signal != nil {
			f := func(point *core.QPointF) {
				(*(*func(*core.QPointF))(signal))(point)
				f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "released")
	}
}

func (ptr *QAreaSeries) Released(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_Released(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QAreaSeries) SetBorderColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetBorderColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAreaSeries) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QAreaSeries) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAreaSeries) SetLowerSeries(series QLineSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetLowerSeries(ptr.Pointer(), PointerFromQLineSeries(series))
	}
}

func (ptr *QAreaSeries) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QAreaSeries) SetPointLabelsClipping(enabled bool) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetPointLabelsClipping(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QAreaSeries) SetPointLabelsColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetPointLabelsColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAreaSeries) SetPointLabelsFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetPointLabelsFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QAreaSeries) SetPointLabelsFormat(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QAreaSeries_SetPointLabelsFormat(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

func (ptr *QAreaSeries) SetPointLabelsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetPointLabelsVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAreaSeries) SetPointsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetPointsVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAreaSeries) SetUpperSeries(series QLineSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_SetUpperSeries(ptr.Pointer(), PointerFromQLineSeries(series))
	}
}

//export callbackQAreaSeries_Type
func callbackQAreaSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQAreaSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QAreaSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QAreaSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QAreaSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAreaSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QAreaSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAreaSeries) UpperSeries() *QLineSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQLineSeriesFromPointer(C.QAreaSeries_UpperSeries(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAreaSeries_DestroyQAreaSeries
func callbackQAreaSeries_DestroyQAreaSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAreaSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAreaSeriesFromPointer(ptr).DestroyQAreaSeriesDefault()
	}
}

func (ptr *QAreaSeries) ConnectDestroyQAreaSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAreaSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAreaSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAreaSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAreaSeries) DisconnectDestroyQAreaSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAreaSeries")
	}
}

func (ptr *QAreaSeries) DestroyQAreaSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAreaSeries_DestroyQAreaSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAreaSeries) DestroyQAreaSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAreaSeries_DestroyQAreaSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QBarCategoryAxis struct {
	QAbstractAxis
}

type QBarCategoryAxis_ITF interface {
	QAbstractAxis_ITF
	QBarCategoryAxis_PTR() *QBarCategoryAxis
}

func (ptr *QBarCategoryAxis) QBarCategoryAxis_PTR() *QBarCategoryAxis {
	return ptr
}

func (ptr *QBarCategoryAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractAxis_PTR().Pointer()
	}
	return nil
}

func (ptr *QBarCategoryAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractAxis_PTR().SetPointer(p)
	}
}

func PointerFromQBarCategoryAxis(ptr QBarCategoryAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarCategoryAxis_PTR().Pointer()
	}
	return nil
}

func NewQBarCategoryAxisFromPointer(ptr unsafe.Pointer) (n *QBarCategoryAxis) {
	n = new(QBarCategoryAxis)
	n.SetPointer(ptr)
	return
}
func NewQBarCategoryAxis(parent core.QObject_ITF) *QBarCategoryAxis {
	tmpValue := NewQBarCategoryAxisFromPointer(C.QBarCategoryAxis_NewQBarCategoryAxis(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBarCategoryAxis) Append(categories []string) {
	if ptr.Pointer() != nil {
		categoriesC := C.CString(strings.Join(categories, "¡¦!"))
		defer C.free(unsafe.Pointer(categoriesC))
		C.QBarCategoryAxis_Append(ptr.Pointer(), C.struct_QtCharts_PackedString{data: categoriesC, len: C.longlong(len(strings.Join(categories, "¡¦!")))})
	}
}

func (ptr *QBarCategoryAxis) Append2(category string) {
	if ptr.Pointer() != nil {
		var categoryC *C.char
		if category != "" {
			categoryC = C.CString(category)
			defer C.free(unsafe.Pointer(categoryC))
		}
		C.QBarCategoryAxis_Append2(ptr.Pointer(), C.struct_QtCharts_PackedString{data: categoryC, len: C.longlong(len(category))})
	}
}

func (ptr *QBarCategoryAxis) At(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBarCategoryAxis_At(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QBarCategoryAxis) Categories() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QBarCategoryAxis_Categories(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQBarCategoryAxis_CategoriesChanged
func callbackQBarCategoryAxis_CategoriesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "categoriesChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBarCategoryAxis) ConnectCategoriesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "categoriesChanged") {
			C.QBarCategoryAxis_ConnectCategoriesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "categoriesChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "categoriesChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "categoriesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "categoriesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarCategoryAxis) DisconnectCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DisconnectCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "categoriesChanged")
	}
}

func (ptr *QBarCategoryAxis) CategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_CategoriesChanged(ptr.Pointer())
	}
}

func (ptr *QBarCategoryAxis) Clear() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_Clear(ptr.Pointer())
	}
}

func (ptr *QBarCategoryAxis) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBarCategoryAxis_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQBarCategoryAxis_CountChanged
func callbackQBarCategoryAxis_CountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "countChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBarCategoryAxis) ConnectCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "countChanged") {
			C.QBarCategoryAxis_ConnectCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "countChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "countChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "countChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "countChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarCategoryAxis) DisconnectCountChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DisconnectCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "countChanged")
	}
}

func (ptr *QBarCategoryAxis) CountChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_CountChanged(ptr.Pointer())
	}
}

func (ptr *QBarCategoryAxis) Insert(index int, category string) {
	if ptr.Pointer() != nil {
		var categoryC *C.char
		if category != "" {
			categoryC = C.CString(category)
			defer C.free(unsafe.Pointer(categoryC))
		}
		C.QBarCategoryAxis_Insert(ptr.Pointer(), C.int(int32(index)), C.struct_QtCharts_PackedString{data: categoryC, len: C.longlong(len(category))})
	}
}

func (ptr *QBarCategoryAxis) Max() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBarCategoryAxis_Max(ptr.Pointer()))
	}
	return ""
}

//export callbackQBarCategoryAxis_MaxChanged
func callbackQBarCategoryAxis_MaxChanged(ptr unsafe.Pointer, max C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "maxChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(max))
	}

}

func (ptr *QBarCategoryAxis) ConnectMaxChanged(f func(max string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "maxChanged") {
			C.QBarCategoryAxis_ConnectMaxChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "maxChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "maxChanged"); signal != nil {
			f := func(max string) {
				(*(*func(string))(signal))(max)
				f(max)
			}
			qt.ConnectSignal(ptr.Pointer(), "maxChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maxChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarCategoryAxis) DisconnectMaxChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DisconnectMaxChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "maxChanged")
	}
}

func (ptr *QBarCategoryAxis) MaxChanged(max string) {
	if ptr.Pointer() != nil {
		var maxC *C.char
		if max != "" {
			maxC = C.CString(max)
			defer C.free(unsafe.Pointer(maxC))
		}
		C.QBarCategoryAxis_MaxChanged(ptr.Pointer(), C.struct_QtCharts_PackedString{data: maxC, len: C.longlong(len(max))})
	}
}

func (ptr *QBarCategoryAxis) Min() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBarCategoryAxis_Min(ptr.Pointer()))
	}
	return ""
}

//export callbackQBarCategoryAxis_MinChanged
func callbackQBarCategoryAxis_MinChanged(ptr unsafe.Pointer, min C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "minChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(min))
	}

}

func (ptr *QBarCategoryAxis) ConnectMinChanged(f func(min string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minChanged") {
			C.QBarCategoryAxis_ConnectMinChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minChanged"); signal != nil {
			f := func(min string) {
				(*(*func(string))(signal))(min)
				f(min)
			}
			qt.ConnectSignal(ptr.Pointer(), "minChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarCategoryAxis) DisconnectMinChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DisconnectMinChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minChanged")
	}
}

func (ptr *QBarCategoryAxis) MinChanged(min string) {
	if ptr.Pointer() != nil {
		var minC *C.char
		if min != "" {
			minC = C.CString(min)
			defer C.free(unsafe.Pointer(minC))
		}
		C.QBarCategoryAxis_MinChanged(ptr.Pointer(), C.struct_QtCharts_PackedString{data: minC, len: C.longlong(len(min))})
	}
}

//export callbackQBarCategoryAxis_RangeChanged
func callbackQBarCategoryAxis_RangeChanged(ptr unsafe.Pointer, min C.struct_QtCharts_PackedString, max C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "rangeChanged"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(min), cGoUnpackString(max))
	}

}

func (ptr *QBarCategoryAxis) ConnectRangeChanged(f func(min string, max string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rangeChanged") {
			C.QBarCategoryAxis_ConnectRangeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "rangeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rangeChanged"); signal != nil {
			f := func(min string, max string) {
				(*(*func(string, string))(signal))(min, max)
				f(min, max)
			}
			qt.ConnectSignal(ptr.Pointer(), "rangeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rangeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarCategoryAxis) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rangeChanged")
	}
}

func (ptr *QBarCategoryAxis) RangeChanged(min string, max string) {
	if ptr.Pointer() != nil {
		var minC *C.char
		if min != "" {
			minC = C.CString(min)
			defer C.free(unsafe.Pointer(minC))
		}
		var maxC *C.char
		if max != "" {
			maxC = C.CString(max)
			defer C.free(unsafe.Pointer(maxC))
		}
		C.QBarCategoryAxis_RangeChanged(ptr.Pointer(), C.struct_QtCharts_PackedString{data: minC, len: C.longlong(len(min))}, C.struct_QtCharts_PackedString{data: maxC, len: C.longlong(len(max))})
	}
}

func (ptr *QBarCategoryAxis) Remove(category string) {
	if ptr.Pointer() != nil {
		var categoryC *C.char
		if category != "" {
			categoryC = C.CString(category)
			defer C.free(unsafe.Pointer(categoryC))
		}
		C.QBarCategoryAxis_Remove(ptr.Pointer(), C.struct_QtCharts_PackedString{data: categoryC, len: C.longlong(len(category))})
	}
}

func (ptr *QBarCategoryAxis) Replace(oldCategory string, newCategory string) {
	if ptr.Pointer() != nil {
		var oldCategoryC *C.char
		if oldCategory != "" {
			oldCategoryC = C.CString(oldCategory)
			defer C.free(unsafe.Pointer(oldCategoryC))
		}
		var newCategoryC *C.char
		if newCategory != "" {
			newCategoryC = C.CString(newCategory)
			defer C.free(unsafe.Pointer(newCategoryC))
		}
		C.QBarCategoryAxis_Replace(ptr.Pointer(), C.struct_QtCharts_PackedString{data: oldCategoryC, len: C.longlong(len(oldCategory))}, C.struct_QtCharts_PackedString{data: newCategoryC, len: C.longlong(len(newCategory))})
	}
}

func (ptr *QBarCategoryAxis) SetCategories(categories []string) {
	if ptr.Pointer() != nil {
		categoriesC := C.CString(strings.Join(categories, "¡¦!"))
		defer C.free(unsafe.Pointer(categoriesC))
		C.QBarCategoryAxis_SetCategories(ptr.Pointer(), C.struct_QtCharts_PackedString{data: categoriesC, len: C.longlong(len(strings.Join(categories, "¡¦!")))})
	}
}

func (ptr *QBarCategoryAxis) SetMax(max string) {
	if ptr.Pointer() != nil {
		var maxC *C.char
		if max != "" {
			maxC = C.CString(max)
			defer C.free(unsafe.Pointer(maxC))
		}
		C.QBarCategoryAxis_SetMax(ptr.Pointer(), C.struct_QtCharts_PackedString{data: maxC, len: C.longlong(len(max))})
	}
}

func (ptr *QBarCategoryAxis) SetMin(min string) {
	if ptr.Pointer() != nil {
		var minC *C.char
		if min != "" {
			minC = C.CString(min)
			defer C.free(unsafe.Pointer(minC))
		}
		C.QBarCategoryAxis_SetMin(ptr.Pointer(), C.struct_QtCharts_PackedString{data: minC, len: C.longlong(len(min))})
	}
}

func (ptr *QBarCategoryAxis) SetRange(minCategory string, maxCategory string) {
	if ptr.Pointer() != nil {
		var minCategoryC *C.char
		if minCategory != "" {
			minCategoryC = C.CString(minCategory)
			defer C.free(unsafe.Pointer(minCategoryC))
		}
		var maxCategoryC *C.char
		if maxCategory != "" {
			maxCategoryC = C.CString(maxCategory)
			defer C.free(unsafe.Pointer(maxCategoryC))
		}
		C.QBarCategoryAxis_SetRange(ptr.Pointer(), C.struct_QtCharts_PackedString{data: minCategoryC, len: C.longlong(len(minCategory))}, C.struct_QtCharts_PackedString{data: maxCategoryC, len: C.longlong(len(maxCategory))})
	}
}

//export callbackQBarCategoryAxis_Type
func callbackQBarCategoryAxis_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractAxis__AxisType)(signal))())
	}

	return C.longlong(NewQBarCategoryAxisFromPointer(ptr).TypeDefault())
}

func (ptr *QBarCategoryAxis) ConnectType(f func() QAbstractAxis__AxisType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractAxis__AxisType {
				(*(*func() QAbstractAxis__AxisType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarCategoryAxis) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QBarCategoryAxis) Type() QAbstractAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstractAxis__AxisType(C.QBarCategoryAxis_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBarCategoryAxis) TypeDefault() QAbstractAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstractAxis__AxisType(C.QBarCategoryAxis_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBarCategoryAxis_DestroyQBarCategoryAxis
func callbackQBarCategoryAxis_DestroyQBarCategoryAxis(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBarCategoryAxis"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBarCategoryAxisFromPointer(ptr).DestroyQBarCategoryAxisDefault()
	}
}

func (ptr *QBarCategoryAxis) ConnectDestroyQBarCategoryAxis(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBarCategoryAxis"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBarCategoryAxis", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBarCategoryAxis", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarCategoryAxis) DisconnectDestroyQBarCategoryAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBarCategoryAxis")
	}
}

func (ptr *QBarCategoryAxis) DestroyQBarCategoryAxis() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBarCategoryAxis_DestroyQBarCategoryAxis(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarCategoryAxis) DestroyQBarCategoryAxisDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBarCategoryAxis_DestroyQBarCategoryAxisDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QBarLegendMarker struct {
	QLegendMarker
}

type QBarLegendMarker_ITF interface {
	QLegendMarker_ITF
	QBarLegendMarker_PTR() *QBarLegendMarker
}

func (ptr *QBarLegendMarker) QBarLegendMarker_PTR() *QBarLegendMarker {
	return ptr
}

func (ptr *QBarLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLegendMarker_PTR().Pointer()
	}
	return nil
}

func (ptr *QBarLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLegendMarker_PTR().SetPointer(p)
	}
}

func PointerFromQBarLegendMarker(ptr QBarLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQBarLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QBarLegendMarker) {
	n = new(QBarLegendMarker)
	n.SetPointer(ptr)
	return
}
func (ptr *QBarLegendMarker) Barset() *QBarSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQBarSetFromPointer(C.QBarLegendMarker_Barset(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQBarLegendMarker_Series
func callbackQBarLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "series"); signal != nil {
		return PointerFromQAbstractBarSeries((*(*func() *QAbstractBarSeries)(signal))())
	}

	return PointerFromQAbstractBarSeries(NewQBarLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QBarLegendMarker) ConnectSeries(f func() *QAbstractBarSeries) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "series"); signal != nil {
			f := func() *QAbstractBarSeries {
				(*(*func() *QAbstractBarSeries)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "series")
	}
}

func (ptr *QBarLegendMarker) Series() *QAbstractBarSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractBarSeriesFromPointer(C.QBarLegendMarker_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBarLegendMarker) SeriesDefault() *QAbstractBarSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractBarSeriesFromPointer(C.QBarLegendMarker_SeriesDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQBarLegendMarker_Type
func callbackQBarLegendMarker_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QLegendMarker__LegendMarkerType)(signal))())
	}

	return C.longlong(NewQBarLegendMarkerFromPointer(ptr).TypeDefault())
}

func (ptr *QBarLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QLegendMarker__LegendMarkerType {
				(*(*func() QLegendMarker__LegendMarkerType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarLegendMarker) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QBarLegendMarker) Type() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QBarLegendMarker_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBarLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QBarLegendMarker_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBarLegendMarker_DestroyQBarLegendMarker
func callbackQBarLegendMarker_DestroyQBarLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBarLegendMarker"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBarLegendMarkerFromPointer(ptr).DestroyQBarLegendMarkerDefault()
	}
}

func (ptr *QBarLegendMarker) ConnectDestroyQBarLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBarLegendMarker"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBarLegendMarker", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBarLegendMarker", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarLegendMarker) DisconnectDestroyQBarLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBarLegendMarker")
	}
}

func (ptr *QBarLegendMarker) DestroyQBarLegendMarker() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBarLegendMarker_DestroyQBarLegendMarker(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarLegendMarker) DestroyQBarLegendMarkerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBarLegendMarker_DestroyQBarLegendMarkerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QBarSeries struct {
	QAbstractBarSeries
}

type QBarSeries_ITF interface {
	QAbstractBarSeries_ITF
	QBarSeries_PTR() *QBarSeries
}

func (ptr *QBarSeries) QBarSeries_PTR() *QBarSeries {
	return ptr
}

func (ptr *QBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractBarSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractBarSeries_PTR().SetPointer(p)
	}
}

func PointerFromQBarSeries(ptr QBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQBarSeriesFromPointer(ptr unsafe.Pointer) (n *QBarSeries) {
	n = new(QBarSeries)
	n.SetPointer(ptr)
	return
}
func NewQBarSeries(parent core.QObject_ITF) *QBarSeries {
	tmpValue := NewQBarSeriesFromPointer(C.QBarSeries_NewQBarSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBarSeries_Type
func callbackQBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QBarSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QBarSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBarSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QBarSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBarSeries_DestroyQBarSeries
func callbackQBarSeries_DestroyQBarSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBarSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBarSeriesFromPointer(ptr).DestroyQBarSeriesDefault()
	}
}

func (ptr *QBarSeries) ConnectDestroyQBarSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBarSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBarSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBarSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSeries) DisconnectDestroyQBarSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBarSeries")
	}
}

func (ptr *QBarSeries) DestroyQBarSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBarSeries_DestroyQBarSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarSeries) DestroyQBarSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBarSeries_DestroyQBarSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QBarSet struct {
	core.QObject
}

type QBarSet_ITF interface {
	core.QObject_ITF
	QBarSet_PTR() *QBarSet
}

func (ptr *QBarSet) QBarSet_PTR() *QBarSet {
	return ptr
}

func (ptr *QBarSet) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBarSet) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQBarSet(ptr QBarSet_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarSet_PTR().Pointer()
	}
	return nil
}

func NewQBarSetFromPointer(ptr unsafe.Pointer) (n *QBarSet) {
	n = new(QBarSet)
	n.SetPointer(ptr)
	return
}
func NewQBarSet(label string, parent core.QObject_ITF) *QBarSet {
	var labelC *C.char
	if label != "" {
		labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
	}
	tmpValue := NewQBarSetFromPointer(C.QBarSet_NewQBarSet(C.struct_QtCharts_PackedString{data: labelC, len: C.longlong(len(label))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBarSet) Append(value float64) {
	if ptr.Pointer() != nil {
		C.QBarSet_Append(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QBarSet) Append2(values []float64) {
	if ptr.Pointer() != nil {
		C.QBarSet_Append2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQBarSetFromPointer(NewQBarSetFromPointer(nil).__append_values_newList2())
			for _, v := range values {
				tmpList.__append_values_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QBarSet) At(index int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QBarSet_At(ptr.Pointer(), C.int(int32(index))))
	}
	return 0
}

func (ptr *QBarSet) BorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QBarSet_BorderColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQBarSet_BorderColorChanged
func callbackQBarSet_BorderColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "borderColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QBarSet) ConnectBorderColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "borderColorChanged") {
			C.QBarSet_ConnectBorderColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "borderColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "borderColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "borderColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "borderColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "borderColorChanged")
	}
}

func (ptr *QBarSet) BorderColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_BorderColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QBarSet) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QBarSet_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQBarSet_BrushChanged
func callbackQBarSet_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "brushChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBarSet) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "brushChanged") {
			C.QBarSet_ConnectBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "brushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "brushChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "brushChanged")
	}
}

func (ptr *QBarSet) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_BrushChanged(ptr.Pointer())
	}
}

//export callbackQBarSet_Clicked
func callbackQBarSet_Clicked(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectClicked(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QBarSet_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QBarSet) Clicked(index int) {
	if ptr.Pointer() != nil {
		C.QBarSet_Clicked(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QBarSet) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QBarSet_Color(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQBarSet_ColorChanged
func callbackQBarSet_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "colorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QBarSet) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "colorChanged") {
			C.QBarSet_ConnectColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "colorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "colorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "colorChanged")
	}
}

func (ptr *QBarSet) ColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QBarSet) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBarSet_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQBarSet_DoubleClicked
func callbackQBarSet_DoubleClicked(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectDoubleClicked(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "doubleClicked") {
			C.QBarSet_ConnectDoubleClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "doubleClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "doubleClicked"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "doubleClicked")
	}
}

func (ptr *QBarSet) DoubleClicked(index int) {
	if ptr.Pointer() != nil {
		C.QBarSet_DoubleClicked(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQBarSet_Hovered
func callbackQBarSet_Hovered(ptr unsafe.Pointer, status C.char, index C.int) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(bool, int))(signal))(int8(status) != 0, int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectHovered(f func(status bool, index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QBarSet_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(status bool, index int) {
				(*(*func(bool, int))(signal))(status, index)
				f(status, index)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QBarSet) Hovered(status bool, index int) {
	if ptr.Pointer() != nil {
		C.QBarSet_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))), C.int(int32(index)))
	}
}

func (ptr *QBarSet) Insert(index int, value float64) {
	if ptr.Pointer() != nil {
		C.QBarSet_Insert(ptr.Pointer(), C.int(int32(index)), C.double(value))
	}
}

func (ptr *QBarSet) Label() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBarSet_Label(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBarSet) LabelBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QBarSet_LabelBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQBarSet_LabelBrushChanged
func callbackQBarSet_LabelBrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelBrushChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBarSet) ConnectLabelBrushChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelBrushChanged") {
			C.QBarSet_ConnectLabelBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelBrushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelBrushChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelBrushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelBrushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectLabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectLabelBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelBrushChanged")
	}
}

func (ptr *QBarSet) LabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_LabelBrushChanged(ptr.Pointer())
	}
}

//export callbackQBarSet_LabelChanged
func callbackQBarSet_LabelChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBarSet) ConnectLabelChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelChanged") {
			C.QBarSet_ConnectLabelChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectLabelChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectLabelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelChanged")
	}
}

func (ptr *QBarSet) LabelChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_LabelChanged(ptr.Pointer())
	}
}

func (ptr *QBarSet) LabelColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QBarSet_LabelColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQBarSet_LabelColorChanged
func callbackQBarSet_LabelColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QBarSet) ConnectLabelColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelColorChanged") {
			C.QBarSet_ConnectLabelColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectLabelColorChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectLabelColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelColorChanged")
	}
}

func (ptr *QBarSet) LabelColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_LabelColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QBarSet) LabelFont() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QBarSet_LabelFont(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

//export callbackQBarSet_LabelFontChanged
func callbackQBarSet_LabelFontChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelFontChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBarSet) ConnectLabelFontChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelFontChanged") {
			C.QBarSet_ConnectLabelFontChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelFontChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelFontChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelFontChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelFontChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectLabelFontChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectLabelFontChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelFontChanged")
	}
}

func (ptr *QBarSet) LabelFontChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_LabelFontChanged(ptr.Pointer())
	}
}

func (ptr *QBarSet) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QBarSet_Pen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQBarSet_PenChanged
func callbackQBarSet_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "penChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBarSet) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "penChanged") {
			C.QBarSet_ConnectPenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "penChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "penChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "penChanged")
	}
}

func (ptr *QBarSet) PenChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_PenChanged(ptr.Pointer())
	}
}

//export callbackQBarSet_Pressed
func callbackQBarSet_Pressed(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectPressed(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pressed") {
			C.QBarSet_ConnectPressed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pressed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pressed"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pressed")
	}
}

func (ptr *QBarSet) Pressed(index int) {
	if ptr.Pointer() != nil {
		C.QBarSet_Pressed(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQBarSet_Released
func callbackQBarSet_Released(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "released"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectReleased(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "released") {
			C.QBarSet_ConnectReleased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "released")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "released"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "released")
	}
}

func (ptr *QBarSet) Released(index int) {
	if ptr.Pointer() != nil {
		C.QBarSet_Released(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QBarSet) Remove(index int, count int) {
	if ptr.Pointer() != nil {
		C.QBarSet_Remove(ptr.Pointer(), C.int(int32(index)), C.int(int32(count)))
	}
}

func (ptr *QBarSet) Replace(index int, value float64) {
	if ptr.Pointer() != nil {
		C.QBarSet_Replace(ptr.Pointer(), C.int(int32(index)), C.double(value))
	}
}

func (ptr *QBarSet) SetBorderColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_SetBorderColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QBarSet) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QBarSet) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QBarSet) SetLabel(label string) {
	if ptr.Pointer() != nil {
		var labelC *C.char
		if label != "" {
			labelC = C.CString(label)
			defer C.free(unsafe.Pointer(labelC))
		}
		C.QBarSet_SetLabel(ptr.Pointer(), C.struct_QtCharts_PackedString{data: labelC, len: C.longlong(len(label))})
	}
}

func (ptr *QBarSet) SetLabelBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_SetLabelBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QBarSet) SetLabelColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_SetLabelColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QBarSet) SetLabelFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_SetLabelFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QBarSet) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QBarSet) Sum() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QBarSet_Sum(ptr.Pointer()))
	}
	return 0
}

//export callbackQBarSet_ValueChanged
func callbackQBarSet_ValueChanged(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "valueChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectValueChanged(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valueChanged") {
			C.QBarSet_ConnectValueChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "valueChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valueChanged"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valueChanged")
	}
}

func (ptr *QBarSet) ValueChanged(index int) {
	if ptr.Pointer() != nil {
		C.QBarSet_ValueChanged(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQBarSet_ValuesAdded
func callbackQBarSet_ValuesAdded(ptr unsafe.Pointer, index C.int, count C.int) {
	if signal := qt.GetSignal(ptr, "valuesAdded"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(index)), int(int32(count)))
	}

}

func (ptr *QBarSet) ConnectValuesAdded(f func(index int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valuesAdded") {
			C.QBarSet_ConnectValuesAdded(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "valuesAdded")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valuesAdded"); signal != nil {
			f := func(index int, count int) {
				(*(*func(int, int))(signal))(index, count)
				f(index, count)
			}
			qt.ConnectSignal(ptr.Pointer(), "valuesAdded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valuesAdded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectValuesAdded() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectValuesAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valuesAdded")
	}
}

func (ptr *QBarSet) ValuesAdded(index int, count int) {
	if ptr.Pointer() != nil {
		C.QBarSet_ValuesAdded(ptr.Pointer(), C.int(int32(index)), C.int(int32(count)))
	}
}

//export callbackQBarSet_ValuesRemoved
func callbackQBarSet_ValuesRemoved(ptr unsafe.Pointer, index C.int, count C.int) {
	if signal := qt.GetSignal(ptr, "valuesRemoved"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(index)), int(int32(count)))
	}

}

func (ptr *QBarSet) ConnectValuesRemoved(f func(index int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valuesRemoved") {
			C.QBarSet_ConnectValuesRemoved(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "valuesRemoved")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valuesRemoved"); signal != nil {
			f := func(index int, count int) {
				(*(*func(int, int))(signal))(index, count)
				f(index, count)
			}
			qt.ConnectSignal(ptr.Pointer(), "valuesRemoved", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valuesRemoved", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectValuesRemoved() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectValuesRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valuesRemoved")
	}
}

func (ptr *QBarSet) ValuesRemoved(index int, count int) {
	if ptr.Pointer() != nil {
		C.QBarSet_ValuesRemoved(ptr.Pointer(), C.int(int32(index)), C.int(int32(count)))
	}
}

//export callbackQBarSet_DestroyQBarSet
func callbackQBarSet_DestroyQBarSet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBarSet"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBarSetFromPointer(ptr).DestroyQBarSetDefault()
	}
}

func (ptr *QBarSet) ConnectDestroyQBarSet(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBarSet"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBarSet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBarSet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBarSet) DisconnectDestroyQBarSet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBarSet")
	}
}

func (ptr *QBarSet) DestroyQBarSet() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBarSet_DestroyQBarSet(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarSet) DestroyQBarSetDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBarSet_DestroyQBarSetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarSet) __append_values_atList2(i int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QBarSet___append_values_atList2(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QBarSet) __append_values_setList2(i float64) {
	if ptr.Pointer() != nil {
		C.QBarSet___append_values_setList2(ptr.Pointer(), C.double(i))
	}
}

func (ptr *QBarSet) __append_values_newList2() unsafe.Pointer {
	return C.QBarSet___append_values_newList2(ptr.Pointer())
}

func (ptr *QBarSet) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBarSet___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBarSet) __children_newList() unsafe.Pointer {
	return C.QBarSet___children_newList(ptr.Pointer())
}

func (ptr *QBarSet) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QBarSet___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QBarSet) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QBarSet___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QBarSet) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBarSet___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBarSet) __findChildren_newList() unsafe.Pointer {
	return C.QBarSet___findChildren_newList(ptr.Pointer())
}

func (ptr *QBarSet) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBarSet___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBarSet) __findChildren_newList3() unsafe.Pointer {
	return C.QBarSet___findChildren_newList3(ptr.Pointer())
}

//export callbackQBarSet_ChildEvent
func callbackQBarSet_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBarSetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBarSet) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBarSet_ConnectNotify
func callbackQBarSet_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBarSetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBarSet) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBarSet_CustomEvent
func callbackQBarSet_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQBarSetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBarSet) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBarSet_DeleteLater
func callbackQBarSet_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBarSetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBarSet) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBarSet_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQBarSet_Destroyed
func callbackQBarSet_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBarSet_DisconnectNotify
func callbackQBarSet_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBarSetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBarSet) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBarSet_Event
func callbackQBarSet_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBarSetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBarSet) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBarSet_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQBarSet_EventFilter
func callbackQBarSet_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBarSetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBarSet) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBarSet_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQBarSet_MetaObject
func callbackQBarSet_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQBarSetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBarSet) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBarSet_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQBarSet_ObjectNameChanged
func callbackQBarSet_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQBarSet_TimerEvent
func callbackQBarSet_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBarSetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBarSet) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QBoxPlotLegendMarker struct {
	QLegendMarker
}

type QBoxPlotLegendMarker_ITF interface {
	QLegendMarker_ITF
	QBoxPlotLegendMarker_PTR() *QBoxPlotLegendMarker
}

func (ptr *QBoxPlotLegendMarker) QBoxPlotLegendMarker_PTR() *QBoxPlotLegendMarker {
	return ptr
}

func (ptr *QBoxPlotLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLegendMarker_PTR().Pointer()
	}
	return nil
}

func (ptr *QBoxPlotLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLegendMarker_PTR().SetPointer(p)
	}
}

func PointerFromQBoxPlotLegendMarker(ptr QBoxPlotLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxPlotLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQBoxPlotLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QBoxPlotLegendMarker) {
	n = new(QBoxPlotLegendMarker)
	n.SetPointer(ptr)
	return
}

//export callbackQBoxPlotLegendMarker_Series
func callbackQBoxPlotLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "series"); signal != nil {
		return PointerFromQBoxPlotSeries((*(*func() *QBoxPlotSeries)(signal))())
	}

	return PointerFromQBoxPlotSeries(NewQBoxPlotLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QBoxPlotLegendMarker) ConnectSeries(f func() *QBoxPlotSeries) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "series"); signal != nil {
			f := func() *QBoxPlotSeries {
				(*(*func() *QBoxPlotSeries)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "series")
	}
}

func (ptr *QBoxPlotLegendMarker) Series() *QBoxPlotSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQBoxPlotSeriesFromPointer(C.QBoxPlotLegendMarker_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBoxPlotLegendMarker) SeriesDefault() *QBoxPlotSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQBoxPlotSeriesFromPointer(C.QBoxPlotLegendMarker_SeriesDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQBoxPlotLegendMarker_Type
func callbackQBoxPlotLegendMarker_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QLegendMarker__LegendMarkerType)(signal))())
	}

	return C.longlong(NewQBoxPlotLegendMarkerFromPointer(ptr).TypeDefault())
}

func (ptr *QBoxPlotLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QLegendMarker__LegendMarkerType {
				(*(*func() QLegendMarker__LegendMarkerType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotLegendMarker) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QBoxPlotLegendMarker) Type() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QBoxPlotLegendMarker_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxPlotLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QBoxPlotLegendMarker_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBoxPlotLegendMarker_DestroyQBoxPlotLegendMarker
func callbackQBoxPlotLegendMarker_DestroyQBoxPlotLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBoxPlotLegendMarker"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBoxPlotLegendMarkerFromPointer(ptr).DestroyQBoxPlotLegendMarkerDefault()
	}
}

func (ptr *QBoxPlotLegendMarker) ConnectDestroyQBoxPlotLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBoxPlotLegendMarker"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBoxPlotLegendMarker", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBoxPlotLegendMarker", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotLegendMarker) DisconnectDestroyQBoxPlotLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBoxPlotLegendMarker")
	}
}

func (ptr *QBoxPlotLegendMarker) DestroyQBoxPlotLegendMarker() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBoxPlotLegendMarker_DestroyQBoxPlotLegendMarker(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBoxPlotLegendMarker) DestroyQBoxPlotLegendMarkerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBoxPlotLegendMarker_DestroyQBoxPlotLegendMarkerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QBoxPlotSeries struct {
	QAbstractSeries
}

type QBoxPlotSeries_ITF interface {
	QAbstractSeries_ITF
	QBoxPlotSeries_PTR() *QBoxPlotSeries
}

func (ptr *QBoxPlotSeries) QBoxPlotSeries_PTR() *QBoxPlotSeries {
	return ptr
}

func (ptr *QBoxPlotSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QBoxPlotSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSeries_PTR().SetPointer(p)
	}
}

func PointerFromQBoxPlotSeries(ptr QBoxPlotSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxPlotSeries_PTR().Pointer()
	}
	return nil
}

func NewQBoxPlotSeriesFromPointer(ptr unsafe.Pointer) (n *QBoxPlotSeries) {
	n = new(QBoxPlotSeries)
	n.SetPointer(ptr)
	return
}
func NewQBoxPlotSeries(parent core.QObject_ITF) *QBoxPlotSeries {
	tmpValue := NewQBoxPlotSeriesFromPointer(C.QBoxPlotSeries_NewQBoxPlotSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBoxPlotSeries) Append(set QBoxSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBoxPlotSeries_Append(ptr.Pointer(), PointerFromQBoxSet(set))) != 0
	}
	return false
}

func (ptr *QBoxPlotSeries) Append2(sets []*QBoxSet) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBoxPlotSeries_Append2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQBoxPlotSeriesFromPointer(NewQBoxPlotSeriesFromPointer(nil).__append_sets_newList2())
			for _, v := range sets {
				tmpList.__append_sets_setList2(v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackQBoxPlotSeries_BoxOutlineVisibilityChanged
func callbackQBoxPlotSeries_BoxOutlineVisibilityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "boxOutlineVisibilityChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxPlotSeries) ConnectBoxOutlineVisibilityChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "boxOutlineVisibilityChanged") {
			C.QBoxPlotSeries_ConnectBoxOutlineVisibilityChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "boxOutlineVisibilityChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "boxOutlineVisibilityChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "boxOutlineVisibilityChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "boxOutlineVisibilityChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectBoxOutlineVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectBoxOutlineVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "boxOutlineVisibilityChanged")
	}
}

func (ptr *QBoxPlotSeries) BoxOutlineVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_BoxOutlineVisibilityChanged(ptr.Pointer())
	}
}

func (ptr *QBoxPlotSeries) BoxOutlineVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBoxPlotSeries_BoxOutlineVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBoxPlotSeries) BoxSets() []*QBoxSet {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QBoxSet {
			out := make([]*QBoxSet, int(l.len))
			tmpList := NewQBoxPlotSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__boxSets_atList(i)
			}
			return out
		}(C.QBoxPlotSeries_BoxSets(ptr.Pointer()))
	}
	return make([]*QBoxSet, 0)
}

func (ptr *QBoxPlotSeries) BoxWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QBoxPlotSeries_BoxWidth(ptr.Pointer()))
	}
	return 0
}

//export callbackQBoxPlotSeries_BoxWidthChanged
func callbackQBoxPlotSeries_BoxWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "boxWidthChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxPlotSeries) ConnectBoxWidthChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "boxWidthChanged") {
			C.QBoxPlotSeries_ConnectBoxWidthChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "boxWidthChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "boxWidthChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "boxWidthChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "boxWidthChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectBoxWidthChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectBoxWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "boxWidthChanged")
	}
}

func (ptr *QBoxPlotSeries) BoxWidthChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_BoxWidthChanged(ptr.Pointer())
	}
}

//export callbackQBoxPlotSeries_BoxsetsAdded
func callbackQBoxPlotSeries_BoxsetsAdded(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(ptr, "boxsetsAdded"); signal != nil {
		(*(*func([]*QBoxSet))(signal))(func(l C.struct_QtCharts_PackedList) []*QBoxSet {
			out := make([]*QBoxSet, int(l.len))
			tmpList := NewQBoxPlotSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__boxsetsAdded_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QBoxPlotSeries) ConnectBoxsetsAdded(f func(sets []*QBoxSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "boxsetsAdded") {
			C.QBoxPlotSeries_ConnectBoxsetsAdded(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "boxsetsAdded")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "boxsetsAdded"); signal != nil {
			f := func(sets []*QBoxSet) {
				(*(*func([]*QBoxSet))(signal))(sets)
				f(sets)
			}
			qt.ConnectSignal(ptr.Pointer(), "boxsetsAdded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "boxsetsAdded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectBoxsetsAdded() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectBoxsetsAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "boxsetsAdded")
	}
}

func (ptr *QBoxPlotSeries) BoxsetsAdded(sets []*QBoxSet) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_BoxsetsAdded(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQBoxPlotSeriesFromPointer(NewQBoxPlotSeriesFromPointer(nil).__boxsetsAdded_sets_newList())
			for _, v := range sets {
				tmpList.__boxsetsAdded_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQBoxPlotSeries_BoxsetsRemoved
func callbackQBoxPlotSeries_BoxsetsRemoved(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(ptr, "boxsetsRemoved"); signal != nil {
		(*(*func([]*QBoxSet))(signal))(func(l C.struct_QtCharts_PackedList) []*QBoxSet {
			out := make([]*QBoxSet, int(l.len))
			tmpList := NewQBoxPlotSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__boxsetsRemoved_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QBoxPlotSeries) ConnectBoxsetsRemoved(f func(sets []*QBoxSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "boxsetsRemoved") {
			C.QBoxPlotSeries_ConnectBoxsetsRemoved(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "boxsetsRemoved")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "boxsetsRemoved"); signal != nil {
			f := func(sets []*QBoxSet) {
				(*(*func([]*QBoxSet))(signal))(sets)
				f(sets)
			}
			qt.ConnectSignal(ptr.Pointer(), "boxsetsRemoved", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "boxsetsRemoved", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectBoxsetsRemoved() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectBoxsetsRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "boxsetsRemoved")
	}
}

func (ptr *QBoxPlotSeries) BoxsetsRemoved(sets []*QBoxSet) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_BoxsetsRemoved(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQBoxPlotSeriesFromPointer(NewQBoxPlotSeriesFromPointer(nil).__boxsetsRemoved_sets_newList())
			for _, v := range sets {
				tmpList.__boxsetsRemoved_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QBoxPlotSeries) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QBoxPlotSeries_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQBoxPlotSeries_BrushChanged
func callbackQBoxPlotSeries_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "brushChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxPlotSeries) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "brushChanged") {
			C.QBoxPlotSeries_ConnectBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "brushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "brushChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "brushChanged")
	}
}

func (ptr *QBoxPlotSeries) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_BrushChanged(ptr.Pointer())
	}
}

func (ptr *QBoxPlotSeries) Clear() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_Clear(ptr.Pointer())
	}
}

//export callbackQBoxPlotSeries_Clicked
func callbackQBoxPlotSeries_Clicked(ptr unsafe.Pointer, boxset unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(*QBoxSet))(signal))(NewQBoxSetFromPointer(boxset))
	}

}

func (ptr *QBoxPlotSeries) ConnectClicked(f func(boxset *QBoxSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QBoxPlotSeries_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(boxset *QBoxSet) {
				(*(*func(*QBoxSet))(signal))(boxset)
				f(boxset)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QBoxPlotSeries) Clicked(boxset QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_Clicked(ptr.Pointer(), PointerFromQBoxSet(boxset))
	}
}

func (ptr *QBoxPlotSeries) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBoxPlotSeries_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQBoxPlotSeries_CountChanged
func callbackQBoxPlotSeries_CountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "countChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxPlotSeries) ConnectCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "countChanged") {
			C.QBoxPlotSeries_ConnectCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "countChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "countChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "countChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "countChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectCountChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "countChanged")
	}
}

func (ptr *QBoxPlotSeries) CountChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_CountChanged(ptr.Pointer())
	}
}

//export callbackQBoxPlotSeries_DoubleClicked
func callbackQBoxPlotSeries_DoubleClicked(ptr unsafe.Pointer, boxset unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		(*(*func(*QBoxSet))(signal))(NewQBoxSetFromPointer(boxset))
	}

}

func (ptr *QBoxPlotSeries) ConnectDoubleClicked(f func(boxset *QBoxSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "doubleClicked") {
			C.QBoxPlotSeries_ConnectDoubleClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "doubleClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "doubleClicked"); signal != nil {
			f := func(boxset *QBoxSet) {
				(*(*func(*QBoxSet))(signal))(boxset)
				f(boxset)
			}
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "doubleClicked")
	}
}

func (ptr *QBoxPlotSeries) DoubleClicked(boxset QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DoubleClicked(ptr.Pointer(), PointerFromQBoxSet(boxset))
	}
}

//export callbackQBoxPlotSeries_Hovered
func callbackQBoxPlotSeries_Hovered(ptr unsafe.Pointer, status C.char, boxset unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(bool, *QBoxSet))(signal))(int8(status) != 0, NewQBoxSetFromPointer(boxset))
	}

}

func (ptr *QBoxPlotSeries) ConnectHovered(f func(status bool, boxset *QBoxSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QBoxPlotSeries_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(status bool, boxset *QBoxSet) {
				(*(*func(bool, *QBoxSet))(signal))(status, boxset)
				f(status, boxset)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QBoxPlotSeries) Hovered(status bool, boxset QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))), PointerFromQBoxSet(boxset))
	}
}

func (ptr *QBoxPlotSeries) Insert(index int, set QBoxSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBoxPlotSeries_Insert(ptr.Pointer(), C.int(int32(index)), PointerFromQBoxSet(set))) != 0
	}
	return false
}

func (ptr *QBoxPlotSeries) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QBoxPlotSeries_Pen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQBoxPlotSeries_PenChanged
func callbackQBoxPlotSeries_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "penChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxPlotSeries) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "penChanged") {
			C.QBoxPlotSeries_ConnectPenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "penChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "penChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "penChanged")
	}
}

func (ptr *QBoxPlotSeries) PenChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_PenChanged(ptr.Pointer())
	}
}

//export callbackQBoxPlotSeries_Pressed
func callbackQBoxPlotSeries_Pressed(ptr unsafe.Pointer, boxset unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func(*QBoxSet))(signal))(NewQBoxSetFromPointer(boxset))
	}

}

func (ptr *QBoxPlotSeries) ConnectPressed(f func(boxset *QBoxSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pressed") {
			C.QBoxPlotSeries_ConnectPressed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pressed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pressed"); signal != nil {
			f := func(boxset *QBoxSet) {
				(*(*func(*QBoxSet))(signal))(boxset)
				f(boxset)
			}
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pressed")
	}
}

func (ptr *QBoxPlotSeries) Pressed(boxset QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_Pressed(ptr.Pointer(), PointerFromQBoxSet(boxset))
	}
}

//export callbackQBoxPlotSeries_Released
func callbackQBoxPlotSeries_Released(ptr unsafe.Pointer, boxset unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "released"); signal != nil {
		(*(*func(*QBoxSet))(signal))(NewQBoxSetFromPointer(boxset))
	}

}

func (ptr *QBoxPlotSeries) ConnectReleased(f func(boxset *QBoxSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "released") {
			C.QBoxPlotSeries_ConnectReleased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "released")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "released"); signal != nil {
			f := func(boxset *QBoxSet) {
				(*(*func(*QBoxSet))(signal))(boxset)
				f(boxset)
			}
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "released")
	}
}

func (ptr *QBoxPlotSeries) Released(boxset QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_Released(ptr.Pointer(), PointerFromQBoxSet(boxset))
	}
}

func (ptr *QBoxPlotSeries) Remove(set QBoxSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBoxPlotSeries_Remove(ptr.Pointer(), PointerFromQBoxSet(set))) != 0
	}
	return false
}

func (ptr *QBoxPlotSeries) SetBoxOutlineVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_SetBoxOutlineVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QBoxPlotSeries) SetBoxWidth(width float64) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_SetBoxWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QBoxPlotSeries) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QBoxPlotSeries) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QBoxPlotSeries) Take(set QBoxSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBoxPlotSeries_Take(ptr.Pointer(), PointerFromQBoxSet(set))) != 0
	}
	return false
}

//export callbackQBoxPlotSeries_Type
func callbackQBoxPlotSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQBoxPlotSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QBoxPlotSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QBoxPlotSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QBoxPlotSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxPlotSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QBoxPlotSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBoxPlotSeries_DestroyQBoxPlotSeries
func callbackQBoxPlotSeries_DestroyQBoxPlotSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBoxPlotSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBoxPlotSeriesFromPointer(ptr).DestroyQBoxPlotSeriesDefault()
	}
}

func (ptr *QBoxPlotSeries) ConnectDestroyQBoxPlotSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBoxPlotSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBoxPlotSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBoxPlotSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxPlotSeries) DisconnectDestroyQBoxPlotSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBoxPlotSeries")
	}
}

func (ptr *QBoxPlotSeries) DestroyQBoxPlotSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBoxPlotSeries_DestroyQBoxPlotSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBoxPlotSeries) DestroyQBoxPlotSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBoxPlotSeries_DestroyQBoxPlotSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBoxPlotSeries) __append_sets_atList2(i int) *QBoxSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQBoxSetFromPointer(C.QBoxPlotSeries___append_sets_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBoxPlotSeries) __append_sets_setList2(i QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries___append_sets_setList2(ptr.Pointer(), PointerFromQBoxSet(i))
	}
}

func (ptr *QBoxPlotSeries) __append_sets_newList2() unsafe.Pointer {
	return C.QBoxPlotSeries___append_sets_newList2(ptr.Pointer())
}

func (ptr *QBoxPlotSeries) __boxSets_atList(i int) *QBoxSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQBoxSetFromPointer(C.QBoxPlotSeries___boxSets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBoxPlotSeries) __boxSets_setList(i QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries___boxSets_setList(ptr.Pointer(), PointerFromQBoxSet(i))
	}
}

func (ptr *QBoxPlotSeries) __boxSets_newList() unsafe.Pointer {
	return C.QBoxPlotSeries___boxSets_newList(ptr.Pointer())
}

func (ptr *QBoxPlotSeries) __boxsetsAdded_sets_atList(i int) *QBoxSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQBoxSetFromPointer(C.QBoxPlotSeries___boxsetsAdded_sets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBoxPlotSeries) __boxsetsAdded_sets_setList(i QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries___boxsetsAdded_sets_setList(ptr.Pointer(), PointerFromQBoxSet(i))
	}
}

func (ptr *QBoxPlotSeries) __boxsetsAdded_sets_newList() unsafe.Pointer {
	return C.QBoxPlotSeries___boxsetsAdded_sets_newList(ptr.Pointer())
}

func (ptr *QBoxPlotSeries) __boxsetsRemoved_sets_atList(i int) *QBoxSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQBoxSetFromPointer(C.QBoxPlotSeries___boxsetsRemoved_sets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBoxPlotSeries) __boxsetsRemoved_sets_setList(i QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries___boxsetsRemoved_sets_setList(ptr.Pointer(), PointerFromQBoxSet(i))
	}
}

func (ptr *QBoxPlotSeries) __boxsetsRemoved_sets_newList() unsafe.Pointer {
	return C.QBoxPlotSeries___boxsetsRemoved_sets_newList(ptr.Pointer())
}

type QBoxSet struct {
	core.QObject
}

type QBoxSet_ITF interface {
	core.QObject_ITF
	QBoxSet_PTR() *QBoxSet
}

func (ptr *QBoxSet) QBoxSet_PTR() *QBoxSet {
	return ptr
}

func (ptr *QBoxSet) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBoxSet) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQBoxSet(ptr QBoxSet_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxSet_PTR().Pointer()
	}
	return nil
}

func NewQBoxSetFromPointer(ptr unsafe.Pointer) (n *QBoxSet) {
	n = new(QBoxSet)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QBoxSet__ValuePositions
//QBoxSet::ValuePositions
type QBoxSet__ValuePositions int64

const (
	QBoxSet__LowerExtreme  QBoxSet__ValuePositions = QBoxSet__ValuePositions(0)
	QBoxSet__LowerQuartile QBoxSet__ValuePositions = QBoxSet__ValuePositions(1)
	QBoxSet__Median        QBoxSet__ValuePositions = QBoxSet__ValuePositions(2)
	QBoxSet__UpperQuartile QBoxSet__ValuePositions = QBoxSet__ValuePositions(3)
	QBoxSet__UpperExtreme  QBoxSet__ValuePositions = QBoxSet__ValuePositions(4)
)

func NewQBoxSet(label string, parent core.QObject_ITF) *QBoxSet {
	var labelC *C.char
	if label != "" {
		labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
	}
	tmpValue := NewQBoxSetFromPointer(C.QBoxSet_NewQBoxSet(C.struct_QtCharts_PackedString{data: labelC, len: C.longlong(len(label))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBoxSet2(le float64, lq float64, m float64, uq float64, ue float64, label string, parent core.QObject_ITF) *QBoxSet {
	var labelC *C.char
	if label != "" {
		labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
	}
	tmpValue := NewQBoxSetFromPointer(C.QBoxSet_NewQBoxSet2(C.double(le), C.double(lq), C.double(m), C.double(uq), C.double(ue), C.struct_QtCharts_PackedString{data: labelC, len: C.longlong(len(label))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBoxSet) Append(value float64) {
	if ptr.Pointer() != nil {
		C.QBoxSet_Append(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QBoxSet) Append2(values []float64) {
	if ptr.Pointer() != nil {
		C.QBoxSet_Append2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQBoxSetFromPointer(NewQBoxSetFromPointer(nil).__append_values_newList2())
			for _, v := range values {
				tmpList.__append_values_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QBoxSet) At(index int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QBoxSet_At(ptr.Pointer(), C.int(int32(index))))
	}
	return 0
}

func (ptr *QBoxSet) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QBoxSet_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQBoxSet_BrushChanged
func callbackQBoxSet_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "brushChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxSet) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "brushChanged") {
			C.QBoxSet_ConnectBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "brushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "brushChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "brushChanged")
	}
}

func (ptr *QBoxSet) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_BrushChanged(ptr.Pointer())
	}
}

func (ptr *QBoxSet) Clear() {
	if ptr.Pointer() != nil {
		C.QBoxSet_Clear(ptr.Pointer())
	}
}

//export callbackQBoxSet_Cleared
func callbackQBoxSet_Cleared(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cleared"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxSet) ConnectCleared(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "cleared") {
			C.QBoxSet_ConnectCleared(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "cleared")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "cleared"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "cleared", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cleared", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectCleared() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectCleared(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "cleared")
	}
}

func (ptr *QBoxSet) Cleared() {
	if ptr.Pointer() != nil {
		C.QBoxSet_Cleared(ptr.Pointer())
	}
}

//export callbackQBoxSet_Clicked
func callbackQBoxSet_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxSet) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QBoxSet_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QBoxSet) Clicked() {
	if ptr.Pointer() != nil {
		C.QBoxSet_Clicked(ptr.Pointer())
	}
}

func (ptr *QBoxSet) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBoxSet_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQBoxSet_DoubleClicked
func callbackQBoxSet_DoubleClicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxSet) ConnectDoubleClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "doubleClicked") {
			C.QBoxSet_ConnectDoubleClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "doubleClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "doubleClicked"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "doubleClicked")
	}
}

func (ptr *QBoxSet) DoubleClicked() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DoubleClicked(ptr.Pointer())
	}
}

//export callbackQBoxSet_Hovered
func callbackQBoxSet_Hovered(ptr unsafe.Pointer, status C.char) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(bool))(signal))(int8(status) != 0)
	}

}

func (ptr *QBoxSet) ConnectHovered(f func(status bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QBoxSet_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(status bool) {
				(*(*func(bool))(signal))(status)
				f(status)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QBoxSet) Hovered(status bool) {
	if ptr.Pointer() != nil {
		C.QBoxSet_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))))
	}
}

func (ptr *QBoxSet) Label() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBoxSet_Label(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBoxSet) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QBoxSet_Pen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQBoxSet_PenChanged
func callbackQBoxSet_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "penChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxSet) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "penChanged") {
			C.QBoxSet_ConnectPenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "penChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "penChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "penChanged")
	}
}

func (ptr *QBoxSet) PenChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_PenChanged(ptr.Pointer())
	}
}

//export callbackQBoxSet_Pressed
func callbackQBoxSet_Pressed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxSet) ConnectPressed(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pressed") {
			C.QBoxSet_ConnectPressed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pressed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pressed"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pressed")
	}
}

func (ptr *QBoxSet) Pressed() {
	if ptr.Pointer() != nil {
		C.QBoxSet_Pressed(ptr.Pointer())
	}
}

//export callbackQBoxSet_Released
func callbackQBoxSet_Released(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "released"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxSet) ConnectReleased(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "released") {
			C.QBoxSet_ConnectReleased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "released")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "released"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "released")
	}
}

func (ptr *QBoxSet) Released() {
	if ptr.Pointer() != nil {
		C.QBoxSet_Released(ptr.Pointer())
	}
}

func (ptr *QBoxSet) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QBoxSet) SetLabel(label string) {
	if ptr.Pointer() != nil {
		var labelC *C.char
		if label != "" {
			labelC = C.CString(label)
			defer C.free(unsafe.Pointer(labelC))
		}
		C.QBoxSet_SetLabel(ptr.Pointer(), C.struct_QtCharts_PackedString{data: labelC, len: C.longlong(len(label))})
	}
}

func (ptr *QBoxSet) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QBoxSet) SetValue(index int, value float64) {
	if ptr.Pointer() != nil {
		C.QBoxSet_SetValue(ptr.Pointer(), C.int(int32(index)), C.double(value))
	}
}

//export callbackQBoxSet_ValueChanged
func callbackQBoxSet_ValueChanged(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "valueChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QBoxSet) ConnectValueChanged(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valueChanged") {
			C.QBoxSet_ConnectValueChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "valueChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valueChanged"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valueChanged")
	}
}

func (ptr *QBoxSet) ValueChanged(index int) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ValueChanged(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQBoxSet_ValuesChanged
func callbackQBoxSet_ValuesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "valuesChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBoxSet) ConnectValuesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valuesChanged") {
			C.QBoxSet_ConnectValuesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "valuesChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valuesChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "valuesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valuesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectValuesChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectValuesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valuesChanged")
	}
}

func (ptr *QBoxSet) ValuesChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_ValuesChanged(ptr.Pointer())
	}
}

//export callbackQBoxSet_DestroyQBoxSet
func callbackQBoxSet_DestroyQBoxSet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBoxSet"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBoxSetFromPointer(ptr).DestroyQBoxSetDefault()
	}
}

func (ptr *QBoxSet) ConnectDestroyQBoxSet(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBoxSet"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBoxSet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBoxSet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxSet) DisconnectDestroyQBoxSet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBoxSet")
	}
}

func (ptr *QBoxSet) DestroyQBoxSet() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBoxSet_DestroyQBoxSet(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBoxSet) DestroyQBoxSetDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBoxSet_DestroyQBoxSetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBoxSet) __append_values_atList2(i int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QBoxSet___append_values_atList2(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QBoxSet) __append_values_setList2(i float64) {
	if ptr.Pointer() != nil {
		C.QBoxSet___append_values_setList2(ptr.Pointer(), C.double(i))
	}
}

func (ptr *QBoxSet) __append_values_newList2() unsafe.Pointer {
	return C.QBoxSet___append_values_newList2(ptr.Pointer())
}

func (ptr *QBoxSet) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBoxSet___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBoxSet) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBoxSet) __children_newList() unsafe.Pointer {
	return C.QBoxSet___children_newList(ptr.Pointer())
}

func (ptr *QBoxSet) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QBoxSet___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QBoxSet) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QBoxSet) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QBoxSet___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QBoxSet) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBoxSet___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBoxSet) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBoxSet) __findChildren_newList() unsafe.Pointer {
	return C.QBoxSet___findChildren_newList(ptr.Pointer())
}

func (ptr *QBoxSet) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBoxSet___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBoxSet) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBoxSet) __findChildren_newList3() unsafe.Pointer {
	return C.QBoxSet___findChildren_newList3(ptr.Pointer())
}

//export callbackQBoxSet_ChildEvent
func callbackQBoxSet_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBoxSetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBoxSet) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBoxSet_ConnectNotify
func callbackQBoxSet_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBoxSetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBoxSet) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBoxSet_CustomEvent
func callbackQBoxSet_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQBoxSetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBoxSet) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBoxSet_DeleteLater
func callbackQBoxSet_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBoxSetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBoxSet) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBoxSet_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQBoxSet_Destroyed
func callbackQBoxSet_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBoxSet_DisconnectNotify
func callbackQBoxSet_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBoxSetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBoxSet) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBoxSet_Event
func callbackQBoxSet_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBoxSetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBoxSet) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBoxSet_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQBoxSet_EventFilter
func callbackQBoxSet_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBoxSetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBoxSet) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBoxSet_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQBoxSet_MetaObject
func callbackQBoxSet_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQBoxSetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBoxSet) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBoxSet_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQBoxSet_ObjectNameChanged
func callbackQBoxSet_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQBoxSet_TimerEvent
func callbackQBoxSet_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBoxSetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBoxSet) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxSet_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QCandlestickLegendMarker struct {
	QLegendMarker
}

type QCandlestickLegendMarker_ITF interface {
	QLegendMarker_ITF
	QCandlestickLegendMarker_PTR() *QCandlestickLegendMarker
}

func (ptr *QCandlestickLegendMarker) QCandlestickLegendMarker_PTR() *QCandlestickLegendMarker {
	return ptr
}

func (ptr *QCandlestickLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLegendMarker_PTR().Pointer()
	}
	return nil
}

func (ptr *QCandlestickLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLegendMarker_PTR().SetPointer(p)
	}
}

func PointerFromQCandlestickLegendMarker(ptr QCandlestickLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCandlestickLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQCandlestickLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QCandlestickLegendMarker) {
	n = new(QCandlestickLegendMarker)
	n.SetPointer(ptr)
	return
}

//export callbackQCandlestickLegendMarker_Series
func callbackQCandlestickLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "series"); signal != nil {
		return PointerFromQCandlestickSeries((*(*func() *QCandlestickSeries)(signal))())
	}

	return PointerFromQCandlestickSeries(NewQCandlestickLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QCandlestickLegendMarker) ConnectSeries(f func() *QCandlestickSeries) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "series"); signal != nil {
			f := func() *QCandlestickSeries {
				(*(*func() *QCandlestickSeries)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "series")
	}
}

func (ptr *QCandlestickLegendMarker) Series() *QCandlestickSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQCandlestickSeriesFromPointer(C.QCandlestickLegendMarker_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickLegendMarker) SeriesDefault() *QCandlestickSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQCandlestickSeriesFromPointer(C.QCandlestickLegendMarker_SeriesDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQCandlestickLegendMarker_Type
func callbackQCandlestickLegendMarker_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QLegendMarker__LegendMarkerType)(signal))())
	}

	return C.longlong(NewQCandlestickLegendMarkerFromPointer(ptr).TypeDefault())
}

func (ptr *QCandlestickLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QLegendMarker__LegendMarkerType {
				(*(*func() QLegendMarker__LegendMarkerType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickLegendMarker) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QCandlestickLegendMarker) Type() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QCandlestickLegendMarker_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QCandlestickLegendMarker_TypeDefault(ptr.Pointer()))
	}
	return 0
}

type QCandlestickModelMapper struct {
	core.QObject
}

type QCandlestickModelMapper_ITF interface {
	core.QObject_ITF
	QCandlestickModelMapper_PTR() *QCandlestickModelMapper
}

func (ptr *QCandlestickModelMapper) QCandlestickModelMapper_PTR() *QCandlestickModelMapper {
	return ptr
}

func (ptr *QCandlestickModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QCandlestickModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQCandlestickModelMapper(ptr QCandlestickModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCandlestickModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQCandlestickModelMapperFromPointer(ptr unsafe.Pointer) (n *QCandlestickModelMapper) {
	n = new(QCandlestickModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQCandlestickModelMapper(parent core.QObject_ITF) *QCandlestickModelMapper {
	tmpValue := NewQCandlestickModelMapperFromPointer(C.QCandlestickModelMapper_NewQCandlestickModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QCandlestickModelMapper) Close() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickModelMapper_Close(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCandlestickModelMapper) FirstSetSection() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickModelMapper_FirstSetSection(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCandlestickModelMapper) High() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickModelMapper_High(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCandlestickModelMapper) LastSetSection() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickModelMapper_LastSetSection(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCandlestickModelMapper) Low() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickModelMapper_Low(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCandlestickModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QCandlestickModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQCandlestickModelMapper_ModelReplaced
func callbackQCandlestickModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "modelReplaced") {
			C.QCandlestickModelMapper_ConnectModelReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "modelReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "modelReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "modelReplaced")
	}
}

func (ptr *QCandlestickModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_ModelReplaced(ptr.Pointer())
	}
}

func (ptr *QCandlestickModelMapper) Open() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickModelMapper_Open(ptr.Pointer())))
	}
	return 0
}

//export callbackQCandlestickModelMapper_Orientation
func callbackQCandlestickModelMapper_Orientation(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "orientation"); signal != nil {
		return C.longlong((*(*func() core.Qt__Orientation)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QCandlestickModelMapper) ConnectOrientation(f func() core.Qt__Orientation) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "orientation"); signal != nil {
			f := func() core.Qt__Orientation {
				(*(*func() core.Qt__Orientation)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "orientation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "orientation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickModelMapper) DisconnectOrientation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "orientation")
	}
}

func (ptr *QCandlestickModelMapper) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QCandlestickModelMapper_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickModelMapper) Series() *QCandlestickSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQCandlestickSeriesFromPointer(C.QCandlestickModelMapper_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQCandlestickModelMapper_SeriesReplaced
func callbackQCandlestickModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "seriesReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "seriesReplaced") {
			C.QCandlestickModelMapper_ConnectSeriesReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "seriesReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "seriesReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "seriesReplaced")
	}
}

func (ptr *QCandlestickModelMapper) SeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_SeriesReplaced(ptr.Pointer())
	}
}

func (ptr *QCandlestickModelMapper) SetClose(close int) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_SetClose(ptr.Pointer(), C.int(int32(close)))
	}
}

func (ptr *QCandlestickModelMapper) SetFirstSetSection(firstSetSection int) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_SetFirstSetSection(ptr.Pointer(), C.int(int32(firstSetSection)))
	}
}

func (ptr *QCandlestickModelMapper) SetHigh(high int) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_SetHigh(ptr.Pointer(), C.int(int32(high)))
	}
}

func (ptr *QCandlestickModelMapper) SetLastSetSection(lastSetSection int) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_SetLastSetSection(ptr.Pointer(), C.int(int32(lastSetSection)))
	}
}

func (ptr *QCandlestickModelMapper) SetLow(low int) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_SetLow(ptr.Pointer(), C.int(int32(low)))
	}
}

func (ptr *QCandlestickModelMapper) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QCandlestickModelMapper) SetOpen(open int) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_SetOpen(ptr.Pointer(), C.int(int32(open)))
	}
}

func (ptr *QCandlestickModelMapper) SetSeries(series QCandlestickSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_SetSeries(ptr.Pointer(), PointerFromQCandlestickSeries(series))
	}
}

func (ptr *QCandlestickModelMapper) SetTimestamp(timestamp int) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_SetTimestamp(ptr.Pointer(), C.int(int32(timestamp)))
	}
}

func (ptr *QCandlestickModelMapper) Timestamp() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickModelMapper_Timestamp(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCandlestickModelMapper) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QCandlestickModelMapper___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickModelMapper) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCandlestickModelMapper) __children_newList() unsafe.Pointer {
	return C.QCandlestickModelMapper___children_newList(ptr.Pointer())
}

func (ptr *QCandlestickModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QCandlestickModelMapper___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QCandlestickModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QCandlestickModelMapper___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QCandlestickModelMapper) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QCandlestickModelMapper___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickModelMapper) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCandlestickModelMapper) __findChildren_newList() unsafe.Pointer {
	return C.QCandlestickModelMapper___findChildren_newList(ptr.Pointer())
}

func (ptr *QCandlestickModelMapper) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QCandlestickModelMapper___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickModelMapper) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCandlestickModelMapper) __findChildren_newList3() unsafe.Pointer {
	return C.QCandlestickModelMapper___findChildren_newList3(ptr.Pointer())
}

//export callbackQCandlestickModelMapper_ChildEvent
func callbackQCandlestickModelMapper_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCandlestickModelMapperFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCandlestickModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCandlestickModelMapper_ConnectNotify
func callbackQCandlestickModelMapper_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCandlestickModelMapperFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCandlestickModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCandlestickModelMapper_CustomEvent
func callbackQCandlestickModelMapper_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQCandlestickModelMapperFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCandlestickModelMapper) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCandlestickModelMapper_DeleteLater
func callbackQCandlestickModelMapper_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQCandlestickModelMapperFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCandlestickModelMapper) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCandlestickModelMapper_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQCandlestickModelMapper_Destroyed
func callbackQCandlestickModelMapper_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQCandlestickModelMapper_DisconnectNotify
func callbackQCandlestickModelMapper_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCandlestickModelMapperFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCandlestickModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCandlestickModelMapper_Event
func callbackQCandlestickModelMapper_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCandlestickModelMapperFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QCandlestickModelMapper) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickModelMapper_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQCandlestickModelMapper_EventFilter
func callbackQCandlestickModelMapper_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCandlestickModelMapperFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QCandlestickModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickModelMapper_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQCandlestickModelMapper_MetaObject
func callbackQCandlestickModelMapper_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQCandlestickModelMapperFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QCandlestickModelMapper) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCandlestickModelMapper_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQCandlestickModelMapper_ObjectNameChanged
func callbackQCandlestickModelMapper_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQCandlestickModelMapper_TimerEvent
func callbackQCandlestickModelMapper_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCandlestickModelMapperFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCandlestickModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QCandlestickSeries struct {
	QAbstractSeries
}

type QCandlestickSeries_ITF interface {
	QAbstractSeries_ITF
	QCandlestickSeries_PTR() *QCandlestickSeries
}

func (ptr *QCandlestickSeries) QCandlestickSeries_PTR() *QCandlestickSeries {
	return ptr
}

func (ptr *QCandlestickSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QCandlestickSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSeries_PTR().SetPointer(p)
	}
}

func PointerFromQCandlestickSeries(ptr QCandlestickSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCandlestickSeries_PTR().Pointer()
	}
	return nil
}

func NewQCandlestickSeriesFromPointer(ptr unsafe.Pointer) (n *QCandlestickSeries) {
	n = new(QCandlestickSeries)
	n.SetPointer(ptr)
	return
}
func NewQCandlestickSeries(parent core.QObject_ITF) *QCandlestickSeries {
	tmpValue := NewQCandlestickSeriesFromPointer(C.QCandlestickSeries_NewQCandlestickSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QCandlestickSeries) Append(set QCandlestickSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickSeries_Append(ptr.Pointer(), PointerFromQCandlestickSet(set))) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) Append2(sets []*QCandlestickSet) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickSeries_Append2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQCandlestickSeriesFromPointer(NewQCandlestickSeriesFromPointer(nil).__append_sets_newList2())
			for _, v := range sets {
				tmpList.__append_sets_setList2(v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackQCandlestickSeries_BodyOutlineVisibilityChanged
func callbackQCandlestickSeries_BodyOutlineVisibilityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "bodyOutlineVisibilityChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectBodyOutlineVisibilityChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "bodyOutlineVisibilityChanged") {
			C.QCandlestickSeries_ConnectBodyOutlineVisibilityChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "bodyOutlineVisibilityChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "bodyOutlineVisibilityChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "bodyOutlineVisibilityChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "bodyOutlineVisibilityChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectBodyOutlineVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectBodyOutlineVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "bodyOutlineVisibilityChanged")
	}
}

func (ptr *QCandlestickSeries) BodyOutlineVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_BodyOutlineVisibilityChanged(ptr.Pointer())
	}
}

func (ptr *QCandlestickSeries) BodyOutlineVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickSeries_BodyOutlineVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) BodyWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSeries_BodyWidth(ptr.Pointer()))
	}
	return 0
}

//export callbackQCandlestickSeries_BodyWidthChanged
func callbackQCandlestickSeries_BodyWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "bodyWidthChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectBodyWidthChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "bodyWidthChanged") {
			C.QCandlestickSeries_ConnectBodyWidthChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "bodyWidthChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "bodyWidthChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "bodyWidthChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "bodyWidthChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectBodyWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectBodyWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "bodyWidthChanged")
	}
}

func (ptr *QCandlestickSeries) BodyWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_BodyWidthChanged(ptr.Pointer())
	}
}

func (ptr *QCandlestickSeries) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QCandlestickSeries_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQCandlestickSeries_BrushChanged
func callbackQCandlestickSeries_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "brushChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "brushChanged") {
			C.QCandlestickSeries_ConnectBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "brushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "brushChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "brushChanged")
	}
}

func (ptr *QCandlestickSeries) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_BrushChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_CandlestickSetsAdded
func callbackQCandlestickSeries_CandlestickSetsAdded(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(ptr, "candlestickSetsAdded"); signal != nil {
		(*(*func([]*QCandlestickSet))(signal))(func(l C.struct_QtCharts_PackedList) []*QCandlestickSet {
			out := make([]*QCandlestickSet, int(l.len))
			tmpList := NewQCandlestickSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__candlestickSetsAdded_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QCandlestickSeries) ConnectCandlestickSetsAdded(f func(sets []*QCandlestickSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "candlestickSetsAdded") {
			C.QCandlestickSeries_ConnectCandlestickSetsAdded(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "candlestickSetsAdded")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "candlestickSetsAdded"); signal != nil {
			f := func(sets []*QCandlestickSet) {
				(*(*func([]*QCandlestickSet))(signal))(sets)
				f(sets)
			}
			qt.ConnectSignal(ptr.Pointer(), "candlestickSetsAdded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "candlestickSetsAdded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectCandlestickSetsAdded() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectCandlestickSetsAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "candlestickSetsAdded")
	}
}

func (ptr *QCandlestickSeries) CandlestickSetsAdded(sets []*QCandlestickSet) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_CandlestickSetsAdded(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQCandlestickSeriesFromPointer(NewQCandlestickSeriesFromPointer(nil).__candlestickSetsAdded_sets_newList())
			for _, v := range sets {
				tmpList.__candlestickSetsAdded_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQCandlestickSeries_CandlestickSetsRemoved
func callbackQCandlestickSeries_CandlestickSetsRemoved(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(ptr, "candlestickSetsRemoved"); signal != nil {
		(*(*func([]*QCandlestickSet))(signal))(func(l C.struct_QtCharts_PackedList) []*QCandlestickSet {
			out := make([]*QCandlestickSet, int(l.len))
			tmpList := NewQCandlestickSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__candlestickSetsRemoved_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QCandlestickSeries) ConnectCandlestickSetsRemoved(f func(sets []*QCandlestickSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "candlestickSetsRemoved") {
			C.QCandlestickSeries_ConnectCandlestickSetsRemoved(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "candlestickSetsRemoved")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "candlestickSetsRemoved"); signal != nil {
			f := func(sets []*QCandlestickSet) {
				(*(*func([]*QCandlestickSet))(signal))(sets)
				f(sets)
			}
			qt.ConnectSignal(ptr.Pointer(), "candlestickSetsRemoved", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "candlestickSetsRemoved", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectCandlestickSetsRemoved() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectCandlestickSetsRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "candlestickSetsRemoved")
	}
}

func (ptr *QCandlestickSeries) CandlestickSetsRemoved(sets []*QCandlestickSet) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_CandlestickSetsRemoved(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQCandlestickSeriesFromPointer(NewQCandlestickSeriesFromPointer(nil).__candlestickSetsRemoved_sets_newList())
			for _, v := range sets {
				tmpList.__candlestickSetsRemoved_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQCandlestickSeries_CapsVisibilityChanged
func callbackQCandlestickSeries_CapsVisibilityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "capsVisibilityChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectCapsVisibilityChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "capsVisibilityChanged") {
			C.QCandlestickSeries_ConnectCapsVisibilityChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "capsVisibilityChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "capsVisibilityChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "capsVisibilityChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "capsVisibilityChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectCapsVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectCapsVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "capsVisibilityChanged")
	}
}

func (ptr *QCandlestickSeries) CapsVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_CapsVisibilityChanged(ptr.Pointer())
	}
}

func (ptr *QCandlestickSeries) CapsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickSeries_CapsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) CapsWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSeries_CapsWidth(ptr.Pointer()))
	}
	return 0
}

//export callbackQCandlestickSeries_CapsWidthChanged
func callbackQCandlestickSeries_CapsWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "capsWidthChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectCapsWidthChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "capsWidthChanged") {
			C.QCandlestickSeries_ConnectCapsWidthChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "capsWidthChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "capsWidthChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "capsWidthChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "capsWidthChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectCapsWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectCapsWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "capsWidthChanged")
	}
}

func (ptr *QCandlestickSeries) CapsWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_CapsWidthChanged(ptr.Pointer())
	}
}

func (ptr *QCandlestickSeries) Clear() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_Clear(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_Clicked
func callbackQCandlestickSeries_Clicked(ptr unsafe.Pointer, set unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(*QCandlestickSet))(signal))(NewQCandlestickSetFromPointer(set))
	}

}

func (ptr *QCandlestickSeries) ConnectClicked(f func(set *QCandlestickSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QCandlestickSeries_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(set *QCandlestickSet) {
				(*(*func(*QCandlestickSet))(signal))(set)
				f(set)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QCandlestickSeries) Clicked(set QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_Clicked(ptr.Pointer(), PointerFromQCandlestickSet(set))
	}
}

func (ptr *QCandlestickSeries) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickSeries_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQCandlestickSeries_CountChanged
func callbackQCandlestickSeries_CountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "countChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "countChanged") {
			C.QCandlestickSeries_ConnectCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "countChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "countChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "countChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "countChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectCountChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "countChanged")
	}
}

func (ptr *QCandlestickSeries) CountChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_CountChanged(ptr.Pointer())
	}
}

func (ptr *QCandlestickSeries) DecreasingColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QCandlestickSeries_DecreasingColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQCandlestickSeries_DecreasingColorChanged
func callbackQCandlestickSeries_DecreasingColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "decreasingColorChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectDecreasingColorChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "decreasingColorChanged") {
			C.QCandlestickSeries_ConnectDecreasingColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "decreasingColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "decreasingColorChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "decreasingColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "decreasingColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectDecreasingColorChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectDecreasingColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "decreasingColorChanged")
	}
}

func (ptr *QCandlestickSeries) DecreasingColorChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DecreasingColorChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_DoubleClicked
func callbackQCandlestickSeries_DoubleClicked(ptr unsafe.Pointer, set unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		(*(*func(*QCandlestickSet))(signal))(NewQCandlestickSetFromPointer(set))
	}

}

func (ptr *QCandlestickSeries) ConnectDoubleClicked(f func(set *QCandlestickSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "doubleClicked") {
			C.QCandlestickSeries_ConnectDoubleClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "doubleClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "doubleClicked"); signal != nil {
			f := func(set *QCandlestickSet) {
				(*(*func(*QCandlestickSet))(signal))(set)
				f(set)
			}
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "doubleClicked")
	}
}

func (ptr *QCandlestickSeries) DoubleClicked(set QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DoubleClicked(ptr.Pointer(), PointerFromQCandlestickSet(set))
	}
}

//export callbackQCandlestickSeries_Hovered
func callbackQCandlestickSeries_Hovered(ptr unsafe.Pointer, status C.char, set unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(bool, *QCandlestickSet))(signal))(int8(status) != 0, NewQCandlestickSetFromPointer(set))
	}

}

func (ptr *QCandlestickSeries) ConnectHovered(f func(status bool, set *QCandlestickSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QCandlestickSeries_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(status bool, set *QCandlestickSet) {
				(*(*func(bool, *QCandlestickSet))(signal))(status, set)
				f(status, set)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QCandlestickSeries) Hovered(status bool, set QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))), PointerFromQCandlestickSet(set))
	}
}

func (ptr *QCandlestickSeries) IncreasingColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QCandlestickSeries_IncreasingColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQCandlestickSeries_IncreasingColorChanged
func callbackQCandlestickSeries_IncreasingColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "increasingColorChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectIncreasingColorChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "increasingColorChanged") {
			C.QCandlestickSeries_ConnectIncreasingColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "increasingColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "increasingColorChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "increasingColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "increasingColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectIncreasingColorChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectIncreasingColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "increasingColorChanged")
	}
}

func (ptr *QCandlestickSeries) IncreasingColorChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_IncreasingColorChanged(ptr.Pointer())
	}
}

func (ptr *QCandlestickSeries) Insert(index int, set QCandlestickSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickSeries_Insert(ptr.Pointer(), C.int(int32(index)), PointerFromQCandlestickSet(set))) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) MaximumColumnWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSeries_MaximumColumnWidth(ptr.Pointer()))
	}
	return 0
}

//export callbackQCandlestickSeries_MaximumColumnWidthChanged
func callbackQCandlestickSeries_MaximumColumnWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "maximumColumnWidthChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectMaximumColumnWidthChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "maximumColumnWidthChanged") {
			C.QCandlestickSeries_ConnectMaximumColumnWidthChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "maximumColumnWidthChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "maximumColumnWidthChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "maximumColumnWidthChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maximumColumnWidthChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectMaximumColumnWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectMaximumColumnWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "maximumColumnWidthChanged")
	}
}

func (ptr *QCandlestickSeries) MaximumColumnWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_MaximumColumnWidthChanged(ptr.Pointer())
	}
}

func (ptr *QCandlestickSeries) MinimumColumnWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSeries_MinimumColumnWidth(ptr.Pointer()))
	}
	return 0
}

//export callbackQCandlestickSeries_MinimumColumnWidthChanged
func callbackQCandlestickSeries_MinimumColumnWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "minimumColumnWidthChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectMinimumColumnWidthChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minimumColumnWidthChanged") {
			C.QCandlestickSeries_ConnectMinimumColumnWidthChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minimumColumnWidthChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minimumColumnWidthChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumColumnWidthChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumColumnWidthChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectMinimumColumnWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectMinimumColumnWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minimumColumnWidthChanged")
	}
}

func (ptr *QCandlestickSeries) MinimumColumnWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_MinimumColumnWidthChanged(ptr.Pointer())
	}
}

func (ptr *QCandlestickSeries) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QCandlestickSeries_Pen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQCandlestickSeries_PenChanged
func callbackQCandlestickSeries_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "penChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSeries) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "penChanged") {
			C.QCandlestickSeries_ConnectPenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "penChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "penChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "penChanged")
	}
}

func (ptr *QCandlestickSeries) PenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_PenChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_Pressed
func callbackQCandlestickSeries_Pressed(ptr unsafe.Pointer, set unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func(*QCandlestickSet))(signal))(NewQCandlestickSetFromPointer(set))
	}

}

func (ptr *QCandlestickSeries) ConnectPressed(f func(set *QCandlestickSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pressed") {
			C.QCandlestickSeries_ConnectPressed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pressed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pressed"); signal != nil {
			f := func(set *QCandlestickSet) {
				(*(*func(*QCandlestickSet))(signal))(set)
				f(set)
			}
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pressed")
	}
}

func (ptr *QCandlestickSeries) Pressed(set QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_Pressed(ptr.Pointer(), PointerFromQCandlestickSet(set))
	}
}

//export callbackQCandlestickSeries_Released
func callbackQCandlestickSeries_Released(ptr unsafe.Pointer, set unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "released"); signal != nil {
		(*(*func(*QCandlestickSet))(signal))(NewQCandlestickSetFromPointer(set))
	}

}

func (ptr *QCandlestickSeries) ConnectReleased(f func(set *QCandlestickSet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "released") {
			C.QCandlestickSeries_ConnectReleased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "released")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "released"); signal != nil {
			f := func(set *QCandlestickSet) {
				(*(*func(*QCandlestickSet))(signal))(set)
				f(set)
			}
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "released")
	}
}

func (ptr *QCandlestickSeries) Released(set QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_Released(ptr.Pointer(), PointerFromQCandlestickSet(set))
	}
}

func (ptr *QCandlestickSeries) Remove(set QCandlestickSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickSeries_Remove(ptr.Pointer(), PointerFromQCandlestickSet(set))) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) Remove2(sets []*QCandlestickSet) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickSeries_Remove2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQCandlestickSeriesFromPointer(NewQCandlestickSeriesFromPointer(nil).__remove_sets_newList2())
			for _, v := range sets {
				tmpList.__remove_sets_setList2(v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) SetBodyOutlineVisible(bodyOutlineVisible bool) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_SetBodyOutlineVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(bodyOutlineVisible))))
	}
}

func (ptr *QCandlestickSeries) SetBodyWidth(bodyWidth float64) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_SetBodyWidth(ptr.Pointer(), C.double(bodyWidth))
	}
}

func (ptr *QCandlestickSeries) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QCandlestickSeries) SetCapsVisible(capsVisible bool) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_SetCapsVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(capsVisible))))
	}
}

func (ptr *QCandlestickSeries) SetCapsWidth(capsWidth float64) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_SetCapsWidth(ptr.Pointer(), C.double(capsWidth))
	}
}

func (ptr *QCandlestickSeries) SetDecreasingColor(decreasingColor gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_SetDecreasingColor(ptr.Pointer(), gui.PointerFromQColor(decreasingColor))
	}
}

func (ptr *QCandlestickSeries) SetIncreasingColor(increasingColor gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_SetIncreasingColor(ptr.Pointer(), gui.PointerFromQColor(increasingColor))
	}
}

func (ptr *QCandlestickSeries) SetMaximumColumnWidth(maximumColumnWidth float64) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_SetMaximumColumnWidth(ptr.Pointer(), C.double(maximumColumnWidth))
	}
}

func (ptr *QCandlestickSeries) SetMinimumColumnWidth(minimumColumnWidth float64) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_SetMinimumColumnWidth(ptr.Pointer(), C.double(minimumColumnWidth))
	}
}

func (ptr *QCandlestickSeries) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QCandlestickSeries) Sets() []*QCandlestickSet {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QCandlestickSet {
			out := make([]*QCandlestickSet, int(l.len))
			tmpList := NewQCandlestickSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sets_atList(i)
			}
			return out
		}(C.QCandlestickSeries_Sets(ptr.Pointer()))
	}
	return make([]*QCandlestickSet, 0)
}

func (ptr *QCandlestickSeries) Take(set QCandlestickSet_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickSeries_Take(ptr.Pointer(), PointerFromQCandlestickSet(set))) != 0
	}
	return false
}

//export callbackQCandlestickSeries_Type
func callbackQCandlestickSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQCandlestickSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QCandlestickSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QCandlestickSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QCandlestickSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QCandlestickSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQCandlestickSeries_DestroyQCandlestickSeries
func callbackQCandlestickSeries_DestroyQCandlestickSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QCandlestickSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQCandlestickSeriesFromPointer(ptr).DestroyQCandlestickSeriesDefault()
	}
}

func (ptr *QCandlestickSeries) ConnectDestroyQCandlestickSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QCandlestickSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QCandlestickSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QCandlestickSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSeries) DisconnectDestroyQCandlestickSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QCandlestickSeries")
	}
}

func (ptr *QCandlestickSeries) DestroyQCandlestickSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCandlestickSeries_DestroyQCandlestickSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCandlestickSeries) DestroyQCandlestickSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCandlestickSeries_DestroyQCandlestickSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCandlestickSeries) __append_sets_atList2(i int) *QCandlestickSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQCandlestickSetFromPointer(C.QCandlestickSeries___append_sets_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSeries) __append_sets_setList2(i QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries___append_sets_setList2(ptr.Pointer(), PointerFromQCandlestickSet(i))
	}
}

func (ptr *QCandlestickSeries) __append_sets_newList2() unsafe.Pointer {
	return C.QCandlestickSeries___append_sets_newList2(ptr.Pointer())
}

func (ptr *QCandlestickSeries) __candlestickSetsAdded_sets_atList(i int) *QCandlestickSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQCandlestickSetFromPointer(C.QCandlestickSeries___candlestickSetsAdded_sets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSeries) __candlestickSetsAdded_sets_setList(i QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries___candlestickSetsAdded_sets_setList(ptr.Pointer(), PointerFromQCandlestickSet(i))
	}
}

func (ptr *QCandlestickSeries) __candlestickSetsAdded_sets_newList() unsafe.Pointer {
	return C.QCandlestickSeries___candlestickSetsAdded_sets_newList(ptr.Pointer())
}

func (ptr *QCandlestickSeries) __candlestickSetsRemoved_sets_atList(i int) *QCandlestickSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQCandlestickSetFromPointer(C.QCandlestickSeries___candlestickSetsRemoved_sets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSeries) __candlestickSetsRemoved_sets_setList(i QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries___candlestickSetsRemoved_sets_setList(ptr.Pointer(), PointerFromQCandlestickSet(i))
	}
}

func (ptr *QCandlestickSeries) __candlestickSetsRemoved_sets_newList() unsafe.Pointer {
	return C.QCandlestickSeries___candlestickSetsRemoved_sets_newList(ptr.Pointer())
}

func (ptr *QCandlestickSeries) __remove_sets_atList2(i int) *QCandlestickSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQCandlestickSetFromPointer(C.QCandlestickSeries___remove_sets_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSeries) __remove_sets_setList2(i QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries___remove_sets_setList2(ptr.Pointer(), PointerFromQCandlestickSet(i))
	}
}

func (ptr *QCandlestickSeries) __remove_sets_newList2() unsafe.Pointer {
	return C.QCandlestickSeries___remove_sets_newList2(ptr.Pointer())
}

func (ptr *QCandlestickSeries) __sets_atList(i int) *QCandlestickSet {
	if ptr.Pointer() != nil {
		tmpValue := NewQCandlestickSetFromPointer(C.QCandlestickSeries___sets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSeries) __sets_setList(i QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries___sets_setList(ptr.Pointer(), PointerFromQCandlestickSet(i))
	}
}

func (ptr *QCandlestickSeries) __sets_newList() unsafe.Pointer {
	return C.QCandlestickSeries___sets_newList(ptr.Pointer())
}

type QCandlestickSet struct {
	core.QObject
}

type QCandlestickSet_ITF interface {
	core.QObject_ITF
	QCandlestickSet_PTR() *QCandlestickSet
}

func (ptr *QCandlestickSet) QCandlestickSet_PTR() *QCandlestickSet {
	return ptr
}

func (ptr *QCandlestickSet) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QCandlestickSet) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQCandlestickSet(ptr QCandlestickSet_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCandlestickSet_PTR().Pointer()
	}
	return nil
}

func NewQCandlestickSetFromPointer(ptr unsafe.Pointer) (n *QCandlestickSet) {
	n = new(QCandlestickSet)
	n.SetPointer(ptr)
	return
}
func NewQCandlestickSet(timestamp float64, parent core.QObject_ITF) *QCandlestickSet {
	tmpValue := NewQCandlestickSetFromPointer(C.QCandlestickSet_NewQCandlestickSet(C.double(timestamp), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQCandlestickSet2(open float64, high float64, low float64, close float64, timestamp float64, parent core.QObject_ITF) *QCandlestickSet {
	tmpValue := NewQCandlestickSetFromPointer(C.QCandlestickSet_NewQCandlestickSet2(C.double(open), C.double(high), C.double(low), C.double(close), C.double(timestamp), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QCandlestickSet) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QCandlestickSet_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQCandlestickSet_BrushChanged
func callbackQCandlestickSet_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "brushChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "brushChanged") {
			C.QCandlestickSet_ConnectBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "brushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "brushChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "brushChanged")
	}
}

func (ptr *QCandlestickSet) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_BrushChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_Clicked
func callbackQCandlestickSet_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QCandlestickSet_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QCandlestickSet) Clicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_Clicked(ptr.Pointer())
	}
}

func (ptr *QCandlestickSet) Close() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSet_Close(ptr.Pointer()))
	}
	return 0
}

//export callbackQCandlestickSet_CloseChanged
func callbackQCandlestickSet_CloseChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectCloseChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "closeChanged") {
			C.QCandlestickSet_ConnectCloseChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "closeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "closeChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "closeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "closeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectCloseChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectCloseChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "closeChanged")
	}
}

func (ptr *QCandlestickSet) CloseChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_CloseChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_DoubleClicked
func callbackQCandlestickSet_DoubleClicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectDoubleClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "doubleClicked") {
			C.QCandlestickSet_ConnectDoubleClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "doubleClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "doubleClicked"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "doubleClicked")
	}
}

func (ptr *QCandlestickSet) DoubleClicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DoubleClicked(ptr.Pointer())
	}
}

func (ptr *QCandlestickSet) High() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSet_High(ptr.Pointer()))
	}
	return 0
}

//export callbackQCandlestickSet_HighChanged
func callbackQCandlestickSet_HighChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "highChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectHighChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "highChanged") {
			C.QCandlestickSet_ConnectHighChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "highChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "highChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "highChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "highChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectHighChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectHighChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "highChanged")
	}
}

func (ptr *QCandlestickSet) HighChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_HighChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_Hovered
func callbackQCandlestickSet_Hovered(ptr unsafe.Pointer, status C.char) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(bool))(signal))(int8(status) != 0)
	}

}

func (ptr *QCandlestickSet) ConnectHovered(f func(status bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QCandlestickSet_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(status bool) {
				(*(*func(bool))(signal))(status)
				f(status)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QCandlestickSet) Hovered(status bool) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))))
	}
}

func (ptr *QCandlestickSet) Low() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSet_Low(ptr.Pointer()))
	}
	return 0
}

//export callbackQCandlestickSet_LowChanged
func callbackQCandlestickSet_LowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectLowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lowChanged") {
			C.QCandlestickSet_ConnectLowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectLowChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectLowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lowChanged")
	}
}

func (ptr *QCandlestickSet) LowChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_LowChanged(ptr.Pointer())
	}
}

func (ptr *QCandlestickSet) Open() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSet_Open(ptr.Pointer()))
	}
	return 0
}

//export callbackQCandlestickSet_OpenChanged
func callbackQCandlestickSet_OpenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "openChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectOpenChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "openChanged") {
			C.QCandlestickSet_ConnectOpenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "openChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "openChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "openChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "openChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectOpenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectOpenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "openChanged")
	}
}

func (ptr *QCandlestickSet) OpenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_OpenChanged(ptr.Pointer())
	}
}

func (ptr *QCandlestickSet) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QCandlestickSet_Pen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQCandlestickSet_PenChanged
func callbackQCandlestickSet_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "penChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "penChanged") {
			C.QCandlestickSet_ConnectPenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "penChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "penChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "penChanged")
	}
}

func (ptr *QCandlestickSet) PenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_PenChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_Pressed
func callbackQCandlestickSet_Pressed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectPressed(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pressed") {
			C.QCandlestickSet_ConnectPressed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pressed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pressed"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pressed")
	}
}

func (ptr *QCandlestickSet) Pressed() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_Pressed(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_Released
func callbackQCandlestickSet_Released(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "released"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectReleased(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "released") {
			C.QCandlestickSet_ConnectReleased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "released")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "released"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "released")
	}
}

func (ptr *QCandlestickSet) Released() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_Released(ptr.Pointer())
	}
}

func (ptr *QCandlestickSet) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QCandlestickSet) SetClose(close float64) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_SetClose(ptr.Pointer(), C.double(close))
	}
}

func (ptr *QCandlestickSet) SetHigh(high float64) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_SetHigh(ptr.Pointer(), C.double(high))
	}
}

func (ptr *QCandlestickSet) SetLow(low float64) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_SetLow(ptr.Pointer(), C.double(low))
	}
}

func (ptr *QCandlestickSet) SetOpen(open float64) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_SetOpen(ptr.Pointer(), C.double(open))
	}
}

func (ptr *QCandlestickSet) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QCandlestickSet) SetTimestamp(timestamp float64) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_SetTimestamp(ptr.Pointer(), C.double(timestamp))
	}
}

func (ptr *QCandlestickSet) Timestamp() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSet_Timestamp(ptr.Pointer()))
	}
	return 0
}

//export callbackQCandlestickSet_TimestampChanged
func callbackQCandlestickSet_TimestampChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timestampChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCandlestickSet) ConnectTimestampChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "timestampChanged") {
			C.QCandlestickSet_ConnectTimestampChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "timestampChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "timestampChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "timestampChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "timestampChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectTimestampChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectTimestampChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "timestampChanged")
	}
}

func (ptr *QCandlestickSet) TimestampChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_TimestampChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_DestroyQCandlestickSet
func callbackQCandlestickSet_DestroyQCandlestickSet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QCandlestickSet"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQCandlestickSetFromPointer(ptr).DestroyQCandlestickSetDefault()
	}
}

func (ptr *QCandlestickSet) ConnectDestroyQCandlestickSet(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QCandlestickSet"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QCandlestickSet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QCandlestickSet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCandlestickSet) DisconnectDestroyQCandlestickSet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QCandlestickSet")
	}
}

func (ptr *QCandlestickSet) DestroyQCandlestickSet() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCandlestickSet_DestroyQCandlestickSet(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCandlestickSet) DestroyQCandlestickSetDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCandlestickSet_DestroyQCandlestickSetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCandlestickSet) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QCandlestickSet___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSet) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCandlestickSet) __children_newList() unsafe.Pointer {
	return C.QCandlestickSet___children_newList(ptr.Pointer())
}

func (ptr *QCandlestickSet) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QCandlestickSet___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSet) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QCandlestickSet) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QCandlestickSet___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QCandlestickSet) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QCandlestickSet___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSet) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCandlestickSet) __findChildren_newList() unsafe.Pointer {
	return C.QCandlestickSet___findChildren_newList(ptr.Pointer())
}

func (ptr *QCandlestickSet) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QCandlestickSet___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSet) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCandlestickSet) __findChildren_newList3() unsafe.Pointer {
	return C.QCandlestickSet___findChildren_newList3(ptr.Pointer())
}

//export callbackQCandlestickSet_ChildEvent
func callbackQCandlestickSet_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCandlestickSetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCandlestickSet) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCandlestickSet_ConnectNotify
func callbackQCandlestickSet_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCandlestickSetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCandlestickSet) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCandlestickSet_CustomEvent
func callbackQCandlestickSet_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQCandlestickSetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCandlestickSet) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCandlestickSet_DeleteLater
func callbackQCandlestickSet_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQCandlestickSetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCandlestickSet) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCandlestickSet_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_Destroyed
func callbackQCandlestickSet_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQCandlestickSet_DisconnectNotify
func callbackQCandlestickSet_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCandlestickSetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCandlestickSet) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCandlestickSet_Event
func callbackQCandlestickSet_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCandlestickSetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QCandlestickSet) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickSet_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQCandlestickSet_EventFilter
func callbackQCandlestickSet_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCandlestickSetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QCandlestickSet) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCandlestickSet_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQCandlestickSet_MetaObject
func callbackQCandlestickSet_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQCandlestickSetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QCandlestickSet) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCandlestickSet_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQCandlestickSet_ObjectNameChanged
func callbackQCandlestickSet_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQCandlestickSet_TimerEvent
func callbackQCandlestickSet_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCandlestickSetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCandlestickSet) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QCategoryAxis struct {
	QValueAxis
}

type QCategoryAxis_ITF interface {
	QValueAxis_ITF
	QCategoryAxis_PTR() *QCategoryAxis
}

func (ptr *QCategoryAxis) QCategoryAxis_PTR() *QCategoryAxis {
	return ptr
}

func (ptr *QCategoryAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QValueAxis_PTR().Pointer()
	}
	return nil
}

func (ptr *QCategoryAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QValueAxis_PTR().SetPointer(p)
	}
}

func PointerFromQCategoryAxis(ptr QCategoryAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCategoryAxis_PTR().Pointer()
	}
	return nil
}

func NewQCategoryAxisFromPointer(ptr unsafe.Pointer) (n *QCategoryAxis) {
	n = new(QCategoryAxis)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QCategoryAxis__AxisLabelsPosition
//QCategoryAxis::AxisLabelsPosition
type QCategoryAxis__AxisLabelsPosition int64

const (
	QCategoryAxis__AxisLabelsPositionCenter  QCategoryAxis__AxisLabelsPosition = QCategoryAxis__AxisLabelsPosition(0x0)
	QCategoryAxis__AxisLabelsPositionOnValue QCategoryAxis__AxisLabelsPosition = QCategoryAxis__AxisLabelsPosition(0x1)
)

func NewQCategoryAxis(parent core.QObject_ITF) *QCategoryAxis {
	tmpValue := NewQCategoryAxisFromPointer(C.QCategoryAxis_NewQCategoryAxis(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QCategoryAxis) Append(categoryLabel string, categoryEndValue float64) {
	if ptr.Pointer() != nil {
		var categoryLabelC *C.char
		if categoryLabel != "" {
			categoryLabelC = C.CString(categoryLabel)
			defer C.free(unsafe.Pointer(categoryLabelC))
		}
		C.QCategoryAxis_Append(ptr.Pointer(), C.struct_QtCharts_PackedString{data: categoryLabelC, len: C.longlong(len(categoryLabel))}, C.double(categoryEndValue))
	}
}

//export callbackQCategoryAxis_CategoriesChanged
func callbackQCategoryAxis_CategoriesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "categoriesChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QCategoryAxis) ConnectCategoriesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "categoriesChanged") {
			C.QCategoryAxis_ConnectCategoriesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "categoriesChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "categoriesChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "categoriesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "categoriesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCategoryAxis) DisconnectCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_DisconnectCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "categoriesChanged")
	}
}

func (ptr *QCategoryAxis) CategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_CategoriesChanged(ptr.Pointer())
	}
}

func (ptr *QCategoryAxis) CategoriesLabels() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QCategoryAxis_CategoriesLabels(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QCategoryAxis) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCategoryAxis_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCategoryAxis) EndValue(categoryLabel string) float64 {
	if ptr.Pointer() != nil {
		var categoryLabelC *C.char
		if categoryLabel != "" {
			categoryLabelC = C.CString(categoryLabel)
			defer C.free(unsafe.Pointer(categoryLabelC))
		}
		return float64(C.QCategoryAxis_EndValue(ptr.Pointer(), C.struct_QtCharts_PackedString{data: categoryLabelC, len: C.longlong(len(categoryLabel))}))
	}
	return 0
}

func (ptr *QCategoryAxis) LabelsPosition() QCategoryAxis__AxisLabelsPosition {
	if ptr.Pointer() != nil {
		return QCategoryAxis__AxisLabelsPosition(C.QCategoryAxis_LabelsPosition(ptr.Pointer()))
	}
	return 0
}

//export callbackQCategoryAxis_LabelsPositionChanged
func callbackQCategoryAxis_LabelsPositionChanged(ptr unsafe.Pointer, position C.longlong) {
	if signal := qt.GetSignal(ptr, "labelsPositionChanged"); signal != nil {
		(*(*func(QCategoryAxis__AxisLabelsPosition))(signal))(QCategoryAxis__AxisLabelsPosition(position))
	}

}

func (ptr *QCategoryAxis) ConnectLabelsPositionChanged(f func(position QCategoryAxis__AxisLabelsPosition)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsPositionChanged") {
			C.QCategoryAxis_ConnectLabelsPositionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsPositionChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsPositionChanged"); signal != nil {
			f := func(position QCategoryAxis__AxisLabelsPosition) {
				(*(*func(QCategoryAxis__AxisLabelsPosition))(signal))(position)
				f(position)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsPositionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsPositionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCategoryAxis) DisconnectLabelsPositionChanged() {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_DisconnectLabelsPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsPositionChanged")
	}
}

func (ptr *QCategoryAxis) LabelsPositionChanged(position QCategoryAxis__AxisLabelsPosition) {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_LabelsPositionChanged(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QCategoryAxis) Remove(categoryLabel string) {
	if ptr.Pointer() != nil {
		var categoryLabelC *C.char
		if categoryLabel != "" {
			categoryLabelC = C.CString(categoryLabel)
			defer C.free(unsafe.Pointer(categoryLabelC))
		}
		C.QCategoryAxis_Remove(ptr.Pointer(), C.struct_QtCharts_PackedString{data: categoryLabelC, len: C.longlong(len(categoryLabel))})
	}
}

func (ptr *QCategoryAxis) ReplaceLabel(oldLabel string, newLabel string) {
	if ptr.Pointer() != nil {
		var oldLabelC *C.char
		if oldLabel != "" {
			oldLabelC = C.CString(oldLabel)
			defer C.free(unsafe.Pointer(oldLabelC))
		}
		var newLabelC *C.char
		if newLabel != "" {
			newLabelC = C.CString(newLabel)
			defer C.free(unsafe.Pointer(newLabelC))
		}
		C.QCategoryAxis_ReplaceLabel(ptr.Pointer(), C.struct_QtCharts_PackedString{data: oldLabelC, len: C.longlong(len(oldLabel))}, C.struct_QtCharts_PackedString{data: newLabelC, len: C.longlong(len(newLabel))})
	}
}

func (ptr *QCategoryAxis) SetLabelsPosition(position QCategoryAxis__AxisLabelsPosition) {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_SetLabelsPosition(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QCategoryAxis) SetStartValue(min float64) {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_SetStartValue(ptr.Pointer(), C.double(min))
	}
}

func (ptr *QCategoryAxis) StartValue(categoryLabel string) float64 {
	if ptr.Pointer() != nil {
		var categoryLabelC *C.char
		if categoryLabel != "" {
			categoryLabelC = C.CString(categoryLabel)
			defer C.free(unsafe.Pointer(categoryLabelC))
		}
		return float64(C.QCategoryAxis_StartValue(ptr.Pointer(), C.struct_QtCharts_PackedString{data: categoryLabelC, len: C.longlong(len(categoryLabel))}))
	}
	return 0
}

//export callbackQCategoryAxis_DestroyQCategoryAxis
func callbackQCategoryAxis_DestroyQCategoryAxis(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QCategoryAxis"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQCategoryAxisFromPointer(ptr).DestroyQCategoryAxisDefault()
	}
}

func (ptr *QCategoryAxis) ConnectDestroyQCategoryAxis(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QCategoryAxis"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QCategoryAxis", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QCategoryAxis", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCategoryAxis) DisconnectDestroyQCategoryAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QCategoryAxis")
	}
}

func (ptr *QCategoryAxis) DestroyQCategoryAxis() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCategoryAxis_DestroyQCategoryAxis(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCategoryAxis) DestroyQCategoryAxisDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCategoryAxis_DestroyQCategoryAxisDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QChart struct {
	widgets.QGraphicsWidget
}

type QChart_ITF interface {
	widgets.QGraphicsWidget_ITF
	QChart_PTR() *QChart
}

func (ptr *QChart) QChart_PTR() *QChart {
	return ptr
}

func (ptr *QChart) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QChart) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsWidget_PTR().SetPointer(p)
	}
}

func PointerFromQChart(ptr QChart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QChart_PTR().Pointer()
	}
	return nil
}

func NewQChartFromPointer(ptr unsafe.Pointer) (n *QChart) {
	n = new(QChart)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QChart__ChartType
//QChart::ChartType
type QChart__ChartType int64

const (
	QChart__ChartTypeUndefined QChart__ChartType = QChart__ChartType(0)
	QChart__ChartTypeCartesian QChart__ChartType = QChart__ChartType(1)
	QChart__ChartTypePolar     QChart__ChartType = QChart__ChartType(2)
)

//go:generate stringer -type=QChart__ChartTheme
//QChart::ChartTheme
type QChart__ChartTheme int64

const (
	QChart__ChartThemeLight        QChart__ChartTheme = QChart__ChartTheme(0)
	QChart__ChartThemeBlueCerulean QChart__ChartTheme = QChart__ChartTheme(1)
	QChart__ChartThemeDark         QChart__ChartTheme = QChart__ChartTheme(2)
	QChart__ChartThemeBrownSand    QChart__ChartTheme = QChart__ChartTheme(3)
	QChart__ChartThemeBlueNcs      QChart__ChartTheme = QChart__ChartTheme(4)
	QChart__ChartThemeHighContrast QChart__ChartTheme = QChart__ChartTheme(5)
	QChart__ChartThemeBlueIcy      QChart__ChartTheme = QChart__ChartTheme(6)
	QChart__ChartThemeQt           QChart__ChartTheme = QChart__ChartTheme(7)
)

//go:generate stringer -type=QChart__AnimationOption
//QChart::AnimationOption
type QChart__AnimationOption int64

const (
	QChart__NoAnimation        QChart__AnimationOption = QChart__AnimationOption(0x0)
	QChart__GridAxisAnimations QChart__AnimationOption = QChart__AnimationOption(0x1)
	QChart__SeriesAnimations   QChart__AnimationOption = QChart__AnimationOption(0x2)
	QChart__AllAnimations      QChart__AnimationOption = QChart__AnimationOption(0x3)
)

func NewQChart(parent widgets.QGraphicsItem_ITF, wFlags core.Qt__WindowType) *QChart {
	tmpValue := NewQChartFromPointer(C.QChart_NewQChart(widgets.PointerFromQGraphicsItem(parent), C.longlong(wFlags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QChart) AddAxis(axis QAbstractAxis_ITF, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QChart_AddAxis(ptr.Pointer(), PointerFromQAbstractAxis(axis), C.longlong(alignment))
	}
}

func (ptr *QChart) AddSeries(series QAbstractSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_AddSeries(ptr.Pointer(), PointerFromQAbstractSeries(series))
	}
}

func (ptr *QChart) AnimationDuration() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QChart_AnimationDuration(ptr.Pointer())))
	}
	return 0
}

func (ptr *QChart) AnimationEasingCurve() *core.QEasingCurve {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQEasingCurveFromPointer(C.QChart_AnimationEasingCurve(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QEasingCurve).DestroyQEasingCurve)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) AnimationOptions() QChart__AnimationOption {
	if ptr.Pointer() != nil {
		return QChart__AnimationOption(C.QChart_AnimationOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChart) Axes(orientation core.Qt__Orientation, series QAbstractSeries_ITF) []*QAbstractAxis {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QAbstractAxis {
			out := make([]*QAbstractAxis, int(l.len))
			tmpList := NewQChartFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__axes_atList(i)
			}
			return out
		}(C.QChart_Axes(ptr.Pointer(), C.longlong(orientation), PointerFromQAbstractSeries(series)))
	}
	return make([]*QAbstractAxis, 0)
}

func (ptr *QChart) BackgroundBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QChart_BackgroundBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) BackgroundPen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QChart_BackgroundPen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) BackgroundRoundness() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QChart_BackgroundRoundness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChart) ChartType() QChart__ChartType {
	if ptr.Pointer() != nil {
		return QChart__ChartType(C.QChart_ChartType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChart) CreateDefaultAxes() {
	if ptr.Pointer() != nil {
		C.QChart_CreateDefaultAxes(ptr.Pointer())
	}
}

func (ptr *QChart) IsBackgroundVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_IsBackgroundVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QChart) IsDropShadowEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_IsDropShadowEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QChart) IsPlotAreaBackgroundVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_IsPlotAreaBackgroundVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QChart) IsZoomed() bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_IsZoomed(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QChart) Legend() *QLegend {
	if ptr.Pointer() != nil {
		tmpValue := NewQLegendFromPointer(C.QChart_Legend(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) Locale() *core.QLocale {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLocaleFromPointer(C.QChart_Locale(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) LocalizeNumbers() bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_LocalizeNumbers(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QChart) MapToPosition(value core.QPointF_ITF, series QAbstractSeries_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QChart_MapToPosition(ptr.Pointer(), core.PointerFromQPointF(value), PointerFromQAbstractSeries(series)))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) MapToValue(position core.QPointF_ITF, series QAbstractSeries_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QChart_MapToValue(ptr.Pointer(), core.PointerFromQPointF(position), PointerFromQAbstractSeries(series)))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) Margins() *core.QMargins {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQMarginsFromPointer(C.QChart_Margins(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QMargins).DestroyQMargins)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) PlotArea() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QChart_PlotArea(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) PlotAreaBackgroundBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QChart_PlotAreaBackgroundBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) PlotAreaBackgroundPen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QChart_PlotAreaBackgroundPen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQChart_PlotAreaChanged
func callbackQChart_PlotAreaChanged(ptr unsafe.Pointer, plotArea unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "plotAreaChanged"); signal != nil {
		(*(*func(*core.QRectF))(signal))(core.NewQRectFFromPointer(plotArea))
	}

}

func (ptr *QChart) ConnectPlotAreaChanged(f func(plotArea *core.QRectF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "plotAreaChanged") {
			C.QChart_ConnectPlotAreaChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "plotAreaChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "plotAreaChanged"); signal != nil {
			f := func(plotArea *core.QRectF) {
				(*(*func(*core.QRectF))(signal))(plotArea)
				f(plotArea)
			}
			qt.ConnectSignal(ptr.Pointer(), "plotAreaChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "plotAreaChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QChart) DisconnectPlotAreaChanged() {
	if ptr.Pointer() != nil {
		C.QChart_DisconnectPlotAreaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "plotAreaChanged")
	}
}

func (ptr *QChart) PlotAreaChanged(plotArea core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_PlotAreaChanged(ptr.Pointer(), core.PointerFromQRectF(plotArea))
	}
}

func (ptr *QChart) RemoveAllSeries() {
	if ptr.Pointer() != nil {
		C.QChart_RemoveAllSeries(ptr.Pointer())
	}
}

func (ptr *QChart) RemoveAxis(axis QAbstractAxis_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_RemoveAxis(ptr.Pointer(), PointerFromQAbstractAxis(axis))
	}
}

func (ptr *QChart) RemoveSeries(series QAbstractSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_RemoveSeries(ptr.Pointer(), PointerFromQAbstractSeries(series))
	}
}

func (ptr *QChart) Scroll(dx float64, dy float64) {
	if ptr.Pointer() != nil {
		C.QChart_Scroll(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QChart) Series() []*QAbstractSeries {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QAbstractSeries {
			out := make([]*QAbstractSeries, int(l.len))
			tmpList := NewQChartFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__series_atList(i)
			}
			return out
		}(C.QChart_Series(ptr.Pointer()))
	}
	return make([]*QAbstractSeries, 0)
}

func (ptr *QChart) SetAnimationDuration(msecs int) {
	if ptr.Pointer() != nil {
		C.QChart_SetAnimationDuration(ptr.Pointer(), C.int(int32(msecs)))
	}
}

func (ptr *QChart) SetAnimationEasingCurve(curve core.QEasingCurve_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetAnimationEasingCurve(ptr.Pointer(), core.PointerFromQEasingCurve(curve))
	}
}

func (ptr *QChart) SetAnimationOptions(options QChart__AnimationOption) {
	if ptr.Pointer() != nil {
		C.QChart_SetAnimationOptions(ptr.Pointer(), C.longlong(options))
	}
}

func (ptr *QChart) SetBackgroundBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetBackgroundBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QChart) SetBackgroundPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetBackgroundPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QChart) SetBackgroundRoundness(diameter float64) {
	if ptr.Pointer() != nil {
		C.QChart_SetBackgroundRoundness(ptr.Pointer(), C.double(diameter))
	}
}

func (ptr *QChart) SetBackgroundVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QChart_SetBackgroundVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QChart) SetDropShadowEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QChart_SetDropShadowEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QChart) SetLocale(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QChart) SetLocalizeNumbers(localize bool) {
	if ptr.Pointer() != nil {
		C.QChart_SetLocalizeNumbers(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(localize))))
	}
}

func (ptr *QChart) SetMargins(margins core.QMargins_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetMargins(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QChart) SetPlotArea(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetPlotArea(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QChart) SetPlotAreaBackgroundBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetPlotAreaBackgroundBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QChart) SetPlotAreaBackgroundPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetPlotAreaBackgroundPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QChart) SetPlotAreaBackgroundVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QChart_SetPlotAreaBackgroundVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QChart) SetTheme(theme QChart__ChartTheme) {
	if ptr.Pointer() != nil {
		C.QChart_SetTheme(ptr.Pointer(), C.longlong(theme))
	}
}

func (ptr *QChart) SetTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QChart_SetTitle(ptr.Pointer(), C.struct_QtCharts_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *QChart) SetTitleBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetTitleBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QChart) SetTitleFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetTitleFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QChart) Theme() QChart__ChartTheme {
	if ptr.Pointer() != nil {
		return QChart__ChartTheme(C.QChart_Theme(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChart) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QChart_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QChart) TitleBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QChart_TitleBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) TitleFont() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QChart_TitleFont(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) Zoom(factor float64) {
	if ptr.Pointer() != nil {
		C.QChart_Zoom(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QChart) ZoomIn() {
	if ptr.Pointer() != nil {
		C.QChart_ZoomIn(ptr.Pointer())
	}
}

func (ptr *QChart) ZoomIn2(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_ZoomIn2(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QChart) ZoomOut() {
	if ptr.Pointer() != nil {
		C.QChart_ZoomOut(ptr.Pointer())
	}
}

func (ptr *QChart) ZoomReset() {
	if ptr.Pointer() != nil {
		C.QChart_ZoomReset(ptr.Pointer())
	}
}

//export callbackQChart_DestroyQChart
func callbackQChart_DestroyQChart(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QChart"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartFromPointer(ptr).DestroyQChartDefault()
	}
}

func (ptr *QChart) ConnectDestroyQChart(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QChart"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QChart", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QChart", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QChart) DisconnectDestroyQChart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QChart")
	}
}

func (ptr *QChart) DestroyQChart() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QChart_DestroyQChart(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QChart) DestroyQChartDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QChart_DestroyQChartDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QChart) __axes_atList(i int) *QAbstractAxis {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractAxisFromPointer(C.QChart___axes_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __axes_setList(i QAbstractAxis_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___axes_setList(ptr.Pointer(), PointerFromQAbstractAxis(i))
	}
}

func (ptr *QChart) __axes_newList() unsafe.Pointer {
	return C.QChart___axes_newList(ptr.Pointer())
}

func (ptr *QChart) __series_atList(i int) *QAbstractSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractSeriesFromPointer(C.QChart___series_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __series_setList(i QAbstractSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___series_setList(ptr.Pointer(), PointerFromQAbstractSeries(i))
	}
}

func (ptr *QChart) __series_newList() unsafe.Pointer {
	return C.QChart___series_newList(ptr.Pointer())
}

func (ptr *QChart) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QChart___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QChart) __actions_newList() unsafe.Pointer {
	return C.QChart___actions_newList(ptr.Pointer())
}

func (ptr *QChart) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QChart___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QChart) __addActions_actions_newList() unsafe.Pointer {
	return C.QChart___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QChart) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QChart___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QChart) __insertActions_actions_newList() unsafe.Pointer {
	return C.QChart___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QChart) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QChart___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QChart) __children_newList() unsafe.Pointer {
	return C.QChart___children_newList(ptr.Pointer())
}

func (ptr *QChart) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QChart___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QChart) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QChart___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QChart) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QChart___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QChart) __findChildren_newList() unsafe.Pointer {
	return C.QChart___findChildren_newList(ptr.Pointer())
}

func (ptr *QChart) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QChart___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QChart) __findChildren_newList3() unsafe.Pointer {
	return C.QChart___findChildren_newList3(ptr.Pointer())
}

func (ptr *QChart) __childItems_atList(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QChart___childItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChart) __childItems_setList(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___childItems_setList(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QChart) __childItems_newList() unsafe.Pointer {
	return C.QChart___childItems_newList(ptr.Pointer())
}

func (ptr *QChart) __collidingItems_atList(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QChart___collidingItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChart) __collidingItems_setList(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___collidingItems_setList(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QChart) __collidingItems_newList() unsafe.Pointer {
	return C.QChart___collidingItems_newList(ptr.Pointer())
}

func (ptr *QChart) __setTransformations_transformations_atList(i int) *widgets.QGraphicsTransform {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQGraphicsTransformFromPointer(C.QChart___setTransformations_transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __setTransformations_transformations_setList(i widgets.QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___setTransformations_transformations_setList(ptr.Pointer(), widgets.PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QChart) __setTransformations_transformations_newList() unsafe.Pointer {
	return C.QChart___setTransformations_transformations_newList(ptr.Pointer())
}

func (ptr *QChart) __transformations_atList(i int) *widgets.QGraphicsTransform {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQGraphicsTransformFromPointer(C.QChart___transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChart) __transformations_setList(i widgets.QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___transformations_setList(ptr.Pointer(), widgets.PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QChart) __transformations_newList() unsafe.Pointer {
	return C.QChart___transformations_newList(ptr.Pointer())
}

//export callbackQChart_BoundingRect
func callbackQChart_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF((*(*func() *core.QRectF)(signal))())
	}

	return core.PointerFromQRectF(NewQChartFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QChart) BoundingRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QChart_BoundingRectDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQChart_ChangeEvent
func callbackQChart_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QChart) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQChart_Close
func callbackQChart_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).CloseDefault())))
}

func (ptr *QChart) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQChart_CloseEvent
func callbackQChart_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*gui.QCloseEvent))(signal))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QChart) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQChart_Event
func callbackQChart_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QChart) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQChart_FocusInEvent
func callbackQChart_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QChart) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQChart_FocusNextPrevChild
func callbackQChart_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QChart) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQChart_FocusOutEvent
func callbackQChart_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QChart) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQChart_GeometryChanged
func callbackQChart_GeometryChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "geometryChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_GetContentsMargins
func callbackQChart_GetContentsMargins(ptr unsafe.Pointer, left C.double, top C.double, right C.double, bottom C.double) {
	if signal := qt.GetSignal(ptr, "getContentsMargins"); signal != nil {
		(*(*func(float64, float64, float64, float64))(signal))(float64(left), float64(top), float64(right), float64(bottom))
	} else {
		NewQChartFromPointer(ptr).GetContentsMarginsDefault(float64(left), float64(top), float64(right), float64(bottom))
	}
}

func (ptr *QChart) GetContentsMarginsDefault(left float64, top float64, right float64, bottom float64) {
	if ptr.Pointer() != nil {
		C.QChart_GetContentsMarginsDefault(ptr.Pointer(), C.double(left), C.double(top), C.double(right), C.double(bottom))
	}
}

//export callbackQChart_GrabKeyboardEvent
func callbackQChart_GrabKeyboardEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "grabKeyboardEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).GrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QChart) GrabKeyboardEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_GrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQChart_GrabMouseEvent
func callbackQChart_GrabMouseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "grabMouseEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).GrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QChart) GrabMouseEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_GrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQChart_HideEvent
func callbackQChart_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QChart) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQChart_HoverLeaveEvent
func callbackQChart_HoverLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverLeaveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneHoverEvent))(signal))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).HoverLeaveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QChart) HoverLeaveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_HoverLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQChart_HoverMoveEvent
func callbackQChart_HoverMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverMoveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneHoverEvent))(signal))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).HoverMoveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QChart) HoverMoveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_HoverMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQChart_InitStyleOption
func callbackQChart_InitStyleOption(ptr unsafe.Pointer, option unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initStyleOption"); signal != nil {
		(*(*func(*widgets.QStyleOption))(signal))(widgets.NewQStyleOptionFromPointer(option))
	} else {
		NewQChartFromPointer(ptr).InitStyleOptionDefault(widgets.NewQStyleOptionFromPointer(option))
	}
}

func (ptr *QChart) InitStyleOptionDefault(option widgets.QStyleOption_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_InitStyleOptionDefault(ptr.Pointer(), widgets.PointerFromQStyleOption(option))
	}
}

//export callbackQChart_ItemChange
func callbackQChart_ItemChange(ptr unsafe.Pointer, change C.longlong, value unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemChange"); signal != nil {
		return core.PointerFromQVariant((*(*func(widgets.QGraphicsItem__GraphicsItemChange, *core.QVariant) *core.QVariant)(signal))(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
	}

	return core.PointerFromQVariant(NewQChartFromPointer(ptr).ItemChangeDefault(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
}

func (ptr *QChart) ItemChangeDefault(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QChart_ItemChangeDefault(ptr.Pointer(), C.longlong(change), core.PointerFromQVariant(value)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQChart_LayoutChanged
func callbackQChart_LayoutChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_MoveEvent
func callbackQChart_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneMoveEvent))(signal))(widgets.NewQGraphicsSceneMoveEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).MoveEventDefault(widgets.NewQGraphicsSceneMoveEventFromPointer(event))
	}
}

func (ptr *QChart) MoveEventDefault(event widgets.QGraphicsSceneMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_MoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMoveEvent(event))
	}
}

//export callbackQChart_Paint
func callbackQChart_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		(*(*func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(signal))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQChartFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QChart) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQChart_PaintWindowFrame
func callbackQChart_PaintWindowFrame(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintWindowFrame"); signal != nil {
		(*(*func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(signal))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQChartFromPointer(ptr).PaintWindowFrameDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QChart) PaintWindowFrameDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_PaintWindowFrameDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQChart_PolishEvent
func callbackQChart_PolishEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "polishEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartFromPointer(ptr).PolishEventDefault()
	}
}

func (ptr *QChart) PolishEventDefault() {
	if ptr.Pointer() != nil {
		C.QChart_PolishEventDefault(ptr.Pointer())
	}
}

//export callbackQChart_ResizeEvent
func callbackQChart_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneResizeEvent))(signal))(widgets.NewQGraphicsSceneResizeEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).ResizeEventDefault(widgets.NewQGraphicsSceneResizeEventFromPointer(event))
	}
}

func (ptr *QChart) ResizeEventDefault(event widgets.QGraphicsSceneResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_ResizeEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneResizeEvent(event))
	}
}

//export callbackQChart_SceneEvent
func callbackQChart_SceneEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "sceneEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).SceneEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QChart) SceneEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_SceneEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQChart_SetGeometry
func callbackQChart_SetGeometry(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		(*(*func(*core.QRectF))(signal))(core.NewQRectFFromPointer(rect))
	} else {
		NewQChartFromPointer(ptr).SetGeometryDefault(core.NewQRectFFromPointer(rect))
	}
}

func (ptr *QChart) SetGeometryDefault(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

//export callbackQChart_Shape
func callbackQChart_Shape(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "shape"); signal != nil {
		return gui.PointerFromQPainterPath((*(*func() *gui.QPainterPath)(signal))())
	}

	return gui.PointerFromQPainterPath(NewQChartFromPointer(ptr).ShapeDefault())
}

func (ptr *QChart) ShapeDefault() *gui.QPainterPath {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPainterPathFromPointer(C.QChart_ShapeDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

//export callbackQChart_ShowEvent
func callbackQChart_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QChart) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQChart_SizeHint
func callbackQChart_SizeHint(ptr unsafe.Pointer, which C.longlong, constraint unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSizeF((*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
	}

	return core.PointerFromQSizeF(NewQChartFromPointer(ptr).SizeHintDefault(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
}

func (ptr *QChart) SizeHintDefault(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QChart_SizeHintDefault(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

//export callbackQChart_Type
func callbackQChart_Type(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQChartFromPointer(ptr).TypeDefault()))
}

func (ptr *QChart) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QChart_TypeDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQChart_UngrabKeyboardEvent
func callbackQChart_UngrabKeyboardEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ungrabKeyboardEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).UngrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QChart) UngrabKeyboardEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_UngrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQChart_UngrabMouseEvent
func callbackQChart_UngrabMouseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ungrabMouseEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).UngrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QChart) UngrabMouseEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_UngrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQChart_UpdateGeometry
func callbackQChart_UpdateGeometry(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateGeometry"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartFromPointer(ptr).UpdateGeometryDefault()
	}
}

func (ptr *QChart) UpdateGeometryDefault() {
	if ptr.Pointer() != nil {
		C.QChart_UpdateGeometryDefault(ptr.Pointer())
	}
}

//export callbackQChart_WindowFrameEvent
func callbackQChart_WindowFrameEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "windowFrameEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).WindowFrameEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QChart) WindowFrameEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_WindowFrameEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQChart_WindowFrameSectionAt
func callbackQChart_WindowFrameSectionAt(ptr unsafe.Pointer, pos unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "windowFrameSectionAt"); signal != nil {
		return C.longlong((*(*func(*core.QPointF) core.Qt__WindowFrameSection)(signal))(core.NewQPointFFromPointer(pos)))
	}

	return C.longlong(NewQChartFromPointer(ptr).WindowFrameSectionAtDefault(core.NewQPointFFromPointer(pos)))
}

func (ptr *QChart) WindowFrameSectionAtDefault(pos core.QPointF_ITF) core.Qt__WindowFrameSection {
	if ptr.Pointer() != nil {
		return core.Qt__WindowFrameSection(C.QChart_WindowFrameSectionAtDefault(ptr.Pointer(), core.PointerFromQPointF(pos)))
	}
	return 0
}

//export callbackQChart_EnabledChanged
func callbackQChart_EnabledChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enabledChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_OpacityChanged
func callbackQChart_OpacityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_ParentChanged
func callbackQChart_ParentChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "parentChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_RotationChanged
func callbackQChart_RotationChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rotationChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_ScaleChanged
func callbackQChart_ScaleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scaleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_UpdateMicroFocus
func callbackQChart_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QChart) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QChart_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQChart_VisibleChanged
func callbackQChart_VisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_XChanged
func callbackQChart_XChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "xChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_YChanged
func callbackQChart_YChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "yChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_ZChanged
func callbackQChart_ZChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "zChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQChart_ChildEvent
func callbackQChart_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QChart) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQChart_ConnectNotify
func callbackQChart_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQChartFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QChart) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQChart_CustomEvent
func callbackQChart_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QChart) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQChart_DeleteLater
func callbackQChart_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QChart) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QChart_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQChart_Destroyed
func callbackQChart_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQChart_DisconnectNotify
func callbackQChart_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQChartFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QChart) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQChart_EventFilter
func callbackQChart_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QChart) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQChart_MetaObject
func callbackQChart_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQChartFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QChart) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QChart_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQChart_ObjectNameChanged
func callbackQChart_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQChart_TimerEvent
func callbackQChart_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QChart) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQChart_Advance
func callbackQChart_Advance(ptr unsafe.Pointer, phase C.int) {
	if signal := qt.GetSignal(ptr, "advance"); signal != nil {
		(*(*func(int))(signal))(int(int32(phase)))
	} else {
		NewQChartFromPointer(ptr).AdvanceDefault(int(int32(phase)))
	}
}

func (ptr *QChart) AdvanceDefault(phase int) {
	if ptr.Pointer() != nil {
		C.QChart_AdvanceDefault(ptr.Pointer(), C.int(int32(phase)))
	}
}

//export callbackQChart_CollidesWithItem
func callbackQChart_CollidesWithItem(ptr unsafe.Pointer, other unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "collidesWithItem"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*widgets.QGraphicsItem, core.Qt__ItemSelectionMode) bool)(signal))(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).CollidesWithItemDefault(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode)))))
}

func (ptr *QChart) CollidesWithItemDefault(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_CollidesWithItemDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(other), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQChart_CollidesWithPath
func callbackQChart_CollidesWithPath(ptr unsafe.Pointer, path unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "collidesWithPath"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*gui.QPainterPath, core.Qt__ItemSelectionMode) bool)(signal))(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).CollidesWithPathDefault(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode)))))
}

func (ptr *QChart) CollidesWithPathDefault(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_CollidesWithPathDefault(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQChart_Contains
func callbackQChart_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QPointF) bool)(signal))(core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point)))))
}

func (ptr *QChart) ContainsDefault(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point))) != 0
	}
	return false
}

//export callbackQChart_ContextMenuEvent
func callbackQChart_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneContextMenuEvent))(signal))(widgets.NewQGraphicsSceneContextMenuEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).ContextMenuEventDefault(widgets.NewQGraphicsSceneContextMenuEventFromPointer(event))
	}
}

func (ptr *QChart) ContextMenuEventDefault(event widgets.QGraphicsSceneContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_ContextMenuEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

//export callbackQChart_DragEnterEvent
func callbackQChart_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneDragDropEvent))(signal))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).DragEnterEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QChart) DragEnterEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_DragEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQChart_DragLeaveEvent
func callbackQChart_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneDragDropEvent))(signal))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).DragLeaveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QChart) DragLeaveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_DragLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQChart_DragMoveEvent
func callbackQChart_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneDragDropEvent))(signal))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).DragMoveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QChart) DragMoveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_DragMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQChart_DropEvent
func callbackQChart_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneDragDropEvent))(signal))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).DropEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QChart) DropEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_DropEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQChart_HoverEnterEvent
func callbackQChart_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverEnterEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneHoverEvent))(signal))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).HoverEnterEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QChart) HoverEnterEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_HoverEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQChart_InputMethodEvent
func callbackQChart_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*gui.QInputMethodEvent))(signal))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QChart) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQChart_InputMethodQuery
func callbackQChart_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant((*(*func(core.Qt__InputMethodQuery) *core.QVariant)(signal))(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQChartFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QChart) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QChart_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQChart_IsObscuredBy
func callbackQChart_IsObscuredBy(ptr unsafe.Pointer, item unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isObscuredBy"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*widgets.QGraphicsItem) bool)(signal))(widgets.NewQGraphicsItemFromPointer(item)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).IsObscuredByDefault(widgets.NewQGraphicsItemFromPointer(item)))))
}

func (ptr *QChart) IsObscuredByDefault(item widgets.QGraphicsItem_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_IsObscuredByDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

//export callbackQChart_KeyPressEvent
func callbackQChart_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QChart) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQChart_KeyReleaseEvent
func callbackQChart_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QChart) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQChart_MouseDoubleClickEvent
func callbackQChart_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneMouseEvent))(signal))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).MouseDoubleClickEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QChart) MouseDoubleClickEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_MouseDoubleClickEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQChart_MouseMoveEvent
func callbackQChart_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneMouseEvent))(signal))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).MouseMoveEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QChart) MouseMoveEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_MouseMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQChart_MousePressEvent
func callbackQChart_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneMouseEvent))(signal))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).MousePressEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QChart) MousePressEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_MousePressEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQChart_MouseReleaseEvent
func callbackQChart_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneMouseEvent))(signal))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).MouseReleaseEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QChart) MouseReleaseEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_MouseReleaseEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQChart_OpaqueArea
func callbackQChart_OpaqueArea(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "opaqueArea"); signal != nil {
		return gui.PointerFromQPainterPath((*(*func() *gui.QPainterPath)(signal))())
	}

	return gui.PointerFromQPainterPath(NewQChartFromPointer(ptr).OpaqueAreaDefault())
}

func (ptr *QChart) OpaqueAreaDefault() *gui.QPainterPath {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPainterPathFromPointer(C.QChart_OpaqueAreaDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

//export callbackQChart_SceneEventFilter
func callbackQChart_SceneEventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "sceneEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*widgets.QGraphicsItem, *core.QEvent) bool)(signal))(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartFromPointer(ptr).SceneEventFilterDefault(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QChart) SceneEventFilterDefault(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChart_SceneEventFilterDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQChart_WheelEvent
func callbackQChart_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneWheelEvent))(signal))(widgets.NewQGraphicsSceneWheelEventFromPointer(event))
	} else {
		NewQChartFromPointer(ptr).WheelEventDefault(widgets.NewQGraphicsSceneWheelEventFromPointer(event))
	}
}

func (ptr *QChart) WheelEventDefault(event widgets.QGraphicsSceneWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_WheelEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneWheelEvent(event))
	}
}

type QChartView struct {
	widgets.QGraphicsView
}

type QChartView_ITF interface {
	widgets.QGraphicsView_ITF
	QChartView_PTR() *QChartView
}

func (ptr *QChartView) QChartView_PTR() *QChartView {
	return ptr
}

func (ptr *QChartView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsView_PTR().Pointer()
	}
	return nil
}

func (ptr *QChartView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsView_PTR().SetPointer(p)
	}
}

func PointerFromQChartView(ptr QChartView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QChartView_PTR().Pointer()
	}
	return nil
}

func NewQChartViewFromPointer(ptr unsafe.Pointer) (n *QChartView) {
	n = new(QChartView)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QChartView__RubberBand
//QChartView::RubberBand
type QChartView__RubberBand int64

const (
	QChartView__NoRubberBand         QChartView__RubberBand = QChartView__RubberBand(0x0)
	QChartView__VerticalRubberBand   QChartView__RubberBand = QChartView__RubberBand(0x1)
	QChartView__HorizontalRubberBand QChartView__RubberBand = QChartView__RubberBand(0x2)
	QChartView__RectangleRubberBand  QChartView__RubberBand = QChartView__RubberBand(0x3)
)

func NewQChartView(parent widgets.QWidget_ITF) *QChartView {
	tmpValue := NewQChartViewFromPointer(C.QChartView_NewQChartView(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQChartView2(chart QChart_ITF, parent widgets.QWidget_ITF) *QChartView {
	tmpValue := NewQChartViewFromPointer(C.QChartView_NewQChartView2(PointerFromQChart(chart), widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QChartView) Chart() *QChart {
	if ptr.Pointer() != nil {
		tmpValue := NewQChartFromPointer(C.QChartView_Chart(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQChartView_MouseMoveEvent
func callbackQChartView_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QChartView) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQChartView_MousePressEvent
func callbackQChartView_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QChartView) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQChartView_MouseReleaseEvent
func callbackQChartView_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QChartView) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQChartView_ResizeEvent
func callbackQChartView_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*gui.QResizeEvent))(signal))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QChartView) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QChartView) RubberBand() QChartView__RubberBand {
	if ptr.Pointer() != nil {
		return QChartView__RubberBand(C.QChartView_RubberBand(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChartView) SetChart(chart QChart_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_SetChart(ptr.Pointer(), PointerFromQChart(chart))
	}
}

func (ptr *QChartView) SetRubberBand(rubberBand QChartView__RubberBand) {
	if ptr.Pointer() != nil {
		C.QChartView_SetRubberBand(ptr.Pointer(), C.longlong(rubberBand))
	}
}

//export callbackQChartView_DestroyQChartView
func callbackQChartView_DestroyQChartView(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QChartView"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).DestroyQChartViewDefault()
	}
}

func (ptr *QChartView) ConnectDestroyQChartView(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QChartView"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QChartView", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QChartView", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QChartView) DisconnectDestroyQChartView() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QChartView")
	}
}

func (ptr *QChartView) DestroyQChartView() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QChartView_DestroyQChartView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QChartView) DestroyQChartViewDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QChartView_DestroyQChartViewDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QChartView) __items_atList(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QChartView___items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChartView) __items_setList(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___items_setList(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QChartView) __items_newList() unsafe.Pointer {
	return C.QChartView___items_newList(ptr.Pointer())
}

func (ptr *QChartView) __items_atList2(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QChartView___items_atList2(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChartView) __items_setList2(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___items_setList2(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QChartView) __items_newList2() unsafe.Pointer {
	return C.QChartView___items_newList2(ptr.Pointer())
}

func (ptr *QChartView) __items_atList3(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QChartView___items_atList3(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChartView) __items_setList3(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___items_setList3(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QChartView) __items_newList3() unsafe.Pointer {
	return C.QChartView___items_newList3(ptr.Pointer())
}

func (ptr *QChartView) __items_atList4(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QChartView___items_atList4(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChartView) __items_setList4(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___items_setList4(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QChartView) __items_newList4() unsafe.Pointer {
	return C.QChartView___items_newList4(ptr.Pointer())
}

func (ptr *QChartView) __items_atList5(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QChartView___items_atList5(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChartView) __items_setList5(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___items_setList5(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QChartView) __items_newList5() unsafe.Pointer {
	return C.QChartView___items_newList5(ptr.Pointer())
}

func (ptr *QChartView) __items_atList6(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QChartView___items_atList6(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChartView) __items_setList6(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___items_setList6(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QChartView) __items_newList6() unsafe.Pointer {
	return C.QChartView___items_newList6(ptr.Pointer())
}

func (ptr *QChartView) __items_atList7(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QChartView___items_atList7(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChartView) __items_setList7(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___items_setList7(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QChartView) __items_newList7() unsafe.Pointer {
	return C.QChartView___items_newList7(ptr.Pointer())
}

func (ptr *QChartView) __updateScene_rects_atList(i int) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QChartView___updateScene_rects_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QChartView) __updateScene_rects_setList(i core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___updateScene_rects_setList(ptr.Pointer(), core.PointerFromQRectF(i))
	}
}

func (ptr *QChartView) __updateScene_rects_newList() unsafe.Pointer {
	return C.QChartView___updateScene_rects_newList(ptr.Pointer())
}

func (ptr *QChartView) __scrollBarWidgets_atList(i int) *widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQWidgetFromPointer(C.QChartView___scrollBarWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChartView) __scrollBarWidgets_setList(i widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___scrollBarWidgets_setList(ptr.Pointer(), widgets.PointerFromQWidget(i))
	}
}

func (ptr *QChartView) __scrollBarWidgets_newList() unsafe.Pointer {
	return C.QChartView___scrollBarWidgets_newList(ptr.Pointer())
}

func (ptr *QChartView) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QChartView___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChartView) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QChartView) __actions_newList() unsafe.Pointer {
	return C.QChartView___actions_newList(ptr.Pointer())
}

func (ptr *QChartView) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QChartView___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChartView) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QChartView) __addActions_actions_newList() unsafe.Pointer {
	return C.QChartView___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QChartView) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QChartView___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChartView) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QChartView) __insertActions_actions_newList() unsafe.Pointer {
	return C.QChartView___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QChartView) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QChartView___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChartView) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QChartView) __children_newList() unsafe.Pointer {
	return C.QChartView___children_newList(ptr.Pointer())
}

func (ptr *QChartView) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QChartView___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QChartView) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QChartView) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QChartView___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QChartView) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QChartView___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChartView) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QChartView) __findChildren_newList() unsafe.Pointer {
	return C.QChartView___findChildren_newList(ptr.Pointer())
}

func (ptr *QChartView) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QChartView___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QChartView) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QChartView) __findChildren_newList3() unsafe.Pointer {
	return C.QChartView___findChildren_newList3(ptr.Pointer())
}

//export callbackQChartView_ContextMenuEvent
func callbackQChartView_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*gui.QContextMenuEvent))(signal))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QChartView) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQChartView_DragEnterEvent
func callbackQChartView_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*gui.QDragEnterEvent))(signal))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QChartView) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQChartView_DragLeaveEvent
func callbackQChartView_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*gui.QDragLeaveEvent))(signal))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QChartView) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQChartView_DragMoveEvent
func callbackQChartView_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*gui.QDragMoveEvent))(signal))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QChartView) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQChartView_DrawBackground
func callbackQChartView_DrawBackground(ptr unsafe.Pointer, painter unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawBackground"); signal != nil {
		(*(*func(*gui.QPainter, *core.QRectF))(signal))(gui.NewQPainterFromPointer(painter), core.NewQRectFFromPointer(rect))
	} else {
		NewQChartViewFromPointer(ptr).DrawBackgroundDefault(gui.NewQPainterFromPointer(painter), core.NewQRectFFromPointer(rect))
	}
}

func (ptr *QChartView) DrawBackgroundDefault(painter gui.QPainter_ITF, rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_DrawBackgroundDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRectF(rect))
	}
}

//export callbackQChartView_DrawForeground
func callbackQChartView_DrawForeground(ptr unsafe.Pointer, painter unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawForeground"); signal != nil {
		(*(*func(*gui.QPainter, *core.QRectF))(signal))(gui.NewQPainterFromPointer(painter), core.NewQRectFFromPointer(rect))
	} else {
		NewQChartViewFromPointer(ptr).DrawForegroundDefault(gui.NewQPainterFromPointer(painter), core.NewQRectFFromPointer(rect))
	}
}

func (ptr *QChartView) DrawForegroundDefault(painter gui.QPainter_ITF, rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_DrawForegroundDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRectF(rect))
	}
}

//export callbackQChartView_DropEvent
func callbackQChartView_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*gui.QDropEvent))(signal))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QChartView) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQChartView_Event
func callbackQChartView_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartViewFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QChartView) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChartView_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQChartView_FocusInEvent
func callbackQChartView_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QChartView) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQChartView_FocusNextPrevChild
func callbackQChartView_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartViewFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QChartView) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChartView_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQChartView_FocusOutEvent
func callbackQChartView_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QChartView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQChartView_InputMethodEvent
func callbackQChartView_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*gui.QInputMethodEvent))(signal))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QChartView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQChartView_InputMethodQuery
func callbackQChartView_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant((*(*func(core.Qt__InputMethodQuery) *core.QVariant)(signal))(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQChartViewFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QChartView) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QChartView_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQChartView_InvalidateScene
func callbackQChartView_InvalidateScene(ptr unsafe.Pointer, rect unsafe.Pointer, layers C.longlong) {
	if signal := qt.GetSignal(ptr, "invalidateScene"); signal != nil {
		(*(*func(*core.QRectF, widgets.QGraphicsScene__SceneLayer))(signal))(core.NewQRectFFromPointer(rect), widgets.QGraphicsScene__SceneLayer(layers))
	} else {
		NewQChartViewFromPointer(ptr).InvalidateSceneDefault(core.NewQRectFFromPointer(rect), widgets.QGraphicsScene__SceneLayer(layers))
	}
}

func (ptr *QChartView) InvalidateSceneDefault(rect core.QRectF_ITF, layers widgets.QGraphicsScene__SceneLayer) {
	if ptr.Pointer() != nil {
		C.QChartView_InvalidateSceneDefault(ptr.Pointer(), core.PointerFromQRectF(rect), C.longlong(layers))
	}
}

//export callbackQChartView_KeyPressEvent
func callbackQChartView_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QChartView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQChartView_KeyReleaseEvent
func callbackQChartView_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QChartView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQChartView_MouseDoubleClickEvent
func callbackQChartView_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QChartView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQChartView_PaintEvent
func callbackQChartView_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*gui.QPaintEvent))(signal))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QChartView) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQChartView_RubberBandChanged
func callbackQChartView_RubberBandChanged(ptr unsafe.Pointer, rubberBandRect unsafe.Pointer, fromScenePoint unsafe.Pointer, toScenePoint unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rubberBandChanged"); signal != nil {
		(*(*func(*core.QRect, *core.QPointF, *core.QPointF))(signal))(func() *core.QRect {
			tmpValue := core.NewQRectFromPointer(rubberBandRect)
			qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
			return tmpValue
		}(), func() *core.QPointF {
			tmpValue := core.NewQPointFFromPointer(fromScenePoint)
			qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
			return tmpValue
		}(), func() *core.QPointF {
			tmpValue := core.NewQPointFFromPointer(toScenePoint)
			qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
			return tmpValue
		}())
	}

}

//export callbackQChartView_ScrollContentsBy
func callbackQChartView_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {
	if signal := qt.GetSignal(ptr, "scrollContentsBy"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(dx)), int(int32(dy)))
	} else {
		NewQChartViewFromPointer(ptr).ScrollContentsByDefault(int(int32(dx)), int(int32(dy)))
	}
}

func (ptr *QChartView) ScrollContentsByDefault(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QChartView_ScrollContentsByDefault(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

//export callbackQChartView_SetupViewport
func callbackQChartView_SetupViewport(ptr unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setupViewport"); signal != nil {
		(*(*func(*widgets.QWidget))(signal))(widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQChartViewFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QChartView) SetupViewportDefault(widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQChartView_ShowEvent
func callbackQChartView_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QChartView) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQChartView_SizeHint
func callbackQChartView_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQChartViewFromPointer(ptr).SizeHintDefault())
}

func (ptr *QChartView) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QChartView_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQChartView_UpdateScene
func callbackQChartView_UpdateScene(ptr unsafe.Pointer, rects C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(ptr, "updateScene"); signal != nil {
		(*(*func([]*core.QRectF))(signal))(func(l C.struct_QtCharts_PackedList) []*core.QRectF {
			out := make([]*core.QRectF, int(l.len))
			tmpList := NewQChartViewFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__updateScene_rects_atList(i)
			}
			return out
		}(rects))
	} else {
		NewQChartViewFromPointer(ptr).UpdateSceneDefault(func(l C.struct_QtCharts_PackedList) []*core.QRectF {
			out := make([]*core.QRectF, int(l.len))
			tmpList := NewQChartViewFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__updateScene_rects_atList(i)
			}
			return out
		}(rects))
	}
}

func (ptr *QChartView) UpdateSceneDefault(rects []*core.QRectF) {
	if ptr.Pointer() != nil {
		C.QChartView_UpdateSceneDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQChartViewFromPointer(NewQChartViewFromPointer(nil).__updateScene_rects_newList())
			for _, v := range rects {
				tmpList.__updateScene_rects_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQChartView_UpdateSceneRect
func callbackQChartView_UpdateSceneRect(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateSceneRect"); signal != nil {
		(*(*func(*core.QRectF))(signal))(core.NewQRectFFromPointer(rect))
	} else {
		NewQChartViewFromPointer(ptr).UpdateSceneRectDefault(core.NewQRectFFromPointer(rect))
	}
}

func (ptr *QChartView) UpdateSceneRectDefault(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_UpdateSceneRectDefault(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

//export callbackQChartView_ViewportEvent
func callbackQChartView_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartViewFromPointer(ptr).ViewportEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QChartView) ViewportEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChartView_ViewportEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQChartView_WheelEvent
func callbackQChartView_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*gui.QWheelEvent))(signal))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QChartView) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQChartView_MinimumSizeHint
func callbackQChartView_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQChartViewFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QChartView) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QChartView_MinimumSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQChartView_ViewportSizeHint
func callbackQChartView_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewportSizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQChartViewFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *QChartView) ViewportSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QChartView_ViewportSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQChartView_ChangeEvent
func callbackQChartView_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(ev))
	} else {
		NewQChartViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QChartView) ChangeEventDefault(ev core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackQChartView_ActionEvent
func callbackQChartView_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*gui.QActionEvent))(signal))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QChartView) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQChartView_Close
func callbackQChartView_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartViewFromPointer(ptr).CloseDefault())))
}

func (ptr *QChartView) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QChartView_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQChartView_CloseEvent
func callbackQChartView_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*gui.QCloseEvent))(signal))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QChartView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQChartView_CustomContextMenuRequested
func callbackQChartView_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*core.QPoint))(signal))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQChartView_EnterEvent
func callbackQChartView_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QChartView) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQChartView_HasHeightForWidth
func callbackQChartView_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartViewFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QChartView) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QChartView_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQChartView_HeightForWidth
func callbackQChartView_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewQChartViewFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QChartView) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QChartView_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQChartView_Hide
func callbackQChartView_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).HideDefault()
	}
}

func (ptr *QChartView) HideDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_HideDefault(ptr.Pointer())
	}
}

//export callbackQChartView_HideEvent
func callbackQChartView_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QChartView) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQChartView_InitPainter
func callbackQChartView_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*gui.QPainter))(signal))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQChartViewFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QChartView) InitPainterDefault(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQChartView_LeaveEvent
func callbackQChartView_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QChartView) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQChartView_Lower
func callbackQChartView_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QChartView) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_LowerDefault(ptr.Pointer())
	}
}

//export callbackQChartView_Metric
func callbackQChartView_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(gui.QPaintDevice__PaintDeviceMetric) int)(signal))(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQChartViewFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QChartView) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QChartView_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQChartView_MoveEvent
func callbackQChartView_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*gui.QMoveEvent))(signal))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QChartView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQChartView_NativeEvent
func callbackQChartView_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QByteArray, unsafe.Pointer, *int) bool)(signal))(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartViewFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *QChartView) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.QChartView_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackQChartView_PaintEngine
func callbackQChartView_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine((*(*func() *gui.QPaintEngine)(signal))())
	}

	return gui.PointerFromQPaintEngine(NewQChartViewFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QChartView) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QChartView_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQChartView_Raise
func callbackQChartView_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QChartView) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQChartView_Repaint
func callbackQChartView_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QChartView) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQChartView_SetDisabled
func callbackQChartView_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewQChartViewFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QChartView) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QChartView_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQChartView_SetEnabled
func callbackQChartView_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewQChartViewFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QChartView) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QChartView_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQChartView_SetFocus2
func callbackQChartView_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QChartView) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QChartView_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQChartView_SetHidden
func callbackQChartView_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewQChartViewFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QChartView) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QChartView_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQChartView_SetStyleSheet
func callbackQChartView_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewQChartViewFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QChartView) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QChartView_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtCharts_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQChartView_SetVisible
func callbackQChartView_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewQChartViewFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QChartView) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QChartView_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQChartView_SetWindowModified
func callbackQChartView_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewQChartViewFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QChartView) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QChartView_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQChartView_SetWindowTitle
func callbackQChartView_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewQChartViewFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QChartView) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QChartView_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtCharts_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQChartView_Show
func callbackQChartView_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QChartView) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_ShowDefault(ptr.Pointer())
	}
}

//export callbackQChartView_ShowFullScreen
func callbackQChartView_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QChartView) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQChartView_ShowMaximized
func callbackQChartView_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QChartView) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQChartView_ShowMinimized
func callbackQChartView_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QChartView) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQChartView_ShowNormal
func callbackQChartView_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QChartView) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQChartView_TabletEvent
func callbackQChartView_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*gui.QTabletEvent))(signal))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QChartView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQChartView_Update
func callbackQChartView_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QChartView) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQChartView_UpdateMicroFocus
func callbackQChartView_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QChartView) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QChartView_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQChartView_WindowIconChanged
func callbackQChartView_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*gui.QIcon))(signal))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQChartView_WindowTitleChanged
func callbackQChartView_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackQChartView_ChildEvent
func callbackQChartView_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QChartView) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQChartView_ConnectNotify
func callbackQChartView_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQChartViewFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QChartView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQChartView_CustomEvent
func callbackQChartView_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QChartView) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQChartView_DeleteLater
func callbackQChartView_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQChartViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QChartView) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QChartView_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQChartView_Destroyed
func callbackQChartView_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQChartView_DisconnectNotify
func callbackQChartView_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQChartViewFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QChartView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQChartView_EventFilter
func callbackQChartView_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQChartViewFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QChartView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QChartView_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQChartView_MetaObject
func callbackQChartView_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQChartViewFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QChartView) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QChartView_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQChartView_ObjectNameChanged
func callbackQChartView_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQChartView_TimerEvent
func callbackQChartView_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQChartViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QChartView) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QDateTimeAxis struct {
	QAbstractAxis
}

type QDateTimeAxis_ITF interface {
	QAbstractAxis_ITF
	QDateTimeAxis_PTR() *QDateTimeAxis
}

func (ptr *QDateTimeAxis) QDateTimeAxis_PTR() *QDateTimeAxis {
	return ptr
}

func (ptr *QDateTimeAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractAxis_PTR().Pointer()
	}
	return nil
}

func (ptr *QDateTimeAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractAxis_PTR().SetPointer(p)
	}
}

func PointerFromQDateTimeAxis(ptr QDateTimeAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateTimeAxis_PTR().Pointer()
	}
	return nil
}

func NewQDateTimeAxisFromPointer(ptr unsafe.Pointer) (n *QDateTimeAxis) {
	n = new(QDateTimeAxis)
	n.SetPointer(ptr)
	return
}
func NewQDateTimeAxis(parent core.QObject_ITF) *QDateTimeAxis {
	tmpValue := NewQDateTimeAxisFromPointer(C.QDateTimeAxis_NewQDateTimeAxis(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDateTimeAxis) Format() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDateTimeAxis_Format(ptr.Pointer()))
	}
	return ""
}

//export callbackQDateTimeAxis_FormatChanged
func callbackQDateTimeAxis_FormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "formatChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(format))
	}

}

func (ptr *QDateTimeAxis) ConnectFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "formatChanged") {
			C.QDateTimeAxis_ConnectFormatChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "formatChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "formatChanged"); signal != nil {
			f := func(format string) {
				(*(*func(string))(signal))(format)
				f(format)
			}
			qt.ConnectSignal(ptr.Pointer(), "formatChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "formatChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDateTimeAxis) DisconnectFormatChanged() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DisconnectFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "formatChanged")
	}
}

func (ptr *QDateTimeAxis) FormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QDateTimeAxis_FormatChanged(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

func (ptr *QDateTimeAxis) Max() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QDateTimeAxis_Max(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

//export callbackQDateTimeAxis_MaxChanged
func callbackQDateTimeAxis_MaxChanged(ptr unsafe.Pointer, max unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "maxChanged"); signal != nil {
		(*(*func(*core.QDateTime))(signal))(func() *core.QDateTime {
			tmpValue := core.NewQDateTimeFromPointer(max)
			qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
			return tmpValue
		}())
	}

}

func (ptr *QDateTimeAxis) ConnectMaxChanged(f func(max *core.QDateTime)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "maxChanged") {
			C.QDateTimeAxis_ConnectMaxChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "maxChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "maxChanged"); signal != nil {
			f := func(max *core.QDateTime) {
				(*(*func(*core.QDateTime))(signal))(max)
				f(max)
			}
			qt.ConnectSignal(ptr.Pointer(), "maxChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maxChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDateTimeAxis) DisconnectMaxChanged() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DisconnectMaxChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "maxChanged")
	}
}

func (ptr *QDateTimeAxis) MaxChanged(max core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_MaxChanged(ptr.Pointer(), core.PointerFromQDateTime(max))
	}
}

func (ptr *QDateTimeAxis) Min() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QDateTimeAxis_Min(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

//export callbackQDateTimeAxis_MinChanged
func callbackQDateTimeAxis_MinChanged(ptr unsafe.Pointer, min unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "minChanged"); signal != nil {
		(*(*func(*core.QDateTime))(signal))(func() *core.QDateTime {
			tmpValue := core.NewQDateTimeFromPointer(min)
			qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
			return tmpValue
		}())
	}

}

func (ptr *QDateTimeAxis) ConnectMinChanged(f func(min *core.QDateTime)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minChanged") {
			C.QDateTimeAxis_ConnectMinChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minChanged"); signal != nil {
			f := func(min *core.QDateTime) {
				(*(*func(*core.QDateTime))(signal))(min)
				f(min)
			}
			qt.ConnectSignal(ptr.Pointer(), "minChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDateTimeAxis) DisconnectMinChanged() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DisconnectMinChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minChanged")
	}
}

func (ptr *QDateTimeAxis) MinChanged(min core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_MinChanged(ptr.Pointer(), core.PointerFromQDateTime(min))
	}
}

//export callbackQDateTimeAxis_RangeChanged
func callbackQDateTimeAxis_RangeChanged(ptr unsafe.Pointer, min unsafe.Pointer, max unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rangeChanged"); signal != nil {
		(*(*func(*core.QDateTime, *core.QDateTime))(signal))(func() *core.QDateTime {
			tmpValue := core.NewQDateTimeFromPointer(min)
			qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
			return tmpValue
		}(), func() *core.QDateTime {
			tmpValue := core.NewQDateTimeFromPointer(max)
			qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
			return tmpValue
		}())
	}

}

func (ptr *QDateTimeAxis) ConnectRangeChanged(f func(min *core.QDateTime, max *core.QDateTime)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rangeChanged") {
			C.QDateTimeAxis_ConnectRangeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "rangeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rangeChanged"); signal != nil {
			f := func(min *core.QDateTime, max *core.QDateTime) {
				(*(*func(*core.QDateTime, *core.QDateTime))(signal))(min, max)
				f(min, max)
			}
			qt.ConnectSignal(ptr.Pointer(), "rangeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rangeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDateTimeAxis) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rangeChanged")
	}
}

func (ptr *QDateTimeAxis) RangeChanged(min core.QDateTime_ITF, max core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_RangeChanged(ptr.Pointer(), core.PointerFromQDateTime(min), core.PointerFromQDateTime(max))
	}
}

func (ptr *QDateTimeAxis) SetFormat(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QDateTimeAxis_SetFormat(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

func (ptr *QDateTimeAxis) SetMax(max core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_SetMax(ptr.Pointer(), core.PointerFromQDateTime(max))
	}
}

func (ptr *QDateTimeAxis) SetMin(min core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_SetMin(ptr.Pointer(), core.PointerFromQDateTime(min))
	}
}

func (ptr *QDateTimeAxis) SetRange(min core.QDateTime_ITF, max core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_SetRange(ptr.Pointer(), core.PointerFromQDateTime(min), core.PointerFromQDateTime(max))
	}
}

func (ptr *QDateTimeAxis) SetTickCount(count int) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_SetTickCount(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QDateTimeAxis) TickCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDateTimeAxis_TickCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQDateTimeAxis_TickCountChanged
func callbackQDateTimeAxis_TickCountChanged(ptr unsafe.Pointer, tickCount C.int) {
	if signal := qt.GetSignal(ptr, "tickCountChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(tickCount)))
	}

}

func (ptr *QDateTimeAxis) ConnectTickCountChanged(f func(tickCount int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "tickCountChanged") {
			C.QDateTimeAxis_ConnectTickCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "tickCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "tickCountChanged"); signal != nil {
			f := func(tickCount int) {
				(*(*func(int))(signal))(tickCount)
				f(tickCount)
			}
			qt.ConnectSignal(ptr.Pointer(), "tickCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "tickCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDateTimeAxis) DisconnectTickCountChanged() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DisconnectTickCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "tickCountChanged")
	}
}

func (ptr *QDateTimeAxis) TickCountChanged(tickCount int) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_TickCountChanged(ptr.Pointer(), C.int(int32(tickCount)))
	}
}

//export callbackQDateTimeAxis_Type
func callbackQDateTimeAxis_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractAxis__AxisType)(signal))())
	}

	return C.longlong(NewQDateTimeAxisFromPointer(ptr).TypeDefault())
}

func (ptr *QDateTimeAxis) ConnectType(f func() QAbstractAxis__AxisType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractAxis__AxisType {
				(*(*func() QAbstractAxis__AxisType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDateTimeAxis) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QDateTimeAxis) Type() QAbstractAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstractAxis__AxisType(C.QDateTimeAxis_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTimeAxis) TypeDefault() QAbstractAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstractAxis__AxisType(C.QDateTimeAxis_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQDateTimeAxis_DestroyQDateTimeAxis
func callbackQDateTimeAxis_DestroyQDateTimeAxis(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDateTimeAxis"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDateTimeAxisFromPointer(ptr).DestroyQDateTimeAxisDefault()
	}
}

func (ptr *QDateTimeAxis) ConnectDestroyQDateTimeAxis(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDateTimeAxis"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QDateTimeAxis", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDateTimeAxis", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDateTimeAxis) DisconnectDestroyQDateTimeAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDateTimeAxis")
	}
}

func (ptr *QDateTimeAxis) DestroyQDateTimeAxis() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDateTimeAxis_DestroyQDateTimeAxis(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDateTimeAxis) DestroyQDateTimeAxisDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDateTimeAxis_DestroyQDateTimeAxisDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QHBarModelMapper struct {
	core.QObject
}

type QHBarModelMapper_ITF interface {
	core.QObject_ITF
	QHBarModelMapper_PTR() *QHBarModelMapper
}

func (ptr *QHBarModelMapper) QHBarModelMapper_PTR() *QHBarModelMapper {
	return ptr
}

func (ptr *QHBarModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QHBarModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQHBarModelMapper(ptr QHBarModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBarModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQHBarModelMapperFromPointer(ptr unsafe.Pointer) (n *QHBarModelMapper) {
	n = new(QHBarModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQHBarModelMapper(parent core.QObject_ITF) *QHBarModelMapper {
	tmpValue := NewQHBarModelMapperFromPointer(C.QHBarModelMapper_NewQHBarModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHBarModelMapper) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBarModelMapper_ColumnCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQHBarModelMapper_ColumnCountChanged
func callbackQHBarModelMapper_ColumnCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "columnCountChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBarModelMapper) ConnectColumnCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "columnCountChanged") {
			C.QHBarModelMapper_ConnectColumnCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "columnCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "columnCountChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "columnCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "columnCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBarModelMapper) DisconnectColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectColumnCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "columnCountChanged")
	}
}

func (ptr *QHBarModelMapper) ColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ColumnCountChanged(ptr.Pointer())
	}
}

func (ptr *QHBarModelMapper) FirstBarSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBarModelMapper_FirstBarSetRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQHBarModelMapper_FirstBarSetRowChanged
func callbackQHBarModelMapper_FirstBarSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstBarSetRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBarModelMapper) ConnectFirstBarSetRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstBarSetRowChanged") {
			C.QHBarModelMapper_ConnectFirstBarSetRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstBarSetRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstBarSetRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstBarSetRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstBarSetRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBarModelMapper) DisconnectFirstBarSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectFirstBarSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstBarSetRowChanged")
	}
}

func (ptr *QHBarModelMapper) FirstBarSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_FirstBarSetRowChanged(ptr.Pointer())
	}
}

func (ptr *QHBarModelMapper) FirstColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBarModelMapper_FirstColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQHBarModelMapper_FirstColumnChanged
func callbackQHBarModelMapper_FirstColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBarModelMapper) ConnectFirstColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstColumnChanged") {
			C.QHBarModelMapper_ConnectFirstColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBarModelMapper) DisconnectFirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectFirstColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstColumnChanged")
	}
}

func (ptr *QHBarModelMapper) FirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_FirstColumnChanged(ptr.Pointer())
	}
}

func (ptr *QHBarModelMapper) LastBarSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBarModelMapper_LastBarSetRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQHBarModelMapper_LastBarSetRowChanged
func callbackQHBarModelMapper_LastBarSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lastBarSetRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBarModelMapper) ConnectLastBarSetRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lastBarSetRowChanged") {
			C.QHBarModelMapper_ConnectLastBarSetRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lastBarSetRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lastBarSetRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lastBarSetRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastBarSetRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBarModelMapper) DisconnectLastBarSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectLastBarSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lastBarSetRowChanged")
	}
}

func (ptr *QHBarModelMapper) LastBarSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_LastBarSetRowChanged(ptr.Pointer())
	}
}

func (ptr *QHBarModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QHBarModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHBarModelMapper_ModelReplaced
func callbackQHBarModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBarModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "modelReplaced") {
			C.QHBarModelMapper_ConnectModelReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "modelReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "modelReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBarModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "modelReplaced")
	}
}

func (ptr *QHBarModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ModelReplaced(ptr.Pointer())
	}
}

func (ptr *QHBarModelMapper) Series() *QAbstractBarSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractBarSeriesFromPointer(C.QHBarModelMapper_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHBarModelMapper_SeriesReplaced
func callbackQHBarModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "seriesReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBarModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "seriesReplaced") {
			C.QHBarModelMapper_ConnectSeriesReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "seriesReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "seriesReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBarModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "seriesReplaced")
	}
}

func (ptr *QHBarModelMapper) SeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_SeriesReplaced(ptr.Pointer())
	}
}

func (ptr *QHBarModelMapper) SetColumnCount(columnCount int) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_SetColumnCount(ptr.Pointer(), C.int(int32(columnCount)))
	}
}

func (ptr *QHBarModelMapper) SetFirstBarSetRow(firstBarSetRow int) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_SetFirstBarSetRow(ptr.Pointer(), C.int(int32(firstBarSetRow)))
	}
}

func (ptr *QHBarModelMapper) SetFirstColumn(firstColumn int) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_SetFirstColumn(ptr.Pointer(), C.int(int32(firstColumn)))
	}
}

func (ptr *QHBarModelMapper) SetLastBarSetRow(lastBarSetRow int) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_SetLastBarSetRow(ptr.Pointer(), C.int(int32(lastBarSetRow)))
	}
}

func (ptr *QHBarModelMapper) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHBarModelMapper) SetSeries(series QAbstractBarSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_SetSeries(ptr.Pointer(), PointerFromQAbstractBarSeries(series))
	}
}

func (ptr *QHBarModelMapper) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHBarModelMapper___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHBarModelMapper) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHBarModelMapper) __children_newList() unsafe.Pointer {
	return C.QHBarModelMapper___children_newList(ptr.Pointer())
}

func (ptr *QHBarModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHBarModelMapper___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHBarModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHBarModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHBarModelMapper___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHBarModelMapper) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHBarModelMapper___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHBarModelMapper) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHBarModelMapper) __findChildren_newList() unsafe.Pointer {
	return C.QHBarModelMapper___findChildren_newList(ptr.Pointer())
}

func (ptr *QHBarModelMapper) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHBarModelMapper___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHBarModelMapper) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHBarModelMapper) __findChildren_newList3() unsafe.Pointer {
	return C.QHBarModelMapper___findChildren_newList3(ptr.Pointer())
}

//export callbackQHBarModelMapper_ChildEvent
func callbackQHBarModelMapper_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHBarModelMapperFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHBarModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHBarModelMapper_ConnectNotify
func callbackQHBarModelMapper_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHBarModelMapperFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHBarModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHBarModelMapper_CustomEvent
func callbackQHBarModelMapper_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQHBarModelMapperFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHBarModelMapper) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHBarModelMapper_DeleteLater
func callbackQHBarModelMapper_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHBarModelMapperFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHBarModelMapper) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHBarModelMapper_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHBarModelMapper_Destroyed
func callbackQHBarModelMapper_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHBarModelMapper_DisconnectNotify
func callbackQHBarModelMapper_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHBarModelMapperFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHBarModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHBarModelMapper_Event
func callbackQHBarModelMapper_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHBarModelMapperFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHBarModelMapper) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHBarModelMapper_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQHBarModelMapper_EventFilter
func callbackQHBarModelMapper_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHBarModelMapperFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHBarModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHBarModelMapper_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHBarModelMapper_MetaObject
func callbackQHBarModelMapper_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHBarModelMapperFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHBarModelMapper) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHBarModelMapper_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHBarModelMapper_ObjectNameChanged
func callbackQHBarModelMapper_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQHBarModelMapper_TimerEvent
func callbackQHBarModelMapper_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHBarModelMapperFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHBarModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QHBoxPlotModelMapper struct {
	core.QObject
}

type QHBoxPlotModelMapper_ITF interface {
	core.QObject_ITF
	QHBoxPlotModelMapper_PTR() *QHBoxPlotModelMapper
}

func (ptr *QHBoxPlotModelMapper) QHBoxPlotModelMapper_PTR() *QHBoxPlotModelMapper {
	return ptr
}

func (ptr *QHBoxPlotModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QHBoxPlotModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQHBoxPlotModelMapper(ptr QHBoxPlotModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBoxPlotModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQHBoxPlotModelMapperFromPointer(ptr unsafe.Pointer) (n *QHBoxPlotModelMapper) {
	n = new(QHBoxPlotModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQHBoxPlotModelMapper(parent core.QObject_ITF) *QHBoxPlotModelMapper {
	tmpValue := NewQHBoxPlotModelMapperFromPointer(C.QHBoxPlotModelMapper_NewQHBoxPlotModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHBoxPlotModelMapper) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBoxPlotModelMapper_ColumnCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQHBoxPlotModelMapper_ColumnCountChanged
func callbackQHBoxPlotModelMapper_ColumnCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "columnCountChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectColumnCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "columnCountChanged") {
			C.QHBoxPlotModelMapper_ConnectColumnCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "columnCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "columnCountChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "columnCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "columnCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectColumnCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "columnCountChanged")
	}
}

func (ptr *QHBoxPlotModelMapper) ColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ColumnCountChanged(ptr.Pointer())
	}
}

func (ptr *QHBoxPlotModelMapper) FirstBoxSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBoxPlotModelMapper_FirstBoxSetRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQHBoxPlotModelMapper_FirstBoxSetRowChanged
func callbackQHBoxPlotModelMapper_FirstBoxSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstBoxSetRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectFirstBoxSetRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstBoxSetRowChanged") {
			C.QHBoxPlotModelMapper_ConnectFirstBoxSetRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstBoxSetRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstBoxSetRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstBoxSetRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstBoxSetRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectFirstBoxSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectFirstBoxSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstBoxSetRowChanged")
	}
}

func (ptr *QHBoxPlotModelMapper) FirstBoxSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_FirstBoxSetRowChanged(ptr.Pointer())
	}
}

func (ptr *QHBoxPlotModelMapper) FirstColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBoxPlotModelMapper_FirstColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQHBoxPlotModelMapper_FirstColumnChanged
func callbackQHBoxPlotModelMapper_FirstColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectFirstColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstColumnChanged") {
			C.QHBoxPlotModelMapper_ConnectFirstColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectFirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectFirstColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstColumnChanged")
	}
}

func (ptr *QHBoxPlotModelMapper) FirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_FirstColumnChanged(ptr.Pointer())
	}
}

func (ptr *QHBoxPlotModelMapper) LastBoxSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBoxPlotModelMapper_LastBoxSetRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQHBoxPlotModelMapper_LastBoxSetRowChanged
func callbackQHBoxPlotModelMapper_LastBoxSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lastBoxSetRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectLastBoxSetRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lastBoxSetRowChanged") {
			C.QHBoxPlotModelMapper_ConnectLastBoxSetRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lastBoxSetRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lastBoxSetRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lastBoxSetRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastBoxSetRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectLastBoxSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectLastBoxSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lastBoxSetRowChanged")
	}
}

func (ptr *QHBoxPlotModelMapper) LastBoxSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_LastBoxSetRowChanged(ptr.Pointer())
	}
}

func (ptr *QHBoxPlotModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QHBoxPlotModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHBoxPlotModelMapper_ModelReplaced
func callbackQHBoxPlotModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "modelReplaced") {
			C.QHBoxPlotModelMapper_ConnectModelReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "modelReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "modelReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "modelReplaced")
	}
}

func (ptr *QHBoxPlotModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ModelReplaced(ptr.Pointer())
	}
}

func (ptr *QHBoxPlotModelMapper) Series() *QBoxPlotSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQBoxPlotSeriesFromPointer(C.QHBoxPlotModelMapper_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHBoxPlotModelMapper_SeriesReplaced
func callbackQHBoxPlotModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "seriesReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "seriesReplaced") {
			C.QHBoxPlotModelMapper_ConnectSeriesReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "seriesReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "seriesReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "seriesReplaced")
	}
}

func (ptr *QHBoxPlotModelMapper) SeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_SeriesReplaced(ptr.Pointer())
	}
}

func (ptr *QHBoxPlotModelMapper) SetColumnCount(rowCount int) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_SetColumnCount(ptr.Pointer(), C.int(int32(rowCount)))
	}
}

func (ptr *QHBoxPlotModelMapper) SetFirstBoxSetRow(firstBoxSetRow int) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_SetFirstBoxSetRow(ptr.Pointer(), C.int(int32(firstBoxSetRow)))
	}
}

func (ptr *QHBoxPlotModelMapper) SetFirstColumn(firstColumn int) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_SetFirstColumn(ptr.Pointer(), C.int(int32(firstColumn)))
	}
}

func (ptr *QHBoxPlotModelMapper) SetLastBoxSetRow(lastBoxSetRow int) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_SetLastBoxSetRow(ptr.Pointer(), C.int(int32(lastBoxSetRow)))
	}
}

func (ptr *QHBoxPlotModelMapper) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHBoxPlotModelMapper) SetSeries(series QBoxPlotSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_SetSeries(ptr.Pointer(), PointerFromQBoxPlotSeries(series))
	}
}

func (ptr *QHBoxPlotModelMapper) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHBoxPlotModelMapper___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHBoxPlotModelMapper) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHBoxPlotModelMapper) __children_newList() unsafe.Pointer {
	return C.QHBoxPlotModelMapper___children_newList(ptr.Pointer())
}

func (ptr *QHBoxPlotModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QHBoxPlotModelMapper___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHBoxPlotModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QHBoxPlotModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QHBoxPlotModelMapper___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QHBoxPlotModelMapper) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHBoxPlotModelMapper___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHBoxPlotModelMapper) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHBoxPlotModelMapper) __findChildren_newList() unsafe.Pointer {
	return C.QHBoxPlotModelMapper___findChildren_newList(ptr.Pointer())
}

func (ptr *QHBoxPlotModelMapper) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QHBoxPlotModelMapper___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHBoxPlotModelMapper) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QHBoxPlotModelMapper) __findChildren_newList3() unsafe.Pointer {
	return C.QHBoxPlotModelMapper___findChildren_newList3(ptr.Pointer())
}

//export callbackQHBoxPlotModelMapper_ChildEvent
func callbackQHBoxPlotModelMapper_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHBoxPlotModelMapperFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHBoxPlotModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHBoxPlotModelMapper_ConnectNotify
func callbackQHBoxPlotModelMapper_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHBoxPlotModelMapperFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHBoxPlotModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHBoxPlotModelMapper_CustomEvent
func callbackQHBoxPlotModelMapper_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQHBoxPlotModelMapperFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHBoxPlotModelMapper) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHBoxPlotModelMapper_DeleteLater
func callbackQHBoxPlotModelMapper_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHBoxPlotModelMapperFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHBoxPlotModelMapper) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHBoxPlotModelMapper_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQHBoxPlotModelMapper_Destroyed
func callbackQHBoxPlotModelMapper_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQHBoxPlotModelMapper_DisconnectNotify
func callbackQHBoxPlotModelMapper_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHBoxPlotModelMapperFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHBoxPlotModelMapper_Event
func callbackQHBoxPlotModelMapper_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHBoxPlotModelMapperFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHBoxPlotModelMapper) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHBoxPlotModelMapper_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQHBoxPlotModelMapper_EventFilter
func callbackQHBoxPlotModelMapper_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHBoxPlotModelMapperFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHBoxPlotModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHBoxPlotModelMapper_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQHBoxPlotModelMapper_MetaObject
func callbackQHBoxPlotModelMapper_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQHBoxPlotModelMapperFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHBoxPlotModelMapper) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHBoxPlotModelMapper_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQHBoxPlotModelMapper_ObjectNameChanged
func callbackQHBoxPlotModelMapper_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQHBoxPlotModelMapper_TimerEvent
func callbackQHBoxPlotModelMapper_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHBoxPlotModelMapperFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHBoxPlotModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QHCandlestickModelMapper struct {
	QCandlestickModelMapper
}

type QHCandlestickModelMapper_ITF interface {
	QCandlestickModelMapper_ITF
	QHCandlestickModelMapper_PTR() *QHCandlestickModelMapper
}

func (ptr *QHCandlestickModelMapper) QHCandlestickModelMapper_PTR() *QHCandlestickModelMapper {
	return ptr
}

func (ptr *QHCandlestickModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QCandlestickModelMapper_PTR().Pointer()
	}
	return nil
}

func (ptr *QHCandlestickModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QCandlestickModelMapper_PTR().SetPointer(p)
	}
}

func PointerFromQHCandlestickModelMapper(ptr QHCandlestickModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHCandlestickModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQHCandlestickModelMapperFromPointer(ptr unsafe.Pointer) (n *QHCandlestickModelMapper) {
	n = new(QHCandlestickModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQHCandlestickModelMapper(parent core.QObject_ITF) *QHCandlestickModelMapper {
	tmpValue := NewQHCandlestickModelMapperFromPointer(C.QHCandlestickModelMapper_NewQHCandlestickModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHCandlestickModelMapper) CloseColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_CloseColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQHCandlestickModelMapper_CloseColumnChanged
func callbackQHCandlestickModelMapper_CloseColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectCloseColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "closeColumnChanged") {
			C.QHCandlestickModelMapper_ConnectCloseColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "closeColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "closeColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "closeColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "closeColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectCloseColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectCloseColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "closeColumnChanged")
	}
}

func (ptr *QHCandlestickModelMapper) CloseColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_CloseColumnChanged(ptr.Pointer())
	}
}

func (ptr *QHCandlestickModelMapper) FirstSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_FirstSetRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQHCandlestickModelMapper_FirstSetRowChanged
func callbackQHCandlestickModelMapper_FirstSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstSetRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectFirstSetRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstSetRowChanged") {
			C.QHCandlestickModelMapper_ConnectFirstSetRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstSetRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstSetRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstSetRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstSetRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectFirstSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectFirstSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstSetRowChanged")
	}
}

func (ptr *QHCandlestickModelMapper) FirstSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_FirstSetRowChanged(ptr.Pointer())
	}
}

func (ptr *QHCandlestickModelMapper) HighColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_HighColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQHCandlestickModelMapper_HighColumnChanged
func callbackQHCandlestickModelMapper_HighColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "highColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectHighColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "highColumnChanged") {
			C.QHCandlestickModelMapper_ConnectHighColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "highColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "highColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "highColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "highColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectHighColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectHighColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "highColumnChanged")
	}
}

func (ptr *QHCandlestickModelMapper) HighColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_HighColumnChanged(ptr.Pointer())
	}
}

func (ptr *QHCandlestickModelMapper) LastSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_LastSetRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQHCandlestickModelMapper_LastSetRowChanged
func callbackQHCandlestickModelMapper_LastSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lastSetRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectLastSetRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lastSetRowChanged") {
			C.QHCandlestickModelMapper_ConnectLastSetRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lastSetRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lastSetRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lastSetRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastSetRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectLastSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectLastSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lastSetRowChanged")
	}
}

func (ptr *QHCandlestickModelMapper) LastSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_LastSetRowChanged(ptr.Pointer())
	}
}

func (ptr *QHCandlestickModelMapper) LowColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_LowColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQHCandlestickModelMapper_LowColumnChanged
func callbackQHCandlestickModelMapper_LowColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lowColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectLowColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lowColumnChanged") {
			C.QHCandlestickModelMapper_ConnectLowColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lowColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lowColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lowColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lowColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectLowColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectLowColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lowColumnChanged")
	}
}

func (ptr *QHCandlestickModelMapper) LowColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_LowColumnChanged(ptr.Pointer())
	}
}

func (ptr *QHCandlestickModelMapper) OpenColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_OpenColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQHCandlestickModelMapper_OpenColumnChanged
func callbackQHCandlestickModelMapper_OpenColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "openColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectOpenColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "openColumnChanged") {
			C.QHCandlestickModelMapper_ConnectOpenColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "openColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "openColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "openColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "openColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectOpenColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectOpenColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "openColumnChanged")
	}
}

func (ptr *QHCandlestickModelMapper) OpenColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_OpenColumnChanged(ptr.Pointer())
	}
}

//export callbackQHCandlestickModelMapper_Orientation
func callbackQHCandlestickModelMapper_Orientation(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "orientation"); signal != nil {
		return C.longlong((*(*func() core.Qt__Orientation)(signal))())
	}

	return C.longlong(NewQHCandlestickModelMapperFromPointer(ptr).OrientationDefault())
}

func (ptr *QHCandlestickModelMapper) ConnectOrientation(f func() core.Qt__Orientation) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "orientation"); signal != nil {
			f := func() core.Qt__Orientation {
				(*(*func() core.Qt__Orientation)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "orientation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "orientation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectOrientation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "orientation")
	}
}

func (ptr *QHCandlestickModelMapper) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QHCandlestickModelMapper_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHCandlestickModelMapper) OrientationDefault() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QHCandlestickModelMapper_OrientationDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHCandlestickModelMapper) SetCloseColumn(closeColumn int) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_SetCloseColumn(ptr.Pointer(), C.int(int32(closeColumn)))
	}
}

func (ptr *QHCandlestickModelMapper) SetFirstSetRow(firstSetRow int) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_SetFirstSetRow(ptr.Pointer(), C.int(int32(firstSetRow)))
	}
}

func (ptr *QHCandlestickModelMapper) SetHighColumn(highColumn int) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_SetHighColumn(ptr.Pointer(), C.int(int32(highColumn)))
	}
}

func (ptr *QHCandlestickModelMapper) SetLastSetRow(lastSetRow int) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_SetLastSetRow(ptr.Pointer(), C.int(int32(lastSetRow)))
	}
}

func (ptr *QHCandlestickModelMapper) SetLowColumn(lowColumn int) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_SetLowColumn(ptr.Pointer(), C.int(int32(lowColumn)))
	}
}

func (ptr *QHCandlestickModelMapper) SetOpenColumn(openColumn int) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_SetOpenColumn(ptr.Pointer(), C.int(int32(openColumn)))
	}
}

func (ptr *QHCandlestickModelMapper) SetTimestampColumn(timestampColumn int) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_SetTimestampColumn(ptr.Pointer(), C.int(int32(timestampColumn)))
	}
}

func (ptr *QHCandlestickModelMapper) TimestampColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_TimestampColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQHCandlestickModelMapper_TimestampColumnChanged
func callbackQHCandlestickModelMapper_TimestampColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timestampColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectTimestampColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "timestampColumnChanged") {
			C.QHCandlestickModelMapper_ConnectTimestampColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "timestampColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "timestampColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "timestampColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "timestampColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectTimestampColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectTimestampColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "timestampColumnChanged")
	}
}

func (ptr *QHCandlestickModelMapper) TimestampColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_TimestampColumnChanged(ptr.Pointer())
	}
}

type QHPieModelMapper struct {
	QPieModelMapper
}

type QHPieModelMapper_ITF interface {
	QPieModelMapper_ITF
	QHPieModelMapper_PTR() *QHPieModelMapper
}

func (ptr *QHPieModelMapper) QHPieModelMapper_PTR() *QHPieModelMapper {
	return ptr
}

func (ptr *QHPieModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPieModelMapper_PTR().Pointer()
	}
	return nil
}

func (ptr *QHPieModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPieModelMapper_PTR().SetPointer(p)
	}
}

func PointerFromQHPieModelMapper(ptr QHPieModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHPieModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQHPieModelMapperFromPointer(ptr unsafe.Pointer) (n *QHPieModelMapper) {
	n = new(QHPieModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQHPieModelMapper(parent core.QObject_ITF) *QHPieModelMapper {
	tmpValue := NewQHPieModelMapperFromPointer(C.QHPieModelMapper_NewQHPieModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHPieModelMapper) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHPieModelMapper_ColumnCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQHPieModelMapper_ColumnCountChanged
func callbackQHPieModelMapper_ColumnCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "columnCountChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHPieModelMapper) ConnectColumnCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "columnCountChanged") {
			C.QHPieModelMapper_ConnectColumnCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "columnCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "columnCountChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "columnCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "columnCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHPieModelMapper) DisconnectColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectColumnCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "columnCountChanged")
	}
}

func (ptr *QHPieModelMapper) ColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ColumnCountChanged(ptr.Pointer())
	}
}

func (ptr *QHPieModelMapper) FirstColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHPieModelMapper_FirstColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQHPieModelMapper_FirstColumnChanged
func callbackQHPieModelMapper_FirstColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHPieModelMapper) ConnectFirstColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstColumnChanged") {
			C.QHPieModelMapper_ConnectFirstColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHPieModelMapper) DisconnectFirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectFirstColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstColumnChanged")
	}
}

func (ptr *QHPieModelMapper) FirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_FirstColumnChanged(ptr.Pointer())
	}
}

func (ptr *QHPieModelMapper) LabelsRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHPieModelMapper_LabelsRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQHPieModelMapper_LabelsRowChanged
func callbackQHPieModelMapper_LabelsRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelsRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHPieModelMapper) ConnectLabelsRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsRowChanged") {
			C.QHPieModelMapper_ConnectLabelsRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHPieModelMapper) DisconnectLabelsRowChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectLabelsRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsRowChanged")
	}
}

func (ptr *QHPieModelMapper) LabelsRowChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_LabelsRowChanged(ptr.Pointer())
	}
}

func (ptr *QHPieModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QHPieModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHPieModelMapper_ModelReplaced
func callbackQHPieModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHPieModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "modelReplaced") {
			C.QHPieModelMapper_ConnectModelReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "modelReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "modelReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHPieModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "modelReplaced")
	}
}

func (ptr *QHPieModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ModelReplaced(ptr.Pointer())
	}
}

func (ptr *QHPieModelMapper) Series() *QPieSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQPieSeriesFromPointer(C.QHPieModelMapper_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHPieModelMapper_SeriesReplaced
func callbackQHPieModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "seriesReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHPieModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "seriesReplaced") {
			C.QHPieModelMapper_ConnectSeriesReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "seriesReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "seriesReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHPieModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "seriesReplaced")
	}
}

func (ptr *QHPieModelMapper) SeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_SeriesReplaced(ptr.Pointer())
	}
}

func (ptr *QHPieModelMapper) SetColumnCount(columnCount int) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_SetColumnCount(ptr.Pointer(), C.int(int32(columnCount)))
	}
}

func (ptr *QHPieModelMapper) SetFirstColumn(firstColumn int) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_SetFirstColumn(ptr.Pointer(), C.int(int32(firstColumn)))
	}
}

func (ptr *QHPieModelMapper) SetLabelsRow(labelsRow int) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_SetLabelsRow(ptr.Pointer(), C.int(int32(labelsRow)))
	}
}

func (ptr *QHPieModelMapper) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHPieModelMapper) SetSeries(series QPieSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_SetSeries(ptr.Pointer(), PointerFromQPieSeries(series))
	}
}

func (ptr *QHPieModelMapper) SetValuesRow(valuesRow int) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_SetValuesRow(ptr.Pointer(), C.int(int32(valuesRow)))
	}
}

func (ptr *QHPieModelMapper) ValuesRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHPieModelMapper_ValuesRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQHPieModelMapper_ValuesRowChanged
func callbackQHPieModelMapper_ValuesRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "valuesRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHPieModelMapper) ConnectValuesRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valuesRowChanged") {
			C.QHPieModelMapper_ConnectValuesRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "valuesRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valuesRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "valuesRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valuesRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHPieModelMapper) DisconnectValuesRowChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectValuesRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valuesRowChanged")
	}
}

func (ptr *QHPieModelMapper) ValuesRowChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ValuesRowChanged(ptr.Pointer())
	}
}

type QHXYModelMapper struct {
	QXYModelMapper
}

type QHXYModelMapper_ITF interface {
	QXYModelMapper_ITF
	QHXYModelMapper_PTR() *QHXYModelMapper
}

func (ptr *QHXYModelMapper) QHXYModelMapper_PTR() *QHXYModelMapper {
	return ptr
}

func (ptr *QHXYModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QXYModelMapper_PTR().Pointer()
	}
	return nil
}

func (ptr *QHXYModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QXYModelMapper_PTR().SetPointer(p)
	}
}

func PointerFromQHXYModelMapper(ptr QHXYModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHXYModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQHXYModelMapperFromPointer(ptr unsafe.Pointer) (n *QHXYModelMapper) {
	n = new(QHXYModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQHXYModelMapper(parent core.QObject_ITF) *QHXYModelMapper {
	tmpValue := NewQHXYModelMapperFromPointer(C.QHXYModelMapper_NewQHXYModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHXYModelMapper) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHXYModelMapper_ColumnCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQHXYModelMapper_ColumnCountChanged
func callbackQHXYModelMapper_ColumnCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "columnCountChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHXYModelMapper) ConnectColumnCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "columnCountChanged") {
			C.QHXYModelMapper_ConnectColumnCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "columnCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "columnCountChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "columnCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "columnCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHXYModelMapper) DisconnectColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectColumnCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "columnCountChanged")
	}
}

func (ptr *QHXYModelMapper) ColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_ColumnCountChanged(ptr.Pointer())
	}
}

func (ptr *QHXYModelMapper) FirstColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHXYModelMapper_FirstColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQHXYModelMapper_FirstColumnChanged
func callbackQHXYModelMapper_FirstColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHXYModelMapper) ConnectFirstColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstColumnChanged") {
			C.QHXYModelMapper_ConnectFirstColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHXYModelMapper) DisconnectFirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectFirstColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstColumnChanged")
	}
}

func (ptr *QHXYModelMapper) FirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_FirstColumnChanged(ptr.Pointer())
	}
}

func (ptr *QHXYModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QHXYModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHXYModelMapper_ModelReplaced
func callbackQHXYModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHXYModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "modelReplaced") {
			C.QHXYModelMapper_ConnectModelReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "modelReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "modelReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHXYModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "modelReplaced")
	}
}

func (ptr *QHXYModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_ModelReplaced(ptr.Pointer())
	}
}

func (ptr *QHXYModelMapper) Series() *QXYSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQXYSeriesFromPointer(C.QHXYModelMapper_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHXYModelMapper_SeriesReplaced
func callbackQHXYModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "seriesReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHXYModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "seriesReplaced") {
			C.QHXYModelMapper_ConnectSeriesReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "seriesReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "seriesReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHXYModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "seriesReplaced")
	}
}

func (ptr *QHXYModelMapper) SeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_SeriesReplaced(ptr.Pointer())
	}
}

func (ptr *QHXYModelMapper) SetColumnCount(columnCount int) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_SetColumnCount(ptr.Pointer(), C.int(int32(columnCount)))
	}
}

func (ptr *QHXYModelMapper) SetFirstColumn(firstColumn int) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_SetFirstColumn(ptr.Pointer(), C.int(int32(firstColumn)))
	}
}

func (ptr *QHXYModelMapper) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHXYModelMapper) SetSeries(series QXYSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_SetSeries(ptr.Pointer(), PointerFromQXYSeries(series))
	}
}

func (ptr *QHXYModelMapper) SetXRow(xRow int) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_SetXRow(ptr.Pointer(), C.int(int32(xRow)))
	}
}

func (ptr *QHXYModelMapper) SetYRow(yRow int) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_SetYRow(ptr.Pointer(), C.int(int32(yRow)))
	}
}

func (ptr *QHXYModelMapper) XRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHXYModelMapper_XRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQHXYModelMapper_XRowChanged
func callbackQHXYModelMapper_XRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "xRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHXYModelMapper) ConnectXRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "xRowChanged") {
			C.QHXYModelMapper_ConnectXRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "xRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "xRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "xRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "xRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHXYModelMapper) DisconnectXRowChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectXRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "xRowChanged")
	}
}

func (ptr *QHXYModelMapper) XRowChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_XRowChanged(ptr.Pointer())
	}
}

func (ptr *QHXYModelMapper) YRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHXYModelMapper_YRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQHXYModelMapper_YRowChanged
func callbackQHXYModelMapper_YRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "yRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QHXYModelMapper) ConnectYRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "yRowChanged") {
			C.QHXYModelMapper_ConnectYRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "yRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "yRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "yRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "yRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHXYModelMapper) DisconnectYRowChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectYRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "yRowChanged")
	}
}

func (ptr *QHXYModelMapper) YRowChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_YRowChanged(ptr.Pointer())
	}
}

type QHorizontalBarSeries struct {
	QAbstractBarSeries
}

type QHorizontalBarSeries_ITF interface {
	QAbstractBarSeries_ITF
	QHorizontalBarSeries_PTR() *QHorizontalBarSeries
}

func (ptr *QHorizontalBarSeries) QHorizontalBarSeries_PTR() *QHorizontalBarSeries {
	return ptr
}

func (ptr *QHorizontalBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractBarSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QHorizontalBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractBarSeries_PTR().SetPointer(p)
	}
}

func PointerFromQHorizontalBarSeries(ptr QHorizontalBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHorizontalBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQHorizontalBarSeriesFromPointer(ptr unsafe.Pointer) (n *QHorizontalBarSeries) {
	n = new(QHorizontalBarSeries)
	n.SetPointer(ptr)
	return
}
func NewQHorizontalBarSeries(parent core.QObject_ITF) *QHorizontalBarSeries {
	tmpValue := NewQHorizontalBarSeriesFromPointer(C.QHorizontalBarSeries_NewQHorizontalBarSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQHorizontalBarSeries_Type
func callbackQHorizontalBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQHorizontalBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QHorizontalBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHorizontalBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QHorizontalBarSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QHorizontalBarSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHorizontalBarSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QHorizontalBarSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHorizontalBarSeries_DestroyQHorizontalBarSeries
func callbackQHorizontalBarSeries_DestroyQHorizontalBarSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHorizontalBarSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHorizontalBarSeriesFromPointer(ptr).DestroyQHorizontalBarSeriesDefault()
	}
}

func (ptr *QHorizontalBarSeries) ConnectDestroyQHorizontalBarSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHorizontalBarSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHorizontalBarSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHorizontalBarSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHorizontalBarSeries) DisconnectDestroyQHorizontalBarSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHorizontalBarSeries")
	}
}

func (ptr *QHorizontalBarSeries) DestroyQHorizontalBarSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHorizontalBarSeries_DestroyQHorizontalBarSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHorizontalBarSeries) DestroyQHorizontalBarSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHorizontalBarSeries_DestroyQHorizontalBarSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QHorizontalPercentBarSeries struct {
	QAbstractBarSeries
}

type QHorizontalPercentBarSeries_ITF interface {
	QAbstractBarSeries_ITF
	QHorizontalPercentBarSeries_PTR() *QHorizontalPercentBarSeries
}

func (ptr *QHorizontalPercentBarSeries) QHorizontalPercentBarSeries_PTR() *QHorizontalPercentBarSeries {
	return ptr
}

func (ptr *QHorizontalPercentBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractBarSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QHorizontalPercentBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractBarSeries_PTR().SetPointer(p)
	}
}

func PointerFromQHorizontalPercentBarSeries(ptr QHorizontalPercentBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHorizontalPercentBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQHorizontalPercentBarSeriesFromPointer(ptr unsafe.Pointer) (n *QHorizontalPercentBarSeries) {
	n = new(QHorizontalPercentBarSeries)
	n.SetPointer(ptr)
	return
}
func NewQHorizontalPercentBarSeries(parent core.QObject_ITF) *QHorizontalPercentBarSeries {
	tmpValue := NewQHorizontalPercentBarSeriesFromPointer(C.QHorizontalPercentBarSeries_NewQHorizontalPercentBarSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQHorizontalPercentBarSeries_Type
func callbackQHorizontalPercentBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQHorizontalPercentBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QHorizontalPercentBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHorizontalPercentBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QHorizontalPercentBarSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QHorizontalPercentBarSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHorizontalPercentBarSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QHorizontalPercentBarSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHorizontalPercentBarSeries_DestroyQHorizontalPercentBarSeries
func callbackQHorizontalPercentBarSeries_DestroyQHorizontalPercentBarSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHorizontalPercentBarSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHorizontalPercentBarSeriesFromPointer(ptr).DestroyQHorizontalPercentBarSeriesDefault()
	}
}

func (ptr *QHorizontalPercentBarSeries) ConnectDestroyQHorizontalPercentBarSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHorizontalPercentBarSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHorizontalPercentBarSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHorizontalPercentBarSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHorizontalPercentBarSeries) DisconnectDestroyQHorizontalPercentBarSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHorizontalPercentBarSeries")
	}
}

func (ptr *QHorizontalPercentBarSeries) DestroyQHorizontalPercentBarSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHorizontalPercentBarSeries_DestroyQHorizontalPercentBarSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHorizontalPercentBarSeries) DestroyQHorizontalPercentBarSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHorizontalPercentBarSeries_DestroyQHorizontalPercentBarSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QHorizontalStackedBarSeries struct {
	QAbstractBarSeries
}

type QHorizontalStackedBarSeries_ITF interface {
	QAbstractBarSeries_ITF
	QHorizontalStackedBarSeries_PTR() *QHorizontalStackedBarSeries
}

func (ptr *QHorizontalStackedBarSeries) QHorizontalStackedBarSeries_PTR() *QHorizontalStackedBarSeries {
	return ptr
}

func (ptr *QHorizontalStackedBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractBarSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QHorizontalStackedBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractBarSeries_PTR().SetPointer(p)
	}
}

func PointerFromQHorizontalStackedBarSeries(ptr QHorizontalStackedBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHorizontalStackedBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQHorizontalStackedBarSeriesFromPointer(ptr unsafe.Pointer) (n *QHorizontalStackedBarSeries) {
	n = new(QHorizontalStackedBarSeries)
	n.SetPointer(ptr)
	return
}
func NewQHorizontalStackedBarSeries(parent core.QObject_ITF) *QHorizontalStackedBarSeries {
	tmpValue := NewQHorizontalStackedBarSeriesFromPointer(C.QHorizontalStackedBarSeries_NewQHorizontalStackedBarSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQHorizontalStackedBarSeries_Type
func callbackQHorizontalStackedBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQHorizontalStackedBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QHorizontalStackedBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHorizontalStackedBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QHorizontalStackedBarSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QHorizontalStackedBarSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHorizontalStackedBarSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QHorizontalStackedBarSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHorizontalStackedBarSeries_DestroyQHorizontalStackedBarSeries
func callbackQHorizontalStackedBarSeries_DestroyQHorizontalStackedBarSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHorizontalStackedBarSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHorizontalStackedBarSeriesFromPointer(ptr).DestroyQHorizontalStackedBarSeriesDefault()
	}
}

func (ptr *QHorizontalStackedBarSeries) ConnectDestroyQHorizontalStackedBarSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHorizontalStackedBarSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHorizontalStackedBarSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHorizontalStackedBarSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHorizontalStackedBarSeries) DisconnectDestroyQHorizontalStackedBarSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHorizontalStackedBarSeries")
	}
}

func (ptr *QHorizontalStackedBarSeries) DestroyQHorizontalStackedBarSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHorizontalStackedBarSeries_DestroyQHorizontalStackedBarSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHorizontalStackedBarSeries) DestroyQHorizontalStackedBarSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHorizontalStackedBarSeries_DestroyQHorizontalStackedBarSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLegend struct {
	widgets.QGraphicsWidget
}

type QLegend_ITF interface {
	widgets.QGraphicsWidget_ITF
	QLegend_PTR() *QLegend
}

func (ptr *QLegend) QLegend_PTR() *QLegend {
	return ptr
}

func (ptr *QLegend) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QLegend) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsWidget_PTR().SetPointer(p)
	}
}

func PointerFromQLegend(ptr QLegend_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLegend_PTR().Pointer()
	}
	return nil
}

func NewQLegendFromPointer(ptr unsafe.Pointer) (n *QLegend) {
	n = new(QLegend)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QLegend__MarkerShape
//QLegend::MarkerShape
type QLegend__MarkerShape int64

const (
	QLegend__MarkerShapeDefault    QLegend__MarkerShape = QLegend__MarkerShape(0)
	QLegend__MarkerShapeRectangle  QLegend__MarkerShape = QLegend__MarkerShape(1)
	QLegend__MarkerShapeCircle     QLegend__MarkerShape = QLegend__MarkerShape(2)
	QLegend__MarkerShapeFromSeries QLegend__MarkerShape = QLegend__MarkerShape(3)
)

func (ptr *QLegend) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLegend_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLegend) AttachToChart() {
	if ptr.Pointer() != nil {
		C.QLegend_AttachToChart(ptr.Pointer())
	}
}

//export callbackQLegend_BackgroundVisibleChanged
func callbackQLegend_BackgroundVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "backgroundVisibleChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QLegend) ConnectBackgroundVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "backgroundVisibleChanged") {
			C.QLegend_ConnectBackgroundVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "backgroundVisibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "backgroundVisibleChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "backgroundVisibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "backgroundVisibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegend) DisconnectBackgroundVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectBackgroundVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "backgroundVisibleChanged")
	}
}

func (ptr *QLegend) BackgroundVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QLegend_BackgroundVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QLegend) BorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QLegend_BorderColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQLegend_BorderColorChanged
func callbackQLegend_BorderColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "borderColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QLegend) ConnectBorderColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "borderColorChanged") {
			C.QLegend_ConnectBorderColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "borderColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "borderColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "borderColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "borderColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegend) DisconnectBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "borderColorChanged")
	}
}

func (ptr *QLegend) BorderColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_BorderColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QLegend) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QLegend_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QLegend_Color(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQLegend_ColorChanged
func callbackQLegend_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "colorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QLegend) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "colorChanged") {
			C.QLegend_ConnectColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "colorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "colorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegend) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "colorChanged")
	}
}

func (ptr *QLegend) ColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QLegend) DetachFromChart() {
	if ptr.Pointer() != nil {
		C.QLegend_DetachFromChart(ptr.Pointer())
	}
}

//export callbackQLegend_FontChanged
func callbackQLegend_FontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fontChanged"); signal != nil {
		(*(*func(*gui.QFont))(signal))(func() *gui.QFont {
			tmpValue := gui.NewQFontFromPointer(font)
			qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
			return tmpValue
		}())
	}

}

func (ptr *QLegend) ConnectFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fontChanged") {
			C.QLegend_ConnectFontChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "fontChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fontChanged"); signal != nil {
			f := func(font *gui.QFont) {
				(*(*func(*gui.QFont))(signal))(font)
				f(font)
			}
			qt.ConnectSignal(ptr.Pointer(), "fontChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fontChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegend) DisconnectFontChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectFontChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fontChanged")
	}
}

func (ptr *QLegend) FontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_FontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QLegend) IsAttachedToChart() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_IsAttachedToChart(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLegend) IsBackgroundVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_IsBackgroundVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLegend) LabelBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QLegend_LabelBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) LabelColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QLegend_LabelColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQLegend_LabelColorChanged
func callbackQLegend_LabelColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QLegend) ConnectLabelColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelColorChanged") {
			C.QLegend_ConnectLabelColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegend) DisconnectLabelColorChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectLabelColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelColorChanged")
	}
}

func (ptr *QLegend) LabelColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_LabelColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QLegend) MarkerShape() QLegend__MarkerShape {
	if ptr.Pointer() != nil {
		return QLegend__MarkerShape(C.QLegend_MarkerShape(ptr.Pointer()))
	}
	return 0
}

//export callbackQLegend_MarkerShapeChanged
func callbackQLegend_MarkerShapeChanged(ptr unsafe.Pointer, shape C.longlong) {
	if signal := qt.GetSignal(ptr, "markerShapeChanged"); signal != nil {
		(*(*func(QLegend__MarkerShape))(signal))(QLegend__MarkerShape(shape))
	}

}

func (ptr *QLegend) ConnectMarkerShapeChanged(f func(shape QLegend__MarkerShape)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "markerShapeChanged") {
			C.QLegend_ConnectMarkerShapeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "markerShapeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "markerShapeChanged"); signal != nil {
			f := func(shape QLegend__MarkerShape) {
				(*(*func(QLegend__MarkerShape))(signal))(shape)
				f(shape)
			}
			qt.ConnectSignal(ptr.Pointer(), "markerShapeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "markerShapeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegend) DisconnectMarkerShapeChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectMarkerShapeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "markerShapeChanged")
	}
}

func (ptr *QLegend) MarkerShapeChanged(shape QLegend__MarkerShape) {
	if ptr.Pointer() != nil {
		C.QLegend_MarkerShapeChanged(ptr.Pointer(), C.longlong(shape))
	}
}

func (ptr *QLegend) Markers(series QAbstractSeries_ITF) []*QLegendMarker {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QLegendMarker {
			out := make([]*QLegendMarker, int(l.len))
			tmpList := NewQLegendFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__markers_atList(i)
			}
			return out
		}(C.QLegend_Markers(ptr.Pointer(), PointerFromQAbstractSeries(series)))
	}
	return make([]*QLegendMarker, 0)
}

func (ptr *QLegend) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QLegend_Pen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) ReverseMarkers() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_ReverseMarkers(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLegend_ReverseMarkersChanged
func callbackQLegend_ReverseMarkersChanged(ptr unsafe.Pointer, reverseMarkers C.char) {
	if signal := qt.GetSignal(ptr, "reverseMarkersChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(reverseMarkers) != 0)
	}

}

func (ptr *QLegend) ConnectReverseMarkersChanged(f func(reverseMarkers bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "reverseMarkersChanged") {
			C.QLegend_ConnectReverseMarkersChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "reverseMarkersChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "reverseMarkersChanged"); signal != nil {
			f := func(reverseMarkers bool) {
				(*(*func(bool))(signal))(reverseMarkers)
				f(reverseMarkers)
			}
			qt.ConnectSignal(ptr.Pointer(), "reverseMarkersChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reverseMarkersChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegend) DisconnectReverseMarkersChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectReverseMarkersChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "reverseMarkersChanged")
	}
}

func (ptr *QLegend) ReverseMarkersChanged(reverseMarkers bool) {
	if ptr.Pointer() != nil {
		C.QLegend_ReverseMarkersChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(reverseMarkers))))
	}
}

func (ptr *QLegend) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QLegend_SetAlignment(ptr.Pointer(), C.longlong(alignment))
	}
}

func (ptr *QLegend) SetBackgroundVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QLegend_SetBackgroundVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QLegend) SetBorderColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_SetBorderColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QLegend) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QLegend) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QLegend) SetLabelBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_SetLabelBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QLegend) SetLabelColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_SetLabelColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QLegend) SetMarkerShape(shape QLegend__MarkerShape) {
	if ptr.Pointer() != nil {
		C.QLegend_SetMarkerShape(ptr.Pointer(), C.longlong(shape))
	}
}

func (ptr *QLegend) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QLegend) SetReverseMarkers(reverseMarkers bool) {
	if ptr.Pointer() != nil {
		C.QLegend_SetReverseMarkers(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(reverseMarkers))))
	}
}

func (ptr *QLegend) SetShowToolTips(show bool) {
	if ptr.Pointer() != nil {
		C.QLegend_SetShowToolTips(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(show))))
	}
}

func (ptr *QLegend) ShowToolTips() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_ShowToolTips(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLegend_ShowToolTipsChanged
func callbackQLegend_ShowToolTipsChanged(ptr unsafe.Pointer, showToolTips C.char) {
	if signal := qt.GetSignal(ptr, "showToolTipsChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(showToolTips) != 0)
	}

}

func (ptr *QLegend) ConnectShowToolTipsChanged(f func(showToolTips bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "showToolTipsChanged") {
			C.QLegend_ConnectShowToolTipsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "showToolTipsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "showToolTipsChanged"); signal != nil {
			f := func(showToolTips bool) {
				(*(*func(bool))(signal))(showToolTips)
				f(showToolTips)
			}
			qt.ConnectSignal(ptr.Pointer(), "showToolTipsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "showToolTipsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegend) DisconnectShowToolTipsChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectShowToolTipsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "showToolTipsChanged")
	}
}

func (ptr *QLegend) ShowToolTipsChanged(showToolTips bool) {
	if ptr.Pointer() != nil {
		C.QLegend_ShowToolTipsChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(showToolTips))))
	}
}

//export callbackQLegend_DestroyQLegend
func callbackQLegend_DestroyQLegend(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLegend"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLegendFromPointer(ptr).DestroyQLegendDefault()
	}
}

func (ptr *QLegend) ConnectDestroyQLegend(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLegend"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLegend", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLegend", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegend) DisconnectDestroyQLegend() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLegend")
	}
}

func (ptr *QLegend) DestroyQLegend() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLegend_DestroyQLegend(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLegend) DestroyQLegendDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLegend_DestroyQLegendDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLegend) __markers_atList(i int) *QLegendMarker {
	if ptr.Pointer() != nil {
		tmpValue := NewQLegendMarkerFromPointer(C.QLegend___markers_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) __markers_setList(i QLegendMarker_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___markers_setList(ptr.Pointer(), PointerFromQLegendMarker(i))
	}
}

func (ptr *QLegend) __markers_newList() unsafe.Pointer {
	return C.QLegend___markers_newList(ptr.Pointer())
}

func (ptr *QLegend) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QLegend___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QLegend) __actions_newList() unsafe.Pointer {
	return C.QLegend___actions_newList(ptr.Pointer())
}

func (ptr *QLegend) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QLegend___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QLegend) __addActions_actions_newList() unsafe.Pointer {
	return C.QLegend___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QLegend) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QLegend___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QLegend) __insertActions_actions_newList() unsafe.Pointer {
	return C.QLegend___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QLegend) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLegend___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLegend) __children_newList() unsafe.Pointer {
	return C.QLegend___children_newList(ptr.Pointer())
}

func (ptr *QLegend) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLegend___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QLegend) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QLegend___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QLegend) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLegend___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLegend) __findChildren_newList() unsafe.Pointer {
	return C.QLegend___findChildren_newList(ptr.Pointer())
}

func (ptr *QLegend) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLegend___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLegend) __findChildren_newList3() unsafe.Pointer {
	return C.QLegend___findChildren_newList3(ptr.Pointer())
}

func (ptr *QLegend) __childItems_atList(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QLegend___childItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QLegend) __childItems_setList(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___childItems_setList(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QLegend) __childItems_newList() unsafe.Pointer {
	return C.QLegend___childItems_newList(ptr.Pointer())
}

func (ptr *QLegend) __collidingItems_atList(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QLegend___collidingItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QLegend) __collidingItems_setList(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___collidingItems_setList(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QLegend) __collidingItems_newList() unsafe.Pointer {
	return C.QLegend___collidingItems_newList(ptr.Pointer())
}

func (ptr *QLegend) __setTransformations_transformations_atList(i int) *widgets.QGraphicsTransform {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQGraphicsTransformFromPointer(C.QLegend___setTransformations_transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) __setTransformations_transformations_setList(i widgets.QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___setTransformations_transformations_setList(ptr.Pointer(), widgets.PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QLegend) __setTransformations_transformations_newList() unsafe.Pointer {
	return C.QLegend___setTransformations_transformations_newList(ptr.Pointer())
}

func (ptr *QLegend) __transformations_atList(i int) *widgets.QGraphicsTransform {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQGraphicsTransformFromPointer(C.QLegend___transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) __transformations_setList(i widgets.QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___transformations_setList(ptr.Pointer(), widgets.PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QLegend) __transformations_newList() unsafe.Pointer {
	return C.QLegend___transformations_newList(ptr.Pointer())
}

//export callbackQLegend_BoundingRect
func callbackQLegend_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF((*(*func() *core.QRectF)(signal))())
	}

	return core.PointerFromQRectF(NewQLegendFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QLegend) BoundingRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QLegend_BoundingRectDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQLegend_ChangeEvent
func callbackQLegend_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLegend) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLegend_Close
func callbackQLegend_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).CloseDefault())))
}

func (ptr *QLegend) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLegend_CloseEvent
func callbackQLegend_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*gui.QCloseEvent))(signal))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QLegend) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQLegend_Event
func callbackQLegend_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QLegend) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLegend_FocusInEvent
func callbackQLegend_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QLegend) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQLegend_FocusNextPrevChild
func callbackQLegend_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QLegend) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackQLegend_FocusOutEvent
func callbackQLegend_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*gui.QFocusEvent))(signal))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QLegend) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQLegend_GeometryChanged
func callbackQLegend_GeometryChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "geometryChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_GetContentsMargins
func callbackQLegend_GetContentsMargins(ptr unsafe.Pointer, left C.double, top C.double, right C.double, bottom C.double) {
	if signal := qt.GetSignal(ptr, "getContentsMargins"); signal != nil {
		(*(*func(float64, float64, float64, float64))(signal))(float64(left), float64(top), float64(right), float64(bottom))
	} else {
		NewQLegendFromPointer(ptr).GetContentsMarginsDefault(float64(left), float64(top), float64(right), float64(bottom))
	}
}

func (ptr *QLegend) GetContentsMarginsDefault(left float64, top float64, right float64, bottom float64) {
	if ptr.Pointer() != nil {
		C.QLegend_GetContentsMarginsDefault(ptr.Pointer(), C.double(left), C.double(top), C.double(right), C.double(bottom))
	}
}

//export callbackQLegend_GrabKeyboardEvent
func callbackQLegend_GrabKeyboardEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "grabKeyboardEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).GrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLegend) GrabKeyboardEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_GrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLegend_GrabMouseEvent
func callbackQLegend_GrabMouseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "grabMouseEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).GrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLegend) GrabMouseEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_GrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLegend_HideEvent
func callbackQLegend_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QLegend) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQLegend_HoverLeaveEvent
func callbackQLegend_HoverLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverLeaveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneHoverEvent))(signal))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).HoverLeaveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QLegend) HoverLeaveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_HoverLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQLegend_HoverMoveEvent
func callbackQLegend_HoverMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverMoveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneHoverEvent))(signal))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).HoverMoveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QLegend) HoverMoveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_HoverMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQLegend_InitStyleOption
func callbackQLegend_InitStyleOption(ptr unsafe.Pointer, option unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initStyleOption"); signal != nil {
		(*(*func(*widgets.QStyleOption))(signal))(widgets.NewQStyleOptionFromPointer(option))
	} else {
		NewQLegendFromPointer(ptr).InitStyleOptionDefault(widgets.NewQStyleOptionFromPointer(option))
	}
}

func (ptr *QLegend) InitStyleOptionDefault(option widgets.QStyleOption_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_InitStyleOptionDefault(ptr.Pointer(), widgets.PointerFromQStyleOption(option))
	}
}

//export callbackQLegend_ItemChange
func callbackQLegend_ItemChange(ptr unsafe.Pointer, change C.longlong, value unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemChange"); signal != nil {
		return core.PointerFromQVariant((*(*func(widgets.QGraphicsItem__GraphicsItemChange, *core.QVariant) *core.QVariant)(signal))(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
	}

	return core.PointerFromQVariant(NewQLegendFromPointer(ptr).ItemChangeDefault(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
}

func (ptr *QLegend) ItemChangeDefault(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QLegend_ItemChangeDefault(ptr.Pointer(), C.longlong(change), core.PointerFromQVariant(value)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQLegend_LayoutChanged
func callbackQLegend_LayoutChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_MoveEvent
func callbackQLegend_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneMoveEvent))(signal))(widgets.NewQGraphicsSceneMoveEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).MoveEventDefault(widgets.NewQGraphicsSceneMoveEventFromPointer(event))
	}
}

func (ptr *QLegend) MoveEventDefault(event widgets.QGraphicsSceneMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_MoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMoveEvent(event))
	}
}

//export callbackQLegend_Paint
func callbackQLegend_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		(*(*func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(signal))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQLegendFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QLegend) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQLegend_PaintWindowFrame
func callbackQLegend_PaintWindowFrame(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintWindowFrame"); signal != nil {
		(*(*func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(signal))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQLegendFromPointer(ptr).PaintWindowFrameDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QLegend) PaintWindowFrameDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_PaintWindowFrameDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQLegend_PolishEvent
func callbackQLegend_PolishEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "polishEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLegendFromPointer(ptr).PolishEventDefault()
	}
}

func (ptr *QLegend) PolishEventDefault() {
	if ptr.Pointer() != nil {
		C.QLegend_PolishEventDefault(ptr.Pointer())
	}
}

//export callbackQLegend_ResizeEvent
func callbackQLegend_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneResizeEvent))(signal))(widgets.NewQGraphicsSceneResizeEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).ResizeEventDefault(widgets.NewQGraphicsSceneResizeEventFromPointer(event))
	}
}

func (ptr *QLegend) ResizeEventDefault(event widgets.QGraphicsSceneResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_ResizeEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneResizeEvent(event))
	}
}

//export callbackQLegend_SceneEvent
func callbackQLegend_SceneEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "sceneEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).SceneEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QLegend) SceneEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_SceneEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLegend_SetGeometry
func callbackQLegend_SetGeometry(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		(*(*func(*core.QRectF))(signal))(core.NewQRectFFromPointer(rect))
	} else {
		NewQLegendFromPointer(ptr).SetGeometryDefault(core.NewQRectFFromPointer(rect))
	}
}

func (ptr *QLegend) SetGeometryDefault(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

//export callbackQLegend_Shape
func callbackQLegend_Shape(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "shape"); signal != nil {
		return gui.PointerFromQPainterPath((*(*func() *gui.QPainterPath)(signal))())
	}

	return gui.PointerFromQPainterPath(NewQLegendFromPointer(ptr).ShapeDefault())
}

func (ptr *QLegend) ShapeDefault() *gui.QPainterPath {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPainterPathFromPointer(C.QLegend_ShapeDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

//export callbackQLegend_ShowEvent
func callbackQLegend_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QLegend) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQLegend_SizeHint
func callbackQLegend_SizeHint(ptr unsafe.Pointer, which C.longlong, constraint unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSizeF((*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
	}

	return core.PointerFromQSizeF(NewQLegendFromPointer(ptr).SizeHintDefault(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
}

func (ptr *QLegend) SizeHintDefault(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QLegend_SizeHintDefault(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

//export callbackQLegend_Type
func callbackQLegend_Type(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQLegendFromPointer(ptr).TypeDefault()))
}

func (ptr *QLegend) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLegend_TypeDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQLegend_UngrabKeyboardEvent
func callbackQLegend_UngrabKeyboardEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ungrabKeyboardEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).UngrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLegend) UngrabKeyboardEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_UngrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLegend_UngrabMouseEvent
func callbackQLegend_UngrabMouseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ungrabMouseEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).UngrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLegend) UngrabMouseEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_UngrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLegend_UpdateGeometry
func callbackQLegend_UpdateGeometry(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateGeometry"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLegendFromPointer(ptr).UpdateGeometryDefault()
	}
}

func (ptr *QLegend) UpdateGeometryDefault() {
	if ptr.Pointer() != nil {
		C.QLegend_UpdateGeometryDefault(ptr.Pointer())
	}
}

//export callbackQLegend_WindowFrameEvent
func callbackQLegend_WindowFrameEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "windowFrameEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).WindowFrameEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QLegend) WindowFrameEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_WindowFrameEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLegend_WindowFrameSectionAt
func callbackQLegend_WindowFrameSectionAt(ptr unsafe.Pointer, pos unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "windowFrameSectionAt"); signal != nil {
		return C.longlong((*(*func(*core.QPointF) core.Qt__WindowFrameSection)(signal))(core.NewQPointFFromPointer(pos)))
	}

	return C.longlong(NewQLegendFromPointer(ptr).WindowFrameSectionAtDefault(core.NewQPointFFromPointer(pos)))
}

func (ptr *QLegend) WindowFrameSectionAtDefault(pos core.QPointF_ITF) core.Qt__WindowFrameSection {
	if ptr.Pointer() != nil {
		return core.Qt__WindowFrameSection(C.QLegend_WindowFrameSectionAtDefault(ptr.Pointer(), core.PointerFromQPointF(pos)))
	}
	return 0
}

//export callbackQLegend_EnabledChanged
func callbackQLegend_EnabledChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enabledChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_OpacityChanged
func callbackQLegend_OpacityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_ParentChanged
func callbackQLegend_ParentChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "parentChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_RotationChanged
func callbackQLegend_RotationChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rotationChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_ScaleChanged
func callbackQLegend_ScaleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scaleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_UpdateMicroFocus
func callbackQLegend_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLegendFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QLegend) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QLegend_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQLegend_VisibleChanged
func callbackQLegend_VisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_XChanged
func callbackQLegend_XChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "xChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_YChanged
func callbackQLegend_YChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "yChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_ZChanged
func callbackQLegend_ZChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "zChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQLegend_ChildEvent
func callbackQLegend_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLegend) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLegend_ConnectNotify
func callbackQLegend_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLegendFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLegend) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLegend_CustomEvent
func callbackQLegend_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLegend) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLegend_DeleteLater
func callbackQLegend_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLegendFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLegend) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLegend_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQLegend_Destroyed
func callbackQLegend_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQLegend_DisconnectNotify
func callbackQLegend_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLegendFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLegend) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLegend_EventFilter
func callbackQLegend_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLegend) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLegend_MetaObject
func callbackQLegend_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQLegendFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLegend) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLegend_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQLegend_ObjectNameChanged
func callbackQLegend_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQLegend_TimerEvent
func callbackQLegend_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLegend) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLegend_Advance
func callbackQLegend_Advance(ptr unsafe.Pointer, phase C.int) {
	if signal := qt.GetSignal(ptr, "advance"); signal != nil {
		(*(*func(int))(signal))(int(int32(phase)))
	} else {
		NewQLegendFromPointer(ptr).AdvanceDefault(int(int32(phase)))
	}
}

func (ptr *QLegend) AdvanceDefault(phase int) {
	if ptr.Pointer() != nil {
		C.QLegend_AdvanceDefault(ptr.Pointer(), C.int(int32(phase)))
	}
}

//export callbackQLegend_CollidesWithItem
func callbackQLegend_CollidesWithItem(ptr unsafe.Pointer, other unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "collidesWithItem"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*widgets.QGraphicsItem, core.Qt__ItemSelectionMode) bool)(signal))(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).CollidesWithItemDefault(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode)))))
}

func (ptr *QLegend) CollidesWithItemDefault(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_CollidesWithItemDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(other), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQLegend_CollidesWithPath
func callbackQLegend_CollidesWithPath(ptr unsafe.Pointer, path unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "collidesWithPath"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*gui.QPainterPath, core.Qt__ItemSelectionMode) bool)(signal))(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).CollidesWithPathDefault(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode)))))
}

func (ptr *QLegend) CollidesWithPathDefault(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_CollidesWithPathDefault(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQLegend_Contains
func callbackQLegend_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QPointF) bool)(signal))(core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point)))))
}

func (ptr *QLegend) ContainsDefault(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point))) != 0
	}
	return false
}

//export callbackQLegend_ContextMenuEvent
func callbackQLegend_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneContextMenuEvent))(signal))(widgets.NewQGraphicsSceneContextMenuEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).ContextMenuEventDefault(widgets.NewQGraphicsSceneContextMenuEventFromPointer(event))
	}
}

func (ptr *QLegend) ContextMenuEventDefault(event widgets.QGraphicsSceneContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_ContextMenuEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

//export callbackQLegend_DragEnterEvent
func callbackQLegend_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneDragDropEvent))(signal))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).DragEnterEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QLegend) DragEnterEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_DragEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQLegend_DragLeaveEvent
func callbackQLegend_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneDragDropEvent))(signal))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).DragLeaveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QLegend) DragLeaveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_DragLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQLegend_DragMoveEvent
func callbackQLegend_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneDragDropEvent))(signal))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).DragMoveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QLegend) DragMoveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_DragMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQLegend_DropEvent
func callbackQLegend_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneDragDropEvent))(signal))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).DropEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QLegend) DropEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_DropEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQLegend_HoverEnterEvent
func callbackQLegend_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverEnterEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneHoverEvent))(signal))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).HoverEnterEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QLegend) HoverEnterEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_HoverEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQLegend_InputMethodEvent
func callbackQLegend_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*gui.QInputMethodEvent))(signal))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QLegend) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQLegend_InputMethodQuery
func callbackQLegend_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant((*(*func(core.Qt__InputMethodQuery) *core.QVariant)(signal))(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQLegendFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QLegend) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QLegend_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQLegend_IsObscuredBy
func callbackQLegend_IsObscuredBy(ptr unsafe.Pointer, item unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isObscuredBy"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*widgets.QGraphicsItem) bool)(signal))(widgets.NewQGraphicsItemFromPointer(item)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).IsObscuredByDefault(widgets.NewQGraphicsItemFromPointer(item)))))
}

func (ptr *QLegend) IsObscuredByDefault(item widgets.QGraphicsItem_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_IsObscuredByDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

//export callbackQLegend_KeyPressEvent
func callbackQLegend_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QLegend) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQLegend_KeyReleaseEvent
func callbackQLegend_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QLegend) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQLegend_MouseDoubleClickEvent
func callbackQLegend_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneMouseEvent))(signal))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).MouseDoubleClickEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QLegend) MouseDoubleClickEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_MouseDoubleClickEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQLegend_MouseMoveEvent
func callbackQLegend_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneMouseEvent))(signal))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).MouseMoveEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QLegend) MouseMoveEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_MouseMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQLegend_MousePressEvent
func callbackQLegend_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneMouseEvent))(signal))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).MousePressEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QLegend) MousePressEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_MousePressEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQLegend_MouseReleaseEvent
func callbackQLegend_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneMouseEvent))(signal))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).MouseReleaseEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QLegend) MouseReleaseEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_MouseReleaseEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQLegend_OpaqueArea
func callbackQLegend_OpaqueArea(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "opaqueArea"); signal != nil {
		return gui.PointerFromQPainterPath((*(*func() *gui.QPainterPath)(signal))())
	}

	return gui.PointerFromQPainterPath(NewQLegendFromPointer(ptr).OpaqueAreaDefault())
}

func (ptr *QLegend) OpaqueAreaDefault() *gui.QPainterPath {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPainterPathFromPointer(C.QLegend_OpaqueAreaDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

//export callbackQLegend_SceneEventFilter
func callbackQLegend_SceneEventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "sceneEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*widgets.QGraphicsItem, *core.QEvent) bool)(signal))(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendFromPointer(ptr).SceneEventFilterDefault(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLegend) SceneEventFilterDefault(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegend_SceneEventFilterDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLegend_WheelEvent
func callbackQLegend_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*widgets.QGraphicsSceneWheelEvent))(signal))(widgets.NewQGraphicsSceneWheelEventFromPointer(event))
	} else {
		NewQLegendFromPointer(ptr).WheelEventDefault(widgets.NewQGraphicsSceneWheelEventFromPointer(event))
	}
}

func (ptr *QLegend) WheelEventDefault(event widgets.QGraphicsSceneWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_WheelEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneWheelEvent(event))
	}
}

type QLegendMarker struct {
	core.QObject
}

type QLegendMarker_ITF interface {
	core.QObject_ITF
	QLegendMarker_PTR() *QLegendMarker
}

func (ptr *QLegendMarker) QLegendMarker_PTR() *QLegendMarker {
	return ptr
}

func (ptr *QLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQLegendMarker(ptr QLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QLegendMarker) {
	n = new(QLegendMarker)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QLegendMarker__LegendMarkerType
//QLegendMarker::LegendMarkerType
type QLegendMarker__LegendMarkerType int64

const (
	QLegendMarker__LegendMarkerTypeArea        QLegendMarker__LegendMarkerType = QLegendMarker__LegendMarkerType(0)
	QLegendMarker__LegendMarkerTypeBar         QLegendMarker__LegendMarkerType = QLegendMarker__LegendMarkerType(1)
	QLegendMarker__LegendMarkerTypePie         QLegendMarker__LegendMarkerType = QLegendMarker__LegendMarkerType(2)
	QLegendMarker__LegendMarkerTypeXY          QLegendMarker__LegendMarkerType = QLegendMarker__LegendMarkerType(3)
	QLegendMarker__LegendMarkerTypeBoxPlot     QLegendMarker__LegendMarkerType = QLegendMarker__LegendMarkerType(4)
	QLegendMarker__LegendMarkerTypeCandlestick QLegendMarker__LegendMarkerType = QLegendMarker__LegendMarkerType(5)
)

func (ptr *QLegendMarker) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QLegendMarker_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQLegendMarker_BrushChanged
func callbackQLegendMarker_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "brushChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLegendMarker) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "brushChanged") {
			C.QLegendMarker_ConnectBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "brushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "brushChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "brushChanged")
	}
}

func (ptr *QLegendMarker) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_BrushChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_Clicked
func callbackQLegendMarker_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLegendMarker) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QLegendMarker_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QLegendMarker) Clicked() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_Clicked(ptr.Pointer())
	}
}

func (ptr *QLegendMarker) Font() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QLegendMarker_Font(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

//export callbackQLegendMarker_FontChanged
func callbackQLegendMarker_FontChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fontChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLegendMarker) ConnectFontChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fontChanged") {
			C.QLegendMarker_ConnectFontChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "fontChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fontChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "fontChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fontChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectFontChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectFontChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fontChanged")
	}
}

func (ptr *QLegendMarker) FontChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_FontChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_Hovered
func callbackQLegendMarker_Hovered(ptr unsafe.Pointer, status C.char) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(bool))(signal))(int8(status) != 0)
	}

}

func (ptr *QLegendMarker) ConnectHovered(f func(status bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QLegendMarker_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(status bool) {
				(*(*func(bool))(signal))(status)
				f(status)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QLegendMarker) Hovered(status bool) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))))
	}
}

func (ptr *QLegendMarker) IsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegendMarker_IsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLegendMarker) Label() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLegendMarker_Label(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLegendMarker) LabelBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QLegendMarker_LabelBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQLegendMarker_LabelBrushChanged
func callbackQLegendMarker_LabelBrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelBrushChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLegendMarker) ConnectLabelBrushChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelBrushChanged") {
			C.QLegendMarker_ConnectLabelBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelBrushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelBrushChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelBrushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelBrushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectLabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectLabelBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelBrushChanged")
	}
}

func (ptr *QLegendMarker) LabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_LabelBrushChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_LabelChanged
func callbackQLegendMarker_LabelChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLegendMarker) ConnectLabelChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelChanged") {
			C.QLegendMarker_ConnectLabelChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectLabelChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectLabelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelChanged")
	}
}

func (ptr *QLegendMarker) LabelChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_LabelChanged(ptr.Pointer())
	}
}

func (ptr *QLegendMarker) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QLegendMarker_Pen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQLegendMarker_PenChanged
func callbackQLegendMarker_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "penChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLegendMarker) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "penChanged") {
			C.QLegendMarker_ConnectPenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "penChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "penChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "penChanged")
	}
}

func (ptr *QLegendMarker) PenChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_PenChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_Series
func callbackQLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "series"); signal != nil {
		return PointerFromQAbstractSeries((*(*func() *QAbstractSeries)(signal))())
	}

	return PointerFromQAbstractSeries(nil)
}

func (ptr *QLegendMarker) ConnectSeries(f func() *QAbstractSeries) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "series"); signal != nil {
			f := func() *QAbstractSeries {
				(*(*func() *QAbstractSeries)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "series")
	}
}

func (ptr *QLegendMarker) Series() *QAbstractSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractSeriesFromPointer(C.QLegendMarker_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegendMarker) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QLegendMarker) SetFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QLegendMarker) SetLabel(label string) {
	if ptr.Pointer() != nil {
		var labelC *C.char
		if label != "" {
			labelC = C.CString(label)
			defer C.free(unsafe.Pointer(labelC))
		}
		C.QLegendMarker_SetLabel(ptr.Pointer(), C.struct_QtCharts_PackedString{data: labelC, len: C.longlong(len(label))})
	}
}

func (ptr *QLegendMarker) SetLabelBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_SetLabelBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QLegendMarker) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QLegendMarker) SetShape(shape QLegend__MarkerShape) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_SetShape(ptr.Pointer(), C.longlong(shape))
	}
}

func (ptr *QLegendMarker) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QLegendMarker) Shape() QLegend__MarkerShape {
	if ptr.Pointer() != nil {
		return QLegend__MarkerShape(C.QLegendMarker_Shape(ptr.Pointer()))
	}
	return 0
}

//export callbackQLegendMarker_ShapeChanged
func callbackQLegendMarker_ShapeChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "shapeChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLegendMarker) ConnectShapeChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "shapeChanged") {
			C.QLegendMarker_ConnectShapeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "shapeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "shapeChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "shapeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shapeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectShapeChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectShapeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "shapeChanged")
	}
}

func (ptr *QLegendMarker) ShapeChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ShapeChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_Type
func callbackQLegendMarker_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QLegendMarker__LegendMarkerType)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QLegendMarker__LegendMarkerType {
				(*(*func() QLegendMarker__LegendMarkerType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QLegendMarker) Type() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QLegendMarker_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQLegendMarker_VisibleChanged
func callbackQLegendMarker_VisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLegendMarker) ConnectVisibleChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "visibleChanged") {
			C.QLegendMarker_ConnectVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "visibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "visibleChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "visibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "visibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "visibleChanged")
	}
}

func (ptr *QLegendMarker) VisibleChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_VisibleChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_DestroyQLegendMarker
func callbackQLegendMarker_DestroyQLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLegendMarker"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLegendMarkerFromPointer(ptr).DestroyQLegendMarkerDefault()
	}
}

func (ptr *QLegendMarker) ConnectDestroyQLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLegendMarker"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLegendMarker", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLegendMarker", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLegendMarker) DisconnectDestroyQLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLegendMarker")
	}
}

func (ptr *QLegendMarker) DestroyQLegendMarker() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLegendMarker_DestroyQLegendMarker(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLegendMarker) DestroyQLegendMarkerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLegendMarker_DestroyQLegendMarkerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLegendMarker) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLegendMarker___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegendMarker) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLegendMarker) __children_newList() unsafe.Pointer {
	return C.QLegendMarker___children_newList(ptr.Pointer())
}

func (ptr *QLegendMarker) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLegendMarker___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLegendMarker) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QLegendMarker) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QLegendMarker___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QLegendMarker) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLegendMarker___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegendMarker) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLegendMarker) __findChildren_newList() unsafe.Pointer {
	return C.QLegendMarker___findChildren_newList(ptr.Pointer())
}

func (ptr *QLegendMarker) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLegendMarker___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLegendMarker) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLegendMarker) __findChildren_newList3() unsafe.Pointer {
	return C.QLegendMarker___findChildren_newList3(ptr.Pointer())
}

//export callbackQLegendMarker_ChildEvent
func callbackQLegendMarker_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLegendMarkerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLegendMarker) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLegendMarker_ConnectNotify
func callbackQLegendMarker_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLegendMarkerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLegendMarker) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLegendMarker_CustomEvent
func callbackQLegendMarker_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQLegendMarkerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLegendMarker) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLegendMarker_DeleteLater
func callbackQLegendMarker_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLegendMarkerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLegendMarker) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLegendMarker_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQLegendMarker_Destroyed
func callbackQLegendMarker_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQLegendMarker_DisconnectNotify
func callbackQLegendMarker_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLegendMarkerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLegendMarker) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLegendMarker_Event
func callbackQLegendMarker_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendMarkerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLegendMarker) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegendMarker_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQLegendMarker_EventFilter
func callbackQLegendMarker_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLegendMarkerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLegendMarker) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLegendMarker_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLegendMarker_MetaObject
func callbackQLegendMarker_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQLegendMarkerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLegendMarker) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLegendMarker_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQLegendMarker_ObjectNameChanged
func callbackQLegendMarker_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQLegendMarker_TimerEvent
func callbackQLegendMarker_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLegendMarkerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLegendMarker) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QLineSeries struct {
	QXYSeries
}

type QLineSeries_ITF interface {
	QXYSeries_ITF
	QLineSeries_PTR() *QLineSeries
}

func (ptr *QLineSeries) QLineSeries_PTR() *QLineSeries {
	return ptr
}

func (ptr *QLineSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QXYSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QLineSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QXYSeries_PTR().SetPointer(p)
	}
}

func PointerFromQLineSeries(ptr QLineSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLineSeries_PTR().Pointer()
	}
	return nil
}

func NewQLineSeriesFromPointer(ptr unsafe.Pointer) (n *QLineSeries) {
	n = new(QLineSeries)
	n.SetPointer(ptr)
	return
}
func NewQLineSeries(parent core.QObject_ITF) *QLineSeries {
	tmpValue := NewQLineSeriesFromPointer(C.QLineSeries_NewQLineSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQLineSeries_Type
func callbackQLineSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQLineSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QLineSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLineSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QLineSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QLineSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QLineSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLineSeries_DestroyQLineSeries
func callbackQLineSeries_DestroyQLineSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLineSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLineSeriesFromPointer(ptr).DestroyQLineSeriesDefault()
	}
}

func (ptr *QLineSeries) ConnectDestroyQLineSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLineSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLineSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLineSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLineSeries) DisconnectDestroyQLineSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLineSeries")
	}
}

func (ptr *QLineSeries) DestroyQLineSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLineSeries_DestroyQLineSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLineSeries) DestroyQLineSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLineSeries_DestroyQLineSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLogValueAxis struct {
	QAbstractAxis
}

type QLogValueAxis_ITF interface {
	QAbstractAxis_ITF
	QLogValueAxis_PTR() *QLogValueAxis
}

func (ptr *QLogValueAxis) QLogValueAxis_PTR() *QLogValueAxis {
	return ptr
}

func (ptr *QLogValueAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractAxis_PTR().Pointer()
	}
	return nil
}

func (ptr *QLogValueAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractAxis_PTR().SetPointer(p)
	}
}

func PointerFromQLogValueAxis(ptr QLogValueAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLogValueAxis_PTR().Pointer()
	}
	return nil
}

func NewQLogValueAxisFromPointer(ptr unsafe.Pointer) (n *QLogValueAxis) {
	n = new(QLogValueAxis)
	n.SetPointer(ptr)
	return
}
func NewQLogValueAxis(parent core.QObject_ITF) *QLogValueAxis {
	tmpValue := NewQLogValueAxisFromPointer(C.QLogValueAxis_NewQLogValueAxis(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLogValueAxis) Base() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLogValueAxis_Base(ptr.Pointer()))
	}
	return 0
}

//export callbackQLogValueAxis_BaseChanged
func callbackQLogValueAxis_BaseChanged(ptr unsafe.Pointer, base C.double) {
	if signal := qt.GetSignal(ptr, "baseChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(base))
	}

}

func (ptr *QLogValueAxis) ConnectBaseChanged(f func(base float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "baseChanged") {
			C.QLogValueAxis_ConnectBaseChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "baseChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "baseChanged"); signal != nil {
			f := func(base float64) {
				(*(*func(float64))(signal))(base)
				f(base)
			}
			qt.ConnectSignal(ptr.Pointer(), "baseChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "baseChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLogValueAxis) DisconnectBaseChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectBaseChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "baseChanged")
	}
}

func (ptr *QLogValueAxis) BaseChanged(base float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_BaseChanged(ptr.Pointer(), C.double(base))
	}
}

func (ptr *QLogValueAxis) LabelFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLogValueAxis_LabelFormat(ptr.Pointer()))
	}
	return ""
}

//export callbackQLogValueAxis_LabelFormatChanged
func callbackQLogValueAxis_LabelFormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "labelFormatChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(format))
	}

}

func (ptr *QLogValueAxis) ConnectLabelFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelFormatChanged") {
			C.QLogValueAxis_ConnectLabelFormatChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelFormatChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelFormatChanged"); signal != nil {
			f := func(format string) {
				(*(*func(string))(signal))(format)
				f(format)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelFormatChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelFormatChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLogValueAxis) DisconnectLabelFormatChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectLabelFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelFormatChanged")
	}
}

func (ptr *QLogValueAxis) LabelFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QLogValueAxis_LabelFormatChanged(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

func (ptr *QLogValueAxis) Max() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLogValueAxis_Max(ptr.Pointer()))
	}
	return 0
}

//export callbackQLogValueAxis_MaxChanged
func callbackQLogValueAxis_MaxChanged(ptr unsafe.Pointer, max C.double) {
	if signal := qt.GetSignal(ptr, "maxChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(max))
	}

}

func (ptr *QLogValueAxis) ConnectMaxChanged(f func(max float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "maxChanged") {
			C.QLogValueAxis_ConnectMaxChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "maxChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "maxChanged"); signal != nil {
			f := func(max float64) {
				(*(*func(float64))(signal))(max)
				f(max)
			}
			qt.ConnectSignal(ptr.Pointer(), "maxChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maxChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLogValueAxis) DisconnectMaxChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectMaxChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "maxChanged")
	}
}

func (ptr *QLogValueAxis) MaxChanged(max float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_MaxChanged(ptr.Pointer(), C.double(max))
	}
}

func (ptr *QLogValueAxis) Min() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLogValueAxis_Min(ptr.Pointer()))
	}
	return 0
}

//export callbackQLogValueAxis_MinChanged
func callbackQLogValueAxis_MinChanged(ptr unsafe.Pointer, min C.double) {
	if signal := qt.GetSignal(ptr, "minChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(min))
	}

}

func (ptr *QLogValueAxis) ConnectMinChanged(f func(min float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minChanged") {
			C.QLogValueAxis_ConnectMinChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minChanged"); signal != nil {
			f := func(min float64) {
				(*(*func(float64))(signal))(min)
				f(min)
			}
			qt.ConnectSignal(ptr.Pointer(), "minChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLogValueAxis) DisconnectMinChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectMinChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minChanged")
	}
}

func (ptr *QLogValueAxis) MinChanged(min float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_MinChanged(ptr.Pointer(), C.double(min))
	}
}

func (ptr *QLogValueAxis) MinorTickCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLogValueAxis_MinorTickCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQLogValueAxis_MinorTickCountChanged
func callbackQLogValueAxis_MinorTickCountChanged(ptr unsafe.Pointer, minorTickCount C.int) {
	if signal := qt.GetSignal(ptr, "minorTickCountChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(minorTickCount)))
	}

}

func (ptr *QLogValueAxis) ConnectMinorTickCountChanged(f func(minorTickCount int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minorTickCountChanged") {
			C.QLogValueAxis_ConnectMinorTickCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minorTickCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minorTickCountChanged"); signal != nil {
			f := func(minorTickCount int) {
				(*(*func(int))(signal))(minorTickCount)
				f(minorTickCount)
			}
			qt.ConnectSignal(ptr.Pointer(), "minorTickCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minorTickCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLogValueAxis) DisconnectMinorTickCountChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectMinorTickCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minorTickCountChanged")
	}
}

func (ptr *QLogValueAxis) MinorTickCountChanged(minorTickCount int) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_MinorTickCountChanged(ptr.Pointer(), C.int(int32(minorTickCount)))
	}
}

//export callbackQLogValueAxis_RangeChanged
func callbackQLogValueAxis_RangeChanged(ptr unsafe.Pointer, min C.double, max C.double) {
	if signal := qt.GetSignal(ptr, "rangeChanged"); signal != nil {
		(*(*func(float64, float64))(signal))(float64(min), float64(max))
	}

}

func (ptr *QLogValueAxis) ConnectRangeChanged(f func(min float64, max float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rangeChanged") {
			C.QLogValueAxis_ConnectRangeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "rangeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rangeChanged"); signal != nil {
			f := func(min float64, max float64) {
				(*(*func(float64, float64))(signal))(min, max)
				f(min, max)
			}
			qt.ConnectSignal(ptr.Pointer(), "rangeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rangeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLogValueAxis) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rangeChanged")
	}
}

func (ptr *QLogValueAxis) RangeChanged(min float64, max float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_RangeChanged(ptr.Pointer(), C.double(min), C.double(max))
	}
}

func (ptr *QLogValueAxis) SetBase(base float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_SetBase(ptr.Pointer(), C.double(base))
	}
}

func (ptr *QLogValueAxis) SetLabelFormat(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QLogValueAxis_SetLabelFormat(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

func (ptr *QLogValueAxis) SetMax(max float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_SetMax(ptr.Pointer(), C.double(max))
	}
}

func (ptr *QLogValueAxis) SetMin(min float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_SetMin(ptr.Pointer(), C.double(min))
	}
}

func (ptr *QLogValueAxis) SetMinorTickCount(minorTickCount int) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_SetMinorTickCount(ptr.Pointer(), C.int(int32(minorTickCount)))
	}
}

func (ptr *QLogValueAxis) SetRange(min float64, max float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_SetRange(ptr.Pointer(), C.double(min), C.double(max))
	}
}

func (ptr *QLogValueAxis) TickCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLogValueAxis_TickCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQLogValueAxis_TickCountChanged
func callbackQLogValueAxis_TickCountChanged(ptr unsafe.Pointer, tickCount C.int) {
	if signal := qt.GetSignal(ptr, "tickCountChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(tickCount)))
	}

}

func (ptr *QLogValueAxis) ConnectTickCountChanged(f func(tickCount int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "tickCountChanged") {
			C.QLogValueAxis_ConnectTickCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "tickCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "tickCountChanged"); signal != nil {
			f := func(tickCount int) {
				(*(*func(int))(signal))(tickCount)
				f(tickCount)
			}
			qt.ConnectSignal(ptr.Pointer(), "tickCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "tickCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLogValueAxis) DisconnectTickCountChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectTickCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "tickCountChanged")
	}
}

func (ptr *QLogValueAxis) TickCountChanged(tickCount int) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_TickCountChanged(ptr.Pointer(), C.int(int32(tickCount)))
	}
}

//export callbackQLogValueAxis_Type
func callbackQLogValueAxis_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractAxis__AxisType)(signal))())
	}

	return C.longlong(NewQLogValueAxisFromPointer(ptr).TypeDefault())
}

func (ptr *QLogValueAxis) ConnectType(f func() QAbstractAxis__AxisType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractAxis__AxisType {
				(*(*func() QAbstractAxis__AxisType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLogValueAxis) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QLogValueAxis) Type() QAbstractAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstractAxis__AxisType(C.QLogValueAxis_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLogValueAxis) TypeDefault() QAbstractAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstractAxis__AxisType(C.QLogValueAxis_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLogValueAxis_DestroyQLogValueAxis
func callbackQLogValueAxis_DestroyQLogValueAxis(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLogValueAxis"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLogValueAxisFromPointer(ptr).DestroyQLogValueAxisDefault()
	}
}

func (ptr *QLogValueAxis) ConnectDestroyQLogValueAxis(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLogValueAxis"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLogValueAxis", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLogValueAxis", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLogValueAxis) DisconnectDestroyQLogValueAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLogValueAxis")
	}
}

func (ptr *QLogValueAxis) DestroyQLogValueAxis() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLogValueAxis_DestroyQLogValueAxis(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLogValueAxis) DestroyQLogValueAxisDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLogValueAxis_DestroyQLogValueAxisDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QPercentBarSeries struct {
	QAbstractBarSeries
}

type QPercentBarSeries_ITF interface {
	QAbstractBarSeries_ITF
	QPercentBarSeries_PTR() *QPercentBarSeries
}

func (ptr *QPercentBarSeries) QPercentBarSeries_PTR() *QPercentBarSeries {
	return ptr
}

func (ptr *QPercentBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractBarSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QPercentBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractBarSeries_PTR().SetPointer(p)
	}
}

func PointerFromQPercentBarSeries(ptr QPercentBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPercentBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQPercentBarSeriesFromPointer(ptr unsafe.Pointer) (n *QPercentBarSeries) {
	n = new(QPercentBarSeries)
	n.SetPointer(ptr)
	return
}
func NewQPercentBarSeries(parent core.QObject_ITF) *QPercentBarSeries {
	tmpValue := NewQPercentBarSeriesFromPointer(C.QPercentBarSeries_NewQPercentBarSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQPercentBarSeries_Type
func callbackQPercentBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQPercentBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QPercentBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPercentBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QPercentBarSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QPercentBarSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPercentBarSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QPercentBarSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQPercentBarSeries_DestroyQPercentBarSeries
func callbackQPercentBarSeries_DestroyQPercentBarSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPercentBarSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPercentBarSeriesFromPointer(ptr).DestroyQPercentBarSeriesDefault()
	}
}

func (ptr *QPercentBarSeries) ConnectDestroyQPercentBarSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPercentBarSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPercentBarSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPercentBarSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPercentBarSeries) DisconnectDestroyQPercentBarSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPercentBarSeries")
	}
}

func (ptr *QPercentBarSeries) DestroyQPercentBarSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPercentBarSeries_DestroyQPercentBarSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPercentBarSeries) DestroyQPercentBarSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPercentBarSeries_DestroyQPercentBarSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QPieLegendMarker struct {
	QLegendMarker
}

type QPieLegendMarker_ITF interface {
	QLegendMarker_ITF
	QPieLegendMarker_PTR() *QPieLegendMarker
}

func (ptr *QPieLegendMarker) QPieLegendMarker_PTR() *QPieLegendMarker {
	return ptr
}

func (ptr *QPieLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLegendMarker_PTR().Pointer()
	}
	return nil
}

func (ptr *QPieLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLegendMarker_PTR().SetPointer(p)
	}
}

func PointerFromQPieLegendMarker(ptr QPieLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPieLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQPieLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QPieLegendMarker) {
	n = new(QPieLegendMarker)
	n.SetPointer(ptr)
	return
}

//export callbackQPieLegendMarker_Series
func callbackQPieLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "series"); signal != nil {
		return PointerFromQPieSeries((*(*func() *QPieSeries)(signal))())
	}

	return PointerFromQPieSeries(NewQPieLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QPieLegendMarker) ConnectSeries(f func() *QPieSeries) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "series"); signal != nil {
			f := func() *QPieSeries {
				(*(*func() *QPieSeries)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "series")
	}
}

func (ptr *QPieLegendMarker) Series() *QPieSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQPieSeriesFromPointer(C.QPieLegendMarker_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieLegendMarker) SeriesDefault() *QPieSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQPieSeriesFromPointer(C.QPieLegendMarker_SeriesDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieLegendMarker) Slice() *QPieSlice {
	if ptr.Pointer() != nil {
		tmpValue := NewQPieSliceFromPointer(C.QPieLegendMarker_Slice(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQPieLegendMarker_Type
func callbackQPieLegendMarker_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QLegendMarker__LegendMarkerType)(signal))())
	}

	return C.longlong(NewQPieLegendMarkerFromPointer(ptr).TypeDefault())
}

func (ptr *QPieLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QLegendMarker__LegendMarkerType {
				(*(*func() QLegendMarker__LegendMarkerType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieLegendMarker) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QPieLegendMarker) Type() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QPieLegendMarker_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QPieLegendMarker_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieLegendMarker_DestroyQPieLegendMarker
func callbackQPieLegendMarker_DestroyQPieLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPieLegendMarker"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPieLegendMarkerFromPointer(ptr).DestroyQPieLegendMarkerDefault()
	}
}

func (ptr *QPieLegendMarker) ConnectDestroyQPieLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPieLegendMarker"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPieLegendMarker", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPieLegendMarker", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieLegendMarker) DisconnectDestroyQPieLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPieLegendMarker")
	}
}

func (ptr *QPieLegendMarker) DestroyQPieLegendMarker() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPieLegendMarker_DestroyQPieLegendMarker(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPieLegendMarker) DestroyQPieLegendMarkerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPieLegendMarker_DestroyQPieLegendMarkerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QPieModelMapper struct {
	core.QObject
}

type QPieModelMapper_ITF interface {
	core.QObject_ITF
	QPieModelMapper_PTR() *QPieModelMapper
}

func (ptr *QPieModelMapper) QPieModelMapper_PTR() *QPieModelMapper {
	return ptr
}

func (ptr *QPieModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QPieModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQPieModelMapper(ptr QPieModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPieModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQPieModelMapperFromPointer(ptr unsafe.Pointer) (n *QPieModelMapper) {
	n = new(QPieModelMapper)
	n.SetPointer(ptr)
	return
}
func (ptr *QPieModelMapper) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPieModelMapper_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPieModelMapper) First() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPieModelMapper_First(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPieModelMapper) LabelsSection() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPieModelMapper_LabelsSection(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPieModelMapper) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QPieModelMapper_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieModelMapper) SetCount(count int) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper_SetCount(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QPieModelMapper) SetFirst(first int) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper_SetFirst(ptr.Pointer(), C.int(int32(first)))
	}
}

func (ptr *QPieModelMapper) SetLabelsSection(labelsSection int) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper_SetLabelsSection(ptr.Pointer(), C.int(int32(labelsSection)))
	}
}

func (ptr *QPieModelMapper) SetOrientation(orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper_SetOrientation(ptr.Pointer(), C.longlong(orientation))
	}
}

func (ptr *QPieModelMapper) SetValuesSection(valuesSection int) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper_SetValuesSection(ptr.Pointer(), C.int(int32(valuesSection)))
	}
}

func (ptr *QPieModelMapper) ValuesSection() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPieModelMapper_ValuesSection(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPieModelMapper) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPieModelMapper___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieModelMapper) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPieModelMapper) __children_newList() unsafe.Pointer {
	return C.QPieModelMapper___children_newList(ptr.Pointer())
}

func (ptr *QPieModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QPieModelMapper___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QPieModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QPieModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QPieModelMapper___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QPieModelMapper) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPieModelMapper___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieModelMapper) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPieModelMapper) __findChildren_newList() unsafe.Pointer {
	return C.QPieModelMapper___findChildren_newList(ptr.Pointer())
}

func (ptr *QPieModelMapper) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPieModelMapper___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieModelMapper) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPieModelMapper) __findChildren_newList3() unsafe.Pointer {
	return C.QPieModelMapper___findChildren_newList3(ptr.Pointer())
}

//export callbackQPieModelMapper_ChildEvent
func callbackQPieModelMapper_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPieModelMapperFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPieModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPieModelMapper_ConnectNotify
func callbackQPieModelMapper_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPieModelMapperFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPieModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPieModelMapper_CustomEvent
func callbackQPieModelMapper_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQPieModelMapperFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPieModelMapper) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPieModelMapper_DeleteLater
func callbackQPieModelMapper_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPieModelMapperFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPieModelMapper) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPieModelMapper_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQPieModelMapper_Destroyed
func callbackQPieModelMapper_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQPieModelMapper_DisconnectNotify
func callbackQPieModelMapper_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPieModelMapperFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPieModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPieModelMapper_Event
func callbackQPieModelMapper_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPieModelMapperFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QPieModelMapper) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieModelMapper_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQPieModelMapper_EventFilter
func callbackQPieModelMapper_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPieModelMapperFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QPieModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieModelMapper_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQPieModelMapper_MetaObject
func callbackQPieModelMapper_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQPieModelMapperFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPieModelMapper) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPieModelMapper_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQPieModelMapper_ObjectNameChanged
func callbackQPieModelMapper_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQPieModelMapper_TimerEvent
func callbackQPieModelMapper_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPieModelMapperFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPieModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPieModelMapper_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QPieSeries struct {
	QAbstractSeries
}

type QPieSeries_ITF interface {
	QAbstractSeries_ITF
	QPieSeries_PTR() *QPieSeries
}

func (ptr *QPieSeries) QPieSeries_PTR() *QPieSeries {
	return ptr
}

func (ptr *QPieSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QPieSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSeries_PTR().SetPointer(p)
	}
}

func PointerFromQPieSeries(ptr QPieSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPieSeries_PTR().Pointer()
	}
	return nil
}

func NewQPieSeriesFromPointer(ptr unsafe.Pointer) (n *QPieSeries) {
	n = new(QPieSeries)
	n.SetPointer(ptr)
	return
}
func NewQPieSeries(parent core.QObject_ITF) *QPieSeries {
	tmpValue := NewQPieSeriesFromPointer(C.QPieSeries_NewQPieSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQPieSeries_Added
func callbackQPieSeries_Added(ptr unsafe.Pointer, slices C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(ptr, "added"); signal != nil {
		(*(*func([]*QPieSlice))(signal))(func(l C.struct_QtCharts_PackedList) []*QPieSlice {
			out := make([]*QPieSlice, int(l.len))
			tmpList := NewQPieSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__added_slices_atList(i)
			}
			return out
		}(slices))
	}

}

func (ptr *QPieSeries) ConnectAdded(f func(slices []*QPieSlice)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "added") {
			C.QPieSeries_ConnectAdded(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "added")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "added"); signal != nil {
			f := func(slices []*QPieSlice) {
				(*(*func([]*QPieSlice))(signal))(slices)
				f(slices)
			}
			qt.ConnectSignal(ptr.Pointer(), "added", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "added", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectAdded() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "added")
	}
}

func (ptr *QPieSeries) Added(slices []*QPieSlice) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Added(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQPieSeriesFromPointer(NewQPieSeriesFromPointer(nil).__added_slices_newList())
			for _, v := range slices {
				tmpList.__added_slices_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QPieSeries) Append(slice QPieSlice_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieSeries_Append(ptr.Pointer(), PointerFromQPieSlice(slice))) != 0
	}
	return false
}

func (ptr *QPieSeries) Append2(slices []*QPieSlice) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieSeries_Append2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQPieSeriesFromPointer(NewQPieSeriesFromPointer(nil).__append_slices_newList2())
			for _, v := range slices {
				tmpList.__append_slices_setList2(v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

func (ptr *QPieSeries) Append3(label string, value float64) *QPieSlice {
	if ptr.Pointer() != nil {
		var labelC *C.char
		if label != "" {
			labelC = C.CString(label)
			defer C.free(unsafe.Pointer(labelC))
		}
		tmpValue := NewQPieSliceFromPointer(C.QPieSeries_Append3(ptr.Pointer(), C.struct_QtCharts_PackedString{data: labelC, len: C.longlong(len(label))}, C.double(value)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieSeries) Clear() {
	if ptr.Pointer() != nil {
		C.QPieSeries_Clear(ptr.Pointer())
	}
}

//export callbackQPieSeries_Clicked
func callbackQPieSeries_Clicked(ptr unsafe.Pointer, slice unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(*QPieSlice))(signal))(NewQPieSliceFromPointer(slice))
	}

}

func (ptr *QPieSeries) ConnectClicked(f func(slice *QPieSlice)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QPieSeries_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(slice *QPieSlice) {
				(*(*func(*QPieSlice))(signal))(slice)
				f(slice)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QPieSeries) Clicked(slice QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Clicked(ptr.Pointer(), PointerFromQPieSlice(slice))
	}
}

func (ptr *QPieSeries) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPieSeries_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQPieSeries_CountChanged
func callbackQPieSeries_CountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "countChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSeries) ConnectCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "countChanged") {
			C.QPieSeries_ConnectCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "countChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "countChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "countChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "countChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectCountChanged() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "countChanged")
	}
}

func (ptr *QPieSeries) CountChanged() {
	if ptr.Pointer() != nil {
		C.QPieSeries_CountChanged(ptr.Pointer())
	}
}

//export callbackQPieSeries_DoubleClicked
func callbackQPieSeries_DoubleClicked(ptr unsafe.Pointer, slice unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		(*(*func(*QPieSlice))(signal))(NewQPieSliceFromPointer(slice))
	}

}

func (ptr *QPieSeries) ConnectDoubleClicked(f func(slice *QPieSlice)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "doubleClicked") {
			C.QPieSeries_ConnectDoubleClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "doubleClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "doubleClicked"); signal != nil {
			f := func(slice *QPieSlice) {
				(*(*func(*QPieSlice))(signal))(slice)
				f(slice)
			}
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "doubleClicked")
	}
}

func (ptr *QPieSeries) DoubleClicked(slice QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries_DoubleClicked(ptr.Pointer(), PointerFromQPieSlice(slice))
	}
}

func (ptr *QPieSeries) HoleSize() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSeries_HoleSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSeries) HorizontalPosition() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSeries_HorizontalPosition(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieSeries_Hovered
func callbackQPieSeries_Hovered(ptr unsafe.Pointer, slice unsafe.Pointer, state C.char) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(*QPieSlice, bool))(signal))(NewQPieSliceFromPointer(slice), int8(state) != 0)
	}

}

func (ptr *QPieSeries) ConnectHovered(f func(slice *QPieSlice, state bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QPieSeries_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(slice *QPieSlice, state bool) {
				(*(*func(*QPieSlice, bool))(signal))(slice, state)
				f(slice, state)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QPieSeries) Hovered(slice QPieSlice_ITF, state bool) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Hovered(ptr.Pointer(), PointerFromQPieSlice(slice), C.char(int8(qt.GoBoolToInt(state))))
	}
}

func (ptr *QPieSeries) Insert(index int, slice QPieSlice_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieSeries_Insert(ptr.Pointer(), C.int(int32(index)), PointerFromQPieSlice(slice))) != 0
	}
	return false
}

func (ptr *QPieSeries) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieSeries_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPieSeries) PieEndAngle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSeries_PieEndAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSeries) PieSize() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSeries_PieSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSeries) PieStartAngle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSeries_PieStartAngle(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieSeries_Pressed
func callbackQPieSeries_Pressed(ptr unsafe.Pointer, slice unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func(*QPieSlice))(signal))(NewQPieSliceFromPointer(slice))
	}

}

func (ptr *QPieSeries) ConnectPressed(f func(slice *QPieSlice)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pressed") {
			C.QPieSeries_ConnectPressed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pressed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pressed"); signal != nil {
			f := func(slice *QPieSlice) {
				(*(*func(*QPieSlice))(signal))(slice)
				f(slice)
			}
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pressed")
	}
}

func (ptr *QPieSeries) Pressed(slice QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Pressed(ptr.Pointer(), PointerFromQPieSlice(slice))
	}
}

//export callbackQPieSeries_Released
func callbackQPieSeries_Released(ptr unsafe.Pointer, slice unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "released"); signal != nil {
		(*(*func(*QPieSlice))(signal))(NewQPieSliceFromPointer(slice))
	}

}

func (ptr *QPieSeries) ConnectReleased(f func(slice *QPieSlice)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "released") {
			C.QPieSeries_ConnectReleased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "released")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "released"); signal != nil {
			f := func(slice *QPieSlice) {
				(*(*func(*QPieSlice))(signal))(slice)
				f(slice)
			}
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "released")
	}
}

func (ptr *QPieSeries) Released(slice QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Released(ptr.Pointer(), PointerFromQPieSlice(slice))
	}
}

func (ptr *QPieSeries) Remove(slice QPieSlice_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieSeries_Remove(ptr.Pointer(), PointerFromQPieSlice(slice))) != 0
	}
	return false
}

//export callbackQPieSeries_Removed
func callbackQPieSeries_Removed(ptr unsafe.Pointer, slices C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(ptr, "removed"); signal != nil {
		(*(*func([]*QPieSlice))(signal))(func(l C.struct_QtCharts_PackedList) []*QPieSlice {
			out := make([]*QPieSlice, int(l.len))
			tmpList := NewQPieSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__removed_slices_atList(i)
			}
			return out
		}(slices))
	}

}

func (ptr *QPieSeries) ConnectRemoved(f func(slices []*QPieSlice)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "removed") {
			C.QPieSeries_ConnectRemoved(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "removed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "removed"); signal != nil {
			f := func(slices []*QPieSlice) {
				(*(*func([]*QPieSlice))(signal))(slices)
				f(slices)
			}
			qt.ConnectSignal(ptr.Pointer(), "removed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectRemoved() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "removed")
	}
}

func (ptr *QPieSeries) Removed(slices []*QPieSlice) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Removed(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQPieSeriesFromPointer(NewQPieSeriesFromPointer(nil).__removed_slices_newList())
			for _, v := range slices {
				tmpList.__removed_slices_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QPieSeries) SetHoleSize(holeSize float64) {
	if ptr.Pointer() != nil {
		C.QPieSeries_SetHoleSize(ptr.Pointer(), C.double(holeSize))
	}
}

func (ptr *QPieSeries) SetHorizontalPosition(relativePosition float64) {
	if ptr.Pointer() != nil {
		C.QPieSeries_SetHorizontalPosition(ptr.Pointer(), C.double(relativePosition))
	}
}

func (ptr *QPieSeries) SetLabelsPosition(position QPieSlice__LabelPosition) {
	if ptr.Pointer() != nil {
		C.QPieSeries_SetLabelsPosition(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QPieSeries) SetLabelsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QPieSeries_SetLabelsVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QPieSeries) SetPieEndAngle(angle float64) {
	if ptr.Pointer() != nil {
		C.QPieSeries_SetPieEndAngle(ptr.Pointer(), C.double(angle))
	}
}

func (ptr *QPieSeries) SetPieSize(relativeSize float64) {
	if ptr.Pointer() != nil {
		C.QPieSeries_SetPieSize(ptr.Pointer(), C.double(relativeSize))
	}
}

func (ptr *QPieSeries) SetPieStartAngle(startAngle float64) {
	if ptr.Pointer() != nil {
		C.QPieSeries_SetPieStartAngle(ptr.Pointer(), C.double(startAngle))
	}
}

func (ptr *QPieSeries) SetVerticalPosition(relativePosition float64) {
	if ptr.Pointer() != nil {
		C.QPieSeries_SetVerticalPosition(ptr.Pointer(), C.double(relativePosition))
	}
}

func (ptr *QPieSeries) Slices() []*QPieSlice {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QPieSlice {
			out := make([]*QPieSlice, int(l.len))
			tmpList := NewQPieSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__slices_atList(i)
			}
			return out
		}(C.QPieSeries_Slices(ptr.Pointer()))
	}
	return make([]*QPieSlice, 0)
}

func (ptr *QPieSeries) Sum() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSeries_Sum(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieSeries_SumChanged
func callbackQPieSeries_SumChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sumChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSeries) ConnectSumChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sumChanged") {
			C.QPieSeries_ConnectSumChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sumChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sumChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sumChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sumChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectSumChanged() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectSumChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sumChanged")
	}
}

func (ptr *QPieSeries) SumChanged() {
	if ptr.Pointer() != nil {
		C.QPieSeries_SumChanged(ptr.Pointer())
	}
}

func (ptr *QPieSeries) Take(slice QPieSlice_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieSeries_Take(ptr.Pointer(), PointerFromQPieSlice(slice))) != 0
	}
	return false
}

//export callbackQPieSeries_Type
func callbackQPieSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQPieSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QPieSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QPieSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QPieSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QPieSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSeries) VerticalPosition() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSeries_VerticalPosition(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieSeries_DestroyQPieSeries
func callbackQPieSeries_DestroyQPieSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPieSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPieSeriesFromPointer(ptr).DestroyQPieSeriesDefault()
	}
}

func (ptr *QPieSeries) ConnectDestroyQPieSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPieSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPieSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPieSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSeries) DisconnectDestroyQPieSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPieSeries")
	}
}

func (ptr *QPieSeries) DestroyQPieSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPieSeries_DestroyQPieSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPieSeries) DestroyQPieSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPieSeries_DestroyQPieSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPieSeries) __added_slices_atList(i int) *QPieSlice {
	if ptr.Pointer() != nil {
		tmpValue := NewQPieSliceFromPointer(C.QPieSeries___added_slices_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieSeries) __added_slices_setList(i QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries___added_slices_setList(ptr.Pointer(), PointerFromQPieSlice(i))
	}
}

func (ptr *QPieSeries) __added_slices_newList() unsafe.Pointer {
	return C.QPieSeries___added_slices_newList(ptr.Pointer())
}

func (ptr *QPieSeries) __append_slices_atList2(i int) *QPieSlice {
	if ptr.Pointer() != nil {
		tmpValue := NewQPieSliceFromPointer(C.QPieSeries___append_slices_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieSeries) __append_slices_setList2(i QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries___append_slices_setList2(ptr.Pointer(), PointerFromQPieSlice(i))
	}
}

func (ptr *QPieSeries) __append_slices_newList2() unsafe.Pointer {
	return C.QPieSeries___append_slices_newList2(ptr.Pointer())
}

func (ptr *QPieSeries) __removed_slices_atList(i int) *QPieSlice {
	if ptr.Pointer() != nil {
		tmpValue := NewQPieSliceFromPointer(C.QPieSeries___removed_slices_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieSeries) __removed_slices_setList(i QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries___removed_slices_setList(ptr.Pointer(), PointerFromQPieSlice(i))
	}
}

func (ptr *QPieSeries) __removed_slices_newList() unsafe.Pointer {
	return C.QPieSeries___removed_slices_newList(ptr.Pointer())
}

func (ptr *QPieSeries) __slices_atList(i int) *QPieSlice {
	if ptr.Pointer() != nil {
		tmpValue := NewQPieSliceFromPointer(C.QPieSeries___slices_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieSeries) __slices_setList(i QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries___slices_setList(ptr.Pointer(), PointerFromQPieSlice(i))
	}
}

func (ptr *QPieSeries) __slices_newList() unsafe.Pointer {
	return C.QPieSeries___slices_newList(ptr.Pointer())
}

type QPieSlice struct {
	core.QObject
}

type QPieSlice_ITF interface {
	core.QObject_ITF
	QPieSlice_PTR() *QPieSlice
}

func (ptr *QPieSlice) QPieSlice_PTR() *QPieSlice {
	return ptr
}

func (ptr *QPieSlice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QPieSlice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQPieSlice(ptr QPieSlice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPieSlice_PTR().Pointer()
	}
	return nil
}

func NewQPieSliceFromPointer(ptr unsafe.Pointer) (n *QPieSlice) {
	n = new(QPieSlice)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPieSlice__LabelPosition
//QPieSlice::LabelPosition
type QPieSlice__LabelPosition int64

const (
	QPieSlice__LabelOutside          QPieSlice__LabelPosition = QPieSlice__LabelPosition(0)
	QPieSlice__LabelInsideHorizontal QPieSlice__LabelPosition = QPieSlice__LabelPosition(1)
	QPieSlice__LabelInsideTangential QPieSlice__LabelPosition = QPieSlice__LabelPosition(2)
	QPieSlice__LabelInsideNormal     QPieSlice__LabelPosition = QPieSlice__LabelPosition(3)
)

func NewQPieSlice(parent core.QObject_ITF) *QPieSlice {
	tmpValue := NewQPieSliceFromPointer(C.QPieSlice_NewQPieSlice(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQPieSlice2(label string, value float64, parent core.QObject_ITF) *QPieSlice {
	var labelC *C.char
	if label != "" {
		labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
	}
	tmpValue := NewQPieSliceFromPointer(C.QPieSlice_NewQPieSlice2(C.struct_QtCharts_PackedString{data: labelC, len: C.longlong(len(label))}, C.double(value), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QPieSlice) AngleSpan() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_AngleSpan(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieSlice_AngleSpanChanged
func callbackQPieSlice_AngleSpanChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "angleSpanChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectAngleSpanChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "angleSpanChanged") {
			C.QPieSlice_ConnectAngleSpanChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "angleSpanChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "angleSpanChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "angleSpanChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "angleSpanChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectAngleSpanChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectAngleSpanChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "angleSpanChanged")
	}
}

func (ptr *QPieSlice) AngleSpanChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_AngleSpanChanged(ptr.Pointer())
	}
}

func (ptr *QPieSlice) BorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QPieSlice_BorderColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQPieSlice_BorderColorChanged
func callbackQPieSlice_BorderColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "borderColorChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectBorderColorChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "borderColorChanged") {
			C.QPieSlice_ConnectBorderColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "borderColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "borderColorChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "borderColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "borderColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "borderColorChanged")
	}
}

func (ptr *QPieSlice) BorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_BorderColorChanged(ptr.Pointer())
	}
}

func (ptr *QPieSlice) BorderWidth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPieSlice_BorderWidth(ptr.Pointer())))
	}
	return 0
}

//export callbackQPieSlice_BorderWidthChanged
func callbackQPieSlice_BorderWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "borderWidthChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectBorderWidthChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "borderWidthChanged") {
			C.QPieSlice_ConnectBorderWidthChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "borderWidthChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "borderWidthChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "borderWidthChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "borderWidthChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectBorderWidthChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectBorderWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "borderWidthChanged")
	}
}

func (ptr *QPieSlice) BorderWidthChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_BorderWidthChanged(ptr.Pointer())
	}
}

func (ptr *QPieSlice) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QPieSlice_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQPieSlice_BrushChanged
func callbackQPieSlice_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "brushChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "brushChanged") {
			C.QPieSlice_ConnectBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "brushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "brushChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "brushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "brushChanged")
	}
}

func (ptr *QPieSlice) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_BrushChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_Clicked
func callbackQPieSlice_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QPieSlice_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QPieSlice) Clicked() {
	if ptr.Pointer() != nil {
		C.QPieSlice_Clicked(ptr.Pointer())
	}
}

func (ptr *QPieSlice) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QPieSlice_Color(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQPieSlice_ColorChanged
func callbackQPieSlice_ColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "colorChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectColorChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "colorChanged") {
			C.QPieSlice_ConnectColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "colorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "colorChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "colorChanged")
	}
}

func (ptr *QPieSlice) ColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_ColorChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_DoubleClicked
func callbackQPieSlice_DoubleClicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectDoubleClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "doubleClicked") {
			C.QPieSlice_ConnectDoubleClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "doubleClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "doubleClicked"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "doubleClicked")
	}
}

func (ptr *QPieSlice) DoubleClicked() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DoubleClicked(ptr.Pointer())
	}
}

func (ptr *QPieSlice) ExplodeDistanceFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_ExplodeDistanceFactor(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieSlice_Hovered
func callbackQPieSlice_Hovered(ptr unsafe.Pointer, state C.char) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(bool))(signal))(int8(state) != 0)
	}

}

func (ptr *QPieSlice) ConnectHovered(f func(state bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QPieSlice_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(state bool) {
				(*(*func(bool))(signal))(state)
				f(state)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QPieSlice) Hovered(state bool) {
	if ptr.Pointer() != nil {
		C.QPieSlice_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(state))))
	}
}

func (ptr *QPieSlice) IsExploded() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieSlice_IsExploded(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPieSlice) IsLabelVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieSlice_IsLabelVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPieSlice) Label() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPieSlice_Label(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPieSlice) LabelArmLengthFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_LabelArmLengthFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSlice) LabelBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QPieSlice_LabelBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQPieSlice_LabelBrushChanged
func callbackQPieSlice_LabelBrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelBrushChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectLabelBrushChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelBrushChanged") {
			C.QPieSlice_ConnectLabelBrushChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelBrushChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelBrushChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelBrushChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelBrushChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectLabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectLabelBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelBrushChanged")
	}
}

func (ptr *QPieSlice) LabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_LabelBrushChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_LabelChanged
func callbackQPieSlice_LabelChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectLabelChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelChanged") {
			C.QPieSlice_ConnectLabelChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectLabelChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectLabelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelChanged")
	}
}

func (ptr *QPieSlice) LabelChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_LabelChanged(ptr.Pointer())
	}
}

func (ptr *QPieSlice) LabelColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QPieSlice_LabelColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQPieSlice_LabelColorChanged
func callbackQPieSlice_LabelColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelColorChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectLabelColorChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelColorChanged") {
			C.QPieSlice_ConnectLabelColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelColorChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectLabelColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectLabelColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelColorChanged")
	}
}

func (ptr *QPieSlice) LabelColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_LabelColorChanged(ptr.Pointer())
	}
}

func (ptr *QPieSlice) LabelFont() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QPieSlice_LabelFont(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

//export callbackQPieSlice_LabelFontChanged
func callbackQPieSlice_LabelFontChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelFontChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectLabelFontChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelFontChanged") {
			C.QPieSlice_ConnectLabelFontChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelFontChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelFontChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelFontChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelFontChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectLabelFontChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectLabelFontChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelFontChanged")
	}
}

func (ptr *QPieSlice) LabelFontChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_LabelFontChanged(ptr.Pointer())
	}
}

func (ptr *QPieSlice) LabelPosition() QPieSlice__LabelPosition {
	if ptr.Pointer() != nil {
		return QPieSlice__LabelPosition(C.QPieSlice_LabelPosition(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieSlice_LabelVisibleChanged
func callbackQPieSlice_LabelVisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelVisibleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectLabelVisibleChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelVisibleChanged") {
			C.QPieSlice_ConnectLabelVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelVisibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelVisibleChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelVisibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelVisibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectLabelVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectLabelVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelVisibleChanged")
	}
}

func (ptr *QPieSlice) LabelVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_LabelVisibleChanged(ptr.Pointer())
	}
}

func (ptr *QPieSlice) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QPieSlice_Pen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQPieSlice_PenChanged
func callbackQPieSlice_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "penChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "penChanged") {
			C.QPieSlice_ConnectPenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "penChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "penChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "penChanged")
	}
}

func (ptr *QPieSlice) PenChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_PenChanged(ptr.Pointer())
	}
}

func (ptr *QPieSlice) Percentage() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_Percentage(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieSlice_PercentageChanged
func callbackQPieSlice_PercentageChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "percentageChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectPercentageChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "percentageChanged") {
			C.QPieSlice_ConnectPercentageChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "percentageChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "percentageChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "percentageChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "percentageChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectPercentageChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectPercentageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "percentageChanged")
	}
}

func (ptr *QPieSlice) PercentageChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_PercentageChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_Pressed
func callbackQPieSlice_Pressed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectPressed(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pressed") {
			C.QPieSlice_ConnectPressed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pressed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pressed"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pressed")
	}
}

func (ptr *QPieSlice) Pressed() {
	if ptr.Pointer() != nil {
		C.QPieSlice_Pressed(ptr.Pointer())
	}
}

//export callbackQPieSlice_Released
func callbackQPieSlice_Released(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "released"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectReleased(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "released") {
			C.QPieSlice_ConnectReleased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "released")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "released"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "released")
	}
}

func (ptr *QPieSlice) Released() {
	if ptr.Pointer() != nil {
		C.QPieSlice_Released(ptr.Pointer())
	}
}

func (ptr *QPieSlice) Series() *QPieSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQPieSeriesFromPointer(C.QPieSlice_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) SetBorderColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetBorderColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QPieSlice) SetBorderWidth(width int) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetBorderWidth(ptr.Pointer(), C.int(int32(width)))
	}
}

func (ptr *QPieSlice) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QPieSlice) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QPieSlice) SetExplodeDistanceFactor(factor float64) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetExplodeDistanceFactor(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QPieSlice) SetExploded(exploded bool) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetExploded(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(exploded))))
	}
}

func (ptr *QPieSlice) SetLabel(label string) {
	if ptr.Pointer() != nil {
		var labelC *C.char
		if label != "" {
			labelC = C.CString(label)
			defer C.free(unsafe.Pointer(labelC))
		}
		C.QPieSlice_SetLabel(ptr.Pointer(), C.struct_QtCharts_PackedString{data: labelC, len: C.longlong(len(label))})
	}
}

func (ptr *QPieSlice) SetLabelArmLengthFactor(factor float64) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetLabelArmLengthFactor(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QPieSlice) SetLabelBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetLabelBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QPieSlice) SetLabelColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetLabelColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QPieSlice) SetLabelFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetLabelFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QPieSlice) SetLabelPosition(position QPieSlice__LabelPosition) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetLabelPosition(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QPieSlice) SetLabelVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetLabelVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QPieSlice) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QPieSlice) SetValue(value float64) {
	if ptr.Pointer() != nil {
		C.QPieSlice_SetValue(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPieSlice) StartAngle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_StartAngle(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieSlice_StartAngleChanged
func callbackQPieSlice_StartAngleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startAngleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectStartAngleChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "startAngleChanged") {
			C.QPieSlice_ConnectStartAngleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "startAngleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "startAngleChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "startAngleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startAngleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectStartAngleChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectStartAngleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "startAngleChanged")
	}
}

func (ptr *QPieSlice) StartAngleChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_StartAngleChanged(ptr.Pointer())
	}
}

func (ptr *QPieSlice) Value() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_Value(ptr.Pointer()))
	}
	return 0
}

//export callbackQPieSlice_ValueChanged
func callbackQPieSlice_ValueChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "valueChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QPieSlice) ConnectValueChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valueChanged") {
			C.QPieSlice_ConnectValueChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "valueChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valueChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valueChanged")
	}
}

func (ptr *QPieSlice) ValueChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_ValueChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_DestroyQPieSlice
func callbackQPieSlice_DestroyQPieSlice(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPieSlice"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPieSliceFromPointer(ptr).DestroyQPieSliceDefault()
	}
}

func (ptr *QPieSlice) ConnectDestroyQPieSlice(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPieSlice"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPieSlice", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPieSlice", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPieSlice) DisconnectDestroyQPieSlice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPieSlice")
	}
}

func (ptr *QPieSlice) DestroyQPieSlice() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPieSlice_DestroyQPieSlice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPieSlice) DestroyQPieSliceDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPieSlice_DestroyQPieSliceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPieSlice) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPieSlice___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPieSlice) __children_newList() unsafe.Pointer {
	return C.QPieSlice___children_newList(ptr.Pointer())
}

func (ptr *QPieSlice) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QPieSlice___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QPieSlice) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QPieSlice___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QPieSlice) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPieSlice___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPieSlice) __findChildren_newList() unsafe.Pointer {
	return C.QPieSlice___findChildren_newList(ptr.Pointer())
}

func (ptr *QPieSlice) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPieSlice___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPieSlice) __findChildren_newList3() unsafe.Pointer {
	return C.QPieSlice___findChildren_newList3(ptr.Pointer())
}

//export callbackQPieSlice_ChildEvent
func callbackQPieSlice_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPieSliceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPieSlice) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPieSlice_ConnectNotify
func callbackQPieSlice_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPieSliceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPieSlice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPieSlice_CustomEvent
func callbackQPieSlice_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQPieSliceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPieSlice) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPieSlice_DeleteLater
func callbackQPieSlice_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPieSliceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPieSlice) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPieSlice_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQPieSlice_Destroyed
func callbackQPieSlice_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQPieSlice_DisconnectNotify
func callbackQPieSlice_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPieSliceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPieSlice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPieSlice_Event
func callbackQPieSlice_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPieSliceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QPieSlice) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieSlice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQPieSlice_EventFilter
func callbackQPieSlice_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPieSliceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QPieSlice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPieSlice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQPieSlice_MetaObject
func callbackQPieSlice_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQPieSliceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPieSlice) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPieSlice_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQPieSlice_ObjectNameChanged
func callbackQPieSlice_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQPieSlice_TimerEvent
func callbackQPieSlice_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPieSliceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPieSlice) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSlice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QPolarChart struct {
	QChart
}

type QPolarChart_ITF interface {
	QChart_ITF
	QPolarChart_PTR() *QPolarChart
}

func (ptr *QPolarChart) QPolarChart_PTR() *QPolarChart {
	return ptr
}

func (ptr *QPolarChart) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QChart_PTR().Pointer()
	}
	return nil
}

func (ptr *QPolarChart) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QChart_PTR().SetPointer(p)
	}
}

func PointerFromQPolarChart(ptr QPolarChart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPolarChart_PTR().Pointer()
	}
	return nil
}

func NewQPolarChartFromPointer(ptr unsafe.Pointer) (n *QPolarChart) {
	n = new(QPolarChart)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPolarChart__PolarOrientation
//QPolarChart::PolarOrientation
type QPolarChart__PolarOrientation int64

const (
	QPolarChart__PolarOrientationRadial  QPolarChart__PolarOrientation = QPolarChart__PolarOrientation(0x1)
	QPolarChart__PolarOrientationAngular QPolarChart__PolarOrientation = QPolarChart__PolarOrientation(0x2)
)

func NewQPolarChart(parent widgets.QGraphicsItem_ITF, wFlags core.Qt__WindowType) *QPolarChart {
	tmpValue := NewQPolarChartFromPointer(C.QPolarChart_NewQPolarChart(widgets.PointerFromQGraphicsItem(parent), C.longlong(wFlags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QPolarChart) AddAxis(axis QAbstractAxis_ITF, polarOrientation QPolarChart__PolarOrientation) {
	if ptr.Pointer() != nil {
		C.QPolarChart_AddAxis(ptr.Pointer(), PointerFromQAbstractAxis(axis), C.longlong(polarOrientation))
	}
}

func (ptr *QPolarChart) Axes(polarOrientation QPolarChart__PolarOrientation, series QAbstractSeries_ITF) []*QAbstractAxis {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QAbstractAxis {
			out := make([]*QAbstractAxis, int(l.len))
			tmpList := NewQPolarChartFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__axes_atList(i)
			}
			return out
		}(C.QPolarChart_Axes(ptr.Pointer(), C.longlong(polarOrientation), PointerFromQAbstractSeries(series)))
	}
	return make([]*QAbstractAxis, 0)
}

func QPolarChart_AxisPolarOrientation(axis QAbstractAxis_ITF) QPolarChart__PolarOrientation {
	return QPolarChart__PolarOrientation(C.QPolarChart_QPolarChart_AxisPolarOrientation(PointerFromQAbstractAxis(axis)))
}

func (ptr *QPolarChart) AxisPolarOrientation(axis QAbstractAxis_ITF) QPolarChart__PolarOrientation {
	return QPolarChart__PolarOrientation(C.QPolarChart_QPolarChart_AxisPolarOrientation(PointerFromQAbstractAxis(axis)))
}

//export callbackQPolarChart_DestroyQPolarChart
func callbackQPolarChart_DestroyQPolarChart(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPolarChart"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPolarChartFromPointer(ptr).DestroyQPolarChartDefault()
	}
}

func (ptr *QPolarChart) ConnectDestroyQPolarChart(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPolarChart"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPolarChart", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPolarChart", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPolarChart) DisconnectDestroyQPolarChart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPolarChart")
	}
}

func (ptr *QPolarChart) DestroyQPolarChart() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPolarChart_DestroyQPolarChart(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPolarChart) DestroyQPolarChartDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPolarChart_DestroyQPolarChartDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScatterSeries struct {
	QXYSeries
}

type QScatterSeries_ITF interface {
	QXYSeries_ITF
	QScatterSeries_PTR() *QScatterSeries
}

func (ptr *QScatterSeries) QScatterSeries_PTR() *QScatterSeries {
	return ptr
}

func (ptr *QScatterSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QXYSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QScatterSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QXYSeries_PTR().SetPointer(p)
	}
}

func PointerFromQScatterSeries(ptr QScatterSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScatterSeries_PTR().Pointer()
	}
	return nil
}

func NewQScatterSeriesFromPointer(ptr unsafe.Pointer) (n *QScatterSeries) {
	n = new(QScatterSeries)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QScatterSeries__MarkerShape
//QScatterSeries::MarkerShape
type QScatterSeries__MarkerShape int64

const (
	QScatterSeries__MarkerShapeCircle    QScatterSeries__MarkerShape = QScatterSeries__MarkerShape(0)
	QScatterSeries__MarkerShapeRectangle QScatterSeries__MarkerShape = QScatterSeries__MarkerShape(1)
)

func NewQScatterSeries(parent core.QObject_ITF) *QScatterSeries {
	tmpValue := NewQScatterSeriesFromPointer(C.QScatterSeries_NewQScatterSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QScatterSeries) BorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QScatterSeries_BorderColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQScatterSeries_BorderColorChanged
func callbackQScatterSeries_BorderColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "borderColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QScatterSeries) ConnectBorderColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "borderColorChanged") {
			C.QScatterSeries_ConnectBorderColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "borderColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "borderColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "borderColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "borderColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScatterSeries) DisconnectBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QScatterSeries_DisconnectBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "borderColorChanged")
	}
}

func (ptr *QScatterSeries) BorderColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_BorderColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QScatterSeries) MarkerShape() QScatterSeries__MarkerShape {
	if ptr.Pointer() != nil {
		return QScatterSeries__MarkerShape(C.QScatterSeries_MarkerShape(ptr.Pointer()))
	}
	return 0
}

//export callbackQScatterSeries_MarkerShapeChanged
func callbackQScatterSeries_MarkerShapeChanged(ptr unsafe.Pointer, shape C.longlong) {
	if signal := qt.GetSignal(ptr, "markerShapeChanged"); signal != nil {
		(*(*func(QScatterSeries__MarkerShape))(signal))(QScatterSeries__MarkerShape(shape))
	}

}

func (ptr *QScatterSeries) ConnectMarkerShapeChanged(f func(shape QScatterSeries__MarkerShape)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "markerShapeChanged") {
			C.QScatterSeries_ConnectMarkerShapeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "markerShapeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "markerShapeChanged"); signal != nil {
			f := func(shape QScatterSeries__MarkerShape) {
				(*(*func(QScatterSeries__MarkerShape))(signal))(shape)
				f(shape)
			}
			qt.ConnectSignal(ptr.Pointer(), "markerShapeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "markerShapeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScatterSeries) DisconnectMarkerShapeChanged() {
	if ptr.Pointer() != nil {
		C.QScatterSeries_DisconnectMarkerShapeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "markerShapeChanged")
	}
}

func (ptr *QScatterSeries) MarkerShapeChanged(shape QScatterSeries__MarkerShape) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_MarkerShapeChanged(ptr.Pointer(), C.longlong(shape))
	}
}

func (ptr *QScatterSeries) MarkerSize() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QScatterSeries_MarkerSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQScatterSeries_MarkerSizeChanged
func callbackQScatterSeries_MarkerSizeChanged(ptr unsafe.Pointer, size C.double) {
	if signal := qt.GetSignal(ptr, "markerSizeChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(size))
	}

}

func (ptr *QScatterSeries) ConnectMarkerSizeChanged(f func(size float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "markerSizeChanged") {
			C.QScatterSeries_ConnectMarkerSizeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "markerSizeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "markerSizeChanged"); signal != nil {
			f := func(size float64) {
				(*(*func(float64))(signal))(size)
				f(size)
			}
			qt.ConnectSignal(ptr.Pointer(), "markerSizeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "markerSizeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScatterSeries) DisconnectMarkerSizeChanged() {
	if ptr.Pointer() != nil {
		C.QScatterSeries_DisconnectMarkerSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "markerSizeChanged")
	}
}

func (ptr *QScatterSeries) MarkerSizeChanged(size float64) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_MarkerSizeChanged(ptr.Pointer(), C.double(size))
	}
}

func (ptr *QScatterSeries) SetBorderColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_SetBorderColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QScatterSeries) SetMarkerShape(shape QScatterSeries__MarkerShape) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_SetMarkerShape(ptr.Pointer(), C.longlong(shape))
	}
}

func (ptr *QScatterSeries) SetMarkerSize(size float64) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_SetMarkerSize(ptr.Pointer(), C.double(size))
	}
}

//export callbackQScatterSeries_Type
func callbackQScatterSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQScatterSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QScatterSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScatterSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QScatterSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QScatterSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScatterSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QScatterSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQScatterSeries_DestroyQScatterSeries
func callbackQScatterSeries_DestroyQScatterSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScatterSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScatterSeriesFromPointer(ptr).DestroyQScatterSeriesDefault()
	}
}

func (ptr *QScatterSeries) ConnectDestroyQScatterSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScatterSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QScatterSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScatterSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScatterSeries) DisconnectDestroyQScatterSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QScatterSeries")
	}
}

func (ptr *QScatterSeries) DestroyQScatterSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScatterSeries_DestroyQScatterSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScatterSeries) DestroyQScatterSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScatterSeries_DestroyQScatterSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSplineSeries struct {
	QLineSeries
}

type QSplineSeries_ITF interface {
	QLineSeries_ITF
	QSplineSeries_PTR() *QSplineSeries
}

func (ptr *QSplineSeries) QSplineSeries_PTR() *QSplineSeries {
	return ptr
}

func (ptr *QSplineSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLineSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QSplineSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLineSeries_PTR().SetPointer(p)
	}
}

func PointerFromQSplineSeries(ptr QSplineSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplineSeries_PTR().Pointer()
	}
	return nil
}

func NewQSplineSeriesFromPointer(ptr unsafe.Pointer) (n *QSplineSeries) {
	n = new(QSplineSeries)
	n.SetPointer(ptr)
	return
}
func NewQSplineSeries(parent core.QObject_ITF) *QSplineSeries {
	tmpValue := NewQSplineSeriesFromPointer(C.QSplineSeries_NewQSplineSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSplineSeries_DestroyQSplineSeries
func callbackQSplineSeries_DestroyQSplineSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSplineSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSplineSeriesFromPointer(ptr).DestroyQSplineSeriesDefault()
	}
}

func (ptr *QSplineSeries) ConnectDestroyQSplineSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSplineSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSplineSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSplineSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSplineSeries) DisconnectDestroyQSplineSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSplineSeries")
	}
}

func (ptr *QSplineSeries) DestroyQSplineSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSplineSeries_DestroyQSplineSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSplineSeries) DestroyQSplineSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSplineSeries_DestroyQSplineSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QStackedBarSeries struct {
	QAbstractBarSeries
}

type QStackedBarSeries_ITF interface {
	QAbstractBarSeries_ITF
	QStackedBarSeries_PTR() *QStackedBarSeries
}

func (ptr *QStackedBarSeries) QStackedBarSeries_PTR() *QStackedBarSeries {
	return ptr
}

func (ptr *QStackedBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractBarSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QStackedBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractBarSeries_PTR().SetPointer(p)
	}
}

func PointerFromQStackedBarSeries(ptr QStackedBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStackedBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQStackedBarSeriesFromPointer(ptr unsafe.Pointer) (n *QStackedBarSeries) {
	n = new(QStackedBarSeries)
	n.SetPointer(ptr)
	return
}
func NewQStackedBarSeries(parent core.QObject_ITF) *QStackedBarSeries {
	tmpValue := NewQStackedBarSeriesFromPointer(C.QStackedBarSeries_NewQStackedBarSeries(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQStackedBarSeries_Type
func callbackQStackedBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQStackedBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QStackedBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractSeries__SeriesType {
				(*(*func() QAbstractSeries__SeriesType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStackedBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QStackedBarSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QStackedBarSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedBarSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QStackedBarSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQStackedBarSeries_DestroyQStackedBarSeries
func callbackQStackedBarSeries_DestroyQStackedBarSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QStackedBarSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQStackedBarSeriesFromPointer(ptr).DestroyQStackedBarSeriesDefault()
	}
}

func (ptr *QStackedBarSeries) ConnectDestroyQStackedBarSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QStackedBarSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QStackedBarSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QStackedBarSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStackedBarSeries) DisconnectDestroyQStackedBarSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QStackedBarSeries")
	}
}

func (ptr *QStackedBarSeries) DestroyQStackedBarSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStackedBarSeries_DestroyQStackedBarSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStackedBarSeries) DestroyQStackedBarSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStackedBarSeries_DestroyQStackedBarSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QVBarModelMapper struct {
	core.QObject
}

type QVBarModelMapper_ITF interface {
	core.QObject_ITF
	QVBarModelMapper_PTR() *QVBarModelMapper
}

func (ptr *QVBarModelMapper) QVBarModelMapper_PTR() *QVBarModelMapper {
	return ptr
}

func (ptr *QVBarModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QVBarModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQVBarModelMapper(ptr QVBarModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBarModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQVBarModelMapperFromPointer(ptr unsafe.Pointer) (n *QVBarModelMapper) {
	n = new(QVBarModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQVBarModelMapper(parent core.QObject_ITF) *QVBarModelMapper {
	tmpValue := NewQVBarModelMapperFromPointer(C.QVBarModelMapper_NewQVBarModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QVBarModelMapper) FirstBarSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBarModelMapper_FirstBarSetColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQVBarModelMapper_FirstBarSetColumnChanged
func callbackQVBarModelMapper_FirstBarSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstBarSetColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBarModelMapper) ConnectFirstBarSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstBarSetColumnChanged") {
			C.QVBarModelMapper_ConnectFirstBarSetColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstBarSetColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstBarSetColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstBarSetColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstBarSetColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBarModelMapper) DisconnectFirstBarSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectFirstBarSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstBarSetColumnChanged")
	}
}

func (ptr *QVBarModelMapper) FirstBarSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_FirstBarSetColumnChanged(ptr.Pointer())
	}
}

func (ptr *QVBarModelMapper) FirstRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBarModelMapper_FirstRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQVBarModelMapper_FirstRowChanged
func callbackQVBarModelMapper_FirstRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBarModelMapper) ConnectFirstRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstRowChanged") {
			C.QVBarModelMapper_ConnectFirstRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBarModelMapper) DisconnectFirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectFirstRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstRowChanged")
	}
}

func (ptr *QVBarModelMapper) FirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_FirstRowChanged(ptr.Pointer())
	}
}

func (ptr *QVBarModelMapper) LastBarSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBarModelMapper_LastBarSetColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQVBarModelMapper_LastBarSetColumnChanged
func callbackQVBarModelMapper_LastBarSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lastBarSetColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBarModelMapper) ConnectLastBarSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lastBarSetColumnChanged") {
			C.QVBarModelMapper_ConnectLastBarSetColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lastBarSetColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lastBarSetColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lastBarSetColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastBarSetColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBarModelMapper) DisconnectLastBarSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectLastBarSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lastBarSetColumnChanged")
	}
}

func (ptr *QVBarModelMapper) LastBarSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_LastBarSetColumnChanged(ptr.Pointer())
	}
}

func (ptr *QVBarModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QVBarModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVBarModelMapper_ModelReplaced
func callbackQVBarModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBarModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "modelReplaced") {
			C.QVBarModelMapper_ConnectModelReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "modelReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "modelReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBarModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "modelReplaced")
	}
}

func (ptr *QVBarModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_ModelReplaced(ptr.Pointer())
	}
}

func (ptr *QVBarModelMapper) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBarModelMapper_RowCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQVBarModelMapper_RowCountChanged
func callbackQVBarModelMapper_RowCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rowCountChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBarModelMapper) ConnectRowCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rowCountChanged") {
			C.QVBarModelMapper_ConnectRowCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "rowCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rowCountChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "rowCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rowCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBarModelMapper) DisconnectRowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectRowCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rowCountChanged")
	}
}

func (ptr *QVBarModelMapper) RowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_RowCountChanged(ptr.Pointer())
	}
}

func (ptr *QVBarModelMapper) Series() *QAbstractBarSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractBarSeriesFromPointer(C.QVBarModelMapper_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVBarModelMapper_SeriesReplaced
func callbackQVBarModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "seriesReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBarModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "seriesReplaced") {
			C.QVBarModelMapper_ConnectSeriesReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "seriesReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "seriesReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBarModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "seriesReplaced")
	}
}

func (ptr *QVBarModelMapper) SeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_SeriesReplaced(ptr.Pointer())
	}
}

func (ptr *QVBarModelMapper) SetFirstBarSetColumn(firstBarSetColumn int) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_SetFirstBarSetColumn(ptr.Pointer(), C.int(int32(firstBarSetColumn)))
	}
}

func (ptr *QVBarModelMapper) SetFirstRow(firstRow int) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_SetFirstRow(ptr.Pointer(), C.int(int32(firstRow)))
	}
}

func (ptr *QVBarModelMapper) SetLastBarSetColumn(lastBarSetColumn int) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_SetLastBarSetColumn(ptr.Pointer(), C.int(int32(lastBarSetColumn)))
	}
}

func (ptr *QVBarModelMapper) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QVBarModelMapper) SetRowCount(rowCount int) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_SetRowCount(ptr.Pointer(), C.int(int32(rowCount)))
	}
}

func (ptr *QVBarModelMapper) SetSeries(series QAbstractBarSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_SetSeries(ptr.Pointer(), PointerFromQAbstractBarSeries(series))
	}
}

func (ptr *QVBarModelMapper) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVBarModelMapper___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBarModelMapper) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVBarModelMapper) __children_newList() unsafe.Pointer {
	return C.QVBarModelMapper___children_newList(ptr.Pointer())
}

func (ptr *QVBarModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QVBarModelMapper___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QVBarModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QVBarModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QVBarModelMapper___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QVBarModelMapper) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVBarModelMapper___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBarModelMapper) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVBarModelMapper) __findChildren_newList() unsafe.Pointer {
	return C.QVBarModelMapper___findChildren_newList(ptr.Pointer())
}

func (ptr *QVBarModelMapper) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVBarModelMapper___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBarModelMapper) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVBarModelMapper) __findChildren_newList3() unsafe.Pointer {
	return C.QVBarModelMapper___findChildren_newList3(ptr.Pointer())
}

//export callbackQVBarModelMapper_ChildEvent
func callbackQVBarModelMapper_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVBarModelMapperFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVBarModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQVBarModelMapper_ConnectNotify
func callbackQVBarModelMapper_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVBarModelMapperFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVBarModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVBarModelMapper_CustomEvent
func callbackQVBarModelMapper_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQVBarModelMapperFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVBarModelMapper) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQVBarModelMapper_DeleteLater
func callbackQVBarModelMapper_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVBarModelMapperFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVBarModelMapper) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVBarModelMapper_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQVBarModelMapper_Destroyed
func callbackQVBarModelMapper_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQVBarModelMapper_DisconnectNotify
func callbackQVBarModelMapper_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVBarModelMapperFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVBarModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVBarModelMapper_Event
func callbackQVBarModelMapper_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVBarModelMapperFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QVBarModelMapper) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVBarModelMapper_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQVBarModelMapper_EventFilter
func callbackQVBarModelMapper_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVBarModelMapperFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QVBarModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVBarModelMapper_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQVBarModelMapper_MetaObject
func callbackQVBarModelMapper_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQVBarModelMapperFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QVBarModelMapper) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVBarModelMapper_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVBarModelMapper_ObjectNameChanged
func callbackQVBarModelMapper_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQVBarModelMapper_TimerEvent
func callbackQVBarModelMapper_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVBarModelMapperFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVBarModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QVBoxPlotModelMapper struct {
	core.QObject
}

type QVBoxPlotModelMapper_ITF interface {
	core.QObject_ITF
	QVBoxPlotModelMapper_PTR() *QVBoxPlotModelMapper
}

func (ptr *QVBoxPlotModelMapper) QVBoxPlotModelMapper_PTR() *QVBoxPlotModelMapper {
	return ptr
}

func (ptr *QVBoxPlotModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QVBoxPlotModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQVBoxPlotModelMapper(ptr QVBoxPlotModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBoxPlotModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQVBoxPlotModelMapperFromPointer(ptr unsafe.Pointer) (n *QVBoxPlotModelMapper) {
	n = new(QVBoxPlotModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQVBoxPlotModelMapper(parent core.QObject_ITF) *QVBoxPlotModelMapper {
	tmpValue := NewQVBoxPlotModelMapperFromPointer(C.QVBoxPlotModelMapper_NewQVBoxPlotModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QVBoxPlotModelMapper) FirstBoxSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxPlotModelMapper_FirstBoxSetColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQVBoxPlotModelMapper_FirstBoxSetColumnChanged
func callbackQVBoxPlotModelMapper_FirstBoxSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstBoxSetColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectFirstBoxSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstBoxSetColumnChanged") {
			C.QVBoxPlotModelMapper_ConnectFirstBoxSetColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstBoxSetColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstBoxSetColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstBoxSetColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstBoxSetColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectFirstBoxSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectFirstBoxSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstBoxSetColumnChanged")
	}
}

func (ptr *QVBoxPlotModelMapper) FirstBoxSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_FirstBoxSetColumnChanged(ptr.Pointer())
	}
}

func (ptr *QVBoxPlotModelMapper) FirstRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxPlotModelMapper_FirstRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQVBoxPlotModelMapper_FirstRowChanged
func callbackQVBoxPlotModelMapper_FirstRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectFirstRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstRowChanged") {
			C.QVBoxPlotModelMapper_ConnectFirstRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectFirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectFirstRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstRowChanged")
	}
}

func (ptr *QVBoxPlotModelMapper) FirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_FirstRowChanged(ptr.Pointer())
	}
}

func (ptr *QVBoxPlotModelMapper) LastBoxSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxPlotModelMapper_LastBoxSetColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQVBoxPlotModelMapper_LastBoxSetColumnChanged
func callbackQVBoxPlotModelMapper_LastBoxSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lastBoxSetColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectLastBoxSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lastBoxSetColumnChanged") {
			C.QVBoxPlotModelMapper_ConnectLastBoxSetColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lastBoxSetColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lastBoxSetColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lastBoxSetColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastBoxSetColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectLastBoxSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectLastBoxSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lastBoxSetColumnChanged")
	}
}

func (ptr *QVBoxPlotModelMapper) LastBoxSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_LastBoxSetColumnChanged(ptr.Pointer())
	}
}

func (ptr *QVBoxPlotModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QVBoxPlotModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVBoxPlotModelMapper_ModelReplaced
func callbackQVBoxPlotModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "modelReplaced") {
			C.QVBoxPlotModelMapper_ConnectModelReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "modelReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "modelReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "modelReplaced")
	}
}

func (ptr *QVBoxPlotModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_ModelReplaced(ptr.Pointer())
	}
}

func (ptr *QVBoxPlotModelMapper) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxPlotModelMapper_RowCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQVBoxPlotModelMapper_RowCountChanged
func callbackQVBoxPlotModelMapper_RowCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rowCountChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectRowCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rowCountChanged") {
			C.QVBoxPlotModelMapper_ConnectRowCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "rowCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rowCountChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "rowCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rowCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectRowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectRowCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rowCountChanged")
	}
}

func (ptr *QVBoxPlotModelMapper) RowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_RowCountChanged(ptr.Pointer())
	}
}

func (ptr *QVBoxPlotModelMapper) Series() *QBoxPlotSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQBoxPlotSeriesFromPointer(C.QVBoxPlotModelMapper_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVBoxPlotModelMapper_SeriesReplaced
func callbackQVBoxPlotModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "seriesReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "seriesReplaced") {
			C.QVBoxPlotModelMapper_ConnectSeriesReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "seriesReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "seriesReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "seriesReplaced")
	}
}

func (ptr *QVBoxPlotModelMapper) SeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_SeriesReplaced(ptr.Pointer())
	}
}

func (ptr *QVBoxPlotModelMapper) SetFirstBoxSetColumn(firstBoxSetColumn int) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_SetFirstBoxSetColumn(ptr.Pointer(), C.int(int32(firstBoxSetColumn)))
	}
}

func (ptr *QVBoxPlotModelMapper) SetFirstRow(firstRow int) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_SetFirstRow(ptr.Pointer(), C.int(int32(firstRow)))
	}
}

func (ptr *QVBoxPlotModelMapper) SetLastBoxSetColumn(lastBoxSetColumn int) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_SetLastBoxSetColumn(ptr.Pointer(), C.int(int32(lastBoxSetColumn)))
	}
}

func (ptr *QVBoxPlotModelMapper) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QVBoxPlotModelMapper) SetRowCount(rowCount int) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_SetRowCount(ptr.Pointer(), C.int(int32(rowCount)))
	}
}

func (ptr *QVBoxPlotModelMapper) SetSeries(series QBoxPlotSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_SetSeries(ptr.Pointer(), PointerFromQBoxPlotSeries(series))
	}
}

func (ptr *QVBoxPlotModelMapper) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVBoxPlotModelMapper___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBoxPlotModelMapper) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVBoxPlotModelMapper) __children_newList() unsafe.Pointer {
	return C.QVBoxPlotModelMapper___children_newList(ptr.Pointer())
}

func (ptr *QVBoxPlotModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QVBoxPlotModelMapper___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QVBoxPlotModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QVBoxPlotModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QVBoxPlotModelMapper___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QVBoxPlotModelMapper) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVBoxPlotModelMapper___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBoxPlotModelMapper) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVBoxPlotModelMapper) __findChildren_newList() unsafe.Pointer {
	return C.QVBoxPlotModelMapper___findChildren_newList(ptr.Pointer())
}

func (ptr *QVBoxPlotModelMapper) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVBoxPlotModelMapper___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBoxPlotModelMapper) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVBoxPlotModelMapper) __findChildren_newList3() unsafe.Pointer {
	return C.QVBoxPlotModelMapper___findChildren_newList3(ptr.Pointer())
}

//export callbackQVBoxPlotModelMapper_ChildEvent
func callbackQVBoxPlotModelMapper_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVBoxPlotModelMapperFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVBoxPlotModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQVBoxPlotModelMapper_ConnectNotify
func callbackQVBoxPlotModelMapper_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVBoxPlotModelMapperFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVBoxPlotModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVBoxPlotModelMapper_CustomEvent
func callbackQVBoxPlotModelMapper_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQVBoxPlotModelMapperFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVBoxPlotModelMapper) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQVBoxPlotModelMapper_DeleteLater
func callbackQVBoxPlotModelMapper_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVBoxPlotModelMapperFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVBoxPlotModelMapper) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVBoxPlotModelMapper_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQVBoxPlotModelMapper_Destroyed
func callbackQVBoxPlotModelMapper_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQVBoxPlotModelMapper_DisconnectNotify
func callbackQVBoxPlotModelMapper_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVBoxPlotModelMapperFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVBoxPlotModelMapper_Event
func callbackQVBoxPlotModelMapper_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVBoxPlotModelMapperFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QVBoxPlotModelMapper) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVBoxPlotModelMapper_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQVBoxPlotModelMapper_EventFilter
func callbackQVBoxPlotModelMapper_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVBoxPlotModelMapperFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QVBoxPlotModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVBoxPlotModelMapper_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQVBoxPlotModelMapper_MetaObject
func callbackQVBoxPlotModelMapper_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQVBoxPlotModelMapperFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QVBoxPlotModelMapper) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVBoxPlotModelMapper_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVBoxPlotModelMapper_ObjectNameChanged
func callbackQVBoxPlotModelMapper_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQVBoxPlotModelMapper_TimerEvent
func callbackQVBoxPlotModelMapper_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVBoxPlotModelMapperFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVBoxPlotModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QVCandlestickModelMapper struct {
	QCandlestickModelMapper
}

type QVCandlestickModelMapper_ITF interface {
	QCandlestickModelMapper_ITF
	QVCandlestickModelMapper_PTR() *QVCandlestickModelMapper
}

func (ptr *QVCandlestickModelMapper) QVCandlestickModelMapper_PTR() *QVCandlestickModelMapper {
	return ptr
}

func (ptr *QVCandlestickModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QCandlestickModelMapper_PTR().Pointer()
	}
	return nil
}

func (ptr *QVCandlestickModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QCandlestickModelMapper_PTR().SetPointer(p)
	}
}

func PointerFromQVCandlestickModelMapper(ptr QVCandlestickModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVCandlestickModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQVCandlestickModelMapperFromPointer(ptr unsafe.Pointer) (n *QVCandlestickModelMapper) {
	n = new(QVCandlestickModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQVCandlestickModelMapper(parent core.QObject_ITF) *QVCandlestickModelMapper {
	tmpValue := NewQVCandlestickModelMapperFromPointer(C.QVCandlestickModelMapper_NewQVCandlestickModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QVCandlestickModelMapper) CloseRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_CloseRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQVCandlestickModelMapper_CloseRowChanged
func callbackQVCandlestickModelMapper_CloseRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectCloseRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "closeRowChanged") {
			C.QVCandlestickModelMapper_ConnectCloseRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "closeRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "closeRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "closeRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "closeRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectCloseRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectCloseRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "closeRowChanged")
	}
}

func (ptr *QVCandlestickModelMapper) CloseRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_CloseRowChanged(ptr.Pointer())
	}
}

func (ptr *QVCandlestickModelMapper) FirstSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_FirstSetColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQVCandlestickModelMapper_FirstSetColumnChanged
func callbackQVCandlestickModelMapper_FirstSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstSetColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectFirstSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstSetColumnChanged") {
			C.QVCandlestickModelMapper_ConnectFirstSetColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstSetColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstSetColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstSetColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstSetColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectFirstSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectFirstSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstSetColumnChanged")
	}
}

func (ptr *QVCandlestickModelMapper) FirstSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_FirstSetColumnChanged(ptr.Pointer())
	}
}

func (ptr *QVCandlestickModelMapper) HighRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_HighRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQVCandlestickModelMapper_HighRowChanged
func callbackQVCandlestickModelMapper_HighRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "highRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectHighRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "highRowChanged") {
			C.QVCandlestickModelMapper_ConnectHighRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "highRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "highRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "highRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "highRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectHighRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectHighRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "highRowChanged")
	}
}

func (ptr *QVCandlestickModelMapper) HighRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_HighRowChanged(ptr.Pointer())
	}
}

func (ptr *QVCandlestickModelMapper) LastSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_LastSetColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQVCandlestickModelMapper_LastSetColumnChanged
func callbackQVCandlestickModelMapper_LastSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lastSetColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectLastSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lastSetColumnChanged") {
			C.QVCandlestickModelMapper_ConnectLastSetColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lastSetColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lastSetColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lastSetColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastSetColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectLastSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectLastSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lastSetColumnChanged")
	}
}

func (ptr *QVCandlestickModelMapper) LastSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_LastSetColumnChanged(ptr.Pointer())
	}
}

func (ptr *QVCandlestickModelMapper) LowRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_LowRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQVCandlestickModelMapper_LowRowChanged
func callbackQVCandlestickModelMapper_LowRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lowRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectLowRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lowRowChanged") {
			C.QVCandlestickModelMapper_ConnectLowRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lowRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lowRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lowRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lowRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectLowRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectLowRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lowRowChanged")
	}
}

func (ptr *QVCandlestickModelMapper) LowRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_LowRowChanged(ptr.Pointer())
	}
}

func (ptr *QVCandlestickModelMapper) OpenRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_OpenRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQVCandlestickModelMapper_OpenRowChanged
func callbackQVCandlestickModelMapper_OpenRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "openRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectOpenRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "openRowChanged") {
			C.QVCandlestickModelMapper_ConnectOpenRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "openRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "openRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "openRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "openRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectOpenRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectOpenRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "openRowChanged")
	}
}

func (ptr *QVCandlestickModelMapper) OpenRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_OpenRowChanged(ptr.Pointer())
	}
}

//export callbackQVCandlestickModelMapper_Orientation
func callbackQVCandlestickModelMapper_Orientation(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "orientation"); signal != nil {
		return C.longlong((*(*func() core.Qt__Orientation)(signal))())
	}

	return C.longlong(NewQVCandlestickModelMapperFromPointer(ptr).OrientationDefault())
}

func (ptr *QVCandlestickModelMapper) ConnectOrientation(f func() core.Qt__Orientation) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "orientation"); signal != nil {
			f := func() core.Qt__Orientation {
				(*(*func() core.Qt__Orientation)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "orientation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "orientation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectOrientation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "orientation")
	}
}

func (ptr *QVCandlestickModelMapper) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QVCandlestickModelMapper_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVCandlestickModelMapper) OrientationDefault() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QVCandlestickModelMapper_OrientationDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVCandlestickModelMapper) SetCloseRow(closeRow int) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_SetCloseRow(ptr.Pointer(), C.int(int32(closeRow)))
	}
}

func (ptr *QVCandlestickModelMapper) SetFirstSetColumn(firstSetColumn int) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_SetFirstSetColumn(ptr.Pointer(), C.int(int32(firstSetColumn)))
	}
}

func (ptr *QVCandlestickModelMapper) SetHighRow(highRow int) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_SetHighRow(ptr.Pointer(), C.int(int32(highRow)))
	}
}

func (ptr *QVCandlestickModelMapper) SetLastSetColumn(lastSetColumn int) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_SetLastSetColumn(ptr.Pointer(), C.int(int32(lastSetColumn)))
	}
}

func (ptr *QVCandlestickModelMapper) SetLowRow(lowRow int) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_SetLowRow(ptr.Pointer(), C.int(int32(lowRow)))
	}
}

func (ptr *QVCandlestickModelMapper) SetOpenRow(openRow int) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_SetOpenRow(ptr.Pointer(), C.int(int32(openRow)))
	}
}

func (ptr *QVCandlestickModelMapper) SetTimestampRow(timestampRow int) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_SetTimestampRow(ptr.Pointer(), C.int(int32(timestampRow)))
	}
}

func (ptr *QVCandlestickModelMapper) TimestampRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_TimestampRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQVCandlestickModelMapper_TimestampRowChanged
func callbackQVCandlestickModelMapper_TimestampRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timestampRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectTimestampRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "timestampRowChanged") {
			C.QVCandlestickModelMapper_ConnectTimestampRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "timestampRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "timestampRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "timestampRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "timestampRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectTimestampRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectTimestampRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "timestampRowChanged")
	}
}

func (ptr *QVCandlestickModelMapper) TimestampRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_TimestampRowChanged(ptr.Pointer())
	}
}

type QVPieModelMapper struct {
	QPieModelMapper
}

type QVPieModelMapper_ITF interface {
	QPieModelMapper_ITF
	QVPieModelMapper_PTR() *QVPieModelMapper
}

func (ptr *QVPieModelMapper) QVPieModelMapper_PTR() *QVPieModelMapper {
	return ptr
}

func (ptr *QVPieModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPieModelMapper_PTR().Pointer()
	}
	return nil
}

func (ptr *QVPieModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPieModelMapper_PTR().SetPointer(p)
	}
}

func PointerFromQVPieModelMapper(ptr QVPieModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVPieModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQVPieModelMapperFromPointer(ptr unsafe.Pointer) (n *QVPieModelMapper) {
	n = new(QVPieModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQVPieModelMapper(parent core.QObject_ITF) *QVPieModelMapper {
	tmpValue := NewQVPieModelMapperFromPointer(C.QVPieModelMapper_NewQVPieModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QVPieModelMapper) FirstRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVPieModelMapper_FirstRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQVPieModelMapper_FirstRowChanged
func callbackQVPieModelMapper_FirstRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVPieModelMapper) ConnectFirstRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstRowChanged") {
			C.QVPieModelMapper_ConnectFirstRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVPieModelMapper) DisconnectFirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectFirstRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstRowChanged")
	}
}

func (ptr *QVPieModelMapper) FirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_FirstRowChanged(ptr.Pointer())
	}
}

func (ptr *QVPieModelMapper) LabelsColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVPieModelMapper_LabelsColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQVPieModelMapper_LabelsColumnChanged
func callbackQVPieModelMapper_LabelsColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "labelsColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVPieModelMapper) ConnectLabelsColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelsColumnChanged") {
			C.QVPieModelMapper_ConnectLabelsColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelsColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelsColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "labelsColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelsColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVPieModelMapper) DisconnectLabelsColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectLabelsColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelsColumnChanged")
	}
}

func (ptr *QVPieModelMapper) LabelsColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_LabelsColumnChanged(ptr.Pointer())
	}
}

func (ptr *QVPieModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QVPieModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVPieModelMapper_ModelReplaced
func callbackQVPieModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVPieModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "modelReplaced") {
			C.QVPieModelMapper_ConnectModelReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "modelReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "modelReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVPieModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "modelReplaced")
	}
}

func (ptr *QVPieModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_ModelReplaced(ptr.Pointer())
	}
}

func (ptr *QVPieModelMapper) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVPieModelMapper_RowCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQVPieModelMapper_RowCountChanged
func callbackQVPieModelMapper_RowCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rowCountChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVPieModelMapper) ConnectRowCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rowCountChanged") {
			C.QVPieModelMapper_ConnectRowCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "rowCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rowCountChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "rowCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rowCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVPieModelMapper) DisconnectRowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectRowCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rowCountChanged")
	}
}

func (ptr *QVPieModelMapper) RowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_RowCountChanged(ptr.Pointer())
	}
}

func (ptr *QVPieModelMapper) Series() *QPieSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQPieSeriesFromPointer(C.QVPieModelMapper_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVPieModelMapper_SeriesReplaced
func callbackQVPieModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "seriesReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVPieModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "seriesReplaced") {
			C.QVPieModelMapper_ConnectSeriesReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "seriesReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "seriesReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVPieModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "seriesReplaced")
	}
}

func (ptr *QVPieModelMapper) SeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_SeriesReplaced(ptr.Pointer())
	}
}

func (ptr *QVPieModelMapper) SetFirstRow(firstRow int) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_SetFirstRow(ptr.Pointer(), C.int(int32(firstRow)))
	}
}

func (ptr *QVPieModelMapper) SetLabelsColumn(labelsColumn int) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_SetLabelsColumn(ptr.Pointer(), C.int(int32(labelsColumn)))
	}
}

func (ptr *QVPieModelMapper) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QVPieModelMapper) SetRowCount(rowCount int) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_SetRowCount(ptr.Pointer(), C.int(int32(rowCount)))
	}
}

func (ptr *QVPieModelMapper) SetSeries(series QPieSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_SetSeries(ptr.Pointer(), PointerFromQPieSeries(series))
	}
}

func (ptr *QVPieModelMapper) SetValuesColumn(valuesColumn int) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_SetValuesColumn(ptr.Pointer(), C.int(int32(valuesColumn)))
	}
}

func (ptr *QVPieModelMapper) ValuesColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVPieModelMapper_ValuesColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQVPieModelMapper_ValuesColumnChanged
func callbackQVPieModelMapper_ValuesColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "valuesColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVPieModelMapper) ConnectValuesColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valuesColumnChanged") {
			C.QVPieModelMapper_ConnectValuesColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "valuesColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valuesColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "valuesColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valuesColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVPieModelMapper) DisconnectValuesColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectValuesColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valuesColumnChanged")
	}
}

func (ptr *QVPieModelMapper) ValuesColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_ValuesColumnChanged(ptr.Pointer())
	}
}

type QVXYModelMapper struct {
	QXYModelMapper
}

type QVXYModelMapper_ITF interface {
	QXYModelMapper_ITF
	QVXYModelMapper_PTR() *QVXYModelMapper
}

func (ptr *QVXYModelMapper) QVXYModelMapper_PTR() *QVXYModelMapper {
	return ptr
}

func (ptr *QVXYModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QXYModelMapper_PTR().Pointer()
	}
	return nil
}

func (ptr *QVXYModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QXYModelMapper_PTR().SetPointer(p)
	}
}

func PointerFromQVXYModelMapper(ptr QVXYModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVXYModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQVXYModelMapperFromPointer(ptr unsafe.Pointer) (n *QVXYModelMapper) {
	n = new(QVXYModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQVXYModelMapper(parent core.QObject_ITF) *QVXYModelMapper {
	tmpValue := NewQVXYModelMapperFromPointer(C.QVXYModelMapper_NewQVXYModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QVXYModelMapper) FirstRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVXYModelMapper_FirstRow(ptr.Pointer())))
	}
	return 0
}

//export callbackQVXYModelMapper_FirstRowChanged
func callbackQVXYModelMapper_FirstRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "firstRowChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVXYModelMapper) ConnectFirstRowChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstRowChanged") {
			C.QVXYModelMapper_ConnectFirstRowChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstRowChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstRowChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstRowChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstRowChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVXYModelMapper) DisconnectFirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectFirstRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstRowChanged")
	}
}

func (ptr *QVXYModelMapper) FirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_FirstRowChanged(ptr.Pointer())
	}
}

func (ptr *QVXYModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QVXYModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVXYModelMapper_ModelReplaced
func callbackQVXYModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVXYModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "modelReplaced") {
			C.QVXYModelMapper_ConnectModelReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "modelReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "modelReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modelReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVXYModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "modelReplaced")
	}
}

func (ptr *QVXYModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_ModelReplaced(ptr.Pointer())
	}
}

func (ptr *QVXYModelMapper) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVXYModelMapper_RowCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQVXYModelMapper_RowCountChanged
func callbackQVXYModelMapper_RowCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rowCountChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVXYModelMapper) ConnectRowCountChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rowCountChanged") {
			C.QVXYModelMapper_ConnectRowCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "rowCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rowCountChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "rowCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rowCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVXYModelMapper) DisconnectRowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectRowCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rowCountChanged")
	}
}

func (ptr *QVXYModelMapper) RowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_RowCountChanged(ptr.Pointer())
	}
}

func (ptr *QVXYModelMapper) Series() *QXYSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQXYSeriesFromPointer(C.QVXYModelMapper_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVXYModelMapper_SeriesReplaced
func callbackQVXYModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "seriesReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVXYModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "seriesReplaced") {
			C.QVXYModelMapper_ConnectSeriesReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "seriesReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "seriesReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seriesReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVXYModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "seriesReplaced")
	}
}

func (ptr *QVXYModelMapper) SeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_SeriesReplaced(ptr.Pointer())
	}
}

func (ptr *QVXYModelMapper) SetFirstRow(firstRow int) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_SetFirstRow(ptr.Pointer(), C.int(int32(firstRow)))
	}
}

func (ptr *QVXYModelMapper) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QVXYModelMapper) SetRowCount(rowCount int) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_SetRowCount(ptr.Pointer(), C.int(int32(rowCount)))
	}
}

func (ptr *QVXYModelMapper) SetSeries(series QXYSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_SetSeries(ptr.Pointer(), PointerFromQXYSeries(series))
	}
}

func (ptr *QVXYModelMapper) SetXColumn(xColumn int) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_SetXColumn(ptr.Pointer(), C.int(int32(xColumn)))
	}
}

func (ptr *QVXYModelMapper) SetYColumn(yColumn int) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_SetYColumn(ptr.Pointer(), C.int(int32(yColumn)))
	}
}

func (ptr *QVXYModelMapper) XColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVXYModelMapper_XColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQVXYModelMapper_XColumnChanged
func callbackQVXYModelMapper_XColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "xColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVXYModelMapper) ConnectXColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "xColumnChanged") {
			C.QVXYModelMapper_ConnectXColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "xColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "xColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "xColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "xColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVXYModelMapper) DisconnectXColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectXColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "xColumnChanged")
	}
}

func (ptr *QVXYModelMapper) XColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_XColumnChanged(ptr.Pointer())
	}
}

func (ptr *QVXYModelMapper) YColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVXYModelMapper_YColumn(ptr.Pointer())))
	}
	return 0
}

//export callbackQVXYModelMapper_YColumnChanged
func callbackQVXYModelMapper_YColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "yColumnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVXYModelMapper) ConnectYColumnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "yColumnChanged") {
			C.QVXYModelMapper_ConnectYColumnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "yColumnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "yColumnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "yColumnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "yColumnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVXYModelMapper) DisconnectYColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectYColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "yColumnChanged")
	}
}

func (ptr *QVXYModelMapper) YColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_YColumnChanged(ptr.Pointer())
	}
}

type QValueAxis struct {
	QAbstractAxis
}

type QValueAxis_ITF interface {
	QAbstractAxis_ITF
	QValueAxis_PTR() *QValueAxis
}

func (ptr *QValueAxis) QValueAxis_PTR() *QValueAxis {
	return ptr
}

func (ptr *QValueAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractAxis_PTR().Pointer()
	}
	return nil
}

func (ptr *QValueAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractAxis_PTR().SetPointer(p)
	}
}

func PointerFromQValueAxis(ptr QValueAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QValueAxis_PTR().Pointer()
	}
	return nil
}

func NewQValueAxisFromPointer(ptr unsafe.Pointer) (n *QValueAxis) {
	n = new(QValueAxis)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QValueAxis__TickType
//QValueAxis::TickType
type QValueAxis__TickType int64

const (
	QValueAxis__TicksDynamic QValueAxis__TickType = QValueAxis__TickType(0)
	QValueAxis__TicksFixed   QValueAxis__TickType = QValueAxis__TickType(1)
)

func NewQValueAxis(parent core.QObject_ITF) *QValueAxis {
	tmpValue := NewQValueAxisFromPointer(C.QValueAxis_NewQValueAxis(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQValueAxis_ApplyNiceNumbers
func callbackQValueAxis_ApplyNiceNumbers(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "applyNiceNumbers"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQValueAxisFromPointer(ptr).ApplyNiceNumbersDefault()
	}
}

func (ptr *QValueAxis) ConnectApplyNiceNumbers(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "applyNiceNumbers"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "applyNiceNumbers", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "applyNiceNumbers", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectApplyNiceNumbers() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "applyNiceNumbers")
	}
}

func (ptr *QValueAxis) ApplyNiceNumbers() {
	if ptr.Pointer() != nil {
		C.QValueAxis_ApplyNiceNumbers(ptr.Pointer())
	}
}

func (ptr *QValueAxis) ApplyNiceNumbersDefault() {
	if ptr.Pointer() != nil {
		C.QValueAxis_ApplyNiceNumbersDefault(ptr.Pointer())
	}
}

func (ptr *QValueAxis) LabelFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QValueAxis_LabelFormat(ptr.Pointer()))
	}
	return ""
}

//export callbackQValueAxis_LabelFormatChanged
func callbackQValueAxis_LabelFormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "labelFormatChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(format))
	}

}

func (ptr *QValueAxis) ConnectLabelFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "labelFormatChanged") {
			C.QValueAxis_ConnectLabelFormatChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "labelFormatChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "labelFormatChanged"); signal != nil {
			f := func(format string) {
				(*(*func(string))(signal))(format)
				f(format)
			}
			qt.ConnectSignal(ptr.Pointer(), "labelFormatChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "labelFormatChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectLabelFormatChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectLabelFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "labelFormatChanged")
	}
}

func (ptr *QValueAxis) LabelFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QValueAxis_LabelFormatChanged(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

func (ptr *QValueAxis) Max() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QValueAxis_Max(ptr.Pointer()))
	}
	return 0
}

//export callbackQValueAxis_MaxChanged
func callbackQValueAxis_MaxChanged(ptr unsafe.Pointer, max C.double) {
	if signal := qt.GetSignal(ptr, "maxChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(max))
	}

}

func (ptr *QValueAxis) ConnectMaxChanged(f func(max float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "maxChanged") {
			C.QValueAxis_ConnectMaxChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "maxChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "maxChanged"); signal != nil {
			f := func(max float64) {
				(*(*func(float64))(signal))(max)
				f(max)
			}
			qt.ConnectSignal(ptr.Pointer(), "maxChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maxChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectMaxChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectMaxChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "maxChanged")
	}
}

func (ptr *QValueAxis) MaxChanged(max float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_MaxChanged(ptr.Pointer(), C.double(max))
	}
}

func (ptr *QValueAxis) Min() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QValueAxis_Min(ptr.Pointer()))
	}
	return 0
}

//export callbackQValueAxis_MinChanged
func callbackQValueAxis_MinChanged(ptr unsafe.Pointer, min C.double) {
	if signal := qt.GetSignal(ptr, "minChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(min))
	}

}

func (ptr *QValueAxis) ConnectMinChanged(f func(min float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minChanged") {
			C.QValueAxis_ConnectMinChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minChanged"); signal != nil {
			f := func(min float64) {
				(*(*func(float64))(signal))(min)
				f(min)
			}
			qt.ConnectSignal(ptr.Pointer(), "minChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectMinChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectMinChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minChanged")
	}
}

func (ptr *QValueAxis) MinChanged(min float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_MinChanged(ptr.Pointer(), C.double(min))
	}
}

func (ptr *QValueAxis) MinorTickCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QValueAxis_MinorTickCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQValueAxis_MinorTickCountChanged
func callbackQValueAxis_MinorTickCountChanged(ptr unsafe.Pointer, minorTickCount C.int) {
	if signal := qt.GetSignal(ptr, "minorTickCountChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(minorTickCount)))
	}

}

func (ptr *QValueAxis) ConnectMinorTickCountChanged(f func(minorTickCount int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minorTickCountChanged") {
			C.QValueAxis_ConnectMinorTickCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minorTickCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minorTickCountChanged"); signal != nil {
			f := func(minorTickCount int) {
				(*(*func(int))(signal))(minorTickCount)
				f(minorTickCount)
			}
			qt.ConnectSignal(ptr.Pointer(), "minorTickCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minorTickCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectMinorTickCountChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectMinorTickCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minorTickCountChanged")
	}
}

func (ptr *QValueAxis) MinorTickCountChanged(minorTickCount int) {
	if ptr.Pointer() != nil {
		C.QValueAxis_MinorTickCountChanged(ptr.Pointer(), C.int(int32(minorTickCount)))
	}
}

//export callbackQValueAxis_RangeChanged
func callbackQValueAxis_RangeChanged(ptr unsafe.Pointer, min C.double, max C.double) {
	if signal := qt.GetSignal(ptr, "rangeChanged"); signal != nil {
		(*(*func(float64, float64))(signal))(float64(min), float64(max))
	}

}

func (ptr *QValueAxis) ConnectRangeChanged(f func(min float64, max float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rangeChanged") {
			C.QValueAxis_ConnectRangeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "rangeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rangeChanged"); signal != nil {
			f := func(min float64, max float64) {
				(*(*func(float64, float64))(signal))(min, max)
				f(min, max)
			}
			qt.ConnectSignal(ptr.Pointer(), "rangeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rangeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rangeChanged")
	}
}

func (ptr *QValueAxis) RangeChanged(min float64, max float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_RangeChanged(ptr.Pointer(), C.double(min), C.double(max))
	}
}

func (ptr *QValueAxis) SetLabelFormat(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QValueAxis_SetLabelFormat(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

func (ptr *QValueAxis) SetMax(max float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_SetMax(ptr.Pointer(), C.double(max))
	}
}

func (ptr *QValueAxis) SetMin(min float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_SetMin(ptr.Pointer(), C.double(min))
	}
}

func (ptr *QValueAxis) SetMinorTickCount(count int) {
	if ptr.Pointer() != nil {
		C.QValueAxis_SetMinorTickCount(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QValueAxis) SetRange(min float64, max float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_SetRange(ptr.Pointer(), C.double(min), C.double(max))
	}
}

func (ptr *QValueAxis) SetTickAnchor(anchor float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_SetTickAnchor(ptr.Pointer(), C.double(anchor))
	}
}

func (ptr *QValueAxis) SetTickCount(count int) {
	if ptr.Pointer() != nil {
		C.QValueAxis_SetTickCount(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QValueAxis) SetTickInterval(insterval float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_SetTickInterval(ptr.Pointer(), C.double(insterval))
	}
}

func (ptr *QValueAxis) SetTickType(ty QValueAxis__TickType) {
	if ptr.Pointer() != nil {
		C.QValueAxis_SetTickType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QValueAxis) TickAnchor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QValueAxis_TickAnchor(ptr.Pointer()))
	}
	return 0
}

//export callbackQValueAxis_TickAnchorChanged
func callbackQValueAxis_TickAnchorChanged(ptr unsafe.Pointer, anchor C.double) {
	if signal := qt.GetSignal(ptr, "tickAnchorChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(anchor))
	}

}

func (ptr *QValueAxis) ConnectTickAnchorChanged(f func(anchor float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "tickAnchorChanged") {
			C.QValueAxis_ConnectTickAnchorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "tickAnchorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "tickAnchorChanged"); signal != nil {
			f := func(anchor float64) {
				(*(*func(float64))(signal))(anchor)
				f(anchor)
			}
			qt.ConnectSignal(ptr.Pointer(), "tickAnchorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "tickAnchorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectTickAnchorChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectTickAnchorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "tickAnchorChanged")
	}
}

func (ptr *QValueAxis) TickAnchorChanged(anchor float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_TickAnchorChanged(ptr.Pointer(), C.double(anchor))
	}
}

func (ptr *QValueAxis) TickCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QValueAxis_TickCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQValueAxis_TickCountChanged
func callbackQValueAxis_TickCountChanged(ptr unsafe.Pointer, tickCount C.int) {
	if signal := qt.GetSignal(ptr, "tickCountChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(tickCount)))
	}

}

func (ptr *QValueAxis) ConnectTickCountChanged(f func(tickCount int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "tickCountChanged") {
			C.QValueAxis_ConnectTickCountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "tickCountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "tickCountChanged"); signal != nil {
			f := func(tickCount int) {
				(*(*func(int))(signal))(tickCount)
				f(tickCount)
			}
			qt.ConnectSignal(ptr.Pointer(), "tickCountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "tickCountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectTickCountChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectTickCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "tickCountChanged")
	}
}

func (ptr *QValueAxis) TickCountChanged(tickCount int) {
	if ptr.Pointer() != nil {
		C.QValueAxis_TickCountChanged(ptr.Pointer(), C.int(int32(tickCount)))
	}
}

func (ptr *QValueAxis) TickInterval() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QValueAxis_TickInterval(ptr.Pointer()))
	}
	return 0
}

//export callbackQValueAxis_TickIntervalChanged
func callbackQValueAxis_TickIntervalChanged(ptr unsafe.Pointer, interval C.double) {
	if signal := qt.GetSignal(ptr, "tickIntervalChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(interval))
	}

}

func (ptr *QValueAxis) ConnectTickIntervalChanged(f func(interval float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "tickIntervalChanged") {
			C.QValueAxis_ConnectTickIntervalChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "tickIntervalChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "tickIntervalChanged"); signal != nil {
			f := func(interval float64) {
				(*(*func(float64))(signal))(interval)
				f(interval)
			}
			qt.ConnectSignal(ptr.Pointer(), "tickIntervalChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "tickIntervalChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectTickIntervalChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectTickIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "tickIntervalChanged")
	}
}

func (ptr *QValueAxis) TickIntervalChanged(interval float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_TickIntervalChanged(ptr.Pointer(), C.double(interval))
	}
}

func (ptr *QValueAxis) TickType() QValueAxis__TickType {
	if ptr.Pointer() != nil {
		return QValueAxis__TickType(C.QValueAxis_TickType(ptr.Pointer()))
	}
	return 0
}

//export callbackQValueAxis_TickTypeChanged
func callbackQValueAxis_TickTypeChanged(ptr unsafe.Pointer, ty C.longlong) {
	if signal := qt.GetSignal(ptr, "tickTypeChanged"); signal != nil {
		(*(*func(QValueAxis__TickType))(signal))(QValueAxis__TickType(ty))
	}

}

func (ptr *QValueAxis) ConnectTickTypeChanged(f func(ty QValueAxis__TickType)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "tickTypeChanged") {
			C.QValueAxis_ConnectTickTypeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "tickTypeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "tickTypeChanged"); signal != nil {
			f := func(ty QValueAxis__TickType) {
				(*(*func(QValueAxis__TickType))(signal))(ty)
				f(ty)
			}
			qt.ConnectSignal(ptr.Pointer(), "tickTypeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "tickTypeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectTickTypeChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectTickTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "tickTypeChanged")
	}
}

func (ptr *QValueAxis) TickTypeChanged(ty QValueAxis__TickType) {
	if ptr.Pointer() != nil {
		C.QValueAxis_TickTypeChanged(ptr.Pointer(), C.longlong(ty))
	}
}

//export callbackQValueAxis_Type
func callbackQValueAxis_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractAxis__AxisType)(signal))())
	}

	return C.longlong(NewQValueAxisFromPointer(ptr).TypeDefault())
}

func (ptr *QValueAxis) ConnectType(f func() QAbstractAxis__AxisType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QAbstractAxis__AxisType {
				(*(*func() QAbstractAxis__AxisType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QValueAxis) Type() QAbstractAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstractAxis__AxisType(C.QValueAxis_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QValueAxis) TypeDefault() QAbstractAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstractAxis__AxisType(C.QValueAxis_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQValueAxis_DestroyQValueAxis
func callbackQValueAxis_DestroyQValueAxis(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QValueAxis"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQValueAxisFromPointer(ptr).DestroyQValueAxisDefault()
	}
}

func (ptr *QValueAxis) ConnectDestroyQValueAxis(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QValueAxis"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QValueAxis", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QValueAxis", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValueAxis) DisconnectDestroyQValueAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QValueAxis")
	}
}

func (ptr *QValueAxis) DestroyQValueAxis() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QValueAxis_DestroyQValueAxis(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QValueAxis) DestroyQValueAxisDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QValueAxis_DestroyQValueAxisDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXYLegendMarker struct {
	QLegendMarker
}

type QXYLegendMarker_ITF interface {
	QLegendMarker_ITF
	QXYLegendMarker_PTR() *QXYLegendMarker
}

func (ptr *QXYLegendMarker) QXYLegendMarker_PTR() *QXYLegendMarker {
	return ptr
}

func (ptr *QXYLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLegendMarker_PTR().Pointer()
	}
	return nil
}

func (ptr *QXYLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLegendMarker_PTR().SetPointer(p)
	}
}

func PointerFromQXYLegendMarker(ptr QXYLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXYLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQXYLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QXYLegendMarker) {
	n = new(QXYLegendMarker)
	n.SetPointer(ptr)
	return
}

//export callbackQXYLegendMarker_Series
func callbackQXYLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "series"); signal != nil {
		return PointerFromQXYSeries((*(*func() *QXYSeries)(signal))())
	}

	return PointerFromQXYSeries(NewQXYLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QXYLegendMarker) ConnectSeries(f func() *QXYSeries) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "series"); signal != nil {
			f := func() *QXYSeries {
				(*(*func() *QXYSeries)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "series", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "series")
	}
}

func (ptr *QXYLegendMarker) Series() *QXYSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQXYSeriesFromPointer(C.QXYLegendMarker_Series(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXYLegendMarker) SeriesDefault() *QXYSeries {
	if ptr.Pointer() != nil {
		tmpValue := NewQXYSeriesFromPointer(C.QXYLegendMarker_SeriesDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQXYLegendMarker_Type
func callbackQXYLegendMarker_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QLegendMarker__LegendMarkerType)(signal))())
	}

	return C.longlong(NewQXYLegendMarkerFromPointer(ptr).TypeDefault())
}

func (ptr *QXYLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QLegendMarker__LegendMarkerType {
				(*(*func() QLegendMarker__LegendMarkerType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYLegendMarker) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QXYLegendMarker) Type() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QXYLegendMarker_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXYLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QXYLegendMarker_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQXYLegendMarker_DestroyQXYLegendMarker
func callbackQXYLegendMarker_DestroyQXYLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXYLegendMarker"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQXYLegendMarkerFromPointer(ptr).DestroyQXYLegendMarkerDefault()
	}
}

func (ptr *QXYLegendMarker) ConnectDestroyQXYLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXYLegendMarker"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QXYLegendMarker", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXYLegendMarker", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYLegendMarker) DisconnectDestroyQXYLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXYLegendMarker")
	}
}

func (ptr *QXYLegendMarker) DestroyQXYLegendMarker() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXYLegendMarker_DestroyQXYLegendMarker(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXYLegendMarker) DestroyQXYLegendMarkerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXYLegendMarker_DestroyQXYLegendMarkerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXYModelMapper struct {
	core.QObject
}

type QXYModelMapper_ITF interface {
	core.QObject_ITF
	QXYModelMapper_PTR() *QXYModelMapper
}

func (ptr *QXYModelMapper) QXYModelMapper_PTR() *QXYModelMapper {
	return ptr
}

func (ptr *QXYModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QXYModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQXYModelMapper(ptr QXYModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXYModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQXYModelMapperFromPointer(ptr unsafe.Pointer) (n *QXYModelMapper) {
	n = new(QXYModelMapper)
	n.SetPointer(ptr)
	return
}
func NewQXYModelMapper(parent core.QObject_ITF) *QXYModelMapper {
	tmpValue := NewQXYModelMapperFromPointer(C.QXYModelMapper_NewQXYModelMapper(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QXYModelMapper) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QXYModelMapper_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXYModelMapper) SetOrientation(orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper_SetOrientation(ptr.Pointer(), C.longlong(orientation))
	}
}

func (ptr *QXYModelMapper) SetXSection(xSection int) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper_SetXSection(ptr.Pointer(), C.int(int32(xSection)))
	}
}

func (ptr *QXYModelMapper) SetYSection(ySection int) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper_SetYSection(ptr.Pointer(), C.int(int32(ySection)))
	}
}

func (ptr *QXYModelMapper) XSection() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXYModelMapper_XSection(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXYModelMapper) YSection() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXYModelMapper_YSection(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXYModelMapper) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QXYModelMapper___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXYModelMapper) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QXYModelMapper) __children_newList() unsafe.Pointer {
	return C.QXYModelMapper___children_newList(ptr.Pointer())
}

func (ptr *QXYModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QXYModelMapper___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QXYModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QXYModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QXYModelMapper___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QXYModelMapper) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QXYModelMapper___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXYModelMapper) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QXYModelMapper) __findChildren_newList() unsafe.Pointer {
	return C.QXYModelMapper___findChildren_newList(ptr.Pointer())
}

func (ptr *QXYModelMapper) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QXYModelMapper___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXYModelMapper) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QXYModelMapper) __findChildren_newList3() unsafe.Pointer {
	return C.QXYModelMapper___findChildren_newList3(ptr.Pointer())
}

//export callbackQXYModelMapper_ChildEvent
func callbackQXYModelMapper_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQXYModelMapperFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QXYModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQXYModelMapper_ConnectNotify
func callbackQXYModelMapper_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQXYModelMapperFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QXYModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQXYModelMapper_CustomEvent
func callbackQXYModelMapper_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQXYModelMapperFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QXYModelMapper) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQXYModelMapper_DeleteLater
func callbackQXYModelMapper_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQXYModelMapperFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QXYModelMapper) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXYModelMapper_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQXYModelMapper_Destroyed
func callbackQXYModelMapper_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQXYModelMapper_DisconnectNotify
func callbackQXYModelMapper_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQXYModelMapperFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QXYModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQXYModelMapper_Event
func callbackQXYModelMapper_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXYModelMapperFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QXYModelMapper) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXYModelMapper_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQXYModelMapper_EventFilter
func callbackQXYModelMapper_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXYModelMapperFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QXYModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXYModelMapper_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQXYModelMapper_MetaObject
func callbackQXYModelMapper_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQXYModelMapperFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QXYModelMapper) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QXYModelMapper_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXYModelMapper_ObjectNameChanged
func callbackQXYModelMapper_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQXYModelMapper_TimerEvent
func callbackQXYModelMapper_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQXYModelMapperFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QXYModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QXYModelMapper_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QXYSeries struct {
	QAbstractSeries
}

type QXYSeries_ITF interface {
	QAbstractSeries_ITF
	QXYSeries_PTR() *QXYSeries
}

func (ptr *QXYSeries) QXYSeries_PTR() *QXYSeries {
	return ptr
}

func (ptr *QXYSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSeries_PTR().Pointer()
	}
	return nil
}

func (ptr *QXYSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSeries_PTR().SetPointer(p)
	}
}

func PointerFromQXYSeries(ptr QXYSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXYSeries_PTR().Pointer()
	}
	return nil
}

func NewQXYSeriesFromPointer(ptr unsafe.Pointer) (n *QXYSeries) {
	n = new(QXYSeries)
	n.SetPointer(ptr)
	return
}
func (ptr *QXYSeries) Append(x float64, y float64) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Append(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QXYSeries) Append2(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Append2(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QXYSeries) Append3(points []*core.QPointF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Append3(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQXYSeriesFromPointer(NewQXYSeriesFromPointer(nil).__append_points_newList3())
			for _, v := range points {
				tmpList.__append_points_setList3(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QXYSeries) At(index int) *core.QPointF {
	if ptr.Pointer() != nil {
		return core.NewQPointFFromPointer(C.QXYSeries_At(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QXYSeries) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQBrushFromPointer(C.QXYSeries_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) Clear() {
	if ptr.Pointer() != nil {
		C.QXYSeries_Clear(ptr.Pointer())
	}
}

//export callbackQXYSeries_Clicked
func callbackQXYSeries_Clicked(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(*core.QPointF))(signal))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QXYSeries) ConnectClicked(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QXYSeries_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(point *core.QPointF) {
				(*(*func(*core.QPointF))(signal))(point)
				f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QXYSeries) Clicked(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Clicked(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQXYSeries_Color
func callbackQXYSeries_Color(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "color"); signal != nil {
		return gui.PointerFromQColor((*(*func() *gui.QColor)(signal))())
	}

	return gui.PointerFromQColor(NewQXYSeriesFromPointer(ptr).ColorDefault())
}

func (ptr *QXYSeries) ConnectColor(f func() *gui.QColor) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "color"); signal != nil {
			f := func() *gui.QColor {
				(*(*func() *gui.QColor)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "color", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "color", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectColor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "color")
	}
}

func (ptr *QXYSeries) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QXYSeries_Color(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) ColorDefault() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QXYSeries_ColorDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQXYSeries_ColorChanged
func callbackQXYSeries_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "colorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(func() *gui.QColor {
			tmpValue := gui.NewQColorFromPointer(color)
			qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
			return tmpValue
		}())
	}

}

func (ptr *QXYSeries) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "colorChanged") {
			C.QXYSeries_ConnectColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "colorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "colorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "colorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "colorChanged")
	}
}

func (ptr *QXYSeries) ColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QXYSeries) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXYSeries_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQXYSeries_DoubleClicked
func callbackQXYSeries_DoubleClicked(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "doubleClicked"); signal != nil {
		(*(*func(*core.QPointF))(signal))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QXYSeries) ConnectDoubleClicked(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "doubleClicked") {
			C.QXYSeries_ConnectDoubleClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "doubleClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "doubleClicked"); signal != nil {
			f := func(point *core.QPointF) {
				(*(*func(*core.QPointF))(signal))(point)
				f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "doubleClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "doubleClicked")
	}
}

func (ptr *QXYSeries) DoubleClicked(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_DoubleClicked(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQXYSeries_Hovered
func callbackQXYSeries_Hovered(ptr unsafe.Pointer, point unsafe.Pointer, state C.char) {
	if signal := qt.GetSignal(ptr, "hovered"); signal != nil {
		(*(*func(*core.QPointF, bool))(signal))(core.NewQPointFFromPointer(point), int8(state) != 0)
	}

}

func (ptr *QXYSeries) ConnectHovered(f func(point *core.QPointF, state bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hovered") {
			C.QXYSeries_ConnectHovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hovered"); signal != nil {
			f := func(point *core.QPointF, state bool) {
				(*(*func(*core.QPointF, bool))(signal))(point, state)
				f(point, state)
			}
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hovered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hovered")
	}
}

func (ptr *QXYSeries) Hovered(point core.QPointF_ITF, state bool) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Hovered(ptr.Pointer(), core.PointerFromQPointF(point), C.char(int8(qt.GoBoolToInt(state))))
	}
}

func (ptr *QXYSeries) Insert(index int, point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Insert(ptr.Pointer(), C.int(int32(index)), core.PointerFromQPointF(point))
	}
}

func (ptr *QXYSeries) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPenFromPointer(C.QXYSeries_Pen(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

//export callbackQXYSeries_PenChanged
func callbackQXYSeries_PenChanged(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "penChanged"); signal != nil {
		(*(*func(*gui.QPen))(signal))(gui.NewQPenFromPointer(pen))
	}

}

func (ptr *QXYSeries) ConnectPenChanged(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "penChanged") {
			C.QXYSeries_ConnectPenChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "penChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "penChanged"); signal != nil {
			f := func(pen *gui.QPen) {
				(*(*func(*gui.QPen))(signal))(pen)
				f(pen)
			}
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "penChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "penChanged")
	}
}

func (ptr *QXYSeries) PenChanged(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PenChanged(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

//export callbackQXYSeries_PointAdded
func callbackQXYSeries_PointAdded(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "pointAdded"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QXYSeries) ConnectPointAdded(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointAdded") {
			C.QXYSeries_ConnectPointAdded(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointAdded")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointAdded"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointAdded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointAdded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPointAdded() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointAdded")
	}
}

func (ptr *QXYSeries) PointAdded(index int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointAdded(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QXYSeries) PointLabelsClipping() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXYSeries_PointLabelsClipping(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXYSeries_PointLabelsClippingChanged
func callbackQXYSeries_PointLabelsClippingChanged(ptr unsafe.Pointer, clipping C.char) {
	if signal := qt.GetSignal(ptr, "pointLabelsClippingChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(clipping) != 0)
	}

}

func (ptr *QXYSeries) ConnectPointLabelsClippingChanged(f func(clipping bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointLabelsClippingChanged") {
			C.QXYSeries_ConnectPointLabelsClippingChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointLabelsClippingChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointLabelsClippingChanged"); signal != nil {
			f := func(clipping bool) {
				(*(*func(bool))(signal))(clipping)
				f(clipping)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsClippingChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsClippingChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPointLabelsClippingChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointLabelsClippingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointLabelsClippingChanged")
	}
}

func (ptr *QXYSeries) PointLabelsClippingChanged(clipping bool) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointLabelsClippingChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(clipping))))
	}
}

func (ptr *QXYSeries) PointLabelsColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QXYSeries_PointLabelsColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQXYSeries_PointLabelsColorChanged
func callbackQXYSeries_PointLabelsColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pointLabelsColorChanged"); signal != nil {
		(*(*func(*gui.QColor))(signal))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QXYSeries) ConnectPointLabelsColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointLabelsColorChanged") {
			C.QXYSeries_ConnectPointLabelsColorChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointLabelsColorChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointLabelsColorChanged"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsColorChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsColorChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPointLabelsColorChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointLabelsColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointLabelsColorChanged")
	}
}

func (ptr *QXYSeries) PointLabelsColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointLabelsColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QXYSeries) PointLabelsFont() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QXYSeries_PointLabelsFont(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

//export callbackQXYSeries_PointLabelsFontChanged
func callbackQXYSeries_PointLabelsFontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pointLabelsFontChanged"); signal != nil {
		(*(*func(*gui.QFont))(signal))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *QXYSeries) ConnectPointLabelsFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointLabelsFontChanged") {
			C.QXYSeries_ConnectPointLabelsFontChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointLabelsFontChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointLabelsFontChanged"); signal != nil {
			f := func(font *gui.QFont) {
				(*(*func(*gui.QFont))(signal))(font)
				f(font)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsFontChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsFontChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPointLabelsFontChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointLabelsFontChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointLabelsFontChanged")
	}
}

func (ptr *QXYSeries) PointLabelsFontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointLabelsFontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QXYSeries) PointLabelsFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXYSeries_PointLabelsFormat(ptr.Pointer()))
	}
	return ""
}

//export callbackQXYSeries_PointLabelsFormatChanged
func callbackQXYSeries_PointLabelsFormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(ptr, "pointLabelsFormatChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(format))
	}

}

func (ptr *QXYSeries) ConnectPointLabelsFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointLabelsFormatChanged") {
			C.QXYSeries_ConnectPointLabelsFormatChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointLabelsFormatChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointLabelsFormatChanged"); signal != nil {
			f := func(format string) {
				(*(*func(string))(signal))(format)
				f(format)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsFormatChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsFormatChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPointLabelsFormatChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointLabelsFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointLabelsFormatChanged")
	}
}

func (ptr *QXYSeries) PointLabelsFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QXYSeries_PointLabelsFormatChanged(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

//export callbackQXYSeries_PointLabelsVisibilityChanged
func callbackQXYSeries_PointLabelsVisibilityChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "pointLabelsVisibilityChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QXYSeries) ConnectPointLabelsVisibilityChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointLabelsVisibilityChanged") {
			C.QXYSeries_ConnectPointLabelsVisibilityChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointLabelsVisibilityChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointLabelsVisibilityChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsVisibilityChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointLabelsVisibilityChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPointLabelsVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointLabelsVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointLabelsVisibilityChanged")
	}
}

func (ptr *QXYSeries) PointLabelsVisibilityChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointLabelsVisibilityChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QXYSeries) PointLabelsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXYSeries_PointLabelsVisible(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXYSeries_PointRemoved
func callbackQXYSeries_PointRemoved(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "pointRemoved"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QXYSeries) ConnectPointRemoved(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointRemoved") {
			C.QXYSeries_ConnectPointRemoved(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointRemoved")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointRemoved"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointRemoved", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointRemoved", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPointRemoved() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointRemoved")
	}
}

func (ptr *QXYSeries) PointRemoved(index int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointRemoved(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQXYSeries_PointReplaced
func callbackQXYSeries_PointReplaced(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "pointReplaced"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QXYSeries) ConnectPointReplaced(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointReplaced") {
			C.QXYSeries_ConnectPointReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointReplaced"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPointReplaced() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointReplaced")
	}
}

func (ptr *QXYSeries) PointReplaced(index int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointReplaced(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QXYSeries) Points() []*core.QPointF {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*core.QPointF {
			out := make([]*core.QPointF, int(l.len))
			tmpList := NewQXYSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__points_atList(i)
			}
			return out
		}(C.QXYSeries_Points(ptr.Pointer()))
	}
	return make([]*core.QPointF, 0)
}

//export callbackQXYSeries_PointsRemoved
func callbackQXYSeries_PointsRemoved(ptr unsafe.Pointer, index C.int, count C.int) {
	if signal := qt.GetSignal(ptr, "pointsRemoved"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(index)), int(int32(count)))
	}

}

func (ptr *QXYSeries) ConnectPointsRemoved(f func(index int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointsRemoved") {
			C.QXYSeries_ConnectPointsRemoved(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointsRemoved")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointsRemoved"); signal != nil {
			f := func(index int, count int) {
				(*(*func(int, int))(signal))(index, count)
				f(index, count)
			}
			qt.ConnectSignal(ptr.Pointer(), "pointsRemoved", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointsRemoved", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPointsRemoved() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointsRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointsRemoved")
	}
}

func (ptr *QXYSeries) PointsRemoved(index int, count int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointsRemoved(ptr.Pointer(), C.int(int32(index)), C.int(int32(count)))
	}
}

//export callbackQXYSeries_PointsReplaced
func callbackQXYSeries_PointsReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pointsReplaced"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QXYSeries) ConnectPointsReplaced(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pointsReplaced") {
			C.QXYSeries_ConnectPointsReplaced(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pointsReplaced")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pointsReplaced"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "pointsReplaced", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pointsReplaced", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPointsReplaced() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointsReplaced(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pointsReplaced")
	}
}

func (ptr *QXYSeries) PointsReplaced() {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointsReplaced(ptr.Pointer())
	}
}

func (ptr *QXYSeries) PointsVector() []*core.QPointF {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*core.QPointF {
			out := make([]*core.QPointF, int(l.len))
			tmpList := NewQXYSeriesFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__pointsVector_atList(i)
			}
			return out
		}(C.QXYSeries_PointsVector(ptr.Pointer()))
	}
	return make([]*core.QPointF, 0)
}

func (ptr *QXYSeries) PointsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXYSeries_PointsVisible(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXYSeries_Pressed
func callbackQXYSeries_Pressed(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pressed"); signal != nil {
		(*(*func(*core.QPointF))(signal))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QXYSeries) ConnectPressed(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pressed") {
			C.QXYSeries_ConnectPressed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pressed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pressed"); signal != nil {
			f := func(point *core.QPointF) {
				(*(*func(*core.QPointF))(signal))(point)
				f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pressed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pressed")
	}
}

func (ptr *QXYSeries) Pressed(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Pressed(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQXYSeries_Released
func callbackQXYSeries_Released(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "released"); signal != nil {
		(*(*func(*core.QPointF))(signal))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QXYSeries) ConnectReleased(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "released") {
			C.QXYSeries_ConnectReleased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "released")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "released"); signal != nil {
			f := func(point *core.QPointF) {
				(*(*func(*core.QPointF))(signal))(point)
				f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "released", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "released")
	}
}

func (ptr *QXYSeries) Released(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Released(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QXYSeries) Remove(x float64, y float64) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Remove(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QXYSeries) Remove2(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Remove2(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QXYSeries) Remove3(index int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Remove3(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QXYSeries) RemovePoints(index int, count int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_RemovePoints(ptr.Pointer(), C.int(int32(index)), C.int(int32(count)))
	}
}

func (ptr *QXYSeries) Replace(oldX float64, oldY float64, newX float64, newY float64) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace(ptr.Pointer(), C.double(oldX), C.double(oldY), C.double(newX), C.double(newY))
	}
}

func (ptr *QXYSeries) Replace2(oldPoint core.QPointF_ITF, newPoint core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace2(ptr.Pointer(), core.PointerFromQPointF(oldPoint), core.PointerFromQPointF(newPoint))
	}
}

func (ptr *QXYSeries) Replace3(index int, newX float64, newY float64) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace3(ptr.Pointer(), C.int(int32(index)), C.double(newX), C.double(newY))
	}
}

func (ptr *QXYSeries) Replace4(index int, newPoint core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace4(ptr.Pointer(), C.int(int32(index)), core.PointerFromQPointF(newPoint))
	}
}

func (ptr *QXYSeries) Replace5(points []*core.QPointF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace5(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQXYSeriesFromPointer(NewQXYSeriesFromPointer(nil).__replace_points_newList5())
			for _, v := range points {
				tmpList.__replace_points_setList5(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QXYSeries) Replace6(points []*core.QPointF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace6(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQXYSeriesFromPointer(NewQXYSeriesFromPointer(nil).__replace_points_newList6())
			for _, v := range points {
				tmpList.__replace_points_setList6(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQXYSeries_SetBrush
func callbackQXYSeries_SetBrush(ptr unsafe.Pointer, brush unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setBrush"); signal != nil {
		(*(*func(*gui.QBrush))(signal))(gui.NewQBrushFromPointer(brush))
	} else {
		NewQXYSeriesFromPointer(ptr).SetBrushDefault(gui.NewQBrushFromPointer(brush))
	}
}

func (ptr *QXYSeries) ConnectSetBrush(f func(brush *gui.QBrush)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setBrush"); signal != nil {
			f := func(brush *gui.QBrush) {
				(*(*func(*gui.QBrush))(signal))(brush)
				f(brush)
			}
			qt.ConnectSignal(ptr.Pointer(), "setBrush", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setBrush", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectSetBrush() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setBrush")
	}
}

func (ptr *QXYSeries) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QXYSeries) SetBrushDefault(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetBrushDefault(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

//export callbackQXYSeries_SetColor
func callbackQXYSeries_SetColor(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setColor"); signal != nil {
		(*(*func(*gui.QColor))(signal))(gui.NewQColorFromPointer(color))
	} else {
		NewQXYSeriesFromPointer(ptr).SetColorDefault(gui.NewQColorFromPointer(color))
	}
}

func (ptr *QXYSeries) ConnectSetColor(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setColor"); signal != nil {
			f := func(color *gui.QColor) {
				(*(*func(*gui.QColor))(signal))(color)
				f(color)
			}
			qt.ConnectSignal(ptr.Pointer(), "setColor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setColor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectSetColor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setColor")
	}
}

func (ptr *QXYSeries) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QXYSeries) SetColorDefault(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetColorDefault(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQXYSeries_SetPen
func callbackQXYSeries_SetPen(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setPen"); signal != nil {
		(*(*func(*gui.QPen))(signal))(gui.NewQPenFromPointer(pen))
	} else {
		NewQXYSeriesFromPointer(ptr).SetPenDefault(gui.NewQPenFromPointer(pen))
	}
}

func (ptr *QXYSeries) ConnectSetPen(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPen"); signal != nil {
			f := func(pen *gui.QPen) {
				(*(*func(*gui.QPen))(signal))(pen)
				f(pen)
			}
			qt.ConnectSignal(ptr.Pointer(), "setPen", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPen", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectSetPen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPen")
	}
}

func (ptr *QXYSeries) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QXYSeries) SetPenDefault(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetPenDefault(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QXYSeries) SetPointLabelsClipping(enabled bool) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetPointLabelsClipping(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QXYSeries) SetPointLabelsColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetPointLabelsColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QXYSeries) SetPointLabelsFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetPointLabelsFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QXYSeries) SetPointLabelsFormat(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QXYSeries_SetPointLabelsFormat(ptr.Pointer(), C.struct_QtCharts_PackedString{data: formatC, len: C.longlong(len(format))})
	}
}

func (ptr *QXYSeries) SetPointLabelsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetPointLabelsVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QXYSeries) SetPointsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QXYSeries_SetPointsVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQXYSeries_DestroyQXYSeries
func callbackQXYSeries_DestroyQXYSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXYSeries"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQXYSeriesFromPointer(ptr).DestroyQXYSeriesDefault()
	}
}

func (ptr *QXYSeries) ConnectDestroyQXYSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXYSeries"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QXYSeries", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXYSeries", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXYSeries) DisconnectDestroyQXYSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXYSeries")
	}
}

func (ptr *QXYSeries) DestroyQXYSeries() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXYSeries_DestroyQXYSeries(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXYSeries) DestroyQXYSeriesDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXYSeries_DestroyQXYSeriesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXYSeries) __append_points_atList3(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QXYSeries___append_points_atList3(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) __append_points_setList3(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries___append_points_setList3(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QXYSeries) __append_points_newList3() unsafe.Pointer {
	return C.QXYSeries___append_points_newList3(ptr.Pointer())
}

func (ptr *QXYSeries) __points_atList(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QXYSeries___points_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) __points_setList(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries___points_setList(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QXYSeries) __points_newList() unsafe.Pointer {
	return C.QXYSeries___points_newList(ptr.Pointer())
}

func (ptr *QXYSeries) __pointsVector_atList(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QXYSeries___pointsVector_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) __pointsVector_setList(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries___pointsVector_setList(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QXYSeries) __pointsVector_newList() unsafe.Pointer {
	return C.QXYSeries___pointsVector_newList(ptr.Pointer())
}

func (ptr *QXYSeries) __replace_points_atList5(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QXYSeries___replace_points_atList5(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) __replace_points_setList5(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries___replace_points_setList5(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QXYSeries) __replace_points_newList5() unsafe.Pointer {
	return C.QXYSeries___replace_points_newList5(ptr.Pointer())
}

func (ptr *QXYSeries) __replace_points_atList6(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QXYSeries___replace_points_atList6(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) __replace_points_setList6(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries___replace_points_setList6(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QXYSeries) __replace_points_newList6() unsafe.Pointer {
	return C.QXYSeries___replace_points_newList6(ptr.Pointer())
}

//export callbackQXYSeries_Type
func callbackQXYSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QAbstractSeries__SeriesType)(signal))())
	}

	return C.longlong(NewQXYSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QXYSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QXYSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXYSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QXYSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

func init() {
	qt.ItfMap["charts.QAbstractAxis_ITF"] = QAbstractAxis{}
	qt.EnumMap["charts.QAbstractAxis__AxisTypeNoAxis"] = int64(QAbstractAxis__AxisTypeNoAxis)
	qt.EnumMap["charts.QAbstractAxis__AxisTypeValue"] = int64(QAbstractAxis__AxisTypeValue)
	qt.EnumMap["charts.QAbstractAxis__AxisTypeBarCategory"] = int64(QAbstractAxis__AxisTypeBarCategory)
	qt.EnumMap["charts.QAbstractAxis__AxisTypeCategory"] = int64(QAbstractAxis__AxisTypeCategory)
	qt.EnumMap["charts.QAbstractAxis__AxisTypeDateTime"] = int64(QAbstractAxis__AxisTypeDateTime)
	qt.EnumMap["charts.QAbstractAxis__AxisTypeLogValue"] = int64(QAbstractAxis__AxisTypeLogValue)
	qt.ItfMap["charts.QAbstractBarSeries_ITF"] = QAbstractBarSeries{}
	qt.EnumMap["charts.QAbstractBarSeries__LabelsCenter"] = int64(QAbstractBarSeries__LabelsCenter)
	qt.EnumMap["charts.QAbstractBarSeries__LabelsInsideEnd"] = int64(QAbstractBarSeries__LabelsInsideEnd)
	qt.EnumMap["charts.QAbstractBarSeries__LabelsInsideBase"] = int64(QAbstractBarSeries__LabelsInsideBase)
	qt.EnumMap["charts.QAbstractBarSeries__LabelsOutsideEnd"] = int64(QAbstractBarSeries__LabelsOutsideEnd)
	qt.ItfMap["charts.QAbstractSeries_ITF"] = QAbstractSeries{}
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeLine"] = int64(QAbstractSeries__SeriesTypeLine)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeArea"] = int64(QAbstractSeries__SeriesTypeArea)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeBar"] = int64(QAbstractSeries__SeriesTypeBar)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeStackedBar"] = int64(QAbstractSeries__SeriesTypeStackedBar)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypePercentBar"] = int64(QAbstractSeries__SeriesTypePercentBar)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypePie"] = int64(QAbstractSeries__SeriesTypePie)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeScatter"] = int64(QAbstractSeries__SeriesTypeScatter)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeSpline"] = int64(QAbstractSeries__SeriesTypeSpline)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeHorizontalBar"] = int64(QAbstractSeries__SeriesTypeHorizontalBar)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeHorizontalStackedBar"] = int64(QAbstractSeries__SeriesTypeHorizontalStackedBar)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeHorizontalPercentBar"] = int64(QAbstractSeries__SeriesTypeHorizontalPercentBar)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeBoxPlot"] = int64(QAbstractSeries__SeriesTypeBoxPlot)
	qt.EnumMap["charts.QAbstractSeries__SeriesTypeCandlestick"] = int64(QAbstractSeries__SeriesTypeCandlestick)
	qt.ItfMap["charts.QAreaLegendMarker_ITF"] = QAreaLegendMarker{}
	qt.ItfMap["charts.QAreaSeries_ITF"] = QAreaSeries{}
	qt.FuncMap["charts.NewQAreaSeries"] = NewQAreaSeries
	qt.FuncMap["charts.NewQAreaSeries2"] = NewQAreaSeries2
	qt.ItfMap["charts.QBarCategoryAxis_ITF"] = QBarCategoryAxis{}
	qt.FuncMap["charts.NewQBarCategoryAxis"] = NewQBarCategoryAxis
	qt.ItfMap["charts.QBarLegendMarker_ITF"] = QBarLegendMarker{}
	qt.ItfMap["charts.QBarSeries_ITF"] = QBarSeries{}
	qt.FuncMap["charts.NewQBarSeries"] = NewQBarSeries
	qt.ItfMap["charts.QBarSet_ITF"] = QBarSet{}
	qt.FuncMap["charts.NewQBarSet"] = NewQBarSet
	qt.ItfMap["charts.QBoxPlotLegendMarker_ITF"] = QBoxPlotLegendMarker{}
	qt.ItfMap["charts.QBoxPlotSeries_ITF"] = QBoxPlotSeries{}
	qt.FuncMap["charts.NewQBoxPlotSeries"] = NewQBoxPlotSeries
	qt.ItfMap["charts.QBoxSet_ITF"] = QBoxSet{}
	qt.FuncMap["charts.NewQBoxSet"] = NewQBoxSet
	qt.FuncMap["charts.NewQBoxSet2"] = NewQBoxSet2
	qt.EnumMap["charts.QBoxSet__LowerExtreme"] = int64(QBoxSet__LowerExtreme)
	qt.EnumMap["charts.QBoxSet__LowerQuartile"] = int64(QBoxSet__LowerQuartile)
	qt.EnumMap["charts.QBoxSet__Median"] = int64(QBoxSet__Median)
	qt.EnumMap["charts.QBoxSet__UpperQuartile"] = int64(QBoxSet__UpperQuartile)
	qt.EnumMap["charts.QBoxSet__UpperExtreme"] = int64(QBoxSet__UpperExtreme)
	qt.ItfMap["charts.QCandlestickLegendMarker_ITF"] = QCandlestickLegendMarker{}
	qt.ItfMap["charts.QCandlestickModelMapper_ITF"] = QCandlestickModelMapper{}
	qt.FuncMap["charts.NewQCandlestickModelMapper"] = NewQCandlestickModelMapper
	qt.ItfMap["charts.QCandlestickSeries_ITF"] = QCandlestickSeries{}
	qt.FuncMap["charts.NewQCandlestickSeries"] = NewQCandlestickSeries
	qt.ItfMap["charts.QCandlestickSet_ITF"] = QCandlestickSet{}
	qt.FuncMap["charts.NewQCandlestickSet"] = NewQCandlestickSet
	qt.FuncMap["charts.NewQCandlestickSet2"] = NewQCandlestickSet2
	qt.ItfMap["charts.QCategoryAxis_ITF"] = QCategoryAxis{}
	qt.FuncMap["charts.NewQCategoryAxis"] = NewQCategoryAxis
	qt.EnumMap["charts.QCategoryAxis__AxisLabelsPositionCenter"] = int64(QCategoryAxis__AxisLabelsPositionCenter)
	qt.EnumMap["charts.QCategoryAxis__AxisLabelsPositionOnValue"] = int64(QCategoryAxis__AxisLabelsPositionOnValue)
	qt.ItfMap["charts.QChart_ITF"] = QChart{}
	qt.FuncMap["charts.NewQChart"] = NewQChart
	qt.EnumMap["charts.QChart__ChartTypeUndefined"] = int64(QChart__ChartTypeUndefined)
	qt.EnumMap["charts.QChart__ChartTypeCartesian"] = int64(QChart__ChartTypeCartesian)
	qt.EnumMap["charts.QChart__ChartTypePolar"] = int64(QChart__ChartTypePolar)
	qt.EnumMap["charts.QChart__ChartThemeLight"] = int64(QChart__ChartThemeLight)
	qt.EnumMap["charts.QChart__ChartThemeBlueCerulean"] = int64(QChart__ChartThemeBlueCerulean)
	qt.EnumMap["charts.QChart__ChartThemeDark"] = int64(QChart__ChartThemeDark)
	qt.EnumMap["charts.QChart__ChartThemeBrownSand"] = int64(QChart__ChartThemeBrownSand)
	qt.EnumMap["charts.QChart__ChartThemeBlueNcs"] = int64(QChart__ChartThemeBlueNcs)
	qt.EnumMap["charts.QChart__ChartThemeHighContrast"] = int64(QChart__ChartThemeHighContrast)
	qt.EnumMap["charts.QChart__ChartThemeBlueIcy"] = int64(QChart__ChartThemeBlueIcy)
	qt.EnumMap["charts.QChart__ChartThemeQt"] = int64(QChart__ChartThemeQt)
	qt.EnumMap["charts.QChart__NoAnimation"] = int64(QChart__NoAnimation)
	qt.EnumMap["charts.QChart__GridAxisAnimations"] = int64(QChart__GridAxisAnimations)
	qt.EnumMap["charts.QChart__SeriesAnimations"] = int64(QChart__SeriesAnimations)
	qt.EnumMap["charts.QChart__AllAnimations"] = int64(QChart__AllAnimations)
	qt.ItfMap["charts.QChartView_ITF"] = QChartView{}
	qt.FuncMap["charts.NewQChartView"] = NewQChartView
	qt.FuncMap["charts.NewQChartView2"] = NewQChartView2
	qt.EnumMap["charts.QChartView__NoRubberBand"] = int64(QChartView__NoRubberBand)
	qt.EnumMap["charts.QChartView__VerticalRubberBand"] = int64(QChartView__VerticalRubberBand)
	qt.EnumMap["charts.QChartView__HorizontalRubberBand"] = int64(QChartView__HorizontalRubberBand)
	qt.EnumMap["charts.QChartView__RectangleRubberBand"] = int64(QChartView__RectangleRubberBand)
	qt.ItfMap["charts.QDateTimeAxis_ITF"] = QDateTimeAxis{}
	qt.FuncMap["charts.NewQDateTimeAxis"] = NewQDateTimeAxis
	qt.ItfMap["charts.QHBarModelMapper_ITF"] = QHBarModelMapper{}
	qt.FuncMap["charts.NewQHBarModelMapper"] = NewQHBarModelMapper
	qt.ItfMap["charts.QHBoxPlotModelMapper_ITF"] = QHBoxPlotModelMapper{}
	qt.FuncMap["charts.NewQHBoxPlotModelMapper"] = NewQHBoxPlotModelMapper
	qt.ItfMap["charts.QHCandlestickModelMapper_ITF"] = QHCandlestickModelMapper{}
	qt.FuncMap["charts.NewQHCandlestickModelMapper"] = NewQHCandlestickModelMapper
	qt.ItfMap["charts.QHPieModelMapper_ITF"] = QHPieModelMapper{}
	qt.FuncMap["charts.NewQHPieModelMapper"] = NewQHPieModelMapper
	qt.ItfMap["charts.QHXYModelMapper_ITF"] = QHXYModelMapper{}
	qt.FuncMap["charts.NewQHXYModelMapper"] = NewQHXYModelMapper
	qt.ItfMap["charts.QHorizontalBarSeries_ITF"] = QHorizontalBarSeries{}
	qt.FuncMap["charts.NewQHorizontalBarSeries"] = NewQHorizontalBarSeries
	qt.ItfMap["charts.QHorizontalPercentBarSeries_ITF"] = QHorizontalPercentBarSeries{}
	qt.FuncMap["charts.NewQHorizontalPercentBarSeries"] = NewQHorizontalPercentBarSeries
	qt.ItfMap["charts.QHorizontalStackedBarSeries_ITF"] = QHorizontalStackedBarSeries{}
	qt.FuncMap["charts.NewQHorizontalStackedBarSeries"] = NewQHorizontalStackedBarSeries
	qt.ItfMap["charts.QLegend_ITF"] = QLegend{}
	qt.EnumMap["charts.QLegend__MarkerShapeDefault"] = int64(QLegend__MarkerShapeDefault)
	qt.EnumMap["charts.QLegend__MarkerShapeRectangle"] = int64(QLegend__MarkerShapeRectangle)
	qt.EnumMap["charts.QLegend__MarkerShapeCircle"] = int64(QLegend__MarkerShapeCircle)
	qt.EnumMap["charts.QLegend__MarkerShapeFromSeries"] = int64(QLegend__MarkerShapeFromSeries)
	qt.ItfMap["charts.QLegendMarker_ITF"] = QLegendMarker{}
	qt.EnumMap["charts.QLegendMarker__LegendMarkerTypeArea"] = int64(QLegendMarker__LegendMarkerTypeArea)
	qt.EnumMap["charts.QLegendMarker__LegendMarkerTypeBar"] = int64(QLegendMarker__LegendMarkerTypeBar)
	qt.EnumMap["charts.QLegendMarker__LegendMarkerTypePie"] = int64(QLegendMarker__LegendMarkerTypePie)
	qt.EnumMap["charts.QLegendMarker__LegendMarkerTypeXY"] = int64(QLegendMarker__LegendMarkerTypeXY)
	qt.EnumMap["charts.QLegendMarker__LegendMarkerTypeBoxPlot"] = int64(QLegendMarker__LegendMarkerTypeBoxPlot)
	qt.EnumMap["charts.QLegendMarker__LegendMarkerTypeCandlestick"] = int64(QLegendMarker__LegendMarkerTypeCandlestick)
	qt.ItfMap["charts.QLineSeries_ITF"] = QLineSeries{}
	qt.FuncMap["charts.NewQLineSeries"] = NewQLineSeries
	qt.ItfMap["charts.QLogValueAxis_ITF"] = QLogValueAxis{}
	qt.FuncMap["charts.NewQLogValueAxis"] = NewQLogValueAxis
	qt.ItfMap["charts.QPercentBarSeries_ITF"] = QPercentBarSeries{}
	qt.FuncMap["charts.NewQPercentBarSeries"] = NewQPercentBarSeries
	qt.ItfMap["charts.QPieLegendMarker_ITF"] = QPieLegendMarker{}
	qt.ItfMap["charts.QPieModelMapper_ITF"] = QPieModelMapper{}
	qt.ItfMap["charts.QPieSeries_ITF"] = QPieSeries{}
	qt.FuncMap["charts.NewQPieSeries"] = NewQPieSeries
	qt.ItfMap["charts.QPieSlice_ITF"] = QPieSlice{}
	qt.FuncMap["charts.NewQPieSlice"] = NewQPieSlice
	qt.FuncMap["charts.NewQPieSlice2"] = NewQPieSlice2
	qt.EnumMap["charts.QPieSlice__LabelOutside"] = int64(QPieSlice__LabelOutside)
	qt.EnumMap["charts.QPieSlice__LabelInsideHorizontal"] = int64(QPieSlice__LabelInsideHorizontal)
	qt.EnumMap["charts.QPieSlice__LabelInsideTangential"] = int64(QPieSlice__LabelInsideTangential)
	qt.EnumMap["charts.QPieSlice__LabelInsideNormal"] = int64(QPieSlice__LabelInsideNormal)
	qt.ItfMap["charts.QPolarChart_ITF"] = QPolarChart{}
	qt.FuncMap["charts.NewQPolarChart"] = NewQPolarChart
	qt.FuncMap["charts.QPolarChart_AxisPolarOrientation"] = QPolarChart_AxisPolarOrientation
	qt.EnumMap["charts.QPolarChart__PolarOrientationRadial"] = int64(QPolarChart__PolarOrientationRadial)
	qt.EnumMap["charts.QPolarChart__PolarOrientationAngular"] = int64(QPolarChart__PolarOrientationAngular)
	qt.ItfMap["charts.QScatterSeries_ITF"] = QScatterSeries{}
	qt.FuncMap["charts.NewQScatterSeries"] = NewQScatterSeries
	qt.EnumMap["charts.QScatterSeries__MarkerShapeCircle"] = int64(QScatterSeries__MarkerShapeCircle)
	qt.EnumMap["charts.QScatterSeries__MarkerShapeRectangle"] = int64(QScatterSeries__MarkerShapeRectangle)
	qt.ItfMap["charts.QSplineSeries_ITF"] = QSplineSeries{}
	qt.FuncMap["charts.NewQSplineSeries"] = NewQSplineSeries
	qt.ItfMap["charts.QStackedBarSeries_ITF"] = QStackedBarSeries{}
	qt.FuncMap["charts.NewQStackedBarSeries"] = NewQStackedBarSeries
	qt.ItfMap["charts.QVBarModelMapper_ITF"] = QVBarModelMapper{}
	qt.FuncMap["charts.NewQVBarModelMapper"] = NewQVBarModelMapper
	qt.ItfMap["charts.QVBoxPlotModelMapper_ITF"] = QVBoxPlotModelMapper{}
	qt.FuncMap["charts.NewQVBoxPlotModelMapper"] = NewQVBoxPlotModelMapper
	qt.ItfMap["charts.QVCandlestickModelMapper_ITF"] = QVCandlestickModelMapper{}
	qt.FuncMap["charts.NewQVCandlestickModelMapper"] = NewQVCandlestickModelMapper
	qt.ItfMap["charts.QVPieModelMapper_ITF"] = QVPieModelMapper{}
	qt.FuncMap["charts.NewQVPieModelMapper"] = NewQVPieModelMapper
	qt.ItfMap["charts.QVXYModelMapper_ITF"] = QVXYModelMapper{}
	qt.FuncMap["charts.NewQVXYModelMapper"] = NewQVXYModelMapper
	qt.ItfMap["charts.QValueAxis_ITF"] = QValueAxis{}
	qt.FuncMap["charts.NewQValueAxis"] = NewQValueAxis
	qt.EnumMap["charts.QValueAxis__TicksDynamic"] = int64(QValueAxis__TicksDynamic)
	qt.EnumMap["charts.QValueAxis__TicksFixed"] = int64(QValueAxis__TicksFixed)
	qt.ItfMap["charts.QXYLegendMarker_ITF"] = QXYLegendMarker{}
	qt.ItfMap["charts.QXYModelMapper_ITF"] = QXYModelMapper{}
	qt.FuncMap["charts.NewQXYModelMapper"] = NewQXYModelMapper
	qt.ItfMap["charts.QXYSeries_ITF"] = QXYSeries{}
}
