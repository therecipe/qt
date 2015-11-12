#include "qinputmethod.h"
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QTransform>
#include <QMetaObject>
#include <QUrl>
#include <QObject>
#include <QRectF>
#include <QRect>
#include <QInputMethod>
#include "_cgo_export.h"

class MyQInputMethod: public QInputMethod {
public:
void Signal_AnimatingChanged(){callbackQInputMethodAnimatingChanged(this->objectName().toUtf8().data());};
void Signal_CursorRectangleChanged(){callbackQInputMethodCursorRectangleChanged(this->objectName().toUtf8().data());};
void Signal_InputDirectionChanged(Qt::LayoutDirection newDirection){callbackQInputMethodInputDirectionChanged(this->objectName().toUtf8().data(), newDirection);};
void Signal_KeyboardRectangleChanged(){callbackQInputMethodKeyboardRectangleChanged(this->objectName().toUtf8().data());};
void Signal_LocaleChanged(){callbackQInputMethodLocaleChanged(this->objectName().toUtf8().data());};
void Signal_VisibleChanged(){callbackQInputMethodVisibleChanged(this->objectName().toUtf8().data());};
};

int QInputMethod_InputDirection(void* ptr){
	return static_cast<QInputMethod*>(ptr)->inputDirection();
}

int QInputMethod_IsAnimating(void* ptr){
	return static_cast<QInputMethod*>(ptr)->isAnimating();
}

int QInputMethod_IsVisible(void* ptr){
	return static_cast<QInputMethod*>(ptr)->isVisible();
}

void QInputMethod_ConnectAnimatingChanged(void* ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::animatingChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_AnimatingChanged));;
}

void QInputMethod_DisconnectAnimatingChanged(void* ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::animatingChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_AnimatingChanged));;
}

void QInputMethod_Commit(void* ptr){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "commit");
}

void QInputMethod_ConnectCursorRectangleChanged(void* ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::cursorRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_CursorRectangleChanged));;
}

void QInputMethod_DisconnectCursorRectangleChanged(void* ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::cursorRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_CursorRectangleChanged));;
}

void QInputMethod_Hide(void* ptr){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "hide");
}

void QInputMethod_ConnectInputDirectionChanged(void* ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)(Qt::LayoutDirection)>(&QInputMethod::inputDirectionChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)(Qt::LayoutDirection)>(&MyQInputMethod::Signal_InputDirectionChanged));;
}

void QInputMethod_DisconnectInputDirectionChanged(void* ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)(Qt::LayoutDirection)>(&QInputMethod::inputDirectionChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)(Qt::LayoutDirection)>(&MyQInputMethod::Signal_InputDirectionChanged));;
}

void QInputMethod_InvokeAction(void* ptr, int a, int cursorPosition){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "invokeAction", Q_ARG(QInputMethod::Action, static_cast<QInputMethod::Action>(a)), Q_ARG(int, cursorPosition));
}

void QInputMethod_ConnectKeyboardRectangleChanged(void* ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::keyboardRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_KeyboardRectangleChanged));;
}

void QInputMethod_DisconnectKeyboardRectangleChanged(void* ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::keyboardRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_KeyboardRectangleChanged));;
}

void QInputMethod_ConnectLocaleChanged(void* ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::localeChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_LocaleChanged));;
}

void QInputMethod_DisconnectLocaleChanged(void* ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::localeChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_LocaleChanged));;
}

void* QInputMethod_QInputMethod_QueryFocusObject(int query, void* argument){
	return new QVariant(QInputMethod::queryFocusObject(static_cast<Qt::InputMethodQuery>(query), *static_cast<QVariant*>(argument)));
}

void QInputMethod_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "reset");
}

void QInputMethod_SetInputItemRectangle(void* ptr, void* rect){
	static_cast<QInputMethod*>(ptr)->setInputItemRectangle(*static_cast<QRectF*>(rect));
}

void QInputMethod_SetInputItemTransform(void* ptr, void* transform){
	static_cast<QInputMethod*>(ptr)->setInputItemTransform(*static_cast<QTransform*>(transform));
}

void QInputMethod_SetVisible(void* ptr, int visible){
	static_cast<QInputMethod*>(ptr)->setVisible(visible != 0);
}

void QInputMethod_Show(void* ptr){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "show");
}

void QInputMethod_Update(void* ptr, int queries){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "update", Q_ARG(Qt::InputMethodQuery, static_cast<Qt::InputMethodQuery>(queries)));
}

void QInputMethod_ConnectVisibleChanged(void* ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::visibleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_VisibleChanged));;
}

void QInputMethod_DisconnectVisibleChanged(void* ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::visibleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_VisibleChanged));;
}

