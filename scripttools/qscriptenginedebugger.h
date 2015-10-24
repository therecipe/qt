#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScriptEngineDebugger_NewQScriptEngineDebugger(QtObjectPtr parent);
QtObjectPtr QScriptEngineDebugger_Action(QtObjectPtr ptr, int action);
void QScriptEngineDebugger_AttachTo(QtObjectPtr ptr, QtObjectPtr engine);
int QScriptEngineDebugger_AutoShowStandardWindow(QtObjectPtr ptr);
QtObjectPtr QScriptEngineDebugger_CreateStandardMenu(QtObjectPtr ptr, QtObjectPtr parent);
QtObjectPtr QScriptEngineDebugger_CreateStandardToolBar(QtObjectPtr ptr, QtObjectPtr parent);
void QScriptEngineDebugger_Detach(QtObjectPtr ptr);
void QScriptEngineDebugger_ConnectEvaluationResumed(QtObjectPtr ptr);
void QScriptEngineDebugger_DisconnectEvaluationResumed(QtObjectPtr ptr);
void QScriptEngineDebugger_ConnectEvaluationSuspended(QtObjectPtr ptr);
void QScriptEngineDebugger_DisconnectEvaluationSuspended(QtObjectPtr ptr);
void QScriptEngineDebugger_SetAutoShowStandardWindow(QtObjectPtr ptr, int autoShow);
QtObjectPtr QScriptEngineDebugger_StandardWindow(QtObjectPtr ptr);
QtObjectPtr QScriptEngineDebugger_Widget(QtObjectPtr ptr, int widget);
void QScriptEngineDebugger_DestroyQScriptEngineDebugger(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif