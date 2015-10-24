#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDnsLookup_NewQDnsLookup3(int ty, char* name, QtObjectPtr nameserver, QtObjectPtr parent);
int QDnsLookup_Error(QtObjectPtr ptr);
char* QDnsLookup_ErrorString(QtObjectPtr ptr);
char* QDnsLookup_Name(QtObjectPtr ptr);
void QDnsLookup_SetName(QtObjectPtr ptr, char* name);
void QDnsLookup_SetNameserver(QtObjectPtr ptr, QtObjectPtr nameserver);
void QDnsLookup_SetType(QtObjectPtr ptr, int v);
int QDnsLookup_Type(QtObjectPtr ptr);
QtObjectPtr QDnsLookup_NewQDnsLookup(QtObjectPtr parent);
QtObjectPtr QDnsLookup_NewQDnsLookup2(int ty, char* name, QtObjectPtr parent);
void QDnsLookup_Abort(QtObjectPtr ptr);
void QDnsLookup_ConnectFinished(QtObjectPtr ptr);
void QDnsLookup_DisconnectFinished(QtObjectPtr ptr);
int QDnsLookup_IsFinished(QtObjectPtr ptr);
void QDnsLookup_Lookup(QtObjectPtr ptr);
void QDnsLookup_ConnectNameChanged(QtObjectPtr ptr);
void QDnsLookup_DisconnectNameChanged(QtObjectPtr ptr);
void QDnsLookup_ConnectTypeChanged(QtObjectPtr ptr);
void QDnsLookup_DisconnectTypeChanged(QtObjectPtr ptr);
void QDnsLookup_DestroyQDnsLookup(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif