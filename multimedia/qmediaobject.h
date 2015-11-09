#ifdef __cplusplus
extern "C" {
#endif

int QMediaObject_NotifyInterval(void* ptr);
void QMediaObject_SetNotifyInterval(void* ptr, int milliSeconds);
void QMediaObject_ConnectAvailabilityChanged(void* ptr);
void QMediaObject_DisconnectAvailabilityChanged(void* ptr);
char* QMediaObject_AvailableMetaData(void* ptr);
int QMediaObject_Bind(void* ptr, void* object);
int QMediaObject_IsAvailable(void* ptr);
int QMediaObject_IsMetaDataAvailable(void* ptr);
void* QMediaObject_MetaData(void* ptr, char* key);
void QMediaObject_ConnectMetaDataAvailableChanged(void* ptr);
void QMediaObject_DisconnectMetaDataAvailableChanged(void* ptr);
void QMediaObject_ConnectMetaDataChanged(void* ptr);
void QMediaObject_DisconnectMetaDataChanged(void* ptr);
void QMediaObject_ConnectNotifyIntervalChanged(void* ptr);
void QMediaObject_DisconnectNotifyIntervalChanged(void* ptr);
void* QMediaObject_Service(void* ptr);
void QMediaObject_Unbind(void* ptr, void* object);
void QMediaObject_DestroyQMediaObject(void* ptr);

#ifdef __cplusplus
}
#endif