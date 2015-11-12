#include "qtextframeformat.h"
#include <QUrl>
#include <QModelIndex>
#include <QTextLength>
#include <QTextFrame>
#include <QTextFormat>
#include <QBrush>
#include <QString>
#include <QVariant>
#include <QTextFrameFormat>
#include "_cgo_export.h"

class MyQTextFrameFormat: public QTextFrameFormat {
public:
};

void* QTextFrameFormat_NewQTextFrameFormat(){
	return new QTextFrameFormat();
}

double QTextFrameFormat_BottomMargin(void* ptr){
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->bottomMargin());
}

double QTextFrameFormat_LeftMargin(void* ptr){
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->leftMargin());
}

double QTextFrameFormat_RightMargin(void* ptr){
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->rightMargin());
}

void QTextFrameFormat_SetMargin(void* ptr, double margin){
	static_cast<QTextFrameFormat*>(ptr)->setMargin(static_cast<qreal>(margin));
}

double QTextFrameFormat_TopMargin(void* ptr){
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->topMargin());
}

double QTextFrameFormat_Border(void* ptr){
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->border());
}

void* QTextFrameFormat_BorderBrush(void* ptr){
	return new QBrush(static_cast<QTextFrameFormat*>(ptr)->borderBrush());
}

int QTextFrameFormat_BorderStyle(void* ptr){
	return static_cast<QTextFrameFormat*>(ptr)->borderStyle();
}

int QTextFrameFormat_IsValid(void* ptr){
	return static_cast<QTextFrameFormat*>(ptr)->isValid();
}

double QTextFrameFormat_Margin(void* ptr){
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->margin());
}

double QTextFrameFormat_Padding(void* ptr){
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->padding());
}

int QTextFrameFormat_PageBreakPolicy(void* ptr){
	return static_cast<QTextFrameFormat*>(ptr)->pageBreakPolicy();
}

int QTextFrameFormat_Position(void* ptr){
	return static_cast<QTextFrameFormat*>(ptr)->position();
}

void QTextFrameFormat_SetBorder(void* ptr, double width){
	static_cast<QTextFrameFormat*>(ptr)->setBorder(static_cast<qreal>(width));
}

void QTextFrameFormat_SetBorderBrush(void* ptr, void* brush){
	static_cast<QTextFrameFormat*>(ptr)->setBorderBrush(*static_cast<QBrush*>(brush));
}

void QTextFrameFormat_SetBorderStyle(void* ptr, int style){
	static_cast<QTextFrameFormat*>(ptr)->setBorderStyle(static_cast<QTextFrameFormat::BorderStyle>(style));
}

void QTextFrameFormat_SetBottomMargin(void* ptr, double margin){
	static_cast<QTextFrameFormat*>(ptr)->setBottomMargin(static_cast<qreal>(margin));
}

void QTextFrameFormat_SetHeight(void* ptr, void* height){
	static_cast<QTextFrameFormat*>(ptr)->setHeight(*static_cast<QTextLength*>(height));
}

void QTextFrameFormat_SetHeight2(void* ptr, double height){
	static_cast<QTextFrameFormat*>(ptr)->setHeight(static_cast<qreal>(height));
}

void QTextFrameFormat_SetLeftMargin(void* ptr, double margin){
	static_cast<QTextFrameFormat*>(ptr)->setLeftMargin(static_cast<qreal>(margin));
}

void QTextFrameFormat_SetPadding(void* ptr, double width){
	static_cast<QTextFrameFormat*>(ptr)->setPadding(static_cast<qreal>(width));
}

void QTextFrameFormat_SetPageBreakPolicy(void* ptr, int policy){
	static_cast<QTextFrameFormat*>(ptr)->setPageBreakPolicy(static_cast<QTextFormat::PageBreakFlag>(policy));
}

void QTextFrameFormat_SetPosition(void* ptr, int policy){
	static_cast<QTextFrameFormat*>(ptr)->setPosition(static_cast<QTextFrameFormat::Position>(policy));
}

void QTextFrameFormat_SetRightMargin(void* ptr, double margin){
	static_cast<QTextFrameFormat*>(ptr)->setRightMargin(static_cast<qreal>(margin));
}

void QTextFrameFormat_SetTopMargin(void* ptr, double margin){
	static_cast<QTextFrameFormat*>(ptr)->setTopMargin(static_cast<qreal>(margin));
}

void QTextFrameFormat_SetWidth(void* ptr, void* width){
	static_cast<QTextFrameFormat*>(ptr)->setWidth(*static_cast<QTextLength*>(width));
}

void QTextFrameFormat_SetWidth2(void* ptr, double width){
	static_cast<QTextFrameFormat*>(ptr)->setWidth(static_cast<qreal>(width));
}

