#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QInputMethodEvent_NewQInputMethodEvent();
QtObjectPtr QInputMethodEvent_NewQInputMethodEvent3(QtObjectPtr other);
char* QInputMethodEvent_CommitString(QtObjectPtr ptr);
char* QInputMethodEvent_PreeditString(QtObjectPtr ptr);
int QInputMethodEvent_ReplacementLength(QtObjectPtr ptr);
int QInputMethodEvent_ReplacementStart(QtObjectPtr ptr);
void QInputMethodEvent_SetCommitString(QtObjectPtr ptr, char* commitString, int replaceFrom, int replaceLength);

#ifdef __cplusplus
}
#endif