#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAbstractListModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent);
int QAbstractListModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent);
int QAbstractListModel_Flags(QtObjectPtr ptr, QtObjectPtr index);
QtObjectPtr QAbstractListModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx);
void QAbstractListModel_DestroyQAbstractListModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif