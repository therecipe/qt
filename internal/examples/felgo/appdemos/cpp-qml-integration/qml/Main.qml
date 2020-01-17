import Felgo 3.0
import QtQuick 2.5

// custom import to use MyQMLType, a C++ based QML type
// NOTE: the import identifier, version and QML type name are set in main.cpp at qmlRegisterType(...)
import com.yourcompany.xyz 1.0

App {
  id: app

  NavigationStack {
    Page {
      title: "Integrate C++ and QML"

      // Example 1 - Global Context Property
      // NOTE: myGlobalObject is available here because it is set as a context property in main.cpp
      Column {

        // 1.1: Calling myGlobalObject.doSomething() function
        AppButton {
          text: "myGlobalObject.doSomething()"
          onClicked: myGlobalObject.doSomething("TEXT FROM QML")
        }

        // 1.2: Increasing myGlobalObject.counter property
        // NOTE: the defined setter function of the property is used automatically and triggers the counterChanged signal
        AppButton {
          text: "myGlobalObject.counter + 1"
          onClicked: {
            myGlobalObject.counter = myGlobalObject.counter + 1
          }
        }

        // 1.3: Showing myGlobalObject counter value in a QML text
        // NOTE: property bindings are supported, as the counter property definition includes the counterChanged signal,
        // which is fired in the implementation of MyGlobalObject::setCounter() for each property change
        AppText {
          text: "Global Context Property Counter: "+myGlobalObject.counter
        }
      // Example 1 ends here...


      // Example 2: Custom QML Type implemented with C++
      // NOTE: This type is declared in main.cpp and available after using "import com.yourcompany.xyz 1.0"
      // To create a type that also has a visual representation and may contain child items, derive from QQuickItem instead of QObject
      MyQMLType {
        id: typeFromCpp

        // 2.1: Property Binding for MyQMLType::message property
        // NOTE: Similar to types created purely with QML, you may use property bindings to keep your property values updated
        message: "counter / 2 = " + Math.floor(myGlobalObject.counter / 2)

        // 2.2: Reacting to property changes
        // NOTE: With the onMessageChanged signal, you can add code to handle property changes
        onMessageChanged: console.log("typeFromCpp message changed to '"+typeFromCpp.message+"'")

        // 2.3: Run code at creation of the QML component
        // NOTE: The Component.onCompleted signal is available for every QML item, even for items defined with C++.
        // The signal is fired when the QML Engine creates the item at runtime.
        Component.onCompleted: myGlobalObject.counter = typeFromCpp.increment(myGlobalObject.counter)

        // 2.4: Handling a custom signal
        onCppTaskFinished: {
          myGlobalObject.counter = 0 // reset counter to zero, this will also update the message
        }
      }

      // 2.1: Show typeFromCpp.message value, which is calculated automatically based on the myGlobalObject.counter value
      AppText {
        text: "Custom QML Type Message:\n" + typeFromCpp.message
      }

      // 2.4: Button to start cpp task
      AppButton {
        text: "typeFromCpp.startCppTask()"
        onClicked: {
            typeFromCpp.startCppTask()
        }
      }
    }

      // 2.5: Connections allow to add signal handlers for global context property objects
      Connections {
          target: myGlobalObject
          onCounterChanged: console.log("Counter changed to "+myGlobalObject.counter)
      }
    }

  }
}

