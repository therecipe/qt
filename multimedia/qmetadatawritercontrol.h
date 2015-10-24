#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QMetaDataWriterControl_AvailableMetaData(QtObjectPtr ptr);
int QMetaDataWriterControl_IsMetaDataAvailable(QtObjectPtr ptr);
int QMetaDataWriterControl_IsWritable(QtObjectPtr ptr);
char* QMetaDataWriterControl_MetaData(QtObjectPtr ptr, char* key);
void QMetaDataWriterControl_ConnectMetaDataAvailableChanged(QtObjectPtr ptr);
void QMetaDataWriterControl_DisconnectMetaDataAvailableChanged(QtObjectPtr ptr);
void QMetaDataWriterControl_ConnectMetaDataChanged(QtObjectPtr ptr);
void QMetaDataWriterControl_DisconnectMetaDataChanged(QtObjectPtr ptr);
void QMetaDataWriterControl_SetMetaData(QtObjectPtr ptr, char* key, char* value);
void QMetaDataWriterControl_ConnectWritableChanged(QtObjectPtr ptr);
void QMetaDataWriterControl_DisconnectWritableChanged(QtObjectPtr ptr);
void QMetaDataWriterControl_DestroyQMetaDataWriterControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif