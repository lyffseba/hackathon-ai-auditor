import React from 'react';
import './ReportDetail.css';

const ReportDetail = ({ report, onBack }) => {
  if (!report) return null;

  return (
    <div className="report-detail">
      <button className="back-button" onClick={onBack}>
        Back to Reports
      </button>
      
      <h2>Report Details</h2>
      <div className="report-header">
        <p><strong>ID:</strong> {report.id}</p>
        <p><strong>Repository:</strong> {report.repository}</p>
        <p><strong>Timestamp:</strong> {new Date(report.timestamp).toLocaleString()}</p>
      </div>
      
      <div className="report-summary">
        <h3>Summary</h3>
        <p>High Severity Issues: {report.summary.high}</p>
        <p>Medium Severity Issues: {report.summary.medium}</p>
        <p>Low Severity Issues: {report.summary.low}</p>
      </div>
      
      <div className="findings-list">
        <h3>Findings</h3>
        {report.findings.map((finding, index) => (
          <div key={index} className={`finding-item ${finding.severity.toLowerCase()}`}>
            <p><strong>Type:</strong> {finding.type}</p>
            <p><strong>File:</strong> {finding.file}:{finding.line}</p>
            <p><strong>Message:</strong> {finding.message}</p>
            <pre><code>{finding.code}</code></pre>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ReportDetail;