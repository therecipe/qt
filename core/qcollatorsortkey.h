#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QCollatorSortKey_NewQCollatorSortKey(QtObjectPtr other);
void QCollatorSortKey_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QCollatorSortKey_DestroyQCollatorSortKey(QtObjectPtr ptr);
int QCollatorSortKey_Compare(QtObjectPtr ptr, QtObjectPtr otherKey);

#ifdef __cplusplus
}
#endif