#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSoundEffect_IsLoaded(QtObjectPtr ptr);
int QSoundEffect_LoopsRemaining(QtObjectPtr ptr);
void QSoundEffect_Play(QtObjectPtr ptr);
void QSoundEffect_Stop(QtObjectPtr ptr);
char* QSoundEffect_QSoundEffect_SupportedMimeTypes();
QtObjectPtr QSoundEffect_NewQSoundEffect(QtObjectPtr parent);
char* QSoundEffect_Category(QtObjectPtr ptr);
void QSoundEffect_ConnectCategoryChanged(QtObjectPtr ptr);
void QSoundEffect_DisconnectCategoryChanged(QtObjectPtr ptr);
int QSoundEffect_IsMuted(QtObjectPtr ptr);
int QSoundEffect_IsPlaying(QtObjectPtr ptr);
void QSoundEffect_ConnectLoadedChanged(QtObjectPtr ptr);
void QSoundEffect_DisconnectLoadedChanged(QtObjectPtr ptr);
int QSoundEffect_LoopCount(QtObjectPtr ptr);
void QSoundEffect_ConnectLoopCountChanged(QtObjectPtr ptr);
void QSoundEffect_DisconnectLoopCountChanged(QtObjectPtr ptr);
void QSoundEffect_ConnectLoopsRemainingChanged(QtObjectPtr ptr);
void QSoundEffect_DisconnectLoopsRemainingChanged(QtObjectPtr ptr);
void QSoundEffect_ConnectMutedChanged(QtObjectPtr ptr);
void QSoundEffect_DisconnectMutedChanged(QtObjectPtr ptr);
void QSoundEffect_ConnectPlayingChanged(QtObjectPtr ptr);
void QSoundEffect_DisconnectPlayingChanged(QtObjectPtr ptr);
void QSoundEffect_SetCategory(QtObjectPtr ptr, char* category);
void QSoundEffect_SetLoopCount(QtObjectPtr ptr, int loopCount);
void QSoundEffect_SetMuted(QtObjectPtr ptr, int muted);
void QSoundEffect_SetSource(QtObjectPtr ptr, char* url);
char* QSoundEffect_Source(QtObjectPtr ptr);
void QSoundEffect_ConnectSourceChanged(QtObjectPtr ptr);
void QSoundEffect_DisconnectSourceChanged(QtObjectPtr ptr);
int QSoundEffect_Status(QtObjectPtr ptr);
void QSoundEffect_ConnectStatusChanged(QtObjectPtr ptr);
void QSoundEffect_DisconnectStatusChanged(QtObjectPtr ptr);
void QSoundEffect_ConnectVolumeChanged(QtObjectPtr ptr);
void QSoundEffect_DisconnectVolumeChanged(QtObjectPtr ptr);
void QSoundEffect_DestroyQSoundEffect(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif