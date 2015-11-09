#ifdef __cplusplus
extern "C" {
#endif

void QAbstractEventDispatcher_ConnectAboutToBlock(void* ptr);
void QAbstractEventDispatcher_DisconnectAboutToBlock(void* ptr);
void QAbstractEventDispatcher_ConnectAwake(void* ptr);
void QAbstractEventDispatcher_DisconnectAwake(void* ptr);
void QAbstractEventDispatcher_Flush(void* ptr);
void QAbstractEventDispatcher_InstallNativeEventFilter(void* ptr, void* filterObj);
void* QAbstractEventDispatcher_QAbstractEventDispatcher_Instance(void* thread);
void QAbstractEventDispatcher_Interrupt(void* ptr);
int QAbstractEventDispatcher_ProcessEvents(void* ptr, int flags);
void QAbstractEventDispatcher_RegisterSocketNotifier(void* ptr, void* notifier);
int QAbstractEventDispatcher_RemainingTime(void* ptr, int timerId);
void QAbstractEventDispatcher_RemoveNativeEventFilter(void* ptr, void* filter);
void QAbstractEventDispatcher_UnregisterSocketNotifier(void* ptr, void* notifier);
int QAbstractEventDispatcher_UnregisterTimer(void* ptr, int timerId);
int QAbstractEventDispatcher_UnregisterTimers(void* ptr, void* object);
void QAbstractEventDispatcher_WakeUp(void* ptr);
void QAbstractEventDispatcher_DestroyQAbstractEventDispatcher(void* ptr);

#ifdef __cplusplus
}
#endif