#include "qstylepainter.h"
#include <QUrl>
#include <QWidget>
#include <QStyle>
#include <QPalette>
#include <QPaintDevice>
#include <QStyleOption>
#include <QPixmap>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QStyleOptionComplex>
#include <QRect>
#include <QStylePainter>
#include "_cgo_export.h"

class MyQStylePainter: public QStylePainter {
public:
};

QtObjectPtr QStylePainter_NewQStylePainter(){
	return new QStylePainter();
}

QtObjectPtr QStylePainter_NewQStylePainter3(QtObjectPtr pd, QtObjectPtr widget){
	return new QStylePainter(static_cast<QPaintDevice*>(pd), static_cast<QWidget*>(widget));
}

QtObjectPtr QStylePainter_NewQStylePainter2(QtObjectPtr widget){
	return new QStylePainter(static_cast<QWidget*>(widget));
}

int QStylePainter_Begin2(QtObjectPtr ptr, QtObjectPtr pd, QtObjectPtr widget){
	return static_cast<QStylePainter*>(ptr)->begin(static_cast<QPaintDevice*>(pd), static_cast<QWidget*>(widget));
}

int QStylePainter_Begin(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QStylePainter*>(ptr)->begin(static_cast<QWidget*>(widget));
}

void QStylePainter_DrawComplexControl(QtObjectPtr ptr, int cc, QtObjectPtr option){
	static_cast<QStylePainter*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(cc), *static_cast<QStyleOptionComplex*>(option));
}

void QStylePainter_DrawControl(QtObjectPtr ptr, int ce, QtObjectPtr option){
	static_cast<QStylePainter*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(ce), *static_cast<QStyleOption*>(option));
}

void QStylePainter_DrawItemPixmap(QtObjectPtr ptr, QtObjectPtr rect, int flags, QtObjectPtr pixmap){
	static_cast<QStylePainter*>(ptr)->drawItemPixmap(*static_cast<QRect*>(rect), flags, *static_cast<QPixmap*>(pixmap));
}

void QStylePainter_DrawItemText(QtObjectPtr ptr, QtObjectPtr rect, int flags, QtObjectPtr pal, int enabled, char* text, int textRole){
	static_cast<QStylePainter*>(ptr)->drawItemText(*static_cast<QRect*>(rect), flags, *static_cast<QPalette*>(pal), enabled != 0, QString(text), static_cast<QPalette::ColorRole>(textRole));
}

void QStylePainter_DrawPrimitive(QtObjectPtr ptr, int pe, QtObjectPtr option){
	static_cast<QStylePainter*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(pe), *static_cast<QStyleOption*>(option));
}

QtObjectPtr QStylePainter_Style(QtObjectPtr ptr){
	return static_cast<QStylePainter*>(ptr)->style();
}

