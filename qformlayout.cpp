#include "qformlayout.h"
#include <QFormLayout>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QFormLayout_New_QWidget(QtObjectPtr parent)
{
	return new QFormLayout(((QWidget*)(parent)));
}

void QFormLayout_Destroy(QtObjectPtr ptr)
{
	((QFormLayout*)(ptr))->~QFormLayout();
}

void QFormLayout_AddRow_QWidget_QWidget(QtObjectPtr ptr, QtObjectPtr label, QtObjectPtr field)
{
	((QFormLayout*)(ptr))->addRow(((QWidget*)(label)), ((QWidget*)(field)));
}

void QFormLayout_AddRow_QWidget_QLayout(QtObjectPtr ptr, QtObjectPtr label, QtObjectPtr field)
{
	((QFormLayout*)(ptr))->addRow(((QWidget*)(label)), ((QLayout*)(field)));
}

void QFormLayout_AddRow_String_QWidget(QtObjectPtr ptr, char* labelText, QtObjectPtr field)
{
	((QFormLayout*)(ptr))->addRow(QString(labelText), ((QWidget*)(field)));
}

void QFormLayout_AddRow_String_QLayout(QtObjectPtr ptr, char* labelText, QtObjectPtr field)
{
	((QFormLayout*)(ptr))->addRow(QString(labelText), ((QLayout*)(field)));
}

void QFormLayout_AddRow_QWidget(QtObjectPtr ptr, QtObjectPtr widget)
{
	((QFormLayout*)(ptr))->addRow(((QWidget*)(widget)));
}

void QFormLayout_AddRow_QLayout(QtObjectPtr ptr, QtObjectPtr layout)
{
	((QFormLayout*)(ptr))->addRow(((QLayout*)(layout)));
}

int QFormLayout_FormAlignment(QtObjectPtr ptr)
{
	return ((QFormLayout*)(ptr))->formAlignment();
}

int QFormLayout_HorizontalSpacing(QtObjectPtr ptr)
{
	return ((QFormLayout*)(ptr))->horizontalSpacing();
}

void QFormLayout_InsertRow_Int_QWidget_QWidget(QtObjectPtr ptr, int row, QtObjectPtr label, QtObjectPtr field)
{
	((QFormLayout*)(ptr))->insertRow(row, ((QWidget*)(label)), ((QWidget*)(field)));
}

void QFormLayout_InsertRow_Int_QWidget_QLayout(QtObjectPtr ptr, int row, QtObjectPtr label, QtObjectPtr field)
{
	((QFormLayout*)(ptr))->insertRow(row, ((QWidget*)(label)), ((QLayout*)(field)));
}

void QFormLayout_InsertRow_Int_String_QWidget(QtObjectPtr ptr, int row, char* labelText, QtObjectPtr field)
{
	((QFormLayout*)(ptr))->insertRow(row, QString(labelText), ((QWidget*)(field)));
}

void QFormLayout_InsertRow_Int_String_QLayout(QtObjectPtr ptr, int row, char* labelText, QtObjectPtr field)
{
	((QFormLayout*)(ptr))->insertRow(row, QString(labelText), ((QLayout*)(field)));
}

void QFormLayout_InsertRow_Int_QWidget(QtObjectPtr ptr, int row, QtObjectPtr widget)
{
	((QFormLayout*)(ptr))->insertRow(row, ((QWidget*)(widget)));
}

void QFormLayout_InsertRow_Int_QLayout(QtObjectPtr ptr, int row, QtObjectPtr layout)
{
	((QFormLayout*)(ptr))->insertRow(row, ((QLayout*)(layout)));
}

int QFormLayout_LabelAlignment(QtObjectPtr ptr)
{
	return ((QFormLayout*)(ptr))->labelAlignment();
}

QtObjectPtr QFormLayout_LabelForField_QWidget(QtObjectPtr ptr, QtObjectPtr field)
{
	return ((QFormLayout*)(ptr))->labelForField(((QWidget*)(field)));
}

int QFormLayout_RowCount(QtObjectPtr ptr)
{
	return ((QFormLayout*)(ptr))->rowCount();
}

void QFormLayout_SetFormAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment)
{
	((QFormLayout*)(ptr))->setFormAlignment(((Qt::AlignmentFlag)(alignment)));
}

void QFormLayout_SetHorizontalSpacing_Int(QtObjectPtr ptr, int spacing)
{
	((QFormLayout*)(ptr))->setHorizontalSpacing(spacing);
}

void QFormLayout_SetLabelAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment)
{
	((QFormLayout*)(ptr))->setLabelAlignment(((Qt::AlignmentFlag)(alignment)));
}

void QFormLayout_SetVerticalSpacing_Int(QtObjectPtr ptr, int spacing)
{
	((QFormLayout*)(ptr))->setVerticalSpacing(spacing);
}

int QFormLayout_VerticalSpacing(QtObjectPtr ptr)
{
	return ((QFormLayout*)(ptr))->verticalSpacing();
}

