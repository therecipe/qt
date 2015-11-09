#include "qmediaserviceproviderplugin.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaService>
#include <QMediaServiceProviderPlugin>
#include "_cgo_export.h"

class MyQMediaServiceProviderPlugin: public QMediaServiceProviderPlugin {
public:
};

void* QMediaServiceProviderPlugin_Create(void* ptr, char* key){
	return static_cast<QMediaServiceProviderPlugin*>(ptr)->create(QString(key));
}

void QMediaServiceProviderPlugin_Release(void* ptr, void* service){
	static_cast<QMediaServiceProviderPlugin*>(ptr)->release(static_cast<QMediaService*>(service));
}

