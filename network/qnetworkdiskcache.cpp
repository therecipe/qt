#include "qnetworkdiskcache.h"
#include <QUrl>
#include <QModelIndex>
#include <QNetworkCacheMetaData>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QNetworkDiskCache>
#include "_cgo_export.h"

class MyQNetworkDiskCache: public QNetworkDiskCache {
public:
};

QtObjectPtr QNetworkDiskCache_NewQNetworkDiskCache(QtObjectPtr parent){
	return new QNetworkDiskCache(static_cast<QObject*>(parent));
}

char* QNetworkDiskCache_CacheDirectory(QtObjectPtr ptr){
	return static_cast<QNetworkDiskCache*>(ptr)->cacheDirectory().toUtf8().data();
}

void QNetworkDiskCache_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkDiskCache*>(ptr), "clear");
}

QtObjectPtr QNetworkDiskCache_Data(QtObjectPtr ptr, char* url){
	return static_cast<QNetworkDiskCache*>(ptr)->data(QUrl(QString(url)));
}

QtObjectPtr QNetworkDiskCache_Prepare(QtObjectPtr ptr, QtObjectPtr metaData){
	return static_cast<QNetworkDiskCache*>(ptr)->prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QNetworkDiskCache_SetCacheDirectory(QtObjectPtr ptr, char* cacheDir){
	static_cast<QNetworkDiskCache*>(ptr)->setCacheDirectory(QString(cacheDir));
}

void QNetworkDiskCache_UpdateMetaData(QtObjectPtr ptr, QtObjectPtr metaData){
	static_cast<QNetworkDiskCache*>(ptr)->updateMetaData(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QNetworkDiskCache_DestroyQNetworkDiskCache(QtObjectPtr ptr){
	static_cast<QNetworkDiskCache*>(ptr)->~QNetworkDiskCache();
}

