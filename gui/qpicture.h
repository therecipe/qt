#ifdef __cplusplus
extern "C" {
#endif

int QPicture_IsNull(void* ptr);
int QPicture_Load2(void* ptr, void* dev, char* format);
int QPicture_Load(void* ptr, char* fileName, char* format);
int QPicture_Play(void* ptr, void* painter);
int QPicture_Save2(void* ptr, void* dev, char* format);
int QPicture_Save(void* ptr, char* fileName, char* format);
void QPicture_SetBoundingRect(void* ptr, void* r);
void QPicture_Swap(void* ptr, void* other);
void QPicture_DestroyQPicture(void* ptr);

#ifdef __cplusplus
}
#endif