#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QHostAddress_NewQHostAddress();
QtObjectPtr QHostAddress_NewQHostAddress9(int address);
QtObjectPtr QHostAddress_NewQHostAddress8(QtObjectPtr address);
QtObjectPtr QHostAddress_NewQHostAddress7(char* address);
void QHostAddress_Clear(QtObjectPtr ptr);
int QHostAddress_IsInSubnet(QtObjectPtr ptr, QtObjectPtr subnet, int netmask);
int QHostAddress_IsLoopback(QtObjectPtr ptr);
int QHostAddress_IsNull(QtObjectPtr ptr);
int QHostAddress_Protocol(QtObjectPtr ptr);
char* QHostAddress_ScopeId(QtObjectPtr ptr);
int QHostAddress_SetAddress5(QtObjectPtr ptr, char* address);
void QHostAddress_SetScopeId(QtObjectPtr ptr, char* id);
char* QHostAddress_ToString(QtObjectPtr ptr);
void QHostAddress_DestroyQHostAddress(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif