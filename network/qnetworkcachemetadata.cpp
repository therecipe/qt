#include "qnetworkcachemetadata.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDate>
#include <QDateTime>
#include <QNetworkCacheMetaData>
#include "_cgo_export.h"

class MyQNetworkCacheMetaData: public QNetworkCacheMetaData {
public:
};

void* QNetworkCacheMetaData_NewQNetworkCacheMetaData(){
	return new QNetworkCacheMetaData();
}

void* QNetworkCacheMetaData_NewQNetworkCacheMetaData2(void* other){
	return new QNetworkCacheMetaData(*static_cast<QNetworkCacheMetaData*>(other));
}

void* QNetworkCacheMetaData_ExpirationDate(void* ptr){
	return new QDateTime(static_cast<QNetworkCacheMetaData*>(ptr)->expirationDate());
}

int QNetworkCacheMetaData_IsValid(void* ptr){
	return static_cast<QNetworkCacheMetaData*>(ptr)->isValid();
}

void* QNetworkCacheMetaData_LastModified(void* ptr){
	return new QDateTime(static_cast<QNetworkCacheMetaData*>(ptr)->lastModified());
}

int QNetworkCacheMetaData_SaveToDisk(void* ptr){
	return static_cast<QNetworkCacheMetaData*>(ptr)->saveToDisk();
}

void QNetworkCacheMetaData_SetExpirationDate(void* ptr, void* dateTime){
	static_cast<QNetworkCacheMetaData*>(ptr)->setExpirationDate(*static_cast<QDateTime*>(dateTime));
}

void QNetworkCacheMetaData_SetLastModified(void* ptr, void* dateTime){
	static_cast<QNetworkCacheMetaData*>(ptr)->setLastModified(*static_cast<QDateTime*>(dateTime));
}

void QNetworkCacheMetaData_SetSaveToDisk(void* ptr, int allow){
	static_cast<QNetworkCacheMetaData*>(ptr)->setSaveToDisk(allow != 0);
}

void QNetworkCacheMetaData_SetUrl(void* ptr, void* url){
	static_cast<QNetworkCacheMetaData*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QNetworkCacheMetaData_Swap(void* ptr, void* other){
	static_cast<QNetworkCacheMetaData*>(ptr)->swap(*static_cast<QNetworkCacheMetaData*>(other));
}

void QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(void* ptr){
	static_cast<QNetworkCacheMetaData*>(ptr)->~QNetworkCacheMetaData();
}

