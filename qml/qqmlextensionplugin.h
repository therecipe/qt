#ifdef __cplusplus
extern "C" {
#endif

void QQmlExtensionPlugin_InitializeEngine(void* ptr, void* engine, char* uri);
void QQmlExtensionPlugin_RegisterTypes(void* ptr, char* uri);

#ifdef __cplusplus
}
#endif