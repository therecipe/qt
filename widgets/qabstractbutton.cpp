#include "qabstractbutton.h"
#include <QUrl>
#include <QSize>
#include <QIcon>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QKeySequence>
#include <QMetaObject>
#include <QObject>
#include <QAbstractButton>
#include "_cgo_export.h"

class MyQAbstractButton: public QAbstractButton {
public:
void Signal_Clicked(bool checked){callbackQAbstractButtonClicked(this->objectName().toUtf8().data(), checked);};
void Signal_Pressed(){callbackQAbstractButtonPressed(this->objectName().toUtf8().data());};
void Signal_Released(){callbackQAbstractButtonReleased(this->objectName().toUtf8().data());};
void Signal_Toggled(bool checked){callbackQAbstractButtonToggled(this->objectName().toUtf8().data(), checked);};
};

int QAbstractButton_AutoExclusive(QtObjectPtr ptr){
	return static_cast<QAbstractButton*>(ptr)->autoExclusive();
}

int QAbstractButton_AutoRepeat(QtObjectPtr ptr){
	return static_cast<QAbstractButton*>(ptr)->autoRepeat();
}

int QAbstractButton_AutoRepeatDelay(QtObjectPtr ptr){
	return static_cast<QAbstractButton*>(ptr)->autoRepeatDelay();
}

int QAbstractButton_AutoRepeatInterval(QtObjectPtr ptr){
	return static_cast<QAbstractButton*>(ptr)->autoRepeatInterval();
}

int QAbstractButton_IsCheckable(QtObjectPtr ptr){
	return static_cast<QAbstractButton*>(ptr)->isCheckable();
}

int QAbstractButton_IsChecked(QtObjectPtr ptr){
	return static_cast<QAbstractButton*>(ptr)->isChecked();
}

int QAbstractButton_IsDown(QtObjectPtr ptr){
	return static_cast<QAbstractButton*>(ptr)->isDown();
}

void QAbstractButton_SetAutoExclusive(QtObjectPtr ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoExclusive(v != 0);
}

void QAbstractButton_SetAutoRepeat(QtObjectPtr ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoRepeat(v != 0);
}

void QAbstractButton_SetAutoRepeatDelay(QtObjectPtr ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoRepeatDelay(v);
}

void QAbstractButton_SetAutoRepeatInterval(QtObjectPtr ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoRepeatInterval(v);
}

void QAbstractButton_SetCheckable(QtObjectPtr ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setCheckable(v != 0);
}

void QAbstractButton_SetChecked(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "setChecked", Q_ARG(bool, v != 0));
}

void QAbstractButton_SetDown(QtObjectPtr ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setDown(v != 0);
}

void QAbstractButton_SetIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QAbstractButton*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QAbstractButton_SetIconSize(QtObjectPtr ptr, QtObjectPtr size){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "setIconSize", Q_ARG(QSize, *static_cast<QSize*>(size)));
}

void QAbstractButton_SetShortcut(QtObjectPtr ptr, QtObjectPtr key){
	static_cast<QAbstractButton*>(ptr)->setShortcut(*static_cast<QKeySequence*>(key));
}

void QAbstractButton_SetText(QtObjectPtr ptr, char* text){
	static_cast<QAbstractButton*>(ptr)->setText(QString(text));
}

char* QAbstractButton_Text(QtObjectPtr ptr){
	return static_cast<QAbstractButton*>(ptr)->text().toUtf8().data();
}

void QAbstractButton_Toggle(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "toggle");
}

void QAbstractButton_AnimateClick(QtObjectPtr ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "animateClick", Q_ARG(int, msec));
}

void QAbstractButton_Click(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "click");
}

void QAbstractButton_ConnectClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Clicked));;
}

void QAbstractButton_DisconnectClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Clicked));;
}

QtObjectPtr QAbstractButton_Group(QtObjectPtr ptr){
	return static_cast<QAbstractButton*>(ptr)->group();
}

void QAbstractButton_ConnectPressed(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::pressed), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Pressed));;
}

void QAbstractButton_DisconnectPressed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::pressed), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Pressed));;
}

void QAbstractButton_ConnectReleased(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::released), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Released));;
}

void QAbstractButton_DisconnectReleased(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::released), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Released));;
}

void QAbstractButton_ConnectToggled(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::toggled), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Toggled));;
}

void QAbstractButton_DisconnectToggled(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::toggled), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Toggled));;
}

void QAbstractButton_DestroyQAbstractButton(QtObjectPtr ptr){
	static_cast<QAbstractButton*>(ptr)->~QAbstractButton();
}

