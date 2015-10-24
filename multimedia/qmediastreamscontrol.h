#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QMediaStreamsControl_ConnectActiveStreamsChanged(QtObjectPtr ptr);
void QMediaStreamsControl_DisconnectActiveStreamsChanged(QtObjectPtr ptr);
int QMediaStreamsControl_IsActive(QtObjectPtr ptr, int stream);
char* QMediaStreamsControl_MetaData(QtObjectPtr ptr, int stream, char* key);
void QMediaStreamsControl_SetActive(QtObjectPtr ptr, int stream, int state);
int QMediaStreamsControl_StreamCount(QtObjectPtr ptr);
int QMediaStreamsControl_StreamType(QtObjectPtr ptr, int stream);
void QMediaStreamsControl_ConnectStreamsChanged(QtObjectPtr ptr);
void QMediaStreamsControl_DisconnectStreamsChanged(QtObjectPtr ptr);
void QMediaStreamsControl_DestroyQMediaStreamsControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif