// Package pools provides deduplication pools for strings
// and lists of ints and strings.
package pools

import "log"

// Default pools: These pools are defined globally
// and are defined per intended usage

// Neighbors stores neighbor IDs
var Neighbors *StringPool

// Networks4 stores network ip v4 addresses
var Networks4 *StringPool

// Networks6 stores network ip v6 addresses
var Networks6 *StringPool

// Interfaces stores interfaces like: eth0, bond0 etc...
var Interfaces *StringPool

// Gateways4 store ip v4 gateway addresses
var Gateways4 *StringPool

// Gateways6 store ip v6 gateway addresses
var Gateways6 *StringPool

// Origins is a store for 'IGP'
var Origins *StringPool

// ASPaths stores lists of ASNs
var ASPaths *IntListPool

// Types stores a list of types (['BGP', 'univ'])
var Types *StringListPool

// Communities store a list of BGP communities
var Communities *CommunitiesPool

// ExtCommunities stores a list of extended communities
var ExtCommunities *CommunitiesPool

// LargeCommunities store a list of large BGP communities
var LargeCommunities *CommunitiesPool

// Initialize global pools
func init() {
	log.Println("initializing memory pools")

	Neighbors = NewStringPool()
	Networks4 = NewStringPool()
	Networks6 = NewStringPool()
	Interfaces = NewStringPool()
	Gateways4 = NewStringPool()
	Gateways6 = NewStringPool()
	Origins = NewStringPool()
	ASPaths = NewIntListPool()
	Types = NewStringListPool()
	Communities = NewCommunitiesPool()
	ExtCommunities = NewExtCommunitiesPool()
	LargeCommunities = NewCommunitiesPool()
}
