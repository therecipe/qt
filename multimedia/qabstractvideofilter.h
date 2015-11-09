#ifdef __cplusplus
extern "C" {
#endif

int QAbstractVideoFilter_IsActive(void* ptr);
void QAbstractVideoFilter_SetActive(void* ptr, int v);
void QAbstractVideoFilter_ConnectActiveChanged(void* ptr);
void QAbstractVideoFilter_DisconnectActiveChanged(void* ptr);
void* QAbstractVideoFilter_CreateFilterRunnable(void* ptr);

#ifdef __cplusplus
}
#endif