#ifdef __cplusplus
extern "C" {
#endif

void* QPen_NewQPen4(void* brush, double width, int style, int cap, int join);
void* QPen_NewQPen5(void* pen);
void* QPen_Color(void* ptr);
void QPen_SetCapStyle(void* ptr, int style);
void QPen_SetColor(void* ptr, void* color);
void QPen_SetJoinStyle(void* ptr, int style);
void QPen_SetStyle(void* ptr, int style);
void QPen_SetWidth(void* ptr, int width);
int QPen_Style(void* ptr);
int QPen_Width(void* ptr);
double QPen_WidthF(void* ptr);
void* QPen_NewQPen();
void* QPen_NewQPen6(void* pen);
void* QPen_NewQPen2(int style);
void* QPen_NewQPen3(void* color);
void* QPen_Brush(void* ptr);
int QPen_CapStyle(void* ptr);
double QPen_DashOffset(void* ptr);
int QPen_IsCosmetic(void* ptr);
int QPen_IsSolid(void* ptr);
int QPen_JoinStyle(void* ptr);
double QPen_MiterLimit(void* ptr);
void QPen_SetBrush(void* ptr, void* brush);
void QPen_SetCosmetic(void* ptr, int cosmetic);
void QPen_SetDashOffset(void* ptr, double offset);
void QPen_SetMiterLimit(void* ptr, double limit);
void QPen_SetWidthF(void* ptr, double width);
void QPen_Swap(void* ptr, void* other);
void QPen_DestroyQPen(void* ptr);

#ifdef __cplusplus
}
#endif