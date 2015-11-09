#include "qthreadpool.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QThread>
#include <QRunnable>
#include <QString>
#include <QVariant>
#include <QThreadPool>
#include "_cgo_export.h"

class MyQThreadPool: public QThreadPool {
public:
};

int QThreadPool_ActiveThreadCount(void* ptr){
	return static_cast<QThreadPool*>(ptr)->activeThreadCount();
}

int QThreadPool_ExpiryTimeout(void* ptr){
	return static_cast<QThreadPool*>(ptr)->expiryTimeout();
}

int QThreadPool_MaxThreadCount(void* ptr){
	return static_cast<QThreadPool*>(ptr)->maxThreadCount();
}

void QThreadPool_SetExpiryTimeout(void* ptr, int expiryTimeout){
	static_cast<QThreadPool*>(ptr)->setExpiryTimeout(expiryTimeout);
}

void QThreadPool_SetMaxThreadCount(void* ptr, int maxThreadCount){
	static_cast<QThreadPool*>(ptr)->setMaxThreadCount(maxThreadCount);
}

void* QThreadPool_NewQThreadPool(void* parent){
	return new QThreadPool(static_cast<QObject*>(parent));
}

void QThreadPool_Cancel(void* ptr, void* runnable){
	static_cast<QThreadPool*>(ptr)->cancel(static_cast<QRunnable*>(runnable));
}

void QThreadPool_Clear(void* ptr){
	static_cast<QThreadPool*>(ptr)->clear();
}

void* QThreadPool_QThreadPool_GlobalInstance(){
	return QThreadPool::globalInstance();
}

void QThreadPool_ReleaseThread(void* ptr){
	static_cast<QThreadPool*>(ptr)->releaseThread();
}

void QThreadPool_ReserveThread(void* ptr){
	static_cast<QThreadPool*>(ptr)->reserveThread();
}

void QThreadPool_Start(void* ptr, void* runnable, int priority){
	static_cast<QThreadPool*>(ptr)->start(static_cast<QRunnable*>(runnable), priority);
}

int QThreadPool_TryStart(void* ptr, void* runnable){
	return static_cast<QThreadPool*>(ptr)->tryStart(static_cast<QRunnable*>(runnable));
}

int QThreadPool_WaitForDone(void* ptr, int msecs){
	return static_cast<QThreadPool*>(ptr)->waitForDone(msecs);
}

void QThreadPool_DestroyQThreadPool(void* ptr){
	static_cast<QThreadPool*>(ptr)->~QThreadPool();
}

