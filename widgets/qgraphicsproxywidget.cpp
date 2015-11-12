#include "qgraphicsproxywidget.h"
#include <QRectF>
#include <QRect>
#include <QPainter>
#include <QStyle>
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsItem>
#include <QGraphicsProxyWidget>
#include "_cgo_export.h"

class MyQGraphicsProxyWidget: public QGraphicsProxyWidget {
public:
};

void* QGraphicsProxyWidget_NewQGraphicsProxyWidget(void* parent, int wFlags){
	return new QGraphicsProxyWidget(static_cast<QGraphicsItem*>(parent), static_cast<Qt::WindowType>(wFlags));
}

void* QGraphicsProxyWidget_CreateProxyForChildWidget(void* ptr, void* child){
	return static_cast<QGraphicsProxyWidget*>(ptr)->createProxyForChildWidget(static_cast<QWidget*>(child));
}

void QGraphicsProxyWidget_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsProxyWidget*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsProxyWidget_SetGeometry(void* ptr, void* rect){
	static_cast<QGraphicsProxyWidget*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsProxyWidget_SetWidget(void* ptr, void* widget){
	static_cast<QGraphicsProxyWidget*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

int QGraphicsProxyWidget_Type(void* ptr){
	return static_cast<QGraphicsProxyWidget*>(ptr)->type();
}

void* QGraphicsProxyWidget_Widget(void* ptr){
	return static_cast<QGraphicsProxyWidget*>(ptr)->widget();
}

void QGraphicsProxyWidget_DestroyQGraphicsProxyWidget(void* ptr){
	static_cast<QGraphicsProxyWidget*>(ptr)->~QGraphicsProxyWidget();
}

