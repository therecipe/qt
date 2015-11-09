#ifdef __cplusplus
extern "C" {
#endif

int QMovie_CacheMode(void* ptr);
void QMovie_SetCacheMode(void* ptr, int mode);
void QMovie_SetSpeed(void* ptr, int percentSpeed);
int QMovie_Speed(void* ptr);
void* QMovie_NewQMovie2(void* device, void* format, void* parent);
void* QMovie_NewQMovie(void* parent);
void* QMovie_NewQMovie3(char* fileName, void* format, void* parent);
void* QMovie_BackgroundColor(void* ptr);
int QMovie_CurrentFrameNumber(void* ptr);
void* QMovie_Device(void* ptr);
void QMovie_ConnectError(void* ptr);
void QMovie_DisconnectError(void* ptr);
char* QMovie_FileName(void* ptr);
void QMovie_ConnectFinished(void* ptr);
void QMovie_DisconnectFinished(void* ptr);
void* QMovie_Format(void* ptr);
void QMovie_ConnectFrameChanged(void* ptr);
void QMovie_DisconnectFrameChanged(void* ptr);
int QMovie_FrameCount(void* ptr);
int QMovie_IsValid(void* ptr);
int QMovie_JumpToFrame(void* ptr, int frameNumber);
int QMovie_JumpToNextFrame(void* ptr);
int QMovie_LoopCount(void* ptr);
int QMovie_NextFrameDelay(void* ptr);
void QMovie_SetBackgroundColor(void* ptr, void* color);
void QMovie_SetDevice(void* ptr, void* device);
void QMovie_SetFileName(void* ptr, char* fileName);
void QMovie_SetFormat(void* ptr, void* format);
void QMovie_SetPaused(void* ptr, int paused);
void QMovie_SetScaledSize(void* ptr, void* size);
void QMovie_Start(void* ptr);
void QMovie_ConnectStarted(void* ptr);
void QMovie_DisconnectStarted(void* ptr);
void QMovie_ConnectStateChanged(void* ptr);
void QMovie_DisconnectStateChanged(void* ptr);
void QMovie_Stop(void* ptr);
void QMovie_DestroyQMovie(void* ptr);

#ifdef __cplusplus
}
#endif