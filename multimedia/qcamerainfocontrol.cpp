#include "qcamerainfocontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraInfo>
#include <QCameraInfoControl>
#include "_cgo_export.h"

class MyQCameraInfoControl: public QCameraInfoControl {
public:
};

int QCameraInfoControl_CameraOrientation(void* ptr, char* deviceName){
	return static_cast<QCameraInfoControl*>(ptr)->cameraOrientation(QString(deviceName));
}

int QCameraInfoControl_CameraPosition(void* ptr, char* deviceName){
	return static_cast<QCameraInfoControl*>(ptr)->cameraPosition(QString(deviceName));
}

void QCameraInfoControl_DestroyQCameraInfoControl(void* ptr){
	static_cast<QCameraInfoControl*>(ptr)->~QCameraInfoControl();
}

