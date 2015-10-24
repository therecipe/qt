#include "qspinbox.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QSpinBox>
#include "_cgo_export.h"

class MyQSpinBox: public QSpinBox {
public:
void Signal_ValueChanged(int i){callbackQSpinBoxValueChanged(this->objectName().toUtf8().data(), i);};
};

char* QSpinBox_CleanText(QtObjectPtr ptr){
	return static_cast<QSpinBox*>(ptr)->cleanText().toUtf8().data();
}

int QSpinBox_DisplayIntegerBase(QtObjectPtr ptr){
	return static_cast<QSpinBox*>(ptr)->displayIntegerBase();
}

int QSpinBox_Maximum(QtObjectPtr ptr){
	return static_cast<QSpinBox*>(ptr)->maximum();
}

int QSpinBox_Minimum(QtObjectPtr ptr){
	return static_cast<QSpinBox*>(ptr)->minimum();
}

char* QSpinBox_Prefix(QtObjectPtr ptr){
	return static_cast<QSpinBox*>(ptr)->prefix().toUtf8().data();
}

void QSpinBox_SetDisplayIntegerBase(QtObjectPtr ptr, int base){
	static_cast<QSpinBox*>(ptr)->setDisplayIntegerBase(base);
}

void QSpinBox_SetMaximum(QtObjectPtr ptr, int max){
	static_cast<QSpinBox*>(ptr)->setMaximum(max);
}

void QSpinBox_SetMinimum(QtObjectPtr ptr, int min){
	static_cast<QSpinBox*>(ptr)->setMinimum(min);
}

void QSpinBox_SetPrefix(QtObjectPtr ptr, char* prefix){
	static_cast<QSpinBox*>(ptr)->setPrefix(QString(prefix));
}

void QSpinBox_SetSingleStep(QtObjectPtr ptr, int val){
	static_cast<QSpinBox*>(ptr)->setSingleStep(val);
}

void QSpinBox_SetSuffix(QtObjectPtr ptr, char* suffix){
	static_cast<QSpinBox*>(ptr)->setSuffix(QString(suffix));
}

void QSpinBox_SetValue(QtObjectPtr ptr, int val){
	QMetaObject::invokeMethod(static_cast<QSpinBox*>(ptr), "setValue", Q_ARG(int, val));
}

int QSpinBox_SingleStep(QtObjectPtr ptr){
	return static_cast<QSpinBox*>(ptr)->singleStep();
}

char* QSpinBox_Suffix(QtObjectPtr ptr){
	return static_cast<QSpinBox*>(ptr)->suffix().toUtf8().data();
}

int QSpinBox_Value(QtObjectPtr ptr){
	return static_cast<QSpinBox*>(ptr)->value();
}

QtObjectPtr QSpinBox_NewQSpinBox(QtObjectPtr parent){
	return new QSpinBox(static_cast<QWidget*>(parent));
}

void QSpinBox_SetRange(QtObjectPtr ptr, int minimum, int maximum){
	static_cast<QSpinBox*>(ptr)->setRange(minimum, maximum);
}

void QSpinBox_ConnectValueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSpinBox*>(ptr), static_cast<void (QSpinBox::*)(int)>(&QSpinBox::valueChanged), static_cast<MyQSpinBox*>(ptr), static_cast<void (MyQSpinBox::*)(int)>(&MyQSpinBox::Signal_ValueChanged));;
}

void QSpinBox_DisconnectValueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSpinBox*>(ptr), static_cast<void (QSpinBox::*)(int)>(&QSpinBox::valueChanged), static_cast<MyQSpinBox*>(ptr), static_cast<void (MyQSpinBox::*)(int)>(&MyQSpinBox::Signal_ValueChanged));;
}

void QSpinBox_DestroyQSpinBox(QtObjectPtr ptr){
	static_cast<QSpinBox*>(ptr)->~QSpinBox();
}

