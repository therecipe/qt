#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkRequest_NewQNetworkRequest2(QtObjectPtr other);
QtObjectPtr QNetworkRequest_NewQNetworkRequest(char* url);
char* QNetworkRequest_Attribute(QtObjectPtr ptr, int code, char* defaultValue);
int QNetworkRequest_HasRawHeader(QtObjectPtr ptr, QtObjectPtr headerName);
char* QNetworkRequest_Header(QtObjectPtr ptr, int header);
QtObjectPtr QNetworkRequest_OriginatingObject(QtObjectPtr ptr);
int QNetworkRequest_Priority(QtObjectPtr ptr);
void QNetworkRequest_SetAttribute(QtObjectPtr ptr, int code, char* value);
void QNetworkRequest_SetHeader(QtObjectPtr ptr, int header, char* value);
void QNetworkRequest_SetOriginatingObject(QtObjectPtr ptr, QtObjectPtr object);
void QNetworkRequest_SetPriority(QtObjectPtr ptr, int priority);
void QNetworkRequest_SetRawHeader(QtObjectPtr ptr, QtObjectPtr headerName, QtObjectPtr headerValue);
void QNetworkRequest_SetSslConfiguration(QtObjectPtr ptr, QtObjectPtr config);
void QNetworkRequest_SetUrl(QtObjectPtr ptr, char* url);
void QNetworkRequest_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QNetworkRequest_Url(QtObjectPtr ptr);
void QNetworkRequest_DestroyQNetworkRequest(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif