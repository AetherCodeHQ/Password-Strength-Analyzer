# Password Strength Analyzer

![CI](https://github.com/Qyroxen/Password-Strength-Analyzer/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Password-Strength-Analyzer/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Password-Strength-Analyzer?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Password-Strength-Analyzer)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Password-Strength-Analyzer)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Password-Strength-Analyzer?style=social)](https://github.com/Qyroxen/Password-Strength-Analyzer/stargazers)

## What is it?

Password Strength Analyzer is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Password-Strength-Analyzer.git
cd Password-Strength-Analyzer
go build -o passwordstrengthanalyzer .

# Run
./passwordstrengthanalyzer --help
```

## CLI Usage

```bash
# Basic usage
./passwordstrengthanalyzer

# With flags
./passwordstrengthanalyzer --verbose --output json

# Get help
./passwordstrengthanalyzer --help
```

## Examples

```bash
# Example 1
./passwordstrengthanalyzer example1

# Example 2
./passwordstrengthanalyzer example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o passwordstrengthanalyzer .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Password-Strength-Analyzer/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Password-Strength-Analyzer?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Password-Strength-Analyzer/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Password-Strength-Analyzer?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Password-Strength-Analyzer/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Password-Strength-Analyzer" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Password-Strength-Analyzer/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Password-Strength-Analyzer" alt="Pull Requests">
  </a>
</p>
