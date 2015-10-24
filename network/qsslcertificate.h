#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSslCertificate_NewQSslCertificate3(QtObjectPtr other);
void QSslCertificate_Clear(QtObjectPtr ptr);
int QSslCertificate_IsBlacklisted(QtObjectPtr ptr);
void QSslCertificate_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QSslCertificate_DestroyQSslCertificate(QtObjectPtr ptr);
int QSslCertificate_IsNull(QtObjectPtr ptr);
int QSslCertificate_IsSelfSigned(QtObjectPtr ptr);
char* QSslCertificate_IssuerInfo(QtObjectPtr ptr, int subject);
char* QSslCertificate_IssuerInfo2(QtObjectPtr ptr, QtObjectPtr attribute);
char* QSslCertificate_SubjectInfo(QtObjectPtr ptr, int subject);
char* QSslCertificate_SubjectInfo2(QtObjectPtr ptr, QtObjectPtr attribute);
char* QSslCertificate_ToText(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif