#ifdef __cplusplus
extern "C" {
#endif

void* QUdpSocket_NewQUdpSocket(void* parent);
int QUdpSocket_HasPendingDatagrams(void* ptr);
int QUdpSocket_JoinMulticastGroup(void* ptr, void* groupAddress);
int QUdpSocket_JoinMulticastGroup2(void* ptr, void* groupAddress, void* iface);
int QUdpSocket_LeaveMulticastGroup(void* ptr, void* groupAddress);
int QUdpSocket_LeaveMulticastGroup2(void* ptr, void* groupAddress, void* iface);
void QUdpSocket_SetMulticastInterface(void* ptr, void* iface);
void QUdpSocket_DestroyQUdpSocket(void* ptr);

#ifdef __cplusplus
}
#endif