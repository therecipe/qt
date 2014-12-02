#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QSpinBox_New_QWidget(QtObjectPtr parent);
void QSpinBox_Destroy(QtObjectPtr ptr);
char* QSpinBox_CleanText(QtObjectPtr ptr);
int QSpinBox_DisplayIntegerBase(QtObjectPtr ptr);
int QSpinBox_Maximum(QtObjectPtr ptr);
int QSpinBox_Minimum(QtObjectPtr ptr);
char* QSpinBox_Prefix(QtObjectPtr ptr);
void QSpinBox_SetDisplayIntegerBase_Int(QtObjectPtr ptr, int base);
void QSpinBox_SetMaximum_Int(QtObjectPtr ptr, int max);
void QSpinBox_SetMinimum_Int(QtObjectPtr ptr, int min);
void QSpinBox_SetPrefix_String(QtObjectPtr ptr, char* prefix);
void QSpinBox_SetRange_Int_Int(QtObjectPtr ptr, int minimum, int maximum);
void QSpinBox_SetSingleStep_Int(QtObjectPtr ptr, int val);
void QSpinBox_SetSuffix_String(QtObjectPtr ptr, char* suffix);
int QSpinBox_SingleStep(QtObjectPtr ptr);
char* QSpinBox_Suffix(QtObjectPtr ptr);
int QSpinBox_Value(QtObjectPtr ptr);
//Public Slots
void QSpinBox_ConnectSlotSetValue(QtObjectPtr ptr);
void QSpinBox_DisconnectSlotSetValue(QtObjectPtr ptr);
void QSpinBox_SetValue_Int(QtObjectPtr ptr, int val);

#ifdef __cplusplus
}
#endif
