// +build !minimal

#pragma once

#ifndef GO_QTSCRIPTTOOLS_H
#define GO_QTSCRIPTTOOLS_H

#include <stdint.h>

#ifdef __cplusplus
int QScriptEngineDebugger_QScriptEngineDebugger_QRegisterMetaType();
extern "C" {
#endif

struct QtScriptTools_PackedString { char* data; long long len; void* ptr; };
struct QtScriptTools_PackedList { void* data; long long len; };
void* QScriptEngineDebugger_NewQScriptEngineDebugger(void* parent);
void* QScriptEngineDebugger_Action(void* ptr, long long action);
void QScriptEngineDebugger_AttachTo(void* ptr, void* engine);
char QScriptEngineDebugger_AutoShowStandardWindow(void* ptr);
void* QScriptEngineDebugger_CreateStandardMenu(void* ptr, void* parent);
void* QScriptEngineDebugger_CreateStandardToolBar(void* ptr, void* parent);
void QScriptEngineDebugger_Detach(void* ptr);
void QScriptEngineDebugger_ConnectEvaluationResumed(void* ptr, long long t);
void QScriptEngineDebugger_DisconnectEvaluationResumed(void* ptr);
void QScriptEngineDebugger_EvaluationResumed(void* ptr);
void QScriptEngineDebugger_ConnectEvaluationSuspended(void* ptr, long long t);
void QScriptEngineDebugger_DisconnectEvaluationSuspended(void* ptr);
void QScriptEngineDebugger_EvaluationSuspended(void* ptr);
void QScriptEngineDebugger_SetAutoShowStandardWindow(void* ptr, char autoShow);
void* QScriptEngineDebugger_StandardWindow(void* ptr);
long long QScriptEngineDebugger_State(void* ptr);
void* QScriptEngineDebugger_Widget(void* ptr, long long widget);
void QScriptEngineDebugger_DestroyQScriptEngineDebugger(void* ptr);
void QScriptEngineDebugger_DestroyQScriptEngineDebuggerDefault(void* ptr);
void* QScriptEngineDebugger___children_atList(void* ptr, int i);
void QScriptEngineDebugger___children_setList(void* ptr, void* i);
void* QScriptEngineDebugger___children_newList(void* ptr);
void* QScriptEngineDebugger___dynamicPropertyNames_atList(void* ptr, int i);
void QScriptEngineDebugger___dynamicPropertyNames_setList(void* ptr, void* i);
void* QScriptEngineDebugger___dynamicPropertyNames_newList(void* ptr);
void* QScriptEngineDebugger___findChildren_atList(void* ptr, int i);
void QScriptEngineDebugger___findChildren_setList(void* ptr, void* i);
void* QScriptEngineDebugger___findChildren_newList(void* ptr);
void* QScriptEngineDebugger___findChildren_atList3(void* ptr, int i);
void QScriptEngineDebugger___findChildren_setList3(void* ptr, void* i);
void* QScriptEngineDebugger___findChildren_newList3(void* ptr);
void QScriptEngineDebugger_ChildEventDefault(void* ptr, void* event);
void QScriptEngineDebugger_ConnectNotifyDefault(void* ptr, void* sign);
void QScriptEngineDebugger_CustomEventDefault(void* ptr, void* event);
void QScriptEngineDebugger_DeleteLaterDefault(void* ptr);
void QScriptEngineDebugger_DisconnectNotifyDefault(void* ptr, void* sign);
char QScriptEngineDebugger_EventDefault(void* ptr, void* e);
char QScriptEngineDebugger_EventFilterDefault(void* ptr, void* watched, void* event);
void* QScriptEngineDebugger_MetaObjectDefault(void* ptr);
void QScriptEngineDebugger_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif