#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAbstractTableModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent);
int QAbstractTableModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent);
int QAbstractTableModel_Flags(QtObjectPtr ptr, QtObjectPtr index);
QtObjectPtr QAbstractTableModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx);
void QAbstractTableModel_DestroyQAbstractTableModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif