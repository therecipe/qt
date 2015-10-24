#include "qtextframeformat.h"
#include <QUrl>
#include <QModelIndex>
#include <QTextFormat>
#include <QBrush>
#include <QTextFrame>
#include <QTextLength>
#include <QString>
#include <QVariant>
#include <QTextFrameFormat>
#include "_cgo_export.h"

class MyQTextFrameFormat: public QTextFrameFormat {
public:
};

QtObjectPtr QTextFrameFormat_NewQTextFrameFormat(){
	return new QTextFrameFormat();
}

int QTextFrameFormat_BorderStyle(QtObjectPtr ptr){
	return static_cast<QTextFrameFormat*>(ptr)->borderStyle();
}

int QTextFrameFormat_IsValid(QtObjectPtr ptr){
	return static_cast<QTextFrameFormat*>(ptr)->isValid();
}

int QTextFrameFormat_PageBreakPolicy(QtObjectPtr ptr){
	return static_cast<QTextFrameFormat*>(ptr)->pageBreakPolicy();
}

int QTextFrameFormat_Position(QtObjectPtr ptr){
	return static_cast<QTextFrameFormat*>(ptr)->position();
}

void QTextFrameFormat_SetBorderBrush(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QTextFrameFormat*>(ptr)->setBorderBrush(*static_cast<QBrush*>(brush));
}

void QTextFrameFormat_SetBorderStyle(QtObjectPtr ptr, int style){
	static_cast<QTextFrameFormat*>(ptr)->setBorderStyle(static_cast<QTextFrameFormat::BorderStyle>(style));
}

void QTextFrameFormat_SetHeight(QtObjectPtr ptr, QtObjectPtr height){
	static_cast<QTextFrameFormat*>(ptr)->setHeight(*static_cast<QTextLength*>(height));
}

void QTextFrameFormat_SetPageBreakPolicy(QtObjectPtr ptr, int policy){
	static_cast<QTextFrameFormat*>(ptr)->setPageBreakPolicy(static_cast<QTextFormat::PageBreakFlag>(policy));
}

void QTextFrameFormat_SetPosition(QtObjectPtr ptr, int policy){
	static_cast<QTextFrameFormat*>(ptr)->setPosition(static_cast<QTextFrameFormat::Position>(policy));
}

void QTextFrameFormat_SetWidth(QtObjectPtr ptr, QtObjectPtr width){
	static_cast<QTextFrameFormat*>(ptr)->setWidth(*static_cast<QTextLength*>(width));
}

