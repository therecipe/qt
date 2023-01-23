// +build !minimal

package macextras

import (
	"github.com/bluszcz/cutego"
	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
	"strings"
	"unsafe"
)

type QMacPasteboardMime struct {
	internal.Internal
}

type QMacPasteboardMime_ITF interface {
	QMacPasteboardMime_PTR() *QMacPasteboardMime
}

func (ptr *QMacPasteboardMime) QMacPasteboardMime_PTR() *QMacPasteboardMime {
	return ptr
}

func (ptr *QMacPasteboardMime) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QMacPasteboardMime) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQMacPasteboardMime(ptr QMacPasteboardMime_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacPasteboardMime_PTR().Pointer()
	}
	return nil
}

func (n *QMacPasteboardMime) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQMacPasteboardMimeFromPointer(ptr unsafe.Pointer) (n *QMacPasteboardMime) {
	n = new(QMacPasteboardMime)
	n.InitFromInternal(uintptr(ptr), "macextras.QMacPasteboardMime")
	return
}

func (ptr *QMacPasteboardMime) DestroyQMacPasteboardMime() {
}

func (ptr *QMacPasteboardMime) ConnectCanConvert(f func(mime string, flav string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCanConvert", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMacPasteboardMime) DisconnectCanConvert() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCanConvert"})
}

func (ptr *QMacPasteboardMime) CanConvert(mime string, flav string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanConvert", mime, flav}).(bool)
}

func (ptr *QMacPasteboardMime) ConnectConvertFromMime(f func(mime string, data *core.QVariant, flav string) []*core.QByteArray) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConvertFromMime", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMacPasteboardMime) DisconnectConvertFromMime() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConvertFromMime"})
}

func (ptr *QMacPasteboardMime) ConvertFromMime(mime string, data core.QVariant_ITF, flav string) []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConvertFromMime", mime, data, flav}).([]*core.QByteArray)
}

func (ptr *QMacPasteboardMime) ConnectConvertToMime(f func(mime string, data []*core.QByteArray, flav string) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConvertToMime", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMacPasteboardMime) DisconnectConvertToMime() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConvertToMime"})
}

func (ptr *QMacPasteboardMime) ConvertToMime(mime string, data []*core.QByteArray, flav string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConvertToMime", mime, data, flav}).(*core.QVariant)
}

func (ptr *QMacPasteboardMime) ConnectConvertorName(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConvertorName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMacPasteboardMime) DisconnectConvertorName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConvertorName"})
}

func (ptr *QMacPasteboardMime) ConvertorName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConvertorName"}).(string)
}

func (ptr *QMacPasteboardMime) ConnectFlavorFor(f func(mime string) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFlavorFor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMacPasteboardMime) DisconnectFlavorFor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFlavorFor"})
}

func (ptr *QMacPasteboardMime) FlavorFor(mime string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlavorFor", mime}).(string)
}

func (ptr *QMacPasteboardMime) ConnectMimeFor(f func(flav string) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMimeFor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMacPasteboardMime) DisconnectMimeFor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMimeFor"})
}

func (ptr *QMacPasteboardMime) MimeFor(flav string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeFor", flav}).(string)
}

func (ptr *QMacPasteboardMime) __convertFromMime_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__convertFromMime_atList", i}).(*core.QByteArray)
}

func (ptr *QMacPasteboardMime) __convertFromMime_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__convertFromMime_setList", i})
}

func (ptr *QMacPasteboardMime) __convertFromMime_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__convertFromMime_newList"}).(unsafe.Pointer)
}

func (ptr *QMacPasteboardMime) __convertToMime_data_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__convertToMime_data_atList", i}).(*core.QByteArray)
}

func (ptr *QMacPasteboardMime) __convertToMime_data_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__convertToMime_data_setList", i})
}

func (ptr *QMacPasteboardMime) __convertToMime_data_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__convertToMime_data_newList"}).(unsafe.Pointer)
}

type QMacToolBar struct {
	core.QObject
}

type QMacToolBar_ITF interface {
	core.QObject_ITF
	QMacToolBar_PTR() *QMacToolBar
}

func (ptr *QMacToolBar) QMacToolBar_PTR() *QMacToolBar {
	return ptr
}

func (ptr *QMacToolBar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QMacToolBar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQMacToolBar(ptr QMacToolBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBar_PTR().Pointer()
	}
	return nil
}

func (n *QMacToolBar) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QMacToolBar) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQMacToolBarFromPointer(ptr unsafe.Pointer) (n *QMacToolBar) {
	n = new(QMacToolBar)
	n.InitFromInternal(uintptr(ptr), "macextras.QMacToolBar")
	return
}

func (ptr *QMacToolBar) DestroyQMacToolBar() {
}

func (ptr *QMacToolBar) __allowedItems_atList(i int) *QMacToolBarItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allowedItems_atList", i}).(*QMacToolBarItem)
}

func (ptr *QMacToolBar) __allowedItems_setList(i QMacToolBarItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allowedItems_setList", i})
}

func (ptr *QMacToolBar) __allowedItems_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allowedItems_newList"}).(unsafe.Pointer)
}

func (ptr *QMacToolBar) __items_atList(i int) *QMacToolBarItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_atList", i}).(*QMacToolBarItem)
}

func (ptr *QMacToolBar) __items_setList(i QMacToolBarItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_setList", i})
}

func (ptr *QMacToolBar) __items_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__items_newList"}).(unsafe.Pointer)
}

func (ptr *QMacToolBar) __setAllowedItems_allowedItems_atList(i int) *QMacToolBarItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAllowedItems_allowedItems_atList", i}).(*QMacToolBarItem)
}

func (ptr *QMacToolBar) __setAllowedItems_allowedItems_setList(i QMacToolBarItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAllowedItems_allowedItems_setList", i})
}

func (ptr *QMacToolBar) __setAllowedItems_allowedItems_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAllowedItems_allowedItems_newList"}).(unsafe.Pointer)
}

func (ptr *QMacToolBar) __setItems_items_atList(i int) *QMacToolBarItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItems_items_atList", i}).(*QMacToolBarItem)
}

func (ptr *QMacToolBar) __setItems_items_setList(i QMacToolBarItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItems_items_setList", i})
}

func (ptr *QMacToolBar) __setItems_items_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItems_items_newList"}).(unsafe.Pointer)
}

func (ptr *QMacToolBar) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QMacToolBar) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QMacToolBar) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QMacToolBar) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QMacToolBar) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QMacToolBar) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QMacToolBar) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QMacToolBar) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QMacToolBar) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QMacToolBar) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QMacToolBar) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QMacToolBar) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QMacToolBar) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QMacToolBar) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QMacToolBar) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QMacToolBar) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QMacToolBar) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QMacToolBar) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QMacToolBar) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QMacToolBar) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QMacToolBar) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QMacToolBarItem struct {
	core.QObject
}

type QMacToolBarItem_ITF interface {
	core.QObject_ITF
	QMacToolBarItem_PTR() *QMacToolBarItem
}

func (ptr *QMacToolBarItem) QMacToolBarItem_PTR() *QMacToolBarItem {
	return ptr
}

func (ptr *QMacToolBarItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QMacToolBarItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQMacToolBarItem(ptr QMacToolBarItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBarItem_PTR().Pointer()
	}
	return nil
}

func (n *QMacToolBarItem) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QMacToolBarItem) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQMacToolBarItemFromPointer(ptr unsafe.Pointer) (n *QMacToolBarItem) {
	n = new(QMacToolBarItem)
	n.InitFromInternal(uintptr(ptr), "macextras.QMacToolBarItem")
	return
}

//go:generate stringer -type=QMacToolBarItem__StandardItem
//QMacToolBarItem::StandardItem
type QMacToolBarItem__StandardItem int64

const (
	QMacToolBarItem__NoStandardItem QMacToolBarItem__StandardItem = QMacToolBarItem__StandardItem(0)
	QMacToolBarItem__Space          QMacToolBarItem__StandardItem = QMacToolBarItem__StandardItem(1)
	QMacToolBarItem__FlexibleSpace  QMacToolBarItem__StandardItem = QMacToolBarItem__StandardItem(2)
)

func NewQMacToolBarItem(parent core.QObject_ITF) *QMacToolBarItem {

	return internal.CallLocalFunction([]interface{}{"", "", "macextras.NewQMacToolBarItem", "", parent}).(*QMacToolBarItem)
}

func (ptr *QMacToolBarItem) ConnectActivated(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActivated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMacToolBarItem) DisconnectActivated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActivated"})
}

func (ptr *QMacToolBarItem) Activated() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Activated"})
}

func (ptr *QMacToolBarItem) Icon() *gui.QIcon {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Icon"}).(*gui.QIcon)
}

func (ptr *QMacToolBarItem) Selectable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Selectable"}).(bool)
}

func (ptr *QMacToolBarItem) SetIcon(icon gui.QIcon_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIcon", icon})
}

func (ptr *QMacToolBarItem) SetSelectable(selectable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectable", selectable})
}

func (ptr *QMacToolBarItem) SetStandardItem(standardItem QMacToolBarItem__StandardItem) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStandardItem", standardItem})
}

func (ptr *QMacToolBarItem) SetText(text string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetText", text})
}

func (ptr *QMacToolBarItem) StandardItem() QMacToolBarItem__StandardItem {

	return QMacToolBarItem__StandardItem(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StandardItem"}).(float64))
}

func (ptr *QMacToolBarItem) Text() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Text"}).(string)
}

func (ptr *QMacToolBarItem) ConnectDestroyQMacToolBarItem(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQMacToolBarItem", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMacToolBarItem) DisconnectDestroyQMacToolBarItem() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQMacToolBarItem"})
}

func (ptr *QMacToolBarItem) DestroyQMacToolBarItem() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQMacToolBarItem"})
}

func (ptr *QMacToolBarItem) DestroyQMacToolBarItemDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQMacToolBarItemDefault"})
}

func (ptr *QMacToolBarItem) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QMacToolBarItem) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QMacToolBarItem) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QMacToolBarItem) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QMacToolBarItem) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QMacToolBarItem) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QMacToolBarItem) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QMacToolBarItem) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QMacToolBarItem) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QMacToolBarItem) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QMacToolBarItem) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QMacToolBarItem) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QMacToolBarItem) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QMacToolBarItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QMacToolBarItem) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QMacToolBarItem) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QMacToolBarItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QMacToolBarItem) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QMacToolBarItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QMacToolBarItem) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QMacToolBarItem) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func init() {
	internal.ConstructorTable["macextras.QMacPasteboardMime"] = NewQMacPasteboardMimeFromPointer
	internal.ConstructorTable["macextras.QMacToolBar"] = NewQMacToolBarFromPointer
	internal.ConstructorTable["macextras.QMacToolBarItem"] = NewQMacToolBarItemFromPointer
}
