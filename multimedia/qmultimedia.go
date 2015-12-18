package multimedia

//#include "multimedia.h"
import "C"
import ()

//QMultimedia::AvailabilityStatus
type QMultimedia__AvailabilityStatus int64

const (
	QMultimedia__Available      = QMultimedia__AvailabilityStatus(0)
	QMultimedia__ServiceMissing = QMultimedia__AvailabilityStatus(1)
	QMultimedia__Busy           = QMultimedia__AvailabilityStatus(2)
	QMultimedia__ResourceError  = QMultimedia__AvailabilityStatus(3)
)

//QMultimedia::EncodingMode
type QMultimedia__EncodingMode int64

const (
	QMultimedia__ConstantQualityEncoding = QMultimedia__EncodingMode(0)
	QMultimedia__ConstantBitRateEncoding = QMultimedia__EncodingMode(1)
	QMultimedia__AverageBitRateEncoding  = QMultimedia__EncodingMode(2)
	QMultimedia__TwoPassEncoding         = QMultimedia__EncodingMode(3)
)

//QMultimedia::EncodingQuality
type QMultimedia__EncodingQuality int64

const (
	QMultimedia__VeryLowQuality  = QMultimedia__EncodingQuality(0)
	QMultimedia__LowQuality      = QMultimedia__EncodingQuality(1)
	QMultimedia__NormalQuality   = QMultimedia__EncodingQuality(2)
	QMultimedia__HighQuality     = QMultimedia__EncodingQuality(3)
	QMultimedia__VeryHighQuality = QMultimedia__EncodingQuality(4)
)

//QMultimedia::SupportEstimate
type QMultimedia__SupportEstimate int64

const (
	QMultimedia__NotSupported      = QMultimedia__SupportEstimate(0)
	QMultimedia__MaybeSupported    = QMultimedia__SupportEstimate(1)
	QMultimedia__ProbablySupported = QMultimedia__SupportEstimate(2)
	QMultimedia__PreferredService  = QMultimedia__SupportEstimate(3)
)
