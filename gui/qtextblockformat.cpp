#include "qtextblockformat.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextFormat>
#include <QTextBlock>
#include <QTextBlockFormat>
#include "_cgo_export.h"

class MyQTextBlockFormat: public QTextBlockFormat {
public:
};

QtObjectPtr QTextBlockFormat_NewQTextBlockFormat(){
	return new QTextBlockFormat();
}

int QTextBlockFormat_Alignment(QtObjectPtr ptr){
	return static_cast<QTextBlockFormat*>(ptr)->alignment();
}

int QTextBlockFormat_Indent(QtObjectPtr ptr){
	return static_cast<QTextBlockFormat*>(ptr)->indent();
}

int QTextBlockFormat_IsValid(QtObjectPtr ptr){
	return static_cast<QTextBlockFormat*>(ptr)->isValid();
}

int QTextBlockFormat_LineHeightType(QtObjectPtr ptr){
	return static_cast<QTextBlockFormat*>(ptr)->lineHeightType();
}

int QTextBlockFormat_NonBreakableLines(QtObjectPtr ptr){
	return static_cast<QTextBlockFormat*>(ptr)->nonBreakableLines();
}

int QTextBlockFormat_PageBreakPolicy(QtObjectPtr ptr){
	return static_cast<QTextBlockFormat*>(ptr)->pageBreakPolicy();
}

void QTextBlockFormat_SetAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QTextBlockFormat*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextBlockFormat_SetIndent(QtObjectPtr ptr, int indentation){
	static_cast<QTextBlockFormat*>(ptr)->setIndent(indentation);
}

void QTextBlockFormat_SetNonBreakableLines(QtObjectPtr ptr, int b){
	static_cast<QTextBlockFormat*>(ptr)->setNonBreakableLines(b != 0);
}

void QTextBlockFormat_SetPageBreakPolicy(QtObjectPtr ptr, int policy){
	static_cast<QTextBlockFormat*>(ptr)->setPageBreakPolicy(static_cast<QTextFormat::PageBreakFlag>(policy));
}

