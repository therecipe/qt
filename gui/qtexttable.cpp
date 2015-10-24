#include "qtexttable.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextCursor>
#include <QTextTableFormat>
#include <QTextTable>
#include "_cgo_export.h"

class MyQTextTable: public QTextTable {
public:
};

void QTextTable_InsertColumns(QtObjectPtr ptr, int index, int columns){
	static_cast<QTextTable*>(ptr)->insertColumns(index, columns);
}

void QTextTable_InsertRows(QtObjectPtr ptr, int index, int rows){
	static_cast<QTextTable*>(ptr)->insertRows(index, rows);
}

void QTextTable_RemoveColumns(QtObjectPtr ptr, int index, int columns){
	static_cast<QTextTable*>(ptr)->removeColumns(index, columns);
}

void QTextTable_RemoveRows(QtObjectPtr ptr, int index, int rows){
	static_cast<QTextTable*>(ptr)->removeRows(index, rows);
}

void QTextTable_Resize(QtObjectPtr ptr, int rows, int columns){
	static_cast<QTextTable*>(ptr)->resize(rows, columns);
}

void QTextTable_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextTable*>(ptr)->setFormat(*static_cast<QTextTableFormat*>(format));
}

void QTextTable_AppendColumns(QtObjectPtr ptr, int count){
	static_cast<QTextTable*>(ptr)->appendColumns(count);
}

void QTextTable_AppendRows(QtObjectPtr ptr, int count){
	static_cast<QTextTable*>(ptr)->appendRows(count);
}

int QTextTable_Columns(QtObjectPtr ptr){
	return static_cast<QTextTable*>(ptr)->columns();
}

void QTextTable_MergeCells2(QtObjectPtr ptr, QtObjectPtr cursor){
	static_cast<QTextTable*>(ptr)->mergeCells(*static_cast<QTextCursor*>(cursor));
}

void QTextTable_MergeCells(QtObjectPtr ptr, int row, int column, int numRows, int numCols){
	static_cast<QTextTable*>(ptr)->mergeCells(row, column, numRows, numCols);
}

int QTextTable_Rows(QtObjectPtr ptr){
	return static_cast<QTextTable*>(ptr)->rows();
}

void QTextTable_SplitCell(QtObjectPtr ptr, int row, int column, int numRows, int numCols){
	static_cast<QTextTable*>(ptr)->splitCell(row, column, numRows, numCols);
}

