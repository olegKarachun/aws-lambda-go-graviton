# AWS Lambda Go on ARM64 (Graviton)

[![Go Version](https://img.shields.io/github/go-mod/go-version/olegKarachun/aws-lambda-go-graviton?filename=src%2Fgo-lambda%2Fgo.mod)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This repository contains a highly optimized Go function for AWS Lambda, running on **ARM64 (Graviton)** architecture with the **Amazon Linux 2023 (AL2023)** runtime.

This is **Part 1** of my series: _"Battle of the Titans: Go vs .NET Native AOT"_.

## 🚀 Optimization Highlights

To achieve the best possible performance and the lowest "Cold Start" times, this project implements:

- **Custom Runtime (`provided.al2023`):** Using the most minimalist OS available in AWS for faster boot times.
- **ARM64 (Graviton):** Specifically compiled for Graviton2/3 processors to get ~20% better price/performance.
- **Selective Compilation (`-tags lambda.norpc`):** Strips away legacy RPC code, reducing the binary size.
- **Linker Stripping (`-ldflags="-s -w"`):** Removes debug information and symbol tables for a leaner binary.

## 🛠 Project Structure

- `main.go`: Business logic (JSON parsing, SHA-256 hashing, filtering).
- `template.yaml`: AWS SAM template defining the Infrastructure as Code (IaC).
- `Makefile`: Automated build and deploy commands.

## 🏁 Quick Start

### Prerequisites

- Go 1.21+
- AWS SAM CLI
- Make (optional, but recommended)

### Build & Deploy

1. **Compile the optimized binary:**
   ```bash
   make build
   ```
