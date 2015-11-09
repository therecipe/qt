#include "qmediaservicedefaultdeviceinterface.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaService>
#include <QByteArray>
#include <QString>
#include <QMediaServiceDefaultDeviceInterface>
#include "_cgo_export.h"

class MyQMediaServiceDefaultDeviceInterface: public QMediaServiceDefaultDeviceInterface {
public:
};

void* QMediaServiceDefaultDeviceInterface_DefaultDevice(void* ptr, void* service){
	return new QByteArray(static_cast<QMediaServiceDefaultDeviceInterface*>(ptr)->defaultDevice(*static_cast<QByteArray*>(service)));
}

void QMediaServiceDefaultDeviceInterface_DestroyQMediaServiceDefaultDeviceInterface(void* ptr){
	static_cast<QMediaServiceDefaultDeviceInterface*>(ptr)->~QMediaServiceDefaultDeviceInterface();
}

