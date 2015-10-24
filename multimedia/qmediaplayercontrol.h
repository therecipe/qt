#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QMediaPlayerControl_ConnectAudioAvailableChanged(QtObjectPtr ptr);
void QMediaPlayerControl_DisconnectAudioAvailableChanged(QtObjectPtr ptr);
int QMediaPlayerControl_BufferStatus(QtObjectPtr ptr);
void QMediaPlayerControl_ConnectBufferStatusChanged(QtObjectPtr ptr);
void QMediaPlayerControl_DisconnectBufferStatusChanged(QtObjectPtr ptr);
void QMediaPlayerControl_ConnectError(QtObjectPtr ptr);
void QMediaPlayerControl_DisconnectError(QtObjectPtr ptr);
int QMediaPlayerControl_IsAudioAvailable(QtObjectPtr ptr);
int QMediaPlayerControl_IsMuted(QtObjectPtr ptr);
int QMediaPlayerControl_IsSeekable(QtObjectPtr ptr);
int QMediaPlayerControl_IsVideoAvailable(QtObjectPtr ptr);
int QMediaPlayerControl_MediaStatus(QtObjectPtr ptr);
void QMediaPlayerControl_ConnectMediaStatusChanged(QtObjectPtr ptr);
void QMediaPlayerControl_DisconnectMediaStatusChanged(QtObjectPtr ptr);
QtObjectPtr QMediaPlayerControl_MediaStream(QtObjectPtr ptr);
void QMediaPlayerControl_ConnectMutedChanged(QtObjectPtr ptr);
void QMediaPlayerControl_DisconnectMutedChanged(QtObjectPtr ptr);
void QMediaPlayerControl_Pause(QtObjectPtr ptr);
void QMediaPlayerControl_Play(QtObjectPtr ptr);
void QMediaPlayerControl_ConnectSeekableChanged(QtObjectPtr ptr);
void QMediaPlayerControl_DisconnectSeekableChanged(QtObjectPtr ptr);
void QMediaPlayerControl_SetMedia(QtObjectPtr ptr, QtObjectPtr media, QtObjectPtr stream);
void QMediaPlayerControl_SetMuted(QtObjectPtr ptr, int mute);
void QMediaPlayerControl_SetVolume(QtObjectPtr ptr, int volume);
void QMediaPlayerControl_ConnectStateChanged(QtObjectPtr ptr);
void QMediaPlayerControl_DisconnectStateChanged(QtObjectPtr ptr);
void QMediaPlayerControl_Stop(QtObjectPtr ptr);
void QMediaPlayerControl_ConnectVideoAvailableChanged(QtObjectPtr ptr);
void QMediaPlayerControl_DisconnectVideoAvailableChanged(QtObjectPtr ptr);
int QMediaPlayerControl_Volume(QtObjectPtr ptr);
void QMediaPlayerControl_ConnectVolumeChanged(QtObjectPtr ptr);
void QMediaPlayerControl_DisconnectVolumeChanged(QtObjectPtr ptr);
void QMediaPlayerControl_DestroyQMediaPlayerControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif