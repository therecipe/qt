// +build !minimal

package felgo

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "felgo.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtFelgo_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtFelgo_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type FelgoApplication struct {
	core.QObject
}

type FelgoApplication_ITF interface {
	core.QObject_ITF
	FelgoApplication_PTR() *FelgoApplication
}

func (ptr *FelgoApplication) FelgoApplication_PTR() *FelgoApplication {
	return ptr
}

func (ptr *FelgoApplication) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *FelgoApplication) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromFelgoApplication(ptr FelgoApplication_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.FelgoApplication_PTR().Pointer()
	}
	return nil
}

func NewFelgoApplicationFromPointer(ptr unsafe.Pointer) (n *FelgoApplication) {
	n = new(FelgoApplication)
	n.SetPointer(ptr)
	return
}

//export callbackFelgoApplication_ObjectCreatedHandler
func callbackFelgoApplication_ObjectCreatedHandler(ptr unsafe.Pointer, object unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "objectCreatedHandler"); signal != nil {
		(*(*func(*core.QObject, *core.QUrl))(signal))(core.NewQObjectFromPointer(object), core.NewQUrlFromPointer(url))
	} else {
		NewFelgoApplicationFromPointer(ptr).ObjectCreatedHandlerDefault(core.NewQObjectFromPointer(object), core.NewQUrlFromPointer(url))
	}
}

func (ptr *FelgoApplication) ConnectObjectCreatedHandler(f func(object *core.QObject, url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "objectCreatedHandler"); signal != nil {
			f := func(object *core.QObject, url *core.QUrl) {
				(*(*func(*core.QObject, *core.QUrl))(signal))(object, url)
				f(object, url)
			}
			qt.ConnectSignal(ptr.Pointer(), "objectCreatedHandler", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "objectCreatedHandler", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoApplication) DisconnectObjectCreatedHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "objectCreatedHandler")
	}
}

func (ptr *FelgoApplication) ObjectCreatedHandler(object core.QObject_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_ObjectCreatedHandler(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQUrl(url))
	}
}

func (ptr *FelgoApplication) ObjectCreatedHandlerDefault(object core.QObject_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_ObjectCreatedHandlerDefault(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQUrl(url))
	}
}

func NewFelgoApplication(parent core.QObject_ITF) *FelgoApplication {
	tmpValue := NewFelgoApplicationFromPointer(C.FelgoApplication_NewFelgoApplication(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *FelgoApplication) AddCustomImportPaths() {
	if ptr.Pointer() != nil {
		C.FelgoApplication_AddCustomImportPaths(ptr.Pointer())
	}
}

func (ptr *FelgoApplication) AddImportPath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.FelgoApplication_AddImportPath(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *FelgoApplication) AdjustImportPathToCanonical(vqs string) string {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		return cGoUnpackString(C.FelgoApplication_AdjustImportPathToCanonical(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: vqsC, len: C.longlong(len(vqs))}))
	}
	return ""
}

func (ptr *FelgoApplication) ContentScale() float64 {
	if ptr.Pointer() != nil {
		return float64(C.FelgoApplication_ContentScale(ptr.Pointer()))
	}
	return 0
}

func (ptr *FelgoApplication) ContentScaleThresholds() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtFelgo_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewFelgoApplicationFromPointer(l.data)
			for i, v := range tmpList.__contentScaleThresholds_keyList() {
				out[v] = tmpList.__contentScaleThresholds_atList(v, i)
			}
			return out
		}(C.FelgoApplication_ContentScaleThresholds(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *FelgoApplication) FileSelectorList() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.FelgoApplication_FileSelectorList(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *FelgoApplication) Initialize(engine qml.QQmlEngine_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_Initialize(ptr.Pointer(), qml.PointerFromQQmlEngine(engine))
	}
}

func (ptr *FelgoApplication) InitializeEngine(engine qml.QQmlEngine_ITF, uri string) {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		C.FelgoApplication_InitializeEngine(ptr.Pointer(), qml.PointerFromQQmlEngine(engine), uriC)
	}
}

func (ptr *FelgoApplication) MainQmlFileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.FelgoApplication_MainQmlFileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *FelgoApplication) QmlEngine() *qml.QQmlEngine {
	if ptr.Pointer() != nil {
		tmpValue := qml.NewQQmlEngineFromPointer(C.FelgoApplication_QmlEngine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FelgoApplication) RegisterTypes(uri string) {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		C.FelgoApplication_RegisterTypes(ptr.Pointer(), uriC)
	}
}

func (ptr *FelgoApplication) ScaleFactorForImages() float64 {
	if ptr.Pointer() != nil {
		return float64(C.FelgoApplication_ScaleFactorForImages(ptr.Pointer()))
	}
	return 0
}

func (ptr *FelgoApplication) SetContentScaleAndFileSelectors(contentScale float64) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_SetContentScaleAndFileSelectors(ptr.Pointer(), C.double(contentScale))
	}
}

func (ptr *FelgoApplication) SetContentScaleThresholds(contentScaleThresholds map[string]*core.QVariant) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_SetContentScaleThresholds(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewFelgoApplicationFromPointer(NewFelgoApplicationFromPointer(nil).__setContentScaleThresholds_contentScaleThresholds_newList())
			for k, v := range contentScaleThresholds {
				tmpList.__setContentScaleThresholds_contentScaleThresholds_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *FelgoApplication) SetDefaultContentScalingAndFileSelectors() {
	if ptr.Pointer() != nil {
		C.FelgoApplication_SetDefaultContentScalingAndFileSelectors(ptr.Pointer())
	}
}

func (ptr *FelgoApplication) SetFileSelectorList(fileSelectorList []string) {
	if ptr.Pointer() != nil {
		fileSelectorListC := C.CString(strings.Join(fileSelectorList, "¡¦!"))
		defer C.free(unsafe.Pointer(fileSelectorListC))
		C.FelgoApplication_SetFileSelectorList(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: fileSelectorListC, len: C.longlong(len(strings.Join(fileSelectorList, "¡¦!")))})
	}
}

func (ptr *FelgoApplication) SetGameWindowIsItem() {
	if ptr.Pointer() != nil {
		C.FelgoApplication_SetGameWindowIsItem(ptr.Pointer())
	}
}

func (ptr *FelgoApplication) SetLicenseKey(licenseKey string) {
	if ptr.Pointer() != nil {
		var licenseKeyC *C.char
		if licenseKey != "" {
			licenseKeyC = C.CString(licenseKey)
			defer C.free(unsafe.Pointer(licenseKeyC))
		}
		C.FelgoApplication_SetLicenseKey(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: licenseKeyC, len: C.longlong(len(licenseKey))})
	}
}

func (ptr *FelgoApplication) SetMainQmlFileName(qmlFileName string) {
	if ptr.Pointer() != nil {
		var qmlFileNameC *C.char
		if qmlFileName != "" {
			qmlFileNameC = C.CString(qmlFileName)
			defer C.free(unsafe.Pointer(qmlFileNameC))
		}
		C.FelgoApplication_SetMainQmlFileName(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: qmlFileNameC, len: C.longlong(len(qmlFileName))})
	}
}

func (ptr *FelgoApplication) SetPreservePlatformFonts(preserve bool) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_SetPreservePlatformFonts(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(preserve))))
	}
}

func (ptr *FelgoApplication) VplayFileSelector() *core.QFileSelector {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQFileSelectorFromPointer(C.FelgoApplication_VplayFileSelector(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FelgoApplication) DestroyFelgoApplication() {
	if ptr.Pointer() != nil {
		C.FelgoApplication_DestroyFelgoApplication(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *FelgoApplication) __contentScaleThresholds_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.FelgoApplication___contentScaleThresholds_atList(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *FelgoApplication) __contentScaleThresholds_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.FelgoApplication___contentScaleThresholds_setList(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *FelgoApplication) __contentScaleThresholds_newList() unsafe.Pointer {
	return C.FelgoApplication___contentScaleThresholds_newList(ptr.Pointer())
}

func (ptr *FelgoApplication) __contentScaleThresholds_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtFelgo_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewFelgoApplicationFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____contentScaleThresholds_keyList_atList(i)
			}
			return out
		}(C.FelgoApplication___contentScaleThresholds_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *FelgoApplication) __setContentScaleThresholds_contentScaleThresholds_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.FelgoApplication___setContentScaleThresholds_contentScaleThresholds_atList(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *FelgoApplication) __setContentScaleThresholds_contentScaleThresholds_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.FelgoApplication___setContentScaleThresholds_contentScaleThresholds_setList(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *FelgoApplication) __setContentScaleThresholds_contentScaleThresholds_newList() unsafe.Pointer {
	return C.FelgoApplication___setContentScaleThresholds_contentScaleThresholds_newList(ptr.Pointer())
}

func (ptr *FelgoApplication) __setContentScaleThresholds_contentScaleThresholds_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtFelgo_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewFelgoApplicationFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setContentScaleThresholds_contentScaleThresholds_keyList_atList(i)
			}
			return out
		}(C.FelgoApplication___setContentScaleThresholds_contentScaleThresholds_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *FelgoApplication) ____contentScaleThresholds_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.FelgoApplication_____contentScaleThresholds_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *FelgoApplication) ____contentScaleThresholds_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.FelgoApplication_____contentScaleThresholds_keyList_setList(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *FelgoApplication) ____contentScaleThresholds_keyList_newList() unsafe.Pointer {
	return C.FelgoApplication_____contentScaleThresholds_keyList_newList(ptr.Pointer())
}

func (ptr *FelgoApplication) ____setContentScaleThresholds_contentScaleThresholds_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.FelgoApplication_____setContentScaleThresholds_contentScaleThresholds_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *FelgoApplication) ____setContentScaleThresholds_contentScaleThresholds_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.FelgoApplication_____setContentScaleThresholds_contentScaleThresholds_keyList_setList(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *FelgoApplication) ____setContentScaleThresholds_contentScaleThresholds_keyList_newList() unsafe.Pointer {
	return C.FelgoApplication_____setContentScaleThresholds_contentScaleThresholds_keyList_newList(ptr.Pointer())
}

func (ptr *FelgoApplication) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.FelgoApplication___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *FelgoApplication) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *FelgoApplication) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.FelgoApplication___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *FelgoApplication) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.FelgoApplication___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FelgoApplication) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *FelgoApplication) __findChildren_newList2() unsafe.Pointer {
	return C.FelgoApplication___findChildren_newList2(ptr.Pointer())
}

func (ptr *FelgoApplication) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.FelgoApplication___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FelgoApplication) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *FelgoApplication) __findChildren_newList3() unsafe.Pointer {
	return C.FelgoApplication___findChildren_newList3(ptr.Pointer())
}

func (ptr *FelgoApplication) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.FelgoApplication___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FelgoApplication) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *FelgoApplication) __findChildren_newList() unsafe.Pointer {
	return C.FelgoApplication___findChildren_newList(ptr.Pointer())
}

func (ptr *FelgoApplication) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.FelgoApplication___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FelgoApplication) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *FelgoApplication) __children_newList() unsafe.Pointer {
	return C.FelgoApplication___children_newList(ptr.Pointer())
}

//export callbackFelgoApplication_Event
func callbackFelgoApplication_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFelgoApplicationFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *FelgoApplication) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FelgoApplication_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackFelgoApplication_EventFilter
func callbackFelgoApplication_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFelgoApplicationFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *FelgoApplication) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FelgoApplication_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackFelgoApplication_ChildEvent
func callbackFelgoApplication_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewFelgoApplicationFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *FelgoApplication) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackFelgoApplication_ConnectNotify
func callbackFelgoApplication_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewFelgoApplicationFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *FelgoApplication) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackFelgoApplication_CustomEvent
func callbackFelgoApplication_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewFelgoApplicationFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *FelgoApplication) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackFelgoApplication_DeleteLater
func callbackFelgoApplication_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewFelgoApplicationFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *FelgoApplication) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.FelgoApplication_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackFelgoApplication_Destroyed
func callbackFelgoApplication_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackFelgoApplication_DisconnectNotify
func callbackFelgoApplication_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewFelgoApplicationFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *FelgoApplication) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackFelgoApplication_ObjectNameChanged
func callbackFelgoApplication_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackFelgoApplication_TimerEvent
func callbackFelgoApplication_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewFelgoApplicationFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *FelgoApplication) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoApplication_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackFelgoApplication_MetaObject
func callbackFelgoApplication_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewFelgoApplicationFromPointer(ptr).MetaObjectDefault())
}

func (ptr *FelgoApplication) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.FelgoApplication_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type FelgoLiveClient struct {
	core.QObject
}

type FelgoLiveClient_ITF interface {
	core.QObject_ITF
	FelgoLiveClient_PTR() *FelgoLiveClient
}

func (ptr *FelgoLiveClient) FelgoLiveClient_PTR() *FelgoLiveClient {
	return ptr
}

func (ptr *FelgoLiveClient) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *FelgoLiveClient) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromFelgoLiveClient(ptr FelgoLiveClient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.FelgoLiveClient_PTR().Pointer()
	}
	return nil
}

func NewFelgoLiveClientFromPointer(ptr unsafe.Pointer) (n *FelgoLiveClient) {
	n = new(FelgoLiveClient)
	n.SetPointer(ptr)
	return
}
func NewFelgoLiveClient(engine qml.QQmlEngine_ITF, parent core.QObject_ITF) *FelgoLiveClient {
	tmpValue := NewFelgoLiveClientFromPointer(C.FelgoLiveClient_NewFelgoLiveClient(qml.PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *FelgoLiveClient) AddShakeDetection() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_AddShakeDetection(ptr.Pointer())
	}
}

//export callbackFelgoLiveClient_ClearProject
func callbackFelgoLiveClient_ClearProject(ptr unsafe.Pointer, projectName C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "clearProject"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(projectName))
	} else {
		NewFelgoLiveClientFromPointer(ptr).ClearProjectDefault(cGoUnpackString(projectName))
	}
}

func (ptr *FelgoLiveClient) ConnectClearProject(f func(projectName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clearProject"); signal != nil {
			f := func(projectName string) {
				(*(*func(string))(signal))(projectName)
				f(projectName)
			}
			qt.ConnectSignal(ptr.Pointer(), "clearProject", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clearProject", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectClearProject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clearProject")
	}
}

func (ptr *FelgoLiveClient) ClearProject(projectName string) {
	if ptr.Pointer() != nil {
		var projectNameC *C.char
		if projectName != "" {
			projectNameC = C.CString(projectName)
			defer C.free(unsafe.Pointer(projectNameC))
		}
		C.FelgoLiveClient_ClearProject(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: projectNameC, len: C.longlong(len(projectName))})
	}
}

func (ptr *FelgoLiveClient) ClearProjectDefault(projectName string) {
	if ptr.Pointer() != nil {
		var projectNameC *C.char
		if projectName != "" {
			projectNameC = C.CString(projectName)
			defer C.free(unsafe.Pointer(projectNameC))
		}
		C.FelgoLiveClient_ClearProjectDefault(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: projectNameC, len: C.longlong(len(projectName))})
	}
}

//export callbackFelgoLiveClient_ClearProjects
func callbackFelgoLiveClient_ClearProjects(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearProjects"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewFelgoLiveClientFromPointer(ptr).ClearProjectsDefault()
	}
}

func (ptr *FelgoLiveClient) ConnectClearProjects(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clearProjects"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clearProjects", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clearProjects", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectClearProjects() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clearProjects")
	}
}

func (ptr *FelgoLiveClient) ClearProjects() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_ClearProjects(ptr.Pointer())
	}
}

func (ptr *FelgoLiveClient) ClearProjectsDefault() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_ClearProjectsDefault(ptr.Pointer())
	}
}

func (ptr *FelgoLiveClient) ClientName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.FelgoLiveClient_ClientName(ptr.Pointer()))
	}
	return ""
}

//export callbackFelgoLiveClient_ClientNameChanged
func callbackFelgoLiveClient_ClientNameChanged(ptr unsafe.Pointer, clientName C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "clientNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(clientName))
	}

}

func (ptr *FelgoLiveClient) ConnectClientNameChanged(f func(clientName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clientNameChanged") {
			C.FelgoLiveClient_ConnectClientNameChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clientNameChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clientNameChanged"); signal != nil {
			f := func(clientName string) {
				(*(*func(string))(signal))(clientName)
				f(clientName)
			}
			qt.ConnectSignal(ptr.Pointer(), "clientNameChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clientNameChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectClientNameChanged() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_DisconnectClientNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clientNameChanged")
	}
}

func (ptr *FelgoLiveClient) ClientNameChanged(clientName string) {
	if ptr.Pointer() != nil {
		var clientNameC *C.char
		if clientName != "" {
			clientNameC = C.CString(clientName)
			defer C.free(unsafe.Pointer(clientNameC))
		}
		C.FelgoLiveClient_ClientNameChanged(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: clientNameC, len: C.longlong(len(clientName))})
	}
}

//export callbackFelgoLiveClient_ConnectAsLocalDesktopClient
func callbackFelgoLiveClient_ConnectAsLocalDesktopClient(ptr unsafe.Pointer, serverUrl C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "connectAsLocalDesktopClient"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(serverUrl))
	} else {
		NewFelgoLiveClientFromPointer(ptr).ConnectAsLocalDesktopClientDefault(cGoUnpackString(serverUrl))
	}
}

func (ptr *FelgoLiveClient) ConnectConnectAsLocalDesktopClient(f func(serverUrl string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "connectAsLocalDesktopClient"); signal != nil {
			f := func(serverUrl string) {
				(*(*func(string))(signal))(serverUrl)
				f(serverUrl)
			}
			qt.ConnectSignal(ptr.Pointer(), "connectAsLocalDesktopClient", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectAsLocalDesktopClient", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectConnectAsLocalDesktopClient() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "connectAsLocalDesktopClient")
	}
}

func (ptr *FelgoLiveClient) ConnectAsLocalDesktopClient(serverUrl string) {
	if ptr.Pointer() != nil {
		var serverUrlC *C.char
		if serverUrl != "" {
			serverUrlC = C.CString(serverUrl)
			defer C.free(unsafe.Pointer(serverUrlC))
		}
		C.FelgoLiveClient_ConnectAsLocalDesktopClient(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: serverUrlC, len: C.longlong(len(serverUrl))})
	}
}

func (ptr *FelgoLiveClient) ConnectAsLocalDesktopClientDefault(serverUrl string) {
	if ptr.Pointer() != nil {
		var serverUrlC *C.char
		if serverUrl != "" {
			serverUrlC = C.CString(serverUrl)
			defer C.free(unsafe.Pointer(serverUrlC))
		}
		C.FelgoLiveClient_ConnectAsLocalDesktopClientDefault(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: serverUrlC, len: C.longlong(len(serverUrl))})
	}
}

//export callbackFelgoLiveClient_ConnectToLocalServer
func callbackFelgoLiveClient_ConnectToLocalServer(ptr unsafe.Pointer, serverUrl C.struct_QtFelgo_PackedString, userId C.struct_QtFelgo_PackedString, deviceId C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "connectToLocalServer"); signal != nil {
		(*(*func(string, string, string))(signal))(cGoUnpackString(serverUrl), cGoUnpackString(userId), cGoUnpackString(deviceId))
	} else {
		NewFelgoLiveClientFromPointer(ptr).ConnectToLocalServerDefault(cGoUnpackString(serverUrl), cGoUnpackString(userId), cGoUnpackString(deviceId))
	}
}

func (ptr *FelgoLiveClient) ConnectConnectToLocalServer(f func(serverUrl string, userId string, deviceId string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "connectToLocalServer"); signal != nil {
			f := func(serverUrl string, userId string, deviceId string) {
				(*(*func(string, string, string))(signal))(serverUrl, userId, deviceId)
				f(serverUrl, userId, deviceId)
			}
			qt.ConnectSignal(ptr.Pointer(), "connectToLocalServer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectToLocalServer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectConnectToLocalServer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "connectToLocalServer")
	}
}

func (ptr *FelgoLiveClient) ConnectToLocalServer(serverUrl string, userId string, deviceId string) {
	if ptr.Pointer() != nil {
		var serverUrlC *C.char
		if serverUrl != "" {
			serverUrlC = C.CString(serverUrl)
			defer C.free(unsafe.Pointer(serverUrlC))
		}
		var userIdC *C.char
		if userId != "" {
			userIdC = C.CString(userId)
			defer C.free(unsafe.Pointer(userIdC))
		}
		var deviceIdC *C.char
		if deviceId != "" {
			deviceIdC = C.CString(deviceId)
			defer C.free(unsafe.Pointer(deviceIdC))
		}
		C.FelgoLiveClient_ConnectToLocalServer(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: serverUrlC, len: C.longlong(len(serverUrl))}, C.struct_QtFelgo_PackedString{data: userIdC, len: C.longlong(len(userId))}, C.struct_QtFelgo_PackedString{data: deviceIdC, len: C.longlong(len(deviceId))})
	}
}

func (ptr *FelgoLiveClient) ConnectToLocalServerDefault(serverUrl string, userId string, deviceId string) {
	if ptr.Pointer() != nil {
		var serverUrlC *C.char
		if serverUrl != "" {
			serverUrlC = C.CString(serverUrl)
			defer C.free(unsafe.Pointer(serverUrlC))
		}
		var userIdC *C.char
		if userId != "" {
			userIdC = C.CString(userId)
			defer C.free(unsafe.Pointer(userIdC))
		}
		var deviceIdC *C.char
		if deviceId != "" {
			deviceIdC = C.CString(deviceId)
			defer C.free(unsafe.Pointer(deviceIdC))
		}
		C.FelgoLiveClient_ConnectToLocalServerDefault(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: serverUrlC, len: C.longlong(len(serverUrl))}, C.struct_QtFelgo_PackedString{data: userIdC, len: C.longlong(len(userId))}, C.struct_QtFelgo_PackedString{data: deviceIdC, len: C.longlong(len(deviceId))})
	}
}

func (ptr *FelgoLiveClient) ConnectToWebServer(userId string, deviceId string) {
	if ptr.Pointer() != nil {
		var userIdC *C.char
		if userId != "" {
			userIdC = C.CString(userId)
			defer C.free(unsafe.Pointer(userIdC))
		}
		var deviceIdC *C.char
		if deviceId != "" {
			deviceIdC = C.CString(deviceId)
			defer C.free(unsafe.Pointer(deviceIdC))
		}
		C.FelgoLiveClient_ConnectToWebServer(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: userIdC, len: C.longlong(len(userId))}, C.struct_QtFelgo_PackedString{data: deviceIdC, len: C.longlong(len(deviceId))})
	}
}

func (ptr *FelgoLiveClient) LoadProject(projectName string, projectMainFile string) {
	if ptr.Pointer() != nil {
		var projectNameC *C.char
		if projectName != "" {
			projectNameC = C.CString(projectName)
			defer C.free(unsafe.Pointer(projectNameC))
		}
		var projectMainFileC *C.char
		if projectMainFile != "" {
			projectMainFileC = C.CString(projectMainFile)
			defer C.free(unsafe.Pointer(projectMainFileC))
		}
		C.FelgoLiveClient_LoadProject(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: projectNameC, len: C.longlong(len(projectName))}, C.struct_QtFelgo_PackedString{data: projectMainFileC, len: C.longlong(len(projectMainFile))})
	}
}

func (ptr *FelgoLiveClient) DisableDisplaySleep() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_DisableDisplaySleep(ptr.Pointer())
	}
}

//export callbackFelgoLiveClient_DisconnectWebReceiver
func callbackFelgoLiveClient_DisconnectWebReceiver(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectWebReceiver"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewFelgoLiveClientFromPointer(ptr).DisconnectWebReceiverDefault()
	}
}

func (ptr *FelgoLiveClient) ConnectDisconnectWebReceiver(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "disconnectWebReceiver"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "disconnectWebReceiver", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnectWebReceiver", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectDisconnectWebReceiver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "disconnectWebReceiver")
	}
}

func (ptr *FelgoLiveClient) DisconnectWebReceiver() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_DisconnectWebReceiver(ptr.Pointer())
	}
}

func (ptr *FelgoLiveClient) DisconnectWebReceiverDefault() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_DisconnectWebReceiverDefault(ptr.Pointer())
	}
}

//export callbackFelgoLiveClient_ModifyLoadedDocument
func callbackFelgoLiveClient_ModifyLoadedDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modifyLoadedDocument"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewFelgoLiveClientFromPointer(ptr).ModifyLoadedDocumentDefault()
	}
}

func (ptr *FelgoLiveClient) ConnectModifyLoadedDocument(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "modifyLoadedDocument"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "modifyLoadedDocument", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "modifyLoadedDocument", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectModifyLoadedDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "modifyLoadedDocument")
	}
}

func (ptr *FelgoLiveClient) ModifyLoadedDocument() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_ModifyLoadedDocument(ptr.Pointer())
	}
}

func (ptr *FelgoLiveClient) ModifyLoadedDocumentDefault() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_ModifyLoadedDocumentDefault(ptr.Pointer())
	}
}

//export callbackFelgoLiveClient_OnDocumentUpdated
func callbackFelgoLiveClient_OnDocumentUpdated(ptr unsafe.Pointer, document C.struct_QtFelgo_PackedString, content C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "onDocumentUpdated"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(document), cGoUnpackString(content))
	} else {
		NewFelgoLiveClientFromPointer(ptr).OnDocumentUpdatedDefault(cGoUnpackString(document), cGoUnpackString(content))
	}
}

func (ptr *FelgoLiveClient) ConnectOnDocumentUpdated(f func(document string, content string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "onDocumentUpdated"); signal != nil {
			f := func(document string, content string) {
				(*(*func(string, string))(signal))(document, content)
				f(document, content)
			}
			qt.ConnectSignal(ptr.Pointer(), "onDocumentUpdated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onDocumentUpdated", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectOnDocumentUpdated() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "onDocumentUpdated")
	}
}

func (ptr *FelgoLiveClient) OnDocumentUpdated(document string, content string) {
	if ptr.Pointer() != nil {
		var documentC *C.char
		if document != "" {
			documentC = C.CString(document)
			defer C.free(unsafe.Pointer(documentC))
		}
		var contentC *C.char
		if content != "" {
			contentC = C.CString(content)
			defer C.free(unsafe.Pointer(contentC))
		}
		C.FelgoLiveClient_OnDocumentUpdated(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: documentC, len: C.longlong(len(document))}, C.struct_QtFelgo_PackedString{data: contentC, len: C.longlong(len(content))})
	}
}

func (ptr *FelgoLiveClient) OnDocumentUpdatedDefault(document string, content string) {
	if ptr.Pointer() != nil {
		var documentC *C.char
		if document != "" {
			documentC = C.CString(document)
			defer C.free(unsafe.Pointer(documentC))
		}
		var contentC *C.char
		if content != "" {
			contentC = C.CString(content)
			defer C.free(unsafe.Pointer(contentC))
		}
		C.FelgoLiveClient_OnDocumentUpdatedDefault(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: documentC, len: C.longlong(len(document))}, C.struct_QtFelgo_PackedString{data: contentC, len: C.longlong(len(content))})
	}
}

//export callbackFelgoLiveClient_OnProjectChanged
func callbackFelgoLiveClient_OnProjectChanged(ptr unsafe.Pointer, projectName C.struct_QtFelgo_PackedString, projectMainFile C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "onProjectChanged"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(projectName), cGoUnpackString(projectMainFile))
	} else {
		NewFelgoLiveClientFromPointer(ptr).OnProjectChangedDefault(cGoUnpackString(projectName), cGoUnpackString(projectMainFile))
	}
}

func (ptr *FelgoLiveClient) ConnectOnProjectChanged(f func(projectName string, projectMainFile string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "onProjectChanged"); signal != nil {
			f := func(projectName string, projectMainFile string) {
				(*(*func(string, string))(signal))(projectName, projectMainFile)
				f(projectName, projectMainFile)
			}
			qt.ConnectSignal(ptr.Pointer(), "onProjectChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onProjectChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectOnProjectChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "onProjectChanged")
	}
}

func (ptr *FelgoLiveClient) OnProjectChanged(projectName string, projectMainFile string) {
	if ptr.Pointer() != nil {
		var projectNameC *C.char
		if projectName != "" {
			projectNameC = C.CString(projectName)
			defer C.free(unsafe.Pointer(projectNameC))
		}
		var projectMainFileC *C.char
		if projectMainFile != "" {
			projectMainFileC = C.CString(projectMainFile)
			defer C.free(unsafe.Pointer(projectMainFileC))
		}
		C.FelgoLiveClient_OnProjectChanged(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: projectNameC, len: C.longlong(len(projectName))}, C.struct_QtFelgo_PackedString{data: projectMainFileC, len: C.longlong(len(projectMainFile))})
	}
}

func (ptr *FelgoLiveClient) OnProjectChangedDefault(projectName string, projectMainFile string) {
	if ptr.Pointer() != nil {
		var projectNameC *C.char
		if projectName != "" {
			projectNameC = C.CString(projectName)
			defer C.free(unsafe.Pointer(projectNameC))
		}
		var projectMainFileC *C.char
		if projectMainFile != "" {
			projectMainFileC = C.CString(projectMainFile)
			defer C.free(unsafe.Pointer(projectMainFileC))
		}
		C.FelgoLiveClient_OnProjectChangedDefault(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: projectNameC, len: C.longlong(len(projectName))}, C.struct_QtFelgo_PackedString{data: projectMainFileC, len: C.longlong(len(projectMainFile))})
	}
}

//export callbackFelgoLiveClient_PendingProject
func callbackFelgoLiveClient_PendingProject(ptr unsafe.Pointer, projectName C.struct_QtFelgo_PackedString, projectMainFile C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "pendingProject"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(projectName), cGoUnpackString(projectMainFile))
	}

}

func (ptr *FelgoLiveClient) ConnectPendingProject(f func(projectName string, projectMainFile string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pendingProject") {
			C.FelgoLiveClient_ConnectPendingProject(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pendingProject")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pendingProject"); signal != nil {
			f := func(projectName string, projectMainFile string) {
				(*(*func(string, string))(signal))(projectName, projectMainFile)
				f(projectName, projectMainFile)
			}
			qt.ConnectSignal(ptr.Pointer(), "pendingProject", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pendingProject", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectPendingProject() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_DisconnectPendingProject(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pendingProject")
	}
}

func (ptr *FelgoLiveClient) PendingProject(projectName string, projectMainFile string) {
	if ptr.Pointer() != nil {
		var projectNameC *C.char
		if projectName != "" {
			projectNameC = C.CString(projectName)
			defer C.free(unsafe.Pointer(projectNameC))
		}
		var projectMainFileC *C.char
		if projectMainFile != "" {
			projectMainFileC = C.CString(projectMainFile)
			defer C.free(unsafe.Pointer(projectMainFileC))
		}
		C.FelgoLiveClient_PendingProject(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: projectNameC, len: C.longlong(len(projectName))}, C.struct_QtFelgo_PackedString{data: projectMainFileC, len: C.longlong(len(projectMainFile))})
	}
}

//export callbackFelgoLiveClient_PrepareLoadingDocument
func callbackFelgoLiveClient_PrepareLoadingDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "prepareLoadingDocument"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewFelgoLiveClientFromPointer(ptr).PrepareLoadingDocumentDefault()
	}
}

func (ptr *FelgoLiveClient) ConnectPrepareLoadingDocument(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "prepareLoadingDocument"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "prepareLoadingDocument", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "prepareLoadingDocument", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectPrepareLoadingDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "prepareLoadingDocument")
	}
}

func (ptr *FelgoLiveClient) PrepareLoadingDocument() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_PrepareLoadingDocument(ptr.Pointer())
	}
}

func (ptr *FelgoLiveClient) PrepareLoadingDocumentDefault() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_PrepareLoadingDocumentDefault(ptr.Pointer())
	}
}

//export callbackFelgoLiveClient_WebReceiverConnectionRefused
func callbackFelgoLiveClient_WebReceiverConnectionRefused(ptr unsafe.Pointer, reason C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "webReceiverConnectionRefused"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(reason))
	}

}

func (ptr *FelgoLiveClient) ConnectWebReceiverConnectionRefused(f func(reason string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "webReceiverConnectionRefused") {
			C.FelgoLiveClient_ConnectWebReceiverConnectionRefused(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "webReceiverConnectionRefused")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "webReceiverConnectionRefused"); signal != nil {
			f := func(reason string) {
				(*(*func(string))(signal))(reason)
				f(reason)
			}
			qt.ConnectSignal(ptr.Pointer(), "webReceiverConnectionRefused", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "webReceiverConnectionRefused", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectWebReceiverConnectionRefused() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_DisconnectWebReceiverConnectionRefused(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "webReceiverConnectionRefused")
	}
}

func (ptr *FelgoLiveClient) WebReceiverConnectionRefused(reason string) {
	if ptr.Pointer() != nil {
		var reasonC *C.char
		if reason != "" {
			reasonC = C.CString(reason)
			defer C.free(unsafe.Pointer(reasonC))
		}
		C.FelgoLiveClient_WebReceiverConnectionRefused(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: reasonC, len: C.longlong(len(reason))})
	}
}

//export callbackFelgoLiveClient_ReceivedMatchId
func callbackFelgoLiveClient_ReceivedMatchId(ptr unsafe.Pointer, matchId C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "receivedMatchId"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(matchId))
	}

}

func (ptr *FelgoLiveClient) ConnectReceivedMatchId(f func(matchId string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "receivedMatchId") {
			C.FelgoLiveClient_ConnectReceivedMatchId(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "receivedMatchId")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "receivedMatchId"); signal != nil {
			f := func(matchId string) {
				(*(*func(string))(signal))(matchId)
				f(matchId)
			}
			qt.ConnectSignal(ptr.Pointer(), "receivedMatchId", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "receivedMatchId", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectReceivedMatchId() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_DisconnectReceivedMatchId(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "receivedMatchId")
	}
}

func (ptr *FelgoLiveClient) ReceivedMatchId(matchId string) {
	if ptr.Pointer() != nil {
		var matchIdC *C.char
		if matchId != "" {
			matchIdC = C.CString(matchId)
			defer C.free(unsafe.Pointer(matchIdC))
		}
		C.FelgoLiveClient_ReceivedMatchId(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: matchIdC, len: C.longlong(len(matchId))})
	}
}

func (ptr *FelgoLiveClient) SetClientName(clientName string) {
	if ptr.Pointer() != nil {
		var clientNameC *C.char
		if clientName != "" {
			clientNameC = C.CString(clientName)
			defer C.free(unsafe.Pointer(clientNameC))
		}
		C.FelgoLiveClient_SetClientName(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: clientNameC, len: C.longlong(len(clientName))})
	}
}

func (ptr *FelgoLiveClient) SetQmlEngine(engine qml.QQmlEngine_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_SetQmlEngine(ptr.Pointer(), qml.PointerFromQQmlEngine(engine))
	}
}

func (ptr *FelgoLiveClient) SetWelcomeQmlFile(file string) {
	if ptr.Pointer() != nil {
		var fileC *C.char
		if file != "" {
			fileC = C.CString(file)
			defer C.free(unsafe.Pointer(fileC))
		}
		C.FelgoLiveClient_SetWelcomeQmlFile(ptr.Pointer(), C.struct_QtFelgo_PackedString{data: fileC, len: C.longlong(len(file))})
	}
}

//export callbackFelgoLiveClient_ShakeDetected
func callbackFelgoLiveClient_ShakeDetected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "shakeDetected"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewFelgoLiveClientFromPointer(ptr).ShakeDetectedDefault()
	}
}

func (ptr *FelgoLiveClient) ConnectShakeDetected(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "shakeDetected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "shakeDetected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shakeDetected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectShakeDetected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "shakeDetected")
	}
}

func (ptr *FelgoLiveClient) ShakeDetected() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_ShakeDetected(ptr.Pointer())
	}
}

func (ptr *FelgoLiveClient) ShakeDetectedDefault() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_ShakeDetectedDefault(ptr.Pointer())
	}
}

//export callbackFelgoLiveClient_WebReceiverConnected
func callbackFelgoLiveClient_WebReceiverConnected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "webReceiverConnected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *FelgoLiveClient) ConnectWebReceiverConnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "webReceiverConnected") {
			C.FelgoLiveClient_ConnectWebReceiverConnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "webReceiverConnected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "webReceiverConnected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "webReceiverConnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "webReceiverConnected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *FelgoLiveClient) DisconnectWebReceiverConnected() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_DisconnectWebReceiverConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "webReceiverConnected")
	}
}

func (ptr *FelgoLiveClient) WebReceiverConnected() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_WebReceiverConnected(ptr.Pointer())
	}
}

func (ptr *FelgoLiveClient) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.FelgoLiveClient___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *FelgoLiveClient) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *FelgoLiveClient) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.FelgoLiveClient___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *FelgoLiveClient) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.FelgoLiveClient___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FelgoLiveClient) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *FelgoLiveClient) __findChildren_newList2() unsafe.Pointer {
	return C.FelgoLiveClient___findChildren_newList2(ptr.Pointer())
}

func (ptr *FelgoLiveClient) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.FelgoLiveClient___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FelgoLiveClient) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *FelgoLiveClient) __findChildren_newList3() unsafe.Pointer {
	return C.FelgoLiveClient___findChildren_newList3(ptr.Pointer())
}

func (ptr *FelgoLiveClient) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.FelgoLiveClient___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FelgoLiveClient) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *FelgoLiveClient) __findChildren_newList() unsafe.Pointer {
	return C.FelgoLiveClient___findChildren_newList(ptr.Pointer())
}

func (ptr *FelgoLiveClient) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.FelgoLiveClient___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *FelgoLiveClient) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *FelgoLiveClient) __children_newList() unsafe.Pointer {
	return C.FelgoLiveClient___children_newList(ptr.Pointer())
}

//export callbackFelgoLiveClient_Event
func callbackFelgoLiveClient_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFelgoLiveClientFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *FelgoLiveClient) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FelgoLiveClient_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackFelgoLiveClient_EventFilter
func callbackFelgoLiveClient_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewFelgoLiveClientFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *FelgoLiveClient) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.FelgoLiveClient_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackFelgoLiveClient_ChildEvent
func callbackFelgoLiveClient_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewFelgoLiveClientFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *FelgoLiveClient) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackFelgoLiveClient_ConnectNotify
func callbackFelgoLiveClient_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewFelgoLiveClientFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *FelgoLiveClient) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackFelgoLiveClient_CustomEvent
func callbackFelgoLiveClient_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewFelgoLiveClientFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *FelgoLiveClient) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackFelgoLiveClient_DeleteLater
func callbackFelgoLiveClient_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewFelgoLiveClientFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *FelgoLiveClient) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackFelgoLiveClient_Destroyed
func callbackFelgoLiveClient_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackFelgoLiveClient_DisconnectNotify
func callbackFelgoLiveClient_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewFelgoLiveClientFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *FelgoLiveClient) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackFelgoLiveClient_ObjectNameChanged
func callbackFelgoLiveClient_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtFelgo_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackFelgoLiveClient_TimerEvent
func callbackFelgoLiveClient_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewFelgoLiveClientFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *FelgoLiveClient) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.FelgoLiveClient_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackFelgoLiveClient_MetaObject
func callbackFelgoLiveClient_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewFelgoLiveClientFromPointer(ptr).MetaObjectDefault())
}

func (ptr *FelgoLiveClient) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.FelgoLiveClient_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func init() {
	qt.ItfMap["felgo.FelgoApplication_ITF"] = FelgoApplication{}
	qt.FuncMap["felgo.NewFelgoApplication"] = NewFelgoApplication
	qt.ItfMap["felgo.FelgoLiveClient_ITF"] = FelgoLiveClient{}
	qt.FuncMap["felgo.NewFelgoLiveClient"] = NewFelgoLiveClient
}
