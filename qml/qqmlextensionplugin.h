#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QQmlExtensionPlugin_InitializeEngine(QtObjectPtr ptr, QtObjectPtr engine, char* uri);
char* QQmlExtensionPlugin_BaseUrl(QtObjectPtr ptr);
void QQmlExtensionPlugin_RegisterTypes(QtObjectPtr ptr, char* uri);

#ifdef __cplusplus
}
#endif