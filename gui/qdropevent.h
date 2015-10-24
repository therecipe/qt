#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QDropEvent_SetDropAction(QtObjectPtr ptr, int action);
QtObjectPtr QDropEvent_NewQDropEvent(QtObjectPtr pos, int actions, QtObjectPtr data, int buttons, int modifiers, int ty);
void QDropEvent_AcceptProposedAction(QtObjectPtr ptr);
int QDropEvent_DropAction(QtObjectPtr ptr);
int QDropEvent_KeyboardModifiers(QtObjectPtr ptr);
QtObjectPtr QDropEvent_MimeData(QtObjectPtr ptr);
int QDropEvent_MouseButtons(QtObjectPtr ptr);
int QDropEvent_PossibleActions(QtObjectPtr ptr);
int QDropEvent_ProposedAction(QtObjectPtr ptr);
QtObjectPtr QDropEvent_Source(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif