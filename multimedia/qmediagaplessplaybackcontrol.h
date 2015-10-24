#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QMediaGaplessPlaybackControl_ConnectAdvancedToNextMedia(QtObjectPtr ptr);
void QMediaGaplessPlaybackControl_DisconnectAdvancedToNextMedia(QtObjectPtr ptr);
int QMediaGaplessPlaybackControl_IsCrossfadeSupported(QtObjectPtr ptr);
void QMediaGaplessPlaybackControl_SetNextMedia(QtObjectPtr ptr, QtObjectPtr media);
void QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif