# sekrets
Kubernetes secrets yaml generator command line tool

## Install

Use `go get` to install: `go get github.com/g4s8/sekrets`

## Example

```sh
# sekrets -o=test.yaml -name=secretname       
Enter key (empty to complete): one
Enter value: qwerty1234
Enter key (empty to complete): 
Done: test.yaml
```
generates this file:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: secretname
  namespace: default
Type: Opaque
data:
  one: cXdlcnR5MTIzNA==
```
