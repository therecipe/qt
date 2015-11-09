#ifdef __cplusplus
extern "C" {
#endif

void* QPlatformSurfaceEvent_NewQPlatformSurfaceEvent(int surfaceEventType);
int QPlatformSurfaceEvent_SurfaceEventType(void* ptr);

#ifdef __cplusplus
}
#endif