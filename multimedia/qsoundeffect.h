#ifdef __cplusplus
extern "C" {
#endif

int QSoundEffect_IsLoaded(void* ptr);
int QSoundEffect_LoopsRemaining(void* ptr);
void QSoundEffect_Play(void* ptr);
void QSoundEffect_Stop(void* ptr);
char* QSoundEffect_QSoundEffect_SupportedMimeTypes();
void* QSoundEffect_NewQSoundEffect(void* parent);
char* QSoundEffect_Category(void* ptr);
void QSoundEffect_ConnectCategoryChanged(void* ptr);
void QSoundEffect_DisconnectCategoryChanged(void* ptr);
int QSoundEffect_IsMuted(void* ptr);
int QSoundEffect_IsPlaying(void* ptr);
void QSoundEffect_ConnectLoadedChanged(void* ptr);
void QSoundEffect_DisconnectLoadedChanged(void* ptr);
int QSoundEffect_LoopCount(void* ptr);
void QSoundEffect_ConnectLoopCountChanged(void* ptr);
void QSoundEffect_DisconnectLoopCountChanged(void* ptr);
void QSoundEffect_ConnectLoopsRemainingChanged(void* ptr);
void QSoundEffect_DisconnectLoopsRemainingChanged(void* ptr);
void QSoundEffect_ConnectMutedChanged(void* ptr);
void QSoundEffect_DisconnectMutedChanged(void* ptr);
void QSoundEffect_ConnectPlayingChanged(void* ptr);
void QSoundEffect_DisconnectPlayingChanged(void* ptr);
void QSoundEffect_SetCategory(void* ptr, char* category);
void QSoundEffect_SetLoopCount(void* ptr, int loopCount);
void QSoundEffect_SetMuted(void* ptr, int muted);
void QSoundEffect_SetSource(void* ptr, void* url);
void QSoundEffect_SetVolume(void* ptr, double volume);
void QSoundEffect_ConnectSourceChanged(void* ptr);
void QSoundEffect_DisconnectSourceChanged(void* ptr);
int QSoundEffect_Status(void* ptr);
void QSoundEffect_ConnectStatusChanged(void* ptr);
void QSoundEffect_DisconnectStatusChanged(void* ptr);
double QSoundEffect_Volume(void* ptr);
void QSoundEffect_ConnectVolumeChanged(void* ptr);
void QSoundEffect_DisconnectVolumeChanged(void* ptr);
void QSoundEffect_DestroyQSoundEffect(void* ptr);

#ifdef __cplusplus
}
#endif