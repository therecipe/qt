#ifdef __cplusplus
extern "C" {
#endif

int QAudioDecoderControl_BufferAvailable(void* ptr);
void QAudioDecoderControl_ConnectBufferAvailableChanged(void* ptr);
void QAudioDecoderControl_DisconnectBufferAvailableChanged(void* ptr);
void QAudioDecoderControl_ConnectBufferReady(void* ptr);
void QAudioDecoderControl_DisconnectBufferReady(void* ptr);
void QAudioDecoderControl_ConnectError(void* ptr);
void QAudioDecoderControl_DisconnectError(void* ptr);
void QAudioDecoderControl_ConnectFinished(void* ptr);
void QAudioDecoderControl_DisconnectFinished(void* ptr);
void QAudioDecoderControl_SetAudioFormat(void* ptr, void* format);
void QAudioDecoderControl_SetSourceDevice(void* ptr, void* device);
void QAudioDecoderControl_SetSourceFilename(void* ptr, char* fileName);
void QAudioDecoderControl_ConnectSourceChanged(void* ptr);
void QAudioDecoderControl_DisconnectSourceChanged(void* ptr);
void* QAudioDecoderControl_SourceDevice(void* ptr);
char* QAudioDecoderControl_SourceFilename(void* ptr);
void QAudioDecoderControl_Start(void* ptr);
void QAudioDecoderControl_ConnectStateChanged(void* ptr);
void QAudioDecoderControl_DisconnectStateChanged(void* ptr);
void QAudioDecoderControl_Stop(void* ptr);
void QAudioDecoderControl_DestroyQAudioDecoderControl(void* ptr);

#ifdef __cplusplus
}
#endif