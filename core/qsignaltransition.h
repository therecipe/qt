#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSignalTransition_NewQSignalTransition(QtObjectPtr sourceState);
QtObjectPtr QSignalTransition_NewQSignalTransition2(QtObjectPtr sender, char* signal, QtObjectPtr sourceState);
QtObjectPtr QSignalTransition_SenderObject(QtObjectPtr ptr);
void QSignalTransition_ConnectSenderObjectChanged(QtObjectPtr ptr);
void QSignalTransition_DisconnectSenderObjectChanged(QtObjectPtr ptr);
void QSignalTransition_SetSenderObject(QtObjectPtr ptr, QtObjectPtr sender);
void QSignalTransition_SetSignal(QtObjectPtr ptr, QtObjectPtr signal);
void QSignalTransition_ConnectSignalChanged(QtObjectPtr ptr);
void QSignalTransition_DisconnectSignalChanged(QtObjectPtr ptr);
void QSignalTransition_DestroyQSignalTransition(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif