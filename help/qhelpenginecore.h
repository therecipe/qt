#ifdef __cplusplus
extern "C" {
#endif

int QHelpEngineCore_AutoSaveFilter(void* ptr);
char* QHelpEngineCore_CollectionFile(void* ptr);
char* QHelpEngineCore_CurrentFilter(void* ptr);
void QHelpEngineCore_SetAutoSaveFilter(void* ptr, int save);
void QHelpEngineCore_SetCollectionFile(void* ptr, char* fileName);
void QHelpEngineCore_SetCurrentFilter(void* ptr, char* filterName);
void* QHelpEngineCore_NewQHelpEngineCore(char* collectionFile, void* parent);
int QHelpEngineCore_AddCustomFilter(void* ptr, char* filterName, char* attributes);
int QHelpEngineCore_CopyCollectionFile(void* ptr, char* fileName);
void QHelpEngineCore_ConnectCurrentFilterChanged(void* ptr);
void QHelpEngineCore_DisconnectCurrentFilterChanged(void* ptr);
char* QHelpEngineCore_CustomFilters(void* ptr);
void* QHelpEngineCore_CustomValue(void* ptr, char* key, void* defaultValue);
char* QHelpEngineCore_DocumentationFileName(void* ptr, char* namespaceName);
char* QHelpEngineCore_Error(void* ptr);
void* QHelpEngineCore_FileData(void* ptr, void* url);
char* QHelpEngineCore_FilterAttributes(void* ptr);
char* QHelpEngineCore_FilterAttributes2(void* ptr, char* filterName);
void* QHelpEngineCore_QHelpEngineCore_MetaData(char* documentationFileName, char* name);
char* QHelpEngineCore_QHelpEngineCore_NamespaceName(char* documentationFileName);
void QHelpEngineCore_ConnectReadersAboutToBeInvalidated(void* ptr);
void QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(void* ptr);
int QHelpEngineCore_RegisterDocumentation(void* ptr, char* documentationFileName);
char* QHelpEngineCore_RegisteredDocumentations(void* ptr);
int QHelpEngineCore_RemoveCustomFilter(void* ptr, char* filterName);
int QHelpEngineCore_RemoveCustomValue(void* ptr, char* key);
int QHelpEngineCore_SetCustomValue(void* ptr, char* key, void* value);
int QHelpEngineCore_SetupData(void* ptr);
void QHelpEngineCore_ConnectSetupFinished(void* ptr);
void QHelpEngineCore_DisconnectSetupFinished(void* ptr);
void QHelpEngineCore_ConnectSetupStarted(void* ptr);
void QHelpEngineCore_DisconnectSetupStarted(void* ptr);
int QHelpEngineCore_UnregisterDocumentation(void* ptr, char* namespaceName);
void QHelpEngineCore_ConnectWarning(void* ptr);
void QHelpEngineCore_DisconnectWarning(void* ptr);
void QHelpEngineCore_DestroyQHelpEngineCore(void* ptr);

#ifdef __cplusplus
}
#endif