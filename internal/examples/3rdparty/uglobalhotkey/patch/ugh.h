#pragma once

#ifndef GO_UGH_H
#define GO_UGH_H

#ifdef __cplusplus
extern "C" {
#endif

struct UGlobalHotkeys_PackedString { char* data; long long len; };

void* UGlobalHotkeys_NewUGlobalHotkeys(void* parent);
void UGlobalHotkeys_RegisterHotkey(void* ptr, struct UGlobalHotkeys_PackedString keySeq, unsigned long id);
void UGlobalHotkeys_UnregisterHotkey(void* ptr, unsigned long id);
void UGlobalHotkeys_UnregisterAllHotkeys(void* ptr);
void UGlobalHotkeys_DestroyUGlobalHotkeys(void* ptr);
void UGlobalHotkeys_ConnectActivated(void* ptr);
void UGlobalHotkeys_DisconnectActivated(void* ptr);
void UGlobalHotkeys_Activated(void* ptr, unsigned long id);

#ifdef __cplusplus
}
#endif

#endif