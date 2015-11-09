#include "qstylehints.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
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

int QStyleHints_CursorFlashTime(void* ptr){
	return static_cast<QStyleHints*>(ptr)->cursorFlashTime();
}

double QStyleHints_FontSmoothingGamma(void* ptr){
	return static_cast<double>(static_cast<QStyleHints*>(ptr)->fontSmoothingGamma());
}

int QStyleHints_KeyboardAutoRepeatRate(void* ptr){
	return static_cast<QStyleHints*>(ptr)->keyboardAutoRepeatRate();
}

int QStyleHints_KeyboardInputInterval(void* ptr){
	return static_cast<QStyleHints*>(ptr)->keyboardInputInterval();
}

int QStyleHints_MouseDoubleClickInterval(void* ptr){
	return static_cast<QStyleHints*>(ptr)->mouseDoubleClickInterval();
}

int QStyleHints_MousePressAndHoldInterval(void* ptr){
	return static_cast<QStyleHints*>(ptr)->mousePressAndHoldInterval();
}

int QStyleHints_PasswordMaskDelay(void* ptr){
	return static_cast<QStyleHints*>(ptr)->passwordMaskDelay();
}

int QStyleHints_SetFocusOnTouchRelease(void* ptr){
	return static_cast<QStyleHints*>(ptr)->setFocusOnTouchRelease();
}

int QStyleHints_ShowIsFullScreen(void* ptr){
	return static_cast<QStyleHints*>(ptr)->showIsFullScreen();
}

int QStyleHints_SingleClickActivation(void* ptr){
	return static_cast<QStyleHints*>(ptr)->singleClickActivation();
}

int QStyleHints_StartDragDistance(void* ptr){
	return static_cast<QStyleHints*>(ptr)->startDragDistance();
}

int QStyleHints_StartDragTime(void* ptr){
	return static_cast<QStyleHints*>(ptr)->startDragTime();
}

int QStyleHints_StartDragVelocity(void* ptr){
	return static_cast<QStyleHints*>(ptr)->startDragVelocity();
}

int QStyleHints_TabFocusBehavior(void* ptr){
	return static_cast<QStyleHints*>(ptr)->tabFocusBehavior();
}

int QStyleHints_UseRtlExtensions(void* ptr){
	return static_cast<QStyleHints*>(ptr)->useRtlExtensions();
}

void QStyleHints_ConnectCursorFlashTimeChanged(void* ptr){
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::cursorFlashTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_CursorFlashTimeChanged));;
}

void QStyleHints_DisconnectCursorFlashTimeChanged(void* ptr){
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::cursorFlashTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_CursorFlashTimeChanged));;
}

void QStyleHints_ConnectKeyboardInputIntervalChanged(void* ptr){
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::keyboardInputIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_KeyboardInputIntervalChanged));;
}

void QStyleHints_DisconnectKeyboardInputIntervalChanged(void* ptr){
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::keyboardInputIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_KeyboardInputIntervalChanged));;
}

void QStyleHints_ConnectMouseDoubleClickIntervalChanged(void* ptr){
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::mouseDoubleClickIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_MouseDoubleClickIntervalChanged));;
}

void QStyleHints_DisconnectMouseDoubleClickIntervalChanged(void* ptr){
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::mouseDoubleClickIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_MouseDoubleClickIntervalChanged));;
}

void QStyleHints_ConnectStartDragDistanceChanged(void* ptr){
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragDistanceChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragDistanceChanged));;
}

void QStyleHints_DisconnectStartDragDistanceChanged(void* ptr){
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragDistanceChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragDistanceChanged));;
}

void QStyleHints_ConnectStartDragTimeChanged(void* ptr){
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragTimeChanged));;
}

void QStyleHints_DisconnectStartDragTimeChanged(void* ptr){
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragTimeChanged));;
}

