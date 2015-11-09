#include "qnetworkproxyfactory.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkProxy>
#include <QNetworkProxyFactory>
#include "_cgo_export.h"

class MyQNetworkProxyFactory: public QNetworkProxyFactory {
public:
};

void QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(void* factory){
	QNetworkProxyFactory::setApplicationProxyFactory(static_cast<QNetworkProxyFactory*>(factory));
}

void QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(int enable){
	QNetworkProxyFactory::setUseSystemConfiguration(enable != 0);
}

void QNetworkProxyFactory_DestroyQNetworkProxyFactory(void* ptr){
	static_cast<QNetworkProxyFactory*>(ptr)->~QNetworkProxyFactory();
}

