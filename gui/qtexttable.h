#ifdef __cplusplus
extern "C" {
#endif

void QTextTable_InsertColumns(void* ptr, int index, int columns);
void QTextTable_InsertRows(void* ptr, int index, int rows);
void QTextTable_RemoveColumns(void* ptr, int index, int columns);
void QTextTable_RemoveRows(void* ptr, int index, int rows);
void QTextTable_Resize(void* ptr, int rows, int columns);
void QTextTable_SetFormat(void* ptr, void* format);
void QTextTable_AppendColumns(void* ptr, int count);
void QTextTable_AppendRows(void* ptr, int count);
int QTextTable_Columns(void* ptr);
void QTextTable_MergeCells2(void* ptr, void* cursor);
void QTextTable_MergeCells(void* ptr, int row, int column, int numRows, int numCols);
int QTextTable_Rows(void* ptr);
void QTextTable_SplitCell(void* ptr, int row, int column, int numRows, int numCols);

#ifdef __cplusplus
}
#endif