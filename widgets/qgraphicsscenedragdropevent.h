#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QGraphicsSceneDragDropEvent_AcceptProposedAction(QtObjectPtr ptr);
int QGraphicsSceneDragDropEvent_Buttons(QtObjectPtr ptr);
int QGraphicsSceneDragDropEvent_DropAction(QtObjectPtr ptr);
QtObjectPtr QGraphicsSceneDragDropEvent_MimeData(QtObjectPtr ptr);
int QGraphicsSceneDragDropEvent_Modifiers(QtObjectPtr ptr);
int QGraphicsSceneDragDropEvent_PossibleActions(QtObjectPtr ptr);
int QGraphicsSceneDragDropEvent_ProposedAction(QtObjectPtr ptr);
void QGraphicsSceneDragDropEvent_SetDropAction(QtObjectPtr ptr, int action);
QtObjectPtr QGraphicsSceneDragDropEvent_Source(QtObjectPtr ptr);
void QGraphicsSceneDragDropEvent_DestroyQGraphicsSceneDragDropEvent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif