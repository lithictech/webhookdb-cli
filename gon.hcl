source = ["./dist/webhookdb-macos_darwin_amd64/webhookdb"]
bundle_id = "com.lithictech.webhookdbcli"

apple_id {
  username = "root@lithic.tech"
  password = "@env:AC_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Lithic Technology LLC"
}