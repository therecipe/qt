#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTableWidgetSelectionRange_NewQTableWidgetSelectionRange();
QtObjectPtr QTableWidgetSelectionRange_NewQTableWidgetSelectionRange3(QtObjectPtr other);
QtObjectPtr QTableWidgetSelectionRange_NewQTableWidgetSelectionRange2(int top, int left, int bottom, int right);
int QTableWidgetSelectionRange_BottomRow(QtObjectPtr ptr);
int QTableWidgetSelectionRange_ColumnCount(QtObjectPtr ptr);
int QTableWidgetSelectionRange_LeftColumn(QtObjectPtr ptr);
int QTableWidgetSelectionRange_RightColumn(QtObjectPtr ptr);
int QTableWidgetSelectionRange_RowCount(QtObjectPtr ptr);
int QTableWidgetSelectionRange_TopRow(QtObjectPtr ptr);
void QTableWidgetSelectionRange_DestroyQTableWidgetSelectionRange(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif