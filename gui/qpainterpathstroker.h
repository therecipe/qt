#ifdef __cplusplus
extern "C" {
#endif

void* QPainterPathStroker_NewQPainterPathStroker();
void* QPainterPathStroker_NewQPainterPathStroker2(void* pen);
int QPainterPathStroker_CapStyle(void* ptr);
double QPainterPathStroker_CurveThreshold(void* ptr);
double QPainterPathStroker_DashOffset(void* ptr);
int QPainterPathStroker_JoinStyle(void* ptr);
double QPainterPathStroker_MiterLimit(void* ptr);
void QPainterPathStroker_SetCapStyle(void* ptr, int style);
void QPainterPathStroker_SetCurveThreshold(void* ptr, double threshold);
void QPainterPathStroker_SetDashOffset(void* ptr, double offset);
void QPainterPathStroker_SetDashPattern(void* ptr, int style);
void QPainterPathStroker_SetJoinStyle(void* ptr, int style);
void QPainterPathStroker_SetMiterLimit(void* ptr, double limit);
void QPainterPathStroker_SetWidth(void* ptr, double width);
double QPainterPathStroker_Width(void* ptr);
void QPainterPathStroker_DestroyQPainterPathStroker(void* ptr);

#ifdef __cplusplus
}
#endif