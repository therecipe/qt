#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkCacheMetaData_NewQNetworkCacheMetaData();
void* QNetworkCacheMetaData_NewQNetworkCacheMetaData2(void* other);
void* QNetworkCacheMetaData_ExpirationDate(void* ptr);
int QNetworkCacheMetaData_IsValid(void* ptr);
void* QNetworkCacheMetaData_LastModified(void* ptr);
int QNetworkCacheMetaData_SaveToDisk(void* ptr);
void QNetworkCacheMetaData_SetExpirationDate(void* ptr, void* dateTime);
void QNetworkCacheMetaData_SetLastModified(void* ptr, void* dateTime);
void QNetworkCacheMetaData_SetSaveToDisk(void* ptr, int allow);
void QNetworkCacheMetaData_SetUrl(void* ptr, void* url);
void QNetworkCacheMetaData_Swap(void* ptr, void* other);
void QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(void* ptr);

#ifdef __cplusplus
}
#endif