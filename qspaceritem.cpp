#include "qspaceritem.h"
#include <QSpacerItem>
#include "cgoexport.h"



//Public Functions
void QSpacerItem_Destroy(QtObjectPtr ptr)
{
	((QSpacerItem*)(ptr))->~QSpacerItem();
}

