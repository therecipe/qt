#ifdef __cplusplus
extern "C" {
#endif

void* QQmlExpression_NewQQmlExpression();
void* QQmlExpression_NewQQmlExpression2(void* ctxt, void* scope, char* expression, void* parent);
void* QQmlExpression_NewQQmlExpression3(void* script, void* ctxt, void* scope, void* parent);
void QQmlExpression_ClearError(void* ptr);
int QQmlExpression_ColumnNumber(void* ptr);
void* QQmlExpression_Context(void* ptr);
void* QQmlExpression_Engine(void* ptr);
void* QQmlExpression_Evaluate(void* ptr, int valueIsUndefined);
char* QQmlExpression_Expression(void* ptr);
int QQmlExpression_HasError(void* ptr);
int QQmlExpression_LineNumber(void* ptr);
int QQmlExpression_NotifyOnValueChanged(void* ptr);
void* QQmlExpression_ScopeObject(void* ptr);
void QQmlExpression_SetExpression(void* ptr, char* expression);
void QQmlExpression_SetNotifyOnValueChanged(void* ptr, int notifyOnChange);
void QQmlExpression_SetSourceLocation(void* ptr, char* url, int line, int column);
char* QQmlExpression_SourceFile(void* ptr);
void QQmlExpression_ConnectValueChanged(void* ptr);
void QQmlExpression_DisconnectValueChanged(void* ptr);
void QQmlExpression_DestroyQQmlExpression(void* ptr);

#ifdef __cplusplus
}
#endif