#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextTableCell_NewQTextTableCell();
QtObjectPtr QTextTableCell_NewQTextTableCell2(QtObjectPtr other);
int QTextTableCell_Column(QtObjectPtr ptr);
int QTextTableCell_ColumnSpan(QtObjectPtr ptr);
int QTextTableCell_IsValid(QtObjectPtr ptr);
int QTextTableCell_Row(QtObjectPtr ptr);
int QTextTableCell_RowSpan(QtObjectPtr ptr);
void QTextTableCell_SetFormat(QtObjectPtr ptr, QtObjectPtr format);
int QTextTableCell_TableCellFormatIndex(QtObjectPtr ptr);
void QTextTableCell_DestroyQTextTableCell(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif