#ifdef __cplusplus
extern "C" {
#endif

void QDBusPendingCallWatcher_WaitForFinished(void* ptr);
void* QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(void* call, void* parent);
void QDBusPendingCallWatcher_ConnectFinished(void* ptr);
void QDBusPendingCallWatcher_DisconnectFinished(void* ptr);
int QDBusPendingCallWatcher_IsFinished(void* ptr);
void QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(void* ptr);

#ifdef __cplusplus
}
#endif