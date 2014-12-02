#include "qsizegrip.h"
#include <QSizeGrip>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QSizeGrip_New_QWidget(QtObjectPtr parent)
{
	return new QSizeGrip(((QWidget*)(parent)));
}

void QSizeGrip_Destroy(QtObjectPtr ptr)
{
	((QSizeGrip*)(ptr))->~QSizeGrip();
}

