#ifdef __cplusplus
extern "C" {
#endif

void* QHostInfo_NewQHostInfo2(void* other);
void* QHostInfo_NewQHostInfo(int id);
void QHostInfo_QHostInfo_AbortHostLookup(int id);
int QHostInfo_Error(void* ptr);
char* QHostInfo_ErrorString(void* ptr);
char* QHostInfo_HostName(void* ptr);
int QHostInfo_QHostInfo_LookupHost(char* name, void* receiver, char* member);
int QHostInfo_LookupId(void* ptr);
void QHostInfo_SetError(void* ptr, int error);
void QHostInfo_SetErrorString(void* ptr, char* str);
void QHostInfo_SetHostName(void* ptr, char* hostName);
void QHostInfo_SetLookupId(void* ptr, int id);
void QHostInfo_DestroyQHostInfo(void* ptr);
char* QHostInfo_QHostInfo_LocalHostName();
char* QHostInfo_QHostInfo_LocalDomainName();

#ifdef __cplusplus
}
#endif