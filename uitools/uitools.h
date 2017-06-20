// +build !minimal

#pragma once

#ifndef GO_QTUITOOLS_H
#define GO_QTUITOOLS_H

#include <stdint.h>

#ifdef __cplusplus
int QUiLoader_QUiLoader_QRegisterMetaType();
extern "C" {
#endif

struct QtUiTools_PackedString { char* data; long long len; };
struct QtUiTools_PackedList { void* data; long long len; };
void* QUiLoader_CreateAction(void* ptr, void* parent, struct QtUiTools_PackedString name);
void* QUiLoader_CreateActionDefault(void* ptr, void* parent, struct QtUiTools_PackedString name);
void* QUiLoader_CreateActionGroup(void* ptr, void* parent, struct QtUiTools_PackedString name);
void* QUiLoader_CreateActionGroupDefault(void* ptr, void* parent, struct QtUiTools_PackedString name);
void* QUiLoader_CreateLayout(void* ptr, struct QtUiTools_PackedString className, void* parent, struct QtUiTools_PackedString name);
void* QUiLoader_CreateLayoutDefault(void* ptr, struct QtUiTools_PackedString className, void* parent, struct QtUiTools_PackedString name);
void* QUiLoader_NewQUiLoader(void* parent);
void* QUiLoader_CreateWidget(void* ptr, struct QtUiTools_PackedString className, void* parent, struct QtUiTools_PackedString name);
void* QUiLoader_CreateWidgetDefault(void* ptr, struct QtUiTools_PackedString className, void* parent, struct QtUiTools_PackedString name);
void* QUiLoader_Load(void* ptr, void* device, void* parentWidget);
void QUiLoader_AddPluginPath(void* ptr, struct QtUiTools_PackedString path);
void QUiLoader_ClearPluginPaths(void* ptr);
void QUiLoader_SetLanguageChangeEnabled(void* ptr, char enabled);
void QUiLoader_SetWorkingDirectory(void* ptr, void* dir);
void QUiLoader_DestroyQUiLoader(void* ptr);
void QUiLoader_DestroyQUiLoaderDefault(void* ptr);
void* QUiLoader_WorkingDirectory(void* ptr);
struct QtUiTools_PackedString QUiLoader_ErrorString(void* ptr);
struct QtUiTools_PackedString QUiLoader_AvailableLayouts(void* ptr);
struct QtUiTools_PackedString QUiLoader_AvailableWidgets(void* ptr);
struct QtUiTools_PackedString QUiLoader_PluginPaths(void* ptr);
char QUiLoader_IsLanguageChangeEnabled(void* ptr);
void* QUiLoader___dynamicPropertyNames_atList(void* ptr, int i);
void QUiLoader___dynamicPropertyNames_setList(void* ptr, void* i);
void* QUiLoader___dynamicPropertyNames_newList(void* ptr);
void* QUiLoader___findChildren_atList2(void* ptr, int i);
void QUiLoader___findChildren_setList2(void* ptr, void* i);
void* QUiLoader___findChildren_newList2(void* ptr);
void* QUiLoader___findChildren_atList3(void* ptr, int i);
void QUiLoader___findChildren_setList3(void* ptr, void* i);
void* QUiLoader___findChildren_newList3(void* ptr);
void* QUiLoader___findChildren_atList(void* ptr, int i);
void QUiLoader___findChildren_setList(void* ptr, void* i);
void* QUiLoader___findChildren_newList(void* ptr);
void* QUiLoader___children_atList(void* ptr, int i);
void QUiLoader___children_setList(void* ptr, void* i);
void* QUiLoader___children_newList(void* ptr);
char QUiLoader_EventDefault(void* ptr, void* e);
char QUiLoader_EventFilterDefault(void* ptr, void* watched, void* event);
void QUiLoader_ChildEventDefault(void* ptr, void* event);
void QUiLoader_ConnectNotifyDefault(void* ptr, void* sign);
void QUiLoader_CustomEventDefault(void* ptr, void* event);
void QUiLoader_DeleteLaterDefault(void* ptr);
void QUiLoader_DisconnectNotifyDefault(void* ptr, void* sign);
void QUiLoader_TimerEventDefault(void* ptr, void* event);
void* QUiLoader_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif