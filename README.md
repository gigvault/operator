# GigVault Operator

Kubernetes operator for managing GigVault Certificate Authority lifecycle.

## Features

- Automated CA deployment and configuration
- Certificate lifecycle management
- Automatic rotation and renewal
- Custom Resource Definitions (CRDs)

## CRDs

### CertificateAuthority

```yaml
apiVersion: gigvault.io/v1alpha1
kind: CertificateAuthority
metadata:
  name: intermediate-ca
spec:
  type: intermediate
  validity: 10y
  keyAlgorithm: ECDSA-P384
  parent:
    secretRef:
      name: root-ca-secret
```

### Certificate

```yaml
apiVersion: gigvault.io/v1alpha1
kind: Certificate
metadata:
  name: example-cert
spec:
  commonName: example.com
  dnsNames:
    - example.com
    - www.example.com
  validity: 90d
  issuerRef:
    name: intermediate-ca
```

## Installation

```bash
# Install CRDs
kubectl apply -f manifests/crds/

# Deploy operator
kubectl apply -f manifests/operator.yaml

# Or use Helm
helm install gigvault-operator charts/operator
```

## Development

```bash
# Build
make build

# Run locally
make run

# Deploy to cluster
make deploy
```

## License

Copyright © 2025 GigVault

