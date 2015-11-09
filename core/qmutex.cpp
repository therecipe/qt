#include "qmutex.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMutex>
#include "_cgo_export.h"

class MyQMutex: public QMutex {
public:
};

void QMutex_Lock(void* ptr){
	static_cast<QMutex*>(ptr)->lock();
}

int QMutex_TryLock(void* ptr, int timeout){
	return static_cast<QMutex*>(ptr)->tryLock(timeout);
}

void QMutex_Unlock(void* ptr){
	static_cast<QMutex*>(ptr)->unlock();
}

void* QMutex_NewQMutex(int mode){
	return new QMutex(static_cast<QMutex::RecursionMode>(mode));
}

int QMutex_IsRecursive(void* ptr){
	return static_cast<QMutex*>(ptr)->isRecursive();
}

