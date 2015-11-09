#ifdef __cplusplus
extern "C" {
#endif

void QRegExpValidator_SetRegExp(void* ptr, void* rx);
void* QRegExpValidator_NewQRegExpValidator(void* parent);
void* QRegExpValidator_NewQRegExpValidator2(void* rx, void* parent);
void* QRegExpValidator_RegExp(void* ptr);
void QRegExpValidator_DestroyQRegExpValidator(void* ptr);

#ifdef __cplusplus
}
#endif