#ifdef __cplusplus
extern "C" {
#endif

char* QSpinBox_CleanText(void* ptr);
int QSpinBox_DisplayIntegerBase(void* ptr);
int QSpinBox_Maximum(void* ptr);
int QSpinBox_Minimum(void* ptr);
char* QSpinBox_Prefix(void* ptr);
void QSpinBox_SetDisplayIntegerBase(void* ptr, int base);
void QSpinBox_SetMaximum(void* ptr, int max);
void QSpinBox_SetMinimum(void* ptr, int min);
void QSpinBox_SetPrefix(void* ptr, char* prefix);
void QSpinBox_SetSingleStep(void* ptr, int val);
void QSpinBox_SetSuffix(void* ptr, char* suffix);
void QSpinBox_SetValue(void* ptr, int val);
int QSpinBox_SingleStep(void* ptr);
char* QSpinBox_Suffix(void* ptr);
int QSpinBox_Value(void* ptr);
void* QSpinBox_NewQSpinBox(void* parent);
void QSpinBox_SetRange(void* ptr, int minimum, int maximum);
void QSpinBox_ConnectValueChanged(void* ptr);
void QSpinBox_DisconnectValueChanged(void* ptr);
void QSpinBox_DestroyQSpinBox(void* ptr);

#ifdef __cplusplus
}
#endif