#ifdef __cplusplus
extern "C" {
#endif

void* QVideoProbe_NewQVideoProbe(void* parent);
void QVideoProbe_ConnectFlush(void* ptr);
void QVideoProbe_DisconnectFlush(void* ptr);
int QVideoProbe_IsActive(void* ptr);
int QVideoProbe_SetSource(void* ptr, void* source);
int QVideoProbe_SetSource2(void* ptr, void* mediaRecorder);
void QVideoProbe_DestroyQVideoProbe(void* ptr);

#ifdef __cplusplus
}
#endif