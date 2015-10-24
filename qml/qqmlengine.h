#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QQmlEngine_OfflineStoragePath(QtObjectPtr ptr);
void QQmlEngine_SetOfflineStoragePath(QtObjectPtr ptr, char* dir);
QtObjectPtr QQmlEngine_NewQQmlEngine(QtObjectPtr parent);
void QQmlEngine_AddImageProvider(QtObjectPtr ptr, char* providerId, QtObjectPtr provider);
void QQmlEngine_AddImportPath(QtObjectPtr ptr, char* path);
void QQmlEngine_AddPluginPath(QtObjectPtr ptr, char* path);
char* QQmlEngine_BaseUrl(QtObjectPtr ptr);
void QQmlEngine_ClearComponentCache(QtObjectPtr ptr);
QtObjectPtr QQmlEngine_QQmlEngine_ContextForObject(QtObjectPtr object);
QtObjectPtr QQmlEngine_ImageProvider(QtObjectPtr ptr, char* providerId);
char* QQmlEngine_ImportPathList(QtObjectPtr ptr);
QtObjectPtr QQmlEngine_IncubationController(QtObjectPtr ptr);
QtObjectPtr QQmlEngine_NetworkAccessManager(QtObjectPtr ptr);
QtObjectPtr QQmlEngine_NetworkAccessManagerFactory(QtObjectPtr ptr);
int QQmlEngine_QQmlEngine_ObjectOwnership(QtObjectPtr object);
int QQmlEngine_OutputWarningsToStandardError(QtObjectPtr ptr);
char* QQmlEngine_PluginPathList(QtObjectPtr ptr);
void QQmlEngine_ConnectQuit(QtObjectPtr ptr);
void QQmlEngine_DisconnectQuit(QtObjectPtr ptr);
void QQmlEngine_RemoveImageProvider(QtObjectPtr ptr, char* providerId);
QtObjectPtr QQmlEngine_RootContext(QtObjectPtr ptr);
void QQmlEngine_SetBaseUrl(QtObjectPtr ptr, char* url);
void QQmlEngine_QQmlEngine_SetContextForObject(QtObjectPtr object, QtObjectPtr context);
void QQmlEngine_SetImportPathList(QtObjectPtr ptr, char* paths);
void QQmlEngine_SetIncubationController(QtObjectPtr ptr, QtObjectPtr controller);
void QQmlEngine_SetNetworkAccessManagerFactory(QtObjectPtr ptr, QtObjectPtr factory);
void QQmlEngine_QQmlEngine_SetObjectOwnership(QtObjectPtr object, int ownership);
void QQmlEngine_SetOutputWarningsToStandardError(QtObjectPtr ptr, int enabled);
void QQmlEngine_SetPluginPathList(QtObjectPtr ptr, char* paths);
void QQmlEngine_TrimComponentCache(QtObjectPtr ptr);
void QQmlEngine_DestroyQQmlEngine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif