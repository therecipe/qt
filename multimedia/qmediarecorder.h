#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QMediaRecorder_ActualLocation(QtObjectPtr ptr);
int QMediaRecorder_IsMetaDataAvailable(QtObjectPtr ptr);
int QMediaRecorder_IsMetaDataWritable(QtObjectPtr ptr);
int QMediaRecorder_IsMuted(QtObjectPtr ptr);
char* QMediaRecorder_OutputLocation(QtObjectPtr ptr);
void QMediaRecorder_SetMuted(QtObjectPtr ptr, int muted);
int QMediaRecorder_SetOutputLocation(QtObjectPtr ptr, char* location);
QtObjectPtr QMediaRecorder_NewQMediaRecorder(QtObjectPtr mediaObject, QtObjectPtr parent);
void QMediaRecorder_ConnectActualLocationChanged(QtObjectPtr ptr);
void QMediaRecorder_DisconnectActualLocationChanged(QtObjectPtr ptr);
char* QMediaRecorder_AudioCodecDescription(QtObjectPtr ptr, char* codec);
void QMediaRecorder_ConnectAvailabilityChanged(QtObjectPtr ptr);
void QMediaRecorder_DisconnectAvailabilityChanged(QtObjectPtr ptr);
char* QMediaRecorder_AvailableMetaData(QtObjectPtr ptr);
char* QMediaRecorder_ContainerDescription(QtObjectPtr ptr, char* format);
char* QMediaRecorder_ContainerFormat(QtObjectPtr ptr);
int QMediaRecorder_Error(QtObjectPtr ptr);
char* QMediaRecorder_ErrorString(QtObjectPtr ptr);
int QMediaRecorder_IsAvailable(QtObjectPtr ptr);
QtObjectPtr QMediaRecorder_MediaObject(QtObjectPtr ptr);
char* QMediaRecorder_MetaData(QtObjectPtr ptr, char* key);
void QMediaRecorder_ConnectMetaDataAvailableChanged(QtObjectPtr ptr);
void QMediaRecorder_DisconnectMetaDataAvailableChanged(QtObjectPtr ptr);
void QMediaRecorder_ConnectMetaDataChanged(QtObjectPtr ptr);
void QMediaRecorder_DisconnectMetaDataChanged(QtObjectPtr ptr);
void QMediaRecorder_ConnectMetaDataWritableChanged(QtObjectPtr ptr);
void QMediaRecorder_DisconnectMetaDataWritableChanged(QtObjectPtr ptr);
void QMediaRecorder_ConnectMutedChanged(QtObjectPtr ptr);
void QMediaRecorder_DisconnectMutedChanged(QtObjectPtr ptr);
void QMediaRecorder_Pause(QtObjectPtr ptr);
void QMediaRecorder_Record(QtObjectPtr ptr);
void QMediaRecorder_SetAudioSettings(QtObjectPtr ptr, QtObjectPtr settings);
void QMediaRecorder_SetContainerFormat(QtObjectPtr ptr, char* container);
void QMediaRecorder_SetEncodingSettings(QtObjectPtr ptr, QtObjectPtr audio, QtObjectPtr video, char* container);
void QMediaRecorder_SetMetaData(QtObjectPtr ptr, char* key, char* value);
void QMediaRecorder_SetVideoSettings(QtObjectPtr ptr, QtObjectPtr settings);
void QMediaRecorder_ConnectStateChanged(QtObjectPtr ptr);
void QMediaRecorder_DisconnectStateChanged(QtObjectPtr ptr);
int QMediaRecorder_Status(QtObjectPtr ptr);
void QMediaRecorder_ConnectStatusChanged(QtObjectPtr ptr);
void QMediaRecorder_DisconnectStatusChanged(QtObjectPtr ptr);
void QMediaRecorder_Stop(QtObjectPtr ptr);
char* QMediaRecorder_SupportedAudioCodecs(QtObjectPtr ptr);
char* QMediaRecorder_SupportedContainers(QtObjectPtr ptr);
char* QMediaRecorder_SupportedVideoCodecs(QtObjectPtr ptr);
char* QMediaRecorder_VideoCodecDescription(QtObjectPtr ptr, char* codec);
void QMediaRecorder_DestroyQMediaRecorder(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif