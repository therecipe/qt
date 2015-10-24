#include "qabstractnetworkcache.h"
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkCacheMetaData>
#include <QAbstractNetworkCache>
#include "_cgo_export.h"

class MyQAbstractNetworkCache: public QAbstractNetworkCache {
public:
};

void QAbstractNetworkCache_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractNetworkCache*>(ptr), "clear");
}

QtObjectPtr QAbstractNetworkCache_Data(QtObjectPtr ptr, char* url){
	return static_cast<QAbstractNetworkCache*>(ptr)->data(QUrl(QString(url)));
}

QtObjectPtr QAbstractNetworkCache_Prepare(QtObjectPtr ptr, QtObjectPtr metaData){
	return static_cast<QAbstractNetworkCache*>(ptr)->prepare(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QAbstractNetworkCache_UpdateMetaData(QtObjectPtr ptr, QtObjectPtr metaData){
	static_cast<QAbstractNetworkCache*>(ptr)->updateMetaData(*static_cast<QNetworkCacheMetaData*>(metaData));
}

void QAbstractNetworkCache_DestroyQAbstractNetworkCache(QtObjectPtr ptr){
	static_cast<QAbstractNetworkCache*>(ptr)->~QAbstractNetworkCache();
}

