#include "qcameraimageprocessing.h"
#include <QCamera>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCameraImageProcessing>
#include "_cgo_export.h"

class MyQCameraImageProcessing: public QCameraImageProcessing {
public:
};

int QCameraImageProcessing_ColorFilter(void* ptr){
	return static_cast<QCameraImageProcessing*>(ptr)->colorFilter();
}

double QCameraImageProcessing_Contrast(void* ptr){
	return static_cast<double>(static_cast<QCameraImageProcessing*>(ptr)->contrast());
}

double QCameraImageProcessing_DenoisingLevel(void* ptr){
	return static_cast<double>(static_cast<QCameraImageProcessing*>(ptr)->denoisingLevel());
}

int QCameraImageProcessing_IsAvailable(void* ptr){
	return static_cast<QCameraImageProcessing*>(ptr)->isAvailable();
}

int QCameraImageProcessing_IsColorFilterSupported(void* ptr, int filter){
	return static_cast<QCameraImageProcessing*>(ptr)->isColorFilterSupported(static_cast<QCameraImageProcessing::ColorFilter>(filter));
}

int QCameraImageProcessing_IsWhiteBalanceModeSupported(void* ptr, int mode){
	return static_cast<QCameraImageProcessing*>(ptr)->isWhiteBalanceModeSupported(static_cast<QCameraImageProcessing::WhiteBalanceMode>(mode));
}

double QCameraImageProcessing_ManualWhiteBalance(void* ptr){
	return static_cast<double>(static_cast<QCameraImageProcessing*>(ptr)->manualWhiteBalance());
}

double QCameraImageProcessing_Saturation(void* ptr){
	return static_cast<double>(static_cast<QCameraImageProcessing*>(ptr)->saturation());
}

void QCameraImageProcessing_SetColorFilter(void* ptr, int filter){
	static_cast<QCameraImageProcessing*>(ptr)->setColorFilter(static_cast<QCameraImageProcessing::ColorFilter>(filter));
}

void QCameraImageProcessing_SetContrast(void* ptr, double value){
	static_cast<QCameraImageProcessing*>(ptr)->setContrast(static_cast<qreal>(value));
}

void QCameraImageProcessing_SetDenoisingLevel(void* ptr, double level){
	static_cast<QCameraImageProcessing*>(ptr)->setDenoisingLevel(static_cast<qreal>(level));
}

void QCameraImageProcessing_SetManualWhiteBalance(void* ptr, double colorTemperature){
	static_cast<QCameraImageProcessing*>(ptr)->setManualWhiteBalance(static_cast<qreal>(colorTemperature));
}

void QCameraImageProcessing_SetSaturation(void* ptr, double value){
	static_cast<QCameraImageProcessing*>(ptr)->setSaturation(static_cast<qreal>(value));
}

void QCameraImageProcessing_SetSharpeningLevel(void* ptr, double level){
	static_cast<QCameraImageProcessing*>(ptr)->setSharpeningLevel(static_cast<qreal>(level));
}

void QCameraImageProcessing_SetWhiteBalanceMode(void* ptr, int mode){
	static_cast<QCameraImageProcessing*>(ptr)->setWhiteBalanceMode(static_cast<QCameraImageProcessing::WhiteBalanceMode>(mode));
}

double QCameraImageProcessing_SharpeningLevel(void* ptr){
	return static_cast<double>(static_cast<QCameraImageProcessing*>(ptr)->sharpeningLevel());
}

int QCameraImageProcessing_WhiteBalanceMode(void* ptr){
	return static_cast<QCameraImageProcessing*>(ptr)->whiteBalanceMode();
}

