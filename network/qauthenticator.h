#ifdef __cplusplus
extern "C" {
#endif

void* QAuthenticator_NewQAuthenticator();
void* QAuthenticator_NewQAuthenticator2(void* other);
int QAuthenticator_IsNull(void* ptr);
void* QAuthenticator_Option(void* ptr, char* opt);
char* QAuthenticator_Password(void* ptr);
char* QAuthenticator_Realm(void* ptr);
void QAuthenticator_SetOption(void* ptr, char* opt, void* value);
void QAuthenticator_SetPassword(void* ptr, char* password);
void QAuthenticator_SetUser(void* ptr, char* user);
char* QAuthenticator_User(void* ptr);
void QAuthenticator_DestroyQAuthenticator(void* ptr);

#ifdef __cplusplus
}
#endif