#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QAbstractEventDispatcher_ConnectAboutToBlock(QtObjectPtr ptr);
void QAbstractEventDispatcher_DisconnectAboutToBlock(QtObjectPtr ptr);
void QAbstractEventDispatcher_ConnectAwake(QtObjectPtr ptr);
void QAbstractEventDispatcher_DisconnectAwake(QtObjectPtr ptr);
void QAbstractEventDispatcher_Flush(QtObjectPtr ptr);
void QAbstractEventDispatcher_InstallNativeEventFilter(QtObjectPtr ptr, QtObjectPtr filterObj);
QtObjectPtr QAbstractEventDispatcher_QAbstractEventDispatcher_Instance(QtObjectPtr thread);
void QAbstractEventDispatcher_Interrupt(QtObjectPtr ptr);
int QAbstractEventDispatcher_ProcessEvents(QtObjectPtr ptr, int flags);
void QAbstractEventDispatcher_RegisterSocketNotifier(QtObjectPtr ptr, QtObjectPtr notifier);
int QAbstractEventDispatcher_RemainingTime(QtObjectPtr ptr, int timerId);
void QAbstractEventDispatcher_RemoveNativeEventFilter(QtObjectPtr ptr, QtObjectPtr filter);
void QAbstractEventDispatcher_UnregisterSocketNotifier(QtObjectPtr ptr, QtObjectPtr notifier);
int QAbstractEventDispatcher_UnregisterTimer(QtObjectPtr ptr, int timerId);
int QAbstractEventDispatcher_UnregisterTimers(QtObjectPtr ptr, QtObjectPtr object);
void QAbstractEventDispatcher_WakeUp(QtObjectPtr ptr);
void QAbstractEventDispatcher_DestroyQAbstractEventDispatcher(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif