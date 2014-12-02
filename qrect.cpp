#include "qrect.h"
#include <QRect>
#include "cgoexport.h"



//Public Functions
void QRect_Adjust_Int_Int_Int_Int(QtObjectPtr ptr, int dx1, int dy1, int dx2, int dy2)
{
	((QRect*)(ptr))->adjust(dx1, dy1, dx2, dy2);
}

int QRect_Bottom(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->bottom();
}

int QRect_Contains_Int_Int_Bool(QtObjectPtr ptr, int x, int y, int proper)
{
	return ((QRect*)(ptr))->contains(x, y, proper != 0);
}

int QRect_Contains_Int_Int(QtObjectPtr ptr, int x, int y)
{
	return ((QRect*)(ptr))->contains(x, y);
}

int QRect_Height(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->height();
}

int QRect_IsEmpty(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->isEmpty();
}

int QRect_IsNull(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->isNull();
}

int QRect_IsValid(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->isValid();
}

int QRect_Left(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->left();
}

void QRect_MoveBottom_Int(QtObjectPtr ptr, int y)
{
	((QRect*)(ptr))->moveBottom(y);
}

void QRect_MoveLeft_Int(QtObjectPtr ptr, int x)
{
	((QRect*)(ptr))->moveLeft(x);
}

void QRect_MoveRight_Int(QtObjectPtr ptr, int x)
{
	((QRect*)(ptr))->moveRight(x);
}

void QRect_MoveTo_Int_Int(QtObjectPtr ptr, int x, int y)
{
	((QRect*)(ptr))->moveTo(x, y);
}

void QRect_MoveTop_Int(QtObjectPtr ptr, int y)
{
	((QRect*)(ptr))->moveTop(y);
}

int QRect_Right(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->right();
}

void QRect_SetBottom_Int(QtObjectPtr ptr, int y)
{
	((QRect*)(ptr))->setBottom(y);
}

void QRect_SetCoords_Int_Int_Int_Int(QtObjectPtr ptr, int x1, int y1, int x2, int y2)
{
	((QRect*)(ptr))->setCoords(x1, y1, x2, y2);
}

void QRect_SetHeight_Int(QtObjectPtr ptr, int height)
{
	((QRect*)(ptr))->setHeight(height);
}

void QRect_SetLeft_Int(QtObjectPtr ptr, int x)
{
	((QRect*)(ptr))->setLeft(x);
}

void QRect_SetRect_Int_Int_Int_Int(QtObjectPtr ptr, int x, int y, int width, int height)
{
	((QRect*)(ptr))->setRect(x, y, width, height);
}

void QRect_SetRight_Int(QtObjectPtr ptr, int x)
{
	((QRect*)(ptr))->setRight(x);
}

void QRect_SetTop_Int(QtObjectPtr ptr, int y)
{
	((QRect*)(ptr))->setTop(y);
}

void QRect_SetWidth_Int(QtObjectPtr ptr, int width)
{
	((QRect*)(ptr))->setWidth(width);
}

void QRect_SetX_Int(QtObjectPtr ptr, int x)
{
	((QRect*)(ptr))->setX(x);
}

void QRect_SetY_Int(QtObjectPtr ptr, int y)
{
	((QRect*)(ptr))->setY(y);
}

int QRect_Top(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->top();
}

void QRect_Translate_Int_Int(QtObjectPtr ptr, int dx, int dy)
{
	((QRect*)(ptr))->translate(dx, dy);
}

int QRect_Width(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->width();
}

int QRect_X(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->x();
}

int QRect_Y(QtObjectPtr ptr)
{
	return ((QRect*)(ptr))->y();
}

