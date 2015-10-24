#include "qtexttableformat.h"
#include <QTextTable>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextTableFormat>
#include "_cgo_export.h"

class MyQTextTableFormat: public QTextTableFormat {
public:
};

QtObjectPtr QTextTableFormat_NewQTextTableFormat(){
	return new QTextTableFormat();
}

int QTextTableFormat_Alignment(QtObjectPtr ptr){
	return static_cast<QTextTableFormat*>(ptr)->alignment();
}

void QTextTableFormat_ClearColumnWidthConstraints(QtObjectPtr ptr){
	static_cast<QTextTableFormat*>(ptr)->clearColumnWidthConstraints();
}

int QTextTableFormat_Columns(QtObjectPtr ptr){
	return static_cast<QTextTableFormat*>(ptr)->columns();
}

int QTextTableFormat_HeaderRowCount(QtObjectPtr ptr){
	return static_cast<QTextTableFormat*>(ptr)->headerRowCount();
}

int QTextTableFormat_IsValid(QtObjectPtr ptr){
	return static_cast<QTextTableFormat*>(ptr)->isValid();
}

void QTextTableFormat_SetAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QTextTableFormat*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextTableFormat_SetHeaderRowCount(QtObjectPtr ptr, int count){
	static_cast<QTextTableFormat*>(ptr)->setHeaderRowCount(count);
}

