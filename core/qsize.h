#ifdef __cplusplus
extern "C" {
#endif

void* QSize_NewQSize();
void* QSize_NewQSize2(int width, int height);
int QSize_Height(void* ptr);
int QSize_IsEmpty(void* ptr);
int QSize_IsNull(void* ptr);
int QSize_IsValid(void* ptr);
int QSize_Rheight(void* ptr);
int QSize_Rwidth(void* ptr);
void QSize_Scale2(void* ptr, void* size, int mode);
void QSize_Scale(void* ptr, int width, int height, int mode);
void QSize_SetHeight(void* ptr, int height);
void QSize_SetWidth(void* ptr, int width);
void QSize_Transpose(void* ptr);
int QSize_Width(void* ptr);

#ifdef __cplusplus
}
#endif