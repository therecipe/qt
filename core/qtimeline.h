#ifdef __cplusplus
extern "C" {
#endif

int QTimeLine_CurrentTime(void* ptr);
int QTimeLine_CurveShape(void* ptr);
int QTimeLine_Direction(void* ptr);
int QTimeLine_Duration(void* ptr);
void* QTimeLine_EasingCurve(void* ptr);
int QTimeLine_LoopCount(void* ptr);
void QTimeLine_SetCurrentTime(void* ptr, int msec);
void QTimeLine_SetCurveShape(void* ptr, int shape);
void QTimeLine_SetDirection(void* ptr, int direction);
void QTimeLine_SetDuration(void* ptr, int duration);
void QTimeLine_SetEasingCurve(void* ptr, void* curve);
void QTimeLine_SetLoopCount(void* ptr, int count);
void QTimeLine_SetUpdateInterval(void* ptr, int interval);
int QTimeLine_UpdateInterval(void* ptr);
void* QTimeLine_NewQTimeLine(int duration, void* parent);
int QTimeLine_CurrentFrame(void* ptr);
double QTimeLine_CurrentValue(void* ptr);
int QTimeLine_EndFrame(void* ptr);
void QTimeLine_ConnectFinished(void* ptr);
void QTimeLine_DisconnectFinished(void* ptr);
void QTimeLine_ConnectFrameChanged(void* ptr);
void QTimeLine_DisconnectFrameChanged(void* ptr);
int QTimeLine_FrameForTime(void* ptr, int msec);
void QTimeLine_Resume(void* ptr);
void QTimeLine_SetEndFrame(void* ptr, int frame);
void QTimeLine_SetFrameRange(void* ptr, int startFrame, int endFrame);
void QTimeLine_SetPaused(void* ptr, int paused);
void QTimeLine_SetStartFrame(void* ptr, int frame);
void QTimeLine_Start(void* ptr);
int QTimeLine_StartFrame(void* ptr);
void QTimeLine_ConnectStateChanged(void* ptr);
void QTimeLine_DisconnectStateChanged(void* ptr);
void QTimeLine_Stop(void* ptr);
void QTimeLine_ToggleDirection(void* ptr);
double QTimeLine_ValueForTime(void* ptr, int msec);
void QTimeLine_DestroyQTimeLine(void* ptr);

#ifdef __cplusplus
}
#endif