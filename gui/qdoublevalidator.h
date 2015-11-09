#ifdef __cplusplus
extern "C" {
#endif

int QDoubleValidator_Notation(void* ptr);
void QDoubleValidator_SetDecimals(void* ptr, int v);
void QDoubleValidator_SetNotation(void* ptr, int v);
void* QDoubleValidator_NewQDoubleValidator(void* parent);
int QDoubleValidator_Decimals(void* ptr);
void QDoubleValidator_DestroyQDoubleValidator(void* ptr);

#ifdef __cplusplus
}
#endif