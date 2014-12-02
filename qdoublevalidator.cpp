#include "qdoublevalidator.h"
#include <QDoubleValidator>
#include "cgoexport.h"


//Public Types
int QDoubleValidator_StandardNotation() { return QDoubleValidator::StandardNotation; }
int QDoubleValidator_ScientificNotation() { return QDoubleValidator::ScientificNotation; }

//Public Functions
QtObjectPtr QDoubleValidator_New_QObject(QtObjectPtr parent)
{
	return new QDoubleValidator(((QObject*)(parent)));
}

void QDoubleValidator_Destroy(QtObjectPtr ptr)
{
	((QDoubleValidator*)(ptr))->~QDoubleValidator();
}

int QDoubleValidator_Decimals(QtObjectPtr ptr)
{
	return ((QDoubleValidator*)(ptr))->decimals();
}

int QDoubleValidator_Notation(QtObjectPtr ptr)
{
	return ((QDoubleValidator*)(ptr))->notation();
}

void QDoubleValidator_SetDecimals_Int(QtObjectPtr ptr, int in)
{
	((QDoubleValidator*)(ptr))->setDecimals(in);
}

void QDoubleValidator_SetNotation_Notation(QtObjectPtr ptr, int No)
{
	((QDoubleValidator*)(ptr))->setNotation(((QDoubleValidator::Notation)(No)));
}

