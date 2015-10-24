#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QHelpEngineCore_AutoSaveFilter(QtObjectPtr ptr);
char* QHelpEngineCore_CollectionFile(QtObjectPtr ptr);
char* QHelpEngineCore_CurrentFilter(QtObjectPtr ptr);
void QHelpEngineCore_SetAutoSaveFilter(QtObjectPtr ptr, int save);
void QHelpEngineCore_SetCollectionFile(QtObjectPtr ptr, char* fileName);
void QHelpEngineCore_SetCurrentFilter(QtObjectPtr ptr, char* filterName);
QtObjectPtr QHelpEngineCore_NewQHelpEngineCore(char* collectionFile, QtObjectPtr parent);
int QHelpEngineCore_AddCustomFilter(QtObjectPtr ptr, char* filterName, char* attributes);
int QHelpEngineCore_CopyCollectionFile(QtObjectPtr ptr, char* fileName);
void QHelpEngineCore_ConnectCurrentFilterChanged(QtObjectPtr ptr);
void QHelpEngineCore_DisconnectCurrentFilterChanged(QtObjectPtr ptr);
char* QHelpEngineCore_CustomFilters(QtObjectPtr ptr);
char* QHelpEngineCore_CustomValue(QtObjectPtr ptr, char* key, char* defaultValue);
char* QHelpEngineCore_DocumentationFileName(QtObjectPtr ptr, char* namespaceName);
char* QHelpEngineCore_Error(QtObjectPtr ptr);
char* QHelpEngineCore_FilterAttributes(QtObjectPtr ptr);
char* QHelpEngineCore_FilterAttributes2(QtObjectPtr ptr, char* filterName);
char* QHelpEngineCore_FindFile(QtObjectPtr ptr, char* url);
char* QHelpEngineCore_QHelpEngineCore_MetaData(char* documentationFileName, char* name);
char* QHelpEngineCore_QHelpEngineCore_NamespaceName(char* documentationFileName);
void QHelpEngineCore_ConnectReadersAboutToBeInvalidated(QtObjectPtr ptr);
void QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(QtObjectPtr ptr);
int QHelpEngineCore_RegisterDocumentation(QtObjectPtr ptr, char* documentationFileName);
char* QHelpEngineCore_RegisteredDocumentations(QtObjectPtr ptr);
int QHelpEngineCore_RemoveCustomFilter(QtObjectPtr ptr, char* filterName);
int QHelpEngineCore_RemoveCustomValue(QtObjectPtr ptr, char* key);
int QHelpEngineCore_SetCustomValue(QtObjectPtr ptr, char* key, char* value);
int QHelpEngineCore_SetupData(QtObjectPtr ptr);
void QHelpEngineCore_ConnectSetupFinished(QtObjectPtr ptr);
void QHelpEngineCore_DisconnectSetupFinished(QtObjectPtr ptr);
void QHelpEngineCore_ConnectSetupStarted(QtObjectPtr ptr);
void QHelpEngineCore_DisconnectSetupStarted(QtObjectPtr ptr);
int QHelpEngineCore_UnregisterDocumentation(QtObjectPtr ptr, char* namespaceName);
void QHelpEngineCore_ConnectWarning(QtObjectPtr ptr);
void QHelpEngineCore_DisconnectWarning(QtObjectPtr ptr);
void QHelpEngineCore_DestroyQHelpEngineCore(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif