#ifdef __cplusplus
extern "C" {
#endif

void* QStringListModel_Data(void* ptr, void* index, int role);
int QStringListModel_Flags(void* ptr, void* index);
int QStringListModel_InsertRows(void* ptr, int row, int count, void* parent);
int QStringListModel_RemoveRows(void* ptr, int row, int count, void* parent);
int QStringListModel_RowCount(void* ptr, void* parent);
int QStringListModel_SetData(void* ptr, void* index, void* value, int role);
void QStringListModel_SetStringList(void* ptr, char* strin);
void* QStringListModel_Sibling(void* ptr, int row, int column, void* idx);
void QStringListModel_Sort(void* ptr, int column, int order);
char* QStringListModel_StringList(void* ptr);
int QStringListModel_SupportedDropActions(void* ptr);

#ifdef __cplusplus
}
#endif