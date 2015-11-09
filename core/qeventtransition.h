#ifdef __cplusplus
extern "C" {
#endif

void* QEventTransition_NewQEventTransition2(void* object, int ty, void* sourceState);
void* QEventTransition_NewQEventTransition(void* sourceState);
void* QEventTransition_EventSource(void* ptr);
int QEventTransition_EventType(void* ptr);
void QEventTransition_SetEventSource(void* ptr, void* object);
void QEventTransition_SetEventType(void* ptr, int ty);
void QEventTransition_DestroyQEventTransition(void* ptr);

#ifdef __cplusplus
}
#endif