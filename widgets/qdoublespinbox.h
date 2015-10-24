#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QDoubleSpinBox_CleanText(QtObjectPtr ptr);
int QDoubleSpinBox_Decimals(QtObjectPtr ptr);
char* QDoubleSpinBox_Prefix(QtObjectPtr ptr);
void QDoubleSpinBox_SetDecimals(QtObjectPtr ptr, int prec);
void QDoubleSpinBox_SetPrefix(QtObjectPtr ptr, char* prefix);
void QDoubleSpinBox_SetSuffix(QtObjectPtr ptr, char* suffix);
char* QDoubleSpinBox_Suffix(QtObjectPtr ptr);
QtObjectPtr QDoubleSpinBox_NewQDoubleSpinBox(QtObjectPtr parent);
void QDoubleSpinBox_DestroyQDoubleSpinBox(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif