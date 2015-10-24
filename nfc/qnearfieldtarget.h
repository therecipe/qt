#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QNearFieldTarget_AccessMethods(QtObjectPtr ptr);
void QNearFieldTarget_ConnectDisconnected(QtObjectPtr ptr);
void QNearFieldTarget_DisconnectDisconnected(QtObjectPtr ptr);
int QNearFieldTarget_HasNdefMessage(QtObjectPtr ptr);
int QNearFieldTarget_IsProcessingCommand(QtObjectPtr ptr);
void QNearFieldTarget_ConnectNdefMessagesWritten(QtObjectPtr ptr);
void QNearFieldTarget_DisconnectNdefMessagesWritten(QtObjectPtr ptr);
int QNearFieldTarget_Type(QtObjectPtr ptr);
char* QNearFieldTarget_Url(QtObjectPtr ptr);
void QNearFieldTarget_DestroyQNearFieldTarget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif