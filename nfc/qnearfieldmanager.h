#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QNearFieldManager_RegisterNdefMessageHandler(QtObjectPtr ptr, QtObjectPtr object, char* method);
int QNearFieldManager_StartTargetDetection(QtObjectPtr ptr);
QtObjectPtr QNearFieldManager_NewQNearFieldManager(QtObjectPtr parent);
int QNearFieldManager_IsAvailable(QtObjectPtr ptr);
int QNearFieldManager_RegisterNdefMessageHandler2(QtObjectPtr ptr, int typeNameFormat, QtObjectPtr ty, QtObjectPtr object, char* method);
int QNearFieldManager_RegisterNdefMessageHandler3(QtObjectPtr ptr, QtObjectPtr filter, QtObjectPtr object, char* method);
void QNearFieldManager_SetTargetAccessModes(QtObjectPtr ptr, int accessModes);
void QNearFieldManager_StopTargetDetection(QtObjectPtr ptr);
int QNearFieldManager_TargetAccessModes(QtObjectPtr ptr);
void QNearFieldManager_ConnectTargetDetected(QtObjectPtr ptr);
void QNearFieldManager_DisconnectTargetDetected(QtObjectPtr ptr);
void QNearFieldManager_ConnectTargetLost(QtObjectPtr ptr);
void QNearFieldManager_DisconnectTargetLost(QtObjectPtr ptr);
int QNearFieldManager_UnregisterNdefMessageHandler(QtObjectPtr ptr, int handlerId);
void QNearFieldManager_DestroyQNearFieldManager(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif