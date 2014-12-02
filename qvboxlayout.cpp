#include "qvboxlayout.h"
#include <QVBoxLayout>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QVBoxLayout_New()
{
	return new QVBoxLayout();
}

QtObjectPtr QVBoxLayout_New_QWidget(QtObjectPtr parent)
{
	return new QVBoxLayout(((QWidget*)(parent)));
}

void QVBoxLayout_Destroy(QtObjectPtr ptr)
{
	((QVBoxLayout*)(ptr))->~QVBoxLayout();
}

