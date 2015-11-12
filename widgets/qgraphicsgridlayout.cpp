#include "qgraphicsgridlayout.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsLayout>
#include <QGraphicsLayoutItem>
#include <QRectF>
#include <QRect>
#include <QString>
#include <QGraphicsGridLayout>
#include "_cgo_export.h"

class MyQGraphicsGridLayout: public QGraphicsGridLayout {
public:
};

void* QGraphicsGridLayout_NewQGraphicsGridLayout(void* parent){
	return new QGraphicsGridLayout(static_cast<QGraphicsLayoutItem*>(parent));
}

void QGraphicsGridLayout_AddItem2(void* ptr, void* item, int row, int column, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->addItem(static_cast<QGraphicsLayoutItem*>(item), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_AddItem(void* ptr, void* item, int row, int column, int rowSpan, int columnSpan, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->addItem(static_cast<QGraphicsLayoutItem*>(item), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

int QGraphicsGridLayout_Alignment(void* ptr, void* item){
	return static_cast<QGraphicsGridLayout*>(ptr)->alignment(static_cast<QGraphicsLayoutItem*>(item));
}

int QGraphicsGridLayout_ColumnAlignment(void* ptr, int column){
	return static_cast<QGraphicsGridLayout*>(ptr)->columnAlignment(column);
}

int QGraphicsGridLayout_ColumnCount(void* ptr){
	return static_cast<QGraphicsGridLayout*>(ptr)->columnCount();
}

double QGraphicsGridLayout_ColumnMaximumWidth(void* ptr, int column){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->columnMaximumWidth(column));
}

double QGraphicsGridLayout_ColumnMinimumWidth(void* ptr, int column){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->columnMinimumWidth(column));
}

double QGraphicsGridLayout_ColumnPreferredWidth(void* ptr, int column){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->columnPreferredWidth(column));
}

double QGraphicsGridLayout_ColumnSpacing(void* ptr, int column){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->columnSpacing(column));
}

int QGraphicsGridLayout_ColumnStretchFactor(void* ptr, int column){
	return static_cast<QGraphicsGridLayout*>(ptr)->columnStretchFactor(column);
}

int QGraphicsGridLayout_Count(void* ptr){
	return static_cast<QGraphicsGridLayout*>(ptr)->count();
}

double QGraphicsGridLayout_HorizontalSpacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->horizontalSpacing());
}

void QGraphicsGridLayout_Invalidate(void* ptr){
	static_cast<QGraphicsGridLayout*>(ptr)->invalidate();
}

void* QGraphicsGridLayout_ItemAt2(void* ptr, int index){
	return static_cast<QGraphicsGridLayout*>(ptr)->itemAt(index);
}

void* QGraphicsGridLayout_ItemAt(void* ptr, int row, int column){
	return static_cast<QGraphicsGridLayout*>(ptr)->itemAt(row, column);
}

void QGraphicsGridLayout_RemoveAt(void* ptr, int index){
	static_cast<QGraphicsGridLayout*>(ptr)->removeAt(index);
}

void QGraphicsGridLayout_RemoveItem(void* ptr, void* item){
	static_cast<QGraphicsGridLayout*>(ptr)->removeItem(static_cast<QGraphicsLayoutItem*>(item));
}

int QGraphicsGridLayout_RowAlignment(void* ptr, int row){
	return static_cast<QGraphicsGridLayout*>(ptr)->rowAlignment(row);
}

int QGraphicsGridLayout_RowCount(void* ptr){
	return static_cast<QGraphicsGridLayout*>(ptr)->rowCount();
}

double QGraphicsGridLayout_RowMaximumHeight(void* ptr, int row){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->rowMaximumHeight(row));
}

double QGraphicsGridLayout_RowMinimumHeight(void* ptr, int row){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->rowMinimumHeight(row));
}

double QGraphicsGridLayout_RowPreferredHeight(void* ptr, int row){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->rowPreferredHeight(row));
}

double QGraphicsGridLayout_RowSpacing(void* ptr, int row){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->rowSpacing(row));
}

int QGraphicsGridLayout_RowStretchFactor(void* ptr, int row){
	return static_cast<QGraphicsGridLayout*>(ptr)->rowStretchFactor(row);
}

void QGraphicsGridLayout_SetAlignment(void* ptr, void* item, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->setAlignment(static_cast<QGraphicsLayoutItem*>(item), static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_SetColumnAlignment(void* ptr, int column, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnAlignment(column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_SetColumnFixedWidth(void* ptr, int column, double width){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnFixedWidth(column, static_cast<qreal>(width));
}

void QGraphicsGridLayout_SetColumnMaximumWidth(void* ptr, int column, double width){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnMaximumWidth(column, static_cast<qreal>(width));
}

void QGraphicsGridLayout_SetColumnMinimumWidth(void* ptr, int column, double width){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnMinimumWidth(column, static_cast<qreal>(width));
}

void QGraphicsGridLayout_SetColumnPreferredWidth(void* ptr, int column, double width){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnPreferredWidth(column, static_cast<qreal>(width));
}

void QGraphicsGridLayout_SetColumnSpacing(void* ptr, int column, double spacing){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnSpacing(column, static_cast<qreal>(spacing));
}

void QGraphicsGridLayout_SetColumnStretchFactor(void* ptr, int column, int stretch){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnStretchFactor(column, stretch);
}

void QGraphicsGridLayout_SetGeometry(void* ptr, void* rect){
	static_cast<QGraphicsGridLayout*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsGridLayout_SetHorizontalSpacing(void* ptr, double spacing){
	static_cast<QGraphicsGridLayout*>(ptr)->setHorizontalSpacing(static_cast<qreal>(spacing));
}

void QGraphicsGridLayout_SetRowAlignment(void* ptr, int row, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowAlignment(row, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_SetRowFixedHeight(void* ptr, int row, double height){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowFixedHeight(row, static_cast<qreal>(height));
}

void QGraphicsGridLayout_SetRowMaximumHeight(void* ptr, int row, double height){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowMaximumHeight(row, static_cast<qreal>(height));
}

void QGraphicsGridLayout_SetRowMinimumHeight(void* ptr, int row, double height){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowMinimumHeight(row, static_cast<qreal>(height));
}

void QGraphicsGridLayout_SetRowPreferredHeight(void* ptr, int row, double height){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowPreferredHeight(row, static_cast<qreal>(height));
}

void QGraphicsGridLayout_SetRowSpacing(void* ptr, int row, double spacing){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowSpacing(row, static_cast<qreal>(spacing));
}

void QGraphicsGridLayout_SetRowStretchFactor(void* ptr, int row, int stretch){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowStretchFactor(row, stretch);
}

void QGraphicsGridLayout_SetSpacing(void* ptr, double spacing){
	static_cast<QGraphicsGridLayout*>(ptr)->setSpacing(static_cast<qreal>(spacing));
}

void QGraphicsGridLayout_SetVerticalSpacing(void* ptr, double spacing){
	static_cast<QGraphicsGridLayout*>(ptr)->setVerticalSpacing(static_cast<qreal>(spacing));
}

double QGraphicsGridLayout_VerticalSpacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->verticalSpacing());
}

void QGraphicsGridLayout_DestroyQGraphicsGridLayout(void* ptr){
	static_cast<QGraphicsGridLayout*>(ptr)->~QGraphicsGridLayout();
}

