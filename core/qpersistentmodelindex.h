#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPersistentModelIndex_NewQPersistentModelIndex3(QtObjectPtr other);
int QPersistentModelIndex_Column(QtObjectPtr ptr);
int QPersistentModelIndex_IsValid(QtObjectPtr ptr);
int QPersistentModelIndex_Row(QtObjectPtr ptr);
QtObjectPtr QPersistentModelIndex_NewQPersistentModelIndex4(QtObjectPtr other);
QtObjectPtr QPersistentModelIndex_NewQPersistentModelIndex(QtObjectPtr index);
QtObjectPtr QPersistentModelIndex_Child(QtObjectPtr ptr, int row, int column);
char* QPersistentModelIndex_Data(QtObjectPtr ptr, int role);
int QPersistentModelIndex_Flags(QtObjectPtr ptr);
QtObjectPtr QPersistentModelIndex_Model(QtObjectPtr ptr);
QtObjectPtr QPersistentModelIndex_Parent(QtObjectPtr ptr);
QtObjectPtr QPersistentModelIndex_Sibling(QtObjectPtr ptr, int row, int column);
void QPersistentModelIndex_Swap(QtObjectPtr ptr, QtObjectPtr other);

#ifdef __cplusplus
}
#endif