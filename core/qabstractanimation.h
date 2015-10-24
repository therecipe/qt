#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAbstractAnimation_CurrentLoop(QtObjectPtr ptr);
int QAbstractAnimation_CurrentTime(QtObjectPtr ptr);
int QAbstractAnimation_Direction(QtObjectPtr ptr);
int QAbstractAnimation_LoopCount(QtObjectPtr ptr);
void QAbstractAnimation_SetCurrentTime(QtObjectPtr ptr, int msecs);
void QAbstractAnimation_SetDirection(QtObjectPtr ptr, int direction);
void QAbstractAnimation_SetLoopCount(QtObjectPtr ptr, int loopCount);
void QAbstractAnimation_ConnectCurrentLoopChanged(QtObjectPtr ptr);
void QAbstractAnimation_DisconnectCurrentLoopChanged(QtObjectPtr ptr);
int QAbstractAnimation_CurrentLoopTime(QtObjectPtr ptr);
void QAbstractAnimation_ConnectDirectionChanged(QtObjectPtr ptr);
void QAbstractAnimation_DisconnectDirectionChanged(QtObjectPtr ptr);
int QAbstractAnimation_Duration(QtObjectPtr ptr);
void QAbstractAnimation_ConnectFinished(QtObjectPtr ptr);
void QAbstractAnimation_DisconnectFinished(QtObjectPtr ptr);
QtObjectPtr QAbstractAnimation_Group(QtObjectPtr ptr);
void QAbstractAnimation_Pause(QtObjectPtr ptr);
void QAbstractAnimation_Resume(QtObjectPtr ptr);
void QAbstractAnimation_SetPaused(QtObjectPtr ptr, int paused);
void QAbstractAnimation_Start(QtObjectPtr ptr, int policy);
void QAbstractAnimation_ConnectStateChanged(QtObjectPtr ptr);
void QAbstractAnimation_DisconnectStateChanged(QtObjectPtr ptr);
void QAbstractAnimation_Stop(QtObjectPtr ptr);
int QAbstractAnimation_TotalDuration(QtObjectPtr ptr);
void QAbstractAnimation_DestroyQAbstractAnimation(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif