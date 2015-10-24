#include "qaccessiblevalueinterface.h"
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QAccessibleValueInterface>
#include "_cgo_export.h"

class MyQAccessibleValueInterface: public QAccessibleValueInterface {
public:
};

char* QAccessibleValueInterface_CurrentValue(QtObjectPtr ptr){
	return static_cast<QAccessibleValueInterface*>(ptr)->currentValue().toString().toUtf8().data();
}

char* QAccessibleValueInterface_MaximumValue(QtObjectPtr ptr){
	return static_cast<QAccessibleValueInterface*>(ptr)->maximumValue().toString().toUtf8().data();
}

char* QAccessibleValueInterface_MinimumStepSize(QtObjectPtr ptr){
	return static_cast<QAccessibleValueInterface*>(ptr)->minimumStepSize().toString().toUtf8().data();
}

char* QAccessibleValueInterface_MinimumValue(QtObjectPtr ptr){
	return static_cast<QAccessibleValueInterface*>(ptr)->minimumValue().toString().toUtf8().data();
}

void QAccessibleValueInterface_SetCurrentValue(QtObjectPtr ptr, char* value){
	static_cast<QAccessibleValueInterface*>(ptr)->setCurrentValue(QVariant(value));
}

void QAccessibleValueInterface_DestroyQAccessibleValueInterface(QtObjectPtr ptr){
	static_cast<QAccessibleValueInterface*>(ptr)->~QAccessibleValueInterface();
}

