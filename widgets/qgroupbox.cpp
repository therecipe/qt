#include "qgroupbox.h"
#include <QWidget>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QGroupBox>
#include "_cgo_export.h"

class MyQGroupBox: public QGroupBox {
public:
void Signal_Clicked(bool checked){callbackQGroupBoxClicked(this->objectName().toUtf8().data(), checked);};
void Signal_Toggled(bool on){callbackQGroupBoxToggled(this->objectName().toUtf8().data(), on);};
};

int QGroupBox_Alignment(void* ptr){
	return static_cast<QGroupBox*>(ptr)->alignment();
}

int QGroupBox_IsCheckable(void* ptr){
	return static_cast<QGroupBox*>(ptr)->isCheckable();
}

int QGroupBox_IsChecked(void* ptr){
	return static_cast<QGroupBox*>(ptr)->isChecked();
}

int QGroupBox_IsFlat(void* ptr){
	return static_cast<QGroupBox*>(ptr)->isFlat();
}

void QGroupBox_SetAlignment(void* ptr, int alignment){
	static_cast<QGroupBox*>(ptr)->setAlignment(alignment);
}

void QGroupBox_SetCheckable(void* ptr, int checkable){
	static_cast<QGroupBox*>(ptr)->setCheckable(checkable != 0);
}

void QGroupBox_SetChecked(void* ptr, int checked){
	QMetaObject::invokeMethod(static_cast<QGroupBox*>(ptr), "setChecked", Q_ARG(bool, checked != 0));
}

void QGroupBox_SetFlat(void* ptr, int flat){
	static_cast<QGroupBox*>(ptr)->setFlat(flat != 0);
}

void QGroupBox_SetTitle(void* ptr, char* title){
	static_cast<QGroupBox*>(ptr)->setTitle(QString(title));
}

char* QGroupBox_Title(void* ptr){
	return static_cast<QGroupBox*>(ptr)->title().toUtf8().data();
}

void* QGroupBox_NewQGroupBox(void* parent){
	return new QGroupBox(static_cast<QWidget*>(parent));
}

void* QGroupBox_NewQGroupBox2(char* title, void* parent){
	return new QGroupBox(QString(title), static_cast<QWidget*>(parent));
}

void QGroupBox_ConnectClicked(void* ptr){
	QObject::connect(static_cast<QGroupBox*>(ptr), static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::clicked), static_cast<MyQGroupBox*>(ptr), static_cast<void (MyQGroupBox::*)(bool)>(&MyQGroupBox::Signal_Clicked));;
}

void QGroupBox_DisconnectClicked(void* ptr){
	QObject::disconnect(static_cast<QGroupBox*>(ptr), static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::clicked), static_cast<MyQGroupBox*>(ptr), static_cast<void (MyQGroupBox::*)(bool)>(&MyQGroupBox::Signal_Clicked));;
}

void QGroupBox_ConnectToggled(void* ptr){
	QObject::connect(static_cast<QGroupBox*>(ptr), static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::toggled), static_cast<MyQGroupBox*>(ptr), static_cast<void (MyQGroupBox::*)(bool)>(&MyQGroupBox::Signal_Toggled));;
}

void QGroupBox_DisconnectToggled(void* ptr){
	QObject::disconnect(static_cast<QGroupBox*>(ptr), static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::toggled), static_cast<MyQGroupBox*>(ptr), static_cast<void (MyQGroupBox::*)(bool)>(&MyQGroupBox::Signal_Toggled));;
}

void QGroupBox_DestroyQGroupBox(void* ptr){
	static_cast<QGroupBox*>(ptr)->~QGroupBox();
}

