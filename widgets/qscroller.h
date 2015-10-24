#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QScroller_SetScrollerProperties(QtObjectPtr ptr, QtObjectPtr prop);
int QScroller_QScroller_GrabGesture(QtObjectPtr target, int scrollGestureType);
int QScroller_QScroller_GrabbedGesture(QtObjectPtr target);
int QScroller_QScroller_HasScroller(QtObjectPtr target);
void QScroller_ResendPrepareEvent(QtObjectPtr ptr);
void QScroller_ScrollTo(QtObjectPtr ptr, QtObjectPtr pos);
void QScroller_ScrollTo2(QtObjectPtr ptr, QtObjectPtr pos, int scrollTime);
QtObjectPtr QScroller_QScroller_Scroller(QtObjectPtr target);
QtObjectPtr QScroller_QScroller_Scroller2(QtObjectPtr target);
void QScroller_ConnectStateChanged(QtObjectPtr ptr);
void QScroller_DisconnectStateChanged(QtObjectPtr ptr);
void QScroller_Stop(QtObjectPtr ptr);
QtObjectPtr QScroller_Target(QtObjectPtr ptr);
void QScroller_QScroller_UngrabGesture(QtObjectPtr target);

#ifdef __cplusplus
}
#endif