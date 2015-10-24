#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QScriptExtensionPlugin_Initialize(QtObjectPtr ptr, char* key, QtObjectPtr engine);
char* QScriptExtensionPlugin_Keys(QtObjectPtr ptr);
void QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif