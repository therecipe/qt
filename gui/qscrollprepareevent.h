#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScrollPrepareEvent_NewQScrollPrepareEvent(QtObjectPtr startPos);
void QScrollPrepareEvent_SetContentPos(QtObjectPtr ptr, QtObjectPtr pos);
void QScrollPrepareEvent_SetContentPosRange(QtObjectPtr ptr, QtObjectPtr rect);
void QScrollPrepareEvent_SetViewportSize(QtObjectPtr ptr, QtObjectPtr size);
void QScrollPrepareEvent_DestroyQScrollPrepareEvent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif