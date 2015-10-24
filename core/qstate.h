#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QState_NewQState2(int childMode, QtObjectPtr parent);
QtObjectPtr QState_NewQState(QtObjectPtr parent);
QtObjectPtr QState_AddTransition3(QtObjectPtr ptr, QtObjectPtr target);
QtObjectPtr QState_AddTransition2(QtObjectPtr ptr, QtObjectPtr sender, char* signal, QtObjectPtr target);
void QState_AddTransition(QtObjectPtr ptr, QtObjectPtr transition);
void QState_AssignProperty(QtObjectPtr ptr, QtObjectPtr object, char* name, char* value);
int QState_ChildMode(QtObjectPtr ptr);
void QState_ConnectChildModeChanged(QtObjectPtr ptr);
void QState_DisconnectChildModeChanged(QtObjectPtr ptr);
QtObjectPtr QState_ErrorState(QtObjectPtr ptr);
void QState_ConnectErrorStateChanged(QtObjectPtr ptr);
void QState_DisconnectErrorStateChanged(QtObjectPtr ptr);
void QState_ConnectFinished(QtObjectPtr ptr);
void QState_DisconnectFinished(QtObjectPtr ptr);
QtObjectPtr QState_InitialState(QtObjectPtr ptr);
void QState_ConnectInitialStateChanged(QtObjectPtr ptr);
void QState_DisconnectInitialStateChanged(QtObjectPtr ptr);
void QState_ConnectPropertiesAssigned(QtObjectPtr ptr);
void QState_DisconnectPropertiesAssigned(QtObjectPtr ptr);
void QState_RemoveTransition(QtObjectPtr ptr, QtObjectPtr transition);
void QState_SetChildMode(QtObjectPtr ptr, int mode);
void QState_SetErrorState(QtObjectPtr ptr, QtObjectPtr state);
void QState_SetInitialState(QtObjectPtr ptr, QtObjectPtr state);
void QState_DestroyQState(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif