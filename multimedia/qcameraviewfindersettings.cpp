#include "qcameraviewfindersettings.h"
#include <QVideoFrame>
#include <QCamera>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCameraViewfinderSettings>
#include "_cgo_export.h"

class MyQCameraViewfinderSettings: public QCameraViewfinderSettings {
public:
};

void* QCameraViewfinderSettings_NewQCameraViewfinderSettings(){
	return new QCameraViewfinderSettings();
}

void* QCameraViewfinderSettings_NewQCameraViewfinderSettings2(void* other){
	return new QCameraViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(other));
}

int QCameraViewfinderSettings_IsNull(void* ptr){
	return static_cast<QCameraViewfinderSettings*>(ptr)->isNull();
}

double QCameraViewfinderSettings_MaximumFrameRate(void* ptr){
	return static_cast<double>(static_cast<QCameraViewfinderSettings*>(ptr)->maximumFrameRate());
}

double QCameraViewfinderSettings_MinimumFrameRate(void* ptr){
	return static_cast<double>(static_cast<QCameraViewfinderSettings*>(ptr)->minimumFrameRate());
}

int QCameraViewfinderSettings_PixelFormat(void* ptr){
	return static_cast<QCameraViewfinderSettings*>(ptr)->pixelFormat();
}

void QCameraViewfinderSettings_SetMaximumFrameRate(void* ptr, double rate){
	static_cast<QCameraViewfinderSettings*>(ptr)->setMaximumFrameRate(static_cast<qreal>(rate));
}

void QCameraViewfinderSettings_SetMinimumFrameRate(void* ptr, double rate){
	static_cast<QCameraViewfinderSettings*>(ptr)->setMinimumFrameRate(static_cast<qreal>(rate));
}

void QCameraViewfinderSettings_SetPixelAspectRatio(void* ptr, void* ratio){
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelAspectRatio(*static_cast<QSize*>(ratio));
}

void QCameraViewfinderSettings_SetPixelAspectRatio2(void* ptr, int horizontal, int vertical){
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelAspectRatio(horizontal, vertical);
}

void QCameraViewfinderSettings_SetPixelFormat(void* ptr, int format){
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraViewfinderSettings_SetResolution(void* ptr, void* resolution){
	static_cast<QCameraViewfinderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QCameraViewfinderSettings_SetResolution2(void* ptr, int width, int height){
	static_cast<QCameraViewfinderSettings*>(ptr)->setResolution(width, height);
}

void QCameraViewfinderSettings_Swap(void* ptr, void* other){
	static_cast<QCameraViewfinderSettings*>(ptr)->swap(*static_cast<QCameraViewfinderSettings*>(other));
}

void QCameraViewfinderSettings_DestroyQCameraViewfinderSettings(void* ptr){
	static_cast<QCameraViewfinderSettings*>(ptr)->~QCameraViewfinderSettings();
}

