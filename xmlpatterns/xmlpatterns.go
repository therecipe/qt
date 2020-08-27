// +build !minimal

package xmlpatterns

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QAbstractMessageHandler struct {
	core.QObject
}

type QAbstractMessageHandler_ITF interface {
	core.QObject_ITF
	QAbstractMessageHandler_PTR() *QAbstractMessageHandler
}

func (ptr *QAbstractMessageHandler) QAbstractMessageHandler_PTR() *QAbstractMessageHandler {
	return ptr
}

func (ptr *QAbstractMessageHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractMessageHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractMessageHandler(ptr QAbstractMessageHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractMessageHandler_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractMessageHandler) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAbstractMessageHandler) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQAbstractMessageHandlerFromPointer(ptr unsafe.Pointer) (n *QAbstractMessageHandler) {
	n = new(QAbstractMessageHandler)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QAbstractMessageHandler")
	return
}
func (ptr *QAbstractMessageHandler) ConnectDestroyQAbstractMessageHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractMessageHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractMessageHandler) DisconnectDestroyQAbstractMessageHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractMessageHandler"})
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractMessageHandler"})
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractMessageHandlerDefault"})
}

func (ptr *QAbstractMessageHandler) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QAbstractMessageHandler) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QAbstractMessageHandler) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractMessageHandler) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QAbstractMessageHandler) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QAbstractMessageHandler) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractMessageHandler) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QAbstractMessageHandler) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QAbstractMessageHandler) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractMessageHandler) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QAbstractMessageHandler) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QAbstractMessageHandler) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QAbstractMessageHandler) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QAbstractMessageHandler) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QAbstractMessageHandler) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QAbstractMessageHandler) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QAbstractMessageHandler) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QAbstractMessageHandler) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QAbstractMessageHandler) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QAbstractMessageHandler) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QAbstractMessageHandler) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QAbstractUriResolver struct {
	core.QObject
}

type QAbstractUriResolver_ITF interface {
	core.QObject_ITF
	QAbstractUriResolver_PTR() *QAbstractUriResolver
}

func (ptr *QAbstractUriResolver) QAbstractUriResolver_PTR() *QAbstractUriResolver {
	return ptr
}

func (ptr *QAbstractUriResolver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractUriResolver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractUriResolver(ptr QAbstractUriResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractUriResolver_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractUriResolver) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAbstractUriResolver) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQAbstractUriResolverFromPointer(ptr unsafe.Pointer) (n *QAbstractUriResolver) {
	n = new(QAbstractUriResolver)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QAbstractUriResolver")
	return
}
func NewQAbstractUriResolver(parent core.QObject_ITF) *QAbstractUriResolver {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQAbstractUriResolver", "", parent}).(*QAbstractUriResolver)
}

func (ptr *QAbstractUriResolver) ConnectResolve(f func(relative *core.QUrl, baseURI *core.QUrl) *core.QUrl) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectResolve", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractUriResolver) DisconnectResolve() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectResolve"})
}

func (ptr *QAbstractUriResolver) Resolve(relative core.QUrl_ITF, baseURI core.QUrl_ITF) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Resolve", relative, baseURI}).(*core.QUrl)
}

func (ptr *QAbstractUriResolver) ConnectDestroyQAbstractUriResolver(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractUriResolver", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractUriResolver) DisconnectDestroyQAbstractUriResolver() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractUriResolver"})
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolver() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractUriResolver"})
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolverDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractUriResolverDefault"})
}

func (ptr *QAbstractUriResolver) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QAbstractUriResolver) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QAbstractUriResolver) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractUriResolver) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QAbstractUriResolver) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QAbstractUriResolver) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractUriResolver) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QAbstractUriResolver) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QAbstractUriResolver) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractUriResolver) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QAbstractUriResolver) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QAbstractUriResolver) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QAbstractUriResolver) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QAbstractUriResolver) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QAbstractUriResolver) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QAbstractUriResolver) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QAbstractUriResolver) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QAbstractUriResolver) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QAbstractUriResolver) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QAbstractUriResolver) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QAbstractUriResolver) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QAbstractXmlNodeModel struct {
	core.QSharedData
}

type QAbstractXmlNodeModel_ITF interface {
	core.QSharedData_ITF
	QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel
}

func (ptr *QAbstractXmlNodeModel) QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel {
	return ptr
}

func (ptr *QAbstractXmlNodeModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedData_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSharedData_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractXmlNodeModel(ptr QAbstractXmlNodeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractXmlNodeModel) InitFromInternal(ptr uintptr, name string) {
	n.QSharedData_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAbstractXmlNodeModel) ClassNameInternalF() string {
	return n.QSharedData_PTR().ClassNameInternalF()
}

func NewQAbstractXmlNodeModelFromPointer(ptr unsafe.Pointer) (n *QAbstractXmlNodeModel) {
	n = new(QAbstractXmlNodeModel)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QAbstractXmlNodeModel")
	return
}

//go:generate stringer -type=QAbstractXmlNodeModel__SimpleAxis
//QAbstractXmlNodeModel::SimpleAxis
type QAbstractXmlNodeModel__SimpleAxis int64

const (
	QAbstractXmlNodeModel__Parent          QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(0)
	QAbstractXmlNodeModel__FirstChild      QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(1)
	QAbstractXmlNodeModel__PreviousSibling QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(2)
	QAbstractXmlNodeModel__NextSibling     QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(3)
)

func (ptr *QAbstractXmlNodeModel) ConnectBaseUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBaseUri", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectBaseUri() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBaseUri"})
}

func (ptr *QAbstractXmlNodeModel) BaseUri(n QXmlNodeModelIndex_ITF) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseUri", n}).(*core.QUrl)
}

func (ptr *QAbstractXmlNodeModel) ConnectCompareOrder(f func(ni1 *QXmlNodeModelIndex, ni2 *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCompareOrder", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectCompareOrder() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCompareOrder"})
}

func (ptr *QAbstractXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {

	return QXmlNodeModelIndex__DocumentOrder(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CompareOrder", ni1, ni2}).(float64))
}

func (ptr *QAbstractXmlNodeModel) CreateIndex(data int64) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateIndex", data}).(*QXmlNodeModelIndex)
}

func (ptr *QAbstractXmlNodeModel) CreateIndex2(pointer unsafe.Pointer, additionalData int64) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateIndex2", pointer, additionalData}).(*QXmlNodeModelIndex)
}

func (ptr *QAbstractXmlNodeModel) CreateIndex3(data int64, additionalData int64) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateIndex3", data, additionalData}).(*QXmlNodeModelIndex)
}

func (ptr *QAbstractXmlNodeModel) ConnectDocumentUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDocumentUri", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectDocumentUri() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDocumentUri"})
}

func (ptr *QAbstractXmlNodeModel) DocumentUri(n QXmlNodeModelIndex_ITF) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DocumentUri", n}).(*core.QUrl)
}

func (ptr *QAbstractXmlNodeModel) ConnectElementById(f func(id *QXmlName) *QXmlNodeModelIndex) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectElementById", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectElementById() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectElementById"})
}

func (ptr *QAbstractXmlNodeModel) ElementById(id QXmlName_ITF) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ElementById", id}).(*QXmlNodeModelIndex)
}

func (ptr *QAbstractXmlNodeModel) ConnectKind(f func(ni *QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectKind", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectKind() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectKind"})
}

func (ptr *QAbstractXmlNodeModel) Kind(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {

	return QXmlNodeModelIndex__NodeKind(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Kind", ni}).(float64))
}

func (ptr *QAbstractXmlNodeModel) ConnectName(f func(ni *QXmlNodeModelIndex) *QXmlName) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectName"})
}

func (ptr *QAbstractXmlNodeModel) Name(ni QXmlNodeModelIndex_ITF) *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name", ni}).(*QXmlName)
}

func (ptr *QAbstractXmlNodeModel) ConnectNamespaceBindings(f func(n *QXmlNodeModelIndex) []*QXmlName) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNamespaceBindings", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectNamespaceBindings() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNamespaceBindings"})
}

func (ptr *QAbstractXmlNodeModel) NamespaceBindings(n QXmlNodeModelIndex_ITF) []*QXmlName {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceBindings", n}).([]*QXmlName)
}

func (ptr *QAbstractXmlNodeModel) ConnectNextFromSimpleAxis(f func(axis QAbstractXmlNodeModel__SimpleAxis, origin *QXmlNodeModelIndex) *QXmlNodeModelIndex) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNextFromSimpleAxis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectNextFromSimpleAxis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNextFromSimpleAxis"})
}

func (ptr *QAbstractXmlNodeModel) NextFromSimpleAxis(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextFromSimpleAxis", axis, origin}).(*QXmlNodeModelIndex)
}

func (ptr *QAbstractXmlNodeModel) ConnectNodesByIdref(f func(idref *QXmlName) []*QXmlNodeModelIndex) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNodesByIdref", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectNodesByIdref() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNodesByIdref"})
}

func (ptr *QAbstractXmlNodeModel) NodesByIdref(idref QXmlName_ITF) []*QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NodesByIdref", idref}).([]*QXmlNodeModelIndex)
}

func (ptr *QAbstractXmlNodeModel) ConnectRoot(f func(n *QXmlNodeModelIndex) *QXmlNodeModelIndex) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRoot", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectRoot() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRoot"})
}

func (ptr *QAbstractXmlNodeModel) Root(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Root", n}).(*QXmlNodeModelIndex)
}

func (ptr *QAbstractXmlNodeModel) SourceLocation(index QXmlNodeModelIndex_ITF) *QSourceLocation {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceLocation", index}).(*QSourceLocation)
}

func (ptr *QAbstractXmlNodeModel) ConnectStringValue(f func(n *QXmlNodeModelIndex) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStringValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectStringValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStringValue"})
}

func (ptr *QAbstractXmlNodeModel) StringValue(n QXmlNodeModelIndex_ITF) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StringValue", n}).(string)
}

func (ptr *QAbstractXmlNodeModel) ConnectTypedValue(f func(node *QXmlNodeModelIndex) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTypedValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectTypedValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTypedValue"})
}

func (ptr *QAbstractXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypedValue", node}).(*core.QVariant)
}

func (ptr *QAbstractXmlNodeModel) ConnectDestroyQAbstractXmlNodeModel(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractXmlNodeModel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlNodeModel) DisconnectDestroyQAbstractXmlNodeModel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractXmlNodeModel"})
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractXmlNodeModel"})
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModelDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractXmlNodeModelDefault"})
}

func (ptr *QAbstractXmlNodeModel) __namespaceBindings_atList(i int) *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceBindings_atList", i}).(*QXmlName)
}

func (ptr *QAbstractXmlNodeModel) __namespaceBindings_setList(i QXmlName_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceBindings_setList", i})
}

func (ptr *QAbstractXmlNodeModel) __namespaceBindings_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceBindings_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractXmlNodeModel) __nodesByIdref_atList(i int) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__nodesByIdref_atList", i}).(*QXmlNodeModelIndex)
}

func (ptr *QAbstractXmlNodeModel) __nodesByIdref_setList(i QXmlNodeModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__nodesByIdref_setList", i})
}

func (ptr *QAbstractXmlNodeModel) __nodesByIdref_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__nodesByIdref_newList"}).(unsafe.Pointer)
}

type QAbstractXmlReceiver struct {
	internal.Internal
}

type QAbstractXmlReceiver_ITF interface {
	QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver
}

func (ptr *QAbstractXmlReceiver) QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver {
	return ptr
}

func (ptr *QAbstractXmlReceiver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAbstractXmlReceiver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAbstractXmlReceiver(ptr QAbstractXmlReceiver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlReceiver_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractXmlReceiver) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAbstractXmlReceiverFromPointer(ptr unsafe.Pointer) (n *QAbstractXmlReceiver) {
	n = new(QAbstractXmlReceiver)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QAbstractXmlReceiver")
	return
}
func NewQAbstractXmlReceiver() *QAbstractXmlReceiver {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQAbstractXmlReceiver", ""}).(*QAbstractXmlReceiver)
}

func (ptr *QAbstractXmlReceiver) ConnectAtomicValue(f func(value *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAtomicValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectAtomicValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAtomicValue"})
}

func (ptr *QAbstractXmlReceiver) AtomicValue(value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AtomicValue", value})
}

func (ptr *QAbstractXmlReceiver) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAttribute", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectAttribute() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAttribute"})
}

func (ptr *QAbstractXmlReceiver) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Attribute", name, value})
}

func (ptr *QAbstractXmlReceiver) ConnectCharacters(f func(value *core.QStringRef)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCharacters", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectCharacters() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCharacters"})
}

func (ptr *QAbstractXmlReceiver) Characters(value core.QStringRef_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Characters", value})
}

func (ptr *QAbstractXmlReceiver) ConnectComment(f func(value string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectComment", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectComment() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectComment"})
}

func (ptr *QAbstractXmlReceiver) Comment(value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Comment", value})
}

func (ptr *QAbstractXmlReceiver) ConnectEndDocument(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndDocument", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectEndDocument() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndDocument"})
}

func (ptr *QAbstractXmlReceiver) EndDocument() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndDocument"})
}

func (ptr *QAbstractXmlReceiver) ConnectEndElement(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndElement", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectEndElement() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndElement"})
}

func (ptr *QAbstractXmlReceiver) EndElement() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndElement"})
}

func (ptr *QAbstractXmlReceiver) ConnectEndOfSequence(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndOfSequence", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectEndOfSequence() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndOfSequence"})
}

func (ptr *QAbstractXmlReceiver) EndOfSequence() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndOfSequence"})
}

func (ptr *QAbstractXmlReceiver) ConnectNamespaceBinding(f func(name *QXmlName)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNamespaceBinding", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectNamespaceBinding() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNamespaceBinding"})
}

func (ptr *QAbstractXmlReceiver) NamespaceBinding(name QXmlName_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceBinding", name})
}

func (ptr *QAbstractXmlReceiver) ConnectProcessingInstruction(f func(target *QXmlName, value string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProcessingInstruction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectProcessingInstruction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProcessingInstruction"})
}

func (ptr *QAbstractXmlReceiver) ProcessingInstruction(target QXmlName_ITF, value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessingInstruction", target, value})
}

func (ptr *QAbstractXmlReceiver) ConnectStartDocument(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartDocument", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectStartDocument() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartDocument"})
}

func (ptr *QAbstractXmlReceiver) StartDocument() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDocument"})
}

func (ptr *QAbstractXmlReceiver) ConnectStartElement(f func(name *QXmlName)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartElement", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectStartElement() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartElement"})
}

func (ptr *QAbstractXmlReceiver) StartElement(name QXmlName_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartElement", name})
}

func (ptr *QAbstractXmlReceiver) ConnectStartOfSequence(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartOfSequence", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectStartOfSequence() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartOfSequence"})
}

func (ptr *QAbstractXmlReceiver) StartOfSequence() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartOfSequence"})
}

func (ptr *QAbstractXmlReceiver) ConnectDestroyQAbstractXmlReceiver(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractXmlReceiver", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractXmlReceiver) DisconnectDestroyQAbstractXmlReceiver() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractXmlReceiver"})
}

func (ptr *QAbstractXmlReceiver) DestroyQAbstractXmlReceiver() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractXmlReceiver"})
}

func (ptr *QAbstractXmlReceiver) DestroyQAbstractXmlReceiverDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractXmlReceiverDefault"})
}

type QSimpleXmlNodeModel struct {
	QAbstractXmlNodeModel
}

type QSimpleXmlNodeModel_ITF interface {
	QAbstractXmlNodeModel_ITF
	QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel
}

func (ptr *QSimpleXmlNodeModel) QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel {
	return ptr
}

func (ptr *QSimpleXmlNodeModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractXmlNodeModel_PTR().SetPointer(p)
	}
}

func PointerFromQSimpleXmlNodeModel(ptr QSimpleXmlNodeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSimpleXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func (n *QSimpleXmlNodeModel) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractXmlNodeModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSimpleXmlNodeModel) ClassNameInternalF() string {
	return n.QAbstractXmlNodeModel_PTR().ClassNameInternalF()
}

func NewQSimpleXmlNodeModelFromPointer(ptr unsafe.Pointer) (n *QSimpleXmlNodeModel) {
	n = new(QSimpleXmlNodeModel)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QSimpleXmlNodeModel")
	return
}
func (ptr *QSimpleXmlNodeModel) ConnectBaseUri(f func(node *QXmlNodeModelIndex) *core.QUrl) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBaseUri", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSimpleXmlNodeModel) DisconnectBaseUri() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBaseUri"})
}

func (ptr *QSimpleXmlNodeModel) BaseUri(node QXmlNodeModelIndex_ITF) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseUri", node}).(*core.QUrl)
}

func (ptr *QSimpleXmlNodeModel) BaseUriDefault(node QXmlNodeModelIndex_ITF) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaseUriDefault", node}).(*core.QUrl)
}

func (ptr *QSimpleXmlNodeModel) ConnectElementById(f func(id *QXmlName) *QXmlNodeModelIndex) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectElementById", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSimpleXmlNodeModel) DisconnectElementById() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectElementById"})
}

func (ptr *QSimpleXmlNodeModel) ElementById(id QXmlName_ITF) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ElementById", id}).(*QXmlNodeModelIndex)
}

func (ptr *QSimpleXmlNodeModel) ElementByIdDefault(id QXmlName_ITF) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ElementByIdDefault", id}).(*QXmlNodeModelIndex)
}

func (ptr *QSimpleXmlNodeModel) NamePool() *QXmlNamePool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamePool"}).(*QXmlNamePool)
}

func (ptr *QSimpleXmlNodeModel) ConnectNamespaceBindings(f func(node *QXmlNodeModelIndex) []*QXmlName) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNamespaceBindings", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSimpleXmlNodeModel) DisconnectNamespaceBindings() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNamespaceBindings"})
}

func (ptr *QSimpleXmlNodeModel) NamespaceBindings(node QXmlNodeModelIndex_ITF) []*QXmlName {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceBindings", node}).([]*QXmlName)
}

func (ptr *QSimpleXmlNodeModel) NamespaceBindingsDefault(node QXmlNodeModelIndex_ITF) []*QXmlName {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceBindingsDefault", node}).([]*QXmlName)
}

func (ptr *QSimpleXmlNodeModel) ConnectNodesByIdref(f func(idref *QXmlName) []*QXmlNodeModelIndex) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNodesByIdref", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSimpleXmlNodeModel) DisconnectNodesByIdref() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNodesByIdref"})
}

func (ptr *QSimpleXmlNodeModel) NodesByIdref(idref QXmlName_ITF) []*QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NodesByIdref", idref}).([]*QXmlNodeModelIndex)
}

func (ptr *QSimpleXmlNodeModel) NodesByIdrefDefault(idref QXmlName_ITF) []*QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NodesByIdrefDefault", idref}).([]*QXmlNodeModelIndex)
}

func (ptr *QSimpleXmlNodeModel) ConnectStringValue(f func(node *QXmlNodeModelIndex) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStringValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSimpleXmlNodeModel) DisconnectStringValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStringValue"})
}

func (ptr *QSimpleXmlNodeModel) StringValue(node QXmlNodeModelIndex_ITF) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StringValue", node}).(string)
}

func (ptr *QSimpleXmlNodeModel) StringValueDefault(node QXmlNodeModelIndex_ITF) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StringValueDefault", node}).(string)
}

func (ptr *QSimpleXmlNodeModel) ConnectDestroyQSimpleXmlNodeModel(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSimpleXmlNodeModel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSimpleXmlNodeModel) DisconnectDestroyQSimpleXmlNodeModel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSimpleXmlNodeModel"})
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSimpleXmlNodeModel"})
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModelDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSimpleXmlNodeModelDefault"})
}

func (ptr *QSimpleXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {

	return QXmlNodeModelIndex__DocumentOrder(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CompareOrder", ni1, ni2}).(float64))
}

func (ptr *QSimpleXmlNodeModel) CompareOrderDefault(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {

	return QXmlNodeModelIndex__DocumentOrder(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CompareOrderDefault", ni1, ni2}).(float64))
}

func (ptr *QSimpleXmlNodeModel) DocumentUri(n QXmlNodeModelIndex_ITF) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DocumentUri", n}).(*core.QUrl)
}

func (ptr *QSimpleXmlNodeModel) DocumentUriDefault(n QXmlNodeModelIndex_ITF) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DocumentUriDefault", n}).(*core.QUrl)
}

func (ptr *QSimpleXmlNodeModel) Kind(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {

	return QXmlNodeModelIndex__NodeKind(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Kind", ni}).(float64))
}

func (ptr *QSimpleXmlNodeModel) KindDefault(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {

	return QXmlNodeModelIndex__NodeKind(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KindDefault", ni}).(float64))
}

func (ptr *QSimpleXmlNodeModel) Name(ni QXmlNodeModelIndex_ITF) *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name", ni}).(*QXmlName)
}

func (ptr *QSimpleXmlNodeModel) NameDefault(ni QXmlNodeModelIndex_ITF) *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NameDefault", ni}).(*QXmlName)
}

func (ptr *QSimpleXmlNodeModel) NextFromSimpleAxis(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextFromSimpleAxis", axis, origin}).(*QXmlNodeModelIndex)
}

func (ptr *QSimpleXmlNodeModel) NextFromSimpleAxisDefault(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextFromSimpleAxisDefault", axis, origin}).(*QXmlNodeModelIndex)
}

func (ptr *QSimpleXmlNodeModel) Root(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Root", n}).(*QXmlNodeModelIndex)
}

func (ptr *QSimpleXmlNodeModel) RootDefault(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RootDefault", n}).(*QXmlNodeModelIndex)
}

func (ptr *QSimpleXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypedValue", node}).(*core.QVariant)
}

func (ptr *QSimpleXmlNodeModel) TypedValueDefault(node QXmlNodeModelIndex_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypedValueDefault", node}).(*core.QVariant)
}

type QSourceLocation struct {
	internal.Internal
}

type QSourceLocation_ITF interface {
	QSourceLocation_PTR() *QSourceLocation
}

func (ptr *QSourceLocation) QSourceLocation_PTR() *QSourceLocation {
	return ptr
}

func (ptr *QSourceLocation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSourceLocation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSourceLocation(ptr QSourceLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSourceLocation_PTR().Pointer()
	}
	return nil
}

func (n *QSourceLocation) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSourceLocationFromPointer(ptr unsafe.Pointer) (n *QSourceLocation) {
	n = new(QSourceLocation)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QSourceLocation")
	return
}
func NewQSourceLocation() *QSourceLocation {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQSourceLocation", ""}).(*QSourceLocation)
}

func NewQSourceLocation2(other QSourceLocation_ITF) *QSourceLocation {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQSourceLocation2", "", other}).(*QSourceLocation)
}

func NewQSourceLocation3(u core.QUrl_ITF, l int, c int) *QSourceLocation {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQSourceLocation3", "", u, l, c}).(*QSourceLocation)
}

func (ptr *QSourceLocation) Column() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Column"}).(float64))
}

func (ptr *QSourceLocation) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QSourceLocation) Line() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Line"}).(float64))
}

func (ptr *QSourceLocation) SetColumn(newColumn int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColumn", newColumn})
}

func (ptr *QSourceLocation) SetLine(newLine int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLine", newLine})
}

func (ptr *QSourceLocation) SetUri(newUri core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUri", newUri})
}

func (ptr *QSourceLocation) Uri() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uri"}).(*core.QUrl)
}

func (ptr *QSourceLocation) DestroyQSourceLocation() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSourceLocation"})
}

type QXmlFormatter struct {
	QXmlSerializer
}

type QXmlFormatter_ITF interface {
	QXmlSerializer_ITF
	QXmlFormatter_PTR() *QXmlFormatter
}

func (ptr *QXmlFormatter) QXmlFormatter_PTR() *QXmlFormatter {
	return ptr
}

func (ptr *QXmlFormatter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSerializer_PTR().Pointer()
	}
	return nil
}

func (ptr *QXmlFormatter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QXmlSerializer_PTR().SetPointer(p)
	}
}

func PointerFromQXmlFormatter(ptr QXmlFormatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlFormatter_PTR().Pointer()
	}
	return nil
}

func (n *QXmlFormatter) InitFromInternal(ptr uintptr, name string) {
	n.QXmlSerializer_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QXmlFormatter) ClassNameInternalF() string {
	return n.QXmlSerializer_PTR().ClassNameInternalF()
}

func NewQXmlFormatterFromPointer(ptr unsafe.Pointer) (n *QXmlFormatter) {
	n = new(QXmlFormatter)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QXmlFormatter")
	return
}

func (ptr *QXmlFormatter) DestroyQXmlFormatter() {
}

func NewQXmlFormatter(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlFormatter {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlFormatter", "", query, outputDevice}).(*QXmlFormatter)
}

func (ptr *QXmlFormatter) IndentationDepth() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndentationDepth"}).(float64))
}

func (ptr *QXmlFormatter) SetIndentationDepth(depth int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIndentationDepth", depth})
}

type QXmlItem struct {
	internal.Internal
}

type QXmlItem_ITF interface {
	QXmlItem_PTR() *QXmlItem
}

func (ptr *QXmlItem) QXmlItem_PTR() *QXmlItem {
	return ptr
}

func (ptr *QXmlItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlItem(ptr QXmlItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlItem_PTR().Pointer()
	}
	return nil
}

func (n *QXmlItem) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlItemFromPointer(ptr unsafe.Pointer) (n *QXmlItem) {
	n = new(QXmlItem)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QXmlItem")
	return
}
func NewQXmlItem() *QXmlItem {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlItem", ""}).(*QXmlItem)
}

func NewQXmlItem2(other QXmlItem_ITF) *QXmlItem {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlItem2", "", other}).(*QXmlItem)
}

func NewQXmlItem3(node QXmlNodeModelIndex_ITF) *QXmlItem {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlItem3", "", node}).(*QXmlItem)
}

func NewQXmlItem4(atomicValue core.QVariant_ITF) *QXmlItem {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlItem4", "", atomicValue}).(*QXmlItem)
}

func (ptr *QXmlItem) IsAtomicValue() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAtomicValue"}).(bool)
}

func (ptr *QXmlItem) IsNode() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNode"}).(bool)
}

func (ptr *QXmlItem) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QXmlItem) ToAtomicValue() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToAtomicValue"}).(*core.QVariant)
}

func (ptr *QXmlItem) ToNodeModelIndex() *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToNodeModelIndex"}).(*QXmlNodeModelIndex)
}

func (ptr *QXmlItem) DestroyQXmlItem() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlItem"})
}

type QXmlName struct {
	internal.Internal
}

type QXmlName_ITF interface {
	QXmlName_PTR() *QXmlName
}

func (ptr *QXmlName) QXmlName_PTR() *QXmlName {
	return ptr
}

func (ptr *QXmlName) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlName) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlName(ptr QXmlName_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlName_PTR().Pointer()
	}
	return nil
}

func (n *QXmlName) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlNameFromPointer(ptr unsafe.Pointer) (n *QXmlName) {
	n = new(QXmlName)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QXmlName")
	return
}

func (ptr *QXmlName) DestroyQXmlName() {
}

func NewQXmlName() *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlName", ""}).(*QXmlName)
}

func NewQXmlName2(namePool QXmlNamePool_ITF, localName string, namespaceURI string, prefix string) *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlName2", "", namePool, localName, namespaceURI, prefix}).(*QXmlName)
}

func NewQXmlName3(other QXmlName_ITF) *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlName3", "", other}).(*QXmlName)
}

func QXmlName_FromClarkName(clarkName string, namePool QXmlNamePool_ITF) *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.QXmlName_FromClarkName", "", clarkName, namePool}).(*QXmlName)
}

func (ptr *QXmlName) FromClarkName(clarkName string, namePool QXmlNamePool_ITF) *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.QXmlName_FromClarkName", "", clarkName, namePool}).(*QXmlName)
}

func QXmlName_IsNCName(candidate string) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.QXmlName_IsNCName", "", candidate}).(bool)
}

func (ptr *QXmlName) IsNCName(candidate string) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.QXmlName_IsNCName", "", candidate}).(bool)
}

func (ptr *QXmlName) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QXmlName) LocalName(namePool QXmlNamePool_ITF) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalName", namePool}).(string)
}

func (ptr *QXmlName) NamespaceUri(namePool QXmlNamePool_ITF) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceUri", namePool}).(string)
}

func (ptr *QXmlName) Prefix(namePool QXmlNamePool_ITF) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prefix", namePool}).(string)
}

func (ptr *QXmlName) ToClarkName(namePool QXmlNamePool_ITF) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToClarkName", namePool}).(string)
}

type QXmlNamePool struct {
	internal.Internal
}

type QXmlNamePool_ITF interface {
	QXmlNamePool_PTR() *QXmlNamePool
}

func (ptr *QXmlNamePool) QXmlNamePool_PTR() *QXmlNamePool {
	return ptr
}

func (ptr *QXmlNamePool) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlNamePool) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlNamePool(ptr QXmlNamePool_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamePool_PTR().Pointer()
	}
	return nil
}

func (n *QXmlNamePool) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlNamePoolFromPointer(ptr unsafe.Pointer) (n *QXmlNamePool) {
	n = new(QXmlNamePool)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QXmlNamePool")
	return
}
func NewQXmlNamePool() *QXmlNamePool {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlNamePool", ""}).(*QXmlNamePool)
}

func NewQXmlNamePool2(other QXmlNamePool_ITF) *QXmlNamePool {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlNamePool2", "", other}).(*QXmlNamePool)
}

func (ptr *QXmlNamePool) DestroyQXmlNamePool() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlNamePool"})
}

type QXmlNodeModelIndex struct {
	internal.Internal
}

type QXmlNodeModelIndex_ITF interface {
	QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex
}

func (ptr *QXmlNodeModelIndex) QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex {
	return ptr
}

func (ptr *QXmlNodeModelIndex) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlNodeModelIndex(ptr QXmlNodeModelIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNodeModelIndex_PTR().Pointer()
	}
	return nil
}

func (n *QXmlNodeModelIndex) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlNodeModelIndexFromPointer(ptr unsafe.Pointer) (n *QXmlNodeModelIndex) {
	n = new(QXmlNodeModelIndex)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QXmlNodeModelIndex")
	return
}

func (ptr *QXmlNodeModelIndex) DestroyQXmlNodeModelIndex() {
}

//go:generate stringer -type=QXmlNodeModelIndex__NodeKind
//QXmlNodeModelIndex::NodeKind
type QXmlNodeModelIndex__NodeKind int64

const (
	QXmlNodeModelIndex__Attribute             QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(1)
	QXmlNodeModelIndex__Comment               QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(2)
	QXmlNodeModelIndex__Document              QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(4)
	QXmlNodeModelIndex__Element               QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(8)
	QXmlNodeModelIndex__Namespace             QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(16)
	QXmlNodeModelIndex__ProcessingInstruction QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(32)
	QXmlNodeModelIndex__Text                  QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(64)
)

//go:generate stringer -type=QXmlNodeModelIndex__DocumentOrder
//QXmlNodeModelIndex::DocumentOrder
type QXmlNodeModelIndex__DocumentOrder int64

const (
	QXmlNodeModelIndex__Precedes QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(-1)
	QXmlNodeModelIndex__Is       QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(0)
	QXmlNodeModelIndex__Follows  QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(1)
)

func NewQXmlNodeModelIndex() *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlNodeModelIndex", ""}).(*QXmlNodeModelIndex)
}

func NewQXmlNodeModelIndex2(other QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlNodeModelIndex2", "", other}).(*QXmlNodeModelIndex)
}

func (ptr *QXmlNodeModelIndex) AdditionalData() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AdditionalData"}).(float64))
}

func (ptr *QXmlNodeModelIndex) Data() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data"}).(float64))
}

func (ptr *QXmlNodeModelIndex) InternalPointer() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InternalPointer"}).(unsafe.Pointer)
}

func (ptr *QXmlNodeModelIndex) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QXmlNodeModelIndex) Model() *QAbstractXmlNodeModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*QAbstractXmlNodeModel)
}

func (ptr *QXmlNodeModelIndex) __namespaceBindings_atList(i int) *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceBindings_atList", i}).(*QXmlName)
}

func (ptr *QXmlNodeModelIndex) __namespaceBindings_setList(i QXmlName_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceBindings_setList", i})
}

func (ptr *QXmlNodeModelIndex) __namespaceBindings_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__namespaceBindings_newList"}).(unsafe.Pointer)
}

type QXmlQuery struct {
	internal.Internal
}

type QXmlQuery_ITF interface {
	QXmlQuery_PTR() *QXmlQuery
}

func (ptr *QXmlQuery) QXmlQuery_PTR() *QXmlQuery {
	return ptr
}

func (ptr *QXmlQuery) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlQuery) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlQuery(ptr QXmlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlQuery_PTR().Pointer()
	}
	return nil
}

func (n *QXmlQuery) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlQueryFromPointer(ptr unsafe.Pointer) (n *QXmlQuery) {
	n = new(QXmlQuery)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QXmlQuery")
	return
}

//go:generate stringer -type=QXmlQuery__QueryLanguage
//QXmlQuery::QueryLanguage
type QXmlQuery__QueryLanguage int64

const (
	QXmlQuery__XQuery10                              QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(1)
	QXmlQuery__XSLT20                                QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(2)
	QXmlQuery__XmlSchema11IdentityConstraintSelector QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(1024)
	QXmlQuery__XmlSchema11IdentityConstraintField    QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(2048)
	QXmlQuery__XPath20                               QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(4096)
)

func NewQXmlQuery() *QXmlQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlQuery", ""}).(*QXmlQuery)
}

func NewQXmlQuery2(other QXmlQuery_ITF) *QXmlQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlQuery2", "", other}).(*QXmlQuery)
}

func NewQXmlQuery3(np QXmlNamePool_ITF) *QXmlQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlQuery3", "", np}).(*QXmlQuery)
}

func NewQXmlQuery4(queryLanguage QXmlQuery__QueryLanguage, np QXmlNamePool_ITF) *QXmlQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlQuery4", "", queryLanguage, np}).(*QXmlQuery)
}

func (ptr *QXmlQuery) BindVariable(name QXmlName_ITF, value QXmlItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindVariable", name, value})
}

func (ptr *QXmlQuery) BindVariable2(localName string, value QXmlItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindVariable2", localName, value})
}

func (ptr *QXmlQuery) BindVariable3(name QXmlName_ITF, device core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindVariable3", name, device})
}

func (ptr *QXmlQuery) BindVariable4(localName string, device core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindVariable4", localName, device})
}

func (ptr *QXmlQuery) BindVariable5(name QXmlName_ITF, query QXmlQuery_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindVariable5", name, query})
}

func (ptr *QXmlQuery) BindVariable6(localName string, query QXmlQuery_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindVariable6", localName, query})
}

func (ptr *QXmlQuery) EvaluateTo(result QXmlResultItems_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EvaluateTo", result})
}

func (ptr *QXmlQuery) EvaluateTo2(callback QAbstractXmlReceiver_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EvaluateTo2", callback}).(bool)
}

func (ptr *QXmlQuery) EvaluateTo3(target []string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EvaluateTo3", target}).(bool)
}

func (ptr *QXmlQuery) EvaluateTo4(target core.QIODevice_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EvaluateTo4", target}).(bool)
}

func (ptr *QXmlQuery) EvaluateTo5(output string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EvaluateTo5", output}).(bool)
}

func (ptr *QXmlQuery) InitialTemplateName() *QXmlName {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitialTemplateName"}).(*QXmlName)
}

func (ptr *QXmlQuery) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QXmlQuery) MessageHandler() *QAbstractMessageHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MessageHandler"}).(*QAbstractMessageHandler)
}

func (ptr *QXmlQuery) NamePool() *QXmlNamePool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamePool"}).(*QXmlNamePool)
}

func (ptr *QXmlQuery) NetworkAccessManager() *network.QNetworkAccessManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NetworkAccessManager"}).(*network.QNetworkAccessManager)
}

func (ptr *QXmlQuery) QueryLanguage() QXmlQuery__QueryLanguage {

	return QXmlQuery__QueryLanguage(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QueryLanguage"}).(float64))
}

func (ptr *QXmlQuery) SetFocus(item QXmlItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus", item})
}

func (ptr *QXmlQuery) SetFocus2(documentURI core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2", documentURI}).(bool)
}

func (ptr *QXmlQuery) SetFocus3(document core.QIODevice_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus3", document}).(bool)
}

func (ptr *QXmlQuery) SetFocus4(focus string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus4", focus}).(bool)
}

func (ptr *QXmlQuery) SetInitialTemplateName(name QXmlName_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInitialTemplateName", name})
}

func (ptr *QXmlQuery) SetInitialTemplateName2(localName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInitialTemplateName2", localName})
}

func (ptr *QXmlQuery) SetMessageHandler(aMessageHandler QAbstractMessageHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMessageHandler", aMessageHandler})
}

func (ptr *QXmlQuery) SetNetworkAccessManager(newManager network.QNetworkAccessManager_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNetworkAccessManager", newManager})
}

func (ptr *QXmlQuery) SetQuery(sourceCode core.QIODevice_ITF, documentURI core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetQuery", sourceCode, documentURI})
}

func (ptr *QXmlQuery) SetQuery2(sourceCode string, documentURI core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetQuery2", sourceCode, documentURI})
}

func (ptr *QXmlQuery) SetQuery3(queryURI core.QUrl_ITF, baseURI core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetQuery3", queryURI, baseURI})
}

func (ptr *QXmlQuery) SetUriResolver(resolver QAbstractUriResolver_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUriResolver", resolver})
}

func (ptr *QXmlQuery) UriResolver() *QAbstractUriResolver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UriResolver"}).(*QAbstractUriResolver)
}

func (ptr *QXmlQuery) DestroyQXmlQuery() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlQuery"})
}

type QXmlResultItems struct {
	internal.Internal
}

type QXmlResultItems_ITF interface {
	QXmlResultItems_PTR() *QXmlResultItems
}

func (ptr *QXmlResultItems) QXmlResultItems_PTR() *QXmlResultItems {
	return ptr
}

func (ptr *QXmlResultItems) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlResultItems) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlResultItems(ptr QXmlResultItems_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlResultItems_PTR().Pointer()
	}
	return nil
}

func (n *QXmlResultItems) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlResultItemsFromPointer(ptr unsafe.Pointer) (n *QXmlResultItems) {
	n = new(QXmlResultItems)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QXmlResultItems")
	return
}
func NewQXmlResultItems() *QXmlResultItems {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlResultItems", ""}).(*QXmlResultItems)
}

func (ptr *QXmlResultItems) Current() *QXmlItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Current"}).(*QXmlItem)
}

func (ptr *QXmlResultItems) HasError() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasError"}).(bool)
}

func (ptr *QXmlResultItems) Next() *QXmlItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Next"}).(*QXmlItem)
}

func (ptr *QXmlResultItems) ConnectDestroyQXmlResultItems(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlResultItems", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlResultItems) DisconnectDestroyQXmlResultItems() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlResultItems"})
}

func (ptr *QXmlResultItems) DestroyQXmlResultItems() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlResultItems"})
}

func (ptr *QXmlResultItems) DestroyQXmlResultItemsDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlResultItemsDefault"})
}

type QXmlSchema struct {
	internal.Internal
}

type QXmlSchema_ITF interface {
	QXmlSchema_PTR() *QXmlSchema
}

func (ptr *QXmlSchema) QXmlSchema_PTR() *QXmlSchema {
	return ptr
}

func (ptr *QXmlSchema) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlSchema) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlSchema(ptr QXmlSchema_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchema_PTR().Pointer()
	}
	return nil
}

func (n *QXmlSchema) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlSchemaFromPointer(ptr unsafe.Pointer) (n *QXmlSchema) {
	n = new(QXmlSchema)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QXmlSchema")
	return
}
func NewQXmlSchema() *QXmlSchema {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlSchema", ""}).(*QXmlSchema)
}

func NewQXmlSchema2(other QXmlSchema_ITF) *QXmlSchema {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlSchema2", "", other}).(*QXmlSchema)
}

func (ptr *QXmlSchema) DocumentUri() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DocumentUri"}).(*core.QUrl)
}

func (ptr *QXmlSchema) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QXmlSchema) Load(source core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load", source}).(bool)
}

func (ptr *QXmlSchema) Load2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load2", source, documentUri}).(bool)
}

func (ptr *QXmlSchema) Load3(data core.QByteArray_ITF, documentUri core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load3", data, documentUri}).(bool)
}

func (ptr *QXmlSchema) MessageHandler() *QAbstractMessageHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MessageHandler"}).(*QAbstractMessageHandler)
}

func (ptr *QXmlSchema) NamePool() *QXmlNamePool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamePool"}).(*QXmlNamePool)
}

func (ptr *QXmlSchema) NetworkAccessManager() *network.QNetworkAccessManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NetworkAccessManager"}).(*network.QNetworkAccessManager)
}

func (ptr *QXmlSchema) SetMessageHandler(handler QAbstractMessageHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMessageHandler", handler})
}

func (ptr *QXmlSchema) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNetworkAccessManager", manager})
}

func (ptr *QXmlSchema) SetUriResolver(resolver QAbstractUriResolver_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUriResolver", resolver})
}

func (ptr *QXmlSchema) UriResolver() *QAbstractUriResolver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UriResolver"}).(*QAbstractUriResolver)
}

func (ptr *QXmlSchema) DestroyQXmlSchema() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlSchema"})
}

type QXmlSchemaValidator struct {
	internal.Internal
}

type QXmlSchemaValidator_ITF interface {
	QXmlSchemaValidator_PTR() *QXmlSchemaValidator
}

func (ptr *QXmlSchemaValidator) QXmlSchemaValidator_PTR() *QXmlSchemaValidator {
	return ptr
}

func (ptr *QXmlSchemaValidator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlSchemaValidator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlSchemaValidator(ptr QXmlSchemaValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchemaValidator_PTR().Pointer()
	}
	return nil
}

func (n *QXmlSchemaValidator) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlSchemaValidatorFromPointer(ptr unsafe.Pointer) (n *QXmlSchemaValidator) {
	n = new(QXmlSchemaValidator)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QXmlSchemaValidator")
	return
}
func NewQXmlSchemaValidator() *QXmlSchemaValidator {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlSchemaValidator", ""}).(*QXmlSchemaValidator)
}

func NewQXmlSchemaValidator2(schema QXmlSchema_ITF) *QXmlSchemaValidator {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlSchemaValidator2", "", schema}).(*QXmlSchemaValidator)
}

func (ptr *QXmlSchemaValidator) MessageHandler() *QAbstractMessageHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MessageHandler"}).(*QAbstractMessageHandler)
}

func (ptr *QXmlSchemaValidator) NamePool() *QXmlNamePool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamePool"}).(*QXmlNamePool)
}

func (ptr *QXmlSchemaValidator) NetworkAccessManager() *network.QNetworkAccessManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NetworkAccessManager"}).(*network.QNetworkAccessManager)
}

func (ptr *QXmlSchemaValidator) Schema() *QXmlSchema {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Schema"}).(*QXmlSchema)
}

func (ptr *QXmlSchemaValidator) SetMessageHandler(handler QAbstractMessageHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMessageHandler", handler})
}

func (ptr *QXmlSchemaValidator) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNetworkAccessManager", manager})
}

func (ptr *QXmlSchemaValidator) SetSchema(schema QXmlSchema_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSchema", schema})
}

func (ptr *QXmlSchemaValidator) SetUriResolver(resolver QAbstractUriResolver_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUriResolver", resolver})
}

func (ptr *QXmlSchemaValidator) UriResolver() *QAbstractUriResolver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UriResolver"}).(*QAbstractUriResolver)
}

func (ptr *QXmlSchemaValidator) Validate(source core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Validate", source}).(bool)
}

func (ptr *QXmlSchemaValidator) Validate2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Validate2", source, documentUri}).(bool)
}

func (ptr *QXmlSchemaValidator) Validate3(data core.QByteArray_ITF, documentUri core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Validate3", data, documentUri}).(bool)
}

func (ptr *QXmlSchemaValidator) DestroyQXmlSchemaValidator() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlSchemaValidator"})
}

type QXmlSerializer struct {
	QAbstractXmlReceiver
}

type QXmlSerializer_ITF interface {
	QAbstractXmlReceiver_ITF
	QXmlSerializer_PTR() *QXmlSerializer
}

func (ptr *QXmlSerializer) QXmlSerializer_PTR() *QXmlSerializer {
	return ptr
}

func (ptr *QXmlSerializer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlReceiver_PTR().Pointer()
	}
	return nil
}

func (ptr *QXmlSerializer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractXmlReceiver_PTR().SetPointer(p)
	}
}

func PointerFromQXmlSerializer(ptr QXmlSerializer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSerializer_PTR().Pointer()
	}
	return nil
}

func (n *QXmlSerializer) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractXmlReceiver_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QXmlSerializer) ClassNameInternalF() string {
	return n.QAbstractXmlReceiver_PTR().ClassNameInternalF()
}

func NewQXmlSerializerFromPointer(ptr unsafe.Pointer) (n *QXmlSerializer) {
	n = new(QXmlSerializer)
	n.InitFromInternal(uintptr(ptr), "xmlpatterns.QXmlSerializer")
	return
}

func (ptr *QXmlSerializer) DestroyQXmlSerializer() {
}

func NewQXmlSerializer(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlSerializer {

	return internal.CallLocalFunction([]interface{}{"", "", "xmlpatterns.NewQXmlSerializer", "", query, outputDevice}).(*QXmlSerializer)
}

func (ptr *QXmlSerializer) ConnectAtomicValue(f func(value *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAtomicValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectAtomicValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAtomicValue"})
}

func (ptr *QXmlSerializer) AtomicValue(value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AtomicValue", value})
}

func (ptr *QXmlSerializer) AtomicValueDefault(value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AtomicValueDefault", value})
}

func (ptr *QXmlSerializer) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAttribute", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectAttribute() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAttribute"})
}

func (ptr *QXmlSerializer) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Attribute", name, value})
}

func (ptr *QXmlSerializer) AttributeDefault(name QXmlName_ITF, value core.QStringRef_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AttributeDefault", name, value})
}

func (ptr *QXmlSerializer) ConnectCharacters(f func(value *core.QStringRef)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCharacters", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectCharacters() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCharacters"})
}

func (ptr *QXmlSerializer) Characters(value core.QStringRef_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Characters", value})
}

func (ptr *QXmlSerializer) CharactersDefault(value core.QStringRef_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CharactersDefault", value})
}

func (ptr *QXmlSerializer) Codec() *core.QTextCodec {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Codec"}).(*core.QTextCodec)
}

func (ptr *QXmlSerializer) ConnectComment(f func(value string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectComment", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectComment() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectComment"})
}

func (ptr *QXmlSerializer) Comment(value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Comment", value})
}

func (ptr *QXmlSerializer) CommentDefault(value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CommentDefault", value})
}

func (ptr *QXmlSerializer) ConnectEndDocument(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndDocument", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectEndDocument() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndDocument"})
}

func (ptr *QXmlSerializer) EndDocument() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndDocument"})
}

func (ptr *QXmlSerializer) EndDocumentDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndDocumentDefault"})
}

func (ptr *QXmlSerializer) ConnectEndElement(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndElement", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectEndElement() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndElement"})
}

func (ptr *QXmlSerializer) EndElement() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndElement"})
}

func (ptr *QXmlSerializer) EndElementDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndElementDefault"})
}

func (ptr *QXmlSerializer) ConnectEndOfSequence(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndOfSequence", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectEndOfSequence() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndOfSequence"})
}

func (ptr *QXmlSerializer) EndOfSequence() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndOfSequence"})
}

func (ptr *QXmlSerializer) EndOfSequenceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndOfSequenceDefault"})
}

func (ptr *QXmlSerializer) ConnectNamespaceBinding(f func(nb *QXmlName)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNamespaceBinding", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectNamespaceBinding() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNamespaceBinding"})
}

func (ptr *QXmlSerializer) NamespaceBinding(nb QXmlName_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceBinding", nb})
}

func (ptr *QXmlSerializer) NamespaceBindingDefault(nb QXmlName_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceBindingDefault", nb})
}

func (ptr *QXmlSerializer) OutputDevice() *core.QIODevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OutputDevice"}).(*core.QIODevice)
}

func (ptr *QXmlSerializer) ConnectProcessingInstruction(f func(name *QXmlName, value string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProcessingInstruction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectProcessingInstruction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProcessingInstruction"})
}

func (ptr *QXmlSerializer) ProcessingInstruction(name QXmlName_ITF, value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessingInstruction", name, value})
}

func (ptr *QXmlSerializer) ProcessingInstructionDefault(name QXmlName_ITF, value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessingInstructionDefault", name, value})
}

func (ptr *QXmlSerializer) SetCodec(outputCodec core.QTextCodec_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCodec", outputCodec})
}

func (ptr *QXmlSerializer) ConnectStartDocument(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartDocument", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectStartDocument() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartDocument"})
}

func (ptr *QXmlSerializer) StartDocument() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDocument"})
}

func (ptr *QXmlSerializer) StartDocumentDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDocumentDefault"})
}

func (ptr *QXmlSerializer) ConnectStartElement(f func(name *QXmlName)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartElement", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectStartElement() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartElement"})
}

func (ptr *QXmlSerializer) StartElement(name QXmlName_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartElement", name})
}

func (ptr *QXmlSerializer) StartElementDefault(name QXmlName_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartElementDefault", name})
}

func (ptr *QXmlSerializer) ConnectStartOfSequence(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartOfSequence", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSerializer) DisconnectStartOfSequence() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartOfSequence"})
}

func (ptr *QXmlSerializer) StartOfSequence() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartOfSequence"})
}

func (ptr *QXmlSerializer) StartOfSequenceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartOfSequenceDefault"})
}

func init() {
	internal.ConstructorTable["xmlpatterns.QAbstractMessageHandler"] = NewQAbstractMessageHandlerFromPointer
	internal.ConstructorTable["xmlpatterns.QAbstractUriResolver"] = NewQAbstractUriResolverFromPointer
	internal.ConstructorTable["xmlpatterns.QAbstractXmlNodeModel"] = NewQAbstractXmlNodeModelFromPointer
	internal.ConstructorTable["xmlpatterns.QAbstractXmlReceiver"] = NewQAbstractXmlReceiverFromPointer
	internal.ConstructorTable["xmlpatterns.QSimpleXmlNodeModel"] = NewQSimpleXmlNodeModelFromPointer
	internal.ConstructorTable["xmlpatterns.QSourceLocation"] = NewQSourceLocationFromPointer
	internal.ConstructorTable["xmlpatterns.QXmlFormatter"] = NewQXmlFormatterFromPointer
	internal.ConstructorTable["xmlpatterns.QXmlItem"] = NewQXmlItemFromPointer
	internal.ConstructorTable["xmlpatterns.QXmlName"] = NewQXmlNameFromPointer
	internal.ConstructorTable["xmlpatterns.QXmlNamePool"] = NewQXmlNamePoolFromPointer
	internal.ConstructorTable["xmlpatterns.QXmlNodeModelIndex"] = NewQXmlNodeModelIndexFromPointer
	internal.ConstructorTable["xmlpatterns.QXmlQuery"] = NewQXmlQueryFromPointer
	internal.ConstructorTable["xmlpatterns.QXmlResultItems"] = NewQXmlResultItemsFromPointer
	internal.ConstructorTable["xmlpatterns.QXmlSchema"] = NewQXmlSchemaFromPointer
	internal.ConstructorTable["xmlpatterns.QXmlSchemaValidator"] = NewQXmlSchemaValidatorFromPointer
	internal.ConstructorTable["xmlpatterns.QXmlSerializer"] = NewQXmlSerializerFromPointer
}
