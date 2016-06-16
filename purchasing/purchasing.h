// +build !minimal

#pragma once

#ifndef GO_QTPURCHASING_H
#define GO_QTPURCHASING_H

#ifdef __cplusplus
extern "C" {
#endif

char* QInAppProduct_Description(void* ptr);
char* QInAppProduct_Identifier(void* ptr);
char* QInAppProduct_Price(void* ptr);
int QInAppProduct_ProductType(void* ptr);
char* QInAppProduct_Title(void* ptr);
void QInAppProduct_Purchase(void* ptr);
void QInAppProduct_TimerEvent(void* ptr, void* event);
void QInAppProduct_TimerEventDefault(void* ptr, void* event);
void QInAppProduct_ChildEvent(void* ptr, void* event);
void QInAppProduct_ChildEventDefault(void* ptr, void* event);
void QInAppProduct_ConnectNotify(void* ptr, void* sign);
void QInAppProduct_ConnectNotifyDefault(void* ptr, void* sign);
void QInAppProduct_CustomEvent(void* ptr, void* event);
void QInAppProduct_CustomEventDefault(void* ptr, void* event);
void QInAppProduct_DeleteLater(void* ptr);
void QInAppProduct_DeleteLaterDefault(void* ptr);
void QInAppProduct_DisconnectNotify(void* ptr, void* sign);
void QInAppProduct_DisconnectNotifyDefault(void* ptr, void* sign);
int QInAppProduct_Event(void* ptr, void* e);
int QInAppProduct_EventDefault(void* ptr, void* e);
int QInAppProduct_EventFilter(void* ptr, void* watched, void* event);
int QInAppProduct_EventFilterDefault(void* ptr, void* watched, void* event);
void* QInAppProduct_MetaObject(void* ptr);
void* QInAppProduct_MetaObjectDefault(void* ptr);
void* QInAppStore_NewQInAppStore(void* parent);
void QInAppStore_ConnectProductRegistered(void* ptr);
void QInAppStore_DisconnectProductRegistered(void* ptr);
void QInAppStore_ProductRegistered(void* ptr, void* product);
void QInAppStore_ConnectProductUnknown(void* ptr);
void QInAppStore_DisconnectProductUnknown(void* ptr);
void QInAppStore_ProductUnknown(void* ptr, int productType, char* identifier);
void QInAppStore_RegisterProduct(void* ptr, int productType, char* identifier);
void* QInAppStore_RegisteredProduct(void* ptr, char* identifier);
void QInAppStore_RestorePurchases(void* ptr);
void QInAppStore_SetPlatformProperty(void* ptr, char* propertyName, char* value);
void QInAppStore_ConnectTransactionReady(void* ptr);
void QInAppStore_DisconnectTransactionReady(void* ptr);
void QInAppStore_TransactionReady(void* ptr, void* transaction);
void QInAppStore_DestroyQInAppStore(void* ptr);
void QInAppStore_TimerEvent(void* ptr, void* event);
void QInAppStore_TimerEventDefault(void* ptr, void* event);
void QInAppStore_ChildEvent(void* ptr, void* event);
void QInAppStore_ChildEventDefault(void* ptr, void* event);
void QInAppStore_ConnectNotify(void* ptr, void* sign);
void QInAppStore_ConnectNotifyDefault(void* ptr, void* sign);
void QInAppStore_CustomEvent(void* ptr, void* event);
void QInAppStore_CustomEventDefault(void* ptr, void* event);
void QInAppStore_DeleteLater(void* ptr);
void QInAppStore_DeleteLaterDefault(void* ptr);
void QInAppStore_DisconnectNotify(void* ptr, void* sign);
void QInAppStore_DisconnectNotifyDefault(void* ptr, void* sign);
int QInAppStore_Event(void* ptr, void* e);
int QInAppStore_EventDefault(void* ptr, void* e);
int QInAppStore_EventFilter(void* ptr, void* watched, void* event);
int QInAppStore_EventFilterDefault(void* ptr, void* watched, void* event);
void* QInAppStore_MetaObject(void* ptr);
void* QInAppStore_MetaObjectDefault(void* ptr);
char* QInAppTransaction_ErrorString(void* ptr);
char* QInAppTransaction_ErrorStringDefault(void* ptr);
int QInAppTransaction_FailureReason(void* ptr);
int QInAppTransaction_FailureReasonDefault(void* ptr);
char* QInAppTransaction_OrderId(void* ptr);
char* QInAppTransaction_OrderIdDefault(void* ptr);
void* QInAppTransaction_Product(void* ptr);
int QInAppTransaction_Status(void* ptr);
void* QInAppTransaction_Timestamp(void* ptr);
void* QInAppTransaction_TimestampDefault(void* ptr);
void QInAppTransaction_Finalize(void* ptr);
char* QInAppTransaction_PlatformProperty(void* ptr, char* propertyName);
char* QInAppTransaction_PlatformPropertyDefault(void* ptr, char* propertyName);
void QInAppTransaction_TimerEvent(void* ptr, void* event);
void QInAppTransaction_TimerEventDefault(void* ptr, void* event);
void QInAppTransaction_ChildEvent(void* ptr, void* event);
void QInAppTransaction_ChildEventDefault(void* ptr, void* event);
void QInAppTransaction_ConnectNotify(void* ptr, void* sign);
void QInAppTransaction_ConnectNotifyDefault(void* ptr, void* sign);
void QInAppTransaction_CustomEvent(void* ptr, void* event);
void QInAppTransaction_CustomEventDefault(void* ptr, void* event);
void QInAppTransaction_DeleteLater(void* ptr);
void QInAppTransaction_DeleteLaterDefault(void* ptr);
void QInAppTransaction_DisconnectNotify(void* ptr, void* sign);
void QInAppTransaction_DisconnectNotifyDefault(void* ptr, void* sign);
int QInAppTransaction_Event(void* ptr, void* e);
int QInAppTransaction_EventDefault(void* ptr, void* e);
int QInAppTransaction_EventFilter(void* ptr, void* watched, void* event);
int QInAppTransaction_EventFilterDefault(void* ptr, void* watched, void* event);
void* QInAppTransaction_MetaObject(void* ptr);
void* QInAppTransaction_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif