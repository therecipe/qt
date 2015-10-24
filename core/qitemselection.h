#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QItemSelection_NewQItemSelection();
QtObjectPtr QItemSelection_NewQItemSelection2(QtObjectPtr topLeft, QtObjectPtr bottomRight);
int QItemSelection_Contains(QtObjectPtr ptr, QtObjectPtr index);
void QItemSelection_Merge(QtObjectPtr ptr, QtObjectPtr other, int command);
void QItemSelection_Select(QtObjectPtr ptr, QtObjectPtr topLeft, QtObjectPtr bottomRight);
void QItemSelection_QItemSelection_Split(QtObjectPtr ran, QtObjectPtr other, QtObjectPtr result);

#ifdef __cplusplus
}
#endif