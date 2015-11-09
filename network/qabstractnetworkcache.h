#ifdef __cplusplus
extern "C" {
#endif

void QAbstractNetworkCache_Clear(void* ptr);
void* QAbstractNetworkCache_Data(void* ptr, void* url);
void* QAbstractNetworkCache_Prepare(void* ptr, void* metaData);
void QAbstractNetworkCache_UpdateMetaData(void* ptr, void* metaData);
void QAbstractNetworkCache_DestroyQAbstractNetworkCache(void* ptr);

#ifdef __cplusplus
}
#endif