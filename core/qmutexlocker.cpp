#include "qmutexlocker.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMutex>
#include <QMutexLocker>
#include "_cgo_export.h"

class MyQMutexLocker: public QMutexLocker {
public:
};

void* QMutexLocker_NewQMutexLocker(void* mutex){
	return new QMutexLocker(static_cast<QMutex*>(mutex));
}

void* QMutexLocker_Mutex(void* ptr){
	return static_cast<QMutexLocker*>(ptr)->mutex();
}

void QMutexLocker_Relock(void* ptr){
	static_cast<QMutexLocker*>(ptr)->relock();
}

void QMutexLocker_Unlock(void* ptr){
	static_cast<QMutexLocker*>(ptr)->unlock();
}

void QMutexLocker_DestroyQMutexLocker(void* ptr){
	static_cast<QMutexLocker*>(ptr)->~QMutexLocker();
}

