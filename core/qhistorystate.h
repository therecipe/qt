#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QHistoryState_NewQHistoryState2(int ty, QtObjectPtr parent);
QtObjectPtr QHistoryState_NewQHistoryState(QtObjectPtr parent);
QtObjectPtr QHistoryState_DefaultState(QtObjectPtr ptr);
void QHistoryState_ConnectDefaultStateChanged(QtObjectPtr ptr);
void QHistoryState_DisconnectDefaultStateChanged(QtObjectPtr ptr);
int QHistoryState_HistoryType(QtObjectPtr ptr);
void QHistoryState_ConnectHistoryTypeChanged(QtObjectPtr ptr);
void QHistoryState_DisconnectHistoryTypeChanged(QtObjectPtr ptr);
void QHistoryState_SetDefaultState(QtObjectPtr ptr, QtObjectPtr state);
void QHistoryState_SetHistoryType(QtObjectPtr ptr, int ty);
void QHistoryState_DestroyQHistoryState(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif