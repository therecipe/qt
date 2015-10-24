#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QAbstractNetworkCache_Clear(QtObjectPtr ptr);
QtObjectPtr QAbstractNetworkCache_Data(QtObjectPtr ptr, char* url);
QtObjectPtr QAbstractNetworkCache_Prepare(QtObjectPtr ptr, QtObjectPtr metaData);
void QAbstractNetworkCache_UpdateMetaData(QtObjectPtr ptr, QtObjectPtr metaData);
void QAbstractNetworkCache_DestroyQAbstractNetworkCache(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif