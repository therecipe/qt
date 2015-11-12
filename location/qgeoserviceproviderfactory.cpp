#include "qgeoserviceproviderfactory.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoServiceProvider>
#include <QGeoServiceProviderFactory>
#include "_cgo_export.h"

class MyQGeoServiceProviderFactory: public QGeoServiceProviderFactory {
public:
};

void QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(void* ptr){
	static_cast<QGeoServiceProviderFactory*>(ptr)->~QGeoServiceProviderFactory();
}

