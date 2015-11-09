#ifdef __cplusplus
extern "C" {
#endif

void* QTableWidgetSelectionRange_NewQTableWidgetSelectionRange();
void* QTableWidgetSelectionRange_NewQTableWidgetSelectionRange3(void* other);
void* QTableWidgetSelectionRange_NewQTableWidgetSelectionRange2(int top, int left, int bottom, int right);
int QTableWidgetSelectionRange_BottomRow(void* ptr);
int QTableWidgetSelectionRange_ColumnCount(void* ptr);
int QTableWidgetSelectionRange_LeftColumn(void* ptr);
int QTableWidgetSelectionRange_RightColumn(void* ptr);
int QTableWidgetSelectionRange_RowCount(void* ptr);
int QTableWidgetSelectionRange_TopRow(void* ptr);
void QTableWidgetSelectionRange_DestroyQTableWidgetSelectionRange(void* ptr);

#ifdef __cplusplus
}
#endif