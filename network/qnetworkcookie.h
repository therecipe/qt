#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkCookie_NewQNetworkCookie(void* name, void* value);
void* QNetworkCookie_NewQNetworkCookie2(void* other);
char* QNetworkCookie_Domain(void* ptr);
void* QNetworkCookie_ExpirationDate(void* ptr);
int QNetworkCookie_HasSameIdentifier(void* ptr, void* other);
int QNetworkCookie_IsHttpOnly(void* ptr);
int QNetworkCookie_IsSecure(void* ptr);
int QNetworkCookie_IsSessionCookie(void* ptr);
void* QNetworkCookie_Name(void* ptr);
void QNetworkCookie_Normalize(void* ptr, void* url);
char* QNetworkCookie_Path(void* ptr);
void QNetworkCookie_SetDomain(void* ptr, char* domain);
void QNetworkCookie_SetExpirationDate(void* ptr, void* date);
void QNetworkCookie_SetHttpOnly(void* ptr, int enable);
void QNetworkCookie_SetName(void* ptr, void* cookieName);
void QNetworkCookie_SetPath(void* ptr, char* path);
void QNetworkCookie_SetSecure(void* ptr, int enable);
void QNetworkCookie_SetValue(void* ptr, void* value);
void QNetworkCookie_Swap(void* ptr, void* other);
void* QNetworkCookie_ToRawForm(void* ptr, int form);
void* QNetworkCookie_Value(void* ptr);
void QNetworkCookie_DestroyQNetworkCookie(void* ptr);

#ifdef __cplusplus
}
#endif