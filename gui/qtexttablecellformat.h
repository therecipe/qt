#ifdef __cplusplus
extern "C" {
#endif

void* QTextTableCellFormat_NewQTextTableCellFormat();
double QTextTableCellFormat_BottomPadding(void* ptr);
int QTextTableCellFormat_IsValid(void* ptr);
double QTextTableCellFormat_LeftPadding(void* ptr);
double QTextTableCellFormat_RightPadding(void* ptr);
void QTextTableCellFormat_SetBottomPadding(void* ptr, double padding);
void QTextTableCellFormat_SetLeftPadding(void* ptr, double padding);
void QTextTableCellFormat_SetPadding(void* ptr, double padding);
void QTextTableCellFormat_SetRightPadding(void* ptr, double padding);
void QTextTableCellFormat_SetTopPadding(void* ptr, double padding);
double QTextTableCellFormat_TopPadding(void* ptr);

#ifdef __cplusplus
}
#endif