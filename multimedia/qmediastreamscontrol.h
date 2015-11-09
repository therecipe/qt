#ifdef __cplusplus
extern "C" {
#endif

void QMediaStreamsControl_ConnectActiveStreamsChanged(void* ptr);
void QMediaStreamsControl_DisconnectActiveStreamsChanged(void* ptr);
int QMediaStreamsControl_IsActive(void* ptr, int stream);
void* QMediaStreamsControl_MetaData(void* ptr, int stream, char* key);
void QMediaStreamsControl_SetActive(void* ptr, int stream, int state);
int QMediaStreamsControl_StreamCount(void* ptr);
int QMediaStreamsControl_StreamType(void* ptr, int stream);
void QMediaStreamsControl_ConnectStreamsChanged(void* ptr);
void QMediaStreamsControl_DisconnectStreamsChanged(void* ptr);
void QMediaStreamsControl_DestroyQMediaStreamsControl(void* ptr);

#ifdef __cplusplus
}
#endif