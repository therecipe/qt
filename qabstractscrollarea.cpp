#include "qabstractscrollarea.h"
#include <QAbstractScrollArea>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QAbstractScrollArea_New_QWidget(QtObjectPtr parent)
{
	return new QAbstractScrollArea(((QWidget*)(parent)));
}

void QAbstractScrollArea_Destroy(QtObjectPtr ptr)
{
	((QAbstractScrollArea*)(ptr))->~QAbstractScrollArea();
}

void QAbstractScrollArea_AddScrollBarWidget_QWidget_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr widget, int alignment)
{
	((QAbstractScrollArea*)(ptr))->addScrollBarWidget(((QWidget*)(widget)), ((Qt::AlignmentFlag)(alignment)));
}

QtObjectPtr QAbstractScrollArea_CornerWidget(QtObjectPtr ptr)
{
	return ((QAbstractScrollArea*)(ptr))->cornerWidget();
}

int QAbstractScrollArea_HorizontalScrollBarPolicy(QtObjectPtr ptr)
{
	return ((QAbstractScrollArea*)(ptr))->horizontalScrollBarPolicy();
}

void QAbstractScrollArea_SetCornerWidget_QWidget(QtObjectPtr ptr, QtObjectPtr widget)
{
	((QAbstractScrollArea*)(ptr))->setCornerWidget(((QWidget*)(widget)));
}

void QAbstractScrollArea_SetHorizontalScrollBarPolicy_ScrollBarPolicy(QtObjectPtr ptr, int ScrollBarPolicy)
{
	((QAbstractScrollArea*)(ptr))->setHorizontalScrollBarPolicy(((Qt::ScrollBarPolicy)(ScrollBarPolicy)));
}

void QAbstractScrollArea_SetVerticalScrollBarPolicy_ScrollBarPolicy(QtObjectPtr ptr, int ScrollBarPolicy)
{
	((QAbstractScrollArea*)(ptr))->setVerticalScrollBarPolicy(((Qt::ScrollBarPolicy)(ScrollBarPolicy)));
}

void QAbstractScrollArea_SetViewport_QWidget(QtObjectPtr ptr, QtObjectPtr widget)
{
	((QAbstractScrollArea*)(ptr))->setViewport(((QWidget*)(widget)));
}

void QAbstractScrollArea_SetupViewport_QWidget(QtObjectPtr ptr, QtObjectPtr viewport)
{
	((QAbstractScrollArea*)(ptr))->setupViewport(((QWidget*)(viewport)));
}

int QAbstractScrollArea_VerticalScrollBarPolicy(QtObjectPtr ptr)
{
	return ((QAbstractScrollArea*)(ptr))->verticalScrollBarPolicy();
}

QtObjectPtr QAbstractScrollArea_Viewport(QtObjectPtr ptr)
{
	return ((QAbstractScrollArea*)(ptr))->viewport();
}

