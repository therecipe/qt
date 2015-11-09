#ifdef __cplusplus
extern "C" {
#endif

int QSwipeGesture_HorizontalDirection(void* ptr);
void QSwipeGesture_SetSwipeAngle(void* ptr, double value);
double QSwipeGesture_SwipeAngle(void* ptr);
int QSwipeGesture_VerticalDirection(void* ptr);
void QSwipeGesture_DestroyQSwipeGesture(void* ptr);

#ifdef __cplusplus
}
#endif