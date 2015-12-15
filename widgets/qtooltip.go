package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QToolTip struct {
	ptr unsafe.Pointer
}

type QToolTip_ITF interface {
	QToolTip_PTR() *QToolTip
}

func (p *QToolTip) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QToolTip) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQToolTip(ptr QToolTip_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolTip_PTR().Pointer()
	}
	return nil
}

func NewQToolTipFromPointer(ptr unsafe.Pointer) *QToolTip {
	var n = new(QToolTip)
	n.SetPointer(ptr)
	return n
}

func (ptr *QToolTip) QToolTip_PTR() *QToolTip {
	return ptr
}

func QToolTip_HideText() {
	defer qt.Recovering("QToolTip::hideText")

	C.QToolTip_QToolTip_HideText()
}

func QToolTip_IsVisible() bool {
	defer qt.Recovering("QToolTip::isVisible")

	return C.QToolTip_QToolTip_IsVisible() != 0
}

func QToolTip_SetFont(font gui.QFont_ITF) {
	defer qt.Recovering("QToolTip::setFont")

	C.QToolTip_QToolTip_SetFont(gui.PointerFromQFont(font))
}

func QToolTip_SetPalette(palette gui.QPalette_ITF) {
	defer qt.Recovering("QToolTip::setPalette")

	C.QToolTip_QToolTip_SetPalette(gui.PointerFromQPalette(palette))
}

func QToolTip_ShowText3(pos core.QPoint_ITF, text string, w QWidget_ITF) {
	defer qt.Recovering("QToolTip::showText")

	C.QToolTip_QToolTip_ShowText3(core.PointerFromQPoint(pos), C.CString(text), PointerFromQWidget(w))
}

func QToolTip_ShowText(pos core.QPoint_ITF, text string, w QWidget_ITF, rect core.QRect_ITF) {
	defer qt.Recovering("QToolTip::showText")

	C.QToolTip_QToolTip_ShowText(core.PointerFromQPoint(pos), C.CString(text), PointerFromQWidget(w), core.PointerFromQRect(rect))
}

func QToolTip_ShowText2(pos core.QPoint_ITF, text string, w QWidget_ITF, rect core.QRect_ITF, msecDisplayTime int) {
	defer qt.Recovering("QToolTip::showText")

	C.QToolTip_QToolTip_ShowText2(core.PointerFromQPoint(pos), C.CString(text), PointerFromQWidget(w), core.PointerFromQRect(rect), C.int(msecDisplayTime))
}

func QToolTip_Text() string {
	defer qt.Recovering("QToolTip::text")

	return C.GoString(C.QToolTip_QToolTip_Text())
}
