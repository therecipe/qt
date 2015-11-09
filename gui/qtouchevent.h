#ifdef __cplusplus
extern "C" {
#endif

void* QTouchEvent_Device(void* ptr);
void* QTouchEvent_Target(void* ptr);
int QTouchEvent_TouchPointStates(void* ptr);
void* QTouchEvent_Window(void* ptr);
void QTouchEvent_DestroyQTouchEvent(void* ptr);

#ifdef __cplusplus
}
#endif