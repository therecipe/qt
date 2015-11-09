#ifdef __cplusplus
extern "C" {
#endif

void* QQmlFileSelector_NewQQmlFileSelector(void* engine, void* parent);
void* QQmlFileSelector_QQmlFileSelector_Get(void* engine);
void QQmlFileSelector_SetExtraSelectors(void* ptr, char* strin);
void QQmlFileSelector_SetExtraSelectors2(void* ptr, char* strin);
void QQmlFileSelector_SetSelector(void* ptr, void* selector);
void QQmlFileSelector_DestroyQQmlFileSelector(void* ptr);

#ifdef __cplusplus
}
#endif