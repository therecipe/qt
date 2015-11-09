#include "qreadlocker.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QReadWriteLock>
#include <QReadLocker>
#include "_cgo_export.h"

class MyQReadLocker: public QReadLocker {
public:
};

void* QReadLocker_NewQReadLocker(void* lock){
	return new QReadLocker(static_cast<QReadWriteLock*>(lock));
}

void* QReadLocker_ReadWriteLock(void* ptr){
	return static_cast<QReadLocker*>(ptr)->readWriteLock();
}

void QReadLocker_Relock(void* ptr){
	static_cast<QReadLocker*>(ptr)->relock();
}

void QReadLocker_Unlock(void* ptr){
	static_cast<QReadLocker*>(ptr)->unlock();
}

void QReadLocker_DestroyQReadLocker(void* ptr){
	static_cast<QReadLocker*>(ptr)->~QReadLocker();
}

