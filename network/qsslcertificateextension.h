#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSslCertificateExtension_NewQSslCertificateExtension();
QtObjectPtr QSslCertificateExtension_NewQSslCertificateExtension2(QtObjectPtr other);
int QSslCertificateExtension_IsCritical(QtObjectPtr ptr);
int QSslCertificateExtension_IsSupported(QtObjectPtr ptr);
char* QSslCertificateExtension_Name(QtObjectPtr ptr);
char* QSslCertificateExtension_Oid(QtObjectPtr ptr);
void QSslCertificateExtension_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QSslCertificateExtension_Value(QtObjectPtr ptr);
void QSslCertificateExtension_DestroyQSslCertificateExtension(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif