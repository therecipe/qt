#ifdef __cplusplus
extern "C" {
#endif

void* QSslEllipticCurve_NewQSslEllipticCurve();
int QSslEllipticCurve_IsValid(void* ptr);
int QSslEllipticCurve_IsTlsNamedCurve(void* ptr);
char* QSslEllipticCurve_LongName(void* ptr);
char* QSslEllipticCurve_ShortName(void* ptr);

#ifdef __cplusplus
}
#endif