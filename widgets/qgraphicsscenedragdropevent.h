#ifdef __cplusplus
extern "C" {
#endif

void QGraphicsSceneDragDropEvent_AcceptProposedAction(void* ptr);
int QGraphicsSceneDragDropEvent_Buttons(void* ptr);
int QGraphicsSceneDragDropEvent_DropAction(void* ptr);
void* QGraphicsSceneDragDropEvent_MimeData(void* ptr);
int QGraphicsSceneDragDropEvent_Modifiers(void* ptr);
int QGraphicsSceneDragDropEvent_PossibleActions(void* ptr);
int QGraphicsSceneDragDropEvent_ProposedAction(void* ptr);
void QGraphicsSceneDragDropEvent_SetDropAction(void* ptr, int action);
void* QGraphicsSceneDragDropEvent_Source(void* ptr);
void QGraphicsSceneDragDropEvent_DestroyQGraphicsSceneDragDropEvent(void* ptr);

#ifdef __cplusplus
}
#endif