#include "qcommonstyle.h"
#include <QModelIndex>
#include <QPainter>
#include <QStyleOption>
#include <QApplication>
#include <QSize>
#include <QSizePolicy>
#include <QPoint>
#include <QPalette>
#include <QStyleHintReturn>
#include <QVariant>
#include <QStyle>
#include <QWidget>
#include <QStyleOptionComplex>
#include <QString>
#include <QUrl>
#include <QCommonStyle>
#include "_cgo_export.h"

class MyQCommonStyle: public QCommonStyle {
public:
};

void QCommonStyle_DrawControl(QtObjectPtr ptr, int element, QtObjectPtr opt, QtObjectPtr p, QtObjectPtr widget){
	static_cast<QCommonStyle*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(element), static_cast<QStyleOption*>(opt), static_cast<QPainter*>(p), static_cast<QWidget*>(widget));
}

void QCommonStyle_DrawPrimitive(QtObjectPtr ptr, int pe, QtObjectPtr opt, QtObjectPtr p, QtObjectPtr widget){
	static_cast<QCommonStyle*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(pe), static_cast<QStyleOption*>(opt), static_cast<QPainter*>(p), static_cast<QWidget*>(widget));
}

void QCommonStyle_DrawComplexControl(QtObjectPtr ptr, int cc, QtObjectPtr opt, QtObjectPtr p, QtObjectPtr widget){
	static_cast<QCommonStyle*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(cc), static_cast<QStyleOptionComplex*>(opt), static_cast<QPainter*>(p), static_cast<QWidget*>(widget));
}

int QCommonStyle_HitTestComplexControl(QtObjectPtr ptr, int cc, QtObjectPtr opt, QtObjectPtr pt, QtObjectPtr widget){
	return static_cast<QCommonStyle*>(ptr)->hitTestComplexControl(static_cast<QStyle::ComplexControl>(cc), static_cast<QStyleOptionComplex*>(opt), *static_cast<QPoint*>(pt), static_cast<QWidget*>(widget));
}

int QCommonStyle_LayoutSpacing(QtObjectPtr ptr, int control1, int control2, int orientation, QtObjectPtr option, QtObjectPtr widget){
	return static_cast<QCommonStyle*>(ptr)->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

int QCommonStyle_PixelMetric(QtObjectPtr ptr, int m, QtObjectPtr opt, QtObjectPtr widget){
	return static_cast<QCommonStyle*>(ptr)->pixelMetric(static_cast<QStyle::PixelMetric>(m), static_cast<QStyleOption*>(opt), static_cast<QWidget*>(widget));
}

void QCommonStyle_Polish2(QtObjectPtr ptr, QtObjectPtr app){
	static_cast<QCommonStyle*>(ptr)->polish(static_cast<QApplication*>(app));
}

void QCommonStyle_Polish(QtObjectPtr ptr, QtObjectPtr pal){
	static_cast<QCommonStyle*>(ptr)->polish(*static_cast<QPalette*>(pal));
}

void QCommonStyle_Polish3(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QCommonStyle*>(ptr)->polish(static_cast<QWidget*>(widget));
}

int QCommonStyle_StyleHint(QtObjectPtr ptr, int sh, QtObjectPtr opt, QtObjectPtr widget, QtObjectPtr hret){
	return static_cast<QCommonStyle*>(ptr)->styleHint(static_cast<QStyle::StyleHint>(sh), static_cast<QStyleOption*>(opt), static_cast<QWidget*>(widget), static_cast<QStyleHintReturn*>(hret));
}

void QCommonStyle_Unpolish2(QtObjectPtr ptr, QtObjectPtr application){
	static_cast<QCommonStyle*>(ptr)->unpolish(static_cast<QApplication*>(application));
}

void QCommonStyle_Unpolish(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QCommonStyle*>(ptr)->unpolish(static_cast<QWidget*>(widget));
}

void QCommonStyle_DestroyQCommonStyle(QtObjectPtr ptr){
	static_cast<QCommonStyle*>(ptr)->~QCommonStyle();
}

