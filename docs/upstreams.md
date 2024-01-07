This project has 3 gRPC servers that are upstream (called by) the License
worker.

## Communicating

Some quick curl commands to communicate with each of the 3 microservices.

### Org Service

```shell
pkgx http POST \
  http://localhost:9090/org.v1beta1.OrgService/CreateOrg \
    name="Org1"

pkgx http http://localhost:9090/orgs/"$(pkgx gum input --placeholder "id")"
```

### Profile Service

```shell
pkgx http POST \
  http://localhost:9091/profile.v1beta1.ProfileService/CreateProfile \
    full_name="Kevin Chen"

pkgx http http://localhost:9091/profiles/"$(pkgx gum input --placeholder "id")"
```

### License Service

```shell
pkgx http POST \
  http://localhost:9092/license.v1beta1.LicenseService/GetLicense \
    id="$(pkgx gum input --placeholder "id")"

pkgx http http://localhost:9092/licenses/"$(pkgx gum input --placeholder "id")"
```