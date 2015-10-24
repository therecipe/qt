#include "qtextlayout.h"
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPaintDevice>
#include <QPointF>
#include <QPainter>
#include <QTextOption>
#include <QFont>
#include <QTextLayout>
#include "_cgo_export.h"

class MyQTextLayout: public QTextLayout {
public:
};

void QTextLayout_DrawCursor2(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr position, int cursorPosition){
	static_cast<QTextLayout*>(ptr)->drawCursor(static_cast<QPainter*>(painter), *static_cast<QPointF*>(position), cursorPosition);
}

void QTextLayout_DrawCursor(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr position, int cursorPosition, int width){
	static_cast<QTextLayout*>(ptr)->drawCursor(static_cast<QPainter*>(painter), *static_cast<QPointF*>(position), cursorPosition, width);
}

QtObjectPtr QTextLayout_NewQTextLayout(){
	return new QTextLayout();
}

QtObjectPtr QTextLayout_NewQTextLayout2(char* text){
	return new QTextLayout(QString(text));
}

QtObjectPtr QTextLayout_NewQTextLayout3(char* text, QtObjectPtr font, QtObjectPtr paintdevice){
	return new QTextLayout(QString(text), *static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
}

void QTextLayout_BeginLayout(QtObjectPtr ptr){
	static_cast<QTextLayout*>(ptr)->beginLayout();
}

int QTextLayout_CacheEnabled(QtObjectPtr ptr){
	return static_cast<QTextLayout*>(ptr)->cacheEnabled();
}

void QTextLayout_ClearAdditionalFormats(QtObjectPtr ptr){
	static_cast<QTextLayout*>(ptr)->clearAdditionalFormats();
}

void QTextLayout_ClearLayout(QtObjectPtr ptr){
	static_cast<QTextLayout*>(ptr)->clearLayout();
}

int QTextLayout_CursorMoveStyle(QtObjectPtr ptr){
	return static_cast<QTextLayout*>(ptr)->cursorMoveStyle();
}

void QTextLayout_EndLayout(QtObjectPtr ptr){
	static_cast<QTextLayout*>(ptr)->endLayout();
}

int QTextLayout_IsValidCursorPosition(QtObjectPtr ptr, int pos){
	return static_cast<QTextLayout*>(ptr)->isValidCursorPosition(pos);
}

int QTextLayout_LeftCursorPosition(QtObjectPtr ptr, int oldPos){
	return static_cast<QTextLayout*>(ptr)->leftCursorPosition(oldPos);
}

int QTextLayout_LineCount(QtObjectPtr ptr){
	return static_cast<QTextLayout*>(ptr)->lineCount();
}

int QTextLayout_NextCursorPosition(QtObjectPtr ptr, int oldPos, int mode){
	return static_cast<QTextLayout*>(ptr)->nextCursorPosition(oldPos, static_cast<QTextLayout::CursorMode>(mode));
}

int QTextLayout_PreeditAreaPosition(QtObjectPtr ptr){
	return static_cast<QTextLayout*>(ptr)->preeditAreaPosition();
}

char* QTextLayout_PreeditAreaText(QtObjectPtr ptr){
	return static_cast<QTextLayout*>(ptr)->preeditAreaText().toUtf8().data();
}

int QTextLayout_PreviousCursorPosition(QtObjectPtr ptr, int oldPos, int mode){
	return static_cast<QTextLayout*>(ptr)->previousCursorPosition(oldPos, static_cast<QTextLayout::CursorMode>(mode));
}

int QTextLayout_RightCursorPosition(QtObjectPtr ptr, int oldPos){
	return static_cast<QTextLayout*>(ptr)->rightCursorPosition(oldPos);
}

void QTextLayout_SetCacheEnabled(QtObjectPtr ptr, int enable){
	static_cast<QTextLayout*>(ptr)->setCacheEnabled(enable != 0);
}

void QTextLayout_SetCursorMoveStyle(QtObjectPtr ptr, int style){
	static_cast<QTextLayout*>(ptr)->setCursorMoveStyle(static_cast<Qt::CursorMoveStyle>(style));
}

void QTextLayout_SetFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QTextLayout*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTextLayout_SetPosition(QtObjectPtr ptr, QtObjectPtr p){
	static_cast<QTextLayout*>(ptr)->setPosition(*static_cast<QPointF*>(p));
}

void QTextLayout_SetPreeditArea(QtObjectPtr ptr, int position, char* text){
	static_cast<QTextLayout*>(ptr)->setPreeditArea(position, QString(text));
}

void QTextLayout_SetText(QtObjectPtr ptr, char* stri){
	static_cast<QTextLayout*>(ptr)->setText(QString(stri));
}

void QTextLayout_SetTextOption(QtObjectPtr ptr, QtObjectPtr option){
	static_cast<QTextLayout*>(ptr)->setTextOption(*static_cast<QTextOption*>(option));
}

char* QTextLayout_Text(QtObjectPtr ptr){
	return static_cast<QTextLayout*>(ptr)->text().toUtf8().data();
}

void QTextLayout_DestroyQTextLayout(QtObjectPtr ptr){
	static_cast<QTextLayout*>(ptr)->~QTextLayout();
}

