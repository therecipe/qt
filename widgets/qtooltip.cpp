#include "qtooltip.h"
#include <QUrl>
#include <QPoint>
#include <QRect>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QFont>
#include <QWidget>
#include <QPalette>
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

void QToolTip_QToolTip_SetFont(QtObjectPtr font){
	QToolTip::setFont(*static_cast<QFont*>(font));
}

void QToolTip_QToolTip_SetPalette(QtObjectPtr palette){
	QToolTip::setPalette(*static_cast<QPalette*>(palette));
}

void QToolTip_QToolTip_ShowText3(QtObjectPtr pos, char* text, QtObjectPtr w){
	QToolTip::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w));
}

void QToolTip_QToolTip_ShowText(QtObjectPtr pos, char* text, QtObjectPtr w, QtObjectPtr rect){
	QToolTip::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w), *static_cast<QRect*>(rect));
}

void QToolTip_QToolTip_ShowText2(QtObjectPtr pos, char* text, QtObjectPtr w, QtObjectPtr rect, int msecDisplayTime){
	QToolTip::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w), *static_cast<QRect*>(rect), msecDisplayTime);
}

char* QToolTip_QToolTip_Text(){
	return QToolTip::text().toUtf8().data();
}

