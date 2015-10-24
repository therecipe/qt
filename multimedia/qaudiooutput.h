#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAudioOutput_NewQAudioOutput2(QtObjectPtr audioDevice, QtObjectPtr format, QtObjectPtr parent);
QtObjectPtr QAudioOutput_NewQAudioOutput(QtObjectPtr format, QtObjectPtr parent);
int QAudioOutput_BufferSize(QtObjectPtr ptr);
int QAudioOutput_BytesFree(QtObjectPtr ptr);
char* QAudioOutput_Category(QtObjectPtr ptr);
void QAudioOutput_ConnectNotify(QtObjectPtr ptr);
void QAudioOutput_DisconnectNotify(QtObjectPtr ptr);
int QAudioOutput_NotifyInterval(QtObjectPtr ptr);
int QAudioOutput_PeriodSize(QtObjectPtr ptr);
void QAudioOutput_Reset(QtObjectPtr ptr);
void QAudioOutput_Resume(QtObjectPtr ptr);
void QAudioOutput_SetBufferSize(QtObjectPtr ptr, int value);
void QAudioOutput_SetCategory(QtObjectPtr ptr, char* category);
void QAudioOutput_SetNotifyInterval(QtObjectPtr ptr, int ms);
QtObjectPtr QAudioOutput_Start2(QtObjectPtr ptr);
void QAudioOutput_Start(QtObjectPtr ptr, QtObjectPtr device);
void QAudioOutput_Stop(QtObjectPtr ptr);
void QAudioOutput_Suspend(QtObjectPtr ptr);
void QAudioOutput_DestroyQAudioOutput(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif