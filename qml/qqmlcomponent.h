#ifdef __cplusplus
extern "C" {
#endif

double QQmlComponent_Progress(void* ptr);
int QQmlComponent_Status(void* ptr);
void* QQmlComponent_NewQQmlComponent(void* engine, void* parent);
void* QQmlComponent_NewQQmlComponent4(void* engine, char* fileName, int mode, void* parent);
void* QQmlComponent_NewQQmlComponent3(void* engine, char* fileName, void* parent);
void* QQmlComponent_NewQQmlComponent6(void* engine, void* url, int mode, void* parent);
void* QQmlComponent_NewQQmlComponent5(void* engine, void* url, void* parent);
void* QQmlComponent_BeginCreate(void* ptr, void* publicContext);
void QQmlComponent_CompleteCreate(void* ptr);
void* QQmlComponent_Create(void* ptr, void* context);
void QQmlComponent_Create2(void* ptr, void* incubator, void* context, void* forContext);
void* QQmlComponent_CreationContext(void* ptr);
int QQmlComponent_IsError(void* ptr);
int QQmlComponent_IsLoading(void* ptr);
int QQmlComponent_IsNull(void* ptr);
int QQmlComponent_IsReady(void* ptr);
void QQmlComponent_LoadUrl(void* ptr, void* url);
void QQmlComponent_LoadUrl2(void* ptr, void* url, int mode);
void QQmlComponent_SetData(void* ptr, void* data, void* url);
void QQmlComponent_ConnectStatusChanged(void* ptr);
void QQmlComponent_DisconnectStatusChanged(void* ptr);
void QQmlComponent_DestroyQQmlComponent(void* ptr);

#ifdef __cplusplus
}
#endif