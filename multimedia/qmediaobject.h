#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMediaObject_NotifyInterval(QtObjectPtr ptr);
void QMediaObject_SetNotifyInterval(QtObjectPtr ptr, int milliSeconds);
void QMediaObject_ConnectAvailabilityChanged(QtObjectPtr ptr);
void QMediaObject_DisconnectAvailabilityChanged(QtObjectPtr ptr);
char* QMediaObject_AvailableMetaData(QtObjectPtr ptr);
int QMediaObject_Bind(QtObjectPtr ptr, QtObjectPtr object);
int QMediaObject_IsAvailable(QtObjectPtr ptr);
int QMediaObject_IsMetaDataAvailable(QtObjectPtr ptr);
char* QMediaObject_MetaData(QtObjectPtr ptr, char* key);
void QMediaObject_ConnectMetaDataAvailableChanged(QtObjectPtr ptr);
void QMediaObject_DisconnectMetaDataAvailableChanged(QtObjectPtr ptr);
void QMediaObject_ConnectMetaDataChanged(QtObjectPtr ptr);
void QMediaObject_DisconnectMetaDataChanged(QtObjectPtr ptr);
void QMediaObject_ConnectNotifyIntervalChanged(QtObjectPtr ptr);
void QMediaObject_DisconnectNotifyIntervalChanged(QtObjectPtr ptr);
QtObjectPtr QMediaObject_Service(QtObjectPtr ptr);
void QMediaObject_Unbind(QtObjectPtr ptr, QtObjectPtr object);
void QMediaObject_DestroyQMediaObject(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif