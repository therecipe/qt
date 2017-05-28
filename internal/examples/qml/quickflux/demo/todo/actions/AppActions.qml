pragma Singleton
import QtQuick 2.0
import QuickFlux 1.1
import "./"

ActionCreator {

    // Add a new task
    signal addTask(string title);

    // Set/unset done on a task
    signal setTaskDone(var uid, bool done)

    // Show/hide completed task
    signal setShowCompletedTasks(bool value)

}

