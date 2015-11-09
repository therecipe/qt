#ifdef __cplusplus
extern "C" {
#endif

void* QNearFieldShareManager_NewQNearFieldShareManager(void* parent);
void QNearFieldShareManager_ConnectError(void* ptr);
void QNearFieldShareManager_DisconnectError(void* ptr);
void QNearFieldShareManager_SetShareModes(void* ptr, int mode);
int QNearFieldShareManager_ShareError(void* ptr);
int QNearFieldShareManager_ShareModes(void* ptr);
void QNearFieldShareManager_ConnectShareModesChanged(void* ptr);
void QNearFieldShareManager_DisconnectShareModesChanged(void* ptr);
int QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes();
void QNearFieldShareManager_ConnectTargetDetected(void* ptr);
void QNearFieldShareManager_DisconnectTargetDetected(void* ptr);
void QNearFieldShareManager_DestroyQNearFieldShareManager(void* ptr);

#ifdef __cplusplus
}
#endif