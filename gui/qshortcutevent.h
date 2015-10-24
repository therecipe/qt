#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QShortcutEvent_NewQShortcutEvent(QtObjectPtr key, int id, int ambiguous);
int QShortcutEvent_IsAmbiguous(QtObjectPtr ptr);
int QShortcutEvent_ShortcutId(QtObjectPtr ptr);
void QShortcutEvent_DestroyQShortcutEvent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif