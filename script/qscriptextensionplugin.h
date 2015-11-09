#ifdef __cplusplus
extern "C" {
#endif

void QScriptExtensionPlugin_Initialize(void* ptr, char* key, void* engine);
char* QScriptExtensionPlugin_Keys(void* ptr);
void* QScriptExtensionPlugin_SetupPackage(void* ptr, char* key, void* engine);
void QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(void* ptr);

#ifdef __cplusplus
}
#endif