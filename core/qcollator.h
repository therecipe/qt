#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCollator_CaseSensitivity(QtObjectPtr ptr);
int QCollator_IgnorePunctuation(QtObjectPtr ptr);
int QCollator_NumericMode(QtObjectPtr ptr);
void QCollator_SetCaseSensitivity(QtObjectPtr ptr, int sensitivity);
void QCollator_SetIgnorePunctuation(QtObjectPtr ptr, int on);
void QCollator_SetNumericMode(QtObjectPtr ptr, int on);
QtObjectPtr QCollator_NewQCollator3(QtObjectPtr other);
QtObjectPtr QCollator_NewQCollator2(QtObjectPtr other);
QtObjectPtr QCollator_NewQCollator(QtObjectPtr locale);
void QCollator_SetLocale(QtObjectPtr ptr, QtObjectPtr locale);
void QCollator_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QCollator_DestroyQCollator(QtObjectPtr ptr);
int QCollator_Compare3(QtObjectPtr ptr, QtObjectPtr s1, int len1, QtObjectPtr s2, int len2);
int QCollator_Compare(QtObjectPtr ptr, char* s1, char* s2);
int QCollator_Compare2(QtObjectPtr ptr, QtObjectPtr s1, QtObjectPtr s2);

#ifdef __cplusplus
}
#endif