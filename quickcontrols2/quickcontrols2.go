// +build !minimal

package quickcontrols2

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "quickcontrols2.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtQuickControls2_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtQuickControls2_PackedString) []byte {
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

type QQuickStyle struct {
	ptr unsafe.Pointer
}

type QQuickStyle_ITF interface {
	QQuickStyle_PTR() *QQuickStyle
}

func (ptr *QQuickStyle) QQuickStyle_PTR() *QQuickStyle {
	return ptr
}

func (ptr *QQuickStyle) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQuickStyle) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQuickStyle(ptr QQuickStyle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickStyle_PTR().Pointer()
	}
	return nil
}

func NewQQuickStyleFromPointer(ptr unsafe.Pointer) (n *QQuickStyle) {
	n = new(QQuickStyle)
	n.SetPointer(ptr)
	return
}
func (ptr *QQuickStyle) DestroyQQuickStyle() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func QQuickStyle_AddStylePath(path string) {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	C.QQuickStyle_QQuickStyle_AddStylePath(C.struct_QtQuickControls2_PackedString{data: pathC, len: C.longlong(len(path))})
}

func (ptr *QQuickStyle) AddStylePath(path string) {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	C.QQuickStyle_QQuickStyle_AddStylePath(C.struct_QtQuickControls2_PackedString{data: pathC, len: C.longlong(len(path))})
}

func QQuickStyle_AvailableStyles() []string {
	return unpackStringList(cGoUnpackString(C.QQuickStyle_QQuickStyle_AvailableStyles()))
}

func (ptr *QQuickStyle) AvailableStyles() []string {
	return unpackStringList(cGoUnpackString(C.QQuickStyle_QQuickStyle_AvailableStyles()))
}

func QQuickStyle_Name() string {
	return cGoUnpackString(C.QQuickStyle_QQuickStyle_Name())
}

func (ptr *QQuickStyle) Name() string {
	return cGoUnpackString(C.QQuickStyle_QQuickStyle_Name())
}

func QQuickStyle_Path() string {
	return cGoUnpackString(C.QQuickStyle_QQuickStyle_Path())
}

func (ptr *QQuickStyle) Path() string {
	return cGoUnpackString(C.QQuickStyle_QQuickStyle_Path())
}

func QQuickStyle_SetFallbackStyle(style string) {
	var styleC *C.char
	if style != "" {
		styleC = C.CString(style)
		defer C.free(unsafe.Pointer(styleC))
	}
	C.QQuickStyle_QQuickStyle_SetFallbackStyle(C.struct_QtQuickControls2_PackedString{data: styleC, len: C.longlong(len(style))})
}

func (ptr *QQuickStyle) SetFallbackStyle(style string) {
	var styleC *C.char
	if style != "" {
		styleC = C.CString(style)
		defer C.free(unsafe.Pointer(styleC))
	}
	C.QQuickStyle_QQuickStyle_SetFallbackStyle(C.struct_QtQuickControls2_PackedString{data: styleC, len: C.longlong(len(style))})
}

func QQuickStyle_SetStyle(style string) {
	var styleC *C.char
	if style != "" {
		styleC = C.CString(style)
		defer C.free(unsafe.Pointer(styleC))
	}
	C.QQuickStyle_QQuickStyle_SetStyle(C.struct_QtQuickControls2_PackedString{data: styleC, len: C.longlong(len(style))})
}

func (ptr *QQuickStyle) SetStyle(style string) {
	var styleC *C.char
	if style != "" {
		styleC = C.CString(style)
		defer C.free(unsafe.Pointer(styleC))
	}
	C.QQuickStyle_QQuickStyle_SetStyle(C.struct_QtQuickControls2_PackedString{data: styleC, len: C.longlong(len(style))})
}

func QQuickStyle_StylePathList() []string {
	return unpackStringList(cGoUnpackString(C.QQuickStyle_QQuickStyle_StylePathList()))
}

func (ptr *QQuickStyle) StylePathList() []string {
	return unpackStringList(cGoUnpackString(C.QQuickStyle_QQuickStyle_StylePathList()))
}

func init() {
	qt.ItfMap["quickcontrols2.QQuickStyle_ITF"] = QQuickStyle{}
	qt.FuncMap["quickcontrols2.QQuickStyle_AddStylePath"] = QQuickStyle_AddStylePath
	qt.FuncMap["quickcontrols2.QQuickStyle_AvailableStyles"] = QQuickStyle_AvailableStyles
	qt.FuncMap["quickcontrols2.QQuickStyle_Name"] = QQuickStyle_Name
	qt.FuncMap["quickcontrols2.QQuickStyle_Path"] = QQuickStyle_Path
	qt.FuncMap["quickcontrols2.QQuickStyle_SetFallbackStyle"] = QQuickStyle_SetFallbackStyle
	qt.FuncMap["quickcontrols2.QQuickStyle_SetStyle"] = QQuickStyle_SetStyle
	qt.FuncMap["quickcontrols2.QQuickStyle_StylePathList"] = QQuickStyle_StylePathList
}
