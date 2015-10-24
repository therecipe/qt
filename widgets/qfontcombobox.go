package widgets

//#include "qfontcombobox.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QFontComboBox struct {
	QComboBox
}

type QFontComboBoxITF interface {
	QComboBoxITF
	QFontComboBoxPTR() *QFontComboBox
}

func PointerFromQFontComboBox(ptr QFontComboBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontComboBoxPTR().Pointer()
	}
	return nil
}

func QFontComboBoxFromPointer(ptr unsafe.Pointer) *QFontComboBox {
	var n = new(QFontComboBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFontComboBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFontComboBox) QFontComboBoxPTR() *QFontComboBox {
	return ptr
}

//QFontComboBox::FontFilter
type QFontComboBox__FontFilter int

var (
	QFontComboBox__AllFonts          = QFontComboBox__FontFilter(0)
	QFontComboBox__ScalableFonts     = QFontComboBox__FontFilter(0x1)
	QFontComboBox__NonScalableFonts  = QFontComboBox__FontFilter(0x2)
	QFontComboBox__MonospacedFonts   = QFontComboBox__FontFilter(0x4)
	QFontComboBox__ProportionalFonts = QFontComboBox__FontFilter(0x8)
)

func (ptr *QFontComboBox) FontFilters() QFontComboBox__FontFilter {
	if ptr.Pointer() != nil {
		return QFontComboBox__FontFilter(C.QFontComboBox_FontFilters(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontComboBox) SetCurrentFont(font gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QFontComboBox_SetCurrentFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(font)))
	}
}

func (ptr *QFontComboBox) SetFontFilters(filters QFontComboBox__FontFilter) {
	if ptr.Pointer() != nil {
		C.QFontComboBox_SetFontFilters(C.QtObjectPtr(ptr.Pointer()), C.int(filters))
	}
}

func (ptr *QFontComboBox) SetWritingSystem(script gui.QFontDatabase__WritingSystem) {
	if ptr.Pointer() != nil {
		C.QFontComboBox_SetWritingSystem(C.QtObjectPtr(ptr.Pointer()), C.int(script))
	}
}

func (ptr *QFontComboBox) WritingSystem() gui.QFontDatabase__WritingSystem {
	if ptr.Pointer() != nil {
		return gui.QFontDatabase__WritingSystem(C.QFontComboBox_WritingSystem(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQFontComboBox(parent QWidgetITF) *QFontComboBox {
	return QFontComboBoxFromPointer(unsafe.Pointer(C.QFontComboBox_NewQFontComboBox(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QFontComboBox) DestroyQFontComboBox() {
	if ptr.Pointer() != nil {
		C.QFontComboBox_DestroyQFontComboBox(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
