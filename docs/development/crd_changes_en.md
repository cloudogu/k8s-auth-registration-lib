# CRD/API Changes

These steps apply whenever the API model in `api/v1` changes, for example:

- new field in `AuthRegistrationSpec`
- new field in `AuthRegistrationStatus`
- new validation markers (`+kubebuilder:validation:*`)

## Workflow

1. Update API types in `api/v1/authregistration_types.go`.
2. Regenerate generated code and CRD manifests:
   `make generate-crd-api`
3. Verify changes:
   `go test ./...`

## Notes

- `make generate-crd-api` runs `make generate-deepcopy`, `make generate-client`, and `make manifests` in sequence.
- `make generate-client` is required for structural API changes.
- For pure field changes, this is often not technically required, but still recommended to avoid drift in generated code.
- Generated CRD manifests are written to `k8s/helm-crd/templates`.
