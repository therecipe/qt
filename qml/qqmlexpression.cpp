#include "qqmlexpression.h"
#include <QUrl>
#include <QModelIndex>
#include <QQmlScriptString>
#include <QQmlContext>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QQmlExpression>
#include "_cgo_export.h"

class MyQQmlExpression: public QQmlExpression {
public:
void Signal_ValueChanged(){callbackQQmlExpressionValueChanged(this->objectName().toUtf8().data());};
};

QtObjectPtr QQmlExpression_NewQQmlExpression(){
	return new QQmlExpression();
}

QtObjectPtr QQmlExpression_NewQQmlExpression2(QtObjectPtr ctxt, QtObjectPtr scope, char* expression, QtObjectPtr parent){
	return new QQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QObject*>(scope), QString(expression), static_cast<QObject*>(parent));
}

QtObjectPtr QQmlExpression_NewQQmlExpression3(QtObjectPtr script, QtObjectPtr ctxt, QtObjectPtr scope, QtObjectPtr parent){
	return new QQmlExpression(*static_cast<QQmlScriptString*>(script), static_cast<QQmlContext*>(ctxt), static_cast<QObject*>(scope), static_cast<QObject*>(parent));
}

void QQmlExpression_ClearError(QtObjectPtr ptr){
	static_cast<QQmlExpression*>(ptr)->clearError();
}

int QQmlExpression_ColumnNumber(QtObjectPtr ptr){
	return static_cast<QQmlExpression*>(ptr)->columnNumber();
}

QtObjectPtr QQmlExpression_Context(QtObjectPtr ptr){
	return static_cast<QQmlExpression*>(ptr)->context();
}

QtObjectPtr QQmlExpression_Engine(QtObjectPtr ptr){
	return static_cast<QQmlExpression*>(ptr)->engine();
}

char* QQmlExpression_Evaluate(QtObjectPtr ptr, int valueIsUndefined){
	return static_cast<QQmlExpression*>(ptr)->evaluate(NULL).toString().toUtf8().data();
}

char* QQmlExpression_Expression(QtObjectPtr ptr){
	return static_cast<QQmlExpression*>(ptr)->expression().toUtf8().data();
}

int QQmlExpression_HasError(QtObjectPtr ptr){
	return static_cast<QQmlExpression*>(ptr)->hasError();
}

int QQmlExpression_LineNumber(QtObjectPtr ptr){
	return static_cast<QQmlExpression*>(ptr)->lineNumber();
}

int QQmlExpression_NotifyOnValueChanged(QtObjectPtr ptr){
	return static_cast<QQmlExpression*>(ptr)->notifyOnValueChanged();
}

QtObjectPtr QQmlExpression_ScopeObject(QtObjectPtr ptr){
	return static_cast<QQmlExpression*>(ptr)->scopeObject();
}

void QQmlExpression_SetExpression(QtObjectPtr ptr, char* expression){
	static_cast<QQmlExpression*>(ptr)->setExpression(QString(expression));
}

void QQmlExpression_SetNotifyOnValueChanged(QtObjectPtr ptr, int notifyOnChange){
	static_cast<QQmlExpression*>(ptr)->setNotifyOnValueChanged(notifyOnChange != 0);
}

void QQmlExpression_SetSourceLocation(QtObjectPtr ptr, char* url, int line, int column){
	static_cast<QQmlExpression*>(ptr)->setSourceLocation(QString(url), line, column);
}

char* QQmlExpression_SourceFile(QtObjectPtr ptr){
	return static_cast<QQmlExpression*>(ptr)->sourceFile().toUtf8().data();
}

void QQmlExpression_ConnectValueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQmlExpression*>(ptr), static_cast<void (QQmlExpression::*)()>(&QQmlExpression::valueChanged), static_cast<MyQQmlExpression*>(ptr), static_cast<void (MyQQmlExpression::*)()>(&MyQQmlExpression::Signal_ValueChanged));;
}

void QQmlExpression_DisconnectValueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQmlExpression*>(ptr), static_cast<void (QQmlExpression::*)()>(&QQmlExpression::valueChanged), static_cast<MyQQmlExpression*>(ptr), static_cast<void (MyQQmlExpression::*)()>(&MyQQmlExpression::Signal_ValueChanged));;
}

void QQmlExpression_DestroyQQmlExpression(QtObjectPtr ptr){
	static_cast<QQmlExpression*>(ptr)->~QQmlExpression();
}

