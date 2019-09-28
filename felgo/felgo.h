// +build !minimal

#pragma once

#ifndef GO_QTFELGO_H
#define GO_QTFELGO_H

#include <stdint.h>

#ifdef __cplusplus
int FelgoApplication_FelgoApplication_QRegisterMetaType();
int FelgoLiveClient_FelgoLiveClient_QRegisterMetaType();
extern "C" {
#endif

struct QtFelgo_PackedString { char* data; long long len; };
struct QtFelgo_PackedList { void* data; long long len; };
void FelgoApplication_ObjectCreatedHandler(void* ptr, void* object, void* url);
void FelgoApplication_ObjectCreatedHandlerDefault(void* ptr, void* object, void* url);
void* FelgoApplication_NewFelgoApplication(void* parent);
void FelgoApplication_AddCustomImportPaths(void* ptr);
void FelgoApplication_AddImportPath(void* ptr, struct QtFelgo_PackedString path);
struct QtFelgo_PackedString FelgoApplication_AdjustImportPathToCanonical(void* ptr, struct QtFelgo_PackedString vqs);
double FelgoApplication_ContentScale(void* ptr);
struct QtFelgo_PackedList FelgoApplication_ContentScaleThresholds(void* ptr);
struct QtFelgo_PackedString FelgoApplication_FileSelectorList(void* ptr);
void FelgoApplication_Initialize(void* ptr, void* engine);
void FelgoApplication_InitializeEngine(void* ptr, void* engine, char* uri);
struct QtFelgo_PackedString FelgoApplication_MainQmlFileName(void* ptr);
void* FelgoApplication_QmlEngine(void* ptr);
void FelgoApplication_RegisterTypes(void* ptr, char* uri);
double FelgoApplication_ScaleFactorForImages(void* ptr);
void FelgoApplication_SetContentScaleAndFileSelectors(void* ptr, double contentScale);
void FelgoApplication_SetContentScaleThresholds(void* ptr, void* contentScaleThresholds);
void FelgoApplication_SetDefaultContentScalingAndFileSelectors(void* ptr);
void FelgoApplication_SetFileSelectorList(void* ptr, struct QtFelgo_PackedString fileSelectorList);
void FelgoApplication_SetGameWindowIsItem(void* ptr);
void FelgoApplication_SetLicenseKey(void* ptr, struct QtFelgo_PackedString licenseKey);
void FelgoApplication_SetMainQmlFileName(void* ptr, struct QtFelgo_PackedString qmlFileName);
void FelgoApplication_SetPreservePlatformFonts(void* ptr, char preserve);
void* FelgoApplication_VplayFileSelector(void* ptr);
void FelgoApplication_DestroyFelgoApplication(void* ptr);
void* FelgoApplication___contentScaleThresholds_atList(void* ptr, struct QtFelgo_PackedString v, int i);
void FelgoApplication___contentScaleThresholds_setList(void* ptr, struct QtFelgo_PackedString key, void* i);
void* FelgoApplication___contentScaleThresholds_newList(void* ptr);
struct QtFelgo_PackedList FelgoApplication___contentScaleThresholds_keyList(void* ptr);
void* FelgoApplication___setContentScaleThresholds_contentScaleThresholds_atList(void* ptr, struct QtFelgo_PackedString v, int i);
void FelgoApplication___setContentScaleThresholds_contentScaleThresholds_setList(void* ptr, struct QtFelgo_PackedString key, void* i);
void* FelgoApplication___setContentScaleThresholds_contentScaleThresholds_newList(void* ptr);
struct QtFelgo_PackedList FelgoApplication___setContentScaleThresholds_contentScaleThresholds_keyList(void* ptr);
struct QtFelgo_PackedString FelgoApplication_____contentScaleThresholds_keyList_atList(void* ptr, int i);
void FelgoApplication_____contentScaleThresholds_keyList_setList(void* ptr, struct QtFelgo_PackedString i);
void* FelgoApplication_____contentScaleThresholds_keyList_newList(void* ptr);
struct QtFelgo_PackedString FelgoApplication_____setContentScaleThresholds_contentScaleThresholds_keyList_atList(void* ptr, int i);
void FelgoApplication_____setContentScaleThresholds_contentScaleThresholds_keyList_setList(void* ptr, struct QtFelgo_PackedString i);
void* FelgoApplication_____setContentScaleThresholds_contentScaleThresholds_keyList_newList(void* ptr);
void* FelgoApplication___children_atList(void* ptr, int i);
void FelgoApplication___children_setList(void* ptr, void* i);
void* FelgoApplication___children_newList(void* ptr);
void* FelgoApplication___dynamicPropertyNames_atList(void* ptr, int i);
void FelgoApplication___dynamicPropertyNames_setList(void* ptr, void* i);
void* FelgoApplication___dynamicPropertyNames_newList(void* ptr);
void* FelgoApplication___findChildren_atList(void* ptr, int i);
void FelgoApplication___findChildren_setList(void* ptr, void* i);
void* FelgoApplication___findChildren_newList(void* ptr);
void* FelgoApplication___findChildren_atList3(void* ptr, int i);
void FelgoApplication___findChildren_setList3(void* ptr, void* i);
void* FelgoApplication___findChildren_newList3(void* ptr);
void* FelgoApplication___qFindChildren_atList2(void* ptr, int i);
void FelgoApplication___qFindChildren_setList2(void* ptr, void* i);
void* FelgoApplication___qFindChildren_newList2(void* ptr);
void FelgoApplication_ChildEventDefault(void* ptr, void* event);
void FelgoApplication_ConnectNotifyDefault(void* ptr, void* sign);
void FelgoApplication_CustomEventDefault(void* ptr, void* event);
void FelgoApplication_DeleteLaterDefault(void* ptr);
void FelgoApplication_DisconnectNotifyDefault(void* ptr, void* sign);
char FelgoApplication_EventDefault(void* ptr, void* e);
char FelgoApplication_EventFilterDefault(void* ptr, void* watched, void* event);
void* FelgoApplication_MetaObjectDefault(void* ptr);
void FelgoApplication_TimerEventDefault(void* ptr, void* event);
void* FelgoLiveClient_NewFelgoLiveClient(void* engine, void* parent);
void FelgoLiveClient_AddShakeDetection(void* ptr);
void FelgoLiveClient_ClearProject(void* ptr, struct QtFelgo_PackedString projectName);
void FelgoLiveClient_ClearProjectDefault(void* ptr, struct QtFelgo_PackedString projectName);
void FelgoLiveClient_ClearProjects(void* ptr);
void FelgoLiveClient_ClearProjectsDefault(void* ptr);
struct QtFelgo_PackedString FelgoLiveClient_ClientName(void* ptr);
void FelgoLiveClient_ConnectClientNameChanged(void* ptr);
void FelgoLiveClient_DisconnectClientNameChanged(void* ptr);
void FelgoLiveClient_ClientNameChanged(void* ptr, struct QtFelgo_PackedString clientName);
void FelgoLiveClient_ConnectAsLocalDesktopClient(void* ptr, struct QtFelgo_PackedString serverUrl);
void FelgoLiveClient_ConnectAsLocalDesktopClientDefault(void* ptr, struct QtFelgo_PackedString serverUrl);
void FelgoLiveClient_ConnectToLocalServer(void* ptr, struct QtFelgo_PackedString serverUrl, struct QtFelgo_PackedString userId, struct QtFelgo_PackedString deviceId);
void FelgoLiveClient_ConnectToLocalServerDefault(void* ptr, struct QtFelgo_PackedString serverUrl, struct QtFelgo_PackedString userId, struct QtFelgo_PackedString deviceId);
void FelgoLiveClient_ConnectToWebServer(void* ptr, struct QtFelgo_PackedString userId, struct QtFelgo_PackedString deviceId);
void FelgoLiveClient_LoadProject(void* ptr, struct QtFelgo_PackedString projectName, struct QtFelgo_PackedString projectMainFile);
void FelgoLiveClient_DisableDisplaySleep(void* ptr);
void FelgoLiveClient_DisconnectWebReceiver(void* ptr);
void FelgoLiveClient_DisconnectWebReceiverDefault(void* ptr);
void FelgoLiveClient_ModifyLoadedDocument(void* ptr);
void FelgoLiveClient_ModifyLoadedDocumentDefault(void* ptr);
void FelgoLiveClient_OnDocumentUpdated(void* ptr, struct QtFelgo_PackedString document, struct QtFelgo_PackedString content);
void FelgoLiveClient_OnDocumentUpdatedDefault(void* ptr, struct QtFelgo_PackedString document, struct QtFelgo_PackedString content);
void FelgoLiveClient_OnProjectChanged(void* ptr, struct QtFelgo_PackedString projectName, struct QtFelgo_PackedString projectMainFile);
void FelgoLiveClient_OnProjectChangedDefault(void* ptr, struct QtFelgo_PackedString projectName, struct QtFelgo_PackedString projectMainFile);
void FelgoLiveClient_ConnectPendingProject(void* ptr);
void FelgoLiveClient_DisconnectPendingProject(void* ptr);
void FelgoLiveClient_PendingProject(void* ptr, struct QtFelgo_PackedString projectName, struct QtFelgo_PackedString projectMainFile);
void FelgoLiveClient_PrepareLoadingDocument(void* ptr);
void FelgoLiveClient_PrepareLoadingDocumentDefault(void* ptr);
void FelgoLiveClient_ConnectWebReceiverConnectionRefused(void* ptr);
void FelgoLiveClient_DisconnectWebReceiverConnectionRefused(void* ptr);
void FelgoLiveClient_WebReceiverConnectionRefused(void* ptr, struct QtFelgo_PackedString reason);
void FelgoLiveClient_ConnectReceivedMatchId(void* ptr);
void FelgoLiveClient_DisconnectReceivedMatchId(void* ptr);
void FelgoLiveClient_ReceivedMatchId(void* ptr, struct QtFelgo_PackedString matchId);
void FelgoLiveClient_SetClientName(void* ptr, struct QtFelgo_PackedString clientName);
void FelgoLiveClient_SetQmlEngine(void* ptr, void* engine);
void FelgoLiveClient_SetWelcomeQmlFile(void* ptr, struct QtFelgo_PackedString file);
void FelgoLiveClient_ShakeDetected(void* ptr);
void FelgoLiveClient_ShakeDetectedDefault(void* ptr);
void FelgoLiveClient_ConnectWebReceiverConnected(void* ptr);
void FelgoLiveClient_DisconnectWebReceiverConnected(void* ptr);
void FelgoLiveClient_WebReceiverConnected(void* ptr);
void* FelgoLiveClient___children_atList(void* ptr, int i);
void FelgoLiveClient___children_setList(void* ptr, void* i);
void* FelgoLiveClient___children_newList(void* ptr);
void* FelgoLiveClient___dynamicPropertyNames_atList(void* ptr, int i);
void FelgoLiveClient___dynamicPropertyNames_setList(void* ptr, void* i);
void* FelgoLiveClient___dynamicPropertyNames_newList(void* ptr);
void* FelgoLiveClient___findChildren_atList(void* ptr, int i);
void FelgoLiveClient___findChildren_setList(void* ptr, void* i);
void* FelgoLiveClient___findChildren_newList(void* ptr);
void* FelgoLiveClient___findChildren_atList3(void* ptr, int i);
void FelgoLiveClient___findChildren_setList3(void* ptr, void* i);
void* FelgoLiveClient___findChildren_newList3(void* ptr);
void* FelgoLiveClient___qFindChildren_atList2(void* ptr, int i);
void FelgoLiveClient___qFindChildren_setList2(void* ptr, void* i);
void* FelgoLiveClient___qFindChildren_newList2(void* ptr);
void FelgoLiveClient_ChildEventDefault(void* ptr, void* event);
void FelgoLiveClient_ConnectNotifyDefault(void* ptr, void* sign);
void FelgoLiveClient_CustomEventDefault(void* ptr, void* event);
void FelgoLiveClient_DeleteLaterDefault(void* ptr);
void FelgoLiveClient_DisconnectNotifyDefault(void* ptr, void* sign);
char FelgoLiveClient_EventDefault(void* ptr, void* e);
char FelgoLiveClient_EventFilterDefault(void* ptr, void* watched, void* event);
void* FelgoLiveClient_MetaObjectDefault(void* ptr);
void FelgoLiveClient_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif