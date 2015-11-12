#include "qtexttable.h"
#include <QModelIndex>
#include <QTextTableFormat>
#include <QTextCursor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTextTable>
#include "_cgo_export.h"

class MyQTextTable: public QTextTable {
public:
};

void QTextTable_InsertColumns(void* ptr, int index, int columns){
	static_cast<QTextTable*>(ptr)->insertColumns(index, columns);
}

void QTextTable_InsertRows(void* ptr, int index, int rows){
	static_cast<QTextTable*>(ptr)->insertRows(index, rows);
}

void QTextTable_RemoveColumns(void* ptr, int index, int columns){
	static_cast<QTextTable*>(ptr)->removeColumns(index, columns);
}

void QTextTable_RemoveRows(void* ptr, int index, int rows){
	static_cast<QTextTable*>(ptr)->removeRows(index, rows);
}

void QTextTable_Resize(void* ptr, int rows, int columns){
	static_cast<QTextTable*>(ptr)->resize(rows, columns);
}

void QTextTable_SetFormat(void* ptr, void* format){
	static_cast<QTextTable*>(ptr)->setFormat(*static_cast<QTextTableFormat*>(format));
}

void QTextTable_AppendColumns(void* ptr, int count){
	static_cast<QTextTable*>(ptr)->appendColumns(count);
}

void QTextTable_AppendRows(void* ptr, int count){
	static_cast<QTextTable*>(ptr)->appendRows(count);
}

int QTextTable_Columns(void* ptr){
	return static_cast<QTextTable*>(ptr)->columns();
}

void QTextTable_MergeCells2(void* ptr, void* cursor){
	static_cast<QTextTable*>(ptr)->mergeCells(*static_cast<QTextCursor*>(cursor));
}

void QTextTable_MergeCells(void* ptr, int row, int column, int numRows, int numCols){
	static_cast<QTextTable*>(ptr)->mergeCells(row, column, numRows, numCols);
}

int QTextTable_Rows(void* ptr){
	return static_cast<QTextTable*>(ptr)->rows();
}

void QTextTable_SplitCell(void* ptr, int row, int column, int numRows, int numCols){
	static_cast<QTextTable*>(ptr)->splitCell(row, column, numRows, numCols);
}

