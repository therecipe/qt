#include "qstylepainter.h"
#include <QStyle>
#include <QRect>
#include <QWidget>
#include <QPixmap>
#include <QPalette>
#include <QModelIndex>
#include <QStyleOption>
#include <QPaintDevice>
#include <QStyleOptionComplex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStylePainter>
#include "_cgo_export.h"

class MyQStylePainter: public QStylePainter {
public:
};

void* QStylePainter_NewQStylePainter(){
	return new QStylePainter();
}

void* QStylePainter_NewQStylePainter3(void* pd, void* widget){
	return new QStylePainter(static_cast<QPaintDevice*>(pd), static_cast<QWidget*>(widget));
}

void* QStylePainter_NewQStylePainter2(void* widget){
	return new QStylePainter(static_cast<QWidget*>(widget));
}

int QStylePainter_Begin2(void* ptr, void* pd, void* widget){
	return static_cast<QStylePainter*>(ptr)->begin(static_cast<QPaintDevice*>(pd), static_cast<QWidget*>(widget));
}

int QStylePainter_Begin(void* ptr, void* widget){
	return static_cast<QStylePainter*>(ptr)->begin(static_cast<QWidget*>(widget));
}

void QStylePainter_DrawComplexControl(void* ptr, int cc, void* option){
	static_cast<QStylePainter*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(cc), *static_cast<QStyleOptionComplex*>(option));
}

void QStylePainter_DrawControl(void* ptr, int ce, void* option){
	static_cast<QStylePainter*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(ce), *static_cast<QStyleOption*>(option));
}

void QStylePainter_DrawItemPixmap(void* ptr, void* rect, int flags, void* pixmap){
	static_cast<QStylePainter*>(ptr)->drawItemPixmap(*static_cast<QRect*>(rect), flags, *static_cast<QPixmap*>(pixmap));
}

void QStylePainter_DrawItemText(void* ptr, void* rect, int flags, void* pal, int enabled, char* text, int textRole){
	static_cast<QStylePainter*>(ptr)->drawItemText(*static_cast<QRect*>(rect), flags, *static_cast<QPalette*>(pal), enabled != 0, QString(text), static_cast<QPalette::ColorRole>(textRole));
}

void QStylePainter_DrawPrimitive(void* ptr, int pe, void* option){
	static_cast<QStylePainter*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(pe), *static_cast<QStyleOption*>(option));
}

void* QStylePainter_Style(void* ptr){
	return static_cast<QStylePainter*>(ptr)->style();
}

