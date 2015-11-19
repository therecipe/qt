package widgets

//#include "qprogressbar.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QProgressBar_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QProgressBar_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Format() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressBar_Format(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressBar) InvertedAppearance() bool {
	if ptr.Pointer() != nil {
		return C.QProgressBar_InvertedAppearance(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressBar) IsTextVisible() bool {
	if ptr.Pointer() != nil {
		return C.QProgressBar_IsTextVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressBar) Maximum() int {
	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Minimum() int {
	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QProgressBar_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) ResetFormat() {
	if ptr.Pointer() != nil {
		C.QProgressBar_ResetFormat(ptr.Pointer())
	}
}

func (ptr *QProgressBar) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QProgressBar) SetFormat(format string) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetFormat(ptr.Pointer(), C.CString(format))
	}
}

func (ptr *QProgressBar) SetInvertedAppearance(invert bool) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetInvertedAppearance(ptr.Pointer(), C.int(qt.GoBoolToInt(invert)))
	}
}

func (ptr *QProgressBar) SetMaximum(maximum int) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetMaximum(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QProgressBar) SetMinimum(minimum int) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetMinimum(ptr.Pointer(), C.int(minimum))
	}
}

func (ptr *QProgressBar) SetOrientation(v core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetOrientation(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QProgressBar) SetTextDirection(textDirection QProgressBar__Direction) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetTextDirection(ptr.Pointer(), C.int(textDirection))
	}
}

func (ptr *QProgressBar) SetTextVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetTextVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QProgressBar) SetValue(value int) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetValue(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QProgressBar) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressBar_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressBar) TextDirection() QProgressBar__Direction {
	if ptr.Pointer() != nil {
		return QProgressBar__Direction(C.QProgressBar_TextDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Value() int {
	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Value(ptr.Pointer()))
	}
	return 0
}

func NewQProgressBar(parent QWidget_ITF) *QProgressBar {
	return NewQProgressBarFromPointer(C.QProgressBar_NewQProgressBar(PointerFromQWidget(parent)))
}

func (ptr *QProgressBar) Reset() {
	if ptr.Pointer() != nil {
		C.QProgressBar_Reset(ptr.Pointer())
	}
}

func (ptr *QProgressBar) SetRange(minimum int, maximum int) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetRange(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QProgressBar) ConnectValueChanged(f func(value int)) {
	if ptr.Pointer() != nil {
		C.QProgressBar_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QProgressBar) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QProgressBar_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQProgressBarValueChanged
func callbackQProgressBarValueChanged(ptrName *C.char, value C.int) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(int))(int(value))
}

func (ptr *QProgressBar) DestroyQProgressBar() {
	if ptr.Pointer() != nil {
		C.QProgressBar_DestroyQProgressBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
