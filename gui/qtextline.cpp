#include "qtextline.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QTextLine>
#include "_cgo_export.h"

class MyQTextLine: public QTextLine {
public:
};

int QTextLine_XToCursor(void* ptr, double x, int cpos){
	return static_cast<QTextLine*>(ptr)->xToCursor(static_cast<qreal>(x), static_cast<QTextLine::CursorPosition>(cpos));
}

void* QTextLine_NewQTextLine(){
	return new QTextLine();
}

double QTextLine_Ascent(void* ptr){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->ascent());
}

double QTextLine_CursorToX(void* ptr, int cursorPos, int edge){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->cursorToX(&cursorPos, static_cast<QTextLine::Edge>(edge)));
}

double QTextLine_CursorToX2(void* ptr, int cursorPos, int edge){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->cursorToX(cursorPos, static_cast<QTextLine::Edge>(edge)));
}

double QTextLine_Descent(void* ptr){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->descent());
}

double QTextLine_Height(void* ptr){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->height());
}

double QTextLine_HorizontalAdvance(void* ptr){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->horizontalAdvance());
}

int QTextLine_IsValid(void* ptr){
	return static_cast<QTextLine*>(ptr)->isValid();
}

double QTextLine_Leading(void* ptr){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->leading());
}

int QTextLine_LeadingIncluded(void* ptr){
	return static_cast<QTextLine*>(ptr)->leadingIncluded();
}

int QTextLine_LineNumber(void* ptr){
	return static_cast<QTextLine*>(ptr)->lineNumber();
}

double QTextLine_NaturalTextWidth(void* ptr){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->naturalTextWidth());
}

void QTextLine_SetLeadingIncluded(void* ptr, int included){
	static_cast<QTextLine*>(ptr)->setLeadingIncluded(included != 0);
}

void QTextLine_SetLineWidth(void* ptr, double width){
	static_cast<QTextLine*>(ptr)->setLineWidth(static_cast<qreal>(width));
}

void QTextLine_SetNumColumns(void* ptr, int numColumns){
	static_cast<QTextLine*>(ptr)->setNumColumns(numColumns);
}

void QTextLine_SetNumColumns2(void* ptr, int numColumns, double alignmentWidth){
	static_cast<QTextLine*>(ptr)->setNumColumns(numColumns, static_cast<qreal>(alignmentWidth));
}

void QTextLine_SetPosition(void* ptr, void* pos){
	static_cast<QTextLine*>(ptr)->setPosition(*static_cast<QPointF*>(pos));
}

int QTextLine_TextLength(void* ptr){
	return static_cast<QTextLine*>(ptr)->textLength();
}

int QTextLine_TextStart(void* ptr){
	return static_cast<QTextLine*>(ptr)->textStart();
}

double QTextLine_Width(void* ptr){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->width());
}

double QTextLine_X(void* ptr){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->x());
}

double QTextLine_Y(void* ptr){
	return static_cast<double>(static_cast<QTextLine*>(ptr)->y());
}

