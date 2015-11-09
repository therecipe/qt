#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkAddressEntry_NewQNetworkAddressEntry();
void* QNetworkAddressEntry_NewQNetworkAddressEntry2(void* other);
int QNetworkAddressEntry_PrefixLength(void* ptr);
void QNetworkAddressEntry_SetBroadcast(void* ptr, void* newBroadcast);
void QNetworkAddressEntry_SetIp(void* ptr, void* newIp);
void QNetworkAddressEntry_SetNetmask(void* ptr, void* newNetmask);
void QNetworkAddressEntry_SetPrefixLength(void* ptr, int length);
void QNetworkAddressEntry_Swap(void* ptr, void* other);
void QNetworkAddressEntry_DestroyQNetworkAddressEntry(void* ptr);

#ifdef __cplusplus
}
#endif