#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QGestureEvent_Accept(QtObjectPtr ptr, QtObjectPtr gesture);
void QGestureEvent_Accept2(QtObjectPtr ptr, int gestureType);
QtObjectPtr QGestureEvent_Gesture(QtObjectPtr ptr, int ty);
void QGestureEvent_Ignore(QtObjectPtr ptr, QtObjectPtr gesture);
void QGestureEvent_Ignore2(QtObjectPtr ptr, int gestureType);
int QGestureEvent_IsAccepted(QtObjectPtr ptr, QtObjectPtr gesture);
int QGestureEvent_IsAccepted2(QtObjectPtr ptr, int gestureType);
void QGestureEvent_SetAccepted(QtObjectPtr ptr, QtObjectPtr gesture, int value);
void QGestureEvent_SetAccepted2(QtObjectPtr ptr, int gestureType, int value);
QtObjectPtr QGestureEvent_Widget(QtObjectPtr ptr);
void QGestureEvent_DestroyQGestureEvent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif