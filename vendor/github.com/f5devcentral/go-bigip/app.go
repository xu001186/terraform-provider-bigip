package bigip

import (
	//"time"
	"log"
)

type Appsvcs struct {
	Appsvcs []Appsvc01 `json:"items"`
}
type Appsvc01 struct {
	Class       string `json:"class"`
	Action      string `json:"action"`
	Persist     bool   `json:"persist"`
	Declaration struct {
		Class         string `json:"class"`
		SchemaVersion string `json:"schemaVersion"`
		ID            string `json:"id"`
		Label         string `json:"label"`
		Remark        string `json:"remark"`
		Sample01      struct {
			Class              string `json:"class"`
			DefaultRouteDomain int    `json:"defaultRouteDomain"`
			Application1       struct {
				Class       string `json:"class"`
				Template    string `json:"template"`
				ServiceMain struct {
					Class            string   `json:"class"`
					VirtualAddresses []string `json:"virtualAddresses"`
					Pool             string   `json:"pool"`
				} `json:"serviceMain"`
				WebPool struct {
					Class    string   `json:"class"`
					Monitors []string `json:"monitors"`
					Members  []struct {
						ServicePort     int      `json:"servicePort"`
						ServerAddresses []string `json:"serverAddresses"`
					} `json:"members"`
				} `json:"web_pool"`
			} `json:"Application_1"`
		} `json:"Sample_01,omitempty"`
	} `json:"declaration,omitempty"`
}

type Appsvc02 struct {
	Class       string `json:"class"`
	Action      string `json:"action"`
	Persist     bool   `json:"persist"`
	Declaration struct {
		Class         string `json:"class"`
		SchemaVersion string `json:"schemaVersion"`
		ID            string `json:"id"`
		Label         string `json:"label"`
		Remark        string `json:"remark"`
		Sample02      struct {
			Class string `json:"class"`
			A1    struct {
				Class       string `json:"class"`
				Template    string `json:"template"`
				ServiceMain struct {
					Class            string   `json:"class"`
					VirtualAddresses []string `json:"virtualAddresses"`
					Pool             string   `json:"pool"`
					ServerTLS        string   `json:"serverTLS"`
				} `json:"serviceMain"`
				WebPool struct {
					Class             string   `json:"class"`
					LoadBalancingMode string   `json:"loadBalancingMode"`
					Monitors          []string `json:"monitors"`
					Members           []struct {
						ServicePort     int      `json:"servicePort"`
						ServerAddresses []string `json:"serverAddresses"`
					} `json:"members"`
				} `json:"web_pool"`
				Webtls struct {
					Class        string `json:"class"`
					Certificates []struct {
						Certificate string `json:"certificate"`
					} `json:"certificates"`
				} `json:"webtls"`
				Webcert struct {
					Class       string `json:"class"`
					Remark      string `json:"remark"`
					Certificate string `json:"certificate"`
					PrivateKey  string `json:"privateKey"`
					Passphrase  struct {
						Ciphertext string `json:"ciphertext"`
						Protected  string `json:"protected"`
					} `json:"passphrase"`
				} `json:"webcert"`
			} `json:"A1"`
		} `json:"Sample_02"`
	} `json:"declaration"`
}

const (
	uriSha     = "shared"
	uriAppsvcs = "appsvcs"
	uriDecl    = "declare"
	uriSam01   = "Sample_01"
	uriSam02   = "Sample_02"
)

// Appsvcss returns a list of appsvcs
func (b *BigIP) Appsvc01() (*Appsvc01, error) {
	var appsvc01 Appsvc01
	err, _ := b.getForEntity(uriSam01, uriSha, uriAppsvcs, uriDecl)
	log.Printf("i am here in sdk %+v  ", appsvc01)
	if err != nil {
		return nil, err
	}

	return &appsvc01, nil
}
func (b *BigIP) Appsvc02() (*Appsvc02, error) {
	var appsvc02 Appsvc02
	err, _ := b.getForEntity(uriSam02, uriSha, uriAppsvcs, uriDecl)
	log.Printf("i am here in sdk %+v  ", appsvc02)
	if err != nil {
		return nil, err
	}

	return &appsvc02, nil
}

// CreateAppsvcs creates a new iAppsvcs on the system.
func (b *BigIP) CreateAppsvc01(p *Appsvc01) error {
	log.Printf("++++++ Here is what terraform is sending to bigip ................ : %+v ", p)
	err := b.post(p, uriMgmt, uriSha, uriAppsvcs, uriDecl)
	if err != nil {
		log.Println(" API call not successfull  %v ", err)
	}
	return nil
}
func (b *BigIP) CreateAppsvc02(p *Appsvc02) error {
	log.Printf("++++++ Here is what terraform is sending to bigip ................ : %+v ", p)
	err := b.post(p, uriMgmt, uriSha, uriAppsvcs, uriDecl)
	if err != nil {
		log.Println(" API call not successfull  %v ", err)
	}
	return nil
}
func (b *BigIP) DeleteAppsvc01() error {
	return b.delete(uriMgmt, uriSha, uriAppsvcs, uriDecl, uriSam01)
}
func (b *BigIP) DeleteAppsvc02() error {
	return b.delete(uriMgmt, uriSha, uriAppsvcs, uriDecl, uriSam02)
}

func (b *BigIP) ModifyAppsvc01(p *Appsvc01) error {
	log.Printf("++++++ Here is what terraform is sending to bigip ................ : %+v ", p)
	err := b.patch(p, uriMgmt, uriSha, uriAppsvcs, uriDecl)
	log.Printf("value of p in modify +++++++++++++++", p)
	if err != nil {
		log.Println(" API call not successfull  %v ", err)
	}
	return nil
}
func (b *BigIP) ModifyAppsvc02(p *Appsvc02) error {
	log.Printf("++++++ Here is what terraform is sending to bigip ................ : %+v ", p)
	err := b.patch(p, uriMgmt, uriSha, uriAppsvcs, uriDecl)
	if err != nil {
		log.Println(" API call not successfull  %v ", err)
	}
	return nil
}
