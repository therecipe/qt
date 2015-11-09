#ifdef __cplusplus
extern "C" {
#endif

void QGradient_SetColorAt(void* ptr, double position, void* color);
int QGradient_CoordinateMode(void* ptr);
void QGradient_SetCoordinateMode(void* ptr, int mode);
void QGradient_SetSpread(void* ptr, int method);
int QGradient_Spread(void* ptr);
int QGradient_Type(void* ptr);

#ifdef __cplusplus
}
#endif