import React from 'react';
import './ReportList.css';

const ReportList = ({ reports, onSelectReport }) => {
  return (
    <div className="report-list">
      <h2>Analysis Reports</h2>
      {reports.map(report => (
        <div 
          key={report.id} 
          className="report-card"
          onClick={() => onSelectReport(report)}
        >
          <h3>{report.repository}</h3>
          <p>ID: {report.id}</p>
          <p>Timestamp: {new Date(report.timestamp).toLocaleString()}</p>
          <div className="report-summary">
            <p>High: {report.summary.high}</p>
            <p>Medium: {report.summary.medium}</p>
            <p>Low: {report.summary.low}</p>
          </div>
        </div>
      ))}
    </div>
  );
};

export default ReportList;