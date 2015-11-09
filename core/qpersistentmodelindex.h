#ifdef __cplusplus
extern "C" {
#endif

void* QPersistentModelIndex_NewQPersistentModelIndex3(void* other);
int QPersistentModelIndex_Column(void* ptr);
int QPersistentModelIndex_IsValid(void* ptr);
int QPersistentModelIndex_Row(void* ptr);
void* QPersistentModelIndex_NewQPersistentModelIndex4(void* other);
void* QPersistentModelIndex_NewQPersistentModelIndex(void* index);
void* QPersistentModelIndex_Child(void* ptr, int row, int column);
void* QPersistentModelIndex_Data(void* ptr, int role);
int QPersistentModelIndex_Flags(void* ptr);
void* QPersistentModelIndex_Model(void* ptr);
void* QPersistentModelIndex_Parent(void* ptr);
void* QPersistentModelIndex_Sibling(void* ptr, int row, int column);
void QPersistentModelIndex_Swap(void* ptr, void* other);

#ifdef __cplusplus
}
#endif