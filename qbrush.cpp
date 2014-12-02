#include "qbrush.h"
#include <QBrush>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QBrush_New()
{
	return new QBrush();
}

QtObjectPtr QBrush_New_BrushStyle(int style)
{
	return new QBrush(((Qt::BrushStyle)(style)));
}

QtObjectPtr QBrush_New_GlobalColor_BrushStyle(int color, int style)
{
	return new QBrush(((Qt::GlobalColor)(color)), ((Qt::BrushStyle)(style)));
}

void QBrush_Destroy(QtObjectPtr ptr)
{
	((QBrush*)(ptr))->~QBrush();
}

int QBrush_IsOpaque(QtObjectPtr ptr)
{
	return ((QBrush*)(ptr))->isOpaque();
}

void QBrush_SetColor_GlobalColor(QtObjectPtr ptr, int color)
{
	((QBrush*)(ptr))->setColor(((Qt::GlobalColor)(color)));
}

void QBrush_SetStyle_BrushStyle(QtObjectPtr ptr, int style)
{
	((QBrush*)(ptr))->setStyle(((Qt::BrushStyle)(style)));
}

int QBrush_Style(QtObjectPtr ptr)
{
	return ((QBrush*)(ptr))->style();
}

