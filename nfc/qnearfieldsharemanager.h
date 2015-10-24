#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNearFieldShareManager_NewQNearFieldShareManager(QtObjectPtr parent);
void QNearFieldShareManager_ConnectError(QtObjectPtr ptr);
void QNearFieldShareManager_DisconnectError(QtObjectPtr ptr);
void QNearFieldShareManager_SetShareModes(QtObjectPtr ptr, int mode);
int QNearFieldShareManager_ShareError(QtObjectPtr ptr);
int QNearFieldShareManager_ShareModes(QtObjectPtr ptr);
void QNearFieldShareManager_ConnectShareModesChanged(QtObjectPtr ptr);
void QNearFieldShareManager_DisconnectShareModesChanged(QtObjectPtr ptr);
int QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes();
void QNearFieldShareManager_ConnectTargetDetected(QtObjectPtr ptr);
void QNearFieldShareManager_DisconnectTargetDetected(QtObjectPtr ptr);
void QNearFieldShareManager_DestroyQNearFieldShareManager(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif