#include "qboxlayout.h"
#include <QBoxLayout>
#include "cgoexport.h"



//Public Functions
void QBoxLayout_Destroy(QtObjectPtr ptr)
{
	((QBoxLayout*)(ptr))->~QBoxLayout();
}

void QBoxLayout_AddLayout_QLayout_Int(QtObjectPtr ptr, QtObjectPtr layout, int stretch)
{
	((QBoxLayout*)(ptr))->addLayout(((QLayout*)(layout)), stretch);
}

void QBoxLayout_AddSpacing_Int(QtObjectPtr ptr, int size)
{
	((QBoxLayout*)(ptr))->addSpacing(size);
}

void QBoxLayout_AddStretch_Int(QtObjectPtr ptr, int stretch)
{
	((QBoxLayout*)(ptr))->addStretch(stretch);
}

void QBoxLayout_AddStrut_Int(QtObjectPtr ptr, int size)
{
	((QBoxLayout*)(ptr))->addStrut(size);
}

void QBoxLayout_InsertLayout_Int_QLayout_Int(QtObjectPtr ptr, int index, QtObjectPtr layout, int stretch)
{
	((QBoxLayout*)(ptr))->insertLayout(index, ((QLayout*)(layout)), stretch);
}

void QBoxLayout_InsertSpacing_Int_Int(QtObjectPtr ptr, int index, int size)
{
	((QBoxLayout*)(ptr))->insertSpacing(index, size);
}

void QBoxLayout_InsertStretch_Int_Int(QtObjectPtr ptr, int index, int stretch)
{
	((QBoxLayout*)(ptr))->insertStretch(index, stretch);
}

void QBoxLayout_InsertWidget_Int_QWidget_Int_AlignmentFlag(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch, int alignment)
{
	((QBoxLayout*)(ptr))->insertWidget(index, ((QWidget*)(widget)), stretch, ((Qt::AlignmentFlag)(alignment)));
}

void QBoxLayout_SetStretch_Int_Int(QtObjectPtr ptr, int index, int stretch)
{
	((QBoxLayout*)(ptr))->setStretch(index, stretch);
}

int QBoxLayout_SetStretchFactor_QWidget_Int(QtObjectPtr ptr, QtObjectPtr widget, int stretch)
{
	return ((QBoxLayout*)(ptr))->setStretchFactor(((QWidget*)(widget)), stretch);
}

int QBoxLayout_Stretch_Int(QtObjectPtr ptr, int index)
{
	return ((QBoxLayout*)(ptr))->stretch(index);
}

