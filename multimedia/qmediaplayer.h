#ifdef __cplusplus
extern "C" {
#endif

int QMediaPlayer_BufferStatus(void* ptr);
char* QMediaPlayer_ErrorString(void* ptr);
int QMediaPlayer_IsAudioAvailable(void* ptr);
int QMediaPlayer_IsMuted(void* ptr);
int QMediaPlayer_IsSeekable(void* ptr);
int QMediaPlayer_IsVideoAvailable(void* ptr);
int QMediaPlayer_MediaStatus(void* ptr);
double QMediaPlayer_PlaybackRate(void* ptr);
void* QMediaPlayer_Playlist(void* ptr);
void QMediaPlayer_SetMuted(void* ptr, int muted);
void QMediaPlayer_SetPlaybackRate(void* ptr, double rate);
void QMediaPlayer_SetPlaylist(void* ptr, void* playlist);
void QMediaPlayer_SetVolume(void* ptr, int volume);
int QMediaPlayer_Volume(void* ptr);
void* QMediaPlayer_NewQMediaPlayer(void* parent, int flags);
void QMediaPlayer_ConnectAudioAvailableChanged(void* ptr);
void QMediaPlayer_DisconnectAudioAvailableChanged(void* ptr);
void QMediaPlayer_ConnectBufferStatusChanged(void* ptr);
void QMediaPlayer_DisconnectBufferStatusChanged(void* ptr);
int QMediaPlayer_Error(void* ptr);
void QMediaPlayer_ConnectMediaStatusChanged(void* ptr);
void QMediaPlayer_DisconnectMediaStatusChanged(void* ptr);
void* QMediaPlayer_MediaStream(void* ptr);
void QMediaPlayer_ConnectMutedChanged(void* ptr);
void QMediaPlayer_DisconnectMutedChanged(void* ptr);
void QMediaPlayer_Pause(void* ptr);
void QMediaPlayer_Play(void* ptr);
void QMediaPlayer_ConnectSeekableChanged(void* ptr);
void QMediaPlayer_DisconnectSeekableChanged(void* ptr);
void QMediaPlayer_SetMedia(void* ptr, void* media, void* stream);
void QMediaPlayer_SetVideoOutput3(void* ptr, void* surface);
void QMediaPlayer_ConnectStateChanged(void* ptr);
void QMediaPlayer_DisconnectStateChanged(void* ptr);
void QMediaPlayer_Stop(void* ptr);
void QMediaPlayer_ConnectVideoAvailableChanged(void* ptr);
void QMediaPlayer_DisconnectVideoAvailableChanged(void* ptr);
void QMediaPlayer_ConnectVolumeChanged(void* ptr);
void QMediaPlayer_DisconnectVolumeChanged(void* ptr);
void QMediaPlayer_DestroyQMediaPlayer(void* ptr);

#ifdef __cplusplus
}
#endif