#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Types
int QValidator_Invalid();
int QValidator_Intermediate();
int QValidator_Acceptable();
//Public Functions
void QValidator_Destroy(QtObjectPtr ptr);
//Signals
void QValidator_ConnectSignalChanged(QtObjectPtr ptr);
void QValidator_DisconnectSignalChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
