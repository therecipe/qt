//go:build !minimal
// +build !minimal

package virtualkeyboard

import (
	"unsafe"

	"github.com/akiyosi/qt/core"
	"github.com/akiyosi/qt/internal"
)

type QVirtualKeyboardAbstractInputMethod struct {
	core.QObject
}

type QVirtualKeyboardAbstractInputMethod_ITF interface {
	core.QObject_ITF
	QVirtualKeyboardAbstractInputMethod_PTR() *QVirtualKeyboardAbstractInputMethod
}

func (ptr *QVirtualKeyboardAbstractInputMethod) QVirtualKeyboardAbstractInputMethod_PTR() *QVirtualKeyboardAbstractInputMethod {
	return ptr
}

func (ptr *QVirtualKeyboardAbstractInputMethod) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQVirtualKeyboardAbstractInputMethod(ptr QVirtualKeyboardAbstractInputMethod_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVirtualKeyboardAbstractInputMethod_PTR().Pointer()
	}
	return nil
}

func (n *QVirtualKeyboardAbstractInputMethod) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVirtualKeyboardAbstractInputMethod) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardAbstractInputMethod) {
	n = new(QVirtualKeyboardAbstractInputMethod)
	n.InitFromInternal(uintptr(ptr), "virtualkeyboard.QVirtualKeyboardAbstractInputMethod")
	return
}
func NewQVirtualKeyboardAbstractInputMethod(parent core.QObject_ITF) *QVirtualKeyboardAbstractInputMethod {

	return internal.CallLocalFunction([]interface{}{"", "", "virtualkeyboard.NewQVirtualKeyboardAbstractInputMethod", "", parent}).(*QVirtualKeyboardAbstractInputMethod)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectClickPreeditText(f func(cursorPosition int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClickPreeditText", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectClickPreeditText() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClickPreeditText"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ClickPreeditText(cursorPosition int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClickPreeditText", cursorPosition}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ClickPreeditTextDefault(cursorPosition int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClickPreeditTextDefault", cursorPosition}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) InputContext() *QVirtualKeyboardInputContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputContext"}).(*QVirtualKeyboardInputContext)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) InputEngine() *QVirtualKeyboardInputEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputEngine"}).(*QVirtualKeyboardInputEngine)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectInputModes(f func(locale string) []QVirtualKeyboardInputEngine__InputMode) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInputModes", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectInputModes() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInputModes"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) InputModes(locale string) []QVirtualKeyboardInputEngine__InputMode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputModes", locale}).([]QVirtualKeyboardInputEngine__InputMode)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectKeyEvent(f func(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectKeyEvent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectKeyEvent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectKeyEvent"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) KeyEvent(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyEvent", key, text, modifiers}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectPatternRecognitionModes(f func() []QVirtualKeyboardInputEngine__PatternRecognitionMode) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPatternRecognitionModes", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectPatternRecognitionModes() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPatternRecognitionModes"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) PatternRecognitionModes() []QVirtualKeyboardInputEngine__PatternRecognitionMode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PatternRecognitionModes"}).([]QVirtualKeyboardInputEngine__PatternRecognitionMode)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) PatternRecognitionModesDefault() []QVirtualKeyboardInputEngine__PatternRecognitionMode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PatternRecognitionModesDefault"}).([]QVirtualKeyboardInputEngine__PatternRecognitionMode)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectReselect(f func(cursorPosition int, reselectFlags QVirtualKeyboardInputEngine__ReselectFlag) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReselect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectReselect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReselect"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) Reselect(cursorPosition int, reselectFlags QVirtualKeyboardInputEngine__ReselectFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reselect", cursorPosition, reselectFlags}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ReselectDefault(cursorPosition int, reselectFlags QVirtualKeyboardInputEngine__ReselectFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReselectDefault", cursorPosition, reselectFlags}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectReset(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReset", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectReset() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReset"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) Reset() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reset"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ResetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetDefault"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListActiveItemChanged(f func(ty QVirtualKeyboardSelectionListModel__Type, index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionListActiveItemChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListActiveItemChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionListActiveItemChanged"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListActiveItemChanged(ty QVirtualKeyboardSelectionListModel__Type, index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListActiveItemChanged", ty, index})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListChanged(f func(ty QVirtualKeyboardSelectionListModel__Type)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionListChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionListChanged"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListChanged(ty QVirtualKeyboardSelectionListModel__Type) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListChanged", ty})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListData(f func(ty QVirtualKeyboardSelectionListModel__Type, index int, role QVirtualKeyboardSelectionListModel__Role) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionListData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionListData"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListData(ty QVirtualKeyboardSelectionListModel__Type, index int, role QVirtualKeyboardSelectionListModel__Role) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListData", ty, index, role}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListDataDefault(ty QVirtualKeyboardSelectionListModel__Type, index int, role QVirtualKeyboardSelectionListModel__Role) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListDataDefault", ty, index, role}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListItemCount(f func(ty QVirtualKeyboardSelectionListModel__Type) int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionListItemCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListItemCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionListItemCount"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListItemCount(ty QVirtualKeyboardSelectionListModel__Type) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListItemCount", ty}).(float64))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListItemCountDefault(ty QVirtualKeyboardSelectionListModel__Type) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListItemCountDefault", ty}).(float64))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListItemSelected(f func(ty QVirtualKeyboardSelectionListModel__Type, index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionListItemSelected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListItemSelected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionListItemSelected"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListItemSelected(ty QVirtualKeyboardSelectionListModel__Type, index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListItemSelected", ty, index})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListItemSelectedDefault(ty QVirtualKeyboardSelectionListModel__Type, index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListItemSelectedDefault", ty, index})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListRemoveItem(f func(ty QVirtualKeyboardSelectionListModel__Type, index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionListRemoveItem", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListRemoveItem() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionListRemoveItem"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListRemoveItem(ty QVirtualKeyboardSelectionListModel__Type, index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListRemoveItem", ty, index}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListRemoveItemDefault(ty QVirtualKeyboardSelectionListModel__Type, index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListRemoveItemDefault", ty, index}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionLists(f func() []QVirtualKeyboardSelectionListModel__Type) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionLists", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionLists() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionLists"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionLists() []QVirtualKeyboardSelectionListModel__Type {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionLists"}).([]QVirtualKeyboardSelectionListModel__Type)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListsDefault() []QVirtualKeyboardSelectionListModel__Type {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListsDefault"}).([]QVirtualKeyboardSelectionListModel__Type)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionListsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionListsChanged"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionListsChanged"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSetInputMode(f func(locale string, inputMode QVirtualKeyboardInputEngine__InputMode) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetInputMode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSetInputMode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetInputMode"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SetInputMode(locale string, inputMode QVirtualKeyboardInputEngine__InputMode) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInputMode", locale, inputMode}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSetTextCase(f func(textCase QVirtualKeyboardInputEngine__TextCase) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetTextCase", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSetTextCase() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetTextCase"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SetTextCase(textCase QVirtualKeyboardInputEngine__TextCase) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTextCase", textCase}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectTraceBegin(f func(traceId int, patternRecognitionMode QVirtualKeyboardInputEngine__PatternRecognitionMode, traceCaptureDeviceInfo map[string]*core.QVariant, traceScreenInfo map[string]*core.QVariant) *QVirtualKeyboardTrace) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTraceBegin", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectTraceBegin() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTraceBegin"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) TraceBegin(traceId int, patternRecognitionMode QVirtualKeyboardInputEngine__PatternRecognitionMode, traceCaptureDeviceInfo map[string]*core.QVariant, traceScreenInfo map[string]*core.QVariant) *QVirtualKeyboardTrace {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TraceBegin", traceId, patternRecognitionMode, traceCaptureDeviceInfo, traceScreenInfo}).(*QVirtualKeyboardTrace)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) TraceBeginDefault(traceId int, patternRecognitionMode QVirtualKeyboardInputEngine__PatternRecognitionMode, traceCaptureDeviceInfo map[string]*core.QVariant, traceScreenInfo map[string]*core.QVariant) *QVirtualKeyboardTrace {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TraceBeginDefault", traceId, patternRecognitionMode, traceCaptureDeviceInfo, traceScreenInfo}).(*QVirtualKeyboardTrace)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectTraceEnd(f func(trace *QVirtualKeyboardTrace) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTraceEnd", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectTraceEnd() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTraceEnd"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) TraceEnd(trace QVirtualKeyboardTrace_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TraceEnd", trace}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) TraceEndDefault(trace QVirtualKeyboardTrace_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TraceEndDefault", trace}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectUpdate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectUpdate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdate"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) Update() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Update"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectDestroyQVirtualKeyboardAbstractInputMethod(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQVirtualKeyboardAbstractInputMethod", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectDestroyQVirtualKeyboardAbstractInputMethod() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQVirtualKeyboardAbstractInputMethod"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DestroyQVirtualKeyboardAbstractInputMethod() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQVirtualKeyboardAbstractInputMethod"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DestroyQVirtualKeyboardAbstractInputMethodDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQVirtualKeyboardAbstractInputMethodDefault"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __inputModes_atList(i int) QVirtualKeyboardInputEngine__InputMode {

	return QVirtualKeyboardInputEngine__InputMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__inputModes_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __inputModes_setList(i QVirtualKeyboardInputEngine__InputMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__inputModes_setList", i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __inputModes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__inputModes_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __patternRecognitionModes_atList(i int) QVirtualKeyboardInputEngine__PatternRecognitionMode {

	return QVirtualKeyboardInputEngine__PatternRecognitionMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__patternRecognitionModes_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __patternRecognitionModes_setList(i QVirtualKeyboardInputEngine__PatternRecognitionMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__patternRecognitionModes_setList", i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __patternRecognitionModes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__patternRecognitionModes_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __selectionLists_atList(i int) QVirtualKeyboardSelectionListModel__Type {

	return QVirtualKeyboardSelectionListModel__Type(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__selectionLists_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __selectionLists_setList(i QVirtualKeyboardSelectionListModel__Type) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__selectionLists_setList", i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __selectionLists_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__selectionLists_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceCaptureDeviceInfo_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceCaptureDeviceInfo_atList", v, i}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceCaptureDeviceInfo_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceCaptureDeviceInfo_setList", key, i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceCaptureDeviceInfo_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceCaptureDeviceInfo_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceCaptureDeviceInfo_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceCaptureDeviceInfo_keyList"}).([]string)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceScreenInfo_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceScreenInfo_atList", v, i}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceScreenInfo_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceScreenInfo_setList", key, i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceScreenInfo_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceScreenInfo_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceScreenInfo_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceScreenInfo_keyList"}).([]string)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceCaptureDeviceInfo_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceCaptureDeviceInfo_keyList_atList", i}).(string)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceCaptureDeviceInfo_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceCaptureDeviceInfo_keyList_setList", i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceCaptureDeviceInfo_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceCaptureDeviceInfo_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceScreenInfo_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceScreenInfo_keyList_atList", i}).(string)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceScreenInfo_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceScreenInfo_keyList_setList", i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceScreenInfo_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceScreenInfo_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardAbstractInputMethod) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QVirtualKeyboardExtensionPlugin struct {
	core.QObject
}

type QVirtualKeyboardExtensionPlugin_ITF interface {
	core.QObject_ITF
	QVirtualKeyboardExtensionPlugin_PTR() *QVirtualKeyboardExtensionPlugin
}

func (ptr *QVirtualKeyboardExtensionPlugin) QVirtualKeyboardExtensionPlugin_PTR() *QVirtualKeyboardExtensionPlugin {
	return ptr
}

func (ptr *QVirtualKeyboardExtensionPlugin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QVirtualKeyboardExtensionPlugin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQVirtualKeyboardExtensionPlugin(ptr QVirtualKeyboardExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVirtualKeyboardExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func (n *QVirtualKeyboardExtensionPlugin) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVirtualKeyboardExtensionPlugin) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQVirtualKeyboardExtensionPluginFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardExtensionPlugin) {
	n = new(QVirtualKeyboardExtensionPlugin)
	n.InitFromInternal(uintptr(ptr), "virtualkeyboard.QVirtualKeyboardExtensionPlugin")
	return
}

func (ptr *QVirtualKeyboardExtensionPlugin) DestroyQVirtualKeyboardExtensionPlugin() {
}

func (ptr *QVirtualKeyboardExtensionPlugin) ConnectRegisterTypes(f func(uri string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRegisterTypes", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardExtensionPlugin) DisconnectRegisterTypes() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRegisterTypes"})
}

func (ptr *QVirtualKeyboardExtensionPlugin) RegisterTypes(uri string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterTypes", uri})
}

func (ptr *QVirtualKeyboardExtensionPlugin) RegisterTypesDefault(uri string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterTypesDefault", uri})
}

func (ptr *QVirtualKeyboardExtensionPlugin) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardExtensionPlugin) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QVirtualKeyboardExtensionPlugin) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardExtensionPlugin) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QVirtualKeyboardExtensionPlugin) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QVirtualKeyboardExtensionPlugin) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QVirtualKeyboardExtensionPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QVirtualKeyboardExtensionPlugin) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QVirtualKeyboardExtensionPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardExtensionPlugin) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QVirtualKeyboardExtensionPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QVirtualKeyboardExtensionPlugin) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QVirtualKeyboardExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QVirtualKeyboardInputContext struct {
	core.QObject
}

type QVirtualKeyboardInputContext_ITF interface {
	core.QObject_ITF
	QVirtualKeyboardInputContext_PTR() *QVirtualKeyboardInputContext
}

func (ptr *QVirtualKeyboardInputContext) QVirtualKeyboardInputContext_PTR() *QVirtualKeyboardInputContext {
	return ptr
}

func (ptr *QVirtualKeyboardInputContext) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QVirtualKeyboardInputContext) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQVirtualKeyboardInputContext(ptr QVirtualKeyboardInputContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVirtualKeyboardInputContext_PTR().Pointer()
	}
	return nil
}

func (n *QVirtualKeyboardInputContext) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVirtualKeyboardInputContext) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQVirtualKeyboardInputContextFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardInputContext) {
	n = new(QVirtualKeyboardInputContext)
	n.InitFromInternal(uintptr(ptr), "virtualkeyboard.QVirtualKeyboardInputContext")
	return
}

func (ptr *QVirtualKeyboardInputContext) DestroyQVirtualKeyboardInputContext() {
}

func (ptr *QVirtualKeyboardInputContext) AnchorPosition() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnchorPosition"}).(float64))
}

func (ptr *QVirtualKeyboardInputContext) ConnectAnchorPositionChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAnchorPositionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectAnchorPositionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAnchorPositionChanged"})
}

func (ptr *QVirtualKeyboardInputContext) AnchorPositionChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnchorPositionChanged"})
}

func (ptr *QVirtualKeyboardInputContext) AnchorRectIntersectsClipRect() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnchorRectIntersectsClipRect"}).(bool)
}

func (ptr *QVirtualKeyboardInputContext) ConnectAnchorRectIntersectsClipRectChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAnchorRectIntersectsClipRectChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectAnchorRectIntersectsClipRectChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAnchorRectIntersectsClipRectChanged"})
}

func (ptr *QVirtualKeyboardInputContext) AnchorRectIntersectsClipRectChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnchorRectIntersectsClipRectChanged"})
}

func (ptr *QVirtualKeyboardInputContext) AnchorRectangle() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnchorRectangle"}).(*core.QRectF)
}

func (ptr *QVirtualKeyboardInputContext) ConnectAnchorRectangleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAnchorRectangleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectAnchorRectangleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAnchorRectangleChanged"})
}

func (ptr *QVirtualKeyboardInputContext) AnchorRectangleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnchorRectangleChanged"})
}

func (ptr *QVirtualKeyboardInputContext) ConnectAnimatingChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAnimatingChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectAnimatingChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAnimatingChanged"})
}

func (ptr *QVirtualKeyboardInputContext) AnimatingChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AnimatingChanged"})
}

func (ptr *QVirtualKeyboardInputContext) ConnectCapsLockActiveChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCapsLockActiveChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectCapsLockActiveChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCapsLockActiveChanged"})
}

func (ptr *QVirtualKeyboardInputContext) CapsLockActiveChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CapsLockActiveChanged"})
}

func (ptr *QVirtualKeyboardInputContext) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QVirtualKeyboardInputContext) Commit() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Commit"})
}

func (ptr *QVirtualKeyboardInputContext) Commit2(text string, replaceFrom int, replaceLength int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Commit2", text, replaceFrom, replaceLength})
}

func (ptr *QVirtualKeyboardInputContext) CursorPosition() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CursorPosition"}).(float64))
}

func (ptr *QVirtualKeyboardInputContext) ConnectCursorPositionChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCursorPositionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectCursorPositionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCursorPositionChanged"})
}

func (ptr *QVirtualKeyboardInputContext) CursorPositionChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CursorPositionChanged"})
}

func (ptr *QVirtualKeyboardInputContext) CursorRectIntersectsClipRect() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CursorRectIntersectsClipRect"}).(bool)
}

func (ptr *QVirtualKeyboardInputContext) ConnectCursorRectIntersectsClipRectChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCursorRectIntersectsClipRectChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectCursorRectIntersectsClipRectChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCursorRectIntersectsClipRectChanged"})
}

func (ptr *QVirtualKeyboardInputContext) CursorRectIntersectsClipRectChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CursorRectIntersectsClipRectChanged"})
}

func (ptr *QVirtualKeyboardInputContext) CursorRectangle() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CursorRectangle"}).(*core.QRectF)
}

func (ptr *QVirtualKeyboardInputContext) ConnectCursorRectangleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCursorRectangleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectCursorRectangleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCursorRectangleChanged"})
}

func (ptr *QVirtualKeyboardInputContext) CursorRectangleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CursorRectangleChanged"})
}

func (ptr *QVirtualKeyboardInputContext) InputEngine() *QVirtualKeyboardInputEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputEngine"}).(*QVirtualKeyboardInputEngine)
}

func (ptr *QVirtualKeyboardInputContext) InputItem() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputItem"}).(*core.QObject)
}

func (ptr *QVirtualKeyboardInputContext) ConnectInputItemChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInputItemChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectInputItemChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInputItemChanged"})
}

func (ptr *QVirtualKeyboardInputContext) InputItemChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputItemChanged"})
}

func (ptr *QVirtualKeyboardInputContext) InputMethodHints() core.Qt__InputMethodHint {

	return core.Qt__InputMethodHint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodHints"}).(float64))
}

func (ptr *QVirtualKeyboardInputContext) ConnectInputMethodHintsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInputMethodHintsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectInputMethodHintsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInputMethodHintsChanged"})
}

func (ptr *QVirtualKeyboardInputContext) InputMethodHintsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodHintsChanged"})
}

func (ptr *QVirtualKeyboardInputContext) IsAnimating() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAnimating"}).(bool)
}

func (ptr *QVirtualKeyboardInputContext) IsCapsLockActive() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCapsLockActive"}).(bool)
}

func (ptr *QVirtualKeyboardInputContext) IsSelectionControlVisible() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSelectionControlVisible"}).(bool)
}

func (ptr *QVirtualKeyboardInputContext) IsShiftActive() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsShiftActive"}).(bool)
}

func (ptr *QVirtualKeyboardInputContext) IsUppercase() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsUppercase"}).(bool)
}

func (ptr *QVirtualKeyboardInputContext) Locale() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Locale"}).(string)
}

func (ptr *QVirtualKeyboardInputContext) ConnectLocaleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLocaleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectLocaleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLocaleChanged"})
}

func (ptr *QVirtualKeyboardInputContext) LocaleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocaleChanged"})
}

func (ptr *QVirtualKeyboardInputContext) PreeditText() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreeditText"}).(string)
}

func (ptr *QVirtualKeyboardInputContext) ConnectPreeditTextChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreeditTextChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectPreeditTextChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreeditTextChanged"})
}

func (ptr *QVirtualKeyboardInputContext) PreeditTextChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreeditTextChanged"})
}

func (ptr *QVirtualKeyboardInputContext) SelectedText() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedText"}).(string)
}

func (ptr *QVirtualKeyboardInputContext) ConnectSelectedTextChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectedTextChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectSelectedTextChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectedTextChanged"})
}

func (ptr *QVirtualKeyboardInputContext) SelectedTextChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedTextChanged"})
}

func (ptr *QVirtualKeyboardInputContext) ConnectSelectionControlVisibleChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionControlVisibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectSelectionControlVisibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionControlVisibleChanged"})
}

func (ptr *QVirtualKeyboardInputContext) SelectionControlVisibleChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionControlVisibleChanged"})
}

func (ptr *QVirtualKeyboardInputContext) SendKeyClick(key int, text string, modifiers int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendKeyClick", key, text, modifiers})
}

func (ptr *QVirtualKeyboardInputContext) SetAnimating(isAnimating bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAnimating", isAnimating})
}

func (ptr *QVirtualKeyboardInputContext) ConnectShiftActiveChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShiftActiveChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectShiftActiveChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShiftActiveChanged"})
}

func (ptr *QVirtualKeyboardInputContext) ShiftActiveChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShiftActiveChanged"})
}

func (ptr *QVirtualKeyboardInputContext) SurroundingText() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SurroundingText"}).(string)
}

func (ptr *QVirtualKeyboardInputContext) ConnectSurroundingTextChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSurroundingTextChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectSurroundingTextChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSurroundingTextChanged"})
}

func (ptr *QVirtualKeyboardInputContext) SurroundingTextChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SurroundingTextChanged"})
}

func (ptr *QVirtualKeyboardInputContext) ConnectUppercaseChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUppercaseChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputContext) DisconnectUppercaseChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUppercaseChanged"})
}

func (ptr *QVirtualKeyboardInputContext) UppercaseChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UppercaseChanged"})
}

func (ptr *QVirtualKeyboardInputContext) __preeditTextAttributes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__preeditTextAttributes_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputContext) __setPreeditText_attributes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPreeditText_attributes_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputContext) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardInputContext) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QVirtualKeyboardInputContext) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputContext) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QVirtualKeyboardInputContext) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QVirtualKeyboardInputContext) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputContext) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QVirtualKeyboardInputContext) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardInputContext) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QVirtualKeyboardInputContext) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QVirtualKeyboardInputContext) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardInputContext) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QVirtualKeyboardInputContext) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QVirtualKeyboardInputContext) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QVirtualKeyboardInputContext) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QVirtualKeyboardInputEngine struct {
	core.QObject
}

type QVirtualKeyboardInputEngine_ITF interface {
	core.QObject_ITF
	QVirtualKeyboardInputEngine_PTR() *QVirtualKeyboardInputEngine
}

func (ptr *QVirtualKeyboardInputEngine) QVirtualKeyboardInputEngine_PTR() *QVirtualKeyboardInputEngine {
	return ptr
}

func (ptr *QVirtualKeyboardInputEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QVirtualKeyboardInputEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQVirtualKeyboardInputEngine(ptr QVirtualKeyboardInputEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVirtualKeyboardInputEngine_PTR().Pointer()
	}
	return nil
}

func (n *QVirtualKeyboardInputEngine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVirtualKeyboardInputEngine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQVirtualKeyboardInputEngineFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardInputEngine) {
	n = new(QVirtualKeyboardInputEngine)
	n.InitFromInternal(uintptr(ptr), "virtualkeyboard.QVirtualKeyboardInputEngine")
	return
}

func (ptr *QVirtualKeyboardInputEngine) DestroyQVirtualKeyboardInputEngine() {
}

//go:generate stringer -type=QVirtualKeyboardInputEngine__TextCase
//QVirtualKeyboardInputEngine::TextCase
type QVirtualKeyboardInputEngine__TextCase int64

const (
	QVirtualKeyboardInputEngine__Lower QVirtualKeyboardInputEngine__TextCase = QVirtualKeyboardInputEngine__TextCase(0)
	QVirtualKeyboardInputEngine__Upper QVirtualKeyboardInputEngine__TextCase = QVirtualKeyboardInputEngine__TextCase(1)
)

//go:generate stringer -type=QVirtualKeyboardInputEngine__InputMode
//QVirtualKeyboardInputEngine::InputMode
type QVirtualKeyboardInputEngine__InputMode int64

const (
	QVirtualKeyboardInputEngine__Latin               QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(0)
	QVirtualKeyboardInputEngine__Numeric             QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(1)
	QVirtualKeyboardInputEngine__Dialable            QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(2)
	QVirtualKeyboardInputEngine__Pinyin              QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(3)
	QVirtualKeyboardInputEngine__Cangjie             QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(4)
	QVirtualKeyboardInputEngine__Zhuyin              QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(5)
	QVirtualKeyboardInputEngine__Hangul              QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(6)
	QVirtualKeyboardInputEngine__Hiragana            QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(7)
	QVirtualKeyboardInputEngine__Katakana            QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(8)
	QVirtualKeyboardInputEngine__FullwidthLatin      QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(9)
	QVirtualKeyboardInputEngine__Greek               QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(10)
	QVirtualKeyboardInputEngine__Cyrillic            QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(11)
	QVirtualKeyboardInputEngine__Arabic              QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(12)
	QVirtualKeyboardInputEngine__Hebrew              QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(13)
	QVirtualKeyboardInputEngine__ChineseHandwriting  QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(14)
	QVirtualKeyboardInputEngine__JapaneseHandwriting QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(15)
	QVirtualKeyboardInputEngine__KoreanHandwriting   QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(16)
	QVirtualKeyboardInputEngine__Thai                QVirtualKeyboardInputEngine__InputMode = QVirtualKeyboardInputEngine__InputMode(17)
)

//go:generate stringer -type=QVirtualKeyboardInputEngine__PatternRecognitionMode
//QVirtualKeyboardInputEngine::PatternRecognitionMode
type QVirtualKeyboardInputEngine__PatternRecognitionMode int64

const (
	QVirtualKeyboardInputEngine__None                       QVirtualKeyboardInputEngine__PatternRecognitionMode = QVirtualKeyboardInputEngine__PatternRecognitionMode(0)
	QVirtualKeyboardInputEngine__PatternRecognitionDisabled QVirtualKeyboardInputEngine__PatternRecognitionMode = QVirtualKeyboardInputEngine__PatternRecognitionMode(QVirtualKeyboardInputEngine__None)
	QVirtualKeyboardInputEngine__Handwriting                QVirtualKeyboardInputEngine__PatternRecognitionMode = QVirtualKeyboardInputEngine__PatternRecognitionMode(1)
	QVirtualKeyboardInputEngine__HandwritingRecoginition    QVirtualKeyboardInputEngine__PatternRecognitionMode = QVirtualKeyboardInputEngine__PatternRecognitionMode(QVirtualKeyboardInputEngine__Handwriting)
)

//go:generate stringer -type=QVirtualKeyboardInputEngine__ReselectFlag
//QVirtualKeyboardInputEngine::ReselectFlag
type QVirtualKeyboardInputEngine__ReselectFlag int64

const (
	QVirtualKeyboardInputEngine__WordBeforeCursor QVirtualKeyboardInputEngine__ReselectFlag = QVirtualKeyboardInputEngine__ReselectFlag(0x1)
	QVirtualKeyboardInputEngine__WordAfterCursor  QVirtualKeyboardInputEngine__ReselectFlag = QVirtualKeyboardInputEngine__ReselectFlag(0x2)
	QVirtualKeyboardInputEngine__WordAtCursor     QVirtualKeyboardInputEngine__ReselectFlag = QVirtualKeyboardInputEngine__ReselectFlag(QVirtualKeyboardInputEngine__WordBeforeCursor | QVirtualKeyboardInputEngine__WordAfterCursor)
)

func (ptr *QVirtualKeyboardInputEngine) ActiveKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveKey"}).(float64))
}

func (ptr *QVirtualKeyboardInputEngine) ConnectActiveKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectActiveKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveKeyChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) ActiveKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveKeyChanged", key})
}

func (ptr *QVirtualKeyboardInputEngine) InputContext() *QVirtualKeyboardInputContext {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputContext"}).(*QVirtualKeyboardInputContext)
}

func (ptr *QVirtualKeyboardInputEngine) InputMethod() *QVirtualKeyboardAbstractInputMethod {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethod"}).(*QVirtualKeyboardAbstractInputMethod)
}

func (ptr *QVirtualKeyboardInputEngine) ConnectInputMethodChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInputMethodChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectInputMethodChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInputMethodChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) InputMethodChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) ConnectInputMethodReset(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInputMethodReset", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectInputMethodReset() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInputMethodReset"})
}

func (ptr *QVirtualKeyboardInputEngine) InputMethodReset() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodReset"})
}

func (ptr *QVirtualKeyboardInputEngine) ConnectInputMethodUpdate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInputMethodUpdate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectInputMethodUpdate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInputMethodUpdate"})
}

func (ptr *QVirtualKeyboardInputEngine) InputMethodUpdate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodUpdate"})
}

func (ptr *QVirtualKeyboardInputEngine) InputMode() QVirtualKeyboardInputEngine__InputMode {

	return QVirtualKeyboardInputEngine__InputMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMode"}).(float64))
}

func (ptr *QVirtualKeyboardInputEngine) ConnectInputModeChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInputModeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectInputModeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInputModeChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) InputModeChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputModeChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) InputModes() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputModes"}).([]int)
}

func (ptr *QVirtualKeyboardInputEngine) ConnectInputModesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInputModesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectInputModesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInputModesChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) InputModesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputModesChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) PatternRecognitionModes() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PatternRecognitionModes"}).([]int)
}

func (ptr *QVirtualKeyboardInputEngine) ConnectPatternRecognitionModesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPatternRecognitionModesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectPatternRecognitionModesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPatternRecognitionModesChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) PatternRecognitionModesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PatternRecognitionModesChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) PreviousKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreviousKey"}).(float64))
}

func (ptr *QVirtualKeyboardInputEngine) ConnectPreviousKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreviousKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectPreviousKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreviousKeyChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) PreviousKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreviousKeyChanged", key})
}

func (ptr *QVirtualKeyboardInputEngine) Reselect(cursorPosition int, reselectFlags QVirtualKeyboardInputEngine__ReselectFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reselect", cursorPosition, reselectFlags}).(bool)
}

func (ptr *QVirtualKeyboardInputEngine) SetInputMethod(inputMethod QVirtualKeyboardAbstractInputMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInputMethod", inputMethod})
}

func (ptr *QVirtualKeyboardInputEngine) SetInputMode(inputMode QVirtualKeyboardInputEngine__InputMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInputMode", inputMode})
}

func (ptr *QVirtualKeyboardInputEngine) TraceBegin(traceId int, patternRecognitionMode QVirtualKeyboardInputEngine__PatternRecognitionMode, traceCaptureDeviceInfo map[string]*core.QVariant, traceScreenInfo map[string]*core.QVariant) *QVirtualKeyboardTrace {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TraceBegin", traceId, patternRecognitionMode, traceCaptureDeviceInfo, traceScreenInfo}).(*QVirtualKeyboardTrace)
}

func (ptr *QVirtualKeyboardInputEngine) TraceEnd(trace QVirtualKeyboardTrace_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TraceEnd", trace}).(bool)
}

func (ptr *QVirtualKeyboardInputEngine) VirtualKeyCancel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VirtualKeyCancel"})
}

func (ptr *QVirtualKeyboardInputEngine) VirtualKeyClick(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VirtualKeyClick", key, text, modifiers}).(bool)
}

func (ptr *QVirtualKeyboardInputEngine) ConnectVirtualKeyClicked(f func(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier, isAutoRepeat bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVirtualKeyClicked", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectVirtualKeyClicked() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVirtualKeyClicked"})
}

func (ptr *QVirtualKeyboardInputEngine) VirtualKeyClicked(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier, isAutoRepeat bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VirtualKeyClicked", key, text, modifiers, isAutoRepeat})
}

func (ptr *QVirtualKeyboardInputEngine) VirtualKeyPress(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier, repe bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VirtualKeyPress", key, text, modifiers, repe}).(bool)
}

func (ptr *QVirtualKeyboardInputEngine) VirtualKeyRelease(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VirtualKeyRelease", key, text, modifiers}).(bool)
}

func (ptr *QVirtualKeyboardInputEngine) WordCandidateListModel() *QVirtualKeyboardSelectionListModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WordCandidateListModel"}).(*QVirtualKeyboardSelectionListModel)
}

func (ptr *QVirtualKeyboardInputEngine) ConnectWordCandidateListModelChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWordCandidateListModelChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectWordCandidateListModelChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWordCandidateListModelChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) WordCandidateListModelChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WordCandidateListModelChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) WordCandidateListVisibleHint() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WordCandidateListVisibleHint"}).(bool)
}

func (ptr *QVirtualKeyboardInputEngine) ConnectWordCandidateListVisibleHintChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWordCandidateListVisibleHintChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectWordCandidateListVisibleHintChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWordCandidateListVisibleHintChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) WordCandidateListVisibleHintChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WordCandidateListVisibleHintChanged"})
}

func (ptr *QVirtualKeyboardInputEngine) __inputModes_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__inputModes_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardInputEngine) __inputModes_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__inputModes_setList", i})
}

func (ptr *QVirtualKeyboardInputEngine) __inputModes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__inputModes_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputEngine) __patternRecognitionModes_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__patternRecognitionModes_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardInputEngine) __patternRecognitionModes_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__patternRecognitionModes_setList", i})
}

func (ptr *QVirtualKeyboardInputEngine) __patternRecognitionModes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__patternRecognitionModes_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceCaptureDeviceInfo_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceCaptureDeviceInfo_atList", v, i}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceCaptureDeviceInfo_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceCaptureDeviceInfo_setList", key, i})
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceCaptureDeviceInfo_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceCaptureDeviceInfo_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceCaptureDeviceInfo_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceCaptureDeviceInfo_keyList"}).([]string)
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceScreenInfo_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceScreenInfo_atList", v, i}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceScreenInfo_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceScreenInfo_setList", key, i})
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceScreenInfo_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceScreenInfo_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceScreenInfo_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__traceBegin_traceScreenInfo_keyList"}).([]string)
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceCaptureDeviceInfo_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceCaptureDeviceInfo_keyList_atList", i}).(string)
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceCaptureDeviceInfo_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceCaptureDeviceInfo_keyList_setList", i})
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceCaptureDeviceInfo_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceCaptureDeviceInfo_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceScreenInfo_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceScreenInfo_keyList_atList", i}).(string)
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceScreenInfo_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceScreenInfo_keyList_setList", i})
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceScreenInfo_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____traceBegin_traceScreenInfo_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputEngine) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardInputEngine) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QVirtualKeyboardInputEngine) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QVirtualKeyboardInputEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QVirtualKeyboardInputEngine) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardInputEngine) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QVirtualKeyboardInputEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardInputEngine) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QVirtualKeyboardInputEngine) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardInputEngine) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QVirtualKeyboardInputEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QVirtualKeyboardInputEngine) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QVirtualKeyboardInputEngine) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QVirtualKeyboardSelectionListModel struct {
	core.QAbstractListModel
}

type QVirtualKeyboardSelectionListModel_ITF interface {
	core.QAbstractListModel_ITF
	QVirtualKeyboardSelectionListModel_PTR() *QVirtualKeyboardSelectionListModel
}

func (ptr *QVirtualKeyboardSelectionListModel) QVirtualKeyboardSelectionListModel_PTR() *QVirtualKeyboardSelectionListModel {
	return ptr
}

func (ptr *QVirtualKeyboardSelectionListModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromQVirtualKeyboardSelectionListModel(ptr QVirtualKeyboardSelectionListModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVirtualKeyboardSelectionListModel_PTR().Pointer()
	}
	return nil
}

func (n *QVirtualKeyboardSelectionListModel) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractListModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVirtualKeyboardSelectionListModel) ClassNameInternalF() string {
	return n.QAbstractListModel_PTR().ClassNameInternalF()
}

func NewQVirtualKeyboardSelectionListModelFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardSelectionListModel) {
	n = new(QVirtualKeyboardSelectionListModel)
	n.InitFromInternal(uintptr(ptr), "virtualkeyboard.QVirtualKeyboardSelectionListModel")
	return
}

func (ptr *QVirtualKeyboardSelectionListModel) DestroyQVirtualKeyboardSelectionListModel() {
}

//go:generate stringer -type=QVirtualKeyboardSelectionListModel__Type
//QVirtualKeyboardSelectionListModel::Type
type QVirtualKeyboardSelectionListModel__Type int64

const (
	QVirtualKeyboardSelectionListModel__WordCandidateList QVirtualKeyboardSelectionListModel__Type = QVirtualKeyboardSelectionListModel__Type(0)
)

//go:generate stringer -type=QVirtualKeyboardSelectionListModel__Role
//QVirtualKeyboardSelectionListModel::Role
type QVirtualKeyboardSelectionListModel__Role int64

const (
	QVirtualKeyboardSelectionListModel__Display             QVirtualKeyboardSelectionListModel__Role = QVirtualKeyboardSelectionListModel__Role(core.Qt__DisplayRole)
	QVirtualKeyboardSelectionListModel__DisplayRole         QVirtualKeyboardSelectionListModel__Role = QVirtualKeyboardSelectionListModel__Role(QVirtualKeyboardSelectionListModel__Display)
	QVirtualKeyboardSelectionListModel__Dictionary          QVirtualKeyboardSelectionListModel__Role = QVirtualKeyboardSelectionListModel__Role(258)
	QVirtualKeyboardSelectionListModel__CanRemoveSuggestion QVirtualKeyboardSelectionListModel__Role = QVirtualKeyboardSelectionListModel__Role(259)
)

//go:generate stringer -type=QVirtualKeyboardSelectionListModel__DictionaryType
//QVirtualKeyboardSelectionListModel::DictionaryType
type QVirtualKeyboardSelectionListModel__DictionaryType int64

const (
	QVirtualKeyboardSelectionListModel__Default QVirtualKeyboardSelectionListModel__DictionaryType = QVirtualKeyboardSelectionListModel__DictionaryType(0)
	QVirtualKeyboardSelectionListModel__User    QVirtualKeyboardSelectionListModel__DictionaryType = QVirtualKeyboardSelectionListModel__DictionaryType(1)
)

func (ptr *QVirtualKeyboardSelectionListModel) ConnectActiveItemChanged(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveItemChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardSelectionListModel) DisconnectActiveItemChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveItemChanged"})
}

func (ptr *QVirtualKeyboardSelectionListModel) ActiveItemChanged(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveItemChanged", index})
}

func (ptr *QVirtualKeyboardSelectionListModel) ConnectItemSelected(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectItemSelected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardSelectionListModel) DisconnectItemSelected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectItemSelected"})
}

func (ptr *QVirtualKeyboardSelectionListModel) ItemSelected(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemSelected", index})
}

func (ptr *QVirtualKeyboardSelectionListModel) RemoveItem(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveItem", index})
}

func (ptr *QVirtualKeyboardSelectionListModel) SelectItem(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectItem", index})
}

func (ptr *QVirtualKeyboardSelectionListModel) __roleNames_atList(v int, i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_atList", v, i}).(*core.QByteArray)
}

func (ptr *QVirtualKeyboardSelectionListModel) __roleNames_setList(key int, i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_setList", key, i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __roleNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __roleNames_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_keyList"}).([]int)
}

func (ptr *QVirtualKeyboardSelectionListModel) ____roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) ____roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) ____roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) ____itemData_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) ____itemData_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) ____itemData_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setItemData_roles_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setItemData_roles_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_atList", i}).(*core.QModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_from_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_from_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_to_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_atList", i}).(*core.QModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_to_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_to_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __dataChanged_roles_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) __dataChanged_roles_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __dataChanged_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __itemData_atList(v int, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_atList", v, i}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardSelectionListModel) __itemData_setList(key int, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_setList", key, i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __itemData_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __itemData_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_keyList"}).([]int)
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_atList", i}).(*core.QPersistentModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutAboutToBeChanged_parents_setList(i core.QPersistentModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutChanged_parents_atList(i int) *core.QPersistentModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_atList", i}).(*core.QPersistentModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutChanged_parents_setList(i core.QPersistentModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutChanged_parents_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __match_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_atList", i}).(*core.QModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) __match_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __match_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __mimeData_indexes_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_atList", i}).(*core.QModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) __mimeData_indexes_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __mimeData_indexes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __persistentIndexList_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_atList", i}).(*core.QModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) __persistentIndexList_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __persistentIndexList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __setItemData_roles_atList(v int, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_atList", v, i}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardSelectionListModel) __setItemData_roles_setList(key int, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_setList", key, i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __setItemData_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __setItemData_roles_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_keyList"}).([]int)
}

func (ptr *QVirtualKeyboardSelectionListModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setRoleNames_roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_atList", i}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setRoleNames_roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardSelectionListModel) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QVirtualKeyboardSelectionListModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardSelectionListModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropMimeDataDefault", data, action, row, column, parent}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {

	return core.Qt__ItemFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlagsDefault", index}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexDefault", row, column, parent}).(*core.QModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SiblingDefault", row, column, idx}).(*core.QModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BuddyDefault", index}).(*core.QModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanDropMimeDataDefault", data, action, row, column, parent}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanFetchMoreDefault", parent}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) ColumnCountDefault(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountDefault", parent}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataDefault", index, role}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardSelectionListModel) FetchMoreDefault(parent core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchMoreDefault", parent})
}

func (ptr *QVirtualKeyboardSelectionListModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasChildrenDefault", parent}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeaderDataDefault", section, orientation, role}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardSelectionListModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertColumnsDefault", column, count, parent}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertRowsDefault", row, count, parent}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) ItemDataDefault(index core.QModelIndex_ITF) map[int]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemDataDefault", index}).(map[int]*core.QVariant)
}

func (ptr *QVirtualKeyboardSelectionListModel) MatchDefault(start core.QModelIndex_ITF, role int, value core.QVariant_ITF, hits int, flags core.Qt__MatchFlag) []*core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MatchDefault", start, role, value, hits, flags}).([]*core.QModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) MimeDataDefault(indexes []*core.QModelIndex) *core.QMimeData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeDataDefault", indexes}).(*core.QMimeData)
}

func (ptr *QVirtualKeyboardSelectionListModel) MimeTypesDefault() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeTypesDefault"}).([]string)
}

func (ptr *QVirtualKeyboardSelectionListModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveColumnsDefault", sourceParent, sourceColumn, count, destinationParent, destinationChild}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveRowsDefault", sourceParent, sourceRow, count, destinationParent, destinationChild}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParentDefault", index}).(*core.QModelIndex)
}

func (ptr *QVirtualKeyboardSelectionListModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveColumnsDefault", column, count, parent}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveRowsDefault", row, count, parent}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) ResetInternalDataDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetInternalDataDefault"})
}

func (ptr *QVirtualKeyboardSelectionListModel) RevertDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertDefault"})
}

func (ptr *QVirtualKeyboardSelectionListModel) RoleNamesDefault() map[int]*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RoleNamesDefault"}).(map[int]*core.QByteArray)
}

func (ptr *QVirtualKeyboardSelectionListModel) RowCountDefault(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountDefault", parent}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataDefault", index, value, role}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeaderDataDefault", section, orientation, value, role}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) SetItemDataDefault(index core.QModelIndex_ITF, roles map[int]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemDataDefault", index, roles}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) SortDefault(column int, order core.Qt__SortOrder) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SortDefault", column, order})
}

func (ptr *QVirtualKeyboardSelectionListModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SpanDefault", index}).(*core.QSize)
}

func (ptr *QVirtualKeyboardSelectionListModel) SubmitDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubmitDefault"}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) SupportedDragActionsDefault() core.Qt__DropAction {

	return core.Qt__DropAction(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDragActionsDefault"}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) SupportedDropActionsDefault() core.Qt__DropAction {

	return core.Qt__DropAction(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDropActionsDefault"}).(float64))
}

func (ptr *QVirtualKeyboardSelectionListModel) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QVirtualKeyboardSelectionListModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardSelectionListModel) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QVirtualKeyboardSelectionListModel) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QVirtualKeyboardSelectionListModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardSelectionListModel) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QVirtualKeyboardSelectionListModel) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QVirtualKeyboardSelectionListModel) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QVirtualKeyboardTrace struct {
	core.QObject
}

type QVirtualKeyboardTrace_ITF interface {
	core.QObject_ITF
	QVirtualKeyboardTrace_PTR() *QVirtualKeyboardTrace
}

func (ptr *QVirtualKeyboardTrace) QVirtualKeyboardTrace_PTR() *QVirtualKeyboardTrace {
	return ptr
}

func (ptr *QVirtualKeyboardTrace) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QVirtualKeyboardTrace) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQVirtualKeyboardTrace(ptr QVirtualKeyboardTrace_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVirtualKeyboardTrace_PTR().Pointer()
	}
	return nil
}

func (n *QVirtualKeyboardTrace) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QVirtualKeyboardTrace) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQVirtualKeyboardTraceFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardTrace) {
	n = new(QVirtualKeyboardTrace)
	n.InitFromInternal(uintptr(ptr), "virtualkeyboard.QVirtualKeyboardTrace")
	return
}

func (ptr *QVirtualKeyboardTrace) DestroyQVirtualKeyboardTrace() {
}

func (ptr *QVirtualKeyboardTrace) AddPoint(point core.QPointF_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddPoint", point}).(float64))
}

func (ptr *QVirtualKeyboardTrace) ConnectCanceledChanged(f func(isCanceled bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCanceledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardTrace) DisconnectCanceledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCanceledChanged"})
}

func (ptr *QVirtualKeyboardTrace) CanceledChanged(isCanceled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanceledChanged", isCanceled})
}

func (ptr *QVirtualKeyboardTrace) ChannelData(channel string, pos int, count int) []*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChannelData", channel, pos, count}).([]*core.QVariant)
}

func (ptr *QVirtualKeyboardTrace) Channels() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Channels"}).([]string)
}

func (ptr *QVirtualKeyboardTrace) ConnectChannelsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectChannelsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardTrace) DisconnectChannelsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectChannelsChanged"})
}

func (ptr *QVirtualKeyboardTrace) ChannelsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChannelsChanged"})
}

func (ptr *QVirtualKeyboardTrace) ConnectFinalChanged(f func(isFinal bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinalChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardTrace) DisconnectFinalChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinalChanged"})
}

func (ptr *QVirtualKeyboardTrace) FinalChanged(isFinal bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FinalChanged", isFinal})
}

func (ptr *QVirtualKeyboardTrace) IsCanceled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCanceled"}).(bool)
}

func (ptr *QVirtualKeyboardTrace) IsFinal() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFinal"}).(bool)
}

func (ptr *QVirtualKeyboardTrace) Length() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Length"}).(float64))
}

func (ptr *QVirtualKeyboardTrace) ConnectLengthChanged(f func(length int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLengthChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardTrace) DisconnectLengthChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLengthChanged"})
}

func (ptr *QVirtualKeyboardTrace) LengthChanged(length int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LengthChanged", length})
}

func (ptr *QVirtualKeyboardTrace) Opacity() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Opacity"}).(float64)
}

func (ptr *QVirtualKeyboardTrace) ConnectOpacityChanged(f func(opacity float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpacityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardTrace) DisconnectOpacityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpacityChanged"})
}

func (ptr *QVirtualKeyboardTrace) OpacityChanged(opacity float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpacityChanged", opacity})
}

func (ptr *QVirtualKeyboardTrace) Points(pos int, count int) []*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Points", pos, count}).([]*core.QVariant)
}

func (ptr *QVirtualKeyboardTrace) SetCanceled(canceled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCanceled", canceled})
}

func (ptr *QVirtualKeyboardTrace) SetChannelData(channel string, index int, data core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetChannelData", channel, index, data})
}

func (ptr *QVirtualKeyboardTrace) SetChannels(channels []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetChannels", channels})
}

func (ptr *QVirtualKeyboardTrace) SetFinal(fin bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFinal", fin})
}

func (ptr *QVirtualKeyboardTrace) SetOpacity(opacity float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpacity", opacity})
}

func (ptr *QVirtualKeyboardTrace) SetTraceId(id int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTraceId", id})
}

func (ptr *QVirtualKeyboardTrace) TraceId() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TraceId"}).(float64))
}

func (ptr *QVirtualKeyboardTrace) ConnectTraceIdChanged(f func(traceId int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTraceIdChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QVirtualKeyboardTrace) DisconnectTraceIdChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTraceIdChanged"})
}

func (ptr *QVirtualKeyboardTrace) TraceIdChanged(traceId int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TraceIdChanged", traceId})
}

func (ptr *QVirtualKeyboardTrace) __channelData_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__channelData_atList", i}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardTrace) __channelData_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__channelData_setList", i})
}

func (ptr *QVirtualKeyboardTrace) __channelData_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__channelData_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardTrace) __points_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__points_atList", i}).(*core.QVariant)
}

func (ptr *QVirtualKeyboardTrace) __points_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__points_setList", i})
}

func (ptr *QVirtualKeyboardTrace) __points_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__points_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardTrace) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardTrace) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QVirtualKeyboardTrace) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardTrace) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QVirtualKeyboardTrace) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QVirtualKeyboardTrace) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardTrace) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardTrace) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QVirtualKeyboardTrace) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardTrace) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QVirtualKeyboardTrace) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QVirtualKeyboardTrace) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QVirtualKeyboardTrace) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QVirtualKeyboardTrace) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardTrace) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QVirtualKeyboardTrace) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QVirtualKeyboardTrace) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QVirtualKeyboardTrace) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QVirtualKeyboardTrace) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QVirtualKeyboardTrace) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QVirtualKeyboardTrace) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func init() {
	internal.ConstructorTable["virtualkeyboard.QVirtualKeyboardAbstractInputMethod"] = NewQVirtualKeyboardAbstractInputMethodFromPointer
	internal.ConstructorTable["virtualkeyboard.QVirtualKeyboardExtensionPlugin"] = NewQVirtualKeyboardExtensionPluginFromPointer
	internal.ConstructorTable["virtualkeyboard.QVirtualKeyboardInputContext"] = NewQVirtualKeyboardInputContextFromPointer
	internal.ConstructorTable["virtualkeyboard.QVirtualKeyboardInputEngine"] = NewQVirtualKeyboardInputEngineFromPointer
	internal.ConstructorTable["virtualkeyboard.QVirtualKeyboardSelectionListModel"] = NewQVirtualKeyboardSelectionListModelFromPointer
	internal.ConstructorTable["virtualkeyboard.QVirtualKeyboardTrace"] = NewQVirtualKeyboardTraceFromPointer
}
