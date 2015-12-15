package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QIcon struct {
	ptr unsafe.Pointer
}

type QIcon_ITF interface {
	QIcon_PTR() *QIcon
}

func (p *QIcon) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QIcon) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQIcon(ptr QIcon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIcon_PTR().Pointer()
	}
	return nil
}

func NewQIconFromPointer(ptr unsafe.Pointer) *QIcon {
	var n = new(QIcon)
	n.SetPointer(ptr)
	return n
}

func (ptr *QIcon) QIcon_PTR() *QIcon {
	return ptr
}

//QIcon::Mode
type QIcon__Mode int64

const (
	QIcon__Normal   = QIcon__Mode(0)
	QIcon__Disabled = QIcon__Mode(1)
	QIcon__Active   = QIcon__Mode(2)
	QIcon__Selected = QIcon__Mode(3)
)

//QIcon::State
type QIcon__State int64

const (
	QIcon__On  = QIcon__State(0)
	QIcon__Off = QIcon__State(1)
)

func NewQIcon() *QIcon {
	defer qt.Recovering("QIcon::QIcon")

	return NewQIconFromPointer(C.QIcon_NewQIcon())
}

func NewQIcon4(other QIcon_ITF) *QIcon {
	defer qt.Recovering("QIcon::QIcon")

	return NewQIconFromPointer(C.QIcon_NewQIcon4(PointerFromQIcon(other)))
}

func NewQIcon6(engine QIconEngine_ITF) *QIcon {
	defer qt.Recovering("QIcon::QIcon")

	return NewQIconFromPointer(C.QIcon_NewQIcon6(PointerFromQIconEngine(engine)))
}

func NewQIcon3(other QIcon_ITF) *QIcon {
	defer qt.Recovering("QIcon::QIcon")

	return NewQIconFromPointer(C.QIcon_NewQIcon3(PointerFromQIcon(other)))
}

func NewQIcon2(pixmap QPixmap_ITF) *QIcon {
	defer qt.Recovering("QIcon::QIcon")

	return NewQIconFromPointer(C.QIcon_NewQIcon2(PointerFromQPixmap(pixmap)))
}

func NewQIcon5(fileName string) *QIcon {
	defer qt.Recovering("QIcon::QIcon")

	return NewQIconFromPointer(C.QIcon_NewQIcon5(C.CString(fileName)))
}

func (ptr *QIcon) AddFile(fileName string, size core.QSize_ITF, mode QIcon__Mode, state QIcon__State) {
	defer qt.Recovering("QIcon::addFile")

	if ptr.Pointer() != nil {
		C.QIcon_AddFile(ptr.Pointer(), C.CString(fileName), core.PointerFromQSize(size), C.int(mode), C.int(state))
	}
}

func (ptr *QIcon) AddPixmap(pixmap QPixmap_ITF, mode QIcon__Mode, state QIcon__State) {
	defer qt.Recovering("QIcon::addPixmap")

	if ptr.Pointer() != nil {
		C.QIcon_AddPixmap(ptr.Pointer(), PointerFromQPixmap(pixmap), C.int(mode), C.int(state))
	}
}

func QIcon_HasThemeIcon(name string) bool {
	defer qt.Recovering("QIcon::hasThemeIcon")

	return C.QIcon_QIcon_HasThemeIcon(C.CString(name)) != 0
}

func (ptr *QIcon) IsNull() bool {
	defer qt.Recovering("QIcon::isNull")

	if ptr.Pointer() != nil {
		return C.QIcon_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIcon) Name() string {
	defer qt.Recovering("QIcon::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QIcon_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIcon) Paint(painter QPainter_ITF, rect core.QRect_ITF, alignment core.Qt__AlignmentFlag, mode QIcon__Mode, state QIcon__State) {
	defer qt.Recovering("QIcon::paint")

	if ptr.Pointer() != nil {
		C.QIcon_Paint(ptr.Pointer(), PointerFromQPainter(painter), core.PointerFromQRect(rect), C.int(alignment), C.int(mode), C.int(state))
	}
}

func (ptr *QIcon) Paint2(painter QPainter_ITF, x int, y int, w int, h int, alignment core.Qt__AlignmentFlag, mode QIcon__Mode, state QIcon__State) {
	defer qt.Recovering("QIcon::paint")

	if ptr.Pointer() != nil {
		C.QIcon_Paint2(ptr.Pointer(), PointerFromQPainter(painter), C.int(x), C.int(y), C.int(w), C.int(h), C.int(alignment), C.int(mode), C.int(state))
	}
}

func QIcon_SetThemeName(name string) {
	defer qt.Recovering("QIcon::setThemeName")

	C.QIcon_QIcon_SetThemeName(C.CString(name))
}

func QIcon_SetThemeSearchPaths(paths []string) {
	defer qt.Recovering("QIcon::setThemeSearchPaths")

	C.QIcon_QIcon_SetThemeSearchPaths(C.CString(strings.Join(paths, ",,,")))
}

func (ptr *QIcon) Swap(other QIcon_ITF) {
	defer qt.Recovering("QIcon::swap")

	if ptr.Pointer() != nil {
		C.QIcon_Swap(ptr.Pointer(), PointerFromQIcon(other))
	}
}

func QIcon_ThemeName() string {
	defer qt.Recovering("QIcon::themeName")

	return C.GoString(C.QIcon_QIcon_ThemeName())
}

func QIcon_ThemeSearchPaths() []string {
	defer qt.Recovering("QIcon::themeSearchPaths")

	return strings.Split(C.GoString(C.QIcon_QIcon_ThemeSearchPaths()), ",,,")
}

func (ptr *QIcon) DestroyQIcon() {
	defer qt.Recovering("QIcon::~QIcon")

	if ptr.Pointer() != nil {
		C.QIcon_DestroyQIcon(ptr.Pointer())
	}
}
