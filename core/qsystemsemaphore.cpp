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

QtObjectPtr QSystemSemaphore_NewQSystemSemaphore(char* key, int initialValue, int mode){
	return new QSystemSemaphore(QString(key), initialValue, static_cast<QSystemSemaphore::AccessMode>(mode));
}

int QSystemSemaphore_Acquire(QtObjectPtr ptr){
	return static_cast<QSystemSemaphore*>(ptr)->acquire();
}

int QSystemSemaphore_Error(QtObjectPtr ptr){
	return static_cast<QSystemSemaphore*>(ptr)->error();
}

char* QSystemSemaphore_ErrorString(QtObjectPtr ptr){
	return static_cast<QSystemSemaphore*>(ptr)->errorString().toUtf8().data();
}

char* QSystemSemaphore_Key(QtObjectPtr ptr){
	return static_cast<QSystemSemaphore*>(ptr)->key().toUtf8().data();
}

int QSystemSemaphore_Release(QtObjectPtr ptr, int n){
	return static_cast<QSystemSemaphore*>(ptr)->release(n);
}

void QSystemSemaphore_SetKey(QtObjectPtr ptr, char* key, int initialValue, int mode){
	static_cast<QSystemSemaphore*>(ptr)->setKey(QString(key), initialValue, static_cast<QSystemSemaphore::AccessMode>(mode));
}

void QSystemSemaphore_DestroyQSystemSemaphore(QtObjectPtr ptr){
	static_cast<QSystemSemaphore*>(ptr)->~QSystemSemaphore();
}

