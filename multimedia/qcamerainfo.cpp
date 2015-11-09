#include "qcamerainfo.h"
#include <QModelIndex>
#include <QCamera>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QCameraInfo>
#include "_cgo_export.h"

class MyQCameraInfo: public QCameraInfo {
public:
};

void* QCameraInfo_NewQCameraInfo(void* name){
	return new QCameraInfo(*static_cast<QByteArray*>(name));
}

void* QCameraInfo_NewQCameraInfo2(void* camera){
	return new QCameraInfo(*static_cast<QCamera*>(camera));
}

void* QCameraInfo_NewQCameraInfo3(void* other){
	return new QCameraInfo(*static_cast<QCameraInfo*>(other));
}

char* QCameraInfo_Description(void* ptr){
	return static_cast<QCameraInfo*>(ptr)->description().toUtf8().data();
}

char* QCameraInfo_DeviceName(void* ptr){
	return static_cast<QCameraInfo*>(ptr)->deviceName().toUtf8().data();
}

int QCameraInfo_IsNull(void* ptr){
	return static_cast<QCameraInfo*>(ptr)->isNull();
}

int QCameraInfo_Orientation(void* ptr){
	return static_cast<QCameraInfo*>(ptr)->orientation();
}

int QCameraInfo_Position(void* ptr){
	return static_cast<QCameraInfo*>(ptr)->position();
}

void QCameraInfo_DestroyQCameraInfo(void* ptr){
	static_cast<QCameraInfo*>(ptr)->~QCameraInfo();
}

