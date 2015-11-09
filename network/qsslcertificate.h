#ifdef __cplusplus
extern "C" {
#endif

void* QSslCertificate_NewQSslCertificate3(void* other);
void QSslCertificate_Clear(void* ptr);
void* QSslCertificate_Digest(void* ptr, int algorithm);
int QSslCertificate_IsBlacklisted(void* ptr);
void QSslCertificate_Swap(void* ptr, void* other);
void QSslCertificate_DestroyQSslCertificate(void* ptr);
void* QSslCertificate_EffectiveDate(void* ptr);
void* QSslCertificate_ExpiryDate(void* ptr);
int QSslCertificate_IsNull(void* ptr);
int QSslCertificate_IsSelfSigned(void* ptr);
char* QSslCertificate_IssuerInfo(void* ptr, int subject);
char* QSslCertificate_IssuerInfo2(void* ptr, void* attribute);
void* QSslCertificate_SerialNumber(void* ptr);
char* QSslCertificate_SubjectInfo(void* ptr, int subject);
char* QSslCertificate_SubjectInfo2(void* ptr, void* attribute);
void* QSslCertificate_ToDer(void* ptr);
void* QSslCertificate_ToPem(void* ptr);
char* QSslCertificate_ToText(void* ptr);
void* QSslCertificate_Version(void* ptr);

#ifdef __cplusplus
}
#endif