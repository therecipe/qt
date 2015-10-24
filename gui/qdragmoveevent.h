#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDragMoveEvent_NewQDragMoveEvent(QtObjectPtr pos, int actions, QtObjectPtr data, int buttons, int modifiers, int ty);
void QDragMoveEvent_Accept2(QtObjectPtr ptr);
void QDragMoveEvent_Accept(QtObjectPtr ptr, QtObjectPtr rectangle);
void QDragMoveEvent_Ignore2(QtObjectPtr ptr);
void QDragMoveEvent_Ignore(QtObjectPtr ptr, QtObjectPtr rectangle);
void QDragMoveEvent_DestroyQDragMoveEvent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif