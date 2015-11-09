#ifdef __cplusplus
extern "C" {
#endif

void QDropEvent_SetDropAction(void* ptr, int action);
void* QDropEvent_NewQDropEvent(void* pos, int actions, void* data, int buttons, int modifiers, int ty);
void QDropEvent_AcceptProposedAction(void* ptr);
int QDropEvent_DropAction(void* ptr);
int QDropEvent_KeyboardModifiers(void* ptr);
void* QDropEvent_MimeData(void* ptr);
int QDropEvent_MouseButtons(void* ptr);
int QDropEvent_PossibleActions(void* ptr);
int QDropEvent_ProposedAction(void* ptr);
void* QDropEvent_Source(void* ptr);

#ifdef __cplusplus
}
#endif