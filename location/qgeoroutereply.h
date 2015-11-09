#ifdef __cplusplus
extern "C" {
#endif

void* QGeoRouteReply_NewQGeoRouteReply(int error, char* errorString, void* parent);
void QGeoRouteReply_Abort(void* ptr);
int QGeoRouteReply_Error(void* ptr);
char* QGeoRouteReply_ErrorString(void* ptr);
void QGeoRouteReply_ConnectFinished(void* ptr);
void QGeoRouteReply_DisconnectFinished(void* ptr);
int QGeoRouteReply_IsFinished(void* ptr);
void QGeoRouteReply_DestroyQGeoRouteReply(void* ptr);

#ifdef __cplusplus
}
#endif