#ifdef __cplusplus
extern "C" {
#endif

void* QRegularExpressionValidator_RegularExpression(void* ptr);
void QRegularExpressionValidator_SetRegularExpression(void* ptr, void* re);
void* QRegularExpressionValidator_NewQRegularExpressionValidator(void* parent);
void* QRegularExpressionValidator_NewQRegularExpressionValidator2(void* re, void* parent);
void QRegularExpressionValidator_ConnectRegularExpressionChanged(void* ptr);
void QRegularExpressionValidator_DisconnectRegularExpressionChanged(void* ptr);
void QRegularExpressionValidator_DestroyQRegularExpressionValidator(void* ptr);

#ifdef __cplusplus
}
#endif