#ifdef __cplusplus
extern "C" {
#endif

void QMediaRecorderControl_ApplySettings(void* ptr);
void QMediaRecorderControl_ConnectError(void* ptr);
void QMediaRecorderControl_DisconnectError(void* ptr);
int QMediaRecorderControl_IsMuted(void* ptr);
void QMediaRecorderControl_ConnectMutedChanged(void* ptr);
void QMediaRecorderControl_DisconnectMutedChanged(void* ptr);
void QMediaRecorderControl_SetMuted(void* ptr, int muted);
int QMediaRecorderControl_SetOutputLocation(void* ptr, void* location);
void QMediaRecorderControl_SetState(void* ptr, int state);
void QMediaRecorderControl_SetVolume(void* ptr, double gain);
void QMediaRecorderControl_ConnectStateChanged(void* ptr);
void QMediaRecorderControl_DisconnectStateChanged(void* ptr);
int QMediaRecorderControl_Status(void* ptr);
void QMediaRecorderControl_ConnectStatusChanged(void* ptr);
void QMediaRecorderControl_DisconnectStatusChanged(void* ptr);
double QMediaRecorderControl_Volume(void* ptr);
void QMediaRecorderControl_DestroyQMediaRecorderControl(void* ptr);

#ifdef __cplusplus
}
#endif