package gui

//#include "qicon.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QIcon struct {
	ptr unsafe.Pointer
}

type QIconITF interface {
	QIconPTR() *QIcon
}

func (p *QIcon) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QIcon) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQIcon(ptr QIconITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIconPTR().Pointer()
	}
	return nil
}

func QIconFromPointer(ptr unsafe.Pointer) *QIcon {
	var n = new(QIcon)
	n.SetPointer(ptr)
	return n
}

func (ptr *QIcon) QIconPTR() *QIcon {
	return ptr
}

//QIcon::Mode
type QIcon__Mode int

var (
	QIcon__Normal   = QIcon__Mode(0)
	QIcon__Disabled = QIcon__Mode(1)
	QIcon__Active   = QIcon__Mode(2)
	QIcon__Selected = QIcon__Mode(3)
)

//QIcon::State
type QIcon__State int

var (
	QIcon__On  = QIcon__State(0)
	QIcon__Off = QIcon__State(1)
)

func NewQIcon() *QIcon {
	return QIconFromPointer(unsafe.Pointer(C.QIcon_NewQIcon()))
}

func NewQIcon4(other QIconITF) *QIcon {
	return QIconFromPointer(unsafe.Pointer(C.QIcon_NewQIcon4(C.QtObjectPtr(PointerFromQIcon(other)))))
}

func NewQIcon6(engine QIconEngineITF) *QIcon {
	return QIconFromPointer(unsafe.Pointer(C.QIcon_NewQIcon6(C.QtObjectPtr(PointerFromQIconEngine(engine)))))
}

func NewQIcon3(other QIconITF) *QIcon {
	return QIconFromPointer(unsafe.Pointer(C.QIcon_NewQIcon3(C.QtObjectPtr(PointerFromQIcon(other)))))
}

func NewQIcon2(pixmap QPixmapITF) *QIcon {
	return QIconFromPointer(unsafe.Pointer(C.QIcon_NewQIcon2(C.QtObjectPtr(PointerFromQPixmap(pixmap)))))
}

func NewQIcon5(fileName string) *QIcon {
	return QIconFromPointer(unsafe.Pointer(C.QIcon_NewQIcon5(C.CString(fileName))))
}

func (ptr *QIcon) AddFile(fileName string, size core.QSizeITF, mode QIcon__Mode, state QIcon__State) {
	if ptr.Pointer() != nil {
		C.QIcon_AddFile(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName), C.QtObjectPtr(core.PointerFromQSize(size)), C.int(mode), C.int(state))
	}
}

func (ptr *QIcon) AddPixmap(pixmap QPixmapITF, mode QIcon__Mode, state QIcon__State) {
	if ptr.Pointer() != nil {
		C.QIcon_AddPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.int(mode), C.int(state))
	}
}

func QIcon_HasThemeIcon(name string) bool {
	return C.QIcon_QIcon_HasThemeIcon(C.CString(name)) != 0
}

func (ptr *QIcon) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QIcon_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QIcon) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QIcon_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QIcon) Paint(painter QPainterITF, rect core.QRectITF, alignment core.Qt__AlignmentFlag, mode QIcon__Mode, state QIcon__State) {
	if ptr.Pointer() != nil {
		C.QIcon_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQRect(rect)), C.int(alignment), C.int(mode), C.int(state))
	}
}

func (ptr *QIcon) Paint2(painter QPainterITF, x int, y int, w int, h int, alignment core.Qt__AlignmentFlag, mode QIcon__Mode, state QIcon__State) {
	if ptr.Pointer() != nil {
		C.QIcon_Paint2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainter(painter)), C.int(x), C.int(y), C.int(w), C.int(h), C.int(alignment), C.int(mode), C.int(state))
	}
}

func QIcon_SetThemeName(name string) {
	C.QIcon_QIcon_SetThemeName(C.CString(name))
}

func QIcon_SetThemeSearchPaths(paths []string) {
	C.QIcon_QIcon_SetThemeSearchPaths(C.CString(strings.Join(paths, "|")))
}

func (ptr *QIcon) Swap(other QIconITF) {
	if ptr.Pointer() != nil {
		C.QIcon_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQIcon(other)))
	}
}

func QIcon_ThemeName() string {
	return C.GoString(C.QIcon_QIcon_ThemeName())
}

func QIcon_ThemeSearchPaths() []string {
	return strings.Split(C.GoString(C.QIcon_QIcon_ThemeSearchPaths()), "|")
}

func (ptr *QIcon) DestroyQIcon() {
	if ptr.Pointer() != nil {
		C.QIcon_DestroyQIcon(C.QtObjectPtr(ptr.Pointer()))
	}
}
