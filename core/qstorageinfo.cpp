#include "qstorageinfo.h"
#include <QDir>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStorageInfo>
#include "_cgo_export.h"

class MyQStorageInfo: public QStorageInfo {
public:
};

QtObjectPtr QStorageInfo_NewQStorageInfo(){
	return new QStorageInfo();
}

QtObjectPtr QStorageInfo_NewQStorageInfo3(QtObjectPtr dir){
	return new QStorageInfo(*static_cast<QDir*>(dir));
}

QtObjectPtr QStorageInfo_NewQStorageInfo4(QtObjectPtr other){
	return new QStorageInfo(*static_cast<QStorageInfo*>(other));
}

QtObjectPtr QStorageInfo_NewQStorageInfo2(char* path){
	return new QStorageInfo(QString(path));
}

char* QStorageInfo_DisplayName(QtObjectPtr ptr){
	return static_cast<QStorageInfo*>(ptr)->displayName().toUtf8().data();
}

int QStorageInfo_IsReadOnly(QtObjectPtr ptr){
	return static_cast<QStorageInfo*>(ptr)->isReadOnly();
}

int QStorageInfo_IsReady(QtObjectPtr ptr){
	return static_cast<QStorageInfo*>(ptr)->isReady();
}

int QStorageInfo_IsRoot(QtObjectPtr ptr){
	return static_cast<QStorageInfo*>(ptr)->isRoot();
}

int QStorageInfo_IsValid(QtObjectPtr ptr){
	return static_cast<QStorageInfo*>(ptr)->isValid();
}

char* QStorageInfo_Name(QtObjectPtr ptr){
	return static_cast<QStorageInfo*>(ptr)->name().toUtf8().data();
}

void QStorageInfo_Refresh(QtObjectPtr ptr){
	static_cast<QStorageInfo*>(ptr)->refresh();
}

char* QStorageInfo_RootPath(QtObjectPtr ptr){
	return static_cast<QStorageInfo*>(ptr)->rootPath().toUtf8().data();
}

void QStorageInfo_SetPath(QtObjectPtr ptr, char* path){
	static_cast<QStorageInfo*>(ptr)->setPath(QString(path));
}

void QStorageInfo_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QStorageInfo*>(ptr)->swap(*static_cast<QStorageInfo*>(other));
}

void QStorageInfo_DestroyQStorageInfo(QtObjectPtr ptr){
	static_cast<QStorageInfo*>(ptr)->~QStorageInfo();
}

