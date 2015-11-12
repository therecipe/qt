#include "qmediaservicesupporteddevicesinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QMediaService>
#include <QMediaServiceSupportedDevicesInterface>
#include "_cgo_export.h"

class MyQMediaServiceSupportedDevicesInterface: public QMediaServiceSupportedDevicesInterface {
public:
};

char* QMediaServiceSupportedDevicesInterface_DeviceDescription(void* ptr, void* service, void* device){
	return static_cast<QMediaServiceSupportedDevicesInterface*>(ptr)->deviceDescription(*static_cast<QByteArray*>(service), *static_cast<QByteArray*>(device)).toUtf8().data();
}

void QMediaServiceSupportedDevicesInterface_DestroyQMediaServiceSupportedDevicesInterface(void* ptr){
	static_cast<QMediaServiceSupportedDevicesInterface*>(ptr)->~QMediaServiceSupportedDevicesInterface();
}

