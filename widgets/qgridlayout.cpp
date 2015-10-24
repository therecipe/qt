#include "qgridlayout.h"
#include <QRect>
#include <QWidget>
#include <QLayout>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLayoutItem>
#include <QGridLayout>
#include "_cgo_export.h"

class MyQGridLayout: public QGridLayout {
public:
};

int QGridLayout_HorizontalSpacing(QtObjectPtr ptr){
	return static_cast<QGridLayout*>(ptr)->horizontalSpacing();
}

void QGridLayout_SetHorizontalSpacing(QtObjectPtr ptr, int spacing){
	static_cast<QGridLayout*>(ptr)->setHorizontalSpacing(spacing);
}

void QGridLayout_SetVerticalSpacing(QtObjectPtr ptr, int spacing){
	static_cast<QGridLayout*>(ptr)->setVerticalSpacing(spacing);
}

int QGridLayout_VerticalSpacing(QtObjectPtr ptr){
	return static_cast<QGridLayout*>(ptr)->verticalSpacing();
}

QtObjectPtr QGridLayout_NewQGridLayout2(){
	return new QGridLayout();
}

QtObjectPtr QGridLayout_NewQGridLayout(QtObjectPtr parent){
	return new QGridLayout(static_cast<QWidget*>(parent));
}

void QGridLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item, int row, int column, int rowSpan, int columnSpan, int alignment){
	static_cast<QGridLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddLayout(QtObjectPtr ptr, QtObjectPtr layout, int row, int column, int alignment){
	static_cast<QGridLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddLayout2(QtObjectPtr ptr, QtObjectPtr layout, int row, int column, int rowSpan, int columnSpan, int alignment){
	static_cast<QGridLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddWidget2(QtObjectPtr ptr, QtObjectPtr widget, int fromRow, int fromColumn, int rowSpan, int columnSpan, int alignment){
	static_cast<QGridLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), fromRow, fromColumn, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddWidget(QtObjectPtr ptr, QtObjectPtr widget, int row, int column, int alignment){
	static_cast<QGridLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

int QGridLayout_ColumnCount(QtObjectPtr ptr){
	return static_cast<QGridLayout*>(ptr)->columnCount();
}

int QGridLayout_ColumnMinimumWidth(QtObjectPtr ptr, int column){
	return static_cast<QGridLayout*>(ptr)->columnMinimumWidth(column);
}

int QGridLayout_ColumnStretch(QtObjectPtr ptr, int column){
	return static_cast<QGridLayout*>(ptr)->columnStretch(column);
}

int QGridLayout_Count(QtObjectPtr ptr){
	return static_cast<QGridLayout*>(ptr)->count();
}

int QGridLayout_ExpandingDirections(QtObjectPtr ptr){
	return static_cast<QGridLayout*>(ptr)->expandingDirections();
}

void QGridLayout_GetItemPosition(QtObjectPtr ptr, int index, int row, int column, int rowSpan, int columnSpan){
	static_cast<QGridLayout*>(ptr)->getItemPosition(index, &row, &column, &rowSpan, &columnSpan);
}

int QGridLayout_HasHeightForWidth(QtObjectPtr ptr){
	return static_cast<QGridLayout*>(ptr)->hasHeightForWidth();
}

int QGridLayout_HeightForWidth(QtObjectPtr ptr, int w){
	return static_cast<QGridLayout*>(ptr)->heightForWidth(w);
}

void QGridLayout_Invalidate(QtObjectPtr ptr){
	static_cast<QGridLayout*>(ptr)->invalidate();
}

QtObjectPtr QGridLayout_ItemAt(QtObjectPtr ptr, int index){
	return static_cast<QGridLayout*>(ptr)->itemAt(index);
}

QtObjectPtr QGridLayout_ItemAtPosition(QtObjectPtr ptr, int row, int column){
	return static_cast<QGridLayout*>(ptr)->itemAtPosition(row, column);
}

int QGridLayout_MinimumHeightForWidth(QtObjectPtr ptr, int w){
	return static_cast<QGridLayout*>(ptr)->minimumHeightForWidth(w);
}

int QGridLayout_OriginCorner(QtObjectPtr ptr){
	return static_cast<QGridLayout*>(ptr)->originCorner();
}

int QGridLayout_RowCount(QtObjectPtr ptr){
	return static_cast<QGridLayout*>(ptr)->rowCount();
}

int QGridLayout_RowMinimumHeight(QtObjectPtr ptr, int row){
	return static_cast<QGridLayout*>(ptr)->rowMinimumHeight(row);
}

int QGridLayout_RowStretch(QtObjectPtr ptr, int row){
	return static_cast<QGridLayout*>(ptr)->rowStretch(row);
}

void QGridLayout_SetColumnMinimumWidth(QtObjectPtr ptr, int column, int minSize){
	static_cast<QGridLayout*>(ptr)->setColumnMinimumWidth(column, minSize);
}

void QGridLayout_SetColumnStretch(QtObjectPtr ptr, int column, int stretch){
	static_cast<QGridLayout*>(ptr)->setColumnStretch(column, stretch);
}

void QGridLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QGridLayout*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void QGridLayout_SetOriginCorner(QtObjectPtr ptr, int corner){
	static_cast<QGridLayout*>(ptr)->setOriginCorner(static_cast<Qt::Corner>(corner));
}

void QGridLayout_SetRowMinimumHeight(QtObjectPtr ptr, int row, int minSize){
	static_cast<QGridLayout*>(ptr)->setRowMinimumHeight(row, minSize);
}

void QGridLayout_SetRowStretch(QtObjectPtr ptr, int row, int stretch){
	static_cast<QGridLayout*>(ptr)->setRowStretch(row, stretch);
}

void QGridLayout_SetSpacing(QtObjectPtr ptr, int spacing){
	static_cast<QGridLayout*>(ptr)->setSpacing(spacing);
}

int QGridLayout_Spacing(QtObjectPtr ptr){
	return static_cast<QGridLayout*>(ptr)->spacing();
}

QtObjectPtr QGridLayout_TakeAt(QtObjectPtr ptr, int index){
	return static_cast<QGridLayout*>(ptr)->takeAt(index);
}

void QGridLayout_DestroyQGridLayout(QtObjectPtr ptr){
	static_cast<QGridLayout*>(ptr)->~QGridLayout();
}

