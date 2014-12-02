#include "qslider.h"
#include <QSlider>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QSlider_New_QWidget(QtObjectPtr parent)
{
	return new QSlider(((QWidget*)(parent)));
}

QtObjectPtr QSlider_New_Orientation_QWidget(int orientation, QtObjectPtr parent)
{
	return new QSlider(((Qt::Orientation)(orientation)), ((QWidget*)(parent)));
}

void QSlider_Destroy(QtObjectPtr ptr)
{
	((QSlider*)(ptr))->~QSlider();
}

void QSlider_SetTickInterval_Int(QtObjectPtr ptr, int ti)
{
	((QSlider*)(ptr))->setTickInterval(ti);
}

int QSlider_TickInterval(QtObjectPtr ptr)
{
	return ((QSlider*)(ptr))->tickInterval();
}

