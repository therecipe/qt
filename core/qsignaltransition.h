#ifdef __cplusplus
extern "C" {
#endif

void* QSignalTransition_NewQSignalTransition(void* sourceState);
void* QSignalTransition_NewQSignalTransition2(void* sender, char* signal, void* sourceState);
void* QSignalTransition_SenderObject(void* ptr);
void QSignalTransition_ConnectSenderObjectChanged(void* ptr);
void QSignalTransition_DisconnectSenderObjectChanged(void* ptr);
void QSignalTransition_SetSenderObject(void* ptr, void* sender);
void QSignalTransition_SetSignal(void* ptr, void* signal);
void* QSignalTransition_Signal(void* ptr);
void QSignalTransition_ConnectSignalChanged(void* ptr);
void QSignalTransition_DisconnectSignalChanged(void* ptr);
void QSignalTransition_DestroyQSignalTransition(void* ptr);

#ifdef __cplusplus
}
#endif