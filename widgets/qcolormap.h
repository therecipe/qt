#ifdef __cplusplus
extern "C" {
#endif

void* QColormap_NewQColormap(void* colormap);
int QColormap_Depth(void* ptr);
int QColormap_Mode(void* ptr);
int QColormap_Size(void* ptr);
void QColormap_DestroyQColormap(void* ptr);

#ifdef __cplusplus
}
#endif