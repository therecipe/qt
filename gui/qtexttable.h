#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QTextTable_InsertColumns(QtObjectPtr ptr, int index, int columns);
void QTextTable_InsertRows(QtObjectPtr ptr, int index, int rows);
void QTextTable_RemoveColumns(QtObjectPtr ptr, int index, int columns);
void QTextTable_RemoveRows(QtObjectPtr ptr, int index, int rows);
void QTextTable_Resize(QtObjectPtr ptr, int rows, int columns);
void QTextTable_SetFormat(QtObjectPtr ptr, QtObjectPtr format);
void QTextTable_AppendColumns(QtObjectPtr ptr, int count);
void QTextTable_AppendRows(QtObjectPtr ptr, int count);
int QTextTable_Columns(QtObjectPtr ptr);
void QTextTable_MergeCells2(QtObjectPtr ptr, QtObjectPtr cursor);
void QTextTable_MergeCells(QtObjectPtr ptr, int row, int column, int numRows, int numCols);
int QTextTable_Rows(QtObjectPtr ptr);
void QTextTable_SplitCell(QtObjectPtr ptr, int row, int column, int numRows, int numCols);

#ifdef __cplusplus
}
#endif