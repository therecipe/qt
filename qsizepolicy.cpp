#include "qsizepolicy.h"
#include <QSizePolicy>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QSizePolicy_New()
{
	return new QSizePolicy();
}

int QSizePolicy_ExpandingDirections(QtObjectPtr ptr)
{
	return ((QSizePolicy*)(ptr))->expandingDirections();
}

int QSizePolicy_HasHeightForWidth(QtObjectPtr ptr)
{
	return ((QSizePolicy*)(ptr))->hasHeightForWidth();
}

int QSizePolicy_HasWidthForHeight(QtObjectPtr ptr)
{
	return ((QSizePolicy*)(ptr))->hasWidthForHeight();
}

int QSizePolicy_HorizontalStretch(QtObjectPtr ptr)
{
	return ((QSizePolicy*)(ptr))->horizontalStretch();
}

int QSizePolicy_RetainSizeWhenHidden(QtObjectPtr ptr)
{
	return ((QSizePolicy*)(ptr))->retainSizeWhenHidden();
}

void QSizePolicy_SetHeightForWidth_Bool(QtObjectPtr ptr, int dependent)
{
	((QSizePolicy*)(ptr))->setHeightForWidth(dependent != 0);
}

void QSizePolicy_SetHorizontalStretch_Int(QtObjectPtr ptr, int stretchFactor)
{
	((QSizePolicy*)(ptr))->setHorizontalStretch(stretchFactor);
}

void QSizePolicy_SetRetainSizeWhenHidden_Bool(QtObjectPtr ptr, int retainSize)
{
	((QSizePolicy*)(ptr))->setRetainSizeWhenHidden(retainSize != 0);
}

void QSizePolicy_SetVerticalStretch_Int(QtObjectPtr ptr, int stretchFactor)
{
	((QSizePolicy*)(ptr))->setVerticalStretch(stretchFactor);
}

void QSizePolicy_SetWidthForHeight_Bool(QtObjectPtr ptr, int dependent)
{
	((QSizePolicy*)(ptr))->setWidthForHeight(dependent != 0);
}

void QSizePolicy_Transpose(QtObjectPtr ptr)
{
	((QSizePolicy*)(ptr))->transpose();
}

int QSizePolicy_VerticalStretch(QtObjectPtr ptr)
{
	return ((QSizePolicy*)(ptr))->verticalStretch();
}

