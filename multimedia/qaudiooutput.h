#ifdef __cplusplus
extern "C" {
#endif

void* QAudioOutput_NewQAudioOutput2(void* audioDevice, void* format, void* parent);
void* QAudioOutput_NewQAudioOutput(void* format, void* parent);
int QAudioOutput_BufferSize(void* ptr);
int QAudioOutput_BytesFree(void* ptr);
char* QAudioOutput_Category(void* ptr);
void QAudioOutput_ConnectNotify(void* ptr);
void QAudioOutput_DisconnectNotify(void* ptr);
int QAudioOutput_NotifyInterval(void* ptr);
int QAudioOutput_PeriodSize(void* ptr);
void QAudioOutput_Reset(void* ptr);
void QAudioOutput_Resume(void* ptr);
void QAudioOutput_SetBufferSize(void* ptr, int value);
void QAudioOutput_SetCategory(void* ptr, char* category);
void QAudioOutput_SetNotifyInterval(void* ptr, int ms);
void QAudioOutput_SetVolume(void* ptr, double volume);
void* QAudioOutput_Start2(void* ptr);
void QAudioOutput_Start(void* ptr, void* device);
void QAudioOutput_Stop(void* ptr);
void QAudioOutput_Suspend(void* ptr);
double QAudioOutput_Volume(void* ptr);
void QAudioOutput_DestroyQAudioOutput(void* ptr);

#ifdef __cplusplus
}
#endif