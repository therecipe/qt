#include "qgraphicsgridlayout.h"
#include <QRectF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsLayout>
#include <QGraphicsLayoutItem>
#include <QRect>
#include <QGraphicsGridLayout>
#include "_cgo_export.h"

class MyQGraphicsGridLayout: public QGraphicsGridLayout {
public:
};

QtObjectPtr QGraphicsGridLayout_NewQGraphicsGridLayout(QtObjectPtr parent){
	return new QGraphicsGridLayout(static_cast<QGraphicsLayoutItem*>(parent));
}

void QGraphicsGridLayout_AddItem2(QtObjectPtr ptr, QtObjectPtr item, int row, int column, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->addItem(static_cast<QGraphicsLayoutItem*>(item), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item, int row, int column, int rowSpan, int columnSpan, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->addItem(static_cast<QGraphicsLayoutItem*>(item), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

int QGraphicsGridLayout_Alignment(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsGridLayout*>(ptr)->alignment(static_cast<QGraphicsLayoutItem*>(item));
}

int QGraphicsGridLayout_ColumnAlignment(QtObjectPtr ptr, int column){
	return static_cast<QGraphicsGridLayout*>(ptr)->columnAlignment(column);
}

int QGraphicsGridLayout_ColumnCount(QtObjectPtr ptr){
	return static_cast<QGraphicsGridLayout*>(ptr)->columnCount();
}

int QGraphicsGridLayout_ColumnStretchFactor(QtObjectPtr ptr, int column){
	return static_cast<QGraphicsGridLayout*>(ptr)->columnStretchFactor(column);
}

int QGraphicsGridLayout_Count(QtObjectPtr ptr){
	return static_cast<QGraphicsGridLayout*>(ptr)->count();
}

void QGraphicsGridLayout_Invalidate(QtObjectPtr ptr){
	static_cast<QGraphicsGridLayout*>(ptr)->invalidate();
}

QtObjectPtr QGraphicsGridLayout_ItemAt2(QtObjectPtr ptr, int index){
	return static_cast<QGraphicsGridLayout*>(ptr)->itemAt(index);
}

QtObjectPtr QGraphicsGridLayout_ItemAt(QtObjectPtr ptr, int row, int column){
	return static_cast<QGraphicsGridLayout*>(ptr)->itemAt(row, column);
}

void QGraphicsGridLayout_RemoveAt(QtObjectPtr ptr, int index){
	static_cast<QGraphicsGridLayout*>(ptr)->removeAt(index);
}

void QGraphicsGridLayout_RemoveItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QGraphicsGridLayout*>(ptr)->removeItem(static_cast<QGraphicsLayoutItem*>(item));
}

int QGraphicsGridLayout_RowAlignment(QtObjectPtr ptr, int row){
	return static_cast<QGraphicsGridLayout*>(ptr)->rowAlignment(row);
}

int QGraphicsGridLayout_RowCount(QtObjectPtr ptr){
	return static_cast<QGraphicsGridLayout*>(ptr)->rowCount();
}

int QGraphicsGridLayout_RowStretchFactor(QtObjectPtr ptr, int row){
	return static_cast<QGraphicsGridLayout*>(ptr)->rowStretchFactor(row);
}

void QGraphicsGridLayout_SetAlignment(QtObjectPtr ptr, QtObjectPtr item, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->setAlignment(static_cast<QGraphicsLayoutItem*>(item), static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_SetColumnAlignment(QtObjectPtr ptr, int column, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnAlignment(column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_SetColumnStretchFactor(QtObjectPtr ptr, int column, int stretch){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnStretchFactor(column, stretch);
}

void QGraphicsGridLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QGraphicsGridLayout*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsGridLayout_SetRowAlignment(QtObjectPtr ptr, int row, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowAlignment(row, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_SetRowStretchFactor(QtObjectPtr ptr, int row, int stretch){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowStretchFactor(row, stretch);
}

void QGraphicsGridLayout_DestroyQGraphicsGridLayout(QtObjectPtr ptr){
	static_cast<QGraphicsGridLayout*>(ptr)->~QGraphicsGridLayout();
}

