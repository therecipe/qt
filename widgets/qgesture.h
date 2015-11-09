#ifdef __cplusplus
extern "C" {
#endif

int QGesture_GestureCancelPolicy(void* ptr);
int QGesture_GestureType(void* ptr);
int QGesture_HasHotSpot(void* ptr);
void QGesture_SetGestureCancelPolicy(void* ptr, int policy);
void QGesture_SetHotSpot(void* ptr, void* value);
void QGesture_UnsetHotSpot(void* ptr);
void* QGesture_NewQGesture(void* parent);
void QGesture_DestroyQGesture(void* ptr);

#ifdef __cplusplus
}
#endif