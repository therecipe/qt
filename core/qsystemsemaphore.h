#ifdef __cplusplus
extern "C" {
#endif

void* QSystemSemaphore_NewQSystemSemaphore(char* key, int initialValue, int mode);
int QSystemSemaphore_Acquire(void* ptr);
int QSystemSemaphore_Error(void* ptr);
char* QSystemSemaphore_ErrorString(void* ptr);
char* QSystemSemaphore_Key(void* ptr);
int QSystemSemaphore_Release(void* ptr, int n);
void QSystemSemaphore_SetKey(void* ptr, char* key, int initialValue, int mode);
void QSystemSemaphore_DestroyQSystemSemaphore(void* ptr);

#ifdef __cplusplus
}
#endif