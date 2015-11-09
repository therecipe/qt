#ifdef __cplusplus
extern "C" {
#endif

void QMediaGaplessPlaybackControl_ConnectAdvancedToNextMedia(void* ptr);
void QMediaGaplessPlaybackControl_DisconnectAdvancedToNextMedia(void* ptr);
double QMediaGaplessPlaybackControl_CrossfadeTime(void* ptr);
int QMediaGaplessPlaybackControl_IsCrossfadeSupported(void* ptr);
void QMediaGaplessPlaybackControl_SetCrossfadeTime(void* ptr, double crossfadeTime);
void QMediaGaplessPlaybackControl_SetNextMedia(void* ptr, void* media);
void QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(void* ptr);

#ifdef __cplusplus
}
#endif