#include "qqmlnetworkaccessmanagerfactory.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QQmlNetworkAccessManagerFactory>
#include "_cgo_export.h"

class MyQQmlNetworkAccessManagerFactory: public QQmlNetworkAccessManagerFactory {
public:
};

QtObjectPtr QQmlNetworkAccessManagerFactory_Create(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QQmlNetworkAccessManagerFactory*>(ptr)->create(static_cast<QObject*>(parent));
}

void QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(QtObjectPtr ptr){
	static_cast<QQmlNetworkAccessManagerFactory*>(ptr)->~QQmlNetworkAccessManagerFactory();
}

