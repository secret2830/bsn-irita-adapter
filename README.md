# BSN-IRITA External Adapter

## Get Started

### Install

```bash
make install
```

### Configuration

#### Environment variables

| Key | Description |
|-----|-------------|
| `BA_KEY_MNEMONIC` | BSN-IRITA key mnemonic |
| `BA_CHAIN_ID` | BSN-IRITA endpoint chain id |
| `BA_ENDPOINT_RPC` | Endpoint RPC address for the BSN-IRITA node to connect to |
| `BA_ENDPOINT_GRPC` | Endpoint gRPC address for the BSN-IRITA node to connect to |
| `BA_LISTEN_ADDR` | BSN-IRITA adapter listen address, default to `0.0.0.0:8080` |

### Start

```bash
bsn-irita-adapter
```
