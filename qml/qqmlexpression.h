#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlExpression_NewQQmlExpression();
QtObjectPtr QQmlExpression_NewQQmlExpression2(QtObjectPtr ctxt, QtObjectPtr scope, char* expression, QtObjectPtr parent);
QtObjectPtr QQmlExpression_NewQQmlExpression3(QtObjectPtr script, QtObjectPtr ctxt, QtObjectPtr scope, QtObjectPtr parent);
void QQmlExpression_ClearError(QtObjectPtr ptr);
int QQmlExpression_ColumnNumber(QtObjectPtr ptr);
QtObjectPtr QQmlExpression_Context(QtObjectPtr ptr);
QtObjectPtr QQmlExpression_Engine(QtObjectPtr ptr);
char* QQmlExpression_Evaluate(QtObjectPtr ptr, int valueIsUndefined);
char* QQmlExpression_Expression(QtObjectPtr ptr);
int QQmlExpression_HasError(QtObjectPtr ptr);
int QQmlExpression_LineNumber(QtObjectPtr ptr);
int QQmlExpression_NotifyOnValueChanged(QtObjectPtr ptr);
QtObjectPtr QQmlExpression_ScopeObject(QtObjectPtr ptr);
void QQmlExpression_SetExpression(QtObjectPtr ptr, char* expression);
void QQmlExpression_SetNotifyOnValueChanged(QtObjectPtr ptr, int notifyOnChange);
void QQmlExpression_SetSourceLocation(QtObjectPtr ptr, char* url, int line, int column);
char* QQmlExpression_SourceFile(QtObjectPtr ptr);
void QQmlExpression_ConnectValueChanged(QtObjectPtr ptr);
void QQmlExpression_DisconnectValueChanged(QtObjectPtr ptr);
void QQmlExpression_DestroyQQmlExpression(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif