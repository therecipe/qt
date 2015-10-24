#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QItemSelectionModel_NewQItemSelectionModel(QtObjectPtr model);
QtObjectPtr QItemSelectionModel_NewQItemSelectionModel2(QtObjectPtr model, QtObjectPtr parent);
void QItemSelectionModel_Clear(QtObjectPtr ptr);
void QItemSelectionModel_ClearCurrentIndex(QtObjectPtr ptr);
void QItemSelectionModel_ClearSelection(QtObjectPtr ptr);
int QItemSelectionModel_ColumnIntersectsSelection(QtObjectPtr ptr, int column, QtObjectPtr parent);
void QItemSelectionModel_ConnectCurrentChanged(QtObjectPtr ptr);
void QItemSelectionModel_DisconnectCurrentChanged(QtObjectPtr ptr);
void QItemSelectionModel_ConnectCurrentColumnChanged(QtObjectPtr ptr);
void QItemSelectionModel_DisconnectCurrentColumnChanged(QtObjectPtr ptr);
QtObjectPtr QItemSelectionModel_CurrentIndex(QtObjectPtr ptr);
void QItemSelectionModel_ConnectCurrentRowChanged(QtObjectPtr ptr);
void QItemSelectionModel_DisconnectCurrentRowChanged(QtObjectPtr ptr);
int QItemSelectionModel_HasSelection(QtObjectPtr ptr);
int QItemSelectionModel_IsColumnSelected(QtObjectPtr ptr, int column, QtObjectPtr parent);
int QItemSelectionModel_IsRowSelected(QtObjectPtr ptr, int row, QtObjectPtr parent);
int QItemSelectionModel_IsSelected(QtObjectPtr ptr, QtObjectPtr index);
QtObjectPtr QItemSelectionModel_Model2(QtObjectPtr ptr);
QtObjectPtr QItemSelectionModel_Model(QtObjectPtr ptr);
void QItemSelectionModel_ConnectModelChanged(QtObjectPtr ptr);
void QItemSelectionModel_DisconnectModelChanged(QtObjectPtr ptr);
void QItemSelectionModel_Reset(QtObjectPtr ptr);
int QItemSelectionModel_RowIntersectsSelection(QtObjectPtr ptr, int row, QtObjectPtr parent);
void QItemSelectionModel_Select2(QtObjectPtr ptr, QtObjectPtr selection, int command);
void QItemSelectionModel_Select(QtObjectPtr ptr, QtObjectPtr index, int command);
void QItemSelectionModel_SetCurrentIndex(QtObjectPtr ptr, QtObjectPtr index, int command);
void QItemSelectionModel_SetModel(QtObjectPtr ptr, QtObjectPtr model);
void QItemSelectionModel_DestroyQItemSelectionModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif