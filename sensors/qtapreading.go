package sensors

//#include "qtapreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTapReading struct {
	QSensorReading
}

type QTapReadingITF interface {
	QSensorReadingITF
	QTapReadingPTR() *QTapReading
}

func PointerFromQTapReading(ptr QTapReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapReadingPTR().Pointer()
	}
	return nil
}

func QTapReadingFromPointer(ptr unsafe.Pointer) *QTapReading {
	var n = new(QTapReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTapReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTapReading) QTapReadingPTR() *QTapReading {
	return ptr
}

//QTapReading::TapDirection
type QTapReading__TapDirection int

var (
	QTapReading__Undefined = QTapReading__TapDirection(0)
	QTapReading__X         = QTapReading__TapDirection(0x0001)
	QTapReading__Y         = QTapReading__TapDirection(0x0002)
	QTapReading__Z         = QTapReading__TapDirection(0x0004)
	QTapReading__X_Pos     = QTapReading__TapDirection(0x0011)
	QTapReading__Y_Pos     = QTapReading__TapDirection(0x0022)
	QTapReading__Z_Pos     = QTapReading__TapDirection(0x0044)
	QTapReading__X_Neg     = QTapReading__TapDirection(0x0101)
	QTapReading__Y_Neg     = QTapReading__TapDirection(0x0202)
	QTapReading__Z_Neg     = QTapReading__TapDirection(0x0404)
	QTapReading__X_Both    = QTapReading__TapDirection(0x0111)
	QTapReading__Y_Both    = QTapReading__TapDirection(0x0222)
	QTapReading__Z_Both    = QTapReading__TapDirection(0x0444)
)

func (ptr *QTapReading) IsDoubleTap() bool {
	if ptr.Pointer() != nil {
		return C.QTapReading_IsDoubleTap(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTapReading) TapDirection() QTapReading__TapDirection {
	if ptr.Pointer() != nil {
		return QTapReading__TapDirection(C.QTapReading_TapDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTapReading) SetDoubleTap(doubleTap bool) {
	if ptr.Pointer() != nil {
		C.QTapReading_SetDoubleTap(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(doubleTap)))
	}
}

func (ptr *QTapReading) SetTapDirection(tapDirection QTapReading__TapDirection) {
	if ptr.Pointer() != nil {
		C.QTapReading_SetTapDirection(C.QtObjectPtr(ptr.Pointer()), C.int(tapDirection))
	}
}
