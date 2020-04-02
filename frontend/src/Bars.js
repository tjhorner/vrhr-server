import React from "react"

function Bars(props = { value: 0 }) {
  const color = {
    1: "red",
    2: "yellow",
    3: "green"
  }

  return (
    <div className={`bars-root ${color[props.value]}`}>
      <div className={`bar ${props.value >= 1 ? "active": ""}`}></div>
      <div className={`bar ${props.value >= 2 ? "active": ""}`}></div>
      <div className={`bar ${props.value >= 3 ? "active": ""}`}></div>
    </div>
  )
}

export default Bars