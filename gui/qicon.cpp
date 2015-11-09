#include "qicon.h"
#include <QSize>
#include <QRect>
#include <QIconEngine>
#include <QPainter>
#include <QPixmap>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIcon>
#include "_cgo_export.h"

class MyQIcon: public QIcon {
public:
};

void* QIcon_NewQIcon(){
	return new QIcon();
}

void* QIcon_NewQIcon4(void* other){
	return new QIcon(*static_cast<QIcon*>(other));
}

void* QIcon_NewQIcon6(void* engine){
	return new QIcon(static_cast<QIconEngine*>(engine));
}

void* QIcon_NewQIcon3(void* other){
	return new QIcon(*static_cast<QIcon*>(other));
}

void* QIcon_NewQIcon2(void* pixmap){
	return new QIcon(*static_cast<QPixmap*>(pixmap));
}

void* QIcon_NewQIcon5(char* fileName){
	return new QIcon(QString(fileName));
}

void QIcon_AddFile(void* ptr, char* fileName, void* size, int mode, int state){
	static_cast<QIcon*>(ptr)->addFile(QString(fileName), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIcon_AddPixmap(void* ptr, void* pixmap, int mode, int state){
	static_cast<QIcon*>(ptr)->addPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

int QIcon_QIcon_HasThemeIcon(char* name){
	return QIcon::hasThemeIcon(QString(name));
}

int QIcon_IsNull(void* ptr){
	return static_cast<QIcon*>(ptr)->isNull();
}

char* QIcon_Name(void* ptr){
	return static_cast<QIcon*>(ptr)->name().toUtf8().data();
}

void QIcon_Paint(void* ptr, void* painter, void* rect, int alignment, int mode, int state){
	static_cast<QIcon*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), static_cast<Qt::AlignmentFlag>(alignment), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIcon_Paint2(void* ptr, void* painter, int x, int y, int w, int h, int alignment, int mode, int state){
	static_cast<QIcon*>(ptr)->paint(static_cast<QPainter*>(painter), x, y, w, h, static_cast<Qt::AlignmentFlag>(alignment), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIcon_QIcon_SetThemeName(char* name){
	QIcon::setThemeName(QString(name));
}

void QIcon_QIcon_SetThemeSearchPaths(char* paths){
	QIcon::setThemeSearchPaths(QString(paths).split("|", QString::SkipEmptyParts));
}

void QIcon_Swap(void* ptr, void* other){
	static_cast<QIcon*>(ptr)->swap(*static_cast<QIcon*>(other));
}

char* QIcon_QIcon_ThemeName(){
	return QIcon::themeName().toUtf8().data();
}

char* QIcon_QIcon_ThemeSearchPaths(){
	return QIcon::themeSearchPaths().join("|").toUtf8().data();
}

void QIcon_DestroyQIcon(void* ptr){
	static_cast<QIcon*>(ptr)->~QIcon();
}

