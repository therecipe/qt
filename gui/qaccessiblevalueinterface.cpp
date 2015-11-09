#include "qaccessiblevalueinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QAccessibleValueInterface>
#include "_cgo_export.h"

class MyQAccessibleValueInterface: public QAccessibleValueInterface {
public:
};

void* QAccessibleValueInterface_CurrentValue(void* ptr){
	return new QVariant(static_cast<QAccessibleValueInterface*>(ptr)->currentValue());
}

void* QAccessibleValueInterface_MaximumValue(void* ptr){
	return new QVariant(static_cast<QAccessibleValueInterface*>(ptr)->maximumValue());
}

void* QAccessibleValueInterface_MinimumStepSize(void* ptr){
	return new QVariant(static_cast<QAccessibleValueInterface*>(ptr)->minimumStepSize());
}

void* QAccessibleValueInterface_MinimumValue(void* ptr){
	return new QVariant(static_cast<QAccessibleValueInterface*>(ptr)->minimumValue());
}

void QAccessibleValueInterface_SetCurrentValue(void* ptr, void* value){
	static_cast<QAccessibleValueInterface*>(ptr)->setCurrentValue(*static_cast<QVariant*>(value));
}

void QAccessibleValueInterface_DestroyQAccessibleValueInterface(void* ptr){
	static_cast<QAccessibleValueInterface*>(ptr)->~QAccessibleValueInterface();
}

