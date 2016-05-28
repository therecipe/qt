// +build !minimal

#pragma once

#ifndef GO_QTMACEXTRAS_H
#define GO_QTMACEXTRAS_H

#ifdef __cplusplus
extern "C" {
#endif

int QMacPasteboardMime_CanConvert(void* ptr, char* mime, char* flav);
char* QMacPasteboardMime_ConvertorName(void* ptr);
int QMacPasteboardMime_Count(void* ptr, void* mimeData);
int QMacPasteboardMime_CountDefault(void* ptr, void* mimeData);
char* QMacPasteboardMime_FlavorFor(void* ptr, char* mime);
char* QMacPasteboardMime_MimeFor(void* ptr, char* flav);
void QMacPasteboardMime_DestroyQMacPasteboardMime(void* ptr);
char* QMacPasteboardMime_ObjectNameAbs(void* ptr);
void QMacPasteboardMime_SetObjectNameAbs(void* ptr, char* name);
void* QMacToolBar_NewQMacToolBar(void* parent);
void* QMacToolBar_NewQMacToolBar2(char* identifier, void* parent);
void* QMacToolBar_AddAllowedItem(void* ptr, void* icon, char* text);
void* QMacToolBar_AddItem(void* ptr, void* icon, char* text);
void QMacToolBar_AddSeparator(void* ptr);
void QMacToolBar_AttachToWindow(void* ptr, void* window);
void QMacToolBar_DetachFromWindow(void* ptr);
void QMacToolBar_DestroyQMacToolBar(void* ptr);
void QMacToolBar_TimerEvent(void* ptr, void* event);
void QMacToolBar_TimerEventDefault(void* ptr, void* event);
void QMacToolBar_ChildEvent(void* ptr, void* event);
void QMacToolBar_ChildEventDefault(void* ptr, void* event);
void QMacToolBar_ConnectNotify(void* ptr, void* sign);
void QMacToolBar_ConnectNotifyDefault(void* ptr, void* sign);
void QMacToolBar_CustomEvent(void* ptr, void* event);
void QMacToolBar_CustomEventDefault(void* ptr, void* event);
void QMacToolBar_DeleteLater(void* ptr);
void QMacToolBar_DeleteLaterDefault(void* ptr);
void QMacToolBar_DisconnectNotify(void* ptr, void* sign);
void QMacToolBar_DisconnectNotifyDefault(void* ptr, void* sign);
int QMacToolBar_Event(void* ptr, void* e);
int QMacToolBar_EventDefault(void* ptr, void* e);
int QMacToolBar_EventFilter(void* ptr, void* watched, void* event);
int QMacToolBar_EventFilterDefault(void* ptr, void* watched, void* event);
void* QMacToolBar_MetaObject(void* ptr);
void* QMacToolBar_MetaObjectDefault(void* ptr);
void* QMacToolBarItem_NewQMacToolBarItem(void* parent);
void QMacToolBarItem_ConnectActivated(void* ptr);
void QMacToolBarItem_DisconnectActivated(void* ptr);
void QMacToolBarItem_Activated(void* ptr);
void QMacToolBarItem_DestroyQMacToolBarItem(void* ptr);
void* QMacToolBarItem_Icon(void* ptr);
int QMacToolBarItem_Selectable(void* ptr);
void QMacToolBarItem_SetIcon(void* ptr, void* icon);
void QMacToolBarItem_SetSelectable(void* ptr, int selectable);
void QMacToolBarItem_SetStandardItem(void* ptr, int standardItem);
void QMacToolBarItem_SetText(void* ptr, char* text);
int QMacToolBarItem_StandardItem(void* ptr);
char* QMacToolBarItem_Text(void* ptr);
void QMacToolBarItem_TimerEvent(void* ptr, void* event);
void QMacToolBarItem_TimerEventDefault(void* ptr, void* event);
void QMacToolBarItem_ChildEvent(void* ptr, void* event);
void QMacToolBarItem_ChildEventDefault(void* ptr, void* event);
void QMacToolBarItem_ConnectNotify(void* ptr, void* sign);
void QMacToolBarItem_ConnectNotifyDefault(void* ptr, void* sign);
void QMacToolBarItem_CustomEvent(void* ptr, void* event);
void QMacToolBarItem_CustomEventDefault(void* ptr, void* event);
void QMacToolBarItem_DeleteLater(void* ptr);
void QMacToolBarItem_DeleteLaterDefault(void* ptr);
void QMacToolBarItem_DisconnectNotify(void* ptr, void* sign);
void QMacToolBarItem_DisconnectNotifyDefault(void* ptr, void* sign);
int QMacToolBarItem_Event(void* ptr, void* e);
int QMacToolBarItem_EventDefault(void* ptr, void* e);
int QMacToolBarItem_EventFilter(void* ptr, void* watched, void* event);
int QMacToolBarItem_EventFilterDefault(void* ptr, void* watched, void* event);
void* QMacToolBarItem_MetaObject(void* ptr);
void* QMacToolBarItem_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif