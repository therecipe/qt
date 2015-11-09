#include "qstorageinfo.h"
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDir>
#include <QStorageInfo>
#include "_cgo_export.h"

class MyQStorageInfo: public QStorageInfo {
public:
};

void* QStorageInfo_NewQStorageInfo(){
	return new QStorageInfo();
}

void* QStorageInfo_NewQStorageInfo3(void* dir){
	return new QStorageInfo(*static_cast<QDir*>(dir));
}

void* QStorageInfo_NewQStorageInfo4(void* other){
	return new QStorageInfo(*static_cast<QStorageInfo*>(other));
}

void* QStorageInfo_NewQStorageInfo2(char* path){
	return new QStorageInfo(QString(path));
}

void* QStorageInfo_Device(void* ptr){
	return new QByteArray(static_cast<QStorageInfo*>(ptr)->device());
}

char* QStorageInfo_DisplayName(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->displayName().toUtf8().data();
}

void* QStorageInfo_FileSystemType(void* ptr){
	return new QByteArray(static_cast<QStorageInfo*>(ptr)->fileSystemType());
}

int QStorageInfo_IsReadOnly(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->isReadOnly();
}

int QStorageInfo_IsReady(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->isReady();
}

int QStorageInfo_IsRoot(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->isRoot();
}

int QStorageInfo_IsValid(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->isValid();
}

char* QStorageInfo_Name(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->name().toUtf8().data();
}

void QStorageInfo_Refresh(void* ptr){
	static_cast<QStorageInfo*>(ptr)->refresh();
}

char* QStorageInfo_RootPath(void* ptr){
	return static_cast<QStorageInfo*>(ptr)->rootPath().toUtf8().data();
}

void QStorageInfo_SetPath(void* ptr, char* path){
	static_cast<QStorageInfo*>(ptr)->setPath(QString(path));
}

void QStorageInfo_Swap(void* ptr, void* other){
	static_cast<QStorageInfo*>(ptr)->swap(*static_cast<QStorageInfo*>(other));
}

void QStorageInfo_DestroyQStorageInfo(void* ptr){
	static_cast<QStorageInfo*>(ptr)->~QStorageInfo();
}

