#include "qmediaservicecamerainfointerface.h"
#include <QByteArray>
#include <QMediaService>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaServiceCameraInfoInterface>
#include "_cgo_export.h"

class MyQMediaServiceCameraInfoInterface: public QMediaServiceCameraInfoInterface {
public:
};

int QMediaServiceCameraInfoInterface_CameraOrientation(void* ptr, void* device){
	return static_cast<QMediaServiceCameraInfoInterface*>(ptr)->cameraOrientation(*static_cast<QByteArray*>(device));
}

int QMediaServiceCameraInfoInterface_CameraPosition(void* ptr, void* device){
	return static_cast<QMediaServiceCameraInfoInterface*>(ptr)->cameraPosition(*static_cast<QByteArray*>(device));
}

void QMediaServiceCameraInfoInterface_DestroyQMediaServiceCameraInfoInterface(void* ptr){
	static_cast<QMediaServiceCameraInfoInterface*>(ptr)->~QMediaServiceCameraInfoInterface();
}

