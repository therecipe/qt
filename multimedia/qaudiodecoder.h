#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QAudioDecoder_ErrorString(QtObjectPtr ptr);
QtObjectPtr QAudioDecoder_NewQAudioDecoder(QtObjectPtr parent);
int QAudioDecoder_BufferAvailable(QtObjectPtr ptr);
void QAudioDecoder_ConnectBufferAvailableChanged(QtObjectPtr ptr);
void QAudioDecoder_DisconnectBufferAvailableChanged(QtObjectPtr ptr);
void QAudioDecoder_ConnectBufferReady(QtObjectPtr ptr);
void QAudioDecoder_DisconnectBufferReady(QtObjectPtr ptr);
int QAudioDecoder_Error(QtObjectPtr ptr);
void QAudioDecoder_ConnectFinished(QtObjectPtr ptr);
void QAudioDecoder_DisconnectFinished(QtObjectPtr ptr);
void QAudioDecoder_SetAudioFormat(QtObjectPtr ptr, QtObjectPtr format);
void QAudioDecoder_SetSourceDevice(QtObjectPtr ptr, QtObjectPtr device);
void QAudioDecoder_SetSourceFilename(QtObjectPtr ptr, char* fileName);
void QAudioDecoder_ConnectSourceChanged(QtObjectPtr ptr);
void QAudioDecoder_DisconnectSourceChanged(QtObjectPtr ptr);
QtObjectPtr QAudioDecoder_SourceDevice(QtObjectPtr ptr);
char* QAudioDecoder_SourceFilename(QtObjectPtr ptr);
void QAudioDecoder_Start(QtObjectPtr ptr);
void QAudioDecoder_ConnectStateChanged(QtObjectPtr ptr);
void QAudioDecoder_DisconnectStateChanged(QtObjectPtr ptr);
void QAudioDecoder_Stop(QtObjectPtr ptr);
void QAudioDecoder_DestroyQAudioDecoder(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif