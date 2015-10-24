#include "qcameraimageprocessing.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QString>
#include <QCameraImageProcessing>
#include "_cgo_export.h"

class MyQCameraImageProcessing: public QCameraImageProcessing {
public:
};

int QCameraImageProcessing_ColorFilter(QtObjectPtr ptr){
	return static_cast<QCameraImageProcessing*>(ptr)->colorFilter();
}

int QCameraImageProcessing_IsAvailable(QtObjectPtr ptr){
	return static_cast<QCameraImageProcessing*>(ptr)->isAvailable();
}

int QCameraImageProcessing_IsColorFilterSupported(QtObjectPtr ptr, int filter){
	return static_cast<QCameraImageProcessing*>(ptr)->isColorFilterSupported(static_cast<QCameraImageProcessing::ColorFilter>(filter));
}

int QCameraImageProcessing_IsWhiteBalanceModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCameraImageProcessing*>(ptr)->isWhiteBalanceModeSupported(static_cast<QCameraImageProcessing::WhiteBalanceMode>(mode));
}

void QCameraImageProcessing_SetColorFilter(QtObjectPtr ptr, int filter){
	static_cast<QCameraImageProcessing*>(ptr)->setColorFilter(static_cast<QCameraImageProcessing::ColorFilter>(filter));
}

void QCameraImageProcessing_SetWhiteBalanceMode(QtObjectPtr ptr, int mode){
	static_cast<QCameraImageProcessing*>(ptr)->setWhiteBalanceMode(static_cast<QCameraImageProcessing::WhiteBalanceMode>(mode));
}

int QCameraImageProcessing_WhiteBalanceMode(QtObjectPtr ptr){
	return static_cast<QCameraImageProcessing*>(ptr)->whiteBalanceMode();
}

