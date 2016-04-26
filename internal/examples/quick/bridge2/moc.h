#pragma once

#ifndef GO_MAIN_H
#define GO_MAIN_H

#ifdef __cplusplus
extern "C" {
#endif

void QmlBridge_ConnectSendToQml(void* ptr);
void QmlBridge_DisconnectSendToQml(void* ptr);
void QmlBridge_SendToQml(void* ptr, char* data);
char* QmlBridge_SendToGo(void* ptr, char* data);
void* QmlBridge_NewQmlBridge(void* parent);
void QmlBridge_DestroyQmlBridge(void* ptr);
void QmlBridge_TimerEvent(void* ptr, void* event);
void QmlBridge_TimerEventDefault(void* ptr, void* event);
void QmlBridge_ChildEvent(void* ptr, void* event);
void QmlBridge_ChildEventDefault(void* ptr, void* event);
void QmlBridge_ConnectNotify(void* ptr, void* sign);
void QmlBridge_ConnectNotifyDefault(void* ptr, void* sign);
void QmlBridge_CustomEvent(void* ptr, void* event);
void QmlBridge_CustomEventDefault(void* ptr, void* event);
void QmlBridge_DisconnectNotify(void* ptr, void* sign);
void QmlBridge_DisconnectNotifyDefault(void* ptr, void* sign);

#ifdef __cplusplus
}
#endif

#endif