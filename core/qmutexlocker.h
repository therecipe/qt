#ifdef __cplusplus
extern "C" {
#endif

void* QMutexLocker_NewQMutexLocker(void* mutex);
void* QMutexLocker_Mutex(void* ptr);
void QMutexLocker_Relock(void* ptr);
void QMutexLocker_Unlock(void* ptr);
void QMutexLocker_DestroyQMutexLocker(void* ptr);

#ifdef __cplusplus
}
#endif