#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QHostInfo_NewQHostInfo2(QtObjectPtr other);
QtObjectPtr QHostInfo_NewQHostInfo(int id);
void QHostInfo_QHostInfo_AbortHostLookup(int id);
int QHostInfo_Error(QtObjectPtr ptr);
char* QHostInfo_ErrorString(QtObjectPtr ptr);
char* QHostInfo_HostName(QtObjectPtr ptr);
int QHostInfo_QHostInfo_LookupHost(char* name, QtObjectPtr receiver, char* member);
int QHostInfo_LookupId(QtObjectPtr ptr);
void QHostInfo_SetError(QtObjectPtr ptr, int error);
void QHostInfo_SetErrorString(QtObjectPtr ptr, char* str);
void QHostInfo_SetHostName(QtObjectPtr ptr, char* hostName);
void QHostInfo_SetLookupId(QtObjectPtr ptr, int id);
void QHostInfo_DestroyQHostInfo(QtObjectPtr ptr);
char* QHostInfo_QHostInfo_LocalHostName();
char* QHostInfo_QHostInfo_LocalDomainName();

#ifdef __cplusplus
}
#endif