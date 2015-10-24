#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGesture_GestureCancelPolicy(QtObjectPtr ptr);
int QGesture_GestureType(QtObjectPtr ptr);
int QGesture_HasHotSpot(QtObjectPtr ptr);
void QGesture_SetGestureCancelPolicy(QtObjectPtr ptr, int policy);
void QGesture_SetHotSpot(QtObjectPtr ptr, QtObjectPtr value);
void QGesture_UnsetHotSpot(QtObjectPtr ptr);
QtObjectPtr QGesture_NewQGesture(QtObjectPtr parent);
void QGesture_DestroyQGesture(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif