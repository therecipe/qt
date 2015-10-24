#include "qnetworkcachemetadata.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDate>
#include <QDateTime>
#include <QString>
#include <QNetworkCacheMetaData>
#include "_cgo_export.h"

class MyQNetworkCacheMetaData: public QNetworkCacheMetaData {
public:
};

QtObjectPtr QNetworkCacheMetaData_NewQNetworkCacheMetaData(){
	return new QNetworkCacheMetaData();
}

QtObjectPtr QNetworkCacheMetaData_NewQNetworkCacheMetaData2(QtObjectPtr other){
	return new QNetworkCacheMetaData(*static_cast<QNetworkCacheMetaData*>(other));
}

int QNetworkCacheMetaData_IsValid(QtObjectPtr ptr){
	return static_cast<QNetworkCacheMetaData*>(ptr)->isValid();
}

int QNetworkCacheMetaData_SaveToDisk(QtObjectPtr ptr){
	return static_cast<QNetworkCacheMetaData*>(ptr)->saveToDisk();
}

void QNetworkCacheMetaData_SetExpirationDate(QtObjectPtr ptr, QtObjectPtr dateTime){
	static_cast<QNetworkCacheMetaData*>(ptr)->setExpirationDate(*static_cast<QDateTime*>(dateTime));
}

void QNetworkCacheMetaData_SetLastModified(QtObjectPtr ptr, QtObjectPtr dateTime){
	static_cast<QNetworkCacheMetaData*>(ptr)->setLastModified(*static_cast<QDateTime*>(dateTime));
}

void QNetworkCacheMetaData_SetSaveToDisk(QtObjectPtr ptr, int allow){
	static_cast<QNetworkCacheMetaData*>(ptr)->setSaveToDisk(allow != 0);
}

void QNetworkCacheMetaData_SetUrl(QtObjectPtr ptr, char* url){
	static_cast<QNetworkCacheMetaData*>(ptr)->setUrl(QUrl(QString(url)));
}

void QNetworkCacheMetaData_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QNetworkCacheMetaData*>(ptr)->swap(*static_cast<QNetworkCacheMetaData*>(other));
}

char* QNetworkCacheMetaData_Url(QtObjectPtr ptr){
	return static_cast<QNetworkCacheMetaData*>(ptr)->url().toString().toUtf8().data();
}

void QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(QtObjectPtr ptr){
	static_cast<QNetworkCacheMetaData*>(ptr)->~QNetworkCacheMetaData();
}

