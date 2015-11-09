#ifdef __cplusplus
extern "C" {
#endif

void* QQmlContext_NewQQmlContext2(void* parentContext, void* parent);
void* QQmlContext_NewQQmlContext(void* engine, void* parent);
void* QQmlContext_ContextObject(void* ptr);
void* QQmlContext_ContextProperty(void* ptr, char* name);
void* QQmlContext_Engine(void* ptr);
int QQmlContext_IsValid(void* ptr);
char* QQmlContext_NameForObject(void* ptr, void* object);
void* QQmlContext_ParentContext(void* ptr);
void QQmlContext_SetBaseUrl(void* ptr, void* baseUrl);
void QQmlContext_SetContextObject(void* ptr, void* object);
void QQmlContext_SetContextProperty(void* ptr, char* name, void* value);
void QQmlContext_SetContextProperty2(void* ptr, char* name, void* value);
void QQmlContext_DestroyQQmlContext(void* ptr);

#ifdef __cplusplus
}
#endif