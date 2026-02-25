# AuthRegistration verwenden

Diese Seite beschreibt, wie eine `AuthRegistration`-Resource verwendet wird.

## Beispiel-CR

Beispiel ohne `secretRef`: Der Controller soll selbst ein Secret erzeugen und verwenden.

```yaml
apiVersion: k8s.cloudogu.com/v1
kind: AuthRegistration
metadata:
  name: bluespice-auth
  namespace: ecosystem
spec:
  protocol: OIDC
  consumer: bluespice
  logoutURL: https://bluespice.example.com/logout
  params:
    scope: "openid profile email"
```

Optional kann `spec.secretRef` gesetzt werden. Dann soll der Controller dieses Secret verwenden:

```yaml
spec:
  secretRef: bluespice-auth-credentials
```

## Anwendung

1. Resource anlegen:
```bash
kubectl apply -f authregistration.yaml
```

2. Status beobachten:
```bash
kubectl -n ecosystem get ar bluespice-auth
kubectl -n ecosystem get ar bluespice-auth -o jsonpath='{.status.resolvedSecretRef}{"\n"}'
kubectl -n ecosystem get ar bluespice-auth -o jsonpath='{range .status.conditions[*]}{.type}={.status}{" ("}{.reason}{")\n"}{end}'
```

3. Verwendetes Secret pruefen:
```bash
SECRET_NAME=$(kubectl -n ecosystem get ar bluespice-auth -o jsonpath='{.status.resolvedSecretRef}')
kubectl -n ecosystem get secret "${SECRET_NAME}" -o yaml
```

## Spec-Felder

| Feld             | Typ                 | Pflicht | Beschreibung                                                                                                             |
|------------------|---------------------|---------|--------------------------------------------------------------------------------------------------------------------------|
| `spec.protocol`  | `string`            | Ja      | Authentifizierungsprotokoll. Erlaubte Werte: `CAS`, `OIDC`, `OAUTH`.                                                     |
| `spec.consumer`  | `string`            | Ja      | Name/Identifier des konsumierenden Dienstes.                                                                             |
| `spec.secretRef` | `string`            | Nein    | Name eines bestehenden Ziel-Secrets fuer die Credentials. Wenn nicht gesetzt, erzeugt der Controller ein eigenes Secret. |
| `spec.logoutURL` | `string` (URI)      | Nein    | Optionale Logout-URL fuer Single-Logout-Integrationen.                                                                   |
| `spec.params`    | `map[string]string` | Nein    | Optionale protokollspezifische Zusatzparameter.                                                                          |

## Status-Felder

| Feld                       | Typ                  | Beschreibung                                                                                          |
|----------------------------|----------------------|-------------------------------------------------------------------------------------------------------|
| `status.resolvedSecretRef` | `string`             | Effektiv verwendetes Secret fuer die Credentials (entweder `spec.secretRef` oder generiertes Secret). |
| `status.conditions`        | `[]metav1.Condition` | Zustandsliste der Ressource. Relevante Condition-Typen sind `Completed` und `CredentialsPublished`.   |

## Conditions

`status.conditions[*].status` verwendet die ueblichen Werte `True`, `False`, `Unknown`.

| Condition-Typ          | Bedeutung                                                                            |
|------------------------|--------------------------------------------------------------------------------------|
| `Completed`            | Der gesamte Registrierungsprozess ist abgeschlossen.                                 |
| `CredentialsPublished` | Credentials wurden in das effektive Secret geschrieben (`status.resolvedSecretRef`). |
