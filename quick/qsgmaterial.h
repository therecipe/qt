#ifdef __cplusplus
extern "C" {
#endif

int QSGMaterial_Compare(void* ptr, void* other);
void* QSGMaterial_CreateShader(void* ptr);
int QSGMaterial_Flags(void* ptr);
void QSGMaterial_SetFlag(void* ptr, int flags, int on);
void* QSGMaterial_Type(void* ptr);

#ifdef __cplusplus
}
#endif