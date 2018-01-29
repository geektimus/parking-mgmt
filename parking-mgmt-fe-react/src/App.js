import React, { Component } from "react";
import "./App.css";

import Users from "./components/Users";

const users = [
  { id: "1", firstname: "John", lastname: "Doe" },
  { id: "2", firstname: "Jane", lastname: "Doe" }
];

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">Parking Management</h1>
        </header>
        <Users userList={users} />
      </div>
    );
  }
}

export default App;
