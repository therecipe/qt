#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QQmlComponent_Status(QtObjectPtr ptr);
char* QQmlComponent_Url(QtObjectPtr ptr);
QtObjectPtr QQmlComponent_NewQQmlComponent(QtObjectPtr engine, QtObjectPtr parent);
QtObjectPtr QQmlComponent_NewQQmlComponent4(QtObjectPtr engine, char* fileName, int mode, QtObjectPtr parent);
QtObjectPtr QQmlComponent_NewQQmlComponent3(QtObjectPtr engine, char* fileName, QtObjectPtr parent);
QtObjectPtr QQmlComponent_NewQQmlComponent6(QtObjectPtr engine, char* url, int mode, QtObjectPtr parent);
QtObjectPtr QQmlComponent_NewQQmlComponent5(QtObjectPtr engine, char* url, QtObjectPtr parent);
QtObjectPtr QQmlComponent_BeginCreate(QtObjectPtr ptr, QtObjectPtr publicContext);
void QQmlComponent_CompleteCreate(QtObjectPtr ptr);
QtObjectPtr QQmlComponent_Create(QtObjectPtr ptr, QtObjectPtr context);
void QQmlComponent_Create2(QtObjectPtr ptr, QtObjectPtr incubator, QtObjectPtr context, QtObjectPtr forContext);
QtObjectPtr QQmlComponent_CreationContext(QtObjectPtr ptr);
int QQmlComponent_IsError(QtObjectPtr ptr);
int QQmlComponent_IsLoading(QtObjectPtr ptr);
int QQmlComponent_IsNull(QtObjectPtr ptr);
int QQmlComponent_IsReady(QtObjectPtr ptr);
void QQmlComponent_LoadUrl(QtObjectPtr ptr, char* url);
void QQmlComponent_LoadUrl2(QtObjectPtr ptr, char* url, int mode);
void QQmlComponent_SetData(QtObjectPtr ptr, QtObjectPtr data, char* url);
void QQmlComponent_ConnectStatusChanged(QtObjectPtr ptr);
void QQmlComponent_DisconnectStatusChanged(QtObjectPtr ptr);
void QQmlComponent_DestroyQQmlComponent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif