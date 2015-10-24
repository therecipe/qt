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

type QProgressBarITF interface {
	QWidgetITF
	QProgressBarPTR() *QProgressBar
}

func PointerFromQProgressBar(ptr QProgressBarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProgressBarPTR().Pointer()
	}
	return nil
}

func QProgressBarFromPointer(ptr unsafe.Pointer) *QProgressBar {
	var n = new(QProgressBar)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QProgressBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProgressBar) QProgressBarPTR() *QProgressBar {
	return ptr
}

//QProgressBar::Direction
type QProgressBar__Direction int

var (
	QProgressBar__TopToBottom = QProgressBar__Direction(0)
	QProgressBar__BottomToTop = QProgressBar__Direction(1)
)

func (ptr *QProgressBar) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QProgressBar_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProgressBar) Format() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressBar_Format(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QProgressBar) InvertedAppearance() bool {
	if ptr.Pointer() != nil {
		return C.QProgressBar_InvertedAppearance(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QProgressBar) IsTextVisible() bool {
	if ptr.Pointer() != nil {
		return C.QProgressBar_IsTextVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QProgressBar) Maximum() int {
	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Maximum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProgressBar) Minimum() int {
	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Minimum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProgressBar) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QProgressBar_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProgressBar) ResetFormat() {
	if ptr.Pointer() != nil {
		C.QProgressBar_ResetFormat(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QProgressBar) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QProgressBar) SetFormat(format string) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.CString(format))
	}
}

func (ptr *QProgressBar) SetInvertedAppearance(invert bool) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetInvertedAppearance(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(invert)))
	}
}

func (ptr *QProgressBar) SetMaximum(maximum int) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetMaximum(C.QtObjectPtr(ptr.Pointer()), C.int(maximum))
	}
}

func (ptr *QProgressBar) SetMinimum(minimum int) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetMinimum(C.QtObjectPtr(ptr.Pointer()), C.int(minimum))
	}
}

func (ptr *QProgressBar) SetOrientation(v core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QProgressBar) SetTextDirection(textDirection QProgressBar__Direction) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetTextDirection(C.QtObjectPtr(ptr.Pointer()), C.int(textDirection))
	}
}

func (ptr *QProgressBar) SetTextVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetTextVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QProgressBar) SetValue(value int) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetValue(C.QtObjectPtr(ptr.Pointer()), C.int(value))
	}
}

func (ptr *QProgressBar) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressBar_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QProgressBar) TextDirection() QProgressBar__Direction {
	if ptr.Pointer() != nil {
		return QProgressBar__Direction(C.QProgressBar_TextDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProgressBar) Value() int {
	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Value(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQProgressBar(parent QWidgetITF) *QProgressBar {
	return QProgressBarFromPointer(unsafe.Pointer(C.QProgressBar_NewQProgressBar(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QProgressBar) Reset() {
	if ptr.Pointer() != nil {
		C.QProgressBar_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QProgressBar) SetRange(minimum int, maximum int) {
	if ptr.Pointer() != nil {
		C.QProgressBar_SetRange(C.QtObjectPtr(ptr.Pointer()), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QProgressBar) ConnectValueChanged(f func(value int)) {
	if ptr.Pointer() != nil {
		C.QProgressBar_ConnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QProgressBar) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QProgressBar_DisconnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQProgressBarValueChanged
func callbackQProgressBarValueChanged(ptrName *C.char, value C.int) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(int))(int(value))
}

func (ptr *QProgressBar) DestroyQProgressBar() {
	if ptr.Pointer() != nil {
		C.QProgressBar_DestroyQProgressBar(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
