#include "qmediaservicesupporteddevicesinterface.h"
#include <QUrl>
#include <QModelIndex>
#include <QMediaService>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QMediaServiceSupportedDevicesInterface>
#include "_cgo_export.h"

class MyQMediaServiceSupportedDevicesInterface: public QMediaServiceSupportedDevicesInterface {
public:
};

char* QMediaServiceSupportedDevicesInterface_DeviceDescription(QtObjectPtr ptr, QtObjectPtr service, QtObjectPtr device){
	return static_cast<QMediaServiceSupportedDevicesInterface*>(ptr)->deviceDescription(*static_cast<QByteArray*>(service), *static_cast<QByteArray*>(device)).toUtf8().data();
}

void QMediaServiceSupportedDevicesInterface_DestroyQMediaServiceSupportedDevicesInterface(QtObjectPtr ptr){
	static_cast<QMediaServiceSupportedDevicesInterface*>(ptr)->~QMediaServiceSupportedDevicesInterface();
}

