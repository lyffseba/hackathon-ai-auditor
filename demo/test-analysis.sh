#!/bin/bash

# 🚀 AI Code Auditor Demo Script
# Tests the Cerebras-powered code analysis engine

echo "🎯 Starting AI Code Auditor Demo..."
echo "🧠 Powered by Cerebras qwen-3-235b-a22b-instruct-2507"
echo "=" "=" "=" "=" "=" "=" "=" "=" "=" "=" "=" "=" "=" "=" "=" "=" "="

# Test vulnerable JavaScript code
echo "📝 Testing with vulnerable JavaScript code..."
curl -X POST http://localhost:8080/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "code": "const express = require(\"express\");\nconst app = express();\n\napp.get(\"/user\", (req, res) => {\n  const userId = req.query.id;\n  const query = \"SELECT * FROM users WHERE id = \" + userId;\n  db.query(query, (err, results) => {\n    res.json(results);\n  });\n});\n\napp.listen(3000);",
    "language": "javascript"
  }' | jq .

echo ""
echo "🔍 Testing with insecure Go code..."

# Test vulnerable Go code
curl -X POST http://localhost:8080/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "code": "package main\n\nimport (\n\t\"fmt\"\n\t\"net/http\"\n\t\"os/exec\"\n)\n\nfunc handler(w http.ResponseWriter, r *http.Request) {\n\tcmd := r.URL.Query().Get(\"cmd\")\n\toutput, _ := exec.Command(\"sh\", \"-c\", cmd).Output()\n\tfmt.Fprintf(w, string(output))\n}\n\nfunc main() {\n\thttp.HandleFunc(\"/exec\", handler)\n\thttp.ListenAndServe(\":8080\", nil)\n}",
    "language": "go"
  }' | jq .

echo ""
echo "🐍 Testing with Python security issues..."

# Test vulnerable Python code  
curl -X POST http://localhost:8080/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "code": "import sqlite3\nimport hashlib\n\ndef login(username, password):\n    conn = sqlite3.connect(\"users.db\")\n    cursor = conn.cursor()\n    \n    # Vulnerable to SQL injection\n    query = f\"SELECT * FROM users WHERE username = '\''{username}'\'' AND password = '\''{password}'\''\"\n    cursor.execute(query)\n    \n    user = cursor.fetchone()\n    if user:\n        # Insecure password storage\n        stored_password = user[2]\n        if password == stored_password:  # Plain text comparison\n            return True\n    return False\n\n# Hardcoded credentials\nADMIN_PASSWORD = \"admin123\"\nAPI_KEY = \"sk-1234567890abcdef\"",
    "language": "python"
  }' | jq .

echo ""
echo "✅ Demo completed! Check the analysis results above."
echo "🎯 The Cerebras AI has identified security vulnerabilities, performance issues, and best practice violations!"
