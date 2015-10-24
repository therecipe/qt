#include "qinputmethod.h"
#include <QString>
#include <QUrl>
#include <QRect>
#include <QRectF>
#include <QVariant>
#include <QModelIndex>
#include <QTransform>
#include <QMetaObject>
#include <QObject>
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

int QInputMethod_InputDirection(QtObjectPtr ptr){
	return static_cast<QInputMethod*>(ptr)->inputDirection();
}

int QInputMethod_IsAnimating(QtObjectPtr ptr){
	return static_cast<QInputMethod*>(ptr)->isAnimating();
}

int QInputMethod_IsVisible(QtObjectPtr ptr){
	return static_cast<QInputMethod*>(ptr)->isVisible();
}

void QInputMethod_ConnectAnimatingChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::animatingChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_AnimatingChanged));;
}

void QInputMethod_DisconnectAnimatingChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::animatingChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_AnimatingChanged));;
}

void QInputMethod_Commit(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "commit");
}

void QInputMethod_ConnectCursorRectangleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::cursorRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_CursorRectangleChanged));;
}

void QInputMethod_DisconnectCursorRectangleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::cursorRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_CursorRectangleChanged));;
}

void QInputMethod_Hide(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "hide");
}

void QInputMethod_ConnectInputDirectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)(Qt::LayoutDirection)>(&QInputMethod::inputDirectionChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)(Qt::LayoutDirection)>(&MyQInputMethod::Signal_InputDirectionChanged));;
}

void QInputMethod_DisconnectInputDirectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)(Qt::LayoutDirection)>(&QInputMethod::inputDirectionChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)(Qt::LayoutDirection)>(&MyQInputMethod::Signal_InputDirectionChanged));;
}

void QInputMethod_InvokeAction(QtObjectPtr ptr, int a, int cursorPosition){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "invokeAction", Q_ARG(QInputMethod::Action, static_cast<QInputMethod::Action>(a)), Q_ARG(int, cursorPosition));
}

void QInputMethod_ConnectKeyboardRectangleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::keyboardRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_KeyboardRectangleChanged));;
}

void QInputMethod_DisconnectKeyboardRectangleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::keyboardRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_KeyboardRectangleChanged));;
}

void QInputMethod_ConnectLocaleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::localeChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_LocaleChanged));;
}

void QInputMethod_DisconnectLocaleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::localeChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_LocaleChanged));;
}

char* QInputMethod_QInputMethod_QueryFocusObject(int query, char* argument){
	return QInputMethod::queryFocusObject(static_cast<Qt::InputMethodQuery>(query), QVariant(argument)).toString().toUtf8().data();
}

void QInputMethod_Reset(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "reset");
}

void QInputMethod_SetInputItemRectangle(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QInputMethod*>(ptr)->setInputItemRectangle(*static_cast<QRectF*>(rect));
}

void QInputMethod_SetInputItemTransform(QtObjectPtr ptr, QtObjectPtr transform){
	static_cast<QInputMethod*>(ptr)->setInputItemTransform(*static_cast<QTransform*>(transform));
}

void QInputMethod_SetVisible(QtObjectPtr ptr, int visible){
	static_cast<QInputMethod*>(ptr)->setVisible(visible != 0);
}

void QInputMethod_Show(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "show");
}

void QInputMethod_Update(QtObjectPtr ptr, int queries){
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "update", Q_ARG(Qt::InputMethodQuery, static_cast<Qt::InputMethodQuery>(queries)));
}

void QInputMethod_ConnectVisibleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::visibleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_VisibleChanged));;
}

void QInputMethod_DisconnectVisibleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::visibleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_VisibleChanged));;
}

