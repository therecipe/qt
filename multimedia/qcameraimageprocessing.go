package multimedia

//#include "qcameraimageprocessing.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraImageProcessing struct {
	core.QObject
}

type QCameraImageProcessingITF interface {
	core.QObjectITF
	QCameraImageProcessingPTR() *QCameraImageProcessing
}

func PointerFromQCameraImageProcessing(ptr QCameraImageProcessingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageProcessingPTR().Pointer()
	}
	return nil
}

func QCameraImageProcessingFromPointer(ptr unsafe.Pointer) *QCameraImageProcessing {
	var n = new(QCameraImageProcessing)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraImageProcessing_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraImageProcessing) QCameraImageProcessingPTR() *QCameraImageProcessing {
	return ptr
}

//QCameraImageProcessing::ColorFilter
type QCameraImageProcessing__ColorFilter int

var (
	QCameraImageProcessing__ColorFilterNone       = QCameraImageProcessing__ColorFilter(0)
	QCameraImageProcessing__ColorFilterGrayscale  = QCameraImageProcessing__ColorFilter(1)
	QCameraImageProcessing__ColorFilterNegative   = QCameraImageProcessing__ColorFilter(2)
	QCameraImageProcessing__ColorFilterSolarize   = QCameraImageProcessing__ColorFilter(3)
	QCameraImageProcessing__ColorFilterSepia      = QCameraImageProcessing__ColorFilter(4)
	QCameraImageProcessing__ColorFilterPosterize  = QCameraImageProcessing__ColorFilter(5)
	QCameraImageProcessing__ColorFilterWhiteboard = QCameraImageProcessing__ColorFilter(6)
	QCameraImageProcessing__ColorFilterBlackboard = QCameraImageProcessing__ColorFilter(7)
	QCameraImageProcessing__ColorFilterAqua       = QCameraImageProcessing__ColorFilter(8)
	QCameraImageProcessing__ColorFilterVendor     = QCameraImageProcessing__ColorFilter(1000)
)

//QCameraImageProcessing::WhiteBalanceMode
type QCameraImageProcessing__WhiteBalanceMode int

var (
	QCameraImageProcessing__WhiteBalanceAuto        = QCameraImageProcessing__WhiteBalanceMode(0)
	QCameraImageProcessing__WhiteBalanceManual      = QCameraImageProcessing__WhiteBalanceMode(1)
	QCameraImageProcessing__WhiteBalanceSunlight    = QCameraImageProcessing__WhiteBalanceMode(2)
	QCameraImageProcessing__WhiteBalanceCloudy      = QCameraImageProcessing__WhiteBalanceMode(3)
	QCameraImageProcessing__WhiteBalanceShade       = QCameraImageProcessing__WhiteBalanceMode(4)
	QCameraImageProcessing__WhiteBalanceTungsten    = QCameraImageProcessing__WhiteBalanceMode(5)
	QCameraImageProcessing__WhiteBalanceFluorescent = QCameraImageProcessing__WhiteBalanceMode(6)
	QCameraImageProcessing__WhiteBalanceFlash       = QCameraImageProcessing__WhiteBalanceMode(7)
	QCameraImageProcessing__WhiteBalanceSunset      = QCameraImageProcessing__WhiteBalanceMode(8)
	QCameraImageProcessing__WhiteBalanceVendor      = QCameraImageProcessing__WhiteBalanceMode(1000)
)

func (ptr *QCameraImageProcessing) ColorFilter() QCameraImageProcessing__ColorFilter {
	if ptr.Pointer() != nil {
		return QCameraImageProcessing__ColorFilter(C.QCameraImageProcessing_ColorFilter(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraImageProcessing) IsAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageProcessing_IsAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCameraImageProcessing) IsColorFilterSupported(filter QCameraImageProcessing__ColorFilter) bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageProcessing_IsColorFilterSupported(C.QtObjectPtr(ptr.Pointer()), C.int(filter)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessing) IsWhiteBalanceModeSupported(mode QCameraImageProcessing__WhiteBalanceMode) bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageProcessing_IsWhiteBalanceModeSupported(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessing) SetColorFilter(filter QCameraImageProcessing__ColorFilter) {
	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetColorFilter(C.QtObjectPtr(ptr.Pointer()), C.int(filter))
	}
}

func (ptr *QCameraImageProcessing) SetWhiteBalanceMode(mode QCameraImageProcessing__WhiteBalanceMode) {
	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetWhiteBalanceMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QCameraImageProcessing) WhiteBalanceMode() QCameraImageProcessing__WhiteBalanceMode {
	if ptr.Pointer() != nil {
		return QCameraImageProcessing__WhiteBalanceMode(C.QCameraImageProcessing_WhiteBalanceMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
