// +build !minimal

package virtualkeyboard

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "virtualkeyboard.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtVirtualKeyboard_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtVirtualKeyboard_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
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

func NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardAbstractInputMethod) {
	n = new(QVirtualKeyboardAbstractInputMethod)
	n.SetPointer(ptr)
	return
}
func NewQVirtualKeyboardAbstractInputMethod(parent core.QObject_ITF) *QVirtualKeyboardAbstractInputMethod {
	tmpValue := NewQVirtualKeyboardAbstractInputMethodFromPointer(C.QVirtualKeyboardAbstractInputMethod_NewQVirtualKeyboardAbstractInputMethod(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQVirtualKeyboardAbstractInputMethod_ClickPreeditText
func callbackQVirtualKeyboardAbstractInputMethod_ClickPreeditText(ptr unsafe.Pointer, cursorPosition C.int) C.char {
	if signal := qt.GetSignal(ptr, "clickPreeditText"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(cursorPosition))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).ClickPreeditTextDefault(int(int32(cursorPosition))))))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectClickPreeditText(f func(cursorPosition int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickPreeditText"); signal != nil {
			f := func(cursorPosition int) bool {
				(*(*func(int) bool)(signal))(cursorPosition)
				return f(cursorPosition)
			}
			qt.ConnectSignal(ptr.Pointer(), "clickPreeditText", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickPreeditText", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectClickPreeditText() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickPreeditText")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ClickPreeditText(cursorPosition int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_ClickPreeditText(ptr.Pointer(), C.int(int32(cursorPosition)))) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ClickPreeditTextDefault(cursorPosition int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_ClickPreeditTextDefault(ptr.Pointer(), C.int(int32(cursorPosition)))) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardAbstractInputMethod) InputContext() *QVirtualKeyboardInputContext {
	if ptr.Pointer() != nil {
		tmpValue := NewQVirtualKeyboardInputContextFromPointer(C.QVirtualKeyboardAbstractInputMethod_InputContext(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardAbstractInputMethod) InputEngine() *QVirtualKeyboardInputEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQVirtualKeyboardInputEngineFromPointer(C.QVirtualKeyboardAbstractInputMethod_InputEngine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardAbstractInputMethod_InputModes
func callbackQVirtualKeyboardAbstractInputMethod_InputModes(ptr unsafe.Pointer, locale C.struct_QtVirtualKeyboard_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputModes"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(NewQVirtualKeyboardAbstractInputMethodFromPointer(nil).__inputModes_newList())
			for _, v := range (*(*func(string) []QVirtualKeyboardInputEngine__InputMode)(signal))(cGoUnpackString(locale)) {
				tmpList.__inputModes_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(NewQVirtualKeyboardAbstractInputMethodFromPointer(nil).__inputModes_newList())
		for _, v := range make([]QVirtualKeyboardInputEngine__InputMode, 0) {
			tmpList.__inputModes_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectInputModes(f func(locale string) []QVirtualKeyboardInputEngine__InputMode) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "inputModes"); signal != nil {
			f := func(locale string) []QVirtualKeyboardInputEngine__InputMode {
				(*(*func(string) []QVirtualKeyboardInputEngine__InputMode)(signal))(locale)
				return f(locale)
			}
			qt.ConnectSignal(ptr.Pointer(), "inputModes", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputModes", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectInputModes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "inputModes")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) InputModes(locale string) []QVirtualKeyboardInputEngine__InputMode {
	if ptr.Pointer() != nil {
		var localeC *C.char
		if locale != "" {
			localeC = C.CString(locale)
			defer C.free(unsafe.Pointer(localeC))
		}
		return func(l C.struct_QtVirtualKeyboard_PackedList) []QVirtualKeyboardInputEngine__InputMode {
			out := make([]QVirtualKeyboardInputEngine__InputMode, int(l.len))
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__inputModes_atList(i)
			}
			return out
		}(C.QVirtualKeyboardAbstractInputMethod_InputModes(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: localeC, len: C.longlong(len(locale))}))
	}
	return make([]QVirtualKeyboardInputEngine__InputMode, 0)
}

//export callbackQVirtualKeyboardAbstractInputMethod_KeyEvent
func callbackQVirtualKeyboardAbstractInputMethod_KeyEvent(ptr unsafe.Pointer, key C.longlong, text C.struct_QtVirtualKeyboard_PackedString, modifiers C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "keyEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(core.Qt__Key, string, core.Qt__KeyboardModifier) bool)(signal))(core.Qt__Key(key), cGoUnpackString(text), core.Qt__KeyboardModifier(modifiers)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectKeyEvent(f func(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keyEvent"); signal != nil {
			f := func(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier) bool {
				(*(*func(core.Qt__Key, string, core.Qt__KeyboardModifier) bool)(signal))(key, text, modifiers)
				return f(key, text, modifiers)
			}
			qt.ConnectSignal(ptr.Pointer(), "keyEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keyEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectKeyEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "keyEvent")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) KeyEvent(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier) bool {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		return int8(C.QVirtualKeyboardAbstractInputMethod_KeyEvent(ptr.Pointer(), C.longlong(key), C.struct_QtVirtualKeyboard_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(modifiers))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardAbstractInputMethod_PatternRecognitionModes
func callbackQVirtualKeyboardAbstractInputMethod_PatternRecognitionModes(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "patternRecognitionModes"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(NewQVirtualKeyboardAbstractInputMethodFromPointer(nil).__patternRecognitionModes_newList())
			for _, v := range (*(*func() []QVirtualKeyboardInputEngine__PatternRecognitionMode)(signal))() {
				tmpList.__patternRecognitionModes_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(NewQVirtualKeyboardAbstractInputMethodFromPointer(nil).__patternRecognitionModes_newList())
		for _, v := range NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).PatternRecognitionModesDefault() {
			tmpList.__patternRecognitionModes_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectPatternRecognitionModes(f func() []QVirtualKeyboardInputEngine__PatternRecognitionMode) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "patternRecognitionModes"); signal != nil {
			f := func() []QVirtualKeyboardInputEngine__PatternRecognitionMode {
				(*(*func() []QVirtualKeyboardInputEngine__PatternRecognitionMode)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "patternRecognitionModes", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "patternRecognitionModes", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectPatternRecognitionModes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "patternRecognitionModes")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) PatternRecognitionModes() []QVirtualKeyboardInputEngine__PatternRecognitionMode {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []QVirtualKeyboardInputEngine__PatternRecognitionMode {
			out := make([]QVirtualKeyboardInputEngine__PatternRecognitionMode, int(l.len))
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__patternRecognitionModes_atList(i)
			}
			return out
		}(C.QVirtualKeyboardAbstractInputMethod_PatternRecognitionModes(ptr.Pointer()))
	}
	return make([]QVirtualKeyboardInputEngine__PatternRecognitionMode, 0)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) PatternRecognitionModesDefault() []QVirtualKeyboardInputEngine__PatternRecognitionMode {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []QVirtualKeyboardInputEngine__PatternRecognitionMode {
			out := make([]QVirtualKeyboardInputEngine__PatternRecognitionMode, int(l.len))
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__patternRecognitionModes_atList(i)
			}
			return out
		}(C.QVirtualKeyboardAbstractInputMethod_PatternRecognitionModesDefault(ptr.Pointer()))
	}
	return make([]QVirtualKeyboardInputEngine__PatternRecognitionMode, 0)
}

//export callbackQVirtualKeyboardAbstractInputMethod_Reselect
func callbackQVirtualKeyboardAbstractInputMethod_Reselect(ptr unsafe.Pointer, cursorPosition C.int, reselectFlags C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "reselect"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, QVirtualKeyboardInputEngine__ReselectFlag) bool)(signal))(int(int32(cursorPosition)), QVirtualKeyboardInputEngine__ReselectFlag(reselectFlags)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).ReselectDefault(int(int32(cursorPosition)), QVirtualKeyboardInputEngine__ReselectFlag(reselectFlags)))))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectReselect(f func(cursorPosition int, reselectFlags QVirtualKeyboardInputEngine__ReselectFlag) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reselect"); signal != nil {
			f := func(cursorPosition int, reselectFlags QVirtualKeyboardInputEngine__ReselectFlag) bool {
				(*(*func(int, QVirtualKeyboardInputEngine__ReselectFlag) bool)(signal))(cursorPosition, reselectFlags)
				return f(cursorPosition, reselectFlags)
			}
			qt.ConnectSignal(ptr.Pointer(), "reselect", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reselect", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectReselect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reselect")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) Reselect(cursorPosition int, reselectFlags QVirtualKeyboardInputEngine__ReselectFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_Reselect(ptr.Pointer(), C.int(int32(cursorPosition)), C.longlong(reselectFlags))) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ReselectDefault(cursorPosition int, reselectFlags QVirtualKeyboardInputEngine__ReselectFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_ReselectDefault(ptr.Pointer(), C.int(int32(cursorPosition)), C.longlong(reselectFlags))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardAbstractInputMethod_Reset
func callbackQVirtualKeyboardAbstractInputMethod_Reset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectReset(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reset"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "reset", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reset", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reset")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) Reset() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_Reset(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ResetDefault() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_ResetDefault(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_SelectionListActiveItemChanged
func callbackQVirtualKeyboardAbstractInputMethod_SelectionListActiveItemChanged(ptr unsafe.Pointer, ty C.longlong, index C.int) {
	if signal := qt.GetSignal(ptr, "selectionListActiveItemChanged"); signal != nil {
		(*(*func(QVirtualKeyboardSelectionListModel__Type, int))(signal))(QVirtualKeyboardSelectionListModel__Type(ty), int(int32(index)))
	}

}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListActiveItemChanged(f func(ty QVirtualKeyboardSelectionListModel__Type, index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionListActiveItemChanged") {
			C.QVirtualKeyboardAbstractInputMethod_ConnectSelectionListActiveItemChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectionListActiveItemChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionListActiveItemChanged"); signal != nil {
			f := func(ty QVirtualKeyboardSelectionListModel__Type, index int) {
				(*(*func(QVirtualKeyboardSelectionListModel__Type, int))(signal))(ty, index)
				f(ty, index)
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionListActiveItemChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionListActiveItemChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListActiveItemChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_DisconnectSelectionListActiveItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionListActiveItemChanged")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListActiveItemChanged(ty QVirtualKeyboardSelectionListModel__Type, index int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_SelectionListActiveItemChanged(ptr.Pointer(), C.longlong(ty), C.int(int32(index)))
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_SelectionListChanged
func callbackQVirtualKeyboardAbstractInputMethod_SelectionListChanged(ptr unsafe.Pointer, ty C.longlong) {
	if signal := qt.GetSignal(ptr, "selectionListChanged"); signal != nil {
		(*(*func(QVirtualKeyboardSelectionListModel__Type))(signal))(QVirtualKeyboardSelectionListModel__Type(ty))
	}

}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListChanged(f func(ty QVirtualKeyboardSelectionListModel__Type)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionListChanged") {
			C.QVirtualKeyboardAbstractInputMethod_ConnectSelectionListChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectionListChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionListChanged"); signal != nil {
			f := func(ty QVirtualKeyboardSelectionListModel__Type) {
				(*(*func(QVirtualKeyboardSelectionListModel__Type))(signal))(ty)
				f(ty)
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionListChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionListChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_DisconnectSelectionListChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionListChanged")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListChanged(ty QVirtualKeyboardSelectionListModel__Type) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_SelectionListChanged(ptr.Pointer(), C.longlong(ty))
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_SelectionListData
func callbackQVirtualKeyboardAbstractInputMethod_SelectionListData(ptr unsafe.Pointer, ty C.longlong, index C.int, role C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "selectionListData"); signal != nil {
		return core.PointerFromQVariant((*(*func(QVirtualKeyboardSelectionListModel__Type, int, QVirtualKeyboardSelectionListModel__Role) *core.QVariant)(signal))(QVirtualKeyboardSelectionListModel__Type(ty), int(int32(index)), QVirtualKeyboardSelectionListModel__Role(role)))
	}

	return core.PointerFromQVariant(NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).SelectionListDataDefault(QVirtualKeyboardSelectionListModel__Type(ty), int(int32(index)), QVirtualKeyboardSelectionListModel__Role(role)))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListData(f func(ty QVirtualKeyboardSelectionListModel__Type, index int, role QVirtualKeyboardSelectionListModel__Role) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectionListData"); signal != nil {
			f := func(ty QVirtualKeyboardSelectionListModel__Type, index int, role QVirtualKeyboardSelectionListModel__Role) *core.QVariant {
				(*(*func(QVirtualKeyboardSelectionListModel__Type, int, QVirtualKeyboardSelectionListModel__Role) *core.QVariant)(signal))(ty, index, role)
				return f(ty, index, role)
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionListData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionListData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectionListData")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListData(ty QVirtualKeyboardSelectionListModel__Type, index int, role QVirtualKeyboardSelectionListModel__Role) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardAbstractInputMethod_SelectionListData(ptr.Pointer(), C.longlong(ty), C.int(int32(index)), C.longlong(role)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListDataDefault(ty QVirtualKeyboardSelectionListModel__Type, index int, role QVirtualKeyboardSelectionListModel__Role) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardAbstractInputMethod_SelectionListDataDefault(ptr.Pointer(), C.longlong(ty), C.int(int32(index)), C.longlong(role)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardAbstractInputMethod_SelectionListItemCount
func callbackQVirtualKeyboardAbstractInputMethod_SelectionListItemCount(ptr unsafe.Pointer, ty C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "selectionListItemCount"); signal != nil {
		return C.int(int32((*(*func(QVirtualKeyboardSelectionListModel__Type) int)(signal))(QVirtualKeyboardSelectionListModel__Type(ty))))
	}

	return C.int(int32(NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).SelectionListItemCountDefault(QVirtualKeyboardSelectionListModel__Type(ty))))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListItemCount(f func(ty QVirtualKeyboardSelectionListModel__Type) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectionListItemCount"); signal != nil {
			f := func(ty QVirtualKeyboardSelectionListModel__Type) int {
				(*(*func(QVirtualKeyboardSelectionListModel__Type) int)(signal))(ty)
				return f(ty)
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionListItemCount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionListItemCount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListItemCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectionListItemCount")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListItemCount(ty QVirtualKeyboardSelectionListModel__Type) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardAbstractInputMethod_SelectionListItemCount(ptr.Pointer(), C.longlong(ty))))
	}
	return 0
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListItemCountDefault(ty QVirtualKeyboardSelectionListModel__Type) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardAbstractInputMethod_SelectionListItemCountDefault(ptr.Pointer(), C.longlong(ty))))
	}
	return 0
}

//export callbackQVirtualKeyboardAbstractInputMethod_SelectionListItemSelected
func callbackQVirtualKeyboardAbstractInputMethod_SelectionListItemSelected(ptr unsafe.Pointer, ty C.longlong, index C.int) {
	if signal := qt.GetSignal(ptr, "selectionListItemSelected"); signal != nil {
		(*(*func(QVirtualKeyboardSelectionListModel__Type, int))(signal))(QVirtualKeyboardSelectionListModel__Type(ty), int(int32(index)))
	} else {
		NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).SelectionListItemSelectedDefault(QVirtualKeyboardSelectionListModel__Type(ty), int(int32(index)))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListItemSelected(f func(ty QVirtualKeyboardSelectionListModel__Type, index int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectionListItemSelected"); signal != nil {
			f := func(ty QVirtualKeyboardSelectionListModel__Type, index int) {
				(*(*func(QVirtualKeyboardSelectionListModel__Type, int))(signal))(ty, index)
				f(ty, index)
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionListItemSelected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionListItemSelected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListItemSelected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectionListItemSelected")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListItemSelected(ty QVirtualKeyboardSelectionListModel__Type, index int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_SelectionListItemSelected(ptr.Pointer(), C.longlong(ty), C.int(int32(index)))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListItemSelectedDefault(ty QVirtualKeyboardSelectionListModel__Type, index int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_SelectionListItemSelectedDefault(ptr.Pointer(), C.longlong(ty), C.int(int32(index)))
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_SelectionListRemoveItem
func callbackQVirtualKeyboardAbstractInputMethod_SelectionListRemoveItem(ptr unsafe.Pointer, ty C.longlong, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "selectionListRemoveItem"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(QVirtualKeyboardSelectionListModel__Type, int) bool)(signal))(QVirtualKeyboardSelectionListModel__Type(ty), int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).SelectionListRemoveItemDefault(QVirtualKeyboardSelectionListModel__Type(ty), int(int32(index))))))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListRemoveItem(f func(ty QVirtualKeyboardSelectionListModel__Type, index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectionListRemoveItem"); signal != nil {
			f := func(ty QVirtualKeyboardSelectionListModel__Type, index int) bool {
				(*(*func(QVirtualKeyboardSelectionListModel__Type, int) bool)(signal))(ty, index)
				return f(ty, index)
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionListRemoveItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionListRemoveItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListRemoveItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectionListRemoveItem")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListRemoveItem(ty QVirtualKeyboardSelectionListModel__Type, index int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_SelectionListRemoveItem(ptr.Pointer(), C.longlong(ty), C.int(int32(index)))) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListRemoveItemDefault(ty QVirtualKeyboardSelectionListModel__Type, index int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_SelectionListRemoveItemDefault(ptr.Pointer(), C.longlong(ty), C.int(int32(index)))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardAbstractInputMethod_SelectionLists
func callbackQVirtualKeyboardAbstractInputMethod_SelectionLists(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "selectionLists"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(NewQVirtualKeyboardAbstractInputMethodFromPointer(nil).__selectionLists_newList())
			for _, v := range (*(*func() []QVirtualKeyboardSelectionListModel__Type)(signal))() {
				tmpList.__selectionLists_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(NewQVirtualKeyboardAbstractInputMethodFromPointer(nil).__selectionLists_newList())
		for _, v := range NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).SelectionListsDefault() {
			tmpList.__selectionLists_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionLists(f func() []QVirtualKeyboardSelectionListModel__Type) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectionLists"); signal != nil {
			f := func() []QVirtualKeyboardSelectionListModel__Type {
				(*(*func() []QVirtualKeyboardSelectionListModel__Type)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionLists", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionLists", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionLists() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectionLists")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionLists() []QVirtualKeyboardSelectionListModel__Type {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []QVirtualKeyboardSelectionListModel__Type {
			out := make([]QVirtualKeyboardSelectionListModel__Type, int(l.len))
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__selectionLists_atList(i)
			}
			return out
		}(C.QVirtualKeyboardAbstractInputMethod_SelectionLists(ptr.Pointer()))
	}
	return make([]QVirtualKeyboardSelectionListModel__Type, 0)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListsDefault() []QVirtualKeyboardSelectionListModel__Type {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []QVirtualKeyboardSelectionListModel__Type {
			out := make([]QVirtualKeyboardSelectionListModel__Type, int(l.len))
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__selectionLists_atList(i)
			}
			return out
		}(C.QVirtualKeyboardAbstractInputMethod_SelectionListsDefault(ptr.Pointer()))
	}
	return make([]QVirtualKeyboardSelectionListModel__Type, 0)
}

//export callbackQVirtualKeyboardAbstractInputMethod_SelectionListsChanged
func callbackQVirtualKeyboardAbstractInputMethod_SelectionListsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionListsChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSelectionListsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionListsChanged") {
			C.QVirtualKeyboardAbstractInputMethod_ConnectSelectionListsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectionListsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionListsChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionListsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionListsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSelectionListsChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_DisconnectSelectionListsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionListsChanged")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SelectionListsChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_SelectionListsChanged(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_SetInputMode
func callbackQVirtualKeyboardAbstractInputMethod_SetInputMode(ptr unsafe.Pointer, locale C.struct_QtVirtualKeyboard_PackedString, inputMode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "setInputMode"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(string, QVirtualKeyboardInputEngine__InputMode) bool)(signal))(cGoUnpackString(locale), QVirtualKeyboardInputEngine__InputMode(inputMode)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSetInputMode(f func(locale string, inputMode QVirtualKeyboardInputEngine__InputMode) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setInputMode"); signal != nil {
			f := func(locale string, inputMode QVirtualKeyboardInputEngine__InputMode) bool {
				(*(*func(string, QVirtualKeyboardInputEngine__InputMode) bool)(signal))(locale, inputMode)
				return f(locale, inputMode)
			}
			qt.ConnectSignal(ptr.Pointer(), "setInputMode", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setInputMode", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSetInputMode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setInputMode")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SetInputMode(locale string, inputMode QVirtualKeyboardInputEngine__InputMode) bool {
	if ptr.Pointer() != nil {
		var localeC *C.char
		if locale != "" {
			localeC = C.CString(locale)
			defer C.free(unsafe.Pointer(localeC))
		}
		return int8(C.QVirtualKeyboardAbstractInputMethod_SetInputMode(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: localeC, len: C.longlong(len(locale))}, C.longlong(inputMode))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardAbstractInputMethod_SetTextCase
func callbackQVirtualKeyboardAbstractInputMethod_SetTextCase(ptr unsafe.Pointer, textCase C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "setTextCase"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(QVirtualKeyboardInputEngine__TextCase) bool)(signal))(QVirtualKeyboardInputEngine__TextCase(textCase)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectSetTextCase(f func(textCase QVirtualKeyboardInputEngine__TextCase) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTextCase"); signal != nil {
			f := func(textCase QVirtualKeyboardInputEngine__TextCase) bool {
				(*(*func(QVirtualKeyboardInputEngine__TextCase) bool)(signal))(textCase)
				return f(textCase)
			}
			qt.ConnectSignal(ptr.Pointer(), "setTextCase", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTextCase", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectSetTextCase() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTextCase")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) SetTextCase(textCase QVirtualKeyboardInputEngine__TextCase) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_SetTextCase(ptr.Pointer(), C.longlong(textCase))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardAbstractInputMethod_TraceBegin
func callbackQVirtualKeyboardAbstractInputMethod_TraceBegin(ptr unsafe.Pointer, traceId C.int, patternRecognitionMode C.longlong, traceCaptureDeviceInfo C.struct_QtVirtualKeyboard_PackedList, traceScreenInfo C.struct_QtVirtualKeyboard_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "traceBegin"); signal != nil {
		return PointerFromQVirtualKeyboardTrace((*(*func(int, QVirtualKeyboardInputEngine__PatternRecognitionMode, map[string]*core.QVariant, map[string]*core.QVariant) *QVirtualKeyboardTrace)(signal))(int(int32(traceId)), QVirtualKeyboardInputEngine__PatternRecognitionMode(patternRecognitionMode), func(l C.struct_QtVirtualKeyboard_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
			for i, v := range tmpList.__traceBegin_traceCaptureDeviceInfo_keyList() {
				out[v] = tmpList.__traceBegin_traceCaptureDeviceInfo_atList(v, i)
			}
			return out
		}(traceCaptureDeviceInfo), func(l C.struct_QtVirtualKeyboard_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
			for i, v := range tmpList.__traceBegin_traceScreenInfo_keyList() {
				out[v] = tmpList.__traceBegin_traceScreenInfo_atList(v, i)
			}
			return out
		}(traceScreenInfo)))
	}

	return PointerFromQVirtualKeyboardTrace(NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).TraceBeginDefault(int(int32(traceId)), QVirtualKeyboardInputEngine__PatternRecognitionMode(patternRecognitionMode), func(l C.struct_QtVirtualKeyboard_PackedList) map[string]*core.QVariant {
		out := make(map[string]*core.QVariant, int(l.len))
		tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
		for i, v := range tmpList.__traceBegin_traceCaptureDeviceInfo_keyList() {
			out[v] = tmpList.__traceBegin_traceCaptureDeviceInfo_atList(v, i)
		}
		return out
	}(traceCaptureDeviceInfo), func(l C.struct_QtVirtualKeyboard_PackedList) map[string]*core.QVariant {
		out := make(map[string]*core.QVariant, int(l.len))
		tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
		for i, v := range tmpList.__traceBegin_traceScreenInfo_keyList() {
			out[v] = tmpList.__traceBegin_traceScreenInfo_atList(v, i)
		}
		return out
	}(traceScreenInfo)))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectTraceBegin(f func(traceId int, patternRecognitionMode QVirtualKeyboardInputEngine__PatternRecognitionMode, traceCaptureDeviceInfo map[string]*core.QVariant, traceScreenInfo map[string]*core.QVariant) *QVirtualKeyboardTrace) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "traceBegin"); signal != nil {
			f := func(traceId int, patternRecognitionMode QVirtualKeyboardInputEngine__PatternRecognitionMode, traceCaptureDeviceInfo map[string]*core.QVariant, traceScreenInfo map[string]*core.QVariant) *QVirtualKeyboardTrace {
				(*(*func(int, QVirtualKeyboardInputEngine__PatternRecognitionMode, map[string]*core.QVariant, map[string]*core.QVariant) *QVirtualKeyboardTrace)(signal))(traceId, patternRecognitionMode, traceCaptureDeviceInfo, traceScreenInfo)
				return f(traceId, patternRecognitionMode, traceCaptureDeviceInfo, traceScreenInfo)
			}
			qt.ConnectSignal(ptr.Pointer(), "traceBegin", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "traceBegin", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectTraceBegin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "traceBegin")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) TraceBegin(traceId int, patternRecognitionMode QVirtualKeyboardInputEngine__PatternRecognitionMode, traceCaptureDeviceInfo map[string]*core.QVariant, traceScreenInfo map[string]*core.QVariant) *QVirtualKeyboardTrace {
	if ptr.Pointer() != nil {
		tmpValue := NewQVirtualKeyboardTraceFromPointer(C.QVirtualKeyboardAbstractInputMethod_TraceBegin(ptr.Pointer(), C.int(int32(traceId)), C.longlong(patternRecognitionMode), func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(NewQVirtualKeyboardAbstractInputMethodFromPointer(nil).__traceBegin_traceCaptureDeviceInfo_newList())
			for k, v := range traceCaptureDeviceInfo {
				tmpList.__traceBegin_traceCaptureDeviceInfo_setList(k, v)
			}
			return tmpList.Pointer()
		}(), func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(NewQVirtualKeyboardAbstractInputMethodFromPointer(nil).__traceBegin_traceScreenInfo_newList())
			for k, v := range traceScreenInfo {
				tmpList.__traceBegin_traceScreenInfo_setList(k, v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardAbstractInputMethod) TraceBeginDefault(traceId int, patternRecognitionMode QVirtualKeyboardInputEngine__PatternRecognitionMode, traceCaptureDeviceInfo map[string]*core.QVariant, traceScreenInfo map[string]*core.QVariant) *QVirtualKeyboardTrace {
	if ptr.Pointer() != nil {
		tmpValue := NewQVirtualKeyboardTraceFromPointer(C.QVirtualKeyboardAbstractInputMethod_TraceBeginDefault(ptr.Pointer(), C.int(int32(traceId)), C.longlong(patternRecognitionMode), func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(NewQVirtualKeyboardAbstractInputMethodFromPointer(nil).__traceBegin_traceCaptureDeviceInfo_newList())
			for k, v := range traceCaptureDeviceInfo {
				tmpList.__traceBegin_traceCaptureDeviceInfo_setList(k, v)
			}
			return tmpList.Pointer()
		}(), func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(NewQVirtualKeyboardAbstractInputMethodFromPointer(nil).__traceBegin_traceScreenInfo_newList())
			for k, v := range traceScreenInfo {
				tmpList.__traceBegin_traceScreenInfo_setList(k, v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardAbstractInputMethod_TraceEnd
func callbackQVirtualKeyboardAbstractInputMethod_TraceEnd(ptr unsafe.Pointer, trace unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "traceEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QVirtualKeyboardTrace) bool)(signal))(NewQVirtualKeyboardTraceFromPointer(trace)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).TraceEndDefault(NewQVirtualKeyboardTraceFromPointer(trace)))))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectTraceEnd(f func(trace *QVirtualKeyboardTrace) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "traceEnd"); signal != nil {
			f := func(trace *QVirtualKeyboardTrace) bool {
				(*(*func(*QVirtualKeyboardTrace) bool)(signal))(trace)
				return f(trace)
			}
			qt.ConnectSignal(ptr.Pointer(), "traceEnd", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "traceEnd", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectTraceEnd() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "traceEnd")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) TraceEnd(trace QVirtualKeyboardTrace_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_TraceEnd(ptr.Pointer(), PointerFromQVirtualKeyboardTrace(trace))) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardAbstractInputMethod) TraceEndDefault(trace QVirtualKeyboardTrace_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_TraceEndDefault(ptr.Pointer(), PointerFromQVirtualKeyboardTrace(trace))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardAbstractInputMethod_Update
func callbackQVirtualKeyboardAbstractInputMethod_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "update"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "update", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "update", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "update")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) Update() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_Update(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_DestroyQVirtualKeyboardAbstractInputMethod
func callbackQVirtualKeyboardAbstractInputMethod_DestroyQVirtualKeyboardAbstractInputMethod(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QVirtualKeyboardAbstractInputMethod"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).DestroyQVirtualKeyboardAbstractInputMethodDefault()
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectDestroyQVirtualKeyboardAbstractInputMethod(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QVirtualKeyboardAbstractInputMethod"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QVirtualKeyboardAbstractInputMethod", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QVirtualKeyboardAbstractInputMethod", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectDestroyQVirtualKeyboardAbstractInputMethod() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QVirtualKeyboardAbstractInputMethod")
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DestroyQVirtualKeyboardAbstractInputMethod() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVirtualKeyboardAbstractInputMethod_DestroyQVirtualKeyboardAbstractInputMethod(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DestroyQVirtualKeyboardAbstractInputMethodDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVirtualKeyboardAbstractInputMethod_DestroyQVirtualKeyboardAbstractInputMethodDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __inputModes_atList(i int) QVirtualKeyboardInputEngine__InputMode {
	if ptr.Pointer() != nil {
		return QVirtualKeyboardInputEngine__InputMode(C.QVirtualKeyboardAbstractInputMethod___inputModes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __inputModes_setList(i QVirtualKeyboardInputEngine__InputMode) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod___inputModes_setList(ptr.Pointer(), C.longlong(i))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __inputModes_newList() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod___inputModes_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __patternRecognitionModes_atList(i int) QVirtualKeyboardInputEngine__PatternRecognitionMode {
	if ptr.Pointer() != nil {
		return QVirtualKeyboardInputEngine__PatternRecognitionMode(C.QVirtualKeyboardAbstractInputMethod___patternRecognitionModes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __patternRecognitionModes_setList(i QVirtualKeyboardInputEngine__PatternRecognitionMode) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod___patternRecognitionModes_setList(ptr.Pointer(), C.longlong(i))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __patternRecognitionModes_newList() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod___patternRecognitionModes_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __selectionLists_atList(i int) QVirtualKeyboardSelectionListModel__Type {
	if ptr.Pointer() != nil {
		return QVirtualKeyboardSelectionListModel__Type(C.QVirtualKeyboardAbstractInputMethod___selectionLists_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __selectionLists_setList(i QVirtualKeyboardSelectionListModel__Type) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod___selectionLists_setList(ptr.Pointer(), C.longlong(i))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __selectionLists_newList() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod___selectionLists_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceCaptureDeviceInfo_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardAbstractInputMethod___traceBegin_traceCaptureDeviceInfo_atList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceCaptureDeviceInfo_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QVirtualKeyboardAbstractInputMethod___traceBegin_traceCaptureDeviceInfo_setList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceCaptureDeviceInfo_newList() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod___traceBegin_traceCaptureDeviceInfo_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceCaptureDeviceInfo_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____traceBegin_traceCaptureDeviceInfo_keyList_atList(i)
			}
			return out
		}(C.QVirtualKeyboardAbstractInputMethod___traceBegin_traceCaptureDeviceInfo_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceScreenInfo_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardAbstractInputMethod___traceBegin_traceScreenInfo_atList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceScreenInfo_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QVirtualKeyboardAbstractInputMethod___traceBegin_traceScreenInfo_setList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceScreenInfo_newList() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod___traceBegin_traceScreenInfo_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __traceBegin_traceScreenInfo_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQVirtualKeyboardAbstractInputMethodFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____traceBegin_traceScreenInfo_keyList_atList(i)
			}
			return out
		}(C.QVirtualKeyboardAbstractInputMethod___traceBegin_traceScreenInfo_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceCaptureDeviceInfo_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVirtualKeyboardAbstractInputMethod_____traceBegin_traceCaptureDeviceInfo_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceCaptureDeviceInfo_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QVirtualKeyboardAbstractInputMethod_____traceBegin_traceCaptureDeviceInfo_keyList_setList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceCaptureDeviceInfo_keyList_newList() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod_____traceBegin_traceCaptureDeviceInfo_keyList_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceScreenInfo_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVirtualKeyboardAbstractInputMethod_____traceBegin_traceScreenInfo_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceScreenInfo_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QVirtualKeyboardAbstractInputMethod_____traceBegin_traceScreenInfo_keyList_setList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ____traceBegin_traceScreenInfo_keyList_newList() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod_____traceBegin_traceScreenInfo_keyList_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardAbstractInputMethod___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __children_newList() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod___children_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QVirtualKeyboardAbstractInputMethod___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardAbstractInputMethod___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_newList() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod___findChildren_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardAbstractInputMethod___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) __findChildren_newList3() unsafe.Pointer {
	return C.QVirtualKeyboardAbstractInputMethod___findChildren_newList3(ptr.Pointer())
}

//export callbackQVirtualKeyboardAbstractInputMethod_ChildEvent
func callbackQVirtualKeyboardAbstractInputMethod_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_ConnectNotify
func callbackQVirtualKeyboardAbstractInputMethod_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_CustomEvent
func callbackQVirtualKeyboardAbstractInputMethod_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_DeleteLater
func callbackQVirtualKeyboardAbstractInputMethod_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVirtualKeyboardAbstractInputMethod_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_Destroyed
func callbackQVirtualKeyboardAbstractInputMethod_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQVirtualKeyboardAbstractInputMethod_DisconnectNotify
func callbackQVirtualKeyboardAbstractInputMethod_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardAbstractInputMethod_Event
func callbackQVirtualKeyboardAbstractInputMethod_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardAbstractInputMethod_EventFilter
func callbackQVirtualKeyboardAbstractInputMethod_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QVirtualKeyboardAbstractInputMethod) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardAbstractInputMethod_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardAbstractInputMethod_MetaObject
func callbackQVirtualKeyboardAbstractInputMethod_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QVirtualKeyboardAbstractInputMethod) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVirtualKeyboardAbstractInputMethod_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVirtualKeyboardAbstractInputMethod_ObjectNameChanged
func callbackQVirtualKeyboardAbstractInputMethod_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtVirtualKeyboard_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQVirtualKeyboardAbstractInputMethod_TimerEvent
func callbackQVirtualKeyboardAbstractInputMethod_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVirtualKeyboardAbstractInputMethodFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardAbstractInputMethod) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardAbstractInputMethod_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQVirtualKeyboardExtensionPluginFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardExtensionPlugin) {
	n = new(QVirtualKeyboardExtensionPlugin)
	n.SetPointer(ptr)
	return
}

//export callbackQVirtualKeyboardExtensionPlugin_RegisterTypes
func callbackQVirtualKeyboardExtensionPlugin_RegisterTypes(ptr unsafe.Pointer, uri C.struct_QtVirtualKeyboard_PackedString) {
	if signal := qt.GetSignal(ptr, "registerTypes"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(uri))
	} else {
		NewQVirtualKeyboardExtensionPluginFromPointer(ptr).RegisterTypesDefault(cGoUnpackString(uri))
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) ConnectRegisterTypes(f func(uri string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "registerTypes"); signal != nil {
			f := func(uri string) {
				(*(*func(string))(signal))(uri)
				f(uri)
			}
			qt.ConnectSignal(ptr.Pointer(), "registerTypes", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "registerTypes", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) DisconnectRegisterTypes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "registerTypes")
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) RegisterTypes(uri string) {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		C.QVirtualKeyboardExtensionPlugin_RegisterTypes(ptr.Pointer(), uriC)
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) RegisterTypesDefault(uri string) {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		C.QVirtualKeyboardExtensionPlugin_RegisterTypesDefault(ptr.Pointer(), uriC)
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardExtensionPlugin___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardExtensionPlugin) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardExtensionPlugin___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) __children_newList() unsafe.Pointer {
	return C.QVirtualKeyboardExtensionPlugin___children_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardExtensionPlugin) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QVirtualKeyboardExtensionPlugin___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardExtensionPlugin) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardExtensionPlugin___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QVirtualKeyboardExtensionPlugin___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardExtensionPlugin___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardExtensionPlugin___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_newList() unsafe.Pointer {
	return C.QVirtualKeyboardExtensionPlugin___findChildren_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardExtensionPlugin___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardExtensionPlugin___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) __findChildren_newList3() unsafe.Pointer {
	return C.QVirtualKeyboardExtensionPlugin___findChildren_newList3(ptr.Pointer())
}

//export callbackQVirtualKeyboardExtensionPlugin_ChildEvent
func callbackQVirtualKeyboardExtensionPlugin_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVirtualKeyboardExtensionPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardExtensionPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQVirtualKeyboardExtensionPlugin_ConnectNotify
func callbackQVirtualKeyboardExtensionPlugin_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardExtensionPluginFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardExtensionPlugin_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardExtensionPlugin_CustomEvent
func callbackQVirtualKeyboardExtensionPlugin_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQVirtualKeyboardExtensionPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardExtensionPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQVirtualKeyboardExtensionPlugin_DeleteLater
func callbackQVirtualKeyboardExtensionPlugin_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardExtensionPluginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVirtualKeyboardExtensionPlugin_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardExtensionPlugin_Destroyed
func callbackQVirtualKeyboardExtensionPlugin_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQVirtualKeyboardExtensionPlugin_DisconnectNotify
func callbackQVirtualKeyboardExtensionPlugin_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardExtensionPluginFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardExtensionPlugin_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardExtensionPlugin_Event
func callbackQVirtualKeyboardExtensionPlugin_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardExtensionPluginFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QVirtualKeyboardExtensionPlugin) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardExtensionPlugin_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardExtensionPlugin_EventFilter
func callbackQVirtualKeyboardExtensionPlugin_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardExtensionPluginFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QVirtualKeyboardExtensionPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardExtensionPlugin_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardExtensionPlugin_MetaObject
func callbackQVirtualKeyboardExtensionPlugin_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQVirtualKeyboardExtensionPluginFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QVirtualKeyboardExtensionPlugin) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVirtualKeyboardExtensionPlugin_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVirtualKeyboardExtensionPlugin_ObjectNameChanged
func callbackQVirtualKeyboardExtensionPlugin_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtVirtualKeyboard_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQVirtualKeyboardExtensionPlugin_TimerEvent
func callbackQVirtualKeyboardExtensionPlugin_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVirtualKeyboardExtensionPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardExtensionPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQVirtualKeyboardInputContextFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardInputContext) {
	n = new(QVirtualKeyboardInputContext)
	n.SetPointer(ptr)
	return
}
func (ptr *QVirtualKeyboardInputContext) AnchorPosition() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardInputContext_AnchorPosition(ptr.Pointer())))
	}
	return 0
}

//export callbackQVirtualKeyboardInputContext_AnchorPositionChanged
func callbackQVirtualKeyboardInputContext_AnchorPositionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "anchorPositionChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectAnchorPositionChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "anchorPositionChanged") {
			C.QVirtualKeyboardInputContext_ConnectAnchorPositionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "anchorPositionChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "anchorPositionChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "anchorPositionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "anchorPositionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectAnchorPositionChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectAnchorPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "anchorPositionChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) AnchorPositionChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_AnchorPositionChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) AnchorRectIntersectsClipRect() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputContext_AnchorRectIntersectsClipRect(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQVirtualKeyboardInputContext_AnchorRectIntersectsClipRectChanged
func callbackQVirtualKeyboardInputContext_AnchorRectIntersectsClipRectChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "anchorRectIntersectsClipRectChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectAnchorRectIntersectsClipRectChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "anchorRectIntersectsClipRectChanged") {
			C.QVirtualKeyboardInputContext_ConnectAnchorRectIntersectsClipRectChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "anchorRectIntersectsClipRectChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "anchorRectIntersectsClipRectChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "anchorRectIntersectsClipRectChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "anchorRectIntersectsClipRectChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectAnchorRectIntersectsClipRectChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectAnchorRectIntersectsClipRectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "anchorRectIntersectsClipRectChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) AnchorRectIntersectsClipRectChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_AnchorRectIntersectsClipRectChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) AnchorRectangle() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QVirtualKeyboardInputContext_AnchorRectangle(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardInputContext_AnchorRectangleChanged
func callbackQVirtualKeyboardInputContext_AnchorRectangleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "anchorRectangleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectAnchorRectangleChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "anchorRectangleChanged") {
			C.QVirtualKeyboardInputContext_ConnectAnchorRectangleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "anchorRectangleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "anchorRectangleChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "anchorRectangleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "anchorRectangleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectAnchorRectangleChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectAnchorRectangleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "anchorRectangleChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) AnchorRectangleChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_AnchorRectangleChanged(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardInputContext_AnimatingChanged
func callbackQVirtualKeyboardInputContext_AnimatingChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "animatingChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectAnimatingChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "animatingChanged") {
			C.QVirtualKeyboardInputContext_ConnectAnimatingChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "animatingChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "animatingChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "animatingChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "animatingChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectAnimatingChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectAnimatingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "animatingChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) AnimatingChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_AnimatingChanged(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardInputContext_CapsLockActiveChanged
func callbackQVirtualKeyboardInputContext_CapsLockActiveChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "capsLockActiveChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectCapsLockActiveChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "capsLockActiveChanged") {
			C.QVirtualKeyboardInputContext_ConnectCapsLockActiveChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "capsLockActiveChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "capsLockActiveChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "capsLockActiveChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "capsLockActiveChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectCapsLockActiveChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectCapsLockActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "capsLockActiveChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) CapsLockActiveChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_CapsLockActiveChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) Clear() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_Clear(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) Commit() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_Commit(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) Commit2(text string, replaceFrom int, replaceLength int) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QVirtualKeyboardInputContext_Commit2(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: textC, len: C.longlong(len(text))}, C.int(int32(replaceFrom)), C.int(int32(replaceLength)))
	}
}

func (ptr *QVirtualKeyboardInputContext) CursorPosition() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardInputContext_CursorPosition(ptr.Pointer())))
	}
	return 0
}

//export callbackQVirtualKeyboardInputContext_CursorPositionChanged
func callbackQVirtualKeyboardInputContext_CursorPositionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cursorPositionChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectCursorPositionChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "cursorPositionChanged") {
			C.QVirtualKeyboardInputContext_ConnectCursorPositionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "cursorPositionChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "cursorPositionChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "cursorPositionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cursorPositionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectCursorPositionChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectCursorPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "cursorPositionChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) CursorPositionChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_CursorPositionChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) CursorRectIntersectsClipRect() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputContext_CursorRectIntersectsClipRect(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQVirtualKeyboardInputContext_CursorRectIntersectsClipRectChanged
func callbackQVirtualKeyboardInputContext_CursorRectIntersectsClipRectChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cursorRectIntersectsClipRectChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectCursorRectIntersectsClipRectChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "cursorRectIntersectsClipRectChanged") {
			C.QVirtualKeyboardInputContext_ConnectCursorRectIntersectsClipRectChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "cursorRectIntersectsClipRectChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "cursorRectIntersectsClipRectChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "cursorRectIntersectsClipRectChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cursorRectIntersectsClipRectChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectCursorRectIntersectsClipRectChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectCursorRectIntersectsClipRectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "cursorRectIntersectsClipRectChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) CursorRectIntersectsClipRectChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_CursorRectIntersectsClipRectChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) CursorRectangle() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QVirtualKeyboardInputContext_CursorRectangle(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardInputContext_CursorRectangleChanged
func callbackQVirtualKeyboardInputContext_CursorRectangleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cursorRectangleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectCursorRectangleChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "cursorRectangleChanged") {
			C.QVirtualKeyboardInputContext_ConnectCursorRectangleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "cursorRectangleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "cursorRectangleChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "cursorRectangleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cursorRectangleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectCursorRectangleChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectCursorRectangleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "cursorRectangleChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) CursorRectangleChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_CursorRectangleChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) InputEngine() *QVirtualKeyboardInputEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQVirtualKeyboardInputEngineFromPointer(C.QVirtualKeyboardInputContext_InputEngine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputContext) InputItem() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardInputContext_InputItem(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardInputContext_InputItemChanged
func callbackQVirtualKeyboardInputContext_InputItemChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputItemChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectInputItemChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "inputItemChanged") {
			C.QVirtualKeyboardInputContext_ConnectInputItemChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "inputItemChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "inputItemChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "inputItemChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputItemChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectInputItemChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectInputItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "inputItemChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) InputItemChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_InputItemChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) InputMethodHints() core.Qt__InputMethodHint {
	if ptr.Pointer() != nil {
		return core.Qt__InputMethodHint(C.QVirtualKeyboardInputContext_InputMethodHints(ptr.Pointer()))
	}
	return 0
}

//export callbackQVirtualKeyboardInputContext_InputMethodHintsChanged
func callbackQVirtualKeyboardInputContext_InputMethodHintsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodHintsChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectInputMethodHintsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "inputMethodHintsChanged") {
			C.QVirtualKeyboardInputContext_ConnectInputMethodHintsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "inputMethodHintsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "inputMethodHintsChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "inputMethodHintsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputMethodHintsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectInputMethodHintsChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectInputMethodHintsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "inputMethodHintsChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) InputMethodHintsChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_InputMethodHintsChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) IsAnimating() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputContext_IsAnimating(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardInputContext) IsCapsLockActive() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputContext_IsCapsLockActive(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardInputContext) IsSelectionControlVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputContext_IsSelectionControlVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardInputContext) IsShiftActive() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputContext_IsShiftActive(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardInputContext) IsUppercase() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputContext_IsUppercase(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardInputContext) Locale() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVirtualKeyboardInputContext_Locale(ptr.Pointer()))
	}
	return ""
}

//export callbackQVirtualKeyboardInputContext_LocaleChanged
func callbackQVirtualKeyboardInputContext_LocaleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "localeChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectLocaleChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "localeChanged") {
			C.QVirtualKeyboardInputContext_ConnectLocaleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "localeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "localeChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "localeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "localeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectLocaleChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectLocaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "localeChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) LocaleChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_LocaleChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) PreeditText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVirtualKeyboardInputContext_PreeditText(ptr.Pointer()))
	}
	return ""
}

//export callbackQVirtualKeyboardInputContext_PreeditTextChanged
func callbackQVirtualKeyboardInputContext_PreeditTextChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "preeditTextChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectPreeditTextChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "preeditTextChanged") {
			C.QVirtualKeyboardInputContext_ConnectPreeditTextChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "preeditTextChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "preeditTextChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "preeditTextChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preeditTextChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectPreeditTextChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectPreeditTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "preeditTextChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) PreeditTextChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_PreeditTextChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVirtualKeyboardInputContext_SelectedText(ptr.Pointer()))
	}
	return ""
}

//export callbackQVirtualKeyboardInputContext_SelectedTextChanged
func callbackQVirtualKeyboardInputContext_SelectedTextChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectedTextChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectSelectedTextChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectedTextChanged") {
			C.QVirtualKeyboardInputContext_ConnectSelectedTextChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectedTextChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectedTextChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "selectedTextChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectedTextChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectSelectedTextChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectSelectedTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectedTextChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) SelectedTextChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_SelectedTextChanged(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardInputContext_SelectionControlVisibleChanged
func callbackQVirtualKeyboardInputContext_SelectionControlVisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionControlVisibleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectSelectionControlVisibleChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionControlVisibleChanged") {
			C.QVirtualKeyboardInputContext_ConnectSelectionControlVisibleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectionControlVisibleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionControlVisibleChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionControlVisibleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionControlVisibleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectSelectionControlVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectSelectionControlVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionControlVisibleChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) SelectionControlVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_SelectionControlVisibleChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) SendKeyClick(key int, text string, modifiers int) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QVirtualKeyboardInputContext_SendKeyClick(ptr.Pointer(), C.int(int32(key)), C.struct_QtVirtualKeyboard_PackedString{data: textC, len: C.longlong(len(text))}, C.int(int32(modifiers)))
	}
}

func (ptr *QVirtualKeyboardInputContext) SetAnimating(isAnimating bool) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_SetAnimating(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isAnimating))))
	}
}

//export callbackQVirtualKeyboardInputContext_ShiftActiveChanged
func callbackQVirtualKeyboardInputContext_ShiftActiveChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "shiftActiveChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectShiftActiveChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "shiftActiveChanged") {
			C.QVirtualKeyboardInputContext_ConnectShiftActiveChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "shiftActiveChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "shiftActiveChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "shiftActiveChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shiftActiveChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectShiftActiveChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectShiftActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "shiftActiveChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) ShiftActiveChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_ShiftActiveChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) SurroundingText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVirtualKeyboardInputContext_SurroundingText(ptr.Pointer()))
	}
	return ""
}

//export callbackQVirtualKeyboardInputContext_SurroundingTextChanged
func callbackQVirtualKeyboardInputContext_SurroundingTextChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "surroundingTextChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectSurroundingTextChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "surroundingTextChanged") {
			C.QVirtualKeyboardInputContext_ConnectSurroundingTextChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "surroundingTextChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "surroundingTextChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "surroundingTextChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "surroundingTextChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectSurroundingTextChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectSurroundingTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "surroundingTextChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) SurroundingTextChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_SurroundingTextChanged(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardInputContext_UppercaseChanged
func callbackQVirtualKeyboardInputContext_UppercaseChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "uppercaseChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputContext) ConnectUppercaseChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "uppercaseChanged") {
			C.QVirtualKeyboardInputContext_ConnectUppercaseChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "uppercaseChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "uppercaseChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "uppercaseChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "uppercaseChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectUppercaseChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectUppercaseChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "uppercaseChanged")
	}
}

func (ptr *QVirtualKeyboardInputContext) UppercaseChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_UppercaseChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputContext) __preeditTextAttributes_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputContext___preeditTextAttributes_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputContext) __setPreeditText_attributes_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputContext___setPreeditText_attributes_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputContext) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardInputContext___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputContext) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardInputContext) __children_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputContext___children_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputContext) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QVirtualKeyboardInputContext___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputContext) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QVirtualKeyboardInputContext) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputContext___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardInputContext___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputContext___findChildren_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardInputContext___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardInputContext) __findChildren_newList3() unsafe.Pointer {
	return C.QVirtualKeyboardInputContext___findChildren_newList3(ptr.Pointer())
}

//export callbackQVirtualKeyboardInputContext_ChildEvent
func callbackQVirtualKeyboardInputContext_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVirtualKeyboardInputContextFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardInputContext) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQVirtualKeyboardInputContext_ConnectNotify
func callbackQVirtualKeyboardInputContext_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardInputContextFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardInputContext) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardInputContext_CustomEvent
func callbackQVirtualKeyboardInputContext_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQVirtualKeyboardInputContextFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardInputContext) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQVirtualKeyboardInputContext_DeleteLater
func callbackQVirtualKeyboardInputContext_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardInputContextFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVirtualKeyboardInputContext) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVirtualKeyboardInputContext_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardInputContext_Destroyed
func callbackQVirtualKeyboardInputContext_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQVirtualKeyboardInputContext_DisconnectNotify
func callbackQVirtualKeyboardInputContext_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardInputContextFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardInputContext) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardInputContext_Event
func callbackQVirtualKeyboardInputContext_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardInputContextFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QVirtualKeyboardInputContext) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputContext_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardInputContext_EventFilter
func callbackQVirtualKeyboardInputContext_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardInputContextFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QVirtualKeyboardInputContext) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputContext_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardInputContext_MetaObject
func callbackQVirtualKeyboardInputContext_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQVirtualKeyboardInputContextFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QVirtualKeyboardInputContext) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVirtualKeyboardInputContext_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVirtualKeyboardInputContext_ObjectNameChanged
func callbackQVirtualKeyboardInputContext_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtVirtualKeyboard_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQVirtualKeyboardInputContext_TimerEvent
func callbackQVirtualKeyboardInputContext_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVirtualKeyboardInputContextFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardInputContext) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputContext_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQVirtualKeyboardInputEngineFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardInputEngine) {
	n = new(QVirtualKeyboardInputEngine)
	n.SetPointer(ptr)
	return
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
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QVirtualKeyboardInputEngine_ActiveKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQVirtualKeyboardInputEngine_ActiveKeyChanged
func callbackQVirtualKeyboardInputEngine_ActiveKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "activeKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectActiveKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activeKeyChanged") {
			C.QVirtualKeyboardInputEngine_ConnectActiveKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "activeKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activeKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "activeKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectActiveKeyChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectActiveKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activeKeyChanged")
	}
}

func (ptr *QVirtualKeyboardInputEngine) ActiveKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_ActiveKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QVirtualKeyboardInputEngine) InputContext() *QVirtualKeyboardInputContext {
	if ptr.Pointer() != nil {
		tmpValue := NewQVirtualKeyboardInputContextFromPointer(C.QVirtualKeyboardInputEngine_InputContext(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputEngine) InputMethod() *QVirtualKeyboardAbstractInputMethod {
	if ptr.Pointer() != nil {
		tmpValue := NewQVirtualKeyboardAbstractInputMethodFromPointer(C.QVirtualKeyboardInputEngine_InputMethod(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardInputEngine_InputMethodChanged
func callbackQVirtualKeyboardInputEngine_InputMethodChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectInputMethodChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "inputMethodChanged") {
			C.QVirtualKeyboardInputEngine_ConnectInputMethodChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "inputMethodChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "inputMethodChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "inputMethodChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputMethodChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectInputMethodChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectInputMethodChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "inputMethodChanged")
	}
}

func (ptr *QVirtualKeyboardInputEngine) InputMethodChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_InputMethodChanged(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardInputEngine_InputMethodReset
func callbackQVirtualKeyboardInputEngine_InputMethodReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodReset"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectInputMethodReset(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "inputMethodReset") {
			C.QVirtualKeyboardInputEngine_ConnectInputMethodReset(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "inputMethodReset")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "inputMethodReset"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "inputMethodReset", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputMethodReset", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectInputMethodReset() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectInputMethodReset(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "inputMethodReset")
	}
}

func (ptr *QVirtualKeyboardInputEngine) InputMethodReset() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_InputMethodReset(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardInputEngine_InputMethodUpdate
func callbackQVirtualKeyboardInputEngine_InputMethodUpdate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodUpdate"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectInputMethodUpdate(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "inputMethodUpdate") {
			C.QVirtualKeyboardInputEngine_ConnectInputMethodUpdate(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "inputMethodUpdate")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "inputMethodUpdate"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "inputMethodUpdate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputMethodUpdate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectInputMethodUpdate() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectInputMethodUpdate(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "inputMethodUpdate")
	}
}

func (ptr *QVirtualKeyboardInputEngine) InputMethodUpdate() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_InputMethodUpdate(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputEngine) InputMode() QVirtualKeyboardInputEngine__InputMode {
	if ptr.Pointer() != nil {
		return QVirtualKeyboardInputEngine__InputMode(C.QVirtualKeyboardInputEngine_InputMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQVirtualKeyboardInputEngine_InputModeChanged
func callbackQVirtualKeyboardInputEngine_InputModeChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputModeChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectInputModeChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "inputModeChanged") {
			C.QVirtualKeyboardInputEngine_ConnectInputModeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "inputModeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "inputModeChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "inputModeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputModeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectInputModeChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectInputModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "inputModeChanged")
	}
}

func (ptr *QVirtualKeyboardInputEngine) InputModeChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_InputModeChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputEngine) InputModes() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQVirtualKeyboardInputEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__inputModes_atList(i)
			}
			return out
		}(C.QVirtualKeyboardInputEngine_InputModes(ptr.Pointer()))
	}
	return make([]int, 0)
}

//export callbackQVirtualKeyboardInputEngine_InputModesChanged
func callbackQVirtualKeyboardInputEngine_InputModesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputModesChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectInputModesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "inputModesChanged") {
			C.QVirtualKeyboardInputEngine_ConnectInputModesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "inputModesChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "inputModesChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "inputModesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputModesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectInputModesChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectInputModesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "inputModesChanged")
	}
}

func (ptr *QVirtualKeyboardInputEngine) InputModesChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_InputModesChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputEngine) PatternRecognitionModes() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQVirtualKeyboardInputEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__patternRecognitionModes_atList(i)
			}
			return out
		}(C.QVirtualKeyboardInputEngine_PatternRecognitionModes(ptr.Pointer()))
	}
	return make([]int, 0)
}

//export callbackQVirtualKeyboardInputEngine_PatternRecognitionModesChanged
func callbackQVirtualKeyboardInputEngine_PatternRecognitionModesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "patternRecognitionModesChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectPatternRecognitionModesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "patternRecognitionModesChanged") {
			C.QVirtualKeyboardInputEngine_ConnectPatternRecognitionModesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "patternRecognitionModesChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "patternRecognitionModesChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "patternRecognitionModesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "patternRecognitionModesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectPatternRecognitionModesChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectPatternRecognitionModesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "patternRecognitionModesChanged")
	}
}

func (ptr *QVirtualKeyboardInputEngine) PatternRecognitionModesChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_PatternRecognitionModesChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputEngine) PreviousKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QVirtualKeyboardInputEngine_PreviousKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQVirtualKeyboardInputEngine_PreviousKeyChanged
func callbackQVirtualKeyboardInputEngine_PreviousKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "previousKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectPreviousKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "previousKeyChanged") {
			C.QVirtualKeyboardInputEngine_ConnectPreviousKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "previousKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "previousKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "previousKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "previousKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectPreviousKeyChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectPreviousKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "previousKeyChanged")
	}
}

func (ptr *QVirtualKeyboardInputEngine) PreviousKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_PreviousKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QVirtualKeyboardInputEngine) Reselect(cursorPosition int, reselectFlags QVirtualKeyboardInputEngine__ReselectFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputEngine_Reselect(ptr.Pointer(), C.int(int32(cursorPosition)), C.longlong(reselectFlags))) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardInputEngine) SetInputMethod(inputMethod QVirtualKeyboardAbstractInputMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_SetInputMethod(ptr.Pointer(), PointerFromQVirtualKeyboardAbstractInputMethod(inputMethod))
	}
}

func (ptr *QVirtualKeyboardInputEngine) SetInputMode(inputMode QVirtualKeyboardInputEngine__InputMode) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_SetInputMode(ptr.Pointer(), C.longlong(inputMode))
	}
}

func (ptr *QVirtualKeyboardInputEngine) TraceBegin(traceId int, patternRecognitionMode QVirtualKeyboardInputEngine__PatternRecognitionMode, traceCaptureDeviceInfo map[string]*core.QVariant, traceScreenInfo map[string]*core.QVariant) *QVirtualKeyboardTrace {
	if ptr.Pointer() != nil {
		tmpValue := NewQVirtualKeyboardTraceFromPointer(C.QVirtualKeyboardInputEngine_TraceBegin(ptr.Pointer(), C.int(int32(traceId)), C.longlong(patternRecognitionMode), func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardInputEngineFromPointer(NewQVirtualKeyboardInputEngineFromPointer(nil).__traceBegin_traceCaptureDeviceInfo_newList())
			for k, v := range traceCaptureDeviceInfo {
				tmpList.__traceBegin_traceCaptureDeviceInfo_setList(k, v)
			}
			return tmpList.Pointer()
		}(), func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardInputEngineFromPointer(NewQVirtualKeyboardInputEngineFromPointer(nil).__traceBegin_traceScreenInfo_newList())
			for k, v := range traceScreenInfo {
				tmpList.__traceBegin_traceScreenInfo_setList(k, v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputEngine) TraceEnd(trace QVirtualKeyboardTrace_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputEngine_TraceEnd(ptr.Pointer(), PointerFromQVirtualKeyboardTrace(trace))) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardInputEngine) VirtualKeyCancel() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_VirtualKeyCancel(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputEngine) VirtualKeyClick(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier) bool {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		return int8(C.QVirtualKeyboardInputEngine_VirtualKeyClick(ptr.Pointer(), C.longlong(key), C.struct_QtVirtualKeyboard_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(modifiers))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardInputEngine_VirtualKeyClicked
func callbackQVirtualKeyboardInputEngine_VirtualKeyClicked(ptr unsafe.Pointer, key C.longlong, text C.struct_QtVirtualKeyboard_PackedString, modifiers C.longlong, isAutoRepeat C.char) {
	if signal := qt.GetSignal(ptr, "virtualKeyClicked"); signal != nil {
		(*(*func(core.Qt__Key, string, core.Qt__KeyboardModifier, bool))(signal))(core.Qt__Key(key), cGoUnpackString(text), core.Qt__KeyboardModifier(modifiers), int8(isAutoRepeat) != 0)
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectVirtualKeyClicked(f func(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier, isAutoRepeat bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "virtualKeyClicked") {
			C.QVirtualKeyboardInputEngine_ConnectVirtualKeyClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "virtualKeyClicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "virtualKeyClicked"); signal != nil {
			f := func(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier, isAutoRepeat bool) {
				(*(*func(core.Qt__Key, string, core.Qt__KeyboardModifier, bool))(signal))(key, text, modifiers, isAutoRepeat)
				f(key, text, modifiers, isAutoRepeat)
			}
			qt.ConnectSignal(ptr.Pointer(), "virtualKeyClicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "virtualKeyClicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectVirtualKeyClicked() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectVirtualKeyClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "virtualKeyClicked")
	}
}

func (ptr *QVirtualKeyboardInputEngine) VirtualKeyClicked(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier, isAutoRepeat bool) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QVirtualKeyboardInputEngine_VirtualKeyClicked(ptr.Pointer(), C.longlong(key), C.struct_QtVirtualKeyboard_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(modifiers), C.char(int8(qt.GoBoolToInt(isAutoRepeat))))
	}
}

func (ptr *QVirtualKeyboardInputEngine) VirtualKeyPress(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier, repeat bool) bool {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		return int8(C.QVirtualKeyboardInputEngine_VirtualKeyPress(ptr.Pointer(), C.longlong(key), C.struct_QtVirtualKeyboard_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(modifiers), C.char(int8(qt.GoBoolToInt(repeat))))) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardInputEngine) VirtualKeyRelease(key core.Qt__Key, text string, modifiers core.Qt__KeyboardModifier) bool {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		return int8(C.QVirtualKeyboardInputEngine_VirtualKeyRelease(ptr.Pointer(), C.longlong(key), C.struct_QtVirtualKeyboard_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(modifiers))) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardInputEngine) WordCandidateListModel() *QVirtualKeyboardSelectionListModel {
	if ptr.Pointer() != nil {
		tmpValue := NewQVirtualKeyboardSelectionListModelFromPointer(C.QVirtualKeyboardInputEngine_WordCandidateListModel(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardInputEngine_WordCandidateListModelChanged
func callbackQVirtualKeyboardInputEngine_WordCandidateListModelChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wordCandidateListModelChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectWordCandidateListModelChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "wordCandidateListModelChanged") {
			C.QVirtualKeyboardInputEngine_ConnectWordCandidateListModelChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "wordCandidateListModelChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "wordCandidateListModelChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "wordCandidateListModelChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "wordCandidateListModelChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectWordCandidateListModelChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectWordCandidateListModelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "wordCandidateListModelChanged")
	}
}

func (ptr *QVirtualKeyboardInputEngine) WordCandidateListModelChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_WordCandidateListModelChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputEngine) WordCandidateListVisibleHint() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputEngine_WordCandidateListVisibleHint(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQVirtualKeyboardInputEngine_WordCandidateListVisibleHintChanged
func callbackQVirtualKeyboardInputEngine_WordCandidateListVisibleHintChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wordCandidateListVisibleHintChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardInputEngine) ConnectWordCandidateListVisibleHintChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "wordCandidateListVisibleHintChanged") {
			C.QVirtualKeyboardInputEngine_ConnectWordCandidateListVisibleHintChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "wordCandidateListVisibleHintChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "wordCandidateListVisibleHintChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "wordCandidateListVisibleHintChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "wordCandidateListVisibleHintChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectWordCandidateListVisibleHintChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectWordCandidateListVisibleHintChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "wordCandidateListVisibleHintChanged")
	}
}

func (ptr *QVirtualKeyboardInputEngine) WordCandidateListVisibleHintChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_WordCandidateListVisibleHintChanged(ptr.Pointer())
	}
}

func (ptr *QVirtualKeyboardInputEngine) __inputModes_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardInputEngine___inputModes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QVirtualKeyboardInputEngine) __inputModes_setList(i int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine___inputModes_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QVirtualKeyboardInputEngine) __inputModes_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputEngine___inputModes_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputEngine) __patternRecognitionModes_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardInputEngine___patternRecognitionModes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QVirtualKeyboardInputEngine) __patternRecognitionModes_setList(i int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine___patternRecognitionModes_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QVirtualKeyboardInputEngine) __patternRecognitionModes_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputEngine___patternRecognitionModes_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceCaptureDeviceInfo_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardInputEngine___traceBegin_traceCaptureDeviceInfo_atList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceCaptureDeviceInfo_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QVirtualKeyboardInputEngine___traceBegin_traceCaptureDeviceInfo_setList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceCaptureDeviceInfo_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputEngine___traceBegin_traceCaptureDeviceInfo_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceCaptureDeviceInfo_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQVirtualKeyboardInputEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____traceBegin_traceCaptureDeviceInfo_keyList_atList(i)
			}
			return out
		}(C.QVirtualKeyboardInputEngine___traceBegin_traceCaptureDeviceInfo_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceScreenInfo_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardInputEngine___traceBegin_traceScreenInfo_atList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceScreenInfo_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QVirtualKeyboardInputEngine___traceBegin_traceScreenInfo_setList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceScreenInfo_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputEngine___traceBegin_traceScreenInfo_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputEngine) __traceBegin_traceScreenInfo_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQVirtualKeyboardInputEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____traceBegin_traceScreenInfo_keyList_atList(i)
			}
			return out
		}(C.QVirtualKeyboardInputEngine___traceBegin_traceScreenInfo_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceCaptureDeviceInfo_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVirtualKeyboardInputEngine_____traceBegin_traceCaptureDeviceInfo_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceCaptureDeviceInfo_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QVirtualKeyboardInputEngine_____traceBegin_traceCaptureDeviceInfo_keyList_setList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceCaptureDeviceInfo_keyList_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputEngine_____traceBegin_traceCaptureDeviceInfo_keyList_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceScreenInfo_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVirtualKeyboardInputEngine_____traceBegin_traceScreenInfo_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceScreenInfo_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QVirtualKeyboardInputEngine_____traceBegin_traceScreenInfo_keyList_setList(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QVirtualKeyboardInputEngine) ____traceBegin_traceScreenInfo_keyList_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputEngine_____traceBegin_traceScreenInfo_keyList_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardInputEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardInputEngine) __children_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputEngine___children_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QVirtualKeyboardInputEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QVirtualKeyboardInputEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputEngine___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardInputEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_newList() unsafe.Pointer {
	return C.QVirtualKeyboardInputEngine___findChildren_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardInputEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardInputEngine) __findChildren_newList3() unsafe.Pointer {
	return C.QVirtualKeyboardInputEngine___findChildren_newList3(ptr.Pointer())
}

//export callbackQVirtualKeyboardInputEngine_ChildEvent
func callbackQVirtualKeyboardInputEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVirtualKeyboardInputEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardInputEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQVirtualKeyboardInputEngine_ConnectNotify
func callbackQVirtualKeyboardInputEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardInputEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardInputEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardInputEngine_CustomEvent
func callbackQVirtualKeyboardInputEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQVirtualKeyboardInputEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardInputEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQVirtualKeyboardInputEngine_DeleteLater
func callbackQVirtualKeyboardInputEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardInputEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVirtualKeyboardInputEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVirtualKeyboardInputEngine_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardInputEngine_Destroyed
func callbackQVirtualKeyboardInputEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQVirtualKeyboardInputEngine_DisconnectNotify
func callbackQVirtualKeyboardInputEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardInputEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardInputEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardInputEngine_Event
func callbackQVirtualKeyboardInputEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardInputEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QVirtualKeyboardInputEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardInputEngine_EventFilter
func callbackQVirtualKeyboardInputEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardInputEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QVirtualKeyboardInputEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardInputEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardInputEngine_MetaObject
func callbackQVirtualKeyboardInputEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQVirtualKeyboardInputEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QVirtualKeyboardInputEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVirtualKeyboardInputEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVirtualKeyboardInputEngine_ObjectNameChanged
func callbackQVirtualKeyboardInputEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtVirtualKeyboard_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQVirtualKeyboardInputEngine_TimerEvent
func callbackQVirtualKeyboardInputEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVirtualKeyboardInputEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardInputEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardInputEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQVirtualKeyboardSelectionListModelFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardSelectionListModel) {
	n = new(QVirtualKeyboardSelectionListModel)
	n.SetPointer(ptr)
	return
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

//export callbackQVirtualKeyboardSelectionListModel_ActiveItemChanged
func callbackQVirtualKeyboardSelectionListModel_ActiveItemChanged(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "activeItemChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QVirtualKeyboardSelectionListModel) ConnectActiveItemChanged(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activeItemChanged") {
			C.QVirtualKeyboardSelectionListModel_ConnectActiveItemChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "activeItemChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activeItemChanged"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "activeItemChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeItemChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) DisconnectActiveItemChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_DisconnectActiveItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activeItemChanged")
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) ActiveItemChanged(index int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_ActiveItemChanged(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQVirtualKeyboardSelectionListModel_ItemSelected
func callbackQVirtualKeyboardSelectionListModel_ItemSelected(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "itemSelected"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QVirtualKeyboardSelectionListModel) ConnectItemSelected(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "itemSelected") {
			C.QVirtualKeyboardSelectionListModel_ConnectItemSelected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "itemSelected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "itemSelected"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "itemSelected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "itemSelected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) DisconnectItemSelected() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_DisconnectItemSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "itemSelected")
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) ItemSelected(index int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_ItemSelected(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) RemoveItem(index int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_RemoveItem(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) SelectItem(index int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_SelectItem(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __roleNames_atList(v int, i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QVirtualKeyboardSelectionListModel___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __roleNames_setList(key int, i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___roleNames_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQByteArray(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __roleNames_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___roleNames_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.QVirtualKeyboardSelectionListModel___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QVirtualKeyboardSelectionListModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardSelectionListModel_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QVirtualKeyboardSelectionListModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardSelectionListModel_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QVirtualKeyboardSelectionListModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardSelectionListModel_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_from_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___changePersistentIndexList_from_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_to_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_to_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___changePersistentIndexList_to_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardSelectionListModel___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QVirtualKeyboardSelectionListModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __itemData_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardSelectionListModel___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __itemData_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___itemData_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __itemData_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___itemData_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.QVirtualKeyboardSelectionListModel___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutAboutToBeChanged_parents_setList(i core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutChanged_parents_setList(i core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___layoutChanged_parents_setList(ptr.Pointer(), core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __match_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __match_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___match_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __match_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___match_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __mimeData_indexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __mimeData_indexes_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___mimeData_indexes_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __persistentIndexList_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __persistentIndexList_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___persistentIndexList_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __setItemData_roles_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardSelectionListModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __setItemData_roles_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.QVirtualKeyboardSelectionListModel___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QVirtualKeyboardSelectionListModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardSelectionListModel_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QVirtualKeyboardSelectionListModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardSelectionListModel_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardSelectionListModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __children_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___children_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QVirtualKeyboardSelectionListModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardSelectionListModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_newList() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___findChildren_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardSelectionListModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) __findChildren_newList3() unsafe.Pointer {
	return C.QVirtualKeyboardSelectionListModel___findChildren_newList3(ptr.Pointer())
}

//export callbackQVirtualKeyboardSelectionListModel_DropMimeData
func callbackQVirtualKeyboardSelectionListModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(signal))(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_Flags
func callbackQVirtualKeyboardSelectionListModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*core.QModelIndex) core.Qt__ItemFlag)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QVirtualKeyboardSelectionListModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QVirtualKeyboardSelectionListModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQVirtualKeyboardSelectionListModel_Index
func callbackQVirtualKeyboardSelectionListModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(int, int, *core.QModelIndex) *core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QVirtualKeyboardSelectionListModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardSelectionListModel_Sibling
func callbackQVirtualKeyboardSelectionListModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(int, int, *core.QModelIndex) *core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
	}

	return core.PointerFromQModelIndex(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
}

func (ptr *QVirtualKeyboardSelectionListModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardSelectionListModel_Buddy
func callbackQVirtualKeyboardSelectionListModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QModelIndex) *core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QVirtualKeyboardSelectionListModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardSelectionListModel_CanDropMimeData
func callbackQVirtualKeyboardSelectionListModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(signal))(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_CanFetchMore
func callbackQVirtualKeyboardSelectionListModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).CanFetchMoreDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_CanFetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_ColumnCount
func callbackQVirtualKeyboardSelectionListModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*core.QModelIndex) int)(signal))(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).ColumnCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QVirtualKeyboardSelectionListModel) ColumnCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardSelectionListModel_ColumnCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQVirtualKeyboardSelectionListModel_ColumnsAboutToBeInserted
func callbackQVirtualKeyboardSelectionListModel_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_ColumnsAboutToBeMoved
func callbackQVirtualKeyboardSelectionListModel_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_ColumnsAboutToBeRemoved
func callbackQVirtualKeyboardSelectionListModel_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_ColumnsInserted
func callbackQVirtualKeyboardSelectionListModel_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_ColumnsMoved
func callbackQVirtualKeyboardSelectionListModel_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_ColumnsRemoved
func callbackQVirtualKeyboardSelectionListModel_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_Data
func callbackQVirtualKeyboardSelectionListModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQVariant((*(*func(*core.QModelIndex, int) *core.QVariant)(signal))(core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).DataDefault(core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *QVirtualKeyboardSelectionListModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardSelectionListModel_DataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardSelectionListModel_DataChanged
func callbackQVirtualKeyboardSelectionListModel_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_QtVirtualKeyboard_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*core.QModelIndex, *core.QModelIndex, []int))(signal))(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtVirtualKeyboard_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_FetchMore
func callbackQVirtualKeyboardSelectionListModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQVirtualKeyboardSelectionListModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

//export callbackQVirtualKeyboardSelectionListModel_HasChildren
func callbackQVirtualKeyboardSelectionListModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_HeaderData
func callbackQVirtualKeyboardSelectionListModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return core.PointerFromQVariant((*(*func(int, core.Qt__Orientation, int) *core.QVariant)(signal))(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *QVirtualKeyboardSelectionListModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardSelectionListModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardSelectionListModel_HeaderDataChanged
func callbackQVirtualKeyboardSelectionListModel_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(core.Qt__Orientation, int, int))(signal))(core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_InsertColumns
func callbackQVirtualKeyboardSelectionListModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_InsertRows
func callbackQVirtualKeyboardSelectionListModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_ItemData
func callbackQVirtualKeyboardSelectionListModel_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(NewQVirtualKeyboardSelectionListModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*core.QModelIndex) map[int]*core.QVariant)(signal))(core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(NewQVirtualKeyboardSelectionListModelFromPointer(nil).__itemData_newList())
		for k, v := range NewQVirtualKeyboardSelectionListModelFromPointer(ptr).ItemDataDefault(core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QVirtualKeyboardSelectionListModel) ItemDataDefault(index core.QModelIndex_ITF) map[int]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) map[int]*core.QVariant {
			out := make(map[int]*core.QVariant, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.QVirtualKeyboardSelectionListModel_ItemDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*core.QVariant, 0)
}

//export callbackQVirtualKeyboardSelectionListModel_LayoutAboutToBeChanged
func callbackQVirtualKeyboardSelectionListModel_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_QtVirtualKeyboard_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_QtVirtualKeyboard_PackedList) []*core.QPersistentModelIndex {
			out := make([]*core.QPersistentModelIndex, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_LayoutChanged
func callbackQVirtualKeyboardSelectionListModel_LayoutChanged(ptr unsafe.Pointer, parents C.struct_QtVirtualKeyboard_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_QtVirtualKeyboard_PackedList) []*core.QPersistentModelIndex {
			out := make([]*core.QPersistentModelIndex, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_Match
func callbackQVirtualKeyboardSelectionListModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(NewQVirtualKeyboardSelectionListModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*core.QModelIndex, int, *core.QVariant, int, core.Qt__MatchFlag) []*core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(NewQVirtualKeyboardSelectionListModelFromPointer(nil).__match_newList())
		for _, v := range NewQVirtualKeyboardSelectionListModelFromPointer(ptr).MatchDefault(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QVirtualKeyboardSelectionListModel) MatchDefault(start core.QModelIndex_ITF, role int, value core.QVariant_ITF, hits int, flags core.Qt__MatchFlag) []*core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []*core.QModelIndex {
			out := make([]*core.QModelIndex, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.QVirtualKeyboardSelectionListModel_MatchDefault(ptr.Pointer(), core.PointerFromQModelIndex(start), C.int(int32(role)), core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*core.QModelIndex, 0)
}

//export callbackQVirtualKeyboardSelectionListModel_MimeData
func callbackQVirtualKeyboardSelectionListModel_MimeData(ptr unsafe.Pointer, indexes C.struct_QtVirtualKeyboard_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return core.PointerFromQMimeData((*(*func([]*core.QModelIndex) *core.QMimeData)(signal))(func(l C.struct_QtVirtualKeyboard_PackedList) []*core.QModelIndex {
			out := make([]*core.QModelIndex, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return core.PointerFromQMimeData(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).MimeDataDefault(func(l C.struct_QtVirtualKeyboard_PackedList) []*core.QModelIndex {
		out := make([]*core.QModelIndex, int(l.len))
		tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *QVirtualKeyboardSelectionListModel) MimeDataDefault(indexes []*core.QModelIndex) *core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQMimeDataFromPointer(C.QVirtualKeyboardSelectionListModel_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(NewQVirtualKeyboardSelectionListModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardSelectionListModel_MimeTypes
func callbackQVirtualKeyboardSelectionListModel_MimeTypes(ptr unsafe.Pointer) C.struct_QtVirtualKeyboard_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_QtVirtualKeyboard_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewQVirtualKeyboardSelectionListModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_QtVirtualKeyboard_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *QVirtualKeyboardSelectionListModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QVirtualKeyboardSelectionListModel_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQVirtualKeyboardSelectionListModel_ModelAboutToBeReset
func callbackQVirtualKeyboardSelectionListModel_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQVirtualKeyboardSelectionListModel_ModelReset
func callbackQVirtualKeyboardSelectionListModel_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQVirtualKeyboardSelectionListModel_MoveColumns
func callbackQVirtualKeyboardSelectionListModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QVirtualKeyboardSelectionListModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_MoveRows
func callbackQVirtualKeyboardSelectionListModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QVirtualKeyboardSelectionListModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_Parent
func callbackQVirtualKeyboardSelectionListModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QModelIndex) *core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QVirtualKeyboardSelectionListModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QVirtualKeyboardSelectionListModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardSelectionListModel_RemoveColumns
func callbackQVirtualKeyboardSelectionListModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_RemoveRows
func callbackQVirtualKeyboardSelectionListModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_ResetInternalData
func callbackQVirtualKeyboardSelectionListModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardSelectionListModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardSelectionListModel_Revert
func callbackQVirtualKeyboardSelectionListModel_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardSelectionListModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardSelectionListModel_RoleNames
func callbackQVirtualKeyboardSelectionListModel_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(NewQVirtualKeyboardSelectionListModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(NewQVirtualKeyboardSelectionListModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewQVirtualKeyboardSelectionListModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QVirtualKeyboardSelectionListModel) RoleNamesDefault() map[int]*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) map[int]*core.QByteArray {
			out := make(map[int]*core.QByteArray, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.QVirtualKeyboardSelectionListModel_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*core.QByteArray, 0)
}

//export callbackQVirtualKeyboardSelectionListModel_RowCount
func callbackQVirtualKeyboardSelectionListModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*core.QModelIndex) int)(signal))(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).RowCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QVirtualKeyboardSelectionListModel) RowCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardSelectionListModel_RowCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQVirtualKeyboardSelectionListModel_RowsAboutToBeInserted
func callbackQVirtualKeyboardSelectionListModel_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_RowsAboutToBeMoved
func callbackQVirtualKeyboardSelectionListModel_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_RowsAboutToBeRemoved
func callbackQVirtualKeyboardSelectionListModel_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_RowsInserted
func callbackQVirtualKeyboardSelectionListModel_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_RowsMoved
func callbackQVirtualKeyboardSelectionListModel_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int, *core.QModelIndex, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_RowsRemoved
func callbackQVirtualKeyboardSelectionListModel_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*core.QModelIndex, int, int))(signal))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_SetData
func callbackQVirtualKeyboardSelectionListModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, *core.QVariant, int) bool)(signal))(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QVirtualKeyboardSelectionListModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_SetHeaderData
func callbackQVirtualKeyboardSelectionListModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, core.Qt__Orientation, *core.QVariant, int) bool)(signal))(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QVirtualKeyboardSelectionListModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_SetItemData
func callbackQVirtualKeyboardSelectionListModel_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_QtVirtualKeyboard_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, map[int]*core.QVariant) bool)(signal))(core.NewQModelIndexFromPointer(index), func(l C.struct_QtVirtualKeyboard_PackedList) map[int]*core.QVariant {
			out := make(map[int]*core.QVariant, int(l.len))
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).SetItemDataDefault(core.NewQModelIndexFromPointer(index), func(l C.struct_QtVirtualKeyboard_PackedList) map[int]*core.QVariant {
		out := make(map[int]*core.QVariant, int(l.len))
		tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) SetItemDataDefault(index core.QModelIndex_ITF, roles map[int]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_SetItemDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewQVirtualKeyboardSelectionListModelFromPointer(NewQVirtualKeyboardSelectionListModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_Sort
func callbackQVirtualKeyboardSelectionListModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, core.Qt__SortOrder))(signal))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQVirtualKeyboardSelectionListModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) SortDefault(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQVirtualKeyboardSelectionListModel_Span
func callbackQVirtualKeyboardSelectionListModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return core.PointerFromQSize((*(*func(*core.QModelIndex) *core.QSize)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QVirtualKeyboardSelectionListModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QVirtualKeyboardSelectionListModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQVirtualKeyboardSelectionListModel_Submit
func callbackQVirtualKeyboardSelectionListModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QVirtualKeyboardSelectionListModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_SupportedDragActions
func callbackQVirtualKeyboardSelectionListModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QVirtualKeyboardSelectionListModel) SupportedDragActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QVirtualKeyboardSelectionListModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQVirtualKeyboardSelectionListModel_SupportedDropActions
func callbackQVirtualKeyboardSelectionListModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QVirtualKeyboardSelectionListModel) SupportedDropActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QVirtualKeyboardSelectionListModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQVirtualKeyboardSelectionListModel_ChildEvent
func callbackQVirtualKeyboardSelectionListModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVirtualKeyboardSelectionListModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQVirtualKeyboardSelectionListModel_ConnectNotify
func callbackQVirtualKeyboardSelectionListModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardSelectionListModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardSelectionListModel_CustomEvent
func callbackQVirtualKeyboardSelectionListModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQVirtualKeyboardSelectionListModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQVirtualKeyboardSelectionListModel_DeleteLater
func callbackQVirtualKeyboardSelectionListModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardSelectionListModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVirtualKeyboardSelectionListModel_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardSelectionListModel_Destroyed
func callbackQVirtualKeyboardSelectionListModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_DisconnectNotify
func callbackQVirtualKeyboardSelectionListModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardSelectionListModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardSelectionListModel_Event
func callbackQVirtualKeyboardSelectionListModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_EventFilter
func callbackQVirtualKeyboardSelectionListModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QVirtualKeyboardSelectionListModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardSelectionListModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardSelectionListModel_MetaObject
func callbackQVirtualKeyboardSelectionListModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQVirtualKeyboardSelectionListModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QVirtualKeyboardSelectionListModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVirtualKeyboardSelectionListModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVirtualKeyboardSelectionListModel_ObjectNameChanged
func callbackQVirtualKeyboardSelectionListModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtVirtualKeyboard_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQVirtualKeyboardSelectionListModel_TimerEvent
func callbackQVirtualKeyboardSelectionListModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVirtualKeyboardSelectionListModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardSelectionListModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardSelectionListModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQVirtualKeyboardTraceFromPointer(ptr unsafe.Pointer) (n *QVirtualKeyboardTrace) {
	n = new(QVirtualKeyboardTrace)
	n.SetPointer(ptr)
	return
}
func (ptr *QVirtualKeyboardTrace) AddPoint(point core.QPointF_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardTrace_AddPoint(ptr.Pointer(), core.PointerFromQPointF(point))))
	}
	return 0
}

//export callbackQVirtualKeyboardTrace_CanceledChanged
func callbackQVirtualKeyboardTrace_CanceledChanged(ptr unsafe.Pointer, isCanceled C.char) {
	if signal := qt.GetSignal(ptr, "canceledChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(isCanceled) != 0)
	}

}

func (ptr *QVirtualKeyboardTrace) ConnectCanceledChanged(f func(isCanceled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "canceledChanged") {
			C.QVirtualKeyboardTrace_ConnectCanceledChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "canceledChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "canceledChanged"); signal != nil {
			f := func(isCanceled bool) {
				(*(*func(bool))(signal))(isCanceled)
				f(isCanceled)
			}
			qt.ConnectSignal(ptr.Pointer(), "canceledChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "canceledChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardTrace) DisconnectCanceledChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_DisconnectCanceledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "canceledChanged")
	}
}

func (ptr *QVirtualKeyboardTrace) CanceledChanged(isCanceled bool) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_CanceledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isCanceled))))
	}
}

func (ptr *QVirtualKeyboardTrace) ChannelData(channel string, pos int, count int) []*core.QVariant {
	if ptr.Pointer() != nil {
		var channelC *C.char
		if channel != "" {
			channelC = C.CString(channel)
			defer C.free(unsafe.Pointer(channelC))
		}
		return func(l C.struct_QtVirtualKeyboard_PackedList) []*core.QVariant {
			out := make([]*core.QVariant, int(l.len))
			tmpList := NewQVirtualKeyboardTraceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__channelData_atList(i)
			}
			return out
		}(C.QVirtualKeyboardTrace_ChannelData(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: channelC, len: C.longlong(len(channel))}, C.int(int32(pos)), C.int(int32(count))))
	}
	return make([]*core.QVariant, 0)
}

func (ptr *QVirtualKeyboardTrace) Channels() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QVirtualKeyboardTrace_Channels(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQVirtualKeyboardTrace_ChannelsChanged
func callbackQVirtualKeyboardTrace_ChannelsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "channelsChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QVirtualKeyboardTrace) ConnectChannelsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "channelsChanged") {
			C.QVirtualKeyboardTrace_ConnectChannelsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "channelsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "channelsChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "channelsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "channelsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardTrace) DisconnectChannelsChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_DisconnectChannelsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "channelsChanged")
	}
}

func (ptr *QVirtualKeyboardTrace) ChannelsChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_ChannelsChanged(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardTrace_FinalChanged
func callbackQVirtualKeyboardTrace_FinalChanged(ptr unsafe.Pointer, isFinal C.char) {
	if signal := qt.GetSignal(ptr, "finalChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(isFinal) != 0)
	}

}

func (ptr *QVirtualKeyboardTrace) ConnectFinalChanged(f func(isFinal bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finalChanged") {
			C.QVirtualKeyboardTrace_ConnectFinalChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finalChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finalChanged"); signal != nil {
			f := func(isFinal bool) {
				(*(*func(bool))(signal))(isFinal)
				f(isFinal)
			}
			qt.ConnectSignal(ptr.Pointer(), "finalChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finalChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardTrace) DisconnectFinalChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_DisconnectFinalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finalChanged")
	}
}

func (ptr *QVirtualKeyboardTrace) FinalChanged(isFinal bool) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_FinalChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isFinal))))
	}
}

func (ptr *QVirtualKeyboardTrace) IsCanceled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardTrace_IsCanceled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardTrace) IsFinal() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardTrace_IsFinal(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVirtualKeyboardTrace) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardTrace_Length(ptr.Pointer())))
	}
	return 0
}

//export callbackQVirtualKeyboardTrace_LengthChanged
func callbackQVirtualKeyboardTrace_LengthChanged(ptr unsafe.Pointer, length C.int) {
	if signal := qt.GetSignal(ptr, "lengthChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(length)))
	}

}

func (ptr *QVirtualKeyboardTrace) ConnectLengthChanged(f func(length int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lengthChanged") {
			C.QVirtualKeyboardTrace_ConnectLengthChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lengthChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lengthChanged"); signal != nil {
			f := func(length int) {
				(*(*func(int))(signal))(length)
				f(length)
			}
			qt.ConnectSignal(ptr.Pointer(), "lengthChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lengthChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardTrace) DisconnectLengthChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_DisconnectLengthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lengthChanged")
	}
}

func (ptr *QVirtualKeyboardTrace) LengthChanged(length int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_LengthChanged(ptr.Pointer(), C.int(int32(length)))
	}
}

func (ptr *QVirtualKeyboardTrace) Opacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QVirtualKeyboardTrace_Opacity(ptr.Pointer()))
	}
	return 0
}

//export callbackQVirtualKeyboardTrace_OpacityChanged
func callbackQVirtualKeyboardTrace_OpacityChanged(ptr unsafe.Pointer, opacity C.double) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(opacity))
	}

}

func (ptr *QVirtualKeyboardTrace) ConnectOpacityChanged(f func(opacity float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "opacityChanged") {
			C.QVirtualKeyboardTrace_ConnectOpacityChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "opacityChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "opacityChanged"); signal != nil {
			f := func(opacity float64) {
				(*(*func(float64))(signal))(opacity)
				f(opacity)
			}
			qt.ConnectSignal(ptr.Pointer(), "opacityChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "opacityChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardTrace) DisconnectOpacityChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_DisconnectOpacityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "opacityChanged")
	}
}

func (ptr *QVirtualKeyboardTrace) OpacityChanged(opacity float64) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_OpacityChanged(ptr.Pointer(), C.double(opacity))
	}
}

func (ptr *QVirtualKeyboardTrace) Points(pos int, count int) []*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtVirtualKeyboard_PackedList) []*core.QVariant {
			out := make([]*core.QVariant, int(l.len))
			tmpList := NewQVirtualKeyboardTraceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__points_atList(i)
			}
			return out
		}(C.QVirtualKeyboardTrace_Points(ptr.Pointer(), C.int(int32(pos)), C.int(int32(count))))
	}
	return make([]*core.QVariant, 0)
}

func (ptr *QVirtualKeyboardTrace) SetCanceled(canceled bool) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_SetCanceled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(canceled))))
	}
}

func (ptr *QVirtualKeyboardTrace) SetChannelData(channel string, index int, data core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var channelC *C.char
		if channel != "" {
			channelC = C.CString(channel)
			defer C.free(unsafe.Pointer(channelC))
		}
		C.QVirtualKeyboardTrace_SetChannelData(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: channelC, len: C.longlong(len(channel))}, C.int(int32(index)), core.PointerFromQVariant(data))
	}
}

func (ptr *QVirtualKeyboardTrace) SetChannels(channels []string) {
	if ptr.Pointer() != nil {
		channelsC := C.CString(strings.Join(channels, "¡¦!"))
		defer C.free(unsafe.Pointer(channelsC))
		C.QVirtualKeyboardTrace_SetChannels(ptr.Pointer(), C.struct_QtVirtualKeyboard_PackedString{data: channelsC, len: C.longlong(len(strings.Join(channels, "¡¦!")))})
	}
}

func (ptr *QVirtualKeyboardTrace) SetFinal(final bool) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_SetFinal(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(final))))
	}
}

func (ptr *QVirtualKeyboardTrace) SetOpacity(opacity float64) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_SetOpacity(ptr.Pointer(), C.double(opacity))
	}
}

func (ptr *QVirtualKeyboardTrace) SetTraceId(id int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_SetTraceId(ptr.Pointer(), C.int(int32(id)))
	}
}

func (ptr *QVirtualKeyboardTrace) TraceId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVirtualKeyboardTrace_TraceId(ptr.Pointer())))
	}
	return 0
}

//export callbackQVirtualKeyboardTrace_TraceIdChanged
func callbackQVirtualKeyboardTrace_TraceIdChanged(ptr unsafe.Pointer, traceId C.int) {
	if signal := qt.GetSignal(ptr, "traceIdChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(traceId)))
	}

}

func (ptr *QVirtualKeyboardTrace) ConnectTraceIdChanged(f func(traceId int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "traceIdChanged") {
			C.QVirtualKeyboardTrace_ConnectTraceIdChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "traceIdChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "traceIdChanged"); signal != nil {
			f := func(traceId int) {
				(*(*func(int))(signal))(traceId)
				f(traceId)
			}
			qt.ConnectSignal(ptr.Pointer(), "traceIdChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "traceIdChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVirtualKeyboardTrace) DisconnectTraceIdChanged() {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_DisconnectTraceIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "traceIdChanged")
	}
}

func (ptr *QVirtualKeyboardTrace) TraceIdChanged(traceId int) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_TraceIdChanged(ptr.Pointer(), C.int(int32(traceId)))
	}
}

func (ptr *QVirtualKeyboardTrace) __channelData_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardTrace___channelData_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardTrace) __channelData_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace___channelData_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QVirtualKeyboardTrace) __channelData_newList() unsafe.Pointer {
	return C.QVirtualKeyboardTrace___channelData_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardTrace) __points_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QVirtualKeyboardTrace___points_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardTrace) __points_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace___points_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QVirtualKeyboardTrace) __points_newList() unsafe.Pointer {
	return C.QVirtualKeyboardTrace___points_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardTrace) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardTrace___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardTrace) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardTrace) __children_newList() unsafe.Pointer {
	return C.QVirtualKeyboardTrace___children_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardTrace) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QVirtualKeyboardTrace___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardTrace) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QVirtualKeyboardTrace) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QVirtualKeyboardTrace___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardTrace) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardTrace___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardTrace) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardTrace) __findChildren_newList() unsafe.Pointer {
	return C.QVirtualKeyboardTrace___findChildren_newList(ptr.Pointer())
}

func (ptr *QVirtualKeyboardTrace) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QVirtualKeyboardTrace___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVirtualKeyboardTrace) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QVirtualKeyboardTrace) __findChildren_newList3() unsafe.Pointer {
	return C.QVirtualKeyboardTrace___findChildren_newList3(ptr.Pointer())
}

//export callbackQVirtualKeyboardTrace_ChildEvent
func callbackQVirtualKeyboardTrace_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVirtualKeyboardTraceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardTrace) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQVirtualKeyboardTrace_ConnectNotify
func callbackQVirtualKeyboardTrace_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardTraceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardTrace) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardTrace_CustomEvent
func callbackQVirtualKeyboardTrace_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQVirtualKeyboardTraceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardTrace) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQVirtualKeyboardTrace_DeleteLater
func callbackQVirtualKeyboardTrace_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVirtualKeyboardTraceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVirtualKeyboardTrace) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVirtualKeyboardTrace_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQVirtualKeyboardTrace_Destroyed
func callbackQVirtualKeyboardTrace_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQVirtualKeyboardTrace_DisconnectNotify
func callbackQVirtualKeyboardTrace_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVirtualKeyboardTraceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVirtualKeyboardTrace) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVirtualKeyboardTrace_Event
func callbackQVirtualKeyboardTrace_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardTraceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QVirtualKeyboardTrace) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardTrace_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardTrace_EventFilter
func callbackQVirtualKeyboardTrace_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVirtualKeyboardTraceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QVirtualKeyboardTrace) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVirtualKeyboardTrace_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQVirtualKeyboardTrace_MetaObject
func callbackQVirtualKeyboardTrace_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQVirtualKeyboardTraceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QVirtualKeyboardTrace) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVirtualKeyboardTrace_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVirtualKeyboardTrace_ObjectNameChanged
func callbackQVirtualKeyboardTrace_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtVirtualKeyboard_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQVirtualKeyboardTrace_TimerEvent
func callbackQVirtualKeyboardTrace_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVirtualKeyboardTraceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVirtualKeyboardTrace) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVirtualKeyboardTrace_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func init() {
	qt.ItfMap["virtualkeyboard.QVirtualKeyboardAbstractInputMethod_ITF"] = QVirtualKeyboardAbstractInputMethod{}
	qt.FuncMap["virtualkeyboard.NewQVirtualKeyboardAbstractInputMethod"] = NewQVirtualKeyboardAbstractInputMethod
	qt.ItfMap["virtualkeyboard.QVirtualKeyboardExtensionPlugin_ITF"] = QVirtualKeyboardExtensionPlugin{}
	qt.ItfMap["virtualkeyboard.QVirtualKeyboardInputContext_ITF"] = QVirtualKeyboardInputContext{}
	qt.ItfMap["virtualkeyboard.QVirtualKeyboardInputEngine_ITF"] = QVirtualKeyboardInputEngine{}
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Lower"] = int64(QVirtualKeyboardInputEngine__Lower)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Upper"] = int64(QVirtualKeyboardInputEngine__Upper)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Latin"] = int64(QVirtualKeyboardInputEngine__Latin)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Numeric"] = int64(QVirtualKeyboardInputEngine__Numeric)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Dialable"] = int64(QVirtualKeyboardInputEngine__Dialable)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Pinyin"] = int64(QVirtualKeyboardInputEngine__Pinyin)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Cangjie"] = int64(QVirtualKeyboardInputEngine__Cangjie)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Zhuyin"] = int64(QVirtualKeyboardInputEngine__Zhuyin)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Hangul"] = int64(QVirtualKeyboardInputEngine__Hangul)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Hiragana"] = int64(QVirtualKeyboardInputEngine__Hiragana)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Katakana"] = int64(QVirtualKeyboardInputEngine__Katakana)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__FullwidthLatin"] = int64(QVirtualKeyboardInputEngine__FullwidthLatin)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Greek"] = int64(QVirtualKeyboardInputEngine__Greek)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Cyrillic"] = int64(QVirtualKeyboardInputEngine__Cyrillic)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Arabic"] = int64(QVirtualKeyboardInputEngine__Arabic)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Hebrew"] = int64(QVirtualKeyboardInputEngine__Hebrew)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__ChineseHandwriting"] = int64(QVirtualKeyboardInputEngine__ChineseHandwriting)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__JapaneseHandwriting"] = int64(QVirtualKeyboardInputEngine__JapaneseHandwriting)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__KoreanHandwriting"] = int64(QVirtualKeyboardInputEngine__KoreanHandwriting)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Thai"] = int64(QVirtualKeyboardInputEngine__Thai)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__None"] = int64(QVirtualKeyboardInputEngine__None)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__PatternRecognitionDisabled"] = int64(QVirtualKeyboardInputEngine__PatternRecognitionDisabled)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__Handwriting"] = int64(QVirtualKeyboardInputEngine__Handwriting)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__HandwritingRecoginition"] = int64(QVirtualKeyboardInputEngine__HandwritingRecoginition)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__WordBeforeCursor"] = int64(QVirtualKeyboardInputEngine__WordBeforeCursor)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__WordAfterCursor"] = int64(QVirtualKeyboardInputEngine__WordAfterCursor)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardInputEngine__WordAtCursor"] = int64(QVirtualKeyboardInputEngine__WordAtCursor)
	qt.ItfMap["virtualkeyboard.QVirtualKeyboardSelectionListModel_ITF"] = QVirtualKeyboardSelectionListModel{}
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardSelectionListModel__WordCandidateList"] = int64(QVirtualKeyboardSelectionListModel__WordCandidateList)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardSelectionListModel__Display"] = int64(QVirtualKeyboardSelectionListModel__Display)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardSelectionListModel__DisplayRole"] = int64(QVirtualKeyboardSelectionListModel__DisplayRole)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardSelectionListModel__Dictionary"] = int64(QVirtualKeyboardSelectionListModel__Dictionary)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardSelectionListModel__CanRemoveSuggestion"] = int64(QVirtualKeyboardSelectionListModel__CanRemoveSuggestion)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardSelectionListModel__Default"] = int64(QVirtualKeyboardSelectionListModel__Default)
	qt.EnumMap["virtualkeyboard.QVirtualKeyboardSelectionListModel__User"] = int64(QVirtualKeyboardSelectionListModel__User)
	qt.ItfMap["virtualkeyboard.QVirtualKeyboardTrace_ITF"] = QVirtualKeyboardTrace{}
}
