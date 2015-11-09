#ifdef __cplusplus
extern "C" {
#endif

void* QDBusSignature_NewQDBusSignature();
void* QDBusSignature_NewQDBusSignature3(void* signature);
void* QDBusSignature_NewQDBusSignature4(char* signature);
void* QDBusSignature_NewQDBusSignature2(char* signature);
void QDBusSignature_SetSignature(void* ptr, char* signature);
char* QDBusSignature_Signature(void* ptr);

#ifdef __cplusplus
}
#endif