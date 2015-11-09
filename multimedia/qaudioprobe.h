#ifdef __cplusplus
extern "C" {
#endif

void* QAudioProbe_NewQAudioProbe(void* parent);
void QAudioProbe_ConnectFlush(void* ptr);
void QAudioProbe_DisconnectFlush(void* ptr);
int QAudioProbe_IsActive(void* ptr);
int QAudioProbe_SetSource(void* ptr, void* source);
int QAudioProbe_SetSource2(void* ptr, void* mediaRecorder);
void QAudioProbe_DestroyQAudioProbe(void* ptr);

#ifdef __cplusplus
}
#endif