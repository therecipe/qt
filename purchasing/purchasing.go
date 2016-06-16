// +build !minimal

package purchasing

//#include "purchasing.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

//QInAppProduct::ProductType
type QInAppProduct__ProductType int64

const (
	QInAppProduct__Consumable = QInAppProduct__ProductType(0)
	QInAppProduct__Unlockable = QInAppProduct__ProductType(1)
)

type QInAppProduct struct {
	core.QObject
}

type QInAppProduct_ITF interface {
	core.QObject_ITF
	QInAppProduct_PTR() *QInAppProduct
}

func (p *QInAppProduct) QInAppProduct_PTR() *QInAppProduct {
	return p
}

func (p *QInAppProduct) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QInAppProduct) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQInAppProductFromPointer(ptr unsafe.Pointer) *QInAppProduct {
	var n = NewQInAppProductFromPointer(ptr)
	for len(n.ObjectName()) < len("QInAppProduct_") {
		n.SetObjectName("QInAppProduct_" + qt.Identifier())
	}
	return n
}

func (ptr *QInAppProduct) Description() string {
	defer qt.Recovering("QInAppProduct::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInAppProduct_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppProduct) Identifier() string {
	defer qt.Recovering("QInAppProduct::identifier")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInAppProduct_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppProduct) Price() string {
	defer qt.Recovering("QInAppProduct::price")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInAppProduct_Price(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppProduct) ProductType() QInAppProduct__ProductType {
	defer qt.Recovering("QInAppProduct::productType")

	if ptr.Pointer() != nil {
		return QInAppProduct__ProductType(C.QInAppProduct_ProductType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInAppProduct) Title() string {
	defer qt.Recovering("QInAppProduct::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInAppProduct_Title(ptr.Pointer()))
	}
	return ""
}

//export callbackQInAppProduct_Purchase
func callbackQInAppProduct_Purchase(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QInAppProduct::purchase")

	if signal := qt.GetSignal(C.GoString(ptrName), "purchase"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInAppProduct) ConnectPurchase(f func()) {
	defer qt.Recovering("connect QInAppProduct::purchase")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "purchase", f)
	}
}

func (ptr *QInAppProduct) DisconnectPurchase() {
	defer qt.Recovering("disconnect QInAppProduct::purchase")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "purchase")
	}
}

func (ptr *QInAppProduct) Purchase() {
	defer qt.Recovering("QInAppProduct::purchase")

	if ptr.Pointer() != nil {
		C.QInAppProduct_Purchase(ptr.Pointer())
	}
}

//export callbackQInAppProduct_TimerEvent
func callbackQInAppProduct_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInAppProduct::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInAppProductFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInAppProduct) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QInAppProduct::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QInAppProduct) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QInAppProduct::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QInAppProduct) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QInAppProduct::timerEvent")

	if ptr.Pointer() != nil {
		C.QInAppProduct_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QInAppProduct) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QInAppProduct::timerEvent")

	if ptr.Pointer() != nil {
		C.QInAppProduct_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQInAppProduct_ChildEvent
func callbackQInAppProduct_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInAppProduct::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInAppProductFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInAppProduct) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QInAppProduct::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QInAppProduct) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QInAppProduct::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QInAppProduct) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QInAppProduct::childEvent")

	if ptr.Pointer() != nil {
		C.QInAppProduct_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QInAppProduct) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QInAppProduct::childEvent")

	if ptr.Pointer() != nil {
		C.QInAppProduct_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQInAppProduct_ConnectNotify
func callbackQInAppProduct_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QInAppProduct::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppProductFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppProduct) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QInAppProduct::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QInAppProduct) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QInAppProduct::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QInAppProduct) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppProduct::connectNotify")

	if ptr.Pointer() != nil {
		C.QInAppProduct_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppProduct) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppProduct::connectNotify")

	if ptr.Pointer() != nil {
		C.QInAppProduct_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppProduct_CustomEvent
func callbackQInAppProduct_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInAppProduct::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInAppProductFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInAppProduct) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QInAppProduct::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QInAppProduct) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QInAppProduct::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QInAppProduct) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QInAppProduct::customEvent")

	if ptr.Pointer() != nil {
		C.QInAppProduct_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInAppProduct) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QInAppProduct::customEvent")

	if ptr.Pointer() != nil {
		C.QInAppProduct_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQInAppProduct_DeleteLater
func callbackQInAppProduct_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QInAppProduct::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQInAppProductFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QInAppProduct) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QInAppProduct::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QInAppProduct) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QInAppProduct::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QInAppProduct) DeleteLater() {
	defer qt.Recovering("QInAppProduct::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QInAppProduct_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QInAppProduct) DeleteLaterDefault() {
	defer qt.Recovering("QInAppProduct::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QInAppProduct_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQInAppProduct_DisconnectNotify
func callbackQInAppProduct_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QInAppProduct::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppProductFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppProduct) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QInAppProduct::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QInAppProduct) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QInAppProduct::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QInAppProduct) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppProduct::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QInAppProduct_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppProduct) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppProduct::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QInAppProduct_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppProduct_Event
func callbackQInAppProduct_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QInAppProduct::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQInAppProductFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QInAppProduct) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QInAppProduct::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QInAppProduct) DisconnectEvent() {
	defer qt.Recovering("disconnect QInAppProduct::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QInAppProduct) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppProduct::event")

	if ptr.Pointer() != nil {
		return C.QInAppProduct_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QInAppProduct) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppProduct::event")

	if ptr.Pointer() != nil {
		return C.QInAppProduct_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQInAppProduct_EventFilter
func callbackQInAppProduct_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QInAppProduct::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQInAppProductFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QInAppProduct) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QInAppProduct::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QInAppProduct) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QInAppProduct::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QInAppProduct) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppProduct::eventFilter")

	if ptr.Pointer() != nil {
		return C.QInAppProduct_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QInAppProduct) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppProduct::eventFilter")

	if ptr.Pointer() != nil {
		return C.QInAppProduct_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQInAppProduct_MetaObject
func callbackQInAppProduct_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QInAppProduct::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQInAppProductFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QInAppProduct) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QInAppProduct::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QInAppProduct) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QInAppProduct::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QInAppProduct) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QInAppProduct::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppProduct_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QInAppProduct) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QInAppProduct::metaObject")

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

func (p *QInAppStore) QInAppStore_PTR() *QInAppStore {
	return p
}

func (p *QInAppStore) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QInAppStore) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQInAppStoreFromPointer(ptr unsafe.Pointer) *QInAppStore {
	var n = NewQInAppStoreFromPointer(ptr)
	for len(n.ObjectName()) < len("QInAppStore_") {
		n.SetObjectName("QInAppStore_" + qt.Identifier())
	}
	return n
}

func NewQInAppStore(parent core.QObject_ITF) *QInAppStore {
	defer qt.Recovering("QInAppStore::QInAppStore")

	return newQInAppStoreFromPointer(C.QInAppStore_NewQInAppStore(core.PointerFromQObject(parent)))
}

//export callbackQInAppStore_ProductRegistered
func callbackQInAppStore_ProductRegistered(ptr unsafe.Pointer, ptrName *C.char, product unsafe.Pointer) {
	defer qt.Recovering("callback QInAppStore::productRegistered")

	if signal := qt.GetSignal(C.GoString(ptrName), "productRegistered"); signal != nil {
		signal.(func(*QInAppProduct))(NewQInAppProductFromPointer(product))
	}

}

func (ptr *QInAppStore) ConnectProductRegistered(f func(product *QInAppProduct)) {
	defer qt.Recovering("connect QInAppStore::productRegistered")

	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectProductRegistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "productRegistered", f)
	}
}

func (ptr *QInAppStore) DisconnectProductRegistered() {
	defer qt.Recovering("disconnect QInAppStore::productRegistered")

	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectProductRegistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "productRegistered")
	}
}

func (ptr *QInAppStore) ProductRegistered(product QInAppProduct_ITF) {
	defer qt.Recovering("QInAppStore::productRegistered")

	if ptr.Pointer() != nil {
		C.QInAppStore_ProductRegistered(ptr.Pointer(), PointerFromQInAppProduct(product))
	}
}

//export callbackQInAppStore_ProductUnknown
func callbackQInAppStore_ProductUnknown(ptr unsafe.Pointer, ptrName *C.char, productType C.int, identifier *C.char) {
	defer qt.Recovering("callback QInAppStore::productUnknown")

	if signal := qt.GetSignal(C.GoString(ptrName), "productUnknown"); signal != nil {
		signal.(func(QInAppProduct__ProductType, string))(QInAppProduct__ProductType(productType), C.GoString(identifier))
	}

}

func (ptr *QInAppStore) ConnectProductUnknown(f func(productType QInAppProduct__ProductType, identifier string)) {
	defer qt.Recovering("connect QInAppStore::productUnknown")

	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectProductUnknown(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "productUnknown", f)
	}
}

func (ptr *QInAppStore) DisconnectProductUnknown() {
	defer qt.Recovering("disconnect QInAppStore::productUnknown")

	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectProductUnknown(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "productUnknown")
	}
}

func (ptr *QInAppStore) ProductUnknown(productType QInAppProduct__ProductType, identifier string) {
	defer qt.Recovering("QInAppStore::productUnknown")

	if ptr.Pointer() != nil {
		C.QInAppStore_ProductUnknown(ptr.Pointer(), C.int(productType), C.CString(identifier))
	}
}

func (ptr *QInAppStore) RegisterProduct(productType QInAppProduct__ProductType, identifier string) {
	defer qt.Recovering("QInAppStore::registerProduct")

	if ptr.Pointer() != nil {
		C.QInAppStore_RegisterProduct(ptr.Pointer(), C.int(productType), C.CString(identifier))
	}
}

func (ptr *QInAppStore) RegisteredProduct(identifier string) *QInAppProduct {
	defer qt.Recovering("QInAppStore::registeredProduct")

	if ptr.Pointer() != nil {
		return NewQInAppProductFromPointer(C.QInAppStore_RegisteredProduct(ptr.Pointer(), C.CString(identifier)))
	}
	return nil
}

func (ptr *QInAppStore) RestorePurchases() {
	defer qt.Recovering("QInAppStore::restorePurchases")

	if ptr.Pointer() != nil {
		C.QInAppStore_RestorePurchases(ptr.Pointer())
	}
}

func (ptr *QInAppStore) SetPlatformProperty(propertyName string, value string) {
	defer qt.Recovering("QInAppStore::setPlatformProperty")

	if ptr.Pointer() != nil {
		C.QInAppStore_SetPlatformProperty(ptr.Pointer(), C.CString(propertyName), C.CString(value))
	}
}

//export callbackQInAppStore_TransactionReady
func callbackQInAppStore_TransactionReady(ptr unsafe.Pointer, ptrName *C.char, transaction unsafe.Pointer) {
	defer qt.Recovering("callback QInAppStore::transactionReady")

	if signal := qt.GetSignal(C.GoString(ptrName), "transactionReady"); signal != nil {
		signal.(func(*QInAppTransaction))(NewQInAppTransactionFromPointer(transaction))
	}

}

func (ptr *QInAppStore) ConnectTransactionReady(f func(transaction *QInAppTransaction)) {
	defer qt.Recovering("connect QInAppStore::transactionReady")

	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectTransactionReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "transactionReady", f)
	}
}

func (ptr *QInAppStore) DisconnectTransactionReady() {
	defer qt.Recovering("disconnect QInAppStore::transactionReady")

	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectTransactionReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "transactionReady")
	}
}

func (ptr *QInAppStore) TransactionReady(transaction QInAppTransaction_ITF) {
	defer qt.Recovering("QInAppStore::transactionReady")

	if ptr.Pointer() != nil {
		C.QInAppStore_TransactionReady(ptr.Pointer(), PointerFromQInAppTransaction(transaction))
	}
}

func (ptr *QInAppStore) DestroyQInAppStore() {
	defer qt.Recovering("QInAppStore::~QInAppStore")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QInAppStore_DestroyQInAppStore(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQInAppStore_TimerEvent
func callbackQInAppStore_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInAppStore::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInAppStoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInAppStore) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QInAppStore::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QInAppStore) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QInAppStore::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QInAppStore) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QInAppStore::timerEvent")

	if ptr.Pointer() != nil {
		C.QInAppStore_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QInAppStore) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QInAppStore::timerEvent")

	if ptr.Pointer() != nil {
		C.QInAppStore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQInAppStore_ChildEvent
func callbackQInAppStore_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInAppStore::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInAppStoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInAppStore) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QInAppStore::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QInAppStore) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QInAppStore::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QInAppStore) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QInAppStore::childEvent")

	if ptr.Pointer() != nil {
		C.QInAppStore_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QInAppStore) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QInAppStore::childEvent")

	if ptr.Pointer() != nil {
		C.QInAppStore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQInAppStore_ConnectNotify
func callbackQInAppStore_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QInAppStore::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppStoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppStore) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QInAppStore::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QInAppStore) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QInAppStore::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QInAppStore) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppStore::connectNotify")

	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppStore::connectNotify")

	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppStore_CustomEvent
func callbackQInAppStore_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInAppStore::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInAppStoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInAppStore) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QInAppStore::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QInAppStore) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QInAppStore::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QInAppStore) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QInAppStore::customEvent")

	if ptr.Pointer() != nil {
		C.QInAppStore_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInAppStore) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QInAppStore::customEvent")

	if ptr.Pointer() != nil {
		C.QInAppStore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQInAppStore_DeleteLater
func callbackQInAppStore_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QInAppStore::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQInAppStoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QInAppStore) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QInAppStore::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QInAppStore) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QInAppStore::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QInAppStore) DeleteLater() {
	defer qt.Recovering("QInAppStore::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QInAppStore_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QInAppStore) DeleteLaterDefault() {
	defer qt.Recovering("QInAppStore::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QInAppStore_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQInAppStore_DisconnectNotify
func callbackQInAppStore_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QInAppStore::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppStoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppStore) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QInAppStore::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QInAppStore) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QInAppStore::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QInAppStore) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppStore::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppStore::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppStore_Event
func callbackQInAppStore_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QInAppStore::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQInAppStoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QInAppStore) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QInAppStore::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QInAppStore) DisconnectEvent() {
	defer qt.Recovering("disconnect QInAppStore::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QInAppStore) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppStore::event")

	if ptr.Pointer() != nil {
		return C.QInAppStore_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QInAppStore) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppStore::event")

	if ptr.Pointer() != nil {
		return C.QInAppStore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQInAppStore_EventFilter
func callbackQInAppStore_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QInAppStore::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQInAppStoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QInAppStore) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QInAppStore::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QInAppStore) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QInAppStore::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QInAppStore) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppStore::eventFilter")

	if ptr.Pointer() != nil {
		return C.QInAppStore_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QInAppStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppStore::eventFilter")

	if ptr.Pointer() != nil {
		return C.QInAppStore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQInAppStore_MetaObject
func callbackQInAppStore_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QInAppStore::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQInAppStoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QInAppStore) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QInAppStore::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QInAppStore) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QInAppStore::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QInAppStore) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QInAppStore::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppStore_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QInAppStore) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QInAppStore::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppStore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QInAppTransaction::FailureReason
type QInAppTransaction__FailureReason int64

const (
	QInAppTransaction__NoFailure      = QInAppTransaction__FailureReason(0)
	QInAppTransaction__CanceledByUser = QInAppTransaction__FailureReason(1)
	QInAppTransaction__ErrorOccurred  = QInAppTransaction__FailureReason(2)
)

//QInAppTransaction::TransactionStatus
type QInAppTransaction__TransactionStatus int64

const (
	QInAppTransaction__Unknown          = QInAppTransaction__TransactionStatus(0)
	QInAppTransaction__PurchaseApproved = QInAppTransaction__TransactionStatus(1)
	QInAppTransaction__PurchaseFailed   = QInAppTransaction__TransactionStatus(2)
	QInAppTransaction__PurchaseRestored = QInAppTransaction__TransactionStatus(3)
)

type QInAppTransaction struct {
	core.QObject
}

type QInAppTransaction_ITF interface {
	core.QObject_ITF
	QInAppTransaction_PTR() *QInAppTransaction
}

func (p *QInAppTransaction) QInAppTransaction_PTR() *QInAppTransaction {
	return p
}

func (p *QInAppTransaction) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QInAppTransaction) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQInAppTransactionFromPointer(ptr unsafe.Pointer) *QInAppTransaction {
	var n = NewQInAppTransactionFromPointer(ptr)
	for len(n.ObjectName()) < len("QInAppTransaction_") {
		n.SetObjectName("QInAppTransaction_" + qt.Identifier())
	}
	return n
}

//export callbackQInAppTransaction_ErrorString
func callbackQInAppTransaction_ErrorString(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QInAppTransaction::errorString")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQInAppTransactionFromPointer(ptr).ErrorStringDefault())
}

func (ptr *QInAppTransaction) ConnectErrorString(f func() string) {
	defer qt.Recovering("connect QInAppTransaction::errorString")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "errorString", f)
	}
}

func (ptr *QInAppTransaction) DisconnectErrorString() {
	defer qt.Recovering("disconnect QInAppTransaction::errorString")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "errorString")
	}
}

func (ptr *QInAppTransaction) ErrorString() string {
	defer qt.Recovering("QInAppTransaction::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInAppTransaction_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppTransaction) ErrorStringDefault() string {
	defer qt.Recovering("QInAppTransaction::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInAppTransaction_ErrorStringDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQInAppTransaction_FailureReason
func callbackQInAppTransaction_FailureReason(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QInAppTransaction::failureReason")

	if signal := qt.GetSignal(C.GoString(ptrName), "failureReason"); signal != nil {
		return C.int(signal.(func() QInAppTransaction__FailureReason)())
	}

	return C.int(NewQInAppTransactionFromPointer(ptr).FailureReasonDefault())
}

func (ptr *QInAppTransaction) ConnectFailureReason(f func() QInAppTransaction__FailureReason) {
	defer qt.Recovering("connect QInAppTransaction::failureReason")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "failureReason", f)
	}
}

func (ptr *QInAppTransaction) DisconnectFailureReason() {
	defer qt.Recovering("disconnect QInAppTransaction::failureReason")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "failureReason")
	}
}

func (ptr *QInAppTransaction) FailureReason() QInAppTransaction__FailureReason {
	defer qt.Recovering("QInAppTransaction::failureReason")

	if ptr.Pointer() != nil {
		return QInAppTransaction__FailureReason(C.QInAppTransaction_FailureReason(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInAppTransaction) FailureReasonDefault() QInAppTransaction__FailureReason {
	defer qt.Recovering("QInAppTransaction::failureReason")

	if ptr.Pointer() != nil {
		return QInAppTransaction__FailureReason(C.QInAppTransaction_FailureReasonDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQInAppTransaction_OrderId
func callbackQInAppTransaction_OrderId(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QInAppTransaction::orderId")

	if signal := qt.GetSignal(C.GoString(ptrName), "orderId"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQInAppTransactionFromPointer(ptr).OrderIdDefault())
}

func (ptr *QInAppTransaction) ConnectOrderId(f func() string) {
	defer qt.Recovering("connect QInAppTransaction::orderId")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "orderId", f)
	}
}

func (ptr *QInAppTransaction) DisconnectOrderId() {
	defer qt.Recovering("disconnect QInAppTransaction::orderId")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "orderId")
	}
}

func (ptr *QInAppTransaction) OrderId() string {
	defer qt.Recovering("QInAppTransaction::orderId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInAppTransaction_OrderId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppTransaction) OrderIdDefault() string {
	defer qt.Recovering("QInAppTransaction::orderId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInAppTransaction_OrderIdDefault(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppTransaction) Product() *QInAppProduct {
	defer qt.Recovering("QInAppTransaction::product")

	if ptr.Pointer() != nil {
		return NewQInAppProductFromPointer(C.QInAppTransaction_Product(ptr.Pointer()))
	}
	return nil
}

func (ptr *QInAppTransaction) Status() QInAppTransaction__TransactionStatus {
	defer qt.Recovering("QInAppTransaction::status")

	if ptr.Pointer() != nil {
		return QInAppTransaction__TransactionStatus(C.QInAppTransaction_Status(ptr.Pointer()))
	}
	return 0
}

//export callbackQInAppTransaction_Timestamp
func callbackQInAppTransaction_Timestamp(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QInAppTransaction::timestamp")

	if signal := qt.GetSignal(C.GoString(ptrName), "timestamp"); signal != nil {
		return core.PointerFromQDateTime(signal.(func() *core.QDateTime)())
	}

	return core.PointerFromQDateTime(NewQInAppTransactionFromPointer(ptr).TimestampDefault())
}

func (ptr *QInAppTransaction) ConnectTimestamp(f func() *core.QDateTime) {
	defer qt.Recovering("connect QInAppTransaction::timestamp")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timestamp", f)
	}
}

func (ptr *QInAppTransaction) DisconnectTimestamp() {
	defer qt.Recovering("disconnect QInAppTransaction::timestamp")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timestamp")
	}
}

func (ptr *QInAppTransaction) Timestamp() *core.QDateTime {
	defer qt.Recovering("QInAppTransaction::timestamp")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QInAppTransaction_Timestamp(ptr.Pointer()))
	}
	return nil
}

func (ptr *QInAppTransaction) TimestampDefault() *core.QDateTime {
	defer qt.Recovering("QInAppTransaction::timestamp")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QInAppTransaction_TimestampDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQInAppTransaction_Finalize
func callbackQInAppTransaction_Finalize(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QInAppTransaction::finalize")

	if signal := qt.GetSignal(C.GoString(ptrName), "finalize"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInAppTransaction) ConnectFinalize(f func()) {
	defer qt.Recovering("connect QInAppTransaction::finalize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "finalize", f)
	}
}

func (ptr *QInAppTransaction) DisconnectFinalize() {
	defer qt.Recovering("disconnect QInAppTransaction::finalize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "finalize")
	}
}

func (ptr *QInAppTransaction) Finalize() {
	defer qt.Recovering("QInAppTransaction::finalize")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_Finalize(ptr.Pointer())
	}
}

//export callbackQInAppTransaction_PlatformProperty
func callbackQInAppTransaction_PlatformProperty(ptr unsafe.Pointer, ptrName *C.char, propertyName *C.char) *C.char {
	defer qt.Recovering("callback QInAppTransaction::platformProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "platformProperty"); signal != nil {
		return C.CString(signal.(func(string) string)(C.GoString(propertyName)))
	}

	return C.CString(NewQInAppTransactionFromPointer(ptr).PlatformPropertyDefault(C.GoString(propertyName)))
}

func (ptr *QInAppTransaction) ConnectPlatformProperty(f func(propertyName string) string) {
	defer qt.Recovering("connect QInAppTransaction::platformProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "platformProperty", f)
	}
}

func (ptr *QInAppTransaction) DisconnectPlatformProperty() {
	defer qt.Recovering("disconnect QInAppTransaction::platformProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "platformProperty")
	}
}

func (ptr *QInAppTransaction) PlatformProperty(propertyName string) string {
	defer qt.Recovering("QInAppTransaction::platformProperty")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInAppTransaction_PlatformProperty(ptr.Pointer(), C.CString(propertyName)))
	}
	return ""
}

func (ptr *QInAppTransaction) PlatformPropertyDefault(propertyName string) string {
	defer qt.Recovering("QInAppTransaction::platformProperty")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInAppTransaction_PlatformPropertyDefault(ptr.Pointer(), C.CString(propertyName)))
	}
	return ""
}

//export callbackQInAppTransaction_TimerEvent
func callbackQInAppTransaction_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInAppTransaction::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInAppTransactionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInAppTransaction) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QInAppTransaction::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QInAppTransaction) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QInAppTransaction::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QInAppTransaction) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QInAppTransaction::timerEvent")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QInAppTransaction) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QInAppTransaction::timerEvent")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQInAppTransaction_ChildEvent
func callbackQInAppTransaction_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInAppTransaction::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInAppTransactionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInAppTransaction) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QInAppTransaction::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QInAppTransaction) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QInAppTransaction::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QInAppTransaction) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QInAppTransaction::childEvent")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QInAppTransaction) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QInAppTransaction::childEvent")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQInAppTransaction_ConnectNotify
func callbackQInAppTransaction_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QInAppTransaction::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppTransactionFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppTransaction) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QInAppTransaction::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QInAppTransaction) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QInAppTransaction::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QInAppTransaction) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppTransaction::connectNotify")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppTransaction) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppTransaction::connectNotify")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppTransaction_CustomEvent
func callbackQInAppTransaction_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInAppTransaction::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInAppTransactionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInAppTransaction) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QInAppTransaction::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QInAppTransaction) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QInAppTransaction::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QInAppTransaction) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QInAppTransaction::customEvent")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInAppTransaction) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QInAppTransaction::customEvent")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQInAppTransaction_DeleteLater
func callbackQInAppTransaction_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QInAppTransaction::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQInAppTransactionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QInAppTransaction) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QInAppTransaction::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QInAppTransaction) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QInAppTransaction::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QInAppTransaction) DeleteLater() {
	defer qt.Recovering("QInAppTransaction::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QInAppTransaction_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QInAppTransaction) DeleteLaterDefault() {
	defer qt.Recovering("QInAppTransaction::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QInAppTransaction_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQInAppTransaction_DisconnectNotify
func callbackQInAppTransaction_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QInAppTransaction::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppTransactionFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppTransaction) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QInAppTransaction::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QInAppTransaction) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QInAppTransaction::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QInAppTransaction) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppTransaction::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QInAppTransaction) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QInAppTransaction::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QInAppTransaction_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppTransaction_Event
func callbackQInAppTransaction_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QInAppTransaction::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQInAppTransactionFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QInAppTransaction) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QInAppTransaction::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QInAppTransaction) DisconnectEvent() {
	defer qt.Recovering("disconnect QInAppTransaction::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QInAppTransaction) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppTransaction::event")

	if ptr.Pointer() != nil {
		return C.QInAppTransaction_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QInAppTransaction) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppTransaction::event")

	if ptr.Pointer() != nil {
		return C.QInAppTransaction_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQInAppTransaction_EventFilter
func callbackQInAppTransaction_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QInAppTransaction::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQInAppTransactionFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QInAppTransaction) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QInAppTransaction::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QInAppTransaction) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QInAppTransaction::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QInAppTransaction) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppTransaction::eventFilter")

	if ptr.Pointer() != nil {
		return C.QInAppTransaction_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QInAppTransaction) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QInAppTransaction::eventFilter")

	if ptr.Pointer() != nil {
		return C.QInAppTransaction_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQInAppTransaction_MetaObject
func callbackQInAppTransaction_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QInAppTransaction::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQInAppTransactionFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QInAppTransaction) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QInAppTransaction::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QInAppTransaction) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QInAppTransaction::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QInAppTransaction) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QInAppTransaction::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppTransaction_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QInAppTransaction) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QInAppTransaction::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppTransaction_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
