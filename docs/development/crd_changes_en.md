# CRD/API Changes

These steps apply whenever the API model in `api/v1` changes, for example:

- new field in `AuthRegistrationSpec`
- new field in `AuthRegistrationStatus`
- new validation markers (`+kubebuilder:validation:*`)

## Workflow

1. Update API types in `api/v1/authregistration_types.go`.
2. Regenerate DeepCopy code:
   `make generate-deepcopy`
3. Regenerate typed client:
   `make generate-client`
4. Regenerate CRD manifests:
   `make manifests`
5. Verify changes:
   `go test ./...`

## Notes

- `make generate-client` is required for structural API changes.
- For pure field changes, this is often not technically required, but still recommended to avoid drift in generated code.
- Generated CRD manifests are written to `k8s/helm-crd/templates`.
