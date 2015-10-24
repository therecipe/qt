#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDBusConnection_NewQDBusConnection2(QtObjectPtr other);
QtObjectPtr QDBusConnection_NewQDBusConnection(char* name);
char* QDBusConnection_BaseService(QtObjectPtr ptr);
int QDBusConnection_CallWithCallback(QtObjectPtr ptr, QtObjectPtr message, QtObjectPtr receiver, char* returnMethod, char* errorMethod, int timeout);
int QDBusConnection_Connect(QtObjectPtr ptr, char* service, char* path, char* interfa, char* name, QtObjectPtr receiver, char* slot);
int QDBusConnection_Connect2(QtObjectPtr ptr, char* service, char* path, char* interfa, char* name, char* signature, QtObjectPtr receiver, char* slot);
int QDBusConnection_Connect3(QtObjectPtr ptr, char* service, char* path, char* interfa, char* name, char* argumentMatch, char* signature, QtObjectPtr receiver, char* slot);
int QDBusConnection_ConnectionCapabilities(QtObjectPtr ptr);
void QDBusConnection_QDBusConnection_DisconnectFromBus(char* name);
void QDBusConnection_QDBusConnection_DisconnectFromPeer(char* name);
QtObjectPtr QDBusConnection_Interface(QtObjectPtr ptr);
int QDBusConnection_IsConnected(QtObjectPtr ptr);
char* QDBusConnection_Name(QtObjectPtr ptr);
QtObjectPtr QDBusConnection_ObjectRegisteredAt(QtObjectPtr ptr, char* path);
int QDBusConnection_RegisterObject(QtObjectPtr ptr, char* path, QtObjectPtr object, int options);
int QDBusConnection_RegisterObject2(QtObjectPtr ptr, char* path, char* interfa, QtObjectPtr object, int options);
int QDBusConnection_RegisterService(QtObjectPtr ptr, char* serviceName);
int QDBusConnection_Send(QtObjectPtr ptr, QtObjectPtr message);
void QDBusConnection_UnregisterObject(QtObjectPtr ptr, char* path, int mode);
int QDBusConnection_UnregisterService(QtObjectPtr ptr, char* serviceName);
void QDBusConnection_DestroyQDBusConnection(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif