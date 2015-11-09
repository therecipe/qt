#ifdef __cplusplus
extern "C" {
#endif

void* QSslCertificateExtension_NewQSslCertificateExtension();
void* QSslCertificateExtension_NewQSslCertificateExtension2(void* other);
int QSslCertificateExtension_IsCritical(void* ptr);
int QSslCertificateExtension_IsSupported(void* ptr);
char* QSslCertificateExtension_Name(void* ptr);
char* QSslCertificateExtension_Oid(void* ptr);
void QSslCertificateExtension_Swap(void* ptr, void* other);
void* QSslCertificateExtension_Value(void* ptr);
void QSslCertificateExtension_DestroyQSslCertificateExtension(void* ptr);

#ifdef __cplusplus
}
#endif