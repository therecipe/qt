#include "qtooltip.h"
#include <QVariant>
#include <QUrl>
#include <QRect>
#include <QPalette>
#include <QPoint>
#include <QFont>
#include <QWidget>
#include <QString>
#include <QModelIndex>
#include <QToolTip>
#include "_cgo_export.h"

class MyQToolTip: public QToolTip {
public:
};

void QToolTip_QToolTip_HideText(){
	QToolTip::hideText();
}

int QToolTip_QToolTip_IsVisible(){
	return QToolTip::isVisible();
}

void QToolTip_QToolTip_SetFont(void* font){
	QToolTip::setFont(*static_cast<QFont*>(font));
}

void QToolTip_QToolTip_SetPalette(void* palette){
	QToolTip::setPalette(*static_cast<QPalette*>(palette));
}

void QToolTip_QToolTip_ShowText3(void* pos, char* text, void* w){
	QToolTip::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w));
}

void QToolTip_QToolTip_ShowText(void* pos, char* text, void* w, void* rect){
	QToolTip::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w), *static_cast<QRect*>(rect));
}

void QToolTip_QToolTip_ShowText2(void* pos, char* text, void* w, void* rect, int msecDisplayTime){
	QToolTip::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w), *static_cast<QRect*>(rect), msecDisplayTime);
}

char* QToolTip_QToolTip_Text(){
	return QToolTip::text().toUtf8().data();
}

