/*
AS3 uses a declarative model, meaning you provide a JSON declaration rather than a set of imperative commands. The declaration represents the configuration which AS3 is responsible for creating on a BIG-IP system. AS3 is well-defined according to the rules of JSON Schema, and declarations validate according to JSON Schema. AS3 accepts declaration updates via REST (push), reference (pull), or CLI (flat file editing).
*/
package bigip

import (
	"encoding/json"
	"fmt"
	"github.com/f5devcentral/go-bigip"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"time"
)

func resourceBigipappsvcs() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigipAppSvscsCreate,
		Read:   resourceBigipAppSvscsRead,
		Update: resourceBigipAppSvscsUpdate,
		Delete: resourceBigipAppSvscsDelete,
		Exists: resourceBigipAppSvscsExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"jsonfile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Address of the Iapp which needs to be Iappensed",
			},

			"label": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "what is the as3 label you are using ",
			},
			"ident": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "what is the as3 label you are using ",
			},
		},
	}
} //resource end

func resourceBigipAppSvscsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bigip.BigIP)

	name := d.Get("ident").(string)
	label := d.Get("label").(string)
	log.Printf("[INFO] Creating As3 app and value of d is  %s %s", name, d)
	if label == "Sample 1" {
		p := dataToAppsvc01(name, d)
		log.Printf(" value of p +++++++++++++++++++++++++++++++  %+v ", p)
		err := client.CreateAppsvc01(&p)
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Printf("[ERROR] Unable to Create Appsvc  (%s) (%v) ", name, err)
			return err
		}
	}

	if label == "Sample 2" {
		p := dataToAppsvc02(name, d)
		log.Printf(" value of p +++++++++++++++++++++++++++++++  %+v ", p)
		time.Sleep(2 * time.Second)
		err := client.CreateAppsvc02(&p)
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Printf("[ERROR] Unable to Create Appsvcs02  (%s) (%v) ", name, err)
			return err
		}
	}
	d.SetId(name)

	return resourceBigipAppSvscsRead(d, meta)

}

func resourceBigipAppSvscsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bigip.BigIP)
	name := d.Id()
	label := d.Get("label").(string)
	log.Println("[INFO] Reading Appsvcs " + name)
	if label == "Simple 1" {
		p, err := client.Appsvc01()
		if err != nil {
			log.Printf("[ERROR] Unable to Retrieve Appsvcs  (%s) (%v)", name, err)
			return err
		}
		if p == nil {
			log.Printf("[WARN] IApp (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		if label == "Simple 2" {
			p, err := client.Appsvc02()
			if err != nil {
				log.Printf("[ERROR] Unable to Retrieve Appsvcs  (%s) (%v)", name, err)
				return err
			}
			if p == nil {
				log.Printf("[WARN] IApp (%s) not found, removing from state", d.Id())
				d.SetId("")
				return nil
			}
		}
	}
	d.Set("name", name)
	return nil
}

func resourceBigipAppSvscsExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*bigip.BigIP)
	name := d.Id()
	label := d.Get("label").(string)

	log.Printf("[INFO] Checking if Appsvc (%s) exists", name)
	if label == "Sample 1" {
		p, err := client.Appsvc01()
		time.Sleep(4 * time.Second)
		if err != nil {
			log.Printf("[ERROR] Unable to Read Appsvc  (%s) (%v)", p, err)
			return false, err
		}
	}
	if label == "Sample 2" {
		time.Sleep(4 * time.Second)
		p, err := client.Appsvc02()
		if err != nil {
			log.Printf("[ERROR] Unable to Read Appsvc  (%s) (%v)", p, err)
			return false, err
		}
	}
	return true, nil
}

func resourceBigipAppSvscsUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bigip.BigIP)

	name := d.Get("ident").(string)
	label := d.Get("label").(string)
	log.Printf("[INFO] Modifying As3 app and value of d is  %s %s", name, d)
	if label == "Sample 1" {
		p := dataToAppsvc01(name, d)
		log.Printf(" value of p +++++++++++++++++++++++++++++++  %+v ", p)
		err := client.ModifyAppsvc01(&p)
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Printf("[ERROR] Unable to Create Appsvc  (%s) (%v) ", name, err)
			return err
		}
	}
	if label == "Sample 2" {
		p := dataToAppsvc02(name, d)
		log.Printf(" value of p +++++++++++++++++++++++++++++++  %+v ", p)
		err := client.ModifyAppsvc02(&p)
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Printf("[ERROR] Unable to Create Appsvc  (%s) (%v) ", name, err)
			return err
		}
	}
	return resourceBigipAppSvscsRead(d, meta)
}

func resourceBigipAppSvscsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bigip.BigIP)
	name := d.Id()
	label := d.Get("label").(string)
	if label == "Sample 1" {
		err := client.DeleteAppsvc01()
		time.Sleep(4 * time.Second)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete Appsvc  (%s) (%v)", name, err)
			return err
		}
	}
	if label == "Sample 2" {
		time.Sleep(4 * time.Second)
		err := client.DeleteAppsvc02()
		if err != nil {
			log.Printf("[ERROR] Unable to Delete Appsvc  (%s) (%v)", name, err)
			return err
		}
	}
	d.SetId("")
	return nil
}

func dataToAppsvc01(name string, d *schema.ResourceData) bigip.Appsvc01 {
	var p bigip.Appsvc01
	jj := d.Get("jsonfile").(string)
	jsonblob := []byte(d.Get("jsonfile").(string))
	err := json.Unmarshal(jsonblob, &p)
	p.Declaration.ID = d.Get("ident").(string)
	log.Printf("I am in datatoAppsvc ++++++++++++++++++ p %v %v", p, jj)
	if err != nil {
		fmt.Println("error", err)
	}
	return p
}
func dataToAppsvc02(name string, d *schema.ResourceData) bigip.Appsvc02 {
	var p bigip.Appsvc02
	jj := d.Get("jsonfile").(string)
	jsonblob := []byte(d.Get("jsonfile").(string))
	err := json.Unmarshal(jsonblob, &p)
	p.Declaration.Label = d.Get("ident").(string)
	log.Printf("I am in datatoAppsvc ++++++++++++++++++ p %v %v", p, jj)
	if err != nil {
		fmt.Println("error", err)
	}
	return p
}
