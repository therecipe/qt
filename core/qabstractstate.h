#ifdef __cplusplus
extern "C" {
#endif

int QAbstractState_Active(void* ptr);
void QAbstractState_ConnectActiveChanged(void* ptr);
void QAbstractState_DisconnectActiveChanged(void* ptr);
void QAbstractState_ConnectEntered(void* ptr);
void QAbstractState_DisconnectEntered(void* ptr);
void QAbstractState_ConnectExited(void* ptr);
void QAbstractState_DisconnectExited(void* ptr);
void* QAbstractState_Machine(void* ptr);
void* QAbstractState_ParentState(void* ptr);
void QAbstractState_DestroyQAbstractState(void* ptr);

#ifdef __cplusplus
}
#endif