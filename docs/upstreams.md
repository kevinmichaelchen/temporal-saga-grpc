This project has 3 gRPC servers that are upstream (called by) the License
worker.

## Communicating

Some quick curl commands to communicate with each of the 3 microservices.

### License Service

```shell
curl -v http://localhost:9090/license.v1beta1.LicenseService/CreateLicense \
  -H "Content-Type: application/json" \
  -d '{"name": "L1"}'

http POST \
  http://localhost:9090/license.v1beta1.LicenseService/CreateLicense \
    name="L1"
```

### Org Service

```shell
curl -v http://localhost:9091/org.v1beta1.OrgService/CreateOrg \
  -H "Content-Type: application/json" \
  -d '{"name": "Org1"}'

http POST \
  http://localhost:9091/org.v1beta1.OrgService/CreateOrg \
    name="Org1"
```

### Profile Service

```shell
curl -v http://localhost:9092/profile.v1beta1.ProfileService/CreateProfile \
  -H "Content-Type: application/json" \
  -d '{"name": "Kevin Chen"}'

http POST \
  http://localhost:9092/profile.v1beta1.ProfileService/CreateProfile \
    name="Kevin Chen"
```
