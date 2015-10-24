#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QErrorMessage_NewQErrorMessage(QtObjectPtr parent);
QtObjectPtr QErrorMessage_QErrorMessage_QtHandler();
void QErrorMessage_ShowMessage(QtObjectPtr ptr, char* message);
void QErrorMessage_ShowMessage2(QtObjectPtr ptr, char* message, char* ty);
void QErrorMessage_DestroyQErrorMessage(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif