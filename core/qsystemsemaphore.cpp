#include "qsystemsemaphore.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSystemSemaphore>
#include "_cgo_export.h"

class MyQSystemSemaphore: public QSystemSemaphore {
public:
};

void* QSystemSemaphore_NewQSystemSemaphore(char* key, int initialValue, int mode){
	return new QSystemSemaphore(QString(key), initialValue, static_cast<QSystemSemaphore::AccessMode>(mode));
}

int QSystemSemaphore_Acquire(void* ptr){
	return static_cast<QSystemSemaphore*>(ptr)->acquire();
}

int QSystemSemaphore_Error(void* ptr){
	return static_cast<QSystemSemaphore*>(ptr)->error();
}

char* QSystemSemaphore_ErrorString(void* ptr){
	return static_cast<QSystemSemaphore*>(ptr)->errorString().toUtf8().data();
}

char* QSystemSemaphore_Key(void* ptr){
	return static_cast<QSystemSemaphore*>(ptr)->key().toUtf8().data();
}

int QSystemSemaphore_Release(void* ptr, int n){
	return static_cast<QSystemSemaphore*>(ptr)->release(n);
}

void QSystemSemaphore_SetKey(void* ptr, char* key, int initialValue, int mode){
	static_cast<QSystemSemaphore*>(ptr)->setKey(QString(key), initialValue, static_cast<QSystemSemaphore::AccessMode>(mode));
}

void QSystemSemaphore_DestroyQSystemSemaphore(void* ptr){
	static_cast<QSystemSemaphore*>(ptr)->~QSystemSemaphore();
}

