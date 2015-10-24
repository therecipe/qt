#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QUdpSocket_NewQUdpSocket(QtObjectPtr parent);
int QUdpSocket_HasPendingDatagrams(QtObjectPtr ptr);
int QUdpSocket_JoinMulticastGroup(QtObjectPtr ptr, QtObjectPtr groupAddress);
int QUdpSocket_JoinMulticastGroup2(QtObjectPtr ptr, QtObjectPtr groupAddress, QtObjectPtr iface);
int QUdpSocket_LeaveMulticastGroup(QtObjectPtr ptr, QtObjectPtr groupAddress);
int QUdpSocket_LeaveMulticastGroup2(QtObjectPtr ptr, QtObjectPtr groupAddress, QtObjectPtr iface);
void QUdpSocket_SetMulticastInterface(QtObjectPtr ptr, QtObjectPtr iface);
void QUdpSocket_DestroyQUdpSocket(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif