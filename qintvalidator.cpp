#include "qintvalidator.h"
#include <QIntValidator>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QIntValidator_New_QObject(QtObjectPtr parent)
{
	return new QIntValidator(((QObject*)(parent)));
}

QtObjectPtr QIntValidator_New_Int_Int_QObject(int minimum, int maximum, QtObjectPtr parent)
{
	return new QIntValidator(minimum, maximum, ((QObject*)(parent)));
}

void QIntValidator_Destroy(QtObjectPtr ptr)
{
	((QIntValidator*)(ptr))->~QIntValidator();
}

int QIntValidator_Bottom(QtObjectPtr ptr)
{
	return ((QIntValidator*)(ptr))->bottom();
}

void QIntValidator_SetBottom_Int(QtObjectPtr ptr, int in)
{
	((QIntValidator*)(ptr))->setBottom(in);
}

void QIntValidator_SetRange_Int_Int(QtObjectPtr ptr, int bottom, int top)
{
	((QIntValidator*)(ptr))->setRange(bottom, top);
}

void QIntValidator_SetTop_Int(QtObjectPtr ptr, int in)
{
	((QIntValidator*)(ptr))->setTop(in);
}

int QIntValidator_Top(QtObjectPtr ptr)
{
	return ((QIntValidator*)(ptr))->top();
}

