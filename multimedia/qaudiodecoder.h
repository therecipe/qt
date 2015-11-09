#ifdef __cplusplus
extern "C" {
#endif

char* QAudioDecoder_ErrorString(void* ptr);
void* QAudioDecoder_NewQAudioDecoder(void* parent);
int QAudioDecoder_BufferAvailable(void* ptr);
void QAudioDecoder_ConnectBufferAvailableChanged(void* ptr);
void QAudioDecoder_DisconnectBufferAvailableChanged(void* ptr);
void QAudioDecoder_ConnectBufferReady(void* ptr);
void QAudioDecoder_DisconnectBufferReady(void* ptr);
int QAudioDecoder_Error(void* ptr);
void QAudioDecoder_ConnectFinished(void* ptr);
void QAudioDecoder_DisconnectFinished(void* ptr);
void QAudioDecoder_SetAudioFormat(void* ptr, void* format);
void QAudioDecoder_SetSourceDevice(void* ptr, void* device);
void QAudioDecoder_SetSourceFilename(void* ptr, char* fileName);
void QAudioDecoder_ConnectSourceChanged(void* ptr);
void QAudioDecoder_DisconnectSourceChanged(void* ptr);
void* QAudioDecoder_SourceDevice(void* ptr);
char* QAudioDecoder_SourceFilename(void* ptr);
void QAudioDecoder_Start(void* ptr);
void QAudioDecoder_ConnectStateChanged(void* ptr);
void QAudioDecoder_DisconnectStateChanged(void* ptr);
void QAudioDecoder_Stop(void* ptr);
void QAudioDecoder_DestroyQAudioDecoder(void* ptr);

#ifdef __cplusplus
}
#endif