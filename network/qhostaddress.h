#ifdef __cplusplus
extern "C" {
#endif

void* QHostAddress_NewQHostAddress();
void* QHostAddress_NewQHostAddress9(int address);
void* QHostAddress_NewQHostAddress8(void* address);
void* QHostAddress_NewQHostAddress7(char* address);
void QHostAddress_Clear(void* ptr);
int QHostAddress_IsInSubnet(void* ptr, void* subnet, int netmask);
int QHostAddress_IsLoopback(void* ptr);
int QHostAddress_IsNull(void* ptr);
int QHostAddress_Protocol(void* ptr);
char* QHostAddress_ScopeId(void* ptr);
int QHostAddress_SetAddress5(void* ptr, char* address);
void QHostAddress_SetScopeId(void* ptr, char* id);
char* QHostAddress_ToString(void* ptr);
void QHostAddress_DestroyQHostAddress(void* ptr);

#ifdef __cplusplus
}
#endif