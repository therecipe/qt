#ifdef __cplusplus
extern "C" {
#endif

void QScroller_SetScrollerProperties(void* ptr, void* prop);
void QScroller_EnsureVisible(void* ptr, void* rect, double xmargin, double ymargin);
void QScroller_EnsureVisible2(void* ptr, void* rect, double xmargin, double ymargin, int scrollTime);
int QScroller_QScroller_GrabGesture(void* target, int scrollGestureType);
int QScroller_QScroller_GrabbedGesture(void* target);
int QScroller_QScroller_HasScroller(void* target);
void QScroller_ResendPrepareEvent(void* ptr);
void QScroller_ScrollTo(void* ptr, void* pos);
void QScroller_ScrollTo2(void* ptr, void* pos, int scrollTime);
void* QScroller_QScroller_Scroller(void* target);
void* QScroller_QScroller_Scroller2(void* target);
void QScroller_SetSnapPositionsX2(void* ptr, double first, double interval);
void QScroller_SetSnapPositionsY2(void* ptr, double first, double interval);
void QScroller_ConnectStateChanged(void* ptr);
void QScroller_DisconnectStateChanged(void* ptr);
void QScroller_Stop(void* ptr);
void* QScroller_Target(void* ptr);
void QScroller_QScroller_UngrabGesture(void* target);

#ifdef __cplusplus
}
#endif