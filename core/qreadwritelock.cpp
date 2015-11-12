#include "qreadwritelock.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QReadWriteLock>
#include "_cgo_export.h"

class MyQReadWriteLock: public QReadWriteLock {
public:
};

void* QReadWriteLock_NewQReadWriteLock(int recursionMode){
	return new QReadWriteLock(static_cast<QReadWriteLock::RecursionMode>(recursionMode));
}

void QReadWriteLock_LockForRead(void* ptr){
	static_cast<QReadWriteLock*>(ptr)->lockForRead();
}

void QReadWriteLock_LockForWrite(void* ptr){
	static_cast<QReadWriteLock*>(ptr)->lockForWrite();
}

int QReadWriteLock_TryLockForRead(void* ptr){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForRead();
}

int QReadWriteLock_TryLockForRead2(void* ptr, int timeout){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForRead(timeout);
}

int QReadWriteLock_TryLockForWrite(void* ptr){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForWrite();
}

int QReadWriteLock_TryLockForWrite2(void* ptr, int timeout){
	return static_cast<QReadWriteLock*>(ptr)->tryLockForWrite(timeout);
}

void QReadWriteLock_Unlock(void* ptr){
	static_cast<QReadWriteLock*>(ptr)->unlock();
}

void QReadWriteLock_DestroyQReadWriteLock(void* ptr){
	static_cast<QReadWriteLock*>(ptr)->~QReadWriteLock();
}

