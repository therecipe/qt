#ifdef __cplusplus
extern "C" {
#endif

void QAbstractTransition_AddAnimation(void* ptr, void* animation);
void* QAbstractTransition_Machine(void* ptr);
void QAbstractTransition_RemoveAnimation(void* ptr, void* animation);
void QAbstractTransition_SetTargetState(void* ptr, void* target);
void QAbstractTransition_SetTransitionType(void* ptr, int ty);
void* QAbstractTransition_SourceState(void* ptr);
void* QAbstractTransition_TargetState(void* ptr);
void QAbstractTransition_ConnectTargetStateChanged(void* ptr);
void QAbstractTransition_DisconnectTargetStateChanged(void* ptr);
void QAbstractTransition_ConnectTargetStatesChanged(void* ptr);
void QAbstractTransition_DisconnectTargetStatesChanged(void* ptr);
int QAbstractTransition_TransitionType(void* ptr);
void QAbstractTransition_ConnectTriggered(void* ptr);
void QAbstractTransition_DisconnectTriggered(void* ptr);
void QAbstractTransition_DestroyQAbstractTransition(void* ptr);

#ifdef __cplusplus
}
#endif