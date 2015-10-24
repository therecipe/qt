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

QtObjectPtr QMediaServiceProviderPlugin_Create(QtObjectPtr ptr, char* key){
	return static_cast<QMediaServiceProviderPlugin*>(ptr)->create(QString(key));
}

void QMediaServiceProviderPlugin_Release(QtObjectPtr ptr, QtObjectPtr service){
	static_cast<QMediaServiceProviderPlugin*>(ptr)->release(static_cast<QMediaService*>(service));
}

