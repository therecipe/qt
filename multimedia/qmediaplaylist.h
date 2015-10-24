#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMediaPlaylist_PlaybackMode(QtObjectPtr ptr);
void QMediaPlaylist_SetPlaybackMode(QtObjectPtr ptr, int mode);
QtObjectPtr QMediaPlaylist_NewQMediaPlaylist(QtObjectPtr parent);
int QMediaPlaylist_AddMedia(QtObjectPtr ptr, QtObjectPtr content);
int QMediaPlaylist_Clear(QtObjectPtr ptr);
int QMediaPlaylist_CurrentIndex(QtObjectPtr ptr);
void QMediaPlaylist_ConnectCurrentIndexChanged(QtObjectPtr ptr);
void QMediaPlaylist_DisconnectCurrentIndexChanged(QtObjectPtr ptr);
int QMediaPlaylist_Error(QtObjectPtr ptr);
char* QMediaPlaylist_ErrorString(QtObjectPtr ptr);
int QMediaPlaylist_InsertMedia(QtObjectPtr ptr, int pos, QtObjectPtr content);
int QMediaPlaylist_IsEmpty(QtObjectPtr ptr);
int QMediaPlaylist_IsReadOnly(QtObjectPtr ptr);
void QMediaPlaylist_Load3(QtObjectPtr ptr, QtObjectPtr device, char* format);
void QMediaPlaylist_Load(QtObjectPtr ptr, QtObjectPtr request, char* format);
void QMediaPlaylist_Load2(QtObjectPtr ptr, char* location, char* format);
void QMediaPlaylist_ConnectLoadFailed(QtObjectPtr ptr);
void QMediaPlaylist_DisconnectLoadFailed(QtObjectPtr ptr);
void QMediaPlaylist_ConnectLoaded(QtObjectPtr ptr);
void QMediaPlaylist_DisconnectLoaded(QtObjectPtr ptr);
void QMediaPlaylist_ConnectMediaAboutToBeInserted(QtObjectPtr ptr);
void QMediaPlaylist_DisconnectMediaAboutToBeInserted(QtObjectPtr ptr);
void QMediaPlaylist_ConnectMediaAboutToBeRemoved(QtObjectPtr ptr);
void QMediaPlaylist_DisconnectMediaAboutToBeRemoved(QtObjectPtr ptr);
void QMediaPlaylist_ConnectMediaChanged(QtObjectPtr ptr);
void QMediaPlaylist_DisconnectMediaChanged(QtObjectPtr ptr);
int QMediaPlaylist_MediaCount(QtObjectPtr ptr);
void QMediaPlaylist_ConnectMediaInserted(QtObjectPtr ptr);
void QMediaPlaylist_DisconnectMediaInserted(QtObjectPtr ptr);
QtObjectPtr QMediaPlaylist_MediaObject(QtObjectPtr ptr);
void QMediaPlaylist_ConnectMediaRemoved(QtObjectPtr ptr);
void QMediaPlaylist_DisconnectMediaRemoved(QtObjectPtr ptr);
void QMediaPlaylist_Next(QtObjectPtr ptr);
int QMediaPlaylist_NextIndex(QtObjectPtr ptr, int steps);
void QMediaPlaylist_ConnectPlaybackModeChanged(QtObjectPtr ptr);
void QMediaPlaylist_DisconnectPlaybackModeChanged(QtObjectPtr ptr);
void QMediaPlaylist_Previous(QtObjectPtr ptr);
int QMediaPlaylist_PreviousIndex(QtObjectPtr ptr, int steps);
int QMediaPlaylist_RemoveMedia(QtObjectPtr ptr, int pos);
int QMediaPlaylist_RemoveMedia2(QtObjectPtr ptr, int start, int end);
int QMediaPlaylist_Save2(QtObjectPtr ptr, QtObjectPtr device, char* format);
int QMediaPlaylist_Save(QtObjectPtr ptr, char* location, char* format);
void QMediaPlaylist_SetCurrentIndex(QtObjectPtr ptr, int playlistPosition);
void QMediaPlaylist_Shuffle(QtObjectPtr ptr);
void QMediaPlaylist_DestroyQMediaPlaylist(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif