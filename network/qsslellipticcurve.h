#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSslEllipticCurve_NewQSslEllipticCurve();
int QSslEllipticCurve_IsValid(QtObjectPtr ptr);
int QSslEllipticCurve_IsTlsNamedCurve(QtObjectPtr ptr);
char* QSslEllipticCurve_LongName(QtObjectPtr ptr);
char* QSslEllipticCurve_ShortName(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif