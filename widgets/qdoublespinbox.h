#ifdef __cplusplus
extern "C" {
#endif

char* QDoubleSpinBox_CleanText(void* ptr);
int QDoubleSpinBox_Decimals(void* ptr);
char* QDoubleSpinBox_Prefix(void* ptr);
void QDoubleSpinBox_SetDecimals(void* ptr, int prec);
void QDoubleSpinBox_SetPrefix(void* ptr, char* prefix);
void QDoubleSpinBox_SetSuffix(void* ptr, char* suffix);
char* QDoubleSpinBox_Suffix(void* ptr);
void* QDoubleSpinBox_NewQDoubleSpinBox(void* parent);
void QDoubleSpinBox_DestroyQDoubleSpinBox(void* ptr);

#ifdef __cplusplus
}
#endif