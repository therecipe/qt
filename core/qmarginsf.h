#ifdef __cplusplus
extern "C" {
#endif

void* QMarginsF_NewQMarginsF();
void* QMarginsF_NewQMarginsF3(void* margins);
void* QMarginsF_NewQMarginsF2(double left, double top, double right, double bottom);
double QMarginsF_Bottom(void* ptr);
int QMarginsF_IsNull(void* ptr);
double QMarginsF_Left(void* ptr);
double QMarginsF_Right(void* ptr);
void QMarginsF_SetBottom(void* ptr, double bottom);
void QMarginsF_SetLeft(void* ptr, double left);
void QMarginsF_SetRight(void* ptr, double right);
void QMarginsF_SetTop(void* ptr, double Top);
double QMarginsF_Top(void* ptr);

#ifdef __cplusplus
}
#endif