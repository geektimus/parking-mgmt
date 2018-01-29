import React, { Component } from "react";

class Users extends Component {
  render() {
    let userLIs = this.props.userList.map(user => (
      <li id={user.id} key={user.id}>
        {user.firstname} {user.lastname}
      </li>
    ));
    return <ul>{userLIs}</ul>;
  }
}

export default Users;
