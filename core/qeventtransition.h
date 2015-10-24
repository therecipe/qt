#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QEventTransition_NewQEventTransition2(QtObjectPtr object, int ty, QtObjectPtr sourceState);
QtObjectPtr QEventTransition_NewQEventTransition(QtObjectPtr sourceState);
QtObjectPtr QEventTransition_EventSource(QtObjectPtr ptr);
int QEventTransition_EventType(QtObjectPtr ptr);
void QEventTransition_SetEventSource(QtObjectPtr ptr, QtObjectPtr object);
void QEventTransition_SetEventType(QtObjectPtr ptr, int ty);
void QEventTransition_DestroyQEventTransition(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif