#include "qqmlexpression.h"
#include <QQmlScriptString>
#include <QObject>
#include <QQmlContext>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlExpression>
#include "_cgo_export.h"

class MyQQmlExpression: public QQmlExpression {
public:
void Signal_ValueChanged(){callbackQQmlExpressionValueChanged(this->objectName().toUtf8().data());};
};

void* QQmlExpression_NewQQmlExpression(){
	return new QQmlExpression();
}

void* QQmlExpression_NewQQmlExpression2(void* ctxt, void* scope, char* expression, void* parent){
	return new QQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QObject*>(scope), QString(expression), static_cast<QObject*>(parent));
}

void* QQmlExpression_NewQQmlExpression3(void* script, void* ctxt, void* scope, void* parent){
	return new QQmlExpression(*static_cast<QQmlScriptString*>(script), static_cast<QQmlContext*>(ctxt), static_cast<QObject*>(scope), static_cast<QObject*>(parent));
}

void QQmlExpression_ClearError(void* ptr){
	static_cast<QQmlExpression*>(ptr)->clearError();
}

int QQmlExpression_ColumnNumber(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->columnNumber();
}

void* QQmlExpression_Context(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->context();
}

void* QQmlExpression_Engine(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->engine();
}

void* QQmlExpression_Evaluate(void* ptr, int valueIsUndefined){
	return new QVariant(static_cast<QQmlExpression*>(ptr)->evaluate(NULL));
}

char* QQmlExpression_Expression(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->expression().toUtf8().data();
}

int QQmlExpression_HasError(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->hasError();
}

int QQmlExpression_LineNumber(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->lineNumber();
}

int QQmlExpression_NotifyOnValueChanged(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->notifyOnValueChanged();
}

void* QQmlExpression_ScopeObject(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->scopeObject();
}

void QQmlExpression_SetExpression(void* ptr, char* expression){
	static_cast<QQmlExpression*>(ptr)->setExpression(QString(expression));
}

void QQmlExpression_SetNotifyOnValueChanged(void* ptr, int notifyOnChange){
	static_cast<QQmlExpression*>(ptr)->setNotifyOnValueChanged(notifyOnChange != 0);
}

void QQmlExpression_SetSourceLocation(void* ptr, char* url, int line, int column){
	static_cast<QQmlExpression*>(ptr)->setSourceLocation(QString(url), line, column);
}

char* QQmlExpression_SourceFile(void* ptr){
	return static_cast<QQmlExpression*>(ptr)->sourceFile().toUtf8().data();
}

void QQmlExpression_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QQmlExpression*>(ptr), static_cast<void (QQmlExpression::*)()>(&QQmlExpression::valueChanged), static_cast<MyQQmlExpression*>(ptr), static_cast<void (MyQQmlExpression::*)()>(&MyQQmlExpression::Signal_ValueChanged));;
}

void QQmlExpression_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlExpression*>(ptr), static_cast<void (QQmlExpression::*)()>(&QQmlExpression::valueChanged), static_cast<MyQQmlExpression*>(ptr), static_cast<void (MyQQmlExpression::*)()>(&MyQQmlExpression::Signal_ValueChanged));;
}

void QQmlExpression_DestroyQQmlExpression(void* ptr){
	static_cast<QQmlExpression*>(ptr)->~QQmlExpression();
}

