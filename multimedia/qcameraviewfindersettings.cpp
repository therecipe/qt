#include "qcameraviewfindersettings.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QVideoFrame>
#include <QCamera>
#include <QCameraViewfinderSettings>
#include "_cgo_export.h"

class MyQCameraViewfinderSettings: public QCameraViewfinderSettings {
public:
};

QtObjectPtr QCameraViewfinderSettings_NewQCameraViewfinderSettings(){
	return new QCameraViewfinderSettings();
}

QtObjectPtr QCameraViewfinderSettings_NewQCameraViewfinderSettings2(QtObjectPtr other){
	return new QCameraViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(other));
}

int QCameraViewfinderSettings_IsNull(QtObjectPtr ptr){
	return static_cast<QCameraViewfinderSettings*>(ptr)->isNull();
}

int QCameraViewfinderSettings_PixelFormat(QtObjectPtr ptr){
	return static_cast<QCameraViewfinderSettings*>(ptr)->pixelFormat();
}

void QCameraViewfinderSettings_SetPixelAspectRatio(QtObjectPtr ptr, QtObjectPtr ratio){
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelAspectRatio(*static_cast<QSize*>(ratio));
}

void QCameraViewfinderSettings_SetPixelAspectRatio2(QtObjectPtr ptr, int horizontal, int vertical){
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelAspectRatio(horizontal, vertical);
}

void QCameraViewfinderSettings_SetPixelFormat(QtObjectPtr ptr, int format){
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraViewfinderSettings_SetResolution(QtObjectPtr ptr, QtObjectPtr resolution){
	static_cast<QCameraViewfinderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QCameraViewfinderSettings_SetResolution2(QtObjectPtr ptr, int width, int height){
	static_cast<QCameraViewfinderSettings*>(ptr)->setResolution(width, height);
}

void QCameraViewfinderSettings_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QCameraViewfinderSettings*>(ptr)->swap(*static_cast<QCameraViewfinderSettings*>(other));
}

void QCameraViewfinderSettings_DestroyQCameraViewfinderSettings(QtObjectPtr ptr){
	static_cast<QCameraViewfinderSettings*>(ptr)->~QCameraViewfinderSettings();
}

