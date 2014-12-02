#include "qthreadpool.h"
#include <QThreadPool>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QThreadPool_New_QObject(QtObjectPtr parent)
{
	return new QThreadPool(((QObject*)(parent)));
}

void QThreadPool_Destroy(QtObjectPtr ptr)
{
	((QThreadPool*)(ptr))->~QThreadPool();
}

int QThreadPool_ActiveThreadCount(QtObjectPtr ptr)
{
	return ((QThreadPool*)(ptr))->activeThreadCount();
}

void QThreadPool_Clear(QtObjectPtr ptr)
{
	((QThreadPool*)(ptr))->clear();
}

int QThreadPool_ExpiryTimeout(QtObjectPtr ptr)
{
	return ((QThreadPool*)(ptr))->expiryTimeout();
}

int QThreadPool_MaxThreadCount(QtObjectPtr ptr)
{
	return ((QThreadPool*)(ptr))->maxThreadCount();
}

void QThreadPool_ReleaseThread(QtObjectPtr ptr)
{
	((QThreadPool*)(ptr))->releaseThread();
}

void QThreadPool_ReserveThread(QtObjectPtr ptr)
{
	((QThreadPool*)(ptr))->reserveThread();
}

void QThreadPool_SetExpiryTimeout_Int(QtObjectPtr ptr, int expiryTimeout)
{
	((QThreadPool*)(ptr))->setExpiryTimeout(expiryTimeout);
}

void QThreadPool_SetMaxThreadCount_Int(QtObjectPtr ptr, int maxThreadCount)
{
	((QThreadPool*)(ptr))->setMaxThreadCount(maxThreadCount);
}

int QThreadPool_WaitForDone_Int(QtObjectPtr ptr, int msecs)
{
	return ((QThreadPool*)(ptr))->waitForDone(msecs);
}

//Static Public Members
QtObjectPtr QThreadPool_GlobalInstance()
{
	return QThreadPool::globalInstance();
}

