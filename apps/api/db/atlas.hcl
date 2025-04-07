data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "./cmd/follytics",
    "migrate",
    "generate"
  ]
}

env "dev" {
  src     = data.external_schema.gorm.url
  dev     = "docker://postgres/17/dev?search_path=public"

  migration {
    dir = "file://db/migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }

  diff {
    skip {
      drop_schema  = true
      drop_table   = true
      drop_func    = true
      drop_trigger = true
    }
  }
}
