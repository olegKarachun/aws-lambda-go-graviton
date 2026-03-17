# AWS Lambda Performance Battle: Go vs .NET (Part 1 - Go)

This repository contains the source code for the first part of my benchmarking series: **"Mastering Go on AWS Lambda: The Graviton & AL2023 Way"**.

## Project Goal

To demonstrate how to build a highly optimized Go function for AWS Lambda using the latest `provided.al2023` runtime and ARM64 (Graviton) architecture.

## Optimization Techniques Used

- **Custom Runtime (AL2023):** Using the most minimalist Amazon Linux OS.
- **ARM64 Architecture:** Leveraging Graviton processors for cost-efficiency.
- **Selective Tags:** Using `-tags lambda.norpc` to strip legacy code.
- **Binary Stripping:** Using linker flags (`-s -w`) to minimize binary size.

## Quick Start

### Prerequisites

- Go 1.21+
- AWS SAM CLI
- Docker (for local testing)

### Build

To compile the highly optimized ARM64 binary:

```bash
make build
```
