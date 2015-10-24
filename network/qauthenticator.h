#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAuthenticator_NewQAuthenticator();
QtObjectPtr QAuthenticator_NewQAuthenticator2(QtObjectPtr other);
int QAuthenticator_IsNull(QtObjectPtr ptr);
char* QAuthenticator_Option(QtObjectPtr ptr, char* opt);
char* QAuthenticator_Password(QtObjectPtr ptr);
char* QAuthenticator_Realm(QtObjectPtr ptr);
void QAuthenticator_SetOption(QtObjectPtr ptr, char* opt, char* value);
void QAuthenticator_SetPassword(QtObjectPtr ptr, char* password);
void QAuthenticator_SetUser(QtObjectPtr ptr, char* user);
char* QAuthenticator_User(QtObjectPtr ptr);
void QAuthenticator_DestroyQAuthenticator(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif