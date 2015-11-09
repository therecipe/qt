#ifdef __cplusplus
extern "C" {
#endif

int QHelpContentModel_ColumnCount(void* ptr, void* parent);
void* QHelpContentModel_ContentItemAt(void* ptr, void* index);
void QHelpContentModel_ConnectContentsCreated(void* ptr);
void QHelpContentModel_DisconnectContentsCreated(void* ptr);
void QHelpContentModel_ConnectContentsCreationStarted(void* ptr);
void QHelpContentModel_DisconnectContentsCreationStarted(void* ptr);
void QHelpContentModel_CreateContents(void* ptr, char* customFilterName);
void* QHelpContentModel_Data(void* ptr, void* index, int role);
void* QHelpContentModel_Index(void* ptr, int row, int column, void* parent);
int QHelpContentModel_IsCreatingContents(void* ptr);
void* QHelpContentModel_Parent(void* ptr, void* index);
int QHelpContentModel_RowCount(void* ptr, void* parent);
void QHelpContentModel_DestroyQHelpContentModel(void* ptr);

#ifdef __cplusplus
}
#endif