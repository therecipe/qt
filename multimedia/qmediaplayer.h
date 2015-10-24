#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMediaPlayer_BufferStatus(QtObjectPtr ptr);
char* QMediaPlayer_ErrorString(QtObjectPtr ptr);
int QMediaPlayer_IsAudioAvailable(QtObjectPtr ptr);
int QMediaPlayer_IsMuted(QtObjectPtr ptr);
int QMediaPlayer_IsSeekable(QtObjectPtr ptr);
int QMediaPlayer_IsVideoAvailable(QtObjectPtr ptr);
int QMediaPlayer_MediaStatus(QtObjectPtr ptr);
QtObjectPtr QMediaPlayer_Playlist(QtObjectPtr ptr);
void QMediaPlayer_SetMuted(QtObjectPtr ptr, int muted);
void QMediaPlayer_SetPlaylist(QtObjectPtr ptr, QtObjectPtr playlist);
void QMediaPlayer_SetVolume(QtObjectPtr ptr, int volume);
int QMediaPlayer_Volume(QtObjectPtr ptr);
QtObjectPtr QMediaPlayer_NewQMediaPlayer(QtObjectPtr parent, int flags);
void QMediaPlayer_ConnectAudioAvailableChanged(QtObjectPtr ptr);
void QMediaPlayer_DisconnectAudioAvailableChanged(QtObjectPtr ptr);
void QMediaPlayer_ConnectBufferStatusChanged(QtObjectPtr ptr);
void QMediaPlayer_DisconnectBufferStatusChanged(QtObjectPtr ptr);
int QMediaPlayer_Error(QtObjectPtr ptr);
void QMediaPlayer_ConnectMediaStatusChanged(QtObjectPtr ptr);
void QMediaPlayer_DisconnectMediaStatusChanged(QtObjectPtr ptr);
QtObjectPtr QMediaPlayer_MediaStream(QtObjectPtr ptr);
void QMediaPlayer_ConnectMutedChanged(QtObjectPtr ptr);
void QMediaPlayer_DisconnectMutedChanged(QtObjectPtr ptr);
void QMediaPlayer_Pause(QtObjectPtr ptr);
void QMediaPlayer_Play(QtObjectPtr ptr);
void QMediaPlayer_ConnectSeekableChanged(QtObjectPtr ptr);
void QMediaPlayer_DisconnectSeekableChanged(QtObjectPtr ptr);
void QMediaPlayer_SetMedia(QtObjectPtr ptr, QtObjectPtr media, QtObjectPtr stream);
void QMediaPlayer_SetVideoOutput3(QtObjectPtr ptr, QtObjectPtr surface);
void QMediaPlayer_ConnectStateChanged(QtObjectPtr ptr);
void QMediaPlayer_DisconnectStateChanged(QtObjectPtr ptr);
void QMediaPlayer_Stop(QtObjectPtr ptr);
void QMediaPlayer_ConnectVideoAvailableChanged(QtObjectPtr ptr);
void QMediaPlayer_DisconnectVideoAvailableChanged(QtObjectPtr ptr);
void QMediaPlayer_ConnectVolumeChanged(QtObjectPtr ptr);
void QMediaPlayer_DisconnectVolumeChanged(QtObjectPtr ptr);
void QMediaPlayer_DestroyQMediaPlayer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif