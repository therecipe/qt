#include "qformlayout.h"
#include <QModelIndex>
#include <QLayout>
#include <QLayoutItem>
#include <QRect>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QFormLayout>
#include "_cgo_export.h"

class MyQFormLayout: public QFormLayout {
public:
};

int QFormLayout_FieldGrowthPolicy(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->fieldGrowthPolicy();
}

int QFormLayout_FormAlignment(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->formAlignment();
}

int QFormLayout_HorizontalSpacing(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->horizontalSpacing();
}

int QFormLayout_LabelAlignment(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->labelAlignment();
}

int QFormLayout_RowWrapPolicy(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->rowWrapPolicy();
}

void QFormLayout_SetFieldGrowthPolicy(QtObjectPtr ptr, int policy){
	static_cast<QFormLayout*>(ptr)->setFieldGrowthPolicy(static_cast<QFormLayout::FieldGrowthPolicy>(policy));
}

void QFormLayout_SetFormAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QFormLayout*>(ptr)->setFormAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QFormLayout_SetHorizontalSpacing(QtObjectPtr ptr, int spacing){
	static_cast<QFormLayout*>(ptr)->setHorizontalSpacing(spacing);
}

void QFormLayout_SetLabelAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QFormLayout*>(ptr)->setLabelAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QFormLayout_SetRowWrapPolicy(QtObjectPtr ptr, int policy){
	static_cast<QFormLayout*>(ptr)->setRowWrapPolicy(static_cast<QFormLayout::RowWrapPolicy>(policy));
}

void QFormLayout_SetVerticalSpacing(QtObjectPtr ptr, int spacing){
	static_cast<QFormLayout*>(ptr)->setVerticalSpacing(spacing);
}

int QFormLayout_VerticalSpacing(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->verticalSpacing();
}

QtObjectPtr QFormLayout_NewQFormLayout(QtObjectPtr parent){
	return new QFormLayout(static_cast<QWidget*>(parent));
}

void QFormLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QFormLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

void QFormLayout_AddRow6(QtObjectPtr ptr, QtObjectPtr layout){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QLayout*>(layout));
}

void QFormLayout_AddRow2(QtObjectPtr ptr, QtObjectPtr label, QtObjectPtr field){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QWidget*>(label), static_cast<QLayout*>(field));
}

void QFormLayout_AddRow(QtObjectPtr ptr, QtObjectPtr label, QtObjectPtr field){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QWidget*>(label), static_cast<QWidget*>(field));
}

void QFormLayout_AddRow5(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QWidget*>(widget));
}

void QFormLayout_AddRow4(QtObjectPtr ptr, char* labelText, QtObjectPtr field){
	static_cast<QFormLayout*>(ptr)->addRow(QString(labelText), static_cast<QLayout*>(field));
}

void QFormLayout_AddRow3(QtObjectPtr ptr, char* labelText, QtObjectPtr field){
	static_cast<QFormLayout*>(ptr)->addRow(QString(labelText), static_cast<QWidget*>(field));
}

int QFormLayout_Count(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->count();
}

int QFormLayout_ExpandingDirections(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->expandingDirections();
}

int QFormLayout_HasHeightForWidth(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->hasHeightForWidth();
}

int QFormLayout_HeightForWidth(QtObjectPtr ptr, int width){
	return static_cast<QFormLayout*>(ptr)->heightForWidth(width);
}

void QFormLayout_InsertRow6(QtObjectPtr ptr, int row, QtObjectPtr layout){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QLayout*>(layout));
}

void QFormLayout_InsertRow2(QtObjectPtr ptr, int row, QtObjectPtr label, QtObjectPtr field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QWidget*>(label), static_cast<QLayout*>(field));
}

void QFormLayout_InsertRow(QtObjectPtr ptr, int row, QtObjectPtr label, QtObjectPtr field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QWidget*>(label), static_cast<QWidget*>(field));
}

void QFormLayout_InsertRow5(QtObjectPtr ptr, int row, QtObjectPtr widget){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QWidget*>(widget));
}

void QFormLayout_InsertRow4(QtObjectPtr ptr, int row, char* labelText, QtObjectPtr field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, QString(labelText), static_cast<QLayout*>(field));
}

void QFormLayout_InsertRow3(QtObjectPtr ptr, int row, char* labelText, QtObjectPtr field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, QString(labelText), static_cast<QWidget*>(field));
}

void QFormLayout_Invalidate(QtObjectPtr ptr){
	static_cast<QFormLayout*>(ptr)->invalidate();
}

QtObjectPtr QFormLayout_ItemAt2(QtObjectPtr ptr, int index){
	return static_cast<QFormLayout*>(ptr)->itemAt(index);
}

QtObjectPtr QFormLayout_ItemAt(QtObjectPtr ptr, int row, int role){
	return static_cast<QFormLayout*>(ptr)->itemAt(row, static_cast<QFormLayout::ItemRole>(role));
}

QtObjectPtr QFormLayout_LabelForField2(QtObjectPtr ptr, QtObjectPtr field){
	return static_cast<QFormLayout*>(ptr)->labelForField(static_cast<QLayout*>(field));
}

QtObjectPtr QFormLayout_LabelForField(QtObjectPtr ptr, QtObjectPtr field){
	return static_cast<QFormLayout*>(ptr)->labelForField(static_cast<QWidget*>(field));
}

int QFormLayout_RowCount(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->rowCount();
}

void QFormLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QFormLayout*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void QFormLayout_SetItem(QtObjectPtr ptr, int row, int role, QtObjectPtr item){
	static_cast<QFormLayout*>(ptr)->setItem(row, static_cast<QFormLayout::ItemRole>(role), static_cast<QLayoutItem*>(item));
}

void QFormLayout_SetLayout(QtObjectPtr ptr, int row, int role, QtObjectPtr layout){
	static_cast<QFormLayout*>(ptr)->setLayout(row, static_cast<QFormLayout::ItemRole>(role), static_cast<QLayout*>(layout));
}

void QFormLayout_SetSpacing(QtObjectPtr ptr, int spacing){
	static_cast<QFormLayout*>(ptr)->setSpacing(spacing);
}

void QFormLayout_SetWidget(QtObjectPtr ptr, int row, int role, QtObjectPtr widget){
	static_cast<QFormLayout*>(ptr)->setWidget(row, static_cast<QFormLayout::ItemRole>(role), static_cast<QWidget*>(widget));
}

int QFormLayout_Spacing(QtObjectPtr ptr){
	return static_cast<QFormLayout*>(ptr)->spacing();
}

QtObjectPtr QFormLayout_TakeAt(QtObjectPtr ptr, int index){
	return static_cast<QFormLayout*>(ptr)->takeAt(index);
}

void QFormLayout_DestroyQFormLayout(QtObjectPtr ptr){
	static_cast<QFormLayout*>(ptr)->~QFormLayout();
}

