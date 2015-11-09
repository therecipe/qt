#ifdef __cplusplus
extern "C" {
#endif

void QGestureEvent_Accept(void* ptr, void* gesture);
void QGestureEvent_Accept2(void* ptr, int gestureType);
void* QGestureEvent_Gesture(void* ptr, int ty);
void QGestureEvent_Ignore(void* ptr, void* gesture);
void QGestureEvent_Ignore2(void* ptr, int gestureType);
int QGestureEvent_IsAccepted(void* ptr, void* gesture);
int QGestureEvent_IsAccepted2(void* ptr, int gestureType);
void QGestureEvent_SetAccepted(void* ptr, void* gesture, int value);
void QGestureEvent_SetAccepted2(void* ptr, int gestureType, int value);
void* QGestureEvent_Widget(void* ptr);
void QGestureEvent_DestroyQGestureEvent(void* ptr);

#ifdef __cplusplus
}
#endif