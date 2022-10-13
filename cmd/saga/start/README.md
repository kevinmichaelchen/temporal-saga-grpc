## GraphiQL
Navigate to [localhost:9094/graphql](http://localhost:9094/graphql) in your
browser.

Run the following mutation:
```graphql
mutation foo {
  createLicense(newNumber:1) {
    theNumber
  }
}
```