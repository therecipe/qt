#include "qreadlocker.h"
#include <QUrl>
#include <QModelIndex>
#include <QReadWriteLock>
#include <QString>
#include <QVariant>
#include <QReadLocker>
#include "_cgo_export.h"

class MyQReadLocker: public QReadLocker {
public:
};

QtObjectPtr QReadLocker_NewQReadLocker(QtObjectPtr lock){
	return new QReadLocker(static_cast<QReadWriteLock*>(lock));
}

QtObjectPtr QReadLocker_ReadWriteLock(QtObjectPtr ptr){
	return static_cast<QReadLocker*>(ptr)->readWriteLock();
}

void QReadLocker_Relock(QtObjectPtr ptr){
	static_cast<QReadLocker*>(ptr)->relock();
}

void QReadLocker_Unlock(QtObjectPtr ptr){
	static_cast<QReadLocker*>(ptr)->unlock();
}

void QReadLocker_DestroyQReadLocker(QtObjectPtr ptr){
	static_cast<QReadLocker*>(ptr)->~QReadLocker();
}

