#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkDiskCache_NewQNetworkDiskCache(void* parent);
char* QNetworkDiskCache_CacheDirectory(void* ptr);
void QNetworkDiskCache_Clear(void* ptr);
void* QNetworkDiskCache_Data(void* ptr, void* url);
void* QNetworkDiskCache_Prepare(void* ptr, void* metaData);
void QNetworkDiskCache_SetCacheDirectory(void* ptr, char* cacheDir);
void QNetworkDiskCache_UpdateMetaData(void* ptr, void* metaData);
void QNetworkDiskCache_DestroyQNetworkDiskCache(void* ptr);

#ifdef __cplusplus
}
#endif