#ifdef __cplusplus
extern "C" {
#endif

void* QMediaServiceProviderPlugin_Create(void* ptr, char* key);
void QMediaServiceProviderPlugin_Release(void* ptr, void* service);

#ifdef __cplusplus
}
#endif