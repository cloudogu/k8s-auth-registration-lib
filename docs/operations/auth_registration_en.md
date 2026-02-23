# Using AuthRegistration

This page describes how to use an `AuthRegistration` resource.

## Example CR

Example without `secretRef`: the controller should create and manage a Secret automatically.

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

Optionally, `spec.secretRef` can be set. In this case, the controller should use that Secret:

```yaml
spec:
  secretRef: bluespice-auth-credentials
```

## Usage

1. Create the resource:
```bash
kubectl apply -f authregistration.yaml
```

2. Observe status:
```bash
kubectl -n ecosystem get ar bluespice-auth
kubectl -n ecosystem get ar bluespice-auth -o jsonpath='{.status.resolvedSecretRef}{"\n"}'
kubectl -n ecosystem get ar bluespice-auth -o jsonpath='{range .status.conditions[*]}{.type}={.status}{" ("}{.reason}{")\n"}{end}'
```

3. Inspect the effective Secret:
```bash
SECRET_NAME=$(kubectl -n ecosystem get ar bluespice-auth -o jsonpath='{.status.resolvedSecretRef}')
kubectl -n ecosystem get secret "${SECRET_NAME}" -o yaml
```

## Spec Fields

| Field            | Type                | Required | Description                                                                                               |
|------------------|---------------------|----------|-----------------------------------------------------------------------------------------------------------|
| `spec.protocol`  | `string`            | Yes      | Authentication protocol. Allowed values: `CAS`, `OIDC`, `OAUTH`.                                          |
| `spec.consumer`  | `string`            | Yes      | Name/identifier of the consuming service.                                                                 |
| `spec.secretRef` | `string`            | No       | Name of an existing target Secret for credentials. If omitted, the controller creates a dedicated Secret. |
| `spec.logoutURL` | `string` (URI)      | No       | Optional logout URL for single logout integrations.                                                       |
| `spec.params`    | `map[string]string` | No       | Optional protocol-specific additional parameters.                                                         |

## Status Fields

| Field                      | Type                 | Description                                                                                     |
|----------------------------|----------------------|-------------------------------------------------------------------------------------------------|
| `status.resolvedSecretRef` | `string`             | Effective Secret used for credentials (either `spec.secretRef` or a generated Secret).          |
| `status.conditions`        | `[]metav1.Condition` | Resource state conditions. Relevant condition types are `Completed` and `CredentialsPublished`. |

## Conditions

`status.conditions[*].status` uses the standard values `True`, `False`, `Unknown`.

| Condition Type         | Meaning                                                                        |
|------------------------|--------------------------------------------------------------------------------|
| `Completed`            | The overall registration process has finished.                                 |
| `CredentialsPublished` | Credentials were written to the effective Secret (`status.resolvedSecretRef`). |
