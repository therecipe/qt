#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QAbstractTransition_AddAnimation(QtObjectPtr ptr, QtObjectPtr animation);
QtObjectPtr QAbstractTransition_Machine(QtObjectPtr ptr);
void QAbstractTransition_RemoveAnimation(QtObjectPtr ptr, QtObjectPtr animation);
void QAbstractTransition_SetTargetState(QtObjectPtr ptr, QtObjectPtr target);
void QAbstractTransition_SetTransitionType(QtObjectPtr ptr, int ty);
QtObjectPtr QAbstractTransition_SourceState(QtObjectPtr ptr);
QtObjectPtr QAbstractTransition_TargetState(QtObjectPtr ptr);
void QAbstractTransition_ConnectTargetStateChanged(QtObjectPtr ptr);
void QAbstractTransition_DisconnectTargetStateChanged(QtObjectPtr ptr);
void QAbstractTransition_ConnectTargetStatesChanged(QtObjectPtr ptr);
void QAbstractTransition_DisconnectTargetStatesChanged(QtObjectPtr ptr);
int QAbstractTransition_TransitionType(QtObjectPtr ptr);
void QAbstractTransition_ConnectTriggered(QtObjectPtr ptr);
void QAbstractTransition_DisconnectTriggered(QtObjectPtr ptr);
void QAbstractTransition_DestroyQAbstractTransition(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif