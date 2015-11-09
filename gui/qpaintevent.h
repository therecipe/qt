#ifdef __cplusplus
extern "C" {
#endif

void* QPaintEvent_NewQPaintEvent2(void* paintRect);
void* QPaintEvent_NewQPaintEvent(void* paintRegion);
void* QPaintEvent_Region(void* ptr);

#ifdef __cplusplus
}
#endif