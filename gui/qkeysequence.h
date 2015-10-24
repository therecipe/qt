#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QKeySequence_NewQKeySequence();
QtObjectPtr QKeySequence_NewQKeySequence5(int key);
QtObjectPtr QKeySequence_NewQKeySequence4(QtObjectPtr keysequence);
QtObjectPtr QKeySequence_NewQKeySequence2(char* key, int format);
QtObjectPtr QKeySequence_NewQKeySequence3(int k1, int k2, int k3, int k4);
int QKeySequence_Count(QtObjectPtr ptr);
int QKeySequence_IsEmpty(QtObjectPtr ptr);
int QKeySequence_Matches(QtObjectPtr ptr, QtObjectPtr seq);
void QKeySequence_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QKeySequence_ToString(QtObjectPtr ptr, int format);
void QKeySequence_DestroyQKeySequence(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif