schema "public" {
  comment = "standard public schema"
}

table "license" {
  schema = schema.public
  comment = "A license indicating some kind of purchase, valid for a specific window of time"
  column "id" {
    comment = "UUID primary key"
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "start_time" {
    null = false
    type = timestamptz
  }
  column "end_time" {
    null = false
    type = timestamptz
  }
  column "user_id" {
    comment = "The license's user"
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
}
