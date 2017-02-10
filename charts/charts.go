// +build !minimal

package charts

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "charts.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtCharts_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QAbstractAxis struct {
	ptr unsafe.Pointer
}

type QAbstractAxis_ITF interface {
	QAbstractAxis_PTR() *QAbstractAxis
}

func (ptr *QAbstractAxis) QAbstractAxis_PTR() *QAbstractAxis {
	return ptr
}

func (ptr *QAbstractAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstractAxis(ptr QAbstractAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractAxis_PTR().Pointer()
	}
	return nil
}

func NewQAbstractAxisFromPointer(ptr unsafe.Pointer) *QAbstractAxis {
	var n = new(QAbstractAxis)
	n.SetPointer(ptr)
	return n
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

func (ptr *QAbstractAxis) GridLineColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAbstractAxis_GridLineColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) MinorGridLineColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAbstractAxis_MinorGridLineColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractAxis_ColorChanged
func callbackQAbstractAxis_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstractAxis) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::colorChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::colorChanged")
	}
}

func (ptr *QAbstractAxis) ColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAbstractAxis_GridLineColorChanged
func callbackQAbstractAxis_GridLineColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::gridLineColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstractAxis) ConnectGridLineColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectGridLineColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::gridLineColorChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectGridLineColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectGridLineColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::gridLineColorChanged")
	}
}

func (ptr *QAbstractAxis) GridLineColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_GridLineColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAbstractAxis_GridLinePenChanged
func callbackQAbstractAxis_GridLinePenChanged(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::gridLinePenChanged"); signal != nil {
		signal.(func(*gui.QPen))(gui.NewQPenFromPointer(pen))
	}

}

func (ptr *QAbstractAxis) ConnectGridLinePenChanged(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectGridLinePenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::gridLinePenChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectGridLinePenChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectGridLinePenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::gridLinePenChanged")
	}
}

func (ptr *QAbstractAxis) GridLinePenChanged(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_GridLinePenChanged(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

//export callbackQAbstractAxis_GridVisibleChanged
func callbackQAbstractAxis_GridVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::gridVisibleChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectGridVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectGridVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::gridVisibleChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectGridVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectGridVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::gridVisibleChanged")
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

//export callbackQAbstractAxis_LabelsAngleChanged
func callbackQAbstractAxis_LabelsAngleChanged(ptr unsafe.Pointer, angle C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::labelsAngleChanged"); signal != nil {
		signal.(func(int))(int(int32(angle)))
	}

}

func (ptr *QAbstractAxis) ConnectLabelsAngleChanged(f func(angle int)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectLabelsAngleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::labelsAngleChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsAngleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsAngleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::labelsAngleChanged")
	}
}

func (ptr *QAbstractAxis) LabelsAngleChanged(angle int) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsAngleChanged(ptr.Pointer(), C.int(int32(angle)))
	}
}

//export callbackQAbstractAxis_LabelsBrushChanged
func callbackQAbstractAxis_LabelsBrushChanged(ptr unsafe.Pointer, brush unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::labelsBrushChanged"); signal != nil {
		signal.(func(*gui.QBrush))(gui.NewQBrushFromPointer(brush))
	}

}

func (ptr *QAbstractAxis) ConnectLabelsBrushChanged(f func(brush *gui.QBrush)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectLabelsBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::labelsBrushChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsBrushChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::labelsBrushChanged")
	}
}

func (ptr *QAbstractAxis) LabelsBrushChanged(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsBrushChanged(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

//export callbackQAbstractAxis_LabelsColorChanged
func callbackQAbstractAxis_LabelsColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::labelsColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstractAxis) ConnectLabelsColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectLabelsColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::labelsColorChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::labelsColorChanged")
	}
}

func (ptr *QAbstractAxis) LabelsColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAbstractAxis_LabelsFontChanged
func callbackQAbstractAxis_LabelsFontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::labelsFontChanged"); signal != nil {
		signal.(func(*gui.QFont))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *QAbstractAxis) ConnectLabelsFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectLabelsFontChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::labelsFontChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsFontChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsFontChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::labelsFontChanged")
	}
}

func (ptr *QAbstractAxis) LabelsFontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsFontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

//export callbackQAbstractAxis_LabelsVisibleChanged
func callbackQAbstractAxis_LabelsVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::labelsVisibleChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectLabelsVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectLabelsVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::labelsVisibleChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectLabelsVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLabelsVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::labelsVisibleChanged")
	}
}

func (ptr *QAbstractAxis) LabelsVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LabelsVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstractAxis_LinePenChanged
func callbackQAbstractAxis_LinePenChanged(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::linePenChanged"); signal != nil {
		signal.(func(*gui.QPen))(gui.NewQPenFromPointer(pen))
	}

}

func (ptr *QAbstractAxis) ConnectLinePenChanged(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectLinePenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::linePenChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectLinePenChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLinePenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::linePenChanged")
	}
}

func (ptr *QAbstractAxis) LinePenChanged(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LinePenChanged(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

//export callbackQAbstractAxis_LineVisibleChanged
func callbackQAbstractAxis_LineVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::lineVisibleChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectLineVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectLineVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::lineVisibleChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectLineVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectLineVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::lineVisibleChanged")
	}
}

func (ptr *QAbstractAxis) LineVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_LineVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstractAxis_MinorGridLineColorChanged
func callbackQAbstractAxis_MinorGridLineColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::minorGridLineColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstractAxis) ConnectMinorGridLineColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectMinorGridLineColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::minorGridLineColorChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectMinorGridLineColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectMinorGridLineColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::minorGridLineColorChanged")
	}
}

func (ptr *QAbstractAxis) MinorGridLineColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_MinorGridLineColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAbstractAxis_MinorGridLinePenChanged
func callbackQAbstractAxis_MinorGridLinePenChanged(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::minorGridLinePenChanged"); signal != nil {
		signal.(func(*gui.QPen))(gui.NewQPenFromPointer(pen))
	}

}

func (ptr *QAbstractAxis) ConnectMinorGridLinePenChanged(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectMinorGridLinePenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::minorGridLinePenChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectMinorGridLinePenChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectMinorGridLinePenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::minorGridLinePenChanged")
	}
}

func (ptr *QAbstractAxis) MinorGridLinePenChanged(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_MinorGridLinePenChanged(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

//export callbackQAbstractAxis_MinorGridVisibleChanged
func callbackQAbstractAxis_MinorGridVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::minorGridVisibleChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectMinorGridVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectMinorGridVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::minorGridVisibleChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectMinorGridVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectMinorGridVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::minorGridVisibleChanged")
	}
}

func (ptr *QAbstractAxis) MinorGridVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_MinorGridVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstractAxis_ReverseChanged
func callbackQAbstractAxis_ReverseChanged(ptr unsafe.Pointer, reverse C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::reverseChanged"); signal != nil {
		signal.(func(bool))(int8(reverse) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectReverseChanged(f func(reverse bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectReverseChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::reverseChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectReverseChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectReverseChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::reverseChanged")
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
		var titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
		C.QAbstractAxis_SetTitleText(ptr.Pointer(), titleC)
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

//export callbackQAbstractAxis_ShadesBorderColorChanged
func callbackQAbstractAxis_ShadesBorderColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::shadesBorderColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstractAxis) ConnectShadesBorderColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectShadesBorderColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::shadesBorderColorChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectShadesBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectShadesBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::shadesBorderColorChanged")
	}
}

func (ptr *QAbstractAxis) ShadesBorderColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ShadesBorderColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAbstractAxis_ShadesBrushChanged
func callbackQAbstractAxis_ShadesBrushChanged(ptr unsafe.Pointer, brush unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::shadesBrushChanged"); signal != nil {
		signal.(func(*gui.QBrush))(gui.NewQBrushFromPointer(brush))
	}

}

func (ptr *QAbstractAxis) ConnectShadesBrushChanged(f func(brush *gui.QBrush)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectShadesBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::shadesBrushChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectShadesBrushChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectShadesBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::shadesBrushChanged")
	}
}

func (ptr *QAbstractAxis) ShadesBrushChanged(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ShadesBrushChanged(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

//export callbackQAbstractAxis_ShadesColorChanged
func callbackQAbstractAxis_ShadesColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::shadesColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstractAxis) ConnectShadesColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectShadesColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::shadesColorChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectShadesColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectShadesColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::shadesColorChanged")
	}
}

func (ptr *QAbstractAxis) ShadesColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ShadesColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAbstractAxis_ShadesPenChanged
func callbackQAbstractAxis_ShadesPenChanged(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::shadesPenChanged"); signal != nil {
		signal.(func(*gui.QPen))(gui.NewQPenFromPointer(pen))
	}

}

func (ptr *QAbstractAxis) ConnectShadesPenChanged(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectShadesPenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::shadesPenChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectShadesPenChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectShadesPenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::shadesPenChanged")
	}
}

func (ptr *QAbstractAxis) ShadesPenChanged(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ShadesPenChanged(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

//export callbackQAbstractAxis_ShadesVisibleChanged
func callbackQAbstractAxis_ShadesVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::shadesVisibleChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectShadesVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectShadesVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::shadesVisibleChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectShadesVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectShadesVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::shadesVisibleChanged")
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

//export callbackQAbstractAxis_TitleBrushChanged
func callbackQAbstractAxis_TitleBrushChanged(ptr unsafe.Pointer, brush unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::titleBrushChanged"); signal != nil {
		signal.(func(*gui.QBrush))(gui.NewQBrushFromPointer(brush))
	}

}

func (ptr *QAbstractAxis) ConnectTitleBrushChanged(f func(brush *gui.QBrush)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectTitleBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::titleBrushChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectTitleBrushChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectTitleBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::titleBrushChanged")
	}
}

func (ptr *QAbstractAxis) TitleBrushChanged(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_TitleBrushChanged(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

//export callbackQAbstractAxis_TitleFontChanged
func callbackQAbstractAxis_TitleFontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::titleFontChanged"); signal != nil {
		signal.(func(*gui.QFont))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *QAbstractAxis) ConnectTitleFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectTitleFontChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::titleFontChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectTitleFontChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectTitleFontChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::titleFontChanged")
	}
}

func (ptr *QAbstractAxis) TitleFontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_TitleFontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

//export callbackQAbstractAxis_TitleTextChanged
func callbackQAbstractAxis_TitleTextChanged(ptr unsafe.Pointer, text C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::titleTextChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(text))
	}

}

func (ptr *QAbstractAxis) ConnectTitleTextChanged(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectTitleTextChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::titleTextChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectTitleTextChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectTitleTextChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::titleTextChanged")
	}
}

func (ptr *QAbstractAxis) TitleTextChanged(text string) {
	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		C.QAbstractAxis_TitleTextChanged(ptr.Pointer(), textC)
	}
}

//export callbackQAbstractAxis_TitleVisibleChanged
func callbackQAbstractAxis_TitleVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::titleVisibleChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectTitleVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectTitleVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::titleVisibleChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectTitleVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectTitleVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::titleVisibleChanged")
	}
}

func (ptr *QAbstractAxis) TitleVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_TitleVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstractAxis_VisibleChanged
func callbackQAbstractAxis_VisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::visibleChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAbstractAxis) ConnectVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_ConnectVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::visibleChanged", f)
	}
}

func (ptr *QAbstractAxis) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::visibleChanged")
	}
}

func (ptr *QAbstractAxis) VisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_VisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstractAxis) DestroyQAbstractAxis() {
	if ptr.Pointer() != nil {
		C.QAbstractAxis_DestroyQAbstractAxis(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractAxis_Type
func callbackQAbstractAxis_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractAxis::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractAxis__AxisType)())
	}

	return C.longlong(0)
}

func (ptr *QAbstractAxis) ConnectType(f func() QAbstractAxis__AxisType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::type", f)
	}
}

func (ptr *QAbstractAxis) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractAxis::type")
	}
}

func (ptr *QAbstractAxis) Type() QAbstractAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstractAxis__AxisType(C.QAbstractAxis_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAxis) LabelsBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QAbstractAxis_LabelsBrush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) ShadesBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QAbstractAxis_ShadesBrush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) TitleBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QAbstractAxis_TitleBrush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) LabelsColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAbstractAxis_LabelsColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) LinePenColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAbstractAxis_LinePenColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) ShadesBorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAbstractAxis_ShadesBorderColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) ShadesColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAbstractAxis_ShadesColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) LabelsFont() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.QAbstractAxis_LabelsFont(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) TitleFont() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.QAbstractAxis_TitleFont(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) GridLinePen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QAbstractAxis_GridLinePen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) LinePen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QAbstractAxis_LinePen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) MinorGridLinePen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QAbstractAxis_MinorGridLinePen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) ShadesPen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QAbstractAxis_ShadesPen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractAxis) TitleText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractAxis_TitleText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractAxis) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QAbstractAxis_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAxis) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QAbstractAxis_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAxis) IsGridLineVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractAxis_IsGridLineVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractAxis) IsLineVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractAxis_IsLineVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractAxis) IsMinorGridLineVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractAxis_IsMinorGridLineVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractAxis) IsReverse() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractAxis_IsReverse(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractAxis) IsTitleVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractAxis_IsTitleVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractAxis) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractAxis_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractAxis) LabelsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractAxis_LabelsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractAxis) ShadesVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractAxis_ShadesVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractAxis) LabelsAngle() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractAxis_LabelsAngle(ptr.Pointer())))
	}
	return 0
}

type QAbstractBarSeries struct {
	ptr unsafe.Pointer
}

type QAbstractBarSeries_ITF interface {
	QAbstractBarSeries_PTR() *QAbstractBarSeries
}

func (ptr *QAbstractBarSeries) QAbstractBarSeries_PTR() *QAbstractBarSeries {
	return ptr
}

func (ptr *QAbstractBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstractBarSeries(ptr QAbstractBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQAbstractBarSeriesFromPointer(ptr unsafe.Pointer) *QAbstractBarSeries {
	var n = new(QAbstractBarSeries)
	n.SetPointer(ptr)
	return n
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
		return C.QAbstractBarSeries_Append(ptr.Pointer(), PointerFromQBarSet(set)) != 0
	}
	return false
}

func (ptr *QAbstractBarSeries) Append2(sets []*QBarSet) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractBarSeries_Append2(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQAbstractBarSeriesFromPointer(NewQAbstractBarSeriesFromPointer(unsafe.Pointer(uintptr(1))).__append_sets_newList2())
			for _, v := range sets {
				tmpList.__append_sets_setList2(v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QAbstractBarSeries) Insert(index int, set QBarSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractBarSeries_Insert(ptr.Pointer(), C.int(int32(index)), PointerFromQBarSet(set)) != 0
	}
	return false
}

func (ptr *QAbstractBarSeries) Remove(set QBarSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractBarSeries_Remove(ptr.Pointer(), PointerFromQBarSet(set)) != 0
	}
	return false
}

func (ptr *QAbstractBarSeries) Take(set QBarSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractBarSeries_Take(ptr.Pointer(), PointerFromQBarSet(set)) != 0
	}
	return false
}

//export callbackQAbstractBarSeries_BarsetsAdded
func callbackQAbstractBarSeries_BarsetsAdded(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::barsetsAdded"); signal != nil {
		signal.(func([]*QBarSet))(func(l C.struct_QtCharts_PackedList) []*QBarSet {
			var out = make([]*QBarSet, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQAbstractBarSeriesFromPointer(l.data).__barsetsAdded_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QAbstractBarSeries) ConnectBarsetsAdded(f func(sets []*QBarSet)) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectBarsetsAdded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::barsetsAdded", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectBarsetsAdded() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectBarsetsAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::barsetsAdded")
	}
}

func (ptr *QAbstractBarSeries) BarsetsAdded(sets []*QBarSet) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_BarsetsAdded(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQAbstractBarSeriesFromPointer(NewQAbstractBarSeriesFromPointer(unsafe.Pointer(uintptr(1))).__barsetsAdded_sets_newList())
			for _, v := range sets {
				tmpList.__barsetsAdded_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQAbstractBarSeries_BarsetsRemoved
func callbackQAbstractBarSeries_BarsetsRemoved(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::barsetsRemoved"); signal != nil {
		signal.(func([]*QBarSet))(func(l C.struct_QtCharts_PackedList) []*QBarSet {
			var out = make([]*QBarSet, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQAbstractBarSeriesFromPointer(l.data).__barsetsRemoved_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QAbstractBarSeries) ConnectBarsetsRemoved(f func(sets []*QBarSet)) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectBarsetsRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::barsetsRemoved", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectBarsetsRemoved() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectBarsetsRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::barsetsRemoved")
	}
}

func (ptr *QAbstractBarSeries) BarsetsRemoved(sets []*QBarSet) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_BarsetsRemoved(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQAbstractBarSeriesFromPointer(NewQAbstractBarSeriesFromPointer(unsafe.Pointer(uintptr(1))).__barsetsRemoved_sets_newList())
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::clicked"); signal != nil {
		signal.(func(int, *QBarSet))(int(int32(index)), NewQBarSetFromPointer(barset))
	}

}

func (ptr *QAbstractBarSeries) ConnectClicked(f func(index int, barset *QBarSet)) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::clicked", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::clicked")
	}
}

func (ptr *QAbstractBarSeries) Clicked(index int, barset QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_Clicked(ptr.Pointer(), C.int(int32(index)), PointerFromQBarSet(barset))
	}
}

//export callbackQAbstractBarSeries_CountChanged
func callbackQAbstractBarSeries_CountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::countChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractBarSeries) ConnectCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::countChanged", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectCountChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::countChanged")
	}
}

func (ptr *QAbstractBarSeries) CountChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_CountChanged(ptr.Pointer())
	}
}

//export callbackQAbstractBarSeries_DoubleClicked
func callbackQAbstractBarSeries_DoubleClicked(ptr unsafe.Pointer, index C.int, barset unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::doubleClicked"); signal != nil {
		signal.(func(int, *QBarSet))(int(int32(index)), NewQBarSetFromPointer(barset))
	}

}

func (ptr *QAbstractBarSeries) ConnectDoubleClicked(f func(index int, barset *QBarSet)) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::doubleClicked", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::doubleClicked")
	}
}

func (ptr *QAbstractBarSeries) DoubleClicked(index int, barset QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DoubleClicked(ptr.Pointer(), C.int(int32(index)), PointerFromQBarSet(barset))
	}
}

//export callbackQAbstractBarSeries_Hovered
func callbackQAbstractBarSeries_Hovered(ptr unsafe.Pointer, status C.char, index C.int, barset unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::hovered"); signal != nil {
		signal.(func(bool, int, *QBarSet))(int8(status) != 0, int(int32(index)), NewQBarSetFromPointer(barset))
	}

}

func (ptr *QAbstractBarSeries) ConnectHovered(f func(status bool, index int, barset *QBarSet)) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::hovered", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::hovered")
	}
}

func (ptr *QAbstractBarSeries) Hovered(status bool, index int, barset QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))), C.int(int32(index)), PointerFromQBarSet(barset))
	}
}

//export callbackQAbstractBarSeries_LabelsAngleChanged
func callbackQAbstractBarSeries_LabelsAngleChanged(ptr unsafe.Pointer, angle C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::labelsAngleChanged"); signal != nil {
		signal.(func(float64))(float64(angle))
	}

}

func (ptr *QAbstractBarSeries) ConnectLabelsAngleChanged(f func(angle float64)) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectLabelsAngleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::labelsAngleChanged", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectLabelsAngleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectLabelsAngleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::labelsAngleChanged")
	}
}

func (ptr *QAbstractBarSeries) LabelsAngleChanged(angle float64) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_LabelsAngleChanged(ptr.Pointer(), C.double(angle))
	}
}

//export callbackQAbstractBarSeries_LabelsFormatChanged
func callbackQAbstractBarSeries_LabelsFormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::labelsFormatChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(format))
	}

}

func (ptr *QAbstractBarSeries) ConnectLabelsFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectLabelsFormatChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::labelsFormatChanged", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectLabelsFormatChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectLabelsFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::labelsFormatChanged")
	}
}

func (ptr *QAbstractBarSeries) LabelsFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QAbstractBarSeries_LabelsFormatChanged(ptr.Pointer(), formatC)
	}
}

//export callbackQAbstractBarSeries_LabelsPositionChanged
func callbackQAbstractBarSeries_LabelsPositionChanged(ptr unsafe.Pointer, position C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::labelsPositionChanged"); signal != nil {
		signal.(func(QAbstractBarSeries__LabelsPosition))(QAbstractBarSeries__LabelsPosition(position))
	}

}

func (ptr *QAbstractBarSeries) ConnectLabelsPositionChanged(f func(position QAbstractBarSeries__LabelsPosition)) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectLabelsPositionChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::labelsPositionChanged", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectLabelsPositionChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectLabelsPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::labelsPositionChanged")
	}
}

func (ptr *QAbstractBarSeries) LabelsPositionChanged(position QAbstractBarSeries__LabelsPosition) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_LabelsPositionChanged(ptr.Pointer(), C.longlong(position))
	}
}

//export callbackQAbstractBarSeries_LabelsVisibleChanged
func callbackQAbstractBarSeries_LabelsVisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::labelsVisibleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractBarSeries) ConnectLabelsVisibleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectLabelsVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::labelsVisibleChanged", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectLabelsVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectLabelsVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::labelsVisibleChanged")
	}
}

func (ptr *QAbstractBarSeries) LabelsVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_LabelsVisibleChanged(ptr.Pointer())
	}
}

//export callbackQAbstractBarSeries_Pressed
func callbackQAbstractBarSeries_Pressed(ptr unsafe.Pointer, index C.int, barset unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::pressed"); signal != nil {
		signal.(func(int, *QBarSet))(int(int32(index)), NewQBarSetFromPointer(barset))
	}

}

func (ptr *QAbstractBarSeries) ConnectPressed(f func(index int, barset *QBarSet)) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::pressed", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::pressed")
	}
}

func (ptr *QAbstractBarSeries) Pressed(index int, barset QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_Pressed(ptr.Pointer(), C.int(int32(index)), PointerFromQBarSet(barset))
	}
}

//export callbackQAbstractBarSeries_Released
func callbackQAbstractBarSeries_Released(ptr unsafe.Pointer, index C.int, barset unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::released"); signal != nil {
		signal.(func(int, *QBarSet))(int(int32(index)), NewQBarSetFromPointer(barset))
	}

}

func (ptr *QAbstractBarSeries) ConnectReleased(f func(index int, barset *QBarSet)) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::released", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::released")
	}
}

func (ptr *QAbstractBarSeries) Released(index int, barset QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_Released(ptr.Pointer(), C.int(int32(index)), PointerFromQBarSet(barset))
	}
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
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QAbstractBarSeries_SetLabelsFormat(ptr.Pointer(), formatC)
	}
}

func (ptr *QAbstractBarSeries) SetLabelsPosition(position QAbstractBarSeries__LabelsPosition) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_SetLabelsPosition(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QAbstractBarSeries) SetLabelsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_SetLabelsVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstractBarSeries_DestroyQAbstractBarSeries
func callbackQAbstractBarSeries_DestroyQAbstractBarSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractBarSeries::~QAbstractBarSeries"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractBarSeriesFromPointer(ptr).DestroyQAbstractBarSeriesDefault()
	}
}

func (ptr *QAbstractBarSeries) ConnectDestroyQAbstractBarSeries(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::~QAbstractBarSeries", f)
	}
}

func (ptr *QAbstractBarSeries) DisconnectDestroyQAbstractBarSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractBarSeries::~QAbstractBarSeries")
	}
}

func (ptr *QAbstractBarSeries) DestroyQAbstractBarSeries() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DestroyQAbstractBarSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractBarSeries) DestroyQAbstractBarSeriesDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries_DestroyQAbstractBarSeriesDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractBarSeries) LabelsPosition() QAbstractBarSeries__LabelsPosition {
	if ptr.Pointer() != nil {
		return QAbstractBarSeries__LabelsPosition(C.QAbstractBarSeries_LabelsPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractBarSeries) BarSets() []*QBarSet {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QBarSet {
			var out = make([]*QBarSet, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQAbstractBarSeriesFromPointer(l.data).__barSets_atList(i)
			}
			return out
		}(C.QAbstractBarSeries_BarSets(ptr.Pointer()))
	}
	return make([]*QBarSet, 0)
}

func (ptr *QAbstractBarSeries) LabelsFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractBarSeries_LabelsFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractBarSeries) IsLabelsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractBarSeries_IsLabelsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractBarSeries) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractBarSeries_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractBarSeries) BarWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstractBarSeries_BarWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractBarSeries) LabelsAngle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstractBarSeries_LabelsAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractBarSeries) __append_sets_atList2(i int) *QBarSet {
	if ptr.Pointer() != nil {
		return NewQBarSetFromPointer(C.QAbstractBarSeries___append_sets_atList2(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QAbstractBarSeries) __append_sets_setList2(i QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries___append_sets_setList2(ptr.Pointer(), PointerFromQBarSet(i))
	}
}

func (ptr *QAbstractBarSeries) __append_sets_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAbstractBarSeries___append_sets_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractBarSeries) __barsetsAdded_sets_atList(i int) *QBarSet {
	if ptr.Pointer() != nil {
		return NewQBarSetFromPointer(C.QAbstractBarSeries___barsetsAdded_sets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QAbstractBarSeries) __barsetsAdded_sets_setList(i QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries___barsetsAdded_sets_setList(ptr.Pointer(), PointerFromQBarSet(i))
	}
}

func (ptr *QAbstractBarSeries) __barsetsAdded_sets_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAbstractBarSeries___barsetsAdded_sets_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractBarSeries) __barsetsRemoved_sets_atList(i int) *QBarSet {
	if ptr.Pointer() != nil {
		return NewQBarSetFromPointer(C.QAbstractBarSeries___barsetsRemoved_sets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QAbstractBarSeries) __barsetsRemoved_sets_setList(i QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries___barsetsRemoved_sets_setList(ptr.Pointer(), PointerFromQBarSet(i))
	}
}

func (ptr *QAbstractBarSeries) __barsetsRemoved_sets_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAbstractBarSeries___barsetsRemoved_sets_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractBarSeries) __barSets_atList(i int) *QBarSet {
	if ptr.Pointer() != nil {
		return NewQBarSetFromPointer(C.QAbstractBarSeries___barSets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QAbstractBarSeries) __barSets_setList(i QBarSet_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractBarSeries___barSets_setList(ptr.Pointer(), PointerFromQBarSet(i))
	}
}

func (ptr *QAbstractBarSeries) __barSets_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAbstractBarSeries___barSets_newList(ptr.Pointer()))
	}
	return nil
}

type QAbstractSeries struct {
	ptr unsafe.Pointer
}

type QAbstractSeries_ITF interface {
	QAbstractSeries_PTR() *QAbstractSeries
}

func (ptr *QAbstractSeries) QAbstractSeries_PTR() *QAbstractSeries {
	return ptr
}

func (ptr *QAbstractSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstractSeries(ptr QAbstractSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSeries_PTR().Pointer()
	}
	return nil
}

func NewQAbstractSeriesFromPointer(ptr unsafe.Pointer) *QAbstractSeries {
	var n = new(QAbstractSeries)
	n.SetPointer(ptr)
	return n
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

func (ptr *QAbstractSeries) AttachedAxes() []*QAbstractAxis {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QAbstractAxis {
			var out = make([]*QAbstractAxis, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQAbstractSeriesFromPointer(l.data).__attachedAxes_atList(i)
			}
			return out
		}(C.QAbstractSeries_AttachedAxes(ptr.Pointer()))
	}
	return make([]*QAbstractAxis, 0)
}

func (ptr *QAbstractSeries) AttachAxis(axis QAbstractAxis_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSeries_AttachAxis(ptr.Pointer(), PointerFromQAbstractAxis(axis)) != 0
	}
	return false
}

func (ptr *QAbstractSeries) DetachAxis(axis QAbstractAxis_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSeries_DetachAxis(ptr.Pointer(), PointerFromQAbstractAxis(axis)) != 0
	}
	return false
}

func (ptr *QAbstractSeries) Hide() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_Hide(ptr.Pointer())
	}
}

//export callbackQAbstractSeries_NameChanged
func callbackQAbstractSeries_NameChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSeries::nameChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSeries) ConnectNameChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_ConnectNameChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSeries::nameChanged", f)
	}
}

func (ptr *QAbstractSeries) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSeries::nameChanged")
	}
}

func (ptr *QAbstractSeries) NameChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_NameChanged(ptr.Pointer())
	}
}

//export callbackQAbstractSeries_OpacityChanged
func callbackQAbstractSeries_OpacityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSeries::opacityChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSeries) ConnectOpacityChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_ConnectOpacityChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSeries::opacityChanged", f)
	}
}

func (ptr *QAbstractSeries) DisconnectOpacityChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_DisconnectOpacityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSeries::opacityChanged")
	}
}

func (ptr *QAbstractSeries) OpacityChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_OpacityChanged(ptr.Pointer())
	}
}

func (ptr *QAbstractSeries) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QAbstractSeries_SetName(ptr.Pointer(), nameC)
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

//export callbackQAbstractSeries_UseOpenGLChanged
func callbackQAbstractSeries_UseOpenGLChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSeries::useOpenGLChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSeries) ConnectUseOpenGLChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_ConnectUseOpenGLChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSeries::useOpenGLChanged", f)
	}
}

func (ptr *QAbstractSeries) DisconnectUseOpenGLChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_DisconnectUseOpenGLChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSeries::useOpenGLChanged")
	}
}

func (ptr *QAbstractSeries) UseOpenGLChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_UseOpenGLChanged(ptr.Pointer())
	}
}

//export callbackQAbstractSeries_VisibleChanged
func callbackQAbstractSeries_VisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSeries::visibleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSeries) ConnectVisibleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_ConnectVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSeries::visibleChanged", f)
	}
}

func (ptr *QAbstractSeries) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSeries::visibleChanged")
	}
}

func (ptr *QAbstractSeries) VisibleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_VisibleChanged(ptr.Pointer())
	}
}

func (ptr *QAbstractSeries) DestroyQAbstractSeries() {
	if ptr.Pointer() != nil {
		C.QAbstractSeries_DestroyQAbstractSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSeries) Chart() *QChart {
	if ptr.Pointer() != nil {
		return NewQChartFromPointer(C.QAbstractSeries_Chart(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractSeries) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractSeries_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQAbstractSeries_Type
func callbackQAbstractSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(0)
}

func (ptr *QAbstractSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSeries::type", f)
	}
}

func (ptr *QAbstractSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSeries::type")
	}
}

func (ptr *QAbstractSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QAbstractSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSeries) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSeries_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSeries) UseOpenGL() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSeries_UseOpenGL(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSeries) Opacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstractSeries_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSeries) __attachedAxes_atList(i int) *QAbstractAxis {
	if ptr.Pointer() != nil {
		return NewQAbstractAxisFromPointer(C.QAbstractSeries___attachedAxes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QAbstractSeries) __attachedAxes_setList(i QAbstractAxis_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSeries___attachedAxes_setList(ptr.Pointer(), PointerFromQAbstractAxis(i))
	}
}

func (ptr *QAbstractSeries) __attachedAxes_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAbstractSeries___attachedAxes_newList(ptr.Pointer()))
	}
	return nil
}

type QAreaLegendMarker struct {
	ptr unsafe.Pointer
}

type QAreaLegendMarker_ITF interface {
	QAreaLegendMarker_PTR() *QAreaLegendMarker
}

func (ptr *QAreaLegendMarker) QAreaLegendMarker_PTR() *QAreaLegendMarker {
	return ptr
}

func (ptr *QAreaLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAreaLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAreaLegendMarker(ptr QAreaLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAreaLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQAreaLegendMarkerFromPointer(ptr unsafe.Pointer) *QAreaLegendMarker {
	var n = new(QAreaLegendMarker)
	n.SetPointer(ptr)
	return n
}

//export callbackQAreaLegendMarker_Series
func callbackQAreaLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaLegendMarker::series"); signal != nil {
		return PointerFromQAreaSeries(signal.(func() *QAreaSeries)())
	}

	return PointerFromQAreaSeries(NewQAreaLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QAreaLegendMarker) ConnectSeries(f func() *QAreaSeries) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaLegendMarker::series", f)
	}
}

func (ptr *QAreaLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaLegendMarker::series")
	}
}

func (ptr *QAreaLegendMarker) Series() *QAreaSeries {
	if ptr.Pointer() != nil {
		return NewQAreaSeriesFromPointer(C.QAreaLegendMarker_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAreaLegendMarker) SeriesDefault() *QAreaSeries {
	if ptr.Pointer() != nil {
		return NewQAreaSeriesFromPointer(C.QAreaLegendMarker_SeriesDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQAreaLegendMarker_DestroyQAreaLegendMarker
func callbackQAreaLegendMarker_DestroyQAreaLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaLegendMarker::~QAreaLegendMarker"); signal != nil {
		signal.(func())()
	} else {
		NewQAreaLegendMarkerFromPointer(ptr).DestroyQAreaLegendMarkerDefault()
	}
}

func (ptr *QAreaLegendMarker) ConnectDestroyQAreaLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaLegendMarker::~QAreaLegendMarker", f)
	}
}

func (ptr *QAreaLegendMarker) DisconnectDestroyQAreaLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaLegendMarker::~QAreaLegendMarker")
	}
}

func (ptr *QAreaLegendMarker) DestroyQAreaLegendMarker() {
	if ptr.Pointer() != nil {
		C.QAreaLegendMarker_DestroyQAreaLegendMarker(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAreaLegendMarker) DestroyQAreaLegendMarkerDefault() {
	if ptr.Pointer() != nil {
		C.QAreaLegendMarker_DestroyQAreaLegendMarkerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QAreaSeries struct {
	ptr unsafe.Pointer
}

type QAreaSeries_ITF interface {
	QAreaSeries_PTR() *QAreaSeries
}

func (ptr *QAreaSeries) QAreaSeries_PTR() *QAreaSeries {
	return ptr
}

func (ptr *QAreaSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAreaSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAreaSeries(ptr QAreaSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAreaSeries_PTR().Pointer()
	}
	return nil
}

func NewQAreaSeriesFromPointer(ptr unsafe.Pointer) *QAreaSeries {
	var n = new(QAreaSeries)
	n.SetPointer(ptr)
	return n
}
func NewQAreaSeries2(upperSeries QLineSeries_ITF, lowerSeries QLineSeries_ITF) *QAreaSeries {
	return NewQAreaSeriesFromPointer(C.QAreaSeries_NewQAreaSeries2(PointerFromQLineSeries(upperSeries), PointerFromQLineSeries(lowerSeries)))
}

func NewQAreaSeries(parent core.QObject_ITF) *QAreaSeries {
	return NewQAreaSeriesFromPointer(C.QAreaSeries_NewQAreaSeries(core.PointerFromQObject(parent)))
}

//export callbackQAreaSeries_BorderColorChanged
func callbackQAreaSeries_BorderColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::borderColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAreaSeries) ConnectBorderColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectBorderColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::borderColorChanged", f)
	}
}

func (ptr *QAreaSeries) DisconnectBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::borderColorChanged")
	}
}

func (ptr *QAreaSeries) BorderColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_BorderColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAreaSeries_Clicked
func callbackQAreaSeries_Clicked(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::clicked"); signal != nil {
		signal.(func(*core.QPointF))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QAreaSeries) ConnectClicked(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::clicked", f)
	}
}

func (ptr *QAreaSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::clicked")
	}
}

func (ptr *QAreaSeries) Clicked(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_Clicked(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQAreaSeries_ColorChanged
func callbackQAreaSeries_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAreaSeries) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::colorChanged", f)
	}
}

func (ptr *QAreaSeries) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::colorChanged")
	}
}

func (ptr *QAreaSeries) ColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAreaSeries_DoubleClicked
func callbackQAreaSeries_DoubleClicked(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::doubleClicked"); signal != nil {
		signal.(func(*core.QPointF))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QAreaSeries) ConnectDoubleClicked(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::doubleClicked", f)
	}
}

func (ptr *QAreaSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::doubleClicked")
	}
}

func (ptr *QAreaSeries) DoubleClicked(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DoubleClicked(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQAreaSeries_Hovered
func callbackQAreaSeries_Hovered(ptr unsafe.Pointer, point unsafe.Pointer, state C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::hovered"); signal != nil {
		signal.(func(*core.QPointF, bool))(core.NewQPointFFromPointer(point), int8(state) != 0)
	}

}

func (ptr *QAreaSeries) ConnectHovered(f func(point *core.QPointF, state bool)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::hovered", f)
	}
}

func (ptr *QAreaSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::hovered")
	}
}

func (ptr *QAreaSeries) Hovered(point core.QPointF_ITF, state bool) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_Hovered(ptr.Pointer(), core.PointerFromQPointF(point), C.char(int8(qt.GoBoolToInt(state))))
	}
}

//export callbackQAreaSeries_PointLabelsClippingChanged
func callbackQAreaSeries_PointLabelsClippingChanged(ptr unsafe.Pointer, clipping C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::pointLabelsClippingChanged"); signal != nil {
		signal.(func(bool))(int8(clipping) != 0)
	}

}

func (ptr *QAreaSeries) ConnectPointLabelsClippingChanged(f func(clipping bool)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectPointLabelsClippingChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pointLabelsClippingChanged", f)
	}
}

func (ptr *QAreaSeries) DisconnectPointLabelsClippingChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPointLabelsClippingChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pointLabelsClippingChanged")
	}
}

func (ptr *QAreaSeries) PointLabelsClippingChanged(clipping bool) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_PointLabelsClippingChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(clipping))))
	}
}

//export callbackQAreaSeries_PointLabelsColorChanged
func callbackQAreaSeries_PointLabelsColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::pointLabelsColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAreaSeries) ConnectPointLabelsColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectPointLabelsColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pointLabelsColorChanged", f)
	}
}

func (ptr *QAreaSeries) DisconnectPointLabelsColorChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPointLabelsColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pointLabelsColorChanged")
	}
}

func (ptr *QAreaSeries) PointLabelsColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_PointLabelsColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAreaSeries_PointLabelsFontChanged
func callbackQAreaSeries_PointLabelsFontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::pointLabelsFontChanged"); signal != nil {
		signal.(func(*gui.QFont))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *QAreaSeries) ConnectPointLabelsFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectPointLabelsFontChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pointLabelsFontChanged", f)
	}
}

func (ptr *QAreaSeries) DisconnectPointLabelsFontChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPointLabelsFontChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pointLabelsFontChanged")
	}
}

func (ptr *QAreaSeries) PointLabelsFontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_PointLabelsFontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

//export callbackQAreaSeries_PointLabelsFormatChanged
func callbackQAreaSeries_PointLabelsFormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::pointLabelsFormatChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(format))
	}

}

func (ptr *QAreaSeries) ConnectPointLabelsFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectPointLabelsFormatChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pointLabelsFormatChanged", f)
	}
}

func (ptr *QAreaSeries) DisconnectPointLabelsFormatChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPointLabelsFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pointLabelsFormatChanged")
	}
}

func (ptr *QAreaSeries) PointLabelsFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QAreaSeries_PointLabelsFormatChanged(ptr.Pointer(), formatC)
	}
}

//export callbackQAreaSeries_PointLabelsVisibilityChanged
func callbackQAreaSeries_PointLabelsVisibilityChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::pointLabelsVisibilityChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAreaSeries) ConnectPointLabelsVisibilityChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectPointLabelsVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pointLabelsVisibilityChanged", f)
	}
}

func (ptr *QAreaSeries) DisconnectPointLabelsVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPointLabelsVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pointLabelsVisibilityChanged")
	}
}

func (ptr *QAreaSeries) PointLabelsVisibilityChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_PointLabelsVisibilityChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAreaSeries_Pressed
func callbackQAreaSeries_Pressed(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::pressed"); signal != nil {
		signal.(func(*core.QPointF))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QAreaSeries) ConnectPressed(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pressed", f)
	}
}

func (ptr *QAreaSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::pressed")
	}
}

func (ptr *QAreaSeries) Pressed(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_Pressed(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQAreaSeries_Released
func callbackQAreaSeries_Released(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::released"); signal != nil {
		signal.(func(*core.QPointF))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QAreaSeries) ConnectReleased(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {
		C.QAreaSeries_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::released", f)
	}
}

func (ptr *QAreaSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::released")
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
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QAreaSeries_SetPointLabelsFormat(ptr.Pointer(), formatC)
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

func (ptr *QAreaSeries) DestroyQAreaSeries() {
	if ptr.Pointer() != nil {
		C.QAreaSeries_DestroyQAreaSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAreaSeries_Type
func callbackQAreaSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAreaSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQAreaSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QAreaSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::type", f)
	}
}

func (ptr *QAreaSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAreaSeries::type")
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

func (ptr *QAreaSeries) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QAreaSeries_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QAreaSeries) BorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAreaSeries_BorderColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAreaSeries) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAreaSeries_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAreaSeries) PointLabelsColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAreaSeries_PointLabelsColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAreaSeries) PointLabelsFont() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.QAreaSeries_PointLabelsFont(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QAreaSeries) LowerSeries() *QLineSeries {
	if ptr.Pointer() != nil {
		return NewQLineSeriesFromPointer(C.QAreaSeries_LowerSeries(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAreaSeries) UpperSeries() *QLineSeries {
	if ptr.Pointer() != nil {
		return NewQLineSeriesFromPointer(C.QAreaSeries_UpperSeries(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAreaSeries) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QAreaSeries_Pen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QAreaSeries) PointLabelsFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAreaSeries_PointLabelsFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAreaSeries) PointLabelsClipping() bool {
	if ptr.Pointer() != nil {
		return C.QAreaSeries_PointLabelsClipping(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAreaSeries) PointLabelsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAreaSeries_PointLabelsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAreaSeries) PointsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAreaSeries_PointsVisible(ptr.Pointer()) != 0
	}
	return false
}

type QBarCategoryAxis struct {
	ptr unsafe.Pointer
}

type QBarCategoryAxis_ITF interface {
	QBarCategoryAxis_PTR() *QBarCategoryAxis
}

func (ptr *QBarCategoryAxis) QBarCategoryAxis_PTR() *QBarCategoryAxis {
	return ptr
}

func (ptr *QBarCategoryAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBarCategoryAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBarCategoryAxis(ptr QBarCategoryAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarCategoryAxis_PTR().Pointer()
	}
	return nil
}

func NewQBarCategoryAxisFromPointer(ptr unsafe.Pointer) *QBarCategoryAxis {
	var n = new(QBarCategoryAxis)
	n.SetPointer(ptr)
	return n
}
func NewQBarCategoryAxis(parent core.QObject_ITF) *QBarCategoryAxis {
	return NewQBarCategoryAxisFromPointer(C.QBarCategoryAxis_NewQBarCategoryAxis(core.PointerFromQObject(parent)))
}

func (ptr *QBarCategoryAxis) Categories() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QBarCategoryAxis_Categories(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QBarCategoryAxis) Append2(category string) {
	if ptr.Pointer() != nil {
		var categoryC = C.CString(category)
		defer C.free(unsafe.Pointer(categoryC))
		C.QBarCategoryAxis_Append2(ptr.Pointer(), categoryC)
	}
}

func (ptr *QBarCategoryAxis) Append(categories []string) {
	if ptr.Pointer() != nil {
		var categoriesC = C.CString(strings.Join(categories, "|"))
		defer C.free(unsafe.Pointer(categoriesC))
		C.QBarCategoryAxis_Append(ptr.Pointer(), categoriesC)
	}
}

//export callbackQBarCategoryAxis_CategoriesChanged
func callbackQBarCategoryAxis_CategoriesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarCategoryAxis::categoriesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBarCategoryAxis) ConnectCategoriesChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_ConnectCategoriesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarCategoryAxis::categoriesChanged", f)
	}
}

func (ptr *QBarCategoryAxis) DisconnectCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DisconnectCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarCategoryAxis::categoriesChanged")
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

//export callbackQBarCategoryAxis_CountChanged
func callbackQBarCategoryAxis_CountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarCategoryAxis::countChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBarCategoryAxis) ConnectCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_ConnectCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarCategoryAxis::countChanged", f)
	}
}

func (ptr *QBarCategoryAxis) DisconnectCountChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DisconnectCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarCategoryAxis::countChanged")
	}
}

func (ptr *QBarCategoryAxis) CountChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_CountChanged(ptr.Pointer())
	}
}

func (ptr *QBarCategoryAxis) Insert(index int, category string) {
	if ptr.Pointer() != nil {
		var categoryC = C.CString(category)
		defer C.free(unsafe.Pointer(categoryC))
		C.QBarCategoryAxis_Insert(ptr.Pointer(), C.int(int32(index)), categoryC)
	}
}

//export callbackQBarCategoryAxis_MaxChanged
func callbackQBarCategoryAxis_MaxChanged(ptr unsafe.Pointer, max C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarCategoryAxis::maxChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(max))
	}

}

func (ptr *QBarCategoryAxis) ConnectMaxChanged(f func(max string)) {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_ConnectMaxChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarCategoryAxis::maxChanged", f)
	}
}

func (ptr *QBarCategoryAxis) DisconnectMaxChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DisconnectMaxChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarCategoryAxis::maxChanged")
	}
}

func (ptr *QBarCategoryAxis) MaxChanged(max string) {
	if ptr.Pointer() != nil {
		var maxC = C.CString(max)
		defer C.free(unsafe.Pointer(maxC))
		C.QBarCategoryAxis_MaxChanged(ptr.Pointer(), maxC)
	}
}

//export callbackQBarCategoryAxis_MinChanged
func callbackQBarCategoryAxis_MinChanged(ptr unsafe.Pointer, min C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarCategoryAxis::minChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(min))
	}

}

func (ptr *QBarCategoryAxis) ConnectMinChanged(f func(min string)) {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_ConnectMinChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarCategoryAxis::minChanged", f)
	}
}

func (ptr *QBarCategoryAxis) DisconnectMinChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DisconnectMinChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarCategoryAxis::minChanged")
	}
}

func (ptr *QBarCategoryAxis) MinChanged(min string) {
	if ptr.Pointer() != nil {
		var minC = C.CString(min)
		defer C.free(unsafe.Pointer(minC))
		C.QBarCategoryAxis_MinChanged(ptr.Pointer(), minC)
	}
}

//export callbackQBarCategoryAxis_RangeChanged
func callbackQBarCategoryAxis_RangeChanged(ptr unsafe.Pointer, min C.struct_QtCharts_PackedString, max C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarCategoryAxis::rangeChanged"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(min), cGoUnpackString(max))
	}

}

func (ptr *QBarCategoryAxis) ConnectRangeChanged(f func(min string, max string)) {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_ConnectRangeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarCategoryAxis::rangeChanged", f)
	}
}

func (ptr *QBarCategoryAxis) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarCategoryAxis::rangeChanged")
	}
}

func (ptr *QBarCategoryAxis) RangeChanged(min string, max string) {
	if ptr.Pointer() != nil {
		var minC = C.CString(min)
		defer C.free(unsafe.Pointer(minC))
		var maxC = C.CString(max)
		defer C.free(unsafe.Pointer(maxC))
		C.QBarCategoryAxis_RangeChanged(ptr.Pointer(), minC, maxC)
	}
}

func (ptr *QBarCategoryAxis) Remove(category string) {
	if ptr.Pointer() != nil {
		var categoryC = C.CString(category)
		defer C.free(unsafe.Pointer(categoryC))
		C.QBarCategoryAxis_Remove(ptr.Pointer(), categoryC)
	}
}

func (ptr *QBarCategoryAxis) Replace(oldCategory string, newCategory string) {
	if ptr.Pointer() != nil {
		var oldCategoryC = C.CString(oldCategory)
		defer C.free(unsafe.Pointer(oldCategoryC))
		var newCategoryC = C.CString(newCategory)
		defer C.free(unsafe.Pointer(newCategoryC))
		C.QBarCategoryAxis_Replace(ptr.Pointer(), oldCategoryC, newCategoryC)
	}
}

func (ptr *QBarCategoryAxis) SetCategories(categories []string) {
	if ptr.Pointer() != nil {
		var categoriesC = C.CString(strings.Join(categories, "|"))
		defer C.free(unsafe.Pointer(categoriesC))
		C.QBarCategoryAxis_SetCategories(ptr.Pointer(), categoriesC)
	}
}

func (ptr *QBarCategoryAxis) SetMax(max string) {
	if ptr.Pointer() != nil {
		var maxC = C.CString(max)
		defer C.free(unsafe.Pointer(maxC))
		C.QBarCategoryAxis_SetMax(ptr.Pointer(), maxC)
	}
}

func (ptr *QBarCategoryAxis) SetMin(min string) {
	if ptr.Pointer() != nil {
		var minC = C.CString(min)
		defer C.free(unsafe.Pointer(minC))
		C.QBarCategoryAxis_SetMin(ptr.Pointer(), minC)
	}
}

func (ptr *QBarCategoryAxis) SetRange(minCategory string, maxCategory string) {
	if ptr.Pointer() != nil {
		var minCategoryC = C.CString(minCategory)
		defer C.free(unsafe.Pointer(minCategoryC))
		var maxCategoryC = C.CString(maxCategory)
		defer C.free(unsafe.Pointer(maxCategoryC))
		C.QBarCategoryAxis_SetRange(ptr.Pointer(), minCategoryC, maxCategoryC)
	}
}

func (ptr *QBarCategoryAxis) DestroyQBarCategoryAxis() {
	if ptr.Pointer() != nil {
		C.QBarCategoryAxis_DestroyQBarCategoryAxis(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarCategoryAxis) At(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBarCategoryAxis_At(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QBarCategoryAxis) Max() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBarCategoryAxis_Max(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBarCategoryAxis) Min() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBarCategoryAxis_Min(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBarCategoryAxis) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBarCategoryAxis_Count(ptr.Pointer())))
	}
	return 0
}

type QBarLegendMarker struct {
	ptr unsafe.Pointer
}

type QBarLegendMarker_ITF interface {
	QBarLegendMarker_PTR() *QBarLegendMarker
}

func (ptr *QBarLegendMarker) QBarLegendMarker_PTR() *QBarLegendMarker {
	return ptr
}

func (ptr *QBarLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBarLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBarLegendMarker(ptr QBarLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQBarLegendMarkerFromPointer(ptr unsafe.Pointer) *QBarLegendMarker {
	var n = new(QBarLegendMarker)
	n.SetPointer(ptr)
	return n
}

//export callbackQBarLegendMarker_Series
func callbackQBarLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarLegendMarker::series"); signal != nil {
		return PointerFromQAbstractBarSeries(signal.(func() *QAbstractBarSeries)())
	}

	return PointerFromQAbstractBarSeries(NewQBarLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QBarLegendMarker) ConnectSeries(f func() *QAbstractBarSeries) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarLegendMarker::series", f)
	}
}

func (ptr *QBarLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarLegendMarker::series")
	}
}

func (ptr *QBarLegendMarker) Series() *QAbstractBarSeries {
	if ptr.Pointer() != nil {
		return NewQAbstractBarSeriesFromPointer(C.QBarLegendMarker_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBarLegendMarker) SeriesDefault() *QAbstractBarSeries {
	if ptr.Pointer() != nil {
		return NewQAbstractBarSeriesFromPointer(C.QBarLegendMarker_SeriesDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBarLegendMarker) Barset() *QBarSet {
	if ptr.Pointer() != nil {
		return NewQBarSetFromPointer(C.QBarLegendMarker_Barset(ptr.Pointer()))
	}
	return nil
}

//export callbackQBarLegendMarker_DestroyQBarLegendMarker
func callbackQBarLegendMarker_DestroyQBarLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarLegendMarker::~QBarLegendMarker"); signal != nil {
		signal.(func())()
	} else {
		NewQBarLegendMarkerFromPointer(ptr).DestroyQBarLegendMarkerDefault()
	}
}

func (ptr *QBarLegendMarker) ConnectDestroyQBarLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarLegendMarker::~QBarLegendMarker", f)
	}
}

func (ptr *QBarLegendMarker) DisconnectDestroyQBarLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarLegendMarker::~QBarLegendMarker")
	}
}

func (ptr *QBarLegendMarker) DestroyQBarLegendMarker() {
	if ptr.Pointer() != nil {
		C.QBarLegendMarker_DestroyQBarLegendMarker(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarLegendMarker) DestroyQBarLegendMarkerDefault() {
	if ptr.Pointer() != nil {
		C.QBarLegendMarker_DestroyQBarLegendMarkerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QBarSeries struct {
	ptr unsafe.Pointer
}

type QBarSeries_ITF interface {
	QBarSeries_PTR() *QBarSeries
}

func (ptr *QBarSeries) QBarSeries_PTR() *QBarSeries {
	return ptr
}

func (ptr *QBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBarSeries(ptr QBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQBarSeriesFromPointer(ptr unsafe.Pointer) *QBarSeries {
	var n = new(QBarSeries)
	n.SetPointer(ptr)
	return n
}
func NewQBarSeries(parent core.QObject_ITF) *QBarSeries {
	return NewQBarSeriesFromPointer(C.QBarSeries_NewQBarSeries(core.PointerFromQObject(parent)))
}

func (ptr *QBarSeries) DestroyQBarSeries() {
	if ptr.Pointer() != nil {
		C.QBarSeries_DestroyQBarSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBarSeries_Type
func callbackQBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSeries::type", f)
	}
}

func (ptr *QBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSeries::type")
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

type QBarSet struct {
	ptr unsafe.Pointer
}

type QBarSet_ITF interface {
	QBarSet_PTR() *QBarSet
}

func (ptr *QBarSet) QBarSet_PTR() *QBarSet {
	return ptr
}

func (ptr *QBarSet) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBarSet) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBarSet(ptr QBarSet_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarSet_PTR().Pointer()
	}
	return nil
}

func NewQBarSetFromPointer(ptr unsafe.Pointer) *QBarSet {
	var n = new(QBarSet)
	n.SetPointer(ptr)
	return n
}
func NewQBarSet(label string, parent core.QObject_ITF) *QBarSet {
	var labelC = C.CString(label)
	defer C.free(unsafe.Pointer(labelC))
	return NewQBarSetFromPointer(C.QBarSet_NewQBarSet(labelC, core.PointerFromQObject(parent)))
}

func (ptr *QBarSet) BorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QBarSet_BorderColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QBarSet_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) LabelColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QBarSet_LabelColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) Append2(values []float64) {
	if ptr.Pointer() != nil {
		C.QBarSet_Append2(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQBarSetFromPointer(NewQBarSetFromPointer(unsafe.Pointer(uintptr(1))).__append_values_newList2())
			for _, v := range values {
				tmpList.__append_values_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QBarSet) Append(value float64) {
	if ptr.Pointer() != nil {
		C.QBarSet_Append(ptr.Pointer(), C.double(value))
	}
}

//export callbackQBarSet_BorderColorChanged
func callbackQBarSet_BorderColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::borderColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QBarSet) ConnectBorderColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectBorderColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::borderColorChanged", f)
	}
}

func (ptr *QBarSet) DisconnectBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::borderColorChanged")
	}
}

func (ptr *QBarSet) BorderColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_BorderColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQBarSet_BrushChanged
func callbackQBarSet_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::brushChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBarSet) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::brushChanged", f)
	}
}

func (ptr *QBarSet) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::brushChanged")
	}
}

func (ptr *QBarSet) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_BrushChanged(ptr.Pointer())
	}
}

//export callbackQBarSet_Clicked
func callbackQBarSet_Clicked(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::clicked"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectClicked(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::clicked", f)
	}
}

func (ptr *QBarSet) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::clicked")
	}
}

func (ptr *QBarSet) Clicked(index int) {
	if ptr.Pointer() != nil {
		C.QBarSet_Clicked(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQBarSet_ColorChanged
func callbackQBarSet_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QBarSet) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::colorChanged", f)
	}
}

func (ptr *QBarSet) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::colorChanged")
	}
}

func (ptr *QBarSet) ColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQBarSet_DoubleClicked
func callbackQBarSet_DoubleClicked(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::doubleClicked"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectDoubleClicked(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::doubleClicked", f)
	}
}

func (ptr *QBarSet) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::doubleClicked")
	}
}

func (ptr *QBarSet) DoubleClicked(index int) {
	if ptr.Pointer() != nil {
		C.QBarSet_DoubleClicked(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQBarSet_Hovered
func callbackQBarSet_Hovered(ptr unsafe.Pointer, status C.char, index C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::hovered"); signal != nil {
		signal.(func(bool, int))(int8(status) != 0, int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectHovered(f func(status bool, index int)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::hovered", f)
	}
}

func (ptr *QBarSet) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::hovered")
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

//export callbackQBarSet_LabelBrushChanged
func callbackQBarSet_LabelBrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::labelBrushChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBarSet) ConnectLabelBrushChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectLabelBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::labelBrushChanged", f)
	}
}

func (ptr *QBarSet) DisconnectLabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectLabelBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::labelBrushChanged")
	}
}

func (ptr *QBarSet) LabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_LabelBrushChanged(ptr.Pointer())
	}
}

//export callbackQBarSet_LabelChanged
func callbackQBarSet_LabelChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::labelChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBarSet) ConnectLabelChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectLabelChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::labelChanged", f)
	}
}

func (ptr *QBarSet) DisconnectLabelChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectLabelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::labelChanged")
	}
}

func (ptr *QBarSet) LabelChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_LabelChanged(ptr.Pointer())
	}
}

//export callbackQBarSet_LabelColorChanged
func callbackQBarSet_LabelColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::labelColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QBarSet) ConnectLabelColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectLabelColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::labelColorChanged", f)
	}
}

func (ptr *QBarSet) DisconnectLabelColorChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectLabelColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::labelColorChanged")
	}
}

func (ptr *QBarSet) LabelColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QBarSet_LabelColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQBarSet_LabelFontChanged
func callbackQBarSet_LabelFontChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::labelFontChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBarSet) ConnectLabelFontChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectLabelFontChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::labelFontChanged", f)
	}
}

func (ptr *QBarSet) DisconnectLabelFontChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectLabelFontChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::labelFontChanged")
	}
}

func (ptr *QBarSet) LabelFontChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_LabelFontChanged(ptr.Pointer())
	}
}

//export callbackQBarSet_PenChanged
func callbackQBarSet_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::penChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBarSet) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectPenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::penChanged", f)
	}
}

func (ptr *QBarSet) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::penChanged")
	}
}

func (ptr *QBarSet) PenChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_PenChanged(ptr.Pointer())
	}
}

//export callbackQBarSet_Pressed
func callbackQBarSet_Pressed(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::pressed"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectPressed(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::pressed", f)
	}
}

func (ptr *QBarSet) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::pressed")
	}
}

func (ptr *QBarSet) Pressed(index int) {
	if ptr.Pointer() != nil {
		C.QBarSet_Pressed(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQBarSet_Released
func callbackQBarSet_Released(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::released"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectReleased(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::released", f)
	}
}

func (ptr *QBarSet) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::released")
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
		var labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
		C.QBarSet_SetLabel(ptr.Pointer(), labelC)
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

//export callbackQBarSet_ValueChanged
func callbackQBarSet_ValueChanged(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::valueChanged"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QBarSet) ConnectValueChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::valueChanged", f)
	}
}

func (ptr *QBarSet) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::valueChanged")
	}
}

func (ptr *QBarSet) ValueChanged(index int) {
	if ptr.Pointer() != nil {
		C.QBarSet_ValueChanged(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQBarSet_ValuesAdded
func callbackQBarSet_ValuesAdded(ptr unsafe.Pointer, index C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::valuesAdded"); signal != nil {
		signal.(func(int, int))(int(int32(index)), int(int32(count)))
	}

}

func (ptr *QBarSet) ConnectValuesAdded(f func(index int, count int)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectValuesAdded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::valuesAdded", f)
	}
}

func (ptr *QBarSet) DisconnectValuesAdded() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectValuesAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::valuesAdded")
	}
}

func (ptr *QBarSet) ValuesAdded(index int, count int) {
	if ptr.Pointer() != nil {
		C.QBarSet_ValuesAdded(ptr.Pointer(), C.int(int32(index)), C.int(int32(count)))
	}
}

//export callbackQBarSet_ValuesRemoved
func callbackQBarSet_ValuesRemoved(ptr unsafe.Pointer, index C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::valuesRemoved"); signal != nil {
		signal.(func(int, int))(int(int32(index)), int(int32(count)))
	}

}

func (ptr *QBarSet) ConnectValuesRemoved(f func(index int, count int)) {
	if ptr.Pointer() != nil {
		C.QBarSet_ConnectValuesRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::valuesRemoved", f)
	}
}

func (ptr *QBarSet) DisconnectValuesRemoved() {
	if ptr.Pointer() != nil {
		C.QBarSet_DisconnectValuesRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::valuesRemoved")
	}
}

func (ptr *QBarSet) ValuesRemoved(index int, count int) {
	if ptr.Pointer() != nil {
		C.QBarSet_ValuesRemoved(ptr.Pointer(), C.int(int32(index)), C.int(int32(count)))
	}
}

//export callbackQBarSet_DestroyQBarSet
func callbackQBarSet_DestroyQBarSet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBarSet::~QBarSet"); signal != nil {
		signal.(func())()
	} else {
		NewQBarSetFromPointer(ptr).DestroyQBarSetDefault()
	}
}

func (ptr *QBarSet) ConnectDestroyQBarSet(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::~QBarSet", f)
	}
}

func (ptr *QBarSet) DisconnectDestroyQBarSet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBarSet::~QBarSet")
	}
}

func (ptr *QBarSet) DestroyQBarSet() {
	if ptr.Pointer() != nil {
		C.QBarSet_DestroyQBarSet(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarSet) DestroyQBarSetDefault() {
	if ptr.Pointer() != nil {
		C.QBarSet_DestroyQBarSetDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarSet) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QBarSet_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) LabelBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QBarSet_LabelBrush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) LabelFont() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.QBarSet_LabelFont(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QBarSet_Pen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QBarSet) Label() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBarSet_Label(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBarSet) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBarSet_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBarSet) At(index int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QBarSet_At(ptr.Pointer(), C.int(int32(index))))
	}
	return 0
}

func (ptr *QBarSet) Sum() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QBarSet_Sum(ptr.Pointer()))
	}
	return 0
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QBarSet___append_values_newList2(ptr.Pointer()))
	}
	return nil
}

type QBoxPlotLegendMarker struct {
	ptr unsafe.Pointer
}

type QBoxPlotLegendMarker_ITF interface {
	QBoxPlotLegendMarker_PTR() *QBoxPlotLegendMarker
}

func (ptr *QBoxPlotLegendMarker) QBoxPlotLegendMarker_PTR() *QBoxPlotLegendMarker {
	return ptr
}

func (ptr *QBoxPlotLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBoxPlotLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBoxPlotLegendMarker(ptr QBoxPlotLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxPlotLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQBoxPlotLegendMarkerFromPointer(ptr unsafe.Pointer) *QBoxPlotLegendMarker {
	var n = new(QBoxPlotLegendMarker)
	n.SetPointer(ptr)
	return n
}

//export callbackQBoxPlotLegendMarker_Series
func callbackQBoxPlotLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotLegendMarker::series"); signal != nil {
		return PointerFromQBoxPlotSeries(signal.(func() *QBoxPlotSeries)())
	}

	return PointerFromQBoxPlotSeries(NewQBoxPlotLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QBoxPlotLegendMarker) ConnectSeries(f func() *QBoxPlotSeries) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotLegendMarker::series", f)
	}
}

func (ptr *QBoxPlotLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotLegendMarker::series")
	}
}

func (ptr *QBoxPlotLegendMarker) Series() *QBoxPlotSeries {
	if ptr.Pointer() != nil {
		return NewQBoxPlotSeriesFromPointer(C.QBoxPlotLegendMarker_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBoxPlotLegendMarker) SeriesDefault() *QBoxPlotSeries {
	if ptr.Pointer() != nil {
		return NewQBoxPlotSeriesFromPointer(C.QBoxPlotLegendMarker_SeriesDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQBoxPlotLegendMarker_DestroyQBoxPlotLegendMarker
func callbackQBoxPlotLegendMarker_DestroyQBoxPlotLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotLegendMarker::~QBoxPlotLegendMarker"); signal != nil {
		signal.(func())()
	} else {
		NewQBoxPlotLegendMarkerFromPointer(ptr).DestroyQBoxPlotLegendMarkerDefault()
	}
}

func (ptr *QBoxPlotLegendMarker) ConnectDestroyQBoxPlotLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotLegendMarker::~QBoxPlotLegendMarker", f)
	}
}

func (ptr *QBoxPlotLegendMarker) DisconnectDestroyQBoxPlotLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotLegendMarker::~QBoxPlotLegendMarker")
	}
}

func (ptr *QBoxPlotLegendMarker) DestroyQBoxPlotLegendMarker() {
	if ptr.Pointer() != nil {
		C.QBoxPlotLegendMarker_DestroyQBoxPlotLegendMarker(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBoxPlotLegendMarker) DestroyQBoxPlotLegendMarkerDefault() {
	if ptr.Pointer() != nil {
		C.QBoxPlotLegendMarker_DestroyQBoxPlotLegendMarkerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QBoxPlotSeries struct {
	ptr unsafe.Pointer
}

type QBoxPlotSeries_ITF interface {
	QBoxPlotSeries_PTR() *QBoxPlotSeries
}

func (ptr *QBoxPlotSeries) QBoxPlotSeries_PTR() *QBoxPlotSeries {
	return ptr
}

func (ptr *QBoxPlotSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBoxPlotSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBoxPlotSeries(ptr QBoxPlotSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxPlotSeries_PTR().Pointer()
	}
	return nil
}

func NewQBoxPlotSeriesFromPointer(ptr unsafe.Pointer) *QBoxPlotSeries {
	var n = new(QBoxPlotSeries)
	n.SetPointer(ptr)
	return n
}
func NewQBoxPlotSeries(parent core.QObject_ITF) *QBoxPlotSeries {
	return NewQBoxPlotSeriesFromPointer(C.QBoxPlotSeries_NewQBoxPlotSeries(core.PointerFromQObject(parent)))
}

func (ptr *QBoxPlotSeries) Append(set QBoxSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBoxPlotSeries_Append(ptr.Pointer(), PointerFromQBoxSet(set)) != 0
	}
	return false
}

func (ptr *QBoxPlotSeries) Append2(sets []*QBoxSet) bool {
	if ptr.Pointer() != nil {
		return C.QBoxPlotSeries_Append2(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQBoxPlotSeriesFromPointer(NewQBoxPlotSeriesFromPointer(unsafe.Pointer(uintptr(1))).__append_sets_newList2())
			for _, v := range sets {
				tmpList.__append_sets_setList2(v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QBoxPlotSeries) BoxOutlineVisible() bool {
	if ptr.Pointer() != nil {
		return C.QBoxPlotSeries_BoxOutlineVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBoxPlotSeries) Insert(index int, set QBoxSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBoxPlotSeries_Insert(ptr.Pointer(), C.int(int32(index)), PointerFromQBoxSet(set)) != 0
	}
	return false
}

func (ptr *QBoxPlotSeries) Remove(set QBoxSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBoxPlotSeries_Remove(ptr.Pointer(), PointerFromQBoxSet(set)) != 0
	}
	return false
}

func (ptr *QBoxPlotSeries) Take(set QBoxSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBoxPlotSeries_Take(ptr.Pointer(), PointerFromQBoxSet(set)) != 0
	}
	return false
}

func (ptr *QBoxPlotSeries) BoxWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QBoxPlotSeries_BoxWidth(ptr.Pointer()))
	}
	return 0
}

//export callbackQBoxPlotSeries_BoxOutlineVisibilityChanged
func callbackQBoxPlotSeries_BoxOutlineVisibilityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::boxOutlineVisibilityChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxPlotSeries) ConnectBoxOutlineVisibilityChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectBoxOutlineVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::boxOutlineVisibilityChanged", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectBoxOutlineVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectBoxOutlineVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::boxOutlineVisibilityChanged")
	}
}

func (ptr *QBoxPlotSeries) BoxOutlineVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_BoxOutlineVisibilityChanged(ptr.Pointer())
	}
}

//export callbackQBoxPlotSeries_BoxWidthChanged
func callbackQBoxPlotSeries_BoxWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::boxWidthChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxPlotSeries) ConnectBoxWidthChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectBoxWidthChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::boxWidthChanged", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectBoxWidthChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectBoxWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::boxWidthChanged")
	}
}

func (ptr *QBoxPlotSeries) BoxWidthChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_BoxWidthChanged(ptr.Pointer())
	}
}

//export callbackQBoxPlotSeries_BoxsetsAdded
func callbackQBoxPlotSeries_BoxsetsAdded(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::boxsetsAdded"); signal != nil {
		signal.(func([]*QBoxSet))(func(l C.struct_QtCharts_PackedList) []*QBoxSet {
			var out = make([]*QBoxSet, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBoxPlotSeriesFromPointer(l.data).__boxsetsAdded_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QBoxPlotSeries) ConnectBoxsetsAdded(f func(sets []*QBoxSet)) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectBoxsetsAdded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::boxsetsAdded", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectBoxsetsAdded() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectBoxsetsAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::boxsetsAdded")
	}
}

func (ptr *QBoxPlotSeries) BoxsetsAdded(sets []*QBoxSet) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_BoxsetsAdded(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQBoxPlotSeriesFromPointer(NewQBoxPlotSeriesFromPointer(unsafe.Pointer(uintptr(1))).__boxsetsAdded_sets_newList())
			for _, v := range sets {
				tmpList.__boxsetsAdded_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQBoxPlotSeries_BoxsetsRemoved
func callbackQBoxPlotSeries_BoxsetsRemoved(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::boxsetsRemoved"); signal != nil {
		signal.(func([]*QBoxSet))(func(l C.struct_QtCharts_PackedList) []*QBoxSet {
			var out = make([]*QBoxSet, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBoxPlotSeriesFromPointer(l.data).__boxsetsRemoved_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QBoxPlotSeries) ConnectBoxsetsRemoved(f func(sets []*QBoxSet)) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectBoxsetsRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::boxsetsRemoved", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectBoxsetsRemoved() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectBoxsetsRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::boxsetsRemoved")
	}
}

func (ptr *QBoxPlotSeries) BoxsetsRemoved(sets []*QBoxSet) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_BoxsetsRemoved(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQBoxPlotSeriesFromPointer(NewQBoxPlotSeriesFromPointer(unsafe.Pointer(uintptr(1))).__boxsetsRemoved_sets_newList())
			for _, v := range sets {
				tmpList.__boxsetsRemoved_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQBoxPlotSeries_BrushChanged
func callbackQBoxPlotSeries_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::brushChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxPlotSeries) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::brushChanged", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::brushChanged")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::clicked"); signal != nil {
		signal.(func(*QBoxSet))(NewQBoxSetFromPointer(boxset))
	}

}

func (ptr *QBoxPlotSeries) ConnectClicked(f func(boxset *QBoxSet)) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::clicked", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::clicked")
	}
}

func (ptr *QBoxPlotSeries) Clicked(boxset QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_Clicked(ptr.Pointer(), PointerFromQBoxSet(boxset))
	}
}

//export callbackQBoxPlotSeries_CountChanged
func callbackQBoxPlotSeries_CountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::countChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxPlotSeries) ConnectCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::countChanged", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectCountChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::countChanged")
	}
}

func (ptr *QBoxPlotSeries) CountChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_CountChanged(ptr.Pointer())
	}
}

//export callbackQBoxPlotSeries_DoubleClicked
func callbackQBoxPlotSeries_DoubleClicked(ptr unsafe.Pointer, boxset unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::doubleClicked"); signal != nil {
		signal.(func(*QBoxSet))(NewQBoxSetFromPointer(boxset))
	}

}

func (ptr *QBoxPlotSeries) ConnectDoubleClicked(f func(boxset *QBoxSet)) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::doubleClicked", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::doubleClicked")
	}
}

func (ptr *QBoxPlotSeries) DoubleClicked(boxset QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DoubleClicked(ptr.Pointer(), PointerFromQBoxSet(boxset))
	}
}

//export callbackQBoxPlotSeries_Hovered
func callbackQBoxPlotSeries_Hovered(ptr unsafe.Pointer, status C.char, boxset unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::hovered"); signal != nil {
		signal.(func(bool, *QBoxSet))(int8(status) != 0, NewQBoxSetFromPointer(boxset))
	}

}

func (ptr *QBoxPlotSeries) ConnectHovered(f func(status bool, boxset *QBoxSet)) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::hovered", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::hovered")
	}
}

func (ptr *QBoxPlotSeries) Hovered(status bool, boxset QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))), PointerFromQBoxSet(boxset))
	}
}

//export callbackQBoxPlotSeries_PenChanged
func callbackQBoxPlotSeries_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::penChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxPlotSeries) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectPenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::penChanged", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::penChanged")
	}
}

func (ptr *QBoxPlotSeries) PenChanged() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_PenChanged(ptr.Pointer())
	}
}

//export callbackQBoxPlotSeries_Pressed
func callbackQBoxPlotSeries_Pressed(ptr unsafe.Pointer, boxset unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::pressed"); signal != nil {
		signal.(func(*QBoxSet))(NewQBoxSetFromPointer(boxset))
	}

}

func (ptr *QBoxPlotSeries) ConnectPressed(f func(boxset *QBoxSet)) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::pressed", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::pressed")
	}
}

func (ptr *QBoxPlotSeries) Pressed(boxset QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_Pressed(ptr.Pointer(), PointerFromQBoxSet(boxset))
	}
}

//export callbackQBoxPlotSeries_Released
func callbackQBoxPlotSeries_Released(ptr unsafe.Pointer, boxset unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::released"); signal != nil {
		signal.(func(*QBoxSet))(NewQBoxSetFromPointer(boxset))
	}

}

func (ptr *QBoxPlotSeries) ConnectReleased(f func(boxset *QBoxSet)) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::released", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::released")
	}
}

func (ptr *QBoxPlotSeries) Released(boxset QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_Released(ptr.Pointer(), PointerFromQBoxSet(boxset))
	}
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

func (ptr *QBoxPlotSeries) DestroyQBoxPlotSeries() {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries_DestroyQBoxPlotSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBoxPlotSeries_Type
func callbackQBoxPlotSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxPlotSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQBoxPlotSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QBoxPlotSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::type", f)
	}
}

func (ptr *QBoxPlotSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxPlotSeries::type")
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

func (ptr *QBoxPlotSeries) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QBoxPlotSeries_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QBoxPlotSeries) BoxSets() []*QBoxSet {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QBoxSet {
			var out = make([]*QBoxSet, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBoxPlotSeriesFromPointer(l.data).__boxSets_atList(i)
			}
			return out
		}(C.QBoxPlotSeries_BoxSets(ptr.Pointer()))
	}
	return make([]*QBoxSet, 0)
}

func (ptr *QBoxPlotSeries) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QBoxPlotSeries_Pen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QBoxPlotSeries) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBoxPlotSeries_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBoxPlotSeries) __append_sets_atList2(i int) *QBoxSet {
	if ptr.Pointer() != nil {
		return NewQBoxSetFromPointer(C.QBoxPlotSeries___append_sets_atList2(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QBoxPlotSeries) __append_sets_setList2(i QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries___append_sets_setList2(ptr.Pointer(), PointerFromQBoxSet(i))
	}
}

func (ptr *QBoxPlotSeries) __append_sets_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QBoxPlotSeries___append_sets_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBoxPlotSeries) __boxsetsAdded_sets_atList(i int) *QBoxSet {
	if ptr.Pointer() != nil {
		return NewQBoxSetFromPointer(C.QBoxPlotSeries___boxsetsAdded_sets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QBoxPlotSeries) __boxsetsAdded_sets_setList(i QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries___boxsetsAdded_sets_setList(ptr.Pointer(), PointerFromQBoxSet(i))
	}
}

func (ptr *QBoxPlotSeries) __boxsetsAdded_sets_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QBoxPlotSeries___boxsetsAdded_sets_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBoxPlotSeries) __boxsetsRemoved_sets_atList(i int) *QBoxSet {
	if ptr.Pointer() != nil {
		return NewQBoxSetFromPointer(C.QBoxPlotSeries___boxsetsRemoved_sets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QBoxPlotSeries) __boxsetsRemoved_sets_setList(i QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries___boxsetsRemoved_sets_setList(ptr.Pointer(), PointerFromQBoxSet(i))
	}
}

func (ptr *QBoxPlotSeries) __boxsetsRemoved_sets_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QBoxPlotSeries___boxsetsRemoved_sets_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBoxPlotSeries) __boxSets_atList(i int) *QBoxSet {
	if ptr.Pointer() != nil {
		return NewQBoxSetFromPointer(C.QBoxPlotSeries___boxSets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QBoxPlotSeries) __boxSets_setList(i QBoxSet_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxPlotSeries___boxSets_setList(ptr.Pointer(), PointerFromQBoxSet(i))
	}
}

func (ptr *QBoxPlotSeries) __boxSets_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QBoxPlotSeries___boxSets_newList(ptr.Pointer()))
	}
	return nil
}

type QBoxSet struct {
	ptr unsafe.Pointer
}

type QBoxSet_ITF interface {
	QBoxSet_PTR() *QBoxSet
}

func (ptr *QBoxSet) QBoxSet_PTR() *QBoxSet {
	return ptr
}

func (ptr *QBoxSet) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBoxSet) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBoxSet(ptr QBoxSet_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxSet_PTR().Pointer()
	}
	return nil
}

func NewQBoxSetFromPointer(ptr unsafe.Pointer) *QBoxSet {
	var n = new(QBoxSet)
	n.SetPointer(ptr)
	return n
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
	var labelC = C.CString(label)
	defer C.free(unsafe.Pointer(labelC))
	return NewQBoxSetFromPointer(C.QBoxSet_NewQBoxSet(labelC, core.PointerFromQObject(parent)))
}

func NewQBoxSet2(le float64, lq float64, m float64, uq float64, ue float64, label string, parent core.QObject_ITF) *QBoxSet {
	var labelC = C.CString(label)
	defer C.free(unsafe.Pointer(labelC))
	return NewQBoxSetFromPointer(C.QBoxSet_NewQBoxSet2(C.double(le), C.double(lq), C.double(m), C.double(uq), C.double(ue), labelC, core.PointerFromQObject(parent)))
}

func (ptr *QBoxSet) Append2(values []float64) {
	if ptr.Pointer() != nil {
		C.QBoxSet_Append2(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQBoxSetFromPointer(NewQBoxSetFromPointer(unsafe.Pointer(uintptr(1))).__append_values_newList2())
			for _, v := range values {
				tmpList.__append_values_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QBoxSet) Append(value float64) {
	if ptr.Pointer() != nil {
		C.QBoxSet_Append(ptr.Pointer(), C.double(value))
	}
}

//export callbackQBoxSet_BrushChanged
func callbackQBoxSet_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::brushChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxSet) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::brushChanged", f)
	}
}

func (ptr *QBoxSet) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::brushChanged")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::cleared"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxSet) ConnectCleared(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectCleared(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::cleared", f)
	}
}

func (ptr *QBoxSet) DisconnectCleared() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectCleared(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::cleared")
	}
}

func (ptr *QBoxSet) Cleared() {
	if ptr.Pointer() != nil {
		C.QBoxSet_Cleared(ptr.Pointer())
	}
}

//export callbackQBoxSet_Clicked
func callbackQBoxSet_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::clicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxSet) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::clicked", f)
	}
}

func (ptr *QBoxSet) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::clicked")
	}
}

func (ptr *QBoxSet) Clicked() {
	if ptr.Pointer() != nil {
		C.QBoxSet_Clicked(ptr.Pointer())
	}
}

//export callbackQBoxSet_DoubleClicked
func callbackQBoxSet_DoubleClicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::doubleClicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxSet) ConnectDoubleClicked(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::doubleClicked", f)
	}
}

func (ptr *QBoxSet) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::doubleClicked")
	}
}

func (ptr *QBoxSet) DoubleClicked() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DoubleClicked(ptr.Pointer())
	}
}

//export callbackQBoxSet_Hovered
func callbackQBoxSet_Hovered(ptr unsafe.Pointer, status C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::hovered"); signal != nil {
		signal.(func(bool))(int8(status) != 0)
	}

}

func (ptr *QBoxSet) ConnectHovered(f func(status bool)) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::hovered", f)
	}
}

func (ptr *QBoxSet) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::hovered")
	}
}

func (ptr *QBoxSet) Hovered(status bool) {
	if ptr.Pointer() != nil {
		C.QBoxSet_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))))
	}
}

//export callbackQBoxSet_PenChanged
func callbackQBoxSet_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::penChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxSet) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectPenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::penChanged", f)
	}
}

func (ptr *QBoxSet) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::penChanged")
	}
}

func (ptr *QBoxSet) PenChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_PenChanged(ptr.Pointer())
	}
}

//export callbackQBoxSet_Pressed
func callbackQBoxSet_Pressed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::pressed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxSet) ConnectPressed(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::pressed", f)
	}
}

func (ptr *QBoxSet) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::pressed")
	}
}

func (ptr *QBoxSet) Pressed() {
	if ptr.Pointer() != nil {
		C.QBoxSet_Pressed(ptr.Pointer())
	}
}

//export callbackQBoxSet_Released
func callbackQBoxSet_Released(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::released"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxSet) ConnectReleased(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::released", f)
	}
}

func (ptr *QBoxSet) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::released")
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
		var labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
		C.QBoxSet_SetLabel(ptr.Pointer(), labelC)
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::valueChanged"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QBoxSet) ConnectValueChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::valueChanged", f)
	}
}

func (ptr *QBoxSet) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::valueChanged")
	}
}

func (ptr *QBoxSet) ValueChanged(index int) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ValueChanged(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQBoxSet_ValuesChanged
func callbackQBoxSet_ValuesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::valuesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBoxSet) ConnectValuesChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QBoxSet_ConnectValuesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::valuesChanged", f)
	}
}

func (ptr *QBoxSet) DisconnectValuesChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DisconnectValuesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::valuesChanged")
	}
}

func (ptr *QBoxSet) ValuesChanged() {
	if ptr.Pointer() != nil {
		C.QBoxSet_ValuesChanged(ptr.Pointer())
	}
}

//export callbackQBoxSet_DestroyQBoxSet
func callbackQBoxSet_DestroyQBoxSet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBoxSet::~QBoxSet"); signal != nil {
		signal.(func())()
	} else {
		NewQBoxSetFromPointer(ptr).DestroyQBoxSetDefault()
	}
}

func (ptr *QBoxSet) ConnectDestroyQBoxSet(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::~QBoxSet", f)
	}
}

func (ptr *QBoxSet) DisconnectDestroyQBoxSet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBoxSet::~QBoxSet")
	}
}

func (ptr *QBoxSet) DestroyQBoxSet() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DestroyQBoxSet(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBoxSet) DestroyQBoxSetDefault() {
	if ptr.Pointer() != nil {
		C.QBoxSet_DestroyQBoxSetDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBoxSet) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QBoxSet_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QBoxSet) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QBoxSet_Pen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QBoxSet) Label() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBoxSet_Label(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBoxSet) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBoxSet_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBoxSet) At(index int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QBoxSet_At(ptr.Pointer(), C.int(int32(index))))
	}
	return 0
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QBoxSet___append_values_newList2(ptr.Pointer()))
	}
	return nil
}

type QCandlestickLegendMarker struct {
	ptr unsafe.Pointer
}

type QCandlestickLegendMarker_ITF interface {
	QCandlestickLegendMarker_PTR() *QCandlestickLegendMarker
}

func (ptr *QCandlestickLegendMarker) QCandlestickLegendMarker_PTR() *QCandlestickLegendMarker {
	return ptr
}

func (ptr *QCandlestickLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCandlestickLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCandlestickLegendMarker(ptr QCandlestickLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCandlestickLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQCandlestickLegendMarkerFromPointer(ptr unsafe.Pointer) *QCandlestickLegendMarker {
	var n = new(QCandlestickLegendMarker)
	n.SetPointer(ptr)
	return n
}

//export callbackQCandlestickLegendMarker_Series
func callbackQCandlestickLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickLegendMarker::series"); signal != nil {
		return PointerFromQCandlestickSeries(signal.(func() *QCandlestickSeries)())
	}

	return PointerFromQCandlestickSeries(NewQCandlestickLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QCandlestickLegendMarker) ConnectSeries(f func() *QCandlestickSeries) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickLegendMarker::series", f)
	}
}

func (ptr *QCandlestickLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickLegendMarker::series")
	}
}

func (ptr *QCandlestickLegendMarker) Series() *QCandlestickSeries {
	if ptr.Pointer() != nil {
		return NewQCandlestickSeriesFromPointer(C.QCandlestickLegendMarker_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCandlestickLegendMarker) SeriesDefault() *QCandlestickSeries {
	if ptr.Pointer() != nil {
		return NewQCandlestickSeriesFromPointer(C.QCandlestickLegendMarker_SeriesDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQCandlestickLegendMarker_DestroyQCandlestickLegendMarker
func callbackQCandlestickLegendMarker_DestroyQCandlestickLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickLegendMarker::~QCandlestickLegendMarker"); signal != nil {
		signal.(func())()
	} else {
		NewQCandlestickLegendMarkerFromPointer(ptr).DestroyQCandlestickLegendMarkerDefault()
	}
}

func (ptr *QCandlestickLegendMarker) ConnectDestroyQCandlestickLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickLegendMarker::~QCandlestickLegendMarker", f)
	}
}

func (ptr *QCandlestickLegendMarker) DisconnectDestroyQCandlestickLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickLegendMarker::~QCandlestickLegendMarker")
	}
}

func (ptr *QCandlestickLegendMarker) DestroyQCandlestickLegendMarker() {
	if ptr.Pointer() != nil {
		C.QCandlestickLegendMarker_DestroyQCandlestickLegendMarker(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCandlestickLegendMarker) DestroyQCandlestickLegendMarkerDefault() {
	if ptr.Pointer() != nil {
		C.QCandlestickLegendMarker_DestroyQCandlestickLegendMarkerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QCandlestickModelMapper struct {
	ptr unsafe.Pointer
}

type QCandlestickModelMapper_ITF interface {
	QCandlestickModelMapper_PTR() *QCandlestickModelMapper
}

func (ptr *QCandlestickModelMapper) QCandlestickModelMapper_PTR() *QCandlestickModelMapper {
	return ptr
}

func (ptr *QCandlestickModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCandlestickModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCandlestickModelMapper(ptr QCandlestickModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCandlestickModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQCandlestickModelMapperFromPointer(ptr unsafe.Pointer) *QCandlestickModelMapper {
	var n = new(QCandlestickModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCandlestickModelMapper) DestroyQCandlestickModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQCandlestickModelMapper(parent core.QObject_ITF) *QCandlestickModelMapper {
	return NewQCandlestickModelMapperFromPointer(C.QCandlestickModelMapper_NewQCandlestickModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQCandlestickModelMapper_ModelReplaced
func callbackQCandlestickModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickModelMapper::modelReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_ConnectModelReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickModelMapper::modelReplaced", f)
	}
}

func (ptr *QCandlestickModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickModelMapper::modelReplaced")
	}
}

func (ptr *QCandlestickModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_ModelReplaced(ptr.Pointer())
	}
}

//export callbackQCandlestickModelMapper_SeriesReplaced
func callbackQCandlestickModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickModelMapper::seriesReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_ConnectSeriesReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickModelMapper::seriesReplaced", f)
	}
}

func (ptr *QCandlestickModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QCandlestickModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickModelMapper::seriesReplaced")
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

func (ptr *QCandlestickModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QCandlestickModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickModelMapper) Series() *QCandlestickSeries {
	if ptr.Pointer() != nil {
		return NewQCandlestickSeriesFromPointer(C.QCandlestickModelMapper_Series(ptr.Pointer()))
	}
	return nil
}

//export callbackQCandlestickModelMapper_Orientation
func callbackQCandlestickModelMapper_Orientation(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickModelMapper::orientation"); signal != nil {
		return C.longlong(signal.(func() core.Qt__Orientation)())
	}

	return C.longlong(0)
}

func (ptr *QCandlestickModelMapper) ConnectOrientation(f func() core.Qt__Orientation) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickModelMapper::orientation", f)
	}
}

func (ptr *QCandlestickModelMapper) DisconnectOrientation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickModelMapper::orientation")
	}
}

func (ptr *QCandlestickModelMapper) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QCandlestickModelMapper_Orientation(ptr.Pointer()))
	}
	return 0
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

func (ptr *QCandlestickModelMapper) Open() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickModelMapper_Open(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCandlestickModelMapper) Timestamp() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickModelMapper_Timestamp(ptr.Pointer())))
	}
	return 0
}

type QCandlestickSeries struct {
	ptr unsafe.Pointer
}

type QCandlestickSeries_ITF interface {
	QCandlestickSeries_PTR() *QCandlestickSeries
}

func (ptr *QCandlestickSeries) QCandlestickSeries_PTR() *QCandlestickSeries {
	return ptr
}

func (ptr *QCandlestickSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCandlestickSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCandlestickSeries(ptr QCandlestickSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCandlestickSeries_PTR().Pointer()
	}
	return nil
}

func NewQCandlestickSeriesFromPointer(ptr unsafe.Pointer) *QCandlestickSeries {
	var n = new(QCandlestickSeries)
	n.SetPointer(ptr)
	return n
}
func NewQCandlestickSeries(parent core.QObject_ITF) *QCandlestickSeries {
	return NewQCandlestickSeriesFromPointer(C.QCandlestickSeries_NewQCandlestickSeries(core.PointerFromQObject(parent)))
}

func (ptr *QCandlestickSeries) Append(set QCandlestickSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCandlestickSeries_Append(ptr.Pointer(), PointerFromQCandlestickSet(set)) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) Append2(sets []*QCandlestickSet) bool {
	if ptr.Pointer() != nil {
		return C.QCandlestickSeries_Append2(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQCandlestickSeriesFromPointer(NewQCandlestickSeriesFromPointer(unsafe.Pointer(uintptr(1))).__append_sets_newList2())
			for _, v := range sets {
				tmpList.__append_sets_setList2(v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) Insert(index int, set QCandlestickSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCandlestickSeries_Insert(ptr.Pointer(), C.int(int32(index)), PointerFromQCandlestickSet(set)) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) Remove(set QCandlestickSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCandlestickSeries_Remove(ptr.Pointer(), PointerFromQCandlestickSet(set)) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) Remove2(sets []*QCandlestickSet) bool {
	if ptr.Pointer() != nil {
		return C.QCandlestickSeries_Remove2(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQCandlestickSeriesFromPointer(NewQCandlestickSeriesFromPointer(unsafe.Pointer(uintptr(1))).__remove_sets_newList2())
			for _, v := range sets {
				tmpList.__remove_sets_setList2(v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) Take(set QCandlestickSet_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCandlestickSeries_Take(ptr.Pointer(), PointerFromQCandlestickSet(set)) != 0
	}
	return false
}

//export callbackQCandlestickSeries_BodyOutlineVisibilityChanged
func callbackQCandlestickSeries_BodyOutlineVisibilityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::bodyOutlineVisibilityChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectBodyOutlineVisibilityChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectBodyOutlineVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::bodyOutlineVisibilityChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectBodyOutlineVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectBodyOutlineVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::bodyOutlineVisibilityChanged")
	}
}

func (ptr *QCandlestickSeries) BodyOutlineVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_BodyOutlineVisibilityChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_BodyWidthChanged
func callbackQCandlestickSeries_BodyWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::bodyWidthChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectBodyWidthChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectBodyWidthChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::bodyWidthChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectBodyWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectBodyWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::bodyWidthChanged")
	}
}

func (ptr *QCandlestickSeries) BodyWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_BodyWidthChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_BrushChanged
func callbackQCandlestickSeries_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::brushChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::brushChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::brushChanged")
	}
}

func (ptr *QCandlestickSeries) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_BrushChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_CandlestickSetsAdded
func callbackQCandlestickSeries_CandlestickSetsAdded(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::candlestickSetsAdded"); signal != nil {
		signal.(func([]*QCandlestickSet))(func(l C.struct_QtCharts_PackedList) []*QCandlestickSet {
			var out = make([]*QCandlestickSet, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQCandlestickSeriesFromPointer(l.data).__candlestickSetsAdded_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QCandlestickSeries) ConnectCandlestickSetsAdded(f func(sets []*QCandlestickSet)) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectCandlestickSetsAdded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::candlestickSetsAdded", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectCandlestickSetsAdded() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectCandlestickSetsAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::candlestickSetsAdded")
	}
}

func (ptr *QCandlestickSeries) CandlestickSetsAdded(sets []*QCandlestickSet) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_CandlestickSetsAdded(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQCandlestickSeriesFromPointer(NewQCandlestickSeriesFromPointer(unsafe.Pointer(uintptr(1))).__candlestickSetsAdded_sets_newList())
			for _, v := range sets {
				tmpList.__candlestickSetsAdded_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQCandlestickSeries_CandlestickSetsRemoved
func callbackQCandlestickSeries_CandlestickSetsRemoved(ptr unsafe.Pointer, sets C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::candlestickSetsRemoved"); signal != nil {
		signal.(func([]*QCandlestickSet))(func(l C.struct_QtCharts_PackedList) []*QCandlestickSet {
			var out = make([]*QCandlestickSet, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQCandlestickSeriesFromPointer(l.data).__candlestickSetsRemoved_sets_atList(i)
			}
			return out
		}(sets))
	}

}

func (ptr *QCandlestickSeries) ConnectCandlestickSetsRemoved(f func(sets []*QCandlestickSet)) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectCandlestickSetsRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::candlestickSetsRemoved", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectCandlestickSetsRemoved() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectCandlestickSetsRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::candlestickSetsRemoved")
	}
}

func (ptr *QCandlestickSeries) CandlestickSetsRemoved(sets []*QCandlestickSet) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_CandlestickSetsRemoved(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQCandlestickSeriesFromPointer(NewQCandlestickSeriesFromPointer(unsafe.Pointer(uintptr(1))).__candlestickSetsRemoved_sets_newList())
			for _, v := range sets {
				tmpList.__candlestickSetsRemoved_sets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQCandlestickSeries_CapsVisibilityChanged
func callbackQCandlestickSeries_CapsVisibilityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::capsVisibilityChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectCapsVisibilityChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectCapsVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::capsVisibilityChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectCapsVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectCapsVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::capsVisibilityChanged")
	}
}

func (ptr *QCandlestickSeries) CapsVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_CapsVisibilityChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_CapsWidthChanged
func callbackQCandlestickSeries_CapsWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::capsWidthChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectCapsWidthChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectCapsWidthChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::capsWidthChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectCapsWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectCapsWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::capsWidthChanged")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::clicked"); signal != nil {
		signal.(func(*QCandlestickSet))(NewQCandlestickSetFromPointer(set))
	}

}

func (ptr *QCandlestickSeries) ConnectClicked(f func(set *QCandlestickSet)) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::clicked", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::clicked")
	}
}

func (ptr *QCandlestickSeries) Clicked(set QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_Clicked(ptr.Pointer(), PointerFromQCandlestickSet(set))
	}
}

//export callbackQCandlestickSeries_CountChanged
func callbackQCandlestickSeries_CountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::countChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::countChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectCountChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::countChanged")
	}
}

func (ptr *QCandlestickSeries) CountChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_CountChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_DecreasingColorChanged
func callbackQCandlestickSeries_DecreasingColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::decreasingColorChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectDecreasingColorChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectDecreasingColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::decreasingColorChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectDecreasingColorChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectDecreasingColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::decreasingColorChanged")
	}
}

func (ptr *QCandlestickSeries) DecreasingColorChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DecreasingColorChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_DoubleClicked
func callbackQCandlestickSeries_DoubleClicked(ptr unsafe.Pointer, set unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::doubleClicked"); signal != nil {
		signal.(func(*QCandlestickSet))(NewQCandlestickSetFromPointer(set))
	}

}

func (ptr *QCandlestickSeries) ConnectDoubleClicked(f func(set *QCandlestickSet)) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::doubleClicked", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::doubleClicked")
	}
}

func (ptr *QCandlestickSeries) DoubleClicked(set QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DoubleClicked(ptr.Pointer(), PointerFromQCandlestickSet(set))
	}
}

//export callbackQCandlestickSeries_Hovered
func callbackQCandlestickSeries_Hovered(ptr unsafe.Pointer, status C.char, set unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::hovered"); signal != nil {
		signal.(func(bool, *QCandlestickSet))(int8(status) != 0, NewQCandlestickSetFromPointer(set))
	}

}

func (ptr *QCandlestickSeries) ConnectHovered(f func(status bool, set *QCandlestickSet)) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::hovered", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::hovered")
	}
}

func (ptr *QCandlestickSeries) Hovered(status bool, set QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))), PointerFromQCandlestickSet(set))
	}
}

//export callbackQCandlestickSeries_IncreasingColorChanged
func callbackQCandlestickSeries_IncreasingColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::increasingColorChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectIncreasingColorChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectIncreasingColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::increasingColorChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectIncreasingColorChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectIncreasingColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::increasingColorChanged")
	}
}

func (ptr *QCandlestickSeries) IncreasingColorChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_IncreasingColorChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_MaximumColumnWidthChanged
func callbackQCandlestickSeries_MaximumColumnWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::maximumColumnWidthChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectMaximumColumnWidthChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectMaximumColumnWidthChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::maximumColumnWidthChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectMaximumColumnWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectMaximumColumnWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::maximumColumnWidthChanged")
	}
}

func (ptr *QCandlestickSeries) MaximumColumnWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_MaximumColumnWidthChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_MinimumColumnWidthChanged
func callbackQCandlestickSeries_MinimumColumnWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::minimumColumnWidthChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectMinimumColumnWidthChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectMinimumColumnWidthChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::minimumColumnWidthChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectMinimumColumnWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectMinimumColumnWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::minimumColumnWidthChanged")
	}
}

func (ptr *QCandlestickSeries) MinimumColumnWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_MinimumColumnWidthChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_PenChanged
func callbackQCandlestickSeries_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::penChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSeries) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectPenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::penChanged", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::penChanged")
	}
}

func (ptr *QCandlestickSeries) PenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_PenChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSeries_Pressed
func callbackQCandlestickSeries_Pressed(ptr unsafe.Pointer, set unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::pressed"); signal != nil {
		signal.(func(*QCandlestickSet))(NewQCandlestickSetFromPointer(set))
	}

}

func (ptr *QCandlestickSeries) ConnectPressed(f func(set *QCandlestickSet)) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::pressed", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::pressed")
	}
}

func (ptr *QCandlestickSeries) Pressed(set QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_Pressed(ptr.Pointer(), PointerFromQCandlestickSet(set))
	}
}

//export callbackQCandlestickSeries_Released
func callbackQCandlestickSeries_Released(ptr unsafe.Pointer, set unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::released"); signal != nil {
		signal.(func(*QCandlestickSet))(NewQCandlestickSetFromPointer(set))
	}

}

func (ptr *QCandlestickSeries) ConnectReleased(f func(set *QCandlestickSet)) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::released", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::released")
	}
}

func (ptr *QCandlestickSeries) Released(set QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_Released(ptr.Pointer(), PointerFromQCandlestickSet(set))
	}
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

func (ptr *QCandlestickSeries) DestroyQCandlestickSeries() {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries_DestroyQCandlestickSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQCandlestickSeries_Type
func callbackQCandlestickSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQCandlestickSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QCandlestickSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::type", f)
	}
}

func (ptr *QCandlestickSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSeries::type")
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

func (ptr *QCandlestickSeries) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QCandlestickSeries_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSeries) DecreasingColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QCandlestickSeries_DecreasingColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSeries) IncreasingColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QCandlestickSeries_IncreasingColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSeries) Sets() []*QCandlestickSet {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QCandlestickSet {
			var out = make([]*QCandlestickSet, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQCandlestickSeriesFromPointer(l.data).__sets_atList(i)
			}
			return out
		}(C.QCandlestickSeries_Sets(ptr.Pointer()))
	}
	return make([]*QCandlestickSet, 0)
}

func (ptr *QCandlestickSeries) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QCandlestickSeries_Pen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSeries) BodyOutlineVisible() bool {
	if ptr.Pointer() != nil {
		return C.QCandlestickSeries_BodyOutlineVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) CapsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QCandlestickSeries_CapsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCandlestickSeries) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCandlestickSeries_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCandlestickSeries) BodyWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSeries_BodyWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickSeries) CapsWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSeries_CapsWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickSeries) MaximumColumnWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSeries_MaximumColumnWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickSeries) MinimumColumnWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSeries_MinimumColumnWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickSeries) __append_sets_atList2(i int) *QCandlestickSet {
	if ptr.Pointer() != nil {
		return NewQCandlestickSetFromPointer(C.QCandlestickSeries___append_sets_atList2(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QCandlestickSeries) __append_sets_setList2(i QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries___append_sets_setList2(ptr.Pointer(), PointerFromQCandlestickSet(i))
	}
}

func (ptr *QCandlestickSeries) __append_sets_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QCandlestickSeries___append_sets_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCandlestickSeries) __remove_sets_atList2(i int) *QCandlestickSet {
	if ptr.Pointer() != nil {
		return NewQCandlestickSetFromPointer(C.QCandlestickSeries___remove_sets_atList2(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QCandlestickSeries) __remove_sets_setList2(i QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries___remove_sets_setList2(ptr.Pointer(), PointerFromQCandlestickSet(i))
	}
}

func (ptr *QCandlestickSeries) __remove_sets_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QCandlestickSeries___remove_sets_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCandlestickSeries) __candlestickSetsAdded_sets_atList(i int) *QCandlestickSet {
	if ptr.Pointer() != nil {
		return NewQCandlestickSetFromPointer(C.QCandlestickSeries___candlestickSetsAdded_sets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QCandlestickSeries) __candlestickSetsAdded_sets_setList(i QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries___candlestickSetsAdded_sets_setList(ptr.Pointer(), PointerFromQCandlestickSet(i))
	}
}

func (ptr *QCandlestickSeries) __candlestickSetsAdded_sets_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QCandlestickSeries___candlestickSetsAdded_sets_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCandlestickSeries) __candlestickSetsRemoved_sets_atList(i int) *QCandlestickSet {
	if ptr.Pointer() != nil {
		return NewQCandlestickSetFromPointer(C.QCandlestickSeries___candlestickSetsRemoved_sets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QCandlestickSeries) __candlestickSetsRemoved_sets_setList(i QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries___candlestickSetsRemoved_sets_setList(ptr.Pointer(), PointerFromQCandlestickSet(i))
	}
}

func (ptr *QCandlestickSeries) __candlestickSetsRemoved_sets_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QCandlestickSeries___candlestickSetsRemoved_sets_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCandlestickSeries) __sets_atList(i int) *QCandlestickSet {
	if ptr.Pointer() != nil {
		return NewQCandlestickSetFromPointer(C.QCandlestickSeries___sets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QCandlestickSeries) __sets_setList(i QCandlestickSet_ITF) {
	if ptr.Pointer() != nil {
		C.QCandlestickSeries___sets_setList(ptr.Pointer(), PointerFromQCandlestickSet(i))
	}
}

func (ptr *QCandlestickSeries) __sets_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QCandlestickSeries___sets_newList(ptr.Pointer()))
	}
	return nil
}

type QCandlestickSet struct {
	ptr unsafe.Pointer
}

type QCandlestickSet_ITF interface {
	QCandlestickSet_PTR() *QCandlestickSet
}

func (ptr *QCandlestickSet) QCandlestickSet_PTR() *QCandlestickSet {
	return ptr
}

func (ptr *QCandlestickSet) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCandlestickSet) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCandlestickSet(ptr QCandlestickSet_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCandlestickSet_PTR().Pointer()
	}
	return nil
}

func NewQCandlestickSetFromPointer(ptr unsafe.Pointer) *QCandlestickSet {
	var n = new(QCandlestickSet)
	n.SetPointer(ptr)
	return n
}
func NewQCandlestickSet2(open float64, high float64, low float64, close float64, timestamp float64, parent core.QObject_ITF) *QCandlestickSet {
	return NewQCandlestickSetFromPointer(C.QCandlestickSet_NewQCandlestickSet2(C.double(open), C.double(high), C.double(low), C.double(close), C.double(timestamp), core.PointerFromQObject(parent)))
}

func NewQCandlestickSet(timestamp float64, parent core.QObject_ITF) *QCandlestickSet {
	return NewQCandlestickSetFromPointer(C.QCandlestickSet_NewQCandlestickSet(C.double(timestamp), core.PointerFromQObject(parent)))
}

//export callbackQCandlestickSet_BrushChanged
func callbackQCandlestickSet_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::brushChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::brushChanged", f)
	}
}

func (ptr *QCandlestickSet) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::brushChanged")
	}
}

func (ptr *QCandlestickSet) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_BrushChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_Clicked
func callbackQCandlestickSet_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::clicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::clicked", f)
	}
}

func (ptr *QCandlestickSet) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::clicked")
	}
}

func (ptr *QCandlestickSet) Clicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_Clicked(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_CloseChanged
func callbackQCandlestickSet_CloseChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::closeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectCloseChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectCloseChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::closeChanged", f)
	}
}

func (ptr *QCandlestickSet) DisconnectCloseChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectCloseChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::closeChanged")
	}
}

func (ptr *QCandlestickSet) CloseChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_CloseChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_DoubleClicked
func callbackQCandlestickSet_DoubleClicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::doubleClicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectDoubleClicked(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::doubleClicked", f)
	}
}

func (ptr *QCandlestickSet) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::doubleClicked")
	}
}

func (ptr *QCandlestickSet) DoubleClicked() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DoubleClicked(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_HighChanged
func callbackQCandlestickSet_HighChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::highChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectHighChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectHighChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::highChanged", f)
	}
}

func (ptr *QCandlestickSet) DisconnectHighChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectHighChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::highChanged")
	}
}

func (ptr *QCandlestickSet) HighChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_HighChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_Hovered
func callbackQCandlestickSet_Hovered(ptr unsafe.Pointer, status C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::hovered"); signal != nil {
		signal.(func(bool))(int8(status) != 0)
	}

}

func (ptr *QCandlestickSet) ConnectHovered(f func(status bool)) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::hovered", f)
	}
}

func (ptr *QCandlestickSet) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::hovered")
	}
}

func (ptr *QCandlestickSet) Hovered(status bool) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))))
	}
}

//export callbackQCandlestickSet_LowChanged
func callbackQCandlestickSet_LowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::lowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectLowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectLowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::lowChanged", f)
	}
}

func (ptr *QCandlestickSet) DisconnectLowChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectLowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::lowChanged")
	}
}

func (ptr *QCandlestickSet) LowChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_LowChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_OpenChanged
func callbackQCandlestickSet_OpenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::openChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectOpenChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectOpenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::openChanged", f)
	}
}

func (ptr *QCandlestickSet) DisconnectOpenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectOpenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::openChanged")
	}
}

func (ptr *QCandlestickSet) OpenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_OpenChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_PenChanged
func callbackQCandlestickSet_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::penChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectPenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::penChanged", f)
	}
}

func (ptr *QCandlestickSet) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::penChanged")
	}
}

func (ptr *QCandlestickSet) PenChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_PenChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_Pressed
func callbackQCandlestickSet_Pressed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::pressed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectPressed(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::pressed", f)
	}
}

func (ptr *QCandlestickSet) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::pressed")
	}
}

func (ptr *QCandlestickSet) Pressed() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_Pressed(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_Released
func callbackQCandlestickSet_Released(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::released"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectReleased(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::released", f)
	}
}

func (ptr *QCandlestickSet) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::released")
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

//export callbackQCandlestickSet_TimestampChanged
func callbackQCandlestickSet_TimestampChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::timestampChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCandlestickSet) ConnectTimestampChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_ConnectTimestampChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::timestampChanged", f)
	}
}

func (ptr *QCandlestickSet) DisconnectTimestampChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DisconnectTimestampChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::timestampChanged")
	}
}

func (ptr *QCandlestickSet) TimestampChanged() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_TimestampChanged(ptr.Pointer())
	}
}

//export callbackQCandlestickSet_DestroyQCandlestickSet
func callbackQCandlestickSet_DestroyQCandlestickSet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCandlestickSet::~QCandlestickSet"); signal != nil {
		signal.(func())()
	} else {
		NewQCandlestickSetFromPointer(ptr).DestroyQCandlestickSetDefault()
	}
}

func (ptr *QCandlestickSet) ConnectDestroyQCandlestickSet(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::~QCandlestickSet", f)
	}
}

func (ptr *QCandlestickSet) DisconnectDestroyQCandlestickSet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCandlestickSet::~QCandlestickSet")
	}
}

func (ptr *QCandlestickSet) DestroyQCandlestickSet() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DestroyQCandlestickSet(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCandlestickSet) DestroyQCandlestickSetDefault() {
	if ptr.Pointer() != nil {
		C.QCandlestickSet_DestroyQCandlestickSetDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCandlestickSet) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QCandlestickSet_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSet) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QCandlestickSet_Pen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QCandlestickSet) Close() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSet_Close(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickSet) High() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSet_High(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickSet) Low() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSet_Low(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickSet) Open() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSet_Open(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCandlestickSet) Timestamp() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCandlestickSet_Timestamp(ptr.Pointer()))
	}
	return 0
}

type QCategoryAxis struct {
	ptr unsafe.Pointer
}

type QCategoryAxis_ITF interface {
	QCategoryAxis_PTR() *QCategoryAxis
}

func (ptr *QCategoryAxis) QCategoryAxis_PTR() *QCategoryAxis {
	return ptr
}

func (ptr *QCategoryAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCategoryAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCategoryAxis(ptr QCategoryAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCategoryAxis_PTR().Pointer()
	}
	return nil
}

func NewQCategoryAxisFromPointer(ptr unsafe.Pointer) *QCategoryAxis {
	var n = new(QCategoryAxis)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QCategoryAxis__AxisLabelsPosition
//QCategoryAxis::AxisLabelsPosition
type QCategoryAxis__AxisLabelsPosition int64

const (
	QCategoryAxis__AxisLabelsPositionCenter  QCategoryAxis__AxisLabelsPosition = QCategoryAxis__AxisLabelsPosition(0x0)
	QCategoryAxis__AxisLabelsPositionOnValue QCategoryAxis__AxisLabelsPosition = QCategoryAxis__AxisLabelsPosition(0x1)
)

func NewQCategoryAxis(parent core.QObject_ITF) *QCategoryAxis {
	return NewQCategoryAxisFromPointer(C.QCategoryAxis_NewQCategoryAxis(core.PointerFromQObject(parent)))
}

func (ptr *QCategoryAxis) CategoriesLabels() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QCategoryAxis_CategoriesLabels(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QCategoryAxis) Append(categoryLabel string, categoryEndValue float64) {
	if ptr.Pointer() != nil {
		var categoryLabelC = C.CString(categoryLabel)
		defer C.free(unsafe.Pointer(categoryLabelC))
		C.QCategoryAxis_Append(ptr.Pointer(), categoryLabelC, C.double(categoryEndValue))
	}
}

//export callbackQCategoryAxis_CategoriesChanged
func callbackQCategoryAxis_CategoriesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCategoryAxis::categoriesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCategoryAxis) ConnectCategoriesChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_ConnectCategoriesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCategoryAxis::categoriesChanged", f)
	}
}

func (ptr *QCategoryAxis) DisconnectCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_DisconnectCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCategoryAxis::categoriesChanged")
	}
}

func (ptr *QCategoryAxis) CategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_CategoriesChanged(ptr.Pointer())
	}
}

//export callbackQCategoryAxis_LabelsPositionChanged
func callbackQCategoryAxis_LabelsPositionChanged(ptr unsafe.Pointer, position C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCategoryAxis::labelsPositionChanged"); signal != nil {
		signal.(func(QCategoryAxis__AxisLabelsPosition))(QCategoryAxis__AxisLabelsPosition(position))
	}

}

func (ptr *QCategoryAxis) ConnectLabelsPositionChanged(f func(position QCategoryAxis__AxisLabelsPosition)) {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_ConnectLabelsPositionChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCategoryAxis::labelsPositionChanged", f)
	}
}

func (ptr *QCategoryAxis) DisconnectLabelsPositionChanged() {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_DisconnectLabelsPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCategoryAxis::labelsPositionChanged")
	}
}

func (ptr *QCategoryAxis) LabelsPositionChanged(position QCategoryAxis__AxisLabelsPosition) {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_LabelsPositionChanged(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QCategoryAxis) Remove(categoryLabel string) {
	if ptr.Pointer() != nil {
		var categoryLabelC = C.CString(categoryLabel)
		defer C.free(unsafe.Pointer(categoryLabelC))
		C.QCategoryAxis_Remove(ptr.Pointer(), categoryLabelC)
	}
}

func (ptr *QCategoryAxis) ReplaceLabel(oldLabel string, newLabel string) {
	if ptr.Pointer() != nil {
		var oldLabelC = C.CString(oldLabel)
		defer C.free(unsafe.Pointer(oldLabelC))
		var newLabelC = C.CString(newLabel)
		defer C.free(unsafe.Pointer(newLabelC))
		C.QCategoryAxis_ReplaceLabel(ptr.Pointer(), oldLabelC, newLabelC)
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

func (ptr *QCategoryAxis) DestroyQCategoryAxis() {
	if ptr.Pointer() != nil {
		C.QCategoryAxis_DestroyQCategoryAxis(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCategoryAxis) LabelsPosition() QCategoryAxis__AxisLabelsPosition {
	if ptr.Pointer() != nil {
		return QCategoryAxis__AxisLabelsPosition(C.QCategoryAxis_LabelsPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCategoryAxis) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCategoryAxis_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCategoryAxis) EndValue(categoryLabel string) float64 {
	if ptr.Pointer() != nil {
		var categoryLabelC = C.CString(categoryLabel)
		defer C.free(unsafe.Pointer(categoryLabelC))
		return float64(C.QCategoryAxis_EndValue(ptr.Pointer(), categoryLabelC))
	}
	return 0
}

func (ptr *QCategoryAxis) StartValue(categoryLabel string) float64 {
	if ptr.Pointer() != nil {
		var categoryLabelC = C.CString(categoryLabel)
		defer C.free(unsafe.Pointer(categoryLabelC))
		return float64(C.QCategoryAxis_StartValue(ptr.Pointer(), categoryLabelC))
	}
	return 0
}

type QChart struct {
	ptr unsafe.Pointer
}

type QChart_ITF interface {
	QChart_PTR() *QChart
}

func (ptr *QChart) QChart_PTR() *QChart {
	return ptr
}

func (ptr *QChart) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QChart) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQChart(ptr QChart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QChart_PTR().Pointer()
	}
	return nil
}

func NewQChartFromPointer(ptr unsafe.Pointer) *QChart {
	var n = new(QChart)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QChart__AnimationOption
//QChart::AnimationOption
type QChart__AnimationOption int64

const (
	QChart__NoAnimation        QChart__AnimationOption = QChart__AnimationOption(0x0)
	QChart__GridAxisAnimations QChart__AnimationOption = QChart__AnimationOption(0x1)
	QChart__SeriesAnimations   QChart__AnimationOption = QChart__AnimationOption(0x2)
	QChart__AllAnimations      QChart__AnimationOption = QChart__AnimationOption(0x3)
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

//go:generate stringer -type=QChart__ChartType
//QChart::ChartType
type QChart__ChartType int64

const (
	QChart__ChartTypeUndefined QChart__ChartType = QChart__ChartType(0)
	QChart__ChartTypeCartesian QChart__ChartType = QChart__ChartType(1)
	QChart__ChartTypePolar     QChart__ChartType = QChart__ChartType(2)
)

func (ptr *QChart) SetAnimationEasingCurve(curve core.QEasingCurve_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetAnimationEasingCurve(ptr.Pointer(), core.PointerFromQEasingCurve(curve))
	}
}

func NewQChart(parent widgets.QGraphicsItem_ITF, wFlags core.Qt__WindowType) *QChart {
	return NewQChartFromPointer(C.QChart_NewQChart(widgets.PointerFromQGraphicsItem(parent), C.longlong(wFlags)))
}

func (ptr *QChart) MapToPosition(value core.QPointF_ITF, series QAbstractSeries_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QChart_MapToPosition(ptr.Pointer(), core.PointerFromQPointF(value), PointerFromQAbstractSeries(series)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) MapToValue(position core.QPointF_ITF, series QAbstractSeries_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QChart_MapToValue(ptr.Pointer(), core.PointerFromQPointF(position), PointerFromQAbstractSeries(series)))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) IsZoomed() bool {
	if ptr.Pointer() != nil {
		return C.QChart_IsZoomed(ptr.Pointer()) != 0
	}
	return false
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

func (ptr *QChart) CreateDefaultAxes() {
	if ptr.Pointer() != nil {
		C.QChart_CreateDefaultAxes(ptr.Pointer())
	}
}

//export callbackQChart_PlotAreaChanged
func callbackQChart_PlotAreaChanged(ptr unsafe.Pointer, plotArea unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QChart::plotAreaChanged"); signal != nil {
		signal.(func(*core.QRectF))(core.NewQRectFFromPointer(plotArea))
	}

}

func (ptr *QChart) ConnectPlotAreaChanged(f func(plotArea *core.QRectF)) {
	if ptr.Pointer() != nil {
		C.QChart_ConnectPlotAreaChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QChart::plotAreaChanged", f)
	}
}

func (ptr *QChart) DisconnectPlotAreaChanged() {
	if ptr.Pointer() != nil {
		C.QChart_DisconnectPlotAreaChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QChart::plotAreaChanged")
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

func (ptr *QChart) SetAnimationDuration(msecs int) {
	if ptr.Pointer() != nil {
		C.QChart_SetAnimationDuration(ptr.Pointer(), C.int(int32(msecs)))
	}
}

func (ptr *QChart) SetAnimationOptions(options QChart__AnimationOption) {
	if ptr.Pointer() != nil {
		C.QChart_SetAnimationOptions(ptr.Pointer(), C.longlong(options))
	}
}

func (ptr *QChart) SetAxisX(axis QAbstractAxis_ITF, series QAbstractSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetAxisX(ptr.Pointer(), PointerFromQAbstractAxis(axis), PointerFromQAbstractSeries(series))
	}
}

func (ptr *QChart) SetAxisY(axis QAbstractAxis_ITF, series QAbstractSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QChart_SetAxisY(ptr.Pointer(), PointerFromQAbstractAxis(axis), PointerFromQAbstractSeries(series))
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
		var titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
		C.QChart_SetTitle(ptr.Pointer(), titleC)
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

func (ptr *QChart) DestroyQChart() {
	if ptr.Pointer() != nil {
		C.QChart_DestroyQChart(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QChart) AnimationOptions() QChart__AnimationOption {
	if ptr.Pointer() != nil {
		return QChart__AnimationOption(C.QChart_AnimationOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChart) ChartType() QChart__ChartType {
	if ptr.Pointer() != nil {
		return QChart__ChartType(C.QChart_ChartType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChart) AxisX(series QAbstractSeries_ITF) *QAbstractAxis {
	if ptr.Pointer() != nil {
		return NewQAbstractAxisFromPointer(C.QChart_AxisX(ptr.Pointer(), PointerFromQAbstractSeries(series)))
	}
	return nil
}

func (ptr *QChart) AxisY(series QAbstractSeries_ITF) *QAbstractAxis {
	if ptr.Pointer() != nil {
		return NewQAbstractAxisFromPointer(C.QChart_AxisY(ptr.Pointer(), PointerFromQAbstractSeries(series)))
	}
	return nil
}

func (ptr *QChart) BackgroundBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QChart_BackgroundBrush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) PlotAreaBackgroundBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QChart_PlotAreaBackgroundBrush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) TitleBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QChart_TitleBrush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) Theme() QChart__ChartTheme {
	if ptr.Pointer() != nil {
		return QChart__ChartTheme(C.QChart_Theme(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChart) AnimationEasingCurve() *core.QEasingCurve {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQEasingCurveFromPointer(C.QChart_AnimationEasingCurve(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QEasingCurve).DestroyQEasingCurve)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) TitleFont() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.QChart_TitleFont(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) Legend() *QLegend {
	if ptr.Pointer() != nil {
		return NewQLegendFromPointer(C.QChart_Legend(ptr.Pointer()))
	}
	return nil
}

func (ptr *QChart) Axes(orientation core.Qt__Orientation, series QAbstractSeries_ITF) []*QAbstractAxis {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QAbstractAxis {
			var out = make([]*QAbstractAxis, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQChartFromPointer(l.data).__axes_atList(i)
			}
			return out
		}(C.QChart_Axes(ptr.Pointer(), C.longlong(orientation), PointerFromQAbstractSeries(series)))
	}
	return make([]*QAbstractAxis, 0)
}

func (ptr *QChart) Series() []*QAbstractSeries {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QAbstractSeries {
			var out = make([]*QAbstractSeries, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQChartFromPointer(l.data).__series_atList(i)
			}
			return out
		}(C.QChart_Series(ptr.Pointer()))
	}
	return make([]*QAbstractSeries, 0)
}

func (ptr *QChart) Locale() *core.QLocale {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQLocaleFromPointer(C.QChart_Locale(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) Margins() *core.QMargins {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQMarginsFromPointer(C.QChart_Margins(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QMargins).DestroyQMargins)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) BackgroundPen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QChart_BackgroundPen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) PlotAreaBackgroundPen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QChart_PlotAreaBackgroundPen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) PlotArea() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QChart_PlotArea(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QChart) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QChart_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QChart) IsBackgroundVisible() bool {
	if ptr.Pointer() != nil {
		return C.QChart_IsBackgroundVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChart) IsDropShadowEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QChart_IsDropShadowEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChart) IsPlotAreaBackgroundVisible() bool {
	if ptr.Pointer() != nil {
		return C.QChart_IsPlotAreaBackgroundVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChart) LocalizeNumbers() bool {
	if ptr.Pointer() != nil {
		return C.QChart_LocalizeNumbers(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChart) AnimationDuration() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QChart_AnimationDuration(ptr.Pointer())))
	}
	return 0
}

func (ptr *QChart) BackgroundRoundness() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QChart_BackgroundRoundness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChart) __axes_atList(i int) *QAbstractAxis {
	if ptr.Pointer() != nil {
		return NewQAbstractAxisFromPointer(C.QChart___axes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChart) __axes_setList(i QAbstractAxis_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___axes_setList(ptr.Pointer(), PointerFromQAbstractAxis(i))
	}
}

func (ptr *QChart) __axes_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QChart___axes_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QChart) __series_atList(i int) *QAbstractSeries {
	if ptr.Pointer() != nil {
		return NewQAbstractSeriesFromPointer(C.QChart___series_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QChart) __series_setList(i QAbstractSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QChart___series_setList(ptr.Pointer(), PointerFromQAbstractSeries(i))
	}
}

func (ptr *QChart) __series_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QChart___series_newList(ptr.Pointer()))
	}
	return nil
}

type QChartView struct {
	ptr unsafe.Pointer
}

type QChartView_ITF interface {
	QChartView_PTR() *QChartView
}

func (ptr *QChartView) QChartView_PTR() *QChartView {
	return ptr
}

func (ptr *QChartView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QChartView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQChartView(ptr QChartView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QChartView_PTR().Pointer()
	}
	return nil
}

func NewQChartViewFromPointer(ptr unsafe.Pointer) *QChartView {
	var n = new(QChartView)
	n.SetPointer(ptr)
	return n
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

func NewQChartView2(chart QChart_ITF, parent widgets.QWidget_ITF) *QChartView {
	var tmpValue = NewQChartViewFromPointer(C.QChartView_NewQChartView2(PointerFromQChart(chart), widgets.PointerFromQWidget(parent)))
	runtime.SetFinalizer(tmpValue, (*QChartView).DestroyQChartView)
	return tmpValue
}

func NewQChartView(parent widgets.QWidget_ITF) *QChartView {
	var tmpValue = NewQChartViewFromPointer(C.QChartView_NewQChartView(widgets.PointerFromQWidget(parent)))
	runtime.SetFinalizer(tmpValue, (*QChartView).DestroyQChartView)
	return tmpValue
}

func (ptr *QChartView) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QChartView) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QChartView) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QChartView) ResizeEvent(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QChartView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
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

func (ptr *QChartView) DestroyQChartView() {
	if ptr.Pointer() != nil {
		C.QChartView_DestroyQChartView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QChartView) Chart() *QChart {
	if ptr.Pointer() != nil {
		return NewQChartFromPointer(C.QChartView_Chart(ptr.Pointer()))
	}
	return nil
}

func (ptr *QChartView) RubberBand() QChartView__RubberBand {
	if ptr.Pointer() != nil {
		return QChartView__RubberBand(C.QChartView_RubberBand(ptr.Pointer()))
	}
	return 0
}

type QDateTimeAxis struct {
	ptr unsafe.Pointer
}

type QDateTimeAxis_ITF interface {
	QDateTimeAxis_PTR() *QDateTimeAxis
}

func (ptr *QDateTimeAxis) QDateTimeAxis_PTR() *QDateTimeAxis {
	return ptr
}

func (ptr *QDateTimeAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDateTimeAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDateTimeAxis(ptr QDateTimeAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateTimeAxis_PTR().Pointer()
	}
	return nil
}

func NewQDateTimeAxisFromPointer(ptr unsafe.Pointer) *QDateTimeAxis {
	var n = new(QDateTimeAxis)
	n.SetPointer(ptr)
	return n
}
func NewQDateTimeAxis(parent core.QObject_ITF) *QDateTimeAxis {
	return NewQDateTimeAxisFromPointer(C.QDateTimeAxis_NewQDateTimeAxis(core.PointerFromQObject(parent)))
}

//export callbackQDateTimeAxis_FormatChanged
func callbackQDateTimeAxis_FormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDateTimeAxis::formatChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(format))
	}

}

func (ptr *QDateTimeAxis) ConnectFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_ConnectFormatChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDateTimeAxis::formatChanged", f)
	}
}

func (ptr *QDateTimeAxis) DisconnectFormatChanged() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DisconnectFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDateTimeAxis::formatChanged")
	}
}

func (ptr *QDateTimeAxis) FormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QDateTimeAxis_FormatChanged(ptr.Pointer(), formatC)
	}
}

//export callbackQDateTimeAxis_MaxChanged
func callbackQDateTimeAxis_MaxChanged(ptr unsafe.Pointer, max unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDateTimeAxis::maxChanged"); signal != nil {
		signal.(func(*core.QDateTime))(core.NewQDateTimeFromPointer(max))
	}

}

func (ptr *QDateTimeAxis) ConnectMaxChanged(f func(max *core.QDateTime)) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_ConnectMaxChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDateTimeAxis::maxChanged", f)
	}
}

func (ptr *QDateTimeAxis) DisconnectMaxChanged() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DisconnectMaxChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDateTimeAxis::maxChanged")
	}
}

func (ptr *QDateTimeAxis) MaxChanged(max core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_MaxChanged(ptr.Pointer(), core.PointerFromQDateTime(max))
	}
}

//export callbackQDateTimeAxis_MinChanged
func callbackQDateTimeAxis_MinChanged(ptr unsafe.Pointer, min unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDateTimeAxis::minChanged"); signal != nil {
		signal.(func(*core.QDateTime))(core.NewQDateTimeFromPointer(min))
	}

}

func (ptr *QDateTimeAxis) ConnectMinChanged(f func(min *core.QDateTime)) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_ConnectMinChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDateTimeAxis::minChanged", f)
	}
}

func (ptr *QDateTimeAxis) DisconnectMinChanged() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DisconnectMinChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDateTimeAxis::minChanged")
	}
}

func (ptr *QDateTimeAxis) MinChanged(min core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_MinChanged(ptr.Pointer(), core.PointerFromQDateTime(min))
	}
}

//export callbackQDateTimeAxis_RangeChanged
func callbackQDateTimeAxis_RangeChanged(ptr unsafe.Pointer, min unsafe.Pointer, max unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDateTimeAxis::rangeChanged"); signal != nil {
		signal.(func(*core.QDateTime, *core.QDateTime))(core.NewQDateTimeFromPointer(min), core.NewQDateTimeFromPointer(max))
	}

}

func (ptr *QDateTimeAxis) ConnectRangeChanged(f func(min *core.QDateTime, max *core.QDateTime)) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_ConnectRangeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDateTimeAxis::rangeChanged", f)
	}
}

func (ptr *QDateTimeAxis) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDateTimeAxis::rangeChanged")
	}
}

func (ptr *QDateTimeAxis) RangeChanged(min core.QDateTime_ITF, max core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_RangeChanged(ptr.Pointer(), core.PointerFromQDateTime(min), core.PointerFromQDateTime(max))
	}
}

func (ptr *QDateTimeAxis) SetFormat(format string) {
	if ptr.Pointer() != nil {
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QDateTimeAxis_SetFormat(ptr.Pointer(), formatC)
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

//export callbackQDateTimeAxis_TickCountChanged
func callbackQDateTimeAxis_TickCountChanged(ptr unsafe.Pointer, tickCount C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDateTimeAxis::tickCountChanged"); signal != nil {
		signal.(func(int))(int(int32(tickCount)))
	}

}

func (ptr *QDateTimeAxis) ConnectTickCountChanged(f func(tickCount int)) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_ConnectTickCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDateTimeAxis::tickCountChanged", f)
	}
}

func (ptr *QDateTimeAxis) DisconnectTickCountChanged() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DisconnectTickCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDateTimeAxis::tickCountChanged")
	}
}

func (ptr *QDateTimeAxis) TickCountChanged(tickCount int) {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_TickCountChanged(ptr.Pointer(), C.int(int32(tickCount)))
	}
}

func (ptr *QDateTimeAxis) DestroyQDateTimeAxis() {
	if ptr.Pointer() != nil {
		C.QDateTimeAxis_DestroyQDateTimeAxis(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDateTimeAxis) Max() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QDateTimeAxis_Max(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QDateTimeAxis) Min() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QDateTimeAxis_Min(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QDateTimeAxis) Format() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDateTimeAxis_Format(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDateTimeAxis) TickCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDateTimeAxis_TickCount(ptr.Pointer())))
	}
	return 0
}

type QHBarModelMapper struct {
	ptr unsafe.Pointer
}

type QHBarModelMapper_ITF interface {
	QHBarModelMapper_PTR() *QHBarModelMapper
}

func (ptr *QHBarModelMapper) QHBarModelMapper_PTR() *QHBarModelMapper {
	return ptr
}

func (ptr *QHBarModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHBarModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHBarModelMapper(ptr QHBarModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBarModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQHBarModelMapperFromPointer(ptr unsafe.Pointer) *QHBarModelMapper {
	var n = new(QHBarModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHBarModelMapper) DestroyQHBarModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQHBarModelMapper(parent core.QObject_ITF) *QHBarModelMapper {
	return NewQHBarModelMapperFromPointer(C.QHBarModelMapper_NewQHBarModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQHBarModelMapper_ColumnCountChanged
func callbackQHBarModelMapper_ColumnCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBarModelMapper::columnCountChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBarModelMapper) ConnectColumnCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ConnectColumnCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::columnCountChanged", f)
	}
}

func (ptr *QHBarModelMapper) DisconnectColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectColumnCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::columnCountChanged")
	}
}

func (ptr *QHBarModelMapper) ColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ColumnCountChanged(ptr.Pointer())
	}
}

//export callbackQHBarModelMapper_FirstBarSetRowChanged
func callbackQHBarModelMapper_FirstBarSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBarModelMapper::firstBarSetRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBarModelMapper) ConnectFirstBarSetRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ConnectFirstBarSetRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::firstBarSetRowChanged", f)
	}
}

func (ptr *QHBarModelMapper) DisconnectFirstBarSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectFirstBarSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::firstBarSetRowChanged")
	}
}

func (ptr *QHBarModelMapper) FirstBarSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_FirstBarSetRowChanged(ptr.Pointer())
	}
}

//export callbackQHBarModelMapper_FirstColumnChanged
func callbackQHBarModelMapper_FirstColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBarModelMapper::firstColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBarModelMapper) ConnectFirstColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ConnectFirstColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::firstColumnChanged", f)
	}
}

func (ptr *QHBarModelMapper) DisconnectFirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectFirstColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::firstColumnChanged")
	}
}

func (ptr *QHBarModelMapper) FirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_FirstColumnChanged(ptr.Pointer())
	}
}

//export callbackQHBarModelMapper_LastBarSetRowChanged
func callbackQHBarModelMapper_LastBarSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBarModelMapper::lastBarSetRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBarModelMapper) ConnectLastBarSetRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ConnectLastBarSetRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::lastBarSetRowChanged", f)
	}
}

func (ptr *QHBarModelMapper) DisconnectLastBarSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectLastBarSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::lastBarSetRowChanged")
	}
}

func (ptr *QHBarModelMapper) LastBarSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_LastBarSetRowChanged(ptr.Pointer())
	}
}

//export callbackQHBarModelMapper_ModelReplaced
func callbackQHBarModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBarModelMapper::modelReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBarModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ConnectModelReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::modelReplaced", f)
	}
}

func (ptr *QHBarModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::modelReplaced")
	}
}

func (ptr *QHBarModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ModelReplaced(ptr.Pointer())
	}
}

//export callbackQHBarModelMapper_SeriesReplaced
func callbackQHBarModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBarModelMapper::seriesReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBarModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_ConnectSeriesReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::seriesReplaced", f)
	}
}

func (ptr *QHBarModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHBarModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBarModelMapper::seriesReplaced")
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

func (ptr *QHBarModelMapper) Series() *QAbstractBarSeries {
	if ptr.Pointer() != nil {
		return NewQAbstractBarSeriesFromPointer(C.QHBarModelMapper_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHBarModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QHBarModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHBarModelMapper) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBarModelMapper_ColumnCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHBarModelMapper) FirstBarSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBarModelMapper_FirstBarSetRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHBarModelMapper) FirstColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBarModelMapper_FirstColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHBarModelMapper) LastBarSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBarModelMapper_LastBarSetRow(ptr.Pointer())))
	}
	return 0
}

type QHBoxPlotModelMapper struct {
	ptr unsafe.Pointer
}

type QHBoxPlotModelMapper_ITF interface {
	QHBoxPlotModelMapper_PTR() *QHBoxPlotModelMapper
}

func (ptr *QHBoxPlotModelMapper) QHBoxPlotModelMapper_PTR() *QHBoxPlotModelMapper {
	return ptr
}

func (ptr *QHBoxPlotModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHBoxPlotModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHBoxPlotModelMapper(ptr QHBoxPlotModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBoxPlotModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQHBoxPlotModelMapperFromPointer(ptr unsafe.Pointer) *QHBoxPlotModelMapper {
	var n = new(QHBoxPlotModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHBoxPlotModelMapper) DestroyQHBoxPlotModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQHBoxPlotModelMapper(parent core.QObject_ITF) *QHBoxPlotModelMapper {
	return NewQHBoxPlotModelMapperFromPointer(C.QHBoxPlotModelMapper_NewQHBoxPlotModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQHBoxPlotModelMapper_ColumnCountChanged
func callbackQHBoxPlotModelMapper_ColumnCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBoxPlotModelMapper::columnCountChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectColumnCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ConnectColumnCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::columnCountChanged", f)
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectColumnCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::columnCountChanged")
	}
}

func (ptr *QHBoxPlotModelMapper) ColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ColumnCountChanged(ptr.Pointer())
	}
}

//export callbackQHBoxPlotModelMapper_FirstBoxSetRowChanged
func callbackQHBoxPlotModelMapper_FirstBoxSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBoxPlotModelMapper::firstBoxSetRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectFirstBoxSetRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ConnectFirstBoxSetRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::firstBoxSetRowChanged", f)
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectFirstBoxSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectFirstBoxSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::firstBoxSetRowChanged")
	}
}

func (ptr *QHBoxPlotModelMapper) FirstBoxSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_FirstBoxSetRowChanged(ptr.Pointer())
	}
}

//export callbackQHBoxPlotModelMapper_FirstColumnChanged
func callbackQHBoxPlotModelMapper_FirstColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBoxPlotModelMapper::firstColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectFirstColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ConnectFirstColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::firstColumnChanged", f)
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectFirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectFirstColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::firstColumnChanged")
	}
}

func (ptr *QHBoxPlotModelMapper) FirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_FirstColumnChanged(ptr.Pointer())
	}
}

//export callbackQHBoxPlotModelMapper_LastBoxSetRowChanged
func callbackQHBoxPlotModelMapper_LastBoxSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBoxPlotModelMapper::lastBoxSetRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectLastBoxSetRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ConnectLastBoxSetRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::lastBoxSetRowChanged", f)
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectLastBoxSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectLastBoxSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::lastBoxSetRowChanged")
	}
}

func (ptr *QHBoxPlotModelMapper) LastBoxSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_LastBoxSetRowChanged(ptr.Pointer())
	}
}

//export callbackQHBoxPlotModelMapper_ModelReplaced
func callbackQHBoxPlotModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBoxPlotModelMapper::modelReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ConnectModelReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::modelReplaced", f)
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::modelReplaced")
	}
}

func (ptr *QHBoxPlotModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ModelReplaced(ptr.Pointer())
	}
}

//export callbackQHBoxPlotModelMapper_SeriesReplaced
func callbackQHBoxPlotModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHBoxPlotModelMapper::seriesReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHBoxPlotModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_ConnectSeriesReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::seriesReplaced", f)
	}
}

func (ptr *QHBoxPlotModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHBoxPlotModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHBoxPlotModelMapper::seriesReplaced")
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

func (ptr *QHBoxPlotModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QHBoxPlotModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHBoxPlotModelMapper) Series() *QBoxPlotSeries {
	if ptr.Pointer() != nil {
		return NewQBoxPlotSeriesFromPointer(C.QHBoxPlotModelMapper_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHBoxPlotModelMapper) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBoxPlotModelMapper_ColumnCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHBoxPlotModelMapper) FirstBoxSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBoxPlotModelMapper_FirstBoxSetRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHBoxPlotModelMapper) FirstColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBoxPlotModelMapper_FirstColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHBoxPlotModelMapper) LastBoxSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHBoxPlotModelMapper_LastBoxSetRow(ptr.Pointer())))
	}
	return 0
}

type QHCandlestickModelMapper struct {
	ptr unsafe.Pointer
}

type QHCandlestickModelMapper_ITF interface {
	QHCandlestickModelMapper_PTR() *QHCandlestickModelMapper
}

func (ptr *QHCandlestickModelMapper) QHCandlestickModelMapper_PTR() *QHCandlestickModelMapper {
	return ptr
}

func (ptr *QHCandlestickModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHCandlestickModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHCandlestickModelMapper(ptr QHCandlestickModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHCandlestickModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQHCandlestickModelMapperFromPointer(ptr unsafe.Pointer) *QHCandlestickModelMapper {
	var n = new(QHCandlestickModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHCandlestickModelMapper) DestroyQHCandlestickModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQHCandlestickModelMapper(parent core.QObject_ITF) *QHCandlestickModelMapper {
	return NewQHCandlestickModelMapperFromPointer(C.QHCandlestickModelMapper_NewQHCandlestickModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQHCandlestickModelMapper_CloseColumnChanged
func callbackQHCandlestickModelMapper_CloseColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHCandlestickModelMapper::closeColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectCloseColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_ConnectCloseColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::closeColumnChanged", f)
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectCloseColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectCloseColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::closeColumnChanged")
	}
}

func (ptr *QHCandlestickModelMapper) CloseColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_CloseColumnChanged(ptr.Pointer())
	}
}

//export callbackQHCandlestickModelMapper_FirstSetRowChanged
func callbackQHCandlestickModelMapper_FirstSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHCandlestickModelMapper::firstSetRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectFirstSetRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_ConnectFirstSetRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::firstSetRowChanged", f)
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectFirstSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectFirstSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::firstSetRowChanged")
	}
}

func (ptr *QHCandlestickModelMapper) FirstSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_FirstSetRowChanged(ptr.Pointer())
	}
}

//export callbackQHCandlestickModelMapper_HighColumnChanged
func callbackQHCandlestickModelMapper_HighColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHCandlestickModelMapper::highColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectHighColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_ConnectHighColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::highColumnChanged", f)
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectHighColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectHighColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::highColumnChanged")
	}
}

func (ptr *QHCandlestickModelMapper) HighColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_HighColumnChanged(ptr.Pointer())
	}
}

//export callbackQHCandlestickModelMapper_LastSetRowChanged
func callbackQHCandlestickModelMapper_LastSetRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHCandlestickModelMapper::lastSetRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectLastSetRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_ConnectLastSetRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::lastSetRowChanged", f)
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectLastSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectLastSetRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::lastSetRowChanged")
	}
}

func (ptr *QHCandlestickModelMapper) LastSetRowChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_LastSetRowChanged(ptr.Pointer())
	}
}

//export callbackQHCandlestickModelMapper_LowColumnChanged
func callbackQHCandlestickModelMapper_LowColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHCandlestickModelMapper::lowColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectLowColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_ConnectLowColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::lowColumnChanged", f)
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectLowColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectLowColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::lowColumnChanged")
	}
}

func (ptr *QHCandlestickModelMapper) LowColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_LowColumnChanged(ptr.Pointer())
	}
}

//export callbackQHCandlestickModelMapper_OpenColumnChanged
func callbackQHCandlestickModelMapper_OpenColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHCandlestickModelMapper::openColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectOpenColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_ConnectOpenColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::openColumnChanged", f)
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectOpenColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectOpenColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::openColumnChanged")
	}
}

func (ptr *QHCandlestickModelMapper) OpenColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_OpenColumnChanged(ptr.Pointer())
	}
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

//export callbackQHCandlestickModelMapper_TimestampColumnChanged
func callbackQHCandlestickModelMapper_TimestampColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHCandlestickModelMapper::timestampColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHCandlestickModelMapper) ConnectTimestampColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_ConnectTimestampColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::timestampColumnChanged", f)
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectTimestampColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_DisconnectTimestampColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::timestampColumnChanged")
	}
}

func (ptr *QHCandlestickModelMapper) TimestampColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHCandlestickModelMapper_TimestampColumnChanged(ptr.Pointer())
	}
}

//export callbackQHCandlestickModelMapper_Orientation
func callbackQHCandlestickModelMapper_Orientation(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHCandlestickModelMapper::orientation"); signal != nil {
		return C.longlong(signal.(func() core.Qt__Orientation)())
	}

	return C.longlong(NewQHCandlestickModelMapperFromPointer(ptr).OrientationDefault())
}

func (ptr *QHCandlestickModelMapper) ConnectOrientation(f func() core.Qt__Orientation) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::orientation", f)
	}
}

func (ptr *QHCandlestickModelMapper) DisconnectOrientation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHCandlestickModelMapper::orientation")
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

func (ptr *QHCandlestickModelMapper) CloseColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_CloseColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHCandlestickModelMapper) FirstSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_FirstSetRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHCandlestickModelMapper) HighColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_HighColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHCandlestickModelMapper) LastSetRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_LastSetRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHCandlestickModelMapper) LowColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_LowColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHCandlestickModelMapper) OpenColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_OpenColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHCandlestickModelMapper) TimestampColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHCandlestickModelMapper_TimestampColumn(ptr.Pointer())))
	}
	return 0
}

type QHPieModelMapper struct {
	ptr unsafe.Pointer
}

type QHPieModelMapper_ITF interface {
	QHPieModelMapper_PTR() *QHPieModelMapper
}

func (ptr *QHPieModelMapper) QHPieModelMapper_PTR() *QHPieModelMapper {
	return ptr
}

func (ptr *QHPieModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHPieModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHPieModelMapper(ptr QHPieModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHPieModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQHPieModelMapperFromPointer(ptr unsafe.Pointer) *QHPieModelMapper {
	var n = new(QHPieModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHPieModelMapper) DestroyQHPieModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQHPieModelMapper(parent core.QObject_ITF) *QHPieModelMapper {
	return NewQHPieModelMapperFromPointer(C.QHPieModelMapper_NewQHPieModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQHPieModelMapper_ColumnCountChanged
func callbackQHPieModelMapper_ColumnCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHPieModelMapper::columnCountChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHPieModelMapper) ConnectColumnCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ConnectColumnCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::columnCountChanged", f)
	}
}

func (ptr *QHPieModelMapper) DisconnectColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectColumnCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::columnCountChanged")
	}
}

func (ptr *QHPieModelMapper) ColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ColumnCountChanged(ptr.Pointer())
	}
}

//export callbackQHPieModelMapper_FirstColumnChanged
func callbackQHPieModelMapper_FirstColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHPieModelMapper::firstColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHPieModelMapper) ConnectFirstColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ConnectFirstColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::firstColumnChanged", f)
	}
}

func (ptr *QHPieModelMapper) DisconnectFirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectFirstColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::firstColumnChanged")
	}
}

func (ptr *QHPieModelMapper) FirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_FirstColumnChanged(ptr.Pointer())
	}
}

//export callbackQHPieModelMapper_LabelsRowChanged
func callbackQHPieModelMapper_LabelsRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHPieModelMapper::labelsRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHPieModelMapper) ConnectLabelsRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ConnectLabelsRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::labelsRowChanged", f)
	}
}

func (ptr *QHPieModelMapper) DisconnectLabelsRowChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectLabelsRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::labelsRowChanged")
	}
}

func (ptr *QHPieModelMapper) LabelsRowChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_LabelsRowChanged(ptr.Pointer())
	}
}

//export callbackQHPieModelMapper_ModelReplaced
func callbackQHPieModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHPieModelMapper::modelReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHPieModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ConnectModelReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::modelReplaced", f)
	}
}

func (ptr *QHPieModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::modelReplaced")
	}
}

func (ptr *QHPieModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ModelReplaced(ptr.Pointer())
	}
}

//export callbackQHPieModelMapper_SeriesReplaced
func callbackQHPieModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHPieModelMapper::seriesReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHPieModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ConnectSeriesReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::seriesReplaced", f)
	}
}

func (ptr *QHPieModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::seriesReplaced")
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

//export callbackQHPieModelMapper_ValuesRowChanged
func callbackQHPieModelMapper_ValuesRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHPieModelMapper::valuesRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHPieModelMapper) ConnectValuesRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ConnectValuesRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::valuesRowChanged", f)
	}
}

func (ptr *QHPieModelMapper) DisconnectValuesRowChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_DisconnectValuesRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHPieModelMapper::valuesRowChanged")
	}
}

func (ptr *QHPieModelMapper) ValuesRowChanged() {
	if ptr.Pointer() != nil {
		C.QHPieModelMapper_ValuesRowChanged(ptr.Pointer())
	}
}

func (ptr *QHPieModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QHPieModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHPieModelMapper) Series() *QPieSeries {
	if ptr.Pointer() != nil {
		return NewQPieSeriesFromPointer(C.QHPieModelMapper_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHPieModelMapper) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHPieModelMapper_ColumnCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHPieModelMapper) FirstColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHPieModelMapper_FirstColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHPieModelMapper) LabelsRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHPieModelMapper_LabelsRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHPieModelMapper) ValuesRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHPieModelMapper_ValuesRow(ptr.Pointer())))
	}
	return 0
}

type QHXYModelMapper struct {
	ptr unsafe.Pointer
}

type QHXYModelMapper_ITF interface {
	QHXYModelMapper_PTR() *QHXYModelMapper
}

func (ptr *QHXYModelMapper) QHXYModelMapper_PTR() *QHXYModelMapper {
	return ptr
}

func (ptr *QHXYModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHXYModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHXYModelMapper(ptr QHXYModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHXYModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQHXYModelMapperFromPointer(ptr unsafe.Pointer) *QHXYModelMapper {
	var n = new(QHXYModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHXYModelMapper) DestroyQHXYModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQHXYModelMapper(parent core.QObject_ITF) *QHXYModelMapper {
	return NewQHXYModelMapperFromPointer(C.QHXYModelMapper_NewQHXYModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQHXYModelMapper_ColumnCountChanged
func callbackQHXYModelMapper_ColumnCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHXYModelMapper::columnCountChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHXYModelMapper) ConnectColumnCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_ConnectColumnCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::columnCountChanged", f)
	}
}

func (ptr *QHXYModelMapper) DisconnectColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectColumnCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::columnCountChanged")
	}
}

func (ptr *QHXYModelMapper) ColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_ColumnCountChanged(ptr.Pointer())
	}
}

//export callbackQHXYModelMapper_FirstColumnChanged
func callbackQHXYModelMapper_FirstColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHXYModelMapper::firstColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHXYModelMapper) ConnectFirstColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_ConnectFirstColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::firstColumnChanged", f)
	}
}

func (ptr *QHXYModelMapper) DisconnectFirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectFirstColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::firstColumnChanged")
	}
}

func (ptr *QHXYModelMapper) FirstColumnChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_FirstColumnChanged(ptr.Pointer())
	}
}

//export callbackQHXYModelMapper_ModelReplaced
func callbackQHXYModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHXYModelMapper::modelReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHXYModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_ConnectModelReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::modelReplaced", f)
	}
}

func (ptr *QHXYModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::modelReplaced")
	}
}

func (ptr *QHXYModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_ModelReplaced(ptr.Pointer())
	}
}

//export callbackQHXYModelMapper_SeriesReplaced
func callbackQHXYModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHXYModelMapper::seriesReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHXYModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_ConnectSeriesReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::seriesReplaced", f)
	}
}

func (ptr *QHXYModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::seriesReplaced")
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

//export callbackQHXYModelMapper_XRowChanged
func callbackQHXYModelMapper_XRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHXYModelMapper::xRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHXYModelMapper) ConnectXRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_ConnectXRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::xRowChanged", f)
	}
}

func (ptr *QHXYModelMapper) DisconnectXRowChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectXRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::xRowChanged")
	}
}

func (ptr *QHXYModelMapper) XRowChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_XRowChanged(ptr.Pointer())
	}
}

//export callbackQHXYModelMapper_YRowChanged
func callbackQHXYModelMapper_YRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHXYModelMapper::yRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHXYModelMapper) ConnectYRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_ConnectYRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::yRowChanged", f)
	}
}

func (ptr *QHXYModelMapper) DisconnectYRowChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_DisconnectYRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHXYModelMapper::yRowChanged")
	}
}

func (ptr *QHXYModelMapper) YRowChanged() {
	if ptr.Pointer() != nil {
		C.QHXYModelMapper_YRowChanged(ptr.Pointer())
	}
}

func (ptr *QHXYModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QHXYModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHXYModelMapper) Series() *QXYSeries {
	if ptr.Pointer() != nil {
		return NewQXYSeriesFromPointer(C.QHXYModelMapper_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHXYModelMapper) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHXYModelMapper_ColumnCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHXYModelMapper) FirstColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHXYModelMapper_FirstColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHXYModelMapper) XRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHXYModelMapper_XRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHXYModelMapper) YRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHXYModelMapper_YRow(ptr.Pointer())))
	}
	return 0
}

type QHorizontalBarSeries struct {
	ptr unsafe.Pointer
}

type QHorizontalBarSeries_ITF interface {
	QHorizontalBarSeries_PTR() *QHorizontalBarSeries
}

func (ptr *QHorizontalBarSeries) QHorizontalBarSeries_PTR() *QHorizontalBarSeries {
	return ptr
}

func (ptr *QHorizontalBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHorizontalBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHorizontalBarSeries(ptr QHorizontalBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHorizontalBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQHorizontalBarSeriesFromPointer(ptr unsafe.Pointer) *QHorizontalBarSeries {
	var n = new(QHorizontalBarSeries)
	n.SetPointer(ptr)
	return n
}
func NewQHorizontalBarSeries(parent core.QObject_ITF) *QHorizontalBarSeries {
	return NewQHorizontalBarSeriesFromPointer(C.QHorizontalBarSeries_NewQHorizontalBarSeries(core.PointerFromQObject(parent)))
}

func (ptr *QHorizontalBarSeries) DestroyQHorizontalBarSeries() {
	if ptr.Pointer() != nil {
		C.QHorizontalBarSeries_DestroyQHorizontalBarSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHorizontalBarSeries_Type
func callbackQHorizontalBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHorizontalBarSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQHorizontalBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QHorizontalBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHorizontalBarSeries::type", f)
	}
}

func (ptr *QHorizontalBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHorizontalBarSeries::type")
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

type QHorizontalPercentBarSeries struct {
	ptr unsafe.Pointer
}

type QHorizontalPercentBarSeries_ITF interface {
	QHorizontalPercentBarSeries_PTR() *QHorizontalPercentBarSeries
}

func (ptr *QHorizontalPercentBarSeries) QHorizontalPercentBarSeries_PTR() *QHorizontalPercentBarSeries {
	return ptr
}

func (ptr *QHorizontalPercentBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHorizontalPercentBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHorizontalPercentBarSeries(ptr QHorizontalPercentBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHorizontalPercentBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQHorizontalPercentBarSeriesFromPointer(ptr unsafe.Pointer) *QHorizontalPercentBarSeries {
	var n = new(QHorizontalPercentBarSeries)
	n.SetPointer(ptr)
	return n
}
func NewQHorizontalPercentBarSeries(parent core.QObject_ITF) *QHorizontalPercentBarSeries {
	return NewQHorizontalPercentBarSeriesFromPointer(C.QHorizontalPercentBarSeries_NewQHorizontalPercentBarSeries(core.PointerFromQObject(parent)))
}

func (ptr *QHorizontalPercentBarSeries) DestroyQHorizontalPercentBarSeries() {
	if ptr.Pointer() != nil {
		C.QHorizontalPercentBarSeries_DestroyQHorizontalPercentBarSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHorizontalPercentBarSeries_Type
func callbackQHorizontalPercentBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHorizontalPercentBarSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQHorizontalPercentBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QHorizontalPercentBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHorizontalPercentBarSeries::type", f)
	}
}

func (ptr *QHorizontalPercentBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHorizontalPercentBarSeries::type")
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

type QHorizontalStackedBarSeries struct {
	ptr unsafe.Pointer
}

type QHorizontalStackedBarSeries_ITF interface {
	QHorizontalStackedBarSeries_PTR() *QHorizontalStackedBarSeries
}

func (ptr *QHorizontalStackedBarSeries) QHorizontalStackedBarSeries_PTR() *QHorizontalStackedBarSeries {
	return ptr
}

func (ptr *QHorizontalStackedBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHorizontalStackedBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHorizontalStackedBarSeries(ptr QHorizontalStackedBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHorizontalStackedBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQHorizontalStackedBarSeriesFromPointer(ptr unsafe.Pointer) *QHorizontalStackedBarSeries {
	var n = new(QHorizontalStackedBarSeries)
	n.SetPointer(ptr)
	return n
}
func NewQHorizontalStackedBarSeries(parent core.QObject_ITF) *QHorizontalStackedBarSeries {
	return NewQHorizontalStackedBarSeriesFromPointer(C.QHorizontalStackedBarSeries_NewQHorizontalStackedBarSeries(core.PointerFromQObject(parent)))
}

func (ptr *QHorizontalStackedBarSeries) DestroyQHorizontalStackedBarSeries() {
	if ptr.Pointer() != nil {
		C.QHorizontalStackedBarSeries_DestroyQHorizontalStackedBarSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHorizontalStackedBarSeries_Type
func callbackQHorizontalStackedBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHorizontalStackedBarSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQHorizontalStackedBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QHorizontalStackedBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHorizontalStackedBarSeries::type", f)
	}
}

func (ptr *QHorizontalStackedBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHorizontalStackedBarSeries::type")
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

type QLegend struct {
	ptr unsafe.Pointer
}

type QLegend_ITF interface {
	QLegend_PTR() *QLegend
}

func (ptr *QLegend) QLegend_PTR() *QLegend {
	return ptr
}

func (ptr *QLegend) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLegend) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLegend(ptr QLegend_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLegend_PTR().Pointer()
	}
	return nil
}

func NewQLegendFromPointer(ptr unsafe.Pointer) *QLegend {
	var n = new(QLegend)
	n.SetPointer(ptr)
	return n
}
func (ptr *QLegend) BorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QLegend_BorderColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QLegend_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) IsAttachedToChart() bool {
	if ptr.Pointer() != nil {
		return C.QLegend_IsAttachedToChart(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLegend) ReverseMarkers() bool {
	if ptr.Pointer() != nil {
		return C.QLegend_ReverseMarkers(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLegend) AttachToChart() {
	if ptr.Pointer() != nil {
		C.QLegend_AttachToChart(ptr.Pointer())
	}
}

//export callbackQLegend_BackgroundVisibleChanged
func callbackQLegend_BackgroundVisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegend::backgroundVisibleChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QLegend) ConnectBackgroundVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QLegend_ConnectBackgroundVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::backgroundVisibleChanged", f)
	}
}

func (ptr *QLegend) DisconnectBackgroundVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectBackgroundVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::backgroundVisibleChanged")
	}
}

func (ptr *QLegend) BackgroundVisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QLegend_BackgroundVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQLegend_BorderColorChanged
func callbackQLegend_BorderColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegend::borderColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QLegend) ConnectBorderColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QLegend_ConnectBorderColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::borderColorChanged", f)
	}
}

func (ptr *QLegend) DisconnectBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::borderColorChanged")
	}
}

func (ptr *QLegend) BorderColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_BorderColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQLegend_ColorChanged
func callbackQLegend_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegend::colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QLegend) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QLegend_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::colorChanged", f)
	}
}

func (ptr *QLegend) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::colorChanged")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegend::fontChanged"); signal != nil {
		signal.(func(*gui.QFont))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *QLegend) ConnectFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {
		C.QLegend_ConnectFontChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::fontChanged", f)
	}
}

func (ptr *QLegend) DisconnectFontChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectFontChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::fontChanged")
	}
}

func (ptr *QLegend) FontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_FontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

//export callbackQLegend_LabelColorChanged
func callbackQLegend_LabelColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegend::labelColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QLegend) ConnectLabelColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QLegend_ConnectLabelColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::labelColorChanged", f)
	}
}

func (ptr *QLegend) DisconnectLabelColorChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectLabelColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::labelColorChanged")
	}
}

func (ptr *QLegend) LabelColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_LabelColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQLegend_ReverseMarkersChanged
func callbackQLegend_ReverseMarkersChanged(ptr unsafe.Pointer, reverseMarkers C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegend::reverseMarkersChanged"); signal != nil {
		signal.(func(bool))(int8(reverseMarkers) != 0)
	}

}

func (ptr *QLegend) ConnectReverseMarkersChanged(f func(reverseMarkers bool)) {
	if ptr.Pointer() != nil {
		C.QLegend_ConnectReverseMarkersChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::reverseMarkersChanged", f)
	}
}

func (ptr *QLegend) DisconnectReverseMarkersChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectReverseMarkersChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::reverseMarkersChanged")
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

func (ptr *QLegend) SetFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
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

//export callbackQLegend_ShowToolTipsChanged
func callbackQLegend_ShowToolTipsChanged(ptr unsafe.Pointer, showToolTips C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegend::showToolTipsChanged"); signal != nil {
		signal.(func(bool))(int8(showToolTips) != 0)
	}

}

func (ptr *QLegend) ConnectShowToolTipsChanged(f func(showToolTips bool)) {
	if ptr.Pointer() != nil {
		C.QLegend_ConnectShowToolTipsChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::showToolTipsChanged", f)
	}
}

func (ptr *QLegend) DisconnectShowToolTipsChanged() {
	if ptr.Pointer() != nil {
		C.QLegend_DisconnectShowToolTipsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegend::showToolTipsChanged")
	}
}

func (ptr *QLegend) ShowToolTipsChanged(showToolTips bool) {
	if ptr.Pointer() != nil {
		C.QLegend_ShowToolTipsChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(showToolTips))))
	}
}

func (ptr *QLegend) DestroyQLegend() {
	if ptr.Pointer() != nil {
		C.QLegend_DestroyQLegend(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLegend) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QLegend_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) LabelBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QLegend_LabelBrush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) LabelColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QLegend_LabelColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) Font() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.QLegend_Font(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) Markers(series QAbstractSeries_ITF) []*QLegendMarker {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QLegendMarker {
			var out = make([]*QLegendMarker, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLegendFromPointer(l.data).__markers_atList(i)
			}
			return out
		}(C.QLegend_Markers(ptr.Pointer(), PointerFromQAbstractSeries(series)))
	}
	return make([]*QLegendMarker, 0)
}

func (ptr *QLegend) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QLegend_Pen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QLegend) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLegend_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLegend) IsBackgroundVisible() bool {
	if ptr.Pointer() != nil {
		return C.QLegend_IsBackgroundVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLegend) ShowToolTips() bool {
	if ptr.Pointer() != nil {
		return C.QLegend_ShowToolTips(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLegend) __markers_atList(i int) *QLegendMarker {
	if ptr.Pointer() != nil {
		return NewQLegendMarkerFromPointer(C.QLegend___markers_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QLegend) __markers_setList(i QLegendMarker_ITF) {
	if ptr.Pointer() != nil {
		C.QLegend___markers_setList(ptr.Pointer(), PointerFromQLegendMarker(i))
	}
}

func (ptr *QLegend) __markers_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QLegend___markers_newList(ptr.Pointer()))
	}
	return nil
}

type QLegendMarker struct {
	ptr unsafe.Pointer
}

type QLegendMarker_ITF interface {
	QLegendMarker_PTR() *QLegendMarker
}

func (ptr *QLegendMarker) QLegendMarker_PTR() *QLegendMarker {
	return ptr
}

func (ptr *QLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLegendMarker(ptr QLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQLegendMarkerFromPointer(ptr unsafe.Pointer) *QLegendMarker {
	var n = new(QLegendMarker)
	n.SetPointer(ptr)
	return n
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

//export callbackQLegendMarker_Type
func callbackQLegendMarker_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::type"); signal != nil {
		return C.longlong(signal.(func() QLegendMarker__LegendMarkerType)())
	}

	return C.longlong(0)
}

func (ptr *QLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::type", f)
	}
}

func (ptr *QLegendMarker) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::type")
	}
}

func (ptr *QLegendMarker) Type() QLegendMarker__LegendMarkerType {
	if ptr.Pointer() != nil {
		return QLegendMarker__LegendMarkerType(C.QLegendMarker_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQLegendMarker_Series
func callbackQLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::series"); signal != nil {
		return PointerFromQAbstractSeries(signal.(func() *QAbstractSeries)())
	}

	return PointerFromQAbstractSeries(nil)
}

func (ptr *QLegendMarker) ConnectSeries(f func() *QAbstractSeries) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::series", f)
	}
}

func (ptr *QLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::series")
	}
}

func (ptr *QLegendMarker) Series() *QAbstractSeries {
	if ptr.Pointer() != nil {
		return NewQAbstractSeriesFromPointer(C.QLegendMarker_Series(ptr.Pointer()))
	}
	return nil
}

//export callbackQLegendMarker_BrushChanged
func callbackQLegendMarker_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::brushChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLegendMarker) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ConnectBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::brushChanged", f)
	}
}

func (ptr *QLegendMarker) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::brushChanged")
	}
}

func (ptr *QLegendMarker) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_BrushChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_Clicked
func callbackQLegendMarker_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::clicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLegendMarker) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::clicked", f)
	}
}

func (ptr *QLegendMarker) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::clicked")
	}
}

func (ptr *QLegendMarker) Clicked() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_Clicked(ptr.Pointer())
	}
}

//export callbackQLegendMarker_FontChanged
func callbackQLegendMarker_FontChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::fontChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLegendMarker) ConnectFontChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ConnectFontChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::fontChanged", f)
	}
}

func (ptr *QLegendMarker) DisconnectFontChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectFontChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::fontChanged")
	}
}

func (ptr *QLegendMarker) FontChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_FontChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_Hovered
func callbackQLegendMarker_Hovered(ptr unsafe.Pointer, status C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::hovered"); signal != nil {
		signal.(func(bool))(int8(status) != 0)
	}

}

func (ptr *QLegendMarker) ConnectHovered(f func(status bool)) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::hovered", f)
	}
}

func (ptr *QLegendMarker) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::hovered")
	}
}

func (ptr *QLegendMarker) Hovered(status bool) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(status))))
	}
}

//export callbackQLegendMarker_LabelBrushChanged
func callbackQLegendMarker_LabelBrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::labelBrushChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLegendMarker) ConnectLabelBrushChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ConnectLabelBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::labelBrushChanged", f)
	}
}

func (ptr *QLegendMarker) DisconnectLabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectLabelBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::labelBrushChanged")
	}
}

func (ptr *QLegendMarker) LabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_LabelBrushChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_LabelChanged
func callbackQLegendMarker_LabelChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::labelChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLegendMarker) ConnectLabelChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ConnectLabelChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::labelChanged", f)
	}
}

func (ptr *QLegendMarker) DisconnectLabelChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectLabelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::labelChanged")
	}
}

func (ptr *QLegendMarker) LabelChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_LabelChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_PenChanged
func callbackQLegendMarker_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::penChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLegendMarker) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ConnectPenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::penChanged", f)
	}
}

func (ptr *QLegendMarker) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::penChanged")
	}
}

func (ptr *QLegendMarker) PenChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_PenChanged(ptr.Pointer())
	}
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
		var labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
		C.QLegendMarker_SetLabel(ptr.Pointer(), labelC)
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

func (ptr *QLegendMarker) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQLegendMarker_VisibleChanged
func callbackQLegendMarker_VisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::visibleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLegendMarker) ConnectVisibleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QLegendMarker_ConnectVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::visibleChanged", f)
	}
}

func (ptr *QLegendMarker) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::visibleChanged")
	}
}

func (ptr *QLegendMarker) VisibleChanged() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_VisibleChanged(ptr.Pointer())
	}
}

//export callbackQLegendMarker_DestroyQLegendMarker
func callbackQLegendMarker_DestroyQLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLegendMarker::~QLegendMarker"); signal != nil {
		signal.(func())()
	} else {
		NewQLegendMarkerFromPointer(ptr).DestroyQLegendMarkerDefault()
	}
}

func (ptr *QLegendMarker) ConnectDestroyQLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::~QLegendMarker", f)
	}
}

func (ptr *QLegendMarker) DisconnectDestroyQLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLegendMarker::~QLegendMarker")
	}
}

func (ptr *QLegendMarker) DestroyQLegendMarker() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DestroyQLegendMarker(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLegendMarker) DestroyQLegendMarkerDefault() {
	if ptr.Pointer() != nil {
		C.QLegendMarker_DestroyQLegendMarkerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLegendMarker) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QLegendMarker_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QLegendMarker) LabelBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QLegendMarker_LabelBrush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QLegendMarker) Font() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.QLegendMarker_Font(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QLegendMarker) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QLegendMarker_Pen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QLegendMarker) Label() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLegendMarker_Label(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLegendMarker) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QLegendMarker_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

type QLineSeries struct {
	ptr unsafe.Pointer
}

type QLineSeries_ITF interface {
	QLineSeries_PTR() *QLineSeries
}

func (ptr *QLineSeries) QLineSeries_PTR() *QLineSeries {
	return ptr
}

func (ptr *QLineSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLineSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLineSeries(ptr QLineSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLineSeries_PTR().Pointer()
	}
	return nil
}

func NewQLineSeriesFromPointer(ptr unsafe.Pointer) *QLineSeries {
	var n = new(QLineSeries)
	n.SetPointer(ptr)
	return n
}
func NewQLineSeries(parent core.QObject_ITF) *QLineSeries {
	return NewQLineSeriesFromPointer(C.QLineSeries_NewQLineSeries(core.PointerFromQObject(parent)))
}

func (ptr *QLineSeries) DestroyQLineSeries() {
	if ptr.Pointer() != nil {
		C.QLineSeries_DestroyQLineSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLineSeries_Type
func callbackQLineSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLineSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQLineSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QLineSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLineSeries::type", f)
	}
}

func (ptr *QLineSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLineSeries::type")
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

type QLogValueAxis struct {
	ptr unsafe.Pointer
}

type QLogValueAxis_ITF interface {
	QLogValueAxis_PTR() *QLogValueAxis
}

func (ptr *QLogValueAxis) QLogValueAxis_PTR() *QLogValueAxis {
	return ptr
}

func (ptr *QLogValueAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLogValueAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLogValueAxis(ptr QLogValueAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLogValueAxis_PTR().Pointer()
	}
	return nil
}

func NewQLogValueAxisFromPointer(ptr unsafe.Pointer) *QLogValueAxis {
	var n = new(QLogValueAxis)
	n.SetPointer(ptr)
	return n
}
func NewQLogValueAxis(parent core.QObject_ITF) *QLogValueAxis {
	return NewQLogValueAxisFromPointer(C.QLogValueAxis_NewQLogValueAxis(core.PointerFromQObject(parent)))
}

//export callbackQLogValueAxis_BaseChanged
func callbackQLogValueAxis_BaseChanged(ptr unsafe.Pointer, base C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLogValueAxis::baseChanged"); signal != nil {
		signal.(func(float64))(float64(base))
	}

}

func (ptr *QLogValueAxis) ConnectBaseChanged(f func(base float64)) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_ConnectBaseChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLogValueAxis::baseChanged", f)
	}
}

func (ptr *QLogValueAxis) DisconnectBaseChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectBaseChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLogValueAxis::baseChanged")
	}
}

func (ptr *QLogValueAxis) BaseChanged(base float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_BaseChanged(ptr.Pointer(), C.double(base))
	}
}

//export callbackQLogValueAxis_LabelFormatChanged
func callbackQLogValueAxis_LabelFormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLogValueAxis::labelFormatChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(format))
	}

}

func (ptr *QLogValueAxis) ConnectLabelFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_ConnectLabelFormatChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLogValueAxis::labelFormatChanged", f)
	}
}

func (ptr *QLogValueAxis) DisconnectLabelFormatChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectLabelFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLogValueAxis::labelFormatChanged")
	}
}

func (ptr *QLogValueAxis) LabelFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QLogValueAxis_LabelFormatChanged(ptr.Pointer(), formatC)
	}
}

//export callbackQLogValueAxis_MaxChanged
func callbackQLogValueAxis_MaxChanged(ptr unsafe.Pointer, max C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLogValueAxis::maxChanged"); signal != nil {
		signal.(func(float64))(float64(max))
	}

}

func (ptr *QLogValueAxis) ConnectMaxChanged(f func(max float64)) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_ConnectMaxChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLogValueAxis::maxChanged", f)
	}
}

func (ptr *QLogValueAxis) DisconnectMaxChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectMaxChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLogValueAxis::maxChanged")
	}
}

func (ptr *QLogValueAxis) MaxChanged(max float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_MaxChanged(ptr.Pointer(), C.double(max))
	}
}

//export callbackQLogValueAxis_MinChanged
func callbackQLogValueAxis_MinChanged(ptr unsafe.Pointer, min C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLogValueAxis::minChanged"); signal != nil {
		signal.(func(float64))(float64(min))
	}

}

func (ptr *QLogValueAxis) ConnectMinChanged(f func(min float64)) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_ConnectMinChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLogValueAxis::minChanged", f)
	}
}

func (ptr *QLogValueAxis) DisconnectMinChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectMinChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLogValueAxis::minChanged")
	}
}

func (ptr *QLogValueAxis) MinChanged(min float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_MinChanged(ptr.Pointer(), C.double(min))
	}
}

//export callbackQLogValueAxis_RangeChanged
func callbackQLogValueAxis_RangeChanged(ptr unsafe.Pointer, min C.double, max C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLogValueAxis::rangeChanged"); signal != nil {
		signal.(func(float64, float64))(float64(min), float64(max))
	}

}

func (ptr *QLogValueAxis) ConnectRangeChanged(f func(min float64, max float64)) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_ConnectRangeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLogValueAxis::rangeChanged", f)
	}
}

func (ptr *QLogValueAxis) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLogValueAxis::rangeChanged")
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
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QLogValueAxis_SetLabelFormat(ptr.Pointer(), formatC)
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

func (ptr *QLogValueAxis) SetRange(min float64, max float64) {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_SetRange(ptr.Pointer(), C.double(min), C.double(max))
	}
}

func (ptr *QLogValueAxis) DestroyQLogValueAxis() {
	if ptr.Pointer() != nil {
		C.QLogValueAxis_DestroyQLogValueAxis(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLogValueAxis) LabelFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLogValueAxis_LabelFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLogValueAxis) Base() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLogValueAxis_Base(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLogValueAxis) Max() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLogValueAxis_Max(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLogValueAxis) Min() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLogValueAxis_Min(ptr.Pointer()))
	}
	return 0
}

type QPercentBarSeries struct {
	ptr unsafe.Pointer
}

type QPercentBarSeries_ITF interface {
	QPercentBarSeries_PTR() *QPercentBarSeries
}

func (ptr *QPercentBarSeries) QPercentBarSeries_PTR() *QPercentBarSeries {
	return ptr
}

func (ptr *QPercentBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPercentBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPercentBarSeries(ptr QPercentBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPercentBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQPercentBarSeriesFromPointer(ptr unsafe.Pointer) *QPercentBarSeries {
	var n = new(QPercentBarSeries)
	n.SetPointer(ptr)
	return n
}
func NewQPercentBarSeries(parent core.QObject_ITF) *QPercentBarSeries {
	return NewQPercentBarSeriesFromPointer(C.QPercentBarSeries_NewQPercentBarSeries(core.PointerFromQObject(parent)))
}

func (ptr *QPercentBarSeries) DestroyQPercentBarSeries() {
	if ptr.Pointer() != nil {
		C.QPercentBarSeries_DestroyQPercentBarSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQPercentBarSeries_Type
func callbackQPercentBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPercentBarSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQPercentBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QPercentBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPercentBarSeries::type", f)
	}
}

func (ptr *QPercentBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPercentBarSeries::type")
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

type QPieLegendMarker struct {
	ptr unsafe.Pointer
}

type QPieLegendMarker_ITF interface {
	QPieLegendMarker_PTR() *QPieLegendMarker
}

func (ptr *QPieLegendMarker) QPieLegendMarker_PTR() *QPieLegendMarker {
	return ptr
}

func (ptr *QPieLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPieLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPieLegendMarker(ptr QPieLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPieLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQPieLegendMarkerFromPointer(ptr unsafe.Pointer) *QPieLegendMarker {
	var n = new(QPieLegendMarker)
	n.SetPointer(ptr)
	return n
}

//export callbackQPieLegendMarker_Series
func callbackQPieLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieLegendMarker::series"); signal != nil {
		return PointerFromQPieSeries(signal.(func() *QPieSeries)())
	}

	return PointerFromQPieSeries(NewQPieLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QPieLegendMarker) ConnectSeries(f func() *QPieSeries) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieLegendMarker::series", f)
	}
}

func (ptr *QPieLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieLegendMarker::series")
	}
}

func (ptr *QPieLegendMarker) Series() *QPieSeries {
	if ptr.Pointer() != nil {
		return NewQPieSeriesFromPointer(C.QPieLegendMarker_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPieLegendMarker) SeriesDefault() *QPieSeries {
	if ptr.Pointer() != nil {
		return NewQPieSeriesFromPointer(C.QPieLegendMarker_SeriesDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPieLegendMarker) Slice() *QPieSlice {
	if ptr.Pointer() != nil {
		return NewQPieSliceFromPointer(C.QPieLegendMarker_Slice(ptr.Pointer()))
	}
	return nil
}

//export callbackQPieLegendMarker_DestroyQPieLegendMarker
func callbackQPieLegendMarker_DestroyQPieLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieLegendMarker::~QPieLegendMarker"); signal != nil {
		signal.(func())()
	} else {
		NewQPieLegendMarkerFromPointer(ptr).DestroyQPieLegendMarkerDefault()
	}
}

func (ptr *QPieLegendMarker) ConnectDestroyQPieLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieLegendMarker::~QPieLegendMarker", f)
	}
}

func (ptr *QPieLegendMarker) DisconnectDestroyQPieLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieLegendMarker::~QPieLegendMarker")
	}
}

func (ptr *QPieLegendMarker) DestroyQPieLegendMarker() {
	if ptr.Pointer() != nil {
		C.QPieLegendMarker_DestroyQPieLegendMarker(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QPieLegendMarker) DestroyQPieLegendMarkerDefault() {
	if ptr.Pointer() != nil {
		C.QPieLegendMarker_DestroyQPieLegendMarkerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QPieSeries struct {
	ptr unsafe.Pointer
}

type QPieSeries_ITF interface {
	QPieSeries_PTR() *QPieSeries
}

func (ptr *QPieSeries) QPieSeries_PTR() *QPieSeries {
	return ptr
}

func (ptr *QPieSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPieSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPieSeries(ptr QPieSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPieSeries_PTR().Pointer()
	}
	return nil
}

func NewQPieSeriesFromPointer(ptr unsafe.Pointer) *QPieSeries {
	var n = new(QPieSeries)
	n.SetPointer(ptr)
	return n
}
func NewQPieSeries(parent core.QObject_ITF) *QPieSeries {
	return NewQPieSeriesFromPointer(C.QPieSeries_NewQPieSeries(core.PointerFromQObject(parent)))
}

func (ptr *QPieSeries) Append3(label string, value float64) *QPieSlice {
	if ptr.Pointer() != nil {
		var labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
		return NewQPieSliceFromPointer(C.QPieSeries_Append3(ptr.Pointer(), labelC, C.double(value)))
	}
	return nil
}

func (ptr *QPieSeries) Append2(slices []*QPieSlice) bool {
	if ptr.Pointer() != nil {
		return C.QPieSeries_Append2(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQPieSeriesFromPointer(NewQPieSeriesFromPointer(unsafe.Pointer(uintptr(1))).__append_slices_newList2())
			for _, v := range slices {
				tmpList.__append_slices_setList2(v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QPieSeries) Append(slice QPieSlice_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPieSeries_Append(ptr.Pointer(), PointerFromQPieSlice(slice)) != 0
	}
	return false
}

func (ptr *QPieSeries) Insert(index int, slice QPieSlice_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPieSeries_Insert(ptr.Pointer(), C.int(int32(index)), PointerFromQPieSlice(slice)) != 0
	}
	return false
}

func (ptr *QPieSeries) Remove(slice QPieSlice_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPieSeries_Remove(ptr.Pointer(), PointerFromQPieSlice(slice)) != 0
	}
	return false
}

func (ptr *QPieSeries) Take(slice QPieSlice_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPieSeries_Take(ptr.Pointer(), PointerFromQPieSlice(slice)) != 0
	}
	return false
}

//export callbackQPieSeries_Added
func callbackQPieSeries_Added(ptr unsafe.Pointer, slices C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::added"); signal != nil {
		signal.(func([]*QPieSlice))(func(l C.struct_QtCharts_PackedList) []*QPieSlice {
			var out = make([]*QPieSlice, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQPieSeriesFromPointer(l.data).__added_slices_atList(i)
			}
			return out
		}(slices))
	}

}

func (ptr *QPieSeries) ConnectAdded(f func(slices []*QPieSlice)) {
	if ptr.Pointer() != nil {
		C.QPieSeries_ConnectAdded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::added", f)
	}
}

func (ptr *QPieSeries) DisconnectAdded() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::added")
	}
}

func (ptr *QPieSeries) Added(slices []*QPieSlice) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Added(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQPieSeriesFromPointer(NewQPieSeriesFromPointer(unsafe.Pointer(uintptr(1))).__added_slices_newList())
			for _, v := range slices {
				tmpList.__added_slices_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QPieSeries) Clear() {
	if ptr.Pointer() != nil {
		C.QPieSeries_Clear(ptr.Pointer())
	}
}

//export callbackQPieSeries_Clicked
func callbackQPieSeries_Clicked(ptr unsafe.Pointer, slice unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::clicked"); signal != nil {
		signal.(func(*QPieSlice))(NewQPieSliceFromPointer(slice))
	}

}

func (ptr *QPieSeries) ConnectClicked(f func(slice *QPieSlice)) {
	if ptr.Pointer() != nil {
		C.QPieSeries_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::clicked", f)
	}
}

func (ptr *QPieSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::clicked")
	}
}

func (ptr *QPieSeries) Clicked(slice QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Clicked(ptr.Pointer(), PointerFromQPieSlice(slice))
	}
}

//export callbackQPieSeries_CountChanged
func callbackQPieSeries_CountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::countChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSeries) ConnectCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSeries_ConnectCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::countChanged", f)
	}
}

func (ptr *QPieSeries) DisconnectCountChanged() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::countChanged")
	}
}

func (ptr *QPieSeries) CountChanged() {
	if ptr.Pointer() != nil {
		C.QPieSeries_CountChanged(ptr.Pointer())
	}
}

//export callbackQPieSeries_DoubleClicked
func callbackQPieSeries_DoubleClicked(ptr unsafe.Pointer, slice unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::doubleClicked"); signal != nil {
		signal.(func(*QPieSlice))(NewQPieSliceFromPointer(slice))
	}

}

func (ptr *QPieSeries) ConnectDoubleClicked(f func(slice *QPieSlice)) {
	if ptr.Pointer() != nil {
		C.QPieSeries_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::doubleClicked", f)
	}
}

func (ptr *QPieSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::doubleClicked")
	}
}

func (ptr *QPieSeries) DoubleClicked(slice QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries_DoubleClicked(ptr.Pointer(), PointerFromQPieSlice(slice))
	}
}

//export callbackQPieSeries_Hovered
func callbackQPieSeries_Hovered(ptr unsafe.Pointer, slice unsafe.Pointer, state C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::hovered"); signal != nil {
		signal.(func(*QPieSlice, bool))(NewQPieSliceFromPointer(slice), int8(state) != 0)
	}

}

func (ptr *QPieSeries) ConnectHovered(f func(slice *QPieSlice, state bool)) {
	if ptr.Pointer() != nil {
		C.QPieSeries_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::hovered", f)
	}
}

func (ptr *QPieSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::hovered")
	}
}

func (ptr *QPieSeries) Hovered(slice QPieSlice_ITF, state bool) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Hovered(ptr.Pointer(), PointerFromQPieSlice(slice), C.char(int8(qt.GoBoolToInt(state))))
	}
}

//export callbackQPieSeries_Pressed
func callbackQPieSeries_Pressed(ptr unsafe.Pointer, slice unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::pressed"); signal != nil {
		signal.(func(*QPieSlice))(NewQPieSliceFromPointer(slice))
	}

}

func (ptr *QPieSeries) ConnectPressed(f func(slice *QPieSlice)) {
	if ptr.Pointer() != nil {
		C.QPieSeries_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::pressed", f)
	}
}

func (ptr *QPieSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::pressed")
	}
}

func (ptr *QPieSeries) Pressed(slice QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Pressed(ptr.Pointer(), PointerFromQPieSlice(slice))
	}
}

//export callbackQPieSeries_Released
func callbackQPieSeries_Released(ptr unsafe.Pointer, slice unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::released"); signal != nil {
		signal.(func(*QPieSlice))(NewQPieSliceFromPointer(slice))
	}

}

func (ptr *QPieSeries) ConnectReleased(f func(slice *QPieSlice)) {
	if ptr.Pointer() != nil {
		C.QPieSeries_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::released", f)
	}
}

func (ptr *QPieSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::released")
	}
}

func (ptr *QPieSeries) Released(slice QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Released(ptr.Pointer(), PointerFromQPieSlice(slice))
	}
}

//export callbackQPieSeries_Removed
func callbackQPieSeries_Removed(ptr unsafe.Pointer, slices C.struct_QtCharts_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::removed"); signal != nil {
		signal.(func([]*QPieSlice))(func(l C.struct_QtCharts_PackedList) []*QPieSlice {
			var out = make([]*QPieSlice, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQPieSeriesFromPointer(l.data).__removed_slices_atList(i)
			}
			return out
		}(slices))
	}

}

func (ptr *QPieSeries) ConnectRemoved(f func(slices []*QPieSlice)) {
	if ptr.Pointer() != nil {
		C.QPieSeries_ConnectRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::removed", f)
	}
}

func (ptr *QPieSeries) DisconnectRemoved() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::removed")
	}
}

func (ptr *QPieSeries) Removed(slices []*QPieSlice) {
	if ptr.Pointer() != nil {
		C.QPieSeries_Removed(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQPieSeriesFromPointer(NewQPieSeriesFromPointer(unsafe.Pointer(uintptr(1))).__removed_slices_newList())
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

//export callbackQPieSeries_SumChanged
func callbackQPieSeries_SumChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::sumChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSeries) ConnectSumChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSeries_ConnectSumChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::sumChanged", f)
	}
}

func (ptr *QPieSeries) DisconnectSumChanged() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DisconnectSumChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::sumChanged")
	}
}

func (ptr *QPieSeries) SumChanged() {
	if ptr.Pointer() != nil {
		C.QPieSeries_SumChanged(ptr.Pointer())
	}
}

//export callbackQPieSeries_DestroyQPieSeries
func callbackQPieSeries_DestroyQPieSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::~QPieSeries"); signal != nil {
		signal.(func())()
	} else {
		NewQPieSeriesFromPointer(ptr).DestroyQPieSeriesDefault()
	}
}

func (ptr *QPieSeries) ConnectDestroyQPieSeries(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::~QPieSeries", f)
	}
}

func (ptr *QPieSeries) DisconnectDestroyQPieSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::~QPieSeries")
	}
}

func (ptr *QPieSeries) DestroyQPieSeries() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DestroyQPieSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QPieSeries) DestroyQPieSeriesDefault() {
	if ptr.Pointer() != nil {
		C.QPieSeries_DestroyQPieSeriesDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQPieSeries_Type
func callbackQPieSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQPieSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QPieSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::type", f)
	}
}

func (ptr *QPieSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSeries::type")
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

func (ptr *QPieSeries) Slices() []*QPieSlice {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QPieSlice {
			var out = make([]*QPieSlice, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQPieSeriesFromPointer(l.data).__slices_atList(i)
			}
			return out
		}(C.QPieSeries_Slices(ptr.Pointer()))
	}
	return make([]*QPieSlice, 0)
}

func (ptr *QPieSeries) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QPieSeries_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPieSeries) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPieSeries_Count(ptr.Pointer())))
	}
	return 0
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

func (ptr *QPieSeries) Sum() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSeries_Sum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSeries) VerticalPosition() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSeries_VerticalPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSeries) __append_slices_atList2(i int) *QPieSlice {
	if ptr.Pointer() != nil {
		return NewQPieSliceFromPointer(C.QPieSeries___append_slices_atList2(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QPieSeries) __append_slices_setList2(i QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries___append_slices_setList2(ptr.Pointer(), PointerFromQPieSlice(i))
	}
}

func (ptr *QPieSeries) __append_slices_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QPieSeries___append_slices_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPieSeries) __added_slices_atList(i int) *QPieSlice {
	if ptr.Pointer() != nil {
		return NewQPieSliceFromPointer(C.QPieSeries___added_slices_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QPieSeries) __added_slices_setList(i QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries___added_slices_setList(ptr.Pointer(), PointerFromQPieSlice(i))
	}
}

func (ptr *QPieSeries) __added_slices_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QPieSeries___added_slices_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPieSeries) __removed_slices_atList(i int) *QPieSlice {
	if ptr.Pointer() != nil {
		return NewQPieSliceFromPointer(C.QPieSeries___removed_slices_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QPieSeries) __removed_slices_setList(i QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries___removed_slices_setList(ptr.Pointer(), PointerFromQPieSlice(i))
	}
}

func (ptr *QPieSeries) __removed_slices_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QPieSeries___removed_slices_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPieSeries) __slices_atList(i int) *QPieSlice {
	if ptr.Pointer() != nil {
		return NewQPieSliceFromPointer(C.QPieSeries___slices_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QPieSeries) __slices_setList(i QPieSlice_ITF) {
	if ptr.Pointer() != nil {
		C.QPieSeries___slices_setList(ptr.Pointer(), PointerFromQPieSlice(i))
	}
}

func (ptr *QPieSeries) __slices_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QPieSeries___slices_newList(ptr.Pointer()))
	}
	return nil
}

type QPieSlice struct {
	ptr unsafe.Pointer
}

type QPieSlice_ITF interface {
	QPieSlice_PTR() *QPieSlice
}

func (ptr *QPieSlice) QPieSlice_PTR() *QPieSlice {
	return ptr
}

func (ptr *QPieSlice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPieSlice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPieSlice(ptr QPieSlice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPieSlice_PTR().Pointer()
	}
	return nil
}

func NewQPieSliceFromPointer(ptr unsafe.Pointer) *QPieSlice {
	var n = new(QPieSlice)
	n.SetPointer(ptr)
	return n
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

func (ptr *QPieSlice) LabelPosition() QPieSlice__LabelPosition {
	if ptr.Pointer() != nil {
		return QPieSlice__LabelPosition(C.QPieSlice_LabelPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSlice) BorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QPieSlice_BorderColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QPieSlice_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) LabelColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QPieSlice_LabelColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func NewQPieSlice(parent core.QObject_ITF) *QPieSlice {
	return NewQPieSliceFromPointer(C.QPieSlice_NewQPieSlice(core.PointerFromQObject(parent)))
}

func NewQPieSlice2(label string, value float64, parent core.QObject_ITF) *QPieSlice {
	var labelC = C.CString(label)
	defer C.free(unsafe.Pointer(labelC))
	return NewQPieSliceFromPointer(C.QPieSlice_NewQPieSlice2(labelC, C.double(value), core.PointerFromQObject(parent)))
}

func (ptr *QPieSlice) BorderWidth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPieSlice_BorderWidth(ptr.Pointer())))
	}
	return 0
}

//export callbackQPieSlice_AngleSpanChanged
func callbackQPieSlice_AngleSpanChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::angleSpanChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectAngleSpanChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectAngleSpanChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::angleSpanChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectAngleSpanChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectAngleSpanChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::angleSpanChanged")
	}
}

func (ptr *QPieSlice) AngleSpanChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_AngleSpanChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_BorderColorChanged
func callbackQPieSlice_BorderColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::borderColorChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectBorderColorChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectBorderColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::borderColorChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::borderColorChanged")
	}
}

func (ptr *QPieSlice) BorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_BorderColorChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_BorderWidthChanged
func callbackQPieSlice_BorderWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::borderWidthChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectBorderWidthChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectBorderWidthChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::borderWidthChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectBorderWidthChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectBorderWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::borderWidthChanged")
	}
}

func (ptr *QPieSlice) BorderWidthChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_BorderWidthChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_BrushChanged
func callbackQPieSlice_BrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::brushChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectBrushChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::brushChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectBrushChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::brushChanged")
	}
}

func (ptr *QPieSlice) BrushChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_BrushChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_Clicked
func callbackQPieSlice_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::clicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::clicked", f)
	}
}

func (ptr *QPieSlice) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::clicked")
	}
}

func (ptr *QPieSlice) Clicked() {
	if ptr.Pointer() != nil {
		C.QPieSlice_Clicked(ptr.Pointer())
	}
}

//export callbackQPieSlice_ColorChanged
func callbackQPieSlice_ColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::colorChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectColorChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::colorChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::colorChanged")
	}
}

func (ptr *QPieSlice) ColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_ColorChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_DoubleClicked
func callbackQPieSlice_DoubleClicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::doubleClicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectDoubleClicked(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::doubleClicked", f)
	}
}

func (ptr *QPieSlice) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::doubleClicked")
	}
}

func (ptr *QPieSlice) DoubleClicked() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DoubleClicked(ptr.Pointer())
	}
}

//export callbackQPieSlice_Hovered
func callbackQPieSlice_Hovered(ptr unsafe.Pointer, state C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::hovered"); signal != nil {
		signal.(func(bool))(int8(state) != 0)
	}

}

func (ptr *QPieSlice) ConnectHovered(f func(state bool)) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::hovered", f)
	}
}

func (ptr *QPieSlice) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::hovered")
	}
}

func (ptr *QPieSlice) Hovered(state bool) {
	if ptr.Pointer() != nil {
		C.QPieSlice_Hovered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(state))))
	}
}

//export callbackQPieSlice_LabelBrushChanged
func callbackQPieSlice_LabelBrushChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::labelBrushChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectLabelBrushChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectLabelBrushChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::labelBrushChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectLabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectLabelBrushChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::labelBrushChanged")
	}
}

func (ptr *QPieSlice) LabelBrushChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_LabelBrushChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_LabelChanged
func callbackQPieSlice_LabelChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::labelChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectLabelChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectLabelChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::labelChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectLabelChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectLabelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::labelChanged")
	}
}

func (ptr *QPieSlice) LabelChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_LabelChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_LabelColorChanged
func callbackQPieSlice_LabelColorChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::labelColorChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectLabelColorChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectLabelColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::labelColorChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectLabelColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectLabelColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::labelColorChanged")
	}
}

func (ptr *QPieSlice) LabelColorChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_LabelColorChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_LabelFontChanged
func callbackQPieSlice_LabelFontChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::labelFontChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectLabelFontChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectLabelFontChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::labelFontChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectLabelFontChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectLabelFontChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::labelFontChanged")
	}
}

func (ptr *QPieSlice) LabelFontChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_LabelFontChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_LabelVisibleChanged
func callbackQPieSlice_LabelVisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::labelVisibleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectLabelVisibleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectLabelVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::labelVisibleChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectLabelVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectLabelVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::labelVisibleChanged")
	}
}

func (ptr *QPieSlice) LabelVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_LabelVisibleChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_PenChanged
func callbackQPieSlice_PenChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::penChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectPenChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectPenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::penChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::penChanged")
	}
}

func (ptr *QPieSlice) PenChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_PenChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_PercentageChanged
func callbackQPieSlice_PercentageChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::percentageChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectPercentageChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectPercentageChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::percentageChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectPercentageChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectPercentageChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::percentageChanged")
	}
}

func (ptr *QPieSlice) PercentageChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_PercentageChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_Pressed
func callbackQPieSlice_Pressed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::pressed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectPressed(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::pressed", f)
	}
}

func (ptr *QPieSlice) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::pressed")
	}
}

func (ptr *QPieSlice) Pressed() {
	if ptr.Pointer() != nil {
		C.QPieSlice_Pressed(ptr.Pointer())
	}
}

//export callbackQPieSlice_Released
func callbackQPieSlice_Released(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::released"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectReleased(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::released", f)
	}
}

func (ptr *QPieSlice) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::released")
	}
}

func (ptr *QPieSlice) Released() {
	if ptr.Pointer() != nil {
		C.QPieSlice_Released(ptr.Pointer())
	}
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
		var labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
		C.QPieSlice_SetLabel(ptr.Pointer(), labelC)
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

//export callbackQPieSlice_StartAngleChanged
func callbackQPieSlice_StartAngleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::startAngleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectStartAngleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectStartAngleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::startAngleChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectStartAngleChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectStartAngleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::startAngleChanged")
	}
}

func (ptr *QPieSlice) StartAngleChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_StartAngleChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_ValueChanged
func callbackQPieSlice_ValueChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::valueChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPieSlice) ConnectValueChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPieSlice_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::valueChanged", f)
	}
}

func (ptr *QPieSlice) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::valueChanged")
	}
}

func (ptr *QPieSlice) ValueChanged() {
	if ptr.Pointer() != nil {
		C.QPieSlice_ValueChanged(ptr.Pointer())
	}
}

//export callbackQPieSlice_DestroyQPieSlice
func callbackQPieSlice_DestroyQPieSlice(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPieSlice::~QPieSlice"); signal != nil {
		signal.(func())()
	} else {
		NewQPieSliceFromPointer(ptr).DestroyQPieSliceDefault()
	}
}

func (ptr *QPieSlice) ConnectDestroyQPieSlice(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::~QPieSlice", f)
	}
}

func (ptr *QPieSlice) DisconnectDestroyQPieSlice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPieSlice::~QPieSlice")
	}
}

func (ptr *QPieSlice) DestroyQPieSlice() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DestroyQPieSlice(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QPieSlice) DestroyQPieSliceDefault() {
	if ptr.Pointer() != nil {
		C.QPieSlice_DestroyQPieSliceDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QPieSlice) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QPieSlice_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) LabelBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QPieSlice_LabelBrush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) LabelFont() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.QPieSlice_LabelFont(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QPieSlice_Pen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QPieSlice) Series() *QPieSeries {
	if ptr.Pointer() != nil {
		return NewQPieSeriesFromPointer(C.QPieSlice_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPieSlice) Label() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPieSlice_Label(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPieSlice) IsExploded() bool {
	if ptr.Pointer() != nil {
		return C.QPieSlice_IsExploded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPieSlice) IsLabelVisible() bool {
	if ptr.Pointer() != nil {
		return C.QPieSlice_IsLabelVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPieSlice) AngleSpan() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_AngleSpan(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSlice) ExplodeDistanceFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_ExplodeDistanceFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSlice) LabelArmLengthFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_LabelArmLengthFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSlice) Percentage() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_Percentage(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSlice) StartAngle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_StartAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPieSlice) Value() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPieSlice_Value(ptr.Pointer()))
	}
	return 0
}

type QPolarChart struct {
	ptr unsafe.Pointer
}

type QPolarChart_ITF interface {
	QPolarChart_PTR() *QPolarChart
}

func (ptr *QPolarChart) QPolarChart_PTR() *QPolarChart {
	return ptr
}

func (ptr *QPolarChart) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPolarChart) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPolarChart(ptr QPolarChart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPolarChart_PTR().Pointer()
	}
	return nil
}

func NewQPolarChartFromPointer(ptr unsafe.Pointer) *QPolarChart {
	var n = new(QPolarChart)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QPolarChart__PolarOrientation
//QPolarChart::PolarOrientation
type QPolarChart__PolarOrientation int64

const (
	QPolarChart__PolarOrientationRadial  QPolarChart__PolarOrientation = QPolarChart__PolarOrientation(0x1)
	QPolarChart__PolarOrientationAngular QPolarChart__PolarOrientation = QPolarChart__PolarOrientation(0x2)
)

func QPolarChart_AxisPolarOrientation(axis QAbstractAxis_ITF) QPolarChart__PolarOrientation {
	return QPolarChart__PolarOrientation(C.QPolarChart_QPolarChart_AxisPolarOrientation(PointerFromQAbstractAxis(axis)))
}

func (ptr *QPolarChart) AxisPolarOrientation(axis QAbstractAxis_ITF) QPolarChart__PolarOrientation {
	return QPolarChart__PolarOrientation(C.QPolarChart_QPolarChart_AxisPolarOrientation(PointerFromQAbstractAxis(axis)))
}

func NewQPolarChart(parent widgets.QGraphicsItem_ITF, wFlags core.Qt__WindowType) *QPolarChart {
	var tmpValue = NewQPolarChartFromPointer(C.QPolarChart_NewQPolarChart(widgets.PointerFromQGraphicsItem(parent), C.longlong(wFlags)))
	runtime.SetFinalizer(tmpValue, (*QPolarChart).DestroyQPolarChart)
	return tmpValue
}

func (ptr *QPolarChart) AddAxis(axis QAbstractAxis_ITF, polarOrientation QPolarChart__PolarOrientation) {
	if ptr.Pointer() != nil {
		C.QPolarChart_AddAxis(ptr.Pointer(), PointerFromQAbstractAxis(axis), C.longlong(polarOrientation))
	}
}

func (ptr *QPolarChart) DestroyQPolarChart() {
	if ptr.Pointer() != nil {
		C.QPolarChart_DestroyQPolarChart(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPolarChart) Axes(polarOrientation QPolarChart__PolarOrientation, series QAbstractSeries_ITF) []*QAbstractAxis {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*QAbstractAxis {
			var out = make([]*QAbstractAxis, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQPolarChartFromPointer(l.data).__axes_atList(i)
			}
			return out
		}(C.QPolarChart_Axes(ptr.Pointer(), C.longlong(polarOrientation), PointerFromQAbstractSeries(series)))
	}
	return make([]*QAbstractAxis, 0)
}

func (ptr *QPolarChart) __axes_atList(i int) *QAbstractAxis {
	if ptr.Pointer() != nil {
		return NewQAbstractAxisFromPointer(C.QPolarChart___axes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QPolarChart) __axes_setList(i QAbstractAxis_ITF) {
	if ptr.Pointer() != nil {
		C.QPolarChart___axes_setList(ptr.Pointer(), PointerFromQAbstractAxis(i))
	}
}

func (ptr *QPolarChart) __axes_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QPolarChart___axes_newList(ptr.Pointer()))
	}
	return nil
}

type QScatterSeries struct {
	ptr unsafe.Pointer
}

type QScatterSeries_ITF interface {
	QScatterSeries_PTR() *QScatterSeries
}

func (ptr *QScatterSeries) QScatterSeries_PTR() *QScatterSeries {
	return ptr
}

func (ptr *QScatterSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScatterSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScatterSeries(ptr QScatterSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScatterSeries_PTR().Pointer()
	}
	return nil
}

func NewQScatterSeriesFromPointer(ptr unsafe.Pointer) *QScatterSeries {
	var n = new(QScatterSeries)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QScatterSeries__MarkerShape
//QScatterSeries::MarkerShape
type QScatterSeries__MarkerShape int64

const (
	QScatterSeries__MarkerShapeCircle    QScatterSeries__MarkerShape = QScatterSeries__MarkerShape(0)
	QScatterSeries__MarkerShapeRectangle QScatterSeries__MarkerShape = QScatterSeries__MarkerShape(1)
)

func (ptr *QScatterSeries) BorderColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QScatterSeries_BorderColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func NewQScatterSeries(parent core.QObject_ITF) *QScatterSeries {
	return NewQScatterSeriesFromPointer(C.QScatterSeries_NewQScatterSeries(core.PointerFromQObject(parent)))
}

//export callbackQScatterSeries_BorderColorChanged
func callbackQScatterSeries_BorderColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScatterSeries::borderColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QScatterSeries) ConnectBorderColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_ConnectBorderColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::borderColorChanged", f)
	}
}

func (ptr *QScatterSeries) DisconnectBorderColorChanged() {
	if ptr.Pointer() != nil {
		C.QScatterSeries_DisconnectBorderColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::borderColorChanged")
	}
}

func (ptr *QScatterSeries) BorderColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_BorderColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQScatterSeries_ColorChanged
func callbackQScatterSeries_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScatterSeries::colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QScatterSeries) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::colorChanged", f)
	}
}

func (ptr *QScatterSeries) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QScatterSeries_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::colorChanged")
	}
}

func (ptr *QScatterSeries) ColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQScatterSeries_MarkerShapeChanged
func callbackQScatterSeries_MarkerShapeChanged(ptr unsafe.Pointer, shape C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScatterSeries::markerShapeChanged"); signal != nil {
		signal.(func(QScatterSeries__MarkerShape))(QScatterSeries__MarkerShape(shape))
	}

}

func (ptr *QScatterSeries) ConnectMarkerShapeChanged(f func(shape QScatterSeries__MarkerShape)) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_ConnectMarkerShapeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::markerShapeChanged", f)
	}
}

func (ptr *QScatterSeries) DisconnectMarkerShapeChanged() {
	if ptr.Pointer() != nil {
		C.QScatterSeries_DisconnectMarkerShapeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::markerShapeChanged")
	}
}

func (ptr *QScatterSeries) MarkerShapeChanged(shape QScatterSeries__MarkerShape) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_MarkerShapeChanged(ptr.Pointer(), C.longlong(shape))
	}
}

//export callbackQScatterSeries_MarkerSizeChanged
func callbackQScatterSeries_MarkerSizeChanged(ptr unsafe.Pointer, size C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScatterSeries::markerSizeChanged"); signal != nil {
		signal.(func(float64))(float64(size))
	}

}

func (ptr *QScatterSeries) ConnectMarkerSizeChanged(f func(size float64)) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_ConnectMarkerSizeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::markerSizeChanged", f)
	}
}

func (ptr *QScatterSeries) DisconnectMarkerSizeChanged() {
	if ptr.Pointer() != nil {
		C.QScatterSeries_DisconnectMarkerSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::markerSizeChanged")
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

//export callbackQScatterSeries_SetBrush
func callbackQScatterSeries_SetBrush(ptr unsafe.Pointer, brush unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScatterSeries::setBrush"); signal != nil {
		signal.(func(*gui.QBrush))(gui.NewQBrushFromPointer(brush))
	} else {
		NewQScatterSeriesFromPointer(ptr).SetBrushDefault(gui.NewQBrushFromPointer(brush))
	}
}

func (ptr *QScatterSeries) ConnectSetBrush(f func(brush *gui.QBrush)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::setBrush", f)
	}
}

func (ptr *QScatterSeries) DisconnectSetBrush() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::setBrush")
	}
}

func (ptr *QScatterSeries) SetBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QScatterSeries) SetBrushDefault(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_SetBrushDefault(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

//export callbackQScatterSeries_SetColor
func callbackQScatterSeries_SetColor(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScatterSeries::setColor"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	} else {
		NewQScatterSeriesFromPointer(ptr).SetColorDefault(gui.NewQColorFromPointer(color))
	}
}

func (ptr *QScatterSeries) ConnectSetColor(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::setColor", f)
	}
}

func (ptr *QScatterSeries) DisconnectSetColor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::setColor")
	}
}

func (ptr *QScatterSeries) SetColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QScatterSeries) SetColorDefault(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_SetColorDefault(ptr.Pointer(), gui.PointerFromQColor(color))
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

//export callbackQScatterSeries_SetPen
func callbackQScatterSeries_SetPen(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScatterSeries::setPen"); signal != nil {
		signal.(func(*gui.QPen))(gui.NewQPenFromPointer(pen))
	} else {
		NewQScatterSeriesFromPointer(ptr).SetPenDefault(gui.NewQPenFromPointer(pen))
	}
}

func (ptr *QScatterSeries) ConnectSetPen(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::setPen", f)
	}
}

func (ptr *QScatterSeries) DisconnectSetPen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::setPen")
	}
}

func (ptr *QScatterSeries) SetPen(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QScatterSeries) SetPenDefault(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterSeries_SetPenDefault(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QScatterSeries) DestroyQScatterSeries() {
	if ptr.Pointer() != nil {
		C.QScatterSeries_DestroyQScatterSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScatterSeries) MarkerShape() QScatterSeries__MarkerShape {
	if ptr.Pointer() != nil {
		return QScatterSeries__MarkerShape(C.QScatterSeries_MarkerShape(ptr.Pointer()))
	}
	return 0
}

//export callbackQScatterSeries_Type
func callbackQScatterSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScatterSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQScatterSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QScatterSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::type", f)
	}
}

func (ptr *QScatterSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::type")
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

func (ptr *QScatterSeries) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QScatterSeries_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQScatterSeries_Color
func callbackQScatterSeries_Color(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScatterSeries::color"); signal != nil {
		return gui.PointerFromQColor(signal.(func() *gui.QColor)())
	}

	return gui.PointerFromQColor(NewQScatterSeriesFromPointer(ptr).ColorDefault())
}

func (ptr *QScatterSeries) ConnectColor(f func() *gui.QColor) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::color", f)
	}
}

func (ptr *QScatterSeries) DisconnectColor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScatterSeries::color")
	}
}

func (ptr *QScatterSeries) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QScatterSeries_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QScatterSeries) ColorDefault() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QScatterSeries_ColorDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QScatterSeries) MarkerSize() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QScatterSeries_MarkerSize(ptr.Pointer()))
	}
	return 0
}

type QSplineSeries struct {
	ptr unsafe.Pointer
}

type QSplineSeries_ITF interface {
	QSplineSeries_PTR() *QSplineSeries
}

func (ptr *QSplineSeries) QSplineSeries_PTR() *QSplineSeries {
	return ptr
}

func (ptr *QSplineSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSplineSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSplineSeries(ptr QSplineSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplineSeries_PTR().Pointer()
	}
	return nil
}

func NewQSplineSeriesFromPointer(ptr unsafe.Pointer) *QSplineSeries {
	var n = new(QSplineSeries)
	n.SetPointer(ptr)
	return n
}
func NewQSplineSeries(parent core.QObject_ITF) *QSplineSeries {
	return NewQSplineSeriesFromPointer(C.QSplineSeries_NewQSplineSeries(core.PointerFromQObject(parent)))
}

func (ptr *QSplineSeries) DestroyQSplineSeries() {
	if ptr.Pointer() != nil {
		C.QSplineSeries_DestroyQSplineSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSplineSeries_Type
func callbackQSplineSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSplineSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQSplineSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QSplineSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSplineSeries::type", f)
	}
}

func (ptr *QSplineSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSplineSeries::type")
	}
}

func (ptr *QSplineSeries) Type() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QSplineSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplineSeries) TypeDefault() QAbstractSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstractSeries__SeriesType(C.QSplineSeries_TypeDefault(ptr.Pointer()))
	}
	return 0
}

type QStackedBarSeries struct {
	ptr unsafe.Pointer
}

type QStackedBarSeries_ITF interface {
	QStackedBarSeries_PTR() *QStackedBarSeries
}

func (ptr *QStackedBarSeries) QStackedBarSeries_PTR() *QStackedBarSeries {
	return ptr
}

func (ptr *QStackedBarSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QStackedBarSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQStackedBarSeries(ptr QStackedBarSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStackedBarSeries_PTR().Pointer()
	}
	return nil
}

func NewQStackedBarSeriesFromPointer(ptr unsafe.Pointer) *QStackedBarSeries {
	var n = new(QStackedBarSeries)
	n.SetPointer(ptr)
	return n
}
func NewQStackedBarSeries(parent core.QObject_ITF) *QStackedBarSeries {
	return NewQStackedBarSeriesFromPointer(C.QStackedBarSeries_NewQStackedBarSeries(core.PointerFromQObject(parent)))
}

func (ptr *QStackedBarSeries) DestroyQStackedBarSeries() {
	if ptr.Pointer() != nil {
		C.QStackedBarSeries_DestroyQStackedBarSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQStackedBarSeries_Type
func callbackQStackedBarSeries_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QStackedBarSeries::type"); signal != nil {
		return C.longlong(signal.(func() QAbstractSeries__SeriesType)())
	}

	return C.longlong(NewQStackedBarSeriesFromPointer(ptr).TypeDefault())
}

func (ptr *QStackedBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QStackedBarSeries::type", f)
	}
}

func (ptr *QStackedBarSeries) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QStackedBarSeries::type")
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

type QVBarModelMapper struct {
	ptr unsafe.Pointer
}

type QVBarModelMapper_ITF interface {
	QVBarModelMapper_PTR() *QVBarModelMapper
}

func (ptr *QVBarModelMapper) QVBarModelMapper_PTR() *QVBarModelMapper {
	return ptr
}

func (ptr *QVBarModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVBarModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVBarModelMapper(ptr QVBarModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBarModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQVBarModelMapperFromPointer(ptr unsafe.Pointer) *QVBarModelMapper {
	var n = new(QVBarModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVBarModelMapper) DestroyQVBarModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQVBarModelMapper(parent core.QObject_ITF) *QVBarModelMapper {
	return NewQVBarModelMapperFromPointer(C.QVBarModelMapper_NewQVBarModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQVBarModelMapper_FirstBarSetColumnChanged
func callbackQVBarModelMapper_FirstBarSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBarModelMapper::firstBarSetColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBarModelMapper) ConnectFirstBarSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_ConnectFirstBarSetColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::firstBarSetColumnChanged", f)
	}
}

func (ptr *QVBarModelMapper) DisconnectFirstBarSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectFirstBarSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::firstBarSetColumnChanged")
	}
}

func (ptr *QVBarModelMapper) FirstBarSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_FirstBarSetColumnChanged(ptr.Pointer())
	}
}

//export callbackQVBarModelMapper_FirstRowChanged
func callbackQVBarModelMapper_FirstRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBarModelMapper::firstRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBarModelMapper) ConnectFirstRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_ConnectFirstRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::firstRowChanged", f)
	}
}

func (ptr *QVBarModelMapper) DisconnectFirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectFirstRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::firstRowChanged")
	}
}

func (ptr *QVBarModelMapper) FirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_FirstRowChanged(ptr.Pointer())
	}
}

//export callbackQVBarModelMapper_LastBarSetColumnChanged
func callbackQVBarModelMapper_LastBarSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBarModelMapper::lastBarSetColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBarModelMapper) ConnectLastBarSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_ConnectLastBarSetColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::lastBarSetColumnChanged", f)
	}
}

func (ptr *QVBarModelMapper) DisconnectLastBarSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectLastBarSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::lastBarSetColumnChanged")
	}
}

func (ptr *QVBarModelMapper) LastBarSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_LastBarSetColumnChanged(ptr.Pointer())
	}
}

//export callbackQVBarModelMapper_ModelReplaced
func callbackQVBarModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBarModelMapper::modelReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBarModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_ConnectModelReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::modelReplaced", f)
	}
}

func (ptr *QVBarModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::modelReplaced")
	}
}

func (ptr *QVBarModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_ModelReplaced(ptr.Pointer())
	}
}

//export callbackQVBarModelMapper_RowCountChanged
func callbackQVBarModelMapper_RowCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBarModelMapper::rowCountChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBarModelMapper) ConnectRowCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_ConnectRowCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::rowCountChanged", f)
	}
}

func (ptr *QVBarModelMapper) DisconnectRowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectRowCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::rowCountChanged")
	}
}

func (ptr *QVBarModelMapper) RowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_RowCountChanged(ptr.Pointer())
	}
}

//export callbackQVBarModelMapper_SeriesReplaced
func callbackQVBarModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBarModelMapper::seriesReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBarModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_ConnectSeriesReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::seriesReplaced", f)
	}
}

func (ptr *QVBarModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVBarModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBarModelMapper::seriesReplaced")
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

func (ptr *QVBarModelMapper) Series() *QAbstractBarSeries {
	if ptr.Pointer() != nil {
		return NewQAbstractBarSeriesFromPointer(C.QVBarModelMapper_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVBarModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QVBarModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBarModelMapper) FirstBarSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBarModelMapper_FirstBarSetColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVBarModelMapper) FirstRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBarModelMapper_FirstRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVBarModelMapper) LastBarSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBarModelMapper_LastBarSetColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVBarModelMapper) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBarModelMapper_RowCount(ptr.Pointer())))
	}
	return 0
}

type QVBoxPlotModelMapper struct {
	ptr unsafe.Pointer
}

type QVBoxPlotModelMapper_ITF interface {
	QVBoxPlotModelMapper_PTR() *QVBoxPlotModelMapper
}

func (ptr *QVBoxPlotModelMapper) QVBoxPlotModelMapper_PTR() *QVBoxPlotModelMapper {
	return ptr
}

func (ptr *QVBoxPlotModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVBoxPlotModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVBoxPlotModelMapper(ptr QVBoxPlotModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBoxPlotModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQVBoxPlotModelMapperFromPointer(ptr unsafe.Pointer) *QVBoxPlotModelMapper {
	var n = new(QVBoxPlotModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVBoxPlotModelMapper) DestroyQVBoxPlotModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQVBoxPlotModelMapper(parent core.QObject_ITF) *QVBoxPlotModelMapper {
	return NewQVBoxPlotModelMapperFromPointer(C.QVBoxPlotModelMapper_NewQVBoxPlotModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQVBoxPlotModelMapper_FirstBoxSetColumnChanged
func callbackQVBoxPlotModelMapper_FirstBoxSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBoxPlotModelMapper::firstBoxSetColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectFirstBoxSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_ConnectFirstBoxSetColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::firstBoxSetColumnChanged", f)
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectFirstBoxSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectFirstBoxSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::firstBoxSetColumnChanged")
	}
}

func (ptr *QVBoxPlotModelMapper) FirstBoxSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_FirstBoxSetColumnChanged(ptr.Pointer())
	}
}

//export callbackQVBoxPlotModelMapper_FirstRowChanged
func callbackQVBoxPlotModelMapper_FirstRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBoxPlotModelMapper::firstRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectFirstRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_ConnectFirstRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::firstRowChanged", f)
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectFirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectFirstRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::firstRowChanged")
	}
}

func (ptr *QVBoxPlotModelMapper) FirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_FirstRowChanged(ptr.Pointer())
	}
}

//export callbackQVBoxPlotModelMapper_LastBoxSetColumnChanged
func callbackQVBoxPlotModelMapper_LastBoxSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBoxPlotModelMapper::lastBoxSetColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectLastBoxSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_ConnectLastBoxSetColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::lastBoxSetColumnChanged", f)
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectLastBoxSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectLastBoxSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::lastBoxSetColumnChanged")
	}
}

func (ptr *QVBoxPlotModelMapper) LastBoxSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_LastBoxSetColumnChanged(ptr.Pointer())
	}
}

//export callbackQVBoxPlotModelMapper_ModelReplaced
func callbackQVBoxPlotModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBoxPlotModelMapper::modelReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_ConnectModelReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::modelReplaced", f)
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::modelReplaced")
	}
}

func (ptr *QVBoxPlotModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_ModelReplaced(ptr.Pointer())
	}
}

//export callbackQVBoxPlotModelMapper_RowCountChanged
func callbackQVBoxPlotModelMapper_RowCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBoxPlotModelMapper::rowCountChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectRowCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_ConnectRowCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::rowCountChanged", f)
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectRowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectRowCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::rowCountChanged")
	}
}

func (ptr *QVBoxPlotModelMapper) RowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_RowCountChanged(ptr.Pointer())
	}
}

//export callbackQVBoxPlotModelMapper_SeriesReplaced
func callbackQVBoxPlotModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVBoxPlotModelMapper::seriesReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVBoxPlotModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_ConnectSeriesReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::seriesReplaced", f)
	}
}

func (ptr *QVBoxPlotModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVBoxPlotModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVBoxPlotModelMapper::seriesReplaced")
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

func (ptr *QVBoxPlotModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QVBoxPlotModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBoxPlotModelMapper) Series() *QBoxPlotSeries {
	if ptr.Pointer() != nil {
		return NewQBoxPlotSeriesFromPointer(C.QVBoxPlotModelMapper_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVBoxPlotModelMapper) FirstBoxSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxPlotModelMapper_FirstBoxSetColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVBoxPlotModelMapper) FirstRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxPlotModelMapper_FirstRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVBoxPlotModelMapper) LastBoxSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxPlotModelMapper_LastBoxSetColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVBoxPlotModelMapper) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxPlotModelMapper_RowCount(ptr.Pointer())))
	}
	return 0
}

type QVCandlestickModelMapper struct {
	ptr unsafe.Pointer
}

type QVCandlestickModelMapper_ITF interface {
	QVCandlestickModelMapper_PTR() *QVCandlestickModelMapper
}

func (ptr *QVCandlestickModelMapper) QVCandlestickModelMapper_PTR() *QVCandlestickModelMapper {
	return ptr
}

func (ptr *QVCandlestickModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVCandlestickModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVCandlestickModelMapper(ptr QVCandlestickModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVCandlestickModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQVCandlestickModelMapperFromPointer(ptr unsafe.Pointer) *QVCandlestickModelMapper {
	var n = new(QVCandlestickModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVCandlestickModelMapper) DestroyQVCandlestickModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQVCandlestickModelMapper(parent core.QObject_ITF) *QVCandlestickModelMapper {
	return NewQVCandlestickModelMapperFromPointer(C.QVCandlestickModelMapper_NewQVCandlestickModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQVCandlestickModelMapper_CloseRowChanged
func callbackQVCandlestickModelMapper_CloseRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVCandlestickModelMapper::closeRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectCloseRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_ConnectCloseRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::closeRowChanged", f)
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectCloseRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectCloseRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::closeRowChanged")
	}
}

func (ptr *QVCandlestickModelMapper) CloseRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_CloseRowChanged(ptr.Pointer())
	}
}

//export callbackQVCandlestickModelMapper_FirstSetColumnChanged
func callbackQVCandlestickModelMapper_FirstSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVCandlestickModelMapper::firstSetColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectFirstSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_ConnectFirstSetColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::firstSetColumnChanged", f)
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectFirstSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectFirstSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::firstSetColumnChanged")
	}
}

func (ptr *QVCandlestickModelMapper) FirstSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_FirstSetColumnChanged(ptr.Pointer())
	}
}

//export callbackQVCandlestickModelMapper_HighRowChanged
func callbackQVCandlestickModelMapper_HighRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVCandlestickModelMapper::highRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectHighRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_ConnectHighRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::highRowChanged", f)
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectHighRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectHighRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::highRowChanged")
	}
}

func (ptr *QVCandlestickModelMapper) HighRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_HighRowChanged(ptr.Pointer())
	}
}

//export callbackQVCandlestickModelMapper_LastSetColumnChanged
func callbackQVCandlestickModelMapper_LastSetColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVCandlestickModelMapper::lastSetColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectLastSetColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_ConnectLastSetColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::lastSetColumnChanged", f)
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectLastSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectLastSetColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::lastSetColumnChanged")
	}
}

func (ptr *QVCandlestickModelMapper) LastSetColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_LastSetColumnChanged(ptr.Pointer())
	}
}

//export callbackQVCandlestickModelMapper_LowRowChanged
func callbackQVCandlestickModelMapper_LowRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVCandlestickModelMapper::lowRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectLowRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_ConnectLowRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::lowRowChanged", f)
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectLowRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectLowRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::lowRowChanged")
	}
}

func (ptr *QVCandlestickModelMapper) LowRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_LowRowChanged(ptr.Pointer())
	}
}

//export callbackQVCandlestickModelMapper_OpenRowChanged
func callbackQVCandlestickModelMapper_OpenRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVCandlestickModelMapper::openRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectOpenRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_ConnectOpenRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::openRowChanged", f)
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectOpenRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectOpenRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::openRowChanged")
	}
}

func (ptr *QVCandlestickModelMapper) OpenRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_OpenRowChanged(ptr.Pointer())
	}
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

//export callbackQVCandlestickModelMapper_TimestampRowChanged
func callbackQVCandlestickModelMapper_TimestampRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVCandlestickModelMapper::timestampRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVCandlestickModelMapper) ConnectTimestampRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_ConnectTimestampRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::timestampRowChanged", f)
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectTimestampRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_DisconnectTimestampRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::timestampRowChanged")
	}
}

func (ptr *QVCandlestickModelMapper) TimestampRowChanged() {
	if ptr.Pointer() != nil {
		C.QVCandlestickModelMapper_TimestampRowChanged(ptr.Pointer())
	}
}

//export callbackQVCandlestickModelMapper_Orientation
func callbackQVCandlestickModelMapper_Orientation(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVCandlestickModelMapper::orientation"); signal != nil {
		return C.longlong(signal.(func() core.Qt__Orientation)())
	}

	return C.longlong(NewQVCandlestickModelMapperFromPointer(ptr).OrientationDefault())
}

func (ptr *QVCandlestickModelMapper) ConnectOrientation(f func() core.Qt__Orientation) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::orientation", f)
	}
}

func (ptr *QVCandlestickModelMapper) DisconnectOrientation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVCandlestickModelMapper::orientation")
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

func (ptr *QVCandlestickModelMapper) CloseRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_CloseRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVCandlestickModelMapper) FirstSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_FirstSetColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVCandlestickModelMapper) HighRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_HighRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVCandlestickModelMapper) LastSetColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_LastSetColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVCandlestickModelMapper) LowRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_LowRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVCandlestickModelMapper) OpenRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_OpenRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVCandlestickModelMapper) TimestampRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVCandlestickModelMapper_TimestampRow(ptr.Pointer())))
	}
	return 0
}

type QVPieModelMapper struct {
	ptr unsafe.Pointer
}

type QVPieModelMapper_ITF interface {
	QVPieModelMapper_PTR() *QVPieModelMapper
}

func (ptr *QVPieModelMapper) QVPieModelMapper_PTR() *QVPieModelMapper {
	return ptr
}

func (ptr *QVPieModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVPieModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVPieModelMapper(ptr QVPieModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVPieModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQVPieModelMapperFromPointer(ptr unsafe.Pointer) *QVPieModelMapper {
	var n = new(QVPieModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVPieModelMapper) DestroyQVPieModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQVPieModelMapper(parent core.QObject_ITF) *QVPieModelMapper {
	return NewQVPieModelMapperFromPointer(C.QVPieModelMapper_NewQVPieModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQVPieModelMapper_FirstRowChanged
func callbackQVPieModelMapper_FirstRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVPieModelMapper::firstRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVPieModelMapper) ConnectFirstRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_ConnectFirstRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::firstRowChanged", f)
	}
}

func (ptr *QVPieModelMapper) DisconnectFirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectFirstRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::firstRowChanged")
	}
}

func (ptr *QVPieModelMapper) FirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_FirstRowChanged(ptr.Pointer())
	}
}

//export callbackQVPieModelMapper_LabelsColumnChanged
func callbackQVPieModelMapper_LabelsColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVPieModelMapper::labelsColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVPieModelMapper) ConnectLabelsColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_ConnectLabelsColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::labelsColumnChanged", f)
	}
}

func (ptr *QVPieModelMapper) DisconnectLabelsColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectLabelsColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::labelsColumnChanged")
	}
}

func (ptr *QVPieModelMapper) LabelsColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_LabelsColumnChanged(ptr.Pointer())
	}
}

//export callbackQVPieModelMapper_ModelReplaced
func callbackQVPieModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVPieModelMapper::modelReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVPieModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_ConnectModelReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::modelReplaced", f)
	}
}

func (ptr *QVPieModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::modelReplaced")
	}
}

func (ptr *QVPieModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_ModelReplaced(ptr.Pointer())
	}
}

//export callbackQVPieModelMapper_RowCountChanged
func callbackQVPieModelMapper_RowCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVPieModelMapper::rowCountChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVPieModelMapper) ConnectRowCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_ConnectRowCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::rowCountChanged", f)
	}
}

func (ptr *QVPieModelMapper) DisconnectRowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectRowCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::rowCountChanged")
	}
}

func (ptr *QVPieModelMapper) RowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_RowCountChanged(ptr.Pointer())
	}
}

//export callbackQVPieModelMapper_SeriesReplaced
func callbackQVPieModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVPieModelMapper::seriesReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVPieModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_ConnectSeriesReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::seriesReplaced", f)
	}
}

func (ptr *QVPieModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::seriesReplaced")
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

//export callbackQVPieModelMapper_ValuesColumnChanged
func callbackQVPieModelMapper_ValuesColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVPieModelMapper::valuesColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVPieModelMapper) ConnectValuesColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_ConnectValuesColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::valuesColumnChanged", f)
	}
}

func (ptr *QVPieModelMapper) DisconnectValuesColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_DisconnectValuesColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVPieModelMapper::valuesColumnChanged")
	}
}

func (ptr *QVPieModelMapper) ValuesColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVPieModelMapper_ValuesColumnChanged(ptr.Pointer())
	}
}

func (ptr *QVPieModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QVPieModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVPieModelMapper) Series() *QPieSeries {
	if ptr.Pointer() != nil {
		return NewQPieSeriesFromPointer(C.QVPieModelMapper_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVPieModelMapper) FirstRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVPieModelMapper_FirstRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVPieModelMapper) LabelsColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVPieModelMapper_LabelsColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVPieModelMapper) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVPieModelMapper_RowCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVPieModelMapper) ValuesColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVPieModelMapper_ValuesColumn(ptr.Pointer())))
	}
	return 0
}

type QVXYModelMapper struct {
	ptr unsafe.Pointer
}

type QVXYModelMapper_ITF interface {
	QVXYModelMapper_PTR() *QVXYModelMapper
}

func (ptr *QVXYModelMapper) QVXYModelMapper_PTR() *QVXYModelMapper {
	return ptr
}

func (ptr *QVXYModelMapper) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVXYModelMapper) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVXYModelMapper(ptr QVXYModelMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVXYModelMapper_PTR().Pointer()
	}
	return nil
}

func NewQVXYModelMapperFromPointer(ptr unsafe.Pointer) *QVXYModelMapper {
	var n = new(QVXYModelMapper)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVXYModelMapper) DestroyQVXYModelMapper() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQVXYModelMapper(parent core.QObject_ITF) *QVXYModelMapper {
	return NewQVXYModelMapperFromPointer(C.QVXYModelMapper_NewQVXYModelMapper(core.PointerFromQObject(parent)))
}

//export callbackQVXYModelMapper_FirstRowChanged
func callbackQVXYModelMapper_FirstRowChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVXYModelMapper::firstRowChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVXYModelMapper) ConnectFirstRowChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_ConnectFirstRowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::firstRowChanged", f)
	}
}

func (ptr *QVXYModelMapper) DisconnectFirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectFirstRowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::firstRowChanged")
	}
}

func (ptr *QVXYModelMapper) FirstRowChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_FirstRowChanged(ptr.Pointer())
	}
}

//export callbackQVXYModelMapper_ModelReplaced
func callbackQVXYModelMapper_ModelReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVXYModelMapper::modelReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVXYModelMapper) ConnectModelReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_ConnectModelReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::modelReplaced", f)
	}
}

func (ptr *QVXYModelMapper) DisconnectModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectModelReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::modelReplaced")
	}
}

func (ptr *QVXYModelMapper) ModelReplaced() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_ModelReplaced(ptr.Pointer())
	}
}

//export callbackQVXYModelMapper_RowCountChanged
func callbackQVXYModelMapper_RowCountChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVXYModelMapper::rowCountChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVXYModelMapper) ConnectRowCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_ConnectRowCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::rowCountChanged", f)
	}
}

func (ptr *QVXYModelMapper) DisconnectRowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectRowCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::rowCountChanged")
	}
}

func (ptr *QVXYModelMapper) RowCountChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_RowCountChanged(ptr.Pointer())
	}
}

//export callbackQVXYModelMapper_SeriesReplaced
func callbackQVXYModelMapper_SeriesReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVXYModelMapper::seriesReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVXYModelMapper) ConnectSeriesReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_ConnectSeriesReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::seriesReplaced", f)
	}
}

func (ptr *QVXYModelMapper) DisconnectSeriesReplaced() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectSeriesReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::seriesReplaced")
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

//export callbackQVXYModelMapper_XColumnChanged
func callbackQVXYModelMapper_XColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVXYModelMapper::xColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVXYModelMapper) ConnectXColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_ConnectXColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::xColumnChanged", f)
	}
}

func (ptr *QVXYModelMapper) DisconnectXColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectXColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::xColumnChanged")
	}
}

func (ptr *QVXYModelMapper) XColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_XColumnChanged(ptr.Pointer())
	}
}

//export callbackQVXYModelMapper_YColumnChanged
func callbackQVXYModelMapper_YColumnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QVXYModelMapper::yColumnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVXYModelMapper) ConnectYColumnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_ConnectYColumnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::yColumnChanged", f)
	}
}

func (ptr *QVXYModelMapper) DisconnectYColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_DisconnectYColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QVXYModelMapper::yColumnChanged")
	}
}

func (ptr *QVXYModelMapper) YColumnChanged() {
	if ptr.Pointer() != nil {
		C.QVXYModelMapper_YColumnChanged(ptr.Pointer())
	}
}

func (ptr *QVXYModelMapper) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QVXYModelMapper_Model(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVXYModelMapper) Series() *QXYSeries {
	if ptr.Pointer() != nil {
		return NewQXYSeriesFromPointer(C.QVXYModelMapper_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVXYModelMapper) FirstRow() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVXYModelMapper_FirstRow(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVXYModelMapper) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVXYModelMapper_RowCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVXYModelMapper) XColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVXYModelMapper_XColumn(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVXYModelMapper) YColumn() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVXYModelMapper_YColumn(ptr.Pointer())))
	}
	return 0
}

type QValueAxis struct {
	ptr unsafe.Pointer
}

type QValueAxis_ITF interface {
	QValueAxis_PTR() *QValueAxis
}

func (ptr *QValueAxis) QValueAxis_PTR() *QValueAxis {
	return ptr
}

func (ptr *QValueAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QValueAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQValueAxis(ptr QValueAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QValueAxis_PTR().Pointer()
	}
	return nil
}

func NewQValueAxisFromPointer(ptr unsafe.Pointer) *QValueAxis {
	var n = new(QValueAxis)
	n.SetPointer(ptr)
	return n
}
func NewQValueAxis(parent core.QObject_ITF) *QValueAxis {
	return NewQValueAxisFromPointer(C.QValueAxis_NewQValueAxis(core.PointerFromQObject(parent)))
}

//export callbackQValueAxis_ApplyNiceNumbers
func callbackQValueAxis_ApplyNiceNumbers(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QValueAxis::applyNiceNumbers"); signal != nil {
		signal.(func())()
	} else {
		NewQValueAxisFromPointer(ptr).ApplyNiceNumbersDefault()
	}
}

func (ptr *QValueAxis) ConnectApplyNiceNumbers(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::applyNiceNumbers", f)
	}
}

func (ptr *QValueAxis) DisconnectApplyNiceNumbers() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::applyNiceNumbers")
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

//export callbackQValueAxis_LabelFormatChanged
func callbackQValueAxis_LabelFormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QValueAxis::labelFormatChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(format))
	}

}

func (ptr *QValueAxis) ConnectLabelFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {
		C.QValueAxis_ConnectLabelFormatChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::labelFormatChanged", f)
	}
}

func (ptr *QValueAxis) DisconnectLabelFormatChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectLabelFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::labelFormatChanged")
	}
}

func (ptr *QValueAxis) LabelFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QValueAxis_LabelFormatChanged(ptr.Pointer(), formatC)
	}
}

//export callbackQValueAxis_MaxChanged
func callbackQValueAxis_MaxChanged(ptr unsafe.Pointer, max C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QValueAxis::maxChanged"); signal != nil {
		signal.(func(float64))(float64(max))
	}

}

func (ptr *QValueAxis) ConnectMaxChanged(f func(max float64)) {
	if ptr.Pointer() != nil {
		C.QValueAxis_ConnectMaxChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::maxChanged", f)
	}
}

func (ptr *QValueAxis) DisconnectMaxChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectMaxChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::maxChanged")
	}
}

func (ptr *QValueAxis) MaxChanged(max float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_MaxChanged(ptr.Pointer(), C.double(max))
	}
}

//export callbackQValueAxis_MinChanged
func callbackQValueAxis_MinChanged(ptr unsafe.Pointer, min C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QValueAxis::minChanged"); signal != nil {
		signal.(func(float64))(float64(min))
	}

}

func (ptr *QValueAxis) ConnectMinChanged(f func(min float64)) {
	if ptr.Pointer() != nil {
		C.QValueAxis_ConnectMinChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::minChanged", f)
	}
}

func (ptr *QValueAxis) DisconnectMinChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectMinChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::minChanged")
	}
}

func (ptr *QValueAxis) MinChanged(min float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_MinChanged(ptr.Pointer(), C.double(min))
	}
}

//export callbackQValueAxis_MinorTickCountChanged
func callbackQValueAxis_MinorTickCountChanged(ptr unsafe.Pointer, minorTickCount C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QValueAxis::minorTickCountChanged"); signal != nil {
		signal.(func(int))(int(int32(minorTickCount)))
	}

}

func (ptr *QValueAxis) ConnectMinorTickCountChanged(f func(minorTickCount int)) {
	if ptr.Pointer() != nil {
		C.QValueAxis_ConnectMinorTickCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::minorTickCountChanged", f)
	}
}

func (ptr *QValueAxis) DisconnectMinorTickCountChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectMinorTickCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::minorTickCountChanged")
	}
}

func (ptr *QValueAxis) MinorTickCountChanged(minorTickCount int) {
	if ptr.Pointer() != nil {
		C.QValueAxis_MinorTickCountChanged(ptr.Pointer(), C.int(int32(minorTickCount)))
	}
}

//export callbackQValueAxis_RangeChanged
func callbackQValueAxis_RangeChanged(ptr unsafe.Pointer, min C.double, max C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QValueAxis::rangeChanged"); signal != nil {
		signal.(func(float64, float64))(float64(min), float64(max))
	}

}

func (ptr *QValueAxis) ConnectRangeChanged(f func(min float64, max float64)) {
	if ptr.Pointer() != nil {
		C.QValueAxis_ConnectRangeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::rangeChanged", f)
	}
}

func (ptr *QValueAxis) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::rangeChanged")
	}
}

func (ptr *QValueAxis) RangeChanged(min float64, max float64) {
	if ptr.Pointer() != nil {
		C.QValueAxis_RangeChanged(ptr.Pointer(), C.double(min), C.double(max))
	}
}

func (ptr *QValueAxis) SetLabelFormat(format string) {
	if ptr.Pointer() != nil {
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QValueAxis_SetLabelFormat(ptr.Pointer(), formatC)
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

func (ptr *QValueAxis) SetTickCount(count int) {
	if ptr.Pointer() != nil {
		C.QValueAxis_SetTickCount(ptr.Pointer(), C.int(int32(count)))
	}
}

//export callbackQValueAxis_TickCountChanged
func callbackQValueAxis_TickCountChanged(ptr unsafe.Pointer, tickCount C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QValueAxis::tickCountChanged"); signal != nil {
		signal.(func(int))(int(int32(tickCount)))
	}

}

func (ptr *QValueAxis) ConnectTickCountChanged(f func(tickCount int)) {
	if ptr.Pointer() != nil {
		C.QValueAxis_ConnectTickCountChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::tickCountChanged", f)
	}
}

func (ptr *QValueAxis) DisconnectTickCountChanged() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DisconnectTickCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QValueAxis::tickCountChanged")
	}
}

func (ptr *QValueAxis) TickCountChanged(tickCount int) {
	if ptr.Pointer() != nil {
		C.QValueAxis_TickCountChanged(ptr.Pointer(), C.int(int32(tickCount)))
	}
}

func (ptr *QValueAxis) DestroyQValueAxis() {
	if ptr.Pointer() != nil {
		C.QValueAxis_DestroyQValueAxis(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QValueAxis) LabelFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QValueAxis_LabelFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QValueAxis) MinorTickCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QValueAxis_MinorTickCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QValueAxis) TickCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QValueAxis_TickCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QValueAxis) Max() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QValueAxis_Max(ptr.Pointer()))
	}
	return 0
}

func (ptr *QValueAxis) Min() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QValueAxis_Min(ptr.Pointer()))
	}
	return 0
}

type QXYLegendMarker struct {
	ptr unsafe.Pointer
}

type QXYLegendMarker_ITF interface {
	QXYLegendMarker_PTR() *QXYLegendMarker
}

func (ptr *QXYLegendMarker) QXYLegendMarker_PTR() *QXYLegendMarker {
	return ptr
}

func (ptr *QXYLegendMarker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXYLegendMarker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXYLegendMarker(ptr QXYLegendMarker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXYLegendMarker_PTR().Pointer()
	}
	return nil
}

func NewQXYLegendMarkerFromPointer(ptr unsafe.Pointer) *QXYLegendMarker {
	var n = new(QXYLegendMarker)
	n.SetPointer(ptr)
	return n
}

//export callbackQXYLegendMarker_Series
func callbackQXYLegendMarker_Series(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYLegendMarker::series"); signal != nil {
		return PointerFromQXYSeries(signal.(func() *QXYSeries)())
	}

	return PointerFromQXYSeries(NewQXYLegendMarkerFromPointer(ptr).SeriesDefault())
}

func (ptr *QXYLegendMarker) ConnectSeries(f func() *QXYSeries) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYLegendMarker::series", f)
	}
}

func (ptr *QXYLegendMarker) DisconnectSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYLegendMarker::series")
	}
}

func (ptr *QXYLegendMarker) Series() *QXYSeries {
	if ptr.Pointer() != nil {
		return NewQXYSeriesFromPointer(C.QXYLegendMarker_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXYLegendMarker) SeriesDefault() *QXYSeries {
	if ptr.Pointer() != nil {
		return NewQXYSeriesFromPointer(C.QXYLegendMarker_SeriesDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXYLegendMarker_DestroyQXYLegendMarker
func callbackQXYLegendMarker_DestroyQXYLegendMarker(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYLegendMarker::~QXYLegendMarker"); signal != nil {
		signal.(func())()
	} else {
		NewQXYLegendMarkerFromPointer(ptr).DestroyQXYLegendMarkerDefault()
	}
}

func (ptr *QXYLegendMarker) ConnectDestroyQXYLegendMarker(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYLegendMarker::~QXYLegendMarker", f)
	}
}

func (ptr *QXYLegendMarker) DisconnectDestroyQXYLegendMarker() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYLegendMarker::~QXYLegendMarker")
	}
}

func (ptr *QXYLegendMarker) DestroyQXYLegendMarker() {
	if ptr.Pointer() != nil {
		C.QXYLegendMarker_DestroyQXYLegendMarker(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QXYLegendMarker) DestroyQXYLegendMarkerDefault() {
	if ptr.Pointer() != nil {
		C.QXYLegendMarker_DestroyQXYLegendMarkerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QXYSeries struct {
	ptr unsafe.Pointer
}

type QXYSeries_ITF interface {
	QXYSeries_PTR() *QXYSeries
}

func (ptr *QXYSeries) QXYSeries_PTR() *QXYSeries {
	return ptr
}

func (ptr *QXYSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXYSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXYSeries(ptr QXYSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXYSeries_PTR().Pointer()
	}
	return nil
}

func NewQXYSeriesFromPointer(ptr unsafe.Pointer) *QXYSeries {
	var n = new(QXYSeries)
	n.SetPointer(ptr)
	return n
}
func (ptr *QXYSeries) Append3(points []*core.QPointF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Append3(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQXYSeriesFromPointer(NewQXYSeriesFromPointer(unsafe.Pointer(uintptr(1))).__append_points_newList3())
			for _, v := range points {
				tmpList.__append_points_setList3(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QXYSeries) Append2(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Append2(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QXYSeries) Append(x float64, y float64) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Append(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QXYSeries) Clear() {
	if ptr.Pointer() != nil {
		C.QXYSeries_Clear(ptr.Pointer())
	}
}

//export callbackQXYSeries_Clicked
func callbackQXYSeries_Clicked(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::clicked"); signal != nil {
		signal.(func(*core.QPointF))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QXYSeries) ConnectClicked(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::clicked", f)
	}
}

func (ptr *QXYSeries) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::clicked")
	}
}

func (ptr *QXYSeries) Clicked(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Clicked(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQXYSeries_ColorChanged
func callbackQXYSeries_ColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QXYSeries) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::colorChanged", f)
	}
}

func (ptr *QXYSeries) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::colorChanged")
	}
}

func (ptr *QXYSeries) ColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQXYSeries_DoubleClicked
func callbackQXYSeries_DoubleClicked(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::doubleClicked"); signal != nil {
		signal.(func(*core.QPointF))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QXYSeries) ConnectDoubleClicked(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::doubleClicked", f)
	}
}

func (ptr *QXYSeries) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::doubleClicked")
	}
}

func (ptr *QXYSeries) DoubleClicked(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_DoubleClicked(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQXYSeries_Hovered
func callbackQXYSeries_Hovered(ptr unsafe.Pointer, point unsafe.Pointer, state C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::hovered"); signal != nil {
		signal.(func(*core.QPointF, bool))(core.NewQPointFFromPointer(point), int8(state) != 0)
	}

}

func (ptr *QXYSeries) ConnectHovered(f func(point *core.QPointF, state bool)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::hovered", f)
	}
}

func (ptr *QXYSeries) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::hovered")
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

//export callbackQXYSeries_PenChanged
func callbackQXYSeries_PenChanged(ptr unsafe.Pointer, pen unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::penChanged"); signal != nil {
		signal.(func(*gui.QPen))(gui.NewQPenFromPointer(pen))
	}

}

func (ptr *QXYSeries) ConnectPenChanged(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPenChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::penChanged", f)
	}
}

func (ptr *QXYSeries) DisconnectPenChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPenChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::penChanged")
	}
}

func (ptr *QXYSeries) PenChanged(pen gui.QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PenChanged(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

//export callbackQXYSeries_PointAdded
func callbackQXYSeries_PointAdded(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pointAdded"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QXYSeries) ConnectPointAdded(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPointAdded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointAdded", f)
	}
}

func (ptr *QXYSeries) DisconnectPointAdded() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointAdded")
	}
}

func (ptr *QXYSeries) PointAdded(index int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointAdded(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQXYSeries_PointLabelsClippingChanged
func callbackQXYSeries_PointLabelsClippingChanged(ptr unsafe.Pointer, clipping C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pointLabelsClippingChanged"); signal != nil {
		signal.(func(bool))(int8(clipping) != 0)
	}

}

func (ptr *QXYSeries) ConnectPointLabelsClippingChanged(f func(clipping bool)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPointLabelsClippingChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointLabelsClippingChanged", f)
	}
}

func (ptr *QXYSeries) DisconnectPointLabelsClippingChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointLabelsClippingChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointLabelsClippingChanged")
	}
}

func (ptr *QXYSeries) PointLabelsClippingChanged(clipping bool) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointLabelsClippingChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(clipping))))
	}
}

//export callbackQXYSeries_PointLabelsColorChanged
func callbackQXYSeries_PointLabelsColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pointLabelsColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QXYSeries) ConnectPointLabelsColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPointLabelsColorChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointLabelsColorChanged", f)
	}
}

func (ptr *QXYSeries) DisconnectPointLabelsColorChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointLabelsColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointLabelsColorChanged")
	}
}

func (ptr *QXYSeries) PointLabelsColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointLabelsColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQXYSeries_PointLabelsFontChanged
func callbackQXYSeries_PointLabelsFontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pointLabelsFontChanged"); signal != nil {
		signal.(func(*gui.QFont))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *QXYSeries) ConnectPointLabelsFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPointLabelsFontChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointLabelsFontChanged", f)
	}
}

func (ptr *QXYSeries) DisconnectPointLabelsFontChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointLabelsFontChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointLabelsFontChanged")
	}
}

func (ptr *QXYSeries) PointLabelsFontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointLabelsFontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

//export callbackQXYSeries_PointLabelsFormatChanged
func callbackQXYSeries_PointLabelsFormatChanged(ptr unsafe.Pointer, format C.struct_QtCharts_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pointLabelsFormatChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(format))
	}

}

func (ptr *QXYSeries) ConnectPointLabelsFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPointLabelsFormatChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointLabelsFormatChanged", f)
	}
}

func (ptr *QXYSeries) DisconnectPointLabelsFormatChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointLabelsFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointLabelsFormatChanged")
	}
}

func (ptr *QXYSeries) PointLabelsFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QXYSeries_PointLabelsFormatChanged(ptr.Pointer(), formatC)
	}
}

//export callbackQXYSeries_PointLabelsVisibilityChanged
func callbackQXYSeries_PointLabelsVisibilityChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pointLabelsVisibilityChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QXYSeries) ConnectPointLabelsVisibilityChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPointLabelsVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointLabelsVisibilityChanged", f)
	}
}

func (ptr *QXYSeries) DisconnectPointLabelsVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointLabelsVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointLabelsVisibilityChanged")
	}
}

func (ptr *QXYSeries) PointLabelsVisibilityChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointLabelsVisibilityChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQXYSeries_PointRemoved
func callbackQXYSeries_PointRemoved(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pointRemoved"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QXYSeries) ConnectPointRemoved(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPointRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointRemoved", f)
	}
}

func (ptr *QXYSeries) DisconnectPointRemoved() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointRemoved")
	}
}

func (ptr *QXYSeries) PointRemoved(index int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointRemoved(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQXYSeries_PointReplaced
func callbackQXYSeries_PointReplaced(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pointReplaced"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QXYSeries) ConnectPointReplaced(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPointReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointReplaced", f)
	}
}

func (ptr *QXYSeries) DisconnectPointReplaced() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointReplaced")
	}
}

func (ptr *QXYSeries) PointReplaced(index int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointReplaced(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQXYSeries_PointsRemoved
func callbackQXYSeries_PointsRemoved(ptr unsafe.Pointer, index C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pointsRemoved"); signal != nil {
		signal.(func(int, int))(int(int32(index)), int(int32(count)))
	}

}

func (ptr *QXYSeries) ConnectPointsRemoved(f func(index int, count int)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPointsRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointsRemoved", f)
	}
}

func (ptr *QXYSeries) DisconnectPointsRemoved() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointsRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointsRemoved")
	}
}

func (ptr *QXYSeries) PointsRemoved(index int, count int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointsRemoved(ptr.Pointer(), C.int(int32(index)), C.int(int32(count)))
	}
}

//export callbackQXYSeries_PointsReplaced
func callbackQXYSeries_PointsReplaced(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pointsReplaced"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QXYSeries) ConnectPointsReplaced(f func()) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPointsReplaced(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointsReplaced", f)
	}
}

func (ptr *QXYSeries) DisconnectPointsReplaced() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPointsReplaced(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pointsReplaced")
	}
}

func (ptr *QXYSeries) PointsReplaced() {
	if ptr.Pointer() != nil {
		C.QXYSeries_PointsReplaced(ptr.Pointer())
	}
}

//export callbackQXYSeries_Pressed
func callbackQXYSeries_Pressed(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::pressed"); signal != nil {
		signal.(func(*core.QPointF))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QXYSeries) ConnectPressed(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pressed", f)
	}
}

func (ptr *QXYSeries) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::pressed")
	}
}

func (ptr *QXYSeries) Pressed(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Pressed(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

//export callbackQXYSeries_Released
func callbackQXYSeries_Released(ptr unsafe.Pointer, point unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::released"); signal != nil {
		signal.(func(*core.QPointF))(core.NewQPointFFromPointer(point))
	}

}

func (ptr *QXYSeries) ConnectReleased(f func(point *core.QPointF)) {
	if ptr.Pointer() != nil {
		C.QXYSeries_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::released", f)
	}
}

func (ptr *QXYSeries) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::released")
	}
}

func (ptr *QXYSeries) Released(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Released(ptr.Pointer(), core.PointerFromQPointF(point))
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

func (ptr *QXYSeries) Remove(x float64, y float64) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Remove(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QXYSeries) RemovePoints(index int, count int) {
	if ptr.Pointer() != nil {
		C.QXYSeries_RemovePoints(ptr.Pointer(), C.int(int32(index)), C.int(int32(count)))
	}
}

func (ptr *QXYSeries) Replace5(points []*core.QPointF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace5(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQXYSeriesFromPointer(NewQXYSeriesFromPointer(unsafe.Pointer(uintptr(1))).__replace_points_newList5())
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
			var tmpList = NewQXYSeriesFromPointer(NewQXYSeriesFromPointer(unsafe.Pointer(uintptr(1))).__replace_points_newList6())
			for _, v := range points {
				tmpList.__replace_points_setList6(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QXYSeries) Replace2(oldPoint core.QPointF_ITF, newPoint core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace2(ptr.Pointer(), core.PointerFromQPointF(oldPoint), core.PointerFromQPointF(newPoint))
	}
}

func (ptr *QXYSeries) Replace4(index int, newPoint core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace4(ptr.Pointer(), C.int(int32(index)), core.PointerFromQPointF(newPoint))
	}
}

func (ptr *QXYSeries) Replace3(index int, newX float64, newY float64) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace3(ptr.Pointer(), C.int(int32(index)), C.double(newX), C.double(newY))
	}
}

func (ptr *QXYSeries) Replace(oldX float64, oldY float64, newX float64, newY float64) {
	if ptr.Pointer() != nil {
		C.QXYSeries_Replace(ptr.Pointer(), C.double(oldX), C.double(oldY), C.double(newX), C.double(newY))
	}
}

//export callbackQXYSeries_SetBrush
func callbackQXYSeries_SetBrush(ptr unsafe.Pointer, brush unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::setBrush"); signal != nil {
		signal.(func(*gui.QBrush))(gui.NewQBrushFromPointer(brush))
	} else {
		NewQXYSeriesFromPointer(ptr).SetBrushDefault(gui.NewQBrushFromPointer(brush))
	}
}

func (ptr *QXYSeries) ConnectSetBrush(f func(brush *gui.QBrush)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::setBrush", f)
	}
}

func (ptr *QXYSeries) DisconnectSetBrush() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::setBrush")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::setColor"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	} else {
		NewQXYSeriesFromPointer(ptr).SetColorDefault(gui.NewQColorFromPointer(color))
	}
}

func (ptr *QXYSeries) ConnectSetColor(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::setColor", f)
	}
}

func (ptr *QXYSeries) DisconnectSetColor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::setColor")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::setPen"); signal != nil {
		signal.(func(*gui.QPen))(gui.NewQPenFromPointer(pen))
	} else {
		NewQXYSeriesFromPointer(ptr).SetPenDefault(gui.NewQPenFromPointer(pen))
	}
}

func (ptr *QXYSeries) ConnectSetPen(f func(pen *gui.QPen)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::setPen", f)
	}
}

func (ptr *QXYSeries) DisconnectSetPen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::setPen")
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
		var formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
		C.QXYSeries_SetPointLabelsFormat(ptr.Pointer(), formatC)
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

func (ptr *QXYSeries) DestroyQXYSeries() {
	if ptr.Pointer() != nil {
		C.QXYSeries_DestroyQXYSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QXYSeries) Brush() *gui.QBrush {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQBrushFromPointer(C.QXYSeries_Brush(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

//export callbackQXYSeries_Color
func callbackQXYSeries_Color(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QXYSeries::color"); signal != nil {
		return gui.PointerFromQColor(signal.(func() *gui.QColor)())
	}

	return gui.PointerFromQColor(NewQXYSeriesFromPointer(ptr).ColorDefault())
}

func (ptr *QXYSeries) ConnectColor(f func() *gui.QColor) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::color", f)
	}
}

func (ptr *QXYSeries) DisconnectColor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QXYSeries::color")
	}
}

func (ptr *QXYSeries) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QXYSeries_Color(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) ColorDefault() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QXYSeries_ColorDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) PointLabelsColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QXYSeries_PointLabelsColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) PointLabelsFont() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.QXYSeries_PointLabelsFont(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) Points() []*core.QPointF {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*core.QPointF {
			var out = make([]*core.QPointF, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQXYSeriesFromPointer(l.data).__points_atList(i)
			}
			return out
		}(C.QXYSeries_Points(ptr.Pointer()))
	}
	return make([]*core.QPointF, 0)
}

func (ptr *QXYSeries) Pen() *gui.QPen {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPenFromPointer(C.QXYSeries_Pen(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPen).DestroyQPen)
		return tmpValue
	}
	return nil
}

func (ptr *QXYSeries) PointLabelsFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXYSeries_PointLabelsFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXYSeries) PointsVector() []*core.QPointF {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCharts_PackedList) []*core.QPointF {
			var out = make([]*core.QPointF, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQXYSeriesFromPointer(l.data).__pointsVector_atList(i)
			}
			return out
		}(C.QXYSeries_PointsVector(ptr.Pointer()))
	}
	return make([]*core.QPointF, 0)
}

func (ptr *QXYSeries) PointLabelsClipping() bool {
	if ptr.Pointer() != nil {
		return C.QXYSeries_PointLabelsClipping(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXYSeries) PointLabelsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QXYSeries_PointLabelsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXYSeries) PointsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QXYSeries_PointsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXYSeries) At(index int) *core.QPointF {
	if ptr.Pointer() != nil {
		return core.NewQPointFFromPointer(C.QXYSeries_At(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QXYSeries) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXYSeries_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXYSeries) __append_points_atList3(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QXYSeries___append_points_atList3(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXYSeries___append_points_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXYSeries) __replace_points_atList5(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QXYSeries___replace_points_atList5(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXYSeries___replace_points_newList5(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXYSeries) __replace_points_atList6(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QXYSeries___replace_points_atList6(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXYSeries___replace_points_newList6(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXYSeries) __points_atList(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QXYSeries___points_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXYSeries___points_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXYSeries) __pointsVector_atList(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFFromPointer(C.QXYSeries___pointsVector_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXYSeries___pointsVector_newList(ptr.Pointer()))
	}
	return nil
}
