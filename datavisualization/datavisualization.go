// +build !minimal

package datavisualization

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/internal"
	"unsafe"
)

type Q3DBars struct {
	internal.Internal
}

type Q3DBars_ITF interface {
	Q3DBars_PTR() *Q3DBars
}

func (ptr *Q3DBars) Q3DBars_PTR() *Q3DBars {
	return ptr
}

func (ptr *Q3DBars) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *Q3DBars) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQ3DBars(ptr Q3DBars_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DBars_PTR().Pointer()
	}
	return nil
}

func (n *Q3DBars) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQ3DBarsFromPointer(ptr unsafe.Pointer) (n *Q3DBars) {
	n = new(Q3DBars)
	n.InitFromInternal(uintptr(ptr), "datavisualization.Q3DBars")
	return
}
func NewQ3DBars(format gui.QSurfaceFormat_ITF, parent gui.QWindow_ITF) *Q3DBars {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQ3DBars", "", format, parent}).(*Q3DBars)
}

func (ptr *Q3DBars) AddAxis(axis QAbstract3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddAxis", axis})
}

func (ptr *Q3DBars) AddSeries(series QBar3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddSeries", series})
}

func (ptr *Q3DBars) Axes() []*QAbstract3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Axes"}).([]*QAbstract3DAxis)
}

func (ptr *Q3DBars) BarSpacing() *core.QSizeF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BarSpacing"}).(*core.QSizeF)
}

func (ptr *Q3DBars) ConnectBarSpacingChanged(f func(spacing *core.QSizeF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBarSpacingChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectBarSpacingChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBarSpacingChanged"})
}

func (ptr *Q3DBars) BarSpacingChanged(spacing core.QSizeF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BarSpacingChanged", spacing})
}

func (ptr *Q3DBars) ConnectBarSpacingRelativeChanged(f func(relative bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBarSpacingRelativeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectBarSpacingRelativeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBarSpacingRelativeChanged"})
}

func (ptr *Q3DBars) BarSpacingRelativeChanged(relative bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BarSpacingRelativeChanged", relative})
}

func (ptr *Q3DBars) BarThickness() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BarThickness"}).(float32)
}

func (ptr *Q3DBars) ConnectBarThicknessChanged(f func(thicknessRatio float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBarThicknessChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectBarThicknessChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBarThicknessChanged"})
}

func (ptr *Q3DBars) BarThicknessChanged(thicknessRatio float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BarThicknessChanged", thicknessRatio})
}

func (ptr *Q3DBars) ColumnAxis() *QCategory3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnAxis"}).(*QCategory3DAxis)
}

func (ptr *Q3DBars) ConnectColumnAxisChanged(f func(axis *QCategory3DAxis)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnAxisChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectColumnAxisChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnAxisChanged"})
}

func (ptr *Q3DBars) ColumnAxisChanged(axis QCategory3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnAxisChanged", axis})
}

func (ptr *Q3DBars) FloorLevel() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FloorLevel"}).(float32)
}

func (ptr *Q3DBars) ConnectFloorLevelChanged(f func(level float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFloorLevelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectFloorLevelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFloorLevelChanged"})
}

func (ptr *Q3DBars) FloorLevelChanged(level float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FloorLevelChanged", level})
}

func (ptr *Q3DBars) InsertSeries(index int, series QBar3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertSeries", index, series})
}

func (ptr *Q3DBars) IsBarSpacingRelative() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBarSpacingRelative"}).(bool)
}

func (ptr *Q3DBars) IsMultiSeriesUniform() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsMultiSeriesUniform"}).(bool)
}

func (ptr *Q3DBars) ConnectMultiSeriesUniformChanged(f func(uniform bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMultiSeriesUniformChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectMultiSeriesUniformChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMultiSeriesUniformChanged"})
}

func (ptr *Q3DBars) MultiSeriesUniformChanged(uniform bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiSeriesUniformChanged", uniform})
}

func (ptr *Q3DBars) PrimarySeries() *QBar3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrimarySeries"}).(*QBar3DSeries)
}

func (ptr *Q3DBars) ConnectPrimarySeriesChanged(f func(series *QBar3DSeries)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrimarySeriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectPrimarySeriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrimarySeriesChanged"})
}

func (ptr *Q3DBars) PrimarySeriesChanged(series QBar3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrimarySeriesChanged", series})
}

func (ptr *Q3DBars) ReleaseAxis(axis QAbstract3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReleaseAxis", axis})
}

func (ptr *Q3DBars) RemoveSeries(series QBar3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveSeries", series})
}

func (ptr *Q3DBars) RowAxis() *QCategory3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowAxis"}).(*QCategory3DAxis)
}

func (ptr *Q3DBars) ConnectRowAxisChanged(f func(axis *QCategory3DAxis)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowAxisChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectRowAxisChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowAxisChanged"})
}

func (ptr *Q3DBars) RowAxisChanged(axis QCategory3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowAxisChanged", axis})
}

func (ptr *Q3DBars) SelectedSeries() *QBar3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedSeries"}).(*QBar3DSeries)
}

func (ptr *Q3DBars) ConnectSelectedSeriesChanged(f func(series *QBar3DSeries)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectedSeriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectSelectedSeriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectedSeriesChanged"})
}

func (ptr *Q3DBars) SelectedSeriesChanged(series QBar3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedSeriesChanged", series})
}

func (ptr *Q3DBars) SeriesList() []*QBar3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesList"}).([]*QBar3DSeries)
}

func (ptr *Q3DBars) SetBarSpacing(spacing core.QSizeF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBarSpacing", spacing})
}

func (ptr *Q3DBars) SetBarSpacingRelative(relative bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBarSpacingRelative", relative})
}

func (ptr *Q3DBars) SetBarThickness(thicknessRatio float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBarThickness", thicknessRatio})
}

func (ptr *Q3DBars) SetColumnAxis(axis QCategory3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnAxis", axis})
}

func (ptr *Q3DBars) SetFloorLevel(level float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFloorLevel", level})
}

func (ptr *Q3DBars) SetMultiSeriesUniform(uniform bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMultiSeriesUniform", uniform})
}

func (ptr *Q3DBars) SetPrimarySeries(series QBar3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrimarySeries", series})
}

func (ptr *Q3DBars) SetRowAxis(axis QCategory3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowAxis", axis})
}

func (ptr *Q3DBars) SetValueAxis(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValueAxis", axis})
}

func (ptr *Q3DBars) ValueAxis() *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueAxis"}).(*QValue3DAxis)
}

func (ptr *Q3DBars) ConnectValueAxisChanged(f func(axis *QValue3DAxis)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValueAxisChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectValueAxisChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValueAxisChanged"})
}

func (ptr *Q3DBars) ValueAxisChanged(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueAxisChanged", axis})
}

func (ptr *Q3DBars) ConnectDestroyQ3DBars(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQ3DBars", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DBars) DisconnectDestroyQ3DBars() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQ3DBars"})
}

func (ptr *Q3DBars) DestroyQ3DBars() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DBars"})
}

func (ptr *Q3DBars) DestroyQ3DBarsDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DBarsDefault"})
}

func (ptr *Q3DBars) __axes_atList(i int) *QAbstract3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_atList", i}).(*QAbstract3DAxis)
}

func (ptr *Q3DBars) __axes_setList(i QAbstract3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_setList", i})
}

func (ptr *Q3DBars) __axes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_newList"}).(unsafe.Pointer)
}

func (ptr *Q3DBars) __seriesList_atList(i int) *QBar3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__seriesList_atList", i}).(*QBar3DSeries)
}

func (ptr *Q3DBars) __seriesList_setList(i QBar3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__seriesList_setList", i})
}

func (ptr *Q3DBars) __seriesList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__seriesList_newList"}).(unsafe.Pointer)
}

type Q3DCamera struct {
	internal.Internal
}

type Q3DCamera_ITF interface {
	Q3DCamera_PTR() *Q3DCamera
}

func (ptr *Q3DCamera) Q3DCamera_PTR() *Q3DCamera {
	return ptr
}

func (ptr *Q3DCamera) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *Q3DCamera) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQ3DCamera(ptr Q3DCamera_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DCamera_PTR().Pointer()
	}
	return nil
}

func (n *Q3DCamera) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQ3DCameraFromPointer(ptr unsafe.Pointer) (n *Q3DCamera) {
	n = new(Q3DCamera)
	n.InitFromInternal(uintptr(ptr), "datavisualization.Q3DCamera")
	return
}

//go:generate stringer -type=Q3DCamera__CameraPreset
//Q3DCamera::CameraPreset
type Q3DCamera__CameraPreset int64

const (
	Q3DCamera__CameraPresetNone               Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(-1)
	Q3DCamera__CameraPresetFrontLow           Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(0)
	Q3DCamera__CameraPresetFront              Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(1)
	Q3DCamera__CameraPresetFrontHigh          Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(2)
	Q3DCamera__CameraPresetLeftLow            Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(3)
	Q3DCamera__CameraPresetLeft               Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(4)
	Q3DCamera__CameraPresetLeftHigh           Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(5)
	Q3DCamera__CameraPresetRightLow           Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(6)
	Q3DCamera__CameraPresetRight              Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(7)
	Q3DCamera__CameraPresetRightHigh          Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(8)
	Q3DCamera__CameraPresetBehindLow          Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(9)
	Q3DCamera__CameraPresetBehind             Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(10)
	Q3DCamera__CameraPresetBehindHigh         Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(11)
	Q3DCamera__CameraPresetIsometricLeft      Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(12)
	Q3DCamera__CameraPresetIsometricLeftHigh  Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(13)
	Q3DCamera__CameraPresetIsometricRight     Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(14)
	Q3DCamera__CameraPresetIsometricRightHigh Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(15)
	Q3DCamera__CameraPresetDirectlyAbove      Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(16)
	Q3DCamera__CameraPresetDirectlyAboveCW45  Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(17)
	Q3DCamera__CameraPresetDirectlyAboveCCW45 Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(18)
	Q3DCamera__CameraPresetFrontBelow         Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(19)
	Q3DCamera__CameraPresetLeftBelow          Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(20)
	Q3DCamera__CameraPresetRightBelow         Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(21)
	Q3DCamera__CameraPresetBehindBelow        Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(22)
	Q3DCamera__CameraPresetDirectlyBelow      Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(23)
)

func NewQ3DCamera(parent core.QObject_ITF) *Q3DCamera {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQ3DCamera", "", parent}).(*Q3DCamera)
}

func (ptr *Q3DCamera) CameraPreset() Q3DCamera__CameraPreset {

	return Q3DCamera__CameraPreset(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CameraPreset"}).(float64))
}

func (ptr *Q3DCamera) ConnectCameraPresetChanged(f func(preset Q3DCamera__CameraPreset)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCameraPresetChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectCameraPresetChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCameraPresetChanged"})
}

func (ptr *Q3DCamera) CameraPresetChanged(preset Q3DCamera__CameraPreset) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CameraPresetChanged", preset})
}

func (ptr *Q3DCamera) ConnectCopyValuesFrom(f func(source *Q3DObject)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCopyValuesFrom", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectCopyValuesFrom() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCopyValuesFrom"})
}

func (ptr *Q3DCamera) CopyValuesFrom(source Q3DObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CopyValuesFrom", source})
}

func (ptr *Q3DCamera) CopyValuesFromDefault(source Q3DObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CopyValuesFromDefault", source})
}

func (ptr *Q3DCamera) MaxZoomLevel() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxZoomLevel"}).(float32)
}

func (ptr *Q3DCamera) ConnectMaxZoomLevelChanged(f func(zoomLevel float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMaxZoomLevelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectMaxZoomLevelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMaxZoomLevelChanged"})
}

func (ptr *Q3DCamera) MaxZoomLevelChanged(zoomLevel float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxZoomLevelChanged", zoomLevel})
}

func (ptr *Q3DCamera) MinZoomLevel() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinZoomLevel"}).(float32)
}

func (ptr *Q3DCamera) ConnectMinZoomLevelChanged(f func(zoomLevel float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinZoomLevelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectMinZoomLevelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinZoomLevelChanged"})
}

func (ptr *Q3DCamera) MinZoomLevelChanged(zoomLevel float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinZoomLevelChanged", zoomLevel})
}

func (ptr *Q3DCamera) SetCameraPosition(horizontal float32, vertical float32, zoom float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCameraPosition", horizontal, vertical, zoom})
}

func (ptr *Q3DCamera) SetCameraPreset(preset Q3DCamera__CameraPreset) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCameraPreset", preset})
}

func (ptr *Q3DCamera) SetMaxZoomLevel(zoomLevel float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaxZoomLevel", zoomLevel})
}

func (ptr *Q3DCamera) SetMinZoomLevel(zoomLevel float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinZoomLevel", zoomLevel})
}

func (ptr *Q3DCamera) SetTarget(target gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTarget", target})
}

func (ptr *Q3DCamera) SetWrapXRotation(isEnabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWrapXRotation", isEnabled})
}

func (ptr *Q3DCamera) SetWrapYRotation(isEnabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWrapYRotation", isEnabled})
}

func (ptr *Q3DCamera) SetXRotation(rotation float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXRotation", rotation})
}

func (ptr *Q3DCamera) SetYRotation(rotation float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYRotation", rotation})
}

func (ptr *Q3DCamera) SetZoomLevel(zoomLevel float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZoomLevel", zoomLevel})
}

func (ptr *Q3DCamera) Target() *gui.QVector3D {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Target"}).(*gui.QVector3D)
}

func (ptr *Q3DCamera) ConnectTargetChanged(f func(target *gui.QVector3D)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTargetChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectTargetChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTargetChanged"})
}

func (ptr *Q3DCamera) TargetChanged(target gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TargetChanged", target})
}

func (ptr *Q3DCamera) WrapXRotation() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WrapXRotation"}).(bool)
}

func (ptr *Q3DCamera) ConnectWrapXRotationChanged(f func(isEnabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWrapXRotationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectWrapXRotationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWrapXRotationChanged"})
}

func (ptr *Q3DCamera) WrapXRotationChanged(isEnabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WrapXRotationChanged", isEnabled})
}

func (ptr *Q3DCamera) WrapYRotation() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WrapYRotation"}).(bool)
}

func (ptr *Q3DCamera) ConnectWrapYRotationChanged(f func(isEnabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWrapYRotationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectWrapYRotationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWrapYRotationChanged"})
}

func (ptr *Q3DCamera) WrapYRotationChanged(isEnabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WrapYRotationChanged", isEnabled})
}

func (ptr *Q3DCamera) XRotation() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XRotation"}).(float32)
}

func (ptr *Q3DCamera) ConnectXRotationChanged(f func(rotation float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectXRotationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectXRotationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectXRotationChanged"})
}

func (ptr *Q3DCamera) XRotationChanged(rotation float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XRotationChanged", rotation})
}

func (ptr *Q3DCamera) YRotation() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YRotation"}).(float32)
}

func (ptr *Q3DCamera) ConnectYRotationChanged(f func(rotation float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectYRotationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectYRotationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectYRotationChanged"})
}

func (ptr *Q3DCamera) YRotationChanged(rotation float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YRotationChanged", rotation})
}

func (ptr *Q3DCamera) ZoomLevel() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomLevel"}).(float32)
}

func (ptr *Q3DCamera) ConnectZoomLevelChanged(f func(zoomLevel float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZoomLevelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectZoomLevelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZoomLevelChanged"})
}

func (ptr *Q3DCamera) ZoomLevelChanged(zoomLevel float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomLevelChanged", zoomLevel})
}

func (ptr *Q3DCamera) ConnectDestroyQ3DCamera(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQ3DCamera", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DCamera) DisconnectDestroyQ3DCamera() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQ3DCamera"})
}

func (ptr *Q3DCamera) DestroyQ3DCamera() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DCamera"})
}

func (ptr *Q3DCamera) DestroyQ3DCameraDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DCameraDefault"})
}

type Q3DInputHandler struct {
	internal.Internal
}

type Q3DInputHandler_ITF interface {
	Q3DInputHandler_PTR() *Q3DInputHandler
}

func (ptr *Q3DInputHandler) Q3DInputHandler_PTR() *Q3DInputHandler {
	return ptr
}

func (ptr *Q3DInputHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *Q3DInputHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQ3DInputHandler(ptr Q3DInputHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DInputHandler_PTR().Pointer()
	}
	return nil
}

func (n *Q3DInputHandler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQ3DInputHandlerFromPointer(ptr unsafe.Pointer) (n *Q3DInputHandler) {
	n = new(Q3DInputHandler)
	n.InitFromInternal(uintptr(ptr), "datavisualization.Q3DInputHandler")
	return
}
func NewQ3DInputHandler(parent core.QObject_ITF) *Q3DInputHandler {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQ3DInputHandler", "", parent}).(*Q3DInputHandler)
}

func (ptr *Q3DInputHandler) IsRotationEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRotationEnabled"}).(bool)
}

func (ptr *Q3DInputHandler) IsSelectionEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSelectionEnabled"}).(bool)
}

func (ptr *Q3DInputHandler) IsZoomAtTargetEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsZoomAtTargetEnabled"}).(bool)
}

func (ptr *Q3DInputHandler) IsZoomEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsZoomEnabled"}).(bool)
}

func (ptr *Q3DInputHandler) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMouseMoveEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DInputHandler) DisconnectMouseMoveEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMouseMoveEvent"})
}

func (ptr *Q3DInputHandler) MouseMoveEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEvent", event, mousePos})
}

func (ptr *Q3DInputHandler) MouseMoveEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event, mousePos})
}

func (ptr *Q3DInputHandler) ConnectMousePressEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMousePressEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DInputHandler) DisconnectMousePressEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMousePressEvent"})
}

func (ptr *Q3DInputHandler) MousePressEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEvent", event, mousePos})
}

func (ptr *Q3DInputHandler) MousePressEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event, mousePos})
}

func (ptr *Q3DInputHandler) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMouseReleaseEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DInputHandler) DisconnectMouseReleaseEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMouseReleaseEvent"})
}

func (ptr *Q3DInputHandler) MouseReleaseEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEvent", event, mousePos})
}

func (ptr *Q3DInputHandler) MouseReleaseEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event, mousePos})
}

func (ptr *Q3DInputHandler) ConnectRotationEnabledChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRotationEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DInputHandler) DisconnectRotationEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRotationEnabledChanged"})
}

func (ptr *Q3DInputHandler) RotationEnabledChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationEnabledChanged", enable})
}

func (ptr *Q3DInputHandler) ConnectSelectionEnabledChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DInputHandler) DisconnectSelectionEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionEnabledChanged"})
}

func (ptr *Q3DInputHandler) SelectionEnabledChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionEnabledChanged", enable})
}

func (ptr *Q3DInputHandler) SetRotationEnabled(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotationEnabled", enable})
}

func (ptr *Q3DInputHandler) SetSelectionEnabled(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectionEnabled", enable})
}

func (ptr *Q3DInputHandler) SetZoomAtTargetEnabled(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZoomAtTargetEnabled", enable})
}

func (ptr *Q3DInputHandler) SetZoomEnabled(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZoomEnabled", enable})
}

func (ptr *Q3DInputHandler) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWheelEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DInputHandler) DisconnectWheelEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWheelEvent"})
}

func (ptr *Q3DInputHandler) WheelEvent(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEvent", event})
}

func (ptr *Q3DInputHandler) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *Q3DInputHandler) ConnectZoomAtTargetEnabledChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZoomAtTargetEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DInputHandler) DisconnectZoomAtTargetEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZoomAtTargetEnabledChanged"})
}

func (ptr *Q3DInputHandler) ZoomAtTargetEnabledChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomAtTargetEnabledChanged", enable})
}

func (ptr *Q3DInputHandler) ConnectZoomEnabledChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZoomEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DInputHandler) DisconnectZoomEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZoomEnabledChanged"})
}

func (ptr *Q3DInputHandler) ZoomEnabledChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomEnabledChanged", enable})
}

func (ptr *Q3DInputHandler) ConnectDestroyQ3DInputHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQ3DInputHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DInputHandler) DisconnectDestroyQ3DInputHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQ3DInputHandler"})
}

func (ptr *Q3DInputHandler) DestroyQ3DInputHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DInputHandler"})
}

func (ptr *Q3DInputHandler) DestroyQ3DInputHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DInputHandlerDefault"})
}

type Q3DLight struct {
	internal.Internal
}

type Q3DLight_ITF interface {
	Q3DLight_PTR() *Q3DLight
}

func (ptr *Q3DLight) Q3DLight_PTR() *Q3DLight {
	return ptr
}

func (ptr *Q3DLight) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *Q3DLight) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQ3DLight(ptr Q3DLight_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DLight_PTR().Pointer()
	}
	return nil
}

func (n *Q3DLight) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQ3DLightFromPointer(ptr unsafe.Pointer) (n *Q3DLight) {
	n = new(Q3DLight)
	n.InitFromInternal(uintptr(ptr), "datavisualization.Q3DLight")
	return
}
func NewQ3DLight(parent core.QObject_ITF) *Q3DLight {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQ3DLight", "", parent}).(*Q3DLight)
}

func (ptr *Q3DLight) ConnectAutoPositionChanged(f func(autoPosition bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAutoPositionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DLight) DisconnectAutoPositionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAutoPositionChanged"})
}

func (ptr *Q3DLight) AutoPositionChanged(autoPosition bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoPositionChanged", autoPosition})
}

func (ptr *Q3DLight) IsAutoPosition() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAutoPosition"}).(bool)
}

func (ptr *Q3DLight) SetAutoPosition(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoPosition", enabled})
}

func (ptr *Q3DLight) ConnectDestroyQ3DLight(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQ3DLight", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DLight) DisconnectDestroyQ3DLight() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQ3DLight"})
}

func (ptr *Q3DLight) DestroyQ3DLight() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DLight"})
}

func (ptr *Q3DLight) DestroyQ3DLightDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DLightDefault"})
}

type Q3DObject struct {
	internal.Internal
}

type Q3DObject_ITF interface {
	Q3DObject_PTR() *Q3DObject
}

func (ptr *Q3DObject) Q3DObject_PTR() *Q3DObject {
	return ptr
}

func (ptr *Q3DObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *Q3DObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQ3DObject(ptr Q3DObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DObject_PTR().Pointer()
	}
	return nil
}

func (n *Q3DObject) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQ3DObjectFromPointer(ptr unsafe.Pointer) (n *Q3DObject) {
	n = new(Q3DObject)
	n.InitFromInternal(uintptr(ptr), "datavisualization.Q3DObject")
	return
}
func NewQ3DObject(parent core.QObject_ITF) *Q3DObject {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQ3DObject", "", parent}).(*Q3DObject)
}

func (ptr *Q3DObject) ConnectCopyValuesFrom(f func(source *Q3DObject)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCopyValuesFrom", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DObject) DisconnectCopyValuesFrom() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCopyValuesFrom"})
}

func (ptr *Q3DObject) CopyValuesFrom(source Q3DObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CopyValuesFrom", source})
}

func (ptr *Q3DObject) CopyValuesFromDefault(source Q3DObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CopyValuesFromDefault", source})
}

func (ptr *Q3DObject) IsDirty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDirty"}).(bool)
}

func (ptr *Q3DObject) ParentScene() *Q3DScene {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParentScene"}).(*Q3DScene)
}

func (ptr *Q3DObject) Position() *gui.QVector3D {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Position"}).(*gui.QVector3D)
}

func (ptr *Q3DObject) ConnectPositionChanged(f func(position *gui.QVector3D)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPositionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DObject) DisconnectPositionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPositionChanged"})
}

func (ptr *Q3DObject) PositionChanged(position gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionChanged", position})
}

func (ptr *Q3DObject) SetDirty(dirty bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDirty", dirty})
}

func (ptr *Q3DObject) SetPosition(position gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPosition", position})
}

func (ptr *Q3DObject) ConnectDestroyQ3DObject(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQ3DObject", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DObject) DisconnectDestroyQ3DObject() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQ3DObject"})
}

func (ptr *Q3DObject) DestroyQ3DObject() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DObject"})
}

func (ptr *Q3DObject) DestroyQ3DObjectDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DObjectDefault"})
}

type Q3DScatter struct {
	internal.Internal
}

type Q3DScatter_ITF interface {
	Q3DScatter_PTR() *Q3DScatter
}

func (ptr *Q3DScatter) Q3DScatter_PTR() *Q3DScatter {
	return ptr
}

func (ptr *Q3DScatter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *Q3DScatter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQ3DScatter(ptr Q3DScatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DScatter_PTR().Pointer()
	}
	return nil
}

func (n *Q3DScatter) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQ3DScatterFromPointer(ptr unsafe.Pointer) (n *Q3DScatter) {
	n = new(Q3DScatter)
	n.InitFromInternal(uintptr(ptr), "datavisualization.Q3DScatter")
	return
}
func NewQ3DScatter(format gui.QSurfaceFormat_ITF, parent gui.QWindow_ITF) *Q3DScatter {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQ3DScatter", "", format, parent}).(*Q3DScatter)
}

func (ptr *Q3DScatter) AddAxis(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddAxis", axis})
}

func (ptr *Q3DScatter) AddSeries(series QScatter3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddSeries", series})
}

func (ptr *Q3DScatter) Axes() []*QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Axes"}).([]*QValue3DAxis)
}

func (ptr *Q3DScatter) AxisX() *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisX"}).(*QValue3DAxis)
}

func (ptr *Q3DScatter) ConnectAxisXChanged(f func(axis *QValue3DAxis)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxisXChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScatter) DisconnectAxisXChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxisXChanged"})
}

func (ptr *Q3DScatter) AxisXChanged(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisXChanged", axis})
}

func (ptr *Q3DScatter) AxisY() *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisY"}).(*QValue3DAxis)
}

func (ptr *Q3DScatter) ConnectAxisYChanged(f func(axis *QValue3DAxis)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxisYChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScatter) DisconnectAxisYChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxisYChanged"})
}

func (ptr *Q3DScatter) AxisYChanged(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisYChanged", axis})
}

func (ptr *Q3DScatter) AxisZ() *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisZ"}).(*QValue3DAxis)
}

func (ptr *Q3DScatter) ConnectAxisZChanged(f func(axis *QValue3DAxis)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxisZChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScatter) DisconnectAxisZChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxisZChanged"})
}

func (ptr *Q3DScatter) AxisZChanged(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisZChanged", axis})
}

func (ptr *Q3DScatter) ReleaseAxis(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReleaseAxis", axis})
}

func (ptr *Q3DScatter) RemoveSeries(series QScatter3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveSeries", series})
}

func (ptr *Q3DScatter) SelectedSeries() *QScatter3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedSeries"}).(*QScatter3DSeries)
}

func (ptr *Q3DScatter) ConnectSelectedSeriesChanged(f func(series *QScatter3DSeries)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectedSeriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScatter) DisconnectSelectedSeriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectedSeriesChanged"})
}

func (ptr *Q3DScatter) SelectedSeriesChanged(series QScatter3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedSeriesChanged", series})
}

func (ptr *Q3DScatter) SeriesList() []*QScatter3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesList"}).([]*QScatter3DSeries)
}

func (ptr *Q3DScatter) SetAxisX(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAxisX", axis})
}

func (ptr *Q3DScatter) SetAxisY(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAxisY", axis})
}

func (ptr *Q3DScatter) SetAxisZ(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAxisZ", axis})
}

func (ptr *Q3DScatter) ConnectDestroyQ3DScatter(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQ3DScatter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScatter) DisconnectDestroyQ3DScatter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQ3DScatter"})
}

func (ptr *Q3DScatter) DestroyQ3DScatter() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DScatter"})
}

func (ptr *Q3DScatter) DestroyQ3DScatterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DScatterDefault"})
}

func (ptr *Q3DScatter) __axes_atList(i int) *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_atList", i}).(*QValue3DAxis)
}

func (ptr *Q3DScatter) __axes_setList(i QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_setList", i})
}

func (ptr *Q3DScatter) __axes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_newList"}).(unsafe.Pointer)
}

func (ptr *Q3DScatter) __seriesList_atList(i int) *QScatter3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__seriesList_atList", i}).(*QScatter3DSeries)
}

func (ptr *Q3DScatter) __seriesList_setList(i QScatter3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__seriesList_setList", i})
}

func (ptr *Q3DScatter) __seriesList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__seriesList_newList"}).(unsafe.Pointer)
}

type Q3DScene struct {
	internal.Internal
}

type Q3DScene_ITF interface {
	Q3DScene_PTR() *Q3DScene
}

func (ptr *Q3DScene) Q3DScene_PTR() *Q3DScene {
	return ptr
}

func (ptr *Q3DScene) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *Q3DScene) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQ3DScene(ptr Q3DScene_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DScene_PTR().Pointer()
	}
	return nil
}

func (n *Q3DScene) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQ3DSceneFromPointer(ptr unsafe.Pointer) (n *Q3DScene) {
	n = new(Q3DScene)
	n.InitFromInternal(uintptr(ptr), "datavisualization.Q3DScene")
	return
}
func NewQ3DScene(parent core.QObject_ITF) *Q3DScene {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQ3DScene", "", parent}).(*Q3DScene)
}

func (ptr *Q3DScene) ActiveCamera() *Q3DCamera {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveCamera"}).(*Q3DCamera)
}

func (ptr *Q3DScene) ConnectActiveCameraChanged(f func(camera *Q3DCamera)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveCameraChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectActiveCameraChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveCameraChanged"})
}

func (ptr *Q3DScene) ActiveCameraChanged(camera Q3DCamera_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveCameraChanged", camera})
}

func (ptr *Q3DScene) ActiveLight() *Q3DLight {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveLight"}).(*Q3DLight)
}

func (ptr *Q3DScene) ConnectActiveLightChanged(f func(light *Q3DLight)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveLightChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectActiveLightChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveLightChanged"})
}

func (ptr *Q3DScene) ActiveLightChanged(light Q3DLight_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveLightChanged", light})
}

func (ptr *Q3DScene) DevicePixelRatio() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DevicePixelRatio"}).(float32)
}

func (ptr *Q3DScene) ConnectDevicePixelRatioChanged(f func(pixelRatio float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDevicePixelRatioChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectDevicePixelRatioChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDevicePixelRatioChanged"})
}

func (ptr *Q3DScene) DevicePixelRatioChanged(pixelRatio float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DevicePixelRatioChanged", pixelRatio})
}

func (ptr *Q3DScene) GraphPositionQuery() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GraphPositionQuery"}).(*core.QPoint)
}

func (ptr *Q3DScene) ConnectGraphPositionQueryChanged(f func(position *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGraphPositionQueryChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectGraphPositionQueryChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGraphPositionQueryChanged"})
}

func (ptr *Q3DScene) GraphPositionQueryChanged(position core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GraphPositionQueryChanged", position})
}

func Q3DScene_InvalidSelectionPoint() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.Q3DScene_InvalidSelectionPoint", ""}).(*core.QPoint)
}

func (ptr *Q3DScene) InvalidSelectionPoint() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.Q3DScene_InvalidSelectionPoint", ""}).(*core.QPoint)
}

func (ptr *Q3DScene) IsPointInPrimarySubView(point core.QPoint_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsPointInPrimarySubView", point}).(bool)
}

func (ptr *Q3DScene) IsPointInSecondarySubView(point core.QPoint_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsPointInSecondarySubView", point}).(bool)
}

func (ptr *Q3DScene) IsSecondarySubviewOnTop() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSecondarySubviewOnTop"}).(bool)
}

func (ptr *Q3DScene) IsSlicingActive() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSlicingActive"}).(bool)
}

func (ptr *Q3DScene) PrimarySubViewport() *core.QRect {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrimarySubViewport"}).(*core.QRect)
}

func (ptr *Q3DScene) ConnectPrimarySubViewportChanged(f func(subViewport *core.QRect)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrimarySubViewportChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectPrimarySubViewportChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrimarySubViewportChanged"})
}

func (ptr *Q3DScene) PrimarySubViewportChanged(subViewport core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrimarySubViewportChanged", subViewport})
}

func (ptr *Q3DScene) SecondarySubViewport() *core.QRect {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SecondarySubViewport"}).(*core.QRect)
}

func (ptr *Q3DScene) ConnectSecondarySubViewportChanged(f func(subViewport *core.QRect)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSecondarySubViewportChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectSecondarySubViewportChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSecondarySubViewportChanged"})
}

func (ptr *Q3DScene) SecondarySubViewportChanged(subViewport core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SecondarySubViewportChanged", subViewport})
}

func (ptr *Q3DScene) ConnectSecondarySubviewOnTopChanged(f func(isSecondaryOnTop bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSecondarySubviewOnTopChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectSecondarySubviewOnTopChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSecondarySubviewOnTopChanged"})
}

func (ptr *Q3DScene) SecondarySubviewOnTopChanged(isSecondaryOnTop bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SecondarySubviewOnTopChanged", isSecondaryOnTop})
}

func (ptr *Q3DScene) SelectionQueryPosition() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionQueryPosition"}).(*core.QPoint)
}

func (ptr *Q3DScene) ConnectSelectionQueryPositionChanged(f func(position *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionQueryPositionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectSelectionQueryPositionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionQueryPositionChanged"})
}

func (ptr *Q3DScene) SelectionQueryPositionChanged(position core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionQueryPositionChanged", position})
}

func (ptr *Q3DScene) SetActiveCamera(camera Q3DCamera_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActiveCamera", camera})
}

func (ptr *Q3DScene) SetActiveLight(light Q3DLight_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActiveLight", light})
}

func (ptr *Q3DScene) SetDevicePixelRatio(pixelRatio float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDevicePixelRatio", pixelRatio})
}

func (ptr *Q3DScene) SetGraphPositionQuery(point core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGraphPositionQuery", point})
}

func (ptr *Q3DScene) SetPrimarySubViewport(primarySubViewport core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrimarySubViewport", primarySubViewport})
}

func (ptr *Q3DScene) SetSecondarySubViewport(secondarySubViewport core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSecondarySubViewport", secondarySubViewport})
}

func (ptr *Q3DScene) SetSecondarySubviewOnTop(isSecondaryOnTop bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSecondarySubviewOnTop", isSecondaryOnTop})
}

func (ptr *Q3DScene) SetSelectionQueryPosition(point core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectionQueryPosition", point})
}

func (ptr *Q3DScene) SetSlicingActive(isSlicing bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSlicingActive", isSlicing})
}

func (ptr *Q3DScene) ConnectSlicingActiveChanged(f func(isSlicingActive bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSlicingActiveChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectSlicingActiveChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSlicingActiveChanged"})
}

func (ptr *Q3DScene) SlicingActiveChanged(isSlicingActive bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SlicingActiveChanged", isSlicingActive})
}

func (ptr *Q3DScene) Viewport() *core.QRect {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Viewport"}).(*core.QRect)
}

func (ptr *Q3DScene) ConnectViewportChanged(f func(viewport *core.QRect)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectViewportChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectViewportChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectViewportChanged"})
}

func (ptr *Q3DScene) ViewportChanged(viewport core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewportChanged", viewport})
}

func (ptr *Q3DScene) ConnectDestroyQ3DScene(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQ3DScene", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DScene) DisconnectDestroyQ3DScene() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQ3DScene"})
}

func (ptr *Q3DScene) DestroyQ3DScene() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DScene"})
}

func (ptr *Q3DScene) DestroyQ3DSceneDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DSceneDefault"})
}

type Q3DSurface struct {
	internal.Internal
}

type Q3DSurface_ITF interface {
	Q3DSurface_PTR() *Q3DSurface
}

func (ptr *Q3DSurface) Q3DSurface_PTR() *Q3DSurface {
	return ptr
}

func (ptr *Q3DSurface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *Q3DSurface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQ3DSurface(ptr Q3DSurface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DSurface_PTR().Pointer()
	}
	return nil
}

func (n *Q3DSurface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQ3DSurfaceFromPointer(ptr unsafe.Pointer) (n *Q3DSurface) {
	n = new(Q3DSurface)
	n.InitFromInternal(uintptr(ptr), "datavisualization.Q3DSurface")
	return
}
func NewQ3DSurface(format gui.QSurfaceFormat_ITF, parent gui.QWindow_ITF) *Q3DSurface {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQ3DSurface", "", format, parent}).(*Q3DSurface)
}

func (ptr *Q3DSurface) AddAxis(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddAxis", axis})
}

func (ptr *Q3DSurface) AddSeries(series QSurface3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddSeries", series})
}

func (ptr *Q3DSurface) Axes() []*QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Axes"}).([]*QValue3DAxis)
}

func (ptr *Q3DSurface) AxisX() *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisX"}).(*QValue3DAxis)
}

func (ptr *Q3DSurface) ConnectAxisXChanged(f func(axis *QValue3DAxis)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxisXChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DSurface) DisconnectAxisXChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxisXChanged"})
}

func (ptr *Q3DSurface) AxisXChanged(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisXChanged", axis})
}

func (ptr *Q3DSurface) AxisY() *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisY"}).(*QValue3DAxis)
}

func (ptr *Q3DSurface) ConnectAxisYChanged(f func(axis *QValue3DAxis)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxisYChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DSurface) DisconnectAxisYChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxisYChanged"})
}

func (ptr *Q3DSurface) AxisYChanged(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisYChanged", axis})
}

func (ptr *Q3DSurface) AxisZ() *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisZ"}).(*QValue3DAxis)
}

func (ptr *Q3DSurface) ConnectAxisZChanged(f func(axis *QValue3DAxis)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxisZChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DSurface) DisconnectAxisZChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxisZChanged"})
}

func (ptr *Q3DSurface) AxisZChanged(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisZChanged", axis})
}

func (ptr *Q3DSurface) FlipHorizontalGrid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlipHorizontalGrid"}).(bool)
}

func (ptr *Q3DSurface) ConnectFlipHorizontalGridChanged(f func(flip bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFlipHorizontalGridChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DSurface) DisconnectFlipHorizontalGridChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFlipHorizontalGridChanged"})
}

func (ptr *Q3DSurface) FlipHorizontalGridChanged(flip bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlipHorizontalGridChanged", flip})
}

func (ptr *Q3DSurface) ReleaseAxis(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReleaseAxis", axis})
}

func (ptr *Q3DSurface) RemoveSeries(series QSurface3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveSeries", series})
}

func (ptr *Q3DSurface) SelectedSeries() *QSurface3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedSeries"}).(*QSurface3DSeries)
}

func (ptr *Q3DSurface) ConnectSelectedSeriesChanged(f func(series *QSurface3DSeries)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectedSeriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DSurface) DisconnectSelectedSeriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectedSeriesChanged"})
}

func (ptr *Q3DSurface) SelectedSeriesChanged(series QSurface3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedSeriesChanged", series})
}

func (ptr *Q3DSurface) SeriesList() []*QSurface3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesList"}).([]*QSurface3DSeries)
}

func (ptr *Q3DSurface) SetAxisX(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAxisX", axis})
}

func (ptr *Q3DSurface) SetAxisY(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAxisY", axis})
}

func (ptr *Q3DSurface) SetAxisZ(axis QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAxisZ", axis})
}

func (ptr *Q3DSurface) SetFlipHorizontalGrid(flip bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFlipHorizontalGrid", flip})
}

func (ptr *Q3DSurface) ConnectDestroyQ3DSurface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQ3DSurface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DSurface) DisconnectDestroyQ3DSurface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQ3DSurface"})
}

func (ptr *Q3DSurface) DestroyQ3DSurface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DSurface"})
}

func (ptr *Q3DSurface) DestroyQ3DSurfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DSurfaceDefault"})
}

func (ptr *Q3DSurface) __axes_atList(i int) *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_atList", i}).(*QValue3DAxis)
}

func (ptr *Q3DSurface) __axes_setList(i QValue3DAxis_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_setList", i})
}

func (ptr *Q3DSurface) __axes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__axes_newList"}).(unsafe.Pointer)
}

func (ptr *Q3DSurface) __seriesList_atList(i int) *QSurface3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__seriesList_atList", i}).(*QSurface3DSeries)
}

func (ptr *Q3DSurface) __seriesList_setList(i QSurface3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__seriesList_setList", i})
}

func (ptr *Q3DSurface) __seriesList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__seriesList_newList"}).(unsafe.Pointer)
}

type Q3DTheme struct {
	internal.Internal
}

type Q3DTheme_ITF interface {
	Q3DTheme_PTR() *Q3DTheme
}

func (ptr *Q3DTheme) Q3DTheme_PTR() *Q3DTheme {
	return ptr
}

func (ptr *Q3DTheme) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *Q3DTheme) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQ3DTheme(ptr Q3DTheme_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DTheme_PTR().Pointer()
	}
	return nil
}

func (n *Q3DTheme) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQ3DThemeFromPointer(ptr unsafe.Pointer) (n *Q3DTheme) {
	n = new(Q3DTheme)
	n.InitFromInternal(uintptr(ptr), "datavisualization.Q3DTheme")
	return
}

//go:generate stringer -type=Q3DTheme__ColorStyle
//Q3DTheme::ColorStyle
type Q3DTheme__ColorStyle int64

var (
	Q3DTheme__ColorStyleUniform        Q3DTheme__ColorStyle = Q3DTheme__ColorStyle(0)
	Q3DTheme__ColorStyleObjectGradient Q3DTheme__ColorStyle = Q3DTheme__ColorStyle(1)
	Q3DTheme__ColorStyleRangeGradient  Q3DTheme__ColorStyle = Q3DTheme__ColorStyle(2)
)

//go:generate stringer -type=Q3DTheme__Theme
//Q3DTheme::Theme
type Q3DTheme__Theme int64

const (
	Q3DTheme__ThemeQt            Q3DTheme__Theme = Q3DTheme__Theme(0)
	Q3DTheme__ThemePrimaryColors Q3DTheme__Theme = Q3DTheme__Theme(1)
	Q3DTheme__ThemeDigia         Q3DTheme__Theme = Q3DTheme__Theme(2)
	Q3DTheme__ThemeStoneMoss     Q3DTheme__Theme = Q3DTheme__Theme(3)
	Q3DTheme__ThemeArmyBlue      Q3DTheme__Theme = Q3DTheme__Theme(4)
	Q3DTheme__ThemeRetro         Q3DTheme__Theme = Q3DTheme__Theme(5)
	Q3DTheme__ThemeEbony         Q3DTheme__Theme = Q3DTheme__Theme(6)
	Q3DTheme__ThemeIsabelle      Q3DTheme__Theme = Q3DTheme__Theme(7)
	Q3DTheme__ThemeUserDefined   Q3DTheme__Theme = Q3DTheme__Theme(8)
)

func NewQ3DTheme(parent core.QObject_ITF) *Q3DTheme {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQ3DTheme", "", parent}).(*Q3DTheme)
}

func NewQ3DTheme2(themeType Q3DTheme__Theme, parent core.QObject_ITF) *Q3DTheme {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQ3DTheme2", "", themeType, parent}).(*Q3DTheme)
}

func (ptr *Q3DTheme) AmbientLightStrength() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AmbientLightStrength"}).(float32)
}

func (ptr *Q3DTheme) ConnectAmbientLightStrengthChanged(f func(strength float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAmbientLightStrengthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectAmbientLightStrengthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAmbientLightStrengthChanged"})
}

func (ptr *Q3DTheme) AmbientLightStrengthChanged(strength float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AmbientLightStrengthChanged", strength})
}

func (ptr *Q3DTheme) BackgroundColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundColor"}).(*gui.QColor)
}

func (ptr *Q3DTheme) ConnectBackgroundColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBackgroundColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectBackgroundColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBackgroundColorChanged"})
}

func (ptr *Q3DTheme) BackgroundColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundColorChanged", color})
}

func (ptr *Q3DTheme) ConnectBackgroundEnabledChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBackgroundEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectBackgroundEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBackgroundEnabledChanged"})
}

func (ptr *Q3DTheme) BackgroundEnabledChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundEnabledChanged", enabled})
}

func (ptr *Q3DTheme) BaseColors() []*gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseColors"}).([]*gui.QColor)
}

func (ptr *Q3DTheme) ConnectBaseColorsChanged(f func(colors []*gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBaseColorsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectBaseColorsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBaseColorsChanged"})
}

func (ptr *Q3DTheme) BaseColorsChanged(colors []*gui.QColor) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseColorsChanged", colors})
}

func (ptr *Q3DTheme) BaseGradients() []*gui.QLinearGradient {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseGradients"}).([]*gui.QLinearGradient)
}

func (ptr *Q3DTheme) ConnectBaseGradientsChanged(f func(gradients []*gui.QLinearGradient)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBaseGradientsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectBaseGradientsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBaseGradientsChanged"})
}

func (ptr *Q3DTheme) BaseGradientsChanged(gradients []*gui.QLinearGradient) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseGradientsChanged", gradients})
}

func (ptr *Q3DTheme) ColorStyle() Q3DTheme__ColorStyle {

	return Q3DTheme__ColorStyle(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorStyle"}).(float64))
}

func (ptr *Q3DTheme) ConnectColorStyleChanged(f func(style Q3DTheme__ColorStyle)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColorStyleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectColorStyleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColorStyleChanged"})
}

func (ptr *Q3DTheme) ColorStyleChanged(style Q3DTheme__ColorStyle) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorStyleChanged", style})
}

func (ptr *Q3DTheme) Font() *gui.QFont {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Font"}).(*gui.QFont)
}

func (ptr *Q3DTheme) ConnectFontChanged(f func(font *gui.QFont)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFontChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectFontChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFontChanged"})
}

func (ptr *Q3DTheme) FontChanged(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FontChanged", font})
}

func (ptr *Q3DTheme) ConnectGridEnabledChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGridEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectGridEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGridEnabledChanged"})
}

func (ptr *Q3DTheme) GridEnabledChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GridEnabledChanged", enabled})
}

func (ptr *Q3DTheme) GridLineColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GridLineColor"}).(*gui.QColor)
}

func (ptr *Q3DTheme) ConnectGridLineColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGridLineColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectGridLineColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGridLineColorChanged"})
}

func (ptr *Q3DTheme) GridLineColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GridLineColorChanged", color})
}

func (ptr *Q3DTheme) HighlightLightStrength() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HighlightLightStrength"}).(float32)
}

func (ptr *Q3DTheme) ConnectHighlightLightStrengthChanged(f func(strength float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHighlightLightStrengthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectHighlightLightStrengthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHighlightLightStrengthChanged"})
}

func (ptr *Q3DTheme) HighlightLightStrengthChanged(strength float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HighlightLightStrengthChanged", strength})
}

func (ptr *Q3DTheme) IsBackgroundEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBackgroundEnabled"}).(bool)
}

func (ptr *Q3DTheme) IsGridEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsGridEnabled"}).(bool)
}

func (ptr *Q3DTheme) IsLabelBackgroundEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLabelBackgroundEnabled"}).(bool)
}

func (ptr *Q3DTheme) IsLabelBorderEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLabelBorderEnabled"}).(bool)
}

func (ptr *Q3DTheme) LabelBackgroundColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBackgroundColor"}).(*gui.QColor)
}

func (ptr *Q3DTheme) ConnectLabelBackgroundColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelBackgroundColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectLabelBackgroundColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelBackgroundColorChanged"})
}

func (ptr *Q3DTheme) LabelBackgroundColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBackgroundColorChanged", color})
}

func (ptr *Q3DTheme) ConnectLabelBackgroundEnabledChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelBackgroundEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectLabelBackgroundEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelBackgroundEnabledChanged"})
}

func (ptr *Q3DTheme) LabelBackgroundEnabledChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBackgroundEnabledChanged", enabled})
}

func (ptr *Q3DTheme) ConnectLabelBorderEnabledChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelBorderEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectLabelBorderEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelBorderEnabledChanged"})
}

func (ptr *Q3DTheme) LabelBorderEnabledChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelBorderEnabledChanged", enabled})
}

func (ptr *Q3DTheme) LabelTextColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelTextColor"}).(*gui.QColor)
}

func (ptr *Q3DTheme) ConnectLabelTextColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelTextColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectLabelTextColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelTextColorChanged"})
}

func (ptr *Q3DTheme) LabelTextColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelTextColorChanged", color})
}

func (ptr *Q3DTheme) LightColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LightColor"}).(*gui.QColor)
}

func (ptr *Q3DTheme) ConnectLightColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLightColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectLightColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLightColorChanged"})
}

func (ptr *Q3DTheme) LightColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LightColorChanged", color})
}

func (ptr *Q3DTheme) LightStrength() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LightStrength"}).(float32)
}

func (ptr *Q3DTheme) ConnectLightStrengthChanged(f func(strength float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLightStrengthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectLightStrengthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLightStrengthChanged"})
}

func (ptr *Q3DTheme) LightStrengthChanged(strength float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LightStrengthChanged", strength})
}

func (ptr *Q3DTheme) MultiHighlightColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiHighlightColor"}).(*gui.QColor)
}

func (ptr *Q3DTheme) ConnectMultiHighlightColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMultiHighlightColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectMultiHighlightColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMultiHighlightColorChanged"})
}

func (ptr *Q3DTheme) MultiHighlightColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiHighlightColorChanged", color})
}

func (ptr *Q3DTheme) MultiHighlightGradient() *gui.QLinearGradient {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiHighlightGradient"}).(*gui.QLinearGradient)
}

func (ptr *Q3DTheme) ConnectMultiHighlightGradientChanged(f func(gradient *gui.QLinearGradient)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMultiHighlightGradientChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectMultiHighlightGradientChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMultiHighlightGradientChanged"})
}

func (ptr *Q3DTheme) MultiHighlightGradientChanged(gradient gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiHighlightGradientChanged", gradient})
}

func (ptr *Q3DTheme) SetAmbientLightStrength(strength float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAmbientLightStrength", strength})
}

func (ptr *Q3DTheme) SetBackgroundColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackgroundColor", color})
}

func (ptr *Q3DTheme) SetBackgroundEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackgroundEnabled", enabled})
}

func (ptr *Q3DTheme) SetBaseColors(colors []*gui.QColor) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBaseColors", colors})
}

func (ptr *Q3DTheme) SetBaseGradients(gradients []*gui.QLinearGradient) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBaseGradients", gradients})
}

func (ptr *Q3DTheme) SetColorStyle(style Q3DTheme__ColorStyle) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColorStyle", style})
}

func (ptr *Q3DTheme) SetFont(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFont", font})
}

func (ptr *Q3DTheme) SetGridEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGridEnabled", enabled})
}

func (ptr *Q3DTheme) SetGridLineColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGridLineColor", color})
}

func (ptr *Q3DTheme) SetHighlightLightStrength(strength float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHighlightLightStrength", strength})
}

func (ptr *Q3DTheme) SetLabelBackgroundColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelBackgroundColor", color})
}

func (ptr *Q3DTheme) SetLabelBackgroundEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelBackgroundEnabled", enabled})
}

func (ptr *Q3DTheme) SetLabelBorderEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelBorderEnabled", enabled})
}

func (ptr *Q3DTheme) SetLabelTextColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelTextColor", color})
}

func (ptr *Q3DTheme) SetLightColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLightColor", color})
}

func (ptr *Q3DTheme) SetLightStrength(strength float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLightStrength", strength})
}

func (ptr *Q3DTheme) SetMultiHighlightColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMultiHighlightColor", color})
}

func (ptr *Q3DTheme) SetMultiHighlightGradient(gradient gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMultiHighlightGradient", gradient})
}

func (ptr *Q3DTheme) SetSingleHighlightColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSingleHighlightColor", color})
}

func (ptr *Q3DTheme) SetSingleHighlightGradient(gradient gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSingleHighlightGradient", gradient})
}

func (ptr *Q3DTheme) SetType(themeType Q3DTheme__Theme) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetType", themeType})
}

func (ptr *Q3DTheme) SetWindowColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowColor", color})
}

func (ptr *Q3DTheme) SingleHighlightColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SingleHighlightColor"}).(*gui.QColor)
}

func (ptr *Q3DTheme) ConnectSingleHighlightColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSingleHighlightColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectSingleHighlightColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSingleHighlightColorChanged"})
}

func (ptr *Q3DTheme) SingleHighlightColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SingleHighlightColorChanged", color})
}

func (ptr *Q3DTheme) SingleHighlightGradient() *gui.QLinearGradient {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SingleHighlightGradient"}).(*gui.QLinearGradient)
}

func (ptr *Q3DTheme) ConnectSingleHighlightGradientChanged(f func(gradient *gui.QLinearGradient)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSingleHighlightGradientChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectSingleHighlightGradientChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSingleHighlightGradientChanged"})
}

func (ptr *Q3DTheme) SingleHighlightGradientChanged(gradient gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SingleHighlightGradientChanged", gradient})
}

func (ptr *Q3DTheme) Type() Q3DTheme__Theme {

	return Q3DTheme__Theme(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *Q3DTheme) ConnectTypeChanged(f func(themeType Q3DTheme__Theme)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTypeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectTypeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTypeChanged"})
}

func (ptr *Q3DTheme) TypeChanged(themeType Q3DTheme__Theme) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeChanged", themeType})
}

func (ptr *Q3DTheme) WindowColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WindowColor"}).(*gui.QColor)
}

func (ptr *Q3DTheme) ConnectWindowColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWindowColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectWindowColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWindowColorChanged"})
}

func (ptr *Q3DTheme) WindowColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WindowColorChanged", color})
}

func (ptr *Q3DTheme) ConnectDestroyQ3DTheme(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQ3DTheme", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *Q3DTheme) DisconnectDestroyQ3DTheme() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQ3DTheme"})
}

func (ptr *Q3DTheme) DestroyQ3DTheme() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DTheme"})
}

func (ptr *Q3DTheme) DestroyQ3DThemeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQ3DThemeDefault"})
}

func (ptr *Q3DTheme) __baseColors_atList(i int) *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseColors_atList", i}).(*gui.QColor)
}

func (ptr *Q3DTheme) __baseColors_setList(i gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseColors_setList", i})
}

func (ptr *Q3DTheme) __baseColors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseColors_newList"}).(unsafe.Pointer)
}

func (ptr *Q3DTheme) __baseColorsChanged_colors_atList(i int) *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseColorsChanged_colors_atList", i}).(*gui.QColor)
}

func (ptr *Q3DTheme) __baseColorsChanged_colors_setList(i gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseColorsChanged_colors_setList", i})
}

func (ptr *Q3DTheme) __baseColorsChanged_colors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseColorsChanged_colors_newList"}).(unsafe.Pointer)
}

func (ptr *Q3DTheme) __baseGradients_atList(i int) *gui.QLinearGradient {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseGradients_atList", i}).(*gui.QLinearGradient)
}

func (ptr *Q3DTheme) __baseGradients_setList(i gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseGradients_setList", i})
}

func (ptr *Q3DTheme) __baseGradients_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseGradients_newList"}).(unsafe.Pointer)
}

func (ptr *Q3DTheme) __baseGradientsChanged_gradients_atList(i int) *gui.QLinearGradient {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseGradientsChanged_gradients_atList", i}).(*gui.QLinearGradient)
}

func (ptr *Q3DTheme) __baseGradientsChanged_gradients_setList(i gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseGradientsChanged_gradients_setList", i})
}

func (ptr *Q3DTheme) __baseGradientsChanged_gradients_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__baseGradientsChanged_gradients_newList"}).(unsafe.Pointer)
}

func (ptr *Q3DTheme) __setBaseColors_colors_atList(i int) *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setBaseColors_colors_atList", i}).(*gui.QColor)
}

func (ptr *Q3DTheme) __setBaseColors_colors_setList(i gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setBaseColors_colors_setList", i})
}

func (ptr *Q3DTheme) __setBaseColors_colors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setBaseColors_colors_newList"}).(unsafe.Pointer)
}

func (ptr *Q3DTheme) __setBaseGradients_gradients_atList(i int) *gui.QLinearGradient {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setBaseGradients_gradients_atList", i}).(*gui.QLinearGradient)
}

func (ptr *Q3DTheme) __setBaseGradients_gradients_setList(i gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setBaseGradients_gradients_setList", i})
}

func (ptr *Q3DTheme) __setBaseGradients_gradients_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setBaseGradients_gradients_newList"}).(unsafe.Pointer)
}

type QAbstract3DAxis struct {
	internal.Internal
}

type QAbstract3DAxis_ITF interface {
	QAbstract3DAxis_PTR() *QAbstract3DAxis
}

func (ptr *QAbstract3DAxis) QAbstract3DAxis_PTR() *QAbstract3DAxis {
	return ptr
}

func (ptr *QAbstract3DAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAbstract3DAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAbstract3DAxis(ptr QAbstract3DAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstract3DAxis_PTR().Pointer()
	}
	return nil
}

func (n *QAbstract3DAxis) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAbstract3DAxisFromPointer(ptr unsafe.Pointer) (n *QAbstract3DAxis) {
	n = new(QAbstract3DAxis)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QAbstract3DAxis")
	return
}

//go:generate stringer -type=QAbstract3DAxis__AxisOrientation
//QAbstract3DAxis::AxisOrientation
type QAbstract3DAxis__AxisOrientation int64

const (
	QAbstract3DAxis__AxisOrientationNone QAbstract3DAxis__AxisOrientation = QAbstract3DAxis__AxisOrientation(0)
	QAbstract3DAxis__AxisOrientationX    QAbstract3DAxis__AxisOrientation = QAbstract3DAxis__AxisOrientation(1)
	QAbstract3DAxis__AxisOrientationY    QAbstract3DAxis__AxisOrientation = QAbstract3DAxis__AxisOrientation(2)
	QAbstract3DAxis__AxisOrientationZ    QAbstract3DAxis__AxisOrientation = QAbstract3DAxis__AxisOrientation(4)
)

//go:generate stringer -type=QAbstract3DAxis__AxisType
//QAbstract3DAxis::AxisType
type QAbstract3DAxis__AxisType int64

const (
	QAbstract3DAxis__AxisTypeNone     QAbstract3DAxis__AxisType = QAbstract3DAxis__AxisType(0)
	QAbstract3DAxis__AxisTypeCategory QAbstract3DAxis__AxisType = QAbstract3DAxis__AxisType(1)
	QAbstract3DAxis__AxisTypeValue    QAbstract3DAxis__AxisType = QAbstract3DAxis__AxisType(2)
)

func (ptr *QAbstract3DAxis) ConnectAutoAdjustRangeChanged(f func(autoAdjust bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAutoAdjustRangeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectAutoAdjustRangeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAutoAdjustRangeChanged"})
}

func (ptr *QAbstract3DAxis) AutoAdjustRangeChanged(autoAdjust bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoAdjustRangeChanged", autoAdjust})
}

func (ptr *QAbstract3DAxis) IsAutoAdjustRange() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAutoAdjustRange"}).(bool)
}

func (ptr *QAbstract3DAxis) IsTitleFixed() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsTitleFixed"}).(bool)
}

func (ptr *QAbstract3DAxis) IsTitleVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsTitleVisible"}).(bool)
}

func (ptr *QAbstract3DAxis) LabelAutoRotation() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelAutoRotation"}).(float32)
}

func (ptr *QAbstract3DAxis) ConnectLabelAutoRotationChanged(f func(angle float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelAutoRotationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectLabelAutoRotationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelAutoRotationChanged"})
}

func (ptr *QAbstract3DAxis) LabelAutoRotationChanged(angle float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelAutoRotationChanged", angle})
}

func (ptr *QAbstract3DAxis) Labels() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Labels"}).([]string)
}

func (ptr *QAbstract3DAxis) ConnectLabelsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectLabelsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsChanged"})
}

func (ptr *QAbstract3DAxis) LabelsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsChanged"})
}

func (ptr *QAbstract3DAxis) Max() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Max"}).(float32)
}

func (ptr *QAbstract3DAxis) ConnectMaxChanged(f func(value float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMaxChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectMaxChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMaxChanged"})
}

func (ptr *QAbstract3DAxis) MaxChanged(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxChanged", value})
}

func (ptr *QAbstract3DAxis) Min() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Min"}).(float32)
}

func (ptr *QAbstract3DAxis) ConnectMinChanged(f func(value float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectMinChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinChanged"})
}

func (ptr *QAbstract3DAxis) MinChanged(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinChanged", value})
}

func (ptr *QAbstract3DAxis) Orientation() QAbstract3DAxis__AxisOrientation {

	return QAbstract3DAxis__AxisOrientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Orientation"}).(float64))
}

func (ptr *QAbstract3DAxis) ConnectOrientationChanged(f func(orientation QAbstract3DAxis__AxisOrientation)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOrientationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectOrientationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOrientationChanged"})
}

func (ptr *QAbstract3DAxis) OrientationChanged(orientation QAbstract3DAxis__AxisOrientation) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OrientationChanged", orientation})
}

func (ptr *QAbstract3DAxis) ConnectRangeChanged(f func(min float32, max float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRangeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectRangeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRangeChanged"})
}

func (ptr *QAbstract3DAxis) RangeChanged(min float32, max float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RangeChanged", min, max})
}

func (ptr *QAbstract3DAxis) SetAutoAdjustRange(autoAdjust bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoAdjustRange", autoAdjust})
}

func (ptr *QAbstract3DAxis) SetLabelAutoRotation(angle float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelAutoRotation", angle})
}

func (ptr *QAbstract3DAxis) SetLabels(labels []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabels", labels})
}

func (ptr *QAbstract3DAxis) SetMax(max float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMax", max})
}

func (ptr *QAbstract3DAxis) SetMin(min float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMin", min})
}

func (ptr *QAbstract3DAxis) SetRange(min float32, max float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRange", min, max})
}

func (ptr *QAbstract3DAxis) SetTitle(title string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitle", title})
}

func (ptr *QAbstract3DAxis) SetTitleFixed(fixed bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitleFixed", fixed})
}

func (ptr *QAbstract3DAxis) SetTitleVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitleVisible", visible})
}

func (ptr *QAbstract3DAxis) Title() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Title"}).(string)
}

func (ptr *QAbstract3DAxis) ConnectTitleChanged(f func(newTitle string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTitleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectTitleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTitleChanged"})
}

func (ptr *QAbstract3DAxis) TitleChanged(newTitle string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleChanged", newTitle})
}

func (ptr *QAbstract3DAxis) ConnectTitleFixedChanged(f func(fixed bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTitleFixedChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectTitleFixedChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTitleFixedChanged"})
}

func (ptr *QAbstract3DAxis) TitleFixedChanged(fixed bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleFixedChanged", fixed})
}

func (ptr *QAbstract3DAxis) ConnectTitleVisibilityChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTitleVisibilityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectTitleVisibilityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTitleVisibilityChanged"})
}

func (ptr *QAbstract3DAxis) TitleVisibilityChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TitleVisibilityChanged", visible})
}

func (ptr *QAbstract3DAxis) Type() QAbstract3DAxis__AxisType {

	return QAbstract3DAxis__AxisType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QAbstract3DAxis) ConnectDestroyQAbstract3DAxis(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstract3DAxis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DAxis) DisconnectDestroyQAbstract3DAxis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstract3DAxis"})
}

func (ptr *QAbstract3DAxis) DestroyQAbstract3DAxis() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstract3DAxis"})
}

func (ptr *QAbstract3DAxis) DestroyQAbstract3DAxisDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstract3DAxisDefault"})
}

type QAbstract3DGraph struct {
	internal.Internal
}

type QAbstract3DGraph_ITF interface {
	QAbstract3DGraph_PTR() *QAbstract3DGraph
}

func (ptr *QAbstract3DGraph) QAbstract3DGraph_PTR() *QAbstract3DGraph {
	return ptr
}

func (ptr *QAbstract3DGraph) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAbstract3DGraph) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAbstract3DGraph(ptr QAbstract3DGraph_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstract3DGraph_PTR().Pointer()
	}
	return nil
}

func (n *QAbstract3DGraph) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAbstract3DGraphFromPointer(ptr unsafe.Pointer) (n *QAbstract3DGraph) {
	n = new(QAbstract3DGraph)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QAbstract3DGraph")
	return
}

//go:generate stringer -type=QAbstract3DGraph__SelectionFlag
//QAbstract3DGraph::SelectionFlag
type QAbstract3DGraph__SelectionFlag int64

const (
	QAbstract3DGraph__SelectionNone             QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(0)
	QAbstract3DGraph__SelectionItem             QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(1)
	QAbstract3DGraph__SelectionRow              QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(2)
	QAbstract3DGraph__SelectionItemAndRow       QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(QAbstract3DGraph__SelectionItem | QAbstract3DGraph__SelectionRow)
	QAbstract3DGraph__SelectionColumn           QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(4)
	QAbstract3DGraph__SelectionItemAndColumn    QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(QAbstract3DGraph__SelectionItem | QAbstract3DGraph__SelectionColumn)
	QAbstract3DGraph__SelectionRowAndColumn     QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(QAbstract3DGraph__SelectionRow | QAbstract3DGraph__SelectionColumn)
	QAbstract3DGraph__SelectionItemRowAndColumn QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(QAbstract3DGraph__SelectionItem | QAbstract3DGraph__SelectionRow | QAbstract3DGraph__SelectionColumn)
	QAbstract3DGraph__SelectionSlice            QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(8)
	QAbstract3DGraph__SelectionMultiSeries      QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(16)
)

//go:generate stringer -type=QAbstract3DGraph__ShadowQuality
//QAbstract3DGraph::ShadowQuality
type QAbstract3DGraph__ShadowQuality int64

const (
	QAbstract3DGraph__ShadowQualityNone       QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(0)
	QAbstract3DGraph__ShadowQualityLow        QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(1)
	QAbstract3DGraph__ShadowQualityMedium     QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(2)
	QAbstract3DGraph__ShadowQualityHigh       QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(3)
	QAbstract3DGraph__ShadowQualitySoftLow    QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(4)
	QAbstract3DGraph__ShadowQualitySoftMedium QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(5)
	QAbstract3DGraph__ShadowQualitySoftHigh   QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(6)
)

//go:generate stringer -type=QAbstract3DGraph__ElementType
//QAbstract3DGraph::ElementType
type QAbstract3DGraph__ElementType int64

const (
	QAbstract3DGraph__ElementNone       QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(0)
	QAbstract3DGraph__ElementSeries     QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(1)
	QAbstract3DGraph__ElementAxisXLabel QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(2)
	QAbstract3DGraph__ElementAxisYLabel QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(3)
	QAbstract3DGraph__ElementAxisZLabel QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(4)
	QAbstract3DGraph__ElementCustomItem QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(5)
)

//go:generate stringer -type=QAbstract3DGraph__OptimizationHint
//QAbstract3DGraph::OptimizationHint
type QAbstract3DGraph__OptimizationHint int64

const (
	QAbstract3DGraph__OptimizationDefault QAbstract3DGraph__OptimizationHint = QAbstract3DGraph__OptimizationHint(0)
	QAbstract3DGraph__OptimizationStatic  QAbstract3DGraph__OptimizationHint = QAbstract3DGraph__OptimizationHint(1)
)

type QAbstract3DInputHandler struct {
	internal.Internal
}

type QAbstract3DInputHandler_ITF interface {
	QAbstract3DInputHandler_PTR() *QAbstract3DInputHandler
}

func (ptr *QAbstract3DInputHandler) QAbstract3DInputHandler_PTR() *QAbstract3DInputHandler {
	return ptr
}

func (ptr *QAbstract3DInputHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAbstract3DInputHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAbstract3DInputHandler(ptr QAbstract3DInputHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstract3DInputHandler_PTR().Pointer()
	}
	return nil
}

func (n *QAbstract3DInputHandler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAbstract3DInputHandlerFromPointer(ptr unsafe.Pointer) (n *QAbstract3DInputHandler) {
	n = new(QAbstract3DInputHandler)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QAbstract3DInputHandler")
	return
}

//go:generate stringer -type=QAbstract3DInputHandler__InputView
//QAbstract3DInputHandler::InputView
type QAbstract3DInputHandler__InputView int64

const (
	QAbstract3DInputHandler__InputViewNone        QAbstract3DInputHandler__InputView = QAbstract3DInputHandler__InputView(0)
	QAbstract3DInputHandler__InputViewOnPrimary   QAbstract3DInputHandler__InputView = QAbstract3DInputHandler__InputView(1)
	QAbstract3DInputHandler__InputViewOnSecondary QAbstract3DInputHandler__InputView = QAbstract3DInputHandler__InputView(2)
)

func NewQAbstract3DInputHandler(parent core.QObject_ITF) *QAbstract3DInputHandler {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQAbstract3DInputHandler", "", parent}).(*QAbstract3DInputHandler)
}

func (ptr *QAbstract3DInputHandler) InputPosition() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputPosition"}).(*core.QPoint)
}

func (ptr *QAbstract3DInputHandler) InputView() QAbstract3DInputHandler__InputView {

	return QAbstract3DInputHandler__InputView(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputView"}).(float64))
}

func (ptr *QAbstract3DInputHandler) ConnectInputViewChanged(f func(view QAbstract3DInputHandler__InputView)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInputViewChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DInputHandler) DisconnectInputViewChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInputViewChanged"})
}

func (ptr *QAbstract3DInputHandler) InputViewChanged(view QAbstract3DInputHandler__InputView) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputViewChanged", view})
}

func (ptr *QAbstract3DInputHandler) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMouseDoubleClickEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DInputHandler) DisconnectMouseDoubleClickEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMouseDoubleClickEvent"})
}

func (ptr *QAbstract3DInputHandler) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEvent", event})
}

func (ptr *QAbstract3DInputHandler) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QAbstract3DInputHandler) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMouseMoveEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DInputHandler) DisconnectMouseMoveEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMouseMoveEvent"})
}

func (ptr *QAbstract3DInputHandler) MouseMoveEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEvent", event, mousePos})
}

func (ptr *QAbstract3DInputHandler) MouseMoveEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event, mousePos})
}

func (ptr *QAbstract3DInputHandler) ConnectMousePressEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMousePressEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DInputHandler) DisconnectMousePressEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMousePressEvent"})
}

func (ptr *QAbstract3DInputHandler) MousePressEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEvent", event, mousePos})
}

func (ptr *QAbstract3DInputHandler) MousePressEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event, mousePos})
}

func (ptr *QAbstract3DInputHandler) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMouseReleaseEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DInputHandler) DisconnectMouseReleaseEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMouseReleaseEvent"})
}

func (ptr *QAbstract3DInputHandler) MouseReleaseEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEvent", event, mousePos})
}

func (ptr *QAbstract3DInputHandler) MouseReleaseEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event, mousePos})
}

func (ptr *QAbstract3DInputHandler) ConnectPositionChanged(f func(position *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPositionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DInputHandler) DisconnectPositionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPositionChanged"})
}

func (ptr *QAbstract3DInputHandler) PositionChanged(position core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionChanged", position})
}

func (ptr *QAbstract3DInputHandler) PrevDistance() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrevDistance"}).(float64))
}

func (ptr *QAbstract3DInputHandler) PreviousInputPos() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreviousInputPos"}).(*core.QPoint)
}

func (ptr *QAbstract3DInputHandler) Scene() *Q3DScene {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Scene"}).(*Q3DScene)
}

func (ptr *QAbstract3DInputHandler) ConnectSceneChanged(f func(scene *Q3DScene)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSceneChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DInputHandler) DisconnectSceneChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSceneChanged"})
}

func (ptr *QAbstract3DInputHandler) SceneChanged(scene Q3DScene_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneChanged", scene})
}

func (ptr *QAbstract3DInputHandler) SetInputPosition(position core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInputPosition", position})
}

func (ptr *QAbstract3DInputHandler) SetInputView(inputView QAbstract3DInputHandler__InputView) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInputView", inputView})
}

func (ptr *QAbstract3DInputHandler) SetPrevDistance(distance int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrevDistance", distance})
}

func (ptr *QAbstract3DInputHandler) SetPreviousInputPos(position core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPreviousInputPos", position})
}

func (ptr *QAbstract3DInputHandler) SetScene(scene Q3DScene_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScene", scene})
}

func (ptr *QAbstract3DInputHandler) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTouchEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DInputHandler) DisconnectTouchEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTouchEvent"})
}

func (ptr *QAbstract3DInputHandler) TouchEvent(event gui.QTouchEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TouchEvent", event})
}

func (ptr *QAbstract3DInputHandler) TouchEventDefault(event gui.QTouchEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TouchEventDefault", event})
}

func (ptr *QAbstract3DInputHandler) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWheelEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DInputHandler) DisconnectWheelEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWheelEvent"})
}

func (ptr *QAbstract3DInputHandler) WheelEvent(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEvent", event})
}

func (ptr *QAbstract3DInputHandler) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QAbstract3DInputHandler) ConnectDestroyQAbstract3DInputHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstract3DInputHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DInputHandler) DisconnectDestroyQAbstract3DInputHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstract3DInputHandler"})
}

func (ptr *QAbstract3DInputHandler) DestroyQAbstract3DInputHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstract3DInputHandler"})
}

func (ptr *QAbstract3DInputHandler) DestroyQAbstract3DInputHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstract3DInputHandlerDefault"})
}

type QAbstract3DSeries struct {
	internal.Internal
}

type QAbstract3DSeries_ITF interface {
	QAbstract3DSeries_PTR() *QAbstract3DSeries
}

func (ptr *QAbstract3DSeries) QAbstract3DSeries_PTR() *QAbstract3DSeries {
	return ptr
}

func (ptr *QAbstract3DSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAbstract3DSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAbstract3DSeries(ptr QAbstract3DSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstract3DSeries_PTR().Pointer()
	}
	return nil
}

func (n *QAbstract3DSeries) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAbstract3DSeriesFromPointer(ptr unsafe.Pointer) (n *QAbstract3DSeries) {
	n = new(QAbstract3DSeries)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QAbstract3DSeries")
	return
}

//go:generate stringer -type=QAbstract3DSeries__SeriesType
//QAbstract3DSeries::SeriesType
type QAbstract3DSeries__SeriesType int64

const (
	QAbstract3DSeries__SeriesTypeNone    QAbstract3DSeries__SeriesType = QAbstract3DSeries__SeriesType(0)
	QAbstract3DSeries__SeriesTypeBar     QAbstract3DSeries__SeriesType = QAbstract3DSeries__SeriesType(1)
	QAbstract3DSeries__SeriesTypeScatter QAbstract3DSeries__SeriesType = QAbstract3DSeries__SeriesType(2)
	QAbstract3DSeries__SeriesTypeSurface QAbstract3DSeries__SeriesType = QAbstract3DSeries__SeriesType(4)
)

//go:generate stringer -type=QAbstract3DSeries__Mesh
//QAbstract3DSeries::Mesh
type QAbstract3DSeries__Mesh int64

const (
	QAbstract3DSeries__MeshUserDefined QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(0)
	QAbstract3DSeries__MeshBar         QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(1)
	QAbstract3DSeries__MeshCube        QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(2)
	QAbstract3DSeries__MeshPyramid     QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(3)
	QAbstract3DSeries__MeshCone        QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(4)
	QAbstract3DSeries__MeshCylinder    QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(5)
	QAbstract3DSeries__MeshBevelBar    QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(6)
	QAbstract3DSeries__MeshBevelCube   QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(7)
	QAbstract3DSeries__MeshSphere      QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(8)
	QAbstract3DSeries__MeshMinimal     QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(9)
	QAbstract3DSeries__MeshArrow       QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(10)
	QAbstract3DSeries__MeshPoint       QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(11)
)

func (ptr *QAbstract3DSeries) BaseColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseColor"}).(*gui.QColor)
}

func (ptr *QAbstract3DSeries) ConnectBaseColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBaseColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectBaseColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBaseColorChanged"})
}

func (ptr *QAbstract3DSeries) BaseColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseColorChanged", color})
}

func (ptr *QAbstract3DSeries) BaseGradient() *gui.QLinearGradient {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseGradient"}).(*gui.QLinearGradient)
}

func (ptr *QAbstract3DSeries) ConnectBaseGradientChanged(f func(gradient *gui.QLinearGradient)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBaseGradientChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectBaseGradientChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBaseGradientChanged"})
}

func (ptr *QAbstract3DSeries) BaseGradientChanged(gradient gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseGradientChanged", gradient})
}

func (ptr *QAbstract3DSeries) ColorStyle() Q3DTheme__ColorStyle {

	return Q3DTheme__ColorStyle(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorStyle"}).(float64))
}

func (ptr *QAbstract3DSeries) ConnectColorStyleChanged(f func(style Q3DTheme__ColorStyle)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColorStyleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectColorStyleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColorStyleChanged"})
}

func (ptr *QAbstract3DSeries) ColorStyleChanged(style Q3DTheme__ColorStyle) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorStyleChanged", style})
}

func (ptr *QAbstract3DSeries) IsItemLabelVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsItemLabelVisible"}).(bool)
}

func (ptr *QAbstract3DSeries) IsMeshSmooth() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsMeshSmooth"}).(bool)
}

func (ptr *QAbstract3DSeries) IsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsVisible"}).(bool)
}

func (ptr *QAbstract3DSeries) ItemLabel() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemLabel"}).(string)
}

func (ptr *QAbstract3DSeries) ConnectItemLabelChanged(f func(label string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemLabelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectItemLabelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemLabelChanged"})
}

func (ptr *QAbstract3DSeries) ItemLabelChanged(label string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemLabelChanged", label})
}

func (ptr *QAbstract3DSeries) ItemLabelFormat() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemLabelFormat"}).(string)
}

func (ptr *QAbstract3DSeries) ConnectItemLabelFormatChanged(f func(format string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemLabelFormatChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectItemLabelFormatChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemLabelFormatChanged"})
}

func (ptr *QAbstract3DSeries) ItemLabelFormatChanged(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemLabelFormatChanged", format})
}

func (ptr *QAbstract3DSeries) ConnectItemLabelVisibilityChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemLabelVisibilityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectItemLabelVisibilityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemLabelVisibilityChanged"})
}

func (ptr *QAbstract3DSeries) ItemLabelVisibilityChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemLabelVisibilityChanged", visible})
}

func (ptr *QAbstract3DSeries) Mesh() QAbstract3DSeries__Mesh {

	return QAbstract3DSeries__Mesh(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Mesh"}).(float64))
}

func (ptr *QAbstract3DSeries) ConnectMeshChanged(f func(mesh QAbstract3DSeries__Mesh)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMeshChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectMeshChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMeshChanged"})
}

func (ptr *QAbstract3DSeries) MeshChanged(mesh QAbstract3DSeries__Mesh) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MeshChanged", mesh})
}

func (ptr *QAbstract3DSeries) MeshRotation() *gui.QQuaternion {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MeshRotation"}).(*gui.QQuaternion)
}

func (ptr *QAbstract3DSeries) ConnectMeshRotationChanged(f func(rotation *gui.QQuaternion)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMeshRotationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectMeshRotationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMeshRotationChanged"})
}

func (ptr *QAbstract3DSeries) MeshRotationChanged(rotation gui.QQuaternion_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MeshRotationChanged", rotation})
}

func (ptr *QAbstract3DSeries) ConnectMeshSmoothChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMeshSmoothChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectMeshSmoothChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMeshSmoothChanged"})
}

func (ptr *QAbstract3DSeries) MeshSmoothChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MeshSmoothChanged", enabled})
}

func (ptr *QAbstract3DSeries) MultiHighlightColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiHighlightColor"}).(*gui.QColor)
}

func (ptr *QAbstract3DSeries) ConnectMultiHighlightColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMultiHighlightColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectMultiHighlightColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMultiHighlightColorChanged"})
}

func (ptr *QAbstract3DSeries) MultiHighlightColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiHighlightColorChanged", color})
}

func (ptr *QAbstract3DSeries) MultiHighlightGradient() *gui.QLinearGradient {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiHighlightGradient"}).(*gui.QLinearGradient)
}

func (ptr *QAbstract3DSeries) ConnectMultiHighlightGradientChanged(f func(gradient *gui.QLinearGradient)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMultiHighlightGradientChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectMultiHighlightGradientChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMultiHighlightGradientChanged"})
}

func (ptr *QAbstract3DSeries) MultiHighlightGradientChanged(gradient gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiHighlightGradientChanged", gradient})
}

func (ptr *QAbstract3DSeries) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QAbstract3DSeries) ConnectNameChanged(f func(name string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNameChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectNameChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNameChanged"})
}

func (ptr *QAbstract3DSeries) NameChanged(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NameChanged", name})
}

func (ptr *QAbstract3DSeries) SetBaseColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBaseColor", color})
}

func (ptr *QAbstract3DSeries) SetBaseGradient(gradient gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBaseGradient", gradient})
}

func (ptr *QAbstract3DSeries) SetColorStyle(style Q3DTheme__ColorStyle) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColorStyle", style})
}

func (ptr *QAbstract3DSeries) SetItemLabelFormat(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemLabelFormat", format})
}

func (ptr *QAbstract3DSeries) SetItemLabelVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemLabelVisible", visible})
}

func (ptr *QAbstract3DSeries) SetMesh(mesh QAbstract3DSeries__Mesh) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMesh", mesh})
}

func (ptr *QAbstract3DSeries) SetMeshAxisAndAngle(axis gui.QVector3D_ITF, angle float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMeshAxisAndAngle", axis, angle})
}

func (ptr *QAbstract3DSeries) SetMeshRotation(rotation gui.QQuaternion_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMeshRotation", rotation})
}

func (ptr *QAbstract3DSeries) SetMeshSmooth(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMeshSmooth", enable})
}

func (ptr *QAbstract3DSeries) SetMultiHighlightColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMultiHighlightColor", color})
}

func (ptr *QAbstract3DSeries) SetMultiHighlightGradient(gradient gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMultiHighlightGradient", gradient})
}

func (ptr *QAbstract3DSeries) SetName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", name})
}

func (ptr *QAbstract3DSeries) SetSingleHighlightColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSingleHighlightColor", color})
}

func (ptr *QAbstract3DSeries) SetSingleHighlightGradient(gradient gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSingleHighlightGradient", gradient})
}

func (ptr *QAbstract3DSeries) SetUserDefinedMesh(fileName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUserDefinedMesh", fileName})
}

func (ptr *QAbstract3DSeries) SetVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisible", visible})
}

func (ptr *QAbstract3DSeries) SingleHighlightColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SingleHighlightColor"}).(*gui.QColor)
}

func (ptr *QAbstract3DSeries) ConnectSingleHighlightColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSingleHighlightColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectSingleHighlightColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSingleHighlightColorChanged"})
}

func (ptr *QAbstract3DSeries) SingleHighlightColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SingleHighlightColorChanged", color})
}

func (ptr *QAbstract3DSeries) SingleHighlightGradient() *gui.QLinearGradient {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SingleHighlightGradient"}).(*gui.QLinearGradient)
}

func (ptr *QAbstract3DSeries) ConnectSingleHighlightGradientChanged(f func(gradient *gui.QLinearGradient)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSingleHighlightGradientChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectSingleHighlightGradientChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSingleHighlightGradientChanged"})
}

func (ptr *QAbstract3DSeries) SingleHighlightGradientChanged(gradient gui.QLinearGradient_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SingleHighlightGradientChanged", gradient})
}

func (ptr *QAbstract3DSeries) Type() QAbstract3DSeries__SeriesType {

	return QAbstract3DSeries__SeriesType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QAbstract3DSeries) UserDefinedMesh() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UserDefinedMesh"}).(string)
}

func (ptr *QAbstract3DSeries) ConnectUserDefinedMeshChanged(f func(fileName string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUserDefinedMeshChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectUserDefinedMeshChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUserDefinedMeshChanged"})
}

func (ptr *QAbstract3DSeries) UserDefinedMeshChanged(fileName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UserDefinedMeshChanged", fileName})
}

func (ptr *QAbstract3DSeries) ConnectVisibilityChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVisibilityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectVisibilityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVisibilityChanged"})
}

func (ptr *QAbstract3DSeries) VisibilityChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VisibilityChanged", visible})
}

func (ptr *QAbstract3DSeries) ConnectDestroyQAbstract3DSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstract3DSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstract3DSeries) DisconnectDestroyQAbstract3DSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstract3DSeries"})
}

func (ptr *QAbstract3DSeries) DestroyQAbstract3DSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstract3DSeries"})
}

func (ptr *QAbstract3DSeries) DestroyQAbstract3DSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstract3DSeriesDefault"})
}

type QAbstractDataProxy struct {
	internal.Internal
}

type QAbstractDataProxy_ITF interface {
	QAbstractDataProxy_PTR() *QAbstractDataProxy
}

func (ptr *QAbstractDataProxy) QAbstractDataProxy_PTR() *QAbstractDataProxy {
	return ptr
}

func (ptr *QAbstractDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAbstractDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAbstractDataProxy(ptr QAbstractDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractDataProxy_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractDataProxy) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAbstractDataProxyFromPointer(ptr unsafe.Pointer) (n *QAbstractDataProxy) {
	n = new(QAbstractDataProxy)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QAbstractDataProxy")
	return
}

//go:generate stringer -type=QAbstractDataProxy__DataType
//QAbstractDataProxy::DataType
type QAbstractDataProxy__DataType int64

const (
	QAbstractDataProxy__DataTypeNone    QAbstractDataProxy__DataType = QAbstractDataProxy__DataType(0)
	QAbstractDataProxy__DataTypeBar     QAbstractDataProxy__DataType = QAbstractDataProxy__DataType(1)
	QAbstractDataProxy__DataTypeScatter QAbstractDataProxy__DataType = QAbstractDataProxy__DataType(2)
	QAbstractDataProxy__DataTypeSurface QAbstractDataProxy__DataType = QAbstractDataProxy__DataType(4)
)

func (ptr *QAbstractDataProxy) Type() QAbstractDataProxy__DataType {

	return QAbstractDataProxy__DataType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QAbstractDataProxy) ConnectDestroyQAbstractDataProxy(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractDataProxy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractDataProxy) DisconnectDestroyQAbstractDataProxy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractDataProxy"})
}

func (ptr *QAbstractDataProxy) DestroyQAbstractDataProxy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractDataProxy"})
}

func (ptr *QAbstractDataProxy) DestroyQAbstractDataProxyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractDataProxyDefault"})
}

type QBar3DSeries struct {
	internal.Internal
}

type QBar3DSeries_ITF interface {
	QBar3DSeries_PTR() *QBar3DSeries
}

func (ptr *QBar3DSeries) QBar3DSeries_PTR() *QBar3DSeries {
	return ptr
}

func (ptr *QBar3DSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QBar3DSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQBar3DSeries(ptr QBar3DSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBar3DSeries_PTR().Pointer()
	}
	return nil
}

func (n *QBar3DSeries) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQBar3DSeriesFromPointer(ptr unsafe.Pointer) (n *QBar3DSeries) {
	n = new(QBar3DSeries)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QBar3DSeries")
	return
}
func NewQBar3DSeries(parent core.QObject_ITF) *QBar3DSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQBar3DSeries", "", parent}).(*QBar3DSeries)
}

func NewQBar3DSeries2(dataProxy QBarDataProxy_ITF, parent core.QObject_ITF) *QBar3DSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQBar3DSeries2", "", dataProxy, parent}).(*QBar3DSeries)
}

func (ptr *QBar3DSeries) DataProxy() *QBarDataProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataProxy"}).(*QBarDataProxy)
}

func (ptr *QBar3DSeries) ConnectDataProxyChanged(f func(proxy *QBarDataProxy)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDataProxyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBar3DSeries) DisconnectDataProxyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDataProxyChanged"})
}

func (ptr *QBar3DSeries) DataProxyChanged(proxy QBarDataProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataProxyChanged", proxy})
}

func QBar3DSeries_InvalidSelectionPosition() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.QBar3DSeries_InvalidSelectionPosition", ""}).(*core.QPoint)
}

func (ptr *QBar3DSeries) InvalidSelectionPosition() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.QBar3DSeries_InvalidSelectionPosition", ""}).(*core.QPoint)
}

func (ptr *QBar3DSeries) MeshAngle() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MeshAngle"}).(float32)
}

func (ptr *QBar3DSeries) ConnectMeshAngleChanged(f func(angle float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMeshAngleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBar3DSeries) DisconnectMeshAngleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMeshAngleChanged"})
}

func (ptr *QBar3DSeries) MeshAngleChanged(angle float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MeshAngleChanged", angle})
}

func (ptr *QBar3DSeries) SelectedBar() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedBar"}).(*core.QPoint)
}

func (ptr *QBar3DSeries) ConnectSelectedBarChanged(f func(position *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectedBarChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBar3DSeries) DisconnectSelectedBarChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectedBarChanged"})
}

func (ptr *QBar3DSeries) SelectedBarChanged(position core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedBarChanged", position})
}

func (ptr *QBar3DSeries) SetDataProxy(proxy QBarDataProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataProxy", proxy})
}

func (ptr *QBar3DSeries) SetMeshAngle(angle float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMeshAngle", angle})
}

func (ptr *QBar3DSeries) SetSelectedBar(position core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectedBar", position})
}

func (ptr *QBar3DSeries) ConnectDestroyQBar3DSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBar3DSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBar3DSeries) DisconnectDestroyQBar3DSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBar3DSeries"})
}

func (ptr *QBar3DSeries) DestroyQBar3DSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBar3DSeries"})
}

func (ptr *QBar3DSeries) DestroyQBar3DSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBar3DSeriesDefault"})
}

type QBarDataItem struct {
	internal.Internal
}

type QBarDataItem_ITF interface {
	QBarDataItem_PTR() *QBarDataItem
}

func (ptr *QBarDataItem) QBarDataItem_PTR() *QBarDataItem {
	return ptr
}

func (ptr *QBarDataItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QBarDataItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQBarDataItem(ptr QBarDataItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarDataItem_PTR().Pointer()
	}
	return nil
}

func (n *QBarDataItem) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQBarDataItemFromPointer(ptr unsafe.Pointer) (n *QBarDataItem) {
	n = new(QBarDataItem)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QBarDataItem")
	return
}
func NewQBarDataItem() *QBarDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQBarDataItem", ""}).(*QBarDataItem)
}

func NewQBarDataItem2(value float32) *QBarDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQBarDataItem2", "", value}).(*QBarDataItem)
}

func NewQBarDataItem3(value float32, angle float32) *QBarDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQBarDataItem3", "", value, angle}).(*QBarDataItem)
}

func NewQBarDataItem4(other QBarDataItem_ITF) *QBarDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQBarDataItem4", "", other}).(*QBarDataItem)
}

func (ptr *QBarDataItem) Rotation() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rotation"}).(float32)
}

func (ptr *QBarDataItem) SetRotation(angle float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotation", angle})
}

func (ptr *QBarDataItem) SetValue(val float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", val})
}

func (ptr *QBarDataItem) Value() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(float32)
}

func (ptr *QBarDataItem) DestroyQBarDataItem() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarDataItem"})
}

type QBarDataProxy struct {
	internal.Internal
}

type QBarDataProxy_ITF interface {
	QBarDataProxy_PTR() *QBarDataProxy
}

func (ptr *QBarDataProxy) QBarDataProxy_PTR() *QBarDataProxy {
	return ptr
}

func (ptr *QBarDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QBarDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQBarDataProxy(ptr QBarDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarDataProxy_PTR().Pointer()
	}
	return nil
}

func (n *QBarDataProxy) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQBarDataProxyFromPointer(ptr unsafe.Pointer) (n *QBarDataProxy) {
	n = new(QBarDataProxy)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QBarDataProxy")
	return
}
func NewQBarDataProxy(parent core.QObject_ITF) *QBarDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQBarDataProxy", "", parent}).(*QBarDataProxy)
}

func (ptr *QBarDataProxy) ConnectArrayReset(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectArrayReset", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectArrayReset() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectArrayReset"})
}

func (ptr *QBarDataProxy) ArrayReset() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ArrayReset"})
}

func (ptr *QBarDataProxy) ColumnLabels() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnLabels"}).([]string)
}

func (ptr *QBarDataProxy) ConnectColumnLabelsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnLabelsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectColumnLabelsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnLabelsChanged"})
}

func (ptr *QBarDataProxy) ColumnLabelsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnLabelsChanged"})
}

func (ptr *QBarDataProxy) ItemAt(rowIndex int, columnIndex int) *QBarDataItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemAt", rowIndex, columnIndex}).(*QBarDataItem)
}

func (ptr *QBarDataProxy) ItemAt2(position core.QPoint_ITF) *QBarDataItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemAt2", position}).(*QBarDataItem)
}

func (ptr *QBarDataProxy) ConnectItemChanged(f func(rowIndex int, columnIndex int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectItemChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemChanged"})
}

func (ptr *QBarDataProxy) ItemChanged(rowIndex int, columnIndex int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemChanged", rowIndex, columnIndex})
}

func (ptr *QBarDataProxy) RemoveRows(rowIndex int, removeCount int, removeLabels bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveRows", rowIndex, removeCount, removeLabels})
}

func (ptr *QBarDataProxy) ResetArray() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetArray"})
}

func (ptr *QBarDataProxy) RowCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCount"}).(float64))
}

func (ptr *QBarDataProxy) ConnectRowCountChanged(f func(count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectRowCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowCountChanged"})
}

func (ptr *QBarDataProxy) RowCountChanged(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountChanged", count})
}

func (ptr *QBarDataProxy) RowLabels() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowLabels"}).([]string)
}

func (ptr *QBarDataProxy) ConnectRowLabelsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowLabelsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectRowLabelsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowLabelsChanged"})
}

func (ptr *QBarDataProxy) RowLabelsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowLabelsChanged"})
}

func (ptr *QBarDataProxy) ConnectRowsAdded(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowsAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectRowsAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowsAdded"})
}

func (ptr *QBarDataProxy) RowsAdded(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsAdded", startIndex, count})
}

func (ptr *QBarDataProxy) ConnectRowsChanged(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectRowsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowsChanged"})
}

func (ptr *QBarDataProxy) RowsChanged(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsChanged", startIndex, count})
}

func (ptr *QBarDataProxy) ConnectRowsInserted(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowsInserted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectRowsInserted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowsInserted"})
}

func (ptr *QBarDataProxy) RowsInserted(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsInserted", startIndex, count})
}

func (ptr *QBarDataProxy) ConnectRowsRemoved(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowsRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectRowsRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowsRemoved"})
}

func (ptr *QBarDataProxy) RowsRemoved(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsRemoved", startIndex, count})
}

func (ptr *QBarDataProxy) Series() *QBar3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QBar3DSeries)
}

func (ptr *QBarDataProxy) ConnectSeriesChanged(f func(series *QBar3DSeries)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectSeriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesChanged"})
}

func (ptr *QBarDataProxy) SeriesChanged(series QBar3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesChanged", series})
}

func (ptr *QBarDataProxy) SetColumnLabels(labels []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnLabels", labels})
}

func (ptr *QBarDataProxy) SetItem(rowIndex int, columnIndex int, item QBarDataItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItem", rowIndex, columnIndex, item})
}

func (ptr *QBarDataProxy) SetItem2(position core.QPoint_ITF, item QBarDataItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItem2", position, item})
}

func (ptr *QBarDataProxy) SetRowLabels(labels []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowLabels", labels})
}

func (ptr *QBarDataProxy) ConnectDestroyQBarDataProxy(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBarDataProxy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBarDataProxy) DisconnectDestroyQBarDataProxy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBarDataProxy"})
}

func (ptr *QBarDataProxy) DestroyQBarDataProxy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarDataProxy"})
}

func (ptr *QBarDataProxy) DestroyQBarDataProxyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBarDataProxyDefault"})
}

type QCategory3DAxis struct {
	internal.Internal
}

type QCategory3DAxis_ITF interface {
	QCategory3DAxis_PTR() *QCategory3DAxis
}

func (ptr *QCategory3DAxis) QCategory3DAxis_PTR() *QCategory3DAxis {
	return ptr
}

func (ptr *QCategory3DAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QCategory3DAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQCategory3DAxis(ptr QCategory3DAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCategory3DAxis_PTR().Pointer()
	}
	return nil
}

func (n *QCategory3DAxis) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQCategory3DAxisFromPointer(ptr unsafe.Pointer) (n *QCategory3DAxis) {
	n = new(QCategory3DAxis)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QCategory3DAxis")
	return
}
func NewQCategory3DAxis(parent core.QObject_ITF) *QCategory3DAxis {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQCategory3DAxis", "", parent}).(*QCategory3DAxis)
}

func (ptr *QCategory3DAxis) Labels() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Labels"}).([]string)
}

func (ptr *QCategory3DAxis) ConnectLabelsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCategory3DAxis) DisconnectLabelsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelsChanged"})
}

func (ptr *QCategory3DAxis) LabelsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelsChanged"})
}

func (ptr *QCategory3DAxis) SetLabels(labels []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabels", labels})
}

func (ptr *QCategory3DAxis) ConnectDestroyQCategory3DAxis(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQCategory3DAxis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCategory3DAxis) DisconnectDestroyQCategory3DAxis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQCategory3DAxis"})
}

func (ptr *QCategory3DAxis) DestroyQCategory3DAxis() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCategory3DAxis"})
}

func (ptr *QCategory3DAxis) DestroyQCategory3DAxisDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCategory3DAxisDefault"})
}

type QCustom3DItem struct {
	internal.Internal
}

type QCustom3DItem_ITF interface {
	QCustom3DItem_PTR() *QCustom3DItem
}

func (ptr *QCustom3DItem) QCustom3DItem_PTR() *QCustom3DItem {
	return ptr
}

func (ptr *QCustom3DItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QCustom3DItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQCustom3DItem(ptr QCustom3DItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCustom3DItem_PTR().Pointer()
	}
	return nil
}

func (n *QCustom3DItem) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQCustom3DItemFromPointer(ptr unsafe.Pointer) (n *QCustom3DItem) {
	n = new(QCustom3DItem)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QCustom3DItem")
	return
}
func NewQCustom3DItem(parent core.QObject_ITF) *QCustom3DItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQCustom3DItem", "", parent}).(*QCustom3DItem)
}

func NewQCustom3DItem2(meshFile string, position gui.QVector3D_ITF, scaling gui.QVector3D_ITF, rotation gui.QQuaternion_ITF, texture gui.QImage_ITF, parent core.QObject_ITF) *QCustom3DItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQCustom3DItem2", "", meshFile, position, scaling, rotation, texture, parent}).(*QCustom3DItem)
}

func (ptr *QCustom3DItem) IsPositionAbsolute() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsPositionAbsolute"}).(bool)
}

func (ptr *QCustom3DItem) IsScalingAbsolute() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsScalingAbsolute"}).(bool)
}

func (ptr *QCustom3DItem) IsShadowCasting() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsShadowCasting"}).(bool)
}

func (ptr *QCustom3DItem) IsVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsVisible"}).(bool)
}

func (ptr *QCustom3DItem) MeshFile() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MeshFile"}).(string)
}

func (ptr *QCustom3DItem) ConnectMeshFileChanged(f func(meshFile string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMeshFileChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DItem) DisconnectMeshFileChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMeshFileChanged"})
}

func (ptr *QCustom3DItem) MeshFileChanged(meshFile string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MeshFileChanged", meshFile})
}

func (ptr *QCustom3DItem) Position() *gui.QVector3D {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Position"}).(*gui.QVector3D)
}

func (ptr *QCustom3DItem) ConnectPositionAbsoluteChanged(f func(positionAbsolute bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPositionAbsoluteChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DItem) DisconnectPositionAbsoluteChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPositionAbsoluteChanged"})
}

func (ptr *QCustom3DItem) PositionAbsoluteChanged(positionAbsolute bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionAbsoluteChanged", positionAbsolute})
}

func (ptr *QCustom3DItem) ConnectPositionChanged(f func(position *gui.QVector3D)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPositionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DItem) DisconnectPositionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPositionChanged"})
}

func (ptr *QCustom3DItem) PositionChanged(position gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionChanged", position})
}

func (ptr *QCustom3DItem) Rotation() *gui.QQuaternion {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rotation"}).(*gui.QQuaternion)
}

func (ptr *QCustom3DItem) ConnectRotationChanged(f func(rotation *gui.QQuaternion)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRotationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DItem) DisconnectRotationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRotationChanged"})
}

func (ptr *QCustom3DItem) RotationChanged(rotation gui.QQuaternion_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationChanged", rotation})
}

func (ptr *QCustom3DItem) Scaling() *gui.QVector3D {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Scaling"}).(*gui.QVector3D)
}

func (ptr *QCustom3DItem) ConnectScalingAbsoluteChanged(f func(scalingAbsolute bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectScalingAbsoluteChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DItem) DisconnectScalingAbsoluteChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectScalingAbsoluteChanged"})
}

func (ptr *QCustom3DItem) ScalingAbsoluteChanged(scalingAbsolute bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScalingAbsoluteChanged", scalingAbsolute})
}

func (ptr *QCustom3DItem) ConnectScalingChanged(f func(scaling *gui.QVector3D)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectScalingChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DItem) DisconnectScalingChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectScalingChanged"})
}

func (ptr *QCustom3DItem) ScalingChanged(scaling gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScalingChanged", scaling})
}

func (ptr *QCustom3DItem) SetMeshFile(meshFile string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMeshFile", meshFile})
}

func (ptr *QCustom3DItem) SetPosition(position gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPosition", position})
}

func (ptr *QCustom3DItem) SetPositionAbsolute(positionAbsolute bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPositionAbsolute", positionAbsolute})
}

func (ptr *QCustom3DItem) SetRotation(rotation gui.QQuaternion_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotation", rotation})
}

func (ptr *QCustom3DItem) SetRotationAxisAndAngle(axis gui.QVector3D_ITF, angle float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotationAxisAndAngle", axis, angle})
}

func (ptr *QCustom3DItem) SetScaling(scaling gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScaling", scaling})
}

func (ptr *QCustom3DItem) SetScalingAbsolute(scalingAbsolute bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScalingAbsolute", scalingAbsolute})
}

func (ptr *QCustom3DItem) SetShadowCasting(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShadowCasting", enabled})
}

func (ptr *QCustom3DItem) SetTextureFile(textureFile string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureFile", textureFile})
}

func (ptr *QCustom3DItem) SetTextureImage(textureImage gui.QImage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureImage", textureImage})
}

func (ptr *QCustom3DItem) SetVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisible", visible})
}

func (ptr *QCustom3DItem) ConnectShadowCastingChanged(f func(shadowCasting bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShadowCastingChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DItem) DisconnectShadowCastingChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShadowCastingChanged"})
}

func (ptr *QCustom3DItem) ShadowCastingChanged(shadowCasting bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShadowCastingChanged", shadowCasting})
}

func (ptr *QCustom3DItem) TextureFile() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureFile"}).(string)
}

func (ptr *QCustom3DItem) ConnectTextureFileChanged(f func(textureFile string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureFileChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DItem) DisconnectTextureFileChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureFileChanged"})
}

func (ptr *QCustom3DItem) TextureFileChanged(textureFile string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureFileChanged", textureFile})
}

func (ptr *QCustom3DItem) ConnectVisibleChanged(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DItem) DisconnectVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVisibleChanged"})
}

func (ptr *QCustom3DItem) VisibleChanged(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VisibleChanged", visible})
}

func (ptr *QCustom3DItem) ConnectDestroyQCustom3DItem(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQCustom3DItem", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DItem) DisconnectDestroyQCustom3DItem() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQCustom3DItem"})
}

func (ptr *QCustom3DItem) DestroyQCustom3DItem() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCustom3DItem"})
}

func (ptr *QCustom3DItem) DestroyQCustom3DItemDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCustom3DItemDefault"})
}

type QCustom3DLabel struct {
	internal.Internal
}

type QCustom3DLabel_ITF interface {
	QCustom3DLabel_PTR() *QCustom3DLabel
}

func (ptr *QCustom3DLabel) QCustom3DLabel_PTR() *QCustom3DLabel {
	return ptr
}

func (ptr *QCustom3DLabel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QCustom3DLabel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQCustom3DLabel(ptr QCustom3DLabel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCustom3DLabel_PTR().Pointer()
	}
	return nil
}

func (n *QCustom3DLabel) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQCustom3DLabelFromPointer(ptr unsafe.Pointer) (n *QCustom3DLabel) {
	n = new(QCustom3DLabel)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QCustom3DLabel")
	return
}
func NewQCustom3DLabel(parent core.QObject_ITF) *QCustom3DLabel {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQCustom3DLabel", "", parent}).(*QCustom3DLabel)
}

func NewQCustom3DLabel2(text string, font gui.QFont_ITF, position gui.QVector3D_ITF, scaling gui.QVector3D_ITF, rotation gui.QQuaternion_ITF, parent core.QObject_ITF) *QCustom3DLabel {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQCustom3DLabel2", "", text, font, position, scaling, rotation, parent}).(*QCustom3DLabel)
}

func (ptr *QCustom3DLabel) BackgroundColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundColor"}).(*gui.QColor)
}

func (ptr *QCustom3DLabel) ConnectBackgroundColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBackgroundColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DLabel) DisconnectBackgroundColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBackgroundColorChanged"})
}

func (ptr *QCustom3DLabel) BackgroundColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundColorChanged", color})
}

func (ptr *QCustom3DLabel) ConnectBackgroundEnabledChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBackgroundEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DLabel) DisconnectBackgroundEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBackgroundEnabledChanged"})
}

func (ptr *QCustom3DLabel) BackgroundEnabledChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackgroundEnabledChanged", enabled})
}

func (ptr *QCustom3DLabel) ConnectBorderEnabledChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBorderEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DLabel) DisconnectBorderEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBorderEnabledChanged"})
}

func (ptr *QCustom3DLabel) BorderEnabledChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BorderEnabledChanged", enabled})
}

func (ptr *QCustom3DLabel) ConnectFacingCameraChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFacingCameraChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DLabel) DisconnectFacingCameraChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFacingCameraChanged"})
}

func (ptr *QCustom3DLabel) FacingCameraChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FacingCameraChanged", enabled})
}

func (ptr *QCustom3DLabel) Font() *gui.QFont {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Font"}).(*gui.QFont)
}

func (ptr *QCustom3DLabel) ConnectFontChanged(f func(font *gui.QFont)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFontChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DLabel) DisconnectFontChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFontChanged"})
}

func (ptr *QCustom3DLabel) FontChanged(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FontChanged", font})
}

func (ptr *QCustom3DLabel) IsBackgroundEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBackgroundEnabled"}).(bool)
}

func (ptr *QCustom3DLabel) IsBorderEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBorderEnabled"}).(bool)
}

func (ptr *QCustom3DLabel) IsFacingCamera() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFacingCamera"}).(bool)
}

func (ptr *QCustom3DLabel) SetBackgroundColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackgroundColor", color})
}

func (ptr *QCustom3DLabel) SetBackgroundEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackgroundEnabled", enabled})
}

func (ptr *QCustom3DLabel) SetBorderEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBorderEnabled", enabled})
}

func (ptr *QCustom3DLabel) SetFacingCamera(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFacingCamera", enabled})
}

func (ptr *QCustom3DLabel) SetFont(font gui.QFont_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFont", font})
}

func (ptr *QCustom3DLabel) SetText(text string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetText", text})
}

func (ptr *QCustom3DLabel) SetTextColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextColor", color})
}

func (ptr *QCustom3DLabel) Text() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Text"}).(string)
}

func (ptr *QCustom3DLabel) ConnectTextChanged(f func(text string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DLabel) DisconnectTextChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextChanged"})
}

func (ptr *QCustom3DLabel) TextChanged(text string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextChanged", text})
}

func (ptr *QCustom3DLabel) TextColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextColor"}).(*gui.QColor)
}

func (ptr *QCustom3DLabel) ConnectTextColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DLabel) DisconnectTextColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextColorChanged"})
}

func (ptr *QCustom3DLabel) TextColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextColorChanged", color})
}

func (ptr *QCustom3DLabel) ConnectDestroyQCustom3DLabel(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQCustom3DLabel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DLabel) DisconnectDestroyQCustom3DLabel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQCustom3DLabel"})
}

func (ptr *QCustom3DLabel) DestroyQCustom3DLabel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCustom3DLabel"})
}

func (ptr *QCustom3DLabel) DestroyQCustom3DLabelDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCustom3DLabelDefault"})
}

type QCustom3DVolume struct {
	internal.Internal
}

type QCustom3DVolume_ITF interface {
	QCustom3DVolume_PTR() *QCustom3DVolume
}

func (ptr *QCustom3DVolume) QCustom3DVolume_PTR() *QCustom3DVolume {
	return ptr
}

func (ptr *QCustom3DVolume) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QCustom3DVolume) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQCustom3DVolume(ptr QCustom3DVolume_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCustom3DVolume_PTR().Pointer()
	}
	return nil
}

func (n *QCustom3DVolume) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQCustom3DVolumeFromPointer(ptr unsafe.Pointer) (n *QCustom3DVolume) {
	n = new(QCustom3DVolume)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QCustom3DVolume")
	return
}
func NewQCustom3DVolume(parent core.QObject_ITF) *QCustom3DVolume {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQCustom3DVolume", "", parent}).(*QCustom3DVolume)
}

func (ptr *QCustom3DVolume) AlphaMultiplier() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AlphaMultiplier"}).(float32)
}

func (ptr *QCustom3DVolume) ConnectAlphaMultiplierChanged(f func(mult float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAlphaMultiplierChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectAlphaMultiplierChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAlphaMultiplierChanged"})
}

func (ptr *QCustom3DVolume) AlphaMultiplierChanged(mult float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AlphaMultiplierChanged", mult})
}

func (ptr *QCustom3DVolume) ColorTable() []uint {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorTable"}).([]uint)
}

func (ptr *QCustom3DVolume) ConnectColorTableChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColorTableChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectColorTableChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColorTableChanged"})
}

func (ptr *QCustom3DVolume) ColorTableChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorTableChanged"})
}

func (ptr *QCustom3DVolume) DrawSliceFrames() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawSliceFrames"}).(bool)
}

func (ptr *QCustom3DVolume) ConnectDrawSliceFramesChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDrawSliceFramesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectDrawSliceFramesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDrawSliceFramesChanged"})
}

func (ptr *QCustom3DVolume) DrawSliceFramesChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawSliceFramesChanged", enabled})
}

func (ptr *QCustom3DVolume) DrawSlices() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawSlices"}).(bool)
}

func (ptr *QCustom3DVolume) ConnectDrawSlicesChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDrawSlicesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectDrawSlicesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDrawSlicesChanged"})
}

func (ptr *QCustom3DVolume) DrawSlicesChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawSlicesChanged", enabled})
}

func (ptr *QCustom3DVolume) PreserveOpacity() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreserveOpacity"}).(bool)
}

func (ptr *QCustom3DVolume) ConnectPreserveOpacityChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreserveOpacityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectPreserveOpacityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreserveOpacityChanged"})
}

func (ptr *QCustom3DVolume) PreserveOpacityChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreserveOpacityChanged", enabled})
}

func (ptr *QCustom3DVolume) RenderSlice(axis core.Qt__Axis, index int) *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderSlice", axis, index}).(*gui.QImage)
}

func (ptr *QCustom3DVolume) SetAlphaMultiplier(mult float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAlphaMultiplier", mult})
}

func (ptr *QCustom3DVolume) SetColorTable(colors []uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColorTable", colors})
}

func (ptr *QCustom3DVolume) SetDrawSliceFrames(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDrawSliceFrames", enable})
}

func (ptr *QCustom3DVolume) SetDrawSlices(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDrawSlices", enable})
}

func (ptr *QCustom3DVolume) SetPreserveOpacity(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPreserveOpacity", enable})
}

func (ptr *QCustom3DVolume) SetSliceFrameColor(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSliceFrameColor", color})
}

func (ptr *QCustom3DVolume) SetSliceFrameGaps(values gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSliceFrameGaps", values})
}

func (ptr *QCustom3DVolume) SetSliceFrameThicknesses(values gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSliceFrameThicknesses", values})
}

func (ptr *QCustom3DVolume) SetSliceFrameWidths(values gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSliceFrameWidths", values})
}

func (ptr *QCustom3DVolume) SetSliceIndexX(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSliceIndexX", value})
}

func (ptr *QCustom3DVolume) SetSliceIndexY(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSliceIndexY", value})
}

func (ptr *QCustom3DVolume) SetSliceIndexZ(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSliceIndexZ", value})
}

func (ptr *QCustom3DVolume) SetSliceIndices(x int, y int, z int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSliceIndices", x, y, z})
}

func (ptr *QCustom3DVolume) SetSubTextureData(axis core.Qt__Axis, index int, data string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSubTextureData", axis, index, data})
}

func (ptr *QCustom3DVolume) SetSubTextureData2(axis core.Qt__Axis, index int, image gui.QImage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSubTextureData2", axis, index, image})
}

func (ptr *QCustom3DVolume) SetTextureDepth(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureDepth", value})
}

func (ptr *QCustom3DVolume) SetTextureDimensions(width int, height int, depth int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureDimensions", width, height, depth})
}

func (ptr *QCustom3DVolume) SetTextureFormat(format gui.QImage__Format) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureFormat", format})
}

func (ptr *QCustom3DVolume) SetTextureHeight(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureHeight", value})
}

func (ptr *QCustom3DVolume) SetTextureWidth(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureWidth", value})
}

func (ptr *QCustom3DVolume) SetUseHighDefShader(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUseHighDefShader", enable})
}

func (ptr *QCustom3DVolume) SliceFrameColor() *gui.QColor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceFrameColor"}).(*gui.QColor)
}

func (ptr *QCustom3DVolume) ConnectSliceFrameColorChanged(f func(color *gui.QColor)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSliceFrameColorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectSliceFrameColorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSliceFrameColorChanged"})
}

func (ptr *QCustom3DVolume) SliceFrameColorChanged(color gui.QColor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceFrameColorChanged", color})
}

func (ptr *QCustom3DVolume) SliceFrameGaps() *gui.QVector3D {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceFrameGaps"}).(*gui.QVector3D)
}

func (ptr *QCustom3DVolume) ConnectSliceFrameGapsChanged(f func(values *gui.QVector3D)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSliceFrameGapsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectSliceFrameGapsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSliceFrameGapsChanged"})
}

func (ptr *QCustom3DVolume) SliceFrameGapsChanged(values gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceFrameGapsChanged", values})
}

func (ptr *QCustom3DVolume) SliceFrameThicknesses() *gui.QVector3D {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceFrameThicknesses"}).(*gui.QVector3D)
}

func (ptr *QCustom3DVolume) ConnectSliceFrameThicknessesChanged(f func(values *gui.QVector3D)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSliceFrameThicknessesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectSliceFrameThicknessesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSliceFrameThicknessesChanged"})
}

func (ptr *QCustom3DVolume) SliceFrameThicknessesChanged(values gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceFrameThicknessesChanged", values})
}

func (ptr *QCustom3DVolume) SliceFrameWidths() *gui.QVector3D {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceFrameWidths"}).(*gui.QVector3D)
}

func (ptr *QCustom3DVolume) ConnectSliceFrameWidthsChanged(f func(values *gui.QVector3D)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSliceFrameWidthsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectSliceFrameWidthsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSliceFrameWidthsChanged"})
}

func (ptr *QCustom3DVolume) SliceFrameWidthsChanged(values gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceFrameWidthsChanged", values})
}

func (ptr *QCustom3DVolume) SliceIndexX() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceIndexX"}).(float64))
}

func (ptr *QCustom3DVolume) ConnectSliceIndexXChanged(f func(value int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSliceIndexXChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectSliceIndexXChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSliceIndexXChanged"})
}

func (ptr *QCustom3DVolume) SliceIndexXChanged(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceIndexXChanged", value})
}

func (ptr *QCustom3DVolume) SliceIndexY() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceIndexY"}).(float64))
}

func (ptr *QCustom3DVolume) ConnectSliceIndexYChanged(f func(value int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSliceIndexYChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectSliceIndexYChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSliceIndexYChanged"})
}

func (ptr *QCustom3DVolume) SliceIndexYChanged(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceIndexYChanged", value})
}

func (ptr *QCustom3DVolume) SliceIndexZ() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceIndexZ"}).(float64))
}

func (ptr *QCustom3DVolume) ConnectSliceIndexZChanged(f func(value int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSliceIndexZChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectSliceIndexZChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSliceIndexZChanged"})
}

func (ptr *QCustom3DVolume) SliceIndexZChanged(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SliceIndexZChanged", value})
}

func (ptr *QCustom3DVolume) TextureDepth() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureDepth"}).(float64))
}

func (ptr *QCustom3DVolume) ConnectTextureDepthChanged(f func(value int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureDepthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectTextureDepthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureDepthChanged"})
}

func (ptr *QCustom3DVolume) TextureDepthChanged(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureDepthChanged", value})
}

func (ptr *QCustom3DVolume) TextureFormat() gui.QImage__Format {

	return gui.QImage__Format(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureFormat"}).(float64))
}

func (ptr *QCustom3DVolume) ConnectTextureFormatChanged(f func(format gui.QImage__Format)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureFormatChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectTextureFormatChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureFormatChanged"})
}

func (ptr *QCustom3DVolume) TextureFormatChanged(format gui.QImage__Format) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureFormatChanged", format})
}

func (ptr *QCustom3DVolume) TextureHeight() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureHeight"}).(float64))
}

func (ptr *QCustom3DVolume) ConnectTextureHeightChanged(f func(value int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureHeightChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectTextureHeightChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureHeightChanged"})
}

func (ptr *QCustom3DVolume) TextureHeightChanged(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureHeightChanged", value})
}

func (ptr *QCustom3DVolume) TextureWidth() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureWidth"}).(float64))
}

func (ptr *QCustom3DVolume) ConnectTextureWidthChanged(f func(value int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureWidthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectTextureWidthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureWidthChanged"})
}

func (ptr *QCustom3DVolume) TextureWidthChanged(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureWidthChanged", value})
}

func (ptr *QCustom3DVolume) UseHighDefShader() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UseHighDefShader"}).(bool)
}

func (ptr *QCustom3DVolume) ConnectUseHighDefShaderChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUseHighDefShaderChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectUseHighDefShaderChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUseHighDefShaderChanged"})
}

func (ptr *QCustom3DVolume) UseHighDefShaderChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UseHighDefShaderChanged", enabled})
}

func (ptr *QCustom3DVolume) ConnectDestroyQCustom3DVolume(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQCustom3DVolume", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCustom3DVolume) DisconnectDestroyQCustom3DVolume() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQCustom3DVolume"})
}

func (ptr *QCustom3DVolume) DestroyQCustom3DVolume() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCustom3DVolume"})
}

func (ptr *QCustom3DVolume) DestroyQCustom3DVolumeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCustom3DVolumeDefault"})
}

func (ptr *QCustom3DVolume) __QCustom3DVolume_colorTable_atList2(i int) uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QCustom3DVolume_colorTable_atList2", i}).(float64))
}

func (ptr *QCustom3DVolume) __QCustom3DVolume_colorTable_setList2(i uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QCustom3DVolume_colorTable_setList2", i})
}

func (ptr *QCustom3DVolume) __QCustom3DVolume_colorTable_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QCustom3DVolume_colorTable_newList2"}).(unsafe.Pointer)
}

func (ptr *QCustom3DVolume) __colorTable_atList(i int) uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__colorTable_atList", i}).(float64))
}

func (ptr *QCustom3DVolume) __colorTable_setList(i uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__colorTable_setList", i})
}

func (ptr *QCustom3DVolume) __colorTable_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__colorTable_newList"}).(unsafe.Pointer)
}

func (ptr *QCustom3DVolume) __setColorTable_colors_atList(i int) uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setColorTable_colors_atList", i}).(float64))
}

func (ptr *QCustom3DVolume) __setColorTable_colors_setList(i uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setColorTable_colors_setList", i})
}

func (ptr *QCustom3DVolume) __setColorTable_colors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setColorTable_colors_newList"}).(unsafe.Pointer)
}

type QHeightMapSurfaceDataProxy struct {
	internal.Internal
}

type QHeightMapSurfaceDataProxy_ITF interface {
	QHeightMapSurfaceDataProxy_PTR() *QHeightMapSurfaceDataProxy
}

func (ptr *QHeightMapSurfaceDataProxy) QHeightMapSurfaceDataProxy_PTR() *QHeightMapSurfaceDataProxy {
	return ptr
}

func (ptr *QHeightMapSurfaceDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QHeightMapSurfaceDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQHeightMapSurfaceDataProxy(ptr QHeightMapSurfaceDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHeightMapSurfaceDataProxy_PTR().Pointer()
	}
	return nil
}

func (n *QHeightMapSurfaceDataProxy) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQHeightMapSurfaceDataProxyFromPointer(ptr unsafe.Pointer) (n *QHeightMapSurfaceDataProxy) {
	n = new(QHeightMapSurfaceDataProxy)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QHeightMapSurfaceDataProxy")
	return
}
func NewQHeightMapSurfaceDataProxy(parent core.QObject_ITF) *QHeightMapSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQHeightMapSurfaceDataProxy", "", parent}).(*QHeightMapSurfaceDataProxy)
}

func NewQHeightMapSurfaceDataProxy2(image gui.QImage_ITF, parent core.QObject_ITF) *QHeightMapSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQHeightMapSurfaceDataProxy2", "", image, parent}).(*QHeightMapSurfaceDataProxy)
}

func NewQHeightMapSurfaceDataProxy3(filename string, parent core.QObject_ITF) *QHeightMapSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQHeightMapSurfaceDataProxy3", "", filename, parent}).(*QHeightMapSurfaceDataProxy)
}

func (ptr *QHeightMapSurfaceDataProxy) HeightMap() *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightMap"}).(*gui.QImage)
}

func (ptr *QHeightMapSurfaceDataProxy) ConnectHeightMapChanged(f func(image *gui.QImage)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHeightMapChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectHeightMapChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHeightMapChanged"})
}

func (ptr *QHeightMapSurfaceDataProxy) HeightMapChanged(image gui.QImage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightMapChanged", image})
}

func (ptr *QHeightMapSurfaceDataProxy) HeightMapFile() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightMapFile"}).(string)
}

func (ptr *QHeightMapSurfaceDataProxy) ConnectHeightMapFileChanged(f func(filename string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHeightMapFileChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectHeightMapFileChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHeightMapFileChanged"})
}

func (ptr *QHeightMapSurfaceDataProxy) HeightMapFileChanged(filename string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightMapFileChanged", filename})
}

func (ptr *QHeightMapSurfaceDataProxy) MaxXValue() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxXValue"}).(float32)
}

func (ptr *QHeightMapSurfaceDataProxy) ConnectMaxXValueChanged(f func(value float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMaxXValueChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectMaxXValueChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMaxXValueChanged"})
}

func (ptr *QHeightMapSurfaceDataProxy) MaxXValueChanged(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxXValueChanged", value})
}

func (ptr *QHeightMapSurfaceDataProxy) MaxZValue() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxZValue"}).(float32)
}

func (ptr *QHeightMapSurfaceDataProxy) ConnectMaxZValueChanged(f func(value float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMaxZValueChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectMaxZValueChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMaxZValueChanged"})
}

func (ptr *QHeightMapSurfaceDataProxy) MaxZValueChanged(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxZValueChanged", value})
}

func (ptr *QHeightMapSurfaceDataProxy) MinXValue() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinXValue"}).(float32)
}

func (ptr *QHeightMapSurfaceDataProxy) ConnectMinXValueChanged(f func(value float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinXValueChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectMinXValueChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinXValueChanged"})
}

func (ptr *QHeightMapSurfaceDataProxy) MinXValueChanged(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinXValueChanged", value})
}

func (ptr *QHeightMapSurfaceDataProxy) MinZValue() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinZValue"}).(float32)
}

func (ptr *QHeightMapSurfaceDataProxy) ConnectMinZValueChanged(f func(value float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinZValueChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectMinZValueChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinZValueChanged"})
}

func (ptr *QHeightMapSurfaceDataProxy) MinZValueChanged(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinZValueChanged", value})
}

func (ptr *QHeightMapSurfaceDataProxy) SetHeightMap(image gui.QImage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeightMap", image})
}

func (ptr *QHeightMapSurfaceDataProxy) SetHeightMapFile(filename string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeightMapFile", filename})
}

func (ptr *QHeightMapSurfaceDataProxy) SetMaxXValue(max float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaxXValue", max})
}

func (ptr *QHeightMapSurfaceDataProxy) SetMaxZValue(max float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaxZValue", max})
}

func (ptr *QHeightMapSurfaceDataProxy) SetMinXValue(min float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinXValue", min})
}

func (ptr *QHeightMapSurfaceDataProxy) SetMinZValue(min float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinZValue", min})
}

func (ptr *QHeightMapSurfaceDataProxy) SetValueRanges(minX float32, maxX float32, minZ float32, maxZ float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValueRanges", minX, maxX, minZ, maxZ})
}

func (ptr *QHeightMapSurfaceDataProxy) ConnectDestroyQHeightMapSurfaceDataProxy(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHeightMapSurfaceDataProxy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectDestroyQHeightMapSurfaceDataProxy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHeightMapSurfaceDataProxy"})
}

func (ptr *QHeightMapSurfaceDataProxy) DestroyQHeightMapSurfaceDataProxy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHeightMapSurfaceDataProxy"})
}

func (ptr *QHeightMapSurfaceDataProxy) DestroyQHeightMapSurfaceDataProxyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHeightMapSurfaceDataProxyDefault"})
}

type QItemModelBarDataProxy struct {
	internal.Internal
}

type QItemModelBarDataProxy_ITF interface {
	QItemModelBarDataProxy_PTR() *QItemModelBarDataProxy
}

func (ptr *QItemModelBarDataProxy) QItemModelBarDataProxy_PTR() *QItemModelBarDataProxy {
	return ptr
}

func (ptr *QItemModelBarDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QItemModelBarDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQItemModelBarDataProxy(ptr QItemModelBarDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemModelBarDataProxy_PTR().Pointer()
	}
	return nil
}

func (n *QItemModelBarDataProxy) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQItemModelBarDataProxyFromPointer(ptr unsafe.Pointer) (n *QItemModelBarDataProxy) {
	n = new(QItemModelBarDataProxy)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QItemModelBarDataProxy")
	return
}

//go:generate stringer -type=QItemModelBarDataProxy__MultiMatchBehavior
//QItemModelBarDataProxy::MultiMatchBehavior
type QItemModelBarDataProxy__MultiMatchBehavior int64

const (
	QItemModelBarDataProxy__MMBFirst      QItemModelBarDataProxy__MultiMatchBehavior = QItemModelBarDataProxy__MultiMatchBehavior(0)
	QItemModelBarDataProxy__MMBLast       QItemModelBarDataProxy__MultiMatchBehavior = QItemModelBarDataProxy__MultiMatchBehavior(1)
	QItemModelBarDataProxy__MMBAverage    QItemModelBarDataProxy__MultiMatchBehavior = QItemModelBarDataProxy__MultiMatchBehavior(2)
	QItemModelBarDataProxy__MMBCumulative QItemModelBarDataProxy__MultiMatchBehavior = QItemModelBarDataProxy__MultiMatchBehavior(3)
)

func NewQItemModelBarDataProxy(parent core.QObject_ITF) *QItemModelBarDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelBarDataProxy", "", parent}).(*QItemModelBarDataProxy)
}

func NewQItemModelBarDataProxy2(itemModel core.QAbstractItemModel_ITF, parent core.QObject_ITF) *QItemModelBarDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelBarDataProxy2", "", itemModel, parent}).(*QItemModelBarDataProxy)
}

func NewQItemModelBarDataProxy3(itemModel core.QAbstractItemModel_ITF, valueRole string, parent core.QObject_ITF) *QItemModelBarDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelBarDataProxy3", "", itemModel, valueRole, parent}).(*QItemModelBarDataProxy)
}

func NewQItemModelBarDataProxy4(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, valueRole string, parent core.QObject_ITF) *QItemModelBarDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelBarDataProxy4", "", itemModel, rowRole, columnRole, valueRole, parent}).(*QItemModelBarDataProxy)
}

func NewQItemModelBarDataProxy5(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, valueRole string, rotationRole string, parent core.QObject_ITF) *QItemModelBarDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelBarDataProxy5", "", itemModel, rowRole, columnRole, valueRole, rotationRole, parent}).(*QItemModelBarDataProxy)
}

func NewQItemModelBarDataProxy6(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, valueRole string, rowCategories []string, columnCategories []string, parent core.QObject_ITF) *QItemModelBarDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelBarDataProxy6", "", itemModel, rowRole, columnRole, valueRole, rowCategories, columnCategories, parent}).(*QItemModelBarDataProxy)
}

func NewQItemModelBarDataProxy7(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, valueRole string, rotationRole string, rowCategories []string, columnCategories []string, parent core.QObject_ITF) *QItemModelBarDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelBarDataProxy7", "", itemModel, rowRole, columnRole, valueRole, rotationRole, rowCategories, columnCategories, parent}).(*QItemModelBarDataProxy)
}

func (ptr *QItemModelBarDataProxy) AutoColumnCategories() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoColumnCategories"}).(bool)
}

func (ptr *QItemModelBarDataProxy) ConnectAutoColumnCategoriesChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAutoColumnCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectAutoColumnCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAutoColumnCategoriesChanged"})
}

func (ptr *QItemModelBarDataProxy) AutoColumnCategoriesChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoColumnCategoriesChanged", enable})
}

func (ptr *QItemModelBarDataProxy) AutoRowCategories() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoRowCategories"}).(bool)
}

func (ptr *QItemModelBarDataProxy) ConnectAutoRowCategoriesChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAutoRowCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectAutoRowCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAutoRowCategoriesChanged"})
}

func (ptr *QItemModelBarDataProxy) AutoRowCategoriesChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoRowCategoriesChanged", enable})
}

func (ptr *QItemModelBarDataProxy) ColumnCategories() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCategories"}).([]string)
}

func (ptr *QItemModelBarDataProxy) ConnectColumnCategoriesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectColumnCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnCategoriesChanged"})
}

func (ptr *QItemModelBarDataProxy) ColumnCategoriesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCategoriesChanged"})
}

func (ptr *QItemModelBarDataProxy) ColumnCategoryIndex(category string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCategoryIndex", category}).(float64))
}

func (ptr *QItemModelBarDataProxy) ColumnRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRole"}).(string)
}

func (ptr *QItemModelBarDataProxy) ConnectColumnRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectColumnRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnRoleChanged"})
}

func (ptr *QItemModelBarDataProxy) ColumnRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRoleChanged", role})
}

func (ptr *QItemModelBarDataProxy) ColumnRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelBarDataProxy) ConnectColumnRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectColumnRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnRolePatternChanged"})
}

func (ptr *QItemModelBarDataProxy) ColumnRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRolePatternChanged", pattern})
}

func (ptr *QItemModelBarDataProxy) ColumnRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRoleReplace"}).(string)
}

func (ptr *QItemModelBarDataProxy) ConnectColumnRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectColumnRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnRoleReplaceChanged"})
}

func (ptr *QItemModelBarDataProxy) ColumnRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRoleReplaceChanged", replace})
}

func (ptr *QItemModelBarDataProxy) ItemModel() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemModel"}).(*core.QAbstractItemModel)
}

func (ptr *QItemModelBarDataProxy) ConnectItemModelChanged(f func(itemModel *core.QAbstractItemModel)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemModelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectItemModelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemModelChanged"})
}

func (ptr *QItemModelBarDataProxy) ItemModelChanged(itemModel core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemModelChanged", itemModel})
}

func (ptr *QItemModelBarDataProxy) MultiMatchBehavior() QItemModelBarDataProxy__MultiMatchBehavior {

	return QItemModelBarDataProxy__MultiMatchBehavior(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiMatchBehavior"}).(float64))
}

func (ptr *QItemModelBarDataProxy) ConnectMultiMatchBehaviorChanged(f func(behavior QItemModelBarDataProxy__MultiMatchBehavior)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMultiMatchBehaviorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectMultiMatchBehaviorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMultiMatchBehaviorChanged"})
}

func (ptr *QItemModelBarDataProxy) MultiMatchBehaviorChanged(behavior QItemModelBarDataProxy__MultiMatchBehavior) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiMatchBehaviorChanged", behavior})
}

func (ptr *QItemModelBarDataProxy) Remap(rowRole string, columnRole string, valueRole string, rotationRole string, rowCategories []string, columnCategories []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remap", rowRole, columnRole, valueRole, rotationRole, rowCategories, columnCategories})
}

func (ptr *QItemModelBarDataProxy) RotationRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRole"}).(string)
}

func (ptr *QItemModelBarDataProxy) ConnectRotationRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRotationRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectRotationRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRotationRoleChanged"})
}

func (ptr *QItemModelBarDataProxy) RotationRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRoleChanged", role})
}

func (ptr *QItemModelBarDataProxy) RotationRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelBarDataProxy) ConnectRotationRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRotationRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectRotationRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRotationRolePatternChanged"})
}

func (ptr *QItemModelBarDataProxy) RotationRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRolePatternChanged", pattern})
}

func (ptr *QItemModelBarDataProxy) RotationRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRoleReplace"}).(string)
}

func (ptr *QItemModelBarDataProxy) ConnectRotationRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRotationRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectRotationRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRotationRoleReplaceChanged"})
}

func (ptr *QItemModelBarDataProxy) RotationRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRoleReplaceChanged", replace})
}

func (ptr *QItemModelBarDataProxy) RowCategories() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCategories"}).([]string)
}

func (ptr *QItemModelBarDataProxy) ConnectRowCategoriesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectRowCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowCategoriesChanged"})
}

func (ptr *QItemModelBarDataProxy) RowCategoriesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCategoriesChanged"})
}

func (ptr *QItemModelBarDataProxy) RowCategoryIndex(category string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCategoryIndex", category}).(float64))
}

func (ptr *QItemModelBarDataProxy) RowRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRole"}).(string)
}

func (ptr *QItemModelBarDataProxy) ConnectRowRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectRowRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowRoleChanged"})
}

func (ptr *QItemModelBarDataProxy) RowRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRoleChanged", role})
}

func (ptr *QItemModelBarDataProxy) RowRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelBarDataProxy) ConnectRowRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectRowRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowRolePatternChanged"})
}

func (ptr *QItemModelBarDataProxy) RowRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRolePatternChanged", pattern})
}

func (ptr *QItemModelBarDataProxy) RowRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRoleReplace"}).(string)
}

func (ptr *QItemModelBarDataProxy) ConnectRowRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectRowRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowRoleReplaceChanged"})
}

func (ptr *QItemModelBarDataProxy) RowRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRoleReplaceChanged", replace})
}

func (ptr *QItemModelBarDataProxy) SetAutoColumnCategories(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoColumnCategories", enable})
}

func (ptr *QItemModelBarDataProxy) SetAutoRowCategories(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoRowCategories", enable})
}

func (ptr *QItemModelBarDataProxy) SetColumnCategories(categories []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnCategories", categories})
}

func (ptr *QItemModelBarDataProxy) SetColumnRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnRole", role})
}

func (ptr *QItemModelBarDataProxy) SetColumnRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnRolePattern", pattern})
}

func (ptr *QItemModelBarDataProxy) SetColumnRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnRoleReplace", replace})
}

func (ptr *QItemModelBarDataProxy) SetItemModel(itemModel core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemModel", itemModel})
}

func (ptr *QItemModelBarDataProxy) SetMultiMatchBehavior(behavior QItemModelBarDataProxy__MultiMatchBehavior) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMultiMatchBehavior", behavior})
}

func (ptr *QItemModelBarDataProxy) SetRotationRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotationRole", role})
}

func (ptr *QItemModelBarDataProxy) SetRotationRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotationRolePattern", pattern})
}

func (ptr *QItemModelBarDataProxy) SetRotationRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotationRoleReplace", replace})
}

func (ptr *QItemModelBarDataProxy) SetRowCategories(categories []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowCategories", categories})
}

func (ptr *QItemModelBarDataProxy) SetRowRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowRole", role})
}

func (ptr *QItemModelBarDataProxy) SetRowRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowRolePattern", pattern})
}

func (ptr *QItemModelBarDataProxy) SetRowRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowRoleReplace", replace})
}

func (ptr *QItemModelBarDataProxy) SetUseModelCategories(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUseModelCategories", enable})
}

func (ptr *QItemModelBarDataProxy) SetValueRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValueRole", role})
}

func (ptr *QItemModelBarDataProxy) SetValueRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValueRolePattern", pattern})
}

func (ptr *QItemModelBarDataProxy) SetValueRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValueRoleReplace", replace})
}

func (ptr *QItemModelBarDataProxy) UseModelCategories() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UseModelCategories"}).(bool)
}

func (ptr *QItemModelBarDataProxy) ConnectUseModelCategoriesChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUseModelCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectUseModelCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUseModelCategoriesChanged"})
}

func (ptr *QItemModelBarDataProxy) UseModelCategoriesChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UseModelCategoriesChanged", enable})
}

func (ptr *QItemModelBarDataProxy) ValueRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueRole"}).(string)
}

func (ptr *QItemModelBarDataProxy) ConnectValueRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValueRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectValueRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValueRoleChanged"})
}

func (ptr *QItemModelBarDataProxy) ValueRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueRoleChanged", role})
}

func (ptr *QItemModelBarDataProxy) ValueRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelBarDataProxy) ConnectValueRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValueRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectValueRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValueRolePatternChanged"})
}

func (ptr *QItemModelBarDataProxy) ValueRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueRolePatternChanged", pattern})
}

func (ptr *QItemModelBarDataProxy) ValueRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueRoleReplace"}).(string)
}

func (ptr *QItemModelBarDataProxy) ConnectValueRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValueRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectValueRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValueRoleReplaceChanged"})
}

func (ptr *QItemModelBarDataProxy) ValueRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueRoleReplaceChanged", replace})
}

func (ptr *QItemModelBarDataProxy) ConnectDestroyQItemModelBarDataProxy(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQItemModelBarDataProxy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelBarDataProxy) DisconnectDestroyQItemModelBarDataProxy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQItemModelBarDataProxy"})
}

func (ptr *QItemModelBarDataProxy) DestroyQItemModelBarDataProxy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQItemModelBarDataProxy"})
}

func (ptr *QItemModelBarDataProxy) DestroyQItemModelBarDataProxyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQItemModelBarDataProxyDefault"})
}

type QItemModelScatterDataProxy struct {
	internal.Internal
}

type QItemModelScatterDataProxy_ITF interface {
	QItemModelScatterDataProxy_PTR() *QItemModelScatterDataProxy
}

func (ptr *QItemModelScatterDataProxy) QItemModelScatterDataProxy_PTR() *QItemModelScatterDataProxy {
	return ptr
}

func (ptr *QItemModelScatterDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QItemModelScatterDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQItemModelScatterDataProxy(ptr QItemModelScatterDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemModelScatterDataProxy_PTR().Pointer()
	}
	return nil
}

func (n *QItemModelScatterDataProxy) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQItemModelScatterDataProxyFromPointer(ptr unsafe.Pointer) (n *QItemModelScatterDataProxy) {
	n = new(QItemModelScatterDataProxy)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QItemModelScatterDataProxy")
	return
}
func NewQItemModelScatterDataProxy(parent core.QObject_ITF) *QItemModelScatterDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelScatterDataProxy", "", parent}).(*QItemModelScatterDataProxy)
}

func NewQItemModelScatterDataProxy2(itemModel core.QAbstractItemModel_ITF, parent core.QObject_ITF) *QItemModelScatterDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelScatterDataProxy2", "", itemModel, parent}).(*QItemModelScatterDataProxy)
}

func NewQItemModelScatterDataProxy3(itemModel core.QAbstractItemModel_ITF, xPosRole string, yPosRole string, zPosRole string, parent core.QObject_ITF) *QItemModelScatterDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelScatterDataProxy3", "", itemModel, xPosRole, yPosRole, zPosRole, parent}).(*QItemModelScatterDataProxy)
}

func NewQItemModelScatterDataProxy4(itemModel core.QAbstractItemModel_ITF, xPosRole string, yPosRole string, zPosRole string, rotationRole string, parent core.QObject_ITF) *QItemModelScatterDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelScatterDataProxy4", "", itemModel, xPosRole, yPosRole, zPosRole, rotationRole, parent}).(*QItemModelScatterDataProxy)
}

func (ptr *QItemModelScatterDataProxy) ItemModel() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemModel"}).(*core.QAbstractItemModel)
}

func (ptr *QItemModelScatterDataProxy) ConnectItemModelChanged(f func(itemModel *core.QAbstractItemModel)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemModelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectItemModelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemModelChanged"})
}

func (ptr *QItemModelScatterDataProxy) ItemModelChanged(itemModel core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemModelChanged", itemModel})
}

func (ptr *QItemModelScatterDataProxy) Remap(xPosRole string, yPosRole string, zPosRole string, rotationRole string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remap", xPosRole, yPosRole, zPosRole, rotationRole})
}

func (ptr *QItemModelScatterDataProxy) RotationRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRole"}).(string)
}

func (ptr *QItemModelScatterDataProxy) ConnectRotationRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRotationRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectRotationRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRotationRoleChanged"})
}

func (ptr *QItemModelScatterDataProxy) RotationRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRoleChanged", role})
}

func (ptr *QItemModelScatterDataProxy) RotationRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelScatterDataProxy) ConnectRotationRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRotationRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectRotationRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRotationRolePatternChanged"})
}

func (ptr *QItemModelScatterDataProxy) RotationRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRolePatternChanged", pattern})
}

func (ptr *QItemModelScatterDataProxy) RotationRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRoleReplace"}).(string)
}

func (ptr *QItemModelScatterDataProxy) ConnectRotationRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRotationRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectRotationRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRotationRoleReplaceChanged"})
}

func (ptr *QItemModelScatterDataProxy) RotationRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RotationRoleReplaceChanged", replace})
}

func (ptr *QItemModelScatterDataProxy) SetItemModel(itemModel core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemModel", itemModel})
}

func (ptr *QItemModelScatterDataProxy) SetRotationRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotationRole", role})
}

func (ptr *QItemModelScatterDataProxy) SetRotationRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotationRolePattern", pattern})
}

func (ptr *QItemModelScatterDataProxy) SetRotationRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotationRoleReplace", replace})
}

func (ptr *QItemModelScatterDataProxy) SetXPosRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXPosRole", role})
}

func (ptr *QItemModelScatterDataProxy) SetXPosRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXPosRolePattern", pattern})
}

func (ptr *QItemModelScatterDataProxy) SetXPosRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXPosRoleReplace", replace})
}

func (ptr *QItemModelScatterDataProxy) SetYPosRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYPosRole", role})
}

func (ptr *QItemModelScatterDataProxy) SetYPosRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYPosRolePattern", pattern})
}

func (ptr *QItemModelScatterDataProxy) SetYPosRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYPosRoleReplace", replace})
}

func (ptr *QItemModelScatterDataProxy) SetZPosRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZPosRole", role})
}

func (ptr *QItemModelScatterDataProxy) SetZPosRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZPosRolePattern", pattern})
}

func (ptr *QItemModelScatterDataProxy) SetZPosRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZPosRoleReplace", replace})
}

func (ptr *QItemModelScatterDataProxy) XPosRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRole"}).(string)
}

func (ptr *QItemModelScatterDataProxy) ConnectXPosRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectXPosRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectXPosRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectXPosRoleChanged"})
}

func (ptr *QItemModelScatterDataProxy) XPosRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRoleChanged", role})
}

func (ptr *QItemModelScatterDataProxy) XPosRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelScatterDataProxy) ConnectXPosRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectXPosRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectXPosRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectXPosRolePatternChanged"})
}

func (ptr *QItemModelScatterDataProxy) XPosRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRolePatternChanged", pattern})
}

func (ptr *QItemModelScatterDataProxy) XPosRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRoleReplace"}).(string)
}

func (ptr *QItemModelScatterDataProxy) ConnectXPosRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectXPosRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectXPosRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectXPosRoleReplaceChanged"})
}

func (ptr *QItemModelScatterDataProxy) XPosRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRoleReplaceChanged", replace})
}

func (ptr *QItemModelScatterDataProxy) YPosRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRole"}).(string)
}

func (ptr *QItemModelScatterDataProxy) ConnectYPosRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectYPosRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectYPosRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectYPosRoleChanged"})
}

func (ptr *QItemModelScatterDataProxy) YPosRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRoleChanged", role})
}

func (ptr *QItemModelScatterDataProxy) YPosRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelScatterDataProxy) ConnectYPosRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectYPosRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectYPosRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectYPosRolePatternChanged"})
}

func (ptr *QItemModelScatterDataProxy) YPosRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRolePatternChanged", pattern})
}

func (ptr *QItemModelScatterDataProxy) YPosRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRoleReplace"}).(string)
}

func (ptr *QItemModelScatterDataProxy) ConnectYPosRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectYPosRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectYPosRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectYPosRoleReplaceChanged"})
}

func (ptr *QItemModelScatterDataProxy) YPosRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRoleReplaceChanged", replace})
}

func (ptr *QItemModelScatterDataProxy) ZPosRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRole"}).(string)
}

func (ptr *QItemModelScatterDataProxy) ConnectZPosRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZPosRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectZPosRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZPosRoleChanged"})
}

func (ptr *QItemModelScatterDataProxy) ZPosRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRoleChanged", role})
}

func (ptr *QItemModelScatterDataProxy) ZPosRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelScatterDataProxy) ConnectZPosRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZPosRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectZPosRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZPosRolePatternChanged"})
}

func (ptr *QItemModelScatterDataProxy) ZPosRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRolePatternChanged", pattern})
}

func (ptr *QItemModelScatterDataProxy) ZPosRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRoleReplace"}).(string)
}

func (ptr *QItemModelScatterDataProxy) ConnectZPosRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZPosRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectZPosRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZPosRoleReplaceChanged"})
}

func (ptr *QItemModelScatterDataProxy) ZPosRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRoleReplaceChanged", replace})
}

func (ptr *QItemModelScatterDataProxy) ConnectDestroyQItemModelScatterDataProxy(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQItemModelScatterDataProxy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelScatterDataProxy) DisconnectDestroyQItemModelScatterDataProxy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQItemModelScatterDataProxy"})
}

func (ptr *QItemModelScatterDataProxy) DestroyQItemModelScatterDataProxy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQItemModelScatterDataProxy"})
}

func (ptr *QItemModelScatterDataProxy) DestroyQItemModelScatterDataProxyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQItemModelScatterDataProxyDefault"})
}

type QItemModelSurfaceDataProxy struct {
	internal.Internal
}

type QItemModelSurfaceDataProxy_ITF interface {
	QItemModelSurfaceDataProxy_PTR() *QItemModelSurfaceDataProxy
}

func (ptr *QItemModelSurfaceDataProxy) QItemModelSurfaceDataProxy_PTR() *QItemModelSurfaceDataProxy {
	return ptr
}

func (ptr *QItemModelSurfaceDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QItemModelSurfaceDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQItemModelSurfaceDataProxy(ptr QItemModelSurfaceDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemModelSurfaceDataProxy_PTR().Pointer()
	}
	return nil
}

func (n *QItemModelSurfaceDataProxy) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQItemModelSurfaceDataProxyFromPointer(ptr unsafe.Pointer) (n *QItemModelSurfaceDataProxy) {
	n = new(QItemModelSurfaceDataProxy)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QItemModelSurfaceDataProxy")
	return
}

//go:generate stringer -type=QItemModelSurfaceDataProxy__MultiMatchBehavior
//QItemModelSurfaceDataProxy::MultiMatchBehavior
type QItemModelSurfaceDataProxy__MultiMatchBehavior int64

const (
	QItemModelSurfaceDataProxy__MMBFirst       QItemModelSurfaceDataProxy__MultiMatchBehavior = QItemModelSurfaceDataProxy__MultiMatchBehavior(0)
	QItemModelSurfaceDataProxy__MMBLast        QItemModelSurfaceDataProxy__MultiMatchBehavior = QItemModelSurfaceDataProxy__MultiMatchBehavior(1)
	QItemModelSurfaceDataProxy__MMBAverage     QItemModelSurfaceDataProxy__MultiMatchBehavior = QItemModelSurfaceDataProxy__MultiMatchBehavior(2)
	QItemModelSurfaceDataProxy__MMBCumulativeY QItemModelSurfaceDataProxy__MultiMatchBehavior = QItemModelSurfaceDataProxy__MultiMatchBehavior(3)
)

func NewQItemModelSurfaceDataProxy(parent core.QObject_ITF) *QItemModelSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelSurfaceDataProxy", "", parent}).(*QItemModelSurfaceDataProxy)
}

func NewQItemModelSurfaceDataProxy2(itemModel core.QAbstractItemModel_ITF, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelSurfaceDataProxy2", "", itemModel, parent}).(*QItemModelSurfaceDataProxy)
}

func NewQItemModelSurfaceDataProxy3(itemModel core.QAbstractItemModel_ITF, yPosRole string, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelSurfaceDataProxy3", "", itemModel, yPosRole, parent}).(*QItemModelSurfaceDataProxy)
}

func NewQItemModelSurfaceDataProxy4(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, yPosRole string, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelSurfaceDataProxy4", "", itemModel, rowRole, columnRole, yPosRole, parent}).(*QItemModelSurfaceDataProxy)
}

func NewQItemModelSurfaceDataProxy5(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, xPosRole string, yPosRole string, zPosRole string, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelSurfaceDataProxy5", "", itemModel, rowRole, columnRole, xPosRole, yPosRole, zPosRole, parent}).(*QItemModelSurfaceDataProxy)
}

func NewQItemModelSurfaceDataProxy6(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, yPosRole string, rowCategories []string, columnCategories []string, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelSurfaceDataProxy6", "", itemModel, rowRole, columnRole, yPosRole, rowCategories, columnCategories, parent}).(*QItemModelSurfaceDataProxy)
}

func NewQItemModelSurfaceDataProxy7(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, xPosRole string, yPosRole string, zPosRole string, rowCategories []string, columnCategories []string, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQItemModelSurfaceDataProxy7", "", itemModel, rowRole, columnRole, xPosRole, yPosRole, zPosRole, rowCategories, columnCategories, parent}).(*QItemModelSurfaceDataProxy)
}

func (ptr *QItemModelSurfaceDataProxy) AutoColumnCategories() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoColumnCategories"}).(bool)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectAutoColumnCategoriesChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAutoColumnCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectAutoColumnCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAutoColumnCategoriesChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) AutoColumnCategoriesChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoColumnCategoriesChanged", enable})
}

func (ptr *QItemModelSurfaceDataProxy) AutoRowCategories() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoRowCategories"}).(bool)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectAutoRowCategoriesChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAutoRowCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectAutoRowCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAutoRowCategoriesChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) AutoRowCategoriesChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoRowCategoriesChanged", enable})
}

func (ptr *QItemModelSurfaceDataProxy) ColumnCategories() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCategories"}).([]string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectColumnCategoriesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectColumnCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnCategoriesChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) ColumnCategoriesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCategoriesChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) ColumnCategoryIndex(category string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCategoryIndex", category}).(float64))
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRole"}).(string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectColumnRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectColumnRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnRoleChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRoleChanged", role})
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectColumnRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectColumnRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnRolePatternChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRolePatternChanged", pattern})
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRoleReplace"}).(string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectColumnRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectColumnRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnRoleReplaceChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnRoleReplaceChanged", replace})
}

func (ptr *QItemModelSurfaceDataProxy) ItemModel() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemModel"}).(*core.QAbstractItemModel)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectItemModelChanged(f func(itemModel *core.QAbstractItemModel)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemModelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectItemModelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemModelChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) ItemModelChanged(itemModel core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemModelChanged", itemModel})
}

func (ptr *QItemModelSurfaceDataProxy) MultiMatchBehavior() QItemModelSurfaceDataProxy__MultiMatchBehavior {

	return QItemModelSurfaceDataProxy__MultiMatchBehavior(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiMatchBehavior"}).(float64))
}

func (ptr *QItemModelSurfaceDataProxy) ConnectMultiMatchBehaviorChanged(f func(behavior QItemModelSurfaceDataProxy__MultiMatchBehavior)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMultiMatchBehaviorChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectMultiMatchBehaviorChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMultiMatchBehaviorChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) MultiMatchBehaviorChanged(behavior QItemModelSurfaceDataProxy__MultiMatchBehavior) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MultiMatchBehaviorChanged", behavior})
}

func (ptr *QItemModelSurfaceDataProxy) Remap(rowRole string, columnRole string, xPosRole string, yPosRole string, zPosRole string, rowCategories []string, columnCategories []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remap", rowRole, columnRole, xPosRole, yPosRole, zPosRole, rowCategories, columnCategories})
}

func (ptr *QItemModelSurfaceDataProxy) RowCategories() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCategories"}).([]string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectRowCategoriesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectRowCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowCategoriesChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) RowCategoriesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCategoriesChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) RowCategoryIndex(category string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCategoryIndex", category}).(float64))
}

func (ptr *QItemModelSurfaceDataProxy) RowRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRole"}).(string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectRowRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectRowRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowRoleChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) RowRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRoleChanged", role})
}

func (ptr *QItemModelSurfaceDataProxy) RowRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectRowRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectRowRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowRolePatternChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) RowRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRolePatternChanged", pattern})
}

func (ptr *QItemModelSurfaceDataProxy) RowRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRoleReplace"}).(string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectRowRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectRowRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowRoleReplaceChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) RowRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowRoleReplaceChanged", replace})
}

func (ptr *QItemModelSurfaceDataProxy) SetAutoColumnCategories(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoColumnCategories", enable})
}

func (ptr *QItemModelSurfaceDataProxy) SetAutoRowCategories(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoRowCategories", enable})
}

func (ptr *QItemModelSurfaceDataProxy) SetColumnCategories(categories []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnCategories", categories})
}

func (ptr *QItemModelSurfaceDataProxy) SetColumnRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnRole", role})
}

func (ptr *QItemModelSurfaceDataProxy) SetColumnRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnRolePattern", pattern})
}

func (ptr *QItemModelSurfaceDataProxy) SetColumnRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumnRoleReplace", replace})
}

func (ptr *QItemModelSurfaceDataProxy) SetItemModel(itemModel core.QAbstractItemModel_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemModel", itemModel})
}

func (ptr *QItemModelSurfaceDataProxy) SetMultiMatchBehavior(behavior QItemModelSurfaceDataProxy__MultiMatchBehavior) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMultiMatchBehavior", behavior})
}

func (ptr *QItemModelSurfaceDataProxy) SetRowCategories(categories []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowCategories", categories})
}

func (ptr *QItemModelSurfaceDataProxy) SetRowRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowRole", role})
}

func (ptr *QItemModelSurfaceDataProxy) SetRowRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowRolePattern", pattern})
}

func (ptr *QItemModelSurfaceDataProxy) SetRowRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRowRoleReplace", replace})
}

func (ptr *QItemModelSurfaceDataProxy) SetUseModelCategories(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUseModelCategories", enable})
}

func (ptr *QItemModelSurfaceDataProxy) SetXPosRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXPosRole", role})
}

func (ptr *QItemModelSurfaceDataProxy) SetXPosRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXPosRolePattern", pattern})
}

func (ptr *QItemModelSurfaceDataProxy) SetXPosRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXPosRoleReplace", replace})
}

func (ptr *QItemModelSurfaceDataProxy) SetYPosRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYPosRole", role})
}

func (ptr *QItemModelSurfaceDataProxy) SetYPosRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYPosRolePattern", pattern})
}

func (ptr *QItemModelSurfaceDataProxy) SetYPosRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYPosRoleReplace", replace})
}

func (ptr *QItemModelSurfaceDataProxy) SetZPosRole(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZPosRole", role})
}

func (ptr *QItemModelSurfaceDataProxy) SetZPosRolePattern(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZPosRolePattern", pattern})
}

func (ptr *QItemModelSurfaceDataProxy) SetZPosRoleReplace(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZPosRoleReplace", replace})
}

func (ptr *QItemModelSurfaceDataProxy) UseModelCategories() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UseModelCategories"}).(bool)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectUseModelCategoriesChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUseModelCategoriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectUseModelCategoriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUseModelCategoriesChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) UseModelCategoriesChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UseModelCategoriesChanged", enable})
}

func (ptr *QItemModelSurfaceDataProxy) XPosRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRole"}).(string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectXPosRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectXPosRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectXPosRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectXPosRoleChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) XPosRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRoleChanged", role})
}

func (ptr *QItemModelSurfaceDataProxy) XPosRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectXPosRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectXPosRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectXPosRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectXPosRolePatternChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) XPosRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRolePatternChanged", pattern})
}

func (ptr *QItemModelSurfaceDataProxy) XPosRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRoleReplace"}).(string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectXPosRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectXPosRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectXPosRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectXPosRoleReplaceChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) XPosRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XPosRoleReplaceChanged", replace})
}

func (ptr *QItemModelSurfaceDataProxy) YPosRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRole"}).(string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectYPosRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectYPosRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectYPosRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectYPosRoleChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) YPosRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRoleChanged", role})
}

func (ptr *QItemModelSurfaceDataProxy) YPosRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectYPosRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectYPosRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectYPosRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectYPosRolePatternChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) YPosRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRolePatternChanged", pattern})
}

func (ptr *QItemModelSurfaceDataProxy) YPosRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRoleReplace"}).(string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectYPosRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectYPosRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectYPosRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectYPosRoleReplaceChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) YPosRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YPosRoleReplaceChanged", replace})
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRole() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRole"}).(string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectZPosRoleChanged(f func(role string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZPosRoleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectZPosRoleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZPosRoleChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRoleChanged(role string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRoleChanged", role})
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRolePattern() *core.QRegExp {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRolePattern"}).(*core.QRegExp)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectZPosRolePatternChanged(f func(pattern *core.QRegExp)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZPosRolePatternChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectZPosRolePatternChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZPosRolePatternChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRolePatternChanged(pattern core.QRegExp_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRolePatternChanged", pattern})
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRoleReplace() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRoleReplace"}).(string)
}

func (ptr *QItemModelSurfaceDataProxy) ConnectZPosRoleReplaceChanged(f func(replace string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZPosRoleReplaceChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectZPosRoleReplaceChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZPosRoleReplaceChanged"})
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRoleReplaceChanged(replace string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZPosRoleReplaceChanged", replace})
}

func (ptr *QItemModelSurfaceDataProxy) ConnectDestroyQItemModelSurfaceDataProxy(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQItemModelSurfaceDataProxy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectDestroyQItemModelSurfaceDataProxy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQItemModelSurfaceDataProxy"})
}

func (ptr *QItemModelSurfaceDataProxy) DestroyQItemModelSurfaceDataProxy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQItemModelSurfaceDataProxy"})
}

func (ptr *QItemModelSurfaceDataProxy) DestroyQItemModelSurfaceDataProxyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQItemModelSurfaceDataProxyDefault"})
}

type QLogValue3DAxisFormatter struct {
	internal.Internal
}

type QLogValue3DAxisFormatter_ITF interface {
	QLogValue3DAxisFormatter_PTR() *QLogValue3DAxisFormatter
}

func (ptr *QLogValue3DAxisFormatter) QLogValue3DAxisFormatter_PTR() *QLogValue3DAxisFormatter {
	return ptr
}

func (ptr *QLogValue3DAxisFormatter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLogValue3DAxisFormatter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLogValue3DAxisFormatter(ptr QLogValue3DAxisFormatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLogValue3DAxisFormatter_PTR().Pointer()
	}
	return nil
}

func (n *QLogValue3DAxisFormatter) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLogValue3DAxisFormatterFromPointer(ptr unsafe.Pointer) (n *QLogValue3DAxisFormatter) {
	n = new(QLogValue3DAxisFormatter)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QLogValue3DAxisFormatter")
	return
}
func NewQLogValue3DAxisFormatter2(parent core.QObject_ITF) *QLogValue3DAxisFormatter {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQLogValue3DAxisFormatter2", "", parent}).(*QLogValue3DAxisFormatter)
}

func (ptr *QLogValue3DAxisFormatter) AutoSubGrid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoSubGrid"}).(bool)
}

func (ptr *QLogValue3DAxisFormatter) ConnectAutoSubGridChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAutoSubGridChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValue3DAxisFormatter) DisconnectAutoSubGridChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAutoSubGridChanged"})
}

func (ptr *QLogValue3DAxisFormatter) AutoSubGridChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoSubGridChanged", enabled})
}

func (ptr *QLogValue3DAxisFormatter) Base() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Base"}).(float64)
}

func (ptr *QLogValue3DAxisFormatter) ConnectBaseChanged(f func(base float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBaseChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValue3DAxisFormatter) DisconnectBaseChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBaseChanged"})
}

func (ptr *QLogValue3DAxisFormatter) BaseChanged(base float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseChanged", base})
}

func (ptr *QLogValue3DAxisFormatter) SetAutoSubGrid(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoSubGrid", enabled})
}

func (ptr *QLogValue3DAxisFormatter) SetBase(base float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBase", base})
}

func (ptr *QLogValue3DAxisFormatter) SetShowEdgeLabels(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetShowEdgeLabels", enabled})
}

func (ptr *QLogValue3DAxisFormatter) ShowEdgeLabels() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEdgeLabels"}).(bool)
}

func (ptr *QLogValue3DAxisFormatter) ConnectShowEdgeLabelsChanged(f func(enabled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShowEdgeLabelsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValue3DAxisFormatter) DisconnectShowEdgeLabelsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShowEdgeLabelsChanged"})
}

func (ptr *QLogValue3DAxisFormatter) ShowEdgeLabelsChanged(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEdgeLabelsChanged", enabled})
}

func (ptr *QLogValue3DAxisFormatter) ConnectDestroyQLogValue3DAxisFormatter(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLogValue3DAxisFormatter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLogValue3DAxisFormatter) DisconnectDestroyQLogValue3DAxisFormatter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLogValue3DAxisFormatter"})
}

func (ptr *QLogValue3DAxisFormatter) DestroyQLogValue3DAxisFormatter() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLogValue3DAxisFormatter"})
}

func (ptr *QLogValue3DAxisFormatter) DestroyQLogValue3DAxisFormatterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLogValue3DAxisFormatterDefault"})
}

type QScatter3DSeries struct {
	internal.Internal
}

type QScatter3DSeries_ITF interface {
	QScatter3DSeries_PTR() *QScatter3DSeries
}

func (ptr *QScatter3DSeries) QScatter3DSeries_PTR() *QScatter3DSeries {
	return ptr
}

func (ptr *QScatter3DSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScatter3DSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScatter3DSeries(ptr QScatter3DSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScatter3DSeries_PTR().Pointer()
	}
	return nil
}

func (n *QScatter3DSeries) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScatter3DSeriesFromPointer(ptr unsafe.Pointer) (n *QScatter3DSeries) {
	n = new(QScatter3DSeries)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QScatter3DSeries")
	return
}
func NewQScatter3DSeries(parent core.QObject_ITF) *QScatter3DSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQScatter3DSeries", "", parent}).(*QScatter3DSeries)
}

func NewQScatter3DSeries2(dataProxy QScatterDataProxy_ITF, parent core.QObject_ITF) *QScatter3DSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQScatter3DSeries2", "", dataProxy, parent}).(*QScatter3DSeries)
}

func (ptr *QScatter3DSeries) DataProxy() *QScatterDataProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataProxy"}).(*QScatterDataProxy)
}

func (ptr *QScatter3DSeries) ConnectDataProxyChanged(f func(proxy *QScatterDataProxy)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDataProxyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatter3DSeries) DisconnectDataProxyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDataProxyChanged"})
}

func (ptr *QScatter3DSeries) DataProxyChanged(proxy QScatterDataProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataProxyChanged", proxy})
}

func QScatter3DSeries_InvalidSelectionIndex() int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "datavisualization.QScatter3DSeries_InvalidSelectionIndex", ""}).(float64))
}

func (ptr *QScatter3DSeries) InvalidSelectionIndex() int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "datavisualization.QScatter3DSeries_InvalidSelectionIndex", ""}).(float64))
}

func (ptr *QScatter3DSeries) ItemSize() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemSize"}).(float32)
}

func (ptr *QScatter3DSeries) ConnectItemSizeChanged(f func(size float32)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemSizeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatter3DSeries) DisconnectItemSizeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemSizeChanged"})
}

func (ptr *QScatter3DSeries) ItemSizeChanged(size float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemSizeChanged", size})
}

func (ptr *QScatter3DSeries) SelectedItem() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedItem"}).(float64))
}

func (ptr *QScatter3DSeries) ConnectSelectedItemChanged(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectedItemChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatter3DSeries) DisconnectSelectedItemChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectedItemChanged"})
}

func (ptr *QScatter3DSeries) SelectedItemChanged(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedItemChanged", index})
}

func (ptr *QScatter3DSeries) SetDataProxy(proxy QScatterDataProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataProxy", proxy})
}

func (ptr *QScatter3DSeries) SetItemSize(size float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemSize", size})
}

func (ptr *QScatter3DSeries) SetSelectedItem(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectedItem", index})
}

func (ptr *QScatter3DSeries) ConnectDestroyQScatter3DSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQScatter3DSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatter3DSeries) DisconnectDestroyQScatter3DSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQScatter3DSeries"})
}

func (ptr *QScatter3DSeries) DestroyQScatter3DSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScatter3DSeries"})
}

func (ptr *QScatter3DSeries) DestroyQScatter3DSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScatter3DSeriesDefault"})
}

type QScatterDataItem struct {
	internal.Internal
}

type QScatterDataItem_ITF interface {
	QScatterDataItem_PTR() *QScatterDataItem
}

func (ptr *QScatterDataItem) QScatterDataItem_PTR() *QScatterDataItem {
	return ptr
}

func (ptr *QScatterDataItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScatterDataItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScatterDataItem(ptr QScatterDataItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScatterDataItem_PTR().Pointer()
	}
	return nil
}

func (n *QScatterDataItem) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScatterDataItemFromPointer(ptr unsafe.Pointer) (n *QScatterDataItem) {
	n = new(QScatterDataItem)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QScatterDataItem")
	return
}
func NewQScatterDataItem() *QScatterDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQScatterDataItem", ""}).(*QScatterDataItem)
}

func NewQScatterDataItem2(position gui.QVector3D_ITF) *QScatterDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQScatterDataItem2", "", position}).(*QScatterDataItem)
}

func NewQScatterDataItem3(position gui.QVector3D_ITF, rotation gui.QQuaternion_ITF) *QScatterDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQScatterDataItem3", "", position, rotation}).(*QScatterDataItem)
}

func NewQScatterDataItem4(other QScatterDataItem_ITF) *QScatterDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQScatterDataItem4", "", other}).(*QScatterDataItem)
}

func (ptr *QScatterDataItem) Position() *gui.QVector3D {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Position"}).(*gui.QVector3D)
}

func (ptr *QScatterDataItem) Rotation() *gui.QQuaternion {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rotation"}).(*gui.QQuaternion)
}

func (ptr *QScatterDataItem) SetPosition(pos gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPosition", pos})
}

func (ptr *QScatterDataItem) SetRotation(rot gui.QQuaternion_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRotation", rot})
}

func (ptr *QScatterDataItem) SetX(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetX", value})
}

func (ptr *QScatterDataItem) SetY(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetY", value})
}

func (ptr *QScatterDataItem) SetZ(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZ", value})
}

func (ptr *QScatterDataItem) X() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "X"}).(float32)
}

func (ptr *QScatterDataItem) Y() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Y"}).(float32)
}

func (ptr *QScatterDataItem) Z() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Z"}).(float32)
}

func (ptr *QScatterDataItem) DestroyQScatterDataItem() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScatterDataItem"})
}

type QScatterDataProxy struct {
	internal.Internal
}

type QScatterDataProxy_ITF interface {
	QScatterDataProxy_PTR() *QScatterDataProxy
}

func (ptr *QScatterDataProxy) QScatterDataProxy_PTR() *QScatterDataProxy {
	return ptr
}

func (ptr *QScatterDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QScatterDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQScatterDataProxy(ptr QScatterDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScatterDataProxy_PTR().Pointer()
	}
	return nil
}

func (n *QScatterDataProxy) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQScatterDataProxyFromPointer(ptr unsafe.Pointer) (n *QScatterDataProxy) {
	n = new(QScatterDataProxy)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QScatterDataProxy")
	return
}
func NewQScatterDataProxy(parent core.QObject_ITF) *QScatterDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQScatterDataProxy", "", parent}).(*QScatterDataProxy)
}

func (ptr *QScatterDataProxy) AddItem(item QScatterDataItem_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddItem", item}).(float64))
}

func (ptr *QScatterDataProxy) ConnectArrayReset(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectArrayReset", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterDataProxy) DisconnectArrayReset() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectArrayReset"})
}

func (ptr *QScatterDataProxy) ArrayReset() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ArrayReset"})
}

func (ptr *QScatterDataProxy) InsertItem(index int, item QScatterDataItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertItem", index, item})
}

func (ptr *QScatterDataProxy) ItemAt(index int) *QScatterDataItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemAt", index}).(*QScatterDataItem)
}

func (ptr *QScatterDataProxy) ItemCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemCount"}).(float64))
}

func (ptr *QScatterDataProxy) ConnectItemCountChanged(f func(count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterDataProxy) DisconnectItemCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemCountChanged"})
}

func (ptr *QScatterDataProxy) ItemCountChanged(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemCountChanged", count})
}

func (ptr *QScatterDataProxy) ConnectItemsAdded(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemsAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterDataProxy) DisconnectItemsAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemsAdded"})
}

func (ptr *QScatterDataProxy) ItemsAdded(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemsAdded", startIndex, count})
}

func (ptr *QScatterDataProxy) ConnectItemsChanged(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterDataProxy) DisconnectItemsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemsChanged"})
}

func (ptr *QScatterDataProxy) ItemsChanged(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemsChanged", startIndex, count})
}

func (ptr *QScatterDataProxy) ConnectItemsInserted(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemsInserted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterDataProxy) DisconnectItemsInserted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemsInserted"})
}

func (ptr *QScatterDataProxy) ItemsInserted(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemsInserted", startIndex, count})
}

func (ptr *QScatterDataProxy) ConnectItemsRemoved(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemsRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterDataProxy) DisconnectItemsRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemsRemoved"})
}

func (ptr *QScatterDataProxy) ItemsRemoved(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemsRemoved", startIndex, count})
}

func (ptr *QScatterDataProxy) RemoveItems(index int, removeCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveItems", index, removeCount})
}

func (ptr *QScatterDataProxy) Series() *QScatter3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QScatter3DSeries)
}

func (ptr *QScatterDataProxy) ConnectSeriesChanged(f func(series *QScatter3DSeries)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterDataProxy) DisconnectSeriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesChanged"})
}

func (ptr *QScatterDataProxy) SeriesChanged(series QScatter3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesChanged", series})
}

func (ptr *QScatterDataProxy) SetItem(index int, item QScatterDataItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItem", index, item})
}

func (ptr *QScatterDataProxy) ConnectDestroyQScatterDataProxy(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQScatterDataProxy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScatterDataProxy) DisconnectDestroyQScatterDataProxy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQScatterDataProxy"})
}

func (ptr *QScatterDataProxy) DestroyQScatterDataProxy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScatterDataProxy"})
}

func (ptr *QScatterDataProxy) DestroyQScatterDataProxyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScatterDataProxyDefault"})
}

type QSurface3DSeries struct {
	internal.Internal
}

type QSurface3DSeries_ITF interface {
	QSurface3DSeries_PTR() *QSurface3DSeries
}

func (ptr *QSurface3DSeries) QSurface3DSeries_PTR() *QSurface3DSeries {
	return ptr
}

func (ptr *QSurface3DSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSurface3DSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSurface3DSeries(ptr QSurface3DSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurface3DSeries_PTR().Pointer()
	}
	return nil
}

func (n *QSurface3DSeries) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSurface3DSeriesFromPointer(ptr unsafe.Pointer) (n *QSurface3DSeries) {
	n = new(QSurface3DSeries)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QSurface3DSeries")
	return
}

//go:generate stringer -type=QSurface3DSeries__DrawFlag
//QSurface3DSeries::DrawFlag
type QSurface3DSeries__DrawFlag int64

const (
	QSurface3DSeries__DrawWireframe           QSurface3DSeries__DrawFlag = QSurface3DSeries__DrawFlag(1)
	QSurface3DSeries__DrawSurface             QSurface3DSeries__DrawFlag = QSurface3DSeries__DrawFlag(2)
	QSurface3DSeries__DrawSurfaceAndWireframe QSurface3DSeries__DrawFlag = QSurface3DSeries__DrawFlag(QSurface3DSeries__DrawWireframe | QSurface3DSeries__DrawSurface)
)

func NewQSurface3DSeries(parent core.QObject_ITF) *QSurface3DSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQSurface3DSeries", "", parent}).(*QSurface3DSeries)
}

func NewQSurface3DSeries2(dataProxy QSurfaceDataProxy_ITF, parent core.QObject_ITF) *QSurface3DSeries {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQSurface3DSeries2", "", dataProxy, parent}).(*QSurface3DSeries)
}

func (ptr *QSurface3DSeries) DataProxy() *QSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataProxy"}).(*QSurfaceDataProxy)
}

func (ptr *QSurface3DSeries) ConnectDataProxyChanged(f func(proxy *QSurfaceDataProxy)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDataProxyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurface3DSeries) DisconnectDataProxyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDataProxyChanged"})
}

func (ptr *QSurface3DSeries) DataProxyChanged(proxy QSurfaceDataProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataProxyChanged", proxy})
}

func (ptr *QSurface3DSeries) DrawMode() QSurface3DSeries__DrawFlag {

	return QSurface3DSeries__DrawFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawMode"}).(float64))
}

func (ptr *QSurface3DSeries) ConnectDrawModeChanged(f func(mode QSurface3DSeries__DrawFlag)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDrawModeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurface3DSeries) DisconnectDrawModeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDrawModeChanged"})
}

func (ptr *QSurface3DSeries) DrawModeChanged(mode QSurface3DSeries__DrawFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawModeChanged", mode})
}

func (ptr *QSurface3DSeries) ConnectFlatShadingEnabledChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFlatShadingEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurface3DSeries) DisconnectFlatShadingEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFlatShadingEnabledChanged"})
}

func (ptr *QSurface3DSeries) FlatShadingEnabledChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlatShadingEnabledChanged", enable})
}

func (ptr *QSurface3DSeries) ConnectFlatShadingSupportedChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFlatShadingSupportedChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurface3DSeries) DisconnectFlatShadingSupportedChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFlatShadingSupportedChanged"})
}

func (ptr *QSurface3DSeries) FlatShadingSupportedChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlatShadingSupportedChanged", enable})
}

func QSurface3DSeries_InvalidSelectionPosition() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.QSurface3DSeries_InvalidSelectionPosition", ""}).(*core.QPoint)
}

func (ptr *QSurface3DSeries) InvalidSelectionPosition() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.QSurface3DSeries_InvalidSelectionPosition", ""}).(*core.QPoint)
}

func (ptr *QSurface3DSeries) IsFlatShadingEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFlatShadingEnabled"}).(bool)
}

func (ptr *QSurface3DSeries) IsFlatShadingSupported() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFlatShadingSupported"}).(bool)
}

func (ptr *QSurface3DSeries) SelectedPoint() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedPoint"}).(*core.QPoint)
}

func (ptr *QSurface3DSeries) ConnectSelectedPointChanged(f func(position *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectedPointChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurface3DSeries) DisconnectSelectedPointChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectedPointChanged"})
}

func (ptr *QSurface3DSeries) SelectedPointChanged(position core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedPointChanged", position})
}

func (ptr *QSurface3DSeries) SetDataProxy(proxy QSurfaceDataProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataProxy", proxy})
}

func (ptr *QSurface3DSeries) SetDrawMode(mode QSurface3DSeries__DrawFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDrawMode", mode})
}

func (ptr *QSurface3DSeries) SetFlatShadingEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFlatShadingEnabled", enabled})
}

func (ptr *QSurface3DSeries) SetSelectedPoint(position core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectedPoint", position})
}

func (ptr *QSurface3DSeries) SetTexture(texture gui.QImage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTexture", texture})
}

func (ptr *QSurface3DSeries) SetTextureFile(filename string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextureFile", filename})
}

func (ptr *QSurface3DSeries) Texture() *gui.QImage {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Texture"}).(*gui.QImage)
}

func (ptr *QSurface3DSeries) ConnectTextureChanged(f func(image *gui.QImage)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurface3DSeries) DisconnectTextureChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureChanged"})
}

func (ptr *QSurface3DSeries) TextureChanged(image gui.QImage_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureChanged", image})
}

func (ptr *QSurface3DSeries) TextureFile() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureFile"}).(string)
}

func (ptr *QSurface3DSeries) ConnectTextureFileChanged(f func(filename string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextureFileChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurface3DSeries) DisconnectTextureFileChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextureFileChanged"})
}

func (ptr *QSurface3DSeries) TextureFileChanged(filename string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextureFileChanged", filename})
}

func (ptr *QSurface3DSeries) ConnectDestroyQSurface3DSeries(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSurface3DSeries", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurface3DSeries) DisconnectDestroyQSurface3DSeries() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSurface3DSeries"})
}

func (ptr *QSurface3DSeries) DestroyQSurface3DSeries() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSurface3DSeries"})
}

func (ptr *QSurface3DSeries) DestroyQSurface3DSeriesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSurface3DSeriesDefault"})
}

type QSurfaceDataItem struct {
	internal.Internal
}

type QSurfaceDataItem_ITF interface {
	QSurfaceDataItem_PTR() *QSurfaceDataItem
}

func (ptr *QSurfaceDataItem) QSurfaceDataItem_PTR() *QSurfaceDataItem {
	return ptr
}

func (ptr *QSurfaceDataItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSurfaceDataItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSurfaceDataItem(ptr QSurfaceDataItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurfaceDataItem_PTR().Pointer()
	}
	return nil
}

func (n *QSurfaceDataItem) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSurfaceDataItemFromPointer(ptr unsafe.Pointer) (n *QSurfaceDataItem) {
	n = new(QSurfaceDataItem)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QSurfaceDataItem")
	return
}
func NewQSurfaceDataItem() *QSurfaceDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQSurfaceDataItem", ""}).(*QSurfaceDataItem)
}

func NewQSurfaceDataItem2(position gui.QVector3D_ITF) *QSurfaceDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQSurfaceDataItem2", "", position}).(*QSurfaceDataItem)
}

func NewQSurfaceDataItem3(other QSurfaceDataItem_ITF) *QSurfaceDataItem {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQSurfaceDataItem3", "", other}).(*QSurfaceDataItem)
}

func (ptr *QSurfaceDataItem) Position() *gui.QVector3D {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Position"}).(*gui.QVector3D)
}

func (ptr *QSurfaceDataItem) SetPosition(pos gui.QVector3D_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPosition", pos})
}

func (ptr *QSurfaceDataItem) SetX(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetX", value})
}

func (ptr *QSurfaceDataItem) SetY(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetY", value})
}

func (ptr *QSurfaceDataItem) SetZ(value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZ", value})
}

func (ptr *QSurfaceDataItem) X() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "X"}).(float32)
}

func (ptr *QSurfaceDataItem) Y() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Y"}).(float32)
}

func (ptr *QSurfaceDataItem) Z() float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Z"}).(float32)
}

func (ptr *QSurfaceDataItem) DestroyQSurfaceDataItem() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSurfaceDataItem"})
}

type QSurfaceDataProxy struct {
	internal.Internal
}

type QSurfaceDataProxy_ITF interface {
	QSurfaceDataProxy_PTR() *QSurfaceDataProxy
}

func (ptr *QSurfaceDataProxy) QSurfaceDataProxy_PTR() *QSurfaceDataProxy {
	return ptr
}

func (ptr *QSurfaceDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSurfaceDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSurfaceDataProxy(ptr QSurfaceDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurfaceDataProxy_PTR().Pointer()
	}
	return nil
}

func (n *QSurfaceDataProxy) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSurfaceDataProxyFromPointer(ptr unsafe.Pointer) (n *QSurfaceDataProxy) {
	n = new(QSurfaceDataProxy)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QSurfaceDataProxy")
	return
}
func NewQSurfaceDataProxy(parent core.QObject_ITF) *QSurfaceDataProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQSurfaceDataProxy", "", parent}).(*QSurfaceDataProxy)
}

func (ptr *QSurfaceDataProxy) ConnectArrayReset(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectArrayReset", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurfaceDataProxy) DisconnectArrayReset() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectArrayReset"})
}

func (ptr *QSurfaceDataProxy) ArrayReset() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ArrayReset"})
}

func (ptr *QSurfaceDataProxy) ColumnCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCount"}).(float64))
}

func (ptr *QSurfaceDataProxy) ConnectColumnCountChanged(f func(count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurfaceDataProxy) DisconnectColumnCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnCountChanged"})
}

func (ptr *QSurfaceDataProxy) ColumnCountChanged(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountChanged", count})
}

func (ptr *QSurfaceDataProxy) ItemAt(rowIndex int, columnIndex int) *QSurfaceDataItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemAt", rowIndex, columnIndex}).(*QSurfaceDataItem)
}

func (ptr *QSurfaceDataProxy) ItemAt2(position core.QPoint_ITF) *QSurfaceDataItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemAt2", position}).(*QSurfaceDataItem)
}

func (ptr *QSurfaceDataProxy) ConnectItemChanged(f func(rowIndex int, columnIndex int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurfaceDataProxy) DisconnectItemChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemChanged"})
}

func (ptr *QSurfaceDataProxy) ItemChanged(rowIndex int, columnIndex int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemChanged", rowIndex, columnIndex})
}

func (ptr *QSurfaceDataProxy) RemoveRows(rowIndex int, removeCount int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveRows", rowIndex, removeCount})
}

func (ptr *QSurfaceDataProxy) RowCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCount"}).(float64))
}

func (ptr *QSurfaceDataProxy) ConnectRowCountChanged(f func(count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurfaceDataProxy) DisconnectRowCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowCountChanged"})
}

func (ptr *QSurfaceDataProxy) RowCountChanged(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountChanged", count})
}

func (ptr *QSurfaceDataProxy) ConnectRowsAdded(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowsAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurfaceDataProxy) DisconnectRowsAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowsAdded"})
}

func (ptr *QSurfaceDataProxy) RowsAdded(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsAdded", startIndex, count})
}

func (ptr *QSurfaceDataProxy) ConnectRowsChanged(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurfaceDataProxy) DisconnectRowsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowsChanged"})
}

func (ptr *QSurfaceDataProxy) RowsChanged(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsChanged", startIndex, count})
}

func (ptr *QSurfaceDataProxy) ConnectRowsInserted(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowsInserted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurfaceDataProxy) DisconnectRowsInserted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowsInserted"})
}

func (ptr *QSurfaceDataProxy) RowsInserted(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsInserted", startIndex, count})
}

func (ptr *QSurfaceDataProxy) ConnectRowsRemoved(f func(startIndex int, count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowsRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurfaceDataProxy) DisconnectRowsRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowsRemoved"})
}

func (ptr *QSurfaceDataProxy) RowsRemoved(startIndex int, count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowsRemoved", startIndex, count})
}

func (ptr *QSurfaceDataProxy) Series() *QSurface3DSeries {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Series"}).(*QSurface3DSeries)
}

func (ptr *QSurfaceDataProxy) ConnectSeriesChanged(f func(series *QSurface3DSeries)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeriesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurfaceDataProxy) DisconnectSeriesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeriesChanged"})
}

func (ptr *QSurfaceDataProxy) SeriesChanged(series QSurface3DSeries_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeriesChanged", series})
}

func (ptr *QSurfaceDataProxy) SetItem(rowIndex int, columnIndex int, item QSurfaceDataItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItem", rowIndex, columnIndex, item})
}

func (ptr *QSurfaceDataProxy) SetItem2(position core.QPoint_ITF, item QSurfaceDataItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItem2", position, item})
}

func (ptr *QSurfaceDataProxy) ConnectDestroyQSurfaceDataProxy(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSurfaceDataProxy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSurfaceDataProxy) DisconnectDestroyQSurfaceDataProxy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSurfaceDataProxy"})
}

func (ptr *QSurfaceDataProxy) DestroyQSurfaceDataProxy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSurfaceDataProxy"})
}

func (ptr *QSurfaceDataProxy) DestroyQSurfaceDataProxyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSurfaceDataProxyDefault"})
}

type QTouch3DInputHandler struct {
	internal.Internal
}

type QTouch3DInputHandler_ITF interface {
	QTouch3DInputHandler_PTR() *QTouch3DInputHandler
}

func (ptr *QTouch3DInputHandler) QTouch3DInputHandler_PTR() *QTouch3DInputHandler {
	return ptr
}

func (ptr *QTouch3DInputHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QTouch3DInputHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQTouch3DInputHandler(ptr QTouch3DInputHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTouch3DInputHandler_PTR().Pointer()
	}
	return nil
}

func (n *QTouch3DInputHandler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQTouch3DInputHandlerFromPointer(ptr unsafe.Pointer) (n *QTouch3DInputHandler) {
	n = new(QTouch3DInputHandler)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QTouch3DInputHandler")
	return
}
func NewQTouch3DInputHandler(parent core.QObject_ITF) *QTouch3DInputHandler {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQTouch3DInputHandler", "", parent}).(*QTouch3DInputHandler)
}

func (ptr *QTouch3DInputHandler) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTouchEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTouch3DInputHandler) DisconnectTouchEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTouchEvent"})
}

func (ptr *QTouch3DInputHandler) TouchEvent(event gui.QTouchEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TouchEvent", event})
}

func (ptr *QTouch3DInputHandler) TouchEventDefault(event gui.QTouchEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TouchEventDefault", event})
}

func (ptr *QTouch3DInputHandler) ConnectDestroyQTouch3DInputHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQTouch3DInputHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTouch3DInputHandler) DisconnectDestroyQTouch3DInputHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQTouch3DInputHandler"})
}

func (ptr *QTouch3DInputHandler) DestroyQTouch3DInputHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTouch3DInputHandler"})
}

func (ptr *QTouch3DInputHandler) DestroyQTouch3DInputHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTouch3DInputHandlerDefault"})
}

type QValue3DAxis struct {
	internal.Internal
}

type QValue3DAxis_ITF interface {
	QValue3DAxis_PTR() *QValue3DAxis
}

func (ptr *QValue3DAxis) QValue3DAxis_PTR() *QValue3DAxis {
	return ptr
}

func (ptr *QValue3DAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QValue3DAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQValue3DAxis(ptr QValue3DAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QValue3DAxis_PTR().Pointer()
	}
	return nil
}

func (n *QValue3DAxis) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQValue3DAxisFromPointer(ptr unsafe.Pointer) (n *QValue3DAxis) {
	n = new(QValue3DAxis)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QValue3DAxis")
	return
}
func NewQValue3DAxis(parent core.QObject_ITF) *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQValue3DAxis", "", parent}).(*QValue3DAxis)
}

func (ptr *QValue3DAxis) Formatter() *QValue3DAxisFormatter {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Formatter"}).(*QValue3DAxisFormatter)
}

func (ptr *QValue3DAxis) ConnectFormatterChanged(f func(formatter *QValue3DAxisFormatter)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFormatterChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxis) DisconnectFormatterChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFormatterChanged"})
}

func (ptr *QValue3DAxis) FormatterChanged(formatter QValue3DAxisFormatter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormatterChanged", formatter})
}

func (ptr *QValue3DAxis) LabelFormat() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelFormat"}).(string)
}

func (ptr *QValue3DAxis) ConnectLabelFormatChanged(f func(format string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLabelFormatChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxis) DisconnectLabelFormatChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLabelFormatChanged"})
}

func (ptr *QValue3DAxis) LabelFormatChanged(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelFormatChanged", format})
}

func (ptr *QValue3DAxis) Reversed() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reversed"}).(bool)
}

func (ptr *QValue3DAxis) ConnectReversedChanged(f func(enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReversedChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxis) DisconnectReversedChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReversedChanged"})
}

func (ptr *QValue3DAxis) ReversedChanged(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReversedChanged", enable})
}

func (ptr *QValue3DAxis) SegmentCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SegmentCount"}).(float64))
}

func (ptr *QValue3DAxis) ConnectSegmentCountChanged(f func(count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSegmentCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxis) DisconnectSegmentCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSegmentCountChanged"})
}

func (ptr *QValue3DAxis) SegmentCountChanged(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SegmentCountChanged", count})
}

func (ptr *QValue3DAxis) SetFormatter(formatter QValue3DAxisFormatter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFormatter", formatter})
}

func (ptr *QValue3DAxis) SetLabelFormat(format string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLabelFormat", format})
}

func (ptr *QValue3DAxis) SetReversed(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReversed", enable})
}

func (ptr *QValue3DAxis) SetSegmentCount(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSegmentCount", count})
}

func (ptr *QValue3DAxis) SetSubSegmentCount(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSubSegmentCount", count})
}

func (ptr *QValue3DAxis) SubSegmentCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubSegmentCount"}).(float64))
}

func (ptr *QValue3DAxis) ConnectSubSegmentCountChanged(f func(count int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSubSegmentCountChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxis) DisconnectSubSegmentCountChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSubSegmentCountChanged"})
}

func (ptr *QValue3DAxis) SubSegmentCountChanged(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubSegmentCountChanged", count})
}

func (ptr *QValue3DAxis) ConnectDestroyQValue3DAxis(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQValue3DAxis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxis) DisconnectDestroyQValue3DAxis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQValue3DAxis"})
}

func (ptr *QValue3DAxis) DestroyQValue3DAxis() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQValue3DAxis"})
}

func (ptr *QValue3DAxis) DestroyQValue3DAxisDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQValue3DAxisDefault"})
}

type QValue3DAxisFormatter struct {
	internal.Internal
}

type QValue3DAxisFormatter_ITF interface {
	QValue3DAxisFormatter_PTR() *QValue3DAxisFormatter
}

func (ptr *QValue3DAxisFormatter) QValue3DAxisFormatter_PTR() *QValue3DAxisFormatter {
	return ptr
}

func (ptr *QValue3DAxisFormatter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QValue3DAxisFormatter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQValue3DAxisFormatter(ptr QValue3DAxisFormatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QValue3DAxisFormatter_PTR().Pointer()
	}
	return nil
}

func (n *QValue3DAxisFormatter) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQValue3DAxisFormatterFromPointer(ptr unsafe.Pointer) (n *QValue3DAxisFormatter) {
	n = new(QValue3DAxisFormatter)
	n.InitFromInternal(uintptr(ptr), "datavisualization.QValue3DAxisFormatter")
	return
}
func NewQValue3DAxisFormatter2(parent core.QObject_ITF) *QValue3DAxisFormatter {

	return internal.CallLocalFunction([]interface{}{"", "", "datavisualization.NewQValue3DAxisFormatter2", "", parent}).(*QValue3DAxisFormatter)
}

func (ptr *QValue3DAxisFormatter) AllowNegatives() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AllowNegatives"}).(bool)
}

func (ptr *QValue3DAxisFormatter) AllowZero() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AllowZero"}).(bool)
}

func (ptr *QValue3DAxisFormatter) Axis() *QValue3DAxis {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Axis"}).(*QValue3DAxis)
}

func (ptr *QValue3DAxisFormatter) ConnectCreateNewInstance(f func() *QValue3DAxisFormatter) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateNewInstance", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxisFormatter) DisconnectCreateNewInstance() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateNewInstance"})
}

func (ptr *QValue3DAxisFormatter) CreateNewInstance() *QValue3DAxisFormatter {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateNewInstance"}).(*QValue3DAxisFormatter)
}

func (ptr *QValue3DAxisFormatter) CreateNewInstanceDefault() *QValue3DAxisFormatter {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateNewInstanceDefault"}).(*QValue3DAxisFormatter)
}

func (ptr *QValue3DAxisFormatter) GridPositions() []float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GridPositions"}).([]float32)
}

func (ptr *QValue3DAxisFormatter) LabelPositions() []float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelPositions"}).([]float32)
}

func (ptr *QValue3DAxisFormatter) LabelStrings() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LabelStrings"}).([]string)
}

func (ptr *QValue3DAxisFormatter) Locale() *core.QLocale {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Locale"}).(*core.QLocale)
}

func (ptr *QValue3DAxisFormatter) MarkDirty(labelsChange bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MarkDirty", labelsChange})
}

func (ptr *QValue3DAxisFormatter) ConnectPopulateCopy(f func(copy *QValue3DAxisFormatter)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPopulateCopy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxisFormatter) DisconnectPopulateCopy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPopulateCopy"})
}

func (ptr *QValue3DAxisFormatter) PopulateCopy(copy QValue3DAxisFormatter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PopulateCopy", copy})
}

func (ptr *QValue3DAxisFormatter) PopulateCopyDefault(copy QValue3DAxisFormatter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PopulateCopyDefault", copy})
}

func (ptr *QValue3DAxisFormatter) ConnectPositionAt(f func(value float32) float32) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPositionAt", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxisFormatter) DisconnectPositionAt() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPositionAt"})
}

func (ptr *QValue3DAxisFormatter) PositionAt(value float32) float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionAt", value}).(float32)
}

func (ptr *QValue3DAxisFormatter) PositionAtDefault(value float32) float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionAtDefault", value}).(float32)
}

func (ptr *QValue3DAxisFormatter) ConnectRecalculate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRecalculate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxisFormatter) DisconnectRecalculate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRecalculate"})
}

func (ptr *QValue3DAxisFormatter) Recalculate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Recalculate"})
}

func (ptr *QValue3DAxisFormatter) RecalculateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RecalculateDefault"})
}

func (ptr *QValue3DAxisFormatter) SetAllowNegatives(allow bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAllowNegatives", allow})
}

func (ptr *QValue3DAxisFormatter) SetAllowZero(allow bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAllowZero", allow})
}

func (ptr *QValue3DAxisFormatter) SetLocale(locale core.QLocale_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocale", locale})
}

func (ptr *QValue3DAxisFormatter) ConnectStringForValue(f func(value float64, format string) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStringForValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxisFormatter) DisconnectStringForValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStringForValue"})
}

func (ptr *QValue3DAxisFormatter) StringForValue(value float64, format string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StringForValue", value, format}).(string)
}

func (ptr *QValue3DAxisFormatter) StringForValueDefault(value float64, format string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StringForValueDefault", value, format}).(string)
}

func (ptr *QValue3DAxisFormatter) SubGridPositions() []float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubGridPositions"}).([]float32)
}

func (ptr *QValue3DAxisFormatter) ConnectValueAt(f func(position float32) float32) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValueAt", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxisFormatter) DisconnectValueAt() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValueAt"})
}

func (ptr *QValue3DAxisFormatter) ValueAt(position float32) float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueAt", position}).(float32)
}

func (ptr *QValue3DAxisFormatter) ValueAtDefault(position float32) float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueAtDefault", position}).(float32)
}

func (ptr *QValue3DAxisFormatter) ConnectDestroyQValue3DAxisFormatter(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQValue3DAxisFormatter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QValue3DAxisFormatter) DisconnectDestroyQValue3DAxisFormatter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQValue3DAxisFormatter"})
}

func (ptr *QValue3DAxisFormatter) DestroyQValue3DAxisFormatter() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQValue3DAxisFormatter"})
}

func (ptr *QValue3DAxisFormatter) DestroyQValue3DAxisFormatterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQValue3DAxisFormatterDefault"})
}

func (ptr *QValue3DAxisFormatter) __gridPositions_atList(i int) float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__gridPositions_atList", i}).(float32)
}

func (ptr *QValue3DAxisFormatter) __gridPositions_setList(i float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__gridPositions_setList", i})
}

func (ptr *QValue3DAxisFormatter) __gridPositions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__gridPositions_newList"}).(unsafe.Pointer)
}

func (ptr *QValue3DAxisFormatter) __labelPositions_atList(i int) float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__labelPositions_atList", i}).(float32)
}

func (ptr *QValue3DAxisFormatter) __labelPositions_setList(i float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__labelPositions_setList", i})
}

func (ptr *QValue3DAxisFormatter) __labelPositions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__labelPositions_newList"}).(unsafe.Pointer)
}

func (ptr *QValue3DAxisFormatter) __subGridPositions_atList(i int) float32 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__subGridPositions_atList", i}).(float32)
}

func (ptr *QValue3DAxisFormatter) __subGridPositions_setList(i float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__subGridPositions_setList", i})
}

func (ptr *QValue3DAxisFormatter) __subGridPositions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__subGridPositions_newList"}).(unsafe.Pointer)
}

func init() {
	internal.ConstructorTable["datavisualization.Q3DBars"] = NewQ3DBarsFromPointer
	internal.ConstructorTable["datavisualization.Q3DCamera"] = NewQ3DCameraFromPointer
	internal.ConstructorTable["datavisualization.Q3DInputHandler"] = NewQ3DInputHandlerFromPointer
	internal.ConstructorTable["datavisualization.Q3DLight"] = NewQ3DLightFromPointer
	internal.ConstructorTable["datavisualization.Q3DObject"] = NewQ3DObjectFromPointer
	internal.ConstructorTable["datavisualization.Q3DScatter"] = NewQ3DScatterFromPointer
	internal.ConstructorTable["datavisualization.Q3DScene"] = NewQ3DSceneFromPointer
	internal.ConstructorTable["datavisualization.Q3DSurface"] = NewQ3DSurfaceFromPointer
	internal.ConstructorTable["datavisualization.Q3DTheme"] = NewQ3DThemeFromPointer
	internal.ConstructorTable["datavisualization.QAbstract3DAxis"] = NewQAbstract3DAxisFromPointer
	internal.ConstructorTable["datavisualization.QAbstract3DInputHandler"] = NewQAbstract3DInputHandlerFromPointer
	internal.ConstructorTable["datavisualization.QAbstract3DSeries"] = NewQAbstract3DSeriesFromPointer
	internal.ConstructorTable["datavisualization.QAbstractDataProxy"] = NewQAbstractDataProxyFromPointer
	internal.ConstructorTable["datavisualization.QBar3DSeries"] = NewQBar3DSeriesFromPointer
	internal.ConstructorTable["datavisualization.QBarDataItem"] = NewQBarDataItemFromPointer
	internal.ConstructorTable["datavisualization.QBarDataProxy"] = NewQBarDataProxyFromPointer
	internal.ConstructorTable["datavisualization.QCategory3DAxis"] = NewQCategory3DAxisFromPointer
	internal.ConstructorTable["datavisualization.QCustom3DItem"] = NewQCustom3DItemFromPointer
	internal.ConstructorTable["datavisualization.QCustom3DLabel"] = NewQCustom3DLabelFromPointer
	internal.ConstructorTable["datavisualization.QCustom3DVolume"] = NewQCustom3DVolumeFromPointer
	internal.ConstructorTable["datavisualization.QHeightMapSurfaceDataProxy"] = NewQHeightMapSurfaceDataProxyFromPointer
	internal.ConstructorTable["datavisualization.QItemModelBarDataProxy"] = NewQItemModelBarDataProxyFromPointer
	internal.ConstructorTable["datavisualization.QItemModelScatterDataProxy"] = NewQItemModelScatterDataProxyFromPointer
	internal.ConstructorTable["datavisualization.QItemModelSurfaceDataProxy"] = NewQItemModelSurfaceDataProxyFromPointer
	internal.ConstructorTable["datavisualization.QLogValue3DAxisFormatter"] = NewQLogValue3DAxisFormatterFromPointer
	internal.ConstructorTable["datavisualization.QScatter3DSeries"] = NewQScatter3DSeriesFromPointer
	internal.ConstructorTable["datavisualization.QScatterDataItem"] = NewQScatterDataItemFromPointer
	internal.ConstructorTable["datavisualization.QScatterDataProxy"] = NewQScatterDataProxyFromPointer
	internal.ConstructorTable["datavisualization.QSurface3DSeries"] = NewQSurface3DSeriesFromPointer
	internal.ConstructorTable["datavisualization.QSurfaceDataItem"] = NewQSurfaceDataItemFromPointer
	internal.ConstructorTable["datavisualization.QSurfaceDataProxy"] = NewQSurfaceDataProxyFromPointer
	internal.ConstructorTable["datavisualization.QTouch3DInputHandler"] = NewQTouch3DInputHandlerFromPointer
	internal.ConstructorTable["datavisualization.QValue3DAxis"] = NewQValue3DAxisFromPointer
	internal.ConstructorTable["datavisualization.QValue3DAxisFormatter"] = NewQValue3DAxisFormatterFromPointer
}
