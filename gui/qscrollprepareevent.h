#ifdef __cplusplus
extern "C" {
#endif

void* QScrollPrepareEvent_NewQScrollPrepareEvent(void* startPos);
void QScrollPrepareEvent_SetContentPos(void* ptr, void* pos);
void QScrollPrepareEvent_SetContentPosRange(void* ptr, void* rect);
void QScrollPrepareEvent_SetViewportSize(void* ptr, void* size);
void QScrollPrepareEvent_DestroyQScrollPrepareEvent(void* ptr);

#ifdef __cplusplus
}
#endif