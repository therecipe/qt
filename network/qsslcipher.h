#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSslCipher_NewQSslCipher();
QtObjectPtr QSslCipher_NewQSslCipher4(QtObjectPtr other);
QtObjectPtr QSslCipher_NewQSslCipher2(char* name);
char* QSslCipher_AuthenticationMethod(QtObjectPtr ptr);
char* QSslCipher_EncryptionMethod(QtObjectPtr ptr);
int QSslCipher_IsNull(QtObjectPtr ptr);
char* QSslCipher_KeyExchangeMethod(QtObjectPtr ptr);
char* QSslCipher_Name(QtObjectPtr ptr);
char* QSslCipher_ProtocolString(QtObjectPtr ptr);
int QSslCipher_SupportedBits(QtObjectPtr ptr);
void QSslCipher_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QSslCipher_UsedBits(QtObjectPtr ptr);
void QSslCipher_DestroyQSslCipher(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif