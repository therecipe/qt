#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkAddressEntry_NewQNetworkAddressEntry();
QtObjectPtr QNetworkAddressEntry_NewQNetworkAddressEntry2(QtObjectPtr other);
int QNetworkAddressEntry_PrefixLength(QtObjectPtr ptr);
void QNetworkAddressEntry_SetBroadcast(QtObjectPtr ptr, QtObjectPtr newBroadcast);
void QNetworkAddressEntry_SetIp(QtObjectPtr ptr, QtObjectPtr newIp);
void QNetworkAddressEntry_SetNetmask(QtObjectPtr ptr, QtObjectPtr newNetmask);
void QNetworkAddressEntry_SetPrefixLength(QtObjectPtr ptr, int length);
void QNetworkAddressEntry_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QNetworkAddressEntry_DestroyQNetworkAddressEntry(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif