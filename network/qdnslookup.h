#ifdef __cplusplus
extern "C" {
#endif

void* QDnsLookup_NewQDnsLookup3(int ty, char* name, void* nameserver, void* parent);
int QDnsLookup_Error(void* ptr);
char* QDnsLookup_ErrorString(void* ptr);
char* QDnsLookup_Name(void* ptr);
void QDnsLookup_SetName(void* ptr, char* name);
void QDnsLookup_SetNameserver(void* ptr, void* nameserver);
void QDnsLookup_SetType(void* ptr, int v);
int QDnsLookup_Type(void* ptr);
void* QDnsLookup_NewQDnsLookup(void* parent);
void* QDnsLookup_NewQDnsLookup2(int ty, char* name, void* parent);
void QDnsLookup_Abort(void* ptr);
void QDnsLookup_ConnectFinished(void* ptr);
void QDnsLookup_DisconnectFinished(void* ptr);
int QDnsLookup_IsFinished(void* ptr);
void QDnsLookup_Lookup(void* ptr);
void QDnsLookup_ConnectNameChanged(void* ptr);
void QDnsLookup_DisconnectNameChanged(void* ptr);
void QDnsLookup_ConnectTypeChanged(void* ptr);
void QDnsLookup_DisconnectTypeChanged(void* ptr);
void QDnsLookup_DestroyQDnsLookup(void* ptr);

#ifdef __cplusplus
}
#endif