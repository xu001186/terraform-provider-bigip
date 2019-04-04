
provider "bigip" {
 address = "54.177.207.51"
 username = "admin"
 password = "1LStm3545"
}
// Label is used to identify which Json payload to use.
resource "bigip_app_as3"  "as3-example1" {
   label = "Sample 1"
   ident = "sanjoseid"
   jsonfile = "${file("example1.json")}"
 }

resource "bigip_app_as3"  "as3-example2" {
   label = "Sample 2"
   ident = "newyorkid"
   jsonfile = "${file("example2.json")}"
 }



