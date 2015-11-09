#ifdef __cplusplus
extern "C" {
#endif

char* QMetaDataWriterControl_AvailableMetaData(void* ptr);
int QMetaDataWriterControl_IsMetaDataAvailable(void* ptr);
int QMetaDataWriterControl_IsWritable(void* ptr);
void* QMetaDataWriterControl_MetaData(void* ptr, char* key);
void QMetaDataWriterControl_ConnectMetaDataAvailableChanged(void* ptr);
void QMetaDataWriterControl_DisconnectMetaDataAvailableChanged(void* ptr);
void QMetaDataWriterControl_ConnectMetaDataChanged(void* ptr);
void QMetaDataWriterControl_DisconnectMetaDataChanged(void* ptr);
void QMetaDataWriterControl_SetMetaData(void* ptr, char* key, void* value);
void QMetaDataWriterControl_ConnectWritableChanged(void* ptr);
void QMetaDataWriterControl_DisconnectWritableChanged(void* ptr);
void QMetaDataWriterControl_DestroyQMetaDataWriterControl(void* ptr);

#ifdef __cplusplus
}
#endif