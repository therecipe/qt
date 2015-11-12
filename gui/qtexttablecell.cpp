#include "qtexttablecell.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextTable>
#include <QTextCharFormat>
#include <QTextTableCell>
#include "_cgo_export.h"

class MyQTextTableCell: public QTextTableCell {
public:
};

void* QTextTableCell_NewQTextTableCell(){
	return new QTextTableCell();
}

void* QTextTableCell_NewQTextTableCell2(void* other){
	return new QTextTableCell(*static_cast<QTextTableCell*>(other));
}

int QTextTableCell_Column(void* ptr){
	return static_cast<QTextTableCell*>(ptr)->column();
}

int QTextTableCell_ColumnSpan(void* ptr){
	return static_cast<QTextTableCell*>(ptr)->columnSpan();
}

int QTextTableCell_IsValid(void* ptr){
	return static_cast<QTextTableCell*>(ptr)->isValid();
}

int QTextTableCell_Row(void* ptr){
	return static_cast<QTextTableCell*>(ptr)->row();
}

int QTextTableCell_RowSpan(void* ptr){
	return static_cast<QTextTableCell*>(ptr)->rowSpan();
}

void QTextTableCell_SetFormat(void* ptr, void* format){
	static_cast<QTextTableCell*>(ptr)->setFormat(*static_cast<QTextCharFormat*>(format));
}

int QTextTableCell_TableCellFormatIndex(void* ptr){
	return static_cast<QTextTableCell*>(ptr)->tableCellFormatIndex();
}

void QTextTableCell_DestroyQTextTableCell(void* ptr){
	static_cast<QTextTableCell*>(ptr)->~QTextTableCell();
}

