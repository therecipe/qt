#include "qthreadpool.h"
#include <QRunnable>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QThread>
#include <QThreadPool>
#include "_cgo_export.h"

class MyQThreadPool: public QThreadPool {
public:
};

int QThreadPool_ActiveThreadCount(QtObjectPtr ptr){
	return static_cast<QThreadPool*>(ptr)->activeThreadCount();
}

int QThreadPool_ExpiryTimeout(QtObjectPtr ptr){
	return static_cast<QThreadPool*>(ptr)->expiryTimeout();
}

int QThreadPool_MaxThreadCount(QtObjectPtr ptr){
	return static_cast<QThreadPool*>(ptr)->maxThreadCount();
}

void QThreadPool_SetExpiryTimeout(QtObjectPtr ptr, int expiryTimeout){
	static_cast<QThreadPool*>(ptr)->setExpiryTimeout(expiryTimeout);
}

void QThreadPool_SetMaxThreadCount(QtObjectPtr ptr, int maxThreadCount){
	static_cast<QThreadPool*>(ptr)->setMaxThreadCount(maxThreadCount);
}

QtObjectPtr QThreadPool_NewQThreadPool(QtObjectPtr parent){
	return new QThreadPool(static_cast<QObject*>(parent));
}

void QThreadPool_Cancel(QtObjectPtr ptr, QtObjectPtr runnable){
	static_cast<QThreadPool*>(ptr)->cancel(static_cast<QRunnable*>(runnable));
}

void QThreadPool_Clear(QtObjectPtr ptr){
	static_cast<QThreadPool*>(ptr)->clear();
}

QtObjectPtr QThreadPool_QThreadPool_GlobalInstance(){
	return QThreadPool::globalInstance();
}

void QThreadPool_ReleaseThread(QtObjectPtr ptr){
	static_cast<QThreadPool*>(ptr)->releaseThread();
}

void QThreadPool_ReserveThread(QtObjectPtr ptr){
	static_cast<QThreadPool*>(ptr)->reserveThread();
}

void QThreadPool_Start(QtObjectPtr ptr, QtObjectPtr runnable, int priority){
	static_cast<QThreadPool*>(ptr)->start(static_cast<QRunnable*>(runnable), priority);
}

int QThreadPool_TryStart(QtObjectPtr ptr, QtObjectPtr runnable){
	return static_cast<QThreadPool*>(ptr)->tryStart(static_cast<QRunnable*>(runnable));
}

int QThreadPool_WaitForDone(QtObjectPtr ptr, int msecs){
	return static_cast<QThreadPool*>(ptr)->waitForDone(msecs);
}

void QThreadPool_DestroyQThreadPool(QtObjectPtr ptr){
	static_cast<QThreadPool*>(ptr)->~QThreadPool();
}

