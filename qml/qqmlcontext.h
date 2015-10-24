#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlContext_NewQQmlContext2(QtObjectPtr parentContext, QtObjectPtr parent);
QtObjectPtr QQmlContext_NewQQmlContext(QtObjectPtr engine, QtObjectPtr parent);
char* QQmlContext_BaseUrl(QtObjectPtr ptr);
QtObjectPtr QQmlContext_ContextObject(QtObjectPtr ptr);
char* QQmlContext_ContextProperty(QtObjectPtr ptr, char* name);
QtObjectPtr QQmlContext_Engine(QtObjectPtr ptr);
int QQmlContext_IsValid(QtObjectPtr ptr);
char* QQmlContext_NameForObject(QtObjectPtr ptr, QtObjectPtr object);
QtObjectPtr QQmlContext_ParentContext(QtObjectPtr ptr);
char* QQmlContext_ResolvedUrl(QtObjectPtr ptr, char* src);
void QQmlContext_SetBaseUrl(QtObjectPtr ptr, char* baseUrl);
void QQmlContext_SetContextObject(QtObjectPtr ptr, QtObjectPtr object);
void QQmlContext_SetContextProperty(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QQmlContext_SetContextProperty2(QtObjectPtr ptr, char* name, char* value);
void QQmlContext_DestroyQQmlContext(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif