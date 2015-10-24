#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QMediaRecorderControl_ConnectActualLocationChanged(QtObjectPtr ptr);
void QMediaRecorderControl_DisconnectActualLocationChanged(QtObjectPtr ptr);
void QMediaRecorderControl_ApplySettings(QtObjectPtr ptr);
void QMediaRecorderControl_ConnectError(QtObjectPtr ptr);
void QMediaRecorderControl_DisconnectError(QtObjectPtr ptr);
int QMediaRecorderControl_IsMuted(QtObjectPtr ptr);
void QMediaRecorderControl_ConnectMutedChanged(QtObjectPtr ptr);
void QMediaRecorderControl_DisconnectMutedChanged(QtObjectPtr ptr);
char* QMediaRecorderControl_OutputLocation(QtObjectPtr ptr);
void QMediaRecorderControl_SetMuted(QtObjectPtr ptr, int muted);
int QMediaRecorderControl_SetOutputLocation(QtObjectPtr ptr, char* location);
void QMediaRecorderControl_SetState(QtObjectPtr ptr, int state);
void QMediaRecorderControl_ConnectStateChanged(QtObjectPtr ptr);
void QMediaRecorderControl_DisconnectStateChanged(QtObjectPtr ptr);
int QMediaRecorderControl_Status(QtObjectPtr ptr);
void QMediaRecorderControl_ConnectStatusChanged(QtObjectPtr ptr);
void QMediaRecorderControl_DisconnectStatusChanged(QtObjectPtr ptr);
void QMediaRecorderControl_DestroyQMediaRecorderControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif