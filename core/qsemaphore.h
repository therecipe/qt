#ifdef __cplusplus
extern "C" {
#endif

void* QSemaphore_NewQSemaphore(int n);
void QSemaphore_Acquire(void* ptr, int n);
int QSemaphore_Available(void* ptr);
void QSemaphore_Release(void* ptr, int n);
int QSemaphore_TryAcquire(void* ptr, int n);
int QSemaphore_TryAcquire2(void* ptr, int n, int timeout);
void QSemaphore_DestroyQSemaphore(void* ptr);

#ifdef __cplusplus
}
#endif