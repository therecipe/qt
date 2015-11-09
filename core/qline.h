#ifdef __cplusplus
extern "C" {
#endif

void* QLine_NewQLine();
void* QLine_NewQLine2(void* p1, void* p2);
void* QLine_NewQLine3(int x1, int y1, int x2, int y2);
int QLine_Dx(void* ptr);
int QLine_Dy(void* ptr);
int QLine_IsNull(void* ptr);
void QLine_SetLine(void* ptr, int x1, int y1, int x2, int y2);
void QLine_SetP1(void* ptr, void* p1);
void QLine_SetP2(void* ptr, void* p2);
void QLine_SetPoints(void* ptr, void* p1, void* p2);
void QLine_Translate(void* ptr, void* offset);
void QLine_Translate2(void* ptr, int dx, int dy);
int QLine_X1(void* ptr);
int QLine_X2(void* ptr);
int QLine_Y1(void* ptr);
int QLine_Y2(void* ptr);

#ifdef __cplusplus
}
#endif