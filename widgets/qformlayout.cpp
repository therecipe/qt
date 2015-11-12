#include "qformlayout.h"
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QLayoutItem>
#include <QLayout>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QFormLayout>
#include "_cgo_export.h"

class MyQFormLayout: public QFormLayout {
public:
};

int QFormLayout_FieldGrowthPolicy(void* ptr){
	return static_cast<QFormLayout*>(ptr)->fieldGrowthPolicy();
}

int QFormLayout_FormAlignment(void* ptr){
	return static_cast<QFormLayout*>(ptr)->formAlignment();
}

int QFormLayout_HorizontalSpacing(void* ptr){
	return static_cast<QFormLayout*>(ptr)->horizontalSpacing();
}

int QFormLayout_LabelAlignment(void* ptr){
	return static_cast<QFormLayout*>(ptr)->labelAlignment();
}

int QFormLayout_RowWrapPolicy(void* ptr){
	return static_cast<QFormLayout*>(ptr)->rowWrapPolicy();
}

void QFormLayout_SetFieldGrowthPolicy(void* ptr, int policy){
	static_cast<QFormLayout*>(ptr)->setFieldGrowthPolicy(static_cast<QFormLayout::FieldGrowthPolicy>(policy));
}

void QFormLayout_SetFormAlignment(void* ptr, int alignment){
	static_cast<QFormLayout*>(ptr)->setFormAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QFormLayout_SetHorizontalSpacing(void* ptr, int spacing){
	static_cast<QFormLayout*>(ptr)->setHorizontalSpacing(spacing);
}

void QFormLayout_SetLabelAlignment(void* ptr, int alignment){
	static_cast<QFormLayout*>(ptr)->setLabelAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QFormLayout_SetRowWrapPolicy(void* ptr, int policy){
	static_cast<QFormLayout*>(ptr)->setRowWrapPolicy(static_cast<QFormLayout::RowWrapPolicy>(policy));
}

void QFormLayout_SetVerticalSpacing(void* ptr, int spacing){
	static_cast<QFormLayout*>(ptr)->setVerticalSpacing(spacing);
}

int QFormLayout_VerticalSpacing(void* ptr){
	return static_cast<QFormLayout*>(ptr)->verticalSpacing();
}

void* QFormLayout_NewQFormLayout(void* parent){
	return new QFormLayout(static_cast<QWidget*>(parent));
}

void QFormLayout_AddItem(void* ptr, void* item){
	static_cast<QFormLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

void QFormLayout_AddRow6(void* ptr, void* layout){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QLayout*>(layout));
}

void QFormLayout_AddRow2(void* ptr, void* label, void* field){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QWidget*>(label), static_cast<QLayout*>(field));
}

void QFormLayout_AddRow(void* ptr, void* label, void* field){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QWidget*>(label), static_cast<QWidget*>(field));
}

void QFormLayout_AddRow5(void* ptr, void* widget){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QWidget*>(widget));
}

void QFormLayout_AddRow4(void* ptr, char* labelText, void* field){
	static_cast<QFormLayout*>(ptr)->addRow(QString(labelText), static_cast<QLayout*>(field));
}

void QFormLayout_AddRow3(void* ptr, char* labelText, void* field){
	static_cast<QFormLayout*>(ptr)->addRow(QString(labelText), static_cast<QWidget*>(field));
}

int QFormLayout_Count(void* ptr){
	return static_cast<QFormLayout*>(ptr)->count();
}

int QFormLayout_ExpandingDirections(void* ptr){
	return static_cast<QFormLayout*>(ptr)->expandingDirections();
}

int QFormLayout_HasHeightForWidth(void* ptr){
	return static_cast<QFormLayout*>(ptr)->hasHeightForWidth();
}

int QFormLayout_HeightForWidth(void* ptr, int width){
	return static_cast<QFormLayout*>(ptr)->heightForWidth(width);
}

void QFormLayout_InsertRow6(void* ptr, int row, void* layout){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QLayout*>(layout));
}

void QFormLayout_InsertRow2(void* ptr, int row, void* label, void* field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QWidget*>(label), static_cast<QLayout*>(field));
}

void QFormLayout_InsertRow(void* ptr, int row, void* label, void* field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QWidget*>(label), static_cast<QWidget*>(field));
}

void QFormLayout_InsertRow5(void* ptr, int row, void* widget){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QWidget*>(widget));
}

void QFormLayout_InsertRow4(void* ptr, int row, char* labelText, void* field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, QString(labelText), static_cast<QLayout*>(field));
}

void QFormLayout_InsertRow3(void* ptr, int row, char* labelText, void* field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, QString(labelText), static_cast<QWidget*>(field));
}

void QFormLayout_Invalidate(void* ptr){
	static_cast<QFormLayout*>(ptr)->invalidate();
}

void* QFormLayout_ItemAt2(void* ptr, int index){
	return static_cast<QFormLayout*>(ptr)->itemAt(index);
}

void* QFormLayout_ItemAt(void* ptr, int row, int role){
	return static_cast<QFormLayout*>(ptr)->itemAt(row, static_cast<QFormLayout::ItemRole>(role));
}

void* QFormLayout_LabelForField2(void* ptr, void* field){
	return static_cast<QFormLayout*>(ptr)->labelForField(static_cast<QLayout*>(field));
}

void* QFormLayout_LabelForField(void* ptr, void* field){
	return static_cast<QFormLayout*>(ptr)->labelForField(static_cast<QWidget*>(field));
}

int QFormLayout_RowCount(void* ptr){
	return static_cast<QFormLayout*>(ptr)->rowCount();
}

void QFormLayout_SetGeometry(void* ptr, void* rect){
	static_cast<QFormLayout*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void QFormLayout_SetItem(void* ptr, int row, int role, void* item){
	static_cast<QFormLayout*>(ptr)->setItem(row, static_cast<QFormLayout::ItemRole>(role), static_cast<QLayoutItem*>(item));
}

void QFormLayout_SetLayout(void* ptr, int row, int role, void* layout){
	static_cast<QFormLayout*>(ptr)->setLayout(row, static_cast<QFormLayout::ItemRole>(role), static_cast<QLayout*>(layout));
}

void QFormLayout_SetSpacing(void* ptr, int spacing){
	static_cast<QFormLayout*>(ptr)->setSpacing(spacing);
}

void QFormLayout_SetWidget(void* ptr, int row, int role, void* widget){
	static_cast<QFormLayout*>(ptr)->setWidget(row, static_cast<QFormLayout::ItemRole>(role), static_cast<QWidget*>(widget));
}

int QFormLayout_Spacing(void* ptr){
	return static_cast<QFormLayout*>(ptr)->spacing();
}

void* QFormLayout_TakeAt(void* ptr, int index){
	return static_cast<QFormLayout*>(ptr)->takeAt(index);
}

void QFormLayout_DestroyQFormLayout(void* ptr){
	static_cast<QFormLayout*>(ptr)->~QFormLayout();
}

