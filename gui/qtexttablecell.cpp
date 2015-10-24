#include "qtexttablecell.h"
#include <QTextTable>
#include <QTextCharFormat>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextTableCell>
#include "_cgo_export.h"

class MyQTextTableCell: public QTextTableCell {
public:
};

QtObjectPtr QTextTableCell_NewQTextTableCell(){
	return new QTextTableCell();
}

QtObjectPtr QTextTableCell_NewQTextTableCell2(QtObjectPtr other){
	return new QTextTableCell(*static_cast<QTextTableCell*>(other));
}

int QTextTableCell_Column(QtObjectPtr ptr){
	return static_cast<QTextTableCell*>(ptr)->column();
}

int QTextTableCell_ColumnSpan(QtObjectPtr ptr){
	return static_cast<QTextTableCell*>(ptr)->columnSpan();
}

int QTextTableCell_IsValid(QtObjectPtr ptr){
	return static_cast<QTextTableCell*>(ptr)->isValid();
}

int QTextTableCell_Row(QtObjectPtr ptr){
	return static_cast<QTextTableCell*>(ptr)->row();
}

int QTextTableCell_RowSpan(QtObjectPtr ptr){
	return static_cast<QTextTableCell*>(ptr)->rowSpan();
}

void QTextTableCell_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextTableCell*>(ptr)->setFormat(*static_cast<QTextCharFormat*>(format));
}

int QTextTableCell_TableCellFormatIndex(QtObjectPtr ptr){
	return static_cast<QTextTableCell*>(ptr)->tableCellFormatIndex();
}

void QTextTableCell_DestroyQTextTableCell(QtObjectPtr ptr){
	static_cast<QTextTableCell*>(ptr)->~QTextTableCell();
}

