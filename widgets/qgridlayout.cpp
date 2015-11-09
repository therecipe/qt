#include "qgridlayout.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QLayoutItem>
#include <QRect>
#include <QLayout>
#include <QGridLayout>
#include "_cgo_export.h"

class MyQGridLayout: public QGridLayout {
public:
};

int QGridLayout_HorizontalSpacing(void* ptr){
	return static_cast<QGridLayout*>(ptr)->horizontalSpacing();
}

void QGridLayout_SetHorizontalSpacing(void* ptr, int spacing){
	static_cast<QGridLayout*>(ptr)->setHorizontalSpacing(spacing);
}

void QGridLayout_SetVerticalSpacing(void* ptr, int spacing){
	static_cast<QGridLayout*>(ptr)->setVerticalSpacing(spacing);
}

int QGridLayout_VerticalSpacing(void* ptr){
	return static_cast<QGridLayout*>(ptr)->verticalSpacing();
}

void* QGridLayout_NewQGridLayout2(){
	return new QGridLayout();
}

void* QGridLayout_NewQGridLayout(void* parent){
	return new QGridLayout(static_cast<QWidget*>(parent));
}

void QGridLayout_AddItem(void* ptr, void* item, int row, int column, int rowSpan, int columnSpan, int alignment){
	static_cast<QGridLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddLayout(void* ptr, void* layout, int row, int column, int alignment){
	static_cast<QGridLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddLayout2(void* ptr, void* layout, int row, int column, int rowSpan, int columnSpan, int alignment){
	static_cast<QGridLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddWidget2(void* ptr, void* widget, int fromRow, int fromColumn, int rowSpan, int columnSpan, int alignment){
	static_cast<QGridLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), fromRow, fromColumn, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddWidget(void* ptr, void* widget, int row, int column, int alignment){
	static_cast<QGridLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

int QGridLayout_ColumnCount(void* ptr){
	return static_cast<QGridLayout*>(ptr)->columnCount();
}

int QGridLayout_ColumnMinimumWidth(void* ptr, int column){
	return static_cast<QGridLayout*>(ptr)->columnMinimumWidth(column);
}

int QGridLayout_ColumnStretch(void* ptr, int column){
	return static_cast<QGridLayout*>(ptr)->columnStretch(column);
}

int QGridLayout_Count(void* ptr){
	return static_cast<QGridLayout*>(ptr)->count();
}

int QGridLayout_ExpandingDirections(void* ptr){
	return static_cast<QGridLayout*>(ptr)->expandingDirections();
}

void QGridLayout_GetItemPosition(void* ptr, int index, int row, int column, int rowSpan, int columnSpan){
	static_cast<QGridLayout*>(ptr)->getItemPosition(index, &row, &column, &rowSpan, &columnSpan);
}

int QGridLayout_HasHeightForWidth(void* ptr){
	return static_cast<QGridLayout*>(ptr)->hasHeightForWidth();
}

int QGridLayout_HeightForWidth(void* ptr, int w){
	return static_cast<QGridLayout*>(ptr)->heightForWidth(w);
}

void QGridLayout_Invalidate(void* ptr){
	static_cast<QGridLayout*>(ptr)->invalidate();
}

void* QGridLayout_ItemAt(void* ptr, int index){
	return static_cast<QGridLayout*>(ptr)->itemAt(index);
}

void* QGridLayout_ItemAtPosition(void* ptr, int row, int column){
	return static_cast<QGridLayout*>(ptr)->itemAtPosition(row, column);
}

int QGridLayout_MinimumHeightForWidth(void* ptr, int w){
	return static_cast<QGridLayout*>(ptr)->minimumHeightForWidth(w);
}

int QGridLayout_OriginCorner(void* ptr){
	return static_cast<QGridLayout*>(ptr)->originCorner();
}

int QGridLayout_RowCount(void* ptr){
	return static_cast<QGridLayout*>(ptr)->rowCount();
}

int QGridLayout_RowMinimumHeight(void* ptr, int row){
	return static_cast<QGridLayout*>(ptr)->rowMinimumHeight(row);
}

int QGridLayout_RowStretch(void* ptr, int row){
	return static_cast<QGridLayout*>(ptr)->rowStretch(row);
}

void QGridLayout_SetColumnMinimumWidth(void* ptr, int column, int minSize){
	static_cast<QGridLayout*>(ptr)->setColumnMinimumWidth(column, minSize);
}

void QGridLayout_SetColumnStretch(void* ptr, int column, int stretch){
	static_cast<QGridLayout*>(ptr)->setColumnStretch(column, stretch);
}

void QGridLayout_SetGeometry(void* ptr, void* rect){
	static_cast<QGridLayout*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void QGridLayout_SetOriginCorner(void* ptr, int corner){
	static_cast<QGridLayout*>(ptr)->setOriginCorner(static_cast<Qt::Corner>(corner));
}

void QGridLayout_SetRowMinimumHeight(void* ptr, int row, int minSize){
	static_cast<QGridLayout*>(ptr)->setRowMinimumHeight(row, minSize);
}

void QGridLayout_SetRowStretch(void* ptr, int row, int stretch){
	static_cast<QGridLayout*>(ptr)->setRowStretch(row, stretch);
}

void QGridLayout_SetSpacing(void* ptr, int spacing){
	static_cast<QGridLayout*>(ptr)->setSpacing(spacing);
}

int QGridLayout_Spacing(void* ptr){
	return static_cast<QGridLayout*>(ptr)->spacing();
}

void* QGridLayout_TakeAt(void* ptr, int index){
	return static_cast<QGridLayout*>(ptr)->takeAt(index);
}

void QGridLayout_DestroyQGridLayout(void* ptr){
	static_cast<QGridLayout*>(ptr)->~QGridLayout();
}

