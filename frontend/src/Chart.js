import React from "react"
import { Line } from "react-chartjs-2"

class Chart extends React.Component {
  render() {
    const mockData = {
      labels: [ "1", "2", "1", "1", "2", "1", "1", "2", "1" ],
      datasets: [
        {
          lineTension: 0,
          label: null,
          data: [
            100,
            150,
            90,
            150,
            200
          ]
        }
      ]
    }

    return (
      <Line
        data={mockData}
        legend={{ display: false }}
        options={{
          scales: {
            gridLines: {
              display: false
            }
          }
        }}
        width={1920}
        height={300}
      />
    )
  }
}

export default Chart