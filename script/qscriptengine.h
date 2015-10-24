#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScriptEngine_NewQScriptEngine();
QtObjectPtr QScriptEngine_NewQScriptEngine2(QtObjectPtr parent);
void QScriptEngine_AbortEvaluation(QtObjectPtr ptr, QtObjectPtr result);
QtObjectPtr QScriptEngine_Agent(QtObjectPtr ptr);
char* QScriptEngine_AvailableExtensions(QtObjectPtr ptr);
void QScriptEngine_ClearExceptions(QtObjectPtr ptr);
void QScriptEngine_CollectGarbage(QtObjectPtr ptr);
QtObjectPtr QScriptEngine_CurrentContext(QtObjectPtr ptr);
int QScriptEngine_HasUncaughtException(QtObjectPtr ptr);
char* QScriptEngine_ImportedExtensions(QtObjectPtr ptr);
void QScriptEngine_InstallTranslatorFunctions(QtObjectPtr ptr, QtObjectPtr object);
int QScriptEngine_IsEvaluating(QtObjectPtr ptr);
void QScriptEngine_PopContext(QtObjectPtr ptr);
int QScriptEngine_ProcessEventsInterval(QtObjectPtr ptr);
QtObjectPtr QScriptEngine_PushContext(QtObjectPtr ptr);
void QScriptEngine_ReportAdditionalMemoryCost(QtObjectPtr ptr, int size);
void QScriptEngine_SetAgent(QtObjectPtr ptr, QtObjectPtr agent);
void QScriptEngine_SetDefaultPrototype(QtObjectPtr ptr, int metaTypeId, QtObjectPtr prototype);
void QScriptEngine_SetGlobalObject(QtObjectPtr ptr, QtObjectPtr object);
void QScriptEngine_SetProcessEventsInterval(QtObjectPtr ptr, int interval);
char* QScriptEngine_UncaughtExceptionBacktrace(QtObjectPtr ptr);
int QScriptEngine_UncaughtExceptionLineNumber(QtObjectPtr ptr);
void QScriptEngine_DestroyQScriptEngine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif