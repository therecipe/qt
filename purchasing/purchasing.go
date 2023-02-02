// +build !minimal

package purchasing

import (
	"github.com/bluszcz/cutego"
	"github.com/bluszcz/cutego/core"
	"strings"
	"unsafe"
)

type QInAppProduct struct {
	core.QObject
}

type QInAppProduct_ITF interface {
	core.QObject_ITF
	QInAppProduct_PTR() *QInAppProduct
}

func (ptr *QInAppProduct) QInAppProduct_PTR() *QInAppProduct {
	return ptr
}

func (ptr *QInAppProduct) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QInAppProduct) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQInAppProduct(ptr QInAppProduct_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInAppProduct_PTR().Pointer()
	}
	return nil
}

func (n *QInAppProduct) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QInAppProduct) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQInAppProductFromPointer(ptr unsafe.Pointer) (n *QInAppProduct) {
	n = new(QInAppProduct)
	n.InitFromInternal(uintptr(ptr), "purchasing.QInAppProduct")
	return
}

func (ptr *QInAppProduct) DestroyQInAppProduct() {
}

//go:generate stringer -type=QInAppProduct__ProductType
//QInAppProduct::ProductType
type QInAppProduct__ProductType int64

const (
	QInAppProduct__Consumable QInAppProduct__ProductType = QInAppProduct__ProductType(0)
	QInAppProduct__Unlockable QInAppProduct__ProductType = QInAppProduct__ProductType(1)
)

func (ptr *QInAppProduct) Description() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Description"}).(string)
}

func (ptr *QInAppProduct) Identifier() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Identifier"}).(string)
}

func (ptr *QInAppProduct) Price() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Price"}).(string)
}

func (ptr *QInAppProduct) ProductType() QInAppProduct__ProductType {

	return QInAppProduct__ProductType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProductType"}).(float64))
}

func (ptr *QInAppProduct) ConnectPurchase(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPurchase", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppProduct) DisconnectPurchase() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPurchase"})
}

func (ptr *QInAppProduct) Purchase() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Purchase"})
}

func (ptr *QInAppProduct) Title() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Title"}).(string)
}

func (ptr *QInAppProduct) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QInAppProduct) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QInAppProduct) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QInAppProduct) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QInAppProduct) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QInAppProduct) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QInAppProduct) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QInAppProduct) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QInAppProduct) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QInAppProduct) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QInAppProduct) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QInAppProduct) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QInAppProduct) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QInAppProduct) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QInAppProduct) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QInAppProduct) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QInAppProduct) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QInAppProduct) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QInAppProduct) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QInAppProduct) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QInAppProduct) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QInAppStore struct {
	core.QObject
}

type QInAppStore_ITF interface {
	core.QObject_ITF
	QInAppStore_PTR() *QInAppStore
}

func (ptr *QInAppStore) QInAppStore_PTR() *QInAppStore {
	return ptr
}

func (ptr *QInAppStore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QInAppStore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQInAppStore(ptr QInAppStore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInAppStore_PTR().Pointer()
	}
	return nil
}

func (n *QInAppStore) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QInAppStore) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQInAppStoreFromPointer(ptr unsafe.Pointer) (n *QInAppStore) {
	n = new(QInAppStore)
	n.InitFromInternal(uintptr(ptr), "purchasing.QInAppStore")
	return
}
func NewQInAppStore(parent core.QObject_ITF) *QInAppStore {

	return internal.CallLocalFunction([]interface{}{"", "", "purchasing.NewQInAppStore", "", parent}).(*QInAppStore)
}

func (ptr *QInAppStore) ConnectProductRegistered(f func(product *QInAppProduct)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProductRegistered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppStore) DisconnectProductRegistered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProductRegistered"})
}

func (ptr *QInAppStore) ProductRegistered(product QInAppProduct_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProductRegistered", product})
}

func (ptr *QInAppStore) ConnectProductUnknown(f func(productType QInAppProduct__ProductType, identifier string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProductUnknown", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppStore) DisconnectProductUnknown() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProductUnknown"})
}

func (ptr *QInAppStore) ProductUnknown(productType QInAppProduct__ProductType, identifier string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProductUnknown", productType, identifier})
}

func (ptr *QInAppStore) RegisterProduct(productType QInAppProduct__ProductType, identifier string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterProduct", productType, identifier})
}

func (ptr *QInAppStore) RegisteredProduct(identifier string) *QInAppProduct {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisteredProduct", identifier}).(*QInAppProduct)
}

func (ptr *QInAppStore) RestorePurchases() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RestorePurchases"})
}

func (ptr *QInAppStore) SetPlatformProperty(propertyName string, value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPlatformProperty", propertyName, value})
}

func (ptr *QInAppStore) ConnectTransactionReady(f func(transaction *QInAppTransaction)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTransactionReady", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppStore) DisconnectTransactionReady() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTransactionReady"})
}

func (ptr *QInAppStore) TransactionReady(transaction QInAppTransaction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TransactionReady", transaction})
}

func (ptr *QInAppStore) ConnectDestroyQInAppStore(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQInAppStore", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppStore) DisconnectDestroyQInAppStore() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQInAppStore"})
}

func (ptr *QInAppStore) DestroyQInAppStore() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQInAppStore"})
}

func (ptr *QInAppStore) DestroyQInAppStoreDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQInAppStoreDefault"})
}

func (ptr *QInAppStore) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QInAppStore) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QInAppStore) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QInAppStore) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QInAppStore) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QInAppStore) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QInAppStore) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QInAppStore) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QInAppStore) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QInAppStore) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QInAppStore) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QInAppStore) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QInAppStore) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QInAppStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QInAppStore) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QInAppStore) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QInAppStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QInAppStore) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QInAppStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QInAppStore) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QInAppStore) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QInAppTransaction struct {
	core.QObject
}

type QInAppTransaction_ITF interface {
	core.QObject_ITF
	QInAppTransaction_PTR() *QInAppTransaction
}

func (ptr *QInAppTransaction) QInAppTransaction_PTR() *QInAppTransaction {
	return ptr
}

func (ptr *QInAppTransaction) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QInAppTransaction) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQInAppTransaction(ptr QInAppTransaction_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInAppTransaction_PTR().Pointer()
	}
	return nil
}

func (n *QInAppTransaction) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QInAppTransaction) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQInAppTransactionFromPointer(ptr unsafe.Pointer) (n *QInAppTransaction) {
	n = new(QInAppTransaction)
	n.InitFromInternal(uintptr(ptr), "purchasing.QInAppTransaction")
	return
}

func (ptr *QInAppTransaction) DestroyQInAppTransaction() {
}

//go:generate stringer -type=QInAppTransaction__TransactionStatus
//QInAppTransaction::TransactionStatus
type QInAppTransaction__TransactionStatus int64

const (
	QInAppTransaction__Unknown          QInAppTransaction__TransactionStatus = QInAppTransaction__TransactionStatus(0)
	QInAppTransaction__PurchaseApproved QInAppTransaction__TransactionStatus = QInAppTransaction__TransactionStatus(1)
	QInAppTransaction__PurchaseFailed   QInAppTransaction__TransactionStatus = QInAppTransaction__TransactionStatus(2)
	QInAppTransaction__PurchaseRestored QInAppTransaction__TransactionStatus = QInAppTransaction__TransactionStatus(3)
)

//go:generate stringer -type=QInAppTransaction__FailureReason
//QInAppTransaction::FailureReason
type QInAppTransaction__FailureReason int64

const (
	QInAppTransaction__NoFailure      QInAppTransaction__FailureReason = QInAppTransaction__FailureReason(0)
	QInAppTransaction__CanceledByUser QInAppTransaction__FailureReason = QInAppTransaction__FailureReason(1)
	QInAppTransaction__ErrorOccurred  QInAppTransaction__FailureReason = QInAppTransaction__FailureReason(2)
)

func (ptr *QInAppTransaction) ConnectErrorString(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorString", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppTransaction) DisconnectErrorString() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorString"})
}

func (ptr *QInAppTransaction) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QInAppTransaction) ErrorStringDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorStringDefault"}).(string)
}

func (ptr *QInAppTransaction) ConnectFailureReason(f func() QInAppTransaction__FailureReason) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFailureReason", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppTransaction) DisconnectFailureReason() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFailureReason"})
}

func (ptr *QInAppTransaction) FailureReason() QInAppTransaction__FailureReason {

	return QInAppTransaction__FailureReason(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FailureReason"}).(float64))
}

func (ptr *QInAppTransaction) FailureReasonDefault() QInAppTransaction__FailureReason {

	return QInAppTransaction__FailureReason(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FailureReasonDefault"}).(float64))
}

func (ptr *QInAppTransaction) ConnectFinalize(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinalize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppTransaction) DisconnectFinalize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinalize"})
}

func (ptr *QInAppTransaction) Finalize() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finalize"})
}

func (ptr *QInAppTransaction) ConnectOrderId(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOrderId", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppTransaction) DisconnectOrderId() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOrderId"})
}

func (ptr *QInAppTransaction) OrderId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OrderId"}).(string)
}

func (ptr *QInAppTransaction) OrderIdDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OrderIdDefault"}).(string)
}

func (ptr *QInAppTransaction) ConnectPlatformProperty(f func(propertyName string) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPlatformProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppTransaction) DisconnectPlatformProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPlatformProperty"})
}

func (ptr *QInAppTransaction) PlatformProperty(propertyName string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PlatformProperty", propertyName}).(string)
}

func (ptr *QInAppTransaction) PlatformPropertyDefault(propertyName string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PlatformPropertyDefault", propertyName}).(string)
}

func (ptr *QInAppTransaction) Product() *QInAppProduct {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Product"}).(*QInAppProduct)
}

func (ptr *QInAppTransaction) Status() QInAppTransaction__TransactionStatus {

	return QInAppTransaction__TransactionStatus(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Status"}).(float64))
}

func (ptr *QInAppTransaction) ConnectTimestamp(f func() *core.QDateTime) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTimestamp", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QInAppTransaction) DisconnectTimestamp() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTimestamp"})
}

func (ptr *QInAppTransaction) Timestamp() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Timestamp"}).(*core.QDateTime)
}

func (ptr *QInAppTransaction) TimestampDefault() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimestampDefault"}).(*core.QDateTime)
}

func (ptr *QInAppTransaction) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QInAppTransaction) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QInAppTransaction) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QInAppTransaction) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QInAppTransaction) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QInAppTransaction) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QInAppTransaction) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QInAppTransaction) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QInAppTransaction) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QInAppTransaction) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QInAppTransaction) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QInAppTransaction) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QInAppTransaction) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QInAppTransaction) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QInAppTransaction) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QInAppTransaction) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QInAppTransaction) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QInAppTransaction) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QInAppTransaction) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QInAppTransaction) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QInAppTransaction) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func init() {
	internal.ConstructorTable["purchasing.QInAppProduct"] = NewQInAppProductFromPointer
	internal.ConstructorTable["purchasing.QInAppStore"] = NewQInAppStoreFromPointer
	internal.ConstructorTable["purchasing.QInAppTransaction"] = NewQInAppTransactionFromPointer
}
