#include "qsemaphore.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSemaphore>
#include "_cgo_export.h"

class MyQSemaphore: public QSemaphore {
public:
};

QtObjectPtr QSemaphore_NewQSemaphore(int n){
	return new QSemaphore(n);
}

void QSemaphore_Acquire(QtObjectPtr ptr, int n){
	static_cast<QSemaphore*>(ptr)->acquire(n);
}

int QSemaphore_Available(QtObjectPtr ptr){
	return static_cast<QSemaphore*>(ptr)->available();
}

void QSemaphore_Release(QtObjectPtr ptr, int n){
	static_cast<QSemaphore*>(ptr)->release(n);
}

int QSemaphore_TryAcquire(QtObjectPtr ptr, int n){
	return static_cast<QSemaphore*>(ptr)->tryAcquire(n);
}

int QSemaphore_TryAcquire2(QtObjectPtr ptr, int n, int timeout){
	return static_cast<QSemaphore*>(ptr)->tryAcquire(n, timeout);
}

void QSemaphore_DestroyQSemaphore(QtObjectPtr ptr){
	static_cast<QSemaphore*>(ptr)->~QSemaphore();
}

