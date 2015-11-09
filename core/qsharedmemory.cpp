#include "qsharedmemory.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSharedMemory>
#include "_cgo_export.h"

class MyQSharedMemory: public QSharedMemory {
public:
};

void* QSharedMemory_NewQSharedMemory2(void* parent){
	return new QSharedMemory(static_cast<QObject*>(parent));
}

void* QSharedMemory_NewQSharedMemory(char* key, void* parent){
	return new QSharedMemory(QString(key), static_cast<QObject*>(parent));
}

int QSharedMemory_Attach(void* ptr, int mode){
	return static_cast<QSharedMemory*>(ptr)->attach(static_cast<QSharedMemory::AccessMode>(mode));
}

void* QSharedMemory_ConstData(void* ptr){
	return const_cast<void*>(static_cast<QSharedMemory*>(ptr)->constData());
}

int QSharedMemory_Create(void* ptr, int size, int mode){
	return static_cast<QSharedMemory*>(ptr)->create(size, static_cast<QSharedMemory::AccessMode>(mode));
}

void* QSharedMemory_Data(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->data();
}

void* QSharedMemory_Data2(void* ptr){
	return const_cast<void*>(static_cast<QSharedMemory*>(ptr)->data());
}

int QSharedMemory_Detach(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->detach();
}

int QSharedMemory_Error(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->error();
}

char* QSharedMemory_ErrorString(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->errorString().toUtf8().data();
}

int QSharedMemory_IsAttached(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->isAttached();
}

char* QSharedMemory_Key(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->key().toUtf8().data();
}

int QSharedMemory_Lock(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->lock();
}

char* QSharedMemory_NativeKey(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->nativeKey().toUtf8().data();
}

void QSharedMemory_SetKey(void* ptr, char* key){
	static_cast<QSharedMemory*>(ptr)->setKey(QString(key));
}

void QSharedMemory_SetNativeKey(void* ptr, char* key){
	static_cast<QSharedMemory*>(ptr)->setNativeKey(QString(key));
}

int QSharedMemory_Size(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->size();
}

int QSharedMemory_Unlock(void* ptr){
	return static_cast<QSharedMemory*>(ptr)->unlock();
}

void QSharedMemory_DestroyQSharedMemory(void* ptr){
	static_cast<QSharedMemory*>(ptr)->~QSharedMemory();
}

