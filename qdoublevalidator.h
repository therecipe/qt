#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Types
int QDoubleValidator_StandardNotation();
int QDoubleValidator_ScientificNotation();
//Public Functions
QtObjectPtr QDoubleValidator_New_QObject(QtObjectPtr parent);
void QDoubleValidator_Destroy(QtObjectPtr ptr);
int QDoubleValidator_Decimals(QtObjectPtr ptr);
int QDoubleValidator_Notation(QtObjectPtr ptr);
void QDoubleValidator_SetDecimals_Int(QtObjectPtr ptr, int in);
void QDoubleValidator_SetNotation_Notation(QtObjectPtr ptr, int No);

#ifdef __cplusplus
}
#endif
