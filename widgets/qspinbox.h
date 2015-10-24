#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QSpinBox_CleanText(QtObjectPtr ptr);
int QSpinBox_DisplayIntegerBase(QtObjectPtr ptr);
int QSpinBox_Maximum(QtObjectPtr ptr);
int QSpinBox_Minimum(QtObjectPtr ptr);
char* QSpinBox_Prefix(QtObjectPtr ptr);
void QSpinBox_SetDisplayIntegerBase(QtObjectPtr ptr, int base);
void QSpinBox_SetMaximum(QtObjectPtr ptr, int max);
void QSpinBox_SetMinimum(QtObjectPtr ptr, int min);
void QSpinBox_SetPrefix(QtObjectPtr ptr, char* prefix);
void QSpinBox_SetSingleStep(QtObjectPtr ptr, int val);
void QSpinBox_SetSuffix(QtObjectPtr ptr, char* suffix);
void QSpinBox_SetValue(QtObjectPtr ptr, int val);
int QSpinBox_SingleStep(QtObjectPtr ptr);
char* QSpinBox_Suffix(QtObjectPtr ptr);
int QSpinBox_Value(QtObjectPtr ptr);
QtObjectPtr QSpinBox_NewQSpinBox(QtObjectPtr parent);
void QSpinBox_SetRange(QtObjectPtr ptr, int minimum, int maximum);
void QSpinBox_ConnectValueChanged(QtObjectPtr ptr);
void QSpinBox_DisconnectValueChanged(QtObjectPtr ptr);
void QSpinBox_DestroyQSpinBox(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif