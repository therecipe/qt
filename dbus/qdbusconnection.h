#ifdef __cplusplus
extern "C" {
#endif

void* QDBusConnection_NewQDBusConnection2(void* other);
void* QDBusConnection_NewQDBusConnection(char* name);
char* QDBusConnection_BaseService(void* ptr);
int QDBusConnection_CallWithCallback(void* ptr, void* message, void* receiver, char* returnMethod, char* errorMethod, int timeout);
int QDBusConnection_Connect(void* ptr, char* service, char* path, char* interfa, char* name, void* receiver, char* slot);
int QDBusConnection_Connect2(void* ptr, char* service, char* path, char* interfa, char* name, char* signature, void* receiver, char* slot);
int QDBusConnection_Connect3(void* ptr, char* service, char* path, char* interfa, char* name, char* argumentMatch, char* signature, void* receiver, char* slot);
int QDBusConnection_ConnectionCapabilities(void* ptr);
void QDBusConnection_QDBusConnection_DisconnectFromBus(char* name);
void QDBusConnection_QDBusConnection_DisconnectFromPeer(char* name);
void* QDBusConnection_Interface(void* ptr);
int QDBusConnection_IsConnected(void* ptr);
void* QDBusConnection_QDBusConnection_LocalMachineId();
char* QDBusConnection_Name(void* ptr);
void* QDBusConnection_ObjectRegisteredAt(void* ptr, char* path);
int QDBusConnection_RegisterObject(void* ptr, char* path, void* object, int options);
int QDBusConnection_RegisterObject2(void* ptr, char* path, char* interfa, void* object, int options);
int QDBusConnection_RegisterService(void* ptr, char* serviceName);
int QDBusConnection_Send(void* ptr, void* message);
void QDBusConnection_UnregisterObject(void* ptr, char* path, int mode);
int QDBusConnection_UnregisterService(void* ptr, char* serviceName);
void QDBusConnection_DestroyQDBusConnection(void* ptr);

#ifdef __cplusplus
}
#endif