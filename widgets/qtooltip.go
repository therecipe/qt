package widgets

//#include "qtooltip.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QToolTip struct {
	ptr unsafe.Pointer
}

type QToolTipITF interface {
	QToolTipPTR() *QToolTip
}

func (p *QToolTip) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QToolTip) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQToolTip(ptr QToolTipITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolTipPTR().Pointer()
	}
	return nil
}

func QToolTipFromPointer(ptr unsafe.Pointer) *QToolTip {
	var n = new(QToolTip)
	n.SetPointer(ptr)
	return n
}

func (ptr *QToolTip) QToolTipPTR() *QToolTip {
	return ptr
}

func QToolTip_HideText() {
	C.QToolTip_QToolTip_HideText()
}

func QToolTip_IsVisible() bool {
	return C.QToolTip_QToolTip_IsVisible() != 0
}

func QToolTip_SetFont(font gui.QFontITF) {
	C.QToolTip_QToolTip_SetFont(C.QtObjectPtr(gui.PointerFromQFont(font)))
}

func QToolTip_SetPalette(palette gui.QPaletteITF) {
	C.QToolTip_QToolTip_SetPalette(C.QtObjectPtr(gui.PointerFromQPalette(palette)))
}

func QToolTip_ShowText3(pos core.QPointITF, text string, w QWidgetITF) {
	C.QToolTip_QToolTip_ShowText3(C.QtObjectPtr(core.PointerFromQPoint(pos)), C.CString(text), C.QtObjectPtr(PointerFromQWidget(w)))
}

func QToolTip_ShowText(pos core.QPointITF, text string, w QWidgetITF, rect core.QRectITF) {
	C.QToolTip_QToolTip_ShowText(C.QtObjectPtr(core.PointerFromQPoint(pos)), C.CString(text), C.QtObjectPtr(PointerFromQWidget(w)), C.QtObjectPtr(core.PointerFromQRect(rect)))
}

func QToolTip_ShowText2(pos core.QPointITF, text string, w QWidgetITF, rect core.QRectITF, msecDisplayTime int) {
	C.QToolTip_QToolTip_ShowText2(C.QtObjectPtr(core.PointerFromQPoint(pos)), C.CString(text), C.QtObjectPtr(PointerFromQWidget(w)), C.QtObjectPtr(core.PointerFromQRect(rect)), C.int(msecDisplayTime))
}

func QToolTip_Text() string {
	return C.GoString(C.QToolTip_QToolTip_Text())
}
