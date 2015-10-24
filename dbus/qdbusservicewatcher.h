#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QDBusServiceWatcher_SetWatchMode(QtObjectPtr ptr, int mode);
int QDBusServiceWatcher_WatchMode(QtObjectPtr ptr);
QtObjectPtr QDBusServiceWatcher_NewQDBusServiceWatcher(QtObjectPtr parent);
QtObjectPtr QDBusServiceWatcher_NewQDBusServiceWatcher2(char* service, QtObjectPtr connection, int watchMode, QtObjectPtr parent);
void QDBusServiceWatcher_AddWatchedService(QtObjectPtr ptr, char* newService);
int QDBusServiceWatcher_RemoveWatchedService(QtObjectPtr ptr, char* service);
void QDBusServiceWatcher_ConnectServiceOwnerChanged(QtObjectPtr ptr);
void QDBusServiceWatcher_DisconnectServiceOwnerChanged(QtObjectPtr ptr);
void QDBusServiceWatcher_ConnectServiceRegistered(QtObjectPtr ptr);
void QDBusServiceWatcher_DisconnectServiceRegistered(QtObjectPtr ptr);
void QDBusServiceWatcher_ConnectServiceUnregistered(QtObjectPtr ptr);
void QDBusServiceWatcher_DisconnectServiceUnregistered(QtObjectPtr ptr);
void QDBusServiceWatcher_SetConnection(QtObjectPtr ptr, QtObjectPtr connection);
void QDBusServiceWatcher_SetWatchedServices(QtObjectPtr ptr, char* services);
char* QDBusServiceWatcher_WatchedServices(QtObjectPtr ptr);
void QDBusServiceWatcher_DestroyQDBusServiceWatcher(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif