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

## 📜 Documentation

### ✅ What's Done

- **Cerebras Integration**: Replaced OpenAI with Cerebras for enhanced code analysis.
- **Security Analysis**: Identifies security vulnerabilities, including SQL injection, XSS.
- **Performance Optimization**: Detects memory leaks, inefficient algorithms.
- **Code Quality Checks**: Highlights code smells, best practices, and maintainability issues.
- **GitHub Webhook Support**: Seamless integration with CI/CD pipelines.
- **Demo Scripts**: Ready-to-run scripts to showcase the tool's capabilities.

### 🚀 What's Next

- **Enhance Frontend**: Implement a responsive React dashboard for better visualization.
- **Extend Language Support**: Add support for more programming languages.
- **Advanced Configurations**: Allow user-defined rules and custom setups.
- **Real-time Collaboration**: Live analysis while coding in shared environments.

## 🚀 Getting Started

1. **Environment Setup**

   ```bash
   cp backend/.env.example backend/.env
   # Fill in your credentials and configuration

   # Ensure Cerebras API is configured correctly
   export CEREBRAS_API_KEY="your_cerebras_api_key"
   export CEREBRAS_API_URL="https://api.cerebras.ai/v1/chat/completions"
   ```

2. **Start the Backend**

   ```bash
   cd backend 
   go run main.go
   ```

3. **Run the Demo**

   ```bash
   bash demo/test-analysis.sh
   ```

---
**Built with ❤️ for the Cerebras Hackathon**
