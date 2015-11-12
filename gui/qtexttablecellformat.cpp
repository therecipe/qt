#include "qtexttablecellformat.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextTable>
#include <QTextTableCell>
#include <QTextTableCellFormat>
#include "_cgo_export.h"

class MyQTextTableCellFormat: public QTextTableCellFormat {
public:
};

void* QTextTableCellFormat_NewQTextTableCellFormat(){
	return new QTextTableCellFormat();
}

double QTextTableCellFormat_BottomPadding(void* ptr){
	return static_cast<double>(static_cast<QTextTableCellFormat*>(ptr)->bottomPadding());
}

int QTextTableCellFormat_IsValid(void* ptr){
	return static_cast<QTextTableCellFormat*>(ptr)->isValid();
}

double QTextTableCellFormat_LeftPadding(void* ptr){
	return static_cast<double>(static_cast<QTextTableCellFormat*>(ptr)->leftPadding());
}

double QTextTableCellFormat_RightPadding(void* ptr){
	return static_cast<double>(static_cast<QTextTableCellFormat*>(ptr)->rightPadding());
}

void QTextTableCellFormat_SetBottomPadding(void* ptr, double padding){
	static_cast<QTextTableCellFormat*>(ptr)->setBottomPadding(static_cast<qreal>(padding));
}

void QTextTableCellFormat_SetLeftPadding(void* ptr, double padding){
	static_cast<QTextTableCellFormat*>(ptr)->setLeftPadding(static_cast<qreal>(padding));
}

void QTextTableCellFormat_SetPadding(void* ptr, double padding){
	static_cast<QTextTableCellFormat*>(ptr)->setPadding(static_cast<qreal>(padding));
}

void QTextTableCellFormat_SetRightPadding(void* ptr, double padding){
	static_cast<QTextTableCellFormat*>(ptr)->setRightPadding(static_cast<qreal>(padding));
}

void QTextTableCellFormat_SetTopPadding(void* ptr, double padding){
	static_cast<QTextTableCellFormat*>(ptr)->setTopPadding(static_cast<qreal>(padding));
}

double QTextTableCellFormat_TopPadding(void* ptr){
	return static_cast<double>(static_cast<QTextTableCellFormat*>(ptr)->topPadding());
}

