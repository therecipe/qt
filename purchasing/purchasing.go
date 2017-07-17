// +build !minimal

package purchasing

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "purchasing.h"
import "C"
import (
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

//go:generate stringer -type=QInAppProduct__ProductType
//QInAppProduct::ProductType
type QInAppProduct__ProductType int64

const (
	QInAppProduct__Consumable QInAppProduct__ProductType = QInAppProduct__ProductType(0)
	QInAppProduct__Unlockable QInAppProduct__ProductType = QInAppProduct__ProductType(1)
)

//export callbackQInAppProduct_Purchase
func callbackQInAppProduct_Purchase(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "purchase"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInAppProduct) ConnectPurchase(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "purchase"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "purchase", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "purchase", f)
		}
	}
}

func (ptr *QInAppProduct) DisconnectPurchase() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "purchase")
	}
}

func (ptr *QInAppProduct) Purchase() {
	if ptr.Pointer() != nil {
		C.QInAppProduct_Purchase(ptr.Pointer())
	}
}

func (ptr *QInAppProduct) ProductType() QInAppProduct__ProductType {
	if ptr.Pointer() != nil {
		return QInAppProduct__ProductType(C.QInAppProduct_ProductType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInAppProduct) Identifier() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppProduct_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppProduct) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppProduct_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppProduct) Price() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppProduct_Price(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppProduct) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QInAppProduct_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInAppProduct) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QInAppProduct___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QInAppProduct) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QInAppProduct) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppProduct___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QInAppProduct) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppProduct___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppProduct) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppProduct) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppProduct___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QInAppProduct) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppProduct___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppProduct) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppProduct) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppProduct___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QInAppProduct) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppProduct___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppProduct) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppProduct) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppProduct___findChildren_newList(ptr.Pointer()))
}

func (ptr *QInAppProduct) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppProduct___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppProduct) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppProduct) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppProduct___children_newList(ptr.Pointer()))
}

//export callbackQInAppProduct_Event
func callbackQInAppProduct_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppProductFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QInAppProduct) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppProduct_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQInAppProduct_EventFilter
func callbackQInAppProduct_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppProductFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QInAppProduct) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppProduct_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQInAppProduct_ChildEvent
func callbackQInAppProduct_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInAppProductFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInAppProduct) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQInAppProduct_ConnectNotify
func callbackQInAppProduct_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppProductFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppProduct) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppProduct_CustomEvent
func callbackQInAppProduct_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInAppProductFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInAppProduct) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQInAppProduct_DeleteLater
func callbackQInAppProduct_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQInAppProductFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QInAppProduct) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QInAppProduct_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQInAppProduct_Destroyed
func callbackQInAppProduct_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQInAppProduct_DisconnectNotify
func callbackQInAppProduct_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppProductFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppProduct) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppProduct_ObjectNameChanged
func callbackQInAppProduct_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPurchasing_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQInAppProduct_TimerEvent
func callbackQInAppProduct_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInAppProductFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInAppProduct) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppProduct_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQInAppProduct_MetaObject
func callbackQInAppProduct_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQInAppProductFromPointer(ptr).MetaObjectDefault())
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
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQInAppStore_ProductRegistered
func callbackQInAppStore_ProductRegistered(ptr unsafe.Pointer, product unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "productRegistered"); signal != nil {
		signal.(func(*QInAppProduct))(NewQInAppProductFromPointer(product))
	}

}

func (ptr *QInAppStore) ConnectProductRegistered(f func(product *QInAppProduct)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "productRegistered") {
			C.QInAppStore_ConnectProductRegistered(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "productRegistered"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "productRegistered", func(product *QInAppProduct) {
				signal.(func(*QInAppProduct))(product)
				f(product)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "productRegistered", f)
		}
	}
}

func (ptr *QInAppStore) DisconnectProductRegistered() {
	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectProductRegistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "productRegistered")
	}
}

func (ptr *QInAppStore) ProductRegistered(product QInAppProduct_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ProductRegistered(ptr.Pointer(), PointerFromQInAppProduct(product))
	}
}

//export callbackQInAppStore_ProductUnknown
func callbackQInAppStore_ProductUnknown(ptr unsafe.Pointer, productType C.longlong, identifier C.struct_QtPurchasing_PackedString) {
	if signal := qt.GetSignal(ptr, "productUnknown"); signal != nil {
		signal.(func(QInAppProduct__ProductType, string))(QInAppProduct__ProductType(productType), cGoUnpackString(identifier))
	}

}

func (ptr *QInAppStore) ConnectProductUnknown(f func(productType QInAppProduct__ProductType, identifier string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "productUnknown") {
			C.QInAppStore_ConnectProductUnknown(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "productUnknown"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "productUnknown", func(productType QInAppProduct__ProductType, identifier string) {
				signal.(func(QInAppProduct__ProductType, string))(productType, identifier)
				f(productType, identifier)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "productUnknown", f)
		}
	}
}

func (ptr *QInAppStore) DisconnectProductUnknown() {
	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectProductUnknown(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "productUnknown")
	}
}

func (ptr *QInAppStore) ProductUnknown(productType QInAppProduct__ProductType, identifier string) {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		C.QInAppStore_ProductUnknown(ptr.Pointer(), C.longlong(productType), C.struct_QtPurchasing_PackedString{data: identifierC, len: C.longlong(len(identifier))})
	}
}

func (ptr *QInAppStore) RegisterProduct(productType QInAppProduct__ProductType, identifier string) {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		C.QInAppStore_RegisterProduct(ptr.Pointer(), C.longlong(productType), C.struct_QtPurchasing_PackedString{data: identifierC, len: C.longlong(len(identifier))})
	}
}

func (ptr *QInAppStore) RestorePurchases() {
	if ptr.Pointer() != nil {
		C.QInAppStore_RestorePurchases(ptr.Pointer())
	}
}

func (ptr *QInAppStore) SetPlatformProperty(propertyName string, value string) {
	if ptr.Pointer() != nil {
		var propertyNameC *C.char
		if propertyName != "" {
			propertyNameC = C.CString(propertyName)
			defer C.free(unsafe.Pointer(propertyNameC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QInAppStore_SetPlatformProperty(ptr.Pointer(), C.struct_QtPurchasing_PackedString{data: propertyNameC, len: C.longlong(len(propertyName))}, C.struct_QtPurchasing_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

//export callbackQInAppStore_TransactionReady
func callbackQInAppStore_TransactionReady(ptr unsafe.Pointer, transaction unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "transactionReady"); signal != nil {
		signal.(func(*QInAppTransaction))(NewQInAppTransactionFromPointer(transaction))
	}

}

func (ptr *QInAppStore) ConnectTransactionReady(f func(transaction *QInAppTransaction)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transactionReady") {
			C.QInAppStore_ConnectTransactionReady(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transactionReady"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transactionReady", func(transaction *QInAppTransaction) {
				signal.(func(*QInAppTransaction))(transaction)
				f(transaction)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transactionReady", f)
		}
	}
}

func (ptr *QInAppStore) DisconnectTransactionReady() {
	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectTransactionReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transactionReady")
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
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QInAppStore) RegisteredProduct(identifier string) *QInAppProduct {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		var tmpValue = NewQInAppProductFromPointer(C.QInAppStore_RegisteredProduct(ptr.Pointer(), C.struct_QtPurchasing_PackedString{data: identifierC, len: C.longlong(len(identifier))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppStore) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QInAppStore___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QInAppStore) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QInAppStore) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppStore___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QInAppStore) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppStore___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppStore) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppStore) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppStore___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QInAppStore) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppStore___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppStore) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppStore) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppStore___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QInAppStore) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppStore___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppStore) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppStore) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppStore___findChildren_newList(ptr.Pointer()))
}

func (ptr *QInAppStore) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppStore___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppStore) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppStore) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppStore___children_newList(ptr.Pointer()))
}

//export callbackQInAppStore_Event
func callbackQInAppStore_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppStoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QInAppStore) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppStore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQInAppStore_EventFilter
func callbackQInAppStore_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppStoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QInAppStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppStore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQInAppStore_ChildEvent
func callbackQInAppStore_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInAppStoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInAppStore) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQInAppStore_ConnectNotify
func callbackQInAppStore_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppStoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppStore_CustomEvent
func callbackQInAppStore_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInAppStoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInAppStore) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQInAppStore_DeleteLater
func callbackQInAppStore_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQInAppStoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QInAppStore) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QInAppStore_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQInAppStore_Destroyed
func callbackQInAppStore_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQInAppStore_DisconnectNotify
func callbackQInAppStore_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppStoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppStore_ObjectNameChanged
func callbackQInAppStore_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPurchasing_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQInAppStore_TimerEvent
func callbackQInAppStore_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInAppStoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInAppStore) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppStore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQInAppStore_MetaObject
func callbackQInAppStore_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQInAppStoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QInAppStore) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppStore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQInAppTransactionFromPointer(ptr unsafe.Pointer) *QInAppTransaction {
	var n = new(QInAppTransaction)
	n.SetPointer(ptr)
	return n
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

//export callbackQInAppTransaction_ErrorString
func callbackQInAppTransaction_ErrorString(ptr unsafe.Pointer) C.struct_QtPurchasing_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtPurchasing_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQInAppTransactionFromPointer(ptr).ErrorStringDefault()
	return C.struct_QtPurchasing_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QInAppTransaction) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorString", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", f)
		}
	}
}

func (ptr *QInAppTransaction) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorString")
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

//export callbackQInAppTransaction_OrderId
func callbackQInAppTransaction_OrderId(ptr unsafe.Pointer) C.struct_QtPurchasing_PackedString {
	if signal := qt.GetSignal(ptr, "orderId"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtPurchasing_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQInAppTransactionFromPointer(ptr).OrderIdDefault()
	return C.struct_QtPurchasing_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QInAppTransaction) ConnectOrderId(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "orderId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "orderId", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "orderId", f)
		}
	}
}

func (ptr *QInAppTransaction) DisconnectOrderId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "orderId")
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

//export callbackQInAppTransaction_Finalize
func callbackQInAppTransaction_Finalize(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finalize"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInAppTransaction) ConnectFinalize(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "finalize"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finalize", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finalize", f)
		}
	}
}

func (ptr *QInAppTransaction) DisconnectFinalize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "finalize")
	}
}

func (ptr *QInAppTransaction) Finalize() {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_Finalize(ptr.Pointer())
	}
}

//export callbackQInAppTransaction_FailureReason
func callbackQInAppTransaction_FailureReason(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "failureReason"); signal != nil {
		return C.longlong(signal.(func() QInAppTransaction__FailureReason)())
	}

	return C.longlong(NewQInAppTransactionFromPointer(ptr).FailureReasonDefault())
}

func (ptr *QInAppTransaction) ConnectFailureReason(f func() QInAppTransaction__FailureReason) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "failureReason"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "failureReason", func() QInAppTransaction__FailureReason {
				signal.(func() QInAppTransaction__FailureReason)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "failureReason", f)
		}
	}
}

func (ptr *QInAppTransaction) DisconnectFailureReason() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "failureReason")
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

//export callbackQInAppTransaction_Timestamp
func callbackQInAppTransaction_Timestamp(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "timestamp"); signal != nil {
		return core.PointerFromQDateTime(signal.(func() *core.QDateTime)())
	}

	return core.PointerFromQDateTime(NewQInAppTransactionFromPointer(ptr).TimestampDefault())
}

func (ptr *QInAppTransaction) ConnectTimestamp(f func() *core.QDateTime) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "timestamp"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "timestamp", func() *core.QDateTime {
				signal.(func() *core.QDateTime)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "timestamp", f)
		}
	}
}

func (ptr *QInAppTransaction) DisconnectTimestamp() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "timestamp")
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

func (ptr *QInAppTransaction) Product() *QInAppProduct {
	if ptr.Pointer() != nil {
		var tmpValue = NewQInAppProductFromPointer(C.QInAppTransaction_Product(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQInAppTransaction_PlatformProperty
func callbackQInAppTransaction_PlatformProperty(ptr unsafe.Pointer, propertyName C.struct_QtPurchasing_PackedString) C.struct_QtPurchasing_PackedString {
	if signal := qt.GetSignal(ptr, "platformProperty"); signal != nil {
		tempVal := signal.(func(string) string)(cGoUnpackString(propertyName))
		return C.struct_QtPurchasing_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQInAppTransactionFromPointer(ptr).PlatformPropertyDefault(cGoUnpackString(propertyName))
	return C.struct_QtPurchasing_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QInAppTransaction) ConnectPlatformProperty(f func(propertyName string) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "platformProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "platformProperty", func(propertyName string) string {
				signal.(func(string) string)(propertyName)
				return f(propertyName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "platformProperty", f)
		}
	}
}

func (ptr *QInAppTransaction) DisconnectPlatformProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "platformProperty")
	}
}

func (ptr *QInAppTransaction) PlatformProperty(propertyName string) string {
	if ptr.Pointer() != nil {
		var propertyNameC *C.char
		if propertyName != "" {
			propertyNameC = C.CString(propertyName)
			defer C.free(unsafe.Pointer(propertyNameC))
		}
		return cGoUnpackString(C.QInAppTransaction_PlatformProperty(ptr.Pointer(), C.struct_QtPurchasing_PackedString{data: propertyNameC, len: C.longlong(len(propertyName))}))
	}
	return ""
}

func (ptr *QInAppTransaction) PlatformPropertyDefault(propertyName string) string {
	if ptr.Pointer() != nil {
		var propertyNameC *C.char
		if propertyName != "" {
			propertyNameC = C.CString(propertyName)
			defer C.free(unsafe.Pointer(propertyNameC))
		}
		return cGoUnpackString(C.QInAppTransaction_PlatformPropertyDefault(ptr.Pointer(), C.struct_QtPurchasing_PackedString{data: propertyNameC, len: C.longlong(len(propertyName))}))
	}
	return ""
}

func (ptr *QInAppTransaction) Status() QInAppTransaction__TransactionStatus {
	if ptr.Pointer() != nil {
		return QInAppTransaction__TransactionStatus(C.QInAppTransaction_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInAppTransaction) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QInAppTransaction___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QInAppTransaction) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QInAppTransaction) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppTransaction___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QInAppTransaction) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppTransaction___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppTransaction) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppTransaction) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppTransaction___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QInAppTransaction) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppTransaction___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppTransaction) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppTransaction) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppTransaction___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QInAppTransaction) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppTransaction___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppTransaction) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppTransaction) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppTransaction___findChildren_newList(ptr.Pointer()))
}

func (ptr *QInAppTransaction) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QInAppTransaction___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QInAppTransaction) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QInAppTransaction) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QInAppTransaction___children_newList(ptr.Pointer()))
}

//export callbackQInAppTransaction_Event
func callbackQInAppTransaction_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppTransactionFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QInAppTransaction) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppTransaction_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQInAppTransaction_EventFilter
func callbackQInAppTransaction_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQInAppTransactionFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QInAppTransaction) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QInAppTransaction_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQInAppTransaction_ChildEvent
func callbackQInAppTransaction_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInAppTransactionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInAppTransaction) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQInAppTransaction_ConnectNotify
func callbackQInAppTransaction_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppTransactionFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppTransaction) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppTransaction_CustomEvent
func callbackQInAppTransaction_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInAppTransactionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInAppTransaction) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQInAppTransaction_DeleteLater
func callbackQInAppTransaction_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQInAppTransactionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QInAppTransaction) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQInAppTransaction_Destroyed
func callbackQInAppTransaction_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQInAppTransaction_DisconnectNotify
func callbackQInAppTransaction_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQInAppTransactionFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QInAppTransaction) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQInAppTransaction_ObjectNameChanged
func callbackQInAppTransaction_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPurchasing_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQInAppTransaction_TimerEvent
func callbackQInAppTransaction_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInAppTransactionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInAppTransaction) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QInAppTransaction_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQInAppTransaction_MetaObject
func callbackQInAppTransaction_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQInAppTransactionFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QInAppTransaction) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QInAppTransaction_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
