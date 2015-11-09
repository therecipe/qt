#ifdef __cplusplus
extern "C" {
#endif

void* QTextTableFormat_NewQTextTableFormat();
int QTextTableFormat_Alignment(void* ptr);
double QTextTableFormat_CellPadding(void* ptr);
double QTextTableFormat_CellSpacing(void* ptr);
void QTextTableFormat_ClearColumnWidthConstraints(void* ptr);
int QTextTableFormat_Columns(void* ptr);
int QTextTableFormat_HeaderRowCount(void* ptr);
int QTextTableFormat_IsValid(void* ptr);
void QTextTableFormat_SetAlignment(void* ptr, int alignment);
void QTextTableFormat_SetCellPadding(void* ptr, double padding);
void QTextTableFormat_SetCellSpacing(void* ptr, double spacing);
void QTextTableFormat_SetHeaderRowCount(void* ptr, int count);

#ifdef __cplusplus
}
#endif