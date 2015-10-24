#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QDBusPendingCallWatcher_WaitForFinished(QtObjectPtr ptr);
QtObjectPtr QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(QtObjectPtr call, QtObjectPtr parent);
void QDBusPendingCallWatcher_ConnectFinished(QtObjectPtr ptr);
void QDBusPendingCallWatcher_DisconnectFinished(QtObjectPtr ptr);
int QDBusPendingCallWatcher_IsFinished(QtObjectPtr ptr);
void QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif