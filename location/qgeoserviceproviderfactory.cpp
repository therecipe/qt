#include "qgeoserviceproviderfactory.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoServiceProvider>
#include <QString>
#include <QGeoServiceProviderFactory>
#include "_cgo_export.h"

class MyQGeoServiceProviderFactory: public QGeoServiceProviderFactory {
public:
};

void QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(void* ptr){
	static_cast<QGeoServiceProviderFactory*>(ptr)->~QGeoServiceProviderFactory();
}

