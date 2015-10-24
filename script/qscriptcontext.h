#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QScriptContext_ArgumentCount(QtObjectPtr ptr);
char* QScriptContext_Backtrace(QtObjectPtr ptr);
QtObjectPtr QScriptContext_Engine(QtObjectPtr ptr);
int QScriptContext_IsCalledAsConstructor(QtObjectPtr ptr);
QtObjectPtr QScriptContext_ParentContext(QtObjectPtr ptr);
void QScriptContext_SetActivationObject(QtObjectPtr ptr, QtObjectPtr activation);
void QScriptContext_SetThisObject(QtObjectPtr ptr, QtObjectPtr thisObject);
char* QScriptContext_ToString(QtObjectPtr ptr);
void QScriptContext_DestroyQScriptContext(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif