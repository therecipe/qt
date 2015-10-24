#include "qsharedmemory.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSharedMemory>
#include "_cgo_export.h"

class MyQSharedMemory: public QSharedMemory {
public:
};

QtObjectPtr QSharedMemory_NewQSharedMemory2(QtObjectPtr parent){
	return new QSharedMemory(static_cast<QObject*>(parent));
}

QtObjectPtr QSharedMemory_NewQSharedMemory(char* key, QtObjectPtr parent){
	return new QSharedMemory(QString(key), static_cast<QObject*>(parent));
}

int QSharedMemory_Attach(QtObjectPtr ptr, int mode){
	return static_cast<QSharedMemory*>(ptr)->attach(static_cast<QSharedMemory::AccessMode>(mode));
}

void QSharedMemory_ConstData(QtObjectPtr ptr){
	static_cast<QSharedMemory*>(ptr)->constData();
}

int QSharedMemory_Create(QtObjectPtr ptr, int size, int mode){
	return static_cast<QSharedMemory*>(ptr)->create(size, static_cast<QSharedMemory::AccessMode>(mode));
}

void QSharedMemory_Data(QtObjectPtr ptr){
	static_cast<QSharedMemory*>(ptr)->data();
}

void QSharedMemory_Data2(QtObjectPtr ptr){
	static_cast<QSharedMemory*>(ptr)->data();
}

int QSharedMemory_Detach(QtObjectPtr ptr){
	return static_cast<QSharedMemory*>(ptr)->detach();
}

int QSharedMemory_Error(QtObjectPtr ptr){
	return static_cast<QSharedMemory*>(ptr)->error();
}

char* QSharedMemory_ErrorString(QtObjectPtr ptr){
	return static_cast<QSharedMemory*>(ptr)->errorString().toUtf8().data();
}

int QSharedMemory_IsAttached(QtObjectPtr ptr){
	return static_cast<QSharedMemory*>(ptr)->isAttached();
}

char* QSharedMemory_Key(QtObjectPtr ptr){
	return static_cast<QSharedMemory*>(ptr)->key().toUtf8().data();
}

int QSharedMemory_Lock(QtObjectPtr ptr){
	return static_cast<QSharedMemory*>(ptr)->lock();
}

char* QSharedMemory_NativeKey(QtObjectPtr ptr){
	return static_cast<QSharedMemory*>(ptr)->nativeKey().toUtf8().data();
}

void QSharedMemory_SetKey(QtObjectPtr ptr, char* key){
	static_cast<QSharedMemory*>(ptr)->setKey(QString(key));
}

void QSharedMemory_SetNativeKey(QtObjectPtr ptr, char* key){
	static_cast<QSharedMemory*>(ptr)->setNativeKey(QString(key));
}

int QSharedMemory_Size(QtObjectPtr ptr){
	return static_cast<QSharedMemory*>(ptr)->size();
}

int QSharedMemory_Unlock(QtObjectPtr ptr){
	return static_cast<QSharedMemory*>(ptr)->unlock();
}

void QSharedMemory_DestroyQSharedMemory(QtObjectPtr ptr){
	static_cast<QSharedMemory*>(ptr)->~QSharedMemory();
}

