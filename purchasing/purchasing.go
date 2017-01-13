// +build !minimal

package purchasing

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "purchasing.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtPurchasing_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

//go:generate stringer -type=QInAppProduct__ProductType
//QInAppProduct::ProductType
type QInAppProduct__ProductType int64

const (
	QInAppProduct__Consumable QInAppProduct__ProductType = QInAppProduct__ProductType(0)
	QInAppProduct__Unlockable QInAppProduct__ProductType = QInAppProduct__ProductType(1)
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

func NewQInAppProductFromPointer(ptr unsafe.Pointer) *QInAppProduct {
	var n = new(QInAppProduct)
	n.SetPointer(ptr)
	return n
}

func (ptr *QInAppProduct) DestroyQInAppProduct() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QInAppProduct) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppProduct_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppProduct) Identifier() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppProduct_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppProduct) Price() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppProduct_Price(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppProduct) ProductType() QInAppProduct__ProductType {
	if ptr.Pointer() != nil {
		return QInAppProduct__ProductType(C.QInAppProduct_ProductType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInAppProduct) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppProduct_Title(ptr.Pointer()))
	}
	return ""
}

//export callbackQInAppProduct_Purchase
func callbackQInAppProduct_Purchase(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppProduct::purchase"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInAppProduct) ConnectPurchase(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::purchase", f)
	}
}

func (ptr *QInAppProduct) DisconnectPurchase() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::purchase")
	}
}

func (ptr *QInAppProduct) Purchase() {
	if ptr.Pointer() != nil {
		C.QInAppProduct_Purchase(ptr.Pointer())
	}
}

//export callbackQInAppProduct_TimerEvent
func callbackQInAppProduct_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppProduct::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInAppProductFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInAppProduct) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::timerEvent", f)
	}
}

func (ptr *QInAppProduct) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::timerEvent")
	}
}

func (ptr *QInAppProduct) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QInAppProduct) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQInAppProduct_ChildEvent
func callbackQInAppProduct_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppProduct::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInAppProductFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInAppProduct) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::childEvent", f)
	}
}

func (ptr *QInAppProduct) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::childEvent")
	}
}

func (ptr *QInAppProduct) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QInAppProduct) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQInAppProduct_ConnectNotify
func callbackQInAppProduct_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppProduct::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppProductFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppProduct) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::connectNotify", f)
	}
}

func (ptr *QInAppProduct) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::connectNotify")
	}
}

func (ptr *QInAppProduct) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppProduct) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppProduct_CustomEvent
func callbackQInAppProduct_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppProduct::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInAppProductFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInAppProduct) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::customEvent", f)
	}
}

func (ptr *QInAppProduct) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::customEvent")
	}
}

func (ptr *QInAppProduct) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInAppProduct) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQInAppProduct_DeleteLater
func callbackQInAppProduct_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppProduct::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQInAppProductFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QInAppProduct) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::deleteLater", f)
	}
}

func (ptr *QInAppProduct) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::deleteLater")
	}
}

func (ptr *QInAppProduct) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QInAppProduct_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QInAppProduct) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QInAppProduct_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQInAppProduct_DisconnectNotify
func callbackQInAppProduct_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppProduct::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppProductFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppProduct) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::disconnectNotify", f)
	}
}

func (ptr *QInAppProduct) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::disconnectNotify")
	}
}

func (ptr *QInAppProduct) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppProduct) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppProduct_Event
func callbackQInAppProduct_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppProduct::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppProductFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QInAppProduct) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::event", f)
	}
}

func (ptr *QInAppProduct) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::event")
	}
}

func (ptr *QInAppProduct) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppProduct_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QInAppProduct) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppProduct_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQInAppProduct_EventFilter
func callbackQInAppProduct_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppProduct::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppProductFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QInAppProduct) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::eventFilter", f)
	}
}

func (ptr *QInAppProduct) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::eventFilter")
	}
}

func (ptr *QInAppProduct) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppProduct_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QInAppProduct) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppProduct_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQInAppProduct_MetaObject
func callbackQInAppProduct_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppProduct::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQInAppProductFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QInAppProduct) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::metaObject", f)
	}
}

func (ptr *QInAppProduct) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppProduct::metaObject")
	}
}

func (ptr *QInAppProduct) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppProduct_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QInAppProduct) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppProduct_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQInAppStoreFromPointer(ptr unsafe.Pointer) *QInAppStore {
	var n = new(QInAppStore)
	n.SetPointer(ptr)
	return n
}
func NewQInAppStore(parent core.QObject_ITF) *QInAppStore {
	var tmpValue = NewQInAppStoreFromPointer(C.QInAppStore_NewQInAppStore(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQInAppStore_ProductRegistered
func callbackQInAppStore_ProductRegistered(ptr unsafe.Pointer, product unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::productRegistered"); signal != nil {
		signal.(func(*QInAppProduct))(NewQInAppProductFromPointer(product))
	}

}

func (ptr *QInAppStore) ConnectProductRegistered(f func(product *QInAppProduct)) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectProductRegistered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::productRegistered", f)
	}
}

func (ptr *QInAppStore) DisconnectProductRegistered() {
	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectProductRegistered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::productRegistered")
	}
}

func (ptr *QInAppStore) ProductRegistered(product QInAppProduct_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ProductRegistered(ptr.Pointer(), PointerFromQInAppProduct(product))
	}
}

//export callbackQInAppStore_ProductUnknown
func callbackQInAppStore_ProductUnknown(ptr unsafe.Pointer, productType C.longlong, identifier C.struct_QtPurchasing_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::productUnknown"); signal != nil {
		signal.(func(QInAppProduct__ProductType, string))(QInAppProduct__ProductType(productType), cGoUnpackString(identifier))
	}

}

func (ptr *QInAppStore) ConnectProductUnknown(f func(productType QInAppProduct__ProductType, identifier string)) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectProductUnknown(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::productUnknown", f)
	}
}

func (ptr *QInAppStore) DisconnectProductUnknown() {
	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectProductUnknown(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::productUnknown")
	}
}

func (ptr *QInAppStore) ProductUnknown(productType QInAppProduct__ProductType, identifier string) {
	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		C.QInAppStore_ProductUnknown(ptr.Pointer(), C.longlong(productType), identifierC)
	}
}

func (ptr *QInAppStore) RegisterProduct(productType QInAppProduct__ProductType, identifier string) {
	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		C.QInAppStore_RegisterProduct(ptr.Pointer(), C.longlong(productType), identifierC)
	}
}

func (ptr *QInAppStore) RegisteredProduct(identifier string) *QInAppProduct {
	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		var tmpValue = NewQInAppProductFromPointer(C.QInAppStore_RegisteredProduct(ptr.Pointer(), identifierC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppStore) RestorePurchases() {
	if ptr.Pointer() != nil {
		C.QInAppStore_RestorePurchases(ptr.Pointer())
	}
}

func (ptr *QInAppStore) SetPlatformProperty(propertyName string, value string) {
	if ptr.Pointer() != nil {
		var propertyNameC = C.CString(propertyName)
		defer C.free(unsafe.Pointer(propertyNameC))
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QInAppStore_SetPlatformProperty(ptr.Pointer(), propertyNameC, valueC)
	}
}

//export callbackQInAppStore_TransactionReady
func callbackQInAppStore_TransactionReady(ptr unsafe.Pointer, transaction unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::transactionReady"); signal != nil {
		signal.(func(*QInAppTransaction))(NewQInAppTransactionFromPointer(transaction))
	}

}

func (ptr *QInAppStore) ConnectTransactionReady(f func(transaction *QInAppTransaction)) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectTransactionReady(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::transactionReady", f)
	}
}

func (ptr *QInAppStore) DisconnectTransactionReady() {
	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectTransactionReady(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::transactionReady")
	}
}

func (ptr *QInAppStore) TransactionReady(transaction QInAppTransaction_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_TransactionReady(ptr.Pointer(), PointerFromQInAppTransaction(transaction))
	}
}

func (ptr *QInAppStore) DestroyQInAppStore() {
	if ptr.Pointer() != nil {
		C.QInAppStore_DestroyQInAppStore(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQInAppStore_TimerEvent
func callbackQInAppStore_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInAppStoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInAppStore) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::timerEvent", f)
	}
}

func (ptr *QInAppStore) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::timerEvent")
	}
}

func (ptr *QInAppStore) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QInAppStore) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQInAppStore_ChildEvent
func callbackQInAppStore_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInAppStoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInAppStore) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::childEvent", f)
	}
}

func (ptr *QInAppStore) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::childEvent")
	}
}

func (ptr *QInAppStore) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QInAppStore) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQInAppStore_ConnectNotify
func callbackQInAppStore_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppStoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppStore) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::connectNotify", f)
	}
}

func (ptr *QInAppStore) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::connectNotify")
	}
}

func (ptr *QInAppStore) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppStore_CustomEvent
func callbackQInAppStore_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInAppStoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInAppStore) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::customEvent", f)
	}
}

func (ptr *QInAppStore) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::customEvent")
	}
}

func (ptr *QInAppStore) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInAppStore) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQInAppStore_DeleteLater
func callbackQInAppStore_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQInAppStoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QInAppStore) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::deleteLater", f)
	}
}

func (ptr *QInAppStore) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::deleteLater")
	}
}

func (ptr *QInAppStore) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QInAppStore_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QInAppStore) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QInAppStore_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQInAppStore_DisconnectNotify
func callbackQInAppStore_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppStoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppStore) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::disconnectNotify", f)
	}
}

func (ptr *QInAppStore) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::disconnectNotify")
	}
}

func (ptr *QInAppStore) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppStore_Event
func callbackQInAppStore_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppStoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QInAppStore) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::event", f)
	}
}

func (ptr *QInAppStore) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::event")
	}
}

func (ptr *QInAppStore) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppStore_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QInAppStore) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppStore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQInAppStore_EventFilter
func callbackQInAppStore_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppStoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QInAppStore) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::eventFilter", f)
	}
}

func (ptr *QInAppStore) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::eventFilter")
	}
}

func (ptr *QInAppStore) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppStore_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QInAppStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppStore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQInAppStore_MetaObject
func callbackQInAppStore_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppStore::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQInAppStoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QInAppStore) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::metaObject", f)
	}
}

func (ptr *QInAppStore) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppStore::metaObject")
	}
}

func (ptr *QInAppStore) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppStore_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QInAppStore) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppStore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//go:generate stringer -type=QInAppTransaction__FailureReason
//QInAppTransaction::FailureReason
type QInAppTransaction__FailureReason int64

const (
	QInAppTransaction__NoFailure      QInAppTransaction__FailureReason = QInAppTransaction__FailureReason(0)
	QInAppTransaction__CanceledByUser QInAppTransaction__FailureReason = QInAppTransaction__FailureReason(1)
	QInAppTransaction__ErrorOccurred  QInAppTransaction__FailureReason = QInAppTransaction__FailureReason(2)
)

//go:generate stringer -type=QInAppTransaction__TransactionStatus
//QInAppTransaction::TransactionStatus
type QInAppTransaction__TransactionStatus int64

const (
	QInAppTransaction__Unknown          QInAppTransaction__TransactionStatus = QInAppTransaction__TransactionStatus(0)
	QInAppTransaction__PurchaseApproved QInAppTransaction__TransactionStatus = QInAppTransaction__TransactionStatus(1)
	QInAppTransaction__PurchaseFailed   QInAppTransaction__TransactionStatus = QInAppTransaction__TransactionStatus(2)
	QInAppTransaction__PurchaseRestored QInAppTransaction__TransactionStatus = QInAppTransaction__TransactionStatus(3)
)

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

func NewQInAppTransactionFromPointer(ptr unsafe.Pointer) *QInAppTransaction {
	var n = new(QInAppTransaction)
	n.SetPointer(ptr)
	return n
}

func (ptr *QInAppTransaction) DestroyQInAppTransaction() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQInAppTransaction_ErrorString
func callbackQInAppTransaction_ErrorString(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQInAppTransactionFromPointer(ptr).ErrorStringDefault())
}

func (ptr *QInAppTransaction) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::errorString", f)
	}
}

func (ptr *QInAppTransaction) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::errorString")
	}
}

func (ptr *QInAppTransaction) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppTransaction_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppTransaction) ErrorStringDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppTransaction_ErrorStringDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQInAppTransaction_FailureReason
func callbackQInAppTransaction_FailureReason(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::failureReason"); signal != nil {
		return C.longlong(signal.(func() QInAppTransaction__FailureReason)())
	}

	return C.longlong(NewQInAppTransactionFromPointer(ptr).FailureReasonDefault())
}

func (ptr *QInAppTransaction) ConnectFailureReason(f func() QInAppTransaction__FailureReason) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::failureReason", f)
	}
}

func (ptr *QInAppTransaction) DisconnectFailureReason() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::failureReason")
	}
}

func (ptr *QInAppTransaction) FailureReason() QInAppTransaction__FailureReason {
	if ptr.Pointer() != nil {
		return QInAppTransaction__FailureReason(C.QInAppTransaction_FailureReason(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInAppTransaction) FailureReasonDefault() QInAppTransaction__FailureReason {
	if ptr.Pointer() != nil {
		return QInAppTransaction__FailureReason(C.QInAppTransaction_FailureReasonDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQInAppTransaction_OrderId
func callbackQInAppTransaction_OrderId(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::orderId"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQInAppTransactionFromPointer(ptr).OrderIdDefault())
}

func (ptr *QInAppTransaction) ConnectOrderId(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::orderId", f)
	}
}

func (ptr *QInAppTransaction) DisconnectOrderId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::orderId")
	}
}

func (ptr *QInAppTransaction) OrderId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppTransaction_OrderId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppTransaction) OrderIdDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppTransaction_OrderIdDefault(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppTransaction) Product() *QInAppProduct {
	if ptr.Pointer() != nil {
		var tmpValue = NewQInAppProductFromPointer(C.QInAppTransaction_Product(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppTransaction) Status() QInAppTransaction__TransactionStatus {
	if ptr.Pointer() != nil {
		return QInAppTransaction__TransactionStatus(C.QInAppTransaction_Status(ptr.Pointer()))
	}
	return 0
}

//export callbackQInAppTransaction_Timestamp
func callbackQInAppTransaction_Timestamp(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::timestamp"); signal != nil {
		return core.PointerFromQDateTime(signal.(func() *core.QDateTime)())
	}

	return core.PointerFromQDateTime(NewQInAppTransactionFromPointer(ptr).TimestampDefault())
}

func (ptr *QInAppTransaction) ConnectTimestamp(f func() *core.QDateTime) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::timestamp", f)
	}
}

func (ptr *QInAppTransaction) DisconnectTimestamp() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::timestamp")
	}
}

func (ptr *QInAppTransaction) Timestamp() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QInAppTransaction_Timestamp(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QInAppTransaction) TimestampDefault() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QInAppTransaction_TimestampDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

//export callbackQInAppTransaction_Finalize
func callbackQInAppTransaction_Finalize(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::finalize"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInAppTransaction) ConnectFinalize(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::finalize", f)
	}
}

func (ptr *QInAppTransaction) DisconnectFinalize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::finalize")
	}
}

func (ptr *QInAppTransaction) Finalize() {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_Finalize(ptr.Pointer())
	}
}

//export callbackQInAppTransaction_PlatformProperty
func callbackQInAppTransaction_PlatformProperty(ptr unsafe.Pointer, propertyName C.struct_QtPurchasing_PackedString) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::platformProperty"); signal != nil {
		return C.CString(signal.(func(string) string)(cGoUnpackString(propertyName)))
	}

	return C.CString(NewQInAppTransactionFromPointer(ptr).PlatformPropertyDefault(cGoUnpackString(propertyName)))
}

func (ptr *QInAppTransaction) ConnectPlatformProperty(f func(propertyName string) string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::platformProperty", f)
	}
}

func (ptr *QInAppTransaction) DisconnectPlatformProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::platformProperty")
	}
}

func (ptr *QInAppTransaction) PlatformProperty(propertyName string) string {
	if ptr.Pointer() != nil {
		var propertyNameC = C.CString(propertyName)
		defer C.free(unsafe.Pointer(propertyNameC))
		return cGoUnpackString(C.QInAppTransaction_PlatformProperty(ptr.Pointer(), propertyNameC))
	}
	return ""
}

func (ptr *QInAppTransaction) PlatformPropertyDefault(propertyName string) string {
	if ptr.Pointer() != nil {
		var propertyNameC = C.CString(propertyName)
		defer C.free(unsafe.Pointer(propertyNameC))
		return cGoUnpackString(C.QInAppTransaction_PlatformPropertyDefault(ptr.Pointer(), propertyNameC))
	}
	return ""
}

//export callbackQInAppTransaction_TimerEvent
func callbackQInAppTransaction_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInAppTransactionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInAppTransaction) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::timerEvent", f)
	}
}

func (ptr *QInAppTransaction) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::timerEvent")
	}
}

func (ptr *QInAppTransaction) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QInAppTransaction) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQInAppTransaction_ChildEvent
func callbackQInAppTransaction_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInAppTransactionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInAppTransaction) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::childEvent", f)
	}
}

func (ptr *QInAppTransaction) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::childEvent")
	}
}

func (ptr *QInAppTransaction) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QInAppTransaction) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQInAppTransaction_ConnectNotify
func callbackQInAppTransaction_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppTransactionFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppTransaction) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::connectNotify", f)
	}
}

func (ptr *QInAppTransaction) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::connectNotify")
	}
}

func (ptr *QInAppTransaction) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppTransaction) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppTransaction_CustomEvent
func callbackQInAppTransaction_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInAppTransactionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInAppTransaction) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::customEvent", f)
	}
}

func (ptr *QInAppTransaction) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::customEvent")
	}
}

func (ptr *QInAppTransaction) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInAppTransaction) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQInAppTransaction_DeleteLater
func callbackQInAppTransaction_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQInAppTransactionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QInAppTransaction) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::deleteLater", f)
	}
}

func (ptr *QInAppTransaction) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::deleteLater")
	}
}

func (ptr *QInAppTransaction) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QInAppTransaction) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQInAppTransaction_DisconnectNotify
func callbackQInAppTransaction_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppTransactionFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppTransaction) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::disconnectNotify", f)
	}
}

func (ptr *QInAppTransaction) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::disconnectNotify")
	}
}

func (ptr *QInAppTransaction) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppTransaction) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppTransaction_Event
func callbackQInAppTransaction_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppTransactionFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QInAppTransaction) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::event", f)
	}
}

func (ptr *QInAppTransaction) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::event")
	}
}

func (ptr *QInAppTransaction) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppTransaction_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QInAppTransaction) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppTransaction_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQInAppTransaction_EventFilter
func callbackQInAppTransaction_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppTransactionFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QInAppTransaction) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::eventFilter", f)
	}
}

func (ptr *QInAppTransaction) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::eventFilter")
	}
}

func (ptr *QInAppTransaction) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppTransaction_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QInAppTransaction) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppTransaction_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQInAppTransaction_MetaObject
func callbackQInAppTransaction_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QInAppTransaction::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQInAppTransactionFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QInAppTransaction) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::metaObject", f)
	}
}

func (ptr *QInAppTransaction) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QInAppTransaction::metaObject")
	}
}

func (ptr *QInAppTransaction) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppTransaction_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QInAppTransaction) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppTransaction_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
