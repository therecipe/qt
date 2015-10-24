#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAbstractState_Active(QtObjectPtr ptr);
void QAbstractState_ConnectActiveChanged(QtObjectPtr ptr);
void QAbstractState_DisconnectActiveChanged(QtObjectPtr ptr);
void QAbstractState_ConnectEntered(QtObjectPtr ptr);
void QAbstractState_DisconnectEntered(QtObjectPtr ptr);
void QAbstractState_ConnectExited(QtObjectPtr ptr);
void QAbstractState_DisconnectExited(QtObjectPtr ptr);
QtObjectPtr QAbstractState_Machine(QtObjectPtr ptr);
QtObjectPtr QAbstractState_ParentState(QtObjectPtr ptr);
void QAbstractState_DestroyQAbstractState(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif