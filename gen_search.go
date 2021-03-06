package omise

/*
!!! DO NOT MODIFY !!!

autogenerated
 src: gen_search_job.tmpl
 job: &main.GenSearchJob{Names:[]string{"Charge", "Customer", "Dispute", "Recipient"}}
  on: Tue Nov 08 15:14:15 +0700 2016
  by: chakrit
*/

import (
	"bytes"
)

const ChargeScope SearchScope = "charge"

// ChargeSearchResult represents search result structure returned by Omise's Search API
// that contains Charge struct as result elements.
type ChargeSearchResult struct {
	SearchResult
	Data []*Charge `json:"data"`
}

func (list *ChargeSearchResult) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

const CustomerScope SearchScope = "customer"

// CustomerSearchResult represents search result structure returned by Omise's Search API
// that contains Customer struct as result elements.
type CustomerSearchResult struct {
	SearchResult
	Data []*Customer `json:"data"`
}

func (list *CustomerSearchResult) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

const DisputeScope SearchScope = "dispute"

// DisputeSearchResult represents search result structure returned by Omise's Search API
// that contains Dispute struct as result elements.
type DisputeSearchResult struct {
	SearchResult
	Data []*Dispute `json:"data"`
}

func (list *DisputeSearchResult) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

const RecipientScope SearchScope = "recipient"

// RecipientSearchResult represents search result structure returned by Omise's Search API
// that contains Recipient struct as result elements.
type RecipientSearchResult struct {
	SearchResult
	Data []*Recipient `json:"data"`
}

func (list *RecipientSearchResult) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}
