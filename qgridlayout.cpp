#include "qgridlayout.h"
#include <QGridLayout>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QGridLayout_New_QWidget(QtObjectPtr parent)
{
	return new QGridLayout(((QWidget*)(parent)));
}

QtObjectPtr QGridLayout_New()
{
	return new QGridLayout();
}

void QGridLayout_Destroy(QtObjectPtr ptr)
{
	((QGridLayout*)(ptr))->~QGridLayout();
}

void QGridLayout_AddItem_QLayoutItem_Int_Int_Int_Int_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr item, int row, int column, int rowSpan, int columnSpan, int alignment)
{
	((QGridLayout*)(ptr))->addItem(((QLayoutItem*)(item)), row, column, rowSpan, columnSpan, ((Qt::AlignmentFlag)(alignment)));
}

void QGridLayout_AddLayout_QLayout_Int_Int_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr layout, int row, int column, int alignment)
{
	((QGridLayout*)(ptr))->addLayout(((QLayout*)(layout)), row, column, ((Qt::AlignmentFlag)(alignment)));
}

void QGridLayout_AddLayout_QLayout_Int_Int_Int_Int_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr layout, int row, int column, int rowSpan, int columnSpan, int alignment)
{
	((QGridLayout*)(ptr))->addLayout(((QLayout*)(layout)), row, column, rowSpan, columnSpan, ((Qt::AlignmentFlag)(alignment)));
}

void QGridLayout_AddWidget_QWidget_Int_Int_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr widget, int row, int column, int alignment)
{
	((QGridLayout*)(ptr))->addWidget(((QWidget*)(widget)), row, column, ((Qt::AlignmentFlag)(alignment)));
}

void QGridLayout_AddWidget_QWidget_Int_Int_Int_Int_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr widget, int fromRow, int fromColumn, int rowSpan, int columnSpan, int alignment)
{
	((QGridLayout*)(ptr))->addWidget(((QWidget*)(widget)), fromRow, fromColumn, rowSpan, columnSpan, ((Qt::AlignmentFlag)(alignment)));
}

int QGridLayout_ColumnCount(QtObjectPtr ptr)
{
	return ((QGridLayout*)(ptr))->columnCount();
}

int QGridLayout_ColumnMinimumWidth_Int(QtObjectPtr ptr, int column)
{
	return ((QGridLayout*)(ptr))->columnMinimumWidth(column);
}

int QGridLayout_ColumnStretch_Int(QtObjectPtr ptr, int column)
{
	return ((QGridLayout*)(ptr))->columnStretch(column);
}

int QGridLayout_HorizontalSpacing(QtObjectPtr ptr)
{
	return ((QGridLayout*)(ptr))->horizontalSpacing();
}

int QGridLayout_OriginCorner(QtObjectPtr ptr)
{
	return ((QGridLayout*)(ptr))->originCorner();
}

int QGridLayout_RowCount(QtObjectPtr ptr)
{
	return ((QGridLayout*)(ptr))->rowCount();
}

int QGridLayout_RowMinimumHeight_Int(QtObjectPtr ptr, int row)
{
	return ((QGridLayout*)(ptr))->rowMinimumHeight(row);
}

int QGridLayout_RowStretch_Int(QtObjectPtr ptr, int row)
{
	return ((QGridLayout*)(ptr))->rowStretch(row);
}

void QGridLayout_SetColumnMinimumWidth_Int_Int(QtObjectPtr ptr, int column, int minSize)
{
	((QGridLayout*)(ptr))->setColumnMinimumWidth(column, minSize);
}

void QGridLayout_SetColumnStretch_Int_Int(QtObjectPtr ptr, int column, int stretch)
{
	((QGridLayout*)(ptr))->setColumnStretch(column, stretch);
}

void QGridLayout_SetHorizontalSpacing_Int(QtObjectPtr ptr, int spacing)
{
	((QGridLayout*)(ptr))->setHorizontalSpacing(spacing);
}

void QGridLayout_SetOriginCorner_Corner(QtObjectPtr ptr, int corner)
{
	((QGridLayout*)(ptr))->setOriginCorner(((Qt::Corner)(corner)));
}

void QGridLayout_SetRowMinimumHeight_Int_Int(QtObjectPtr ptr, int row, int minSize)
{
	((QGridLayout*)(ptr))->setRowMinimumHeight(row, minSize);
}

void QGridLayout_SetRowStretch_Int_Int(QtObjectPtr ptr, int row, int stretch)
{
	((QGridLayout*)(ptr))->setRowStretch(row, stretch);
}

void QGridLayout_SetVerticalSpacing_Int(QtObjectPtr ptr, int spacing)
{
	((QGridLayout*)(ptr))->setVerticalSpacing(spacing);
}

int QGridLayout_VerticalSpacing(QtObjectPtr ptr)
{
	return ((QGridLayout*)(ptr))->verticalSpacing();
}

