// +build !minimal

package macextras

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "macextras.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtMacExtras_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtMacExtras_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QMacPasteboardMime struct {
	ptr unsafe.Pointer
}

type QMacPasteboardMime_ITF interface {
	QMacPasteboardMime_PTR() *QMacPasteboardMime
}

func (ptr *QMacPasteboardMime) QMacPasteboardMime_PTR() *QMacPasteboardMime {
	return ptr
}

func (ptr *QMacPasteboardMime) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QMacPasteboardMime) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQMacPasteboardMime(ptr QMacPasteboardMime_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacPasteboardMime_PTR().Pointer()
	}
	return nil
}

func NewQMacPasteboardMimeFromPointer(ptr unsafe.Pointer) (n *QMacPasteboardMime) {
	n = new(QMacPasteboardMime)
	n.SetPointer(ptr)
	return
}

//export callbackQMacPasteboardMime_ConvertFromMime
func callbackQMacPasteboardMime_ConvertFromMime(ptr unsafe.Pointer, mime C.struct_QtMacExtras_PackedString, data unsafe.Pointer, flav C.struct_QtMacExtras_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "convertFromMime"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQMacPasteboardMimeFromPointer(NewQMacPasteboardMimeFromPointer(nil).__convertFromMime_newList())
			for _, v := range signal.(func(string, *core.QVariant, string) []*core.QByteArray)(cGoUnpackString(mime), core.NewQVariantFromPointer(data), cGoUnpackString(flav)) {
				tmpList.__convertFromMime_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQMacPasteboardMimeFromPointer(NewQMacPasteboardMimeFromPointer(nil).__convertFromMime_newList())
		for _, v := range make([]*core.QByteArray, 0) {
			tmpList.__convertFromMime_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QMacPasteboardMime) ConnectConvertFromMime(f func(mime string, data *core.QVariant, flav string) []*core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "convertFromMime"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "convertFromMime", func(mime string, data *core.QVariant, flav string) []*core.QByteArray {
				signal.(func(string, *core.QVariant, string) []*core.QByteArray)(mime, data, flav)
				return f(mime, data, flav)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "convertFromMime", f)
		}
	}
}

func (ptr *QMacPasteboardMime) DisconnectConvertFromMime() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "convertFromMime")
	}
}

func (ptr *QMacPasteboardMime) ConvertFromMime(mime string, data core.QVariant_ITF, flav string) []*core.QByteArray {
	if ptr.Pointer() != nil {
		var mimeC *C.char
		if mime != "" {
			mimeC = C.CString(mime)
			defer C.free(unsafe.Pointer(mimeC))
		}
		var flavC *C.char
		if flav != "" {
			flavC = C.CString(flav)
			defer C.free(unsafe.Pointer(flavC))
		}
		return func(l C.struct_QtMacExtras_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQMacPasteboardMimeFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__convertFromMime_atList(i)
			}
			return out
		}(C.QMacPasteboardMime_ConvertFromMime(ptr.Pointer(), C.struct_QtMacExtras_PackedString{data: mimeC, len: C.longlong(len(mime))}, core.PointerFromQVariant(data), C.struct_QtMacExtras_PackedString{data: flavC, len: C.longlong(len(flav))}))
	}
	return make([]*core.QByteArray, 0)
}

func NewQMacPasteboardMime(vch string) *QMacPasteboardMime {
	var vchC *C.char
	if vch != "" {
		vchC = C.CString(vch)
		defer C.free(unsafe.Pointer(vchC))
	}
	return NewQMacPasteboardMimeFromPointer(C.QMacPasteboardMime_NewQMacPasteboardMime(vchC))
}

//export callbackQMacPasteboardMime_ConvertorName
func callbackQMacPasteboardMime_ConvertorName(ptr unsafe.Pointer) C.struct_QtMacExtras_PackedString {
	if signal := qt.GetSignal(ptr, "convertorName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtMacExtras_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtMacExtras_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QMacPasteboardMime) ConnectConvertorName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "convertorName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "convertorName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "convertorName", f)
		}
	}
}

func (ptr *QMacPasteboardMime) DisconnectConvertorName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "convertorName")
	}
}

func (ptr *QMacPasteboardMime) ConvertorName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QMacPasteboardMime_ConvertorName(ptr.Pointer()))
	}
	return ""
}

//export callbackQMacPasteboardMime_FlavorFor
func callbackQMacPasteboardMime_FlavorFor(ptr unsafe.Pointer, mime C.struct_QtMacExtras_PackedString) C.struct_QtMacExtras_PackedString {
	if signal := qt.GetSignal(ptr, "flavorFor"); signal != nil {
		tempVal := signal.(func(string) string)(cGoUnpackString(mime))
		return C.struct_QtMacExtras_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtMacExtras_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QMacPasteboardMime) ConnectFlavorFor(f func(mime string) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "flavorFor"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "flavorFor", func(mime string) string {
				signal.(func(string) string)(mime)
				return f(mime)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "flavorFor", f)
		}
	}
}

func (ptr *QMacPasteboardMime) DisconnectFlavorFor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "flavorFor")
	}
}

func (ptr *QMacPasteboardMime) FlavorFor(mime string) string {
	if ptr.Pointer() != nil {
		var mimeC *C.char
		if mime != "" {
			mimeC = C.CString(mime)
			defer C.free(unsafe.Pointer(mimeC))
		}
		return cGoUnpackString(C.QMacPasteboardMime_FlavorFor(ptr.Pointer(), C.struct_QtMacExtras_PackedString{data: mimeC, len: C.longlong(len(mime))}))
	}
	return ""
}

//export callbackQMacPasteboardMime_MimeFor
func callbackQMacPasteboardMime_MimeFor(ptr unsafe.Pointer, flav C.struct_QtMacExtras_PackedString) C.struct_QtMacExtras_PackedString {
	if signal := qt.GetSignal(ptr, "mimeFor"); signal != nil {
		tempVal := signal.(func(string) string)(cGoUnpackString(flav))
		return C.struct_QtMacExtras_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtMacExtras_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QMacPasteboardMime) ConnectMimeFor(f func(flav string) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mimeFor"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mimeFor", func(flav string) string {
				signal.(func(string) string)(flav)
				return f(flav)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mimeFor", f)
		}
	}
}

func (ptr *QMacPasteboardMime) DisconnectMimeFor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mimeFor")
	}
}

func (ptr *QMacPasteboardMime) MimeFor(flav string) string {
	if ptr.Pointer() != nil {
		var flavC *C.char
		if flav != "" {
			flavC = C.CString(flav)
			defer C.free(unsafe.Pointer(flavC))
		}
		return cGoUnpackString(C.QMacPasteboardMime_MimeFor(ptr.Pointer(), C.struct_QtMacExtras_PackedString{data: flavC, len: C.longlong(len(flav))}))
	}
	return ""
}

//export callbackQMacPasteboardMime_ConvertToMime
func callbackQMacPasteboardMime_ConvertToMime(ptr unsafe.Pointer, mime C.struct_QtMacExtras_PackedString, data C.struct_QtMacExtras_PackedList, flav C.struct_QtMacExtras_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "convertToMime"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string, []*core.QByteArray, string) *core.QVariant)(cGoUnpackString(mime), func(l C.struct_QtMacExtras_PackedList) []*core.QByteArray {
			out := make([]*core.QByteArray, int(l.len))
			tmpList := NewQMacPasteboardMimeFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__convertToMime_data_atList(i)
			}
			return out
		}(data), cGoUnpackString(flav)))
	}

	return core.PointerFromQVariant(core.NewQVariant())
}

func (ptr *QMacPasteboardMime) ConnectConvertToMime(f func(mime string, data []*core.QByteArray, flav string) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "convertToMime"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "convertToMime", func(mime string, data []*core.QByteArray, flav string) *core.QVariant {
				signal.(func(string, []*core.QByteArray, string) *core.QVariant)(mime, data, flav)
				return f(mime, data, flav)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "convertToMime", f)
		}
	}
}

func (ptr *QMacPasteboardMime) DisconnectConvertToMime() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "convertToMime")
	}
}

func (ptr *QMacPasteboardMime) ConvertToMime(mime string, data []*core.QByteArray, flav string) *core.QVariant {
	if ptr.Pointer() != nil {
		var mimeC *C.char
		if mime != "" {
			mimeC = C.CString(mime)
			defer C.free(unsafe.Pointer(mimeC))
		}
		var flavC *C.char
		if flav != "" {
			flavC = C.CString(flav)
			defer C.free(unsafe.Pointer(flavC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QMacPasteboardMime_ConvertToMime(ptr.Pointer(), C.struct_QtMacExtras_PackedString{data: mimeC, len: C.longlong(len(mime))}, func() unsafe.Pointer {
			tmpList := NewQMacPasteboardMimeFromPointer(NewQMacPasteboardMimeFromPointer(nil).__convertToMime_data_newList())
			for _, v := range data {
				tmpList.__convertToMime_data_setList(v)
			}
			return tmpList.Pointer()
		}(), C.struct_QtMacExtras_PackedString{data: flavC, len: C.longlong(len(flav))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQMacPasteboardMime_CanConvert
func callbackQMacPasteboardMime_CanConvert(ptr unsafe.Pointer, mime C.struct_QtMacExtras_PackedString, flav C.struct_QtMacExtras_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "canConvert"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string) bool)(cGoUnpackString(mime), cGoUnpackString(flav)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QMacPasteboardMime) ConnectCanConvert(f func(mime string, flav string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "canConvert"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "canConvert", func(mime string, flav string) bool {
				signal.(func(string, string) bool)(mime, flav)
				return f(mime, flav)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "canConvert", f)
		}
	}
}

func (ptr *QMacPasteboardMime) DisconnectCanConvert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "canConvert")
	}
}

func (ptr *QMacPasteboardMime) CanConvert(mime string, flav string) bool {
	if ptr.Pointer() != nil {
		var mimeC *C.char
		if mime != "" {
			mimeC = C.CString(mime)
			defer C.free(unsafe.Pointer(mimeC))
		}
		var flavC *C.char
		if flav != "" {
			flavC = C.CString(flav)
			defer C.free(unsafe.Pointer(flavC))
		}
		return int8(C.QMacPasteboardMime_CanConvert(ptr.Pointer(), C.struct_QtMacExtras_PackedString{data: mimeC, len: C.longlong(len(mime))}, C.struct_QtMacExtras_PackedString{data: flavC, len: C.longlong(len(flav))})) != 0
	}
	return false
}

//export callbackQMacPasteboardMime_Count
func callbackQMacPasteboardMime_Count(ptr unsafe.Pointer, mimeData unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32(signal.(func(*core.QMimeData) int)(core.NewQMimeDataFromPointer(mimeData))))
	}

	return C.int(int32(NewQMacPasteboardMimeFromPointer(ptr).CountDefault(core.NewQMimeDataFromPointer(mimeData))))
}

func (ptr *QMacPasteboardMime) ConnectCount(f func(mimeData *core.QMimeData) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "count"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "count", func(mimeData *core.QMimeData) int {
				signal.(func(*core.QMimeData) int)(mimeData)
				return f(mimeData)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "count", f)
		}
	}
}

func (ptr *QMacPasteboardMime) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "count")
	}
}

func (ptr *QMacPasteboardMime) Count(mimeData core.QMimeData_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QMacPasteboardMime_Count(ptr.Pointer(), core.PointerFromQMimeData(mimeData))))
	}
	return 0
}

func (ptr *QMacPasteboardMime) CountDefault(mimeData core.QMimeData_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QMacPasteboardMime_CountDefault(ptr.Pointer(), core.PointerFromQMimeData(mimeData))))
	}
	return 0
}

//export callbackQMacPasteboardMime_DestroyQMacPasteboardMime
func callbackQMacPasteboardMime_DestroyQMacPasteboardMime(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QMacPasteboardMime"); signal != nil {
		signal.(func())()
	} else {
		NewQMacPasteboardMimeFromPointer(ptr).DestroyQMacPasteboardMimeDefault()
	}
}

func (ptr *QMacPasteboardMime) ConnectDestroyQMacPasteboardMime(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QMacPasteboardMime"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QMacPasteboardMime", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QMacPasteboardMime", f)
		}
	}
}

func (ptr *QMacPasteboardMime) DisconnectDestroyQMacPasteboardMime() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QMacPasteboardMime")
	}
}

func (ptr *QMacPasteboardMime) DestroyQMacPasteboardMime() {
	if ptr.Pointer() != nil {
		C.QMacPasteboardMime_DestroyQMacPasteboardMime(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QMacPasteboardMime) DestroyQMacPasteboardMimeDefault() {
	if ptr.Pointer() != nil {
		C.QMacPasteboardMime_DestroyQMacPasteboardMimeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QMacPasteboardMime) __convertFromMime_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QMacPasteboardMime___convertFromMime_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QMacPasteboardMime) __convertFromMime_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMacPasteboardMime___convertFromMime_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QMacPasteboardMime) __convertFromMime_newList() unsafe.Pointer {
	return C.QMacPasteboardMime___convertFromMime_newList(ptr.Pointer())
}

func (ptr *QMacPasteboardMime) __convertToMime_data_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QMacPasteboardMime___convertToMime_data_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QMacPasteboardMime) __convertToMime_data_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMacPasteboardMime___convertToMime_data_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QMacPasteboardMime) __convertToMime_data_newList() unsafe.Pointer {
	return C.QMacPasteboardMime___convertToMime_data_newList(ptr.Pointer())
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

func NewQMacToolBarFromPointer(ptr unsafe.Pointer) (n *QMacToolBar) {
	n = new(QMacToolBar)
	n.SetPointer(ptr)
	return
}
func (ptr *QMacToolBar) AllowedItems() []*QMacToolBarItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtMacExtras_PackedList) []*QMacToolBarItem {
			out := make([]*QMacToolBarItem, int(l.len))
			tmpList := NewQMacToolBarFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__allowedItems_atList(i)
			}
			return out
		}(C.QMacToolBar_AllowedItems(ptr.Pointer()))
	}
	return make([]*QMacToolBarItem, 0)
}

func (ptr *QMacToolBar) Items() []*QMacToolBarItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtMacExtras_PackedList) []*QMacToolBarItem {
			out := make([]*QMacToolBarItem, int(l.len))
			tmpList := NewQMacToolBarFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__items_atList(i)
			}
			return out
		}(C.QMacToolBar_Items(ptr.Pointer()))
	}
	return make([]*QMacToolBarItem, 0)
}

func NewQMacToolBar(parent core.QObject_ITF) *QMacToolBar {
	tmpValue := NewQMacToolBarFromPointer(C.QMacToolBar_NewQMacToolBar(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQMacToolBar2(identifier string, parent core.QObject_ITF) *QMacToolBar {
	var identifierC *C.char
	if identifier != "" {
		identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
	}
	tmpValue := NewQMacToolBarFromPointer(C.QMacToolBar_NewQMacToolBar2(C.struct_QtMacExtras_PackedString{data: identifierC, len: C.longlong(len(identifier))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QMacToolBar) AddAllowedItem(icon gui.QIcon_ITF, text string) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := NewQMacToolBarItemFromPointer(C.QMacToolBar_AddAllowedItem(ptr.Pointer(), gui.PointerFromQIcon(icon), C.struct_QtMacExtras_PackedString{data: textC, len: C.longlong(len(text))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) AddAllowedStandardItem(standardItem QMacToolBarItem__StandardItem) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQMacToolBarItemFromPointer(C.QMacToolBar_AddAllowedStandardItem(ptr.Pointer(), C.longlong(standardItem)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) AddItem(icon gui.QIcon_ITF, text string) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := NewQMacToolBarItemFromPointer(C.QMacToolBar_AddItem(ptr.Pointer(), gui.PointerFromQIcon(icon), C.struct_QtMacExtras_PackedString{data: textC, len: C.longlong(len(text))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) AddStandardItem(standardItem QMacToolBarItem__StandardItem) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQMacToolBarItemFromPointer(C.QMacToolBar_AddStandardItem(ptr.Pointer(), C.longlong(standardItem)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QMacToolBar_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QMacToolBar_QMacToolBar_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QMacToolBar) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QMacToolBar_QMacToolBar_Tr(sC, cC, C.int(int32(n))))
}

func QMacToolBar_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QMacToolBar_QMacToolBar_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QMacToolBar) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QMacToolBar_QMacToolBar_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QMacToolBar) AddSeparator() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_AddSeparator(ptr.Pointer())
	}
}

func (ptr *QMacToolBar) AttachToWindow(window gui.QWindow_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_AttachToWindow(ptr.Pointer(), gui.PointerFromQWindow(window))
	}
}

func (ptr *QMacToolBar) DetachFromWindow() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DetachFromWindow(ptr.Pointer())
	}
}

func (ptr *QMacToolBar) SetAllowedItems(allowedItems []*QMacToolBarItem) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_SetAllowedItems(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQMacToolBarFromPointer(NewQMacToolBarFromPointer(nil).__setAllowedItems_allowedItems_newList())
			for _, v := range allowedItems {
				tmpList.__setAllowedItems_allowedItems_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QMacToolBar) SetItems(items []*QMacToolBarItem) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_SetItems(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQMacToolBarFromPointer(NewQMacToolBarFromPointer(nil).__setItems_items_newList())
			for _, v := range items {
				tmpList.__setItems_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQMacToolBar_DestroyQMacToolBar
func callbackQMacToolBar_DestroyQMacToolBar(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QMacToolBar"); signal != nil {
		signal.(func())()
	} else {
		NewQMacToolBarFromPointer(ptr).DestroyQMacToolBarDefault()
	}
}

func (ptr *QMacToolBar) ConnectDestroyQMacToolBar(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QMacToolBar"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QMacToolBar", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QMacToolBar", f)
		}
	}
}

func (ptr *QMacToolBar) DisconnectDestroyQMacToolBar() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QMacToolBar")
	}
}

func (ptr *QMacToolBar) DestroyQMacToolBar() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DestroyQMacToolBar(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QMacToolBar) DestroyQMacToolBarDefault() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DestroyQMacToolBarDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQMacToolBar_MetaObject
func callbackQMacToolBar_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMacToolBarFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMacToolBar) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMacToolBar_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMacToolBar) __allowedItems_atList(i int) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQMacToolBarItemFromPointer(C.QMacToolBar___allowedItems_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) __allowedItems_setList(i QMacToolBarItem_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar___allowedItems_setList(ptr.Pointer(), PointerFromQMacToolBarItem(i))
	}
}

func (ptr *QMacToolBar) __allowedItems_newList() unsafe.Pointer {
	return C.QMacToolBar___allowedItems_newList(ptr.Pointer())
}

func (ptr *QMacToolBar) __items_atList(i int) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQMacToolBarItemFromPointer(C.QMacToolBar___items_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) __items_setList(i QMacToolBarItem_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar___items_setList(ptr.Pointer(), PointerFromQMacToolBarItem(i))
	}
}

func (ptr *QMacToolBar) __items_newList() unsafe.Pointer {
	return C.QMacToolBar___items_newList(ptr.Pointer())
}

func (ptr *QMacToolBar) __setAllowedItems_allowedItems_atList(i int) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQMacToolBarItemFromPointer(C.QMacToolBar___setAllowedItems_allowedItems_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) __setAllowedItems_allowedItems_setList(i QMacToolBarItem_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar___setAllowedItems_allowedItems_setList(ptr.Pointer(), PointerFromQMacToolBarItem(i))
	}
}

func (ptr *QMacToolBar) __setAllowedItems_allowedItems_newList() unsafe.Pointer {
	return C.QMacToolBar___setAllowedItems_allowedItems_newList(ptr.Pointer())
}

func (ptr *QMacToolBar) __setItems_items_atList(i int) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQMacToolBarItemFromPointer(C.QMacToolBar___setItems_items_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) __setItems_items_setList(i QMacToolBarItem_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar___setItems_items_setList(ptr.Pointer(), PointerFromQMacToolBarItem(i))
	}
}

func (ptr *QMacToolBar) __setItems_items_newList() unsafe.Pointer {
	return C.QMacToolBar___setItems_items_newList(ptr.Pointer())
}

func (ptr *QMacToolBar) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QMacToolBar___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QMacToolBar) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QMacToolBar___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QMacToolBar) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMacToolBar___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMacToolBar) __findChildren_newList2() unsafe.Pointer {
	return C.QMacToolBar___findChildren_newList2(ptr.Pointer())
}

func (ptr *QMacToolBar) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMacToolBar___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMacToolBar) __findChildren_newList3() unsafe.Pointer {
	return C.QMacToolBar___findChildren_newList3(ptr.Pointer())
}

func (ptr *QMacToolBar) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMacToolBar___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMacToolBar) __findChildren_newList() unsafe.Pointer {
	return C.QMacToolBar___findChildren_newList(ptr.Pointer())
}

func (ptr *QMacToolBar) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMacToolBar___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMacToolBar) __children_newList() unsafe.Pointer {
	return C.QMacToolBar___children_newList(ptr.Pointer())
}

//export callbackQMacToolBar_Event
func callbackQMacToolBar_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMacToolBar) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QMacToolBar_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQMacToolBar_EventFilter
func callbackQMacToolBar_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMacToolBar) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QMacToolBar_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQMacToolBar_ChildEvent
func callbackQMacToolBar_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMacToolBar_ConnectNotify
func callbackQMacToolBar_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBar) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBar_CustomEvent
func callbackQMacToolBar_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMacToolBar_DeleteLater
func callbackQMacToolBar_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQMacToolBarFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMacToolBar) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQMacToolBar_Destroyed
func callbackQMacToolBar_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQMacToolBar_DisconnectNotify
func callbackQMacToolBar_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBar) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBar_ObjectNameChanged
func callbackQMacToolBar_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtMacExtras_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQMacToolBar_TimerEvent
func callbackQMacToolBar_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQMacToolBarItemFromPointer(ptr unsafe.Pointer) (n *QMacToolBarItem) {
	n = new(QMacToolBarItem)
	n.SetPointer(ptr)
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
	tmpValue := NewQMacToolBarItemFromPointer(C.QMacToolBarItem_NewQMacToolBarItem(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QMacToolBarItem_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QMacToolBarItem_QMacToolBarItem_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QMacToolBarItem) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QMacToolBarItem_QMacToolBarItem_Tr(sC, cC, C.int(int32(n))))
}

func QMacToolBarItem_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QMacToolBarItem_QMacToolBarItem_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QMacToolBarItem) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QMacToolBarItem_QMacToolBarItem_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQMacToolBarItem_Activated
func callbackQMacToolBarItem_Activated(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMacToolBarItem) ConnectActivated(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activated") {
			C.QMacToolBarItem_ConnectActivated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activated", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activated", f)
		}
	}
}

func (ptr *QMacToolBarItem) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activated")
	}
}

func (ptr *QMacToolBarItem) Activated() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_Activated(ptr.Pointer())
	}
}

func (ptr *QMacToolBarItem) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QMacToolBarItem) SetSelectable(selectable bool) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetSelectable(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(selectable))))
	}
}

func (ptr *QMacToolBarItem) SetStandardItem(standardItem QMacToolBarItem__StandardItem) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetStandardItem(ptr.Pointer(), C.longlong(standardItem))
	}
}

func (ptr *QMacToolBarItem) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QMacToolBarItem_SetText(ptr.Pointer(), C.struct_QtMacExtras_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQMacToolBarItem_DestroyQMacToolBarItem
func callbackQMacToolBarItem_DestroyQMacToolBarItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QMacToolBarItem"); signal != nil {
		signal.(func())()
	} else {
		NewQMacToolBarItemFromPointer(ptr).DestroyQMacToolBarItemDefault()
	}
}

func (ptr *QMacToolBarItem) ConnectDestroyQMacToolBarItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QMacToolBarItem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QMacToolBarItem", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QMacToolBarItem", f)
		}
	}
}

func (ptr *QMacToolBarItem) DisconnectDestroyQMacToolBarItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QMacToolBarItem")
	}
}

func (ptr *QMacToolBarItem) DestroyQMacToolBarItem() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DestroyQMacToolBarItem(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QMacToolBarItem) DestroyQMacToolBarItemDefault() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DestroyQMacToolBarItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QMacToolBarItem) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QMacToolBarItem_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBarItem) StandardItem() QMacToolBarItem__StandardItem {
	if ptr.Pointer() != nil {
		return QMacToolBarItem__StandardItem(C.QMacToolBarItem_StandardItem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMacToolBarItem) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QMacToolBarItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMacToolBarItem) Selectable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QMacToolBarItem_Selectable(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQMacToolBarItem_MetaObject
func callbackQMacToolBarItem_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMacToolBarItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMacToolBarItem) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMacToolBarItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMacToolBarItem) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QMacToolBarItem___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBarItem) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QMacToolBarItem) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QMacToolBarItem___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QMacToolBarItem) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMacToolBarItem___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBarItem) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMacToolBarItem) __findChildren_newList2() unsafe.Pointer {
	return C.QMacToolBarItem___findChildren_newList2(ptr.Pointer())
}

func (ptr *QMacToolBarItem) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMacToolBarItem___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBarItem) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMacToolBarItem) __findChildren_newList3() unsafe.Pointer {
	return C.QMacToolBarItem___findChildren_newList3(ptr.Pointer())
}

func (ptr *QMacToolBarItem) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMacToolBarItem___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBarItem) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMacToolBarItem) __findChildren_newList() unsafe.Pointer {
	return C.QMacToolBarItem___findChildren_newList(ptr.Pointer())
}

func (ptr *QMacToolBarItem) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMacToolBarItem___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBarItem) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMacToolBarItem) __children_newList() unsafe.Pointer {
	return C.QMacToolBarItem___children_newList(ptr.Pointer())
}

//export callbackQMacToolBarItem_Event
func callbackQMacToolBarItem_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarItemFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMacToolBarItem) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QMacToolBarItem_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQMacToolBarItem_EventFilter
func callbackQMacToolBarItem_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMacToolBarItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QMacToolBarItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQMacToolBarItem_ChildEvent
func callbackQMacToolBarItem_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMacToolBarItem_ConnectNotify
func callbackQMacToolBarItem_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBarItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBarItem_CustomEvent
func callbackQMacToolBarItem_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMacToolBarItem_DeleteLater
func callbackQMacToolBarItem_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQMacToolBarItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMacToolBarItem) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQMacToolBarItem_Destroyed
func callbackQMacToolBarItem_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQMacToolBarItem_DisconnectNotify
func callbackQMacToolBarItem_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBarItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBarItem_ObjectNameChanged
func callbackQMacToolBarItem_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtMacExtras_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQMacToolBarItem_TimerEvent
func callbackQMacToolBarItem_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}
