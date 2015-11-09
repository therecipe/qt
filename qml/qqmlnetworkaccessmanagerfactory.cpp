#include "qqmlnetworkaccessmanagerfactory.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlNetworkAccessManagerFactory>
#include "_cgo_export.h"

class MyQQmlNetworkAccessManagerFactory: public QQmlNetworkAccessManagerFactory {
public:
};

void* QQmlNetworkAccessManagerFactory_Create(void* ptr, void* parent){
	return static_cast<QQmlNetworkAccessManagerFactory*>(ptr)->create(static_cast<QObject*>(parent));
}

void QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(void* ptr){
	static_cast<QQmlNetworkAccessManagerFactory*>(ptr)->~QQmlNetworkAccessManagerFactory();
}

