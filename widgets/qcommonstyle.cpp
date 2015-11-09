#include "qcommonstyle.h"
#include <QApplication>
#include <QStyleOptionComplex>
#include <QStyleOption>
#include <QVariant>
#include <QWidget>
#include <QPoint>
#include <QStyle>
#include <QStyleHintReturn>
#include <QUrl>
#include <QModelIndex>
#include <QPalette>
#include <QSize>
#include <QPainter>
#include <QString>
#include <QSizePolicy>
#include <QCommonStyle>
#include "_cgo_export.h"

class MyQCommonStyle: public QCommonStyle {
public:
};

void QCommonStyle_DrawControl(void* ptr, int element, void* opt, void* p, void* widget){
	static_cast<QCommonStyle*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(element), static_cast<QStyleOption*>(opt), static_cast<QPainter*>(p), static_cast<QWidget*>(widget));
}

void QCommonStyle_DrawPrimitive(void* ptr, int pe, void* opt, void* p, void* widget){
	static_cast<QCommonStyle*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(pe), static_cast<QStyleOption*>(opt), static_cast<QPainter*>(p), static_cast<QWidget*>(widget));
}

void QCommonStyle_DrawComplexControl(void* ptr, int cc, void* opt, void* p, void* widget){
	static_cast<QCommonStyle*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(cc), static_cast<QStyleOptionComplex*>(opt), static_cast<QPainter*>(p), static_cast<QWidget*>(widget));
}

int QCommonStyle_HitTestComplexControl(void* ptr, int cc, void* opt, void* pt, void* widget){
	return static_cast<QCommonStyle*>(ptr)->hitTestComplexControl(static_cast<QStyle::ComplexControl>(cc), static_cast<QStyleOptionComplex*>(opt), *static_cast<QPoint*>(pt), static_cast<QWidget*>(widget));
}

int QCommonStyle_LayoutSpacing(void* ptr, int control1, int control2, int orientation, void* option, void* widget){
	return static_cast<QCommonStyle*>(ptr)->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

int QCommonStyle_PixelMetric(void* ptr, int m, void* opt, void* widget){
	return static_cast<QCommonStyle*>(ptr)->pixelMetric(static_cast<QStyle::PixelMetric>(m), static_cast<QStyleOption*>(opt), static_cast<QWidget*>(widget));
}

void QCommonStyle_Polish2(void* ptr, void* app){
	static_cast<QCommonStyle*>(ptr)->polish(static_cast<QApplication*>(app));
}

void QCommonStyle_Polish(void* ptr, void* pal){
	static_cast<QCommonStyle*>(ptr)->polish(*static_cast<QPalette*>(pal));
}

void QCommonStyle_Polish3(void* ptr, void* widget){
	static_cast<QCommonStyle*>(ptr)->polish(static_cast<QWidget*>(widget));
}

int QCommonStyle_StyleHint(void* ptr, int sh, void* opt, void* widget, void* hret){
	return static_cast<QCommonStyle*>(ptr)->styleHint(static_cast<QStyle::StyleHint>(sh), static_cast<QStyleOption*>(opt), static_cast<QWidget*>(widget), static_cast<QStyleHintReturn*>(hret));
}

void QCommonStyle_Unpolish2(void* ptr, void* application){
	static_cast<QCommonStyle*>(ptr)->unpolish(static_cast<QApplication*>(application));
}

void QCommonStyle_Unpolish(void* ptr, void* widget){
	static_cast<QCommonStyle*>(ptr)->unpolish(static_cast<QWidget*>(widget));
}

void QCommonStyle_DestroyQCommonStyle(void* ptr){
	static_cast<QCommonStyle*>(ptr)->~QCommonStyle();
}

