//go:build !minimal
// +build !minimal

package charts

import (
	"unsafe"

	"github.com/akiyosi/qt/core"
	"github.com/akiyosi/qt/gui"
	"github.com/akiyosi/qt/internal"
	"github.com/akiyosi/qt/widgets"
)

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

func (n *QAbstractAxis) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAbstractAxis) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQAbstractAxisFromPointer(ptr unsafe.Pointer) (n *QAbstractAxis) {
	n = new(QAbstractAxis)
	n.InitFromInternal(uintptr(ptr), "charts.QAbstractAxis")
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

	return core.Qt__AlignmentFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Alignment"}).(float64))
}

func (ptr *QAbstractAxis) ConnectColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColorChanged"})
}

func (ptr *QAbstractAxis) ColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorChanged", color})
}

func (ptr *QAbstractAxis) GridLineColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GridLineColor"}).(*gui.QColor)
}

func (ptr *QAbstractAxis) ConnectGridLineColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGridLineColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectGridLineColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGridLineColorChanged"})
}

func (ptr *QAbstractAxis) GridLineColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GridLineColorChanged", color})
}

func (ptr *QAbstractAxis) GridLinePen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GridLinePen"}).(*gui.QPen)
}

func (ptr *QAbstractAxis) ConnectGridLinePenChanged(f func(pen *gui.QPen)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGridLinePenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectGridLinePenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGridLinePenChanged"})
}

func (ptr *QAbstractAxis) GridLinePenChanged(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GridLinePenChanged", pen})
}

func (ptr *QAbstractAxis) ConnectGridVisibleChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGridVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectGridVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGridVisibleChanged"})
}

func (ptr *QAbstractAxis) GridVisibleChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GridVisibleChanged", visible})
}

func (ptr *QAbstractAxis) Hide() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hide"})
}

func (ptr *QAbstractAxis) IsGridLineVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsGridLineVisible"}).(bool)
}

func (ptr *QAbstractAxis) IsLineVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLineVisible"}).(bool)
}

func (ptr *QAbstractAxis) IsMinorGridLineVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsMinorGridLineVisible"}).(bool)
}

func (ptr *QAbstractAxis) IsReverse() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsReverse"}).(bool)
}

func (ptr *QAbstractAxis) IsTitleVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsTitleVisible"}).(bool)
}

func (ptr *QAbstractAxis) IsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsVisible"}).(bool)
}

func (ptr *QAbstractAxis) LabelsAngle() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsAngle"}).(float64))
}

func (ptr *QAbstractAxis) ConnectLabelsAngleChanged(f func(angle int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsAngleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectLabelsAngleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsAngleChanged"})
}

func (ptr *QAbstractAxis) LabelsAngleChanged(angle int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsAngleChanged", angle})
}

func (ptr *QAbstractAxis) LabelsBrush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsBrush"}).(*gui.QBrush)
}

func (ptr *QAbstractAxis) ConnectLabelsBrushChanged(f func(brush *gui.QBrush)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectLabelsBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsBrushChanged"})
}

func (ptr *QAbstractAxis) LabelsBrushChanged(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsBrushChanged", brush})
}

func (ptr *QAbstractAxis) LabelsColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsColor"}).(*gui.QColor)
}

func (ptr *QAbstractAxis) ConnectLabelsColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectLabelsColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsColorChanged"})
}

func (ptr *QAbstractAxis) LabelsColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsColorChanged", color})
}

func (ptr *QAbstractAxis) LabelsEditable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsEditable"}).(bool)
}

func (ptr *QAbstractAxis) ConnectLabelsEditableChanged(f func(editable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsEditableChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectLabelsEditableChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsEditableChanged"})
}

func (ptr *QAbstractAxis) LabelsEditableChanged(editable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsEditableChanged", editable})
}

func (ptr *QAbstractAxis) LabelsFont() *gui.QFont {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsFont"}).(*gui.QFont)
}

func (ptr *QAbstractAxis) ConnectLabelsFontChanged(f func(font *gui.QFont)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsFontChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectLabelsFontChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsFontChanged"})
}

func (ptr *QAbstractAxis) LabelsFontChanged(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsFontChanged", font})
}

func (ptr *QAbstractAxis) LabelsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsVisible"}).(bool)
}

func (ptr *QAbstractAxis) ConnectLabelsVisibleChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectLabelsVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsVisibleChanged"})
}

func (ptr *QAbstractAxis) LabelsVisibleChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsVisibleChanged", visible})
}

func (ptr *QAbstractAxis) LinePen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LinePen"}).(*gui.QPen)
}

func (ptr *QAbstractAxis) ConnectLinePenChanged(f func(pen *gui.QPen)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLinePenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectLinePenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLinePenChanged"})
}

func (ptr *QAbstractAxis) LinePenChanged(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LinePenChanged", pen})
}

func (ptr *QAbstractAxis) LinePenColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LinePenColor"}).(*gui.QColor)
}

func (ptr *QAbstractAxis) ConnectLineVisibleChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLineVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectLineVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLineVisibleChanged"})
}

func (ptr *QAbstractAxis) LineVisibleChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LineVisibleChanged", visible})
}

func (ptr *QAbstractAxis) MinorGridLineColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinorGridLineColor"}).(*gui.QColor)
}

func (ptr *QAbstractAxis) ConnectMinorGridLineColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinorGridLineColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectMinorGridLineColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinorGridLineColorChanged"})
}

func (ptr *QAbstractAxis) MinorGridLineColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinorGridLineColorChanged", color})
}

func (ptr *QAbstractAxis) MinorGridLinePen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinorGridLinePen"}).(*gui.QPen)
}

func (ptr *QAbstractAxis) ConnectMinorGridLinePenChanged(f func(pen *gui.QPen)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinorGridLinePenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectMinorGridLinePenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinorGridLinePenChanged"})
}

func (ptr *QAbstractAxis) MinorGridLinePenChanged(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinorGridLinePenChanged", pen})
}

func (ptr *QAbstractAxis) ConnectMinorGridVisibleChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinorGridVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectMinorGridVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinorGridVisibleChanged"})
}

func (ptr *QAbstractAxis) MinorGridVisibleChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinorGridVisibleChanged", visible})
}

func (ptr *QAbstractAxis) Orientation() core.Qt__Orientation {

	return core.Qt__Orientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Orientation"}).(float64))
}

func (ptr *QAbstractAxis) ConnectReverseChanged(f func(reverse bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReverseChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectReverseChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReverseChanged"})
}

func (ptr *QAbstractAxis) ReverseChanged(reverse bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReverseChanged", reverse})
}

func (ptr *QAbstractAxis) SetGridLineColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGridLineColor", color})
}

func (ptr *QAbstractAxis) SetGridLinePen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGridLinePen", pen})
}

func (ptr *QAbstractAxis) SetGridLineVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGridLineVisible", visible})
}

func (ptr *QAbstractAxis) SetLabelsAngle(angle int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsAngle", angle})
}

func (ptr *QAbstractAxis) SetLabelsBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsBrush", brush})
}

func (ptr *QAbstractAxis) SetLabelsColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsColor", color})
}

func (ptr *QAbstractAxis) SetLabelsEditable(editable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsEditable", editable})
}

func (ptr *QAbstractAxis) SetLabelsFont(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsFont", font})
}

func (ptr *QAbstractAxis) SetLabelsVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsVisible", visible})
}

func (ptr *QAbstractAxis) SetLinePen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLinePen", pen})
}

func (ptr *QAbstractAxis) SetLinePenColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLinePenColor", color})
}

func (ptr *QAbstractAxis) SetLineVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLineVisible", visible})
}

func (ptr *QAbstractAxis) SetMax(max core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMax", max})
}

func (ptr *QAbstractAxis) SetMin(min core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMin", min})
}

func (ptr *QAbstractAxis) SetMinorGridLineColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinorGridLineColor", color})
}

func (ptr *QAbstractAxis) SetMinorGridLinePen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinorGridLinePen", pen})
}

func (ptr *QAbstractAxis) SetMinorGridLineVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinorGridLineVisible", visible})
}

func (ptr *QAbstractAxis) SetRange(min core.QVariant_ITF, max core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRange", min, max})
}

func (ptr *QAbstractAxis) SetReverse(reverse bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReverse", reverse})
}

func (ptr *QAbstractAxis) SetShadesBorderColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShadesBorderColor", color})
}

func (ptr *QAbstractAxis) SetShadesBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShadesBrush", brush})
}

func (ptr *QAbstractAxis) SetShadesColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShadesColor", color})
}

func (ptr *QAbstractAxis) SetShadesPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShadesPen", pen})
}

func (ptr *QAbstractAxis) SetShadesVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShadesVisible", visible})
}

func (ptr *QAbstractAxis) SetTitleBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitleBrush", brush})
}

func (ptr *QAbstractAxis) SetTitleFont(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitleFont", font})
}

func (ptr *QAbstractAxis) SetTitleText(title string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitleText", title})
}

func (ptr *QAbstractAxis) SetTitleVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitleVisible", visible})
}

func (ptr *QAbstractAxis) SetVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisible", visible})
}

func (ptr *QAbstractAxis) ShadesBorderColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadesBorderColor"}).(*gui.QColor)
}

func (ptr *QAbstractAxis) ConnectShadesBorderColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShadesBorderColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectShadesBorderColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShadesBorderColorChanged"})
}

func (ptr *QAbstractAxis) ShadesBorderColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadesBorderColorChanged", color})
}

func (ptr *QAbstractAxis) ShadesBrush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadesBrush"}).(*gui.QBrush)
}

func (ptr *QAbstractAxis) ConnectShadesBrushChanged(f func(brush *gui.QBrush)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShadesBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectShadesBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShadesBrushChanged"})
}

func (ptr *QAbstractAxis) ShadesBrushChanged(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadesBrushChanged", brush})
}

func (ptr *QAbstractAxis) ShadesColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadesColor"}).(*gui.QColor)
}

func (ptr *QAbstractAxis) ConnectShadesColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShadesColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectShadesColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShadesColorChanged"})
}

func (ptr *QAbstractAxis) ShadesColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadesColorChanged", color})
}

func (ptr *QAbstractAxis) ShadesPen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadesPen"}).(*gui.QPen)
}

func (ptr *QAbstractAxis) ConnectShadesPenChanged(f func(pen *gui.QPen)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShadesPenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectShadesPenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShadesPenChanged"})
}

func (ptr *QAbstractAxis) ShadesPenChanged(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadesPenChanged", pen})
}

func (ptr *QAbstractAxis) ShadesVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadesVisible"}).(bool)
}

func (ptr *QAbstractAxis) ConnectShadesVisibleChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShadesVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectShadesVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShadesVisibleChanged"})
}

func (ptr *QAbstractAxis) ShadesVisibleChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadesVisibleChanged", visible})
}

func (ptr *QAbstractAxis) Show() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Show"})
}

func (ptr *QAbstractAxis) TitleBrush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleBrush"}).(*gui.QBrush)
}

func (ptr *QAbstractAxis) ConnectTitleBrushChanged(f func(brush *gui.QBrush)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTitleBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectTitleBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTitleBrushChanged"})
}

func (ptr *QAbstractAxis) TitleBrushChanged(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleBrushChanged", brush})
}

func (ptr *QAbstractAxis) TitleFont() *gui.QFont {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleFont"}).(*gui.QFont)
}

func (ptr *QAbstractAxis) ConnectTitleFontChanged(f func(font *gui.QFont)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTitleFontChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectTitleFontChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTitleFontChanged"})
}

func (ptr *QAbstractAxis) TitleFontChanged(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleFontChanged", font})
}

func (ptr *QAbstractAxis) TitleText() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleText"}).(string)
}

func (ptr *QAbstractAxis) ConnectTitleTextChanged(f func(text string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTitleTextChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectTitleTextChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTitleTextChanged"})
}

func (ptr *QAbstractAxis) TitleTextChanged(text string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleTextChanged", text})
}

func (ptr *QAbstractAxis) ConnectTitleVisibleChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTitleVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectTitleVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTitleVisibleChanged"})
}

func (ptr *QAbstractAxis) TitleVisibleChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleVisibleChanged", visible})
}

func (ptr *QAbstractAxis) ConnectType(f func() QAbstractAxis__AxisType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QAbstractAxis) Type() QAbstractAxis__AxisType {

	return QAbstractAxis__AxisType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QAbstractAxis) ConnectVisibleChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVisibleChanged"})
}

func (ptr *QAbstractAxis) VisibleChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VisibleChanged", visible})
}

func (ptr *QAbstractAxis) ConnectDestroyQAbstractAxis(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractAxis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractAxis) DisconnectDestroyQAbstractAxis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractAxis"})
}

func (ptr *QAbstractAxis) DestroyQAbstractAxis() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractAxis"})
}

func (ptr *QAbstractAxis) DestroyQAbstractAxisDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractAxisDefault"})
}

func (ptr *QAbstractAxis) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QAbstractAxis) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QAbstractAxis) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractAxis) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QAbstractAxis) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QAbstractAxis) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractAxis) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QAbstractAxis) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QAbstractAxis) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractAxis) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QAbstractAxis) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QAbstractAxis) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QAbstractAxis) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QAbstractAxis) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QAbstractAxis) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QAbstractAxis) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QAbstractAxis) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QAbstractAxis) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QAbstractAxis) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QAbstractAxis) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QAbstractAxis) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QAbstractBarSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAbstractBarSeries) ClassNameInternalF() string {
	return n.QAbstractSeries_PTR().ClassNameInternalF()
}

func NewQAbstractBarSeriesFromPointer(ptr unsafe.Pointer) (n *QAbstractBarSeries) {
	n = new(QAbstractBarSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QAbstractBarSeries")
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

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", set}).(bool)
}

func (ptr *QAbstractBarSeries) Append2(sets []*QBarSet) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append2", sets}).(bool)
}

func (ptr *QAbstractBarSeries) BarSets() []*QBarSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BarSets"}).([]*QBarSet)
}

func (ptr *QAbstractBarSeries) BarWidth() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BarWidth"}).(float64)
}

func (ptr *QAbstractBarSeries) ConnectBarsetsAdded(f func(sets []*QBarSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBarsetsAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectBarsetsAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBarsetsAdded"})
}

func (ptr *QAbstractBarSeries) BarsetsAdded(sets []*QBarSet) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BarsetsAdded", sets})
}

func (ptr *QAbstractBarSeries) ConnectBarsetsRemoved(f func(sets []*QBarSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBarsetsRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectBarsetsRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBarsetsRemoved"})
}

func (ptr *QAbstractBarSeries) BarsetsRemoved(sets []*QBarSet) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BarsetsRemoved", sets})
}

func (ptr *QAbstractBarSeries) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QAbstractBarSeries) ConnectClicked(f func(index int, barset *QBarSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QAbstractBarSeries) Clicked(index int, barset QBarSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked", index, barset})
}

func (ptr *QAbstractBarSeries) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QAbstractBarSeries) ConnectCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCountChanged"})
}

func (ptr *QAbstractBarSeries) CountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CountChanged"})
}

func (ptr *QAbstractBarSeries) ConnectDoubleClicked(f func(index int, barset *QBarSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDoubleClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectDoubleClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDoubleClicked"})
}

func (ptr *QAbstractBarSeries) DoubleClicked(index int, barset QBarSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoubleClicked", index, barset})
}

func (ptr *QAbstractBarSeries) ConnectHovered(f func(status bool, index int, barset *QBarSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QAbstractBarSeries) Hovered(status bool, index int, barset QBarSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", status, index, barset})
}

func (ptr *QAbstractBarSeries) Insert(index int, set QBarSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", index, set}).(bool)
}

func (ptr *QAbstractBarSeries) IsLabelsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLabelsVisible"}).(bool)
}

func (ptr *QAbstractBarSeries) LabelsAngle() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsAngle"}).(float64)
}

func (ptr *QAbstractBarSeries) ConnectLabelsAngleChanged(f func(angle float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsAngleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectLabelsAngleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsAngleChanged"})
}

func (ptr *QAbstractBarSeries) LabelsAngleChanged(angle float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsAngleChanged", angle})
}

func (ptr *QAbstractBarSeries) LabelsFormat() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsFormat"}).(string)
}

func (ptr *QAbstractBarSeries) ConnectLabelsFormatChanged(f func(format string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsFormatChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectLabelsFormatChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsFormatChanged"})
}

func (ptr *QAbstractBarSeries) LabelsFormatChanged(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsFormatChanged", format})
}

func (ptr *QAbstractBarSeries) LabelsPosition() QAbstractBarSeries__LabelsPosition {

	return QAbstractBarSeries__LabelsPosition(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsPosition"}).(float64))
}

func (ptr *QAbstractBarSeries) ConnectLabelsPositionChanged(f func(position QAbstractBarSeries__LabelsPosition)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsPositionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectLabelsPositionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsPositionChanged"})
}

func (ptr *QAbstractBarSeries) LabelsPositionChanged(position QAbstractBarSeries__LabelsPosition) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsPositionChanged", position})
}

func (ptr *QAbstractBarSeries) LabelsPrecision() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsPrecision"}).(float64))
}

func (ptr *QAbstractBarSeries) ConnectLabelsPrecisionChanged(f func(precision int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsPrecisionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectLabelsPrecisionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsPrecisionChanged"})
}

func (ptr *QAbstractBarSeries) LabelsPrecisionChanged(precision int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsPrecisionChanged", precision})
}

func (ptr *QAbstractBarSeries) ConnectLabelsVisibleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectLabelsVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsVisibleChanged"})
}

func (ptr *QAbstractBarSeries) LabelsVisibleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsVisibleChanged"})
}

func (ptr *QAbstractBarSeries) ConnectPressed(f func(index int, barset *QBarSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPressed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectPressed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPressed"})
}

func (ptr *QAbstractBarSeries) Pressed(index int, barset QBarSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressed", index, barset})
}

func (ptr *QAbstractBarSeries) ConnectReleased(f func(index int, barset *QBarSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleased", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectReleased() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleased"})
}

func (ptr *QAbstractBarSeries) Released(index int, barset QBarSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Released", index, barset})
}

func (ptr *QAbstractBarSeries) Remove(set QBarSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", set}).(bool)
}

func (ptr *QAbstractBarSeries) SetBarWidth(width float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBarWidth", width})
}

func (ptr *QAbstractBarSeries) SetLabelsAngle(angle float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsAngle", angle})
}

func (ptr *QAbstractBarSeries) SetLabelsFormat(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsFormat", format})
}

func (ptr *QAbstractBarSeries) SetLabelsPosition(position QAbstractBarSeries__LabelsPosition) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsPosition", position})
}

func (ptr *QAbstractBarSeries) SetLabelsPrecision(precision int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsPrecision", precision})
}

func (ptr *QAbstractBarSeries) SetLabelsVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsVisible", visible})
}

func (ptr *QAbstractBarSeries) Take(set QBarSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Take", set}).(bool)
}

func (ptr *QAbstractBarSeries) ConnectDestroyQAbstractBarSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractBarSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractBarSeries) DisconnectDestroyQAbstractBarSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractBarSeries"})
}

func (ptr *QAbstractBarSeries) DestroyQAbstractBarSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractBarSeries"})
}

func (ptr *QAbstractBarSeries) DestroyQAbstractBarSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractBarSeriesDefault"})
}

func (ptr *QAbstractBarSeries) __append_sets_atList2(i int) *QBarSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_sets_atList2", i}).(*QBarSet)
}

func (ptr *QAbstractBarSeries) __append_sets_setList2(i QBarSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_sets_setList2", i})
}

func (ptr *QAbstractBarSeries) __append_sets_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_sets_newList2"}).(unsafe.Pointer)
}

func (ptr *QAbstractBarSeries) __barSets_atList(i int) *QBarSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__barSets_atList", i}).(*QBarSet)
}

func (ptr *QAbstractBarSeries) __barSets_setList(i QBarSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__barSets_setList", i})
}

func (ptr *QAbstractBarSeries) __barSets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__barSets_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractBarSeries) __barsetsAdded_sets_atList(i int) *QBarSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__barsetsAdded_sets_atList", i}).(*QBarSet)
}

func (ptr *QAbstractBarSeries) __barsetsAdded_sets_setList(i QBarSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__barsetsAdded_sets_setList", i})
}

func (ptr *QAbstractBarSeries) __barsetsAdded_sets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__barsetsAdded_sets_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractBarSeries) __barsetsRemoved_sets_atList(i int) *QBarSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__barsetsRemoved_sets_atList", i}).(*QBarSet)
}

func (ptr *QAbstractBarSeries) __barsetsRemoved_sets_setList(i QBarSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__barsetsRemoved_sets_setList", i})
}

func (ptr *QAbstractBarSeries) __barsetsRemoved_sets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__barsetsRemoved_sets_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractBarSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QAbstractBarSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
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

func (n *QAbstractSeries) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAbstractSeries) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQAbstractSeriesFromPointer(ptr unsafe.Pointer) (n *QAbstractSeries) {
	n = new(QAbstractSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QAbstractSeries")
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

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AttachAxis", axis}).(bool)
}

func (ptr *QAbstractSeries) AttachedAxes() []*QAbstractAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AttachedAxes"}).([]*QAbstractAxis)
}

func (ptr *QAbstractSeries) Chart() *QChart {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Chart"}).(*QChart)
}

func (ptr *QAbstractSeries) DetachAxis(axis QAbstractAxis_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DetachAxis", axis}).(bool)
}

func (ptr *QAbstractSeries) Hide() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hide"})
}

func (ptr *QAbstractSeries) IsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsVisible"}).(bool)
}

func (ptr *QAbstractSeries) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QAbstractSeries) ConnectNameChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNameChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSeries) DisconnectNameChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNameChanged"})
}

func (ptr *QAbstractSeries) NameChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NameChanged"})
}

func (ptr *QAbstractSeries) Opacity() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Opacity"}).(float64)
}

func (ptr *QAbstractSeries) ConnectOpacityChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpacityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSeries) DisconnectOpacityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpacityChanged"})
}

func (ptr *QAbstractSeries) OpacityChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpacityChanged"})
}

func (ptr *QAbstractSeries) SetName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", name})
}

func (ptr *QAbstractSeries) SetOpacity(opacity float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpacity", opacity})
}

func (ptr *QAbstractSeries) SetUseOpenGL(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUseOpenGL", enable})
}

func (ptr *QAbstractSeries) SetVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisible", visible})
}

func (ptr *QAbstractSeries) Show() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Show"})
}

func (ptr *QAbstractSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QAbstractSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QAbstractSeries) UseOpenGL() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UseOpenGL"}).(bool)
}

func (ptr *QAbstractSeries) ConnectUseOpenGLChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUseOpenGLChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSeries) DisconnectUseOpenGLChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUseOpenGLChanged"})
}

func (ptr *QAbstractSeries) UseOpenGLChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UseOpenGLChanged"})
}

func (ptr *QAbstractSeries) ConnectVisibleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSeries) DisconnectVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVisibleChanged"})
}

func (ptr *QAbstractSeries) VisibleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VisibleChanged"})
}

func (ptr *QAbstractSeries) ConnectDestroyQAbstractSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSeries) DisconnectDestroyQAbstractSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractSeries"})
}

func (ptr *QAbstractSeries) DestroyQAbstractSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractSeries"})
}

func (ptr *QAbstractSeries) DestroyQAbstractSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractSeriesDefault"})
}

func (ptr *QAbstractSeries) __attachedAxes_atList(i int) *QAbstractAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__attachedAxes_atList", i}).(*QAbstractAxis)
}

func (ptr *QAbstractSeries) __attachedAxes_setList(i QAbstractAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__attachedAxes_setList", i})
}

func (ptr *QAbstractSeries) __attachedAxes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__attachedAxes_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractSeries) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QAbstractSeries) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QAbstractSeries) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractSeries) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QAbstractSeries) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QAbstractSeries) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractSeries) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QAbstractSeries) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QAbstractSeries) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractSeries) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QAbstractSeries) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QAbstractSeries) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QAbstractSeries) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QAbstractSeries) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QAbstractSeries) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QAbstractSeries) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QAbstractSeries) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QAbstractSeries) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QAbstractSeries) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QAbstractSeries) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QAbstractSeries) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QAreaLegendMarker) InitFromInternal(ptr uintptr, name string) {
	n.QLegendMarker_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAreaLegendMarker) ClassNameInternalF() string {
	return n.QLegendMarker_PTR().ClassNameInternalF()
}

func NewQAreaLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QAreaLegendMarker) {
	n = new(QAreaLegendMarker)
	n.InitFromInternal(uintptr(ptr), "charts.QAreaLegendMarker")
	return
}
func (ptr *QAreaLegendMarker) ConnectSeries(f func() *QAreaSeries) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaLegendMarker) DisconnectSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeries"})
}

func (ptr *QAreaLegendMarker) Series() *QAreaSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QAreaSeries)
}

func (ptr *QAreaLegendMarker) SeriesDefault() *QAreaSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesDefault"}).(*QAreaSeries)
}

func (ptr *QAreaLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaLegendMarker) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QAreaLegendMarker) Type() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QAreaLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QAreaLegendMarker) ConnectDestroyQAreaLegendMarker(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAreaLegendMarker", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaLegendMarker) DisconnectDestroyQAreaLegendMarker() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAreaLegendMarker"})
}

func (ptr *QAreaLegendMarker) DestroyQAreaLegendMarker() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAreaLegendMarker"})
}

func (ptr *QAreaLegendMarker) DestroyQAreaLegendMarkerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAreaLegendMarkerDefault"})
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

func (n *QAreaSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAreaSeries) ClassNameInternalF() string {
	return n.QAbstractSeries_PTR().ClassNameInternalF()
}

func NewQAreaSeriesFromPointer(ptr unsafe.Pointer) (n *QAreaSeries) {
	n = new(QAreaSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QAreaSeries")
	return
}
func NewQAreaSeries(parent core.QObject_ITF) *QAreaSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQAreaSeries", "", parent}).(*QAreaSeries)
}

func NewQAreaSeries2(upperSeries QLineSeries_ITF, lowerSeries QLineSeries_ITF) *QAreaSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQAreaSeries2", "", upperSeries, lowerSeries}).(*QAreaSeries)
}

func (ptr *QAreaSeries) BorderColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderColor"}).(*gui.QColor)
}

func (ptr *QAreaSeries) ConnectBorderColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBorderColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectBorderColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBorderColorChanged"})
}

func (ptr *QAreaSeries) BorderColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderColorChanged", color})
}

func (ptr *QAreaSeries) Brush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Brush"}).(*gui.QBrush)
}

func (ptr *QAreaSeries) ConnectClicked(f func(point *core.QPointF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QAreaSeries) Clicked(point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked", point})
}

func (ptr *QAreaSeries) Color() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Color"}).(*gui.QColor)
}

func (ptr *QAreaSeries) ConnectColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColorChanged"})
}

func (ptr *QAreaSeries) ColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorChanged", color})
}

func (ptr *QAreaSeries) ConnectDoubleClicked(f func(point *core.QPointF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDoubleClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectDoubleClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDoubleClicked"})
}

func (ptr *QAreaSeries) DoubleClicked(point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoubleClicked", point})
}

func (ptr *QAreaSeries) ConnectHovered(f func(point *core.QPointF, state bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QAreaSeries) Hovered(point core.QPointF_ITF, state bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", point, state})
}

func (ptr *QAreaSeries) LowerSeries() *QLineSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerSeries"}).(*QLineSeries)
}

func (ptr *QAreaSeries) Pen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pen"}).(*gui.QPen)
}

func (ptr *QAreaSeries) PointLabelsClipping() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsClipping"}).(bool)
}

func (ptr *QAreaSeries) ConnectPointLabelsClippingChanged(f func(clipping bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointLabelsClippingChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectPointLabelsClippingChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointLabelsClippingChanged"})
}

func (ptr *QAreaSeries) PointLabelsClippingChanged(clipping bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsClippingChanged", clipping})
}

func (ptr *QAreaSeries) PointLabelsColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsColor"}).(*gui.QColor)
}

func (ptr *QAreaSeries) ConnectPointLabelsColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointLabelsColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectPointLabelsColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointLabelsColorChanged"})
}

func (ptr *QAreaSeries) PointLabelsColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsColorChanged", color})
}

func (ptr *QAreaSeries) PointLabelsFont() *gui.QFont {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsFont"}).(*gui.QFont)
}

func (ptr *QAreaSeries) ConnectPointLabelsFontChanged(f func(font *gui.QFont)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointLabelsFontChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectPointLabelsFontChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointLabelsFontChanged"})
}

func (ptr *QAreaSeries) PointLabelsFontChanged(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsFontChanged", font})
}

func (ptr *QAreaSeries) PointLabelsFormat() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsFormat"}).(string)
}

func (ptr *QAreaSeries) ConnectPointLabelsFormatChanged(f func(format string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointLabelsFormatChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectPointLabelsFormatChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointLabelsFormatChanged"})
}

func (ptr *QAreaSeries) PointLabelsFormatChanged(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsFormatChanged", format})
}

func (ptr *QAreaSeries) ConnectPointLabelsVisibilityChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointLabelsVisibilityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectPointLabelsVisibilityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointLabelsVisibilityChanged"})
}

func (ptr *QAreaSeries) PointLabelsVisibilityChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsVisibilityChanged", visible})
}

func (ptr *QAreaSeries) PointLabelsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsVisible"}).(bool)
}

func (ptr *QAreaSeries) PointsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointsVisible"}).(bool)
}

func (ptr *QAreaSeries) ConnectPressed(f func(point *core.QPointF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPressed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectPressed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPressed"})
}

func (ptr *QAreaSeries) Pressed(point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressed", point})
}

func (ptr *QAreaSeries) ConnectReleased(f func(point *core.QPointF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleased", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectReleased() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleased"})
}

func (ptr *QAreaSeries) Released(point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Released", point})
}

func (ptr *QAreaSeries) SetBorderColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBorderColor", color})
}

func (ptr *QAreaSeries) SetBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrush", brush})
}

func (ptr *QAreaSeries) SetColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColor", color})
}

func (ptr *QAreaSeries) SetLowerSeries(series QLineSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLowerSeries", series})
}

func (ptr *QAreaSeries) SetPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPen", pen})
}

func (ptr *QAreaSeries) SetPointLabelsClipping(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointLabelsClipping", enabled})
}

func (ptr *QAreaSeries) SetPointLabelsColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointLabelsColor", color})
}

func (ptr *QAreaSeries) SetPointLabelsFont(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointLabelsFont", font})
}

func (ptr *QAreaSeries) SetPointLabelsFormat(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointLabelsFormat", format})
}

func (ptr *QAreaSeries) SetPointLabelsVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointLabelsVisible", visible})
}

func (ptr *QAreaSeries) SetPointsVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointsVisible", visible})
}

func (ptr *QAreaSeries) SetUpperSeries(series QLineSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUpperSeries", series})
}

func (ptr *QAreaSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QAreaSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QAreaSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QAreaSeries) UpperSeries() *QLineSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpperSeries"}).(*QLineSeries)
}

func (ptr *QAreaSeries) ConnectDestroyQAreaSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAreaSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAreaSeries) DisconnectDestroyQAreaSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAreaSeries"})
}

func (ptr *QAreaSeries) DestroyQAreaSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAreaSeries"})
}

func (ptr *QAreaSeries) DestroyQAreaSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAreaSeriesDefault"})
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

func (n *QBarCategoryAxis) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractAxis_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBarCategoryAxis) ClassNameInternalF() string {
	return n.QAbstractAxis_PTR().ClassNameInternalF()
}

func NewQBarCategoryAxisFromPointer(ptr unsafe.Pointer) (n *QBarCategoryAxis) {
	n = new(QBarCategoryAxis)
	n.InitFromInternal(uintptr(ptr), "charts.QBarCategoryAxis")
	return
}
func NewQBarCategoryAxis(parent core.QObject_ITF) *QBarCategoryAxis {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQBarCategoryAxis", "", parent}).(*QBarCategoryAxis)
}

func (ptr *QBarCategoryAxis) Append(categories []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", categories})
}

func (ptr *QBarCategoryAxis) Append2(category string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append2", category})
}

func (ptr *QBarCategoryAxis) At(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "At", index}).(string)
}

func (ptr *QBarCategoryAxis) Categories() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Categories"}).([]string)
}

func (ptr *QBarCategoryAxis) ConnectCategoriesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarCategoryAxis) DisconnectCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCategoriesChanged"})
}

func (ptr *QBarCategoryAxis) CategoriesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CategoriesChanged"})
}

func (ptr *QBarCategoryAxis) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QBarCategoryAxis) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QBarCategoryAxis) ConnectCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarCategoryAxis) DisconnectCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCountChanged"})
}

func (ptr *QBarCategoryAxis) CountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CountChanged"})
}

func (ptr *QBarCategoryAxis) Insert(index int, category string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", index, category})
}

func (ptr *QBarCategoryAxis) Max() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Max"}).(string)
}

func (ptr *QBarCategoryAxis) ConnectMaxChanged(f func(max string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMaxChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarCategoryAxis) DisconnectMaxChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMaxChanged"})
}

func (ptr *QBarCategoryAxis) MaxChanged(max string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxChanged", max})
}

func (ptr *QBarCategoryAxis) Min() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Min"}).(string)
}

func (ptr *QBarCategoryAxis) ConnectMinChanged(f func(min string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarCategoryAxis) DisconnectMinChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinChanged"})
}

func (ptr *QBarCategoryAxis) MinChanged(min string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinChanged", min})
}

func (ptr *QBarCategoryAxis) ConnectRangeChanged(f func(min string, max string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRangeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarCategoryAxis) DisconnectRangeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRangeChanged"})
}

func (ptr *QBarCategoryAxis) RangeChanged(min string, max string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RangeChanged", min, max})
}

func (ptr *QBarCategoryAxis) Remove(category string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", category})
}

func (ptr *QBarCategoryAxis) Replace(oldCategory string, newCategory string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Replace", oldCategory, newCategory})
}

func (ptr *QBarCategoryAxis) SetCategories(categories []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCategories", categories})
}

func (ptr *QBarCategoryAxis) SetMax(max string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMax", max})
}

func (ptr *QBarCategoryAxis) SetMin(min string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMin", min})
}

func (ptr *QBarCategoryAxis) SetRange(minCategory string, maxCategory string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRange", minCategory, maxCategory})
}

func (ptr *QBarCategoryAxis) ConnectType(f func() QAbstractAxis__AxisType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarCategoryAxis) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QBarCategoryAxis) Type() QAbstractAxis__AxisType {

	return QAbstractAxis__AxisType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QBarCategoryAxis) TypeDefault() QAbstractAxis__AxisType {

	return QAbstractAxis__AxisType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QBarCategoryAxis) ConnectDestroyQBarCategoryAxis(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBarCategoryAxis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarCategoryAxis) DisconnectDestroyQBarCategoryAxis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBarCategoryAxis"})
}

func (ptr *QBarCategoryAxis) DestroyQBarCategoryAxis() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarCategoryAxis"})
}

func (ptr *QBarCategoryAxis) DestroyQBarCategoryAxisDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarCategoryAxisDefault"})
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

func (n *QBarLegendMarker) InitFromInternal(ptr uintptr, name string) {
	n.QLegendMarker_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBarLegendMarker) ClassNameInternalF() string {
	return n.QLegendMarker_PTR().ClassNameInternalF()
}

func NewQBarLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QBarLegendMarker) {
	n = new(QBarLegendMarker)
	n.InitFromInternal(uintptr(ptr), "charts.QBarLegendMarker")
	return
}
func (ptr *QBarLegendMarker) Barset() *QBarSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Barset"}).(*QBarSet)
}

func (ptr *QBarLegendMarker) ConnectSeries(f func() *QAbstractBarSeries) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarLegendMarker) DisconnectSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeries"})
}

func (ptr *QBarLegendMarker) Series() *QAbstractBarSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QAbstractBarSeries)
}

func (ptr *QBarLegendMarker) SeriesDefault() *QAbstractBarSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesDefault"}).(*QAbstractBarSeries)
}

func (ptr *QBarLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarLegendMarker) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QBarLegendMarker) Type() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QBarLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QBarLegendMarker) ConnectDestroyQBarLegendMarker(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBarLegendMarker", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarLegendMarker) DisconnectDestroyQBarLegendMarker() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBarLegendMarker"})
}

func (ptr *QBarLegendMarker) DestroyQBarLegendMarker() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarLegendMarker"})
}

func (ptr *QBarLegendMarker) DestroyQBarLegendMarkerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarLegendMarkerDefault"})
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

func (n *QBarSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractBarSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBarSeries) ClassNameInternalF() string {
	return n.QAbstractBarSeries_PTR().ClassNameInternalF()
}

func NewQBarSeriesFromPointer(ptr unsafe.Pointer) (n *QBarSeries) {
	n = new(QBarSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QBarSeries")
	return
}
func NewQBarSeries(parent core.QObject_ITF) *QBarSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQBarSeries", "", parent}).(*QBarSeries)
}

func (ptr *QBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QBarSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QBarSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QBarSeries) ConnectDestroyQBarSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBarSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSeries) DisconnectDestroyQBarSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBarSeries"})
}

func (ptr *QBarSeries) DestroyQBarSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarSeries"})
}

func (ptr *QBarSeries) DestroyQBarSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarSeriesDefault"})
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

func (n *QBarSet) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBarSet) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQBarSetFromPointer(ptr unsafe.Pointer) (n *QBarSet) {
	n = new(QBarSet)
	n.InitFromInternal(uintptr(ptr), "charts.QBarSet")
	return
}
func NewQBarSet(label string, parent core.QObject_ITF) *QBarSet {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQBarSet", "", label, parent}).(*QBarSet)
}

func (ptr *QBarSet) Append(value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", value})
}

func (ptr *QBarSet) Append2(values []float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append2", values})
}

func (ptr *QBarSet) At(index int) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "At", index}).(float64)
}

func (ptr *QBarSet) BorderColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderColor"}).(*gui.QColor)
}

func (ptr *QBarSet) ConnectBorderColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBorderColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectBorderColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBorderColorChanged"})
}

func (ptr *QBarSet) BorderColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderColorChanged", color})
}

func (ptr *QBarSet) Brush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Brush"}).(*gui.QBrush)
}

func (ptr *QBarSet) ConnectBrushChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBrushChanged"})
}

func (ptr *QBarSet) BrushChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BrushChanged"})
}

func (ptr *QBarSet) ConnectClicked(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QBarSet) Clicked(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked", index})
}

func (ptr *QBarSet) Color() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Color"}).(*gui.QColor)
}

func (ptr *QBarSet) ConnectColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColorChanged"})
}

func (ptr *QBarSet) ColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorChanged", color})
}

func (ptr *QBarSet) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QBarSet) ConnectDoubleClicked(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDoubleClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectDoubleClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDoubleClicked"})
}

func (ptr *QBarSet) DoubleClicked(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoubleClicked", index})
}

func (ptr *QBarSet) ConnectHovered(f func(status bool, index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QBarSet) Hovered(status bool, index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", status, index})
}

func (ptr *QBarSet) Insert(index int, value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", index, value})
}

func (ptr *QBarSet) Label() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Label"}).(string)
}

func (ptr *QBarSet) LabelBrush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBrush"}).(*gui.QBrush)
}

func (ptr *QBarSet) ConnectLabelBrushChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectLabelBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelBrushChanged"})
}

func (ptr *QBarSet) LabelBrushChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBrushChanged"})
}

func (ptr *QBarSet) ConnectLabelChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectLabelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelChanged"})
}

func (ptr *QBarSet) LabelChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelChanged"})
}

func (ptr *QBarSet) LabelColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelColor"}).(*gui.QColor)
}

func (ptr *QBarSet) ConnectLabelColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectLabelColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelColorChanged"})
}

func (ptr *QBarSet) LabelColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelColorChanged", color})
}

func (ptr *QBarSet) LabelFont() *gui.QFont {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelFont"}).(*gui.QFont)
}

func (ptr *QBarSet) ConnectLabelFontChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelFontChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectLabelFontChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelFontChanged"})
}

func (ptr *QBarSet) LabelFontChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelFontChanged"})
}

func (ptr *QBarSet) Pen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pen"}).(*gui.QPen)
}

func (ptr *QBarSet) ConnectPenChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectPenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPenChanged"})
}

func (ptr *QBarSet) PenChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PenChanged"})
}

func (ptr *QBarSet) ConnectPressed(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPressed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectPressed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPressed"})
}

func (ptr *QBarSet) Pressed(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressed", index})
}

func (ptr *QBarSet) ConnectReleased(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleased", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectReleased() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleased"})
}

func (ptr *QBarSet) Released(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Released", index})
}

func (ptr *QBarSet) Remove(index int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", index, count})
}

func (ptr *QBarSet) Replace(index int, value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Replace", index, value})
}

func (ptr *QBarSet) SetBorderColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBorderColor", color})
}

func (ptr *QBarSet) SetBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrush", brush})
}

func (ptr *QBarSet) SetColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColor", color})
}

func (ptr *QBarSet) SetLabel(label string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabel", label})
}

func (ptr *QBarSet) SetLabelBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelBrush", brush})
}

func (ptr *QBarSet) SetLabelColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelColor", color})
}

func (ptr *QBarSet) SetLabelFont(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelFont", font})
}

func (ptr *QBarSet) SetPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPen", pen})
}

func (ptr *QBarSet) Sum() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Sum"}).(float64)
}

func (ptr *QBarSet) ConnectValueChanged(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValueChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectValueChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValueChanged"})
}

func (ptr *QBarSet) ValueChanged(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueChanged", index})
}

func (ptr *QBarSet) ConnectValuesAdded(f func(index int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValuesAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectValuesAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValuesAdded"})
}

func (ptr *QBarSet) ValuesAdded(index int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValuesAdded", index, count})
}

func (ptr *QBarSet) ConnectValuesRemoved(f func(index int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValuesRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectValuesRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValuesRemoved"})
}

func (ptr *QBarSet) ValuesRemoved(index int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValuesRemoved", index, count})
}

func (ptr *QBarSet) ConnectDestroyQBarSet(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBarSet", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarSet) DisconnectDestroyQBarSet() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBarSet"})
}

func (ptr *QBarSet) DestroyQBarSet() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarSet"})
}

func (ptr *QBarSet) DestroyQBarSetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarSetDefault"})
}

func (ptr *QBarSet) __append_values_atList2(i int) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_values_atList2", i}).(float64)
}

func (ptr *QBarSet) __append_values_setList2(i float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_values_setList2", i})
}

func (ptr *QBarSet) __append_values_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_values_newList2"}).(unsafe.Pointer)
}

func (ptr *QBarSet) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QBarSet) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QBarSet) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QBarSet) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QBarSet) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QBarSet) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QBarSet) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QBarSet) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QBarSet) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QBarSet) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QBarSet) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QBarSet) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QBarSet) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QBarSet) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QBarSet) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QBarSet) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QBarSet) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QBarSet) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QBarSet) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QBarSet) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QBarSet) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QBoxPlotLegendMarker) InitFromInternal(ptr uintptr, name string) {
	n.QLegendMarker_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBoxPlotLegendMarker) ClassNameInternalF() string {
	return n.QLegendMarker_PTR().ClassNameInternalF()
}

func NewQBoxPlotLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QBoxPlotLegendMarker) {
	n = new(QBoxPlotLegendMarker)
	n.InitFromInternal(uintptr(ptr), "charts.QBoxPlotLegendMarker")
	return
}
func (ptr *QBoxPlotLegendMarker) ConnectSeries(f func() *QBoxPlotSeries) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotLegendMarker) DisconnectSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeries"})
}

func (ptr *QBoxPlotLegendMarker) Series() *QBoxPlotSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QBoxPlotSeries)
}

func (ptr *QBoxPlotLegendMarker) SeriesDefault() *QBoxPlotSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesDefault"}).(*QBoxPlotSeries)
}

func (ptr *QBoxPlotLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotLegendMarker) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QBoxPlotLegendMarker) Type() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QBoxPlotLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QBoxPlotLegendMarker) ConnectDestroyQBoxPlotLegendMarker(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBoxPlotLegendMarker", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotLegendMarker) DisconnectDestroyQBoxPlotLegendMarker() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBoxPlotLegendMarker"})
}

func (ptr *QBoxPlotLegendMarker) DestroyQBoxPlotLegendMarker() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBoxPlotLegendMarker"})
}

func (ptr *QBoxPlotLegendMarker) DestroyQBoxPlotLegendMarkerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBoxPlotLegendMarkerDefault"})
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

func (n *QBoxPlotSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBoxPlotSeries) ClassNameInternalF() string {
	return n.QAbstractSeries_PTR().ClassNameInternalF()
}

func NewQBoxPlotSeriesFromPointer(ptr unsafe.Pointer) (n *QBoxPlotSeries) {
	n = new(QBoxPlotSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QBoxPlotSeries")
	return
}
func NewQBoxPlotSeries(parent core.QObject_ITF) *QBoxPlotSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQBoxPlotSeries", "", parent}).(*QBoxPlotSeries)
}

func (ptr *QBoxPlotSeries) Append(set QBoxSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", set}).(bool)
}

func (ptr *QBoxPlotSeries) Append2(sets []*QBoxSet) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append2", sets}).(bool)
}

func (ptr *QBoxPlotSeries) ConnectBoxOutlineVisibilityChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBoxOutlineVisibilityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectBoxOutlineVisibilityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBoxOutlineVisibilityChanged"})
}

func (ptr *QBoxPlotSeries) BoxOutlineVisibilityChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoxOutlineVisibilityChanged"})
}

func (ptr *QBoxPlotSeries) BoxOutlineVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoxOutlineVisible"}).(bool)
}

func (ptr *QBoxPlotSeries) BoxSets() []*QBoxSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoxSets"}).([]*QBoxSet)
}

func (ptr *QBoxPlotSeries) BoxWidth() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoxWidth"}).(float64)
}

func (ptr *QBoxPlotSeries) ConnectBoxWidthChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBoxWidthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectBoxWidthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBoxWidthChanged"})
}

func (ptr *QBoxPlotSeries) BoxWidthChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoxWidthChanged"})
}

func (ptr *QBoxPlotSeries) ConnectBoxsetsAdded(f func(sets []*QBoxSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBoxsetsAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectBoxsetsAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBoxsetsAdded"})
}

func (ptr *QBoxPlotSeries) BoxsetsAdded(sets []*QBoxSet) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoxsetsAdded", sets})
}

func (ptr *QBoxPlotSeries) ConnectBoxsetsRemoved(f func(sets []*QBoxSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBoxsetsRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectBoxsetsRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBoxsetsRemoved"})
}

func (ptr *QBoxPlotSeries) BoxsetsRemoved(sets []*QBoxSet) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoxsetsRemoved", sets})
}

func (ptr *QBoxPlotSeries) Brush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Brush"}).(*gui.QBrush)
}

func (ptr *QBoxPlotSeries) ConnectBrushChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBrushChanged"})
}

func (ptr *QBoxPlotSeries) BrushChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BrushChanged"})
}

func (ptr *QBoxPlotSeries) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QBoxPlotSeries) ConnectClicked(f func(boxset *QBoxSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QBoxPlotSeries) Clicked(boxset QBoxSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked", boxset})
}

func (ptr *QBoxPlotSeries) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QBoxPlotSeries) ConnectCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCountChanged"})
}

func (ptr *QBoxPlotSeries) CountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CountChanged"})
}

func (ptr *QBoxPlotSeries) ConnectDoubleClicked(f func(boxset *QBoxSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDoubleClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectDoubleClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDoubleClicked"})
}

func (ptr *QBoxPlotSeries) DoubleClicked(boxset QBoxSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoubleClicked", boxset})
}

func (ptr *QBoxPlotSeries) ConnectHovered(f func(status bool, boxset *QBoxSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QBoxPlotSeries) Hovered(status bool, boxset QBoxSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", status, boxset})
}

func (ptr *QBoxPlotSeries) Insert(index int, set QBoxSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", index, set}).(bool)
}

func (ptr *QBoxPlotSeries) Pen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pen"}).(*gui.QPen)
}

func (ptr *QBoxPlotSeries) ConnectPenChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectPenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPenChanged"})
}

func (ptr *QBoxPlotSeries) PenChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PenChanged"})
}

func (ptr *QBoxPlotSeries) ConnectPressed(f func(boxset *QBoxSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPressed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectPressed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPressed"})
}

func (ptr *QBoxPlotSeries) Pressed(boxset QBoxSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressed", boxset})
}

func (ptr *QBoxPlotSeries) ConnectReleased(f func(boxset *QBoxSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleased", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectReleased() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleased"})
}

func (ptr *QBoxPlotSeries) Released(boxset QBoxSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Released", boxset})
}

func (ptr *QBoxPlotSeries) Remove(set QBoxSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", set}).(bool)
}

func (ptr *QBoxPlotSeries) SetBoxOutlineVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBoxOutlineVisible", visible})
}

func (ptr *QBoxPlotSeries) SetBoxWidth(width float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBoxWidth", width})
}

func (ptr *QBoxPlotSeries) SetBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrush", brush})
}

func (ptr *QBoxPlotSeries) SetPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPen", pen})
}

func (ptr *QBoxPlotSeries) Take(set QBoxSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Take", set}).(bool)
}

func (ptr *QBoxPlotSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QBoxPlotSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QBoxPlotSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QBoxPlotSeries) ConnectDestroyQBoxPlotSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBoxPlotSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxPlotSeries) DisconnectDestroyQBoxPlotSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBoxPlotSeries"})
}

func (ptr *QBoxPlotSeries) DestroyQBoxPlotSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBoxPlotSeries"})
}

func (ptr *QBoxPlotSeries) DestroyQBoxPlotSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBoxPlotSeriesDefault"})
}

func (ptr *QBoxPlotSeries) __append_sets_atList2(i int) *QBoxSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_sets_atList2", i}).(*QBoxSet)
}

func (ptr *QBoxPlotSeries) __append_sets_setList2(i QBoxSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_sets_setList2", i})
}

func (ptr *QBoxPlotSeries) __append_sets_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_sets_newList2"}).(unsafe.Pointer)
}

func (ptr *QBoxPlotSeries) __boxSets_atList(i int) *QBoxSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boxSets_atList", i}).(*QBoxSet)
}

func (ptr *QBoxPlotSeries) __boxSets_setList(i QBoxSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boxSets_setList", i})
}

func (ptr *QBoxPlotSeries) __boxSets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boxSets_newList"}).(unsafe.Pointer)
}

func (ptr *QBoxPlotSeries) __boxsetsAdded_sets_atList(i int) *QBoxSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boxsetsAdded_sets_atList", i}).(*QBoxSet)
}

func (ptr *QBoxPlotSeries) __boxsetsAdded_sets_setList(i QBoxSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boxsetsAdded_sets_setList", i})
}

func (ptr *QBoxPlotSeries) __boxsetsAdded_sets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boxsetsAdded_sets_newList"}).(unsafe.Pointer)
}

func (ptr *QBoxPlotSeries) __boxsetsRemoved_sets_atList(i int) *QBoxSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boxsetsRemoved_sets_atList", i}).(*QBoxSet)
}

func (ptr *QBoxPlotSeries) __boxsetsRemoved_sets_setList(i QBoxSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boxsetsRemoved_sets_setList", i})
}

func (ptr *QBoxPlotSeries) __boxsetsRemoved_sets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boxsetsRemoved_sets_newList"}).(unsafe.Pointer)
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

func (n *QBoxSet) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBoxSet) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQBoxSetFromPointer(ptr unsafe.Pointer) (n *QBoxSet) {
	n = new(QBoxSet)
	n.InitFromInternal(uintptr(ptr), "charts.QBoxSet")
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

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQBoxSet", "", label, parent}).(*QBoxSet)
}

func NewQBoxSet2(le float64, lq float64, m float64, uq float64, ue float64, label string, parent core.QObject_ITF) *QBoxSet {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQBoxSet2", "", le, lq, m, uq, ue, label, parent}).(*QBoxSet)
}

func (ptr *QBoxSet) Append(value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", value})
}

func (ptr *QBoxSet) Append2(values []float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append2", values})
}

func (ptr *QBoxSet) At(index int) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "At", index}).(float64)
}

func (ptr *QBoxSet) Brush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Brush"}).(*gui.QBrush)
}

func (ptr *QBoxSet) ConnectBrushChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBrushChanged"})
}

func (ptr *QBoxSet) BrushChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BrushChanged"})
}

func (ptr *QBoxSet) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QBoxSet) ConnectCleared(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCleared", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectCleared() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCleared"})
}

func (ptr *QBoxSet) Cleared() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Cleared"})
}

func (ptr *QBoxSet) ConnectClicked(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QBoxSet) Clicked() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked"})
}

func (ptr *QBoxSet) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QBoxSet) ConnectDoubleClicked(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDoubleClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectDoubleClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDoubleClicked"})
}

func (ptr *QBoxSet) DoubleClicked() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoubleClicked"})
}

func (ptr *QBoxSet) ConnectHovered(f func(status bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QBoxSet) Hovered(status bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", status})
}

func (ptr *QBoxSet) Label() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Label"}).(string)
}

func (ptr *QBoxSet) Pen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pen"}).(*gui.QPen)
}

func (ptr *QBoxSet) ConnectPenChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectPenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPenChanged"})
}

func (ptr *QBoxSet) PenChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PenChanged"})
}

func (ptr *QBoxSet) ConnectPressed(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPressed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectPressed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPressed"})
}

func (ptr *QBoxSet) Pressed() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressed"})
}

func (ptr *QBoxSet) ConnectReleased(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleased", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectReleased() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleased"})
}

func (ptr *QBoxSet) Released() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Released"})
}

func (ptr *QBoxSet) SetBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrush", brush})
}

func (ptr *QBoxSet) SetLabel(label string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabel", label})
}

func (ptr *QBoxSet) SetPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPen", pen})
}

func (ptr *QBoxSet) SetValue(index int, value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", index, value})
}

func (ptr *QBoxSet) ConnectValueChanged(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValueChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectValueChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValueChanged"})
}

func (ptr *QBoxSet) ValueChanged(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueChanged", index})
}

func (ptr *QBoxSet) ConnectValuesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValuesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectValuesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValuesChanged"})
}

func (ptr *QBoxSet) ValuesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValuesChanged"})
}

func (ptr *QBoxSet) ConnectDestroyQBoxSet(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBoxSet", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBoxSet) DisconnectDestroyQBoxSet() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBoxSet"})
}

func (ptr *QBoxSet) DestroyQBoxSet() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBoxSet"})
}

func (ptr *QBoxSet) DestroyQBoxSetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBoxSetDefault"})
}

func (ptr *QBoxSet) __append_values_atList2(i int) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_values_atList2", i}).(float64)
}

func (ptr *QBoxSet) __append_values_setList2(i float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_values_setList2", i})
}

func (ptr *QBoxSet) __append_values_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_values_newList2"}).(unsafe.Pointer)
}

func (ptr *QBoxSet) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QBoxSet) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QBoxSet) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QBoxSet) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QBoxSet) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QBoxSet) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QBoxSet) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QBoxSet) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QBoxSet) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QBoxSet) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QBoxSet) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QBoxSet) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QBoxSet) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QBoxSet) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QBoxSet) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QBoxSet) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QBoxSet) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QBoxSet) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QBoxSet) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QBoxSet) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QBoxSet) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QCandlestickLegendMarker) InitFromInternal(ptr uintptr, name string) {
	n.QLegendMarker_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCandlestickLegendMarker) ClassNameInternalF() string {
	return n.QLegendMarker_PTR().ClassNameInternalF()
}

func NewQCandlestickLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QCandlestickLegendMarker) {
	n = new(QCandlestickLegendMarker)
	n.InitFromInternal(uintptr(ptr), "charts.QCandlestickLegendMarker")
	return
}

func (ptr *QCandlestickLegendMarker) DestroyQCandlestickLegendMarker() {
}

func (ptr *QCandlestickLegendMarker) ConnectSeries(f func() *QCandlestickSeries) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickLegendMarker) DisconnectSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeries"})
}

func (ptr *QCandlestickLegendMarker) Series() *QCandlestickSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QCandlestickSeries)
}

func (ptr *QCandlestickLegendMarker) SeriesDefault() *QCandlestickSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesDefault"}).(*QCandlestickSeries)
}

func (ptr *QCandlestickLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickLegendMarker) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QCandlestickLegendMarker) Type() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QCandlestickLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
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

func (n *QCandlestickModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCandlestickModelMapper) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQCandlestickModelMapperFromPointer(ptr unsafe.Pointer) (n *QCandlestickModelMapper) {
	n = new(QCandlestickModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QCandlestickModelMapper")
	return
}

func (ptr *QCandlestickModelMapper) DestroyQCandlestickModelMapper() {
}

func NewQCandlestickModelMapper(parent core.QObject_ITF) *QCandlestickModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQCandlestickModelMapper", "", parent}).(*QCandlestickModelMapper)
}

func (ptr *QCandlestickModelMapper) Close() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"}).(float64))
}

func (ptr *QCandlestickModelMapper) FirstSetSection() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstSetSection"}).(float64))
}

func (ptr *QCandlestickModelMapper) High() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "High"}).(float64))
}

func (ptr *QCandlestickModelMapper) LastSetSection() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastSetSection"}).(float64))
}

func (ptr *QCandlestickModelMapper) Low() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Low"}).(float64))
}

func (ptr *QCandlestickModelMapper) Model() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*core.QAbstractItemModel)
}

func (ptr *QCandlestickModelMapper) ConnectModelReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModelReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickModelMapper) DisconnectModelReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModelReplaced"})
}

func (ptr *QCandlestickModelMapper) ModelReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModelReplaced"})
}

func (ptr *QCandlestickModelMapper) Open() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(float64))
}

func (ptr *QCandlestickModelMapper) ConnectOrientation(f func() core.Qt__Orientation) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOrientation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickModelMapper) DisconnectOrientation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOrientation"})
}

func (ptr *QCandlestickModelMapper) Orientation() core.Qt__Orientation {

	return core.Qt__Orientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Orientation"}).(float64))
}

func (ptr *QCandlestickModelMapper) Series() *QCandlestickSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QCandlestickSeries)
}

func (ptr *QCandlestickModelMapper) ConnectSeriesReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickModelMapper) DisconnectSeriesReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesReplaced"})
}

func (ptr *QCandlestickModelMapper) SeriesReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesReplaced"})
}

func (ptr *QCandlestickModelMapper) SetClose(close int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetClose", close})
}

func (ptr *QCandlestickModelMapper) SetFirstSetSection(firstSetSection int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstSetSection", firstSetSection})
}

func (ptr *QCandlestickModelMapper) SetHigh(high int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHigh", high})
}

func (ptr *QCandlestickModelMapper) SetLastSetSection(lastSetSection int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastSetSection", lastSetSection})
}

func (ptr *QCandlestickModelMapper) SetLow(low int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLow", low})
}

func (ptr *QCandlestickModelMapper) SetModel(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModel", model})
}

func (ptr *QCandlestickModelMapper) SetOpen(open int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpen", open})
}

func (ptr *QCandlestickModelMapper) SetSeries(series QCandlestickSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSeries", series})
}

func (ptr *QCandlestickModelMapper) SetTimestamp(timestamp int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTimestamp", timestamp})
}

func (ptr *QCandlestickModelMapper) Timestamp() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Timestamp"}).(float64))
}

func (ptr *QCandlestickModelMapper) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QCandlestickModelMapper) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QCandlestickModelMapper) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QCandlestickModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QCandlestickModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QCandlestickModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QCandlestickModelMapper) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QCandlestickModelMapper) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QCandlestickModelMapper) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QCandlestickModelMapper) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QCandlestickModelMapper) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QCandlestickModelMapper) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QCandlestickModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QCandlestickModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QCandlestickModelMapper) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QCandlestickModelMapper) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QCandlestickModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QCandlestickModelMapper) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QCandlestickModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QCandlestickModelMapper) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QCandlestickModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QCandlestickSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCandlestickSeries) ClassNameInternalF() string {
	return n.QAbstractSeries_PTR().ClassNameInternalF()
}

func NewQCandlestickSeriesFromPointer(ptr unsafe.Pointer) (n *QCandlestickSeries) {
	n = new(QCandlestickSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QCandlestickSeries")
	return
}
func NewQCandlestickSeries(parent core.QObject_ITF) *QCandlestickSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQCandlestickSeries", "", parent}).(*QCandlestickSeries)
}

func (ptr *QCandlestickSeries) Append(set QCandlestickSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", set}).(bool)
}

func (ptr *QCandlestickSeries) Append2(sets []*QCandlestickSet) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append2", sets}).(bool)
}

func (ptr *QCandlestickSeries) ConnectBodyOutlineVisibilityChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBodyOutlineVisibilityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectBodyOutlineVisibilityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBodyOutlineVisibilityChanged"})
}

func (ptr *QCandlestickSeries) BodyOutlineVisibilityChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BodyOutlineVisibilityChanged"})
}

func (ptr *QCandlestickSeries) BodyOutlineVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BodyOutlineVisible"}).(bool)
}

func (ptr *QCandlestickSeries) BodyWidth() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BodyWidth"}).(float64)
}

func (ptr *QCandlestickSeries) ConnectBodyWidthChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBodyWidthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectBodyWidthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBodyWidthChanged"})
}

func (ptr *QCandlestickSeries) BodyWidthChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BodyWidthChanged"})
}

func (ptr *QCandlestickSeries) Brush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Brush"}).(*gui.QBrush)
}

func (ptr *QCandlestickSeries) ConnectBrushChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBrushChanged"})
}

func (ptr *QCandlestickSeries) BrushChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BrushChanged"})
}

func (ptr *QCandlestickSeries) ConnectCandlestickSetsAdded(f func(sets []*QCandlestickSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCandlestickSetsAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectCandlestickSetsAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCandlestickSetsAdded"})
}

func (ptr *QCandlestickSeries) CandlestickSetsAdded(sets []*QCandlestickSet) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CandlestickSetsAdded", sets})
}

func (ptr *QCandlestickSeries) ConnectCandlestickSetsRemoved(f func(sets []*QCandlestickSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCandlestickSetsRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectCandlestickSetsRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCandlestickSetsRemoved"})
}

func (ptr *QCandlestickSeries) CandlestickSetsRemoved(sets []*QCandlestickSet) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CandlestickSetsRemoved", sets})
}

func (ptr *QCandlestickSeries) ConnectCapsVisibilityChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCapsVisibilityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectCapsVisibilityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCapsVisibilityChanged"})
}

func (ptr *QCandlestickSeries) CapsVisibilityChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CapsVisibilityChanged"})
}

func (ptr *QCandlestickSeries) CapsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CapsVisible"}).(bool)
}

func (ptr *QCandlestickSeries) CapsWidth() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CapsWidth"}).(float64)
}

func (ptr *QCandlestickSeries) ConnectCapsWidthChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCapsWidthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectCapsWidthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCapsWidthChanged"})
}

func (ptr *QCandlestickSeries) CapsWidthChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CapsWidthChanged"})
}

func (ptr *QCandlestickSeries) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QCandlestickSeries) ConnectClicked(f func(set *QCandlestickSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QCandlestickSeries) Clicked(set QCandlestickSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked", set})
}

func (ptr *QCandlestickSeries) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QCandlestickSeries) ConnectCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCountChanged"})
}

func (ptr *QCandlestickSeries) CountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CountChanged"})
}

func (ptr *QCandlestickSeries) DecreasingColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DecreasingColor"}).(*gui.QColor)
}

func (ptr *QCandlestickSeries) ConnectDecreasingColorChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDecreasingColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectDecreasingColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDecreasingColorChanged"})
}

func (ptr *QCandlestickSeries) DecreasingColorChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DecreasingColorChanged"})
}

func (ptr *QCandlestickSeries) ConnectDoubleClicked(f func(set *QCandlestickSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDoubleClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectDoubleClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDoubleClicked"})
}

func (ptr *QCandlestickSeries) DoubleClicked(set QCandlestickSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoubleClicked", set})
}

func (ptr *QCandlestickSeries) ConnectHovered(f func(status bool, set *QCandlestickSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QCandlestickSeries) Hovered(status bool, set QCandlestickSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", status, set})
}

func (ptr *QCandlestickSeries) IncreasingColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncreasingColor"}).(*gui.QColor)
}

func (ptr *QCandlestickSeries) ConnectIncreasingColorChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIncreasingColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectIncreasingColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIncreasingColorChanged"})
}

func (ptr *QCandlestickSeries) IncreasingColorChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncreasingColorChanged"})
}

func (ptr *QCandlestickSeries) Insert(index int, set QCandlestickSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", index, set}).(bool)
}

func (ptr *QCandlestickSeries) MaximumColumnWidth() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumColumnWidth"}).(float64)
}

func (ptr *QCandlestickSeries) ConnectMaximumColumnWidthChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMaximumColumnWidthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectMaximumColumnWidthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMaximumColumnWidthChanged"})
}

func (ptr *QCandlestickSeries) MaximumColumnWidthChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumColumnWidthChanged"})
}

func (ptr *QCandlestickSeries) MinimumColumnWidth() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumColumnWidth"}).(float64)
}

func (ptr *QCandlestickSeries) ConnectMinimumColumnWidthChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinimumColumnWidthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectMinimumColumnWidthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinimumColumnWidthChanged"})
}

func (ptr *QCandlestickSeries) MinimumColumnWidthChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumColumnWidthChanged"})
}

func (ptr *QCandlestickSeries) Pen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pen"}).(*gui.QPen)
}

func (ptr *QCandlestickSeries) ConnectPenChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectPenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPenChanged"})
}

func (ptr *QCandlestickSeries) PenChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PenChanged"})
}

func (ptr *QCandlestickSeries) ConnectPressed(f func(set *QCandlestickSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPressed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectPressed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPressed"})
}

func (ptr *QCandlestickSeries) Pressed(set QCandlestickSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressed", set})
}

func (ptr *QCandlestickSeries) ConnectReleased(f func(set *QCandlestickSet)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleased", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectReleased() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleased"})
}

func (ptr *QCandlestickSeries) Released(set QCandlestickSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Released", set})
}

func (ptr *QCandlestickSeries) Remove(set QCandlestickSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", set}).(bool)
}

func (ptr *QCandlestickSeries) Remove2(sets []*QCandlestickSet) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove2", sets}).(bool)
}

func (ptr *QCandlestickSeries) SetBodyOutlineVisible(bodyOutlineVisible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBodyOutlineVisible", bodyOutlineVisible})
}

func (ptr *QCandlestickSeries) SetBodyWidth(bodyWidth float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBodyWidth", bodyWidth})
}

func (ptr *QCandlestickSeries) SetBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrush", brush})
}

func (ptr *QCandlestickSeries) SetCapsVisible(capsVisible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCapsVisible", capsVisible})
}

func (ptr *QCandlestickSeries) SetCapsWidth(capsWidth float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCapsWidth", capsWidth})
}

func (ptr *QCandlestickSeries) SetDecreasingColor(decreasingColor gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDecreasingColor", decreasingColor})
}

func (ptr *QCandlestickSeries) SetIncreasingColor(increasingColor gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIncreasingColor", increasingColor})
}

func (ptr *QCandlestickSeries) SetMaximumColumnWidth(maximumColumnWidth float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaximumColumnWidth", maximumColumnWidth})
}

func (ptr *QCandlestickSeries) SetMinimumColumnWidth(minimumColumnWidth float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinimumColumnWidth", minimumColumnWidth})
}

func (ptr *QCandlestickSeries) SetPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPen", pen})
}

func (ptr *QCandlestickSeries) Sets() []*QCandlestickSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Sets"}).([]*QCandlestickSet)
}

func (ptr *QCandlestickSeries) Take(set QCandlestickSet_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Take", set}).(bool)
}

func (ptr *QCandlestickSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QCandlestickSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QCandlestickSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QCandlestickSeries) ConnectDestroyQCandlestickSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQCandlestickSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSeries) DisconnectDestroyQCandlestickSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQCandlestickSeries"})
}

func (ptr *QCandlestickSeries) DestroyQCandlestickSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCandlestickSeries"})
}

func (ptr *QCandlestickSeries) DestroyQCandlestickSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCandlestickSeriesDefault"})
}

func (ptr *QCandlestickSeries) __append_sets_atList2(i int) *QCandlestickSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_sets_atList2", i}).(*QCandlestickSet)
}

func (ptr *QCandlestickSeries) __append_sets_setList2(i QCandlestickSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_sets_setList2", i})
}

func (ptr *QCandlestickSeries) __append_sets_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_sets_newList2"}).(unsafe.Pointer)
}

func (ptr *QCandlestickSeries) __candlestickSetsAdded_sets_atList(i int) *QCandlestickSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__candlestickSetsAdded_sets_atList", i}).(*QCandlestickSet)
}

func (ptr *QCandlestickSeries) __candlestickSetsAdded_sets_setList(i QCandlestickSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__candlestickSetsAdded_sets_setList", i})
}

func (ptr *QCandlestickSeries) __candlestickSetsAdded_sets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__candlestickSetsAdded_sets_newList"}).(unsafe.Pointer)
}

func (ptr *QCandlestickSeries) __candlestickSetsRemoved_sets_atList(i int) *QCandlestickSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__candlestickSetsRemoved_sets_atList", i}).(*QCandlestickSet)
}

func (ptr *QCandlestickSeries) __candlestickSetsRemoved_sets_setList(i QCandlestickSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__candlestickSetsRemoved_sets_setList", i})
}

func (ptr *QCandlestickSeries) __candlestickSetsRemoved_sets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__candlestickSetsRemoved_sets_newList"}).(unsafe.Pointer)
}

func (ptr *QCandlestickSeries) __remove_sets_atList2(i int) *QCandlestickSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__remove_sets_atList2", i}).(*QCandlestickSet)
}

func (ptr *QCandlestickSeries) __remove_sets_setList2(i QCandlestickSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__remove_sets_setList2", i})
}

func (ptr *QCandlestickSeries) __remove_sets_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__remove_sets_newList2"}).(unsafe.Pointer)
}

func (ptr *QCandlestickSeries) __sets_atList(i int) *QCandlestickSet {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sets_atList", i}).(*QCandlestickSet)
}

func (ptr *QCandlestickSeries) __sets_setList(i QCandlestickSet_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sets_setList", i})
}

func (ptr *QCandlestickSeries) __sets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sets_newList"}).(unsafe.Pointer)
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

func (n *QCandlestickSet) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCandlestickSet) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQCandlestickSetFromPointer(ptr unsafe.Pointer) (n *QCandlestickSet) {
	n = new(QCandlestickSet)
	n.InitFromInternal(uintptr(ptr), "charts.QCandlestickSet")
	return
}
func NewQCandlestickSet(timestamp float64, parent core.QObject_ITF) *QCandlestickSet {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQCandlestickSet", "", timestamp, parent}).(*QCandlestickSet)
}

func NewQCandlestickSet2(open float64, high float64, low float64, close float64, timestamp float64, parent core.QObject_ITF) *QCandlestickSet {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQCandlestickSet2", "", open, high, low, close, timestamp, parent}).(*QCandlestickSet)
}

func (ptr *QCandlestickSet) Brush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Brush"}).(*gui.QBrush)
}

func (ptr *QCandlestickSet) ConnectBrushChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBrushChanged"})
}

func (ptr *QCandlestickSet) BrushChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BrushChanged"})
}

func (ptr *QCandlestickSet) ConnectClicked(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QCandlestickSet) Clicked() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked"})
}

func (ptr *QCandlestickSet) Close() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"}).(float64)
}

func (ptr *QCandlestickSet) ConnectCloseChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCloseChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectCloseChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCloseChanged"})
}

func (ptr *QCandlestickSet) CloseChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseChanged"})
}

func (ptr *QCandlestickSet) ConnectDoubleClicked(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDoubleClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectDoubleClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDoubleClicked"})
}

func (ptr *QCandlestickSet) DoubleClicked() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoubleClicked"})
}

func (ptr *QCandlestickSet) High() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "High"}).(float64)
}

func (ptr *QCandlestickSet) ConnectHighChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHighChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectHighChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHighChanged"})
}

func (ptr *QCandlestickSet) HighChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HighChanged"})
}

func (ptr *QCandlestickSet) ConnectHovered(f func(status bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QCandlestickSet) Hovered(status bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", status})
}

func (ptr *QCandlestickSet) Low() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Low"}).(float64)
}

func (ptr *QCandlestickSet) ConnectLowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectLowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLowChanged"})
}

func (ptr *QCandlestickSet) LowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowChanged"})
}

func (ptr *QCandlestickSet) Open() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(float64)
}

func (ptr *QCandlestickSet) ConnectOpenChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectOpenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpenChanged"})
}

func (ptr *QCandlestickSet) OpenChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenChanged"})
}

func (ptr *QCandlestickSet) Pen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pen"}).(*gui.QPen)
}

func (ptr *QCandlestickSet) ConnectPenChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectPenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPenChanged"})
}

func (ptr *QCandlestickSet) PenChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PenChanged"})
}

func (ptr *QCandlestickSet) ConnectPressed(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPressed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectPressed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPressed"})
}

func (ptr *QCandlestickSet) Pressed() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressed"})
}

func (ptr *QCandlestickSet) ConnectReleased(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleased", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectReleased() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleased"})
}

func (ptr *QCandlestickSet) Released() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Released"})
}

func (ptr *QCandlestickSet) SetBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrush", brush})
}

func (ptr *QCandlestickSet) SetClose(close float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetClose", close})
}

func (ptr *QCandlestickSet) SetHigh(high float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHigh", high})
}

func (ptr *QCandlestickSet) SetLow(low float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLow", low})
}

func (ptr *QCandlestickSet) SetOpen(open float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpen", open})
}

func (ptr *QCandlestickSet) SetPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPen", pen})
}

func (ptr *QCandlestickSet) SetTimestamp(timestamp float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTimestamp", timestamp})
}

func (ptr *QCandlestickSet) Timestamp() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Timestamp"}).(float64)
}

func (ptr *QCandlestickSet) ConnectTimestampChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTimestampChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectTimestampChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTimestampChanged"})
}

func (ptr *QCandlestickSet) TimestampChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimestampChanged"})
}

func (ptr *QCandlestickSet) ConnectDestroyQCandlestickSet(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQCandlestickSet", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCandlestickSet) DisconnectDestroyQCandlestickSet() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQCandlestickSet"})
}

func (ptr *QCandlestickSet) DestroyQCandlestickSet() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCandlestickSet"})
}

func (ptr *QCandlestickSet) DestroyQCandlestickSetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCandlestickSetDefault"})
}

func (ptr *QCandlestickSet) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QCandlestickSet) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QCandlestickSet) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QCandlestickSet) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QCandlestickSet) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QCandlestickSet) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QCandlestickSet) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QCandlestickSet) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QCandlestickSet) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QCandlestickSet) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QCandlestickSet) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QCandlestickSet) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QCandlestickSet) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QCandlestickSet) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QCandlestickSet) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QCandlestickSet) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QCandlestickSet) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QCandlestickSet) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QCandlestickSet) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QCandlestickSet) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QCandlestickSet) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QCategoryAxis) InitFromInternal(ptr uintptr, name string) {
	n.QValueAxis_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCategoryAxis) ClassNameInternalF() string {
	return n.QValueAxis_PTR().ClassNameInternalF()
}

func NewQCategoryAxisFromPointer(ptr unsafe.Pointer) (n *QCategoryAxis) {
	n = new(QCategoryAxis)
	n.InitFromInternal(uintptr(ptr), "charts.QCategoryAxis")
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

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQCategoryAxis", "", parent}).(*QCategoryAxis)
}

func (ptr *QCategoryAxis) Append(categoryLabel string, categoryEndValue float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", categoryLabel, categoryEndValue})
}

func (ptr *QCategoryAxis) ConnectCategoriesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCategoryAxis) DisconnectCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCategoriesChanged"})
}

func (ptr *QCategoryAxis) CategoriesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CategoriesChanged"})
}

func (ptr *QCategoryAxis) CategoriesLabels() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CategoriesLabels"}).([]string)
}

func (ptr *QCategoryAxis) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QCategoryAxis) EndValue(categoryLabel string) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndValue", categoryLabel}).(float64)
}

func (ptr *QCategoryAxis) LabelsPosition() QCategoryAxis__AxisLabelsPosition {

	return QCategoryAxis__AxisLabelsPosition(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsPosition"}).(float64))
}

func (ptr *QCategoryAxis) ConnectLabelsPositionChanged(f func(position QCategoryAxis__AxisLabelsPosition)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsPositionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCategoryAxis) DisconnectLabelsPositionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsPositionChanged"})
}

func (ptr *QCategoryAxis) LabelsPositionChanged(position QCategoryAxis__AxisLabelsPosition) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsPositionChanged", position})
}

func (ptr *QCategoryAxis) Remove(categoryLabel string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", categoryLabel})
}

func (ptr *QCategoryAxis) ReplaceLabel(oldLabel string, newLabel string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReplaceLabel", oldLabel, newLabel})
}

func (ptr *QCategoryAxis) SetLabelsPosition(position QCategoryAxis__AxisLabelsPosition) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsPosition", position})
}

func (ptr *QCategoryAxis) SetStartValue(min float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStartValue", min})
}

func (ptr *QCategoryAxis) StartValue(categoryLabel string) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartValue", categoryLabel}).(float64)
}

func (ptr *QCategoryAxis) ConnectDestroyQCategoryAxis(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQCategoryAxis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCategoryAxis) DisconnectDestroyQCategoryAxis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQCategoryAxis"})
}

func (ptr *QCategoryAxis) DestroyQCategoryAxis() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCategoryAxis"})
}

func (ptr *QCategoryAxis) DestroyQCategoryAxisDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCategoryAxisDefault"})
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

func (n *QChart) InitFromInternal(ptr uintptr, name string) {
	n.QGraphicsWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QChart) ClassNameInternalF() string {
	return n.QGraphicsWidget_PTR().ClassNameInternalF()
}

func NewQChartFromPointer(ptr unsafe.Pointer) (n *QChart) {
	n = new(QChart)
	n.InitFromInternal(uintptr(ptr), "charts.QChart")
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

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQChart", "", parent, wFlags}).(*QChart)
}

func (ptr *QChart) AddAxis(axis QAbstractAxis_ITF, alignment core.Qt__AlignmentFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddAxis", axis, alignment})
}

func (ptr *QChart) AddSeries(series QAbstractSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddSeries", series})
}

func (ptr *QChart) AnimationDuration() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnimationDuration"}).(float64))
}

func (ptr *QChart) AnimationEasingCurve() *core.QEasingCurve {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnimationEasingCurve"}).(*core.QEasingCurve)
}

func (ptr *QChart) AnimationOptions() QChart__AnimationOption {

	return QChart__AnimationOption(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnimationOptions"}).(float64))
}

func (ptr *QChart) Axes(orientation core.Qt__Orientation, series QAbstractSeries_ITF) []*QAbstractAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Axes", orientation, series}).([]*QAbstractAxis)
}

func (ptr *QChart) BackgroundBrush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundBrush"}).(*gui.QBrush)
}

func (ptr *QChart) BackgroundPen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundPen"}).(*gui.QPen)
}

func (ptr *QChart) BackgroundRoundness() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundRoundness"}).(float64)
}

func (ptr *QChart) ChartType() QChart__ChartType {

	return QChart__ChartType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChartType"}).(float64))
}

func (ptr *QChart) CreateDefaultAxes() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateDefaultAxes"})
}

func (ptr *QChart) IsBackgroundVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBackgroundVisible"}).(bool)
}

func (ptr *QChart) IsDropShadowEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDropShadowEnabled"}).(bool)
}

func (ptr *QChart) IsPlotAreaBackgroundVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsPlotAreaBackgroundVisible"}).(bool)
}

func (ptr *QChart) IsZoomed() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsZoomed"}).(bool)
}

func (ptr *QChart) Legend() *QLegend {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Legend"}).(*QLegend)
}

func (ptr *QChart) Locale() *core.QLocale {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Locale"}).(*core.QLocale)
}

func (ptr *QChart) LocalizeNumbers() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalizeNumbers"}).(bool)
}

func (ptr *QChart) MapToPosition(value core.QPointF_ITF, series QAbstractSeries_ITF) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MapToPosition", value, series}).(*core.QPointF)
}

func (ptr *QChart) MapToValue(position core.QPointF_ITF, series QAbstractSeries_ITF) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MapToValue", position, series}).(*core.QPointF)
}

func (ptr *QChart) Margins() *core.QMargins {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Margins"}).(*core.QMargins)
}

func (ptr *QChart) PlotArea() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PlotArea"}).(*core.QRectF)
}

func (ptr *QChart) PlotAreaBackgroundBrush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PlotAreaBackgroundBrush"}).(*gui.QBrush)
}

func (ptr *QChart) PlotAreaBackgroundPen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PlotAreaBackgroundPen"}).(*gui.QPen)
}

func (ptr *QChart) ConnectPlotAreaChanged(f func(plotArea *core.QRectF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPlotAreaChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QChart) DisconnectPlotAreaChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPlotAreaChanged"})
}

func (ptr *QChart) PlotAreaChanged(plotArea core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PlotAreaChanged", plotArea})
}

func (ptr *QChart) RemoveAllSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAllSeries"})
}

func (ptr *QChart) RemoveAxis(axis QAbstractAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAxis", axis})
}

func (ptr *QChart) RemoveSeries(series QAbstractSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveSeries", series})
}

func (ptr *QChart) Scroll(dx float64, dy float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Scroll", dx, dy})
}

func (ptr *QChart) Series() []*QAbstractSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).([]*QAbstractSeries)
}

func (ptr *QChart) SetAnimationDuration(msecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAnimationDuration", msecs})
}

func (ptr *QChart) SetAnimationEasingCurve(curve core.QEasingCurve_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAnimationEasingCurve", curve})
}

func (ptr *QChart) SetAnimationOptions(options QChart__AnimationOption) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAnimationOptions", options})
}

func (ptr *QChart) SetBackgroundBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackgroundBrush", brush})
}

func (ptr *QChart) SetBackgroundPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackgroundPen", pen})
}

func (ptr *QChart) SetBackgroundRoundness(diameter float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackgroundRoundness", diameter})
}

func (ptr *QChart) SetBackgroundVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackgroundVisible", visible})
}

func (ptr *QChart) SetDropShadowEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDropShadowEnabled", enabled})
}

func (ptr *QChart) SetLocale(locale core.QLocale_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocale", locale})
}

func (ptr *QChart) SetLocalizeNumbers(localize bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalizeNumbers", localize})
}

func (ptr *QChart) SetMargins(margins core.QMargins_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMargins", margins})
}

func (ptr *QChart) SetPlotArea(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPlotArea", rect})
}

func (ptr *QChart) SetPlotAreaBackgroundBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPlotAreaBackgroundBrush", brush})
}

func (ptr *QChart) SetPlotAreaBackgroundPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPlotAreaBackgroundPen", pen})
}

func (ptr *QChart) SetPlotAreaBackgroundVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPlotAreaBackgroundVisible", visible})
}

func (ptr *QChart) SetTheme(theme QChart__ChartTheme) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTheme", theme})
}

func (ptr *QChart) SetTitle(title string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitle", title})
}

func (ptr *QChart) SetTitleBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitleBrush", brush})
}

func (ptr *QChart) SetTitleFont(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitleFont", font})
}

func (ptr *QChart) Theme() QChart__ChartTheme {

	return QChart__ChartTheme(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Theme"}).(float64))
}

func (ptr *QChart) Title() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Title"}).(string)
}

func (ptr *QChart) TitleBrush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleBrush"}).(*gui.QBrush)
}

func (ptr *QChart) TitleFont() *gui.QFont {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleFont"}).(*gui.QFont)
}

func (ptr *QChart) Zoom(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Zoom", factor})
}

func (ptr *QChart) ZoomIn() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomIn"})
}

func (ptr *QChart) ZoomIn2(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomIn2", rect})
}

func (ptr *QChart) ZoomOut() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomOut"})
}

func (ptr *QChart) ZoomReset() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomReset"})
}

func (ptr *QChart) ConnectDestroyQChart(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQChart", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QChart) DisconnectDestroyQChart() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQChart"})
}

func (ptr *QChart) DestroyQChart() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQChart"})
}

func (ptr *QChart) DestroyQChartDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQChartDefault"})
}

func (ptr *QChart) __axes_atList(i int) *QAbstractAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_atList", i}).(*QAbstractAxis)
}

func (ptr *QChart) __axes_setList(i QAbstractAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_setList", i})
}

func (ptr *QChart) __axes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __series_atList(i int) *QAbstractSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__series_atList", i}).(*QAbstractSeries)
}

func (ptr *QChart) __series_setList(i QAbstractSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__series_setList", i})
}

func (ptr *QChart) __series_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__series_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QChart) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QChart) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QChart) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QChart) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QChart) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QChart) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QChart) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QChart) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QChart) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QChart) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QChart) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QChart) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QChart) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QChart) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QChart) __childItems_atList(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__childItems_atList", i}).(*widgets.QGraphicsItem)
}

func (ptr *QChart) __childItems_setList(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__childItems_setList", i})
}

func (ptr *QChart) __childItems_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__childItems_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __collidingItems_atList(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__collidingItems_atList", i}).(*widgets.QGraphicsItem)
}

func (ptr *QChart) __collidingItems_setList(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__collidingItems_setList", i})
}

func (ptr *QChart) __collidingItems_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__collidingItems_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __setTransformations_transformations_atList(i int) *widgets.QGraphicsTransform {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTransformations_transformations_atList", i}).(*widgets.QGraphicsTransform)
}

func (ptr *QChart) __setTransformations_transformations_setList(i widgets.QGraphicsTransform_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTransformations_transformations_setList", i})
}

func (ptr *QChart) __setTransformations_transformations_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTransformations_transformations_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) __transformations_atList(i int) *widgets.QGraphicsTransform {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__transformations_atList", i}).(*widgets.QGraphicsTransform)
}

func (ptr *QChart) __transformations_setList(i widgets.QGraphicsTransform_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__transformations_setList", i})
}

func (ptr *QChart) __transformations_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__transformations_newList"}).(unsafe.Pointer)
}

func (ptr *QChart) BoundingRectDefault() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundingRectDefault"}).(*core.QRectF)
}

func (ptr *QChart) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QChart) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QChart) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QChart) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QChart) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QChart) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QChart) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QChart) GetContentsMarginsDefault(left float64, top float64, right float64, bottom float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GetContentsMarginsDefault", left, top, right, bottom})
}

func (ptr *QChart) GrabKeyboardEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GrabKeyboardEventDefault", event})
}

func (ptr *QChart) GrabMouseEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GrabMouseEventDefault", event})
}

func (ptr *QChart) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QChart) HoverLeaveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HoverLeaveEventDefault", event})
}

func (ptr *QChart) HoverMoveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HoverMoveEventDefault", event})
}

func (ptr *QChart) InitStyleOptionDefault(option widgets.QStyleOption_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitStyleOptionDefault", option})
}

func (ptr *QChart) ItemChangeDefault(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemChangeDefault", change, value}).(*core.QVariant)
}

func (ptr *QChart) MoveEventDefault(event widgets.QGraphicsSceneMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QChart) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintDefault", painter, option, widget})
}

func (ptr *QChart) PaintWindowFrameDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintWindowFrameDefault", painter, option, widget})
}

func (ptr *QChart) PolishEventDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PolishEventDefault"})
}

func (ptr *QChart) ResizeEventDefault(event widgets.QGraphicsSceneResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QChart) SceneEventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneEventDefault", event}).(bool)
}

func (ptr *QChart) SetGeometryDefault(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGeometryDefault", rect})
}

func (ptr *QChart) ShapeDefault() *gui.QPainterPath {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShapeDefault"}).(*gui.QPainterPath)
}

func (ptr *QChart) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QChart) SizeHintDefault(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault", which, constraint}).(*core.QSizeF)
}

func (ptr *QChart) TypeDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QChart) UngrabKeyboardEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UngrabKeyboardEventDefault", event})
}

func (ptr *QChart) UngrabMouseEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UngrabMouseEventDefault", event})
}

func (ptr *QChart) UpdateGeometryDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateGeometryDefault"})
}

func (ptr *QChart) WindowFrameEventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WindowFrameEventDefault", event}).(bool)
}

func (ptr *QChart) WindowFrameSectionAtDefault(pos core.QPointF_ITF) core.Qt__WindowFrameSection {

	return core.Qt__WindowFrameSection(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WindowFrameSectionAtDefault", pos}).(float64))
}

func (ptr *QChart) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QChart) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QChart) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QChart) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QChart) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QChart) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QChart) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QChart) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QChart) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func (ptr *QChart) AdvanceDefault(phase int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AdvanceDefault", phase})
}

func (ptr *QChart) CollidesWithItemDefault(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollidesWithItemDefault", other, mode}).(bool)
}

func (ptr *QChart) CollidesWithPathDefault(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollidesWithPathDefault", path, mode}).(bool)
}

func (ptr *QChart) ContainsDefault(point core.QPointF_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContainsDefault", point}).(bool)
}

func (ptr *QChart) ContextMenuEventDefault(event widgets.QGraphicsSceneContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QChart) DragEnterEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QChart) DragLeaveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QChart) DragMoveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QChart) DropEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QChart) HoverEnterEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HoverEnterEventDefault", event})
}

func (ptr *QChart) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QChart) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QChart) IsObscuredByDefault(item widgets.QGraphicsItem_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsObscuredByDefault", item}).(bool)
}

func (ptr *QChart) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QChart) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QChart) MouseDoubleClickEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QChart) MouseMoveEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QChart) MousePressEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QChart) MouseReleaseEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QChart) OpaqueAreaDefault() *gui.QPainterPath {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpaqueAreaDefault"}).(*gui.QPainterPath)
}

func (ptr *QChart) SceneEventFilterDefault(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneEventFilterDefault", watched, event}).(bool)
}

func (ptr *QChart) WheelEventDefault(event widgets.QGraphicsSceneWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
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

func (n *QChartView) InitFromInternal(ptr uintptr, name string) {
	n.QGraphicsView_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QChartView) ClassNameInternalF() string {
	return n.QGraphicsView_PTR().ClassNameInternalF()
}

func NewQChartViewFromPointer(ptr unsafe.Pointer) (n *QChartView) {
	n = new(QChartView)
	n.InitFromInternal(uintptr(ptr), "charts.QChartView")
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

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQChartView", "", parent}).(*QChartView)
}

func NewQChartView2(chart QChart_ITF, parent widgets.QWidget_ITF) *QChartView {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQChartView2", "", chart, parent}).(*QChartView)
}

func (ptr *QChartView) Chart() *QChart {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Chart"}).(*QChart)
}

func (ptr *QChartView) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QChartView) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QChartView) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QChartView) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QChartView) RubberBand() QChartView__RubberBand {

	return QChartView__RubberBand(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RubberBand"}).(float64))
}

func (ptr *QChartView) SetChart(chart QChart_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetChart", chart})
}

func (ptr *QChartView) SetRubberBand(rubberBand QChartView__RubberBand) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRubberBand", rubberBand})
}

func (ptr *QChartView) ConnectDestroyQChartView(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQChartView", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QChartView) DisconnectDestroyQChartView() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQChartView"})
}

func (ptr *QChartView) DestroyQChartView() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQChartView"})
}

func (ptr *QChartView) DestroyQChartViewDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQChartViewDefault"})
}

func (ptr *QChartView) __items_atList(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_atList", i}).(*widgets.QGraphicsItem)
}

func (ptr *QChartView) __items_setList(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_setList", i})
}

func (ptr *QChartView) __items_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_newList"}).(unsafe.Pointer)
}

func (ptr *QChartView) __items_atList2(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_atList2", i}).(*widgets.QGraphicsItem)
}

func (ptr *QChartView) __items_setList2(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_setList2", i})
}

func (ptr *QChartView) __items_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_newList2"}).(unsafe.Pointer)
}

func (ptr *QChartView) __items_atList3(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_atList3", i}).(*widgets.QGraphicsItem)
}

func (ptr *QChartView) __items_setList3(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_setList3", i})
}

func (ptr *QChartView) __items_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_newList3"}).(unsafe.Pointer)
}

func (ptr *QChartView) __items_atList4(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_atList4", i}).(*widgets.QGraphicsItem)
}

func (ptr *QChartView) __items_setList4(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_setList4", i})
}

func (ptr *QChartView) __items_newList4() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_newList4"}).(unsafe.Pointer)
}

func (ptr *QChartView) __items_atList5(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_atList5", i}).(*widgets.QGraphicsItem)
}

func (ptr *QChartView) __items_setList5(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_setList5", i})
}

func (ptr *QChartView) __items_newList5() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_newList5"}).(unsafe.Pointer)
}

func (ptr *QChartView) __items_atList6(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_atList6", i}).(*widgets.QGraphicsItem)
}

func (ptr *QChartView) __items_setList6(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_setList6", i})
}

func (ptr *QChartView) __items_newList6() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_newList6"}).(unsafe.Pointer)
}

func (ptr *QChartView) __items_atList7(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_atList7", i}).(*widgets.QGraphicsItem)
}

func (ptr *QChartView) __items_setList7(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_setList7", i})
}

func (ptr *QChartView) __items_newList7() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_newList7"}).(unsafe.Pointer)
}

func (ptr *QChartView) __updateScene_rects_atList(i int) *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__updateScene_rects_atList", i}).(*core.QRectF)
}

func (ptr *QChartView) __updateScene_rects_setList(i core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__updateScene_rects_setList", i})
}

func (ptr *QChartView) __updateScene_rects_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__updateScene_rects_newList"}).(unsafe.Pointer)
}

func (ptr *QChartView) __scrollBarWidgets_atList(i int) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__scrollBarWidgets_atList", i}).(*widgets.QWidget)
}

func (ptr *QChartView) __scrollBarWidgets_setList(i widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__scrollBarWidgets_setList", i})
}

func (ptr *QChartView) __scrollBarWidgets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__scrollBarWidgets_newList"}).(unsafe.Pointer)
}

func (ptr *QChartView) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QChartView) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QChartView) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QChartView) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QChartView) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QChartView) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QChartView) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QChartView) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QChartView) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QChartView) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QChartView) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QChartView) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QChartView) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QChartView) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QChartView) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QChartView) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QChartView) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QChartView) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QChartView) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QChartView) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QChartView) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QChartView) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QChartView) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QChartView) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QChartView) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QChartView) DrawBackgroundDefault(painter gui.QPainter_ITF, rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawBackgroundDefault", painter, rect})
}

func (ptr *QChartView) DrawForegroundDefault(painter gui.QPainter_ITF, rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawForegroundDefault", painter, rect})
}

func (ptr *QChartView) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QChartView) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QChartView) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QChartView) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QChartView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QChartView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QChartView) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QChartView) InvalidateSceneDefault(rect core.QRectF_ITF, layers widgets.QGraphicsScene__SceneLayer) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InvalidateSceneDefault", rect, layers})
}

func (ptr *QChartView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QChartView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QChartView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QChartView) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QChartView) ScrollContentsByDefault(dx int, dy int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScrollContentsByDefault", dx, dy})
}

func (ptr *QChartView) SetupViewportDefault(widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetupViewportDefault", widget})
}

func (ptr *QChartView) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QChartView) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QChartView) UpdateSceneDefault(rects []*core.QRectF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateSceneDefault", rects})
}

func (ptr *QChartView) UpdateSceneRectDefault(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateSceneRectDefault", rect})
}

func (ptr *QChartView) ViewportEventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewportEventDefault", event}).(bool)
}

func (ptr *QChartView) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QChartView) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QChartView) ViewportSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewportSizeHintDefault"}).(*core.QSize)
}

func (ptr *QChartView) ChangeEventDefault(ev core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", ev})
}

func (ptr *QChartView) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QChartView) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QChartView) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QChartView) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QChartView) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QChartView) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QChartView) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QChartView) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QChartView) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QChartView) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QChartView) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QChartView) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QChartView) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QChartView) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QChartView) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QChartView) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QChartView) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QChartView) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QChartView) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QChartView) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QChartView) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QChartView) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QChartView) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QChartView) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QChartView) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QChartView) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QChartView) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QChartView) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QChartView) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QChartView) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QChartView) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QChartView) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QChartView) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QChartView) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QChartView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QChartView) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QChartView) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QChartView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QChartView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QChartView) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QChartView) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QDateTimeAxis) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractAxis_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDateTimeAxis) ClassNameInternalF() string {
	return n.QAbstractAxis_PTR().ClassNameInternalF()
}

func NewQDateTimeAxisFromPointer(ptr unsafe.Pointer) (n *QDateTimeAxis) {
	n = new(QDateTimeAxis)
	n.InitFromInternal(uintptr(ptr), "charts.QDateTimeAxis")
	return
}
func NewQDateTimeAxis(parent core.QObject_ITF) *QDateTimeAxis {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQDateTimeAxis", "", parent}).(*QDateTimeAxis)
}

func (ptr *QDateTimeAxis) Format() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Format"}).(string)
}

func (ptr *QDateTimeAxis) ConnectFormatChanged(f func(format string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFormatChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDateTimeAxis) DisconnectFormatChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFormatChanged"})
}

func (ptr *QDateTimeAxis) FormatChanged(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormatChanged", format})
}

func (ptr *QDateTimeAxis) Max() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Max"}).(*core.QDateTime)
}

func (ptr *QDateTimeAxis) ConnectMaxChanged(f func(max *core.QDateTime)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMaxChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDateTimeAxis) DisconnectMaxChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMaxChanged"})
}

func (ptr *QDateTimeAxis) MaxChanged(max core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxChanged", max})
}

func (ptr *QDateTimeAxis) Min() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Min"}).(*core.QDateTime)
}

func (ptr *QDateTimeAxis) ConnectMinChanged(f func(min *core.QDateTime)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDateTimeAxis) DisconnectMinChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinChanged"})
}

func (ptr *QDateTimeAxis) MinChanged(min core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinChanged", min})
}

func (ptr *QDateTimeAxis) ConnectRangeChanged(f func(min *core.QDateTime, max *core.QDateTime)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRangeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDateTimeAxis) DisconnectRangeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRangeChanged"})
}

func (ptr *QDateTimeAxis) RangeChanged(min core.QDateTime_ITF, max core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RangeChanged", min, max})
}

func (ptr *QDateTimeAxis) SetFormat(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFormat", format})
}

func (ptr *QDateTimeAxis) SetMax(max core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMax", max})
}

func (ptr *QDateTimeAxis) SetMin(min core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMin", min})
}

func (ptr *QDateTimeAxis) SetRange(min core.QDateTime_ITF, max core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRange", min, max})
}

func (ptr *QDateTimeAxis) SetTickCount(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTickCount", count})
}

func (ptr *QDateTimeAxis) TickCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickCount"}).(float64))
}

func (ptr *QDateTimeAxis) ConnectTickCountChanged(f func(tickCount int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTickCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDateTimeAxis) DisconnectTickCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTickCountChanged"})
}

func (ptr *QDateTimeAxis) TickCountChanged(tickCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickCountChanged", tickCount})
}

func (ptr *QDateTimeAxis) ConnectType(f func() QAbstractAxis__AxisType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDateTimeAxis) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QDateTimeAxis) Type() QAbstractAxis__AxisType {

	return QAbstractAxis__AxisType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QDateTimeAxis) TypeDefault() QAbstractAxis__AxisType {

	return QAbstractAxis__AxisType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QDateTimeAxis) ConnectDestroyQDateTimeAxis(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDateTimeAxis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDateTimeAxis) DisconnectDestroyQDateTimeAxis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDateTimeAxis"})
}

func (ptr *QDateTimeAxis) DestroyQDateTimeAxis() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDateTimeAxis"})
}

func (ptr *QDateTimeAxis) DestroyQDateTimeAxisDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDateTimeAxisDefault"})
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

func (n *QHBarModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHBarModelMapper) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQHBarModelMapperFromPointer(ptr unsafe.Pointer) (n *QHBarModelMapper) {
	n = new(QHBarModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QHBarModelMapper")
	return
}

func (ptr *QHBarModelMapper) DestroyQHBarModelMapper() {
}

func NewQHBarModelMapper(parent core.QObject_ITF) *QHBarModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQHBarModelMapper", "", parent}).(*QHBarModelMapper)
}

func (ptr *QHBarModelMapper) ColumnCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCount"}).(float64))
}

func (ptr *QHBarModelMapper) ConnectColumnCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBarModelMapper) DisconnectColumnCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnCountChanged"})
}

func (ptr *QHBarModelMapper) ColumnCountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountChanged"})
}

func (ptr *QHBarModelMapper) FirstBarSetRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstBarSetRow"}).(float64))
}

func (ptr *QHBarModelMapper) ConnectFirstBarSetRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstBarSetRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBarModelMapper) DisconnectFirstBarSetRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstBarSetRowChanged"})
}

func (ptr *QHBarModelMapper) FirstBarSetRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstBarSetRowChanged"})
}

func (ptr *QHBarModelMapper) FirstColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstColumn"}).(float64))
}

func (ptr *QHBarModelMapper) ConnectFirstColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBarModelMapper) DisconnectFirstColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstColumnChanged"})
}

func (ptr *QHBarModelMapper) FirstColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstColumnChanged"})
}

func (ptr *QHBarModelMapper) LastBarSetRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastBarSetRow"}).(float64))
}

func (ptr *QHBarModelMapper) ConnectLastBarSetRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLastBarSetRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBarModelMapper) DisconnectLastBarSetRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLastBarSetRowChanged"})
}

func (ptr *QHBarModelMapper) LastBarSetRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastBarSetRowChanged"})
}

func (ptr *QHBarModelMapper) Model() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*core.QAbstractItemModel)
}

func (ptr *QHBarModelMapper) ConnectModelReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModelReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBarModelMapper) DisconnectModelReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModelReplaced"})
}

func (ptr *QHBarModelMapper) ModelReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModelReplaced"})
}

func (ptr *QHBarModelMapper) Series() *QAbstractBarSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QAbstractBarSeries)
}

func (ptr *QHBarModelMapper) ConnectSeriesReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBarModelMapper) DisconnectSeriesReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesReplaced"})
}

func (ptr *QHBarModelMapper) SeriesReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesReplaced"})
}

func (ptr *QHBarModelMapper) SetColumnCount(columnCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnCount", columnCount})
}

func (ptr *QHBarModelMapper) SetFirstBarSetRow(firstBarSetRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstBarSetRow", firstBarSetRow})
}

func (ptr *QHBarModelMapper) SetFirstColumn(firstColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstColumn", firstColumn})
}

func (ptr *QHBarModelMapper) SetLastBarSetRow(lastBarSetRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastBarSetRow", lastBarSetRow})
}

func (ptr *QHBarModelMapper) SetModel(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModel", model})
}

func (ptr *QHBarModelMapper) SetSeries(series QAbstractBarSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSeries", series})
}

func (ptr *QHBarModelMapper) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHBarModelMapper) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHBarModelMapper) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHBarModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHBarModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHBarModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHBarModelMapper) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHBarModelMapper) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHBarModelMapper) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHBarModelMapper) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHBarModelMapper) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHBarModelMapper) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHBarModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHBarModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHBarModelMapper) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHBarModelMapper) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHBarModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHBarModelMapper) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QHBarModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QHBarModelMapper) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QHBarModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QHBoxPlotModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHBoxPlotModelMapper) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQHBoxPlotModelMapperFromPointer(ptr unsafe.Pointer) (n *QHBoxPlotModelMapper) {
	n = new(QHBoxPlotModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QHBoxPlotModelMapper")
	return
}

func (ptr *QHBoxPlotModelMapper) DestroyQHBoxPlotModelMapper() {
}

func NewQHBoxPlotModelMapper(parent core.QObject_ITF) *QHBoxPlotModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQHBoxPlotModelMapper", "", parent}).(*QHBoxPlotModelMapper)
}

func (ptr *QHBoxPlotModelMapper) ColumnCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCount"}).(float64))
}

func (ptr *QHBoxPlotModelMapper) ConnectColumnCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBoxPlotModelMapper) DisconnectColumnCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnCountChanged"})
}

func (ptr *QHBoxPlotModelMapper) ColumnCountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountChanged"})
}

func (ptr *QHBoxPlotModelMapper) FirstBoxSetRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstBoxSetRow"}).(float64))
}

func (ptr *QHBoxPlotModelMapper) ConnectFirstBoxSetRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstBoxSetRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBoxPlotModelMapper) DisconnectFirstBoxSetRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstBoxSetRowChanged"})
}

func (ptr *QHBoxPlotModelMapper) FirstBoxSetRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstBoxSetRowChanged"})
}

func (ptr *QHBoxPlotModelMapper) FirstColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstColumn"}).(float64))
}

func (ptr *QHBoxPlotModelMapper) ConnectFirstColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBoxPlotModelMapper) DisconnectFirstColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstColumnChanged"})
}

func (ptr *QHBoxPlotModelMapper) FirstColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstColumnChanged"})
}

func (ptr *QHBoxPlotModelMapper) LastBoxSetRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastBoxSetRow"}).(float64))
}

func (ptr *QHBoxPlotModelMapper) ConnectLastBoxSetRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLastBoxSetRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBoxPlotModelMapper) DisconnectLastBoxSetRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLastBoxSetRowChanged"})
}

func (ptr *QHBoxPlotModelMapper) LastBoxSetRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastBoxSetRowChanged"})
}

func (ptr *QHBoxPlotModelMapper) Model() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*core.QAbstractItemModel)
}

func (ptr *QHBoxPlotModelMapper) ConnectModelReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModelReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBoxPlotModelMapper) DisconnectModelReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModelReplaced"})
}

func (ptr *QHBoxPlotModelMapper) ModelReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModelReplaced"})
}

func (ptr *QHBoxPlotModelMapper) Series() *QBoxPlotSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QBoxPlotSeries)
}

func (ptr *QHBoxPlotModelMapper) ConnectSeriesReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHBoxPlotModelMapper) DisconnectSeriesReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesReplaced"})
}

func (ptr *QHBoxPlotModelMapper) SeriesReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesReplaced"})
}

func (ptr *QHBoxPlotModelMapper) SetColumnCount(rowCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnCount", rowCount})
}

func (ptr *QHBoxPlotModelMapper) SetFirstBoxSetRow(firstBoxSetRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstBoxSetRow", firstBoxSetRow})
}

func (ptr *QHBoxPlotModelMapper) SetFirstColumn(firstColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstColumn", firstColumn})
}

func (ptr *QHBoxPlotModelMapper) SetLastBoxSetRow(lastBoxSetRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastBoxSetRow", lastBoxSetRow})
}

func (ptr *QHBoxPlotModelMapper) SetModel(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModel", model})
}

func (ptr *QHBoxPlotModelMapper) SetSeries(series QBoxPlotSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSeries", series})
}

func (ptr *QHBoxPlotModelMapper) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHBoxPlotModelMapper) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHBoxPlotModelMapper) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHBoxPlotModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHBoxPlotModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHBoxPlotModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHBoxPlotModelMapper) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHBoxPlotModelMapper) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHBoxPlotModelMapper) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHBoxPlotModelMapper) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHBoxPlotModelMapper) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHBoxPlotModelMapper) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHBoxPlotModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHBoxPlotModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHBoxPlotModelMapper) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHBoxPlotModelMapper) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHBoxPlotModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHBoxPlotModelMapper) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QHBoxPlotModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QHBoxPlotModelMapper) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QHBoxPlotModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QHCandlestickModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QCandlestickModelMapper_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHCandlestickModelMapper) ClassNameInternalF() string {
	return n.QCandlestickModelMapper_PTR().ClassNameInternalF()
}

func NewQHCandlestickModelMapperFromPointer(ptr unsafe.Pointer) (n *QHCandlestickModelMapper) {
	n = new(QHCandlestickModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QHCandlestickModelMapper")
	return
}

func (ptr *QHCandlestickModelMapper) DestroyQHCandlestickModelMapper() {
}

func NewQHCandlestickModelMapper(parent core.QObject_ITF) *QHCandlestickModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQHCandlestickModelMapper", "", parent}).(*QHCandlestickModelMapper)
}

func (ptr *QHCandlestickModelMapper) CloseColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseColumn"}).(float64))
}

func (ptr *QHCandlestickModelMapper) ConnectCloseColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCloseColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHCandlestickModelMapper) DisconnectCloseColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCloseColumnChanged"})
}

func (ptr *QHCandlestickModelMapper) CloseColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseColumnChanged"})
}

func (ptr *QHCandlestickModelMapper) FirstSetRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstSetRow"}).(float64))
}

func (ptr *QHCandlestickModelMapper) ConnectFirstSetRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstSetRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHCandlestickModelMapper) DisconnectFirstSetRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstSetRowChanged"})
}

func (ptr *QHCandlestickModelMapper) FirstSetRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstSetRowChanged"})
}

func (ptr *QHCandlestickModelMapper) HighColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HighColumn"}).(float64))
}

func (ptr *QHCandlestickModelMapper) ConnectHighColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHighColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHCandlestickModelMapper) DisconnectHighColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHighColumnChanged"})
}

func (ptr *QHCandlestickModelMapper) HighColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HighColumnChanged"})
}

func (ptr *QHCandlestickModelMapper) LastSetRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastSetRow"}).(float64))
}

func (ptr *QHCandlestickModelMapper) ConnectLastSetRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLastSetRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHCandlestickModelMapper) DisconnectLastSetRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLastSetRowChanged"})
}

func (ptr *QHCandlestickModelMapper) LastSetRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastSetRowChanged"})
}

func (ptr *QHCandlestickModelMapper) LowColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowColumn"}).(float64))
}

func (ptr *QHCandlestickModelMapper) ConnectLowColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLowColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHCandlestickModelMapper) DisconnectLowColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLowColumnChanged"})
}

func (ptr *QHCandlestickModelMapper) LowColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowColumnChanged"})
}

func (ptr *QHCandlestickModelMapper) OpenColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenColumn"}).(float64))
}

func (ptr *QHCandlestickModelMapper) ConnectOpenColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpenColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHCandlestickModelMapper) DisconnectOpenColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpenColumnChanged"})
}

func (ptr *QHCandlestickModelMapper) OpenColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenColumnChanged"})
}

func (ptr *QHCandlestickModelMapper) ConnectOrientation(f func() core.Qt__Orientation) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOrientation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHCandlestickModelMapper) DisconnectOrientation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOrientation"})
}

func (ptr *QHCandlestickModelMapper) Orientation() core.Qt__Orientation {

	return core.Qt__Orientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Orientation"}).(float64))
}

func (ptr *QHCandlestickModelMapper) OrientationDefault() core.Qt__Orientation {

	return core.Qt__Orientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OrientationDefault"}).(float64))
}

func (ptr *QHCandlestickModelMapper) SetCloseColumn(closeColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCloseColumn", closeColumn})
}

func (ptr *QHCandlestickModelMapper) SetFirstSetRow(firstSetRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstSetRow", firstSetRow})
}

func (ptr *QHCandlestickModelMapper) SetHighColumn(highColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHighColumn", highColumn})
}

func (ptr *QHCandlestickModelMapper) SetLastSetRow(lastSetRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastSetRow", lastSetRow})
}

func (ptr *QHCandlestickModelMapper) SetLowColumn(lowColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLowColumn", lowColumn})
}

func (ptr *QHCandlestickModelMapper) SetOpenColumn(openColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpenColumn", openColumn})
}

func (ptr *QHCandlestickModelMapper) SetTimestampColumn(timestampColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTimestampColumn", timestampColumn})
}

func (ptr *QHCandlestickModelMapper) TimestampColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimestampColumn"}).(float64))
}

func (ptr *QHCandlestickModelMapper) ConnectTimestampColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTimestampColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHCandlestickModelMapper) DisconnectTimestampColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTimestampColumnChanged"})
}

func (ptr *QHCandlestickModelMapper) TimestampColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimestampColumnChanged"})
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

func (n *QHPieModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QPieModelMapper_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHPieModelMapper) ClassNameInternalF() string {
	return n.QPieModelMapper_PTR().ClassNameInternalF()
}

func NewQHPieModelMapperFromPointer(ptr unsafe.Pointer) (n *QHPieModelMapper) {
	n = new(QHPieModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QHPieModelMapper")
	return
}

func (ptr *QHPieModelMapper) DestroyQHPieModelMapper() {
}

func NewQHPieModelMapper(parent core.QObject_ITF) *QHPieModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQHPieModelMapper", "", parent}).(*QHPieModelMapper)
}

func (ptr *QHPieModelMapper) ColumnCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCount"}).(float64))
}

func (ptr *QHPieModelMapper) ConnectColumnCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHPieModelMapper) DisconnectColumnCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnCountChanged"})
}

func (ptr *QHPieModelMapper) ColumnCountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountChanged"})
}

func (ptr *QHPieModelMapper) FirstColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstColumn"}).(float64))
}

func (ptr *QHPieModelMapper) ConnectFirstColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHPieModelMapper) DisconnectFirstColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstColumnChanged"})
}

func (ptr *QHPieModelMapper) FirstColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstColumnChanged"})
}

func (ptr *QHPieModelMapper) LabelsRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsRow"}).(float64))
}

func (ptr *QHPieModelMapper) ConnectLabelsRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHPieModelMapper) DisconnectLabelsRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsRowChanged"})
}

func (ptr *QHPieModelMapper) LabelsRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsRowChanged"})
}

func (ptr *QHPieModelMapper) Model() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*core.QAbstractItemModel)
}

func (ptr *QHPieModelMapper) ConnectModelReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModelReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHPieModelMapper) DisconnectModelReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModelReplaced"})
}

func (ptr *QHPieModelMapper) ModelReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModelReplaced"})
}

func (ptr *QHPieModelMapper) Series() *QPieSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QPieSeries)
}

func (ptr *QHPieModelMapper) ConnectSeriesReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHPieModelMapper) DisconnectSeriesReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesReplaced"})
}

func (ptr *QHPieModelMapper) SeriesReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesReplaced"})
}

func (ptr *QHPieModelMapper) SetColumnCount(columnCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnCount", columnCount})
}

func (ptr *QHPieModelMapper) SetFirstColumn(firstColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstColumn", firstColumn})
}

func (ptr *QHPieModelMapper) SetLabelsRow(labelsRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsRow", labelsRow})
}

func (ptr *QHPieModelMapper) SetModel(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModel", model})
}

func (ptr *QHPieModelMapper) SetSeries(series QPieSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSeries", series})
}

func (ptr *QHPieModelMapper) SetValuesRow(valuesRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValuesRow", valuesRow})
}

func (ptr *QHPieModelMapper) ValuesRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValuesRow"}).(float64))
}

func (ptr *QHPieModelMapper) ConnectValuesRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValuesRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHPieModelMapper) DisconnectValuesRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValuesRowChanged"})
}

func (ptr *QHPieModelMapper) ValuesRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValuesRowChanged"})
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

func (n *QHXYModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QXYModelMapper_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHXYModelMapper) ClassNameInternalF() string {
	return n.QXYModelMapper_PTR().ClassNameInternalF()
}

func NewQHXYModelMapperFromPointer(ptr unsafe.Pointer) (n *QHXYModelMapper) {
	n = new(QHXYModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QHXYModelMapper")
	return
}

func (ptr *QHXYModelMapper) DestroyQHXYModelMapper() {
}

func NewQHXYModelMapper(parent core.QObject_ITF) *QHXYModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQHXYModelMapper", "", parent}).(*QHXYModelMapper)
}

func (ptr *QHXYModelMapper) ColumnCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCount"}).(float64))
}

func (ptr *QHXYModelMapper) ConnectColumnCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHXYModelMapper) DisconnectColumnCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnCountChanged"})
}

func (ptr *QHXYModelMapper) ColumnCountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountChanged"})
}

func (ptr *QHXYModelMapper) FirstColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstColumn"}).(float64))
}

func (ptr *QHXYModelMapper) ConnectFirstColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHXYModelMapper) DisconnectFirstColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstColumnChanged"})
}

func (ptr *QHXYModelMapper) FirstColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstColumnChanged"})
}

func (ptr *QHXYModelMapper) Model() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*core.QAbstractItemModel)
}

func (ptr *QHXYModelMapper) ConnectModelReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModelReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHXYModelMapper) DisconnectModelReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModelReplaced"})
}

func (ptr *QHXYModelMapper) ModelReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModelReplaced"})
}

func (ptr *QHXYModelMapper) Series() *QXYSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QXYSeries)
}

func (ptr *QHXYModelMapper) ConnectSeriesReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHXYModelMapper) DisconnectSeriesReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesReplaced"})
}

func (ptr *QHXYModelMapper) SeriesReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesReplaced"})
}

func (ptr *QHXYModelMapper) SetColumnCount(columnCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnCount", columnCount})
}

func (ptr *QHXYModelMapper) SetFirstColumn(firstColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstColumn", firstColumn})
}

func (ptr *QHXYModelMapper) SetModel(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModel", model})
}

func (ptr *QHXYModelMapper) SetSeries(series QXYSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSeries", series})
}

func (ptr *QHXYModelMapper) SetXRow(xRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXRow", xRow})
}

func (ptr *QHXYModelMapper) SetYRow(yRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYRow", yRow})
}

func (ptr *QHXYModelMapper) XRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XRow"}).(float64))
}

func (ptr *QHXYModelMapper) ConnectXRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectXRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHXYModelMapper) DisconnectXRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectXRowChanged"})
}

func (ptr *QHXYModelMapper) XRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XRowChanged"})
}

func (ptr *QHXYModelMapper) YRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YRow"}).(float64))
}

func (ptr *QHXYModelMapper) ConnectYRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectYRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHXYModelMapper) DisconnectYRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectYRowChanged"})
}

func (ptr *QHXYModelMapper) YRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YRowChanged"})
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

func (n *QHorizontalBarSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractBarSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHorizontalBarSeries) ClassNameInternalF() string {
	return n.QAbstractBarSeries_PTR().ClassNameInternalF()
}

func NewQHorizontalBarSeriesFromPointer(ptr unsafe.Pointer) (n *QHorizontalBarSeries) {
	n = new(QHorizontalBarSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QHorizontalBarSeries")
	return
}
func NewQHorizontalBarSeries(parent core.QObject_ITF) *QHorizontalBarSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQHorizontalBarSeries", "", parent}).(*QHorizontalBarSeries)
}

func (ptr *QHorizontalBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHorizontalBarSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QHorizontalBarSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QHorizontalBarSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QHorizontalBarSeries) ConnectDestroyQHorizontalBarSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHorizontalBarSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHorizontalBarSeries) DisconnectDestroyQHorizontalBarSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHorizontalBarSeries"})
}

func (ptr *QHorizontalBarSeries) DestroyQHorizontalBarSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHorizontalBarSeries"})
}

func (ptr *QHorizontalBarSeries) DestroyQHorizontalBarSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHorizontalBarSeriesDefault"})
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

func (n *QHorizontalPercentBarSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractBarSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHorizontalPercentBarSeries) ClassNameInternalF() string {
	return n.QAbstractBarSeries_PTR().ClassNameInternalF()
}

func NewQHorizontalPercentBarSeriesFromPointer(ptr unsafe.Pointer) (n *QHorizontalPercentBarSeries) {
	n = new(QHorizontalPercentBarSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QHorizontalPercentBarSeries")
	return
}
func NewQHorizontalPercentBarSeries(parent core.QObject_ITF) *QHorizontalPercentBarSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQHorizontalPercentBarSeries", "", parent}).(*QHorizontalPercentBarSeries)
}

func (ptr *QHorizontalPercentBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHorizontalPercentBarSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QHorizontalPercentBarSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QHorizontalPercentBarSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QHorizontalPercentBarSeries) ConnectDestroyQHorizontalPercentBarSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHorizontalPercentBarSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHorizontalPercentBarSeries) DisconnectDestroyQHorizontalPercentBarSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHorizontalPercentBarSeries"})
}

func (ptr *QHorizontalPercentBarSeries) DestroyQHorizontalPercentBarSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHorizontalPercentBarSeries"})
}

func (ptr *QHorizontalPercentBarSeries) DestroyQHorizontalPercentBarSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHorizontalPercentBarSeriesDefault"})
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

func (n *QHorizontalStackedBarSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractBarSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHorizontalStackedBarSeries) ClassNameInternalF() string {
	return n.QAbstractBarSeries_PTR().ClassNameInternalF()
}

func NewQHorizontalStackedBarSeriesFromPointer(ptr unsafe.Pointer) (n *QHorizontalStackedBarSeries) {
	n = new(QHorizontalStackedBarSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QHorizontalStackedBarSeries")
	return
}
func NewQHorizontalStackedBarSeries(parent core.QObject_ITF) *QHorizontalStackedBarSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQHorizontalStackedBarSeries", "", parent}).(*QHorizontalStackedBarSeries)
}

func (ptr *QHorizontalStackedBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHorizontalStackedBarSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QHorizontalStackedBarSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QHorizontalStackedBarSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QHorizontalStackedBarSeries) ConnectDestroyQHorizontalStackedBarSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHorizontalStackedBarSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHorizontalStackedBarSeries) DisconnectDestroyQHorizontalStackedBarSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHorizontalStackedBarSeries"})
}

func (ptr *QHorizontalStackedBarSeries) DestroyQHorizontalStackedBarSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHorizontalStackedBarSeries"})
}

func (ptr *QHorizontalStackedBarSeries) DestroyQHorizontalStackedBarSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHorizontalStackedBarSeriesDefault"})
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

func (n *QLegend) InitFromInternal(ptr uintptr, name string) {
	n.QGraphicsWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLegend) ClassNameInternalF() string {
	return n.QGraphicsWidget_PTR().ClassNameInternalF()
}

func NewQLegendFromPointer(ptr unsafe.Pointer) (n *QLegend) {
	n = new(QLegend)
	n.InitFromInternal(uintptr(ptr), "charts.QLegend")
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

	return core.Qt__AlignmentFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Alignment"}).(float64))
}

func (ptr *QLegend) AttachToChart() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AttachToChart"})
}

func (ptr *QLegend) ConnectBackgroundVisibleChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBackgroundVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegend) DisconnectBackgroundVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBackgroundVisibleChanged"})
}

func (ptr *QLegend) BackgroundVisibleChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundVisibleChanged", visible})
}

func (ptr *QLegend) BorderColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderColor"}).(*gui.QColor)
}

func (ptr *QLegend) ConnectBorderColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBorderColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegend) DisconnectBorderColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBorderColorChanged"})
}

func (ptr *QLegend) BorderColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderColorChanged", color})
}

func (ptr *QLegend) Brush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Brush"}).(*gui.QBrush)
}

func (ptr *QLegend) Color() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Color"}).(*gui.QColor)
}

func (ptr *QLegend) ConnectColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegend) DisconnectColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColorChanged"})
}

func (ptr *QLegend) ColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorChanged", color})
}

func (ptr *QLegend) DetachFromChart() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DetachFromChart"})
}

func (ptr *QLegend) ConnectFontChanged(f func(font *gui.QFont)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFontChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegend) DisconnectFontChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFontChanged"})
}

func (ptr *QLegend) FontChanged(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FontChanged", font})
}

func (ptr *QLegend) IsAttachedToChart() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAttachedToChart"}).(bool)
}

func (ptr *QLegend) IsBackgroundVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBackgroundVisible"}).(bool)
}

func (ptr *QLegend) LabelBrush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBrush"}).(*gui.QBrush)
}

func (ptr *QLegend) LabelColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelColor"}).(*gui.QColor)
}

func (ptr *QLegend) ConnectLabelColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegend) DisconnectLabelColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelColorChanged"})
}

func (ptr *QLegend) LabelColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelColorChanged", color})
}

func (ptr *QLegend) MarkerShape() QLegend__MarkerShape {

	return QLegend__MarkerShape(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MarkerShape"}).(float64))
}

func (ptr *QLegend) ConnectMarkerShapeChanged(f func(shape QLegend__MarkerShape)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMarkerShapeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegend) DisconnectMarkerShapeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMarkerShapeChanged"})
}

func (ptr *QLegend) MarkerShapeChanged(shape QLegend__MarkerShape) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MarkerShapeChanged", shape})
}

func (ptr *QLegend) Markers(series QAbstractSeries_ITF) []*QLegendMarker {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Markers", series}).([]*QLegendMarker)
}

func (ptr *QLegend) Pen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pen"}).(*gui.QPen)
}

func (ptr *QLegend) ReverseMarkers() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReverseMarkers"}).(bool)
}

func (ptr *QLegend) ConnectReverseMarkersChanged(f func(reverseMarkers bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReverseMarkersChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegend) DisconnectReverseMarkersChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReverseMarkersChanged"})
}

func (ptr *QLegend) ReverseMarkersChanged(reverseMarkers bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReverseMarkersChanged", reverseMarkers})
}

func (ptr *QLegend) SetAlignment(alignment core.Qt__AlignmentFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAlignment", alignment})
}

func (ptr *QLegend) SetBackgroundVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackgroundVisible", visible})
}

func (ptr *QLegend) SetBorderColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBorderColor", color})
}

func (ptr *QLegend) SetBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrush", brush})
}

func (ptr *QLegend) SetColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColor", color})
}

func (ptr *QLegend) SetLabelBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelBrush", brush})
}

func (ptr *QLegend) SetLabelColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelColor", color})
}

func (ptr *QLegend) SetMarkerShape(shape QLegend__MarkerShape) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMarkerShape", shape})
}

func (ptr *QLegend) SetPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPen", pen})
}

func (ptr *QLegend) SetReverseMarkers(reverseMarkers bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReverseMarkers", reverseMarkers})
}

func (ptr *QLegend) SetShowToolTips(show bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShowToolTips", show})
}

func (ptr *QLegend) ShowToolTips() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowToolTips"}).(bool)
}

func (ptr *QLegend) ConnectShowToolTipsChanged(f func(showToolTips bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShowToolTipsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegend) DisconnectShowToolTipsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShowToolTipsChanged"})
}

func (ptr *QLegend) ShowToolTipsChanged(showToolTips bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowToolTipsChanged", showToolTips})
}

func (ptr *QLegend) ConnectDestroyQLegend(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLegend", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegend) DisconnectDestroyQLegend() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLegend"})
}

func (ptr *QLegend) DestroyQLegend() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLegend"})
}

func (ptr *QLegend) DestroyQLegendDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLegendDefault"})
}

func (ptr *QLegend) __markers_atList(i int) *QLegendMarker {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__markers_atList", i}).(*QLegendMarker)
}

func (ptr *QLegend) __markers_setList(i QLegendMarker_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__markers_setList", i})
}

func (ptr *QLegend) __markers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__markers_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QLegend) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QLegend) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QLegend) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QLegend) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QLegend) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QLegend) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QLegend) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QLegend) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QLegend) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QLegend) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QLegend) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QLegend) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QLegend) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QLegend) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QLegend) __childItems_atList(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__childItems_atList", i}).(*widgets.QGraphicsItem)
}

func (ptr *QLegend) __childItems_setList(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__childItems_setList", i})
}

func (ptr *QLegend) __childItems_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__childItems_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) __collidingItems_atList(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__collidingItems_atList", i}).(*widgets.QGraphicsItem)
}

func (ptr *QLegend) __collidingItems_setList(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__collidingItems_setList", i})
}

func (ptr *QLegend) __collidingItems_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__collidingItems_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) __setTransformations_transformations_atList(i int) *widgets.QGraphicsTransform {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTransformations_transformations_atList", i}).(*widgets.QGraphicsTransform)
}

func (ptr *QLegend) __setTransformations_transformations_setList(i widgets.QGraphicsTransform_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTransformations_transformations_setList", i})
}

func (ptr *QLegend) __setTransformations_transformations_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTransformations_transformations_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) __transformations_atList(i int) *widgets.QGraphicsTransform {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__transformations_atList", i}).(*widgets.QGraphicsTransform)
}

func (ptr *QLegend) __transformations_setList(i widgets.QGraphicsTransform_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__transformations_setList", i})
}

func (ptr *QLegend) __transformations_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__transformations_newList"}).(unsafe.Pointer)
}

func (ptr *QLegend) BoundingRectDefault() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundingRectDefault"}).(*core.QRectF)
}

func (ptr *QLegend) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QLegend) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QLegend) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QLegend) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QLegend) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QLegend) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QLegend) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QLegend) GetContentsMarginsDefault(left float64, top float64, right float64, bottom float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GetContentsMarginsDefault", left, top, right, bottom})
}

func (ptr *QLegend) GrabKeyboardEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GrabKeyboardEventDefault", event})
}

func (ptr *QLegend) GrabMouseEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GrabMouseEventDefault", event})
}

func (ptr *QLegend) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QLegend) HoverLeaveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HoverLeaveEventDefault", event})
}

func (ptr *QLegend) HoverMoveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HoverMoveEventDefault", event})
}

func (ptr *QLegend) InitStyleOptionDefault(option widgets.QStyleOption_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitStyleOptionDefault", option})
}

func (ptr *QLegend) ItemChangeDefault(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemChangeDefault", change, value}).(*core.QVariant)
}

func (ptr *QLegend) MoveEventDefault(event widgets.QGraphicsSceneMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QLegend) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintDefault", painter, option, widget})
}

func (ptr *QLegend) PaintWindowFrameDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintWindowFrameDefault", painter, option, widget})
}

func (ptr *QLegend) PolishEventDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PolishEventDefault"})
}

func (ptr *QLegend) ResizeEventDefault(event widgets.QGraphicsSceneResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QLegend) SceneEventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneEventDefault", event}).(bool)
}

func (ptr *QLegend) SetGeometryDefault(rect core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGeometryDefault", rect})
}

func (ptr *QLegend) ShapeDefault() *gui.QPainterPath {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShapeDefault"}).(*gui.QPainterPath)
}

func (ptr *QLegend) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QLegend) SizeHintDefault(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault", which, constraint}).(*core.QSizeF)
}

func (ptr *QLegend) TypeDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QLegend) UngrabKeyboardEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UngrabKeyboardEventDefault", event})
}

func (ptr *QLegend) UngrabMouseEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UngrabMouseEventDefault", event})
}

func (ptr *QLegend) UpdateGeometryDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateGeometryDefault"})
}

func (ptr *QLegend) WindowFrameEventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WindowFrameEventDefault", event}).(bool)
}

func (ptr *QLegend) WindowFrameSectionAtDefault(pos core.QPointF_ITF) core.Qt__WindowFrameSection {

	return core.Qt__WindowFrameSection(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WindowFrameSectionAtDefault", pos}).(float64))
}

func (ptr *QLegend) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QLegend) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QLegend) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QLegend) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QLegend) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QLegend) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QLegend) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QLegend) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QLegend) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func (ptr *QLegend) AdvanceDefault(phase int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AdvanceDefault", phase})
}

func (ptr *QLegend) CollidesWithItemDefault(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollidesWithItemDefault", other, mode}).(bool)
}

func (ptr *QLegend) CollidesWithPathDefault(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollidesWithPathDefault", path, mode}).(bool)
}

func (ptr *QLegend) ContainsDefault(point core.QPointF_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContainsDefault", point}).(bool)
}

func (ptr *QLegend) ContextMenuEventDefault(event widgets.QGraphicsSceneContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QLegend) DragEnterEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QLegend) DragLeaveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QLegend) DragMoveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QLegend) DropEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QLegend) HoverEnterEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HoverEnterEventDefault", event})
}

func (ptr *QLegend) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QLegend) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QLegend) IsObscuredByDefault(item widgets.QGraphicsItem_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsObscuredByDefault", item}).(bool)
}

func (ptr *QLegend) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QLegend) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QLegend) MouseDoubleClickEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QLegend) MouseMoveEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QLegend) MousePressEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QLegend) MouseReleaseEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QLegend) OpaqueAreaDefault() *gui.QPainterPath {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpaqueAreaDefault"}).(*gui.QPainterPath)
}

func (ptr *QLegend) SceneEventFilterDefault(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneEventFilterDefault", watched, event}).(bool)
}

func (ptr *QLegend) WheelEventDefault(event widgets.QGraphicsSceneWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
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

func (n *QLegendMarker) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLegendMarker) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QLegendMarker) {
	n = new(QLegendMarker)
	n.InitFromInternal(uintptr(ptr), "charts.QLegendMarker")
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

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Brush"}).(*gui.QBrush)
}

func (ptr *QLegendMarker) ConnectBrushChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBrushChanged"})
}

func (ptr *QLegendMarker) BrushChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BrushChanged"})
}

func (ptr *QLegendMarker) ConnectClicked(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QLegendMarker) Clicked() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked"})
}

func (ptr *QLegendMarker) Font() *gui.QFont {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Font"}).(*gui.QFont)
}

func (ptr *QLegendMarker) ConnectFontChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFontChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectFontChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFontChanged"})
}

func (ptr *QLegendMarker) FontChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FontChanged"})
}

func (ptr *QLegendMarker) ConnectHovered(f func(status bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QLegendMarker) Hovered(status bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", status})
}

func (ptr *QLegendMarker) IsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsVisible"}).(bool)
}

func (ptr *QLegendMarker) Label() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Label"}).(string)
}

func (ptr *QLegendMarker) LabelBrush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBrush"}).(*gui.QBrush)
}

func (ptr *QLegendMarker) ConnectLabelBrushChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectLabelBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelBrushChanged"})
}

func (ptr *QLegendMarker) LabelBrushChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBrushChanged"})
}

func (ptr *QLegendMarker) ConnectLabelChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectLabelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelChanged"})
}

func (ptr *QLegendMarker) LabelChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelChanged"})
}

func (ptr *QLegendMarker) Pen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pen"}).(*gui.QPen)
}

func (ptr *QLegendMarker) ConnectPenChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectPenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPenChanged"})
}

func (ptr *QLegendMarker) PenChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PenChanged"})
}

func (ptr *QLegendMarker) ConnectSeries(f func() *QAbstractSeries) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeries"})
}

func (ptr *QLegendMarker) Series() *QAbstractSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QAbstractSeries)
}

func (ptr *QLegendMarker) SetBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrush", brush})
}

func (ptr *QLegendMarker) SetFont(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFont", font})
}

func (ptr *QLegendMarker) SetLabel(label string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabel", label})
}

func (ptr *QLegendMarker) SetLabelBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelBrush", brush})
}

func (ptr *QLegendMarker) SetPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPen", pen})
}

func (ptr *QLegendMarker) SetShape(shape QLegend__MarkerShape) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShape", shape})
}

func (ptr *QLegendMarker) SetVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisible", visible})
}

func (ptr *QLegendMarker) Shape() QLegend__MarkerShape {

	return QLegend__MarkerShape(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Shape"}).(float64))
}

func (ptr *QLegendMarker) ConnectShapeChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShapeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectShapeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShapeChanged"})
}

func (ptr *QLegendMarker) ShapeChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShapeChanged"})
}

func (ptr *QLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QLegendMarker) Type() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QLegendMarker) ConnectVisibleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVisibleChanged"})
}

func (ptr *QLegendMarker) VisibleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VisibleChanged"})
}

func (ptr *QLegendMarker) ConnectDestroyQLegendMarker(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLegendMarker", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLegendMarker) DisconnectDestroyQLegendMarker() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLegendMarker"})
}

func (ptr *QLegendMarker) DestroyQLegendMarker() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLegendMarker"})
}

func (ptr *QLegendMarker) DestroyQLegendMarkerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLegendMarkerDefault"})
}

func (ptr *QLegendMarker) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QLegendMarker) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QLegendMarker) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QLegendMarker) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QLegendMarker) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QLegendMarker) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QLegendMarker) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QLegendMarker) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QLegendMarker) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QLegendMarker) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QLegendMarker) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QLegendMarker) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QLegendMarker) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QLegendMarker) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QLegendMarker) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QLegendMarker) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QLegendMarker) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QLegendMarker) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QLegendMarker) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QLegendMarker) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QLegendMarker) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QLineSeries) InitFromInternal(ptr uintptr, name string) {
	n.QXYSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLineSeries) ClassNameInternalF() string {
	return n.QXYSeries_PTR().ClassNameInternalF()
}

func NewQLineSeriesFromPointer(ptr unsafe.Pointer) (n *QLineSeries) {
	n = new(QLineSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QLineSeries")
	return
}
func NewQLineSeries(parent core.QObject_ITF) *QLineSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQLineSeries", "", parent}).(*QLineSeries)
}

func (ptr *QLineSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLineSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QLineSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QLineSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QLineSeries) ConnectDestroyQLineSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLineSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLineSeries) DisconnectDestroyQLineSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLineSeries"})
}

func (ptr *QLineSeries) DestroyQLineSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLineSeries"})
}

func (ptr *QLineSeries) DestroyQLineSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLineSeriesDefault"})
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

func (n *QLogValueAxis) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractAxis_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLogValueAxis) ClassNameInternalF() string {
	return n.QAbstractAxis_PTR().ClassNameInternalF()
}

func NewQLogValueAxisFromPointer(ptr unsafe.Pointer) (n *QLogValueAxis) {
	n = new(QLogValueAxis)
	n.InitFromInternal(uintptr(ptr), "charts.QLogValueAxis")
	return
}
func NewQLogValueAxis(parent core.QObject_ITF) *QLogValueAxis {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQLogValueAxis", "", parent}).(*QLogValueAxis)
}

func (ptr *QLogValueAxis) Base() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Base"}).(float64)
}

func (ptr *QLogValueAxis) ConnectBaseChanged(f func(base float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBaseChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValueAxis) DisconnectBaseChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBaseChanged"})
}

func (ptr *QLogValueAxis) BaseChanged(base float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseChanged", base})
}

func (ptr *QLogValueAxis) LabelFormat() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelFormat"}).(string)
}

func (ptr *QLogValueAxis) ConnectLabelFormatChanged(f func(format string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelFormatChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValueAxis) DisconnectLabelFormatChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelFormatChanged"})
}

func (ptr *QLogValueAxis) LabelFormatChanged(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelFormatChanged", format})
}

func (ptr *QLogValueAxis) Max() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Max"}).(float64)
}

func (ptr *QLogValueAxis) ConnectMaxChanged(f func(max float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMaxChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValueAxis) DisconnectMaxChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMaxChanged"})
}

func (ptr *QLogValueAxis) MaxChanged(max float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxChanged", max})
}

func (ptr *QLogValueAxis) Min() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Min"}).(float64)
}

func (ptr *QLogValueAxis) ConnectMinChanged(f func(min float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValueAxis) DisconnectMinChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinChanged"})
}

func (ptr *QLogValueAxis) MinChanged(min float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinChanged", min})
}

func (ptr *QLogValueAxis) MinorTickCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinorTickCount"}).(float64))
}

func (ptr *QLogValueAxis) ConnectMinorTickCountChanged(f func(minorTickCount int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinorTickCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValueAxis) DisconnectMinorTickCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinorTickCountChanged"})
}

func (ptr *QLogValueAxis) MinorTickCountChanged(minorTickCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinorTickCountChanged", minorTickCount})
}

func (ptr *QLogValueAxis) ConnectRangeChanged(f func(min float64, max float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRangeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValueAxis) DisconnectRangeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRangeChanged"})
}

func (ptr *QLogValueAxis) RangeChanged(min float64, max float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RangeChanged", min, max})
}

func (ptr *QLogValueAxis) SetBase(base float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBase", base})
}

func (ptr *QLogValueAxis) SetLabelFormat(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelFormat", format})
}

func (ptr *QLogValueAxis) SetMax(max float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMax", max})
}

func (ptr *QLogValueAxis) SetMin(min float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMin", min})
}

func (ptr *QLogValueAxis) SetMinorTickCount(minorTickCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinorTickCount", minorTickCount})
}

func (ptr *QLogValueAxis) SetRange(min float64, max float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRange", min, max})
}

func (ptr *QLogValueAxis) TickCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickCount"}).(float64))
}

func (ptr *QLogValueAxis) ConnectTickCountChanged(f func(tickCount int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTickCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValueAxis) DisconnectTickCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTickCountChanged"})
}

func (ptr *QLogValueAxis) TickCountChanged(tickCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickCountChanged", tickCount})
}

func (ptr *QLogValueAxis) ConnectType(f func() QAbstractAxis__AxisType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValueAxis) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QLogValueAxis) Type() QAbstractAxis__AxisType {

	return QAbstractAxis__AxisType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QLogValueAxis) TypeDefault() QAbstractAxis__AxisType {

	return QAbstractAxis__AxisType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QLogValueAxis) ConnectDestroyQLogValueAxis(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLogValueAxis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValueAxis) DisconnectDestroyQLogValueAxis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLogValueAxis"})
}

func (ptr *QLogValueAxis) DestroyQLogValueAxis() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLogValueAxis"})
}

func (ptr *QLogValueAxis) DestroyQLogValueAxisDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLogValueAxisDefault"})
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

func (n *QPercentBarSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractBarSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPercentBarSeries) ClassNameInternalF() string {
	return n.QAbstractBarSeries_PTR().ClassNameInternalF()
}

func NewQPercentBarSeriesFromPointer(ptr unsafe.Pointer) (n *QPercentBarSeries) {
	n = new(QPercentBarSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QPercentBarSeries")
	return
}
func NewQPercentBarSeries(parent core.QObject_ITF) *QPercentBarSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQPercentBarSeries", "", parent}).(*QPercentBarSeries)
}

func (ptr *QPercentBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPercentBarSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QPercentBarSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QPercentBarSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QPercentBarSeries) ConnectDestroyQPercentBarSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPercentBarSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPercentBarSeries) DisconnectDestroyQPercentBarSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPercentBarSeries"})
}

func (ptr *QPercentBarSeries) DestroyQPercentBarSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPercentBarSeries"})
}

func (ptr *QPercentBarSeries) DestroyQPercentBarSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPercentBarSeriesDefault"})
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

func (n *QPieLegendMarker) InitFromInternal(ptr uintptr, name string) {
	n.QLegendMarker_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPieLegendMarker) ClassNameInternalF() string {
	return n.QLegendMarker_PTR().ClassNameInternalF()
}

func NewQPieLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QPieLegendMarker) {
	n = new(QPieLegendMarker)
	n.InitFromInternal(uintptr(ptr), "charts.QPieLegendMarker")
	return
}
func (ptr *QPieLegendMarker) ConnectSeries(f func() *QPieSeries) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieLegendMarker) DisconnectSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeries"})
}

func (ptr *QPieLegendMarker) Series() *QPieSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QPieSeries)
}

func (ptr *QPieLegendMarker) SeriesDefault() *QPieSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesDefault"}).(*QPieSeries)
}

func (ptr *QPieLegendMarker) Slice() *QPieSlice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Slice"}).(*QPieSlice)
}

func (ptr *QPieLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieLegendMarker) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QPieLegendMarker) Type() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QPieLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QPieLegendMarker) ConnectDestroyQPieLegendMarker(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPieLegendMarker", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieLegendMarker) DisconnectDestroyQPieLegendMarker() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPieLegendMarker"})
}

func (ptr *QPieLegendMarker) DestroyQPieLegendMarker() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPieLegendMarker"})
}

func (ptr *QPieLegendMarker) DestroyQPieLegendMarkerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPieLegendMarkerDefault"})
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

func (n *QPieModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPieModelMapper) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQPieModelMapperFromPointer(ptr unsafe.Pointer) (n *QPieModelMapper) {
	n = new(QPieModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QPieModelMapper")
	return
}

func (ptr *QPieModelMapper) DestroyQPieModelMapper() {
}

func (ptr *QPieModelMapper) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QPieModelMapper) First() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "First"}).(float64))
}

func (ptr *QPieModelMapper) LabelsSection() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsSection"}).(float64))
}

func (ptr *QPieModelMapper) Orientation() core.Qt__Orientation {

	return core.Qt__Orientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Orientation"}).(float64))
}

func (ptr *QPieModelMapper) SetCount(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCount", count})
}

func (ptr *QPieModelMapper) SetFirst(first int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirst", first})
}

func (ptr *QPieModelMapper) SetLabelsSection(labelsSection int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsSection", labelsSection})
}

func (ptr *QPieModelMapper) SetOrientation(orientation core.Qt__Orientation) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOrientation", orientation})
}

func (ptr *QPieModelMapper) SetValuesSection(valuesSection int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValuesSection", valuesSection})
}

func (ptr *QPieModelMapper) ValuesSection() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValuesSection"}).(float64))
}

func (ptr *QPieModelMapper) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QPieModelMapper) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QPieModelMapper) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QPieModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QPieModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QPieModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QPieModelMapper) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QPieModelMapper) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QPieModelMapper) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QPieModelMapper) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QPieModelMapper) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QPieModelMapper) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QPieModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QPieModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QPieModelMapper) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QPieModelMapper) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QPieModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QPieModelMapper) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QPieModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QPieModelMapper) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QPieModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QPieSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPieSeries) ClassNameInternalF() string {
	return n.QAbstractSeries_PTR().ClassNameInternalF()
}

func NewQPieSeriesFromPointer(ptr unsafe.Pointer) (n *QPieSeries) {
	n = new(QPieSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QPieSeries")
	return
}
func NewQPieSeries(parent core.QObject_ITF) *QPieSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQPieSeries", "", parent}).(*QPieSeries)
}

func (ptr *QPieSeries) ConnectAdded(f func(slices []*QPieSlice)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAdded"})
}

func (ptr *QPieSeries) Added(slices []*QPieSlice) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Added", slices})
}

func (ptr *QPieSeries) Append(slice QPieSlice_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", slice}).(bool)
}

func (ptr *QPieSeries) Append2(slices []*QPieSlice) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append2", slices}).(bool)
}

func (ptr *QPieSeries) Append3(label string, value float64) *QPieSlice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append3", label, value}).(*QPieSlice)
}

func (ptr *QPieSeries) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QPieSeries) ConnectClicked(f func(slice *QPieSlice)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QPieSeries) Clicked(slice QPieSlice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked", slice})
}

func (ptr *QPieSeries) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QPieSeries) ConnectCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCountChanged"})
}

func (ptr *QPieSeries) CountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CountChanged"})
}

func (ptr *QPieSeries) ConnectDoubleClicked(f func(slice *QPieSlice)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDoubleClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectDoubleClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDoubleClicked"})
}

func (ptr *QPieSeries) DoubleClicked(slice QPieSlice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoubleClicked", slice})
}

func (ptr *QPieSeries) HoleSize() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HoleSize"}).(float64)
}

func (ptr *QPieSeries) HorizontalPosition() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HorizontalPosition"}).(float64)
}

func (ptr *QPieSeries) ConnectHovered(f func(slice *QPieSlice, state bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QPieSeries) Hovered(slice QPieSlice_ITF, state bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", slice, state})
}

func (ptr *QPieSeries) Insert(index int, slice QPieSlice_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", index, slice}).(bool)
}

func (ptr *QPieSeries) IsEmpty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEmpty"}).(bool)
}

func (ptr *QPieSeries) PieEndAngle() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PieEndAngle"}).(float64)
}

func (ptr *QPieSeries) PieSize() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PieSize"}).(float64)
}

func (ptr *QPieSeries) PieStartAngle() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PieStartAngle"}).(float64)
}

func (ptr *QPieSeries) ConnectPressed(f func(slice *QPieSlice)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPressed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectPressed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPressed"})
}

func (ptr *QPieSeries) Pressed(slice QPieSlice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressed", slice})
}

func (ptr *QPieSeries) ConnectReleased(f func(slice *QPieSlice)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleased", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectReleased() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleased"})
}

func (ptr *QPieSeries) Released(slice QPieSlice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Released", slice})
}

func (ptr *QPieSeries) Remove(slice QPieSlice_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", slice}).(bool)
}

func (ptr *QPieSeries) ConnectRemoved(f func(slices []*QPieSlice)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemoved"})
}

func (ptr *QPieSeries) Removed(slices []*QPieSlice) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Removed", slices})
}

func (ptr *QPieSeries) SetHoleSize(holeSize float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHoleSize", holeSize})
}

func (ptr *QPieSeries) SetHorizontalPosition(relativePosition float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHorizontalPosition", relativePosition})
}

func (ptr *QPieSeries) SetLabelsPosition(position QPieSlice__LabelPosition) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsPosition", position})
}

func (ptr *QPieSeries) SetLabelsVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsVisible", visible})
}

func (ptr *QPieSeries) SetPieEndAngle(angle float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPieEndAngle", angle})
}

func (ptr *QPieSeries) SetPieSize(relativeSize float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPieSize", relativeSize})
}

func (ptr *QPieSeries) SetPieStartAngle(startAngle float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPieStartAngle", startAngle})
}

func (ptr *QPieSeries) SetVerticalPosition(relativePosition float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVerticalPosition", relativePosition})
}

func (ptr *QPieSeries) Slices() []*QPieSlice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Slices"}).([]*QPieSlice)
}

func (ptr *QPieSeries) Sum() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Sum"}).(float64)
}

func (ptr *QPieSeries) ConnectSumChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSumChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectSumChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSumChanged"})
}

func (ptr *QPieSeries) SumChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SumChanged"})
}

func (ptr *QPieSeries) Take(slice QPieSlice_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Take", slice}).(bool)
}

func (ptr *QPieSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QPieSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QPieSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QPieSeries) VerticalPosition() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VerticalPosition"}).(float64)
}

func (ptr *QPieSeries) ConnectDestroyQPieSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPieSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSeries) DisconnectDestroyQPieSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPieSeries"})
}

func (ptr *QPieSeries) DestroyQPieSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPieSeries"})
}

func (ptr *QPieSeries) DestroyQPieSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPieSeriesDefault"})
}

func (ptr *QPieSeries) __added_slices_atList(i int) *QPieSlice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__added_slices_atList", i}).(*QPieSlice)
}

func (ptr *QPieSeries) __added_slices_setList(i QPieSlice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__added_slices_setList", i})
}

func (ptr *QPieSeries) __added_slices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__added_slices_newList"}).(unsafe.Pointer)
}

func (ptr *QPieSeries) __append_slices_atList2(i int) *QPieSlice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_slices_atList2", i}).(*QPieSlice)
}

func (ptr *QPieSeries) __append_slices_setList2(i QPieSlice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_slices_setList2", i})
}

func (ptr *QPieSeries) __append_slices_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_slices_newList2"}).(unsafe.Pointer)
}

func (ptr *QPieSeries) __removed_slices_atList(i int) *QPieSlice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__removed_slices_atList", i}).(*QPieSlice)
}

func (ptr *QPieSeries) __removed_slices_setList(i QPieSlice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__removed_slices_setList", i})
}

func (ptr *QPieSeries) __removed_slices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__removed_slices_newList"}).(unsafe.Pointer)
}

func (ptr *QPieSeries) __slices_atList(i int) *QPieSlice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__slices_atList", i}).(*QPieSlice)
}

func (ptr *QPieSeries) __slices_setList(i QPieSlice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__slices_setList", i})
}

func (ptr *QPieSeries) __slices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__slices_newList"}).(unsafe.Pointer)
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

func (n *QPieSlice) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPieSlice) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQPieSliceFromPointer(ptr unsafe.Pointer) (n *QPieSlice) {
	n = new(QPieSlice)
	n.InitFromInternal(uintptr(ptr), "charts.QPieSlice")
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

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQPieSlice", "", parent}).(*QPieSlice)
}

func NewQPieSlice2(label string, value float64, parent core.QObject_ITF) *QPieSlice {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQPieSlice2", "", label, value, parent}).(*QPieSlice)
}

func (ptr *QPieSlice) AngleSpan() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AngleSpan"}).(float64)
}

func (ptr *QPieSlice) ConnectAngleSpanChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAngleSpanChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectAngleSpanChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAngleSpanChanged"})
}

func (ptr *QPieSlice) AngleSpanChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AngleSpanChanged"})
}

func (ptr *QPieSlice) BorderColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderColor"}).(*gui.QColor)
}

func (ptr *QPieSlice) ConnectBorderColorChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBorderColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectBorderColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBorderColorChanged"})
}

func (ptr *QPieSlice) BorderColorChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderColorChanged"})
}

func (ptr *QPieSlice) BorderWidth() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderWidth"}).(float64))
}

func (ptr *QPieSlice) ConnectBorderWidthChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBorderWidthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectBorderWidthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBorderWidthChanged"})
}

func (ptr *QPieSlice) BorderWidthChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderWidthChanged"})
}

func (ptr *QPieSlice) Brush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Brush"}).(*gui.QBrush)
}

func (ptr *QPieSlice) ConnectBrushChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBrushChanged"})
}

func (ptr *QPieSlice) BrushChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BrushChanged"})
}

func (ptr *QPieSlice) ConnectClicked(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QPieSlice) Clicked() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked"})
}

func (ptr *QPieSlice) Color() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Color"}).(*gui.QColor)
}

func (ptr *QPieSlice) ConnectColorChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColorChanged"})
}

func (ptr *QPieSlice) ColorChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorChanged"})
}

func (ptr *QPieSlice) ConnectDoubleClicked(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDoubleClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectDoubleClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDoubleClicked"})
}

func (ptr *QPieSlice) DoubleClicked() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoubleClicked"})
}

func (ptr *QPieSlice) ExplodeDistanceFactor() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExplodeDistanceFactor"}).(float64)
}

func (ptr *QPieSlice) ConnectHovered(f func(state bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QPieSlice) Hovered(state bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", state})
}

func (ptr *QPieSlice) IsExploded() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsExploded"}).(bool)
}

func (ptr *QPieSlice) IsLabelVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLabelVisible"}).(bool)
}

func (ptr *QPieSlice) Label() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Label"}).(string)
}

func (ptr *QPieSlice) LabelArmLengthFactor() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelArmLengthFactor"}).(float64)
}

func (ptr *QPieSlice) LabelBrush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBrush"}).(*gui.QBrush)
}

func (ptr *QPieSlice) ConnectLabelBrushChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelBrushChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectLabelBrushChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelBrushChanged"})
}

func (ptr *QPieSlice) LabelBrushChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBrushChanged"})
}

func (ptr *QPieSlice) ConnectLabelChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectLabelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelChanged"})
}

func (ptr *QPieSlice) LabelChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelChanged"})
}

func (ptr *QPieSlice) LabelColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelColor"}).(*gui.QColor)
}

func (ptr *QPieSlice) ConnectLabelColorChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectLabelColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelColorChanged"})
}

func (ptr *QPieSlice) LabelColorChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelColorChanged"})
}

func (ptr *QPieSlice) LabelFont() *gui.QFont {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelFont"}).(*gui.QFont)
}

func (ptr *QPieSlice) ConnectLabelFontChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelFontChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectLabelFontChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelFontChanged"})
}

func (ptr *QPieSlice) LabelFontChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelFontChanged"})
}

func (ptr *QPieSlice) LabelPosition() QPieSlice__LabelPosition {

	return QPieSlice__LabelPosition(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelPosition"}).(float64))
}

func (ptr *QPieSlice) ConnectLabelVisibleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectLabelVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelVisibleChanged"})
}

func (ptr *QPieSlice) LabelVisibleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelVisibleChanged"})
}

func (ptr *QPieSlice) Pen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pen"}).(*gui.QPen)
}

func (ptr *QPieSlice) ConnectPenChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectPenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPenChanged"})
}

func (ptr *QPieSlice) PenChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PenChanged"})
}

func (ptr *QPieSlice) Percentage() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Percentage"}).(float64)
}

func (ptr *QPieSlice) ConnectPercentageChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPercentageChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectPercentageChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPercentageChanged"})
}

func (ptr *QPieSlice) PercentageChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PercentageChanged"})
}

func (ptr *QPieSlice) ConnectPressed(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPressed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectPressed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPressed"})
}

func (ptr *QPieSlice) Pressed() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressed"})
}

func (ptr *QPieSlice) ConnectReleased(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleased", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectReleased() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleased"})
}

func (ptr *QPieSlice) Released() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Released"})
}

func (ptr *QPieSlice) Series() *QPieSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QPieSeries)
}

func (ptr *QPieSlice) SetBorderColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBorderColor", color})
}

func (ptr *QPieSlice) SetBorderWidth(width int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBorderWidth", width})
}

func (ptr *QPieSlice) SetBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrush", brush})
}

func (ptr *QPieSlice) SetColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColor", color})
}

func (ptr *QPieSlice) SetExplodeDistanceFactor(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExplodeDistanceFactor", factor})
}

func (ptr *QPieSlice) SetExploded(exploded bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExploded", exploded})
}

func (ptr *QPieSlice) SetLabel(label string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabel", label})
}

func (ptr *QPieSlice) SetLabelArmLengthFactor(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelArmLengthFactor", factor})
}

func (ptr *QPieSlice) SetLabelBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelBrush", brush})
}

func (ptr *QPieSlice) SetLabelColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelColor", color})
}

func (ptr *QPieSlice) SetLabelFont(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelFont", font})
}

func (ptr *QPieSlice) SetLabelPosition(position QPieSlice__LabelPosition) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelPosition", position})
}

func (ptr *QPieSlice) SetLabelVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelVisible", visible})
}

func (ptr *QPieSlice) SetPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPen", pen})
}

func (ptr *QPieSlice) SetValue(value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", value})
}

func (ptr *QPieSlice) StartAngle() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartAngle"}).(float64)
}

func (ptr *QPieSlice) ConnectStartAngleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartAngleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectStartAngleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartAngleChanged"})
}

func (ptr *QPieSlice) StartAngleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartAngleChanged"})
}

func (ptr *QPieSlice) Value() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(float64)
}

func (ptr *QPieSlice) ConnectValueChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValueChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectValueChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValueChanged"})
}

func (ptr *QPieSlice) ValueChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueChanged"})
}

func (ptr *QPieSlice) ConnectDestroyQPieSlice(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPieSlice", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPieSlice) DisconnectDestroyQPieSlice() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPieSlice"})
}

func (ptr *QPieSlice) DestroyQPieSlice() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPieSlice"})
}

func (ptr *QPieSlice) DestroyQPieSliceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPieSliceDefault"})
}

func (ptr *QPieSlice) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QPieSlice) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QPieSlice) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QPieSlice) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QPieSlice) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QPieSlice) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QPieSlice) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QPieSlice) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QPieSlice) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QPieSlice) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QPieSlice) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QPieSlice) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QPieSlice) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QPieSlice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QPieSlice) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QPieSlice) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QPieSlice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QPieSlice) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QPieSlice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QPieSlice) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QPieSlice) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QPolarChart) InitFromInternal(ptr uintptr, name string) {
	n.QChart_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPolarChart) ClassNameInternalF() string {
	return n.QChart_PTR().ClassNameInternalF()
}

func NewQPolarChartFromPointer(ptr unsafe.Pointer) (n *QPolarChart) {
	n = new(QPolarChart)
	n.InitFromInternal(uintptr(ptr), "charts.QPolarChart")
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

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQPolarChart", "", parent, wFlags}).(*QPolarChart)
}

func (ptr *QPolarChart) AddAxis(axis QAbstractAxis_ITF, polarOrientation QPolarChart__PolarOrientation) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddAxis", axis, polarOrientation})
}

func (ptr *QPolarChart) Axes(polarOrientation QPolarChart__PolarOrientation, series QAbstractSeries_ITF) []*QAbstractAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Axes", polarOrientation, series}).([]*QAbstractAxis)
}

func QPolarChart_AxisPolarOrientation(axis QAbstractAxis_ITF) QPolarChart__PolarOrientation {

	return QPolarChart__PolarOrientation(internal.CallLocalFunction([]interface{}{"", "", "charts.QPolarChart_AxisPolarOrientation", "", axis}).(float64))
}

func (ptr *QPolarChart) AxisPolarOrientation(axis QAbstractAxis_ITF) QPolarChart__PolarOrientation {

	return QPolarChart__PolarOrientation(internal.CallLocalFunction([]interface{}{"", "", "charts.QPolarChart_AxisPolarOrientation", "", axis}).(float64))
}

func (ptr *QPolarChart) ConnectDestroyQPolarChart(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPolarChart", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPolarChart) DisconnectDestroyQPolarChart() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPolarChart"})
}

func (ptr *QPolarChart) DestroyQPolarChart() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPolarChart"})
}

func (ptr *QPolarChart) DestroyQPolarChartDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPolarChartDefault"})
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

func (n *QScatterSeries) InitFromInternal(ptr uintptr, name string) {
	n.QXYSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScatterSeries) ClassNameInternalF() string {
	return n.QXYSeries_PTR().ClassNameInternalF()
}

func NewQScatterSeriesFromPointer(ptr unsafe.Pointer) (n *QScatterSeries) {
	n = new(QScatterSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QScatterSeries")
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

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQScatterSeries", "", parent}).(*QScatterSeries)
}

func (ptr *QScatterSeries) BorderColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderColor"}).(*gui.QColor)
}

func (ptr *QScatterSeries) ConnectBorderColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBorderColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterSeries) DisconnectBorderColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBorderColorChanged"})
}

func (ptr *QScatterSeries) BorderColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderColorChanged", color})
}

func (ptr *QScatterSeries) MarkerShape() QScatterSeries__MarkerShape {

	return QScatterSeries__MarkerShape(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MarkerShape"}).(float64))
}

func (ptr *QScatterSeries) ConnectMarkerShapeChanged(f func(shape QScatterSeries__MarkerShape)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMarkerShapeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterSeries) DisconnectMarkerShapeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMarkerShapeChanged"})
}

func (ptr *QScatterSeries) MarkerShapeChanged(shape QScatterSeries__MarkerShape) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MarkerShapeChanged", shape})
}

func (ptr *QScatterSeries) MarkerSize() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MarkerSize"}).(float64)
}

func (ptr *QScatterSeries) ConnectMarkerSizeChanged(f func(size float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMarkerSizeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterSeries) DisconnectMarkerSizeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMarkerSizeChanged"})
}

func (ptr *QScatterSeries) MarkerSizeChanged(size float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MarkerSizeChanged", size})
}

func (ptr *QScatterSeries) SetBorderColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBorderColor", color})
}

func (ptr *QScatterSeries) SetMarkerShape(shape QScatterSeries__MarkerShape) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMarkerShape", shape})
}

func (ptr *QScatterSeries) SetMarkerSize(size float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMarkerSize", size})
}

func (ptr *QScatterSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QScatterSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QScatterSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QScatterSeries) ConnectDestroyQScatterSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQScatterSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterSeries) DisconnectDestroyQScatterSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQScatterSeries"})
}

func (ptr *QScatterSeries) DestroyQScatterSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScatterSeries"})
}

func (ptr *QScatterSeries) DestroyQScatterSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScatterSeriesDefault"})
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

func (n *QSplineSeries) InitFromInternal(ptr uintptr, name string) {
	n.QLineSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSplineSeries) ClassNameInternalF() string {
	return n.QLineSeries_PTR().ClassNameInternalF()
}

func NewQSplineSeriesFromPointer(ptr unsafe.Pointer) (n *QSplineSeries) {
	n = new(QSplineSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QSplineSeries")
	return
}
func NewQSplineSeries(parent core.QObject_ITF) *QSplineSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQSplineSeries", "", parent}).(*QSplineSeries)
}

func (ptr *QSplineSeries) ConnectDestroyQSplineSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSplineSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSplineSeries) DisconnectDestroyQSplineSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSplineSeries"})
}

func (ptr *QSplineSeries) DestroyQSplineSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSplineSeries"})
}

func (ptr *QSplineSeries) DestroyQSplineSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSplineSeriesDefault"})
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

func (n *QStackedBarSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractBarSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QStackedBarSeries) ClassNameInternalF() string {
	return n.QAbstractBarSeries_PTR().ClassNameInternalF()
}

func NewQStackedBarSeriesFromPointer(ptr unsafe.Pointer) (n *QStackedBarSeries) {
	n = new(QStackedBarSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QStackedBarSeries")
	return
}
func NewQStackedBarSeries(parent core.QObject_ITF) *QStackedBarSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQStackedBarSeries", "", parent}).(*QStackedBarSeries)
}

func (ptr *QStackedBarSeries) ConnectType(f func() QAbstractSeries__SeriesType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QStackedBarSeries) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QStackedBarSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QStackedBarSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QStackedBarSeries) ConnectDestroyQStackedBarSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQStackedBarSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QStackedBarSeries) DisconnectDestroyQStackedBarSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQStackedBarSeries"})
}

func (ptr *QStackedBarSeries) DestroyQStackedBarSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQStackedBarSeries"})
}

func (ptr *QStackedBarSeries) DestroyQStackedBarSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQStackedBarSeriesDefault"})
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

func (n *QVBarModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVBarModelMapper) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQVBarModelMapperFromPointer(ptr unsafe.Pointer) (n *QVBarModelMapper) {
	n = new(QVBarModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QVBarModelMapper")
	return
}

func (ptr *QVBarModelMapper) DestroyQVBarModelMapper() {
}

func NewQVBarModelMapper(parent core.QObject_ITF) *QVBarModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQVBarModelMapper", "", parent}).(*QVBarModelMapper)
}

func (ptr *QVBarModelMapper) FirstBarSetColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstBarSetColumn"}).(float64))
}

func (ptr *QVBarModelMapper) ConnectFirstBarSetColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstBarSetColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBarModelMapper) DisconnectFirstBarSetColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstBarSetColumnChanged"})
}

func (ptr *QVBarModelMapper) FirstBarSetColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstBarSetColumnChanged"})
}

func (ptr *QVBarModelMapper) FirstRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstRow"}).(float64))
}

func (ptr *QVBarModelMapper) ConnectFirstRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBarModelMapper) DisconnectFirstRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstRowChanged"})
}

func (ptr *QVBarModelMapper) FirstRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstRowChanged"})
}

func (ptr *QVBarModelMapper) LastBarSetColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastBarSetColumn"}).(float64))
}

func (ptr *QVBarModelMapper) ConnectLastBarSetColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLastBarSetColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBarModelMapper) DisconnectLastBarSetColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLastBarSetColumnChanged"})
}

func (ptr *QVBarModelMapper) LastBarSetColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastBarSetColumnChanged"})
}

func (ptr *QVBarModelMapper) Model() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*core.QAbstractItemModel)
}

func (ptr *QVBarModelMapper) ConnectModelReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModelReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBarModelMapper) DisconnectModelReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModelReplaced"})
}

func (ptr *QVBarModelMapper) ModelReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModelReplaced"})
}

func (ptr *QVBarModelMapper) RowCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCount"}).(float64))
}

func (ptr *QVBarModelMapper) ConnectRowCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBarModelMapper) DisconnectRowCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowCountChanged"})
}

func (ptr *QVBarModelMapper) RowCountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountChanged"})
}

func (ptr *QVBarModelMapper) Series() *QAbstractBarSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QAbstractBarSeries)
}

func (ptr *QVBarModelMapper) ConnectSeriesReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBarModelMapper) DisconnectSeriesReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesReplaced"})
}

func (ptr *QVBarModelMapper) SeriesReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesReplaced"})
}

func (ptr *QVBarModelMapper) SetFirstBarSetColumn(firstBarSetColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstBarSetColumn", firstBarSetColumn})
}

func (ptr *QVBarModelMapper) SetFirstRow(firstRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstRow", firstRow})
}

func (ptr *QVBarModelMapper) SetLastBarSetColumn(lastBarSetColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastBarSetColumn", lastBarSetColumn})
}

func (ptr *QVBarModelMapper) SetModel(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModel", model})
}

func (ptr *QVBarModelMapper) SetRowCount(rowCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowCount", rowCount})
}

func (ptr *QVBarModelMapper) SetSeries(series QAbstractBarSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSeries", series})
}

func (ptr *QVBarModelMapper) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QVBarModelMapper) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QVBarModelMapper) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QVBarModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QVBarModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QVBarModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QVBarModelMapper) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QVBarModelMapper) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QVBarModelMapper) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QVBarModelMapper) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QVBarModelMapper) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QVBarModelMapper) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QVBarModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QVBarModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QVBarModelMapper) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QVBarModelMapper) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QVBarModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QVBarModelMapper) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QVBarModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QVBarModelMapper) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QVBarModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QVBoxPlotModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVBoxPlotModelMapper) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQVBoxPlotModelMapperFromPointer(ptr unsafe.Pointer) (n *QVBoxPlotModelMapper) {
	n = new(QVBoxPlotModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QVBoxPlotModelMapper")
	return
}

func (ptr *QVBoxPlotModelMapper) DestroyQVBoxPlotModelMapper() {
}

func NewQVBoxPlotModelMapper(parent core.QObject_ITF) *QVBoxPlotModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQVBoxPlotModelMapper", "", parent}).(*QVBoxPlotModelMapper)
}

func (ptr *QVBoxPlotModelMapper) FirstBoxSetColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstBoxSetColumn"}).(float64))
}

func (ptr *QVBoxPlotModelMapper) ConnectFirstBoxSetColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstBoxSetColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBoxPlotModelMapper) DisconnectFirstBoxSetColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstBoxSetColumnChanged"})
}

func (ptr *QVBoxPlotModelMapper) FirstBoxSetColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstBoxSetColumnChanged"})
}

func (ptr *QVBoxPlotModelMapper) FirstRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstRow"}).(float64))
}

func (ptr *QVBoxPlotModelMapper) ConnectFirstRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBoxPlotModelMapper) DisconnectFirstRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstRowChanged"})
}

func (ptr *QVBoxPlotModelMapper) FirstRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstRowChanged"})
}

func (ptr *QVBoxPlotModelMapper) LastBoxSetColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastBoxSetColumn"}).(float64))
}

func (ptr *QVBoxPlotModelMapper) ConnectLastBoxSetColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLastBoxSetColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBoxPlotModelMapper) DisconnectLastBoxSetColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLastBoxSetColumnChanged"})
}

func (ptr *QVBoxPlotModelMapper) LastBoxSetColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastBoxSetColumnChanged"})
}

func (ptr *QVBoxPlotModelMapper) Model() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*core.QAbstractItemModel)
}

func (ptr *QVBoxPlotModelMapper) ConnectModelReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModelReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBoxPlotModelMapper) DisconnectModelReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModelReplaced"})
}

func (ptr *QVBoxPlotModelMapper) ModelReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModelReplaced"})
}

func (ptr *QVBoxPlotModelMapper) RowCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCount"}).(float64))
}

func (ptr *QVBoxPlotModelMapper) ConnectRowCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBoxPlotModelMapper) DisconnectRowCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowCountChanged"})
}

func (ptr *QVBoxPlotModelMapper) RowCountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountChanged"})
}

func (ptr *QVBoxPlotModelMapper) Series() *QBoxPlotSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QBoxPlotSeries)
}

func (ptr *QVBoxPlotModelMapper) ConnectSeriesReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVBoxPlotModelMapper) DisconnectSeriesReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesReplaced"})
}

func (ptr *QVBoxPlotModelMapper) SeriesReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesReplaced"})
}

func (ptr *QVBoxPlotModelMapper) SetFirstBoxSetColumn(firstBoxSetColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstBoxSetColumn", firstBoxSetColumn})
}

func (ptr *QVBoxPlotModelMapper) SetFirstRow(firstRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstRow", firstRow})
}

func (ptr *QVBoxPlotModelMapper) SetLastBoxSetColumn(lastBoxSetColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastBoxSetColumn", lastBoxSetColumn})
}

func (ptr *QVBoxPlotModelMapper) SetModel(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModel", model})
}

func (ptr *QVBoxPlotModelMapper) SetRowCount(rowCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowCount", rowCount})
}

func (ptr *QVBoxPlotModelMapper) SetSeries(series QBoxPlotSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSeries", series})
}

func (ptr *QVBoxPlotModelMapper) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QVBoxPlotModelMapper) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QVBoxPlotModelMapper) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QVBoxPlotModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QVBoxPlotModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QVBoxPlotModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QVBoxPlotModelMapper) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QVBoxPlotModelMapper) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QVBoxPlotModelMapper) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QVBoxPlotModelMapper) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QVBoxPlotModelMapper) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QVBoxPlotModelMapper) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QVBoxPlotModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QVBoxPlotModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QVBoxPlotModelMapper) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QVBoxPlotModelMapper) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QVBoxPlotModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QVBoxPlotModelMapper) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QVBoxPlotModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QVBoxPlotModelMapper) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QVBoxPlotModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QVCandlestickModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QCandlestickModelMapper_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVCandlestickModelMapper) ClassNameInternalF() string {
	return n.QCandlestickModelMapper_PTR().ClassNameInternalF()
}

func NewQVCandlestickModelMapperFromPointer(ptr unsafe.Pointer) (n *QVCandlestickModelMapper) {
	n = new(QVCandlestickModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QVCandlestickModelMapper")
	return
}

func (ptr *QVCandlestickModelMapper) DestroyQVCandlestickModelMapper() {
}

func NewQVCandlestickModelMapper(parent core.QObject_ITF) *QVCandlestickModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQVCandlestickModelMapper", "", parent}).(*QVCandlestickModelMapper)
}

func (ptr *QVCandlestickModelMapper) CloseRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseRow"}).(float64))
}

func (ptr *QVCandlestickModelMapper) ConnectCloseRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCloseRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVCandlestickModelMapper) DisconnectCloseRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCloseRowChanged"})
}

func (ptr *QVCandlestickModelMapper) CloseRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseRowChanged"})
}

func (ptr *QVCandlestickModelMapper) FirstSetColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstSetColumn"}).(float64))
}

func (ptr *QVCandlestickModelMapper) ConnectFirstSetColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstSetColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVCandlestickModelMapper) DisconnectFirstSetColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstSetColumnChanged"})
}

func (ptr *QVCandlestickModelMapper) FirstSetColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstSetColumnChanged"})
}

func (ptr *QVCandlestickModelMapper) HighRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HighRow"}).(float64))
}

func (ptr *QVCandlestickModelMapper) ConnectHighRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHighRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVCandlestickModelMapper) DisconnectHighRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHighRowChanged"})
}

func (ptr *QVCandlestickModelMapper) HighRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HighRowChanged"})
}

func (ptr *QVCandlestickModelMapper) LastSetColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastSetColumn"}).(float64))
}

func (ptr *QVCandlestickModelMapper) ConnectLastSetColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLastSetColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVCandlestickModelMapper) DisconnectLastSetColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLastSetColumnChanged"})
}

func (ptr *QVCandlestickModelMapper) LastSetColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastSetColumnChanged"})
}

func (ptr *QVCandlestickModelMapper) LowRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowRow"}).(float64))
}

func (ptr *QVCandlestickModelMapper) ConnectLowRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLowRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVCandlestickModelMapper) DisconnectLowRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLowRowChanged"})
}

func (ptr *QVCandlestickModelMapper) LowRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowRowChanged"})
}

func (ptr *QVCandlestickModelMapper) OpenRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenRow"}).(float64))
}

func (ptr *QVCandlestickModelMapper) ConnectOpenRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpenRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVCandlestickModelMapper) DisconnectOpenRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpenRowChanged"})
}

func (ptr *QVCandlestickModelMapper) OpenRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenRowChanged"})
}

func (ptr *QVCandlestickModelMapper) ConnectOrientation(f func() core.Qt__Orientation) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOrientation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVCandlestickModelMapper) DisconnectOrientation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOrientation"})
}

func (ptr *QVCandlestickModelMapper) Orientation() core.Qt__Orientation {

	return core.Qt__Orientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Orientation"}).(float64))
}

func (ptr *QVCandlestickModelMapper) OrientationDefault() core.Qt__Orientation {

	return core.Qt__Orientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OrientationDefault"}).(float64))
}

func (ptr *QVCandlestickModelMapper) SetCloseRow(closeRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCloseRow", closeRow})
}

func (ptr *QVCandlestickModelMapper) SetFirstSetColumn(firstSetColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstSetColumn", firstSetColumn})
}

func (ptr *QVCandlestickModelMapper) SetHighRow(highRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHighRow", highRow})
}

func (ptr *QVCandlestickModelMapper) SetLastSetColumn(lastSetColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastSetColumn", lastSetColumn})
}

func (ptr *QVCandlestickModelMapper) SetLowRow(lowRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLowRow", lowRow})
}

func (ptr *QVCandlestickModelMapper) SetOpenRow(openRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpenRow", openRow})
}

func (ptr *QVCandlestickModelMapper) SetTimestampRow(timestampRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTimestampRow", timestampRow})
}

func (ptr *QVCandlestickModelMapper) TimestampRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimestampRow"}).(float64))
}

func (ptr *QVCandlestickModelMapper) ConnectTimestampRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTimestampRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVCandlestickModelMapper) DisconnectTimestampRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTimestampRowChanged"})
}

func (ptr *QVCandlestickModelMapper) TimestampRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimestampRowChanged"})
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

func (n *QVPieModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QPieModelMapper_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVPieModelMapper) ClassNameInternalF() string {
	return n.QPieModelMapper_PTR().ClassNameInternalF()
}

func NewQVPieModelMapperFromPointer(ptr unsafe.Pointer) (n *QVPieModelMapper) {
	n = new(QVPieModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QVPieModelMapper")
	return
}

func (ptr *QVPieModelMapper) DestroyQVPieModelMapper() {
}

func NewQVPieModelMapper(parent core.QObject_ITF) *QVPieModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQVPieModelMapper", "", parent}).(*QVPieModelMapper)
}

func (ptr *QVPieModelMapper) FirstRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstRow"}).(float64))
}

func (ptr *QVPieModelMapper) ConnectFirstRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVPieModelMapper) DisconnectFirstRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstRowChanged"})
}

func (ptr *QVPieModelMapper) FirstRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstRowChanged"})
}

func (ptr *QVPieModelMapper) LabelsColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsColumn"}).(float64))
}

func (ptr *QVPieModelMapper) ConnectLabelsColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVPieModelMapper) DisconnectLabelsColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsColumnChanged"})
}

func (ptr *QVPieModelMapper) LabelsColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsColumnChanged"})
}

func (ptr *QVPieModelMapper) Model() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*core.QAbstractItemModel)
}

func (ptr *QVPieModelMapper) ConnectModelReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModelReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVPieModelMapper) DisconnectModelReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModelReplaced"})
}

func (ptr *QVPieModelMapper) ModelReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModelReplaced"})
}

func (ptr *QVPieModelMapper) RowCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCount"}).(float64))
}

func (ptr *QVPieModelMapper) ConnectRowCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVPieModelMapper) DisconnectRowCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowCountChanged"})
}

func (ptr *QVPieModelMapper) RowCountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountChanged"})
}

func (ptr *QVPieModelMapper) Series() *QPieSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QPieSeries)
}

func (ptr *QVPieModelMapper) ConnectSeriesReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVPieModelMapper) DisconnectSeriesReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesReplaced"})
}

func (ptr *QVPieModelMapper) SeriesReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesReplaced"})
}

func (ptr *QVPieModelMapper) SetFirstRow(firstRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstRow", firstRow})
}

func (ptr *QVPieModelMapper) SetLabelsColumn(labelsColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelsColumn", labelsColumn})
}

func (ptr *QVPieModelMapper) SetModel(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModel", model})
}

func (ptr *QVPieModelMapper) SetRowCount(rowCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowCount", rowCount})
}

func (ptr *QVPieModelMapper) SetSeries(series QPieSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSeries", series})
}

func (ptr *QVPieModelMapper) SetValuesColumn(valuesColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValuesColumn", valuesColumn})
}

func (ptr *QVPieModelMapper) ValuesColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValuesColumn"}).(float64))
}

func (ptr *QVPieModelMapper) ConnectValuesColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValuesColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVPieModelMapper) DisconnectValuesColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValuesColumnChanged"})
}

func (ptr *QVPieModelMapper) ValuesColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValuesColumnChanged"})
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

func (n *QVXYModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QXYModelMapper_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVXYModelMapper) ClassNameInternalF() string {
	return n.QXYModelMapper_PTR().ClassNameInternalF()
}

func NewQVXYModelMapperFromPointer(ptr unsafe.Pointer) (n *QVXYModelMapper) {
	n = new(QVXYModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QVXYModelMapper")
	return
}

func (ptr *QVXYModelMapper) DestroyQVXYModelMapper() {
}

func NewQVXYModelMapper(parent core.QObject_ITF) *QVXYModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQVXYModelMapper", "", parent}).(*QVXYModelMapper)
}

func (ptr *QVXYModelMapper) FirstRow() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstRow"}).(float64))
}

func (ptr *QVXYModelMapper) ConnectFirstRowChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFirstRowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVXYModelMapper) DisconnectFirstRowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFirstRowChanged"})
}

func (ptr *QVXYModelMapper) FirstRowChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstRowChanged"})
}

func (ptr *QVXYModelMapper) Model() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*core.QAbstractItemModel)
}

func (ptr *QVXYModelMapper) ConnectModelReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModelReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVXYModelMapper) DisconnectModelReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModelReplaced"})
}

func (ptr *QVXYModelMapper) ModelReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModelReplaced"})
}

func (ptr *QVXYModelMapper) RowCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCount"}).(float64))
}

func (ptr *QVXYModelMapper) ConnectRowCountChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVXYModelMapper) DisconnectRowCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowCountChanged"})
}

func (ptr *QVXYModelMapper) RowCountChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountChanged"})
}

func (ptr *QVXYModelMapper) Series() *QXYSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QXYSeries)
}

func (ptr *QVXYModelMapper) ConnectSeriesReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVXYModelMapper) DisconnectSeriesReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesReplaced"})
}

func (ptr *QVXYModelMapper) SeriesReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesReplaced"})
}

func (ptr *QVXYModelMapper) SetFirstRow(firstRow int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstRow", firstRow})
}

func (ptr *QVXYModelMapper) SetModel(model core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModel", model})
}

func (ptr *QVXYModelMapper) SetRowCount(rowCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowCount", rowCount})
}

func (ptr *QVXYModelMapper) SetSeries(series QXYSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSeries", series})
}

func (ptr *QVXYModelMapper) SetXColumn(xColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXColumn", xColumn})
}

func (ptr *QVXYModelMapper) SetYColumn(yColumn int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYColumn", yColumn})
}

func (ptr *QVXYModelMapper) XColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XColumn"}).(float64))
}

func (ptr *QVXYModelMapper) ConnectXColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectXColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVXYModelMapper) DisconnectXColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectXColumnChanged"})
}

func (ptr *QVXYModelMapper) XColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XColumnChanged"})
}

func (ptr *QVXYModelMapper) YColumn() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YColumn"}).(float64))
}

func (ptr *QVXYModelMapper) ConnectYColumnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectYColumnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVXYModelMapper) DisconnectYColumnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectYColumnChanged"})
}

func (ptr *QVXYModelMapper) YColumnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YColumnChanged"})
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

func (n *QValueAxis) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractAxis_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QValueAxis) ClassNameInternalF() string {
	return n.QAbstractAxis_PTR().ClassNameInternalF()
}

func NewQValueAxisFromPointer(ptr unsafe.Pointer) (n *QValueAxis) {
	n = new(QValueAxis)
	n.InitFromInternal(uintptr(ptr), "charts.QValueAxis")
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

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQValueAxis", "", parent}).(*QValueAxis)
}

func (ptr *QValueAxis) ConnectApplyNiceNumbers(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectApplyNiceNumbers", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectApplyNiceNumbers() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectApplyNiceNumbers"})
}

func (ptr *QValueAxis) ApplyNiceNumbers() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ApplyNiceNumbers"})
}

func (ptr *QValueAxis) ApplyNiceNumbersDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ApplyNiceNumbersDefault"})
}

func (ptr *QValueAxis) LabelFormat() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelFormat"}).(string)
}

func (ptr *QValueAxis) ConnectLabelFormatChanged(f func(format string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelFormatChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectLabelFormatChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelFormatChanged"})
}

func (ptr *QValueAxis) LabelFormatChanged(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelFormatChanged", format})
}

func (ptr *QValueAxis) Max() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Max"}).(float64)
}

func (ptr *QValueAxis) ConnectMaxChanged(f func(max float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMaxChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectMaxChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMaxChanged"})
}

func (ptr *QValueAxis) MaxChanged(max float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxChanged", max})
}

func (ptr *QValueAxis) Min() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Min"}).(float64)
}

func (ptr *QValueAxis) ConnectMinChanged(f func(min float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectMinChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinChanged"})
}

func (ptr *QValueAxis) MinChanged(min float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinChanged", min})
}

func (ptr *QValueAxis) MinorTickCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinorTickCount"}).(float64))
}

func (ptr *QValueAxis) ConnectMinorTickCountChanged(f func(minorTickCount int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinorTickCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectMinorTickCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinorTickCountChanged"})
}

func (ptr *QValueAxis) MinorTickCountChanged(minorTickCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinorTickCountChanged", minorTickCount})
}

func (ptr *QValueAxis) ConnectRangeChanged(f func(min float64, max float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRangeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectRangeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRangeChanged"})
}

func (ptr *QValueAxis) RangeChanged(min float64, max float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RangeChanged", min, max})
}

func (ptr *QValueAxis) SetLabelFormat(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelFormat", format})
}

func (ptr *QValueAxis) SetMax(max float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMax", max})
}

func (ptr *QValueAxis) SetMin(min float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMin", min})
}

func (ptr *QValueAxis) SetMinorTickCount(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinorTickCount", count})
}

func (ptr *QValueAxis) SetRange(min float64, max float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRange", min, max})
}

func (ptr *QValueAxis) SetTickAnchor(anchor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTickAnchor", anchor})
}

func (ptr *QValueAxis) SetTickCount(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTickCount", count})
}

func (ptr *QValueAxis) SetTickInterval(insterval float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTickInterval", insterval})
}

func (ptr *QValueAxis) SetTickType(ty QValueAxis__TickType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTickType", ty})
}

func (ptr *QValueAxis) TickAnchor() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickAnchor"}).(float64)
}

func (ptr *QValueAxis) ConnectTickAnchorChanged(f func(anchor float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTickAnchorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectTickAnchorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTickAnchorChanged"})
}

func (ptr *QValueAxis) TickAnchorChanged(anchor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickAnchorChanged", anchor})
}

func (ptr *QValueAxis) TickCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickCount"}).(float64))
}

func (ptr *QValueAxis) ConnectTickCountChanged(f func(tickCount int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTickCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectTickCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTickCountChanged"})
}

func (ptr *QValueAxis) TickCountChanged(tickCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickCountChanged", tickCount})
}

func (ptr *QValueAxis) TickInterval() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickInterval"}).(float64)
}

func (ptr *QValueAxis) ConnectTickIntervalChanged(f func(interval float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTickIntervalChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectTickIntervalChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTickIntervalChanged"})
}

func (ptr *QValueAxis) TickIntervalChanged(interval float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickIntervalChanged", interval})
}

func (ptr *QValueAxis) TickType() QValueAxis__TickType {

	return QValueAxis__TickType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickType"}).(float64))
}

func (ptr *QValueAxis) ConnectTickTypeChanged(f func(ty QValueAxis__TickType)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTickTypeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectTickTypeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTickTypeChanged"})
}

func (ptr *QValueAxis) TickTypeChanged(ty QValueAxis__TickType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TickTypeChanged", ty})
}

func (ptr *QValueAxis) ConnectType(f func() QAbstractAxis__AxisType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QValueAxis) Type() QAbstractAxis__AxisType {

	return QAbstractAxis__AxisType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QValueAxis) TypeDefault() QAbstractAxis__AxisType {

	return QAbstractAxis__AxisType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QValueAxis) ConnectDestroyQValueAxis(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQValueAxis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValueAxis) DisconnectDestroyQValueAxis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQValueAxis"})
}

func (ptr *QValueAxis) DestroyQValueAxis() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQValueAxis"})
}

func (ptr *QValueAxis) DestroyQValueAxisDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQValueAxisDefault"})
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

func (n *QXYLegendMarker) InitFromInternal(ptr uintptr, name string) {
	n.QLegendMarker_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QXYLegendMarker) ClassNameInternalF() string {
	return n.QLegendMarker_PTR().ClassNameInternalF()
}

func NewQXYLegendMarkerFromPointer(ptr unsafe.Pointer) (n *QXYLegendMarker) {
	n = new(QXYLegendMarker)
	n.InitFromInternal(uintptr(ptr), "charts.QXYLegendMarker")
	return
}
func (ptr *QXYLegendMarker) ConnectSeries(f func() *QXYSeries) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYLegendMarker) DisconnectSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeries"})
}

func (ptr *QXYLegendMarker) Series() *QXYSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QXYSeries)
}

func (ptr *QXYLegendMarker) SeriesDefault() *QXYSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesDefault"}).(*QXYSeries)
}

func (ptr *QXYLegendMarker) ConnectType(f func() QLegendMarker__LegendMarkerType) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectType", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYLegendMarker) DisconnectType() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectType"})
}

func (ptr *QXYLegendMarker) Type() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QXYLegendMarker) TypeDefault() QLegendMarker__LegendMarkerType {

	return QLegendMarker__LegendMarkerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QXYLegendMarker) ConnectDestroyQXYLegendMarker(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXYLegendMarker", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYLegendMarker) DisconnectDestroyQXYLegendMarker() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXYLegendMarker"})
}

func (ptr *QXYLegendMarker) DestroyQXYLegendMarker() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXYLegendMarker"})
}

func (ptr *QXYLegendMarker) DestroyQXYLegendMarkerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXYLegendMarkerDefault"})
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

func (n *QXYModelMapper) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QXYModelMapper) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQXYModelMapperFromPointer(ptr unsafe.Pointer) (n *QXYModelMapper) {
	n = new(QXYModelMapper)
	n.InitFromInternal(uintptr(ptr), "charts.QXYModelMapper")
	return
}

func (ptr *QXYModelMapper) DestroyQXYModelMapper() {
}

func NewQXYModelMapper(parent core.QObject_ITF) *QXYModelMapper {

	return internal.CallLocalFunction([]interface{}{"", "", "charts.NewQXYModelMapper", "", parent}).(*QXYModelMapper)
}

func (ptr *QXYModelMapper) Orientation() core.Qt__Orientation {

	return core.Qt__Orientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Orientation"}).(float64))
}

func (ptr *QXYModelMapper) SetOrientation(orientation core.Qt__Orientation) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOrientation", orientation})
}

func (ptr *QXYModelMapper) SetXSection(xSection int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXSection", xSection})
}

func (ptr *QXYModelMapper) SetYSection(ySection int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYSection", ySection})
}

func (ptr *QXYModelMapper) XSection() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XSection"}).(float64))
}

func (ptr *QXYModelMapper) YSection() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YSection"}).(float64))
}

func (ptr *QXYModelMapper) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QXYModelMapper) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QXYModelMapper) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QXYModelMapper) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QXYModelMapper) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QXYModelMapper) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QXYModelMapper) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QXYModelMapper) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QXYModelMapper) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QXYModelMapper) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QXYModelMapper) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QXYModelMapper) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QXYModelMapper) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QXYModelMapper) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QXYModelMapper) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QXYModelMapper) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QXYModelMapper) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QXYModelMapper) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QXYModelMapper) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QXYModelMapper) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QXYModelMapper) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QXYSeries) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractSeries_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QXYSeries) ClassNameInternalF() string {
	return n.QAbstractSeries_PTR().ClassNameInternalF()
}

func NewQXYSeriesFromPointer(ptr unsafe.Pointer) (n *QXYSeries) {
	n = new(QXYSeries)
	n.InitFromInternal(uintptr(ptr), "charts.QXYSeries")
	return
}
func (ptr *QXYSeries) Append(x float64, y float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", x, y})
}

func (ptr *QXYSeries) Append2(point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append2", point})
}

func (ptr *QXYSeries) Append3(points []*core.QPointF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append3", points})
}

func (ptr *QXYSeries) At(index int) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "At", index}).(*core.QPointF)
}

func (ptr *QXYSeries) Brush() *gui.QBrush {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Brush"}).(*gui.QBrush)
}

func (ptr *QXYSeries) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QXYSeries) ConnectClicked(f func(point *core.QPointF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClicked"})
}

func (ptr *QXYSeries) Clicked(point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clicked", point})
}

func (ptr *QXYSeries) ConnectColor(f func() *gui.QColor) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectColor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColor"})
}

func (ptr *QXYSeries) Color() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Color"}).(*gui.QColor)
}

func (ptr *QXYSeries) ColorDefault() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorDefault"}).(*gui.QColor)
}

func (ptr *QXYSeries) ConnectColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColorChanged"})
}

func (ptr *QXYSeries) ColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorChanged", color})
}

func (ptr *QXYSeries) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QXYSeries) ConnectDoubleClicked(f func(point *core.QPointF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDoubleClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectDoubleClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDoubleClicked"})
}

func (ptr *QXYSeries) DoubleClicked(point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoubleClicked", point})
}

func (ptr *QXYSeries) ConnectHovered(f func(point *core.QPointF, state bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectHovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHovered"})
}

func (ptr *QXYSeries) Hovered(point core.QPointF_ITF, state bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hovered", point, state})
}

func (ptr *QXYSeries) Insert(index int, point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", index, point})
}

func (ptr *QXYSeries) Pen() *gui.QPen {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pen"}).(*gui.QPen)
}

func (ptr *QXYSeries) ConnectPenChanged(f func(pen *gui.QPen)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPenChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPenChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPenChanged"})
}

func (ptr *QXYSeries) PenChanged(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PenChanged", pen})
}

func (ptr *QXYSeries) ConnectPointAdded(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPointAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointAdded"})
}

func (ptr *QXYSeries) PointAdded(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointAdded", index})
}

func (ptr *QXYSeries) PointLabelsClipping() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsClipping"}).(bool)
}

func (ptr *QXYSeries) ConnectPointLabelsClippingChanged(f func(clipping bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointLabelsClippingChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPointLabelsClippingChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointLabelsClippingChanged"})
}

func (ptr *QXYSeries) PointLabelsClippingChanged(clipping bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsClippingChanged", clipping})
}

func (ptr *QXYSeries) PointLabelsColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsColor"}).(*gui.QColor)
}

func (ptr *QXYSeries) ConnectPointLabelsColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointLabelsColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPointLabelsColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointLabelsColorChanged"})
}

func (ptr *QXYSeries) PointLabelsColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsColorChanged", color})
}

func (ptr *QXYSeries) PointLabelsFont() *gui.QFont {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsFont"}).(*gui.QFont)
}

func (ptr *QXYSeries) ConnectPointLabelsFontChanged(f func(font *gui.QFont)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointLabelsFontChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPointLabelsFontChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointLabelsFontChanged"})
}

func (ptr *QXYSeries) PointLabelsFontChanged(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsFontChanged", font})
}

func (ptr *QXYSeries) PointLabelsFormat() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsFormat"}).(string)
}

func (ptr *QXYSeries) ConnectPointLabelsFormatChanged(f func(format string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointLabelsFormatChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPointLabelsFormatChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointLabelsFormatChanged"})
}

func (ptr *QXYSeries) PointLabelsFormatChanged(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsFormatChanged", format})
}

func (ptr *QXYSeries) ConnectPointLabelsVisibilityChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointLabelsVisibilityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPointLabelsVisibilityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointLabelsVisibilityChanged"})
}

func (ptr *QXYSeries) PointLabelsVisibilityChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsVisibilityChanged", visible})
}

func (ptr *QXYSeries) PointLabelsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointLabelsVisible"}).(bool)
}

func (ptr *QXYSeries) ConnectPointRemoved(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPointRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointRemoved"})
}

func (ptr *QXYSeries) PointRemoved(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointRemoved", index})
}

func (ptr *QXYSeries) ConnectPointReplaced(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPointReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointReplaced"})
}

func (ptr *QXYSeries) PointReplaced(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointReplaced", index})
}

func (ptr *QXYSeries) Points() []*core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Points"}).([]*core.QPointF)
}

func (ptr *QXYSeries) ConnectPointsRemoved(f func(index int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointsRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPointsRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointsRemoved"})
}

func (ptr *QXYSeries) PointsRemoved(index int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointsRemoved", index, count})
}

func (ptr *QXYSeries) ConnectPointsReplaced(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPointsReplaced", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPointsReplaced() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPointsReplaced"})
}

func (ptr *QXYSeries) PointsReplaced() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointsReplaced"})
}

func (ptr *QXYSeries) PointsVector() []*core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointsVector"}).([]*core.QPointF)
}

func (ptr *QXYSeries) PointsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointsVisible"}).(bool)
}

func (ptr *QXYSeries) ConnectPressed(f func(point *core.QPointF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPressed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectPressed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPressed"})
}

func (ptr *QXYSeries) Pressed(point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressed", point})
}

func (ptr *QXYSeries) ConnectReleased(f func(point *core.QPointF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReleased", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectReleased() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReleased"})
}

func (ptr *QXYSeries) Released(point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Released", point})
}

func (ptr *QXYSeries) Remove(x float64, y float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", x, y})
}

func (ptr *QXYSeries) Remove2(point core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove2", point})
}

func (ptr *QXYSeries) Remove3(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove3", index})
}

func (ptr *QXYSeries) RemovePoints(index int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemovePoints", index, count})
}

func (ptr *QXYSeries) Replace(oldX float64, oldY float64, newX float64, newY float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Replace", oldX, oldY, newX, newY})
}

func (ptr *QXYSeries) Replace2(oldPoint core.QPointF_ITF, newPoint core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Replace2", oldPoint, newPoint})
}

func (ptr *QXYSeries) Replace3(index int, newX float64, newY float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Replace3", index, newX, newY})
}

func (ptr *QXYSeries) Replace4(index int, newPoint core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Replace4", index, newPoint})
}

func (ptr *QXYSeries) Replace5(points []*core.QPointF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Replace5", points})
}

func (ptr *QXYSeries) Replace6(points []*core.QPointF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Replace6", points})
}

func (ptr *QXYSeries) ConnectSetBrush(f func(brush *gui.QBrush)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetBrush", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectSetBrush() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetBrush"})
}

func (ptr *QXYSeries) SetBrush(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrush", brush})
}

func (ptr *QXYSeries) SetBrushDefault(brush gui.QBrush_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBrushDefault", brush})
}

func (ptr *QXYSeries) ConnectSetColor(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetColor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectSetColor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetColor"})
}

func (ptr *QXYSeries) SetColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColor", color})
}

func (ptr *QXYSeries) SetColorDefault(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColorDefault", color})
}

func (ptr *QXYSeries) ConnectSetPen(f func(pen *gui.QPen)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetPen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectSetPen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetPen"})
}

func (ptr *QXYSeries) SetPen(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPen", pen})
}

func (ptr *QXYSeries) SetPenDefault(pen gui.QPen_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPenDefault", pen})
}

func (ptr *QXYSeries) SetPointLabelsClipping(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointLabelsClipping", enabled})
}

func (ptr *QXYSeries) SetPointLabelsColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointLabelsColor", color})
}

func (ptr *QXYSeries) SetPointLabelsFont(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointLabelsFont", font})
}

func (ptr *QXYSeries) SetPointLabelsFormat(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointLabelsFormat", format})
}

func (ptr *QXYSeries) SetPointLabelsVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointLabelsVisible", visible})
}

func (ptr *QXYSeries) SetPointsVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPointsVisible", visible})
}

func (ptr *QXYSeries) ConnectDestroyQXYSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXYSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXYSeries) DisconnectDestroyQXYSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXYSeries"})
}

func (ptr *QXYSeries) DestroyQXYSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXYSeries"})
}

func (ptr *QXYSeries) DestroyQXYSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXYSeriesDefault"})
}

func (ptr *QXYSeries) __append_points_atList3(i int) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_points_atList3", i}).(*core.QPointF)
}

func (ptr *QXYSeries) __append_points_setList3(i core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_points_setList3", i})
}

func (ptr *QXYSeries) __append_points_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__append_points_newList3"}).(unsafe.Pointer)
}

func (ptr *QXYSeries) __points_atList(i int) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__points_atList", i}).(*core.QPointF)
}

func (ptr *QXYSeries) __points_setList(i core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__points_setList", i})
}

func (ptr *QXYSeries) __points_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__points_newList"}).(unsafe.Pointer)
}

func (ptr *QXYSeries) __pointsVector_atList(i int) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__pointsVector_atList", i}).(*core.QPointF)
}

func (ptr *QXYSeries) __pointsVector_setList(i core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__pointsVector_setList", i})
}

func (ptr *QXYSeries) __pointsVector_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__pointsVector_newList"}).(unsafe.Pointer)
}

func (ptr *QXYSeries) __replace_points_atList5(i int) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__replace_points_atList5", i}).(*core.QPointF)
}

func (ptr *QXYSeries) __replace_points_setList5(i core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__replace_points_setList5", i})
}

func (ptr *QXYSeries) __replace_points_newList5() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__replace_points_newList5"}).(unsafe.Pointer)
}

func (ptr *QXYSeries) __replace_points_atList6(i int) *core.QPointF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__replace_points_atList6", i}).(*core.QPointF)
}

func (ptr *QXYSeries) __replace_points_setList6(i core.QPointF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__replace_points_setList6", i})
}

func (ptr *QXYSeries) __replace_points_newList6() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__replace_points_newList6"}).(unsafe.Pointer)
}

func (ptr *QXYSeries) Type() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QXYSeries) TypeDefault() QAbstractSeries__SeriesType {

	return QAbstractSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func init() {
	internal.ConstructorTable["charts.QAbstractAxis"] = NewQAbstractAxisFromPointer
	internal.ConstructorTable["charts.QAbstractBarSeries"] = NewQAbstractBarSeriesFromPointer
	internal.ConstructorTable["charts.QAbstractSeries"] = NewQAbstractSeriesFromPointer
	internal.ConstructorTable["charts.QAreaLegendMarker"] = NewQAreaLegendMarkerFromPointer
	internal.ConstructorTable["charts.QAreaSeries"] = NewQAreaSeriesFromPointer
	internal.ConstructorTable["charts.QBarCategoryAxis"] = NewQBarCategoryAxisFromPointer
	internal.ConstructorTable["charts.QBarLegendMarker"] = NewQBarLegendMarkerFromPointer
	internal.ConstructorTable["charts.QBarSeries"] = NewQBarSeriesFromPointer
	internal.ConstructorTable["charts.QBarSet"] = NewQBarSetFromPointer
	internal.ConstructorTable["charts.QBoxPlotLegendMarker"] = NewQBoxPlotLegendMarkerFromPointer
	internal.ConstructorTable["charts.QBoxPlotSeries"] = NewQBoxPlotSeriesFromPointer
	internal.ConstructorTable["charts.QBoxSet"] = NewQBoxSetFromPointer
	internal.ConstructorTable["charts.QCandlestickLegendMarker"] = NewQCandlestickLegendMarkerFromPointer
	internal.ConstructorTable["charts.QCandlestickModelMapper"] = NewQCandlestickModelMapperFromPointer
	internal.ConstructorTable["charts.QCandlestickSeries"] = NewQCandlestickSeriesFromPointer
	internal.ConstructorTable["charts.QCandlestickSet"] = NewQCandlestickSetFromPointer
	internal.ConstructorTable["charts.QCategoryAxis"] = NewQCategoryAxisFromPointer
	internal.ConstructorTable["charts.QChart"] = NewQChartFromPointer
	internal.ConstructorTable["charts.QChartView"] = NewQChartViewFromPointer
	internal.ConstructorTable["charts.QDateTimeAxis"] = NewQDateTimeAxisFromPointer
	internal.ConstructorTable["charts.QHBarModelMapper"] = NewQHBarModelMapperFromPointer
	internal.ConstructorTable["charts.QHBoxPlotModelMapper"] = NewQHBoxPlotModelMapperFromPointer
	internal.ConstructorTable["charts.QHCandlestickModelMapper"] = NewQHCandlestickModelMapperFromPointer
	internal.ConstructorTable["charts.QHPieModelMapper"] = NewQHPieModelMapperFromPointer
	internal.ConstructorTable["charts.QHXYModelMapper"] = NewQHXYModelMapperFromPointer
	internal.ConstructorTable["charts.QHorizontalBarSeries"] = NewQHorizontalBarSeriesFromPointer
	internal.ConstructorTable["charts.QHorizontalPercentBarSeries"] = NewQHorizontalPercentBarSeriesFromPointer
	internal.ConstructorTable["charts.QHorizontalStackedBarSeries"] = NewQHorizontalStackedBarSeriesFromPointer
	internal.ConstructorTable["charts.QLegend"] = NewQLegendFromPointer
	internal.ConstructorTable["charts.QLegendMarker"] = NewQLegendMarkerFromPointer
	internal.ConstructorTable["charts.QLineSeries"] = NewQLineSeriesFromPointer
	internal.ConstructorTable["charts.QLogValueAxis"] = NewQLogValueAxisFromPointer
	internal.ConstructorTable["charts.QPercentBarSeries"] = NewQPercentBarSeriesFromPointer
	internal.ConstructorTable["charts.QPieLegendMarker"] = NewQPieLegendMarkerFromPointer
	internal.ConstructorTable["charts.QPieModelMapper"] = NewQPieModelMapperFromPointer
	internal.ConstructorTable["charts.QPieSeries"] = NewQPieSeriesFromPointer
	internal.ConstructorTable["charts.QPieSlice"] = NewQPieSliceFromPointer
	internal.ConstructorTable["charts.QPolarChart"] = NewQPolarChartFromPointer
	internal.ConstructorTable["charts.QScatterSeries"] = NewQScatterSeriesFromPointer
	internal.ConstructorTable["charts.QSplineSeries"] = NewQSplineSeriesFromPointer
	internal.ConstructorTable["charts.QStackedBarSeries"] = NewQStackedBarSeriesFromPointer
	internal.ConstructorTable["charts.QVBarModelMapper"] = NewQVBarModelMapperFromPointer
	internal.ConstructorTable["charts.QVBoxPlotModelMapper"] = NewQVBoxPlotModelMapperFromPointer
	internal.ConstructorTable["charts.QVCandlestickModelMapper"] = NewQVCandlestickModelMapperFromPointer
	internal.ConstructorTable["charts.QVPieModelMapper"] = NewQVPieModelMapperFromPointer
	internal.ConstructorTable["charts.QVXYModelMapper"] = NewQVXYModelMapperFromPointer
	internal.ConstructorTable["charts.QValueAxis"] = NewQValueAxisFromPointer
	internal.ConstructorTable["charts.QXYLegendMarker"] = NewQXYLegendMarkerFromPointer
	internal.ConstructorTable["charts.QXYModelMapper"] = NewQXYModelMapperFromPointer
	internal.ConstructorTable["charts.QXYSeries"] = NewQXYSeriesFromPointer
}
