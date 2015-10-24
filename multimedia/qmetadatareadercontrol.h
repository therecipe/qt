#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QMetaDataReaderControl_AvailableMetaData(QtObjectPtr ptr);
int QMetaDataReaderControl_IsMetaDataAvailable(QtObjectPtr ptr);
char* QMetaDataReaderControl_MetaData(QtObjectPtr ptr, char* key);
void QMetaDataReaderControl_ConnectMetaDataAvailableChanged(QtObjectPtr ptr);
void QMetaDataReaderControl_DisconnectMetaDataAvailableChanged(QtObjectPtr ptr);
void QMetaDataReaderControl_ConnectMetaDataChanged(QtObjectPtr ptr);
void QMetaDataReaderControl_DisconnectMetaDataChanged(QtObjectPtr ptr);
void QMetaDataReaderControl_DestroyQMetaDataReaderControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif