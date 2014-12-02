#include "qframe.h"
#include <QFrame>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QFrame_New_QWidget_WindowType(QtObjectPtr parent, int f)
{
	return new QFrame(((QWidget*)(parent)), ((Qt::WindowType)(f)));
}

void QFrame_Destroy(QtObjectPtr ptr)
{
	((QFrame*)(ptr))->~QFrame();
}

int QFrame_FrameStyle(QtObjectPtr ptr)
{
	return ((QFrame*)(ptr))->frameStyle();
}

int QFrame_FrameWidth(QtObjectPtr ptr)
{
	return ((QFrame*)(ptr))->frameWidth();
}

int QFrame_LineWidth(QtObjectPtr ptr)
{
	return ((QFrame*)(ptr))->lineWidth();
}

int QFrame_MidLineWidth(QtObjectPtr ptr)
{
	return ((QFrame*)(ptr))->midLineWidth();
}

void QFrame_SetFrameStyle_Int(QtObjectPtr ptr, int style)
{
	((QFrame*)(ptr))->setFrameStyle(style);
}

void QFrame_SetLineWidth_Int(QtObjectPtr ptr, int width)
{
	((QFrame*)(ptr))->setLineWidth(width);
}

void QFrame_SetMidLineWidth_Int(QtObjectPtr ptr, int width)
{
	((QFrame*)(ptr))->setMidLineWidth(width);
}

