#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAbstractProxyModel_Buddy(QtObjectPtr ptr, QtObjectPtr index);
int QAbstractProxyModel_CanDropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent);
int QAbstractProxyModel_CanFetchMore(QtObjectPtr ptr, QtObjectPtr parent);
char* QAbstractProxyModel_Data(QtObjectPtr ptr, QtObjectPtr proxyIndex, int role);
int QAbstractProxyModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent);
void QAbstractProxyModel_FetchMore(QtObjectPtr ptr, QtObjectPtr parent);
int QAbstractProxyModel_Flags(QtObjectPtr ptr, QtObjectPtr index);
int QAbstractProxyModel_HasChildren(QtObjectPtr ptr, QtObjectPtr parent);
char* QAbstractProxyModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role);
QtObjectPtr QAbstractProxyModel_MapFromSource(QtObjectPtr ptr, QtObjectPtr sourceIndex);
QtObjectPtr QAbstractProxyModel_MapToSource(QtObjectPtr ptr, QtObjectPtr proxyIndex);
char* QAbstractProxyModel_MimeTypes(QtObjectPtr ptr);
void QAbstractProxyModel_Revert(QtObjectPtr ptr);
int QAbstractProxyModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role);
int QAbstractProxyModel_SetHeaderData(QtObjectPtr ptr, int section, int orientation, char* value, int role);
void QAbstractProxyModel_SetSourceModel(QtObjectPtr ptr, QtObjectPtr sourceModel);
QtObjectPtr QAbstractProxyModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx);
void QAbstractProxyModel_Sort(QtObjectPtr ptr, int column, int order);
QtObjectPtr QAbstractProxyModel_SourceModel(QtObjectPtr ptr);
void QAbstractProxyModel_ConnectSourceModelChanged(QtObjectPtr ptr);
void QAbstractProxyModel_DisconnectSourceModelChanged(QtObjectPtr ptr);
int QAbstractProxyModel_Submit(QtObjectPtr ptr);
int QAbstractProxyModel_SupportedDragActions(QtObjectPtr ptr);
int QAbstractProxyModel_SupportedDropActions(QtObjectPtr ptr);
void QAbstractProxyModel_DestroyQAbstractProxyModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif