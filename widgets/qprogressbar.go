package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QProgressBar struct {
	QWidget
}

type QProgressBar_ITF interface {
	QWidget_ITF
	QProgressBar_PTR() *QProgressBar
}

func PointerFromQProgressBar(ptr QProgressBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProgressBar_PTR().Pointer()
	}
	return nil
}

func NewQProgressBarFromPointer(ptr unsafe.Pointer) *QProgressBar {
	var n = new(QProgressBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProgressBar_") {
		n.SetObjectName("QProgressBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QProgressBar) QProgressBar_PTR() *QProgressBar {
	return ptr
}

//QProgressBar::Direction
type QProgressBar__Direction int64

const (
	QProgressBar__TopToBottom = QProgressBar__Direction(0)
	QProgressBar__BottomToTop = QProgressBar__Direction(1)
)

func (ptr *QProgressBar) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QProgressBar::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QProgressBar_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Format() string {
	defer qt.Recovering("QProgressBar::format")

	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressBar_Format(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressBar) InvertedAppearance() bool {
	defer qt.Recovering("QProgressBar::invertedAppearance")

	if ptr.Pointer() != nil {
		return C.QProgressBar_InvertedAppearance(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressBar) IsTextVisible() bool {
	defer qt.Recovering("QProgressBar::isTextVisible")

	if ptr.Pointer() != nil {
		return C.QProgressBar_IsTextVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressBar) Maximum() int {
	defer qt.Recovering("QProgressBar::maximum")

	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Minimum() int {
	defer qt.Recovering("QProgressBar::minimum")

	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QProgressBar::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QProgressBar_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) ResetFormat() {
	defer qt.Recovering("QProgressBar::resetFormat")

	if ptr.Pointer() != nil {
		C.QProgressBar_ResetFormat(ptr.Pointer())
	}
}

func (ptr *QProgressBar) SetAlignment(alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QProgressBar::setAlignment")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QProgressBar) SetFormat(format string) {
	defer qt.Recovering("QProgressBar::setFormat")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetFormat(ptr.Pointer(), C.CString(format))
	}
}

func (ptr *QProgressBar) SetInvertedAppearance(invert bool) {
	defer qt.Recovering("QProgressBar::setInvertedAppearance")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetInvertedAppearance(ptr.Pointer(), C.int(qt.GoBoolToInt(invert)))
	}
}

func (ptr *QProgressBar) SetMaximum(maximum int) {
	defer qt.Recovering("QProgressBar::setMaximum")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetMaximum(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QProgressBar) SetMinimum(minimum int) {
	defer qt.Recovering("QProgressBar::setMinimum")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetMinimum(ptr.Pointer(), C.int(minimum))
	}
}

func (ptr *QProgressBar) SetOrientation(v core.Qt__Orientation) {
	defer qt.Recovering("QProgressBar::setOrientation")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetOrientation(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QProgressBar) SetTextDirection(textDirection QProgressBar__Direction) {
	defer qt.Recovering("QProgressBar::setTextDirection")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetTextDirection(ptr.Pointer(), C.int(textDirection))
	}
}

func (ptr *QProgressBar) SetTextVisible(visible bool) {
	defer qt.Recovering("QProgressBar::setTextVisible")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetTextVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QProgressBar) SetValue(value int) {
	defer qt.Recovering("QProgressBar::setValue")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetValue(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QProgressBar) Text() string {
	defer qt.Recovering("QProgressBar::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressBar_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressBar) TextDirection() QProgressBar__Direction {
	defer qt.Recovering("QProgressBar::textDirection")

	if ptr.Pointer() != nil {
		return QProgressBar__Direction(C.QProgressBar_TextDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Value() int {
	defer qt.Recovering("QProgressBar::value")

	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Value(ptr.Pointer()))
	}
	return 0
}

func NewQProgressBar(parent QWidget_ITF) *QProgressBar {
	defer qt.Recovering("QProgressBar::QProgressBar")

	return NewQProgressBarFromPointer(C.QProgressBar_NewQProgressBar(PointerFromQWidget(parent)))
}

func (ptr *QProgressBar) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QProgressBar::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QProgressBar_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProgressBar) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QProgressBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QProgressBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQProgressBarPaintEvent
func callbackQProgressBarPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressBar::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QProgressBar) Reset() {
	defer qt.Recovering("QProgressBar::reset")

	if ptr.Pointer() != nil {
		C.QProgressBar_Reset(ptr.Pointer())
	}
}

func (ptr *QProgressBar) SetRange(minimum int, maximum int) {
	defer qt.Recovering("QProgressBar::setRange")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetRange(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QProgressBar) SizeHint() *core.QSize {
	defer qt.Recovering("QProgressBar::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QProgressBar_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProgressBar) ConnectValueChanged(f func(value int)) {
	defer qt.Recovering("connect QProgressBar::valueChanged")

	if ptr.Pointer() != nil {
		C.QProgressBar_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QProgressBar) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QProgressBar::valueChanged")

	if ptr.Pointer() != nil {
		C.QProgressBar_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQProgressBarValueChanged
func callbackQProgressBarValueChanged(ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QProgressBar::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged"); signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QProgressBar) DestroyQProgressBar() {
	defer qt.Recovering("QProgressBar::~QProgressBar")

	if ptr.Pointer() != nil {
		C.QProgressBar_DestroyQProgressBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
