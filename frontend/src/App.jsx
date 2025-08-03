import React, { useState, useEffect } from 'react'
import './App.css'
import ReportList from './components/ReportList'
import ReportDetail from './components/ReportDetail'

function App() {
  const [reports, setReports] = useState([])
  const [selectedReport, setSelectedReport] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    // Fetch reports from the backend
    fetch('/api/reports')
      .then(response => {
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`)
        }
        return response.json()
      })
      .then(data => {
        setReports(data.reports)
        setLoading(false)
      })
      .catch(err => {
        setError(err.message)
        setLoading(false)
      })
  }, [])

  const handleReportSelect = (report) => {
    setSelectedReport(report)
  }

  const handleBackToList = () => {
    setSelectedReport(null)
  }

  if (loading) {
    return <div className="app">Loading...</div>
  }

  if (error) {
    return <div className="app">Error: {error}</div>
  }

  return (
    <div className="app">
      <header className="app-header">
        <h1>AI Code Auditor Dashboard</h1>
      </header>
      <main className="app-main">
        {selectedReport ? (
          <ReportDetail report={selectedReport} onBack={handleBackToList} />
        ) : (
          <ReportList reports={reports} onSelectReport={handleReportSelect} />
        )}
      </main>
    </div>
  )
}

export default App