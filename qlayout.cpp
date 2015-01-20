#include "qlayout.h"
#include <QLayout>
#include "cgoexport.h"


//Public Types
int QLayout_SetDefaultConstraint() { return QLayout::SetDefaultConstraint; }
int QLayout_SetFixedSize() { return QLayout::SetFixedSize; }
int QLayout_SetMinimumSize() { return QLayout::SetMinimumSize; }
int QLayout_SetMaximumSize() { return QLayout::SetMaximumSize; }
int QLayout_SetMinAndMaxSize() { return QLayout::SetMinAndMaxSize; }
int QLayout_SetNoConstraint() { return QLayout::SetNoConstraint; }

//Public Functions
int QLayout_Activate(QtObjectPtr ptr)
{
	return ((QLayout*)(ptr))->activate();
}

void QLayout_AddWidget_QWidget(QtObjectPtr ptr, QtObjectPtr w)
{
	((QLayout*)(ptr))->addWidget(((QWidget*)(w)));
}

int QLayout_Count(QtObjectPtr ptr)
{
	return ((QLayout*)(ptr))->count();
}

int QLayout_IndexOf_QWidget(QtObjectPtr ptr, QtObjectPtr widget)
{
	return ((QLayout*)(ptr))->indexOf(((QWidget*)(widget)));
}

int QLayout_IsEnabled(QtObjectPtr ptr)
{
	return ((QLayout*)(ptr))->isEnabled();
}

QtObjectPtr QLayout_MenuBar(QtObjectPtr ptr)
{
	return ((QLayout*)(ptr))->menuBar();
}

QtObjectPtr QLayout_ParentWidget(QtObjectPtr ptr)
{
	return ((QLayout*)(ptr))->parentWidget();
}

void QLayout_RemoveItem_QLayoutItem(QtObjectPtr ptr, QtObjectPtr item)
{
	((QLayout*)(ptr))->removeItem(((QLayoutItem*)(item)));
}

void QLayout_RemoveWidget_QWidget(QtObjectPtr ptr, QtObjectPtr widget)
{
	((QLayout*)(ptr))->removeWidget(((QWidget*)(widget)));
}

int QLayout_SetAlignment_QWidget_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr w, int alignment)
{
	return ((QLayout*)(ptr))->setAlignment(((QWidget*)(w)), ((Qt::AlignmentFlag)(alignment)));
}

void QLayout_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment)
{
	((QLayout*)(ptr))->setAlignment(((Qt::AlignmentFlag)(alignment)));
}

int QLayout_SetAlignment_QLayout_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr l, int alignment)
{
	return ((QLayout*)(ptr))->setAlignment(((QLayout*)(l)), ((Qt::AlignmentFlag)(alignment)));
}

void QLayout_SetContentsMargins_Int_Int_Int_Int(QtObjectPtr ptr, int left, int top, int right, int bottom)
{
	((QLayout*)(ptr))->setContentsMargins(left, top, right, bottom);
}

void QLayout_SetEnabled_Bool(QtObjectPtr ptr, int enable)
{
	((QLayout*)(ptr))->setEnabled(enable != 0);
}

void QLayout_SetMenuBar_QWidget(QtObjectPtr ptr, QtObjectPtr widget)
{
	((QLayout*)(ptr))->setMenuBar(((QWidget*)(widget)));
}

void QLayout_SetSizeConstraint_SizeConstraint(QtObjectPtr ptr, int Si)
{
	((QLayout*)(ptr))->setSizeConstraint(((QLayout::SizeConstraint)(Si)));
}

void QLayout_SetSpacing_Int(QtObjectPtr ptr, int spacing)
{
	((QLayout*)(ptr))->setSpacing(spacing);
}

int QLayout_SizeConstraint(QtObjectPtr ptr)
{
	return ((QLayout*)(ptr))->sizeConstraint();
}

int QLayout_Spacing(QtObjectPtr ptr)
{
	return ((QLayout*)(ptr))->spacing();
}

void QLayout_Update(QtObjectPtr ptr)
{
	((QLayout*)(ptr))->update();
}

