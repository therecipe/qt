#include "qhboxlayout.h"
#include <QHBoxLayout>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QHBoxLayout_New()
{
	return new QHBoxLayout();
}

QtObjectPtr QHBoxLayout_New_QWidget(QtObjectPtr parent)
{
	return new QHBoxLayout(((QWidget*)(parent)));
}

void QHBoxLayout_Destroy(QtObjectPtr ptr)
{
	((QHBoxLayout*)(ptr))->~QHBoxLayout();
}

