#include "qwidgetitem.h"
#include <QWidgetItem>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QWidgetItem_New_QWidget(QtObjectPtr widget)
{
	return new QWidgetItem(((QWidget*)(widget)));
}

void QWidgetItem_Destroy(QtObjectPtr ptr)
{
	((QWidgetItem*)(ptr))->~QWidgetItem();
}

