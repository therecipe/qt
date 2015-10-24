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

QtObjectPtr QMutexLocker_NewQMutexLocker(QtObjectPtr mutex){
	return new QMutexLocker(static_cast<QMutex*>(mutex));
}

QtObjectPtr QMutexLocker_Mutex(QtObjectPtr ptr){
	return static_cast<QMutexLocker*>(ptr)->mutex();
}

void QMutexLocker_Relock(QtObjectPtr ptr){
	static_cast<QMutexLocker*>(ptr)->relock();
}

void QMutexLocker_Unlock(QtObjectPtr ptr){
	static_cast<QMutexLocker*>(ptr)->unlock();
}

void QMutexLocker_DestroyQMutexLocker(QtObjectPtr ptr){
	static_cast<QMutexLocker*>(ptr)->~QMutexLocker();
}

