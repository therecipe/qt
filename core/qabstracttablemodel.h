#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractTableModel_Index(void* ptr, int row, int column, void* parent);
int QAbstractTableModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent);
int QAbstractTableModel_Flags(void* ptr, void* index);
void* QAbstractTableModel_Sibling(void* ptr, int row, int column, void* idx);
void QAbstractTableModel_DestroyQAbstractTableModel(void* ptr);

#ifdef __cplusplus
}
#endif