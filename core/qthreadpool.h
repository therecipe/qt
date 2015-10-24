#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QThreadPool_ActiveThreadCount(QtObjectPtr ptr);
int QThreadPool_ExpiryTimeout(QtObjectPtr ptr);
int QThreadPool_MaxThreadCount(QtObjectPtr ptr);
void QThreadPool_SetExpiryTimeout(QtObjectPtr ptr, int expiryTimeout);
void QThreadPool_SetMaxThreadCount(QtObjectPtr ptr, int maxThreadCount);
QtObjectPtr QThreadPool_NewQThreadPool(QtObjectPtr parent);
void QThreadPool_Cancel(QtObjectPtr ptr, QtObjectPtr runnable);
void QThreadPool_Clear(QtObjectPtr ptr);
QtObjectPtr QThreadPool_QThreadPool_GlobalInstance();
void QThreadPool_ReleaseThread(QtObjectPtr ptr);
void QThreadPool_ReserveThread(QtObjectPtr ptr);
void QThreadPool_Start(QtObjectPtr ptr, QtObjectPtr runnable, int priority);
int QThreadPool_TryStart(QtObjectPtr ptr, QtObjectPtr runnable);
int QThreadPool_WaitForDone(QtObjectPtr ptr, int msecs);
void QThreadPool_DestroyQThreadPool(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif