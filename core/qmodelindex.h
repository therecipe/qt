#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QModelIndex_NewQModelIndex();
QtObjectPtr QModelIndex_Child(QtObjectPtr ptr, int row, int column);
int QModelIndex_Column(QtObjectPtr ptr);
char* QModelIndex_Data(QtObjectPtr ptr, int role);
int QModelIndex_Flags(QtObjectPtr ptr);
void QModelIndex_InternalPointer(QtObjectPtr ptr);
int QModelIndex_IsValid(QtObjectPtr ptr);
QtObjectPtr QModelIndex_Model(QtObjectPtr ptr);
QtObjectPtr QModelIndex_Parent(QtObjectPtr ptr);
int QModelIndex_Row(QtObjectPtr ptr);
QtObjectPtr QModelIndex_Sibling(QtObjectPtr ptr, int row, int column);

#ifdef __cplusplus
}
#endif