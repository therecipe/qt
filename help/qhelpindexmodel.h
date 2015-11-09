#ifdef __cplusplus
extern "C" {
#endif

void QHelpIndexModel_CreateIndex(void* ptr, char* customFilterName);
void* QHelpIndexModel_Filter(void* ptr, char* filter, char* wildcard);
void QHelpIndexModel_ConnectIndexCreated(void* ptr);
void QHelpIndexModel_DisconnectIndexCreated(void* ptr);
void QHelpIndexModel_ConnectIndexCreationStarted(void* ptr);
void QHelpIndexModel_DisconnectIndexCreationStarted(void* ptr);
int QHelpIndexModel_IsCreatingIndex(void* ptr);

#ifdef __cplusplus
}
#endif