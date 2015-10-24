#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPainterPathStroker_NewQPainterPathStroker();
QtObjectPtr QPainterPathStroker_NewQPainterPathStroker2(QtObjectPtr pen);
int QPainterPathStroker_CapStyle(QtObjectPtr ptr);
int QPainterPathStroker_JoinStyle(QtObjectPtr ptr);
void QPainterPathStroker_SetCapStyle(QtObjectPtr ptr, int style);
void QPainterPathStroker_SetDashPattern(QtObjectPtr ptr, int style);
void QPainterPathStroker_SetJoinStyle(QtObjectPtr ptr, int style);
void QPainterPathStroker_DestroyQPainterPathStroker(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif