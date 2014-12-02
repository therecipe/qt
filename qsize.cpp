#include "qsize.h"
#include <QSize>
#include "cgoexport.h"



//Public Functions
int QSize_Height(QtObjectPtr ptr)
{
	return ((QSize*)(ptr))->height();
}

int QSize_IsEmpty(QtObjectPtr ptr)
{
	return ((QSize*)(ptr))->isEmpty();
}

int QSize_IsNull(QtObjectPtr ptr)
{
	return ((QSize*)(ptr))->isNull();
}

int QSize_IsValid(QtObjectPtr ptr)
{
	return ((QSize*)(ptr))->isValid();
}

int QSize_Rheight(QtObjectPtr ptr)
{
	return ((QSize*)(ptr))->rheight();
}

int QSize_Rwidth(QtObjectPtr ptr)
{
	return ((QSize*)(ptr))->rwidth();
}

void QSize_Scale_Int_Int_AspectRatioMode(QtObjectPtr ptr, int width, int height, int mode)
{
	((QSize*)(ptr))->scale(width, height, ((Qt::AspectRatioMode)(mode)));
}

void QSize_SetHeight_Int(QtObjectPtr ptr, int height)
{
	((QSize*)(ptr))->setHeight(height);
}

void QSize_SetWidth_Int(QtObjectPtr ptr, int width)
{
	((QSize*)(ptr))->setWidth(width);
}

void QSize_Transpose(QtObjectPtr ptr)
{
	((QSize*)(ptr))->transpose();
}

int QSize_Width(QtObjectPtr ptr)
{
	return ((QSize*)(ptr))->width();
}

