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

void QMutex_Lock(QtObjectPtr ptr){
	static_cast<QMutex*>(ptr)->lock();
}

int QMutex_TryLock(QtObjectPtr ptr, int timeout){
	return static_cast<QMutex*>(ptr)->tryLock(timeout);
}

void QMutex_Unlock(QtObjectPtr ptr){
	static_cast<QMutex*>(ptr)->unlock();
}

QtObjectPtr QMutex_NewQMutex(int mode){
	return new QMutex(static_cast<QMutex::RecursionMode>(mode));
}

int QMutex_IsRecursive(QtObjectPtr ptr){
	return static_cast<QMutex*>(ptr)->isRecursive();
}

