#ifdef __cplusplus
extern "C" {
#endif

void* QFocusEvent_NewQFocusEvent(int ty, int reason);
int QFocusEvent_GotFocus(void* ptr);
int QFocusEvent_LostFocus(void* ptr);
int QFocusEvent_Reason(void* ptr);

#ifdef __cplusplus
}
#endif