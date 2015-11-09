#ifdef __cplusplus
extern "C" {
#endif

void QRubberBand_SetGeometry(void* ptr, void* rect);
void* QRubberBand_NewQRubberBand(int s, void* p);
void QRubberBand_Move2(void* ptr, void* p);
void QRubberBand_Move(void* ptr, int x, int y);
void QRubberBand_Resize2(void* ptr, void* size);
void QRubberBand_Resize(void* ptr, int width, int height);
void QRubberBand_SetGeometry2(void* ptr, int x, int y, int width, int height);
int QRubberBand_Shape(void* ptr);
void QRubberBand_DestroyQRubberBand(void* ptr);

#ifdef __cplusplus
}
#endif