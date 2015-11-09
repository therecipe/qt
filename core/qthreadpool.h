#ifdef __cplusplus
extern "C" {
#endif

int QThreadPool_ActiveThreadCount(void* ptr);
int QThreadPool_ExpiryTimeout(void* ptr);
int QThreadPool_MaxThreadCount(void* ptr);
void QThreadPool_SetExpiryTimeout(void* ptr, int expiryTimeout);
void QThreadPool_SetMaxThreadCount(void* ptr, int maxThreadCount);
void* QThreadPool_NewQThreadPool(void* parent);
void QThreadPool_Cancel(void* ptr, void* runnable);
void QThreadPool_Clear(void* ptr);
void* QThreadPool_QThreadPool_GlobalInstance();
void QThreadPool_ReleaseThread(void* ptr);
void QThreadPool_ReserveThread(void* ptr);
void QThreadPool_Start(void* ptr, void* runnable, int priority);
int QThreadPool_TryStart(void* ptr, void* runnable);
int QThreadPool_WaitForDone(void* ptr, int msecs);
void QThreadPool_DestroyQThreadPool(void* ptr);

#ifdef __cplusplus
}
#endif