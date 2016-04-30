#pragma once

#ifndef GO_QTWEBCHANNEL_H
#define GO_QTWEBCHANNEL_H

#ifdef __cplusplus
extern "C" {
#endif

int QWebChannel_BlockUpdates(void* ptr);
void QWebChannel_SetBlockUpdates(void* ptr, int block);
void* QWebChannel_NewQWebChannel(void* parent);
void QWebChannel_ConnectBlockUpdatesChanged(void* ptr);
void QWebChannel_DisconnectBlockUpdatesChanged(void* ptr);
void QWebChannel_BlockUpdatesChanged(void* ptr, int block);
void QWebChannel_ConnectTo(void* ptr, void* transport);
void QWebChannel_DeregisterObject(void* ptr, void* object);
void QWebChannel_DisconnectFrom(void* ptr, void* transport);
void QWebChannel_RegisterObject(void* ptr, char* id, void* object);
void QWebChannel_DestroyQWebChannel(void* ptr);
void QWebChannel_TimerEvent(void* ptr, void* event);
void QWebChannel_TimerEventDefault(void* ptr, void* event);
void QWebChannel_ChildEvent(void* ptr, void* event);
void QWebChannel_ChildEventDefault(void* ptr, void* event);
void QWebChannel_ConnectNotify(void* ptr, void* sign);
void QWebChannel_ConnectNotifyDefault(void* ptr, void* sign);
void QWebChannel_CustomEvent(void* ptr, void* event);
void QWebChannel_CustomEventDefault(void* ptr, void* event);
void QWebChannel_DeleteLater(void* ptr);
void QWebChannel_DeleteLaterDefault(void* ptr);
void QWebChannel_DisconnectNotify(void* ptr, void* sign);
void QWebChannel_DisconnectNotifyDefault(void* ptr, void* sign);
int QWebChannel_Event(void* ptr, void* e);
int QWebChannel_EventDefault(void* ptr, void* e);
int QWebChannel_EventFilter(void* ptr, void* watched, void* event);
int QWebChannel_EventFilterDefault(void* ptr, void* watched, void* event);
void* QWebChannel_MetaObject(void* ptr);
void* QWebChannel_MetaObjectDefault(void* ptr);
void* QWebChannelAbstractTransport_NewQWebChannelAbstractTransport(void* parent);
void QWebChannelAbstractTransport_ConnectMessageReceived(void* ptr);
void QWebChannelAbstractTransport_DisconnectMessageReceived(void* ptr);
void QWebChannelAbstractTransport_MessageReceived(void* ptr, void* message, void* transport);
void QWebChannelAbstractTransport_SendMessage(void* ptr, void* message);
void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(void* ptr);
void QWebChannelAbstractTransport_TimerEvent(void* ptr, void* event);
void QWebChannelAbstractTransport_TimerEventDefault(void* ptr, void* event);
void QWebChannelAbstractTransport_ChildEvent(void* ptr, void* event);
void QWebChannelAbstractTransport_ChildEventDefault(void* ptr, void* event);
void QWebChannelAbstractTransport_ConnectNotify(void* ptr, void* sign);
void QWebChannelAbstractTransport_ConnectNotifyDefault(void* ptr, void* sign);
void QWebChannelAbstractTransport_CustomEvent(void* ptr, void* event);
void QWebChannelAbstractTransport_CustomEventDefault(void* ptr, void* event);
void QWebChannelAbstractTransport_DeleteLater(void* ptr);
void QWebChannelAbstractTransport_DeleteLaterDefault(void* ptr);
void QWebChannelAbstractTransport_DisconnectNotify(void* ptr, void* sign);
void QWebChannelAbstractTransport_DisconnectNotifyDefault(void* ptr, void* sign);
int QWebChannelAbstractTransport_Event(void* ptr, void* e);
int QWebChannelAbstractTransport_EventDefault(void* ptr, void* e);
int QWebChannelAbstractTransport_EventFilter(void* ptr, void* watched, void* event);
int QWebChannelAbstractTransport_EventFilterDefault(void* ptr, void* watched, void* event);
void* QWebChannelAbstractTransport_MetaObject(void* ptr);
void* QWebChannelAbstractTransport_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif