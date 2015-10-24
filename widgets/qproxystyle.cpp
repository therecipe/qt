#include "qproxystyle.h"
#include <QWidget>
#include <QPainter>
#include <QPalette>
#include <QRect>
#include <QApplication>
#include <QStyleOptionComplex>
#include <QStyle>
#include <QPoint>
#include <QUrl>
#include <QModelIndex>
#include <QSizePolicy>
#include <QPixmap>
#include <QString>
#include <QVariant>
#include <QSize>
#include <QStyleHintReturn>
#include <QStyleOption>
#include <QProxyStyle>
#include "_cgo_export.h"

class MyQProxyStyle: public QProxyStyle {
public:
};

QtObjectPtr QProxyStyle_BaseStyle(QtObjectPtr ptr){
	return static_cast<QProxyStyle*>(ptr)->baseStyle();
}

void QProxyStyle_DrawComplexControl(QtObjectPtr ptr, int control, QtObjectPtr option, QtObjectPtr painter, QtObjectPtr widget){
	static_cast<QProxyStyle*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void QProxyStyle_DrawControl(QtObjectPtr ptr, int element, QtObjectPtr option, QtObjectPtr painter, QtObjectPtr widget){
	static_cast<QProxyStyle*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(element), static_cast<QStyleOption*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void QProxyStyle_DrawItemPixmap(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr rect, int alignment, QtObjectPtr pixmap){
	static_cast<QProxyStyle*>(ptr)->drawItemPixmap(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), alignment, *static_cast<QPixmap*>(pixmap));
}

void QProxyStyle_DrawItemText(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr rect, int flags, QtObjectPtr pal, int enabled, char* text, int textRole){
	static_cast<QProxyStyle*>(ptr)->drawItemText(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), flags, *static_cast<QPalette*>(pal), enabled != 0, QString(text), static_cast<QPalette::ColorRole>(textRole));
}

void QProxyStyle_DrawPrimitive(QtObjectPtr ptr, int element, QtObjectPtr option, QtObjectPtr painter, QtObjectPtr widget){
	static_cast<QProxyStyle*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(element), static_cast<QStyleOption*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

int QProxyStyle_HitTestComplexControl(QtObjectPtr ptr, int control, QtObjectPtr option, QtObjectPtr pos, QtObjectPtr widget){
	return static_cast<QProxyStyle*>(ptr)->hitTestComplexControl(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), *static_cast<QPoint*>(pos), static_cast<QWidget*>(widget));
}

int QProxyStyle_LayoutSpacing(QtObjectPtr ptr, int control1, int control2, int orientation, QtObjectPtr option, QtObjectPtr widget){
	return static_cast<QProxyStyle*>(ptr)->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

int QProxyStyle_PixelMetric(QtObjectPtr ptr, int metric, QtObjectPtr option, QtObjectPtr widget){
	return static_cast<QProxyStyle*>(ptr)->pixelMetric(static_cast<QStyle::PixelMetric>(metric), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

void QProxyStyle_Polish3(QtObjectPtr ptr, QtObjectPtr app){
	static_cast<QProxyStyle*>(ptr)->polish(static_cast<QApplication*>(app));
}

void QProxyStyle_Polish2(QtObjectPtr ptr, QtObjectPtr pal){
	static_cast<QProxyStyle*>(ptr)->polish(*static_cast<QPalette*>(pal));
}

void QProxyStyle_Polish(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QProxyStyle*>(ptr)->polish(static_cast<QWidget*>(widget));
}

void QProxyStyle_SetBaseStyle(QtObjectPtr ptr, QtObjectPtr style){
	static_cast<QProxyStyle*>(ptr)->setBaseStyle(static_cast<QStyle*>(style));
}

int QProxyStyle_StyleHint(QtObjectPtr ptr, int hint, QtObjectPtr option, QtObjectPtr widget, QtObjectPtr returnData){
	return static_cast<QProxyStyle*>(ptr)->styleHint(static_cast<QStyle::StyleHint>(hint), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget), static_cast<QStyleHintReturn*>(returnData));
}

void QProxyStyle_Unpolish2(QtObjectPtr ptr, QtObjectPtr app){
	static_cast<QProxyStyle*>(ptr)->unpolish(static_cast<QApplication*>(app));
}

void QProxyStyle_Unpolish(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QProxyStyle*>(ptr)->unpolish(static_cast<QWidget*>(widget));
}

void QProxyStyle_DestroyQProxyStyle(QtObjectPtr ptr){
	static_cast<QProxyStyle*>(ptr)->~QProxyStyle();
}

