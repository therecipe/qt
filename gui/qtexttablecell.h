#ifdef __cplusplus
extern "C" {
#endif

void* QTextTableCell_NewQTextTableCell();
void* QTextTableCell_NewQTextTableCell2(void* other);
int QTextTableCell_Column(void* ptr);
int QTextTableCell_ColumnSpan(void* ptr);
int QTextTableCell_IsValid(void* ptr);
int QTextTableCell_Row(void* ptr);
int QTextTableCell_RowSpan(void* ptr);
void QTextTableCell_SetFormat(void* ptr, void* format);
int QTextTableCell_TableCellFormatIndex(void* ptr);
void QTextTableCell_DestroyQTextTableCell(void* ptr);

#ifdef __cplusplus
}
#endif