pragma Singleton
import QtQuick 2.0
import QuickFlux 1.1

RootStore {
    /// Set the source of actions from AppDispatcher
    bindSource: AppDispatcher
}
