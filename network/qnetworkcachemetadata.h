#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkCacheMetaData_NewQNetworkCacheMetaData();
QtObjectPtr QNetworkCacheMetaData_NewQNetworkCacheMetaData2(QtObjectPtr other);
int QNetworkCacheMetaData_IsValid(QtObjectPtr ptr);
int QNetworkCacheMetaData_SaveToDisk(QtObjectPtr ptr);
void QNetworkCacheMetaData_SetExpirationDate(QtObjectPtr ptr, QtObjectPtr dateTime);
void QNetworkCacheMetaData_SetLastModified(QtObjectPtr ptr, QtObjectPtr dateTime);
void QNetworkCacheMetaData_SetSaveToDisk(QtObjectPtr ptr, int allow);
void QNetworkCacheMetaData_SetUrl(QtObjectPtr ptr, char* url);
void QNetworkCacheMetaData_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QNetworkCacheMetaData_Url(QtObjectPtr ptr);
void QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif