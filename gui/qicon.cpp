#include "qicon.h"
#include <QSize>
#include <QRect>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPixmap>
#include <QPainter>
#include <QIconEngine>
#include <QIcon>
#include "_cgo_export.h"

class MyQIcon: public QIcon {
public:
};

QtObjectPtr QIcon_NewQIcon(){
	return new QIcon();
}

QtObjectPtr QIcon_NewQIcon4(QtObjectPtr other){
	return new QIcon(*static_cast<QIcon*>(other));
}

QtObjectPtr QIcon_NewQIcon6(QtObjectPtr engine){
	return new QIcon(static_cast<QIconEngine*>(engine));
}

QtObjectPtr QIcon_NewQIcon3(QtObjectPtr other){
	return new QIcon(*static_cast<QIcon*>(other));
}

QtObjectPtr QIcon_NewQIcon2(QtObjectPtr pixmap){
	return new QIcon(*static_cast<QPixmap*>(pixmap));
}

QtObjectPtr QIcon_NewQIcon5(char* fileName){
	return new QIcon(QString(fileName));
}

void QIcon_AddFile(QtObjectPtr ptr, char* fileName, QtObjectPtr size, int mode, int state){
	static_cast<QIcon*>(ptr)->addFile(QString(fileName), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIcon_AddPixmap(QtObjectPtr ptr, QtObjectPtr pixmap, int mode, int state){
	static_cast<QIcon*>(ptr)->addPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

int QIcon_QIcon_HasThemeIcon(char* name){
	return QIcon::hasThemeIcon(QString(name));
}

int QIcon_IsNull(QtObjectPtr ptr){
	return static_cast<QIcon*>(ptr)->isNull();
}

char* QIcon_Name(QtObjectPtr ptr){
	return static_cast<QIcon*>(ptr)->name().toUtf8().data();
}

void QIcon_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr rect, int alignment, int mode, int state){
	static_cast<QIcon*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), static_cast<Qt::AlignmentFlag>(alignment), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIcon_Paint2(QtObjectPtr ptr, QtObjectPtr painter, int x, int y, int w, int h, int alignment, int mode, int state){
	static_cast<QIcon*>(ptr)->paint(static_cast<QPainter*>(painter), x, y, w, h, static_cast<Qt::AlignmentFlag>(alignment), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIcon_QIcon_SetThemeName(char* name){
	QIcon::setThemeName(QString(name));
}

void QIcon_QIcon_SetThemeSearchPaths(char* paths){
	QIcon::setThemeSearchPaths(QString(paths).split("|", QString::SkipEmptyParts));
}

void QIcon_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QIcon*>(ptr)->swap(*static_cast<QIcon*>(other));
}

char* QIcon_QIcon_ThemeName(){
	return QIcon::themeName().toUtf8().data();
}

char* QIcon_QIcon_ThemeSearchPaths(){
	return QIcon::themeSearchPaths().join("|").toUtf8().data();
}

void QIcon_DestroyQIcon(QtObjectPtr ptr){
	static_cast<QIcon*>(ptr)->~QIcon();
}

