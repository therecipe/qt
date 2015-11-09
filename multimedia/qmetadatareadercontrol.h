#ifdef __cplusplus
extern "C" {
#endif

char* QMetaDataReaderControl_AvailableMetaData(void* ptr);
int QMetaDataReaderControl_IsMetaDataAvailable(void* ptr);
void* QMetaDataReaderControl_MetaData(void* ptr, char* key);
void QMetaDataReaderControl_ConnectMetaDataAvailableChanged(void* ptr);
void QMetaDataReaderControl_DisconnectMetaDataAvailableChanged(void* ptr);
void QMetaDataReaderControl_ConnectMetaDataChanged(void* ptr);
void QMetaDataReaderControl_DisconnectMetaDataChanged(void* ptr);
void QMetaDataReaderControl_DestroyQMetaDataReaderControl(void* ptr);

#ifdef __cplusplus
}
#endif