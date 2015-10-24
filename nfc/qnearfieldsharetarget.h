#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QNearFieldShareTarget_Cancel(QtObjectPtr ptr);
void QNearFieldShareTarget_ConnectError(QtObjectPtr ptr);
void QNearFieldShareTarget_DisconnectError(QtObjectPtr ptr);
int QNearFieldShareTarget_IsShareInProgress(QtObjectPtr ptr);
int QNearFieldShareTarget_Share(QtObjectPtr ptr, QtObjectPtr message);
int QNearFieldShareTarget_ShareError(QtObjectPtr ptr);
void QNearFieldShareTarget_ConnectShareFinished(QtObjectPtr ptr);
void QNearFieldShareTarget_DisconnectShareFinished(QtObjectPtr ptr);
int QNearFieldShareTarget_ShareModes(QtObjectPtr ptr);
void QNearFieldShareTarget_DestroyQNearFieldShareTarget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif