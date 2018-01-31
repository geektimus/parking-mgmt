import React, { Component } from "react";

class Users extends Component {
  constructor(props) {
    super(props);
    this.state = { users: [] };
  }

  render() {
    console.log("Component Rendered");
    let userLIs = this.state.users.map(user => (
      <li id={user.ID} key={user.ID}>
        {user.firstname} {user.lastname}
      </li>
    ));
    return <ul>{userLIs}</ul>;
  }

  componentDidMount() {
    fetch(`http://parking-mgmt-react:9000/api/users`)
      .then(result => result.json())
      .then(users => this.setState({ users }));
  }
}

export default Users;
