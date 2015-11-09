#include "qtexttableformat.h"
#include <QUrl>
#include <QModelIndex>
#include <QTextTable>
#include <QString>
#include <QVariant>
#include <QTextTableFormat>
#include "_cgo_export.h"

class MyQTextTableFormat: public QTextTableFormat {
public:
};

void* QTextTableFormat_NewQTextTableFormat(){
	return new QTextTableFormat();
}

int QTextTableFormat_Alignment(void* ptr){
	return static_cast<QTextTableFormat*>(ptr)->alignment();
}

double QTextTableFormat_CellPadding(void* ptr){
	return static_cast<double>(static_cast<QTextTableFormat*>(ptr)->cellPadding());
}

double QTextTableFormat_CellSpacing(void* ptr){
	return static_cast<double>(static_cast<QTextTableFormat*>(ptr)->cellSpacing());
}

void QTextTableFormat_ClearColumnWidthConstraints(void* ptr){
	static_cast<QTextTableFormat*>(ptr)->clearColumnWidthConstraints();
}

int QTextTableFormat_Columns(void* ptr){
	return static_cast<QTextTableFormat*>(ptr)->columns();
}

int QTextTableFormat_HeaderRowCount(void* ptr){
	return static_cast<QTextTableFormat*>(ptr)->headerRowCount();
}

int QTextTableFormat_IsValid(void* ptr){
	return static_cast<QTextTableFormat*>(ptr)->isValid();
}

void QTextTableFormat_SetAlignment(void* ptr, int alignment){
	static_cast<QTextTableFormat*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextTableFormat_SetCellPadding(void* ptr, double padding){
	static_cast<QTextTableFormat*>(ptr)->setCellPadding(static_cast<qreal>(padding));
}

void QTextTableFormat_SetCellSpacing(void* ptr, double spacing){
	static_cast<QTextTableFormat*>(ptr)->setCellSpacing(static_cast<qreal>(spacing));
}

void QTextTableFormat_SetHeaderRowCount(void* ptr, int count){
	static_cast<QTextTableFormat*>(ptr)->setHeaderRowCount(count);
}

