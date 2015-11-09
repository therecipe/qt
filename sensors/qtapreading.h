#ifdef __cplusplus
extern "C" {
#endif

int QTapReading_IsDoubleTap(void* ptr);
int QTapReading_TapDirection(void* ptr);
void QTapReading_SetDoubleTap(void* ptr, int doubleTap);
void QTapReading_SetTapDirection(void* ptr, int tapDirection);

#ifdef __cplusplus
}
#endif