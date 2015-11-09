#ifdef __cplusplus
extern "C" {
#endif

void* QHistoryState_NewQHistoryState2(int ty, void* parent);
void* QHistoryState_NewQHistoryState(void* parent);
void* QHistoryState_DefaultState(void* ptr);
void QHistoryState_ConnectDefaultStateChanged(void* ptr);
void QHistoryState_DisconnectDefaultStateChanged(void* ptr);
int QHistoryState_HistoryType(void* ptr);
void QHistoryState_ConnectHistoryTypeChanged(void* ptr);
void QHistoryState_DisconnectHistoryTypeChanged(void* ptr);
void QHistoryState_SetDefaultState(void* ptr, void* state);
void QHistoryState_SetHistoryType(void* ptr, int ty);
void QHistoryState_DestroyQHistoryState(void* ptr);

#ifdef __cplusplus
}
#endif