#include "qgraphicsproxywidget.h"
#include <QString>
#include <QModelIndex>
#include <QStyle>
#include <QPainter>
#include <QRect>
#include <QRectF>
#include <QStyleOption>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionGraphicsItem>
#include <QWidget>
#include <QGraphicsItem>
#include <QGraphicsProxyWidget>
#include "_cgo_export.h"

class MyQGraphicsProxyWidget: public QGraphicsProxyWidget {
public:
};

QtObjectPtr QGraphicsProxyWidget_NewQGraphicsProxyWidget(QtObjectPtr parent, int wFlags){
	return new QGraphicsProxyWidget(static_cast<QGraphicsItem*>(parent), static_cast<Qt::WindowType>(wFlags));
}

QtObjectPtr QGraphicsProxyWidget_CreateProxyForChildWidget(QtObjectPtr ptr, QtObjectPtr child){
	return static_cast<QGraphicsProxyWidget*>(ptr)->createProxyForChildWidget(static_cast<QWidget*>(child));
}

void QGraphicsProxyWidget_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsProxyWidget*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsProxyWidget_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QGraphicsProxyWidget*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsProxyWidget_SetWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QGraphicsProxyWidget*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

int QGraphicsProxyWidget_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsProxyWidget*>(ptr)->type();
}

QtObjectPtr QGraphicsProxyWidget_Widget(QtObjectPtr ptr){
	return static_cast<QGraphicsProxyWidget*>(ptr)->widget();
}

void QGraphicsProxyWidget_DestroyQGraphicsProxyWidget(QtObjectPtr ptr){
	static_cast<QGraphicsProxyWidget*>(ptr)->~QGraphicsProxyWidget();
}

