#include "qwritelocker.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QReadWriteLock>
#include <QWriteLocker>
#include "_cgo_export.h"

class MyQWriteLocker: public QWriteLocker {
public:
};

QtObjectPtr QWriteLocker_NewQWriteLocker(QtObjectPtr lock){
	return new QWriteLocker(static_cast<QReadWriteLock*>(lock));
}

QtObjectPtr QWriteLocker_ReadWriteLock(QtObjectPtr ptr){
	return static_cast<QWriteLocker*>(ptr)->readWriteLock();
}

void QWriteLocker_Relock(QtObjectPtr ptr){
	static_cast<QWriteLocker*>(ptr)->relock();
}

void QWriteLocker_Unlock(QtObjectPtr ptr){
	static_cast<QWriteLocker*>(ptr)->unlock();
}

void QWriteLocker_DestroyQWriteLocker(QtObjectPtr ptr){
	static_cast<QWriteLocker*>(ptr)->~QWriteLocker();
}

