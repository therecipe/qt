#ifdef __cplusplus
extern "C" {
#endif

void QMutex_Lock(void* ptr);
int QMutex_TryLock(void* ptr, int timeout);
void QMutex_Unlock(void* ptr);
void* QMutex_NewQMutex(int mode);
int QMutex_IsRecursive(void* ptr);

#ifdef __cplusplus
}
#endif