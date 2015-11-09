#ifdef __cplusplus
extern "C" {
#endif

void* QSslCipher_NewQSslCipher();
void* QSslCipher_NewQSslCipher4(void* other);
void* QSslCipher_NewQSslCipher2(char* name);
char* QSslCipher_AuthenticationMethod(void* ptr);
char* QSslCipher_EncryptionMethod(void* ptr);
int QSslCipher_IsNull(void* ptr);
char* QSslCipher_KeyExchangeMethod(void* ptr);
char* QSslCipher_Name(void* ptr);
char* QSslCipher_ProtocolString(void* ptr);
int QSslCipher_SupportedBits(void* ptr);
void QSslCipher_Swap(void* ptr, void* other);
int QSslCipher_UsedBits(void* ptr);
void QSslCipher_DestroyQSslCipher(void* ptr);

#ifdef __cplusplus
}
#endif