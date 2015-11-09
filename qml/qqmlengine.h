#ifdef __cplusplus
extern "C" {
#endif

char* QQmlEngine_OfflineStoragePath(void* ptr);
void QQmlEngine_SetOfflineStoragePath(void* ptr, char* dir);
void* QQmlEngine_NewQQmlEngine(void* parent);
void QQmlEngine_AddImageProvider(void* ptr, char* providerId, void* provider);
void QQmlEngine_AddImportPath(void* ptr, char* path);
void QQmlEngine_AddPluginPath(void* ptr, char* path);
void QQmlEngine_ClearComponentCache(void* ptr);
void* QQmlEngine_QQmlEngine_ContextForObject(void* object);
void* QQmlEngine_ImageProvider(void* ptr, char* providerId);
char* QQmlEngine_ImportPathList(void* ptr);
void* QQmlEngine_IncubationController(void* ptr);
void* QQmlEngine_NetworkAccessManager(void* ptr);
void* QQmlEngine_NetworkAccessManagerFactory(void* ptr);
int QQmlEngine_QQmlEngine_ObjectOwnership(void* object);
int QQmlEngine_OutputWarningsToStandardError(void* ptr);
char* QQmlEngine_PluginPathList(void* ptr);
void QQmlEngine_ConnectQuit(void* ptr);
void QQmlEngine_DisconnectQuit(void* ptr);
void QQmlEngine_RemoveImageProvider(void* ptr, char* providerId);
void* QQmlEngine_RootContext(void* ptr);
void QQmlEngine_SetBaseUrl(void* ptr, void* url);
void QQmlEngine_QQmlEngine_SetContextForObject(void* object, void* context);
void QQmlEngine_SetImportPathList(void* ptr, char* paths);
void QQmlEngine_SetIncubationController(void* ptr, void* controller);
void QQmlEngine_SetNetworkAccessManagerFactory(void* ptr, void* factory);
void QQmlEngine_QQmlEngine_SetObjectOwnership(void* object, int ownership);
void QQmlEngine_SetOutputWarningsToStandardError(void* ptr, int enabled);
void QQmlEngine_SetPluginPathList(void* ptr, char* paths);
void QQmlEngine_TrimComponentCache(void* ptr);
void QQmlEngine_DestroyQQmlEngine(void* ptr);

#ifdef __cplusplus
}
#endif