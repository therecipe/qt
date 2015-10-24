#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTimeLine_CurrentTime(QtObjectPtr ptr);
int QTimeLine_CurveShape(QtObjectPtr ptr);
int QTimeLine_Direction(QtObjectPtr ptr);
int QTimeLine_Duration(QtObjectPtr ptr);
int QTimeLine_LoopCount(QtObjectPtr ptr);
void QTimeLine_SetCurrentTime(QtObjectPtr ptr, int msec);
void QTimeLine_SetCurveShape(QtObjectPtr ptr, int shape);
void QTimeLine_SetDirection(QtObjectPtr ptr, int direction);
void QTimeLine_SetDuration(QtObjectPtr ptr, int duration);
void QTimeLine_SetEasingCurve(QtObjectPtr ptr, QtObjectPtr curve);
void QTimeLine_SetLoopCount(QtObjectPtr ptr, int count);
void QTimeLine_SetUpdateInterval(QtObjectPtr ptr, int interval);
int QTimeLine_UpdateInterval(QtObjectPtr ptr);
QtObjectPtr QTimeLine_NewQTimeLine(int duration, QtObjectPtr parent);
int QTimeLine_CurrentFrame(QtObjectPtr ptr);
int QTimeLine_EndFrame(QtObjectPtr ptr);
void QTimeLine_ConnectFinished(QtObjectPtr ptr);
void QTimeLine_DisconnectFinished(QtObjectPtr ptr);
void QTimeLine_ConnectFrameChanged(QtObjectPtr ptr);
void QTimeLine_DisconnectFrameChanged(QtObjectPtr ptr);
int QTimeLine_FrameForTime(QtObjectPtr ptr, int msec);
void QTimeLine_Resume(QtObjectPtr ptr);
void QTimeLine_SetEndFrame(QtObjectPtr ptr, int frame);
void QTimeLine_SetFrameRange(QtObjectPtr ptr, int startFrame, int endFrame);
void QTimeLine_SetPaused(QtObjectPtr ptr, int paused);
void QTimeLine_SetStartFrame(QtObjectPtr ptr, int frame);
void QTimeLine_Start(QtObjectPtr ptr);
int QTimeLine_StartFrame(QtObjectPtr ptr);
void QTimeLine_ConnectStateChanged(QtObjectPtr ptr);
void QTimeLine_DisconnectStateChanged(QtObjectPtr ptr);
void QTimeLine_Stop(QtObjectPtr ptr);
void QTimeLine_ToggleDirection(QtObjectPtr ptr);
void QTimeLine_DestroyQTimeLine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif