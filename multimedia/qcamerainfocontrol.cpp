#include "qcamerainfocontrol.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraInfo>
#include <QString>
#include <QCameraInfoControl>
#include "_cgo_export.h"

class MyQCameraInfoControl: public QCameraInfoControl {
public:
};

int QCameraInfoControl_CameraOrientation(QtObjectPtr ptr, char* deviceName){
	return static_cast<QCameraInfoControl*>(ptr)->cameraOrientation(QString(deviceName));
}

int QCameraInfoControl_CameraPosition(QtObjectPtr ptr, char* deviceName){
	return static_cast<QCameraInfoControl*>(ptr)->cameraPosition(QString(deviceName));
}

void QCameraInfoControl_DestroyQCameraInfoControl(QtObjectPtr ptr){
	static_cast<QCameraInfoControl*>(ptr)->~QCameraInfoControl();
}

