#include "qabstractnetworkcache.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkCacheMetaData>
#include <QMetaObject>
#include <QAbstractNetworkCache>
#include "_cgo_export.h"

class MyQAbstractNetworkCache: public QAbstractNetworkCache {
public:
};

void QAbstractNetworkCache_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractNetworkCache*>(ptr), "clear");
}

void* QAbstractNetworkCache_Data(void* ptr, void* url){
	return static_cast<QAbstractNetworkCache*>(ptr)->data(*static_cast<QUrl*>(url));
}

void* QAbstractNetworkCache_Prepare(void* ptr, void* metaData){
	return static_cast<QAbstractNetworkCache*>(ptr)->prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QAbstractNetworkCache_UpdateMetaData(void* ptr, void* metaData){
	static_cast<QAbstractNetworkCache*>(ptr)->updateMetaData(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QAbstractNetworkCache_DestroyQAbstractNetworkCache(void* ptr){
	static_cast<QAbstractNetworkCache*>(ptr)->~QAbstractNetworkCache();
}

