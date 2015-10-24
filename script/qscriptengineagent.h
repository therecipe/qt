#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScriptEngineAgent_NewQScriptEngineAgent(QtObjectPtr engine);
void QScriptEngineAgent_ContextPop(QtObjectPtr ptr);
void QScriptEngineAgent_ContextPush(QtObjectPtr ptr);
QtObjectPtr QScriptEngineAgent_Engine(QtObjectPtr ptr);
char* QScriptEngineAgent_Extension(QtObjectPtr ptr, int extension, char* argument);
int QScriptEngineAgent_SupportsExtension(QtObjectPtr ptr, int extension);
void QScriptEngineAgent_DestroyQScriptEngineAgent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif