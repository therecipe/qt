#include "qgraphicssvgitem.h"
#include <QSvgRenderer>
#include <QSize>
#include <QStyleOption>
#include <QWidget>
#include <QGraphicsItem>
#include <QStyleOptionGraphicsItem>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QPainter>
#include <QGraphicsSvgItem>
#include "_cgo_export.h"

class MyQGraphicsSvgItem: public QGraphicsSvgItem {
public:
};

QtObjectPtr QGraphicsSvgItem_NewQGraphicsSvgItem(QtObjectPtr parent){
	return new QGraphicsSvgItem(static_cast<QGraphicsItem*>(parent));
}

QtObjectPtr QGraphicsSvgItem_NewQGraphicsSvgItem2(char* fileName, QtObjectPtr parent){
	return new QGraphicsSvgItem(QString(fileName), static_cast<QGraphicsItem*>(parent));
}

char* QGraphicsSvgItem_ElementId(QtObjectPtr ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->elementId().toUtf8().data();
}

void QGraphicsSvgItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsSvgItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

QtObjectPtr QGraphicsSvgItem_Renderer(QtObjectPtr ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->renderer();
}

void QGraphicsSvgItem_SetElementId(QtObjectPtr ptr, char* id){
	static_cast<QGraphicsSvgItem*>(ptr)->setElementId(QString(id));
}

void QGraphicsSvgItem_SetMaximumCacheSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QGraphicsSvgItem*>(ptr)->setMaximumCacheSize(*static_cast<QSize*>(size));
}

void QGraphicsSvgItem_SetSharedRenderer(QtObjectPtr ptr, QtObjectPtr renderer){
	static_cast<QGraphicsSvgItem*>(ptr)->setSharedRenderer(static_cast<QSvgRenderer*>(renderer));
}

int QGraphicsSvgItem_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->type();
}

