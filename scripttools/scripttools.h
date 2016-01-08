#pragma once

#ifndef GO_QTSCRIPTTOOLS_H
#define GO_QTSCRIPTTOOLS_H

#ifdef __cplusplus
extern "C" {
#endif

void* QScriptEngineDebugger_NewQScriptEngineDebugger(void* parent);
void* QScriptEngineDebugger_Action(void* ptr, int action);
void QScriptEngineDebugger_AttachTo(void* ptr, void* engine);
int QScriptEngineDebugger_AutoShowStandardWindow(void* ptr);
void* QScriptEngineDebugger_CreateStandardMenu(void* ptr, void* parent);
void* QScriptEngineDebugger_CreateStandardToolBar(void* ptr, void* parent);
void QScriptEngineDebugger_Detach(void* ptr);
void QScriptEngineDebugger_ConnectEvaluationResumed(void* ptr);
void QScriptEngineDebugger_DisconnectEvaluationResumed(void* ptr);
void QScriptEngineDebugger_EvaluationResumed(void* ptr);
void QScriptEngineDebugger_ConnectEvaluationSuspended(void* ptr);
void QScriptEngineDebugger_DisconnectEvaluationSuspended(void* ptr);
void QScriptEngineDebugger_EvaluationSuspended(void* ptr);
void QScriptEngineDebugger_SetAutoShowStandardWindow(void* ptr, int autoShow);
void* QScriptEngineDebugger_StandardWindow(void* ptr);
int QScriptEngineDebugger_State(void* ptr);
void* QScriptEngineDebugger_Widget(void* ptr, int widget);
void QScriptEngineDebugger_DestroyQScriptEngineDebugger(void* ptr);
void QScriptEngineDebugger_TimerEvent(void* ptr, void* event);
void QScriptEngineDebugger_TimerEventDefault(void* ptr, void* event);
void QScriptEngineDebugger_ChildEvent(void* ptr, void* event);
void QScriptEngineDebugger_ChildEventDefault(void* ptr, void* event);
void QScriptEngineDebugger_CustomEvent(void* ptr, void* event);
void QScriptEngineDebugger_CustomEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif