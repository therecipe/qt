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

void* QSemaphore_NewQSemaphore(int n){
	return new QSemaphore(n);
}

void QSemaphore_Acquire(void* ptr, int n){
	static_cast<QSemaphore*>(ptr)->acquire(n);
}

int QSemaphore_Available(void* ptr){
	return static_cast<QSemaphore*>(ptr)->available();
}

void QSemaphore_Release(void* ptr, int n){
	static_cast<QSemaphore*>(ptr)->release(n);
}

int QSemaphore_TryAcquire(void* ptr, int n){
	return static_cast<QSemaphore*>(ptr)->tryAcquire(n);
}

int QSemaphore_TryAcquire2(void* ptr, int n, int timeout){
	return static_cast<QSemaphore*>(ptr)->tryAcquire(n, timeout);
}

void QSemaphore_DestroyQSemaphore(void* ptr){
	static_cast<QSemaphore*>(ptr)->~QSemaphore();
}

