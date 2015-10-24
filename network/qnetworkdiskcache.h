#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkDiskCache_NewQNetworkDiskCache(QtObjectPtr parent);
char* QNetworkDiskCache_CacheDirectory(QtObjectPtr ptr);
void QNetworkDiskCache_Clear(QtObjectPtr ptr);
QtObjectPtr QNetworkDiskCache_Data(QtObjectPtr ptr, char* url);
QtObjectPtr QNetworkDiskCache_Prepare(QtObjectPtr ptr, QtObjectPtr metaData);
void QNetworkDiskCache_SetCacheDirectory(QtObjectPtr ptr, char* cacheDir);
void QNetworkDiskCache_UpdateMetaData(QtObjectPtr ptr, QtObjectPtr metaData);
void QNetworkDiskCache_DestroyQNetworkDiskCache(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif