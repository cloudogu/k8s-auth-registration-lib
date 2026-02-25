# Development Guide

## API/CRD Development

For changes in `api/v1/*` (for example a new field in `Spec` or `Status`), follow this workflow:

- [CRD/API Changes](./crd_changes_en.md)

## Important Directories

- API types: `api/v1`
- Generated client: `client`
- Generated CRD manifests: `k8s/helm-crd/templates`
- Development documentation: `docs/development`
