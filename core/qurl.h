#ifdef __cplusplus
extern "C" {
#endif

void* QUrl_NewQUrl();
void* QUrl_NewQUrl4(void* other);
void* QUrl_NewQUrl3(char* url, int parsingMode);
void* QUrl_NewQUrl2(void* other);
char* QUrl_Authority(void* ptr, int options);
void QUrl_Clear(void* ptr);
char* QUrl_ErrorString(void* ptr);
char* QUrl_FileName(void* ptr, int options);
char* QUrl_Fragment(void* ptr, int options);
char* QUrl_QUrl_FromAce(void* domain);
char* QUrl_QUrl_FromPercentEncoding(void* input);
int QUrl_HasFragment(void* ptr);
int QUrl_HasQuery(void* ptr);
char* QUrl_Host(void* ptr, int options);
char* QUrl_QUrl_IdnWhitelist();
int QUrl_IsEmpty(void* ptr);
int QUrl_IsLocalFile(void* ptr);
int QUrl_IsParentOf(void* ptr, void* childUrl);
int QUrl_IsRelative(void* ptr);
int QUrl_IsValid(void* ptr);
int QUrl_Matches(void* ptr, void* url, int options);
char* QUrl_Password(void* ptr, int options);
char* QUrl_Path(void* ptr, int options);
int QUrl_Port(void* ptr, int defaultPort);
char* QUrl_Query(void* ptr, int options);
char* QUrl_Scheme(void* ptr);
void QUrl_SetAuthority(void* ptr, char* authority, int mode);
void QUrl_SetFragment(void* ptr, char* fragment, int mode);
void QUrl_SetHost(void* ptr, char* host, int mode);
void QUrl_QUrl_SetIdnWhitelist(char* list);
void QUrl_SetPassword(void* ptr, char* password, int mode);
void QUrl_SetPath(void* ptr, char* path, int mode);
void QUrl_SetPort(void* ptr, int port);
void QUrl_SetQuery(void* ptr, char* query, int mode);
void QUrl_SetQuery2(void* ptr, void* query);
void QUrl_SetScheme(void* ptr, char* scheme);
void QUrl_SetUrl(void* ptr, char* url, int parsingMode);
void QUrl_SetUserInfo(void* ptr, char* userInfo, int mode);
void QUrl_SetUserName(void* ptr, char* userName, int mode);
void QUrl_Swap(void* ptr, void* other);
void* QUrl_QUrl_ToAce(char* domain);
char* QUrl_ToDisplayString(void* ptr, int options);
void* QUrl_ToEncoded(void* ptr, int options);
char* QUrl_ToLocalFile(void* ptr);
void* QUrl_QUrl_ToPercentEncoding(char* input, void* exclude, void* include);
char* QUrl_ToString(void* ptr, int options);
char* QUrl_TopLevelDomain(void* ptr, int options);
char* QUrl_Url(void* ptr, int options);
char* QUrl_UserInfo(void* ptr, int options);
char* QUrl_UserName(void* ptr, int options);
void QUrl_DestroyQUrl(void* ptr);

#ifdef __cplusplus
}
#endif