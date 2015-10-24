#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QHelpContentModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent);
QtObjectPtr QHelpContentModel_ContentItemAt(QtObjectPtr ptr, QtObjectPtr index);
void QHelpContentModel_ConnectContentsCreated(QtObjectPtr ptr);
void QHelpContentModel_DisconnectContentsCreated(QtObjectPtr ptr);
void QHelpContentModel_ConnectContentsCreationStarted(QtObjectPtr ptr);
void QHelpContentModel_DisconnectContentsCreationStarted(QtObjectPtr ptr);
void QHelpContentModel_CreateContents(QtObjectPtr ptr, char* customFilterName);
char* QHelpContentModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role);
QtObjectPtr QHelpContentModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent);
int QHelpContentModel_IsCreatingContents(QtObjectPtr ptr);
QtObjectPtr QHelpContentModel_Parent(QtObjectPtr ptr, QtObjectPtr index);
int QHelpContentModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent);
void QHelpContentModel_DestroyQHelpContentModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif