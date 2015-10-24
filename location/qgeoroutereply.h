#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoRouteReply_NewQGeoRouteReply(int error, char* errorString, QtObjectPtr parent);
void QGeoRouteReply_Abort(QtObjectPtr ptr);
int QGeoRouteReply_Error(QtObjectPtr ptr);
char* QGeoRouteReply_ErrorString(QtObjectPtr ptr);
void QGeoRouteReply_ConnectFinished(QtObjectPtr ptr);
void QGeoRouteReply_DisconnectFinished(QtObjectPtr ptr);
int QGeoRouteReply_IsFinished(QtObjectPtr ptr);
void QGeoRouteReply_DestroyQGeoRouteReply(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif