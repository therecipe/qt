#include "qcheckbox.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QCheckBox>
#include "_cgo_export.h"

class MyQCheckBox: public QCheckBox {
public:
void Signal_StateChanged(int state){callbackQCheckBoxStateChanged(this->objectName().toUtf8().data(), state);};
};

int QCheckBox_IsTristate(QtObjectPtr ptr){
	return static_cast<QCheckBox*>(ptr)->isTristate();
}

void QCheckBox_SetTristate(QtObjectPtr ptr, int y){
	static_cast<QCheckBox*>(ptr)->setTristate(y != 0);
}

QtObjectPtr QCheckBox_NewQCheckBox(QtObjectPtr parent){
	return new QCheckBox(static_cast<QWidget*>(parent));
}

QtObjectPtr QCheckBox_NewQCheckBox2(char* text, QtObjectPtr parent){
	return new QCheckBox(QString(text), static_cast<QWidget*>(parent));
}

int QCheckBox_CheckState(QtObjectPtr ptr){
	return static_cast<QCheckBox*>(ptr)->checkState();
}

void QCheckBox_SetCheckState(QtObjectPtr ptr, int state){
	static_cast<QCheckBox*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QCheckBox_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCheckBox*>(ptr), static_cast<void (QCheckBox::*)(int)>(&QCheckBox::stateChanged), static_cast<MyQCheckBox*>(ptr), static_cast<void (MyQCheckBox::*)(int)>(&MyQCheckBox::Signal_StateChanged));;
}

void QCheckBox_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCheckBox*>(ptr), static_cast<void (QCheckBox::*)(int)>(&QCheckBox::stateChanged), static_cast<MyQCheckBox*>(ptr), static_cast<void (MyQCheckBox::*)(int)>(&MyQCheckBox::Signal_StateChanged));;
}

void QCheckBox_DestroyQCheckBox(QtObjectPtr ptr){
	static_cast<QCheckBox*>(ptr)->~QCheckBox();
}

