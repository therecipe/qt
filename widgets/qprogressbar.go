package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::alignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QProgressBar_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Format() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::format")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressBar_Format(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressBar) InvertedAppearance() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::invertedAppearance")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProgressBar_InvertedAppearance(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressBar) IsTextVisible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::isTextVisible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProgressBar_IsTextVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressBar) Maximum() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::maximum")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Minimum() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::minimum")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Orientation() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::orientation")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QProgressBar_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) ResetFormat() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::resetFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_ResetFormat(ptr.Pointer())
	}
}

func (ptr *QProgressBar) SetAlignment(alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QProgressBar) SetFormat(format string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::setFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_SetFormat(ptr.Pointer(), C.CString(format))
	}
}

func (ptr *QProgressBar) SetInvertedAppearance(invert bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::setInvertedAppearance")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_SetInvertedAppearance(ptr.Pointer(), C.int(qt.GoBoolToInt(invert)))
	}
}

func (ptr *QProgressBar) SetMaximum(maximum int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::setMaximum")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_SetMaximum(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QProgressBar) SetMinimum(minimum int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::setMinimum")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_SetMinimum(ptr.Pointer(), C.int(minimum))
	}
}

func (ptr *QProgressBar) SetOrientation(v core.Qt__Orientation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::setOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_SetOrientation(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QProgressBar) SetTextDirection(textDirection QProgressBar__Direction) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::setTextDirection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_SetTextDirection(ptr.Pointer(), C.int(textDirection))
	}
}

func (ptr *QProgressBar) SetTextVisible(visible bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::setTextVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_SetTextVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QProgressBar) SetValue(value int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::setValue")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_SetValue(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QProgressBar) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressBar_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressBar) TextDirection() QProgressBar__Direction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::textDirection")
		}
	}()

	if ptr.Pointer() != nil {
		return QProgressBar__Direction(C.QProgressBar_TextDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Value() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::value")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Value(ptr.Pointer()))
	}
	return 0
}

func NewQProgressBar(parent QWidget_ITF) *QProgressBar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::QProgressBar")
		}
	}()

	return NewQProgressBarFromPointer(C.QProgressBar_NewQProgressBar(PointerFromQWidget(parent)))
}

func (ptr *QProgressBar) Reset() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::reset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_Reset(ptr.Pointer())
	}
}

func (ptr *QProgressBar) SetRange(minimum int, maximum int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::setRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_SetRange(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QProgressBar) ConnectValueChanged(f func(value int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::valueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QProgressBar) DisconnectValueChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::valueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQProgressBarValueChanged
func callbackQProgressBarValueChanged(ptrName *C.char, value C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::valueChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func(int))(int(value))
}

func (ptr *QProgressBar) DestroyQProgressBar() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressBar::~QProgressBar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressBar_DestroyQProgressBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
