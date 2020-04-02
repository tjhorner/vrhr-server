import React from 'react'
import './App.css'
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faHeartbeat } from "@fortawesome/free-solid-svg-icons"
import Bars from "./Bars"

class App extends React.Component {
  constructor() {
    super()
    this.state = {
      heartRate: 0,
      accuracy: -1
    }

    let refreshInterval = setInterval(function() {
      this.refresh()
    }.bind(this), 1000)

    this.setState({ refreshInterval })
  }

  async refresh() {
    const resp = await fetch(`${process.env.REACT_APP_API_BASE_URL || ""}/api/state`)
    const data = await resp.json()
    this.setState({
      heartRate: data.current_heart_rate,
      accuracy: data.current_accuracy
    })
  }

  componentWillUnmount() {
    if(this.state.refreshInterval)
      clearInterval(this.state.refreshInterval)
  }

  render() {
    const accuracyMap = {
      [-1]: "No Data",
      0: "Unreliable",
      1: "Low",
      2: "Medium",
      3: "High"
    }

    return (
      <div className="app">
        <div className="left-side">
          <FontAwesomeIcon icon={faHeartbeat} className="heart-icon" />
        </div>
        <div className="right-side">
          <div className="accuracy">
            <Bars value={this.state.accuracy} />
            <span className="acc-text">Accuracy: {accuracyMap[this.state.accuracy]}</span>
          </div>
          <div className="hr">
            <span className="hr-number">{this.state.heartRate}</span>
            BPM
          </div>
        </div>
      </div>
    )
  }
}

export default App
