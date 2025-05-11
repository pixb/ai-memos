# Protocol Buffers & Buf

This directory contains the Protocol Buffer (protobuf) definitions for the Memos project.
We use [Buf](https://buf.build) to manage and lint our protobuf files, as well as for code generation.

## Installation

To work with the protobuf definitions and generate code, you need to install the `buf` CLI.

### macOS (Homebrew)

```bash
brew install bufbuild/buf/buf
```

### Linux / macOS (Manual)

```bash
VERSION="1.33.0" # Check for the latest version on Buf's GitHub releases
sudo curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "/usr/local/bin/buf" && \
sudo chmod +x "/usr/local/bin/buf"
```

### Windows (Manual)

Download the latest `buf.exe` binary from the [Buf GitHub releases page](https://github.com/bufbuild/buf/releases) and add it to your `PATH`.

For more installation options, refer to the [official Buf installation guide](https://buf.build/docs/installation).

## Usage

### Initialization

To set up Buf for your project, run the following command from the root of your repository:

```bash
buf mod init
```

This will create a `buf.yaml` file in the root of your repository, which you can customize to your liking.

### Linting

To check your protobuf files for style and consistency issues:

```bash
buf lint
```

This command should be run from the root of the repository or a directory containing a `buf.yaml` configuration file (if you set one up).

### Code Generation

To generate code from your protobuf definitions:

```bash
buf generate
```

This command relies on a `buf.gen.yaml` file that defines the generation templates and output locations. You'll need to create this file based on your project's needs (e.g., Go, gRPC, etc.).

Example `buf.gen.yaml` for Go:

```yaml
version: v1
plugins:
  - plugin: buf.build/protocolbuffers/go:v1.31.0
    out: gen/go
    opt: paths=source_relative
  - plugin: buf.build/grpc/go:v1.3.0
    out: gen/go
    opt: paths=source_relative,require_unimplemented_servers=false
```

Make sure to install the necessary plugins (e.g., `protoc-gen-go`, `protoc-gen-go-grpc`).

### Formatting

Buf can also format your `.proto` files. While `buf lint` often catches formatting issues, you might want to integrate a formatter directly. Buf itself doesn't have a standalone format command in the same way as `gofmt` or `prettier`. Instead, linting rules (like those in `buf.yaml` with `use: FILE`) enforce formatting.

For more advanced formatting or specific editor integrations, you might look into tools like `clang-format` with a `.clang-format` file configured for protobuf, or editor-specific protobuf plugins.

However, the primary way to ensure consistent formatting with Buf is through its linting capabilities.

## Buf Configuration

You can configure Buf's behavior using a `buf.yaml` file at the root of your protobuf sources (e.g., in this `proto` directory).

Example `buf.yaml`:

```yaml
version: v1
name: buf.build/your-org/your-repo # Replace with your Buf Schema Registry path if you use it
lint:
  use:
    - DEFAULT
breaking:
  use:
    - FILE
```

Refer to the [Buf documentation](https://buf.build/docs/introduction) for more details on configuration and advanced usage.


## Format

```sh
buf format -w
```