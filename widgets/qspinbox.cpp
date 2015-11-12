#include "qspinbox.h"
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSpinBox>
#include "_cgo_export.h"

class MyQSpinBox: public QSpinBox {
public:
void Signal_ValueChanged(int i){callbackQSpinBoxValueChanged(this->objectName().toUtf8().data(), i);};
};

char* QSpinBox_CleanText(void* ptr){
	return static_cast<QSpinBox*>(ptr)->cleanText().toUtf8().data();
}

int QSpinBox_DisplayIntegerBase(void* ptr){
	return static_cast<QSpinBox*>(ptr)->displayIntegerBase();
}

int QSpinBox_Maximum(void* ptr){
	return static_cast<QSpinBox*>(ptr)->maximum();
}

int QSpinBox_Minimum(void* ptr){
	return static_cast<QSpinBox*>(ptr)->minimum();
}

char* QSpinBox_Prefix(void* ptr){
	return static_cast<QSpinBox*>(ptr)->prefix().toUtf8().data();
}

void QSpinBox_SetDisplayIntegerBase(void* ptr, int base){
	static_cast<QSpinBox*>(ptr)->setDisplayIntegerBase(base);
}

void QSpinBox_SetMaximum(void* ptr, int max){
	static_cast<QSpinBox*>(ptr)->setMaximum(max);
}

void QSpinBox_SetMinimum(void* ptr, int min){
	static_cast<QSpinBox*>(ptr)->setMinimum(min);
}

void QSpinBox_SetPrefix(void* ptr, char* prefix){
	static_cast<QSpinBox*>(ptr)->setPrefix(QString(prefix));
}

void QSpinBox_SetSingleStep(void* ptr, int val){
	static_cast<QSpinBox*>(ptr)->setSingleStep(val);
}

void QSpinBox_SetSuffix(void* ptr, char* suffix){
	static_cast<QSpinBox*>(ptr)->setSuffix(QString(suffix));
}

void QSpinBox_SetValue(void* ptr, int val){
	QMetaObject::invokeMethod(static_cast<QSpinBox*>(ptr), "setValue", Q_ARG(int, val));
}

int QSpinBox_SingleStep(void* ptr){
	return static_cast<QSpinBox*>(ptr)->singleStep();
}

char* QSpinBox_Suffix(void* ptr){
	return static_cast<QSpinBox*>(ptr)->suffix().toUtf8().data();
}

int QSpinBox_Value(void* ptr){
	return static_cast<QSpinBox*>(ptr)->value();
}

void* QSpinBox_NewQSpinBox(void* parent){
	return new QSpinBox(static_cast<QWidget*>(parent));
}

void QSpinBox_SetRange(void* ptr, int minimum, int maximum){
	static_cast<QSpinBox*>(ptr)->setRange(minimum, maximum);
}

void QSpinBox_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QSpinBox*>(ptr), static_cast<void (QSpinBox::*)(int)>(&QSpinBox::valueChanged), static_cast<MyQSpinBox*>(ptr), static_cast<void (MyQSpinBox::*)(int)>(&MyQSpinBox::Signal_ValueChanged));;
}

void QSpinBox_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QSpinBox*>(ptr), static_cast<void (QSpinBox::*)(int)>(&QSpinBox::valueChanged), static_cast<MyQSpinBox*>(ptr), static_cast<void (MyQSpinBox::*)(int)>(&MyQSpinBox::Signal_ValueChanged));;
}

void QSpinBox_DestroyQSpinBox(void* ptr){
	static_cast<QSpinBox*>(ptr)->~QSpinBox();
}

