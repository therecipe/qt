#ifdef __cplusplus
extern "C" {
#endif

void* QItemSelectionModel_NewQItemSelectionModel(void* model);
void* QItemSelectionModel_NewQItemSelectionModel2(void* model, void* parent);
void QItemSelectionModel_Clear(void* ptr);
void QItemSelectionModel_ClearCurrentIndex(void* ptr);
void QItemSelectionModel_ClearSelection(void* ptr);
int QItemSelectionModel_ColumnIntersectsSelection(void* ptr, int column, void* parent);
void QItemSelectionModel_ConnectCurrentChanged(void* ptr);
void QItemSelectionModel_DisconnectCurrentChanged(void* ptr);
void QItemSelectionModel_ConnectCurrentColumnChanged(void* ptr);
void QItemSelectionModel_DisconnectCurrentColumnChanged(void* ptr);
void* QItemSelectionModel_CurrentIndex(void* ptr);
void QItemSelectionModel_ConnectCurrentRowChanged(void* ptr);
void QItemSelectionModel_DisconnectCurrentRowChanged(void* ptr);
int QItemSelectionModel_HasSelection(void* ptr);
int QItemSelectionModel_IsColumnSelected(void* ptr, int column, void* parent);
int QItemSelectionModel_IsRowSelected(void* ptr, int row, void* parent);
int QItemSelectionModel_IsSelected(void* ptr, void* index);
void* QItemSelectionModel_Model2(void* ptr);
void* QItemSelectionModel_Model(void* ptr);
void QItemSelectionModel_ConnectModelChanged(void* ptr);
void QItemSelectionModel_DisconnectModelChanged(void* ptr);
void QItemSelectionModel_Reset(void* ptr);
int QItemSelectionModel_RowIntersectsSelection(void* ptr, int row, void* parent);
void QItemSelectionModel_Select2(void* ptr, void* selection, int command);
void QItemSelectionModel_Select(void* ptr, void* index, int command);
void QItemSelectionModel_SetCurrentIndex(void* ptr, void* index, int command);
void QItemSelectionModel_SetModel(void* ptr, void* model);
void QItemSelectionModel_DestroyQItemSelectionModel(void* ptr);

#ifdef __cplusplus
}
#endif