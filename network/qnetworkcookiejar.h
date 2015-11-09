#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkCookieJar_NewQNetworkCookieJar(void* parent);
int QNetworkCookieJar_DeleteCookie(void* ptr, void* cookie);
int QNetworkCookieJar_InsertCookie(void* ptr, void* cookie);
int QNetworkCookieJar_UpdateCookie(void* ptr, void* cookie);
void QNetworkCookieJar_DestroyQNetworkCookieJar(void* ptr);

#ifdef __cplusplus
}
#endif