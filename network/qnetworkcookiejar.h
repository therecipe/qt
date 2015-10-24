#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkCookieJar_NewQNetworkCookieJar(QtObjectPtr parent);
int QNetworkCookieJar_DeleteCookie(QtObjectPtr ptr, QtObjectPtr cookie);
int QNetworkCookieJar_InsertCookie(QtObjectPtr ptr, QtObjectPtr cookie);
int QNetworkCookieJar_UpdateCookie(QtObjectPtr ptr, QtObjectPtr cookie);
void QNetworkCookieJar_DestroyQNetworkCookieJar(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif