#include "qmodelindex.h"
#include <QModelIndex>
#include "cgoexport.h"



//Public Functions
int QModelIndex_Column(QtObjectPtr ptr)
{
	return ((QModelIndex*)(ptr))->column();
}

char* QModelIndex_Data_Int(QtObjectPtr ptr, int role)
{
	return ((QModelIndex*)(ptr))->data(role).toString().toUtf8().data();
}

int QModelIndex_Flags(QtObjectPtr ptr)
{
	return ((QModelIndex*)(ptr))->flags();
}

int QModelIndex_IsValid(QtObjectPtr ptr)
{
	return ((QModelIndex*)(ptr))->isValid();
}

int QModelIndex_Row(QtObjectPtr ptr)
{
	return ((QModelIndex*)(ptr))->row();
}

