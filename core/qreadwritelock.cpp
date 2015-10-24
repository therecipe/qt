#include "qreadwritelock.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QReadWriteLock>
#include "_cgo_export.h"

class MyQReadWriteLock: public QReadWriteLock {
public:
};

QtObjectPtr QReadWriteLock_NewQReadWriteLock(int recursionMode){
	return new QReadWriteLock(static_cast<QReadWriteLock::RecursionMode>(recursionMode));
}

void QReadWriteLock_LockForRead(QtObjectPtr ptr){
	static_cast<QReadWriteLock*>(ptr)->lockForRead();
}

void QReadWriteLock_LockForWrite(QtObjectPtr ptr){
	static_cast<QReadWriteLock*>(ptr)->lockForWrite();
}

int QReadWriteLock_TryLockForRead(QtObjectPtr ptr){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForRead();
}

int QReadWriteLock_TryLockForRead2(QtObjectPtr ptr, int timeout){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForRead(timeout);
}

int QReadWriteLock_TryLockForWrite(QtObjectPtr ptr){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForWrite();
}

int QReadWriteLock_TryLockForWrite2(QtObjectPtr ptr, int timeout){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForWrite(timeout);
}

void QReadWriteLock_Unlock(QtObjectPtr ptr){
	static_cast<QReadWriteLock*>(ptr)->unlock();
}

void QReadWriteLock_DestroyQReadWriteLock(QtObjectPtr ptr){
	static_cast<QReadWriteLock*>(ptr)->~QReadWriteLock();
}

