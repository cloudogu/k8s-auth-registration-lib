# Leitfaden zur Entwicklung

## Entwicklung am API/CRD-Modell

Bei Änderungen an `api/v1/*` (z. B. neues Feld in `Spec` oder `Status`) den beschriebenen Ablauf ausführen:

- `docs/development/crd_changes_de.md`

## Wichtige Verzeichnisse

- API-Typen: `api/v1`
- Generierter Client: `client`
- Generierte CRD-Manifeste: `k8s/helm-crd/templates`
- Entwicklungsdokumentation: `docs/development`
