#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QDoubleValidator_Notation(QtObjectPtr ptr);
void QDoubleValidator_SetDecimals(QtObjectPtr ptr, int v);
void QDoubleValidator_SetNotation(QtObjectPtr ptr, int v);
QtObjectPtr QDoubleValidator_NewQDoubleValidator(QtObjectPtr parent);
int QDoubleValidator_Decimals(QtObjectPtr ptr);
void QDoubleValidator_DestroyQDoubleValidator(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif