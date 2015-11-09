#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractListModel_Index(void* ptr, int row, int column, void* parent);
int QAbstractListModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent);
int QAbstractListModel_Flags(void* ptr, void* index);
void* QAbstractListModel_Sibling(void* ptr, int row, int column, void* idx);
void QAbstractListModel_DestroyQAbstractListModel(void* ptr);

#ifdef __cplusplus
}
#endif