#include "qproxystyle.h"
#include <QModelIndex>
#include <QWidget>
#include <QPixmap>
#include <QPoint>
#include <QPalette>
#include <QStyle>
#include <QRect>
#include <QString>
#include <QVariant>
#include <QStyleOptionComplex>
#include <QStyleHintReturn>
#include <QStyleOption>
#include <QSizePolicy>
#include <QApplication>
#include <QPainter>
#include <QUrl>
#include <QSize>
#include <QProxyStyle>
#include "_cgo_export.h"

class MyQProxyStyle: public QProxyStyle {
public:
};

void* QProxyStyle_BaseStyle(void* ptr){
	return static_cast<QProxyStyle*>(ptr)->baseStyle();
}

void QProxyStyle_DrawComplexControl(void* ptr, int control, void* option, void* painter, void* widget){
	static_cast<QProxyStyle*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void QProxyStyle_DrawControl(void* ptr, int element, void* option, void* painter, void* widget){
	static_cast<QProxyStyle*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(element), static_cast<QStyleOption*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void QProxyStyle_DrawItemPixmap(void* ptr, void* painter, void* rect, int alignment, void* pixmap){
	static_cast<QProxyStyle*>(ptr)->drawItemPixmap(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), alignment, *static_cast<QPixmap*>(pixmap));
}

void QProxyStyle_DrawItemText(void* ptr, void* painter, void* rect, int flags, void* pal, int enabled, char* text, int textRole){
	static_cast<QProxyStyle*>(ptr)->drawItemText(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), flags, *static_cast<QPalette*>(pal), enabled != 0, QString(text), static_cast<QPalette::ColorRole>(textRole));
}

void QProxyStyle_DrawPrimitive(void* ptr, int element, void* option, void* painter, void* widget){
	static_cast<QProxyStyle*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(element), static_cast<QStyleOption*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

int QProxyStyle_HitTestComplexControl(void* ptr, int control, void* option, void* pos, void* widget){
	return static_cast<QProxyStyle*>(ptr)->hitTestComplexControl(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), *static_cast<QPoint*>(pos), static_cast<QWidget*>(widget));
}

int QProxyStyle_LayoutSpacing(void* ptr, int control1, int control2, int orientation, void* option, void* widget){
	return static_cast<QProxyStyle*>(ptr)->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

int QProxyStyle_PixelMetric(void* ptr, int metric, void* option, void* widget){
	return static_cast<QProxyStyle*>(ptr)->pixelMetric(static_cast<QStyle::PixelMetric>(metric), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

void QProxyStyle_Polish3(void* ptr, void* app){
	static_cast<QProxyStyle*>(ptr)->polish(static_cast<QApplication*>(app));
}

void QProxyStyle_Polish2(void* ptr, void* pal){
	static_cast<QProxyStyle*>(ptr)->polish(*static_cast<QPalette*>(pal));
}

void QProxyStyle_Polish(void* ptr, void* widget){
	static_cast<QProxyStyle*>(ptr)->polish(static_cast<QWidget*>(widget));
}

void QProxyStyle_SetBaseStyle(void* ptr, void* style){
	static_cast<QProxyStyle*>(ptr)->setBaseStyle(static_cast<QStyle*>(style));
}

int QProxyStyle_StyleHint(void* ptr, int hint, void* option, void* widget, void* returnData){
	return static_cast<QProxyStyle*>(ptr)->styleHint(static_cast<QStyle::StyleHint>(hint), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget), static_cast<QStyleHintReturn*>(returnData));
}

void QProxyStyle_Unpolish2(void* ptr, void* app){
	static_cast<QProxyStyle*>(ptr)->unpolish(static_cast<QApplication*>(app));
}

void QProxyStyle_Unpolish(void* ptr, void* widget){
	static_cast<QProxyStyle*>(ptr)->unpolish(static_cast<QWidget*>(widget));
}

void QProxyStyle_DestroyQProxyStyle(void* ptr){
	static_cast<QProxyStyle*>(ptr)->~QProxyStyle();
}

