#ifdef __cplusplus
extern "C" {
#endif

int QNearFieldManager_RegisterNdefMessageHandler(void* ptr, void* object, char* method);
int QNearFieldManager_StartTargetDetection(void* ptr);
void* QNearFieldManager_NewQNearFieldManager(void* parent);
int QNearFieldManager_IsAvailable(void* ptr);
int QNearFieldManager_RegisterNdefMessageHandler2(void* ptr, int typeNameFormat, void* ty, void* object, char* method);
int QNearFieldManager_RegisterNdefMessageHandler3(void* ptr, void* filter, void* object, char* method);
void QNearFieldManager_SetTargetAccessModes(void* ptr, int accessModes);
void QNearFieldManager_StopTargetDetection(void* ptr);
int QNearFieldManager_TargetAccessModes(void* ptr);
void QNearFieldManager_ConnectTargetDetected(void* ptr);
void QNearFieldManager_DisconnectTargetDetected(void* ptr);
void QNearFieldManager_ConnectTargetLost(void* ptr);
void QNearFieldManager_DisconnectTargetLost(void* ptr);
int QNearFieldManager_UnregisterNdefMessageHandler(void* ptr, int handlerId);
void QNearFieldManager_DestroyQNearFieldManager(void* ptr);

#ifdef __cplusplus
}
#endif