#include "qtextline.h"
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QTextLine>
#include "_cgo_export.h"

class MyQTextLine: public QTextLine {
public:
};

QtObjectPtr QTextLine_NewQTextLine(){
	return new QTextLine();
}

int QTextLine_IsValid(QtObjectPtr ptr){
	return static_cast<QTextLine*>(ptr)->isValid();
}

int QTextLine_LeadingIncluded(QtObjectPtr ptr){
	return static_cast<QTextLine*>(ptr)->leadingIncluded();
}

int QTextLine_LineNumber(QtObjectPtr ptr){
	return static_cast<QTextLine*>(ptr)->lineNumber();
}

void QTextLine_SetLeadingIncluded(QtObjectPtr ptr, int included){
	static_cast<QTextLine*>(ptr)->setLeadingIncluded(included != 0);
}

void QTextLine_SetNumColumns(QtObjectPtr ptr, int numColumns){
	static_cast<QTextLine*>(ptr)->setNumColumns(numColumns);
}

void QTextLine_SetPosition(QtObjectPtr ptr, QtObjectPtr pos){
	static_cast<QTextLine*>(ptr)->setPosition(*static_cast<QPointF*>(pos));
}

int QTextLine_TextLength(QtObjectPtr ptr){
	return static_cast<QTextLine*>(ptr)->textLength();
}

int QTextLine_TextStart(QtObjectPtr ptr){
	return static_cast<QTextLine*>(ptr)->textStart();
}

