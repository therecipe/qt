#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QEventLoop_NewQEventLoop(QtObjectPtr parent);
int QEventLoop_Event(QtObjectPtr ptr, QtObjectPtr event);
int QEventLoop_Exec(QtObjectPtr ptr, int flags);
void QEventLoop_Exit(QtObjectPtr ptr, int returnCode);
int QEventLoop_IsRunning(QtObjectPtr ptr);
int QEventLoop_ProcessEvents(QtObjectPtr ptr, int flags);
void QEventLoop_ProcessEvents2(QtObjectPtr ptr, int flags, int maxTime);
void QEventLoop_Quit(QtObjectPtr ptr);
void QEventLoop_WakeUp(QtObjectPtr ptr);
void QEventLoop_DestroyQEventLoop(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif