# CRD/API-Änderungen

Diese Schritte gelten, wenn sich das API-Modell in `api/v1` ändert, z. B.:

- neues Feld in `AuthRegistrationSpec`
- neues Feld in `AuthRegistrationStatus`
- neue Validierungs-Marker (`+kubebuilder:validation:*`)

## Ablauf

1. API-Typen anpassen in `api/v1/authregistration_types.go`.
2. DeepCopy-Code neu generieren:
   `make generate-deepcopy`
3. Typed Client neu generieren:
   `make generate-client`
4. CRD-Manifeste neu generieren:
   `make manifests`
5. Änderungen prüfen:
   `go test ./...`

## Hinweise

- `make generate-client` ist bei strukturellen API-Änderungen erforderlich.
- Bei reinen Feldänderungen ist es technisch oft nicht zwingend, wird aber empfohlen, um Drift in generiertem Code zu vermeiden.
- Die generierte CRD liegt danach unter `k8s/helm-crd/templates`.
