#include "qtextblockformat.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextBlock>
#include <QTextFormat>
#include <QTextBlockFormat>
#include "_cgo_export.h"

class MyQTextBlockFormat: public QTextBlockFormat {
public:
};

void* QTextBlockFormat_NewQTextBlockFormat(){
	return new QTextBlockFormat();
}

int QTextBlockFormat_Alignment(void* ptr){
	return static_cast<QTextBlockFormat*>(ptr)->alignment();
}

double QTextBlockFormat_BottomMargin(void* ptr){
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->bottomMargin());
}

int QTextBlockFormat_Indent(void* ptr){
	return static_cast<QTextBlockFormat*>(ptr)->indent();
}

int QTextBlockFormat_IsValid(void* ptr){
	return static_cast<QTextBlockFormat*>(ptr)->isValid();
}

double QTextBlockFormat_LeftMargin(void* ptr){
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->leftMargin());
}

double QTextBlockFormat_LineHeight2(void* ptr){
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->lineHeight());
}

double QTextBlockFormat_LineHeight(void* ptr, double scriptLineHeight, double scaling){
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->lineHeight(static_cast<qreal>(scriptLineHeight), static_cast<qreal>(scaling)));
}

int QTextBlockFormat_LineHeightType(void* ptr){
	return static_cast<QTextBlockFormat*>(ptr)->lineHeightType();
}

int QTextBlockFormat_NonBreakableLines(void* ptr){
	return static_cast<QTextBlockFormat*>(ptr)->nonBreakableLines();
}

int QTextBlockFormat_PageBreakPolicy(void* ptr){
	return static_cast<QTextBlockFormat*>(ptr)->pageBreakPolicy();
}

double QTextBlockFormat_RightMargin(void* ptr){
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->rightMargin());
}

void QTextBlockFormat_SetAlignment(void* ptr, int alignment){
	static_cast<QTextBlockFormat*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextBlockFormat_SetBottomMargin(void* ptr, double margin){
	static_cast<QTextBlockFormat*>(ptr)->setBottomMargin(static_cast<qreal>(margin));
}

void QTextBlockFormat_SetIndent(void* ptr, int indentation){
	static_cast<QTextBlockFormat*>(ptr)->setIndent(indentation);
}

void QTextBlockFormat_SetLeftMargin(void* ptr, double margin){
	static_cast<QTextBlockFormat*>(ptr)->setLeftMargin(static_cast<qreal>(margin));
}

void QTextBlockFormat_SetLineHeight(void* ptr, double height, int heightType){
	static_cast<QTextBlockFormat*>(ptr)->setLineHeight(static_cast<qreal>(height), heightType);
}

void QTextBlockFormat_SetNonBreakableLines(void* ptr, int b){
	static_cast<QTextBlockFormat*>(ptr)->setNonBreakableLines(b != 0);
}

void QTextBlockFormat_SetPageBreakPolicy(void* ptr, int policy){
	static_cast<QTextBlockFormat*>(ptr)->setPageBreakPolicy(static_cast<QTextFormat::PageBreakFlag>(policy));
}

void QTextBlockFormat_SetRightMargin(void* ptr, double margin){
	static_cast<QTextBlockFormat*>(ptr)->setRightMargin(static_cast<qreal>(margin));
}

void QTextBlockFormat_SetTextIndent(void* ptr, double indent){
	static_cast<QTextBlockFormat*>(ptr)->setTextIndent(static_cast<qreal>(indent));
}

void QTextBlockFormat_SetTopMargin(void* ptr, double margin){
	static_cast<QTextBlockFormat*>(ptr)->setTopMargin(static_cast<qreal>(margin));
}

double QTextBlockFormat_TextIndent(void* ptr){
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->textIndent());
}

double QTextBlockFormat_TopMargin(void* ptr){
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->topMargin());
}

