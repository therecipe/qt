#ifdef __cplusplus
extern "C" {
#endif

int QRunnable_AutoDelete(void* ptr);
void QRunnable_Run(void* ptr);
void QRunnable_SetAutoDelete(void* ptr, int autoDelete);
void QRunnable_DestroyQRunnable(void* ptr);

#ifdef __cplusplus
}
#endif