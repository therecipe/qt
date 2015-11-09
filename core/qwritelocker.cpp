#include "qwritelocker.h"
#include <QReadWriteLock>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWriteLocker>
#include "_cgo_export.h"

class MyQWriteLocker: public QWriteLocker {
public:
};

void* QWriteLocker_NewQWriteLocker(void* lock){
	return new QWriteLocker(static_cast<QReadWriteLock*>(lock));
}

void* QWriteLocker_ReadWriteLock(void* ptr){
	return static_cast<QWriteLocker*>(ptr)->readWriteLock();
}

void QWriteLocker_Relock(void* ptr){
	static_cast<QWriteLocker*>(ptr)->relock();
}

void QWriteLocker_Unlock(void* ptr){
	static_cast<QWriteLocker*>(ptr)->unlock();
}

void QWriteLocker_DestroyQWriteLocker(void* ptr){
	static_cast<QWriteLocker*>(ptr)->~QWriteLocker();
}

