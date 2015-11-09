#ifdef __cplusplus
extern "C" {
#endif

void* QEventLoop_NewQEventLoop(void* parent);
int QEventLoop_Event(void* ptr, void* event);
int QEventLoop_Exec(void* ptr, int flags);
void QEventLoop_Exit(void* ptr, int returnCode);
int QEventLoop_IsRunning(void* ptr);
int QEventLoop_ProcessEvents(void* ptr, int flags);
void QEventLoop_ProcessEvents2(void* ptr, int flags, int maxTime);
void QEventLoop_Quit(void* ptr);
void QEventLoop_WakeUp(void* ptr);
void QEventLoop_DestroyQEventLoop(void* ptr);

#ifdef __cplusplus
}
#endif