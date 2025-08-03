# 🔍 AI Code Auditor - Hackathon Project

**Intelligent Code Review & Security Analysis powered by Cerebras AI**

## 🏆 What We're Building

An AI-powered tool that automatically analyzes codebases for:
- 🛡️ Security vulnerabilities
- 🐛 Code quality issues  
- ⚡ Performance bottlenecks
- 📚 Documentation gaps
- 🎯 Best practice violations

## ✨ Key Features

- **Real-time Analysis**: Instant feedback on code commits
- **Multi-language Support**: Go, JavaScript, Python, Java, and more
- **Security Focus**: OWASP Top 10, injection attacks, auth issues
- **Performance Insights**: Memory leaks, inefficient algorithms
- **GitHub Integration**: Seamless PR reviews
- **Beautiful Reports**: HTML/JSON output with actionable insights

## 🚀 Tech Stack

- **AI Engine**: Cerebras Llama models for code analysis
- **Backend**: Go with Gin framework
- **Frontend**: React with modern UI
- **Analysis**: AST parsing + pattern matching
- **Integration**: GitHub API, Git hooks

## 🎯 Demo Scenarios

1. **Security Audit**: Scan popular repos, find real vulnerabilities
2. **Live Analysis**: Real-time feedback as you type
3. **Batch Processing**: Analyze entire organizations
4. **CI/CD Integration**: Automated PR reviews

## 🏗️ Project Structure

```
hackathon-ai-auditor/
├── backend/           # Go API server
├── frontend/          # React dashboard
├── analyzer/          # Core analysis engine
├── integrations/      # GitHub, Git integrations
├── demo/             # Demo scripts and data
└── docs/             # Documentation
```

## 🚀 Getting Started

```bash
# Start the backend
cd backend && go run main.go

# Start the frontend  
cd frontend && npm start

# Run analysis
./scripts/analyze-repo.sh https://github.com/user/repo
```

---
**Built with ❤️ for the Cerebras Hackathon**
