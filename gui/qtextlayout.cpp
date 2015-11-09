#include "qtextlayout.h"
#include <QModelIndex>
#include <QFont>
#include <QPointF>
#include <QTextOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QPainter>
#include <QPaintDevice>
#include <QPoint>
#include <QTextLayout>
#include "_cgo_export.h"

class MyQTextLayout: public QTextLayout {
public:
};

void QTextLayout_DrawCursor2(void* ptr, void* painter, void* position, int cursorPosition){
	static_cast<QTextLayout*>(ptr)->drawCursor(static_cast<QPainter*>(painter), *static_cast<QPointF*>(position), cursorPosition);
}

void QTextLayout_DrawCursor(void* ptr, void* painter, void* position, int cursorPosition, int width){
	static_cast<QTextLayout*>(ptr)->drawCursor(static_cast<QPainter*>(painter), *static_cast<QPointF*>(position), cursorPosition, width);
}

void* QTextLayout_NewQTextLayout(){
	return new QTextLayout();
}

void* QTextLayout_NewQTextLayout2(char* text){
	return new QTextLayout(QString(text));
}

void* QTextLayout_NewQTextLayout3(char* text, void* font, void* paintdevice){
	return new QTextLayout(QString(text), *static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
}

void QTextLayout_BeginLayout(void* ptr){
	static_cast<QTextLayout*>(ptr)->beginLayout();
}

int QTextLayout_CacheEnabled(void* ptr){
	return static_cast<QTextLayout*>(ptr)->cacheEnabled();
}

void QTextLayout_ClearAdditionalFormats(void* ptr){
	static_cast<QTextLayout*>(ptr)->clearAdditionalFormats();
}

void QTextLayout_ClearLayout(void* ptr){
	static_cast<QTextLayout*>(ptr)->clearLayout();
}

int QTextLayout_CursorMoveStyle(void* ptr){
	return static_cast<QTextLayout*>(ptr)->cursorMoveStyle();
}

void QTextLayout_EndLayout(void* ptr){
	static_cast<QTextLayout*>(ptr)->endLayout();
}

int QTextLayout_IsValidCursorPosition(void* ptr, int pos){
	return static_cast<QTextLayout*>(ptr)->isValidCursorPosition(pos);
}

int QTextLayout_LeftCursorPosition(void* ptr, int oldPos){
	return static_cast<QTextLayout*>(ptr)->leftCursorPosition(oldPos);
}

int QTextLayout_LineCount(void* ptr){
	return static_cast<QTextLayout*>(ptr)->lineCount();
}

double QTextLayout_MaximumWidth(void* ptr){
	return static_cast<double>(static_cast<QTextLayout*>(ptr)->maximumWidth());
}

double QTextLayout_MinimumWidth(void* ptr){
	return static_cast<double>(static_cast<QTextLayout*>(ptr)->minimumWidth());
}

int QTextLayout_NextCursorPosition(void* ptr, int oldPos, int mode){
	return static_cast<QTextLayout*>(ptr)->nextCursorPosition(oldPos, static_cast<QTextLayout::CursorMode>(mode));
}

int QTextLayout_PreeditAreaPosition(void* ptr){
	return static_cast<QTextLayout*>(ptr)->preeditAreaPosition();
}

char* QTextLayout_PreeditAreaText(void* ptr){
	return static_cast<QTextLayout*>(ptr)->preeditAreaText().toUtf8().data();
}

int QTextLayout_PreviousCursorPosition(void* ptr, int oldPos, int mode){
	return static_cast<QTextLayout*>(ptr)->previousCursorPosition(oldPos, static_cast<QTextLayout::CursorMode>(mode));
}

int QTextLayout_RightCursorPosition(void* ptr, int oldPos){
	return static_cast<QTextLayout*>(ptr)->rightCursorPosition(oldPos);
}

void QTextLayout_SetCacheEnabled(void* ptr, int enable){
	static_cast<QTextLayout*>(ptr)->setCacheEnabled(enable != 0);
}

void QTextLayout_SetCursorMoveStyle(void* ptr, int style){
	static_cast<QTextLayout*>(ptr)->setCursorMoveStyle(static_cast<Qt::CursorMoveStyle>(style));
}

void QTextLayout_SetFont(void* ptr, void* font){
	static_cast<QTextLayout*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTextLayout_SetPosition(void* ptr, void* p){
	static_cast<QTextLayout*>(ptr)->setPosition(*static_cast<QPointF*>(p));
}

void QTextLayout_SetPreeditArea(void* ptr, int position, char* text){
	static_cast<QTextLayout*>(ptr)->setPreeditArea(position, QString(text));
}

void QTextLayout_SetText(void* ptr, char* stri){
	static_cast<QTextLayout*>(ptr)->setText(QString(stri));
}

void QTextLayout_SetTextOption(void* ptr, void* option){
	static_cast<QTextLayout*>(ptr)->setTextOption(*static_cast<QTextOption*>(option));
}

char* QTextLayout_Text(void* ptr){
	return static_cast<QTextLayout*>(ptr)->text().toUtf8().data();
}

void QTextLayout_DestroyQTextLayout(void* ptr){
	static_cast<QTextLayout*>(ptr)->~QTextLayout();
}

