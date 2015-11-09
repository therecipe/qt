#ifdef __cplusplus
extern "C" {
#endif

void* QState_NewQState2(int childMode, void* parent);
void* QState_NewQState(void* parent);
void* QState_AddTransition3(void* ptr, void* target);
void* QState_AddTransition2(void* ptr, void* sender, char* signal, void* target);
void QState_AddTransition(void* ptr, void* transition);
void QState_AssignProperty(void* ptr, void* object, char* name, void* value);
int QState_ChildMode(void* ptr);
void QState_ConnectChildModeChanged(void* ptr);
void QState_DisconnectChildModeChanged(void* ptr);
void* QState_ErrorState(void* ptr);
void QState_ConnectErrorStateChanged(void* ptr);
void QState_DisconnectErrorStateChanged(void* ptr);
void QState_ConnectFinished(void* ptr);
void QState_DisconnectFinished(void* ptr);
void* QState_InitialState(void* ptr);
void QState_ConnectInitialStateChanged(void* ptr);
void QState_DisconnectInitialStateChanged(void* ptr);
void QState_ConnectPropertiesAssigned(void* ptr);
void QState_DisconnectPropertiesAssigned(void* ptr);
void QState_RemoveTransition(void* ptr, void* transition);
void QState_SetChildMode(void* ptr, int mode);
void QState_SetErrorState(void* ptr, void* state);
void QState_SetInitialState(void* ptr, void* state);
void QState_DestroyQState(void* ptr);

#ifdef __cplusplus
}
#endif