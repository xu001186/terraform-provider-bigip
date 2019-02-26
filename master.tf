
provider "bigip" {
 address = "10.192.74.57"
 username = "admin"
 password = "admin"
}
// Label is used to identify which Json payload to use
resource "bigip_app_as3"  "as3" {
   label = "Sample 1"
   ident = "someid"
   jsonfile = "${file("sam.json")}"
 }
