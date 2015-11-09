#include "qnetworkdiskcache.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkCacheMetaData>
#include <QMetaObject>
#include <QObject>
#include <QNetworkDiskCache>
#include "_cgo_export.h"

class MyQNetworkDiskCache: public QNetworkDiskCache {
public:
};

void* QNetworkDiskCache_NewQNetworkDiskCache(void* parent){
	return new QNetworkDiskCache(static_cast<QObject*>(parent));
}

char* QNetworkDiskCache_CacheDirectory(void* ptr){
	return static_cast<QNetworkDiskCache*>(ptr)->cacheDirectory().toUtf8().data();
}

void QNetworkDiskCache_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkDiskCache*>(ptr), "clear");
}

void* QNetworkDiskCache_Data(void* ptr, void* url){
	return static_cast<QNetworkDiskCache*>(ptr)->data(*static_cast<QUrl*>(url));
}

void* QNetworkDiskCache_Prepare(void* ptr, void* metaData){
	return static_cast<QNetworkDiskCache*>(ptr)->prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QNetworkDiskCache_SetCacheDirectory(void* ptr, char* cacheDir){
	static_cast<QNetworkDiskCache*>(ptr)->setCacheDirectory(QString(cacheDir));
}

void QNetworkDiskCache_UpdateMetaData(void* ptr, void* metaData){
	static_cast<QNetworkDiskCache*>(ptr)->updateMetaData(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QNetworkDiskCache_DestroyQNetworkDiskCache(void* ptr){
	static_cast<QNetworkDiskCache*>(ptr)->~QNetworkDiskCache();
}

