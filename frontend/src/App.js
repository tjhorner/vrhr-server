import React from 'react'
import { BrowserRouter as Router, Switch, Route, Redirect } from "react-router-dom"
import Status from "./Status"
import Chart from "./Chart"

class App extends React.Component {
  constructor() {
    super()
  }

  render() {
    return (
      <Router>
        <Switch>
          <Route path="/status">
            <Status />
          </Route>

          <Route path="/chart">
            <Chart />
          </Route>

          <Route path="/">
            <Redirect to="/status" />
          </Route>
        </Switch>
      </Router>
    )
  }
}

export default App
