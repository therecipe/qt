#ifdef __cplusplus
extern "C" {
#endif

void QMediaPlayerControl_ConnectAudioAvailableChanged(void* ptr);
void QMediaPlayerControl_DisconnectAudioAvailableChanged(void* ptr);
int QMediaPlayerControl_BufferStatus(void* ptr);
void QMediaPlayerControl_ConnectBufferStatusChanged(void* ptr);
void QMediaPlayerControl_DisconnectBufferStatusChanged(void* ptr);
void QMediaPlayerControl_ConnectError(void* ptr);
void QMediaPlayerControl_DisconnectError(void* ptr);
int QMediaPlayerControl_IsAudioAvailable(void* ptr);
int QMediaPlayerControl_IsMuted(void* ptr);
int QMediaPlayerControl_IsSeekable(void* ptr);
int QMediaPlayerControl_IsVideoAvailable(void* ptr);
int QMediaPlayerControl_MediaStatus(void* ptr);
void QMediaPlayerControl_ConnectMediaStatusChanged(void* ptr);
void QMediaPlayerControl_DisconnectMediaStatusChanged(void* ptr);
void* QMediaPlayerControl_MediaStream(void* ptr);
void QMediaPlayerControl_ConnectMutedChanged(void* ptr);
void QMediaPlayerControl_DisconnectMutedChanged(void* ptr);
void QMediaPlayerControl_Pause(void* ptr);
void QMediaPlayerControl_Play(void* ptr);
double QMediaPlayerControl_PlaybackRate(void* ptr);
void QMediaPlayerControl_ConnectSeekableChanged(void* ptr);
void QMediaPlayerControl_DisconnectSeekableChanged(void* ptr);
void QMediaPlayerControl_SetMedia(void* ptr, void* media, void* stream);
void QMediaPlayerControl_SetMuted(void* ptr, int mute);
void QMediaPlayerControl_SetPlaybackRate(void* ptr, double rate);
void QMediaPlayerControl_SetVolume(void* ptr, int volume);
void QMediaPlayerControl_ConnectStateChanged(void* ptr);
void QMediaPlayerControl_DisconnectStateChanged(void* ptr);
void QMediaPlayerControl_Stop(void* ptr);
void QMediaPlayerControl_ConnectVideoAvailableChanged(void* ptr);
void QMediaPlayerControl_DisconnectVideoAvailableChanged(void* ptr);
int QMediaPlayerControl_Volume(void* ptr);
void QMediaPlayerControl_ConnectVolumeChanged(void* ptr);
void QMediaPlayerControl_DisconnectVolumeChanged(void* ptr);
void QMediaPlayerControl_DestroyQMediaPlayerControl(void* ptr);

#ifdef __cplusplus
}
#endif