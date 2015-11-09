#ifdef __cplusplus
extern "C" {
#endif

double QLineF_AngleTo(void* ptr, void* line);
int QLineF_Intersect(void* ptr, void* line, void* intersectionPoint);
void* QLineF_NewQLineF();
void* QLineF_NewQLineF4(void* line);
void* QLineF_NewQLineF2(void* p1, void* p2);
void* QLineF_NewQLineF3(double x1, double y1, double x2, double y2);
double QLineF_Angle(void* ptr);
double QLineF_Dx(void* ptr);
double QLineF_Dy(void* ptr);
int QLineF_IsNull(void* ptr);
double QLineF_Length(void* ptr);
void QLineF_SetAngle(void* ptr, double angle);
void QLineF_SetLength(void* ptr, double length);
void QLineF_SetLine(void* ptr, double x1, double y1, double x2, double y2);
void QLineF_SetP1(void* ptr, void* p1);
void QLineF_SetP2(void* ptr, void* p2);
void QLineF_SetPoints(void* ptr, void* p1, void* p2);
void QLineF_Translate(void* ptr, void* offset);
void QLineF_Translate2(void* ptr, double dx, double dy);
double QLineF_X1(void* ptr);
double QLineF_X2(void* ptr);
double QLineF_Y1(void* ptr);
double QLineF_Y2(void* ptr);

#ifdef __cplusplus
}
#endif