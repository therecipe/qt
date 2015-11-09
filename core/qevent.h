#ifdef __cplusplus
extern "C" {
#endif

void* QEvent_NewQEvent(int ty);
void QEvent_Accept(void* ptr);
void QEvent_Ignore(void* ptr);
int QEvent_IsAccepted(void* ptr);
int QEvent_QEvent_RegisterEventType(int hint);
void QEvent_SetAccepted(void* ptr, int accepted);
int QEvent_Spontaneous(void* ptr);
int QEvent_Type(void* ptr);
void QEvent_DestroyQEvent(void* ptr);

#ifdef __cplusplus
}
#endif