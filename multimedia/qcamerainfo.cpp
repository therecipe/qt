#include "qcamerainfo.h"
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QCamera>
#include <QString>
#include <QVariant>
#include <QCameraInfo>
#include "_cgo_export.h"

class MyQCameraInfo: public QCameraInfo {
public:
};

QtObjectPtr QCameraInfo_NewQCameraInfo(QtObjectPtr name){
	return new QCameraInfo(*static_cast<QByteArray*>(name));
}

QtObjectPtr QCameraInfo_NewQCameraInfo2(QtObjectPtr camera){
	return new QCameraInfo(*static_cast<QCamera*>(camera));
}

QtObjectPtr QCameraInfo_NewQCameraInfo3(QtObjectPtr other){
	return new QCameraInfo(*static_cast<QCameraInfo*>(other));
}

char* QCameraInfo_Description(QtObjectPtr ptr){
	return static_cast<QCameraInfo*>(ptr)->description().toUtf8().data();
}

char* QCameraInfo_DeviceName(QtObjectPtr ptr){
	return static_cast<QCameraInfo*>(ptr)->deviceName().toUtf8().data();
}

int QCameraInfo_IsNull(QtObjectPtr ptr){
	return static_cast<QCameraInfo*>(ptr)->isNull();
}

int QCameraInfo_Orientation(QtObjectPtr ptr){
	return static_cast<QCameraInfo*>(ptr)->orientation();
}

int QCameraInfo_Position(QtObjectPtr ptr){
	return static_cast<QCameraInfo*>(ptr)->position();
}

void QCameraInfo_DestroyQCameraInfo(QtObjectPtr ptr){
	static_cast<QCameraInfo*>(ptr)->~QCameraInfo();
}

