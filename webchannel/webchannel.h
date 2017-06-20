// +build !minimal

#pragma once

#ifndef GO_QTWEBCHANNEL_H
#define GO_QTWEBCHANNEL_H

#include <stdint.h>

#ifdef __cplusplus
int QWebChannel_QWebChannel_QRegisterMetaType();
int QWebChannelAbstractTransport_QWebChannelAbstractTransport_QRegisterMetaType();
extern "C" {
#endif

struct QtWebChannel_PackedString { char* data; long long len; };
struct QtWebChannel_PackedList { void* data; long long len; };
void* QWebChannel_NewQWebChannel(void* parent);
void QWebChannel_ConnectBlockUpdatesChanged(void* ptr);
void QWebChannel_DisconnectBlockUpdatesChanged(void* ptr);
void QWebChannel_BlockUpdatesChanged(void* ptr, char block);
void QWebChannel_ConnectTo(void* ptr, void* transport);
void QWebChannel_ConnectToDefault(void* ptr, void* transport);
void QWebChannel_DeregisterObject(void* ptr, void* object);
void QWebChannel_DisconnectFrom(void* ptr, void* transport);
void QWebChannel_DisconnectFromDefault(void* ptr, void* transport);
void QWebChannel_RegisterObject(void* ptr, struct QtWebChannel_PackedString id, void* object);
void QWebChannel_RegisterObjects(void* ptr, void* objects);
void QWebChannel_SetBlockUpdates(void* ptr, char block);
void QWebChannel_DestroyQWebChannel(void* ptr);
struct QtWebChannel_PackedList QWebChannel_RegisteredObjects(void* ptr);
char QWebChannel_BlockUpdates(void* ptr);
void* QWebChannel___registerObjects_objects_atList(void* ptr, struct QtWebChannel_PackedString i);
void QWebChannel___registerObjects_objects_setList(void* ptr, struct QtWebChannel_PackedString key, void* i);
void* QWebChannel___registerObjects_objects_newList(void* ptr);
struct QtWebChannel_PackedList QWebChannel___registerObjects_keyList(void* ptr);
void* QWebChannel___registeredObjects_atList(void* ptr, struct QtWebChannel_PackedString i);
void QWebChannel___registeredObjects_setList(void* ptr, struct QtWebChannel_PackedString key, void* i);
void* QWebChannel___registeredObjects_newList(void* ptr);
struct QtWebChannel_PackedList QWebChannel___registeredObjects_keyList(void* ptr);
struct QtWebChannel_PackedString QWebChannel_____registerObjects_keyList_atList(void* ptr, int i);
void QWebChannel_____registerObjects_keyList_setList(void* ptr, struct QtWebChannel_PackedString i);
void* QWebChannel_____registerObjects_keyList_newList(void* ptr);
struct QtWebChannel_PackedString QWebChannel_____registeredObjects_keyList_atList(void* ptr, int i);
void QWebChannel_____registeredObjects_keyList_setList(void* ptr, struct QtWebChannel_PackedString i);
void* QWebChannel_____registeredObjects_keyList_newList(void* ptr);
void* QWebChannel___dynamicPropertyNames_atList(void* ptr, int i);
void QWebChannel___dynamicPropertyNames_setList(void* ptr, void* i);
void* QWebChannel___dynamicPropertyNames_newList(void* ptr);
void* QWebChannel___findChildren_atList2(void* ptr, int i);
void QWebChannel___findChildren_setList2(void* ptr, void* i);
void* QWebChannel___findChildren_newList2(void* ptr);
void* QWebChannel___findChildren_atList3(void* ptr, int i);
void QWebChannel___findChildren_setList3(void* ptr, void* i);
void* QWebChannel___findChildren_newList3(void* ptr);
void* QWebChannel___findChildren_atList(void* ptr, int i);
void QWebChannel___findChildren_setList(void* ptr, void* i);
void* QWebChannel___findChildren_newList(void* ptr);
void* QWebChannel___children_atList(void* ptr, int i);
void QWebChannel___children_setList(void* ptr, void* i);
void* QWebChannel___children_newList(void* ptr);
char QWebChannel_EventDefault(void* ptr, void* e);
char QWebChannel_EventFilterDefault(void* ptr, void* watched, void* event);
void QWebChannel_ChildEventDefault(void* ptr, void* event);
void QWebChannel_ConnectNotifyDefault(void* ptr, void* sign);
void QWebChannel_CustomEventDefault(void* ptr, void* event);
void QWebChannel_DeleteLaterDefault(void* ptr);
void QWebChannel_DisconnectNotifyDefault(void* ptr, void* sign);
void QWebChannel_TimerEventDefault(void* ptr, void* event);
void* QWebChannel_MetaObjectDefault(void* ptr);
void* QWebChannelAbstractTransport_NewQWebChannelAbstractTransport(void* parent);
void QWebChannelAbstractTransport_ConnectMessageReceived(void* ptr);
void QWebChannelAbstractTransport_DisconnectMessageReceived(void* ptr);
void QWebChannelAbstractTransport_MessageReceived(void* ptr, void* message, void* transport);
void QWebChannelAbstractTransport_SendMessage(void* ptr, void* message);
void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(void* ptr);
void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransportDefault(void* ptr);
void* QWebChannelAbstractTransport___dynamicPropertyNames_atList(void* ptr, int i);
void QWebChannelAbstractTransport___dynamicPropertyNames_setList(void* ptr, void* i);
void* QWebChannelAbstractTransport___dynamicPropertyNames_newList(void* ptr);
void* QWebChannelAbstractTransport___findChildren_atList2(void* ptr, int i);
void QWebChannelAbstractTransport___findChildren_setList2(void* ptr, void* i);
void* QWebChannelAbstractTransport___findChildren_newList2(void* ptr);
void* QWebChannelAbstractTransport___findChildren_atList3(void* ptr, int i);
void QWebChannelAbstractTransport___findChildren_setList3(void* ptr, void* i);
void* QWebChannelAbstractTransport___findChildren_newList3(void* ptr);
void* QWebChannelAbstractTransport___findChildren_atList(void* ptr, int i);
void QWebChannelAbstractTransport___findChildren_setList(void* ptr, void* i);
void* QWebChannelAbstractTransport___findChildren_newList(void* ptr);
void* QWebChannelAbstractTransport___children_atList(void* ptr, int i);
void QWebChannelAbstractTransport___children_setList(void* ptr, void* i);
void* QWebChannelAbstractTransport___children_newList(void* ptr);
char QWebChannelAbstractTransport_EventDefault(void* ptr, void* e);
char QWebChannelAbstractTransport_EventFilterDefault(void* ptr, void* watched, void* event);
void QWebChannelAbstractTransport_ChildEventDefault(void* ptr, void* event);
void QWebChannelAbstractTransport_ConnectNotifyDefault(void* ptr, void* sign);
void QWebChannelAbstractTransport_CustomEventDefault(void* ptr, void* event);
void QWebChannelAbstractTransport_DeleteLaterDefault(void* ptr);
void QWebChannelAbstractTransport_DisconnectNotifyDefault(void* ptr, void* sign);
void QWebChannelAbstractTransport_TimerEventDefault(void* ptr, void* event);
void* QWebChannelAbstractTransport_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif