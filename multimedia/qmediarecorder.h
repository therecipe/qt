#ifdef __cplusplus
extern "C" {
#endif

int QMediaRecorder_IsMetaDataAvailable(void* ptr);
int QMediaRecorder_IsMetaDataWritable(void* ptr);
int QMediaRecorder_IsMuted(void* ptr);
void QMediaRecorder_SetMuted(void* ptr, int muted);
int QMediaRecorder_SetOutputLocation(void* ptr, void* location);
void QMediaRecorder_SetVolume(void* ptr, double volume);
double QMediaRecorder_Volume(void* ptr);
void* QMediaRecorder_NewQMediaRecorder(void* mediaObject, void* parent);
char* QMediaRecorder_AudioCodecDescription(void* ptr, char* codec);
void QMediaRecorder_ConnectAvailabilityChanged(void* ptr);
void QMediaRecorder_DisconnectAvailabilityChanged(void* ptr);
char* QMediaRecorder_AvailableMetaData(void* ptr);
char* QMediaRecorder_ContainerDescription(void* ptr, char* format);
char* QMediaRecorder_ContainerFormat(void* ptr);
int QMediaRecorder_Error(void* ptr);
char* QMediaRecorder_ErrorString(void* ptr);
int QMediaRecorder_IsAvailable(void* ptr);
void* QMediaRecorder_MediaObject(void* ptr);
void* QMediaRecorder_MetaData(void* ptr, char* key);
void QMediaRecorder_ConnectMetaDataAvailableChanged(void* ptr);
void QMediaRecorder_DisconnectMetaDataAvailableChanged(void* ptr);
void QMediaRecorder_ConnectMetaDataChanged(void* ptr);
void QMediaRecorder_DisconnectMetaDataChanged(void* ptr);
void QMediaRecorder_ConnectMetaDataWritableChanged(void* ptr);
void QMediaRecorder_DisconnectMetaDataWritableChanged(void* ptr);
void QMediaRecorder_ConnectMutedChanged(void* ptr);
void QMediaRecorder_DisconnectMutedChanged(void* ptr);
void QMediaRecorder_Pause(void* ptr);
void QMediaRecorder_Record(void* ptr);
void QMediaRecorder_SetAudioSettings(void* ptr, void* settings);
void QMediaRecorder_SetContainerFormat(void* ptr, char* container);
void QMediaRecorder_SetEncodingSettings(void* ptr, void* audio, void* video, char* container);
void QMediaRecorder_SetMetaData(void* ptr, char* key, void* value);
void QMediaRecorder_SetVideoSettings(void* ptr, void* settings);
void QMediaRecorder_ConnectStateChanged(void* ptr);
void QMediaRecorder_DisconnectStateChanged(void* ptr);
int QMediaRecorder_Status(void* ptr);
void QMediaRecorder_ConnectStatusChanged(void* ptr);
void QMediaRecorder_DisconnectStatusChanged(void* ptr);
void QMediaRecorder_Stop(void* ptr);
char* QMediaRecorder_SupportedAudioCodecs(void* ptr);
char* QMediaRecorder_SupportedContainers(void* ptr);
char* QMediaRecorder_SupportedVideoCodecs(void* ptr);
char* QMediaRecorder_VideoCodecDescription(void* ptr, char* codec);
void QMediaRecorder_DestroyQMediaRecorder(void* ptr);

#ifdef __cplusplus
}
#endif