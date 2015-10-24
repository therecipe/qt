#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDBusSignature_NewQDBusSignature();
QtObjectPtr QDBusSignature_NewQDBusSignature3(QtObjectPtr signature);
QtObjectPtr QDBusSignature_NewQDBusSignature4(char* signature);
QtObjectPtr QDBusSignature_NewQDBusSignature2(char* signature);
void QDBusSignature_SetSignature(QtObjectPtr ptr, char* signature);
char* QDBusSignature_Signature(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif