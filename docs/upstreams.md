This project has 3 gRPC servers that are upstream (called by) the License
worker.

## Communicating

Some quick curl commands to communicate with each of the 3 microservices.

### Org Service

```shell
curl -v http://localhost:9090/org.v1beta1.OrgService/CreateOrg \
  -H "Content-Type: application/json" \
  -d '{"name": "Org1"}'

http POST \
  http://localhost:9090/org.v1beta1.OrgService/CreateOrg \
    name="Org1"
```

### Profile Service

```shell
curl -v http://localhost:9091/profile.v1beta1.ProfileService/CreateProfile \
  -H "Content-Type: application/json" \
  -d '{"full_name": "Kevin Chen"}'

http POST \
  http://localhost:9091/profile.v1beta1.ProfileService/CreateProfile \
    full_name="Kevin Chen"
```

### License Service

```shell
curl -v http://localhost:9092/license.v1beta1.LicenseService/CreateLicense \
  -H "Content-Type: application/json" \
  -d '{"start": "2023-11-16T12:00:00Z", "end": "2024-01-16T12:00:00Z"}'

http POST \
  http://localhost:9092/license.v1beta1.LicenseService/CreateLicense \
    start="2023-11-16T12:00:00Z" \
    end="2024-01-16T12:00:00Z"
```