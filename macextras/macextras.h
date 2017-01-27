// +build !minimal

#pragma once

#ifndef GO_QTMACEXTRAS_H
#define GO_QTMACEXTRAS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtMacExtras_PackedString { char* data; long long len; };
struct QtMacExtras_PackedList { void* data; long long len; };
void* QMacPasteboardMime_NewQMacPasteboardMime(char* t);
char QMacPasteboardMime_CanConvert(void* ptr, char* mime, char* flav);
struct QtMacExtras_PackedList QMacPasteboardMime_ConvertFromMime(void* ptr, char* mime, void* data, char* flav);
void* QMacPasteboardMime_ConvertToMime(void* ptr, char* mime, void* data, char* flav);
struct QtMacExtras_PackedString QMacPasteboardMime_ConvertorName(void* ptr);
int QMacPasteboardMime_Count(void* ptr, void* mimeData);
int QMacPasteboardMime_CountDefault(void* ptr, void* mimeData);
struct QtMacExtras_PackedString QMacPasteboardMime_FlavorFor(void* ptr, char* mime);
struct QtMacExtras_PackedString QMacPasteboardMime_MimeFor(void* ptr, char* flav);
void QMacPasteboardMime_DestroyQMacPasteboardMime(void* ptr);
void QMacPasteboardMime_DestroyQMacPasteboardMimeDefault(void* ptr);
void* QMacPasteboardMime___convertFromMime_atList(void* ptr, int i);
void QMacPasteboardMime___convertFromMime_setList(void* ptr, void* i);
void* QMacPasteboardMime___convertFromMime_newList(void* ptr);
void* QMacPasteboardMime___convertToMime_data_atList(void* ptr, int i);
void QMacPasteboardMime___convertToMime_data_setList(void* ptr, void* i);
void* QMacPasteboardMime___convertToMime_data_newList(void* ptr);
void* QMacToolBar_NewQMacToolBar(void* parent);
void* QMacToolBar_NewQMacToolBar2(char* identifier, void* parent);
void* QMacToolBar_AddAllowedItem(void* ptr, void* icon, char* text);
void* QMacToolBar_AddItem(void* ptr, void* icon, char* text);
void QMacToolBar_AddSeparator(void* ptr);
struct QtMacExtras_PackedList QMacToolBar_AllowedItems(void* ptr);
void QMacToolBar_AttachToWindow(void* ptr, void* window);
void QMacToolBar_DetachFromWindow(void* ptr);
struct QtMacExtras_PackedList QMacToolBar_Items(void* ptr);
void QMacToolBar_SetAllowedItems(void* ptr, void* allowedItems);
void QMacToolBar_SetItems(void* ptr, void* items);
void QMacToolBar_DestroyQMacToolBar(void* ptr);
void* QMacToolBar___allowedItems_atList(void* ptr, int i);
void QMacToolBar___allowedItems_setList(void* ptr, void* i);
void* QMacToolBar___allowedItems_newList(void* ptr);
void* QMacToolBar___items_atList(void* ptr, int i);
void QMacToolBar___items_setList(void* ptr, void* i);
void* QMacToolBar___items_newList(void* ptr);
void* QMacToolBar___setAllowedItems_allowedItems_atList(void* ptr, int i);
void QMacToolBar___setAllowedItems_allowedItems_setList(void* ptr, void* i);
void* QMacToolBar___setAllowedItems_allowedItems_newList(void* ptr);
void* QMacToolBar___setItems_items_atList(void* ptr, int i);
void QMacToolBar___setItems_items_setList(void* ptr, void* i);
void* QMacToolBar___setItems_items_newList(void* ptr);
void* QMacToolBar___children_atList(void* ptr, int i);
void QMacToolBar___children_setList(void* ptr, void* i);
void* QMacToolBar___children_newList(void* ptr);
void* QMacToolBar___dynamicPropertyNames_atList(void* ptr, int i);
void QMacToolBar___dynamicPropertyNames_setList(void* ptr, void* i);
void* QMacToolBar___dynamicPropertyNames_newList(void* ptr);
void* QMacToolBar___findChildren_atList2(void* ptr, int i);
void QMacToolBar___findChildren_setList2(void* ptr, void* i);
void* QMacToolBar___findChildren_newList2(void* ptr);
void* QMacToolBar___findChildren_atList3(void* ptr, int i);
void QMacToolBar___findChildren_setList3(void* ptr, void* i);
void* QMacToolBar___findChildren_newList3(void* ptr);
void* QMacToolBar___findChildren_atList(void* ptr, int i);
void QMacToolBar___findChildren_setList(void* ptr, void* i);
void* QMacToolBar___findChildren_newList(void* ptr);
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
char QMacToolBar_Event(void* ptr, void* e);
char QMacToolBar_EventDefault(void* ptr, void* e);
char QMacToolBar_EventFilter(void* ptr, void* watched, void* event);
char QMacToolBar_EventFilterDefault(void* ptr, void* watched, void* event);
void* QMacToolBar_MetaObject(void* ptr);
void* QMacToolBar_MetaObjectDefault(void* ptr);
void* QMacToolBarItem_NewQMacToolBarItem(void* parent);
void QMacToolBarItem_ConnectActivated(void* ptr);
void QMacToolBarItem_DisconnectActivated(void* ptr);
void QMacToolBarItem_Activated(void* ptr);
void QMacToolBarItem_DestroyQMacToolBarItem(void* ptr);
void QMacToolBarItem_DestroyQMacToolBarItemDefault(void* ptr);
void* QMacToolBarItem_Icon(void* ptr);
char QMacToolBarItem_Selectable(void* ptr);
void QMacToolBarItem_SetIcon(void* ptr, void* icon);
void QMacToolBarItem_SetSelectable(void* ptr, char selectable);
void QMacToolBarItem_SetStandardItem(void* ptr, long long standardItem);
void QMacToolBarItem_SetText(void* ptr, char* text);
long long QMacToolBarItem_StandardItem(void* ptr);
struct QtMacExtras_PackedString QMacToolBarItem_Text(void* ptr);
void* QMacToolBarItem___children_atList(void* ptr, int i);
void QMacToolBarItem___children_setList(void* ptr, void* i);
void* QMacToolBarItem___children_newList(void* ptr);
void* QMacToolBarItem___dynamicPropertyNames_atList(void* ptr, int i);
void QMacToolBarItem___dynamicPropertyNames_setList(void* ptr, void* i);
void* QMacToolBarItem___dynamicPropertyNames_newList(void* ptr);
void* QMacToolBarItem___findChildren_atList2(void* ptr, int i);
void QMacToolBarItem___findChildren_setList2(void* ptr, void* i);
void* QMacToolBarItem___findChildren_newList2(void* ptr);
void* QMacToolBarItem___findChildren_atList3(void* ptr, int i);
void QMacToolBarItem___findChildren_setList3(void* ptr, void* i);
void* QMacToolBarItem___findChildren_newList3(void* ptr);
void* QMacToolBarItem___findChildren_atList(void* ptr, int i);
void QMacToolBarItem___findChildren_setList(void* ptr, void* i);
void* QMacToolBarItem___findChildren_newList(void* ptr);
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
char QMacToolBarItem_Event(void* ptr, void* e);
char QMacToolBarItem_EventDefault(void* ptr, void* e);
char QMacToolBarItem_EventFilter(void* ptr, void* watched, void* event);
char QMacToolBarItem_EventFilterDefault(void* ptr, void* watched, void* event);
void* QMacToolBarItem_MetaObject(void* ptr);
void* QMacToolBarItem_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif