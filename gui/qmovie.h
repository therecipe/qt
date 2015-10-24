#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMovie_CacheMode(QtObjectPtr ptr);
void QMovie_SetCacheMode(QtObjectPtr ptr, int mode);
void QMovie_SetSpeed(QtObjectPtr ptr, int percentSpeed);
int QMovie_Speed(QtObjectPtr ptr);
QtObjectPtr QMovie_NewQMovie2(QtObjectPtr device, QtObjectPtr format, QtObjectPtr parent);
QtObjectPtr QMovie_NewQMovie(QtObjectPtr parent);
QtObjectPtr QMovie_NewQMovie3(char* fileName, QtObjectPtr format, QtObjectPtr parent);
int QMovie_CurrentFrameNumber(QtObjectPtr ptr);
QtObjectPtr QMovie_Device(QtObjectPtr ptr);
void QMovie_ConnectError(QtObjectPtr ptr);
void QMovie_DisconnectError(QtObjectPtr ptr);
char* QMovie_FileName(QtObjectPtr ptr);
void QMovie_ConnectFinished(QtObjectPtr ptr);
void QMovie_DisconnectFinished(QtObjectPtr ptr);
void QMovie_ConnectFrameChanged(QtObjectPtr ptr);
void QMovie_DisconnectFrameChanged(QtObjectPtr ptr);
int QMovie_FrameCount(QtObjectPtr ptr);
int QMovie_IsValid(QtObjectPtr ptr);
int QMovie_JumpToFrame(QtObjectPtr ptr, int frameNumber);
int QMovie_JumpToNextFrame(QtObjectPtr ptr);
int QMovie_LoopCount(QtObjectPtr ptr);
int QMovie_NextFrameDelay(QtObjectPtr ptr);
void QMovie_SetBackgroundColor(QtObjectPtr ptr, QtObjectPtr color);
void QMovie_SetDevice(QtObjectPtr ptr, QtObjectPtr device);
void QMovie_SetFileName(QtObjectPtr ptr, char* fileName);
void QMovie_SetFormat(QtObjectPtr ptr, QtObjectPtr format);
void QMovie_SetPaused(QtObjectPtr ptr, int paused);
void QMovie_SetScaledSize(QtObjectPtr ptr, QtObjectPtr size);
void QMovie_Start(QtObjectPtr ptr);
void QMovie_ConnectStarted(QtObjectPtr ptr);
void QMovie_DisconnectStarted(QtObjectPtr ptr);
void QMovie_ConnectStateChanged(QtObjectPtr ptr);
void QMovie_DisconnectStateChanged(QtObjectPtr ptr);
void QMovie_Stop(QtObjectPtr ptr);
void QMovie_DestroyQMovie(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif