#include "qmediaservicecamerainfointerface.h"
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QMediaService>
#include <QString>
#include <QVariant>
#include <QMediaServiceCameraInfoInterface>
#include "_cgo_export.h"

class MyQMediaServiceCameraInfoInterface: public QMediaServiceCameraInfoInterface {
public:
};

int QMediaServiceCameraInfoInterface_CameraOrientation(QtObjectPtr ptr, QtObjectPtr device){
	return static_cast<QMediaServiceCameraInfoInterface*>(ptr)->cameraOrientation(*static_cast<QByteArray*>(device));
}

int QMediaServiceCameraInfoInterface_CameraPosition(QtObjectPtr ptr, QtObjectPtr device){
	return static_cast<QMediaServiceCameraInfoInterface*>(ptr)->cameraPosition(*static_cast<QByteArray*>(device));
}

void QMediaServiceCameraInfoInterface_DestroyQMediaServiceCameraInfoInterface(QtObjectPtr ptr){
	static_cast<QMediaServiceCameraInfoInterface*>(ptr)->~QMediaServiceCameraInfoInterface();
}

