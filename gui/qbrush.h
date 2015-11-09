#ifdef __cplusplus
extern "C" {
#endif

void* QBrush_NewQBrush4(int color, int style);
void QBrush_SetColor(void* ptr, void* color);
void* QBrush_NewQBrush();
void* QBrush_NewQBrush2(int style);
void* QBrush_NewQBrush6(int color, void* pixmap);
void* QBrush_NewQBrush9(void* other);
void* QBrush_NewQBrush3(void* color, int style);
void* QBrush_NewQBrush5(void* color, void* pixmap);
void* QBrush_NewQBrush10(void* gradient);
void* QBrush_NewQBrush8(void* image);
void* QBrush_NewQBrush7(void* pixmap);
void* QBrush_Color(void* ptr);
void* QBrush_Gradient(void* ptr);
int QBrush_IsOpaque(void* ptr);
void QBrush_SetColor2(void* ptr, int color);
void QBrush_SetStyle(void* ptr, int style);
void QBrush_SetTexture(void* ptr, void* pixmap);
void QBrush_SetTextureImage(void* ptr, void* image);
void QBrush_SetTransform(void* ptr, void* matrix);
int QBrush_Style(void* ptr);
void QBrush_Swap(void* ptr, void* other);
void QBrush_DestroyQBrush(void* ptr);

#ifdef __cplusplus
}
#endif