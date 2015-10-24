package gui

//#include "qcolor.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QColor struct {
	ptr unsafe.Pointer
}

type QColorITF interface {
	QColorPTR() *QColor
}

func (p *QColor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QColor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQColor(ptr QColorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColorPTR().Pointer()
	}
	return nil
}

func QColorFromPointer(ptr unsafe.Pointer) *QColor {
	var n = new(QColor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QColor) QColorPTR() *QColor {
	return ptr
}

//QColor::NameFormat
type QColor__NameFormat int

var (
	QColor__HexRgb  = QColor__NameFormat(0)
	QColor__HexArgb = QColor__NameFormat(1)
)

//QColor::Spec
type QColor__Spec int

var (
	QColor__Invalid = QColor__Spec(0)
	QColor__Rgb     = QColor__Spec(1)
	QColor__Hsv     = QColor__Spec(2)
	QColor__Cmyk    = QColor__Spec(3)
	QColor__Hsl     = QColor__Spec(4)
)

func NewQColor() *QColor {
	return QColorFromPointer(unsafe.Pointer(C.QColor_NewQColor()))
}

func NewQColor8(color core.Qt__GlobalColor) *QColor {
	return QColorFromPointer(unsafe.Pointer(C.QColor_NewQColor8(C.int(color))))
}

func NewQColor6(color QColorITF) *QColor {
	return QColorFromPointer(unsafe.Pointer(C.QColor_NewQColor6(C.QtObjectPtr(PointerFromQColor(color)))))
}

func NewQColor4(name string) *QColor {
	return QColorFromPointer(unsafe.Pointer(C.QColor_NewQColor4(C.CString(name))))
}

func NewQColor5(name string) *QColor {
	return QColorFromPointer(unsafe.Pointer(C.QColor_NewQColor5(C.CString(name))))
}

func NewQColor2(r int, g int, b int, a int) *QColor {
	return QColorFromPointer(unsafe.Pointer(C.QColor_NewQColor2(C.int(r), C.int(g), C.int(b), C.int(a))))
}

func (ptr *QColor) Alpha() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Alpha(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) Black() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Black(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) Blue() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Blue(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QColor_ColorNames() []string {
	return strings.Split(C.GoString(C.QColor_QColor_ColorNames()), "|")
}

func (ptr *QColor) Cyan() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Cyan(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) GetCmyk(c int, m int, y int, k int, a int) {
	if ptr.Pointer() != nil {
		C.QColor_GetCmyk(C.QtObjectPtr(ptr.Pointer()), C.int(c), C.int(m), C.int(y), C.int(k), C.int(a))
	}
}

func (ptr *QColor) GetHsl(h int, s int, l int, a int) {
	if ptr.Pointer() != nil {
		C.QColor_GetHsl(C.QtObjectPtr(ptr.Pointer()), C.int(h), C.int(s), C.int(l), C.int(a))
	}
}

func (ptr *QColor) GetHsv(h int, s int, v int, a int) {
	if ptr.Pointer() != nil {
		C.QColor_GetHsv(C.QtObjectPtr(ptr.Pointer()), C.int(h), C.int(s), C.int(v), C.int(a))
	}
}

func (ptr *QColor) GetRgb(r int, g int, b int, a int) {
	if ptr.Pointer() != nil {
		C.QColor_GetRgb(C.QtObjectPtr(ptr.Pointer()), C.int(r), C.int(g), C.int(b), C.int(a))
	}
}

func (ptr *QColor) Green() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Green(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) HslHue() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_HslHue(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) HslSaturation() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_HslSaturation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) HsvHue() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_HsvHue(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) HsvSaturation() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_HsvSaturation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) Hue() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Hue(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QColor_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QColor_IsValidColor(name string) bool {
	return C.QColor_QColor_IsValidColor(C.CString(name)) != 0
}

func (ptr *QColor) Lightness() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Lightness(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) Magenta() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Magenta(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QColor_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QColor) Name2(format QColor__NameFormat) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QColor_Name2(C.QtObjectPtr(ptr.Pointer()), C.int(format)))
	}
	return ""
}

func (ptr *QColor) Red() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Red(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) Saturation() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Saturation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) SetAlpha(alpha int) {
	if ptr.Pointer() != nil {
		C.QColor_SetAlpha(C.QtObjectPtr(ptr.Pointer()), C.int(alpha))
	}
}

func (ptr *QColor) SetBlue(blue int) {
	if ptr.Pointer() != nil {
		C.QColor_SetBlue(C.QtObjectPtr(ptr.Pointer()), C.int(blue))
	}
}

func (ptr *QColor) SetCmyk(c int, m int, y int, k int, a int) {
	if ptr.Pointer() != nil {
		C.QColor_SetCmyk(C.QtObjectPtr(ptr.Pointer()), C.int(c), C.int(m), C.int(y), C.int(k), C.int(a))
	}
}

func (ptr *QColor) SetGreen(green int) {
	if ptr.Pointer() != nil {
		C.QColor_SetGreen(C.QtObjectPtr(ptr.Pointer()), C.int(green))
	}
}

func (ptr *QColor) SetHsl(h int, s int, l int, a int) {
	if ptr.Pointer() != nil {
		C.QColor_SetHsl(C.QtObjectPtr(ptr.Pointer()), C.int(h), C.int(s), C.int(l), C.int(a))
	}
}

func (ptr *QColor) SetHsv(h int, s int, v int, a int) {
	if ptr.Pointer() != nil {
		C.QColor_SetHsv(C.QtObjectPtr(ptr.Pointer()), C.int(h), C.int(s), C.int(v), C.int(a))
	}
}

func (ptr *QColor) SetNamedColor(name string) {
	if ptr.Pointer() != nil {
		C.QColor_SetNamedColor(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QColor) SetRed(red int) {
	if ptr.Pointer() != nil {
		C.QColor_SetRed(C.QtObjectPtr(ptr.Pointer()), C.int(red))
	}
}

func (ptr *QColor) SetRgb(r int, g int, b int, a int) {
	if ptr.Pointer() != nil {
		C.QColor_SetRgb(C.QtObjectPtr(ptr.Pointer()), C.int(r), C.int(g), C.int(b), C.int(a))
	}
}

func (ptr *QColor) Spec() QColor__Spec {
	if ptr.Pointer() != nil {
		return QColor__Spec(C.QColor_Spec(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) Value() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Value(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) Yellow() int {
	if ptr.Pointer() != nil {
		return int(C.QColor_Yellow(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
