#include "qradiobutton.h"
#include <QRadioButton>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QRadioButton_New_QWidget(QtObjectPtr parent)
{
	return new QRadioButton(((QWidget*)(parent)));
}

QtObjectPtr QRadioButton_New_String_QWidget(char* text, QtObjectPtr parent)
{
	return new QRadioButton(QString(text), ((QWidget*)(parent)));
}

void QRadioButton_Destroy(QtObjectPtr ptr)
{
	((QRadioButton*)(ptr))->~QRadioButton();
}

