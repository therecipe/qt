#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkCookie_NewQNetworkCookie(QtObjectPtr name, QtObjectPtr value);
QtObjectPtr QNetworkCookie_NewQNetworkCookie2(QtObjectPtr other);
char* QNetworkCookie_Domain(QtObjectPtr ptr);
int QNetworkCookie_HasSameIdentifier(QtObjectPtr ptr, QtObjectPtr other);
int QNetworkCookie_IsHttpOnly(QtObjectPtr ptr);
int QNetworkCookie_IsSecure(QtObjectPtr ptr);
int QNetworkCookie_IsSessionCookie(QtObjectPtr ptr);
void QNetworkCookie_Normalize(QtObjectPtr ptr, char* url);
char* QNetworkCookie_Path(QtObjectPtr ptr);
void QNetworkCookie_SetDomain(QtObjectPtr ptr, char* domain);
void QNetworkCookie_SetExpirationDate(QtObjectPtr ptr, QtObjectPtr date);
void QNetworkCookie_SetHttpOnly(QtObjectPtr ptr, int enable);
void QNetworkCookie_SetName(QtObjectPtr ptr, QtObjectPtr cookieName);
void QNetworkCookie_SetPath(QtObjectPtr ptr, char* path);
void QNetworkCookie_SetSecure(QtObjectPtr ptr, int enable);
void QNetworkCookie_SetValue(QtObjectPtr ptr, QtObjectPtr value);
void QNetworkCookie_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QNetworkCookie_DestroyQNetworkCookie(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif