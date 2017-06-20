// +build !minimal

#pragma once

#ifndef GO_QTPURCHASING_H
#define GO_QTPURCHASING_H

#include <stdint.h>

#ifdef __cplusplus
int QInAppProduct_QInAppProduct_QRegisterMetaType();
int QInAppStore_QInAppStore_QRegisterMetaType();
int QInAppTransaction_QInAppTransaction_QRegisterMetaType();
extern "C" {
#endif

struct QtPurchasing_PackedString { char* data; long long len; };
struct QtPurchasing_PackedList { void* data; long long len; };
void QInAppProduct_Purchase(void* ptr);
long long QInAppProduct_ProductType(void* ptr);
struct QtPurchasing_PackedString QInAppProduct_Identifier(void* ptr);
struct QtPurchasing_PackedString QInAppProduct_Description(void* ptr);
struct QtPurchasing_PackedString QInAppProduct_Price(void* ptr);
struct QtPurchasing_PackedString QInAppProduct_Title(void* ptr);
void* QInAppProduct___dynamicPropertyNames_atList(void* ptr, int i);
void QInAppProduct___dynamicPropertyNames_setList(void* ptr, void* i);
void* QInAppProduct___dynamicPropertyNames_newList(void* ptr);
void* QInAppProduct___findChildren_atList2(void* ptr, int i);
void QInAppProduct___findChildren_setList2(void* ptr, void* i);
void* QInAppProduct___findChildren_newList2(void* ptr);
void* QInAppProduct___findChildren_atList3(void* ptr, int i);
void QInAppProduct___findChildren_setList3(void* ptr, void* i);
void* QInAppProduct___findChildren_newList3(void* ptr);
void* QInAppProduct___findChildren_atList(void* ptr, int i);
void QInAppProduct___findChildren_setList(void* ptr, void* i);
void* QInAppProduct___findChildren_newList(void* ptr);
void* QInAppProduct___children_atList(void* ptr, int i);
void QInAppProduct___children_setList(void* ptr, void* i);
void* QInAppProduct___children_newList(void* ptr);
char QInAppProduct_EventDefault(void* ptr, void* e);
char QInAppProduct_EventFilterDefault(void* ptr, void* watched, void* event);
void QInAppProduct_ChildEventDefault(void* ptr, void* event);
void QInAppProduct_ConnectNotifyDefault(void* ptr, void* sign);
void QInAppProduct_CustomEventDefault(void* ptr, void* event);
void QInAppProduct_DeleteLaterDefault(void* ptr);
void QInAppProduct_DisconnectNotifyDefault(void* ptr, void* sign);
void QInAppProduct_TimerEventDefault(void* ptr, void* event);
void* QInAppProduct_MetaObjectDefault(void* ptr);
void* QInAppStore_NewQInAppStore(void* parent);
void QInAppStore_ConnectProductRegistered(void* ptr);
void QInAppStore_DisconnectProductRegistered(void* ptr);
void QInAppStore_ProductRegistered(void* ptr, void* product);
void QInAppStore_ConnectProductUnknown(void* ptr);
void QInAppStore_DisconnectProductUnknown(void* ptr);
void QInAppStore_ProductUnknown(void* ptr, long long productType, struct QtPurchasing_PackedString identifier);
void QInAppStore_RegisterProduct(void* ptr, long long productType, struct QtPurchasing_PackedString identifier);
void QInAppStore_RestorePurchases(void* ptr);
void QInAppStore_SetPlatformProperty(void* ptr, struct QtPurchasing_PackedString propertyName, struct QtPurchasing_PackedString value);
void QInAppStore_ConnectTransactionReady(void* ptr);
void QInAppStore_DisconnectTransactionReady(void* ptr);
void QInAppStore_TransactionReady(void* ptr, void* transaction);
void QInAppStore_DestroyQInAppStore(void* ptr);
void* QInAppStore_RegisteredProduct(void* ptr, struct QtPurchasing_PackedString identifier);
void* QInAppStore___dynamicPropertyNames_atList(void* ptr, int i);
void QInAppStore___dynamicPropertyNames_setList(void* ptr, void* i);
void* QInAppStore___dynamicPropertyNames_newList(void* ptr);
void* QInAppStore___findChildren_atList2(void* ptr, int i);
void QInAppStore___findChildren_setList2(void* ptr, void* i);
void* QInAppStore___findChildren_newList2(void* ptr);
void* QInAppStore___findChildren_atList3(void* ptr, int i);
void QInAppStore___findChildren_setList3(void* ptr, void* i);
void* QInAppStore___findChildren_newList3(void* ptr);
void* QInAppStore___findChildren_atList(void* ptr, int i);
void QInAppStore___findChildren_setList(void* ptr, void* i);
void* QInAppStore___findChildren_newList(void* ptr);
void* QInAppStore___children_atList(void* ptr, int i);
void QInAppStore___children_setList(void* ptr, void* i);
void* QInAppStore___children_newList(void* ptr);
char QInAppStore_EventDefault(void* ptr, void* e);
char QInAppStore_EventFilterDefault(void* ptr, void* watched, void* event);
void QInAppStore_ChildEventDefault(void* ptr, void* event);
void QInAppStore_ConnectNotifyDefault(void* ptr, void* sign);
void QInAppStore_CustomEventDefault(void* ptr, void* event);
void QInAppStore_DeleteLaterDefault(void* ptr);
void QInAppStore_DisconnectNotifyDefault(void* ptr, void* sign);
void QInAppStore_TimerEventDefault(void* ptr, void* event);
void* QInAppStore_MetaObjectDefault(void* ptr);
struct QtPurchasing_PackedString QInAppTransaction_ErrorString(void* ptr);
struct QtPurchasing_PackedString QInAppTransaction_ErrorStringDefault(void* ptr);
struct QtPurchasing_PackedString QInAppTransaction_OrderId(void* ptr);
struct QtPurchasing_PackedString QInAppTransaction_OrderIdDefault(void* ptr);
void QInAppTransaction_Finalize(void* ptr);
long long QInAppTransaction_FailureReason(void* ptr);
long long QInAppTransaction_FailureReasonDefault(void* ptr);
void* QInAppTransaction_Timestamp(void* ptr);
void* QInAppTransaction_TimestampDefault(void* ptr);
void* QInAppTransaction_Product(void* ptr);
struct QtPurchasing_PackedString QInAppTransaction_PlatformProperty(void* ptr, struct QtPurchasing_PackedString propertyName);
struct QtPurchasing_PackedString QInAppTransaction_PlatformPropertyDefault(void* ptr, struct QtPurchasing_PackedString propertyName);
long long QInAppTransaction_Status(void* ptr);
void* QInAppTransaction___dynamicPropertyNames_atList(void* ptr, int i);
void QInAppTransaction___dynamicPropertyNames_setList(void* ptr, void* i);
void* QInAppTransaction___dynamicPropertyNames_newList(void* ptr);
void* QInAppTransaction___findChildren_atList2(void* ptr, int i);
void QInAppTransaction___findChildren_setList2(void* ptr, void* i);
void* QInAppTransaction___findChildren_newList2(void* ptr);
void* QInAppTransaction___findChildren_atList3(void* ptr, int i);
void QInAppTransaction___findChildren_setList3(void* ptr, void* i);
void* QInAppTransaction___findChildren_newList3(void* ptr);
void* QInAppTransaction___findChildren_atList(void* ptr, int i);
void QInAppTransaction___findChildren_setList(void* ptr, void* i);
void* QInAppTransaction___findChildren_newList(void* ptr);
void* QInAppTransaction___children_atList(void* ptr, int i);
void QInAppTransaction___children_setList(void* ptr, void* i);
void* QInAppTransaction___children_newList(void* ptr);
char QInAppTransaction_EventDefault(void* ptr, void* e);
char QInAppTransaction_EventFilterDefault(void* ptr, void* watched, void* event);
void QInAppTransaction_ChildEventDefault(void* ptr, void* event);
void QInAppTransaction_ConnectNotifyDefault(void* ptr, void* sign);
void QInAppTransaction_CustomEventDefault(void* ptr, void* event);
void QInAppTransaction_DeleteLaterDefault(void* ptr);
void QInAppTransaction_DisconnectNotifyDefault(void* ptr, void* sign);
void QInAppTransaction_TimerEventDefault(void* ptr, void* event);
void* QInAppTransaction_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif