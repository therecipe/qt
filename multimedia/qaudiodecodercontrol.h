#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAudioDecoderControl_BufferAvailable(QtObjectPtr ptr);
void QAudioDecoderControl_ConnectBufferAvailableChanged(QtObjectPtr ptr);
void QAudioDecoderControl_DisconnectBufferAvailableChanged(QtObjectPtr ptr);
void QAudioDecoderControl_ConnectBufferReady(QtObjectPtr ptr);
void QAudioDecoderControl_DisconnectBufferReady(QtObjectPtr ptr);
void QAudioDecoderControl_ConnectError(QtObjectPtr ptr);
void QAudioDecoderControl_DisconnectError(QtObjectPtr ptr);
void QAudioDecoderControl_ConnectFinished(QtObjectPtr ptr);
void QAudioDecoderControl_DisconnectFinished(QtObjectPtr ptr);
void QAudioDecoderControl_SetAudioFormat(QtObjectPtr ptr, QtObjectPtr format);
void QAudioDecoderControl_SetSourceDevice(QtObjectPtr ptr, QtObjectPtr device);
void QAudioDecoderControl_SetSourceFilename(QtObjectPtr ptr, char* fileName);
void QAudioDecoderControl_ConnectSourceChanged(QtObjectPtr ptr);
void QAudioDecoderControl_DisconnectSourceChanged(QtObjectPtr ptr);
QtObjectPtr QAudioDecoderControl_SourceDevice(QtObjectPtr ptr);
char* QAudioDecoderControl_SourceFilename(QtObjectPtr ptr);
void QAudioDecoderControl_Start(QtObjectPtr ptr);
void QAudioDecoderControl_ConnectStateChanged(QtObjectPtr ptr);
void QAudioDecoderControl_DisconnectStateChanged(QtObjectPtr ptr);
void QAudioDecoderControl_Stop(QtObjectPtr ptr);
void QAudioDecoderControl_DestroyQAudioDecoderControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif