#ifdef __cplusplus
extern "C" {
#endif

void* QDragMoveEvent_NewQDragMoveEvent(void* pos, int actions, void* data, int buttons, int modifiers, int ty);
void QDragMoveEvent_Accept2(void* ptr);
void QDragMoveEvent_Accept(void* ptr, void* rectangle);
void QDragMoveEvent_Ignore2(void* ptr);
void QDragMoveEvent_Ignore(void* ptr, void* rectangle);
void QDragMoveEvent_DestroyQDragMoveEvent(void* ptr);

#ifdef __cplusplus
}
#endif