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
void QWebChannel_CustomEvent(void* ptr, void* event);
void QWebChannel_CustomEventDefault(void* ptr, void* event);
void QWebChannelAbstractTransport_ConnectMessageReceived(void* ptr);
void QWebChannelAbstractTransport_DisconnectMessageReceived(void* ptr);
void QWebChannelAbstractTransport_MessageReceived(void* ptr, void* message, void* transport);
void QWebChannelAbstractTransport_SendMessage(void* ptr, void* message);
void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(void* ptr);
void QWebChannelAbstractTransport_TimerEvent(void* ptr, void* event);
void QWebChannelAbstractTransport_TimerEventDefault(void* ptr, void* event);
void QWebChannelAbstractTransport_ChildEvent(void* ptr, void* event);
void QWebChannelAbstractTransport_ChildEventDefault(void* ptr, void* event);
void QWebChannelAbstractTransport_CustomEvent(void* ptr, void* event);
void QWebChannelAbstractTransport_CustomEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif