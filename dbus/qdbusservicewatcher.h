#ifdef __cplusplus
extern "C" {
#endif

void QDBusServiceWatcher_SetWatchMode(void* ptr, int mode);
int QDBusServiceWatcher_WatchMode(void* ptr);
void* QDBusServiceWatcher_NewQDBusServiceWatcher(void* parent);
void* QDBusServiceWatcher_NewQDBusServiceWatcher2(char* service, void* connection, int watchMode, void* parent);
void QDBusServiceWatcher_AddWatchedService(void* ptr, char* newService);
int QDBusServiceWatcher_RemoveWatchedService(void* ptr, char* service);
void QDBusServiceWatcher_ConnectServiceOwnerChanged(void* ptr);
void QDBusServiceWatcher_DisconnectServiceOwnerChanged(void* ptr);
void QDBusServiceWatcher_ConnectServiceRegistered(void* ptr);
void QDBusServiceWatcher_DisconnectServiceRegistered(void* ptr);
void QDBusServiceWatcher_ConnectServiceUnregistered(void* ptr);
void QDBusServiceWatcher_DisconnectServiceUnregistered(void* ptr);
void QDBusServiceWatcher_SetConnection(void* ptr, void* connection);
void QDBusServiceWatcher_SetWatchedServices(void* ptr, char* services);
char* QDBusServiceWatcher_WatchedServices(void* ptr);
void QDBusServiceWatcher_DestroyQDBusServiceWatcher(void* ptr);

#ifdef __cplusplus
}
#endif