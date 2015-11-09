#ifdef __cplusplus
extern "C" {
#endif

void QNearFieldShareTarget_Cancel(void* ptr);
void QNearFieldShareTarget_ConnectError(void* ptr);
void QNearFieldShareTarget_DisconnectError(void* ptr);
int QNearFieldShareTarget_IsShareInProgress(void* ptr);
int QNearFieldShareTarget_Share(void* ptr, void* message);
int QNearFieldShareTarget_ShareError(void* ptr);
void QNearFieldShareTarget_ConnectShareFinished(void* ptr);
void QNearFieldShareTarget_DisconnectShareFinished(void* ptr);
int QNearFieldShareTarget_ShareModes(void* ptr);
void QNearFieldShareTarget_DestroyQNearFieldShareTarget(void* ptr);

#ifdef __cplusplus
}
#endif