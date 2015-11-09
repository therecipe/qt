#ifdef __cplusplus
extern "C" {
#endif

void* QSizeF_NewQSizeF();
void* QSizeF_NewQSizeF2(void* size);
void* QSizeF_NewQSizeF3(double width, double height);
double QSizeF_Height(void* ptr);
int QSizeF_IsEmpty(void* ptr);
int QSizeF_IsNull(void* ptr);
int QSizeF_IsValid(void* ptr);
double QSizeF_Rheight(void* ptr);
double QSizeF_Rwidth(void* ptr);
void QSizeF_Scale2(void* ptr, void* size, int mode);
void QSizeF_Scale(void* ptr, double width, double height, int mode);
void QSizeF_SetHeight(void* ptr, double height);
void QSizeF_SetWidth(void* ptr, double width);
void QSizeF_Transpose(void* ptr);
double QSizeF_Width(void* ptr);

#ifdef __cplusplus
}
#endif