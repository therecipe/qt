#include "qabstractbutton.h"
#include <QString>
#include <QUrl>
#include <QMetaObject>
#include <QIcon>
#include <QKeySequence>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QSize>
#include <QAbstractButton>
#include "_cgo_export.h"

class MyQAbstractButton: public QAbstractButton {
public:
void Signal_Clicked(bool checked){callbackQAbstractButtonClicked(this->objectName().toUtf8().data(), checked);};
void Signal_Pressed(){callbackQAbstractButtonPressed(this->objectName().toUtf8().data());};
void Signal_Released(){callbackQAbstractButtonReleased(this->objectName().toUtf8().data());};
void Signal_Toggled(bool checked){callbackQAbstractButtonToggled(this->objectName().toUtf8().data(), checked);};
};

int QAbstractButton_AutoExclusive(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->autoExclusive();
}

int QAbstractButton_AutoRepeat(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->autoRepeat();
}

int QAbstractButton_AutoRepeatDelay(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->autoRepeatDelay();
}

int QAbstractButton_AutoRepeatInterval(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->autoRepeatInterval();
}

int QAbstractButton_IsCheckable(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->isCheckable();
}

int QAbstractButton_IsChecked(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->isChecked();
}

int QAbstractButton_IsDown(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->isDown();
}

void QAbstractButton_SetAutoExclusive(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoExclusive(v != 0);
}

void QAbstractButton_SetAutoRepeat(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoRepeat(v != 0);
}

void QAbstractButton_SetAutoRepeatDelay(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoRepeatDelay(v);
}

void QAbstractButton_SetAutoRepeatInterval(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoRepeatInterval(v);
}

void QAbstractButton_SetCheckable(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setCheckable(v != 0);
}

void QAbstractButton_SetChecked(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "setChecked", Q_ARG(bool, v != 0));
}

void QAbstractButton_SetDown(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setDown(v != 0);
}

void QAbstractButton_SetIcon(void* ptr, void* icon){
	static_cast<QAbstractButton*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QAbstractButton_SetIconSize(void* ptr, void* size){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "setIconSize", Q_ARG(QSize, *static_cast<QSize*>(size)));
}

void QAbstractButton_SetShortcut(void* ptr, void* key){
	static_cast<QAbstractButton*>(ptr)->setShortcut(*static_cast<QKeySequence*>(key));
}

void QAbstractButton_SetText(void* ptr, char* text){
	static_cast<QAbstractButton*>(ptr)->setText(QString(text));
}

char* QAbstractButton_Text(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->text().toUtf8().data();
}

void QAbstractButton_Toggle(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "toggle");
}

void QAbstractButton_AnimateClick(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "animateClick", Q_ARG(int, msec));
}

void QAbstractButton_Click(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "click");
}

void QAbstractButton_ConnectClicked(void* ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Clicked));;
}

void QAbstractButton_DisconnectClicked(void* ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Clicked));;
}

void* QAbstractButton_Group(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->group();
}

void QAbstractButton_ConnectPressed(void* ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::pressed), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Pressed));;
}

void QAbstractButton_DisconnectPressed(void* ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::pressed), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Pressed));;
}

void QAbstractButton_ConnectReleased(void* ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::released), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Released));;
}

void QAbstractButton_DisconnectReleased(void* ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::released), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Released));;
}

void QAbstractButton_ConnectToggled(void* ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::toggled), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Toggled));;
}

void QAbstractButton_DisconnectToggled(void* ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::toggled), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Toggled));;
}

void QAbstractButton_DestroyQAbstractButton(void* ptr){
	static_cast<QAbstractButton*>(ptr)->~QAbstractButton();
}

