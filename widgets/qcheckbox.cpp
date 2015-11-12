#include "qcheckbox.h"
#include <QWidget>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCheckBox>
#include "_cgo_export.h"

class MyQCheckBox: public QCheckBox {
public:
void Signal_StateChanged(int state){callbackQCheckBoxStateChanged(this->objectName().toUtf8().data(), state);};
};

int QCheckBox_IsTristate(void* ptr){
	return static_cast<QCheckBox*>(ptr)->isTristate();
}

void QCheckBox_SetTristate(void* ptr, int y){
	static_cast<QCheckBox*>(ptr)->setTristate(y != 0);
}

void* QCheckBox_NewQCheckBox(void* parent){
	return new QCheckBox(static_cast<QWidget*>(parent));
}

void* QCheckBox_NewQCheckBox2(char* text, void* parent){
	return new QCheckBox(QString(text), static_cast<QWidget*>(parent));
}

int QCheckBox_CheckState(void* ptr){
	return static_cast<QCheckBox*>(ptr)->checkState();
}

void QCheckBox_SetCheckState(void* ptr, int state){
	static_cast<QCheckBox*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QCheckBox_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QCheckBox*>(ptr), static_cast<void (QCheckBox::*)(int)>(&QCheckBox::stateChanged), static_cast<MyQCheckBox*>(ptr), static_cast<void (MyQCheckBox::*)(int)>(&MyQCheckBox::Signal_StateChanged));;
}

void QCheckBox_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QCheckBox*>(ptr), static_cast<void (QCheckBox::*)(int)>(&QCheckBox::stateChanged), static_cast<MyQCheckBox*>(ptr), static_cast<void (MyQCheckBox::*)(int)>(&MyQCheckBox::Signal_StateChanged));;
}

void QCheckBox_DestroyQCheckBox(void* ptr){
	static_cast<QCheckBox*>(ptr)->~QCheckBox();
}

