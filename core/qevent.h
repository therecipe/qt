#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QEvent_NewQEvent(int ty);
void QEvent_Accept(QtObjectPtr ptr);
void QEvent_Ignore(QtObjectPtr ptr);
int QEvent_IsAccepted(QtObjectPtr ptr);
int QEvent_QEvent_RegisterEventType(int hint);
void QEvent_SetAccepted(QtObjectPtr ptr, int accepted);
int QEvent_Spontaneous(QtObjectPtr ptr);
int QEvent_Type(QtObjectPtr ptr);
void QEvent_DestroyQEvent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif