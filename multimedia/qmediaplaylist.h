#ifdef __cplusplus
extern "C" {
#endif

int QMediaPlaylist_PlaybackMode(void* ptr);
void QMediaPlaylist_SetPlaybackMode(void* ptr, int mode);
void* QMediaPlaylist_NewQMediaPlaylist(void* parent);
int QMediaPlaylist_AddMedia(void* ptr, void* content);
int QMediaPlaylist_Clear(void* ptr);
int QMediaPlaylist_CurrentIndex(void* ptr);
void QMediaPlaylist_ConnectCurrentIndexChanged(void* ptr);
void QMediaPlaylist_DisconnectCurrentIndexChanged(void* ptr);
int QMediaPlaylist_Error(void* ptr);
char* QMediaPlaylist_ErrorString(void* ptr);
int QMediaPlaylist_InsertMedia(void* ptr, int pos, void* content);
int QMediaPlaylist_IsEmpty(void* ptr);
int QMediaPlaylist_IsReadOnly(void* ptr);
void QMediaPlaylist_Load3(void* ptr, void* device, char* format);
void QMediaPlaylist_Load(void* ptr, void* request, char* format);
void QMediaPlaylist_Load2(void* ptr, void* location, char* format);
void QMediaPlaylist_ConnectLoadFailed(void* ptr);
void QMediaPlaylist_DisconnectLoadFailed(void* ptr);
void QMediaPlaylist_ConnectLoaded(void* ptr);
void QMediaPlaylist_DisconnectLoaded(void* ptr);
void QMediaPlaylist_ConnectMediaAboutToBeInserted(void* ptr);
void QMediaPlaylist_DisconnectMediaAboutToBeInserted(void* ptr);
void QMediaPlaylist_ConnectMediaAboutToBeRemoved(void* ptr);
void QMediaPlaylist_DisconnectMediaAboutToBeRemoved(void* ptr);
void QMediaPlaylist_ConnectMediaChanged(void* ptr);
void QMediaPlaylist_DisconnectMediaChanged(void* ptr);
int QMediaPlaylist_MediaCount(void* ptr);
void QMediaPlaylist_ConnectMediaInserted(void* ptr);
void QMediaPlaylist_DisconnectMediaInserted(void* ptr);
void* QMediaPlaylist_MediaObject(void* ptr);
void QMediaPlaylist_ConnectMediaRemoved(void* ptr);
void QMediaPlaylist_DisconnectMediaRemoved(void* ptr);
void QMediaPlaylist_Next(void* ptr);
int QMediaPlaylist_NextIndex(void* ptr, int steps);
void QMediaPlaylist_ConnectPlaybackModeChanged(void* ptr);
void QMediaPlaylist_DisconnectPlaybackModeChanged(void* ptr);
void QMediaPlaylist_Previous(void* ptr);
int QMediaPlaylist_PreviousIndex(void* ptr, int steps);
int QMediaPlaylist_RemoveMedia(void* ptr, int pos);
int QMediaPlaylist_RemoveMedia2(void* ptr, int start, int end);
int QMediaPlaylist_Save2(void* ptr, void* device, char* format);
int QMediaPlaylist_Save(void* ptr, void* location, char* format);
void QMediaPlaylist_SetCurrentIndex(void* ptr, int playlistPosition);
void QMediaPlaylist_Shuffle(void* ptr);
void QMediaPlaylist_DestroyQMediaPlaylist(void* ptr);

#ifdef __cplusplus
}
#endif