#include "qstylehints.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QStyleHints>
#include "_cgo_export.h"

class MyQStyleHints: public QStyleHints {
public:
void Signal_CursorFlashTimeChanged(int cursorFlashTime){callbackQStyleHintsCursorFlashTimeChanged(this->objectName().toUtf8().data(), cursorFlashTime);};
void Signal_KeyboardInputIntervalChanged(int keyboardInputInterval){callbackQStyleHintsKeyboardInputIntervalChanged(this->objectName().toUtf8().data(), keyboardInputInterval);};
void Signal_MouseDoubleClickIntervalChanged(int mouseDoubleClickInterval){callbackQStyleHintsMouseDoubleClickIntervalChanged(this->objectName().toUtf8().data(), mouseDoubleClickInterval);};
void Signal_StartDragDistanceChanged(int startDragDistance){callbackQStyleHintsStartDragDistanceChanged(this->objectName().toUtf8().data(), startDragDistance);};
void Signal_StartDragTimeChanged(int startDragTime){callbackQStyleHintsStartDragTimeChanged(this->objectName().toUtf8().data(), startDragTime);};
};

int QStyleHints_CursorFlashTime(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->cursorFlashTime();
}

int QStyleHints_KeyboardAutoRepeatRate(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->keyboardAutoRepeatRate();
}

int QStyleHints_KeyboardInputInterval(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->keyboardInputInterval();
}

int QStyleHints_MouseDoubleClickInterval(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->mouseDoubleClickInterval();
}

int QStyleHints_MousePressAndHoldInterval(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->mousePressAndHoldInterval();
}

int QStyleHints_PasswordMaskDelay(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->passwordMaskDelay();
}

int QStyleHints_SetFocusOnTouchRelease(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->setFocusOnTouchRelease();
}

int QStyleHints_ShowIsFullScreen(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->showIsFullScreen();
}

int QStyleHints_SingleClickActivation(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->singleClickActivation();
}

int QStyleHints_StartDragDistance(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->startDragDistance();
}

int QStyleHints_StartDragTime(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->startDragTime();
}

int QStyleHints_StartDragVelocity(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->startDragVelocity();
}

int QStyleHints_TabFocusBehavior(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->tabFocusBehavior();
}

int QStyleHints_UseRtlExtensions(QtObjectPtr ptr){
	return static_cast<QStyleHints*>(ptr)->useRtlExtensions();
}

void QStyleHints_ConnectCursorFlashTimeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::cursorFlashTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_CursorFlashTimeChanged));;
}

void QStyleHints_DisconnectCursorFlashTimeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::cursorFlashTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_CursorFlashTimeChanged));;
}

void QStyleHints_ConnectKeyboardInputIntervalChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::keyboardInputIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_KeyboardInputIntervalChanged));;
}

void QStyleHints_DisconnectKeyboardInputIntervalChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::keyboardInputIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_KeyboardInputIntervalChanged));;
}

void QStyleHints_ConnectMouseDoubleClickIntervalChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::mouseDoubleClickIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_MouseDoubleClickIntervalChanged));;
}

void QStyleHints_DisconnectMouseDoubleClickIntervalChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::mouseDoubleClickIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_MouseDoubleClickIntervalChanged));;
}

void QStyleHints_ConnectStartDragDistanceChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragDistanceChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragDistanceChanged));;
}

void QStyleHints_DisconnectStartDragDistanceChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragDistanceChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragDistanceChanged));;
}

void QStyleHints_ConnectStartDragTimeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragTimeChanged));;
}

void QStyleHints_DisconnectStartDragTimeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragTimeChanged));;
}

