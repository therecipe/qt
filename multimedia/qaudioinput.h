#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAudioInput_NewQAudioInput2(QtObjectPtr audioDevice, QtObjectPtr format, QtObjectPtr parent);
QtObjectPtr QAudioInput_NewQAudioInput(QtObjectPtr format, QtObjectPtr parent);
int QAudioInput_BufferSize(QtObjectPtr ptr);
int QAudioInput_BytesReady(QtObjectPtr ptr);
void QAudioInput_ConnectNotify(QtObjectPtr ptr);
void QAudioInput_DisconnectNotify(QtObjectPtr ptr);
int QAudioInput_NotifyInterval(QtObjectPtr ptr);
int QAudioInput_PeriodSize(QtObjectPtr ptr);
void QAudioInput_Reset(QtObjectPtr ptr);
void QAudioInput_Resume(QtObjectPtr ptr);
void QAudioInput_SetBufferSize(QtObjectPtr ptr, int value);
void QAudioInput_SetNotifyInterval(QtObjectPtr ptr, int ms);
QtObjectPtr QAudioInput_Start2(QtObjectPtr ptr);
void QAudioInput_Start(QtObjectPtr ptr, QtObjectPtr device);
void QAudioInput_Stop(QtObjectPtr ptr);
void QAudioInput_Suspend(QtObjectPtr ptr);
void QAudioInput_DestroyQAudioInput(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif