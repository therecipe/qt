#ifdef __cplusplus
extern "C" {
#endif

void* QModelIndex_NewQModelIndex();
void* QModelIndex_Child(void* ptr, int row, int column);
int QModelIndex_Column(void* ptr);
void* QModelIndex_Data(void* ptr, int role);
int QModelIndex_Flags(void* ptr);
void* QModelIndex_InternalPointer(void* ptr);
int QModelIndex_IsValid(void* ptr);
void* QModelIndex_Model(void* ptr);
void* QModelIndex_Parent(void* ptr);
int QModelIndex_Row(void* ptr);
void* QModelIndex_Sibling(void* ptr, int row, int column);

#ifdef __cplusplus
}
#endif