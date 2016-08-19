

#pragma once

#ifndef GO_MAIN_H
#define GO_MAIN_H

#ifdef __cplusplus
extern "C" {
#endif

void* Delegate_NewDelegate(void* parent);
void Delegate_DestroyDelegate(void* ptr);
void* Delegate_CreateEditor(void* ptr, void* parent, void* option, void* index);
void* Delegate_CreateEditorDefault(void* ptr, void* parent, void* option, void* index);
char* Delegate_DisplayText(void* ptr, void* value, void* locale);
char* Delegate_DisplayTextDefault(void* ptr, void* value, void* locale);
int Delegate_EditorEvent(void* ptr, void* event, void* model, void* option, void* index);
int Delegate_EditorEventDefault(void* ptr, void* event, void* model, void* option, void* index);
void Delegate_InitStyleOption(void* ptr, void* option, void* index);
void Delegate_InitStyleOptionDefault(void* ptr, void* option, void* index);
void Delegate_Paint(void* ptr, void* painter, void* option, void* index);
void Delegate_PaintDefault(void* ptr, void* painter, void* option, void* index);
void Delegate_SetEditorData(void* ptr, void* editor, void* index);
void Delegate_SetEditorDataDefault(void* ptr, void* editor, void* index);
void Delegate_SetModelData(void* ptr, void* editor, void* model, void* index);
void Delegate_SetModelDataDefault(void* ptr, void* editor, void* model, void* index);
void* Delegate_SizeHint(void* ptr, void* option, void* index);
void* Delegate_SizeHintDefault(void* ptr, void* option, void* index);
void Delegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index);
void Delegate_UpdateEditorGeometryDefault(void* ptr, void* editor, void* option, void* index);
void Delegate_DestroyEditor(void* ptr, void* editor, void* index);
void Delegate_DestroyEditorDefault(void* ptr, void* editor, void* index);
int Delegate_HelpEvent(void* ptr, void* event, void* view, void* option, void* index);
int Delegate_HelpEventDefault(void* ptr, void* event, void* view, void* option, void* index);
void Delegate_TimerEvent(void* ptr, void* event);
void Delegate_TimerEventDefault(void* ptr, void* event);
void Delegate_ChildEvent(void* ptr, void* event);
void Delegate_ChildEventDefault(void* ptr, void* event);
void Delegate_ConnectNotify(void* ptr, void* sign);
void Delegate_ConnectNotifyDefault(void* ptr, void* sign);
void Delegate_CustomEvent(void* ptr, void* event);
void Delegate_CustomEventDefault(void* ptr, void* event);
void Delegate_DeleteLater(void* ptr);
void Delegate_DeleteLaterDefault(void* ptr);
void Delegate_DisconnectNotify(void* ptr, void* sign);
void Delegate_DisconnectNotifyDefault(void* ptr, void* sign);
int Delegate_Event(void* ptr, void* e);
int Delegate_EventDefault(void* ptr, void* e);
;
;

#ifdef __cplusplus
}
#endif

#endif