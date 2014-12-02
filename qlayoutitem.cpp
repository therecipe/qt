#include "qlayoutitem.h"
#include <QLayoutItem>
#include "cgoexport.h"



//Public Functions
void QLayoutItem_Destroy(QtObjectPtr ptr)
{
	((QLayoutItem*)(ptr))->~QLayoutItem();
}

int QLayoutItem_Alignment(QtObjectPtr ptr)
{
	return ((QLayoutItem*)(ptr))->alignment();
}

int QLayoutItem_ExpandingDirections(QtObjectPtr ptr)
{
	return ((QLayoutItem*)(ptr))->expandingDirections();
}

int QLayoutItem_HasHeightForWidth(QtObjectPtr ptr)
{
	return ((QLayoutItem*)(ptr))->hasHeightForWidth();
}

int QLayoutItem_HeightForWidth_Int(QtObjectPtr ptr, int w)
{
	return ((QLayoutItem*)(ptr))->heightForWidth(w);
}

void QLayoutItem_Invalidate(QtObjectPtr ptr)
{
	((QLayoutItem*)(ptr))->invalidate();
}

int QLayoutItem_IsEmpty(QtObjectPtr ptr)
{
	return ((QLayoutItem*)(ptr))->isEmpty();
}

QtObjectPtr QLayoutItem_Layout(QtObjectPtr ptr)
{
	return ((QLayoutItem*)(ptr))->layout();
}

int QLayoutItem_MinimumHeightForWidth_Int(QtObjectPtr ptr, int w)
{
	return ((QLayoutItem*)(ptr))->minimumHeightForWidth(w);
}

void QLayoutItem_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment)
{
	((QLayoutItem*)(ptr))->setAlignment(((Qt::AlignmentFlag)(alignment)));
}

QtObjectPtr QLayoutItem_Widget(QtObjectPtr ptr)
{
	return ((QLayoutItem*)(ptr))->widget();
}

