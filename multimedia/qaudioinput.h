#ifdef __cplusplus
extern "C" {
#endif

void* QAudioInput_NewQAudioInput2(void* audioDevice, void* format, void* parent);
void* QAudioInput_NewQAudioInput(void* format, void* parent);
int QAudioInput_BufferSize(void* ptr);
int QAudioInput_BytesReady(void* ptr);
void QAudioInput_ConnectNotify(void* ptr);
void QAudioInput_DisconnectNotify(void* ptr);
int QAudioInput_NotifyInterval(void* ptr);
int QAudioInput_PeriodSize(void* ptr);
void QAudioInput_Reset(void* ptr);
void QAudioInput_Resume(void* ptr);
void QAudioInput_SetBufferSize(void* ptr, int value);
void QAudioInput_SetNotifyInterval(void* ptr, int ms);
void QAudioInput_SetVolume(void* ptr, double volume);
void* QAudioInput_Start2(void* ptr);
void QAudioInput_Start(void* ptr, void* device);
void QAudioInput_Stop(void* ptr);
void QAudioInput_Suspend(void* ptr);
double QAudioInput_Volume(void* ptr);
void QAudioInput_DestroyQAudioInput(void* ptr);

#ifdef __cplusplus
}
#endif